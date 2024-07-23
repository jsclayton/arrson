package radarr

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"golift.io/starr/radarr"
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
		if len(args) == 1 && args[0] == "-" {
			stdin := cmd.InOrStdin()
			scanner := bufio.NewScanner(stdin)
			for scanner.Scan() {
				var movie *radarr.Movie
				err := json.Unmarshal(scanner.Bytes(), &movie)
				if err != nil {
					return err
				}
				if movie == nil {
					return fmt.Errorf("Failed to unmarshal movie")
				}
				_, err = deleteMovie(cmd.Context(), movie.TmdbID)
				if err != nil {
					return err
				}
			}
			return nil
		}

		var tmdbID int64 = -1
		if intVal, isInt := strconv.Atoi(args[0]); isInt == nil {
			tmdbID = int64(intVal)
		}

		_, err := deleteMovie(cmd.Context(), tmdbID)
		if err != nil {
			return err
		}

		return nil
	},
}

func deleteMovie(ctx context.Context, tmdbID int64) (bool, error) {
	movies, err := client.GetMovieContext(ctx, tmdbID)
	if err != nil {
		return false, err
	}

	if len(movies) == 0 {
		return false, nil
	}

	if !confirmDelete {
		userConfirmed, err := pterm.DefaultInteractiveConfirm.
			Show(fmt.Sprintf("Delete movie \"%s\"", movies[0].Title))
		if err != nil {
			fmt.Println(err)
			return false, err
		}
		if !userConfirmed {
			fmt.Println("Aborted")
			return false, nil
		}
	}

	fmt.Println(fmt.Sprintf("Deleting movie: %s", movies[0].Title))
	err = client.DeleteMovieContext(ctx, movies[0].ID, true, true)
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to delete movie: %s", err))
	}

	fmt.Println(fmt.Sprintf("Deleted movie: %s", movies[0].Title))
	return true, nil
}
