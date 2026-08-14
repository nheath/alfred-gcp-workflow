package filestore

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/nheath/alfred-gcp-workflow/gcloud"
	"github.com/nheath/alfred-gcp-workflow/parser"
	"github.com/nheath/alfred-gcp-workflow/services"
	"github.com/nheath/alfred-gcp-workflow/workflow/resource"
)

type InstanceSearcher struct{}

func (s *InstanceSearcher) Search(
	wf *aw.Workflow, svc *services.Service, cfg *gc.Config, q *parser.Result,
) error {
	builder := resource.NewBuilder(
		"filestore_instances",
		wf,
		cfg,
		q,
		gc.ListFilestoreInstances,
		func(wf *aw.Workflow, gfsi gc.FilestoreInstance) {
			fsi := FromGCloudInstance(&gfsi)
			resource.NewItem(wf, cfg, fsi, svc.Icon())
		},
	)

	return builder.Build()
}
