#!/bin/sh
set -e
rm -rf manpages
mkdir manpages
go run cmd/mollie/main.go docs man | gzip -c -9 >manpages/mollie.1.gz