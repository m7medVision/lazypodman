package presentation

import "github.com/mohammed/lazypodman/pkg/commands"

func GetVolumeDisplayStrings(volume *commands.Volume) []string {
	return []string{volume.Volume.Driver, volume.Name}
}
