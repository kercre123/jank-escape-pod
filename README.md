# jank-escape-pod

This repo contains my experiments with [chipper](https://github.com/digital-dream-labs/chipper) and [vector-cloud](https://github.com/digital-dream-labs/vector-cloud)

## What this is

chipper is the program which takes Vector's mic stream (after "hey vector") and puts it into a speech-to-text processor, then spits out an intent.

This repo has a chipper which has a custom certificate, which allows it to work on a custom server.

(The certs included are there on purpose and serve as examples, you will need to put in your own. Make sure to make certs with SANs)

## Building

chipper:

```
cd chipper
make build
```

vector-cloud:
```
cd vector-cloud
make docker-builder
make vic-cloud
```

## Use your own configuration

Build your own certs ([I used this guide](https://gist.github.com/KeithYeh/bb07cadd23645a6a62509b1ec8986bbc))

Replace the certificate located near the top of `./vector-cloud/cloud/main.go` with the public cert you built

Replace the certificates located in `./chipper/source.sh` with the ones you built (also change the port if you need to)

You will also need to change Vector's `/anki/data/assets/cozmo_resources/config/server_config.json` to the server you will be using

## Running

To have Vector use the custom vector-cloud: stop the Anki robot processes, SCP the binary in, then restart the Anki robot processes.

```
# replace vectorip with vector's actual ip address
ssh root@vectorip "systemctl stop anki-robot.target"
scp ./vector-cloud/build/vic-cloud root@vectorip:/anki/bin/
ssh root@vectorip "systemctl start anki-robot.target"
```

To run chipper:

```
cd chipper
source source.sh
./chipper
#if you get an error, something is likely taking up the port chipper is trying to use (default :445)
```

## Implement STT

Download [this](https://github.com/coqui-ai/STT/releases/download/v1.3.0/native_client.tflite.Linux.tar.xz) and extract it so it is accessible at ./stt (and stt binary can be laucnhed via ./stt/stt).

Also download [this](https://coqui.gateway.scarf.sh/english/coqui/v1.0.0-large-vocab/model.tflite) and [this](https://coqui.gateway.scarf.sh/english/coqui/v1.0.0-large-vocab/large_vocabulary.scorer) (warning; big downloads) and put them in the same ./stt/ folder

## Status

Currently, on a fast desktop, the speech-to-text itself is pretty snappy and relatively accurate. Intent matching is barebones right now and only the following are implemented:

go home, start exploring (i recommend saying "deploring", he understands it better), go to sleep, change your eye color, how old are you, good robot, bad robot
