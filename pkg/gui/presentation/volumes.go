package presentation

import "github.com/m7medVision/lazypodman/pkg/commands"

func GetVolumeDisplayStrings(volume *commands.Volume) []string {
	return []string{volume.Volume.Driver, volume.Name}
}
