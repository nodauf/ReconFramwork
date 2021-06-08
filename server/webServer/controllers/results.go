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
	c.TplName = "recon/results/resultsOverview.tpl"
}

func (c *ReconController) ListResultsWeb() {
	var results []modelsBeego.Result
	var sizeRowSpan []int
	sizeRowSpan, results = hostsToResults(db.GetAllHostsWhereServices(utils.GetWebService()))
	c.Data["Results"] = results
	c.Data["SizeRowSpan"] = sizeRowSpan
	c.Data["Modal"] = true
	c.Data["DataTables"] = true
	c.Layout = "recon/includes/layout.tpl"
	c.TplName = "recon/results/resultsTable.tpl"
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

			returnValue := m.Call(argument)
			content := returnValue[0]
			html := returnValue[1]
			if html.Bool() {
				c.Data["Html"] = content.String()
			} else {
				c.Data["Content"] = content.String()
			}
		}
	}
	c.TplName = "recon/includes/modal-content.tpl"
}

func (c *ReconController) TreeResults() {
	var nodeData, nodeLink string
	ip := c.Ctx.Input.Param(":ip")
	if ip != "" {
		host := db.GetHost(ip)
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
					url: "` + web.URLFor("ReconController.DetailsResultsWeb", ":ip", host.Address, ":port", port.Port, ":task", comment.Task) + `"
				  }`
				nodeLink += `,{
					from: "p` + strconv.Itoa(i) + `",
					to: "p` + strconv.Itoa(i) + `c` + strconv.Itoa(j) + `"
				  }`

			}

		}
		for i, domain := range host.Domain {
			nodeData += `,{
				key: "d` + strconv.Itoa(i) + `",
				category: "Domain",
				value: "` + domain.Domain + `"
			  }`
			nodeLink += `,{
				from: "domains",
				to: "d` + strconv.Itoa(i) + `"
			  }`
		}

		c.Data["Tree"] = nodeData + "]\n" + nodeLink + "]"
	}
	c.Layout = "recon/includes/layout.tpl"
	c.TplName = "recon/results/resultsTree.tpl"
}

func (c *ReconController) ListAllResults() {
	var results []modelsBeego.Result
	var sizeRowSpan []int
	sizeRowSpan, results = hostsToResults(db.GetAllHosts())
	c.Data["Results"] = results
	c.Data["SizeRowSpan"] = sizeRowSpan
	c.Data["Modal"] = true
	c.Data["DataTables"] = true
	c.Layout = "recon/includes/layout.tpl"
	c.TplName = "recon/results/resultsTable.tpl"
}

func hostsToResults(hosts []modelsDatabases.Host) ([]int, []modelsBeego.Result) {
	var sizeRowSpan []int
	var results []modelsBeego.Result

	// Loop to create the table with a port per line
	for _, host := range hosts {
		var result modelsBeego.Result
		result.Address = host.Address
		for _, domain := range host.Domain {
			result.Domain = append(result.Domain, domain.Domain)
		}
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
					result.Port = port.Port
					result.Task = portComment.Task
					results = append(results, result)
				}
			} else {
				result.Port = port.Port
				results = append(results, result)
			}
		}

	}
	return sizeRowSpan, results
}
