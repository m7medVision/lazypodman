package presentation

import (
	"github.com/m7medVision/lazypodman/pkg/commands"
	"github.com/m7medVision/lazypodman/pkg/utils"
)

func GetImageDisplayStrings(image *commands.Image) []string {
	return []string{
		image.Name,
		image.Tag,
		utils.FormatDecimalBytes(int(image.Image.Size)),
	}
}
