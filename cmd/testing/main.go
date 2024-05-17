package main

import (
	"Stage-2024-dashboard/pkg/alert"
	"Stage-2024-dashboard/pkg/database"
	"Stage-2024-dashboard/pkg/helper"
	"Stage-2024-dashboard/pkg/settings"
)

func main() {
	set, err := settings.Load()
	helper.MaybeDie(err, "Failed to load configs")

	q := database.NewQueries(set.Database)

	err = alert.CheckDeltas(set.Alert, q)

	helper.MaybeDieErr(err)
}
