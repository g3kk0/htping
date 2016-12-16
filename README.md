# HyperText Ping

Inspired by ping, htping is a tool for testing/monitoring HTTP endpoints.

```
$ htping www.google.co.uk
connected to 172.217.23.3:80, seq=0, time=0 ms, response=200
connected to 172.217.23.3:80, seq=1, time=0 ms, response=200
connected to 172.217.23.3:80, seq=2, time=0 ms, response=200
connected to 172.217.23.3:80, seq=3, time=0 ms, response=200
```

## Installation

htping is a is a self-contained Go binary (no external dependencies). Simply download the binary to one of the directories in $PATH and run it.

### Linux
```
cd /usr/local/bin

wget https://github.com/g3kk0/htping/releases/download/1.0.0/htping

chmod 755 htping

$ htping www.google.co.uk
connected to 172.217.23.3:80, seq=0, time=0 ms, response=200
```
