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
	for key, val := range argMap {
		switch key {
		case "--port":
			argMap["port"] = val
		case "--version":
			fmt.Printf("API Version: %s\n", argMap["version"])
		case "--env":
			argMap["environment"] = val
		default:
		}
	}
}
```
