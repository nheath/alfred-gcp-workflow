package netconnectivity

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/nheath/alfred-gcp-workflow/gcloud"
	"github.com/nheath/alfred-gcp-workflow/parser"
	"github.com/nheath/alfred-gcp-workflow/services"
	"github.com/nheath/alfred-gcp-workflow/workflow/resource"
)

type VPNTunnelSearcher struct{}

func (s *VPNTunnelSearcher) Search(
	wf *aw.Workflow, svc *services.Service, cfg *gc.Config, q *parser.Result,
) error {
	builder := resource.NewBuilder(
		"netconnectivity_vpn_tunnels",
		wf,
		cfg,
		q,
		gc.ListVPNTunnels,
		func(wf *aw.Workflow, gvt gc.VPNTunnel) {
			vt := FromGCloudVPNTunnel(&gvt)
			resource.NewItem(wf, cfg, vt, svc.Icon())
		},
	)

	return builder.Build()
}
