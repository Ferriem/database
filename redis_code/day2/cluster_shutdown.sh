#!/bin/bash

for port in {7000..7005}; do
    redis-cli -p $port shutdown
done

sleep 3

echo "Redis is shut down..."