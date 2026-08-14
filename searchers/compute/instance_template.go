package compute

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/nheath/alfred-gcp-workflow/gcloud"
	"github.com/nheath/alfred-gcp-workflow/parser"
	"github.com/nheath/alfred-gcp-workflow/services"
	"github.com/nheath/alfred-gcp-workflow/workflow/resource"
)

type InstanceTmplSearcher struct{}

func (s *InstanceTmplSearcher) Search(
	wf *aw.Workflow, svc *services.Service, cfg *gc.Config, q *parser.Result,
) error {
	builder := resource.NewBuilder(
		"compute_instance_templates",
		wf,
		cfg,
		q,
		gc.ListComputeInstanceTemplates,
		func(wf *aw.Workflow, gcit gc.ComputeInstanceTemplate) {
			cit := FromGCloudComputeInstanceTemplate(&gcit)
			resource.NewItem(wf, cfg, cit, svc.Icon())
		},
	)

	return builder.Build()
}
