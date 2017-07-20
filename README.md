# Octopus

Octopus makes creation of decentralized networks super easy.

	- Connect easily to other nodes

```
$ octopus ts 193.28.39.82
Connecting to 193.28.39.82 ...
Tap 193.28.39.82
Response token : 923bc615
Shake 193.28.39.82
Name snwfdhmp
Connection successful to snwfdhmp
```

Create networks to group nodes

```
$ octopus net home
Created network home
$ octopus net home add snwfdhmp
Invited snwfdhmp to network 'home'
snwfdhmp joined 'home'
```

Take control over the network

```
$ octopus sh home
home> @mute apt-get install duck
Command started on 3 nodes (muted)
home> echo $USER
snwfdhmp: Martin
NAS: root
webserver: root
```

Fusion several networks

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

Access local machines remotely

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

## Usage



`status <network>` to get network status

```
$ octopus status game
NAME		LAST PING	IP
snwfdhmp	3s ago		local
lucas		6s ago		214.29.75.38
landry		13s ago		7.38.99.54
webserver	30s ago		192.168.1.38
NAS			5m ago		192.168.1.32

total: 5 nodes
```

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