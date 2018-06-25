package goargh

import (
	"os"
	"strings"
)

type ArgsMap map[string]string

func GetArgs() map[string]string {
	ArgsMap := make(map[string]string)
	for i := 1; i < len(os.Args); i++ {
		if strings.Contains(os.Args[i], "--") {
			if _, ok := ArgsMap[os.Args[i]]; !ok {
				ArgsMap[os.Args[i]] = ""
				if i < len(os.Args)-1 {
					ArgsMap[os.Args[i]] = os.Args[i+1]
				}
			}
		}
	}
	return ArgsMap
}
