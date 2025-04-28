// internal/art/art.go
package art

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	githubRawBaseURL = "https://raw.githubusercontent.com/JimmyTarson/Myimx/refs/heads/main/internal/art"
)

// GetArt returns the ASCII art for the given key, or an empty string if not found
func GetArt(key string) (string, bool) {
	safeKey := strings.ToLower(strings.TrimSpace(key))

	// First try to get art from local files
	art, found := getLocalArt(safeKey)
	if found {
		return art, true
	}

	// If not found locally, try to get from GitHub
	return getRemoteArt(safeKey)
}

func getLocalArt(key string) (string, bool) {
	execPath, err := os.Executable()
	if err != nil {
		return "", false
	}

	// Try multiple paths to find the art files
	artPaths := []string{
		filepath.Join("internal", "art", key+".md"),
		filepath.Join(filepath.Dir(execPath), "art", key+".md"),
		filepath.Join("/usr", "local", "share", "myimx", "art", key+".md"),
		filepath.Join(os.Getenv("APPDATA"), "myimx", "art", key+".md"),
		filepath.Join(os.Getenv("ProgramFiles"), "myimx", "art", key+".md"),
		filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Local", "Programs", "myimx", "art", key+".md"),
		filepath.Join(os.Getenv("HOME"), ".myimx", "art", key+".md"),
	}

	// Try each path
	for _, path := range artPaths {
		content, err := ioutil.ReadFile(path)
		if err == nil {
			return string(content), true
		}
	}

	return "", false
}

// getRemoteArt tries to get ASCII art from GitHub and saves it locally
func getRemoteArt(key string) (string, bool) {
	url := fmt.Sprintf("%s/%s.md", githubRawBaseURL, key)

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Error connecting to GitHub: %v\n", err)
		return "", false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", false
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return "", false
	}

	saveRemoteArt(key, content)

	return string(content), true
}

func saveRemoteArt(key string, content []byte) {
	var artDir string

	// Check if we can write to various locations
	possibleDirs := []string{
		filepath.Join(os.Getenv("HOME"), ".myimx", "art"),
		filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Local", "Programs", "myimx", "art"),
		filepath.Join(os.Getenv("APPDATA"), "myimx", "art"),
		func() string {
			execPath, err := os.Executable()
			if err != nil {
				return ""
			}
			return filepath.Join(filepath.Dir(execPath), "art")
		}(),
		"/usr/local/share/myimx/art",
	}

	// Find the first directory we can write to
	for _, dir := range possibleDirs {
		if dir == "" {
			continue
		}

		// Try to create directory if it doesn't exist
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err = os.MkdirAll(dir, 0755)
			if err != nil {
				continue
			}
		}

		// Check if we can write to this directory
		testFile := filepath.Join(dir, ".test")
		err := ioutil.WriteFile(testFile, []byte("test"), 0644)
		if err == nil {
			os.Remove(testFile)
			artDir = dir
			break
		}
	}

	// If we found a writable directory, save the file
	if artDir != "" {
		filePath := filepath.Join(artDir, key+".md")
		err := ioutil.WriteFile(filePath, content, 0644)
		if err != nil {
			fmt.Printf("Warning: Could not save art file locally: %v\n", err)
		} else {
			fmt.Printf("Downloaded new ASCII art '%s' from GitHub\n", key)
		}
	}
}

// GetAvailableArt returns a slice of all available art keys
func GetAvailableArt() []string {
	localKeys := getLocalArtKeys()

	remoteKeys := getRemoteArtKeys()

	allKeys := make(map[string]bool)

	for _, key := range localKeys {
		allKeys[key] = true
	}

	for _, key := range remoteKeys {
		allKeys[key] = true
	}

	result := make([]string, 0, len(allKeys))
	for key := range allKeys {
		result = append(result, key)
	}

	return result
}

// getLocalArtKeys returns art keys available locally
func getLocalArtKeys() []string {
	var keys []string

	artDirs := []string{
		filepath.Join("internal", "art"),
		func() string {
			execPath, err := os.Executable()
			if err != nil {
				return ""
			}
			return filepath.Join(filepath.Dir(execPath), "art")
		}(),
		filepath.Join("/usr", "local", "share", "myimx", "art"),
		filepath.Join(os.Getenv("APPDATA"), "myimx", "art"),
		filepath.Join(os.Getenv("ProgramFiles"), "myimx", "art"),
		filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Local", "Programs", "myimx", "art"),
		filepath.Join(os.Getenv("HOME"), ".myimx", "art"),
	}

	// Try to find art files in each directory
	for _, dir := range artDirs {
		if dir == "" {
			continue
		}

		files, err := ioutil.ReadDir(dir)
		if err == nil {
			for _, file := range files {
				if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") {
					// Strip the .md extension to get the key
					key := strings.TrimSuffix(file.Name(), ".md")
					keys = append(keys, key)
				}
			}
		}
	}

	return keys
}

// getRemoteArtKeys tries to get a list of available art from GitHub
func getRemoteArtKeys() []string {
	var keys []string
	return keys
}

// EnsureArtDirectoryExists creates the art directory if it doesn't exist
func EnsureArtDirectoryExists() error {
	var artDir string

	// Check various possible locations
	possibleDirs := []string{
		filepath.Join(os.Getenv("HOME"), ".myimx", "art"),
		filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Local", "Programs", "myimx", "art"),
		func() string {
			execPath, err := os.Executable()
			if err != nil {
				return ""
			}
			return filepath.Join(filepath.Dir(execPath), "art")
		}(),
	}

	// Find the first directory we can write to
	for _, dir := range possibleDirs {
		if dir == "" {
			continue
		}

		// Try to create directory if it doesn't exist
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err = os.MkdirAll(dir, 0755)
			if err == nil {
				artDir = dir
				break
			}
		} else {
			testFile := filepath.Join(dir, ".test")
			err := ioutil.WriteFile(testFile, []byte("test"), 0644)
			if err == nil {
				os.Remove(testFile)
				artDir = dir
				break
			}
		}
	}

	if artDir == "" {
		return fmt.Errorf("could not find or create a writable art directory")
	}

	return nil
}
