package handlers

import (
	"os"
	"time"

	"github.com/hajimehoshi/go-mp3"
	"github.com/ebitengine/oto/v3"
)

type Native struct {
}

func (n *Native) Play(fileName string) error {
	// Read the mp3 file into memory
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	// Decode file
	decodedMp3, err := mp3.NewDecoder(file)
	if err != nil {
		return err
	}

	op := &oto.NewContextOptions{}

    	// Usually 44100 or 48000. Other values might cause distortions in Oto
    	op.SampleRate = 44100

    	// Number of channels (aka locations) to play sounds from. Either 1 or 2.
    	// 1 is mono sound, and 2 is stereo (most speakers are stereo). 
    	op.ChannelCount = 2

    	// Format of the source. go-mp3's format is signed 16bit integers.
    	op.Format = oto.FormatSignedInt16LE

	otoCtx, readyChan, err := oto.NewContext(op)
	if err != nil {
		return err
	}
	<-readyChan

	player := otoCtx.NewPlayer(decodedMp3)

	player.Play()

	for player.IsPlaying() {
		time.Sleep(2*time.Second)
	}

	err = player.Close()

	return err
}
