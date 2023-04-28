package builder

import "github.com/traPtitech/neoshowcase/pkg/domain"

type Task struct {
	ApplicationID string
	BuildID       string
	RepositoryID  string
	Commit        string
	ImageName     string
	ImageTag      string
	BuildConfig   domain.BuildConfig
}

func (t *Task) DestImage() string {
	return t.ImageName + ":" + t.ImageTag
}
