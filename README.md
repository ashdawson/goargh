# goargh

goargh loads any arguments passed to your go application via the command line

### Installation
```sh
$ go get github.com/ashdawson/goargh
```

### Usage
Example usage

```sh
package main

import (
	"github.com/ashdawson/goargh"
)

var args *goargh.ArgsMap

func init() {
	args := goargh.GetArgs()
}

func main() {
	readArgs(args)
}

func readArgs(argMap goargh.ArgsMap) {
	if _, ok := argsMap["--version"]; ok {
		fmt.Printf("API Version: %s\n", argsMap["version"])
	}
	if val, ok := argsMap["--port"]; ok {
		argsMap["port"] = val
	}
	if val, ok := argsMap["--env"]; ok {
		argsMap["environment"] = val
	}
}
```
