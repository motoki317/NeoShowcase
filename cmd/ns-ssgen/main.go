package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/traPtitech/neoshowcase/pkg/cliutil"
	"github.com/traPtitech/neoshowcase/pkg/staticsitegen"
	"time"
)

var (
	version  = "UNKNOWN"
	revision = "UNKNOWN"
)

var (
	configFilePath string
	c              staticsitegen.Config
)

var rootCommand = &cobra.Command{
	Use:     "ns-ssgen",
	Short:   "NeoShowcase StaticSiteGenerator",
	Version: fmt.Sprintf("%s (%s)", version, revision),
}

func runCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "run",
		RunE: func(cmd *cobra.Command, args []string) error {
			service, err := staticsitegen.New(c)
			if err != nil {
				return err
			}

			go func() {
				err := service.Start(context.Background())
				if err != nil {
					log.Fatal(err)
				}
			}()

			log.Info("NeoShowcase StaticSiteGenerator started")
			cliutil.WaitSIGINT()
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			return service.Shutdown(ctx)
		},
	}
	return cmd
}

func init() {
	cobra.OnInitialize(cliutil.CobraOnInitializeFunc(&configFilePath, "NS_SSGEN", &c))

	rootCommand.AddCommand(
		runCommand(),
		cliutil.PrintConfCommand(&c),
	)

	flags := rootCommand.PersistentFlags()
	flags.StringVarP(&configFilePath, "config", "c", "", "config file path")
	cliutil.SetupDebugFlag(flags)

	viper.SetDefault("server", "caddy")
	viper.SetDefault("artifactsRoot", "/srv/artifacts")
	viper.SetDefault("generatedConfDir", "/srv/config")
	viper.SetDefault("caddy.adminEndpoint", "http://localhost:2019")
	viper.SetDefault("grpc.port", 10000)
	viper.SetDefault("db.host", "127.0.0.1")
	viper.SetDefault("db.port", 3306)
	viper.SetDefault("db.username", "root")
	viper.SetDefault("db.password", "password")
	viper.SetDefault("db.database", "neoshowcase")
	viper.SetDefault("db.connection.maxOpen", 0)
	viper.SetDefault("db.connection.maxIdle", 2)
	viper.SetDefault("db.connection.lifetime", 0)
}

func main() {
	if err := rootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}