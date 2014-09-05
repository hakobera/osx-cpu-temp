# OSX CPU Temp

Outputs current CPU temperature in °C for OSX and publis it to MQTT broker.

## Usage 

### Compiling
```bash
make
```

### Running

```bash
SANGO_USERNAME="xxx" SANGO_PASSWORD="yyy" go run main.go
```

### Output example

```
publish:
TOPIC: hakobera@github/testgo
MSG: 54.0°C
subscribe:
TOPIC: hakobera@github/testgo
MSG: 54.0°C
```

### Source 

Apple System Management Control (SMC) Tool 
Copyright (C) 2006

### Inspiration 

 * http://www.eidac.de/smcfancontrol/
 * https://github.com/hholtmann/smcFanControl/tree/master/smc-command
