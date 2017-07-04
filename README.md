# Bodhi

Bodhi is a tool which collects an inventory of system information. It aims to implement some parts of features from It's forked from Kentaro Kuribayashi's [gohai](https://github.com/DataDog/gohai).

## Usage
Require at least Go 1.8.
Running it will dump json formatted output:

```sh
Bodhi
{"cpu":{"cpu_cores":"2","family":"6","mhz":"2600","model":"58","model_name":"Intel(R) Core(TM) i5-3230M CPU @ 2.60GHz","stepping":"9","vendor_id":"GenuineIntel"},"filesystem":[{"kb_size":"244277768","mounted_on":"/","name":"/dev/disk0s2"}],"memory":{"swap_total":"4096.00M","total":"8589934592"},"network":{"ipaddress":"192.168.1.6","ipaddressv6":"fe80::5626:96ff:fed3:5811","macaddress":"54:26:96:d3:58:11"},"platform":{"GOOARCH":"amd64","GOOS":"darwin","goV":"1.2.1","hostname":"new-host.home","kernel_name":"Darwin","kernel_release":"12.5.0","kernel_version":"Darwin Kernel Version 12.5.0: Sun Sep 29 13:33:47 PDT 2013; root:xnu-2050.48.12~1/RELEASE_X86_64","machine":"x86_64","os":"Darwin","processor":"i386","pythonV":"2.7.2"}}
```

## How to build

We use [glide](https://github.com/Masterminds/glide) to pin our dependencies.
```sh
go get github.com/Masterminds/glide
glide install
go build
```

## Build with version info

To build Bodhi with version information, use `make.go`:

```sh
go run make.go
```

It will build Bodhi using the `go build` command, with the version info passed through `-ldflags`.
