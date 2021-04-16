# Build and Run Open Horizon Management Hub Services

This continues the instructions from [Install the Open Horizon Management Hub Services](01-horizon-services-setup.md).

*NOTE*: This assumes you are running as root (with `sudo`) on the Ubuntu VM 
as described in the previous set of instructions and are continuing in the same session. 
If you are not, you may need to export all previous environment variables from step 1.

### Create and Persist Environment Variables

``` bash
export MY_IP=`ifconfig | egrep 'inet ' | sed 's/addr://' | awk '{ print $2 }' | egrep -v '^172.|^10.|^127.' | head -1`
echo "export MY_IP=${MY_IP}" >> ~/.bashrc
echo "export HZN_ORG_ID=dellsg" >> ~/.bashrc
echo "export HZN_EXCHANGE_URL=http://${MY_IP}:3090/v1/" >> ~/.bashrc
echo "export HZN_FSS_CSSURL=http://${MY_IP}:9443" >> ~/.bashrc
echo "export HZN_EXCHANGE_ROOT_USER_AUTH=root/root:Horizon-Rul3s" >> ~/.bashrc
echo "export HZN_EXCHANGE_USER_AUTH=admin:adminpw" >> ~/.bashrc
source ~/.bashrc
```

### Edit Management Hub Services Configuration JSON

NOTE: Again, this is assuming that you cloned this repository and it is located at `./open-horizon-integration/`.

Change to the `oh` folder and and generate the hub config file, config.json.
Edit the `config.json` file if you would like to make changes.

``` bash
cd open-horizon-integration/hub/oh
envsubst '${MY_IP}' < config.json.template > config.json
```

NOTE: The following text refers to nodes in the `config.json` file using XPath-like dot notation.

Your IP address for this machine is found at `.horizon.hostname`.

The AgBot's name and token are specified at `.services.agbot.bot`.

The Exchange's root credentials are at `.exchange.root` and `.exchange.password`.

The admin user for your org is at `.exchange.admin.username` , and `.exchange.admin.password`.

The organization is specified at `.horizon.namespace` and `.exchange.org`.

These instructions are assuming, and will be referring to, the default values specified at the nodes listed above.

### Build and Start the Services

NOTE: I am assuming that you are still in the `horizon-edgex/oh` directory.

``` bash
make
make up
```

NOTE: `openhorizon/amd64_exchange-api` container will exit due to file permission. To verify this, run `docker ps -a`, you will see that this particular container exited. To fix this issue, run below command in current `horizon-edgex/oh` directory

``` bash
chmod -R 777 exchange
make up
```

Install Open Horizon hzn CLI:

Note: We are going to download horizon cli from open-horizon anax repo releases.

``` bash
wget https://github.com/open-horizon/anax/releases/download/v2.28.0-338/horizon-agent-linux-deb-amd64.tar.gz
tar -xzvf horizon-agent-linux-deb-amd64.tar.gz
dpkg -i horizon-cli_2.28.0-338_amd64.deb
```

Let's confirm that the Exchange is running (and our environment variable are configured correctly) by 
requesting the endpoints for `version` and `status`:

``` bash
curl -u ${HZN_EXCHANGE_ROOT_USER_AUTH} ${HZN_EXCHANGE_URL}admin/version
```

``` bash
curl --silent -u ${HZN_EXCHANGE_ROOT_USER_AUTH} ${HZN_EXCHANGE_URL}admin/status | jq .
```

If all is well, let's continue by listing the existing Organizations:

``` bash
curl --silent -u ${HZN_EXCHANGE_ROOT_USER_AUTH} ${HZN_EXCHANGE_URL}orgs | jq .
```

Add an Organization named `dellsg`:

``` bash
curl -sSf -X POST -u ${HZN_EXCHANGE_ROOT_USER_AUTH} -H "Content-Type:application/json" -d '{"label": "dellsg", "description": "Organization for Testing"}' ${HZN_EXCHANGE_URL}orgs/dellsg | jq .
```

And then list the existing Organizations again to see `dellsg` now in the list:

``` bash
curl --silent -u ${HZN_EXCHANGE_ROOT_USER_AUTH} ${HZN_EXCHANGE_URL}orgs | jq .
```

If all is well, let's "prime the pump" by clearing the tables and refilling to a known state:

``` bash
make prime
```

The step you just performed created an admin user for your dellsg and also told your new AgBot `agbot1` to listen for Deployment Patterns and Policies from the `dellsg` Organization.

List the current users in `dellsg`:

``` bash
curl --silent -sSf -u ${HZN_EXCHANGE_ROOT_USER_AUTH} ${HZN_EXCHANGE_URL}orgs/dellsg/users | jq .
```

## Next

[Install the Open Horizon Agent](03-install-agent.md).
