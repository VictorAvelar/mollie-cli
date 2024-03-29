#!/bin/sh
set -e

rm -rf completions
mkdir completions
for sh in bash zsh fish; do
	go run cmd/mollie/main.go completion "$sh" >"completions/goreleaser.$sh"
done