package nexus

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type File struct {
	FileID   int    `json:"file_id"`
	Name     string `json:"name"`
	Category string `json:"category_name"`
	FileName string `json:"file_name"`
}

type apiResponse struct {
	Files []File `json:"files"`
}

func FetchFiles(apiKey, gameSlug, modID string) ([]File, error) {
	url := fmt.Sprintf("https://api.nexusmods.com/v1/games/%s/mods/%s/files.json", gameSlug, modID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("apikey", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	var parsed apiResponse
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return nil, err
	}
	return parsed.Files, nil
}
