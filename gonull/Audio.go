package main

import (
	"bytes"
	"io"
	"null/stagerror"
	"os"
	"slices"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

// Audio Code
// Not used in future code

type AudioHandler struct {
	ctx             *audio.Context
	soundTracks     map[string]*AudioPlayer
	curAudio        []string
	soundEffects    map[string]*AudioPlayer
	curSoundEffects []string
	mute            bool
}

// After culling this struct it looks
// very useless, could be more in future
type AudioPlayer struct {
	audioPlayer *audio.Player
}

type musicType int

const sampleRate int = 44_100

type audioStream interface {
	io.ReadSeeker
	Length() int64
}

func (a *AudioHandler) newPlayer(stream audioStream) (*AudioPlayer, *stagerror.Error) {
	p, err := a.ctx.NewPlayer(stream)
	if err != nil {
		return nil, stagerror.New(16, "Couldn't create player from audio")
	}

	player := &AudioPlayer{
		audioPlayer: p,
	}

	return player, nil
}

func (ap *AudioPlayer) close() error {
	return ap.audioPlayer.Close()
}

func (a *AudioHandler) updateAudio() {
	if a.mute {
		return
	}

	for _, key := range a.curAudio {
		if a.soundTracks[key] != nil {
			a.soundTracks[key].update()
		}
	}

	// for _, key := range a.curSoundEffects {
	// 	if a.soundEffects[key] != nil {
	// 		a.soundEffects[key].update()
	// 	}
	// }
}

func (a *AudioHandler) toggleMute() {
	if a.mute {
		a.mute = false
		for _, key := range a.curAudio {
			s, ok := a.soundTracks[key]
			if ok {
				s.audioPlayer.Play()
			}
		}
	} else {
		a.mute = true
		for _, key := range a.curAudio {
			s, ok := a.soundTracks[key]
			if ok {
				s.audioPlayer.Pause()
			}
		}
	}
}

func (ap *AudioPlayer) update() {
	if ap.audioPlayer.IsPlaying() {
		return
	}

	err := ap.audioPlayer.Rewind()
	if err != nil {
		stagerror.New(19, "Couldn't rewind audio").SaveToLog(IS_RELEASE)
	}

	ap.audioPlayer.Play()

	return
}

func (a *AudioHandler) playAudio(key string) {
	if a.mute {
		return
	}

	m, ok := a.soundTracks[key]
	if !ok {
		stagerror.New(24, "Couldn't find audio: "+key).SaveToLog(IS_RELEASE)
		return
	}

	if !slices.Contains(a.curAudio, key) {
		a.curAudio = append(a.curAudio, key)
	}
	a.soundTracks[key] = m

	err := a.soundTracks[key].audioPlayer.Rewind()
	if err != nil {
		stagerror.New(19, "Couldn't rewind audio: "+key).SaveToLog(IS_RELEASE)
		return
	}

	a.soundTracks[key].audioPlayer.Play()
}

func (a *AudioHandler) stopAudio(key string) {
	if !slices.Contains(a.curAudio, key) {
		return
	}

	m, ok := a.soundTracks[key]
	if !ok {
		stagerror.New(24, "Couldn't find audio: "+key).SaveToLog(IS_RELEASE)
		return
	}

	m.audioPlayer.Pause()
}

func (a *AudioHandler) playEffect(key string) {
	if a.mute {
		return
	}

	m, ok := a.soundEffects[key]
	if !ok {
		stagerror.New(24, "Couldn't find audio: "+key).SaveToLog(IS_RELEASE)
		return
	}

	if !slices.Contains(a.curSoundEffects, key) {
		a.curSoundEffects = append(a.curSoundEffects, key)
	}
	a.soundEffects[key] = m

	err := a.soundEffects[key].audioPlayer.Rewind()
	if err != nil {
		stagerror.New(19, "Couldn't rewind audio: "+key).SaveToLog(IS_RELEASE)
		return
	}
	a.soundEffects[key].audioPlayer.Play()
}

// setVolume will panic if volume is not between 0 and 1
func (a *AudioHandler) setVolume(volume float64) {
	for _, ap := range a.soundTracks {
		ap.audioPlayer.SetVolume(volume)
	}
}

func (a *AudioHandler) setEffectVolume(volume float64) {
	for _, ap := range a.soundEffects {
		ap.audioPlayer.SetVolume(volume)
	}
}

func (a *AudioHandler) init() {
	a.mute = false
	a.ctx = audio.NewContext(sampleRate)
	a.loadAudio()
}

func (a *AudioHandler) loadAudio() {
	a.soundTracks = make(map[string]*AudioPlayer)
	a.soundEffects = make(map[string]*AudioPlayer)

	// a.loadTrack("fighing", "theme")

	// All the layers
	a.loadTrack("themeGuitar", "theme0")
	a.loadTrack("themeDrums", "theme1")
	a.loadTrack("themeBass", "theme2")
	a.loadTrack("themeRythm", "theme3")
	a.loadTrack("themeBackup", "theme4")
	a.loadTrack("themeLead", "theme5")

	a.loadEffect("badDereference", "badDereference")
	a.loadEffect("dereference", "dereference")
	a.loadEffect("jump", "jump")
	a.loadEffect("boost", "boost")
}

func (a *AudioHandler) loadTrack(fName, mName string) {
	var s audioStream

	audio, err := os.ReadFile("assets/sounds/tracks/" + fName + ".ogg")
	if err != nil {
		stagerror.New(20, "Couldn't read file: "+fName+".ogg").SaveToLog(IS_RELEASE)
		return
	}

	s, err = vorbis.DecodeWithoutResampling(bytes.NewReader(audio))
	if err != nil {
		stagerror.New(22, "Couldn't decode audio: "+fName+".ogg").SaveToLog(IS_RELEASE)
		return
	}

	ap, serr := a.newPlayer(s)
	serr.SaveToLog(IS_RELEASE)

	a.soundTracks[mName] = ap
}

func (a *AudioHandler) loadEffect(fName, mName string) {
	var s audioStream

	audio, err := os.ReadFile("assets/sounds/effects/" + fName + ".ogg")
	if err != nil {
		stagerror.New(21, "Couldn't read file: "+fName+".ogg").SaveToLog(IS_RELEASE)
		return
	}

	s, err = vorbis.DecodeWithoutResampling(bytes.NewReader(audio))
	if err != nil {
		stagerror.New(23, "Couldn't decode audio: "+fName+".ogg").SaveToLog(IS_RELEASE)
		return
	}

	ap, serr := a.newPlayer(s)
	serr.SaveToLog(IS_RELEASE)

	a.soundEffects[mName] = ap
}
