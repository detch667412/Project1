#!/bin/bash

set -euo pipefail

printf "\360\237\215\272\t Running API unit tests...  \n"
go test -v ./...
