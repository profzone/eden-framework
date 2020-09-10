package confpostgres

import (
	"github.com/profzone/envconfig"
	"time"

	"github.com/sirupsen/logrus"
)

type Retry struct {
	Repeats  int
	Interval envconfig.Duration
}

func (r *Retry) SetDefaults() {
	if r.Repeats == 0 {
		r.Repeats = 3
	}
	if r.Interval == 0 {
		r.Interval = envconfig.Duration(10 * time.Second)
	}
}

func (r Retry) Do(exec func() error) (err error) {
	if r.Repeats <= 0 {
		err = exec()
		return
	}
	for i := 0; i < r.Repeats; i++ {
		err = exec()
		if err != nil {
			logrus.Warningf("retry in seconds [%d]", r.Interval)
			time.Sleep(time.Duration(r.Interval))
		}
	}
	return
}