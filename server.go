package main

import (
	"context"
	// "git.internal.attains.cn/attains-cloud/service-acs/cmd"
	"log"
)

func main() {
	if err := cmd.Run(context.Background()); err != nil {
		log.Panicln(err)
	}
}
