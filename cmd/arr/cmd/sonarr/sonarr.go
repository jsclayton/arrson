package sonarr

import "github.com/spf13/cobra"

var SonarrCmd = &cobra.Command{
	Use:     "sonarr",
	Aliases: []string{"s"},
	Short:   "Sonarr utilities",
}
