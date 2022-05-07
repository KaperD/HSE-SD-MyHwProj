FROM golang:1.18 AS build
WORKDIR /go/src
COPY internal ./internal
COPY go.mod .
COPY go.sum .
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o myhwproj .

FROM scratch AS runtime
COPY --from=build /go/src/myhwproj ./
EXPOSE 8080/tcp
ENTRYPOINT ["./myhwproj"]
