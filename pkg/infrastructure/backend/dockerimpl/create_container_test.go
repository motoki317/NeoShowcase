package dockerimpl

import (
	"context"
	"testing"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/leandro-lugaresi/hub"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/traPtitech/neoshowcase/pkg/domain"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/eventbus"
)

func TestDockerBackend_CreateContainer(t *testing.T) {
	m, c := prepareManager(t, eventbus.NewLocal(hub.New()))

	t.Run("存在しないイメージを指定", func(t *testing.T) {
		t.Parallel()
		err := m.CreateContainer(context.Background(), domain.ContainerCreateArgs{
			ImageName: "notfoundimage",
		})
		assert.Error(t, err)
	})

	t.Run("コンテナを正常に作成して起動", func(t *testing.T) {
		t.Parallel()
		image := "tianon/sleeping-beauty"
		appID := "pjpjpjoijion"

		err := m.CreateContainer(context.Background(), domain.ContainerCreateArgs{
			ImageName:     image,
			ApplicationID: appID,
		})
		if assert.NoError(t, err) {
			cont, err := c.InspectContainerWithOptions(docker.InspectContainerOptions{
				ID: containerName(appID),
			})
			require.NoError(t, err)

			assert.Equal(t, cont.Config.Image, image+":latest")
			assert.Equal(t, cont.Config.Labels[appContainerLabel], "true")
			assert.Equal(t, cont.Config.Labels[appContainerApplicationIDLabel], appID)

			require.NoError(t, c.RemoveContainer(docker.RemoveContainerOptions{
				ID:            containerName(appID),
				RemoveVolumes: true,
				Force:         true,
			}))
		}
	})

	t.Run("コンテナを正常に作成して起動 (Recreate)", func(t *testing.T) {
		t.Parallel()
		image := "tianon/sleeping-beauty"
		appID := "pij0bij90j20"

		err := m.CreateContainer(context.Background(), domain.ContainerCreateArgs{
			ImageName:     image,
			ApplicationID: appID,
			Recreate:      true,
		})
		require.NoError(t, err)

		err = m.CreateContainer(context.Background(), domain.ContainerCreateArgs{
			ImageName:     image,
			ApplicationID: appID,
			Recreate:      true,
		})
		if assert.NoError(t, err) {
			cont, err := c.InspectContainerWithOptions(docker.InspectContainerOptions{
				ID: containerName(appID),
			})
			require.NoError(t, err)

			assert.Equal(t, cont.Config.Image, image+":latest")
			assert.Equal(t, cont.Config.Labels[appContainerLabel], "true")
			assert.Equal(t, cont.Config.Labels[appContainerApplicationIDLabel], appID)

			require.NoError(t, c.RemoveContainer(docker.RemoveContainerOptions{
				ID:            containerName(appID),
				RemoveVolumes: true,
				Force:         true,
			}))
		}
	})
}
