# OpenWRT-WAN-reconnect

[![GoTemplate](https://img.shields.io/badge/go/template-black?logo=go)](https://github.com/SchwarzIT/go-template)

Reconnecting the PPPoE interface in luci when the Internet is not available.

The project uses `make` to make your life easier. If you're not familiar with Makefiles you can take a look at [this quickstart guide](https://makefiletutorial.com).

Whenever you need help regarding the available actions, just use the following command.

```bash
make help
```

## Setup

To get your setup up and running the only thing you have to do is

```bash
make all
```

This will initialize a git repo, download the dependencies in the latest versions and install all needed tools.
If needed code generation will be triggered in this target as well.

## Run

```bash
make run
```

## Test & lint

Run linting

```bash
make lint
```

Run tests

```bash
make test
```

## detector

http

```yaml
detector:
    type: "http"
    url: "http://baidu.com"
```

## connector

ssh

```yaml
type: "ssh"
host: "192.168.1.1:22"
username: "root"
password: "123456"
```

clash

```yaml
type: "clash"
host: "http://192.168.1.1:9090"
token: "123456"
selector: "selector"
```
