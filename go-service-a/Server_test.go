package main
import (
	"fmt"
	"os/exec"
	"testing"
)

func HelloWorldTest(t *testing.T) {
	cmd := exec.Command("./Server.go")
	if err := cmd.Run(); err != nil{
		fmt.Println(err)
	}
	if true {
		t.Error("Error in Server ")
	}
}