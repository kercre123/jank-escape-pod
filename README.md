# jank-escape-pod

This repo contains a custom Vector escape pod made from [chipper](https://github.com/digital-dream-labs/chipper) and [vector-cloud](https://github.com/digital-dream-labs/vector-cloud).

## Program Descriptions

`chipper` - Chipper is a program used on Digital Dream Lab's servers which takes in a Vector's voice stream, puts it into a speech-to-text processor, and spits out an intent. This is also likely used on the official escape pod. This repo contains an older tree of chipper which does not have the "intent graph" feature (it caused an error upon every new stream), and it now has a working voice processor.

`vector-cloud` - Vector-cloud is the program which runs on Vector himself which uploads the mic stream to a chipper instance. This repo has an older tree of vector-cloud which also does not have the "intent graph" feature and has been modified to allow for a custom CA cert.

## Configuring, Building, Installing

NOTE: This only works with OSKR-unlocked, Dev-unlocked, or Whiskey robots.

`setup.sh` is a prompt-script which guides you through the installation. It can install all necessary packages, get the speech-to-text software, create SSL certificates (with the address/port given), build vector-cloud, build chipper, create a new server config file for Vector, and allow you to copy the new vic-cloud and server config into him.

(This currently only works on Arch or Debian-based Linux)

```
git clone https://github.com/kercre123/jank-escape-pod.git
cd jank-escape-pod
sudo ./setup.sh

# You should be able to just press enter for all of the settings
```

Now install the files created by the script onto the bot:

`sudo ./setup.sh scp <vectorip> <path/to/key>`

Example:

`sudo ./setup.sh scp 192.168.1.150 /home/wire/id_rsa_Vector-R2D2`

If you are on my custom software (WireOS), you do not have to provide an SSH key,

Example:

`sudo ./setup.sh scp 192.168.1.150`

The bot should now be configured to communicate with your server.

To start chipper, run:

```
cd chipper
sudo ./start.sh
```

After all of that, try a voice command.

## Speech Tips

- You have to speak loud and clear for chipper to understand.
- You also need to wait a little longer after Vector's ding (after "hey vector" is said) before saying the rest of the command than you would expect. Maybe a half a second longer. This is an issue and is being worked on.

## Status

OS Support:

- Arch
- Debian/Ubuntu/other APT distros

Architecture Support:

- amd64/x86_64
- arm64/aarch64
- arm32/armv7l

Things It Has Worked On:

- Raspberry Pi 4B+ 4GB RAM with Raspberry Pi OS
	- Very fast, recommended platform
	- Doesn't matter if it is 32-bit or 64-bit
- Raspberry Pi 4B+ 4GB RAM with Manjaro 22.04
- Nintendo Switch with L4T Ubuntu
- Desktop with Ryzen 5 3600, 16 GB RAM with Ubuntu 22.04
- Laptop with mobile i7
- Android Devices
	- Pixel 4, Note 4, Razer Phone, Oculus Quest 2, OnePlus 7 Pro, Moto G6, Pixel 2
	- [Termux](https://github.com/termux/termux-app) proot-distro: Use Ubuntu, make sure to use a port above 1024 and not the default 443.
	- Linux Deploy: Works stock, just make sure to choose the arch that matches your device in settings. Also use a bigger image size, at least 3 GB.

General Notes:

- On a Raspberry Pi 4 4GB, the text is processed very fast, possibly faster than the official escape pod.
- If the architecture is AMD64, the text is processed 4 times so longer phrases get processed fully. Text is only processed once on arm32/arm64 for speed.
- If you get this error when running chipper, you are using a port that is being taken up by a program already: `panic: runtime error: invalid memory address or nil pointer dereference`
	- Run `./setup.sh` with the 5th option to change the port, you will need to push files to the bot again.

Known Issues:

- On Fedora, the STT binary does not start and errors out with "Illegal Instruction (core dumped)"
- Not many intents are currently supported at the moment.
- Intents that need parameters (like "my name is <name>") are not supported currently.
- The audio stream is a little cut off at the beginning.
- Only one robot can be streaming audio to it at a time.

Current Implemented Actions:

- Good robot
- Bad robot
- Change your eye color
	- As well as "Change your eye color to"
	- blue, purple, teal, green, yellow
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
- Move forward
- Turn left
- Turn right
- Roll your cube
- Pop a wheelie
- Fistbump
- Blackjack (say yes/no instead of hit/stand)
- Yes (affirmative)
- No (negative)
- What's my name
- Take a photo
	- As well as "Take a photo of me"

## Credits

[Digital Dream Labs](https://github.com/digital-dream-labs) for saving Vector and for open sourcing chipper which made this possible
