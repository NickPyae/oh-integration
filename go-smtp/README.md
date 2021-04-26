# Email Alert Service
This Go application queries InfluxDB Cloud at a configurable interval and sends an email when temperatures above a configurable threshold is detected.

An example of the alert email:
```
Subject: [ALERT] Critical Temperature
From: No-Reply: Hello Sally <hello-sally-email>
To: <user-email>

[ALERT]
Device Name:	Random Temperature
Temperature:	1055.0F
Status:		    Critical
Event Time:	    Wed, 21 Apr 2021 06:01:07 UTC
Alert Time:	    Wed, 21 Apr 2021 14:01:34 +08
```

## Configuration
User-configurable values are set using environment variables. Default values are shown below. The remaining variables must be specified before running the application.

```sh
# Default values
export HS_SMTP_PORT=587
export HS_INFLUX_BUCKET=hello-sally-frk
export HS_QUERY_SECONDS=15
export HS_ALERT_THRESHOLD=1000

# Must be specified
export HS_SENDER_EMAIL=<HELLO_SALLY_EMAIL>
export HS_SENDER_APP_KEY=<HELLO_SALLY_EMAIL_APP_KEY>
export HS_USER_EMAIL=<RECEIVER_EMAIL>
export HS_SMTP_HOST=<SMTP_HOST>
export HS_INFLUX_URL=<INFLUX_CLOUD_URL>
export HS_INFLUX_TOKEN=<INFLUX_CLOUD_TOKEN>
export HS_INFLUX_ORG=<INFLUX_CLOUD_ORG>
```

Notes:
- Every `n` seconds, the application will query readings in the last `2n` seconds inside InfluxDB Cloud. `HS_QUERY_SECONDS` is the variable that represents `n`.
- `HS_USER_EMAIL` can take multiple emails, however it must be comma-delimited.
```sh
export HS_USER_EMAIL=<RECEIVER_EMAIL_1>,<RECEIVER_EMAIL_2>,<RECEIVER_EMAIL_3>
```

## Deploy Natively
```sh
git clone https://eos2git.cec.lab.emc.com/ISG-Edge/HelloSally.git
cd HelloSally/go-smtp
go mod tidy
go run main.go
```

## Deploy on Docker
Build and push image to artifactory
```sh
docker build -t amaas-eos-mw1.cec.lab.emc.com:5070/hellosally/email-alert:0.0.1 .
docker push amaas-eos-mw1.cec.lab.emc.com:5070/hellosally/email-alert:0.0.1
```

Since Azure cannot access artifactory, we have to manually save and load the Docker image into Azure.
```sh
docker save amaas-eos-mw1.cec.lab.emc.com:5070/hellosally/email-alert:0.0.1 > email-alert.tar
scp email-alert.tar <USER>@<AZURE_VM_IP>:~/

# Inside Azure VM
docker load < email-alert.tar
```

We can specify environment variables using a `.env` file. Prior to running the container, make a copy of `default.env` and edit the values accordingly. (See [Configuration](#configuration))
```sh
cp default.env .env

# Run
docker run --name email-alert-service -d --env-file .env --restart unless-stopped amaas-eos-mw1.cec.lab.emc.com:5070/hellosally/email-alert:0.0.1
```