# godump

godump aims to be a simple executable binary as tcpdump.

## Build

### Mac

```
go build
```

### Linux (for those who had not install `libpcap`)

```
$ go get github.com/ykyuen/xgo
$ xgo -ldflags='-linkmode external -extldflags "-static -s -w"' --deps=http://www.tcpdump.org/release/libpcap-1.9.0.tar.gz --depsargs=--with-pcap=linux --targets=linux/amd64 github.com/iamken1204/godump
```
