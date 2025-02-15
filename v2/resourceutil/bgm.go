package resourceutil

import (
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

type BGMOptions struct {
	LoopLengthOffset int64
}

func CreateBGMPlayer(fs fs.FS, path string, audioContext *audio.Context, opts *BGMOptions) (*audio.Player, error) {
	stream, err := LoadAudioStream(fs, path)
	if err != nil {
		return nil, err
	}

	loop := audio.NewInfiniteLoopF32(stream, stream.Length()+opts.LoopLengthOffset)

	player, err := audioContext.NewPlayerF32(loop)
	if err != nil {
		return nil, err
	}

	return player, nil
}

func ForceCreateBGMPlayer(fs fs.FS, path string, audioContext *audio.Context, opts *BGMOptions) *audio.Player {
	player, err := CreateBGMPlayer(fs, path, audioContext, opts)
	if err != nil {
		panic(err)
	}
	return player
}
