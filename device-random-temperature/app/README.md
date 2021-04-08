# Temperature simulator: device-random-temperature

The application creates device service on start, and creates a pre-defined device named 'Random-Temperature-Generator01'. This device generates random temperature reading between 50 - 200 Fahrenheit (every 1 second).

To check:
```
curl -v <core-svcs-ip>:<core-data-port>/api/v1/event/device/Random-Temperature-Generator01/10
```

It also provides API endpoints for:
- device service to register device
- core command to request action of device with command (via REST PUT) 



Reference: https://docs.edgexfoundry.org/1.3/examples/LinuxTutorial/EdgeX-Foundry-tutorial-ver1.1.pdf (page 34 - 41)


## Pre-requisites

Ensure EdgeX Core Services (Core Metadata, Core Data, Core Command) are running.

Please check:
```
curl -v <core-svcs-ip>:<core-data-port>/api/v1/ping
curl -v <core-svcs-ip>:<core-metadata-port>/api/v1/ping
curl -v <core-svcs-ip>:<core-command-port>/api/v1/ping
```

## Build

```
git clone https://eos2git.cec.lab.emc.com/ISG-Edge/HelloSally.git
```

```
cd HelloSally/device-random-temperature/app
```

set required environment variables:
```
export CORE_SVCS_IP="https://<core-svcs>"
export CORE_DATA_PORT="<core-data-port>"
export CORE_METADATA_PORT="<core-metadata-port>"
export ADDRESSABLE_PORT="<addressable-port>"
```

```
go build .
```

## Run

```
./device-random-temperature.exe
```

open browser and go to https://localhost:49989/

## Usage

- get device PUT command. In Postman, call GET with:
```
https://<core-svcs-ip>/api/v1/device
```

- call device PUT command. In Postman, call PUT with:
```
https://<core-svcs-ip>:<core-command-port>/api/v1/device/<device-id>/command/<command-id>
```

```
{
    "MinTemperature": 90,
    "MaxTemperature": 100
}
```