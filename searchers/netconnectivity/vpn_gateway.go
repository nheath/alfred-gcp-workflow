package netconnectivity

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/nheath/alfred-gcp-workflow/gcloud"
	"github.com/nheath/alfred-gcp-workflow/parser"
	"github.com/nheath/alfred-gcp-workflow/services"
	"github.com/nheath/alfred-gcp-workflow/workflow/resource"
)

type VPNGatewaySearcher struct{}

func (s *VPNGatewaySearcher) Search(
	wf *aw.Workflow, svc *services.Service, cfg *gc.Config, q *parser.Result,
) error {
	builder := resource.NewBuilder(
		"netconnectivity_vpn_gateways",
		wf,
		cfg,
		q,
		gc.ListVPNGateways,
		func(wf *aw.Workflow, gvc gc.VPNGateway) {
			vg := FromGCloudVPNGateway(&gvc)
			resource.NewItem(wf, cfg, vg, svc.Icon())
		},
	)
	return builder.Build()
}
