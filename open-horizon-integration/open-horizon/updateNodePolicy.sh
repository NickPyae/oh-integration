curl -s -w "%{http_code}" -X POST -H "Content-Type: application/json" -d '{
    "properties": [
        {
            "name": "tag",
            "value": "hellosally"
        },
        {
            "name": "openhorizon.allowPrivileged",
            "value": true
        }
    ]
}' http://10.244.14.32:3510/node/policy
