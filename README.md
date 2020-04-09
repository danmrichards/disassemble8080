# Disassemble 8080
A disassembler for the Intel 8080 written in pure Go.

Takes an 8080 ROM file and prints out the resulting assembly code for the Intel
8080.

## Installation
If you are a golang developer you can install the binary like so:
```bash
$ go get -u github.com/danmrichards/disassemble8080/cmd/disassemble8080/...
```

Alternatively you can clone this repo and build from source:

```bash
$ make
```

## Usage
```bash
Usage of disassemble8080:
  -rom string
        Path to 8080 ROM file
```