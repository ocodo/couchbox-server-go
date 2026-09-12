package filesystem

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestGetFilteredFilesSkipsMissingRoot(t *testing.T) {
	root := t.TempDir()

	mustWriteFile(t, filepath.Join(root, "movie.mkv"))

	missing := filepath.Join(root, "does-not-exist")

	got := GetFilteredFiles([]string{missing, root})

	want := []string{"movie.mkv"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("GetFilteredFiles() = %#v, want %#v", got, want)
	}
}

func TestGetFilteredFiles(t *testing.T) {
	root := t.TempDir()

	mustMkdirAll(t, filepath.Join(root, "movies", "nested"))

	mustWriteFile(t, filepath.Join(root, "movie.mkv"))
	mustWriteFile(t, filepath.Join(root, "movies", "nested", "other.mkv"))

	got := GetFilteredFiles([]string{root})

	want := []string{
		"movie.mkv",
		filepath.Join("movies", "nested", "other.mkv"),
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("GetFilteredFiles() = %#v, want %#v", got, want)
	}
}

func TestGetFilteredFilesWithFileRoot(t *testing.T) {
	root := t.TempDir()
	file := filepath.Join(root, "movie.mkv")

	mustWriteFile(t, file)

	got := GetFilteredFiles([]string{file})

	want := []string{"movie.mkv"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("GetFilteredFiles() = %#v, want %#v", got, want)
	}
}

func mustMkdirAll(t *testing.T, path string) {
	t.Helper()

	if err := os.MkdirAll(path, 0o755); err != nil {
		t.Fatalf("MkdirAll(%q): %v", path, err)
	}
}

func mustWriteFile(t *testing.T, path string) {
	t.Helper()

	if err := os.WriteFile(path, nil, 0o644); err != nil {
		t.Fatalf("WriteFile(%q): %v", path, err)
	}
}

func TestFindFileMatches(t *testing.T) {
	root := t.TempDir()

	mustMkdirAll(t, filepath.Join(root, "movies", "nested"))
	mustWriteFile(t, filepath.Join(root, "movies", "nested", "movie.mkv"))
	mustWriteFile(t, filepath.Join(root, "movies", "other.mkv"))

	got := FindFileMatches(
		[]string{root},
		filepath.Join("movies", "nested", "movie.mkv"),
	)

	want := []string{root}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("FindFileMatches() = %#v, want %#v", got, want)
	}
}
