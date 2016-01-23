# addup

Adds integers in columns from stdin (like sort, but for adding things up).

## Install

```
go install github.com/brownjohnf/goutils/addup
```

## TODO

* Support flags:
  * -f, output file
  * -i, input file
* Support floats
* Support M, G, units (etc.)

## Usage

From stdout:

```
df | addup -k 3 -v
```

Local file:

```
addup -k 2 < some_file.txt
```

## Building

```
$ go get github.com/brownjohnf/goutils/addup
$ cd $GOPATH
$ go build # or go install
```

