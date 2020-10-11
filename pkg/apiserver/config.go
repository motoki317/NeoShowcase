package apiserver

import (
	"github.com/traPtitech/neoshowcase/pkg/common"
)

type Config struct {
	Builder common.GRPCClientConfig `mapstructure:"builder" yaml:"builder"`
	SSGen   common.GRPCClientConfig `mapstructure:"ssgen" yaml:"ssgen"`
	GRPC    common.GRPCConfig       `mapstructure:"grpc" yaml:"grpc"`
	DB      common.DBConfig         `mapstructure:"db" yaml:"db"`
}
