package p2p

type HandShaker interface {
	Handshake() error
}

type HandShakeFunc func(Peer) error

type DefaultHandShaker struct {
}

func NOPHandshakeFunc(Peer) error {
	return nil
}
