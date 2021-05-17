package tasks

import (
	"bytes"
	"fmt"
	"os/exec"
)

func RunCmd(taskName, cmd string) (string, string, string, string, error) {
	//cmdSplit := strings.Split(cmd, " ")
	//command := exec.Command(cmdSplit[0], cmdSplit[1:]...)
	command := exec.Command("bash", "-c", cmd)
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
	return taskName, cmd, stdout.String(), stderr.String(), err
	//return "", "", err
}
