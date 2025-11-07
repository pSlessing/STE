package core

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
)

// Settings represents the configuration structure for termbox colors
type Settings struct {
	BGColor          tcell.Color `json:"bg_color"`
	FGColor          tcell.Color `json:"fg_color"`
	StatusBGColor    tcell.Color `json:"status_bg_color"`
	StatusFGColor    tcell.Color `json:"status_fg_color"`
	MsgBGColor       tcell.Color `json:"msg_bg_color"`
	MsgFGColor       tcell.Color `json:"msg_fg_color"`
	LineCountBGColor tcell.Color `json:"line_count_bg_color"`
	LineCountFGColor tcell.Color `json:"line_count_fg_color"`
	ErrorBGColor     tcell.Color `json:"error_bg_color"`
	ErrorFGColor     tcell.Color `json:"error_fg_color"`
}

type StyleSet struct {
	Main      tcell.Style
	Status    tcell.Style
	Linecount tcell.Style
	Message   tcell.Style
	Error     tcell.Style
}

// GetDefaultSettings returns the default configuration
func GetDefaultSettings() Settings {
	return Settings{
		BGColor:          tcell.ColorBlack,
		FGColor:          tcell.ColorWhite,
		StatusBGColor:    tcell.ColorWhite,
		StatusFGColor:    tcell.ColorBlack,
		MsgBGColor:       tcell.ColorWhite,
		MsgFGColor:       tcell.ColorBlack,
		LineCountBGColor: tcell.ColorWhite,
		LineCountFGColor: tcell.ColorLightBlue,
		ErrorBGColor:     tcell.ColorBlack,
		ErrorFGColor:     tcell.ColorRed,
	}
}

// SaveSettings saves the current settings to a JSON file
func SaveSettings(settings Settings) error {
	// Get OS-specific config directory
	configDir, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get user config directory: %w", err)
	}
	configDir = filepath.Join(configDir, "SlessingTextEditor")

	// Ensure the config directory exists
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// Convert settings to JSON string
	jsonStr, err := SettingsToJSON(settings)
	if err != nil {
		return fmt.Errorf("failed to convert settings to JSON: %w", err)
	}

	// Write to file using similar pattern to WriteBufferToFile
	configPath := filepath.Join(configDir, "config.json")
	file, err := os.Create(configPath)
	if err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	_, err = writer.WriteString(jsonStr)
	if err != nil {
		return fmt.Errorf("failed to write settings to file: %w", err)
	}

	return nil
}

// LoadSettings loads settings from a JSON file, creating default config if file doesn't exist
func LoadSettings() (Settings, error) {
	// Get OS-specific config directory
	configDir, err := os.UserConfigDir()
	if err != nil {
		return Settings{}, fmt.Errorf("failed to get user config directory: %w", err)
	}
	configDir = filepath.Join(configDir, "SlessingTextEditor")

	configPath := filepath.Join(configDir, "config.json")

	// Try to open the file, similar to OpenFile pattern
	file, err := os.Open(configPath)
	if err != nil {
		// File doesn't exist, create default config
		defaultSettings := GetDefaultSettings()
		if saveErr := SaveSettings(defaultSettings); saveErr != nil {
			return Settings{}, fmt.Errorf("failed to create default config: %w", saveErr)
		}
		return defaultSettings, nil
	}
	defer file.Close()

	// Read the file content
	var jsonContent string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		jsonContent += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		return Settings{}, fmt.Errorf("failed to read config file: %w", err)
	}

	// Convert JSON to settings
	settings, err := JSONToSettings(jsonContent)
	if err != nil {
		return Settings{}, fmt.Errorf("failed to parse settings JSON: %w", err)
	}

	return settings, nil
}

// SettingsToJSON converts a Settings struct to a JSON string
func SettingsToJSON(settings Settings) (string, error) {
	jsonData, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal settings to JSON: %w", err)
	}
	return string(jsonData), nil
}

// JSONToSettings converts a JSON string to a Settings struct
func JSONToSettings(jsonStr string) (Settings, error) {
	var settings Settings
	if err := json.Unmarshal([]byte(jsonStr), &settings); err != nil {
		return Settings{}, fmt.Errorf("failed to unmarshal JSON to settings: %w", err)
	}
	return settings, nil
}

// ApplySettings applies the loaded settings to the global color variables
func (e *EditorCore) ApplySettings(settings Settings) {
	e.Styles.Main = tcell.StyleDefault.Background(settings.BGColor).Foreground(settings.FGColor)
	e.Styles.Status = tcell.StyleDefault.Background(settings.StatusBGColor).Foreground(settings.StatusFGColor)
	e.Styles.Message = tcell.StyleDefault.Background(settings.MsgBGColor).Foreground(settings.MsgFGColor)
	e.Styles.Linecount = tcell.StyleDefault.Background(settings.LineCountBGColor).Foreground(settings.LineCountFGColor)
	e.Styles.Error = tcell.StyleDefault.Background(settings.ErrorBGColor).Foreground(settings.ErrorFGColor)

}

// GetCurrentSettings creates a Settings struct from the current global variables
func (e *EditorCore) GetCurrentSettings() Settings {
	mainfg, mainbg, _ := e.Styles.Main.Decompose()
	statusfg, statusbg, _ := e.Styles.Linecount.Decompose()
	msgfg, msgbg, _ := e.Styles.Message.Decompose()
	linecountfg, linecountbg, _ := e.Styles.Linecount.Decompose()
	errorfg, errorbg, _ := e.Styles.Error.Decompose()
	return Settings{
		BGColor:          mainbg,
		FGColor:          mainfg,
		StatusBGColor:    statusbg,
		StatusFGColor:    statusfg,
		MsgBGColor:       msgbg,
		MsgFGColor:       msgfg,
		LineCountBGColor: linecountbg,
		LineCountFGColor: linecountfg,
		ErrorBGColor:     errorbg,
		ErrorFGColor:     errorfg,
	}
}

// Example usage:
/*
func main() {
	// Load settings on startup (creates default if doesn't exist)
	settings, err := LoadSettings()
	if err != nil {
		fmt.Printf("Error loading settings: %v\n", err)
		return
	}
	ApplySettings(settings)

	// Your application code here...

	// Save settings when modified
	currentSettings := GetCurrentSettings()
	err = SaveSettings(currentSettings)
	if err != nil {
		fmt.Printf("Error saving settings: %v\n", err)
	}
}
*/
