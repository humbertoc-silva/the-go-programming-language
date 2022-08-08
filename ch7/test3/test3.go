package main

import (
	"fmt"
	"io"
	"time"
)

type Artifact interface {
	Title() string
	Creators() []string
	Created() time.Time
}

type Text interface {
	Pages() int
	Words() int
	PageSize() int
}

type Audio interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string // e.g., "MP3", "WAV"
}

type Video interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string // e.g., "MP4", "WMV"
	Resolution() (x, y int)
}

type Streamer interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string
}

type Album struct {
	title    string
	creators []string
	created  time.Time
}

func (a Album) Title() string {
	return a.title
}

func (a Album) Creators() []string {
	return a.creators
}

func (a Album) Created() time.Time {
	return a.created
}

func main() {
	var artifact Artifact = Album{
		title:    "The Dark Side of the Moon",
		creators: []string{"Nick Mason", "Roger Waters", "Richard Wright", "David Gilmour"},
		created:  time.Date(1973, time.March, 1, 0, 0, 0, 0, time.UTC),
	}
	fmt.Println(artifact.Title())
	fmt.Println(artifact.Creators())
	fmt.Println(artifact.Created())
}
