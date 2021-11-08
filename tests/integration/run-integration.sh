#!/bin/bash
set -euxo pipefail

cleanup() {
  if [[ -n "${COSMZONED_PID:-}" ]]; then
    echo "Stopping cosmzone..."
    kill -7 "$COSMZONED_PID"
    exit
  fi
}

trap cleanup INT TERM EXIT

prepare_env() {
  cosmzoned keys add test-user-1 --keyring-backend "test" # force no password
  cosmzoned keys add test-user-2 --keyring-backend "test" # force no password
  VALIDATOR_ADDR=$(cosmzoned keys show local-validator -a)
  VALIDATOR_PRIV_KEY=$(echo 'y' | cosmzoned keys  export local-validator --unsafe --unarmored-hex 2>&1)
  PERIODIC_VEST_ADDR=$(cosmzoned keys show periodic-vesting-account -a)
  PERIODIC_PRIV_KEY=$(echo 'y' | cosmzoned keys  export periodic-vesting-account --unsafe --unarmored-hex 2>&1)
  USR_1_ADDR=$(cosmzoned keys show test-user-1 -a)
  USR_2_ADDR=$(cosmzoned keys show test-user-2 -a)
  USR_3_ADDR=$(cosmzoned keys show test-user-3 -a)
  USR_1_PRIV_KEY=$(echo 'y' | cosmzoned keys  export test-user-1 --unsafe --unarmored-hex 2>&1)
  USR_2_PRIV_KEY=$(echo 'y' | cosmzoned keys  export test-user-2 --unsafe --unarmored-hex 2>&1)
  USR_3_PRIV_KEY=$(echo 'y' | cosmzoned keys  export test-user-3 --unsafe --unarmored-hex 2>&1)
  DOT_ENV=$(cat <<-EOF
NODE_URL=http://localhost:26657
VALIDATOR_ADDR=${VALIDATOR_ADDR}
VALIDATOR_PRIV_KEY=${VALIDATOR_PRIV_KEY}
PERIODIC_VEST_ADDR=${PERIODIC_VEST_ADDR}
PERIODIC_PRIV_KEY=${PERIODIC_PRIV_KEY}
USR_1_ADDR=${USR_1_ADDR}
USR_1_PRIV_KEY=${USR_1_PRIV_KEY}
USR_2_ADDR=${USR_2_ADDR}
USR_2_PRIV_KEY=${USR_2_PRIV_KEY}
USR_3_ADDR=${USR_3_ADDR}
USR_3_PRIV_KEY=${USR_3_PRIV_KEY}
EOF
  )
  echo "$DOT_ENV" > .env
}


cd ../../
./init.sh prepare integration >/tmp/cosmzone-prepare.log 2>&1
cd tests/integration

cosmzoned start >/tmp/cosmzone-run.log 2>&1 &
COSMZONED_PID=$!
sleep 5

prepare_env
yarn test "$@"

