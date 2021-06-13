package commands

import (
	"github.com/avocatl/admiral/pkg/commander"
)

func profile() *commander.Command {
	p := commander.Builder(
		nil,
		commander.Config{
			Namespace: "profiles",
			ShortDesc: "In order to process payments, you need to create a website profile",
		},
		getProfileCols(),
	)

	getProfileCmd(p)
	currentProfileCmd(p)

	return p
}

func getProfileCols() []string {
	return []string{
		"RESOURCE",
		"ID",
		"MODE",
		"NAME",
		"WEBSITE",
		"EMAIL",
		"PHONE",
		"CATEGORY_CODE",
		"STATUS",
		"REVIEW",
		"CREATED_AT",
	}
}
