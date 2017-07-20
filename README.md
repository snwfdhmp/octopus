Octopus
======

Octopus makes creation of decentralized networks super easy.

----

## Installation

TODO: Describe the installation process

## Getting started

#### Verify installation
 Verify that octopus is installed by running `octopus --version` 
#### Create your node
```
$ octopus node create
Name: snwfdhmp
Port (2048): <Enter>
```
#### Connect easily to other nodes
You can find a list of open nodes [here](#).

Once you've decided which node to connect to, run `octopus ts <ip> <port (default: 2048)>`
```
$octopus ts 192.168.1.42
Connecting to 192.168.1.42:2048...
Tap 192.168.1.42
Response token : 8d5faab2
Shake 192.168.1.42
Name media-center
Connection successful to media-center
```

#### Start networking
Create your networks and add your nodes
```
$ octopus net home
Created network home
$ octopus net home tap snwfdhmp
Invited snwfdhmp to network 'home'
snwfdhmp joined 'home'
```

Or let you be tapped to existing networks
```
$ octopus logs
...
> tap from 14.233.27.9 to dev-team
...

$ octopus nets
NAME      STATUS
dev-team  pending
friends   joined
local     joined

$ octopus join dev-team
Accepted invitation to dev-team
Joining dev-team
Connect to (5/9) nodes...
```


#### DSNet Remote Shell

You can start a shell that will run on remote nodes ...

```
$ octopus sh media-center
media-center> echo $HOME
/home/media
media-center> reboot

$ octopus sh media-center "echo $HOME"
/home/media
```

... or on an entire network
```
$ octopus sh local
local> echo $HOME
media-center: /home/media
snwfdhmp: /Users/Martin
```

### Fusion several networks

You can fusion several networks to make a bigger one.
This will invite every node of each network to the new network created.

```
$ octopus net game -f home friends
Fusioning 'home' and 'friends' into 'game'
home/snwfdhmp joined
friends/lucas joined
home/NAS joined
home/webserver joined
friends/landry joined

5 nodes in 'game'
```

### Access local machines remotely

Situation: Lucas has 2 computers, each one running octopus.

Its first computer (named "lucas") is connected to the internet.

Its second computer (named "second") is just connected to LAN.

You're not connected to Lucas' LAN but you have octopus running connected to lucas' Node.

```
$ octopus ts second -w lucas
Asking 'lucas' permission to access 'second'
'lucas' sent token 3fa06a2d
Shake 192.168.12.56 through 'lucas'
Name second
Connection successful to second through lucas

$ octopus ping computer
ping ... OK (25ms)

$ octopus sh computer
computer> ...
``

## [Duck](https://github.com/snwfdhmp/duck) support

We're actually working on implementing [duck](https://github.com/snwfdhmp/duck) support.

This will permit to run lings on several machines.

#### Example combination of [duck](https://github.com/snwfdhmp/duck) and [octopus](https://github.com/snwfdhmp/octopus)

TODO: write usage examples (duck+octopus)

## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D

## History

TODO: Write history

## Credits

TODO: Write credits

## License

TODO: Write license
