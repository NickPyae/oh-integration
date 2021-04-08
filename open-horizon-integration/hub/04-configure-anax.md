# Configure the Anax Agent

This continues the instructions from [Install the Open Horizon Hub Services](01-horizon-services-setup.md) and 
[Build and Run](02-build-and-run-horizon.md) the Open Horizon Hub Services and 
[Install the Open Horizon Agent](03-install-agent.md) software.

Configure environment variables so the openhorizon CLI can connect to the exchange.

NOTE: Replace `x.x.x.x` with the actual IP address of the machine running the Open Horizon Hub Services.

``` bash
export HZN_EXCHANGE_URL=http://x.x.x.x:3090/v1
export ORG_ID=dellsg
export HZN_ORG_ID=dellsg
export HZN_EXCHANGE_USER_AUTH=admin:adminpw
echo "export HZN_EXCHANGE_USER_AUTH='admin:adminpw'" >> ~/.bashrc
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

Next we will publish an example EdgeX service to the openhorizon hub and then tell the agent to run the service.

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
The digest is recorded in the published service definition.
This example service is composed of several EdgeX microservices.

```
hzn exchange service publish -P -f app-integration/service.json
```

Now check to ensure that the service definition was published and is available in the exchange:

```
hzn exchange service list
```

The above should respond with the following, if successful:

``` json
[
  "dellsg/com.eos2git.cec.lab.emc.hellosally_1.0.1_amd64"
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

Last, let's register the openhorizon agent with the hub, so that it will begin executing the service.

To prepare the node, manually create the volumes that will be used by the edge services:

```
docker volume create db-data
docker volume create log-data
docker volume create consul-data
docker volume create consul-config
```

(Optionally) verify the volumes have been created: `docker volume list`

Manually setup and open the permissions on the host directories being bound:

```
mkdir -p /var/run/edgex/logs
mkdir -p /var/run/edgex/data
mkdir -p /var/run/edgex/consul/data
mkdir -p /var/run/edgex/consul/config
chmod -R a+rwx /var/run/edgex
```

Now register the node:

``` bash
hzn register -p pattern-edgex-amd64 --policy app-integration/node.policy.json
```

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
