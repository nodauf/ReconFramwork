package tasks

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func HelloWorld(args string, args2 string) error {
	fmt.Println("Hello world " + args + args2)
	return nil
}

func RunCmd(cmd string) (string, string, error) {
	cmdSplit := strings.Split(cmd, " ")
	command := exec.Command(cmdSplit[0], cmdSplit[1:]...)
	//var stdout io.Writer
	var stderr bytes.Buffer
	var stdout bytes.Buffer
	command.Stdout = &stdout
	command.Stderr = &stderr
	err := command.Run()
	//out, err := command.CombinedOutput()

	fmt.Println(stdout.String())
	fmt.Println(stderr.String())
	fmt.Println(command.ProcessState.ExitCode())
	//test := reflect.ValueOf("ParserNmap").Interface().(parsers.ParserNmap)
	//var t parsers.Parser
	//reflect.ValueOf(t).MethodByName("ParseNmap").Call(nil)
	//test.Parse(string(out))
	return stdout.String(), stderr.String(), err
	//return "", "", err
}
