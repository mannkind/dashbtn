# dashbtn

[![Software
License](https://img.shields.io/badge/License-MIT-orange.svg?style=flat-square)](https://github.com/mannkind/dashbtn/blob/master/LICENSE.md)
[![Travis CI](https://img.shields.io/travis/mannkind/dashbtn/master.svg?style=flat-square)](https://travis-ci.org/mannkind/dashbtn)
[![Coverage Status](https://img.shields.io/codecov/c/github/mannkind/dashbtn/master.svg)](http://codecov.io/github/mannkind/dashbtn?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/mannkind/dashbtn)](https://goreportcard.com/report/github.com/mannkind/dashbtn)

Making Amazon Dash buttons do different things using dnsmasq

# Installation

* go get github.com/mannkind/dashbtn
* go intall github.com/mannkind/dashbtn
* dashbtn -c */the/path/to/config.yaml*

# Configuration

Configuration happens in the config.yaml file. A full example might look this:

```
'74:75:48:C3:B1:D0':
    'add+old': [ '/usr/local/bin/mosquitto_pub', '-h', 'mosquitto', '-t', 'home/dash_btn', '-m', 'ON' ]
    'del': [ '/usr/local/bin/mosquitto_pub', '-h', 'mosquitto', '-t', 'home/dash_btn', '-m', 'ON' ]

'74:75:48:C3:B1:D1+74:75:48:C3:B1:D2':
    'add+old': [ 'say', 'Someone is at the door. Please go see who it is' ]
```

# Tomato/OpenWRT/etc Script
```
#!/bin/sh
mode="$1" # "add", "del", or "old"
mac="$2"
ip="$3"
host="$4"
wget -O - "http://HOST_RUNNING_DASHBTN:PORT/dash?mode=$mode&mac=$mac&ip=$ip&host=$host" >/dev/null 2>&1 &
```
