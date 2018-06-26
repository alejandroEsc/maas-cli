#  MAAS-Client-Sample

Project consists of two parts:
- MAAS Client
- MAAS CLI tool

One goal of this project is to get usable maas-client library code example so one
can learn to use the api tools provieded by the [https://github.com/juju/gomaasapi](https://github.com/juju/gomaasapi) project.

Note: That make file cannot be trusted as code is still work in progress.

## Building tools


## MAAS CLI
The CLI is a work in progess, you can access the help menu by typing
```
$ ./bin/maas-cli help
```

It is recommended that you export variables associated with your maas deployment and usages,
e.g.,

```
export MAAS_CLI_URL=http://192.168.4.2:5240/MAAS/
export MAAS_CLI_API_VERSION=2.0
export MAAS_CLI_API_KEY=G5YtjXQgjuVu9Yz4FG:NKq4KqHyfSm45fUZ5k:5xt9yatzKnYkMv278fKyzwH7h7n6X4mf
```
### Examples
List machines available
```
$ ./bin/maas-cli list-machines
2018-06-26 18:32:57 INFO pkg.maas maas_client.go:27 Fetch list of machines...
--- ON ---
|	 0 	|	 fpfnhk 	|	 nuc2-1 	|	 ubuntu:ga-16.04 	|	 on 	|	 Deployed 	|
|	 1 	|	 t67tnf 	|	 nuc2-2 	|	 ubuntu:ga-16.04 	|	 on 	|	 Deployed 	|
|	 2 	|	 dddcpt 	|	 nuc2-3 	|	 ubuntu:ga-16.04 	|	 on 	|	 Deployed 	|
|	 3 	|	 rxb4tr 	|	 nuc2-4 	|	 ubuntu:ga-16.04 	|	 on 	|	 Deployed 	|
```

Get individual status
```
$ ./bin/maas-cli machine status fpfnhk t67tnf
|	 fpfnhk 	|	 on 	|	 Deployed 	|
|	 t67tnf 	|	 on 	|	 Deployed 	|
```

## MAAS Client

## Developing