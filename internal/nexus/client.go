package nexus

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	Name     string `json:"name"`
	FileID   int    `json:"file_id"`
	Category string `json:"category_name"`
	FileName string `json:"file_name"`
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

type DownloadLink struct {
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	URI       string `json:"URI"`
}

func GetDownloadLink(apiKey, game, modID string, fileID int) (string, error) {
	url := fmt.Sprintf("https://api.nexusmods.com/v1/games/%s/mods/%s/files/%d/download_link.json", game, modID, fileID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("apikey", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("API request failed: %s\n%s", resp.Status, string(body))
	}

	var links []DownloadLink
	if err := json.NewDecoder(resp.Body).Decode(&links); err != nil {
		return "", err
	}

	if len(links) == 0 {
		return "", fmt.Errorf("no download links returned")
	}

	return links[0].URI, nil
}

func DownloadFile(uri, downloadDir string) (string, error) {
	// Extract filename from the URI (strip query params)
	parts := strings.Split(uri, "/")
	lastPart := parts[len(parts)-1]
	fileName := strings.SplitN(lastPart, "?", 2)[0]

	// Build full path
	outPath := filepath.Join(downloadDir, fileName)

	// Make sure the directory exists
	if err := os.MkdirAll(downloadDir, 0o755); err != nil {
		return "", fmt.Errorf("failed to create download directory: %w", err)
	}

	// GET request to fetch the file
	resp, err := http.Get(uri)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download file: %s", resp.Status)
	}

	// Create file on disk
	out, err := os.Create(outPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	return outPath, nil
}
