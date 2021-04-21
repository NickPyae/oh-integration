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

If deploying on Docker, we can specify environment variables using a `.env` file. Make a copy of `default.env` and edit the values accordingly.
```sh
cp default.env .env
```

## Deploy Natively
```sh
git clone https://eos2git.cec.lab.emc.com/ISG-Edge/HelloSally.git
cd HelloSally/go-smtp
go mod tidy
go run main.go
```

## Deploy on Docker
```sh
docker build -t email-alert:0.0.1 .
docker run --name email-alert-service -d --env-file .env email-alert:0.0.1
```
