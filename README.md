# PrioBike MQTT

This repo contains a Dockerfile to use an [EMQX](https://github.com/emqx/emqx) cluster and a Dockerfile to use an authenticator.


# Quickstart

The easiest way to run a minimal setup containing EMQX and the authenticator is to use the contained `docker-compose.yaml`:
```
docker-compose up --build -d
```

# API and CLI

```
mosquitto_pub \
    -h "localhost" \
    -p 1883 \
    -u "backend" -P "nWK8Am3d2Hbupx" \
    -t "test" -m "hello at $(date)"
```

```
mosquitto_sub \
    -h "localhost" \
    -p 1883 \
    -u "user" -P "mqtt@priobike-2022" \
    -t "test"
```

## What else to know
- MQTT Cluster...

## Contributing

We highly encourage you to open an issue or a pull request. You can also use our repository freely with the `MIT` license.

Every service runs through testing before it is deployed in our release setup. Read more in our [PrioBike deployment readme](https://github.com/priobike/.github/blob/main/wiki/deployment.md) to understand how specific branches/tags are deployed.

## Anything unclear?

Help us improve this documentation. If you have any problems or unclarities, feel free to open an issue.
