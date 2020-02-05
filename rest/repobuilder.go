package rest

import (
	"fmt"
	"net/http"

	"github.com/evergreen-ci/gimlet"
	"github.com/mongodb/amboy"
	"github.com/mongodb/curator/repobuilder"
	"github.com/pkg/errors"
)

////////////////////////////////////////////////////////////////////////
//
// POST /repobuilder

func (s *Service) addRepobuilderJob(rw http.ResponseWriter, r *http.Request) {
	opts := repobuilder.JobOptions{}
	err := gimlet.GetJSON(r.Body, &opts)
	if err != nil {
		gimlet.WriteResponse(rw, gimlet.MakeJSONErrorResponder(gimlet.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    errors.Wrap(err, "problem parsing input for repobuilder options").Error(),
		}))
		return
	}

	// TODO inject options here:
	// - aws keys from config
	// - working directory of cache

	job, err := repobuilder.NewRepoBuilderJob(opts)
	if err != nil {
		gimlet.WriteResponse(rw, gimlet.MakeJSONErrorResponder(gimlet.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    errors.Wrap(err, "problem building repobuilder job").Error(),
		}))
		return
	}

	if err = s.Environment.RemoteQueue().Put(r.Context(), job); err != nil {
		gimlet.WriteResponse(rw, gimlet.MakeJSONErrorResponder(gimlet.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    errors.Wrap(err, "problem building enquing job").Error(),
		}))
		return
	}

	gimlet.WriteJSON(rw, struct {
		ID     string   `json:"id"`
		Scopes []string `json:"scopes"`
	}{
		ID:     job.ID(),
		Scopes: job.Scopes(),
	})
}

////////////////////////////////////////////////////////////////////////
//
// GET /repobuilder/check/{job_id}

func (s *Service) checkRepobuilderJob(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	jobID := gimlet.GetVars(r)["job_id"]
	queue := s.Environment.RemoteQueue()
	job, ok := queue.Get(ctx, jobID)
	if !ok {
		gimlet.WriteResponse(rw, gimlet.MakeJSONErrorResponder(gimlet.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    fmt.Sprintf("job named '%s' was not found", jobID),
		}))
		return
	}

	gimlet.WriteJSON(rw, struct {
		JobStatus   amboy.JobStatusInfo `json:"job_status"`
		JobTiming   amboy.JobTimeInfo   `json:"job_timing"`
		QueueStatus amboy.QueueStats    `json:"queue_status"`
	}{
		JobStatus:   job.Status(),
		JobTiming:   job.TimeInfo(),
		QueueStatus: queue.Stats(ctx),
	})
}
