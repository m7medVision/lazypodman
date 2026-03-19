package presentation

import (
	"github.com/fatih/color"
	"github.com/m7medVision/lazypodman/pkg/commands"
	"github.com/m7medVision/lazypodman/pkg/config"
	"github.com/m7medVision/lazypodman/pkg/utils"
)

func GetServiceDisplayStrings(guiConfig *config.GuiConfig, service *commands.Service) []string {
	if service.Container == nil {
		var containerState string
		switch guiConfig.ContainerStatusHealthStyle {
		case "short":
			containerState = "n"
		case "icon":
			containerState = "."
		case "long":
			fallthrough
		default:
			containerState = "none"
		}

		return []string{
			utils.ColoredString(containerState, color.FgBlue),
			"",
			service.Name,
			"",
			"",
			"",
		}
	}

	container := service.Container
	return []string{
		getContainerDisplayStatus(guiConfig, container),
		getContainerDisplaySubstatus(guiConfig, container),
		service.Name,
		getDisplayCPUPerc(container),
		utils.ColoredString(displayPorts(container), color.FgYellow),
		utils.ColoredString(displayContainerImage(container), color.FgMagenta),
	}
}
