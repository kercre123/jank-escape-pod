# personal-chipper-cloud

A collection of experimental code involving [chipper](https://github.com/digital-dream-labs/chipper) and [vector-cloud](https://github.com/digital-dream-labs/vector-cloud)

There are certificates in here, they are just self-signed ones and serve as examples for how someone else can implement their own.

There is an STT implementation. Download [this](https://github.com/coqui-ai/STT/releases/download/v1.3.0/native_client.tflite.Linux.tar.xz) and extract it so it is accessible at GITROOT/stt (stt binary can be laucnhed via ./stt/stt).

For the STT implementation, also download [this](https://coqui.gateway.scarf.sh/english/coqui/v1.0.0-huge-vocab/huge-vocabulary.scorer) and [this](https://coqui.gateway.scarf.sh/english/coqui/v1.0.0-huge-vocab/model.tflite) (warning, big downloads) and put them in the same ./stt/ folder

It is very janky and I would recommend rewriting the STT implementation if you want to use this yourself.

You need to put your certs in vector-cloud/cloud/main.go and chipper/certs

To launch, run `cd chipper` then `source ./source.sh` then `./chipper`

## Differences from normal vector-cloud and chipper

These come from older trees before intent_graph was added. That seemed to break chipper a bit, at least in the noop voice_processor.

Custom cert now gets added in cloud/main.go rather than in the jdocs client (may lead to memory leak, will look into later)

I added an STT implementation, currently includes "good robot", "bad robot", "change your eye color", and "how old are you"
