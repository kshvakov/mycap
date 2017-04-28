#!/bin/bash

web_path="`dirname $0`/../web/"

server_host="localhost"
server_port=9600

service_host="localhost"
service_port=9700


./web \
  -web_path="$web_path" \
  -server_host="$server_host" \
  -server_port="$server_port" \
  -service_host="$service_host" \
  -service_port="$service_port"
