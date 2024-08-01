package main

import (
	"factorio-mods-manager/factorio"
	"fmt"
)

func main() {
	mod, err := factorio.GetMod("SimpleRadioTowers")

	if err != nil {
		panic(err)
	}

	fmt.Printf(
		"ID: %s\nTitle: %s\nAuthor: %s\nSummary:\n%s\nDescription:\n%s\nReleases:\n",
		mod.ID, mod.Title, mod.Author, mod.Summary, mod.Description,
	)

	for _, release := range mod.Releases {
		fmt.Printf(
			"  Version: %s\n  Factorio version: %s\n  Base mod version: %s\n  Dependencies:\n",
			release.Version.String(), release.FactorioVersion.String(), release.BaseVersion.String(),
		)

		for _, dependency := range release.Dependencies {
			fmt.Printf(
				"    ID: %s\n    Version: %s\n    Type: %s\n",
				dependency.ModID, dependency.Version.String(), dependency.Type.String(),
			)
		}
	}
}
