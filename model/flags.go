package model

import (
	"github.com/evergreen-ci/barque"
	"github.com/mongodb/anser/bsonutil"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

type OperationalFlags struct {
	DisableRepobuilderJobExecution  bool `bson:"disable_repobuilder_job_exeuction" json:"disable_repobuilder_job_exeuction" yaml:"disable_repobuilder_job_exeuction"`
	DisableRepobuilderJobSubmission bool `bson:"disable_repobuilder_job_submission" json:"disable_repobuilder_job_submission" yaml:"disable_repobuilder_job_submission"`
	DisableRepobuilderBackgroundJob bool `bson:"disable_repobuilder_background_job" json:"disable_repobuilder_background_job" yaml:"disable_repobuilder_background_job"`
	DisableInternalMetricsReporting bool `bson:"disalbe_internal_metrics_reporting" json:"disalbe_internal_metrics_reporting" yaml:"disalbe_internal_metrics_reporting"`

	env barque.Environment
}

var (
	opsFlagsDisableRepobuilderJobExecution  = bsonutil.MustHaveTag(OperationalFlags{}, "DisableRepobuilderJobExecution")
	opsFlagsDisableRepobuilderJobSubmission = bsonutil.MustHaveTag(OperationalFlags{}, "DisableRepobuilderJobSubmission")
	opsFlagsDisableRepobuilderBackgroundJob = bsonutil.MustHaveTag(OperationalFlags{}, "DisableRepobuilderBackgroundJob")
	opsFlagsDisableInternalMetricsReporting = bsonutil.MustHaveTag(OperationalFlags{}, "DisableInternalMetricsReporting")
)

func (f *OperationalFlags) findAndSet(name string, v bool) error {
	switch name {
	case "disable_repobuilder_job_submission":
		return f.SetDisableRepobuilderJobSubmission(v)
	case "disable_repobuilder_background_job":
		return f.SetDisableInternalMetricsReporting(v)
	case "disable_repobuilder_job_execution":
		return f.SetDisableRepobuilderJobExecution(v)
	case "disable_internal_metrics_reporting":
		return f.SetDisableInternalMetricsReporting(v)
	default:
		return errors.Errorf("%s is not a known feature flag name", name)
	}
}

func (f *OperationalFlags) SetTrue(name string) error {
	return errors.WithStack(f.findAndSet(name, true))
}

func (f *OperationalFlags) SetFalse(name string) error {
	return errors.WithStack(f.findAndSet(name, false))
}

func (f *OperationalFlags) SetDisableRepobuilderJobSubmission(v bool) error {
	if err := f.update(opsFlagsDisableRepobuilderJobSubmission, v); err != nil {
		return errors.WithStack(err)
	}
	f.DisableRepobuilderJobSubmission = v
	return nil
}

func (f *OperationalFlags) SetDisableRepobuilderJobExecution(v bool) error {
	if err := f.update(opsFlagsDisableRepobuilderJobExecution, v); err != nil {
		return errors.WithStack(err)
	}
	f.DisableRepobuilderJobExecution = v
	return nil
}

func (f *OperationalFlags) SetDsiableRepobuilderBackgroundJob(v bool) error {
	if err := f.update(opsFlagsDisableRepobuilderBackgroundJob, v); err != nil {
		return errors.WithStack(err)
	}
	f.DisableRepobuilderBackgroundJob = v
	return nil
}

func (f *OperationalFlags) SetDisableInternalMetricsReporting(v bool) error {
	if err := f.update(opsFlagsDisableInternalMetricsReporting, v); err != nil {
		return errors.WithStack(err)
	}
	f.DisableInternalMetricsReporting = v
	return nil
}

func (f *OperationalFlags) update(key string, value bool) error {
	ctx, cancel := f.env.Context()
	defer cancel()

	res, err := f.env.DB().Collection(configCollection).UpdateOne(ctx, bson.M{"_id": configID},
		bson.M{"$set": bson.M{bsonutil.GetDottedKeyName(confFlagsKey, key): value}})
	if err != nil {
		return errors.Wrapf(err, "problem setting %s to %t", key, value)
	}

	if res.MatchedCount > 0 && res.ModifiedCount != 1 {
		return errors.New("document found but not modified")
	}

	return nil
}
