/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

const DOWNLOAD_PATH = "https://huggingface.co/datasets/cbdb/cbdb-sqlite/resolve/main/latest.zip"

func downloadZip(dir string) error {
	req, err := http.NewRequest("GET", DOWNLOAD_PATH, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	filePath := filepath.Join(dir, "latest.zip")
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)
	io.Copy(io.MultiWriter(f, bar), resp.Body)
	return nil
}

func createCacheDir() (string, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	projectCacheDir := filepath.Join(cacheDir, "cbdb-tool")
	return projectCacheDir, os.MkdirAll(projectCacheDir, 0777)
}

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		if f.FileInfo().IsDir() {
			continue
		}

		if strings.HasSuffix(f.FileInfo().Name(), ".sqlite3") {
			fpath := filepath.Join(dest, "cbdb.sqlite3")
			if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
				return fmt.Errorf("invalid file path: %s", fpath)
			}

			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}

			dstFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}

			srcFile, err := f.Open()
			if err != nil {
				return err
			}

			io.Copy(dstFile, srcFile)
			dstFile.Close()
			srcFile.Close()
		}

	}
	return nil
}

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download the sqlite database of CBDB",
	Annotations: map[string]string{
		"skip-db": "true",
	},
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := createCacheDir()
		if err != nil {
			log.Fatal(err)
		}

		err = downloadZip(dir)
		if err != nil {
			log.Fatal(err)
		}

		err = unzip(filepath.Join(dir, "latest.zip"), dir)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}
