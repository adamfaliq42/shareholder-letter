//sort companies name lexicographically

package main

import (
	"fmt"
	"sort"
)

func main() {
	company := []string{"Amazon", "Berkshire Hathaway", "Biglari", "Nelnet", "Markel", "JPMorgan", "Southwester Energy", "Starwood", "Trupanion",
		"Tesla", "Strayer", "Koppers", "Old National", "Standard Motors", "Kulicke and Soffa", "Sykes", "Inogen", "Allegiant", "Acadia", "Dime", "Waddell And Reed"}
	sort.Strings(company)
	fmt.Println(company)
}
