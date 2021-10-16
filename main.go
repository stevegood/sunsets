package main

import (
	"github.com/stevegood/sunsets/cmd"
	"log"
)

func init() {

	//setting flags for log level
	log.SetFlags(3)
}

func main() {
	cmd.Execute()
}
