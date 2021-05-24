package orchestrator

import (
	"reflect"
	"sync"
	"time"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/log"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/nodauf/ReconFramwork/server/db"
	"github.com/nodauf/ReconFramwork/server/parsers"
)

func ConsumeEndedTasks(server *machinery.Server, wg *sync.WaitGroup) {
	defer wg.Done()
	jobs := db.GetNonProcessedTasks()

	var t parsers.Parser
	for _, job := range jobs {
		results, _ := server.GetBackend().GetState(job.TaskUUID)
		if results.IsSuccess() {
			t.Job = job
			reflectResults, _ := tasks.ReflectTaskResults(results.Results)
			reflect.ValueOf(t).MethodByName(job.Parser).Call(reflectResults)
			log.INFO.Println("Done")

			db.RemoveJob(&job)
		} else {
			log.INFO.Println("Not success")
		}
	}
}

func executeCommands(server *machinery.Server, host, cmd, parser, taskName string) {
	//fmt.Println(cmd)
	task0 := tasks.Signature{
		Name: "runcmd",
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: taskName,
			},
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
	job, err := db.AddJob(host, parser, res.GetState().TaskUUID)
	if err != nil {
		log.ERROR.Println(err)
		return
	}
	results, _ := res.Get(2 * time.Millisecond)
	var t parsers.Parser
	t.Job = job
	//fmt.Println(res.Signature)
	if results != nil {
		reflect.ValueOf(t).MethodByName(parser).Call(results)
	} else {
		log.ERROR.Println("Task got an error")
	}
	db.RemoveJob(&job)

}
