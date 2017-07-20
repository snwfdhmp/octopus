Octopus
======

Octopus makes creation of decentralized networks super easy.

----

### Connect easily to other nodes

`$ octopus ts 193.28.39.82`


```
Connecting to 193.28.39.82:2048...
Tap 193.28.39.82
Response token : 923bc615
Shake 193.28.39.82
Name snwfdhmp
Connection successful to snwfdhmp
```

### Create networks to group nodes

```
$ octopus net home
Created network home

$ octopus net home add snwfdhmp
Invited snwfdhmp to network 'home'
snwfdhmp joined 'home'
```

### Take control over the network

```
$ octopus sh home
home> @mute apt-get install duck
Command started on 3 nodes (muted)
home> echo $USER
snwfdhmp: Martin
NAS: root
webserver: root
```

### Fusion several networks

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

```
$ octopus ts 192.168.12.56 -w lucas
Asking 'lucas' permission to access 192.168.12.56
'lucas' sent token 3fa06a2d
Shake 192.168.12.56 through 'lucas'
Name computer
Connection successful to computer through lucas

$ octopus ping computer
ping ... OK (25ms)

$ octopus sh computer
computer> ...
```

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
#### Connect to others
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

#### Connect to networks
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


#### Remote shell
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
