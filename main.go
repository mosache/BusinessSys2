package main

import "BusinessSys/cmd"

func main() {

	if err := cmd.RootCmd.Execute(); err != nil {
		panic("run err")
	}
}
