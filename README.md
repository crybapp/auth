![Cryb OSS](.github/cryb.png "Cryb OSS Logo")

_**Auth** — Authentication microservice_

[![GitHub contributors](https://img.shields.io/github/contributors/crybapp/auth)](https://github.com/crybapp/auth/graphs/contributors) [![License](https://img.shields.io/github/license/crybapp/auth)](https://github.com/crybapp/auth/blob/master/LICENSE) [![Patreon Donate](https://img.shields.io/badge/donate-Patreon-red.svg)](https://patreon.com/cryb)

## Docs
* [Info](#info)
    * [Status](#status)
* [Codebase](#codebase)
    * [First time setup](#first-time-setup)
        * [Installation](#installation)
    * [Running the app locally](#running-the-app-locally)
        * [Background services](#background-services)
        * [Starting @cryb/auth](#starting-@cryb/auth)
			* [Building](#building)
* [Questions / Issues](#questions--issues)

## Info
`@cryb/auth` is the microservice used authenticate users and other authenticated data types.

### Status
`@cryb/auth` has been actively developed since December 2019. In January 2020 it was rewritten in Go.

## Codebase
The codebase for `@cryb/auth` is written in Go. MongoDB is used as the primary database.

### First time setup
First, clone the `@cryb/auth` repository locally:

```
git clone https://github.com/crybapp/auth.git
```

#### Installation
The following services need to be installed for `@cryb/auth` to function:

* MongoDB

We recommend that you run the following services alongside `@cryb/auth`, but it's not required.
* `@cryb/api`
* `@cryb/atlas`

You also need to install the required dependencies by running `go get -d ./`.

Ensure that `.env.example` is either copied and renamed to `.env`, or is simply renamed to `.env`.

In this file, you'll need some values. Documentation is available in the `.env.example` file.

### Running the app locally

#### Background Services
Make sure that you have installed MongoDB, and that it is running on port 27017.

The command to start MongoDB is `mongod`.

#### Starting @cryb/auth
To run `@cryb/auth`, run `go run main.go`.

#### Building
To build `@cryb/auth`, run `go build -o main .`.

Once built, run `./main` to run the compiled app.

## Questions / Issues

If you have an issues with `@cryb/auth`, please either open a GitHub issue, contact a maintainer or join the [Cryb Discord Server](https://discord.gg/ShTATH4) and ask in #tech-support
