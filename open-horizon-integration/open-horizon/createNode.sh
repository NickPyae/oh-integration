curl -s -w "%{http_code}" -X PUT -H "Content-Type: application/json" -u "dellsg/admin:adminpw" -d '{
  "token": "abc",
  "name": "node00",
  "nodeType": "device",
  "pattern": "dellsg/pattern-edgex-amd64",
  "arch": "amd64",
  "registeredServices": [
    {
      "url": "dellsg/com.eos2git.cec.lab.emc.hellosally",
      "numAgreements": 1,
      "policy": "{\"header\":{\"name\":\"Policy for dellsg_com.eos2git.cec.lab.emc.hellosally\",\"version\":\"2.0\"},\"apiSpec\":[{\"specRef\":\"com.eos2git.cec.lab.emc.hellosally\",\"organization\":\"dellsg\",\"version\":\"0.0.1\",\"exclusiveAccess\":true,\"arch\":\"amd64\"}],\"valueExchange\":{},\"dataVerification\":{\"metering\":{}},\"proposalRejection\":{},\"properties\":[{\"name\":\"cpus\",\"value\":\"4\"},{\"name\":\"ram\",\"value\":\"7961\"}],\"ha_group\":{},\"nodeHealth\":{}}",
      "properties": [
        {
          "name": "version",
          "value": "0.0.1",
          "propType": "version",
          "op": "in"
        },
        {
          "name": "cpus",
          "value": "4",
          "propType": "string",
          "op": "in"
        },
        {
          "name": "ram",
          "value": "7961",
          "propType": "string",
          "op": "in"
        },
        {
          "name": "arch",
          "value": "amd64",
          "propType": "string",
          "op": "in"
        }
      ]
    }
  ],
  "userInput": [],
  "msgEndPoint": "",
  "softwareVersions": {"horizon": "1.2.3"},
  "publicKey": "ABCDEF",
  "heartbeatIntervals": {
    "minInterval": 10,
    "maxInterval": 120,
    "intervalAdjustment": 10
  }
}' http://10.244.14.32:3090/v1/orgs/dellsg/nodes/node00
