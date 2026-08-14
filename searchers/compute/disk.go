package compute

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/nheath/alfred-gcp-workflow/gcloud"
	"github.com/nheath/alfred-gcp-workflow/parser"
	"github.com/nheath/alfred-gcp-workflow/services"
	"github.com/nheath/alfred-gcp-workflow/workflow/resource"
)

type DiskSearcher struct{}

func (s *DiskSearcher) Search(
	wf *aw.Workflow, svc *services.Service, cfg *gc.Config, q *parser.Result,
) error {
	builder := resource.NewBuilder(
		"compute_disks",
		wf,
		cfg,
		q,
		gc.ListComputeDisks,
		func(wf *aw.Workflow, gcd gc.ComputeDisk) {
			cd := FromGCloudComputeDisk(&gcd)
			resource.NewItem(wf, cfg, cd, svc.Icon())
		},
	)

	return builder.Build()
}
