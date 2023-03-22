package main

import (
	"strings"

	"github.com/traPtitech/neoshowcase/pkg/domain"
	"github.com/traPtitech/neoshowcase/pkg/domain/builder"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/admindb"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/backend/dockerimpl"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/dbmanager"
	"github.com/traPtitech/neoshowcase/pkg/interface/grpc"
	"github.com/traPtitech/neoshowcase/pkg/usecase"
)

const (
	ModeDocker = iota
	ModeK8s
)

type Config struct {
	Debug   bool                                  `mapstructure:"debug" yaml:"debug"`
	Mode    string                                `mapstructure:"mode" yaml:"mode"`
	Builder grpc.BuilderServiceClientConfig       `mapstructure:"builder" yaml:"builder"`
	SS      domain.StaticServerConnectivityConfig `mapstructure:"ss" yaml:"ss"`
	SSGen   grpc.StaticSiteServiceClientConfig    `mapstructure:"ssgen" yaml:"ssgen"`
	DB      admindb.Config                        `mapstructure:"db" yaml:"db"`
	MariaDB dbmanager.MariaDBConfig               `mapstructure:"mariadb" yaml:"mariadb"`
	MongoDB dbmanager.MongoDBConfig               `mapstructure:"mongodb" yaml:"mongodb"`
	Docker  struct {
		ConfDir string `mapstructure:"confDir" yaml:"confDir"`
	} `mapstructure:"docker" yaml:"docker"`
	GRPC struct {
		Port int `mapstructure:"port" yaml:"port"`
	} `mapstructure:"grpc" yaml:"grpc"`
	HTTP struct {
		Debug bool `mapstructure:"debug" yaml:"debug"`
		Port  int  `mapstructure:"port" yaml:"port"`
	} `mapstructure:"http" yaml:"http"`
	Repository struct {
		CacheDir string `mapstructure:"cacheDir" yaml:"cacheDir"`
	} `mapstructure:"repository" yaml:"repository"`
	Image struct {
		Registry   builder.DockerImageRegistryString   `mapstructure:"registry" yaml:"registry"`
		NamePrefix builder.DockerImageNamePrefixString `mapstructure:"namePrefix" yaml:"namePrefix"`
	} `mapstructure:"image" yaml:"image"`
}

func (c *Config) GetMode() int {
	switch strings.ToLower(c.Mode) {
	case "k8s", "kubernetes":
		return ModeK8s
	case "docker":
		return ModeDocker
	default:
		return ModeDocker
	}
}

func provideIngressConfDirPath(c Config) dockerimpl.IngressConfDirPath {
	return dockerimpl.IngressConfDirPath(c.Docker.ConfDir)
}

func provideImageRegistry(c Config) builder.DockerImageRegistryString {
	return c.Image.Registry
}

func provideImagePrefix(c Config) builder.DockerImageNamePrefixString {
	return c.Image.NamePrefix
}

func provideRepositoryFetcherCacheDir(c Config) usecase.RepositoryFetcherCacheDir {
	return usecase.RepositoryFetcherCacheDir(c.Repository.CacheDir)
}
