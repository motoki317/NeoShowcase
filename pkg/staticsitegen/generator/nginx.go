package generator

import (
	"github.com/traPtitech/neoshowcase/pkg/apiserver/api"
)

type Nginx struct {
	ArtifactsRootPath string
	GeneratedFilePath string
}

func (engine *Nginx) Generate(sites []api.StaticSite) error {
	return nil
}
