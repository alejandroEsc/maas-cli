# MAAS-CLI
Client tool developed against the golang maas api library [https://github.com/alejandroEsc/golang-maas-client](https://github.com/alejandroEsc/golang-maas-client).
That project in turn was based out of [https://github.com/juju/gomaasapi](https://github.com/juju/gomaasapi), which is strongly encouraged.

The point of this project is to build a CLI tool as a means to access MAAS.

## Building
First we need to build the vendor directory

```
$ dep ensure
```

To build the projects listed above you run 

```
$ make compile
```

which should build the binary:

```
./bin/mass-cli
```

The client tool being the CLI and the other an example of creating a maas client.

## maas-cli
The CLI is a work in progess, you can access the help menu by typing
```
$ ./bin/maas-cli help
```

It is recommended that you export variables associated with your maas deployment and usages,
e.g.,

```
export MAAS_URL=http://192.168.4.2:5240/MAAS/
export MAAS_VERSION=2.0
export MAAS_APIKEY=G5YtjXQgjuVu9Yz4FG:NKq4KqHyfSm45fUZ5k:5xt9yatzKnYkMv278fKyzwH7h7n6X4mf
```
### Examples
List machines available
```
$ ./bin/maas-cli list-machines
	 0 		 fpfnhk 	nuc2-1 	ubuntu     ga-16.04 	 on 	 Deployed 	
	 1 		 t67tnf 	nuc2-2 	ubuntu     ga-16.04 	 on 	 Deployed 	
	 2 		 dddcpt 	nuc2-3 	ubuntu     ga-16.04 	 on 	 Deployed 	
	 3 		 rxb4tr 	nuc2-4 	ubuntu     ga-16.04 	 on 	 Deployed 	
```

Get individual status
```
$ ./bin/maas-cli machine status fpfnhk t67tnf
	 fpfnhk 	on      Deployed 	
	 t67tnf 	on      Deployed 	
```

## MAAS Client
Sample client code that creates a client that consumes the maas api. 


## Developing


# References

- [1] [MAAS API Documenttion](https://docs.maas.io/2.4/en/api)
- [2] [https://github.com/juju/gomaasapi](https://github.com/juju/gomaasapi)
