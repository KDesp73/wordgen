#!/usr/bin/env bash


go build ./cmd/wordgen

sudo cp ./wordgen /usr/bin
