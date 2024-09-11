package handlers

import (
	"os/exec"
)

type MPlayer struct {}

func (MPlayer *MPlayer) Play(fileName string, flag int) error {
	mplayer := exec.Command("mplayer", "-cache", "8092", "-", fileName)
	_ = flag
	return mplayer.Run()
}
