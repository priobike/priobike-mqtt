# priobike-mqtt

This repository contains a custom [EMQX](https://github.com/emqx/emqx) configuration with access control. It also contains a tiny Go service for authentication at the brokers: https://github.com/priobike/priobike-mqtt/blob/main/authenticator/main.go -- this service loads the `USERNAMES` and `PASSWORDS` env vars (see [`docker-compose.yml`](https://github.com/priobike/priobike-mqtt/blob/main/docker-compose.yml)) and checks all connection attempts to the mqtt cluster against these credentials.

# Quickstart

The easiest way to run a minimal setup containing EMQX and the authenticator is to use the contained `docker-compose.yml`:
```
docker-compose up --build
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
- We use the EMQX broker to make the MQTT more scalable for our services

## Contributing

We highly encourage you to open an issue or a pull request. You can also use our repository freely with the `MIT` license.

Every service runs through testing before it is deployed in our release setup. Read more in our [PrioBike deployment readme](https://github.com/priobike/.github/blob/main/wiki/deployment.md) to understand how specific branches/tags are deployed.

## Anything unclear?

Help us improve this documentation. If you have any problems or unclarities, feel free to open an issue.
