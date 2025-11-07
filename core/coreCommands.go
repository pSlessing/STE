package core

import (
	"os"
)

func (*EditorCore) cmdQuit(*EditorCore, []string) error {
	os.Exit(0)
	return nil
}

func (*EditorCore) cmdWrite(*EditorCore, []string) error {
	os.Exit(0)
	return nil
}

func (*EditorCore) cmdSave(*EditorCore, []string) error {
	os.Exit(0)
	return nil
}

func (*EditorCore) cmdOpen(*EditorCore, []string) error {
	os.Exit(0)
	return nil
}

func (*EditorCore) cmdHelp(*EditorCore, []string) error {
	os.Exit(0)
	return nil
}
