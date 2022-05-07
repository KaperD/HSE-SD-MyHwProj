# Go API Server for myhwproj
[![codecov](https://codecov.io/gh/KaperD/HSE-SD-MyHwProj/branch/02-impl/graph/badge.svg?token=TPI8LNSA9E)](https://codecov.io/gh/KaperD/HSE-SD-MyHwProj)

REST api for MyHwProj

## Overview
This server was generated by the [openapi-generator]
(https://openapi-generator.tech) project.
By using the [OpenAPI-Spec](https://github.com/OAI/OpenAPI-Specification) from a remote server, you can easily generate a server stub.
-

To see how to make this your own, look here:

[README](https://openapi-generator.tech)

- API version: 1.0.0
- Build date: 2022-05-07T15:37:14.927731+05:00[Asia/Yekaterinburg]


### Running the server
To run the server, follow these simple steps:

```
go run main.go
```

To run the server in a docker container
```
docker build --network=host -t myhwproj .
```

Once image is built use
```
docker run --rm -it myhwproj
```
