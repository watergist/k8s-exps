package listener

type Listener struct {
	Address  string
	Port     string
	Protocol string
}

type Server struct {
	UpstreamTimeout    int
	ListenerProperties *Listener
}
