package file_ops

import (
	"os"
	"path/filepath"
)

func RecordNewDir(dirName string) error {
	path := os.Getenv("GWT_NEW_DIR_FILE")

	if path == "" {
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return err
	}

	if err := os.WriteFile(path, []byte(dirName), 0o644); err != nil {
		return err
	}

	return nil
}
