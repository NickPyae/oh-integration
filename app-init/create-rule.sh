#!/bin/sh

echo "[Task 1] Creating edgex stream"
curl -s -X POST http://${KUIPER_IP}:${KUIPER_PORT}/streams \
  -H 'Content-Type: application/json' \
  -d '{
  "sql": "create stream edgex() WITH (FORMAT=\"JSON\", TYPE=\"edgex\")"
}' 2>&1

echo "[Task 2] Creating core_storage_rule rule"
curl -s -X POST http://${KUIPER_IP}:${KUIPER_PORT}/rules \
  -H 'Content-Type: application/json' \
  -d '{
    "id": "core_storage_rule",
    "sql": "SELECT meta(created) AS created, RandomTemperature FROM edgex",
    "actions": [
        {
            "mqtt": {
                "clientId": "ece-ubuntu-kuiper",
                "server": "tcp://192.168.1.152:1883",
                "topic": "oh/kuiper/temp"
            }
        },
        {
            "log": {}
        }
    ]
}' 2>&1

echo "[Task 3] Creating influx cloud rule"
curl -s -X POST http://${KUIPER_IP}:${KUIPER_PORT}/rules \
  -H 'Content-Type: application/json' \
  -d '{
    "id": "influx_cloud_rule",
    "sql": "SELECT RandomTemperature, meta(created) as created FROM edgex WHERE RandomTemperature > 90",
    "actions": [
        {
            "rest": {
                "bodyType": "text",
                "url": "https://eastus-1.azure.cloud2.influxdata.com/api/v2/write?bucket=hello-sally&orgID=78d340669c36cc9a&precision=ms",
                "method": "POST",
                "dataTemplate": "hello-sally temperature={{.RandomTemperature}} {{printf \"%.0f\" .created}}",
                "sendSingle": true,
                "headers": {"Authorization": "Token n_xnacdL6mGdkSaGr4EHgt-csOz49fPbZNLRMKZ0W_L6B8-N7Dd77QrDSzQ_KJHe8ys1zYIUiqGbmV5lxlc25g=="}
            }
        },
        {
            "log": {}
        }
    ]
}' 2>&1
