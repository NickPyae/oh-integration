curl -s -w "%{http_code}" -X POST -H "Content-Type: application/json" -d '[
    {
      "serviceOrgid": "dellsg",
      "serviceUrl": "com.eos2git.cec.lab.emc.hellosally",
      "serviceArch": "amd64",
      "serviceVersionRange": "[0.0.0,INFINITY)",
      "inputs": [
        {
          "name": "KUIPER_IP",
          "value": "172.19.112.49"
        },
        {
          "name": "KUIPER_PORT",
          "value": "48075"
        },
        {
          "name": "INFLUXDB_IP",
          "value": "192.168.1.152"
        },
        {
          "name": "INFLUXDB_PORT",
          "value": "8086"
        },
        {
          "name": "INFLUXDB_TOKEN",
          "value": "m8Btf1vDt5vDx6N7f_Yy0RxwvG0wxzvyjslMZ0XxmREtEpZWBsAhUcd5N020NRt2ZltldQlUiIpLD8Km3hcKjQ=="
        },
        {
          "name": "INFLUXDB_CLOUD_TOKEN",
          "value": "CU4e0jPsdqeSVRL5niynNFJt7pRKG5Xx6ssnPp2Vt4Azd7LFjK0D7Ofg1ElAnfnfW9iJldgLUqDPoFXAbzZ1LA=="
        },
        {
          "name": "BUCKET_NAME",
          "value": "hello-sally"
        },
        {
          "name": "REDIS_IP",
          "value": "172.19.112.49"
        },
        {
          "name": "REDIS_PORT",
          "value": "6379"
        }
      ]
    }
]' http://10.244.14.32:3510/node/userinput
