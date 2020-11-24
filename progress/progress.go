package progress

import "fmt"

// Progress is a struct to describe
// the progress bar
type Progress struct {
	percent int64
	cur     int64
	total   int64
	rate    string
	graph   string
}

// New initializes a new instance of Progress
func (bar *Progress) New(start, total int64) {
	bar.cur = start
	bar.total = total
	if bar.graph == "" {
		bar.graph = "█"
	}
	bar.percent = bar.getPercent()
	for i := 0; i < int(bar.percent); i += 2 {
		bar.rate += bar.graph // initial progress position
	}
}

func (bar *Progress) getPercent() int64 {
	return int64(float32(bar.cur) / float32(bar.total) * 100)
}

// Play increments the progress bar
func (bar *Progress) Play(cur int64) string {
	bar.cur = cur
	last := bar.percent
	bar.percent = bar.getPercent()
	if bar.percent != last && bar.percent%2 == 0 {
		bar.rate += bar.graph
	}
	return fmt.Sprintf("\r[%-50s]%3d%%", bar.rate, bar.percent)
}

// Finish ends printing the bar graph
func (bar *Progress) Finish() {
	fmt.Println()
}
