# How to troubleshoot

## iptables rule reset after reboot

``` bash
iptables -t nat -I PREROUTING -p tcp -d agent_vm_ip/24 --dport 8510 -j DNAT --to-destination 127.0.0.1:8510
sysctl -w net.ipv4.conf.enp0s8.route_localnet=1
```

This iptables rule you set on your Agent VM is not permanent. Once you have restarted your machine, your iptables rule will be gone.
To fix this,

``` bash
apt-get install iptables-persistent
nano /etc/sysctl.conf
```
Paste this: `net.ipv4.conf.enp0s8.route_localnet=1` . Make sure you replace `enp0s8` with correct network adapter that Agent VM IP resides and `agent_vm_ip` with correct Agent VM IP.

## hzn node list, hzn register does not work anymore when exeucted from Management Hub VM

1. Make sure you check your iptables rule in your agent VM and it's properly set in the first place.
2. Make sure you export your `HORIZON_URL` on your management VM with correct Agent VM IP. 
For example, `export HORIZON_URL=http://x.x.x.x:8510`


### Check which API hzn CLI is calling

You can pass `-v` flag to check which API hzn CLI is calling under the hood.

``` bash
hzn node list -v
```