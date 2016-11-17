# web2pic

## Dependence

* [Docker](https://docker.com) 

## Prepare

1. Install docker

2. Build capture image

``` bash
cd capture
docker build . -t "capture"
```

## Run

``` bash
godep run capture.go
```

## Usage

``` bash
## Request
curl -X "GET" "http://localhost:8080/snap?url=http:%2F%2Fwww.example.com"

```

