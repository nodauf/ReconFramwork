package controllers

import (
	"reflect"
	"strconv"

	"github.com/RichardKnop/machinery/v1/log"
	"github.com/beego/beego/v2/server/web"
	"github.com/nodauf/ReconFramwork/server/server/config"
	"github.com/nodauf/ReconFramwork/server/server/db"
	modelsDatabases "github.com/nodauf/ReconFramwork/server/server/models/database"
	parsersTools "github.com/nodauf/ReconFramwork/server/server/parsers/tools"
	"github.com/nodauf/ReconFramwork/server/server/utils"
	modelsBeego "github.com/nodauf/ReconFramwork/server/webServer/models"
)

func (c *ReconController) OverviewResults() {
	web.ReadFromRequest(&c.Controller)
	var results []modelsBeego.ResultTask

	for _, host := range db.GetAllHosts() {
		var result modelsBeego.ResultTask
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
	c.TplName = "recon/results/resultsOverview.tpl"
}

func (c *ReconController) ListResultsWeb() {
	var results []modelsBeego.ResultTask
	results = hostsToResults(db.GetAllHostsWhereServices(utils.GetWebService()))
	c.Data["Results"] = results
	c.Data["Modal"] = true
	c.Data["DataTables"] = true
	c.Layout = "recon/includes/layout.tpl"
	c.TplName = "recon/results/resultsTable.tpl"
}

func (c *ReconController) DetailsResults() {
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

			returnValue := m.Call(argument)
			content := returnValue[0]
			html := returnValue[1]
			if html.Bool() {
				c.Data["Html"] = true
				c.Data["Content"] = content.String()
			} else {
				c.Data["Content"] = content.String()
			}
		}
	}
	c.TplName = "recon/includes/modal-content.tpl"
}

func (c *ReconController) TreeResults() {
	flash := web.ReadFromRequest(&c.Controller)
	var nodeData, nodeLink string
	ip := c.Ctx.Input.Param(":ip")
	if ip != "" {
		host := db.GetHost(ip)
		if host.Address != "" {
			nodeData = `var nodeDataArray = [{
			key: "host",
			value: "` + host.Address + `",
		  },
		  {
			key: "ports",
			value: "Ports",
		  },
		  {
			key: "domains",
			value: "Domains",
		  }`
			nodeLink = `var linkDataArray = [{
			from: "host",
			to: "ports"
		  },
		  {
			from: "host",
			to: "domains"
		  }`
			for i, port := range host.Ports {
				nodeData += `,{
				key: "p` + strconv.Itoa(i) + `",
				category: "Port",
				value: "` + strconv.Itoa(port.Port) + `"
			  }`
				nodeLink += `,{
				from: "ports",
				to: "p` + strconv.Itoa(i) + `"
			  }`
				for j, comment := range port.PortComment {
					nodeData += `,{
					key: "p` + strconv.Itoa(i) + `c` + strconv.Itoa(j) + `",
					category: "PortComment",
					value: "` + comment.Task + `",
					url: "` + web.URLFor("ReconController.DetailsResults", ":ip", host.Address, ":port", port.Port, ":task", comment.Task) + `"
				  }`
					nodeLink += `,{
					from: "p` + strconv.Itoa(i) + `",
					to: "p` + strconv.Itoa(i) + `c` + strconv.Itoa(j) + `"
				  }`
					if comment.DomainID != 0 {
						nodeLink += `,{
						from: "d` + strconv.Itoa(int(comment.DomainID)) + `",
						to: "p` + strconv.Itoa(i) + `c` + strconv.Itoa(j) + `"
					  }`
					}

				}

			}
			for _, domain := range host.Domain {
				nodeData += `,{
				key: "d` + strconv.Itoa(int(domain.ID)) + `",
				category: "Domain",
				value: "` + domain.Domain + `"
			  }`
				nodeLink += `,{
				from: "domains",
				to: "d` + strconv.Itoa(int(domain.ID)) + `"
			  }`
			}

			c.Data["Tree"] = nodeData + "]\n" + nodeLink + "]"
		} else {
			flash.Error("Target not found")
			flash.Store(&c.Controller)
		}
	} else {
		flash.Error("Target argument not found")
		flash.Store(&c.Controller)
	}
	c.Layout = "recon/includes/layout.tpl"
	c.TplName = "recon/results/resultsTree.tpl"
}

func (c *ReconController) ListAllResults() {
	var results []modelsBeego.ResultTask
	results = hostsToResults(db.GetAllHosts())
	c.Data["Results"] = results
	c.Data["Modal"] = true
	c.Data["DataTables"] = true
	c.Layout = "recon/includes/layout.tpl"
	c.TplName = "recon/results/resultsTable.tpl"
}

func hostsToResults(hosts []modelsDatabases.Host) []modelsBeego.ResultTask {
	var results []modelsBeego.ResultTask

	// Loop to create the table with a port per line
	for _, host := range hosts {
		var result modelsBeego.ResultTask
		var domainsOfTheHost []string

		result.Address = host.Address
		for _, domain := range host.Domain {
			domainsOfTheHost = append(domainsOfTheHost, domain.Domain)
		}
		for _, port := range host.Ports {

			if len(port.PortComment) > 0 {
				for _, portComment := range port.PortComment {
					result.Port = port.Port
					result.Task = portComment.Task
					result.Comment = portComment.Comment
					if portComment.DomainID != 0 {
						result.Domain = []string{portComment.Domain.Domain}
					} else {
						result.Domain = domainsOfTheHost
					}
					results = append(results, result)
				}
			} else {
				result.Port = port.Port
				results = append(results, result)
			}
		}

	}
	return results
}
