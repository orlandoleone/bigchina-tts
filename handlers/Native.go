package handlers

import (
	"os"
	"time"
	"fmt"

	"github.com/hajimehoshi/go-mp3"
	"github.com/ebitengine/oto/v3"
)

type Native struct {
}

func (n *Native) Play(fileName string, flag int) error {
	// Read the mp3 file into memory

	fmt.Print("a ")
	file, err := os.Open(fileName)
	fmt.Print("b ")
	if err != nil {
		return err
	}

	fmt.Print("c ")

	// Decode file
	decodedMp3, err := mp3.NewDecoder(file)
	/*
	if err != nil {
		return err
	}
 	*/
	fmt.Print("d ")

	op := &oto.NewContextOptions{}

	fmt.Print("e ")

    	// Usually 44100 or 48000. Other values might cause distortions in Oto
    	op.SampleRate = 4800

	fmt.Print("f ")

    	// Number of channels (aka locations) to play sounds from. Either 1 or 2.
    	// 1 is mono sound, and 2 is stereo (most speakers are stereo). 
    	op.ChannelCount = 2

	fmt.Print("g ")

    	// Format of the source. go-mp3's format is signed 16bit integers.
    	op.Format = oto.FormatSignedInt16LE

	fmt.Print("h ")

	var otoCtx *oto.Context
	var readyChan chan struct{}

	if flag == 0 {
		otoCtx, readyChan, err = oto.NewContext(op)
	}

	/*
	if err != nil {
		return err
	}
 	*/

	<-readyChan

	fmt.Print("i ")

	fmt.Print("j ")

	player := otoCtx.NewPlayer(decodedMp3)

	fmt.Print("k ")

	player.Play()

	fmt.Print("l ")

	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}

	fmt.Print("m ")

	err = player.Close()

	fmt.Print("n ")
	
	return err
}
