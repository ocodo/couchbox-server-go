package filesystem

import (
	"os"
	"path/filepath"
	"sort"
)

func GetFilteredFiles(roots []string) []string {
	var files []string

	for _, root := range roots {
		info, err := os.Stat(root)
		if err != nil {
			continue
		}

		if !info.IsDir() {
			files = append(files, info.Name())
			continue
		}

		err = filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if entry.IsDir() {
				return nil
			}

			relative, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}

			files = append(files, relative)
			return nil
		})

		if err != nil {
			continue
		}
	}

	sort.Strings(files)

	return files
}

func FindFileMatches(roots []string, relativePath string) []string {
	var matches []string

	for _, root := range roots {
		path := filepath.Join(root, relativePath)

		info, err := os.Stat(path)
		if err != nil || info.IsDir() {
			continue
		}

		matches = append(matches, root)
	}

	return matches
}
