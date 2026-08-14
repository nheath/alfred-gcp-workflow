package artifactregistry

import (
	aw "github.com/deanishe/awgo"
	gc "github.com/nheath/alfred-gcp-workflow/gcloud"
	"github.com/nheath/alfred-gcp-workflow/parser"
	"github.com/nheath/alfred-gcp-workflow/services"
	"github.com/nheath/alfred-gcp-workflow/workflow/resource"
)

type RepositorySearcher struct{}

func (s *RepositorySearcher) Search(wf *aw.Workflow, svc *services.Service, cfg *gc.Config, q *parser.Result) error {
	builder := resource.NewBuilder(
		"artifact_registry_repositories",
		wf,
		cfg,
		q,
		gc.ListArtifactRepositories,
		func(wf *aw.Workflow, car gc.ArtifactRepository) {
			cr := FromGCloudRepository(&car)
			resource.NewItem(wf, cfg, cr, svc.Icon())
		},
	)

	return builder.Build()
}
