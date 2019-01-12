# cookiecutter-golang

[![Build Status](https://travis-ci.org/lacion/cookiecutter-golang.svg?branch=master)](https://travis-ci.org/lacion/cookiecutter-golang)

Powered by [Cookiecutter](https://github.com/audreyr/cookiecutter), Cookiecutter Golang is a framework for jumpstarting production-ready go projects quickly.

## Features

- Generous `Makefile` with management commands
- Uses `go dep` (with optional go module support *requires go 1.11*)
- injects build time and git hash at build time.

## Optional Integrations

- Can use [viper](https://github.com/spf13/viper) for env var config
- Can use [cobra](https://github.com/spf13/cobra) for cli tools
- Can use [logrus](https://github.com/sirupsen/logrus) for logging
- Can create dockerfile for building go binary and dockerfile for final go binary (no code in final container)
- If docker is used adds docker management commands to makefile
- Option of TravisCI, CircleCI or None

## Constraints

- Uses `dep` for dependency management
- Only maintained 3rd party libraries are used.

This project now uses docker multistage builds, you need at least docker version v17.05.0-ce to use the docker file in this template, [you can read more about multistage builds here](https://www.critiqus.com/post/multi-stage-docker-builds/).

## Docker

This template uses docker multistage builds to make images slimmer and containers only the final project binary and assets with no source code whatsoever.

You can find the image dokcer file in this [repo](https://github.com/lacion/alpine-golang-buildimage) and more information about docker multistage builds in this [blog post](https://www.critiqus.com/post/multi-stage-docker-builds/).

Apps run under non root user and also with [dumb-init](https://github.com/Yelp/dumb-init).

## Usage

Let's pretend you want to create a project called "echoserver". Rather than starting from scratch maybe copying 
some files and then editing the results to include your name, email, and various configuration issues that always 
get forgotten until the worst possible moment, get cookiecutter to do all the work.

First, get Cookiecutter. Trust me, it's awesome:
```console
$ pip install cookiecutter
```

Alternatively, you can install `cookiecutter` with homebrew:
```console
$ brew install cookiecutter
```

Finally, to run it based on this template, type:
```console
$ cookiecutter https://github.com/lacion/cookiecutter-golang.git
```

You will be asked about your basic info (name, project name, app name, etc.). This info will be used to customize your new project.

Warning: After this point, change 'Luis Morales', 'lacion', etc to your own information.

Answer the prompts with your own desired [options](). For example:
```console
full_name [Luis Morales]: Luis Morales
github_username [lacion]: lacion
app_name [mygolangproject]: echoserver
project_short_description [A Golang project.]: Awesome Echo Server
docker_hub_username [lacion]: lacion
docker_image [lacion/docker-alpine:latest]: lacion/docker-alpine:latest
docker_build_image [lacion/docker-alpine:gobuildimage]: lacion/docker-alpine:gobuildimage
use_docker [y]: y
use_git [y]: y
use_logrus_logging [y]: y
use_viper_config [y]: y
use_cobra_cmd [y]: y
Select use_ci:
1 - travis
2 - circle
3 - none
Choose from 1, 2, 3 [1]: 1
```

Enter the project and take a look around:
```console
$ cd echoserver/
$ ls
```

Run `make help` to see the available management commands, or just run `make build` to build your project.
```console
$ make help
$ make build
$ ./bin/echoserver
```

## Projects build with cookiecutter-golang

- [iothub](https://github.com/lacion/iothub) websocket multiroom server for IoT

## 补充说明

`注意`：

- 本脚手架默认使用读取参数的方式是pflag，后期考虑加入flag的识别
- 有两个组件是耦合在一起的，use_logrus_logging和use_viper_config，两两组合的话，总共有4种可能：

1. use_logrus_logging [y] and use_viper_config [y]，这种情况，将pflag解析到的所有命令行参数绑定到：viper_config，如果有configfile参数，viper_config还会读取该参数对应的配置文件（`注意`：按照这个读取顺序，配置文件中的配置项会覆盖掉命令行中的配置项，比如：json_logs/loglevel/logfile），再用总的配置信息来重置logrus_logging
2. use_logrus_logging [y] and use_viper_config [n]，这种情况，将pflag解析到的所有命令行的参数传给logrus_logging，logrus_logging的时候会根据三个参数来重置配置——json_logs/loglevel/logfile
3. use_logrus_logging [n] and use_viper_config [y]，这种情况，参考情况1，只是不去重置logrus_logging
3. use_logrus_logging [n] and use_viper_config [n]，两个模块都不用的情况

只有两种方式读取配置信息：

1. 通过pflag解析命令行参数
2. 通过配置文件
