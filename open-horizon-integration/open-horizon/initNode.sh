curl -s -w "%{http_code}" -X POST -H "Content-Type: application/json" -d '{
    "id": "node00",
    "organization": "dellsg",
    "pattern": "dellsg/pattern-edgex-amd64",
    "name": "node00",
    "token": "abc"
}' http://10.244.14.32:3510/node
