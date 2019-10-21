package functions

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Delays []time.Duration

type Targets []string

type CoordinatesInt interface {
	GetProcess()
	GetMaster()
	GetTimeDelay()
	GetTarget()
	GetRun()
	GetPort()
}

type Coordinates struct {
	Process   int
	Master    bool
	TimeDelay Delays
	Target    Targets
	Run       string
	Port      string
}

func (i *Delays) String() string {
	return fmt.Sprint(*i)
}

func (i *Targets) String() string {
	return fmt.Sprint(*i)
}

func (i *Delays) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("Delays flag already set")
	}
	for _, dt := range strings.Split(value, ",") {
		duration, err := time.ParseDuration(dt)
		if err != nil {
			return err
		}
		*i = append(*i, duration)
	}
	return nil
}

func (i *Targets) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("Delays flag already set")
	}
	for _, dt := range strings.Split(value, ",") {
		*i = append(*i, dt)
	}
	return nil
}

func (c Coordinates) GetProcess() int {
	return c.Process
}
func (c Coordinates) GetMaster() bool {
	return c.Master
}
func (c Coordinates) GetTimeDelay() Delays {
	return c.TimeDelay
}
func (c Coordinates) GetTarget() []string {
	return c.Target
}
func (c Coordinates) GetRun() string {
	return c.Run
}
func (c Coordinates) GetPort() string {
	return c.Port
}
