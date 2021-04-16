#!/bin/sh

echo "[Task 1] Creating edgex stream"
curl -s -X POST http://${KUIPER_IP}:${KUIPER_PORT}/streams \
  -H 'Content-Type: application/json' \
  -d '{
  "sql": "create stream edgex() WITH (FORMAT=\"JSON\", TYPE=\"edgex\")"
}' 2>&1

echo "[Task 2] Creating influx_core_storage_rule"
curl -s -X POST http://${KUIPER_IP}:${KUIPER_PORT}/rules \
  -H 'Content-Type: application/json' \
  -d '{
    "id": "influx_core_storage_rule",
    "sql": "SELECT RandomTemperature, meta(created) as created FROM edgex",
    "actions": [
        {
            "rest": {
                "bodyType": "text",
                "url": "http://'${INFLUXDB_IP}:${INFLUXDB_PORT}'/api/v2/write?bucket='${BUCKET_NAME}'&org=dell-sg&precision=ms",
                "method": "POST",
                "dataTemplate": "hello-sally temperature={{.RandomTemperature}} {{printf \"%.0f\" .created}}",
                "sendSingle": true,
                "headers": {"Authorization": "'"${INFLUXDB_TOKEN}"'"}
            }
        },
        {
            "log": {}
        }
    ]
}' 2>&1

echo "[Task 3] Creating influx_cloud_rule"
curl -s -X POST http://${KUIPER_IP}:${KUIPER_PORT}/rules \
  -H 'Content-Type: application/json' \
  -d '{
    "id": "influx_cloud_rule",
    "sql": "SELECT RandomTemperature, meta(created) as created FROM edgex WHERE RandomTemperature > 90",
    "actions": [
        {
            "rest": {
                "bodyType": "text",
                "url": "https://eastus-1.azure.cloud2.influxdata.com/api/v2/write?bucket='${BUCKET_NAME}'&orgID=78d340669c36cc9a&precision=ms",
                "method": "POST",
                "dataTemplate": "hello-sally temperature={{.RandomTemperature}} {{printf \"%.0f\" .created}}",
                "sendSingle": true,
                "headers": {"Authorization": "'"${INFLUXDB_CLOUD_TOKEN}"'"}
            }
        },
        {
            "log": {}
        }
    ]
}' 2>&1

tail -f /dev/null
