package logging

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/fatih/color"
	"github.com/lucasb-eyer/go-colorful"
	"golang.org/x/term"
	"os"
	"strconv"
	"strings"
	"time"
)

func Initialize() {
	physicalWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))
	doc := strings.Builder{}

	// Dialog
	{
		question := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render("NETHCON - HOME AI\nVERSION 0.1")
		ui := lipgloss.JoinVertical(lipgloss.Center, question)

		dialog := lipgloss.Place(width, 10,
			lipgloss.Center, lipgloss.Center,
			dialogBoxStyle.Render(ui),
			lipgloss.WithWhitespaceChars("猫咪"),
			lipgloss.WithWhitespaceForeground(subtle),
		)

		doc.WriteString(dialog + "\n\n")
	}

	doc.WriteString(listHeader("Setting up things"))

	if physicalWidth > 0 {
		docStyle = docStyle.MaxWidth(physicalWidth)
	}

	docStyle.PaddingBottom(0)

	fmt.Print(docStyle.Render(doc.String()))
}

func GinFormat(timeStamp time.Time, statusCode int, clientIP, path string, latency time.Duration) string {
	var statusCodeColor *color.Color
	var latencyColor *color.Color
	//var prefixOutputColor *color.Color

	//prefixOutputTextColor := color.New(color.FgHiWhite)
	//prefixOutputColor = prefixOutputTextColor.Add(color.BgBlue)

	statusCodeTextColor := color.New(color.FgHiWhite)
	latencyTextColor := color.New(color.FgHiWhite)

	switch statusCode {
	case 200:
		statusCodeColor = statusCodeTextColor.Add(color.BgGreen)
	case 404:
		statusCodeColor = statusCodeTextColor.Add(color.BgRed)
	}

	switch {
	case latency.Milliseconds() < 250:
		latencyColor = latencyTextColor.Add(color.FgGreen)
	case latency.Milliseconds() >= 250 && latency.Abs() < 1000:
		latencyColor = latencyTextColor.Add(color.FgYellow)
	case latency.Milliseconds() >= 1000:
		latencyColor = latencyTextColor.Add(color.FgRed)
	}

	//prefixOutput := prefixOutputColor.Sprintf("Endpoint Hit") + " "
	statusCodeOutput := statusCodeColor.Sprintf(" " + strconv.Itoa(statusCode) + " ")
	timeStampOutput := timeStamp.Format(time.RFC822)
	clientIPPOutput := clientIP
	latencyOutput := latencyColor.Sprintf(" " + latency.String() + " ")

	formattedOutput := fmt.Sprintf("%s %s %10s %14s %25s \n", timeStampOutput, statusCodeOutput, path, clientIPPOutput, latencyOutput)
	return formattedOutput
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func colorGrid(xSteps, ySteps int) [][]string {
	x0y0, _ := colorful.Hex("#F25D94")
	x1y0, _ := colorful.Hex("#EDFF82")
	x0y1, _ := colorful.Hex("#643AFF")
	x1y1, _ := colorful.Hex("#14F9D5")

	x0 := make([]colorful.Color, ySteps)
	for i := range x0 {
		x0[i] = x0y0.BlendLuv(x0y1, float64(i)/float64(ySteps))
	}

	x1 := make([]colorful.Color, ySteps)
	for i := range x1 {
		x1[i] = x1y0.BlendLuv(x1y1, float64(i)/float64(ySteps))
	}

	grid := make([][]string, ySteps)
	for x := 0; x < ySteps; x++ {
		y0 := x0[x]
		grid[x] = make([]string, xSteps)
		for y := 0; y < xSteps; y++ {
			grid[x][y] = y0.BlendLuv(x1[x], float64(y)/float64(xSteps)).Hex()
		}
	}

	return grid
}
