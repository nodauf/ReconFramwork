package tasks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/nodauf/ReconFramwork/server/server/models"
)

func RunCmd(taskName, cmd string) ([]byte, error) {
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
	var output models.Output
	output.TaskName = taskName
	output.Cmd = cmd
	output.Stdout = stdout.String()
	output.Stderr = stderr.String()
	if err != nil {
		output.Error = err.Error()
	}
	//test := reflect.ValueOf("ParserNmap").Interface().(parsers.ParserNmap)
	//var t parsers.Parser
	//reflect.ValueOf(t).MethodByName("ParseNmap").Call(nil)
	//test.Parse(string(out))
	outputBytes, _ := json.Marshal(output)
	return outputBytes, nil
	//return "", "", err
}
