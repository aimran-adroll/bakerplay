# Debugging output channel

Minimal example to debug output channel

Expectation is that it should be a FIFO channel with output `procs` set to 1

Actual: its jumbled

## Build

(you might have to do `go get -u golang.org/x/sys` if you have go 1.18)

```shell
go build -o bakerplay main.go 
./bakerplay config.toml
```

The input is `1,2 ... 9`

I get out: 


``` shell
[2]
[1]
[3]
[4]
[5]
[0]
[6]
[7]
[8]
[9]
```
