## connector-cpanel (uplink v1.0.5)

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/5da0d3aaeb1c4b608b1b7d79521e0bfb)](https://app.codacy.com/gh/storj-thirdparty/connector-cpanel?utm_source=github.com&utm_medium=referral&utm_content=storj-thirdparty/connector-cpanel&utm_campaign=Badge_Grade_Dashboard)
[![Go Report Card](https://goreportcard.com/badge/github.com/storj-thirdparty/connector-cpanel)](https://goreportcard.com/report/github.com/storj-thirdparty/connector-cpanel)
![Cloud Build](https://storage.googleapis.com/storj-utropic-services-badges/builds/connector-cpanel/branches/master.svg)


## Overview

The cPanel Connector connects to a cPanel server, takes a backup of the specified files and uploads the backup data on Storj network.

```bash
Usage:
  connector-cpanel [command] <flags>

Available Commands:
  help        Help about any command
  store       Command to upload data to a Storj V3 network.
  version     Prints the version of the tool

```

`store` - Connect to the specified cpanel (default: `cpanel_property.json`).  Back-up of the cpanel is generated using tooling provided by cpanel and then uploaded to the Storj network.  Connect to a Storj v3 network using the access specified in the Storj configuration file (default: `storj_config.json`).

The following flags  can be used with the `store` command:

* `accesskey` - Connects to the Storj network using a serialized access key instead of an API key, satellite url and encryption passphrase.
* `shared` - Generates a restricted shareable serialized access with the restrictions specified in the Storj configuration file.

Sample configuration files are provided in the `./config` folder.

## Requirements and Install

To build from scratch, [install the latest Go](https://golang.org/doc/install#install).

> Note: Ensure go modules are enabled (GO111MODULE=on)

#### Option #1: clone this repo (most common)

To clone the repo

```
git clone https://github.com/storj-thirdparty/connector-cpanel.git
```
Then, build the project using the following:

```
cd connector-cpanel
go build
```

#### Option #2:  ``go get`` into your gopath

 To download the project inside your GOPATH use the following command:

```
go get github.com/storj-thirdparty/connector-cpanel
```

## Upload executable files on server

Place the executable file along with configuration files to the user's home directory. 

## Run (short version)

Once you have built the project run the following commands as per your requirement:

##### Get help

```
$ ./connector-cpanel --help
```

##### Check version

```
$ ./connector-cpanel --version
```

##### Create backup from cPanel and upload to Storj

```
$ ./connector-cpanel store
```

## Documentation

For more information on runtime flags, configuration, testing, and diagrams, check out the [Detail](//github.com/storj-thirdparty/connector-cpanel/wiki/) or jump to:

* [Config Files](//github.com/storj-thirdparty/connector-cpanel/wiki/#config-files)
* [Run (long version)](//github.com/utropicmedia/storj-cpanl/wiki/#run)
* [Testing](//github.com/storj-thirdparty/connector-cpanel/wiki/#testing)
* [Flow Diagram](//github.com/storj-thirdparty/connector-cpanel/wiki/#flow-diagram)

