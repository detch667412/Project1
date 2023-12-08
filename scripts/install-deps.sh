#!/bin/bash

set -euo pipefail

printf "🕸️\t Install Backend dependency...  \n"
go mod download
