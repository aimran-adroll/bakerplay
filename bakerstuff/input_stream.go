package bakerstuff

import (
	"fmt"
	"sync/atomic"

	"github.com/AdRoll/baker"
)

var StreamDesc = baker.InputDesc{
	Name:   "Stream",
	New:    NewStream,
	Config: &StreamConfig{},
	Help:   "",
}

type StreamConfig struct{}

type Stream struct {
	numLines int64
	data     chan<- *baker.Data
	done     chan bool
}

func NewStream(cfg baker.InputParams) (baker.Input, error) {

	s := &Stream{}

	return s, nil
}

func (s *Stream) FreeMem(data *baker.Data) {
}

func (s *Stream) Run(inch chan<- *baker.Data) error {
	s.data = inch

	for i := 0; i < 10; i++ {
		logLine := baker.LogLine{
			FieldSeparator: ',',
		}
		num := fmt.Sprintf("%d", i)
		logLine.Set(0, []byte(num))

		s.data <- &baker.Data{
			Bytes: logLine.ToText(nil),
		}
	}

	<-s.done
	return nil
}

func (s *Stream) Stop() {
	fmt.Printf("stopping input")
	// close(s.done)
}

func (s *Stream) Stats() baker.InputStats {
	return baker.InputStats{
		NumProcessedLines: atomic.LoadInt64(&s.numLines),
	}

}
