package presentation

import "github.com/mohammed/lazypodman/pkg/commands"

func GetNetworkDisplayStrings(network *commands.Network) []string {
	return []string{network.Network.Driver, network.Name}
}
