#!/bin/bash

device="lo0"

# bpf_filter="tcp and port 3306"
bpf_filter="tcp"
txt_filter="SELECT"

max_query_len=9192

service_host="localhost"
service_port=9501


./agent \
  -device="$device" \
  -bpf_filter="$bpf_filter" \
  -txt_filter="$txt_filter" \
  -max_query_len="$max_query_len" \
  -service_host="$service_host" \
  -service_port="$service_port"
