#!/bin/bash

path="`pwd`"
binpath="$path/bin/"

echo "MyCAP builder"

echo "Build agent"
cd $path/agent/bin && go build -o $binpath/agent


echo "Build server"
cd $path/server/bin && go build -o $binpath/server

echo "Build web"
cd $path/web/bin && go build -o $binpath/web
