package config

import (
	"bytes"
	"errors"
	"os"
	
	"gopkg.in/yaml.v3"
)

type Config struct {
	Content   ContentConfig `yaml:"content"`
	Data      DataConfig    `yaml:"data"`
	VideoExts []string      `yaml:"video_exts"`
}

type ContentConfig struct {
	Movies []string `yaml:"movies"`
	TV     []string `yaml:"tv"`
}

type DataConfig struct {
	Blank  PosterData `yaml:"blank"`
	Movies MovieData  `yaml:"movies"`
	TV     TVData     `yaml:"tv"`
}

type PosterData struct {
	Posters PathConfig `yaml:"posters"`
}

type MovieData struct {
	Posters PathConfig `yaml:"posters"`
}

type TVData struct {
	Posters   PathConfig `yaml:"posters"`
	Backdrops PathConfig `yaml:"backdrops"`
}

type PathConfig struct {
	Path string `yaml:"path"`
}

func Load(data []byte, dst *Config) error {
	decoder := yaml.NewDecoder(bytes.NewReader(data))
	decoder.KnownFields(true)

	if err := decoder.Decode(dst); err != nil {
		return err
	}

	return dst.Validate()
}

func LoadFile(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var config Config

	if err := Load(data, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func (c Config) Validate() error {
	if len(c.Content.Movies) == 0 && len(c.Content.TV) == 0 {
		return errors.New("config: no content paths configured")
	}

	return nil
}

