package radarr

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List movies in Radarr",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}

		movies, err := client.GetMovieContext(cmd.Context(), 0)
		if err != nil {
			return err
		}
		for _, movie := range movies {
			line, err := json.Marshal(movie)
			if err != nil {
				return err
			}
			fmt.Println(string(line))
		}
		return nil
	},
}
