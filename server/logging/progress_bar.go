package logging

import (
	"github.com/schollz/progressbar/v3"
)

func Progressbar(max int) *progressbar.ProgressBar {
	Bar := progressbar.NewOptions(
		max,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowIts(),
		progressbar.OptionSetWidth(43),
		progressbar.OptionSetDescription("    Generating Trashcans..."),
		progressbar.OptionSetElapsedTime(false),
		progressbar.OptionSetPredictTime(true))

	return Bar
}
