package rest

import (
	"github.com/evergreen-ci/barque"
	"github.com/evergreen-ci/gimlet"
)

type Service struct {
	Environment barque.Environment
	umconf      gimlet.UserMiddlewareConfiguration
}

func New(env Environment) (*gimlet.APIApp, error) {
	app := gimlet.NewApp()

}
