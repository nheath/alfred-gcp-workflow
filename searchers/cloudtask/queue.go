package cloudtask

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/nheath/alfred-gcp-workflow/gcloud"
	"github.com/nheath/alfred-gcp-workflow/parser"
	"github.com/nheath/alfred-gcp-workflow/services"
	"github.com/nheath/alfred-gcp-workflow/workflow/resource"
)

type QueueSearcher struct{}

func (qs *QueueSearcher) Search(wf *aw.Workflow, svc *services.Service, cfg *gc.Config, q *parser.Result,
) error {
	builder := resource.NewBuilder(
		"cloudtasks_queues",
		wf,
		cfg,
		q,
		gc.ListCloudTaskQueues,
		func(wf *aw.Workflow, gctq gc.CloudTaskQueue) {
			tq := FromGCloudCloudTaskQueue(&gctq)
			resource.NewItem(wf, cfg, tq, svc.Icon())
		},
	)

	return builder.Build()
}
