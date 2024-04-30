package main

import (
	"Stage-2024-dashboard/pkg/settings"
	"fmt"

	"gopkg.in/yaml.v3"
)

func main() {
	s := settings.Settings{}
	s.SetDefault()
	y, _ := yaml.Marshal(&s)
	fmt.Print(string(y))
}
