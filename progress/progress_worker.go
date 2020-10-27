package progress

import (
	"github.com/cheggaaa/pb/v3"
)

type ProgressWorker interface {
	GetMax() int64
	SetCurrent(number int64)
	GetCurrent() int64
	Finish()
}

type AsciiProgressWorker struct {
	bar *pb.ProgressBar
	max int64
}

func NewAsciiProgressWorker(max int64) *AsciiProgressWorker {
	return &AsciiProgressWorker{max: max, bar: pb.Start64(max)}
}

func (a *AsciiProgressWorker) GetMax() int64 {
	return a.max
}

func (a *AsciiProgressWorker) SetCurrent(number int64) {
	a.bar.SetCurrent(number)
}

func (a *AsciiProgressWorker) Finish() {
	a.bar.Finish()
}

func (a *AsciiProgressWorker) GetCurrent() int64 {
	return a.bar.Current()
}
