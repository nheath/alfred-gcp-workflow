package monitoring

import (
	aw "github.com/deanishe/awgo"

	gc "github.com/nheath/alfred-gcp-workflow/gcloud"
	"github.com/nheath/alfred-gcp-workflow/parser"
	"github.com/nheath/alfred-gcp-workflow/services"
	"github.com/nheath/alfred-gcp-workflow/workflow/resource"
)

type DashboardSearcher struct{}

func (s *DashboardSearcher) Search(
	wf *aw.Workflow, svc *services.Service, cfg *gc.Config, q *parser.Result,
) error {
	builder := resource.NewBuilder(
		"monitoring_dashboards",
		wf,
		cfg,
		q,
		gc.ListMonitoringDashboards,
		func(wf *aw.Workflow, dash gc.Dashboard) {
			sb := FromGCloudMonitoringDashboard(&dash)
			resource.NewItem(wf, cfg, sb, svc.Icon())
		},
	)

	return builder.Build()
}
