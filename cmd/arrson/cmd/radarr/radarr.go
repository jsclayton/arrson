package radarr

import (
	"fmt"
	gourl "net/url"

	"github.com/spf13/cobra"
	"golift.io/starr"
	"golift.io/starr/radarr"
)

func init() {
	RadarrCmd.PersistentFlags().StringVarP(&rawURL, "url", "u", "", "URL of the Radarr instance, including API token.")
	RadarrCmd.MarkPersistentFlagRequired("url")

	RadarrCmd.AddCommand(listCmd)
	RadarrCmd.AddCommand(deleteCmd)
}

var (
	rawURL string
)

var RadarrCmd = &cobra.Command{
	Use:     "radarr",
	Aliases: []string{"r"},
	Short:   "Radarr utilities",
}

func newClient() (*radarr.Radarr, error) {
	url, err := gourl.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	apiKey, hasKey := url.User.Password()
	if !hasKey {
		return nil, fmt.Errorf("no API key found in URL")
	}
	url.User = nil
	config := starr.New(apiKey, url.String(), starr.DefaultTimeout)
	return radarr.New(config), nil
}
