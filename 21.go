package main

import "fmt"

// MediaPlayer - нужный интерфейс
type MediaPlayer interface {
	PlayVideo(fileName string)
	PlayAudio(fileName string)
}

// AudioPlayer - реализует только часть интерфейса MediaPlayer
type AudioPlayer struct{}

func (a *AudioPlayer) PlayAudio(fileName string) {
	fmt.Printf("Playing audio: %s\n", fileName)
}

// VideoPlayer - реализует только часть интерфейса MediaPlayer
type VideoPlayer struct{}

func (v *VideoPlayer) PlayVideo(fileName string) {
	fmt.Printf("Playing video: %s\n", fileName)
}

// MediaAdapter - стуктура, которая полностью реализует MediaPlayer интерфейс
type MediaAdapter struct {
	audioPlayer *AudioPlayer
	videoPlayer *VideoPlayer
}

func (a *MediaAdapter) PlayAudio(fileName string) {
	a.audioPlayer.PlayAudio(fileName)
}

func (a *MediaAdapter) PlayVideo(fileName string) {
	a.videoPlayer.PlayVideo(fileName)
}

func UserPlayer(media MediaPlayer) {
	media.PlayAudio("audio.mp3")
	media.PlayVideo("video.mp4")
}

func main() {
	audioPlayer := &AudioPlayer{}
	videoPlayer := &VideoPlayer{}

	// mediaAdapter - удовлетворяет интерфейсу MediaPlayer и может использоваться вместо него
	mediaAdapter := &MediaAdapter{audioPlayer: audioPlayer, videoPlayer: videoPlayer}

	UserPlayer(mediaAdapter)
}
