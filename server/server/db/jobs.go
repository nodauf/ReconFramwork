package db

import (
	"errors"

	modelsDatabases "github.com/nodauf/ReconFramwork/server/server/models/database"
)

func AddJob(target, parser, taskUUID, machineryTask, MachineryTaskArgs string) (modelsDatabases.Job, error) {
	var err error
	host := GetHost(target)
	domain := GetDomain(target)
	var job modelsDatabases.Job
	job.TaskUUID = taskUUID
	job.Processed = false
	job.Parser = parser
	job.MachineryTask = machineryTask
	job.MachineryTaskArgs = MachineryTaskArgs
	if host.ID != 0 {
		job.Host = host
		db.Create(&job)
		db.Preload("Host").First(&job)
	} else if domain.ID != 0 {
		job.Domain = domain
		db.Create(&job)
		db.Preload("Domain").First(&job)

	} else {
		err = errors.New("Cannot attach the job to an host or domain")
	}
	return job, err
}

func RemoveJob(job *modelsDatabases.Job) {
	db.Model(&modelsDatabases.Job{}).Where("id = ?", job.ID).Update("processed", true)
}

func GetNonProcessedTasks() []modelsDatabases.Job {
	var jobs []modelsDatabases.Job
	db.Where("processed = ?", false).Preload("Host").Find(&jobs)
	return jobs
}
