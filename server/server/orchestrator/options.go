package orchestrator

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/log"
	"github.com/nodauf/ReconFramwork/utils"
)

var server *machinery.Server
var wg *sync.WaitGroup

type Options struct {
	RecurseOnSubdomain bool
	Target             string
	Task               string
	Workflow           string
	Wg                 *sync.WaitGroup
	Server             *machinery.Server
}

func init() {
	srv, err := utils.GetMachineryServer()
	if err != nil {
		log.ERROR.Fatalln(err)
	}
	server = srv
	wg = &sync.WaitGroup{}

}

func NewOptions() Options {
	var options Options
	options.Wg = wg
	options.Server = server
	return options
}

func (options Options) PrintOptions() {
	fmt.Println("RecurseOnSubdomain => " + strconv.FormatBool(options.RecurseOnSubdomain))
	fmt.Println("Target => " + options.Target)
	fmt.Println("Task => " + options.Task)
	fmt.Println("Workflow => " + options.Workflow)
}
