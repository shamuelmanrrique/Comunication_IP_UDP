package functions

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Delays []time.Duration

type Targets []string

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
