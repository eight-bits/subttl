package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

// HELP
//====================
func help() {
	fmt.Println("The program allows you to get and set the value of ttl.")
	fmt.Println("Example:")
	fmt.Println("To get the current ttl value, enter 'sudo subttl -key=get'")
	fmt.Println("To change the ttl value, enter the command 'sudo subttl -key=set -vol=64'")
}

// SET TTL
//====================
func set_ttl(ttl int) {
	cmd := exec.Command("sysctl", "-w", "net.inet.ip.ttl="+fmt.Sprintf("%d", ttl))
	var buf bytes.Buffer
	cmd.Stdout = &buf
	err := cmd.Start()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	cmd.Wait()
	fmt.Printf(buf.String())
}

// GET TTL
//====================
func get_ttl() {
	cmd := exec.Command("sysctl", "-n", "net.inet.ip.ttl")
	var buf bytes.Buffer
	cmd.Stdout = &buf
	err := cmd.Start()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	cmd.Wait()
	fmt.Printf(buf.String())
}

// FUNCTION MAIN
//===================
func main() {
	var keyArg string
	var volArg int
	flag.StringVar(&keyArg, "key", "zero", "The parameter is set and get. Example -key=set")
	flag.IntVar(&volArg, "vol", 64, "Еhe parameter has values ​​ranging from 1 to 255. Example -vol=64")
	flag.Parse()
	if keyArg == "zero" {
		help()
		os.Exit(0)
	} else if keyArg == "get" {
		get_ttl()
		os.Exit(0)
	} else if keyArg == "set" && volArg > 0 && volArg <= 255 {
		set_ttl(volArg)
		os.Exit(0)
	} else {
		fmt.Println("Incorrect keys or their values")
		os.Exit(0)
	}

}
