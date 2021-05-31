package orchestrator

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/RichardKnop/machinery/v1"
)

type Options struct {
	RecurseOnSubdomain bool
	Target             string
	Task               string
	Workflow           string
	Wg                 *sync.WaitGroup
	Server             *machinery.Server
}

func (options Options) PrintOptions() {
	fmt.Println("RecurseOnSubdomain => " + strconv.FormatBool(options.RecurseOnSubdomain))
	fmt.Println("Target => " + options.Target)
	fmt.Println("Task => " + options.Task)
	fmt.Println("Workflow => " + options.Workflow)
}
