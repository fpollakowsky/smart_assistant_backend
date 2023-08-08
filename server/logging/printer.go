package logging

import (
	"fmt"
	"golang.org/x/term"
	"os"
	"strings"
)

func PrintListDone(text string) {
	physicalWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))
	doc := strings.Builder{}

	doc.WriteString(listDone(text))

	if physicalWidth > 0 {
		docStyle = docStyle.MaxWidth(physicalWidth)
	}

	docStyle.PaddingBottom(0)

	fmt.Print(docStyle.Render(doc.String()))
}

func PrintListFail(text string) {
	physicalWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))
	doc := strings.Builder{}

	doc.WriteString(listFail(text))

	if physicalWidth > 0 {
		docStyle = docStyle.MaxWidth(physicalWidth)
	}

	docStyle.PaddingBottom(0)

	fmt.Print(docStyle.Render(doc.String()))
}
