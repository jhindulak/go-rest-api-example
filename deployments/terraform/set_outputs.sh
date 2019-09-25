#!/usr/bin/env bash
set -euo pipefail

# Invoke this script using `source ./set_outputs.sh` so the parent process gets the environment variables set

function export_variable() {
    export $1=$(terraform output $1)
    echo "$1 set."
}

export_variable "database_address"
export_variable "database_name"
export_variable "database_password"
export_variable "database_port"
export_variable "route53_record"