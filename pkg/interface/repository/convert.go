package repository

import (
	"github.com/traPtitech/neoshowcase/pkg/domain"
	"github.com/traPtitech/neoshowcase/pkg/domain/builder"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/admindb/models"
	"github.com/traPtitech/neoshowcase/pkg/util/optional"
)

func toDomainRepository(repo *models.Repository) domain.Repository {
	return domain.Repository{
		ID:  repo.ID,
		URL: repo.URL,
	}
}

func toDomainApplication(app *models.Application) *domain.Application {
	ret := &domain.Application{
		ID:            app.ID,
		Repository:    toDomainRepository(app.R.Repository),
		BranchName:    app.BranchName,
		BuildType:     builder.BuildTypeFromString(app.BuildType),
		State:         domain.ApplicationStateFromString(app.State),
		CurrentCommit: app.CurrentCommit,
		WantCommit:    app.WantCommit,
	}
	if app.R.Website != nil {
		ret.Website = optional.From(*toDomainWebsite(app.R.Website))
	}
	return ret
}

func toDomainBuild(build *models.Build) *domain.Build {
	ret := &domain.Build{
		ID:            build.ID,
		Commit:        build.Commit,
		Status:        builder.BuildStatusFromString(build.Status),
		ApplicationID: build.ApplicationID,
		StartedAt:     build.StartedAt,
		FinishedAt:    optional.New(build.FinishedAt.Time, build.FinishedAt.Valid),
	}
	if build.R != nil && build.R.Artifact != nil {
		artifact := build.R.Artifact
		ret.Artifact = optional.From(domain.Artifact{
			ID:        artifact.ID,
			Size:      artifact.Size,
			CreatedAt: artifact.CreatedAt,
		})
	}
	return ret
}

func toDomainEnvironment(env *models.Environment) *domain.Environment {
	return &domain.Environment{
		ID:            env.ID,
		ApplicationID: env.ApplicationID,
		Key:           env.Key,
		Value:         env.Value,
	}
}

func toDomainWebsite(website *models.Website) *domain.Website {
	return &domain.Website{
		ID:   website.ID,
		FQDN: website.FQDN,
		Port: website.HTTPPort,
	}
}
