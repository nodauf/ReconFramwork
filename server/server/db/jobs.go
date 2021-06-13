package db

import (
	"errors"

	modelsDatabases "github.com/nodauf/ReconFramwork/server/server/models/database"
)

func AddJob(target, parser, machineryTask, MachineryTaskArgs string) (modelsDatabases.Job, error) {
	var err error
	host := GetHost(target)
	domain := GetDomain(target)
	var job modelsDatabases.Job
	//job.TaskUUID = taskUUID
	job.Processed = false
	job.Parser = parser
	job.MachineryTask = machineryTask
	job.MachineryTaskArgs = MachineryTaskArgs
	// When we do a lot of select with go routing we need to use transaction to lock the table
	tx := db.Begin()

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

func UpdateJob(job *modelsDatabases.Job, taskUUID string) {
	db.Model(&modelsDatabases.Job{}).Where("id = ?", job.ID).Update("task_uuid", taskUUID)
}

func RemoveJob(job *modelsDatabases.Job) {
	db.Model(&modelsDatabases.Job{}).Where("id = ?", job.ID).Update("processed", true)
}

func GetNonProcessedTasks() []modelsDatabases.Job {
	var jobs []modelsDatabases.Job
	db.Where("processed = ?", false).Preload("Host").Find(&jobs)
	return jobs
}
