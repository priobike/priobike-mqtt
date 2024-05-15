# Quickstart

```
docker-compose up --build -d
```

```
mosquitto_pub \
    -h "localhost" \
    -p 1883 \
    -u "backend" -P "nWK8Am3d2Hbupx" \
    -t "test" -m "hello at $(date)"
```

```
mosquitto_sub -h "localhost" -p 1883 -t "test"
```