package settings

type Logger struct {
	Level string `yaml:"level"`
}

func (l *Logger) SetDefault() {
	l.Level = "INFO"
}
