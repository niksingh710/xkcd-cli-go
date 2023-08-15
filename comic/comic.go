package comic

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Comic struct {
	Title      string `json:"title"`
	Num        int    `json:"num"`
	Year       string `json:"year"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Day        string `json:"day"`
}

const URLPrefix = "https://xkcd.com/"

var db = make(map[int]*Comic)

func GetURL(id int) string {
	return URLPrefix + strconv.Itoa(id) + "/info.0.json"
}

func GetComic(id int) (*Comic, error) {
	if comic, ok := db[id]; !ok {
		url := GetURL(id)
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("No Comic with This id is available")
		}
		if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
			return nil, fmt.Errorf("Invalid Json Data")
		}
		db[id] = comic
	}
	return db[id], nil
}
