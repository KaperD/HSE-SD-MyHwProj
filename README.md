# MyHwProj
[![codecov](https://codecov.io/gh/KaperD/HSE-SD-MyHwProj/branch/02-impl/graph/badge.svg?token=TPI8LNSA9E)](https://codecov.io/gh/KaperD/HSE-SD-MyHwProj)

### Running the server
Firstly, you need Go 1.18 to be installed

Secondly, you need Postgres instance. You can create it with Docker

And the last thing, we use RabbitMQ. The last command will do in inside an other Docker image.

```shell
docker run --name myhwproj-db -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -p 5432:5432 -d postgres
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
```

If you specified other parameters, you should change them in [db_config.json](./db_config.json)

After that you can run server
```shell
go run main.go
```

### Running tests
```shell
go install gotest.tools/gotestsum@latest
gotestsum --format testname ./test/...
```
