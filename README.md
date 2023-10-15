# memcached-mult

Memcached fork with added `mult` command.

[Why?](https://quuxplusone.github.io/blog/2022/01/06/memcached-interview/)

### Installation

```sh
./configure
make
# and optionally:
# make test
```

### Usage

```sh
./memcached
```

This will start a memcached server. You can play with it by running

```sh 
telnet 127.0.0.1 11211
# or
nc -c 127.0.0.1 11211
```

In case you don't have a telnet or netcat installed, there is a small client written in Go inside `client` directory.

```sh
./client/client
# or
cd ./client && go mod tidy && go run main.go
```

