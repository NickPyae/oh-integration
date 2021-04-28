# Device controller UI

The application provides a form to call PUT core command of device named 'Random-Temperature-Generator01'. 

By default, the device generates random temperature reading between 50 - 200 Fahrenheit (every 1 second).
Using this form, user can temporary change the Temperature range (Min, Max) for a specified duration.


## Build

```
git clone https://eos2git.cec.lab.emc.com/ISG-Edge/HelloSally.git
```

```
cd HelloSally/device-controller-ui
```

```
rm device-controller-ui

go build .
```


## Run

```
./device-controller-ui
```

open browser and go to http://localhost:49990/

