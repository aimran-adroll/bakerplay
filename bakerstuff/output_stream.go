package bakerstuff

import (
	"fmt"
	"sync/atomic"

	"github.com/AdRoll/baker"
)

var EchoDesc = baker.OutputDesc{
	Name:   "Echo",
	New:    NewEcho,
	Config: &EchoConfig{},
	Help:   "",
}

type Echo struct {
	totaln int64
}

type EchoConfig struct {
}

func NewEcho(cfg baker.OutputParams) (baker.Output, error) {
	return &Echo{}, nil
}

func (e *Echo) CanShard() bool { return false }

func (e *Echo) Run(input <-chan baker.OutputRecord, _ chan<- string) error {
	for ll := range input {
		fmt.Printf("%+v\n", ll.Fields)
		atomic.AddInt64(&e.totaln, 1)
	}

	return nil
}

func (e *Echo) Stats() baker.OutputStats {
	return baker.OutputStats{
		NumProcessedLines: atomic.LoadInt64(&e.totaln),
		NumErrorLines:     0,
	}
}
