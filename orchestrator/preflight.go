package orchestrator

import (
	"errors"

	aw "github.com/deanishe/awgo"
	"github.com/nheath/alfred-gcp-workflow/parser"
	"github.com/nheath/alfred-gcp-workflow/workflow/config"
)

var ErrPreflightCheckFailed = errors.New("preflight check failed")

type PreFlight struct{}

func (r *PreFlight) Check(wf *aw.Workflow, result *parser.Result) error {
	if result.HasIntent() {
		return nil
	}

	cfgFile := config.GetConfigFile()
	if cfgFile.GCloudPath == "" {
		wf.NewItem("Setup gcloud cli path you are golden! 🧙✨").
			Subtitle("Use: gcloud-path /your/gcloud").
			Arg("gcloud-path ").
			Autocomplete("gcloud-path ").
			Icon(aw.IconSettings).
			Valid(false)
		wf.SendFeedback()
		return ErrPreflightCheckFailed
	}

	return nil
}
