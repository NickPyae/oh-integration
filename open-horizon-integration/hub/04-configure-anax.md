# Configure the Anax Agent VM

This continues the instructions from [Install the Open Horizon Management Hub Services](01-horizon-services-setup.md) and 
[Build and Run](02-build-and-run-horizon.md) the Open Horizon Management Hub Services and 
[Install the Open Horizon Agent](03-install-agent.md).

Configure environment variables so the Open Horizon hzn CLI can connect to the exchange from Agent VM.

NOTE: Replace `x.x.x.x` with the actual IP address of the machine running the Open Horizon Management Hub Services.

``` bash
echo "export HZN_EXCHANGE_URL=http://x.x.x.x:3090/v1" >> ~/.bashrc
echo "export ORG_ID=dellsg" >> ~/.bashrc
echo "export HZN_ORG_ID=dellsg" >> ~/.bashrc
echo "export HZN_EXCHANGE_USER_AUTH=admin:adminpw" >> ~/.bashrc
source ~/.bashrc

hzn exchange user list
```

The results of `hzn exchange user list` should be something like the following:

``` json
{
  "dellsg/admin": {
    "admin": true,
    "email": "admin@dellsg",
    "lastUpdated": "2019-08-19T12:25:06.754Z[UTC]",
    "password": "********",
    "updatedBy": "root/root"
  }
}
```

Before we deploy any services using hzn CLI, there are two options:

1. Deploy all services `ONLY from Management Hub VM` to Agent VM using hzn CLI
2. Deploy all services from Agent VM itself using hzn CLI. This is default way of deploying open horizon services so far.

For Option 1, please follow this guide: [Exposing Open Horizon Agent API to Outside](06-expose-agent-api.md) and come back here for the remaining steps below to deploy services.

For Option 2, you can follow the remaining steps below to deploy services.

Next we will publish an example EdgeX service to Open Horizon Management hub and then tell the agent to run the service.

``` bash
git clone https://eos2git.cec.lab.emc.com/ISG-Edge/HelloSally.git
cd ./HelloSally
```

First, we'll generate an RSA key pair to be used for signing the edge service.
A service is one or more microservices that are deployed on an edge device as containers.
A service has to be signed so that the agent can verify that autheticity of the service definition.
This command will put the keys into a default location where other CLI commands know where to find them.

``` bash
hzn key create -l 4096 dellsg john@everywhere.com
```

This will return something like:

```
Created keys:
        /root/.hzn/keys/service.private.key
        /root/.hzn/keys/service.public.pem
```

Publish the _service_ definition to the exchange.
The command will run for a bit, and will pull each container from the container registry so that it can obtain the container digest.
The digest is recorded in the published service definition. Provide your JFrog Artifactory user name and API Key.

Note: this below command is only required if you are publishing your service from HOP or Franklin Labs.
```
hzn exchange service publish -P -r "amaas-eos-mw1.cec.lab.emc.com:5070:your_user_name:your_api_key" -f app-integration/service.json
```

The -r "amaas-eos-mw1.cec.lab.emc.com:5070:your_user_name:your_api_key" argument indicates to Horizon exchange which Docker container registry to use, and indicates your credentials, such as your API key. This will pull the images from private JFrog Artifactory.
This will also put your private Container Registry API Key into Horizon exchange under the service definition so Horizon edge nodes can automatically retrieve the service definition when needed.

If you are testing locally on your VirtualBox, use 
```
hzn exchange service publish -P -f app-integration/service.dockerhub.json
```
This will pull the images from public Docker Hub instead.

Now check to ensure that the service definition was published and is available in the exchange:

```
hzn exchange service list
```

The above should respond with the following, if successful:

``` json
[
  "dellsg/com.eos2git.cec.lab.emc.hellosally_0.0.1_amd64"
]
```

Next, publish a _pattern_ to the exchange.
A pattern is the easiest way for a node to indicate which services it should run.
Policy based service deployment is also supported, but is slightly more complex to setup.

```
hzn exchange pattern publish -f app-integration/pattern.json
```

Now check to ensure the pattern is available:

```
hzn exchange pattern list
```

It should respond with:

``` json
[
  "dellsg/pattern-edgex-amd64"
]
```

Export KUIPER_IP and KUIPER_PORT as environment variables to host shell environment so that these env variables can be injected into app-init container during runtime.

NOTE: Replace `x.x.x.x` with the actual IP address of the machine running kuiper rule engine service.

``` bash
export KUIPER_IP=x.x.x.x
export KUIPER_PORT=48075
```

Last, let's register the openhorizon agent with the hub, so that it will begin executing the service.

``` bash
hzn register -p pattern-edgex-amd64 --policy app-integration/node.policy.json -f app-integration/userinput.json
```

The argument, -f app-integration/userinput.json informs the registration command to use the userinput file to start any configuration variables that are used by any of the services in the registered deployment pattern. The file provides configuration variable values at registration time. When you pass this file on the hzn register ... -f ... command line, the hzn command first runs envsubst to enhance the file by extracting values from the host shell environment. 

To confirm that your edge node is registered for the pattern, run:

``` bash
hzn node list
```

and confirm that the response shows your node ID. 
and that you're configured for the `dellsg/pattern-edgex-amd64` pattern.

To check on the status of the agreement, use:

``` bash
hzn agreement list
```

or to see verbose details:

``` bash 
hzn eventlog list
```

And once the agreement is finalized, you should be able to see the containers running:

``` bash
sudo docker ps
```

# Next

[View the Device Data](05-view-device-data.md).
