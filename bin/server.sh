#!/bin/bash

#nodes_file="../nodes.json"
nodes_file="../nodes-single.json"

service_host="localhost"
service_port=9600


./server \
  -nodes_file="$nodes_file" \
  -service_host="$service_host" \
  -service_port="$service_port"
