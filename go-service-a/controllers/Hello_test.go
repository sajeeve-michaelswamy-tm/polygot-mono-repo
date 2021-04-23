package controllers

import (
	"fmt"
	"os/exec"
	"testing"
)

func HelloWorldTest(t *testing.T) {
	cmd := exec.Command("./Hello.go")
	if err := cmd.Run(); err != nil{
		fmt.Println(err)
	}
	if true {
		t.Error("Error in Hello ")
	}
}