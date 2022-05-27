# jank-escape-pod

This repo contains my experiments with [chipper](https://github.com/digital-dream-labs/chipper) and [vector-cloud](https://github.com/digital-dream-labs/vector-cloud)

## What this is

This is my very-prototype-phase custom escape pod.

chipper is the program which takes Vector's mic stream (after "hey vector") and puts it into a speech-to-text processor, then spits out an intent.

This repo has a chipper which has my voice processing code and is from an older tree. The "intent graph" feature seemed to break it a bit.

vector-cloud is the program running on Vector himself 

This also contains a vector-cloud which is also from an older tree and is modified a little bit to allow for a custom cert. This allows Vector to communicate with your custom chipper.

## Configuring, Building, Installing

`setup.sh` installs all necessary packages, creates the certificates (with the address/port given), puts the cert into the vector-cloud source, builds vector-cloud, builds chipper, creates a new server config file for Vector, and allows you to copy the new vic-cloud and server config into him.

(This currently only works on Debian-based Linux)

```
git clone https://github.com/kercre123/jank-escape-pod.git
cd jank-escape-pod
sudo ./setup.sh
#Press enter once you run it
```

To install the files created by the script onto the bot, run:

`sudo ./setup.sh scp <vectorip> <path/to/key>`

Example:

`sudo ./setup.sh scp 192.168.1.150 /home/wire/id_rsa_Vector-R2D2`

The bot should now be configured to communicate with your server.

To start chipper, run:

```
cd chipper
./start.sh
```

After all of that, try a voice command.


## Status

Right now; This has the same STT software and STT model/scorer as escape pod, and I think processes text in a very similar way, but the official escape pod is currently much better at intent matching.

Currently, on a fast desktop, the speech-to-text itself is pretty snappy and accurate. But, you have to speak loud and clearly.

Intent matching is barebones right now and only the following are implemented:

go home, start exploring (i recommend saying "deploring", he understands it better), go to sleep, change your eye color, how old are you, good robot, bad robot

## Credits

[Digital Dream Labs](https://github.com/digital-dream-labs) for saving Vector and for open sourcing chipper which made this possible
