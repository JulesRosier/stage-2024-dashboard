package settings

type Server struct {
	Port int    `yaml:"port"`
	Bind string `yaml:"bind"`
}

func (s *Server) SetDefault() {
	s.Port = 3000
	s.Bind = "127.0.0.1"
}
