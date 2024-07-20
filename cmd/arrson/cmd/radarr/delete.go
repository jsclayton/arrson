package radarr

import (
	"fmt"
	"strconv"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func init() {
	deleteCmd.Flags().BoolVarP(&confirmDelete, "confirm", "y", false, "Confirm deletion")
}

var (
	confirmDelete bool
)

var deleteCmd = &cobra.Command{
	Use:   "delete [tmdb_id]",
	Short: "Delete a movie from Radarr",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}

		var tmdbID int64 = -1
		if intVal, isInt := strconv.Atoi(args[0]); isInt == nil {
			tmdbID = int64(intVal)
		}

		movies, err := client.GetMovieContext(cmd.Context(), tmdbID)
		if err != nil {
			return err
		}

		if len(movies) == 0 {
			return fmt.Errorf("Movie not found (TMDB ID: %d)", tmdbID)
		}

		if !confirmDelete {
			userConfirmed, err := pterm.DefaultInteractiveConfirm.
				Show(fmt.Sprintf("Delete movie \"%s\"", movies[0].Title))
			if err != nil {
				fmt.Println(err)
				return err
			}
			if !userConfirmed {
				fmt.Println("Aborted")
				return nil
			}
		}

		fmt.Println(fmt.Sprintf("Deleting movie: %s", movies[0].Title))
		err = client.DeleteMovieContext(cmd.Context(), movies[0].ID, true, true)
		if err != nil {
			fmt.Println(fmt.Sprintf("Failed to delete movie: %s", err))
		}

		fmt.Println(fmt.Sprintf("Deleted movie: %s", movies[0].Title))
		return nil
	},
}
