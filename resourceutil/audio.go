package resourceutil

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

type AudioStream interface {
	io.ReadSeeker
	Length() int64
}

func LoadAudioStream(repository fs.FS, filename string, audioContext *audio.Context) (AudioStream, error) {
	f, err := repository.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	switch filename[strings.LastIndex(filename, "."):] {
	case ".mp3":
		return mp3.Decode(audioContext, bytes.NewReader(data))
	case ".ogg":
		return vorbis.Decode(audioContext, bytes.NewReader(data))
	case ".wav":
		return wav.Decode(audioContext, bytes.NewReader(data))
	default:
		return nil, fmt.Errorf("Invalid extension: %s", filename)
	}
}

func SaveDecodedAudio(path string, audioContext *audio.Context) error {
	dir := filepath.Dir(path)
	filename := filepath.Base(path)

	stream, err := LoadAudioStream(os.DirFS(dir), filename, audioContext)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(stream)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path+".dat", data, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func LoadDecodedAudio(repository fs.FS, path string, audioContext *audio.Context) ([]byte, error) {
	f, err := repository.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ForceSaveDecodedAudio(path string, audioContext *audio.Context) {
	err := SaveDecodedAudio(path, audioContext)
	if err != nil {
		panic(err)
	}
}

func ForceLoadDecodedAudio(repository fs.FS, path string, audioContext *audio.Context) []byte {
	data, err := LoadDecodedAudio(repository, path, audioContext)
	if err != nil {
		panic(err)
	}
	return data
}
