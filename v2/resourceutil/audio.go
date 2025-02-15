package resourceutil

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

type AudioStream interface {
	io.ReadSeeker
	Length() int64
}

func LoadAudioStream(fs fs.FS, filename string) (AudioStream, error) {
	f, err := fs.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	switch filename[strings.LastIndex(filename, "."):] {
	case ".mp3":
		return mp3.DecodeF32(f)
	case ".ogg":
		return vorbis.DecodeF32(f)
	case ".wav":
		return wav.DecodeF32(f)
	default:
		return nil, fmt.Errorf("Invalid extension: %s", filename)
	}
}

func SaveDecodedAudio(path string) error {
	dir := filepath.Dir(path)
	filename := filepath.Base(path)

	stream, err := LoadAudioStream(os.DirFS(dir), filename)
	if err != nil {
		return err
	}

	data, err := io.ReadAll(stream)
	if err != nil {
		return err
	}

	if err := os.WriteFile(path+".dat", data, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func LoadDecodedAudio(fs fs.FS, path string) ([]byte, error) {
	f, err := fs.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ForceSaveDecodedAudio(path string) {
	err := SaveDecodedAudio(path)
	if err != nil {
		panic(err)
	}
}

func ForceLoadDecodedAudio(fs fs.FS, path string) []byte {
	data, err := LoadDecodedAudio(fs, path)
	if err != nil {
		panic(err)
	}
	return data
}
