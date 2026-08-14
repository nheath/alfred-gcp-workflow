package cloudrun

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/nheath/alfred-gcp-workflow/gcloud"
	"github.com/nheath/alfred-gcp-workflow/parser"
	"github.com/nheath/alfred-gcp-workflow/services"
	"github.com/nheath/alfred-gcp-workflow/workflow/resource"
)

type ServiceSearcher struct{}

func (s *ServiceSearcher) Search(
	wf *aw.Workflow, svc *services.Service, cfg *gc.Config, q *parser.Result,
) error {
	builder := resource.NewBuilder(
		"cloudrun_services",
		wf,
		cfg,
		q,
		gc.ListCloudRunServices,
		func(wf *aw.Workflow, gcrs gc.CloudRunService) {
			crs := FromGCloudCloudRunService(&gcrs)
			resource.NewItem(wf, cfg, crs, svc.Icon())
		},
	)

	return builder.Build()
}
