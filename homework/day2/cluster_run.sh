#! /bin/bash

rm -f nodes_7000.conf nodes_7001.conf nodes_7002.conf nodes_7003.conf nodes_7004.conf nodes_7005.conf

redis-server redis_7000.conf &
redis-server redis_7001.conf &
redis-server redis_7002.conf &

echo "Redis is running..."

#redis-cli --cluster create 127.0.0.1:7000 127.0.0.1:7001 127.0.0.1:7002 --cluster-replicas 0
