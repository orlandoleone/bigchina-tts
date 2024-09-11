package handlers

type PlayerInterface interface {
	Play(fileName string, flag int) error
}
