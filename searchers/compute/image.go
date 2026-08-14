package compute

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/nheath/alfred-gcp-workflow/gcloud"
	"github.com/nheath/alfred-gcp-workflow/parser"
	"github.com/nheath/alfred-gcp-workflow/services"
	"github.com/nheath/alfred-gcp-workflow/workflow/resource"
)

type ImageSearcher struct{}

func (s *ImageSearcher) Search(wf *aw.Workflow, svc *services.Service, cfg *gc.Config, q *parser.Result) error {
	builder := resource.NewBuilder(
		"compute_images",
		wf,
		cfg,
		q,
		gc.ListComputeImages,
		func(wf *aw.Workflow, gci gc.ComputeImage) {
			ci := FromGCloudComputeImage(&gci)
			resource.NewItem(wf, cfg, ci, svc.Icon())
		},
	)

	return builder.Build()
}
