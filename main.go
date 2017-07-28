package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	runCommand := flag.NewFlagSet("run", flag.ExitOnError)
	psCommand := flag.NewFlagSet("ps", flag.ExitOnError)

	runNamePtr := runCommand.String("name", "", "Container name")
	runDaemonPtr := runCommand.Bool("d", false, "Forground or Backgroud")

	psAllPtr := psCommand.Bool("a", false, "List all container")
	psQuietPtr := psCommand.Bool("q", false, "Only list container ID")

	if len(os.Args) < 2 {
		fmt.Println("run or ps command is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "run":
		runCommand.Parse(os.Args[2:])
	case "ps":
		psCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if runCommand.Parsed() {
		fmt.Printf("runNamePtr is %s, runDaemonPtr is %t\n", *runNamePtr, *runDaemonPtr)
	}
	if psCommand.Parsed() {
		fmt.Printf("psAllPtr is %t, psQuietPtr is %t\n", *psAllPtr, *psQuietPtr)
	}

}
