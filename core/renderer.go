package core

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

// PrintMessage #TODO: should this be able to use any, or standard colors every time?
func (e *EditorCore) PrintMessage(col, row int, fg, bg tcell.Color, msg string) {
	for _, c := range msg {
		currStyle := tcell.StyleDefault.Foreground(fg).Background(bg)
		e.Terminal.SetContent(col, row, c, nil, currStyle)
		col += runewidth.RuneWidth(c)
	}
}

func (e *EditorCore) PrintMessageStyle(col, row int, style tcell.Style, msg string) {
	for _, c := range msg {
		e.Terminal.SetContent(col, row, c, nil, style)
		col += runewidth.RuneWidth(c)
	}
}

func (e *EditorCore) DisplayBuffer() {
	var row, col int

	for row = 0; row <= e.Rows; row++ {
		textBufferRow := row + e.OffsetY

		e.DisplayLineNumber(row, textBufferRow)

		for col = 0; col < e.Cols; col++ {
			textBufferCol := col + e.OffsetX

			if textBufferRow >= 0 &&
				textBufferRow < len(e.TextBuffer) &&
				textBufferCol < len(e.TextBuffer[textBufferRow]) {
				e.Terminal.SetContent(col+e.LineCountWidth, row,
					e.TextBuffer[textBufferRow][textBufferCol],
					nil, e.Styles.Main)
			}
		}
	}
}

func (e *EditorCore) DisplayStatus() {
	var col int

	e.Terminal.SetContent(0, e.Rows+1, ' ', nil, e.Styles.Status)
	e.Terminal.SetContent(1, e.Rows+1, '', nil, e.Styles.Status)
	e.Terminal.SetContent(2, e.Rows+1, '❯', nil, e.Styles.Status)

	BufferOffset := 3
	for col = BufferOffset; col < e.Cols+e.LineCountWidth; col++ {
		e.Terminal.SetContent(col, e.Rows+1, ' ', nil, e.Styles.Status)
		if col-BufferOffset < len(e.InputBuffer) {
			e.Terminal.SetContent(col, e.Rows+1,
				e.InputBuffer[col-BufferOffset],
				nil, e.Styles.Status)
		}
	}

	var currentLine = e.CursorY + e.OffsetY
	var lineNumberStr = strconv.Itoa(currentLine + 1)
	var currentColumn = e.CursorX + e.OffsetX - e.LineCountWidth
	var columnNumberStr = strconv.Itoa(currentColumn + 1)
	// #TODO do the offsets more neat
	e.PrintMessageStyle(e.Cols, e.Rows+1, e.Styles.Status, columnNumberStr)
	e.PrintMessageStyle(e.Cols-4, e.Rows+1, e.Styles.Status, "col")
	e.PrintMessageStyle(e.Cols-8, e.Rows+1, e.Styles.Status, lineNumberStr)
	e.PrintMessageStyle(e.Cols-12, e.Rows+1, e.Styles.Status, "row")
}

func (e *EditorCore) DisplayLineNumber(row int, textBufferRow int) {
	lineNumberStr := "~"

	if textBufferRow < len(e.TextBuffer) {
		lineNumberStr = strconv.Itoa(textBufferRow + 1)
	}

	lineNumberOffset := e.LineCountWidth - len(lineNumberStr)
	if lineNumberOffset > 0 {
		for i := 0; i < lineNumberOffset; i++ {
			e.Terminal.SetContent(i, row, ' ', nil, e.Styles.Linecount)
		}
	}

	e.PrintMessageStyle(lineNumberOffset, row, e.Styles.Linecount, lineNumberStr)
}

//TODO: Refactor all of the below bs
/* func (e *EditorCore) DisplaySettingsLoop(currentPos int) {
	//Offset between setting names and colors
	colorOffset := 2
	e.Terminal.SetContent(0, currentPos, ' ', nil, tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack))
	styleList := e.Styles.AsSlice()
	currDisplayRow := 0

	// Define style names that match your actual styles
	styleNames := []string{"Main", "Status", "Msg", "LineCount"}

	for i := 0; i < len(styleList); i++ {
		fgColor, bgColor, _ := styleList[i].Decompose()
		// Display background setting (this comes first based on your currentPos % 2 == 0 check)
		e.PrintMessage(1, currDisplayRow, tcell.ColorWhite, tcell.ColorDefault, styleNames[i]+" BG")
		DisplayName := "Default"
		if bgColor.Name() != "" {
			DisplayName = bgColor.Name()
		}
		e.PrintMessage(1+len(styleNames[i])+len(" BG")+colorOffset, currDisplayRow, tcell.ColorWhite, tcell.ColorDefault, DisplayName)
		currDisplayRow++
		// Display foreground setting
		DisplayName = "Default"
		if fgColor.Name() != "" {
			DisplayName = fgColor.Name()
		}
		e.PrintMessage(1, currDisplayRow, tcell.ColorWhite, tcell.ColorDefault, styleNames[i]+" FG")
		e.PrintMessage(1+len(styleNames[i])+len(" FG")+colorOffset, currDisplayRow, tcell.ColorWhite, tcell.ColorDefault, DisplayName)
		currDisplayRow++
	}
}

func (e *EditorCore) DisplayColorsLoop(offset int) {
	styleList := e.Styles.AsSlice()
	//LineCount
	e.PrintMessageStyle(0, len(styleList)+offset, e.Styles.Linecount, "~1")
	e.PrintMessageStyle(0, len(styleList)+1+offset, e.Styles.Linecount, "~2")
	e.PrintMessageStyle(0, len(styleList)+2+offset, e.Styles.Linecount, "~3")
	e.PrintMessageStyle(0, len(styleList)+3+offset, e.Styles.Linecount, "~4")
	e.PrintMessageStyle(0, len(styleList)+4+offset, e.Styles.Linecount, "~5")
	e.PrintMessageStyle(0, len(styleList)+5+offset, e.Styles.Linecount, "~6")
	//Main
	e.PrintMessageStyle(2, len(styleList)+0+offset, e.Styles.Main, "This is a piece of text! Some characters for testing: ! # ¤ % & / [] {}")

	//Statusbar
	e.PrintMessageStyle(0, len(styleList)+6+offset, e.Styles.Status, "write                                                     row 0 col 0")
	//MSG
	e.PrintMessageStyle(30, len(styleList)+3+offset, e.Styles.Message, "Open file:")
	e.PrintMessageStyle(30, len(styleList)+4+offset, e.Styles.Message, "file.txt  ")
}
*/
