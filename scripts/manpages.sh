#!/bin/sh
set -e
rm -rf manpages
mkdir manpages
go run ./cmd/sn2ssg-go/ man | gzip -c -9 >manpages/sn2ssg-go.1.gz
