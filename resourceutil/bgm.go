package resourceutil

import (
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

func CreateBGMPlayer(repository fs.FS, path string, audioContext *audio.Context) (*audio.Player, error) {
	stream, err := LoadAudioStream(repository, path, audioContext)
	if err != nil {
		return nil, err
	}

	loop := audio.NewInfiniteLoop(stream, stream.Length())

	player, err := audio.NewPlayer(audioContext, loop)
	if err != nil {
		return nil, err
	}

	return player, nil
}

func ForceCreateBGMPlayer(repository fs.FS, path string, audioContext *audio.Context) *audio.Player {
	player, err := CreateBGMPlayer(repository, path, audioContext)
	if err != nil {
		panic(err)
	}
	return player
}
