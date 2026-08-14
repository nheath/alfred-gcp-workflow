package sql

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
		"sql_instances",
		wf,
		cfg,
		q,
		gc.ListSQLInstances,
		func(wf *aw.Workflow, gsi gc.SQLInstance) {
			si := FromGCloudSQLInstance(&gsi)
			resource.NewItem(wf, cfg, si, svc.Icon())
		},
	)
	return builder.Build()
}
