// core/editor.go
package core

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
)

type EditorCore struct {
	// Buffer and cursor state
	TextBuffer [][]rune
	CursorX    int
	CursorY    int
	OffsetX    int
	OffsetY    int
	SourceFile string

	// Display state
	Terminal tcell.Screen
	Styles   *StyleSet
	Cols     int
	Rows     int

	// Plugin system
	plugins  map[string]Plugin
	commands map[string]Command

	// Other state
	InputBuffer    []rune
	LineCountWidth int
}

func NewEditor() (*EditorCore, error) {
	terminal, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	if err := terminal.Init(); err != nil {
		return nil, err
	}

	editor := &EditorCore{
		TextBuffer:     [][]rune{{}},
		CursorX:        3,
		CursorY:        0,
		Terminal:       terminal,
		plugins:        make(map[string]Plugin),
		commands:       make(map[string]Command),
		LineCountWidth: 3,
		Styles: &StyleSet{
			MAINSTYLE:      tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorDefault),
			STATUSSTYLE:    tcell.StyleDefault.Foreground(tcell.ColorDefault).Background(tcell.ColorWhite),
			MSGSTYLE:       tcell.StyleDefault.Foreground(tcell.ColorDefault).Background(tcell.ColorWhite),
			LINECOUNTSTYLE: tcell.StyleDefault.Foreground(tcell.ColorDarkCyan).Background(tcell.ColorWhite),
		},
	}

	// Register built-in commands
	editor.registerBuiltInCommands()

	return editor, nil
}

func (e *EditorCore) registerBuiltInCommands() {
	builtins := []Command{
		{
			Name:        "quit",
			Aliases:     []string{"q"},
			Description: "Exit the editor",
			Execute:     e.cmdQuit,
		},
		{
			Name:        "write",
			Aliases:     []string{"w"},
			Description: "Enter write mode",
			Execute:     e.cmdWrite,
		},
		{
			Name:        "save",
			Aliases:     []string{"s"},
			Description: "Save current file",
			Execute:     e.cmdSave,
		},
		{
			Name:        "open",
			Aliases:     []string{"o"},
			Description: "Open a file",
			Execute:     e.cmdOpen,
		},
		{
			Name:        "help",
			Aliases:     []string{"h", "?"},
			Description: "Show available commands",
			Execute:     e.cmdHelp,
		},
	}

	for _, cmd := range builtins {
		e.commands[cmd.Name] = cmd
		for _, alias := range cmd.Aliases {
			e.commands[alias] = cmd
		}
	}
}

func (e *EditorCore) Run() {
	// Load plugins before starting
	if err := e.LoadPluginsFromDirectory("./modules"); err != nil {
		fmt.Printf("Warning: %v\n", err)
	}

	e.mainLoop()
}

func (e *EditorCore) mainLoop() {
	// Your existing main loop logic
	// But now handleCommand uses ExecuteCommand
}

func (e *EditorCore) handleCommand() {
	cmdText := string(e.InputBuffer)
	parts := strings.Fields(cmdText)

	if len(parts) == 0 {
		return
	}

	cmdName := strings.ToLower(parts[0])
	args := parts[1:]

	if err := e.ExecuteCommand(cmdName, args); err != nil {
		e.ShowError(err.Error())
	}
}
