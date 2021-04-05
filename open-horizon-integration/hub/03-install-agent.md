# Install the Open Horizon Agent

This continues the instructions from [Install the Open Horizon Hub Services](01-horizon-services-setup.md) and 
[Build and Run](02-build-and-run-horizon.md) the Open Horizon Hub Services.

Stand up your environment in the other tier for the Horizon Agent (Anax) and open a shell.  
Do not attempt this in the same environment as the Horizon Services 
due to port conflicts and environment variable collisions.
Instructions below are for either Ubuntu or OSX.

## For Ubuntu

IMPORTANT, DO NOT SKIP THIS STEP

Become root so the following several steps will work properly

``` bash
sudo -s
```

Update the package manager

``` bash
apt-get -y update
```

Install dependencies

``` bash
apt-get -y install jq
```

Install and test Docker

``` bash
curl -fsSL get.docker.com | sh
docker --version
```

Install the Anax Agent:

Note: We are going to download horizon cli and agent package from open-horizon anax repo releases.

``` bash
wget https://github.com/open-horizon/anax/releases/download/v2.28.0-338/horizon-agent-linux-deb-amd64.tar.gz
tar -xzvf horizon-agent-linux-deb-amd64.tar.gz
dpkg -i horizon-cli_2.28.0-338_amd64.deb
dpkg -i horizon_2.28.0-338_amd64.deb
```

IMPORTANT, exit out of root, back to your user account

``` bash
exit
```

Check the version to confirm that the Horizon CLI is installed and working

``` bash
hzn version
```

Check to ensure that it is working and your machine/node is "unconfigured"

``` bash
hzn node list
```
When you look at the output for `hzn node list`, pay attention to the line for the exchange api:

``` json
"exchange_api": ""
```

When the Agent is properly configured, 
it will point to the public IP address of the Horizon Hub Services that you stood up earlier, and is useful for confirming proper configuration.

To fix this, you will edit the horizon agent configuration file and then restart the agent service.

Please note two important details: first, the protocol is `http` instead of `https` (due to lack of a public signed cert), 
and second, the URL *must* end with a trailing slash, even though the corresponding environment variable does not.  

NOTE: Replace `x.x.x.x` with the actual IP address of the machine running the Open Horizon Hub Services.

Edit the file at `/etc/default/horizon` using `sudo` to set the exchange URL (HZN_EXCHANGE_URL) to `http://x.x.x.x:3090/v1/`.

Also add the MSS URL and set your device id.

Edit the file at `/etc/default/horizon` using `sudo` to set the MMS URL (HZN_FSS_CSSURL) to `http://x.x.x.x:9443`.
Edit the file at `/etc/default/horizon` using `sudo` to set the device id (HZN_DEVICE_ID) to whatever you want, e.g. mynode.

``` bash
nano /etc/default/horizon

HZN_EXCHANGE_URL=http://x.x.x.x:3090/v1/
HZN_FSS_CSSURL=http://x.x.x.x:9443
HZN_MGMT_HUB_CERT_PATH=
HZN_DEVICE_ID=
HZN_AGENT_PORT=8510
```
Restart the Agent service:

``` bash
sudo systemctl restart horizon
```

and confirm that the changes took effect by re-running `hzn node list` and checking the `exchange_api` value.

Also, verify that the "exchange_version" is correct.
If it is an empty string, then the agent does not have network connectivity to the exchange.
You will need to resolve this problem before you continue.

# Next

[Configure the Agent](04-configure-anax.md).