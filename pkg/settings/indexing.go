package settings

import "time"

type Indexing struct {
	Interval time.Duration `yaml:"interval"`
}

func (i *Indexing) SetDefault() {
	i.Interval = time.Hour
}
