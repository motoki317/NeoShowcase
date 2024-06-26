package pbconvert

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/traPtitech/neoshowcase/pkg/domain"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/grpc/pb"
)

func ToPBSimpleCommit(c *domain.RepositoryCommit) *pb.SimpleCommit {
	return &pb.SimpleCommit{
		Hash:       c.Hash,
		AuthorName: c.Author.Name,
		CommitDate: timestamppb.New(c.Committer.Date),
		Message:    c.Message,
	}
}
