package cmd

import (
	"context"

	"github.com/jsclayton/arr-utils/cmd/arr/cmd/radarr"
	"github.com/jsclayton/arr-utils/cmd/arr/cmd/sonarr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(radarr.RadarrCmd)
	rootCmd.AddCommand(sonarr.SonarrCmd)
}

var rootCmd = &cobra.Command{
	Use:   "arr",
	Short: "A collection of utilities for working with arr apps",
}

func Execute(ctx context.Context) error {
	return rootCmd.ExecuteContext(ctx)
}
