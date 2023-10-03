package toprated

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/fschossler/tmdbcli/cmd/serie"
	"github.com/fschossler/tmdbcli/internal"
	"github.com/spf13/cobra"
)

var topRatedCmd = &cobra.Command{
	Use:   "toprated",
	Short: "Shows Top Rated series in TMDB database",
	Long:  `Shows Top Rated series in TMDB database.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := TopRated()
		if err != nil {
			log.Fatal(err)
		}
	},
}

type Root struct {
	Page    int
	Results []Results
}

type Results struct {
	Name        string  `json:"name"`
	VoteAverage float32 `json:"vote_average"`
	Overview    string  `json:"overview"`
}

func TopRated() error {

	BEARER_TOKEN := internal.ValidateBearerToken()

	url := "https://api.themoviedb.org/3/tv/top_rated?language=en-US&page=1"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+BEARER_TOKEN)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var serie Root

	jsonText := string(body)

	err = json.Unmarshal([]byte(jsonText), &serie)
	if err != nil {
		return err
	}

	for _, movie := range serie.Results {
		title := color.New(color.FgHiCyan)
		voteAverage := color.New(color.FgGreen)

		title.Print(movie.Name + ": ")
		voteAverage.Println(movie.VoteAverage)
		fmt.Println(movie.Overview)
		fmt.Println("")
	}

	return nil
}

func init() {
	serie.SerieCmd.AddCommand(topRatedCmd)
}
