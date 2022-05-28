# jank-escape-pod

This repo contains a custom Vector escape pod made from [chipper](https://github.com/digital-dream-labs/chipper) and [vector-cloud](https://github.com/digital-dream-labs/vector-cloud).

## Program Descriptions

`chipper` - Chipper is a program used on Digital Dream Lab's servers which takes in a Vector's voice stream, puts it into a speech-to-text processor, and spits out an intent. This is also likely used on the official escape pod. This repo contains an older tree of chipper which does not have the "intent graph" feature (it caused an error upon every new stream), and it now has a working voice processor.

`vector-cloud` - Vector-cloud is the program which runs on Vector himself which uploads the mic stream to a chipper instance. This repo has an older tree of vector-cloud which also does not have the "intent graph" feature and has been modified to allow for a custom CA cert.

## Configuring, Building, Installing

NOTE: This only works with OSKR-unlocked, Dev-unlocked, or Whiskey robots.

`setup.sh` is a prompt-script which guides you through the installation. It can install all necessary packages, get the speech-to-text software, create SSL certificates (with the address/port given), put the public cert into the vector-cloud source, build vector-cloud, build chipper, create a new server config file for Vector, and allow you to copy the new vic-cloud and server config into him.

(This currently only works on Debian-based Linux)

```
git clone https://github.com/kercre123/jank-escape-pod.git
cd jank-escape-pod
sudo ./setup.sh

# You should be able to just press enter for all of the settings, 
# except the part where you must enter an IP address or domain
```

To install the files created by the script onto the bot, run:

`sudo ./setup.sh scp <vectorip> <path/to/key>`

Example:

`sudo ./setup.sh scp 192.168.1.150 /home/wire/id_rsa_Vector-R2D2`

The bot should now be configured to communicate with your server.

To start chipper, run:

```
cd chipper
sudo ./start.sh
```

After all of that, try a voice command.

## Speech Tips

You have to speak loud and clear for chipper to understand.

You also need to wait a little longer after Vector's ding (after "hey vector" is said) before saying the rest of the command than you would expect. Maybe a half a second longer. This is an issue and is being worked on.

## Status

This has the same STT software and STT model/scorer as escape pod, and I think processes text in a similar way, but the official escape pod is currently much better at intent matching and UX.

Debian-based Linux amd64/aarch64/armv7l is supported (Ubuntu, Linux Mint, Debian, Raspberry Pi OS, anything with APT should work). 

Fedora was attempted but I couldn't get the stt binary to run without lib issues.

If the server is ARM based, it will only process the text once instead of four times like amd64 will. This makes them about the same speed, but aarch64 may not do well with longer phrases.

A Raspberry Pi 4 running this is actually pretty fast, maybe even faster than the official escape pod.

Currently, on a desktop with a Ryzen 5 3600, the speech-to-text itself is about as fast (maybe even faster) than the actual chipper server prod bots use. 

Speech tips: You have to speak loud and clear and you have to wait a little bit longer than you usually would after the "Hey Vector" before you start talking. A quarter of a second longer maybe.

Here is the current list of implemented actions:


- Good Robot
- Bad Robot
- Change your eye color
- How old are you
- Start exploring ("deploring" works better)
- Go home (or "go to your charger")
- Go to sleep
- Good morning
- Good night
- What time is it
- Goodbye
- Happy new year
- Happy holidays
- Hello
- Sign in alexa
- Sign out alexa
- I love you

## Credits

[Digital Dream Labs](https://github.com/digital-dream-labs) for saving Vector and for open sourcing chipper which made this possible
