package presentation

import "github.com/m7medVision/lazypodman/pkg/commands"

func GetProjectDisplayStrings(project *commands.Project) []string {
	return []string{project.Name}
}
