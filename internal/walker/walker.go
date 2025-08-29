package walker

import (
	"io/fs"
	"path/filepath"
	"strings"
)

var skipDirs = map[string]struct{}{
	"node_modules": {},
	".git":         {},
	"dist":         {},
	"build":        {},
}

var allowedExts = map[string]struct{}{
	".html": {},
	".ts":   {},
	".qml":  {},
	".cpp":  {},
	".h":    {},
}

func Walk(root string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			if _, skip := skipDirs[d.Name()]; skip {
				return filepath.SkipDir
			}
			return nil
		}

		ext := strings.ToLower(filepath.Ext(d.Name()))
		if _, ok := allowedExts[ext]; !ok {
			return nil
		}
		files = append(files, path)
		return nil
	})

	return files, err
}
