# Device Service: device-random-temperature

This app:
1) Creates - Addressable, Value Descriptors, Device Profile, Device Service, Device
2) Provides API endpoint to:
- For device service to register device
- For core command to request action of device with command (via REST PUT) 
3) References https://docs.edgexfoundry.org/1.3/examples/LinuxTutorial/EdgeX-Foundry-tutorial-ver1.1.pdf (page 34 - 41)


## Pre-requisites

Ensure EdgeX Core Services (Core Metadata, Core Data, Core Command) are running.

Please check:
- curl -v <server-node-ip>:30800/api/v1/ping
- curl -v <server-node-ip>:30801/api/v1/ping
- curl -v <server-node-ip>:30802/api/v1/ping

## To run app

git clone https://eos2git.cec.lab.emc.com/ISG-Edge/HelloSally.git

cd HelloSally/device-random-temperature/app

go run .

open browser and go to http://localhost:49989/


