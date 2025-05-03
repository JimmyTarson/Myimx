// internal/art/art.go
package art

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func GetArt(key string) (string, bool) {
	safeKey := strings.ToLower(strings.TrimSpace(key))

	execPath, err := os.Executable()
	if err != nil {
		return "", false
	}

	artPaths := []string{
		filepath.Join("internal", "art", safeKey+".md"),
		filepath.Join(filepath.Dir(execPath), "art", safeKey+".md"),
		filepath.Join("/usr", "local", "share", "myimx", "art", safeKey+".md"),
		filepath.Join(os.Getenv("APPDATA"), "myimx", "art", safeKey+".md"),
		filepath.Join(os.Getenv("ProgramFiles"), "myimx", "art", safeKey+".md"),
	}

	for _, path := range artPaths {
		content, err := ioutil.ReadFile(path)
		if err == nil {
			return string(content), true
		}
	}

	return "", false
}

func GetAvailableArt() []string {
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
	}

	for _, dir := range artDirs {
		if dir == "" {
			continue
		}

		files, err := ioutil.ReadDir(dir)
		if err == nil {
			for _, file := range files {
				if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") {
					key := strings.TrimSuffix(file.Name(), ".md")
					keys = append(keys, key)
				}
			}
			if len(keys) > 0 {
				break
			}
		}
	}

	return keys
}

func EnsureArtDirectoryExists() error {
	execPath, err := os.Executable()
	if err != nil {
		return err
	}

	artDir := filepath.Join(filepath.Dir(execPath), "art")

	if _, err := os.Stat(artDir); os.IsNotExist(err) {
		err = os.MkdirAll(artDir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create art directory: %w", err)
		}
	}

	return nil
}
