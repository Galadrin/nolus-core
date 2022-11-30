#!/bin/bash
# Configure a locally installed sentry node for the needs of the test or main networks
#
# arg: home directory of the sentry node, mandatory
# arg: external IP address, mandatory
# arg: P2P port, mandatory
# arg: RPC port, mandatory
# arg: Monitoring port, mandatory
# arg: API port, mandatory
# arg: validator node URL, mandatory
# arg: validator node ID, mandatory
# arg: a comma separated list of sibling sentry node URLs, mandatory
# arg: a comma separated list of sibling sentry node IDs, mandatory
# arg: a comma separated list of other sentry node URLs, mandatory, the list must begin with a comma

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "$SCRIPT_DIR"/lib/lib.sh

declare -r home_dir="$1"
declare -r external_address="$2"
declare -r p2p_port="$3"
declare -r rpc_port="$4"
declare -r monitoring_port="$5"
declare -r api_port="$6"
declare -r grpc_port="$7"
declare -r validator_node_url="$8"
declare -r validator_node_id="$9"
declare -r sibling_sentry_node_urls_str="${10}"
declare -r sibling_sentry_node_ids_str="${11}"
declare -r other_sentry_node_urls_str="${12}"

declare -r config_dir="$home_dir"/config

# although the API endpoint is deprecated it is still required by Keplr
# TBD reevaluate the necessity to remain open
# the grpc endpoint is required by the market data feeder
update_app "$config_dir" '."api"."enable"' "true"
update_app "$config_dir" '."api"."address"' '"tcp://'"$external_address:$api_port"'"'
update_app "$config_dir" '."api"."enabled-unsafe-cors"' "true"
update_app "$config_dir" '."grpc"."enable"' "true" >/dev/null
update_app "$config_dir" '."grpc"."address"' '"0.0.0.0:'"$grpc_port"'"' >/dev/null
update_app "$config_dir" '."grpc-web"."enable"' "false"
update_app "$config_dir" '."telemetry"."enabled"' "true"
update_app "$config_dir" '."telemetry"."prometheus-retention-time"' "1"

update_config "$config_dir" '."rpc"."laddr"' '"tcp://'"$external_address:$rpc_port"'"'
update_config "$config_dir" '."rpc"."cors_allowed_origins"' '["*"]'
update_config "$config_dir" '."p2p"."laddr"' '"tcp://'"$external_address:$p2p_port"'"'
update_config "$config_dir" '."p2p"."seed_mode"' "false"
update_config "$config_dir" '."p2p"."pex"' "true"
update_config "$config_dir" '."p2p"."persistent_peers"' '"'"$validator_node_url","$sibling_sentry_node_urls_str""$other_sentry_node_urls_str"'"'
update_config "$config_dir" '."p2p"."unconditional_peer_ids"' '"'"$validator_node_id","$sibling_sentry_node_ids_str"'"'
update_config "$config_dir" '."p2p"."private_peer_ids"' '"'"$validator_node_id"'"'
update_config "$config_dir" '."p2p"."addr_book_strict"' "false"
update_config "$config_dir" '."p2p"."allow_duplicate_ip"' "false"
update_config "$config_dir" '."instrumentation"."prometheus"' "true"
update_config "$config_dir" '."instrumentation"."prometheus_listen_addr"' '"'":$monitoring_port"'"'
