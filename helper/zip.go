package helper

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/schollz/progressbar/v3"
)

func Unzip(src string, dest string) error {

	reader, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer reader.Close()
	len := int64(len(reader.File))
	bar := progressbar.Default(len, "Unzipped")

	for _, file := range reader.File {
		fpath := filepath.Join(dest, file.Name)

		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("%s: Invalid file path", fpath)
		}

		if file.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		rc, err := file.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
		bar.Add(1)
	}
	return nil
}
