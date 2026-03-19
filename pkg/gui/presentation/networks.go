package presentation

import "github.com/m7medVision/lazypodman/pkg/commands"

func GetNetworkDisplayStrings(network *commands.Network) []string {
	return []string{network.Network.Driver, network.Name}
}
