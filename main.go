package main

import (
	"github.com/ryanwalls/cluster-demo/commands"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	commands.Execute()
}
