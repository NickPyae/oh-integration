# Device Service: device-random-temperature

The device-random-temperature device service creates a pre-defined device named 'Random-Temperature-Generator01'. This device generates random temperature reading between 50 - 200 Fahrenheit (every 1 second).

To check:
- curl -v <server-node-ip>:30800/api/v1/event/device/Random-Temperature-Generator01/10

## Pre-requisites

Ensure EdgeX Core Services (Core Metadata, Core Data, Core Command) are running.

Please check:
- curl -v <server-node-ip>:30800/api/v1/ping
- curl -v <server-node-ip>:30801/api/v1/ping
- curl -v <server-node-ip>:30802/api/v1/ping

### Creating Docker image

git clone https://eos2git.cec.lab.emc.com/ISG-Edge/HelloSally.git
(ensure path is: ~/go/src/HelloSally/device-random-temperature/)

cd HelloSally/device-random-temperature/image

go test -v ./driver/

docker build -t device-random-temperature-go:latest .

#### Creating Kubernetes resources

cd HelloSally/device-random-temperature/k8s

kubectl apply -f deployment.yaml

kubectl apply -f service.yaml
