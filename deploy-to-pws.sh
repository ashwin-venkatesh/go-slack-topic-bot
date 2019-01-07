#!/bin/bash

# Target PWS
cf api https://api.run.pivotal.io

cf login -o garden-windows -s topic-bots

# Build the binaries for the bots
pushd windows-bots
  go build pcf-windows.go
  go build garden-windows.go
popd

cf push

cf stop garden-windows-topic-bot
cf stop pcf-windows-topic-bot
