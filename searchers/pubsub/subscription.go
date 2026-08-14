package pubsub

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/nheath/alfred-gcp-workflow/gcloud"
	"github.com/nheath/alfred-gcp-workflow/parser"
	"github.com/nheath/alfred-gcp-workflow/services"
	"github.com/nheath/alfred-gcp-workflow/workflow/resource"
)

type SubscriptionSearcher struct{}

func (s *SubscriptionSearcher) Search(
	wf *aw.Workflow, svc *services.Service, cfg *gc.Config, q *parser.Result,
) error {
	builder := resource.NewBuilder(
		"pubsub_subscriptions",
		wf,
		cfg,
		q,
		gc.ListSubscriptions,
		func(wf *aw.Workflow, gps gc.PubSubSubscription) {
			ps := FromGCloudSubscription(&gps)
			resource.NewItem(wf, cfg, ps, svc.Icon())
		},
	)

	return builder.Build()
}
