package k8s

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/nheath/alfred-gcp-workflow/gcloud"
	"github.com/nheath/alfred-gcp-workflow/parser"
	"github.com/nheath/alfred-gcp-workflow/services"
	"github.com/nheath/alfred-gcp-workflow/workflow/resource"
)

type ClusterSearcher struct{}

func (s *ClusterSearcher) Search(
	wf *aw.Workflow, svc *services.Service, cfg *gc.Config, pq *parser.Result,
) error {
	builder := resource.NewBuilder(
		"k8s_clusters",
		wf,
		cfg,
		pq,
		gc.ListK8sClusters,
		func(wf *aw.Workflow, gkec gc.K8sCluster) {
			kec := FromGCloudCluster(&gkec)
			resource.NewItem(wf, cfg, kec, svc.Icon())
		},
	)

	return builder.Build()
}
