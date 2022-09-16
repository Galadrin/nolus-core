#!/bin/bash

cleanup_init_network_sh() {
  cleanup_genesis_sh
}

SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
source "$SCRIPT_DIR"/genesis.sh

# arg1: an existing local dir where validator accounts are created, mandatory
init_network() {
  local val_accounts_dir="$1"
  local validators="$2"
  local chain_id="$3"
  local native_currency="$4"
  local val_tokens="$5"
  local val_stake="$6"
  local genesis_accounts_spec="$7"
  local -r wasm_script_path="$8"
  local -r wasm_code_path="$9"
  local -r wasm_admin_addr="${10}"
  local -r treasury_init_tokens_u128="${11}"
  local -r lpp_native="${12}"
  local -r contracts_info_file="${13}"

  node_id_and_val_pubkeys="$(setup_validators "$validators")"
  local final_genesis_file;
  final_genesis_file=$(generate_genesis "$chain_id" "$native_currency" \
                                          "$val_tokens" "$val_stake" \
                                          "$val_accounts_dir" "$genesis_accounts_spec" \
                                          "$wasm_script_path" "$wasm_code_path" \
                                          "$wasm_admin_addr" "$treasury_init_tokens_u128" \
                                          "$node_id_and_val_pubkeys" \
                                          "$lpp_native" "$contracts_info_file")
  propagate_genesis "$final_genesis_file" "$validators"
}
