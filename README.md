# Prometheus Ribbon Exporter

Exposes metrics for your Ribbon SBC from the RESTCONF API, to a Prometheus
compatible endpoint.  Based on the [sonus-metrics-exporter](https://github.com/teliax/sonus-metrics-exporter) code.

## Configuration

This exporter is setup to take input from environment variables:

### Required
* `API_URLS` Space-separated list of URLs for the Ribbon RESTCONF API to be tried in order.  Should appear as `https://{ip1}/api https://{ip2}/restconf/data/`
* `API_USER` The username to use when logging in to the Ribbon RESTCONF API.
* `API_PASSWORD` The password to use when authenticating to the Ribbon RESTCONF.

### Optional
* `API_ADDRESSCONTEXTS` Space-separated list of addressContexts to iterate over, defaults to `default`
* `API_TIMEOUT` HTTP client timeout in seconds, defaults to `10`
* `LISTEN_PORT` The port you wish to run the container on, the Dockerfile defaults this to `9172`
* `METRICS_PATH` the metrics URL path you wish to use, defaults to `/metrics`
* `LOG_LEVEL` The level of logging the exporter will run with, defaults to `debug`


## Install and deploy


Build a docker image:
```
docker build -t <image-name> .
docker run -d --restart=always -p 9172:9172  -e API_USER="username" -e API_PASSWORD="password" <image-name>
```

## Docker compose

```
ribbon-metrics-exporter:
    tty: true
    stdin_open: true
    expose:
      - 9172
    ports:
      - 9172:9172
    image: <image-name>:latest
    environment:
      - API_USER=username
      - API_PASSWORD=password

```

## Metrics

Metrics will be made available on port 9172 by default
An example of these metrics can be found in the `METRICS.md` markdown file in the root of this repository

## Metadata
