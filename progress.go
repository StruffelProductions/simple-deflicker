package main

import (
	"fmt"
	"math"

	"github.com/gosuri/uiprogress"
)

func createProgressBars(numberOfPictures int) map[string]*uiprogress.Bar {

	progressBars := make(map[string]*uiprogress.Bar)
	progressBars["ANALYZE"] = uiprogress.AddBar(numberOfPictures).PrependCompleted().PrependElapsed()
	progressBars["ADJUST"] = uiprogress.AddBar(numberOfPictures).PrependCompleted().PrependElapsed()

	progressBars["ANALYZE"].Width = 20
	progressBars["ADJUST"].Width = 20

	progressBarFunction := func(b *uiprogress.Bar, step string) string {
		//Calculate the number of digits to display
		n := math.Floor(math.Log10(float64(b.Total)) + 1)
		f := fmt.Sprintf("%%-15v %%-%vv/%%-%vv", n, n)
		return fmt.Sprintf(f, step, b.Current(), b.Total)
	}

	progressBarFunctionAnalyze := func(b *uiprogress.Bar) string {
		return progressBarFunction(b, "Analyzing")
	}

	progressBarFunctionAdjust := func(b *uiprogress.Bar) string {
		return progressBarFunction(b, "Adjusting")
	}

	progressBars["ANALYZE"].AppendFunc(progressBarFunctionAnalyze)
	progressBars["ADJUST"].AppendFunc(progressBarFunctionAdjust)

	return progressBars
}
