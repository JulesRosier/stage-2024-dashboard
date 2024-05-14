package view

import (
	"Stage-2024-dashboard/pkg/helper"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

var (
	excludeExts = []string{".woff2"}
	StaticMap   map[string]string
)

func hashFile(file fs.File) (string, error) {
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	sum := hash.Sum(nil)

	return hex.EncodeToString(sum)[:12], nil
}

func hashFiles(rootDir string, excludeExts []string) (map[string]string, error) {
	filesDetails := map[string]string{}

	err := filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			ext := filepath.Ext(path)
			for _, e := range excludeExts {
				if ext == e {
					return nil
				}
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			hash, err := hashFile(file)
			if err != nil {
				return err
			}

			fileName := strings.TrimSuffix(d.Name(), ext)

			newFilename := fmt.Sprintf("%s-%s%s", fileName, hash, ext)
			basePath := strings.TrimSuffix(path, d.Name())
			basePath = strings.ReplaceAll(basePath, "public", "/static")

			p := "/" + strings.ReplaceAll(strings.ReplaceAll(path, "public", "/static"), "\\", "/")
			filesDetails[p] = "/" + strings.ReplaceAll(basePath+newFilename, "\\", "/")
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return filesDetails, nil
}

func HashPublicFS() map[string]string {
	slog.Info("Hashing static files")
	files, err := hashFiles("./static", excludeExts)
	helper.MaybeDieErr(err)

	for k, v := range files {
		slog.Debug("file hashed", "old", k, "new", v)
	}
	StaticMap = files
	swappedMap := make(map[string]string)
	for k, v := range files {
		swappedMap[v] = k
	}
	return swappedMap
}
