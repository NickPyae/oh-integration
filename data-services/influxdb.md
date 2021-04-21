## Install InfluxDB as a service with systemd

1. Download and install the appropriate `.deb` or `.rpm` file using a URL from the InfluxData downloads page with the following commands:

   ```bash
   # Ubuntu/Debian
   wget https://dl.influxdata.com/influxdb/releases/influxdb2-2.0.4-amd64.rpm.deb
   sudo dpkg -i influxdb2-2.0.4-amd64.rpm.deb

   # Red Hat/CentOS/Fedora
   wget https://dl.influxdata.com/influxdb/releases/influxdb2-2.0.3-amd64.rpm
   sudo yum localinstall influxdb2-2.0.3-amd64.rpm
   ```

2. Start the InfluxDB service:

   ```bash
   sudo service influxdb start
   ```

   Installing the InfluxDB package creates a service file at `/lib/systemd/services/influxdb.service` to start InfluxDB as a background service on startup.

3. Restart your system and verify that the service is running correctly:

   ```bash
   sudo service influxdb status
   ```

   When installed as a service, InfluxDB stores data in the following locations:

   - Time series data: `/var/lib/influxdb/engine/`
   - Key-value data: `/var/lib/influxdb/influxd.bolt`
   - influx CLI configurations: `~/.influxdbv2/configs` (see `influx config` for more information)

To customize your InfluxDB configuration, use either [command line flags (arguments)](https://docs.influxdata.com/influxdb/v2.0/get-started/?t=Linux#pass-arguments-to-systemd), environment variables, or an InfluxDB configuration file. See InfluxDB [configuration options](https://docs.influxdata.com/influxdb/v2.0/reference/config-options/) for more information.

### Pass arguments to systemd

1. Add one or more lines like the following containing arguments for `influxd` to `/etc/default/influxdb2`:

   ```bash
   ARG1="--http-bind-address :8087"
   ARG2="<another argument here>"
   ```

2. Edit the `/lib/systemd/system/influxdb.service` file as follows:
   ```bash
   ExecStart=/usr/bin/influxd $ARG1 $ARG2
   ```
