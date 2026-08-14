package storage

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/nheath/alfred-gcp-workflow/gcloud"
	"github.com/nheath/alfred-gcp-workflow/parser"
	"github.com/nheath/alfred-gcp-workflow/services"
	"github.com/nheath/alfred-gcp-workflow/workflow/resource"
)

type BucketSearcher struct{}

func (s *BucketSearcher) Search(
	wf *aw.Workflow, svc *services.Service, cfg *gc.Config, q *parser.Result,
) error {
	builder := resource.NewBuilder(
		"storage_buckets",
		wf,
		cfg,
		q,
		gc.ListCloudStorageBuckets,
		func(wf *aw.Workflow, gsb gc.Bucket) {
			sb := FromGCloudStorageBucket(&gsb)
			resource.NewItem(wf, cfg, sb, svc.Icon())
		},
	)

	return builder.Build()
}
