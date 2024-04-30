#!/bin/sh
set -e
rm -rf completions
mkdir completions
for sh in bash zsh fish; do
	go run ./cmd/sn2ssg-go/ completion "$sh" >"completions/sn2ssg-go.$sh"
done
