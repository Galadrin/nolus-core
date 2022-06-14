#!/bin/bash
# 
# Check if jq and tomlq are accessible
#

if [[ ! -x $(which jq) ]]; then
    echo >&2 "jq processor is not found in PATH=${PATH}"
    exit 1
fi

if [[ ! -x $(which tomlq) ]]; then
    echo >&2 "tomlq processor is not found in PATH=${PATH}"
    exit 1
fi
