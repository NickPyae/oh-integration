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

1. Export KUIPER_IP and KUIPER_PORT as environment variables to host environment. 
2. Replace X.X.X.X with IP address of VM where kuiper rule engine container is running. Default kuiper rule engine container port is 48075.


```sh
export KUIPER_IP=X.X.X.X && export KUIPER_PORT=48075 
docker run -e KUIPER_IP -e KUIPER_PORT <docker_hub_id>/app-init:<tag>
```


