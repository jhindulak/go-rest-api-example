#!/usr/bin/env bash
set -euo pipefail

varfile="$1"
if [ ! -f $varfile ]; then
  echo "Unable to find tfvars file at $varfile, exiting..."
  exit 1
fi

if [ -z "$(which terraform 2>/dev/null)" ]; then
  echo "Unable to find 'terraform' in \$PATH, exiting..."
  exit 1
fi

source $varfile

echo "Initializing terraform for the $env environment."

terraform init -backend=true \
               -backend-config="key=go-contacts/state/${env}/terraform.tfstate"