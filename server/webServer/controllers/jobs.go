package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/beego/beego/v2/server/web"
	"github.com/nodauf/ReconFramwork/server/server/db"
	"github.com/nodauf/ReconFramwork/server/server/models"
	modelsBeego "github.com/nodauf/ReconFramwork/server/webServer/models"
)

func (c *ReconController) ListJobs() {
	web.ReadFromRequest(&c.Controller)
	var results []modelsBeego.ResultJob

	for _, job := range db.GetAllJobs() {
		var result modelsBeego.ResultJob
		result.ID = job.ID
		if job.Domain.Domain != "" {
			result.Target = job.Domain.Domain
		} else if job.Host.Address != "" {
			result.Target = job.Host.Address
		} else {
			result.Target = "The job is not link to a target, should not happened"
		}
		result.ExecutionTime = job.CreatedAt.String()
		result.MachineryTask = job.MachineryTask
		result.MachineryArgs = job.MachineryTaskArgs
		result.Processed = job.Processed
		results = append(results, result)
	}
	c.Data["Results"] = results
	c.Data["ModalTabs"] = true
	c.Data["DataTables"] = true
	c.Layout = "recon/includes/layout.tpl"
	c.TplName = "recon/jobs/jobsOverview.tpl"
}

func (c *ReconController) DetailsJob() {
	id64, err := strconv.ParseUint(c.Ctx.Input.Param(":id"), 10, 32)
	id := uint(id64)
	if id != 0 && err == nil {
		var output models.Output
		job := db.GetJob(id)
		json.Unmarshal(job.RawOutput, &output)
		c.Data["Content"] = output.Stdout

	}
	c.TplName = "recon/includes/modal-content.tpl"
}
