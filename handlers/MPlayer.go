package handlers

import (
	"os/exec"
)

type MPlayer struct {}

func (MPlayer *MPlayer) Play(fileName string, flag int) error {
	mplayer := exec.Command("mplayer", "-cache", "8092", "-", fileName)
	flag == flag
	return mplayer.Run()
}
