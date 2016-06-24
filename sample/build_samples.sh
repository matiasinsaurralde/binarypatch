#!/bin/sh
echo "Building sample programs"
go build -o sample_darwin64 sample.go
GOOS=linux GOARCH=386 go build -o sample_linux32 sample.go
GOOS=linux GOARCH=amd64 go build -o sample_linux64 sample.go
GOOS=windows GOARCH=386 go build -o sample_win32.exe sample.go
GOOS=windows GOARCH=amd64 go build -o sample_win64.exe sample.go
echo "Done."
