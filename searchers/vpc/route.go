package vpc

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/nheath/alfred-gcp-workflow/gcloud"
	"github.com/nheath/alfred-gcp-workflow/parser"
	"github.com/nheath/alfred-gcp-workflow/services"
	"github.com/nheath/alfred-gcp-workflow/workflow/resource"
)

type RouteSearcher struct{}

func (r *RouteSearcher) Search(
	wf *aw.Workflow, svc *services.Service, cfg *gc.Config, pq *parser.Result,
) error {
	builder := resource.NewBuilder(
		"vpc_routes",
		wf,
		cfg,
		pq,
		gc.ListVPCRoutes,
		func(wf *aw.Workflow, gvr gc.VPCRoute) {
			vr := FromGCloudRoute(&gvr)
			resource.NewItem(wf, cfg, vr, svc.Icon())
		},
	)
	return builder.Build()
}
