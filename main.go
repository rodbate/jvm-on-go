package main

import (
	"fmt"
	"os"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("Version 1.0")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJvm(cmd)
	}
}

func startJvm(cmd *Cmd) {
	fmt.Printf("java_home: %s\nclasspath: %s\nclass: %s\nargs: %v\n",
		os.Getenv("JAVA_HOME"), cmd.cpOption, cmd.class, cmd.args)
	newJvm(cmd).Start()
}
