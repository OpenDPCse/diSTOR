package tool

import (
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
)

type PeerNode struct {
	// 节点信息
	NodeInfo *peer.AddrInfo
	// p2p节点标识
	NodeAddr multiaddr.Multiaddr
}

func NewPeerNode(info *peer.AddrInfo, addr multiaddr.Multiaddr) *PeerNode {
	return &PeerNode{
		NodeInfo: info,
		NodeAddr: addr,
	}
}

func (p *PeerNode) String() string {
	return p.NodeAddr.String()
}

func (p *PeerNode) ID() peer.ID {
	return p.NodeInfo.ID
}

func ParsePeerNode(p string) (*PeerNode, error) {
	if p == "" {
		return nil, nil
	}
	addr, err1 := multiaddr.NewMultiaddr(p)
	if err1 != nil {
		return nil, err1
	}
	info, err1 := peer.AddrInfoFromP2pAddr(addr)
	if err1 != nil {
		return nil, err1
	}
	return &PeerNode{
		NodeInfo: info,
		NodeAddr: addr,
	}, nil
}
