# Rot26

Cli tool made for Rot26, twice as secure compare to Rot13

# Install
1. You will need to install go on your machine: https://go.dev/doc/install
2. Setup GOPATH

Add the following to your shell config
```bash
export PATH=${PATH}:$HOME/go/bin
```
More information: https://go.dev/wiki/GOPATH#gopath-variable

3. Install the binary
```bash
go install github.com/epicseven-cup/Rot26@latest 
```

There could be delays between the Goproxy and GitHub binarys, you can use the direct setup
```bash
GOPROXY=direct go install github.com/epicseven-cup/Rot26@latest
```

# How to use Rot26


```bash
Rot26 -h
```

# Example

```bash
Rot26 -t Helloworld
```

```bash
Rot26 Rot13isBad
```

```bash
echo -n "helloworld" | Rot26
```
