package orchestrator

import (
	"reflect"
	"sync"
	"time"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/log"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/nodauf/ReconFramwork/server/server/db"
	parsersCustoms "github.com/nodauf/ReconFramwork/server/server/parsers/customs"
	parsersTools "github.com/nodauf/ReconFramwork/server/server/parsers/tools"
)

func ConsumeEndedTasks(server *machinery.Server, wg *sync.WaitGroup) {
	defer wg.Done()
	jobs := db.GetNonProcessedTasks()

	for _, job := range jobs {
		results, _ := server.GetBackend().GetState(job.TaskUUID)
		if results.IsSuccess() {
			reflectResults, _ := tasks.ReflectTaskResults(results.Results)

			// Two seperate package if the parser is for custom tasks or for runCmd task
			if results.TaskName == "runCmd" {
				var t parsersTools.Parser
				t.Job = job
				reflect.ValueOf(t).MethodByName(job.Parser).Call(reflectResults)
			} else {
				var t parsersCustoms.Parser
				t.Job = job
				reflect.ValueOf(t).MethodByName(job.Parser).Call(reflectResults)
			}
			log.INFO.Println("Done")

			db.ValidateJob(&job, reflectResults)
		} else {
			log.INFO.Println("Not success")
		}
	}
}

func executeCommands(server *machinery.Server, host, cmd, parser, taskName, machineryTask string) {
	//fmt.Println(cmd)

	task0 := tasks.Signature{
		Name: machineryTask,
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
	job, err := db.AddJob(host, parser, machineryTask, cmd)

	if err != nil {
		log.ERROR.Println(err)
		return
	}
	log.DEBUG.Println("Send " + taskName + " and " + cmd + " to worker")
	res, err := server.SendTask(&task0)
	if err != nil {
		log.ERROR.Fatalln(err.Error())
	}
	db.UpdateJob(&job, res.GetState().TaskUUID)

	results, _ := res.Get(2 * time.Millisecond)
	//fmt.Println(res.Signature)
	if results != nil {
		// Two seperate package if the parser is for custom tasks or for runCmd task
		if machineryTask == "runCmd" {
			var t parsersTools.Parser
			t.Job = job
			reflect.ValueOf(t).MethodByName(parser).Call(results)
		} else {
			var t parsersCustoms.Parser
			t.Job = job
			reflect.ValueOf(t).MethodByName(parser).Call(results)
		}
	} else {
		log.ERROR.Println("Task got an error")
	}
	db.ValidateJob(&job, results)

}
