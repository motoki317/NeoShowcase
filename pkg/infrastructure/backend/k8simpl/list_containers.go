package k8simpl

import (
	"context"

	"github.com/friendsofgo/errors"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/traPtitech/neoshowcase/pkg/domain"
)

func (b *k8sBackend) GetContainer(ctx context.Context, appID string) (*domain.Container, error) {
	list, err := b.client.CoreV1().Pods(appNamespace).List(ctx, metav1.ListOptions{
		LabelSelector: labelSelector(appID),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch pods")
	}
	if len(list.Items) == 0 {
		return nil, domain.ErrContainerNotFound
	}

	item := list.Items[0]
	return &domain.Container{
		ApplicationID: appID,
		State:         getContainerState(item.Status),
	}, nil
}

func (b *k8sBackend) ListContainers(ctx context.Context) ([]domain.Container, error) {
	list, err := b.client.CoreV1().Pods(appNamespace).List(ctx, metav1.ListOptions{
		LabelSelector: allSelector(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch pods")
	}

	var result []domain.Container
	for _, item := range list.Items {
		result = append(result, domain.Container{
			ApplicationID: item.Labels[appIDLabel],
			State:         getContainerState(item.Status),
		})
	}
	return result, nil
}

func getContainerState(status apiv1.PodStatus) domain.ContainerState {
	switch status.Phase {
	case apiv1.PodPending:
		return domain.ContainerStateRestarting
	case apiv1.PodRunning:
		return domain.ContainerStateRunning
	case apiv1.PodFailed:
		return domain.ContainerStateStopped
	case apiv1.PodSucceeded:
		return domain.ContainerStateStopped
	case apiv1.PodUnknown:
		return domain.ContainerStateOther
	default:
		return domain.ContainerStateOther
	}
}
