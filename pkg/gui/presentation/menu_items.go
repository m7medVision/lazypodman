package presentation

import "github.com/m7medVision/lazypodman/pkg/gui/types"

func GetMenuItemDisplayStrings(menuItem *types.MenuItem) []string {
	return menuItem.LabelColumns
}
