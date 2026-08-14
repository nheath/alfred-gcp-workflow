package netservices

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/nheath/alfred-gcp-workflow/gcloud"
	"github.com/nheath/alfred-gcp-workflow/parser"
	"github.com/nheath/alfred-gcp-workflow/services"
	"github.com/nheath/alfred-gcp-workflow/workflow/resource"
)

type DNSZoneSearcher struct{}

func (s *DNSZoneSearcher) Search(
	wf *aw.Workflow, svc *services.Service, cfg *gc.Config, q *parser.Result,
) error {
	builder := resource.NewBuilder(
		"netservices_dns_zones",
		wf,
		cfg,
		q,
		gc.ListDNSZones,
		func(wf *aw.Workflow, gdz gc.DNSZone) {
			dz := FromGCloudDNSZone(&gdz)
			resource.NewItem(wf, cfg, dz, svc.Icon())
		},
	)

	return builder.Build()
}
