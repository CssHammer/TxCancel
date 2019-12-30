# Test task by Dmitriy Komisarenko

## Docker

You can use Docker to test the app. Check README.md in docker/ directory

## Project structure

+ cmd - entry points of executables (you can use pre-built ones)
+ config - config validation and parsing
+ dbqueries - interfaces to perform DB queries
+ http - handler for HTTP API
+ middleware - HTTP middlewares
+ migrations - SQL migrations
+ models - basic app models

## Postman collection

[Link](https://www.getpostman.com/collections/30b985805f20c3962643)

## Migrator

Migrator applies migrations to your DB. Migrations are stored in **/migrations** directory.
They can be bundled into executable using [packr2](https://github.com/gobuffalo/packr/tree/master/v2)

If you decide to bundle them:

1. go get -u github.com/gobuffalo/packr/v2/packr2
2. cd /cmd/migrator
3. packr2 build

You can remove generated files with ```packr2 clean```

If you don`t bundle them:

1. cd /cmd/migrator
2. go build

Now you have your executable **migrator** with migrations inside. 
Setup your connection to DB in **config_api.yaml**.
Launch migrator passing config file:

```./migrator -c ../../config_api.yaml up```

If you haven`t bundled your migrations make sure you execute the command from **/cmd/migrator**!

Migrations applied.

## API

HTTP API server

How to build :
1. cd /cmd/api
2. go build

Launch server passing config file:

```./api -c ../../config_api.yaml serve```

## Cron

Cron job service checks DB and cancels transactions

How to build:
1. cd /cmd/cron
2. go build

Launch cron passing config file:

```./cron -c ../../config_cron.yaml serve```