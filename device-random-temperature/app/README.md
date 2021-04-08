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

```
rm device-random-temperature

go build .
```

## Add binary as systemd unit

```
nano /etc/systemd/system/device.service 
```

Copy and paste below text to file. Replace env variables accordingly. 
```
[Unit]
Description=Go Device Random Temperature App
[Service]
WorkingDirectory=/root/HelloSally/device-random-temperature/app
ExecStart=/root/HelloSally/device-random-temperature/app/device-random-temperature
// Env Vars
Environment=CORE_SVCS_IP=<core-svcs-ip>
Environment=DEVICE_SVC_IP=<device-svc-ip>
[Install]
WantedBy=multi-user.target
```

These are the default values used for ADDRESSABLE_PORT, CORE_DATA_PORT, CORE_METADATA_PORT:
- ADDRESSABLE_PORT: 49989
- CORE_DATA_PORT: 3080
- CORE_METADATA_PORT: 3081

To override any of these values, please include in /etc/systemd/system/device.service file

```
Environment=CORE_DATA_PORT=3080
Environment=CORE_METADATA_PORT=3081
Environment=ADDRESSABLE_PORT=49989
```

Then, save file.

Reload systemd manager configuration:
```
systemctl daemon-reload  
```

Start service:
```
systemctl start device
```

Check status:
```
systemctl status device
```


## Run

```
./device-random-temperature
```

open browser and go to https://localhost:49989/
(replace 49989 with addressable-port used)

## Usage

- get device PUT command. In Postman, call GET with:
```
https://<core-svcs-ip>:<core-command-port>/api/v1/device
```

- call device PUT command. In Postman, call PUT with:
```
https://<core-svcs-ip>:<core-command-port>/api/v1/device/<device-id>/command/<command-id>

{
    "MinTemperature": 90,
    "MaxTemperature": 100
}
```