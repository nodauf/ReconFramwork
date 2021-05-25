package customTasks

import "fmt"

func DomainFromCert(taskName, cmd string) (string, string, string, string, error) {
	fmt.Println("Custom task")
	return taskName, cmd, "stdout.String()", "stderr.String()", nil
}
