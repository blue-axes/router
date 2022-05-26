package rpc

type (
	Config struct {
	}
	Server struct {
	}
)

func NewServer(cfg Config) *Server {
	s := &Server{}

	return s
}

func (s Server) Serve() error {
	//TODO implement me
	panic("implement me")
}

func (s Server) Shutdown() error {
	//TODO implement me
	panic("implement me")
}
