package handlers

import (
	"bytes"
	"os"
	"time"
	"fmt"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

type Native struct {
}

func (n *Native) Play(fileName string) error {
	// Read the mp3 file into memory

	fmt.Print("a ")
	fileBytes, err := os.ReadFile(fileName)

	fmt.Print("b ")
	if err != nil {
		return err
	}

	fmt.Print("c ")

	fileBytesReader := bytes.NewReader(fileBytes)

	fmt.Print("d ")


	// Decode file
	decodedMp3, err := mp3.NewDecoder(fileBytesReader)

	fmt.Print("e ")
	if err != nil {
		return err
	}

	fmt.Print("f ")

	numOfChannels := 2
	audioBitDepth := 2

	fmt.Print("g ")

	otoCtx, readyChan, err := oto.NewContext(decodedMp3.SampleRate(), numOfChannels, audioBitDepth)

	fmt.Print("h ")
	if err != nil {
		return err
	}

	fmt.Print("i ")
	<-readyChan

	fmt.Print("j ")

	player := otoCtx.NewPlayer(decodedMp3)

	fmt.Print("k ")

	player.Play()

	fmt.Print("l ")

	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}

	fmt.Print("m ")

	return player.Close()

}
