package game

// From: https://github.com/hajimehoshi/go-inovation/blob/master/ino/audioloop.go

import (
	"io"

	"github.com/hajimehoshi/ebiten/audio"
)

type Loop struct {
	stream audio.ReadSeekCloser
	size   int64
}

func NewLoop(stream audio.ReadSeekCloser, size int64) *Loop {
	return &Loop{
		stream: stream,
		size:   size,
	}
}

func (l *Loop) Read(b []byte) (int, error) {
	n, err := l.stream.Read(b)
	if err == io.EOF {
		if _, err := l.Seek(0, 0); err != nil {
			return 0, err
		}
		err = nil
	}
	return n, err
}

func (l *Loop) Seek(offset int64, whence int) (int64, error) {
	next := int64(0)
	switch whence {
	case 0:
		next = offset
	case 1:
		current, err := l.stream.Seek(0, 1)
		if err != nil {
			return 0, err
		}
		next = current + offset
	case 2:
		panic("whence must be 0 or 1 for a loop stream")
	}
	next %= l.size
	pos, err := l.stream.Seek(next, 0)
	if err != nil {
		return 0, err
	}
	return pos, nil
}

func (l *Loop) Close() error {
	return l.stream.Close()
}
