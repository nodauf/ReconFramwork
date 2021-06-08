package controllers

import (
	"reflect"

	"github.com/RichardKnop/machinery/v1/log"
	"github.com/nodauf/ReconFramwork/server/server/config"
	"github.com/nodauf/ReconFramwork/server/server/db"
	parsersTools "github.com/nodauf/ReconFramwork/server/server/parsers/tools"
	modelsBeego "github.com/nodauf/ReconFramwork/server/webServer/models"
	"github.com/nodauf/ReconFramwork/utils"
)

func (c *ReconController) ListResults() {
	var results []modelsBeego.Result

	for _, host := range db.GetAllHosts() {
		var result modelsBeego.Result
		result.Address = host.Address
		for _, domain := range host.Domain {
			result.Domain = append(result.Domain, domain.Domain)
		}
		result.NbPorts = len(host.Ports)
		results = append(results, result)
	}
	c.Data["Results"] = results
	c.Data["DataTables"] = true
	c.Layout = "recon/includes/layout.tpl"
	c.TplName = "recon/results/resultsList.tpl"
}

func (c *ReconController) ListResultsWeb() {
	var results []modelsBeego.ResultWeb
	var sizeRowSpan []int
	// Loop to create the table with a port per line
	for _, host := range db.GetAllHostsWhereServices(utils.GetWebService()) {
		var result modelsBeego.ResultWeb
		result.Address = host.Address
		// Tricks to display the column domain as if the domain slice is nil it will not display. This behavior is necessary for print rows with rowspan
		//if len(host.Domain) > 0 {
		for _, domain := range host.Domain {
			result.Domain = append(result.Domain, domain.Domain)
		}
		//} else {
		//result.Domain = append(result.Domain, " ")
		//}
		for _, port := range host.Ports {

			// Get max number of the three element
			if len(result.Domain) > len(host.Ports) && len(result.Domain) > len(port.PortComment) {
				sizeRowSpan = append(sizeRowSpan, len(host.Domain))
			} else if len(host.Ports) > len(result.Domain) && len(host.Ports) > len(port.PortComment) {
				sizeRowSpan = append(sizeRowSpan, len(host.Ports))
			} else {
				sizeRowSpan = append(sizeRowSpan, len(port.PortComment))
			}
			if len(port.PortComment) > 0 {
				for _, portComment := range port.PortComment {
					result.WebPort = port.Port
					result.Task = portComment.Task
					results = append(results, result)
				}
			} else {
				result.WebPort = port.Port
				results = append(results, result)
			}
		}

	}
	c.Data["Results"] = results
	c.Data["SizeRowSpan"] = sizeRowSpan
	c.Data["Modal"] = true
	c.Data["DataTables"] = true
	c.Layout = "recon/includes/layout.tpl"
	c.TplName = "recon/results/resultsWeb.tpl"
}

func (c *ReconController) DetailsResultsWeb() {
	ip := c.Ctx.Input.Param(":ip")
	port := c.Ctx.Input.Param(":port")
	task := c.Ctx.Input.Param(":task")
	if ip != "" && port != "" && task != "" {
		commandOutput := db.GetPortComment(ip, port, task)
		v := reflect.ValueOf(parsersTools.Parser{})
		m := v.MethodByName(config.Config.Command[task].PrintFunction)
		if m.Kind() != reflect.Func {
			log.ERROR.Println("The function " + config.Config.Command[task].PrintFunction + " for " + task + " not found")
			c.Data["Content"] = "The function " + config.Config.Command[task].PrintFunction + " for " + task + " not found"
		} else {
			var argument []reflect.Value
			argument = append(argument, reflect.ValueOf(commandOutput))

			c.Data["Content"] = m.Call(argument)[0]
		}
	}
	c.TplName = "recon/includes/modal-content.tpl"
}
