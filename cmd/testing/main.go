package main

import (
	"Stage-2024-dashboard/pkg/alert"
	"Stage-2024-dashboard/pkg/helper"
	"Stage-2024-dashboard/pkg/settings"
)

func main() {
	set, err := settings.Load()
	helper.MaybeDie(err, "Failed to load configs")

	err = alert.SendSlackNotification(set.Alert, "dit is een test :3")
	helper.MaybeDieErr(err)
}
