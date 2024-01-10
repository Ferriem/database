#!/bin/bash

# Function to execute cluster reset and flushall
reset_and_flush() {
    redis-cli -p $1 cluster reset
    redis-cli -p $1 flushall
}

# Remove configuration files
rm -f nodes_{7000..7005}.conf

# Start new Redis instances
for port in {7000..7005}; do
    redis-server redis_${port}.conf &
done

# Wait for Redis instances to start (you may need to adjust the sleep duration)
sleep 5

# Execute cluster reset and flushall for each Redis instance
for port in {7000..7005}; do
    reset_and_flush $port
done

echo "Redis is running..."

# Optionally, you can uncomment the following line to create the cluster
# redis-cli --cluster create 127.0.0.1:7000 127.0.0.1:7001 127.0.0.1:7002 127.0.0.1:7003 127.0.0.1:7004 127.0.0.1:7005 --cluster-replicas 1
