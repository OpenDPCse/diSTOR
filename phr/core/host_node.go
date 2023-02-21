package core

import (
	"context"
	"errors"
	"os/exec"
	"sardines/config"
	"sardines/router"
	"sardines/service"
	"sardines/storage"
	"sardines/tool"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
)

// BootstrapNodes 引导节点
var BootstrapNodes []*tool.PeerNode

func getBootstrapNodes(bn string) ([]*tool.PeerNode, error) {
	bsn := make([]*tool.PeerNode, 0)
	pn, err1 := tool.ParsePeerNode(bn)
	if err1 != nil {
		return bsn, err1
	}
	bsn = append(bsn, pn)
	return bsn, nil
}

type HostNode struct {
	// p2p节点
	Host host.Host
	// 节点信息
	NodeInfo *peer.AddrInfo
	// p2p节点标识
	NodeAddr multiaddr.Multiaddr
	// context
	Ctx context.Context
	// 相关协议服务
	Serv *service.Service
	// 路由表
	Router *router.Router
	// 倒排索引表
	Ktab *storage.KeyTable
	// ipfs daemon
	ipfs *exec.Cmd
	// ipfs api
	api *API
}

func GenerateNode() (*HostNode, error) {

	// 读取配置
	c := &config.Config{}
	err2 := c.LoadAll()
	if err2 != nil {
		return nil, errors.New("your node haven't been configure correctly, please use -help for more guidance")
	}
	node := new(HostNode)
	node.Ctx = context.Background()

	// 初始化引导节点
	BootstrapNodes, err2 = getBootstrapNodes(c.BootstrapNode)
	if err2 != nil {
		return nil, err2
	}

	// 获取节点
	h, err := libp2p.New(
		libp2p.ListenAddrStrings(c.AddrString()),
		libp2p.Identity(c.PrvKey),
		libp2p.Ping(false),
	)
	if err != nil {
		return nil, err
	}
	node.Host = h

	// 获取节点信息
	node.NodeInfo = &peer.AddrInfo{
		ID:    h.ID(),
		Addrs: h.Addrs(),
		// 获取节点标识
	}

	addrs, err := peer.AddrInfoToP2pAddrs(node.NodeInfo)
	if err != nil {
		return nil, err
	}

	node.NodeAddr = addrs[0]

	// 初始化路由表
	node.Router = router.New(node.Host)
	// 初始化倒排索引表
	node.Ktab, err = storage.NewKeyTab(config.Ktab)
	// 初始化协议服务
	node.Serv = service.New(node.Host, node.Router, node.Ktab).ServiceHandlerRegister()
	if err != nil {
		return nil, err
	}
	// 获取ipfs-api
	node.api = NewAPI()

	return node, nil
}
