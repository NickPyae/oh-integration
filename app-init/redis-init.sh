#!/bin/bash

echo "Connecting to Redis..."
connected=false
for try in {1..15}; do
    if [[ `redis-cli -h ${REDIS_IP} -p ${REDIS_PORT} ping` == "PONG" ]]; then
        echo "Connected to Redis."
        connected=true
        break
    else
        sleep 1
    fi
done

if [[ "$connected" = false ]]; then
    echo "Unable to connect to Redis after 15 seconds. Exiting."
    exit 1
fi

echo "Registering RedisGears function..."
redis-cli -h ${REDIS_IP} -p ${REDIS_PORT} RG.PYEXECUTE "`cat transform.py`"

echo "Configure save..."
redis-cli -h ${REDIS_IP} -p ${REDIS_PORT} CONFIG SET SAVE "900 1 300 10 60 10000"