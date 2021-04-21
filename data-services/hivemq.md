# Installation information

HiveMQ CE can be packaged into a zip file, which contains the executables, init scripts and sample configurations.

The zip contains the following directories:

| Folder name          | Description                                                                                                  |
| -------------------- | :----------------------------------------------------------------------------------------------------------- |
| bin                  | The folder with start scripts and binary files.                                                              |
| conf                 | The folder with the [configurations](https://github.com/hivemq/hivemq-community-edition/wiki/Configuration). |
| data                 | Persistent client data and cluster data are located here.                                                    |
| license              | The folder where the HiveMQ License file(s) resides.                                                         |
| log                  | All log files can be found here.                                                                             |
| extensions           | The folder where extensions reside.                                                                          |
| third-party-licenses | Information about the licenses of third party libraries can be found here.                                   |

## Installation on Unix based systems (Linux, BSD, MacOS X, Unix)

The default installation directory is `/opt/hivemq` and the default user to run HiveMQ is named `hivemq`. If you need to install HiveMQ to a custom directory or run it under a custom user please be aware of changing the `$HIVEMQ_DIRECTORY` and/or the `HIVEMQ_USER` in the `$HIVEMQ_DIRECTORY/bin/start.sh` script.

1. Login as root
   Some of the following commands need root privileges, please login as root or use sudo to execute the commands.
2. Check out the git repository and build the binary package.

   ```bash
   git clone https://github.com/hivemq/hivemq-community-edition.git

   cd hivemq-community-edition

   ./gradlew clean packaging
   ```

3. Go to the folder containing the zip file

   ```bash
   cd build/zip/
   ```

4. Extract the files

   ```bash
   unzip hivemq-ce-<version>.zip
   ```

5. Create hivemq symlink
   ```bash
   ln -s /opt/hivemq-ce-<version> /opt/hivemq
   ```
6. Create HiveMQ user
   ```bash
   useradd -d /opt/hivemq hivemq
   ```
7. Make scripts executable and change owner to `hivemq` user
   ```bash
   chown -R hivemq:hivemq /opt/hivemq-ce-<version>
   chown -R hivemq:hivemq /opt/hivemq
   cd /opt/hivemq
   chmod +x ./bin/run.sh
   ```
8. Adjust the configuration properties to your needs.
   See chapter [Configuration](https://github.com/hivemq/hivemq-community-edition/wiki/Configuration) for detailed instructions how to configure HiveMQ.
9. Install the init script (optional)
   For Debian-based linux like Debian, Ubuntu, Raspbian using init.d scripts
   ```bash
   cp /opt/hivemq/bin/init-script/hivemq-debian /etc/init.d/hivemq
   chmod +x /etc/init.d/hivemq
   ```
   For linux systems using systemd
   ```bash
   cp /opt/hivemq/bin/init-script/hivemq.service /etc/systemd/system/hivemq.service
   ```
   For all other linux systems
   ```bash
    cp /opt/hivemq/bin/init-script/hivemq /etc/init.d/hivemq
    chmod +x /etc/init.d/hivemq
   ```
10. Modify /etc/init.d/hivemq (optional)
    Set the HIVEMQ_HOME and the HIVEMQ_USER variable to the correct values for your system.

    By default this would be:

    `HIVEMQ_HOME=/opt/hivemq`

    `HIVEMQ_USER=hivemq`

    If you installed HiveMQ to a different directory than `/opt/hivemq` please point the `HIVEMQ_HOME` in your init script to the correct directory. Otherwise the daemon will not start correctly.

11. Start HiveMQ on boot (optional)
    For Debian-based linux like Debian, Ubuntu, Raspbian
    ```bash
    update-rc.d hivemq defaults
    ```
    For Debian-based linux like Debian, Ubuntu, Raspbian using systemd
    ```bash
    systemctl enable hivemq
    ```

### Starting HiveMQ

The following instructions show how to start HiveMQ after the installation.

Starting manually

1. Change directory to HiveMQ directory
   ```bash
   cd /opt/hivemq
   ```
2. Execute startup script
   ```bash
   ./bin/run.sh
   ```

Starting as daemon

```bash
/etc/init.d/hivemq start
```

Verify if HiveMQ is running
Check if HiveMQ is listening to the default port for MQTT

```bash
netstat -an | grep 1883
```

If you’re running HiveMQ as daemon:

```bash
/etc/init.d/hivemq status
```
