package options

import (
	"github.com/mongodb/grip"
	"github.com/mongodb/grip/level"
)

// Command represents jasper.Command options that are configurable by the
// user.
type Command struct {
	ID              string            `json:"id,omitempty"`
	Commands        [][]string        `json:"commands"`
	Process         Create            `json:"proc_opts,omitempty"`
	Remote          *Remote           `json:"remote_options,omitempty"`
	ContinueOnError bool              `json:"continue_on_error,omitempty"`
	IgnoreError     bool              `json:"ignore_error,omitempty"`
	Priority        level.Priority    `json:"priority,omitempty"`
	RunBackground   bool              `json:"run_background,omitempty"`
	Sudo            bool              `json:"sudo,omitempty"`
	SudoUser        string            `json:"sudo_user,omitempty"`
	Prerequisite    func() bool       `json:"-"`
	Hook            func(error) error `json:"-"`
}

// Validate ensures that the options passed to the command are valid.
func (opts *Command) Validate() error {
	catcher := grip.NewBasicCatcher()
	// The semantics of options.Create expects Args to be non-empty, but Command
	// ignores these args.
	if len(opts.Process.Args) == 0 {
		opts.Process.Args = []string{""}
	}
	catcher.Add(opts.Process.Validate())
	catcher.NewWhen(opts.Priority != 0 && opts.Priority.IsValid(), "priority is not in the valid range of values")
	catcher.NewWhen(len(opts.Commands) == 0, "must specify at least one command")
	return catcher.Resolve()
}
