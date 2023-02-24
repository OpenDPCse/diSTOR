package service

import (
	"context"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/phrara/go-util/tlv"
	"sardines/tool"
)

func JoinApplyHandler(s network.Stream) {
	pn, _ := tool.ParsePeerNode(s.Conn().RemoteMultiaddr().String() + "/p2p/" + s.Conn().RemotePeer().String())
	// fmt.Println("receive a join application from", pn.String())

	// 节点加入路由表
	serv.router.AddNode(pn)

	p := tlv.New(1, []byte("acc"))
	data, _ := p.Wrap()
	s.Write(data)

}

func (s *Service) JoinApply(bootstrapNode *tool.PeerNode) bool {
	s.router.AddNode(bootstrapNode)

	stream, err2 := s.Host.NewStream(context.Background(), bootstrapNode.ID(), JOIN)
	if err2 != nil {
		return false
	}

	packet := &tlv.Packet{}
	err2 = packet.Load(stream)
	if err2 != nil {
		return false
	}

	defer stream.Close()

	if packet.Tag == 1 {
		return true
	} else {
		return false
	}
}
