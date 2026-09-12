package config

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestLoadExampleConfig(t *testing.T) {
	path := filepath.Join("..", "..", "testdata", "config.example.yaml")

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	var got Config

	if err := Load(data, &got); err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	want := Config{
		Content: ContentConfig{
			Movies: []string{
				"./testdata/movies",
				"./testdata/movies-alt",
			},
			TV: []string{
				"./testdata/tv",
			},
		},
		Data: DataConfig{
			Blank: PosterData{
				Posters: PathConfig{
					Path: "./testdata/posters/blank",
				},
			},
			Movies: MovieData{
				Posters: PathConfig{
					Path: "./testdata/posters/movies",
				},
			},
			TV: TVData{
				Posters: PathConfig{
					Path: "./testdata/posters/tv",
				},
				Backdrops: PathConfig{
					Path: "./testdata/posters/backdrops",
				},
			},
		},
		VideoExts: []string{
			".avi",
			".mov",
			".mkv",
			".mp4",
			".webm",
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Load() = %#v, want %#v", got, want)
	}
}

func TestLoadConfigFile(t *testing.T) {
	path := filepath.Join("..", "..", "testdata", "config.example.yaml")

	got, err := LoadFile(path)
	if err != nil {
		t.Fatalf("LoadFile() error = %v", err)
	}

	if len(got.Content.Movies) != 2 {
		t.Fatalf("Content.Movies length = %d, want 2", len(got.Content.Movies))
	}

	if len(got.Content.TV) != 1 {
		t.Fatalf("Content.TV length = %d, want 1", len(got.Content.TV))
	}

	if got.Data.Movies.Posters.Path == "" {
		t.Fatal("Data.Movies.Posters.Path is empty")
	}
}

func TestLoadFileMissing(t *testing.T) {
	_, err := LoadFile(filepath.Join("..", "..", "testdata", "does-not-exist.yaml"))
	if err == nil {
		t.Fatal("LoadFile() error = nil, want an error")
	}
}

func TestLoadRejectsMalformedYAML(t *testing.T) {
	data := []byte(`
content:
  movies:
    - ./testdata/movies
  tv:
    - [broken
`)

	var got Config

	if err := Load(data, &got); err == nil {
		t.Fatal("Load() error = nil, want an error")
	}
}

func TestLoadRejectsUnknownFields(t *testing.T) {
	data := []byte(`
content:
  movies:
    - ./testdata/movies
  tv:
    - ./testdata/tv

data:
  blank:
    posters:
      path: ./testdata/posters/blank

  movies:
    posters:
      path: ./testdata/posters/movies

  tv:
    posters:
      path: ./testdata/posters/tv
    backdrops:
      path: ./testdata/posters/backdrops

  unexpected: true
`)

	var got Config

	if err := Load(data, &got); err == nil {
		t.Fatal("Load() error = nil, want an error for unknown field")
	}
}

func TestLoadExampleConfigVideoExts(t *testing.T) {
	path := filepath.Join("..", "..", "testdata", "config.example.yaml")

	got, err := LoadFile(path)
	if err != nil {
		t.Fatalf("LoadFile() error = %v", err)
	}

	want := []string{
		".avi",
		".mov",
		".mkv",
		".mp4",
		".webm",
	}

	if !reflect.DeepEqual(got.VideoExts, want) {
		t.Fatalf("VideoExts = %#v, want %#v", got.VideoExts, want)
	}
}

func TestLoadRejectsEmptyConfig(t *testing.T) {
	data := []byte(`{}`)

	var got Config

	if err := Load(data, &got); err == nil {
		t.Fatal("Load() error = nil, want an error for empty config")
	}
}

