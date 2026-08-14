package pubsub

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/nheath/alfred-gcp-workflow/gcloud"
	"github.com/nheath/alfred-gcp-workflow/parser"
	"github.com/nheath/alfred-gcp-workflow/services"
	"github.com/nheath/alfred-gcp-workflow/workflow/resource"
)

type TopicSearcher struct{}

func (s *TopicSearcher) Search(
	wf *aw.Workflow, svc *services.Service, cfg *gc.Config, q *parser.Result,
) error {
	builder := resource.NewBuilder(
		"pubsub_topics",
		wf,
		cfg,
		q,
		gc.ListTopics,
		func(wf *aw.Workflow, gpt gc.PubSubTopic) {
			pt := FromGCloudTopic(&gpt)
			resource.NewItem(wf, cfg, pt, svc.Icon())
		},
	)

	return builder.Build()
}
