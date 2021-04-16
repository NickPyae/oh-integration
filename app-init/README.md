# Creating default stream and rules for kuiper rule engine

# Prerequisite
1. Docker installed
2. Kuiper rule engine and Redis containers should be reachable from this app-init container

# How it works
This custom app-init container will run 2 scripts:
- `create-rule.sh`: Create one stream and two new rules for Kuiper rule engine container once it is ready.
- `redis-init.sh`: Register RedisGears function to convert Redis data from EdgeX JSON format to RedisTimeSeries format. This function is triggered as new readings are saved.
    - The created timeseries will have the key `ts:<DEVICE-NAME>:<READING-VALUE>` and can be viewed by running `redis-cli TS.RANGE <TS-KEY> - +` inside the Redis container.

# How to build docker image

```sh
docker build -t <docker_hub_id>/app-init:<tag> .
```

# Run the container from image

1. Export the following environment variables to host environment.

```sh
export KUIPER_IP=x.x.x.x 
export KUIPER_PORT=48075 
export INFLUXDB_IP=x.x.x.x 
export INFLUXDB_PORT=8086
export INFLUXDB_TOKEN=YOUR_INFLUXDB_TOKEN
export INFLUXDB_CLOUD_TOKEN=YOUR_INFLUXDB_CLOUD_TOKEN
export BUCKET_NAME=YOUR_BUCKET_NAME
export REDIS_IP=x.x.x.x
export REDIS_PORT=6379
```

2. Replace `KUIPER_IP=x.x.x.x` with IP address of VM where kuiper rule engine container is running. Default kuiper rule engine container port is `48075`. 
3. Replace `INFLUXDB_IP=x.x.x.x` with IP address of VM where InfluxDB is running and `8086` is default port of InfluxDB.
4. Replace `YOUR_BUCKET_NAME` with correct bucket name that you use for InfluxDB.
5. Replace `YOUR_INFLUXDB_TOKEN` with correct authorization header token that you use for InfluxDB.
6. Replace `YOUR_INFLUXDB_CLOUD_TOKEN` with correct authorization header token that you use for InfluxDB Cloud.
7. Replace `REDIS_IP=x.x.x.x` with IP address of VM where Redis container is running. Default Redis port is `6379`.

```sh
docker run -e KUIPER_IP -e KUIPER_PORT -e INFLUXDB_IP -e INFLUXDB_PORT -e INFLUXDB_TOKEN -e INFLUXDB_CLOUD_TOKEN -e BUCKET_NAME -e REDIS_IP -e REDIS_PORT <docker_hub_id>/app-init:<tag>
```