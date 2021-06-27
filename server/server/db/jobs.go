package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/RichardKnop/machinery/v1/log"
	"github.com/nodauf/ReconFramwork/server/server/models"
	modelsDatabases "github.com/nodauf/ReconFramwork/server/server/models/database"
)

func AddJob(target, parser, machineryTask, machineryTaskArgs string) (modelsDatabases.Job, error) {
	var err error
	host := GetHost(target)
	domain := GetDomain(target)
	var job modelsDatabases.Job
	//job.TaskUUID = taskUUID
	job.Processed = false
	job.Parser = parser
	job.MachineryTask = machineryTask
	job.MachineryTaskArgs = machineryTaskArgs
	// When we do a lot of select with go routing we need to use transaction to lock the table
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if host.ID != 0 {
		job.Host = host
		tx.Create(&job)
		tx.Preload("Host").First(&job)
	} else if domain.ID != 0 {
		job.Domain = domain
		tx.Create(&job)
		tx.Preload("Domain").First(&job)

	} else {
		err = errors.New("Cannot attach the job to an host or domain")
	}
	tx.Commit()
	return job, err
}

func AddJobWithError(target, parser, machineryTask, machineryTaskArgs string, err error) {
	job, errJob := AddJob(target, parser, machineryTask, machineryTaskArgs)
	if errJob == nil {
		result := models.Output{Error: err.Error()}
		resultBytes, errf := json.Marshal(result)

		fmt.Println(errf)
		db.Model(&modelsDatabases.Job{}).Where("id = ?", job.ID).Updates(modelsDatabases.Job{RawOutput: resultBytes})
	} else {
		log.ERROR.Println(err)
	}

}

func UpdateJob(job *modelsDatabases.Job, taskUUID string) {
	db.Model(&modelsDatabases.Job{}).Where("id = ?", job.ID).Update("task_uuid", taskUUID)
}

func ValidateJob(job *modelsDatabases.Job, resultOutput []reflect.Value) {
	if len(resultOutput) > 0 {
		db.Model(&modelsDatabases.Job{}).Where("id = ?", job.ID).Updates(modelsDatabases.Job{Processed: true, RawOutput: resultOutput[0].Bytes()})
	} else {
		log.ERROR.Println("result seems completly empty")
	}
}

func GetNonProcessedTasks() []modelsDatabases.Job {
	var jobs []modelsDatabases.Job
	db.Where("processed = ?", false).Preload("Host").Find(&jobs)
	return jobs
}

func GetAllJobs() []modelsDatabases.Job {
	var listJobs []modelsDatabases.Job

	db.Preload("Host").Preload("Domain").Find(&listJobs)
	return listJobs
}

func GetJob(id uint) modelsDatabases.Job {
	var job modelsDatabases.Job
	db.Preload("Host").Preload("Domain").Where("id = ?", id).Find(&job)
	return job
}
