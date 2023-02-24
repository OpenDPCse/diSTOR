package service

import (
	"context"
	"sardines/tool"
	"sync"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/phrara/go-util/tlv"
)

// KeyTable Distribution
// KeyTable will be distributed periodically
// When peers get the distributed KeyTable, they use then to update their own KeyTable
// This automatically renew the KeyTable info of the decentralized network

func KeyTableDistributeHandler(s network.Stream) {
	pn, _ := tool.ParsePeerNode(s.Conn().RemoteMultiaddr().String() + "/p2p/" + s.Conn().RemotePeer().String())
	serv.router.AddNode(pn)

	//fmt.Println("Get a distributed router table from", pn.String())

	p := &tlv.Packet{}
	err2 := p.Load(s)
	if err2 != nil {
		return
	}

	defer s.Close()

	if p.Tag == 2 {
		serv.ktab.AppendBatchRaw(p.Value)
		p = tlv.New(1, []byte("acc"))
		wrap, _ := p.Wrap()
		s.Write(wrap)
	} else {
		p = tlv.New(0, []byte("err"))
		wrap, _ := p.Wrap()
		s.Write(wrap)
	}

}

type errCounter struct {
	errNum int
	mu     sync.Mutex
}

func (e *errCounter) count() {
	e.mu.Lock()
	e.errNum++
	e.mu.Unlock()
}

func (s *Service) KeyTableDistribute() int {

	if s.router.Sum() <= 1 {
		return 0
	}

	var wg sync.WaitGroup
	ec := errCounter{
		errNum: 0,
		mu:     sync.Mutex{},
	}
	nodes := serv.router.AllNodes()[1:]
	for _, pn := range nodes {

		wg.Add(1)
		go func(p *tool.PeerNode) {
			defer wg.Done()

			stream, err2 := s.Host.NewStream(context.Background(), p.ID(), KD)
			if err2 != nil {
				ec.count()
				return
			}

			packet := tlv.New(2, s.ktab.GetAllRaw())
			wrap, _ := packet.Wrap()

			stream.Write(wrap)

			rcvp := &tlv.Packet{}
			err2 = rcvp.Load(stream)
			if err2 != nil {
				ec.count()
				return
			}

			if rcvp.Tag == 1 {
				return
			} else {
				ec.count()
				return
			}

		}(pn)
	}

	wg.Wait()
	return ec.errNum
}
