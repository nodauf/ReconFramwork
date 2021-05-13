package orchestrator

import (
	"fmt"
	"reflect"
	"time"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/log"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/nodauf/ReconFramwork/server/parsers"
)

func ExecuteCommands(server *machinery.Server, cmd, parser string) []reflect.Value {

	fmt.Println(cmd)
	task0 := tasks.Signature{
		Name: "runcmd",
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: cmd,
			},
		},
	}
	res, err := server.SendTask(&task0)
	if err != nil {
		log.ERROR.Fatalln(err.Error())
	}

	results, _ := res.Get(2 * time.Millisecond)

	var t parsers.Parser
	if results != nil {
		reflect.ValueOf(t).MethodByName(parser).Call(results)
	} else {
		log.ERROR.Println("Task got an error")
	}

	return results
}
