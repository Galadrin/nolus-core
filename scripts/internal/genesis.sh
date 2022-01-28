#!/bin/bash
set -euxo pipefail

SCRIPT_DIR=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)
source "$SCRIPT_DIR"/cmd.sh
"$SCRIPT_DIR"/check-jq.sh

# start "instance" variables
genesis_home_dir=$(mktemp -d)
# end "instance" variables

cleanup_genesis_sh() {
  if [[ -n "${genesis_home_dir:-}" ]]; then
    rm -rf "$genesis_home_dir"
  fi
}

generate_proto_genesis() {
  local chain_id="$1"
  local accounts_file="$2"
  local currency="$3"
  local proto_genesis_file="$4"
  local suspend_admin="$5"

  run_cmd "$genesis_home_dir" init genesis_manager --chain-id "$chain_id"
  run_cmd "$genesis_home_dir" config keyring-backend test
  run_cmd "$genesis_home_dir" config chain-id "$chain_id"

  local genesis_file="$genesis_home_dir/config/genesis.json"
  __set_token_denominations "$genesis_file" "$currency"
  __set_suspend_admin "$genesis_file" "$suspend_admin"

  for i in $(jq '. | keys | .[]' "$accounts_file"); do
    row=$(jq ".[$i]" "$accounts_file")
    add_genesis_account "$row" "$currency" "$genesis_home_dir"
  done

  cp "$genesis_file" "$proto_genesis_file"
}

#
# Takes a json object and creates a genesis account
#
# JSON specification object:
# "address" - mandatory string
# "amount" - mandatory string in the form '<number><currency>[,<number><currency>]*'
# "vesting" - optional object
# "vesting.start-time" - optional string representing a datetime in ISO 8601 format with max precision in seconds,
#                         for example "2022-01-28T13:15:59+02:00"
# "vesting.end-time" - mandatory string representing a datetime in ISO 8601 format with max precision in seconds,
#                         for example "2022-01-30T15:15:59-06:00"
# "vesting.amount" - mandatory number in native currency, e.g. 100 means "100 unolus"
add_genesis_account() {
  local specification="$1"
  local currency="$2"
  # TBD remove the following argument once the periodic vesting testing is deleted
  local home_dir="$3"

  local address
  address=$(jq -r '.address' <<< "$specification")
  local amount
  amount=$(jq -r '.amount' <<< "$specification")
  if [[ "$(jq -r '.vesting' <<< "$specification")" != 'null' ]]; then
    local vesting_start_time=""
    if [[ "$(jq -r '.vesting."start-time"' <<< "$specification")" != 'null' ]]; then
      vesting_start_time="--vesting-start-time $(__read_unix_time "$specification" start-time)"
    fi

    local vesting_end_time
    vesting_end_time=$(echo "$specification" | jq -r '.vesting."end-time"' | __as_unix_time )
    local vesting_amount
    vesting_amount=$(jq -r '.vesting.amount' <<< "$row")
    run_cmd "$home_dir" add-genesis-account "$address" "$amount" \
                --vesting-amount "$vesting_amount$currency" \
                --vesting-end-time "$vesting_end_time" $vesting_start_time
  else
    run_cmd "$home_dir" add-genesis-account "$address" "$amount"
  fi
}

integrate_genesis_txs() {
  local genesis_in_file="$1"
  local txs="$2"
  local genesis_out_file="$3"

  local genesis_basedir="$genesis_home_dir"/config
  local genesis_file="$genesis_basedir"/genesis.json
  cp "$genesis_in_file" "$genesis_file"

  local txs_dir="$genesis_home_dir"/txs
  {
    mkdir "$txs_dir"
    local index=0
    for tx in $txs; do
        echo "$tx" > "$txs_dir"/tx"$index".json
        index=$((index+1))
    done
  }

  run_cmd "$genesis_home_dir" collect-gentxs --gentx-dir "$txs_dir"
  cp "$genesis_file" "$genesis_out_file"
}

#####################
# private functions #
#####################
__set_token_denominations() {
  local genesis_file="$1"
  local currency="$2"

  local genesis_tmp_file="$genesis_file".tmp

  < "$genesis_file" \
    jq '.app_state["staking"]["params"]["bond_denom"]="'"$currency"'"' \
    | jq '.app_state["crisis"]["constant_fee"]["denom"]="'"$currency"'"' \
    | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="'"$currency"'"' \
    | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="'"$currency"'"' \
    | jq '.app_state["mint"]["params"]["mint_denom"]="'"$currency"'"' > "$genesis_tmp_file"
  mv "$genesis_tmp_file" "$genesis_file"
}

__set_suspend_admin() {
  local genesis_file="$1"
  local suspend_admin="$2"
  local genesis_tmp_file="$genesis_file".tmp

  < "$genesis_file" \
    jq '.app_state["suspend"]["state"]["admin_address"]="'"$suspend_admin"'"' > "$genesis_tmp_file"
  mv "$genesis_tmp_file" "$genesis_file"
}

__read_unix_time() {
  local spec="$1"
  local time_prop="$2"

  echo "$spec" | jq -r ".vesting.\"$time_prop\"" | __as_unix_time
}

__as_unix_time() {
  local datetime;
  read -r datetime
  date --date "$datetime" +%s
}