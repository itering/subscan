# Subscan Essentials

![License: GPL](https://img.shields.io/badge/license-GPL-blue.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/itering/subscan)](https://goreportcard.com/report/github.com/itering/subscan)
![subscan](https://github.com/itering/subscan/workflows/subscan/badge.svg)

Subscan Essentials是一个高精度的区块链浏览器脚手架项目，它具有开发人员友好的界面和自定义模块解析功能，支持基于substrate的区块链网络。
它由Subscan团队开发，并在subscan.io中使用。开发人员可以自由使用代码库来扩展功能并为其受众开发独特的用户体验。


## Contents

- [Subscan Essentials](#Subscan-Essentials)
  - [Contents](#contents)
  - [Feature](#Feature)
  - [Requirement](#Requirement)
  - [Installation](#Install)
  - [Usage](#Usage)
  - [Docker](#Docker)
  - [LICENSE](#LICENSE)
  - [Resource](#Resource)


### Feature

1. 支持substrate 网络[自定义](/custom_type.md)type注册 
2. 支持索引block, Extrinsic, Event, log
3. 可自定义[插件](/plugins)索引更多的数据
4. [Gen](https://github.com/itering/subscan-plugin/tree/master/tool) 工具可自动生成插件模版
5. 内置默认的HTTP API [DOC](/docs/index.md)

### Requirement

* Linux / Mac OSX
* Golang 1.12.4+
* Redis 3.0.4+
* MySQL 5.6+

### Install

```bash
./build.sh install
```


### Usage

```
NAME:
   SubScan - SubScan Backend Service, use -h get help

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   1.0

DESCRIPTION:
   SubScan Backend Service, substrate blockchain explorer

COMMANDS:
     start    Start one worker, E.g substrate
     stop     Stop one worker, E.g substrate
     install  Create database and create default conf file
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --conf value   (default: "../configs")
   --help, -h     show help
   --version, -v  print the version


```


### Docker

> db 

```bash
docker-compose -f docker-compose.db.yml up  -d
```

> subscan

```bash
docker-compose build
docker-compose up -d
```

## LICENSE

GPL-3.0


## Resource
 
- [ITERING] https://github.com/itering
- [Darwinia] https://github.com/darwinia-network/darwinia