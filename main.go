package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("hello")

	goback , _ := exec.Command("cd ..").Output()
	fmt.Println(string(goback))

	mkdir , _ := exec.Command("mkdir", "username").Output()
	fmt.Println(string(mkdir))

	checkls , _ := exec.Command("ls").Output()
	fmt.Println(string(checkls))
}

