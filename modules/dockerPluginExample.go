package main

import "slessingTextEditor/core"

// modules/docker/docker.go

type DockerPlugin struct {
	core *core.EditorCore
}

func (d *DockerPlugin) Name() string {
	return "docker"
}

func (d *DockerPlugin) Commands() []core.Command {
	return []core.Command{
		{
			Name:        "docker-ps",
			Aliases:     []string{"dps"},
			Description: "List Docker containers",
			Execute:     d.listContainers,
		},
		{
			Name:        "docker-logs",
			Aliases:     []string{"dlogs"},
			Description: "View container logs",
			Execute:     d.viewLogs,
		},
	}
}

func (d *DockerPlugin) Initialize(editorCore *core.EditorCore) error {
	d.core = editorCore
	return nil
}

func (d *DockerPlugin) Cleanup() error {
	return nil
}

func (d *DockerPlugin) listContainers(e *core.EditorCore, args []string) error {
	// Run docker ps command
	// Display results in the editor
	return nil
}

func (d *DockerPlugin) viewLogs(e *core.EditorCore, args []string) error {
	// Get container name from args
	// Run docker logs
	// Display in buffer
	return nil
}

// CRITICAL: This exported function is how the plugin loader gets your plugin
func NewPlugin() core.Plugin {
	return &DockerPlugin{}
}
