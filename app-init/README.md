# Creating default stream and rules for kuiper rule engine

# Prerequisite
1. Docker installed
2. kuiper rule engine container should be reachable from this app-init container

# How it works
This custom app-init container will create one stream and two new rules for kuiper rule engine container once it is ready.

# How to build docker image

```sh
docker build -t <docker_hub_id>/app-init:<tag> .
```

# Run the container from image

1. Export KUIPER_IP, KUIPER_PORT, INFLUXDB_IP, INFLUXDB_PORT, INFLUXDB_TOKEN, INFLUXDB_CLOUD_TOKEN and BUCKET_NAME as environment variables to host environment. 
2. Replace ` KUIPER_IP=x.x.x.x` with IP address of VM where kuiper rule engine container is running. Default kuiper rule engine container port is `48075`. 
3. Replace `INFLUXDB_IP=x.x.x.x` with IP address of VM where InfluxDB is running and `8086` is default port of InfluxDB.
4. Replace `YOUR_BUCKET_NAME` with correct bucket name that you use for InfluxDB.
5. Replace `YOUR_INFLUXDB_TOKEN` with correct authorization header token that you use for InfluxDB.
6. Replace `YOUR_INFLUXDB_CLOUD_TOKEN` with correct authorization header token that you use for InfluxDB Cloud.


```sh
export KUIPER_IP=x.x.x.x 
export KUIPER_PORT=48075 
export INFLUXDB_IP=x.x.x.x 
export INFLUXDB_PORT=8086
export INFLUXDB_TOKEN=YOUR_INFLUXDB_TOKEN
export INFLUXDB_CLOUD_TOKEN=YOUR_INFLUXDB_CLOUD_TOKEN
export BUCKET_NAME=YOUR_BUCKET_NAME 

docker run -e KUIPER_IP -e KUIPER_PORT -e INFLUXDB_IP -e INFLUXDB_PORT -e INFLUXDB_TOKEN -e INFLUXDB_CLOUD_TOKEN -e BUCKET_NAME <docker_hub_id>/app-init:<tag>
```


