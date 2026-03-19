package presentation

import "github.com/mohammed/lazypodman/pkg/gui/types"

func GetMenuItemDisplayStrings(menuItem *types.MenuItem) []string {
	return menuItem.LabelColumns
}
