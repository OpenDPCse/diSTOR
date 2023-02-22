package router

import (
	"container/list"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"sardines/tool"
)

var rt *Router

func init() {
	rt = &Router{}
}

type Table map[int]*list.List

type Router struct {
	host host.Host
}

func New(hn host.Host) *Router {
	rt.host = hn
	return rt
}

func (r *Router) AddNode(pn *tool.PeerNode) {
	r.host.Peerstore().AddAddrs(pn.ID(), pn.NodeInfo.Addrs, peerstore.PermanentAddrTTL)
}

func (r *Router) DelNode(pn *tool.PeerNode) {
	r.host.Peerstore().ClearAddrs(pn.ID())
	r.host.Peerstore().RemovePeer(pn.ID())
}

func (r *Router) AllNodes() []*tool.PeerNode {

	nodes := make([]*tool.PeerNode, 0, 10)

	for _, v := range r.host.Peerstore().Peers() {
		info := r.host.Peerstore().PeerInfo(v)
		addr, _ := peer.AddrInfoToP2pAddrs(&info)
		node := tool.NewPeerNode(&info, addr[0])
		nodes = append(nodes, node)
	}
	return nodes
}

func (r *Router) Sum() int {
	return r.host.Peerstore().Peers().Len()
}
