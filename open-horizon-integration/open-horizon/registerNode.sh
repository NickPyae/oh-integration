curl -s -w "%{http_code}" -X PUT -H "Content-Type: application/json" -d '{
    "state": "configured"
}' http://10.244.14.32:3510/node/configstate
