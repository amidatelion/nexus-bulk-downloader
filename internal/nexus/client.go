package nexus

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type File struct {
	Name       string `json:"name"`
	FileID     int    `json:"file_id"`
	Category   string `json:"category_name"`
	FileName   string `json:"file_name"`
}

type filesResponse struct {
	Files []File `json:"files"`
	// file_updates is ignored
}

func FetchFiles(apiKey, game, modID string) ([]File, error) {
	url := fmt.Sprintf("https://api.nexusmods.com/v1/games/%s/mods/%s/files.json", game, modID)

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
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed: %s\n%s", resp.Status, string(body))
	}

	var data filesResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	// 🔹 Filter: only include MAIN files
	mainFiles := []File{}
	for _, f := range data.Files {
		if f.Category == "MAIN" {
			mainFiles = append(mainFiles, f)
		}
	}

	return mainFiles, nil
}
