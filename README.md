# HyperText Ping

![Build status](https://travis-ci.org/g3kk0/htping.svg?branch=master)

Inspired by ping, htping is a tool for testing and monitoring HTTP endpoints.

```
$ htping www.google.co.uk
connected to 172.217.23.3:80, seq=0, time=68 ms, response=200
connected to 172.217.23.3:80, seq=1, time=68 ms, response=200
connected to 172.217.23.3:80, seq=2, time=67 ms, response=200
connected to 172.217.23.3:80, seq=3, time=74 ms, response=200
```

## Installation

htping is a is a self-contained Go binary (no external dependencies). Simply download the binary to one of the directories in your $PATH and run it.

### Linux

Download
```
wget https://github.com/g3kk0/htping/releases/download/1.0.0/htping.zip -P /tmp
```

Extract
```
unzip /tmp/htping.zip -d /usr/local/bin
```

Run
```
$ htping www.google.co.uk
connected to 172.217.23.3:80, seq=0, time=18 ms, response=200
```
