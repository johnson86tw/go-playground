package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
)

type commandLine struct{}

func main() {
	defer os.Exit(0)
	cli := commandLine{}
	cli.run()
}

func (cli *commandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" greet")
}

func (cli *commandLine) run() {
	cli.validateArgs()

	// new FlagSet
	sayHiCmd := flag.NewFlagSet("greet", flag.ExitOnError)

	// create flag
	name := sayHiCmd.String("name", "", "the man you want to say hi to.")

	// parse args and its flag
	switch os.Args[1] {
	case "greet":
		err := sayHiCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	default:
		cli.printUsage()
		runtime.Goexit()
	}

	// execute program
	if sayHiCmd.Parsed() {
		if *name == "" {
			sayHiCmd.Usage()
			runtime.Goexit()
		}

		cli.sayHi(*name)
	}
}

func (cli *commandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *commandLine) sayHi(name string) {
	fmt.Printf("Heyo, %s, what's up!\n", name)
}
