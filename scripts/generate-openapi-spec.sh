#!/bin/bash
#
# Copyright SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#
set -e

CODEBASE="$(dirname "$PWD")"
SPEC_LOC="${SPEC_LOC}"
SPEC_META="${SPEC_META:-cmd/aries-agentd/main.go}"
OUTPUT="$SPEC_LOC/openAPI.yml"

# generate and validate commands
GENERATE_COMMAND="generate spec $SPEC_META -o $OUTPUT"
VALIDATE_COMMAND="validate $OUTPUT"

echo "Generating Open API spec"
docker run --rm -it -e GOPATH=$HOME/go:/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger $GENERATE_COMMAND

echo "Validating generated spec"
docker run --rm -it -e GOPATH=$HOME/go:/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger $VALIDATE_COMMAND