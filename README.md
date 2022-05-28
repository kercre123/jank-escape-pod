# jank-escape-pod

This repo contains my experiments with [chipper](https://github.com/digital-dream-labs/chipper) and [vector-cloud](https://github.com/digital-dream-labs/vector-cloud)

## What this is

This is my very-prototype-phase custom escape pod.

chipper is the program which takes Vector's mic stream (after "hey vector") and puts it into a speech-to-text processor, then spits out an intent.

This repo has a chipper which has my voice processing code and is from an older tree. The "intent graph" feature seemed to break it a bit.

vector-cloud is the program running on Vector himself which takes the mic stream and pushes it to a chipper instance.

This also contains a vector-cloud which is also from an older tree and is modified a little bit to allow for a custom cert. This allows Vector to communicate with your custom chipper.

## Configuring, Building, Installing

`setup.sh` installs all necessary packages, gets the speech-to-text software, creates SSL certificates (with the address/port given), puts the public cert into the vector-cloud source, builds vector-cloud, builds chipper, creates a new server config file for Vector, and allows you to copy the new vic-cloud and server config into him.

(This currently only works on Debian-based Linux)

```
git clone https://github.com/kercre123/jank-escape-pod.git
cd jank-escape-pod
sudo ./setup.sh

#You should be able to just press enter for all of the settings, except the part where you enter an IP address or domain
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


## Status

Right now; This has the same STT software and STT model/scorer as escape pod, and I think processes text in a very similar way, but the official escape pod is currently much better at intent matching.

Only Debian-based Linux amd64 and aarch64 is supported (Ubuntu, Linux Mint, Debian, Raspberry Pi OS, anything with APT should work). Fedora was attempted but I couldn't get the stt binary to run without lib issues.

I have tested it with an RPi4 and a Nintendo Switch with L4T Ubuntu and it works fine, but it's slow.

Currently, on a fast desktop, the speech-to-text itself is pretty snappy and accurate. But, you have to speak loud and clearly.

Here is the current list of implemented actions:

Good Robot
Bad Robot
Change your eye color
How old are you
Start exploring ("deploring" works better)
Go home (or "go to your charger")
Go to sleep
Good morning
Good night
What time is it
Goodbye
Happy new year
Happy holidays
Hello
Sign in alexa
Sign out alexa

## Credits

[Digital Dream Labs](https://github.com/digital-dream-labs) for saving Vector and for open sourcing chipper which made this possible
