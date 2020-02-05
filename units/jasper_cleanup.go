package units

import (
	"context"
	"fmt"
	"time"

	"github.com/evergreen-ci/barque"
	"github.com/mongodb/amboy"
	"github.com/mongodb/amboy/dependency"
	"github.com/mongodb/amboy/job"
	"github.com/mongodb/amboy/registry"
)

const jasperManagerCleanupJobName = "jasper-manager-cleanup"

func init() {
	registry.AddJobType(jasperManagerCleanupJobName,
		func() amboy.Job { return makeJasperManagerCleanup() })
}

type jasperManagerCleanup struct {
	job.Base `bson:"job_base" json:"job_base" yaml:"job_base"`
	env      barque.Environment
}

// NewJasperManagerCleanup reports basic system information and a
// report of the go runtime information, as provided by grip.
func NewJasperManagerCleanup(id string, env barque.Environment) amboy.Job {
	j := makeJasperManagerCleanup()
	j.env = env
	j.SetID(fmt.Sprintf("%s.%s", jasperManagerCleanupJobName, id))
	ti := j.TimeInfo()
	ti.MaxTime = time.Minute
	j.UpdateTimeInfo(ti)
	return j
}

func makeJasperManagerCleanup() *jasperManagerCleanup {
	j := &jasperManagerCleanup{
		Base: job.Base{
			JobType: amboy.JobType{
				Name:    jasperManagerCleanupJobName,
				Version: 0,
			},
		},
	}
	j.SetDependency(dependency.NewAlways())
	return j
}

func (j *jasperManagerCleanup) Run(ctx context.Context) {
	defer j.MarkComplete()

	if j.env == nil {
		j.env = barque.GetEnvironment()
	}

	j.env.Jasper().Clear(ctx)
}
