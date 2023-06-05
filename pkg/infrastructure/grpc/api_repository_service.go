package grpc

import (
	"context"

	"github.com/bufbuild/connect-go"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/traPtitech/neoshowcase/pkg/infrastructure/grpc/pb"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/grpc/pbconvert"
	"github.com/traPtitech/neoshowcase/pkg/usecase"
	"github.com/traPtitech/neoshowcase/pkg/util/ds"
	"github.com/traPtitech/neoshowcase/pkg/util/optional"
)

func (s *APIService) CreateRepository(ctx context.Context, req *connect.Request[pb.CreateRepositoryRequest]) (*connect.Response[pb.Repository], error) {
	msg := req.Msg
	repo, err := s.svc.CreateRepository(ctx,
		msg.Name,
		msg.Url,
		pbconvert.FromPBRepositoryAuth(msg.Auth),
	)
	if err != nil {
		return nil, handleUseCaseError(err)
	}
	res := connect.NewResponse(pbconvert.ToPBRepository(repo))
	return res, nil
}

func (s *APIService) GetRepositories(ctx context.Context, req *connect.Request[pb.GetRepositoriesRequest]) (*connect.Response[pb.GetRepositoriesResponse], error) {
	repositories, err := s.svc.GetRepositories(ctx, pbconvert.RepoScopeMapper.FromMust(req.Msg.Scope))
	if err != nil {
		return nil, handleUseCaseError(err)
	}
	res := connect.NewResponse(&pb.GetRepositoriesResponse{
		Repositories: ds.Map(repositories, pbconvert.ToPBRepository),
	})
	return res, nil
}

func (s *APIService) GetRepository(ctx context.Context, req *connect.Request[pb.RepositoryIdRequest]) (*connect.Response[pb.Repository], error) {
	repository, err := s.svc.GetRepository(ctx, req.Msg.RepositoryId)
	if err != nil {
		return nil, handleUseCaseError(err)
	}
	res := connect.NewResponse(pbconvert.ToPBRepository(repository))
	return res, nil
}

func (s *APIService) UpdateRepository(ctx context.Context, req *connect.Request[pb.UpdateRepositoryRequest]) (*connect.Response[emptypb.Empty], error) {
	msg := req.Msg
	args := &usecase.UpdateRepositoryArgs{
		Name:     optional.FromNonZero(msg.Name),
		URL:      optional.FromNonZero(msg.Url),
		Auth:     optional.Map(optional.FromNonZero(msg.Auth), pbconvert.FromPBRepositoryAuth),
		OwnerIDs: optional.FromNonZeroSlice(msg.OwnerIds),
	}
	err := s.svc.UpdateRepository(ctx, msg.Id, args)
	if err != nil {
		return nil, handleUseCaseError(err)
	}
	res := connect.NewResponse(&emptypb.Empty{})
	return res, nil
}

func (s *APIService) RefreshRepository(ctx context.Context, req *connect.Request[pb.RepositoryIdRequest]) (*connect.Response[emptypb.Empty], error) {
	err := s.svc.RefreshRepository(ctx, req.Msg.RepositoryId)
	if err != nil {
		return nil, handleUseCaseError(err)
	}
	res := connect.NewResponse(&emptypb.Empty{})
	return res, nil
}

func (s *APIService) DeleteRepository(ctx context.Context, req *connect.Request[pb.RepositoryIdRequest]) (*connect.Response[emptypb.Empty], error) {
	err := s.svc.DeleteRepository(ctx, req.Msg.RepositoryId)
	if err != nil {
		return nil, handleUseCaseError(err)
	}
	res := connect.NewResponse(&emptypb.Empty{})
	return res, nil
}
