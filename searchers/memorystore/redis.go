package memorystore

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/nheath/alfred-gcp-workflow/gcloud"
	"github.com/nheath/alfred-gcp-workflow/parser"
	"github.com/nheath/alfred-gcp-workflow/services"
	"github.com/nheath/alfred-gcp-workflow/workflow/resource"
)

type RedisInstanceSearcher struct{}

func (s *RedisInstanceSearcher) Search(
	wf *aw.Workflow, svc *services.Service, cfg *gc.Config, q *parser.Result,
) error {
	builder := resource.NewBuilder(
		"memorystore_redis_instances",
		wf,
		cfg,
		q,
		gc.ListRedisInstances,
		func(wf *aw.Workflow, gri gc.RedisInstance) {
			ri := FromGCloudRedisInstance(&gri)
			resource.NewItem(wf, cfg, ri, svc.Icon())
		},
	)

	return builder.Build()
}
