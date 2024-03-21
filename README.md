# op-be-logging
The Logging Services for the Backend Layer of OctopOSPrime

[![Build - Test](https://github.com/octoposprime/op-be-logging/actions/workflows/ci.yml/badge.svg)](https://github.com/octoposprime/op-be-logging/actions/workflows/ci.yml)
[![Docker Image Publish](https://github.com/octoposprime/op-be-logging/actions/workflows/cd.yml/badge.svg)](https://github.com/octoposprime/op-be-logging/actions/workflows/cd.yml)

## Pre-Requirements

## Development Environment
You have to see ![github.com/octoposprime/op-be-docs](https://github.com/octoposprime/op-be-docs) before development.

#### .env
```
POSTGRES_USERNAME=op
POSTGRES_PASSWORD=op
JWT_SECRET_KEY=op
REDIS_PASSWORD=op
```

#### Local Run
```
make local-run
```

#### Docker Run
```
TEST=true POSTGRES_USERNAME=op POSTGRES_PASSWORD=op JWT_SECRET_KEY=op REDIS_PASSWORD=op make docker-build
make docker-run 
```

