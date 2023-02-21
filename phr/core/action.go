package core

import (
	"context"
	"errors"
	"io"
	"os/exec"
	"sardines/err"
	"strings"
	"time"
)

func (h *HostNode) JoinNetwork() <-chan int {
	msg := make(chan int, 1)
	for _, bn := range BootstrapNodes {
		node := bn
		if node == nil {
			msg <- 1
			continue
		}
		go func() {
			res := <-h.Serv.Ping(node)
			if res.Error != nil {
				msg <- 0
				return
			} else {
				// ping 通了
				// 发出加入申请
				if b := h.Serv.JoinApply(node); b {
					msg <- 1
					h.Router.AddNode(node)
				} else {
					msg <- 0
				}
			}
		}()
	}
	return msg
}

func (h *HostNode) Close() error {
	err := h.Host.Close()
	h.Ktab.Close()
	if h.ipfs != nil {
		h.ipfs.Process.Kill()
	}
	return err
}

func (h *HostNode) RunIpfs() error {
	_, er := exec.Command("ipfs", "init").Output()
	if er != nil {
		return er
	}
	cmd := exec.Command("ipfs", "daemon")
	rc, err2 := cmd.StdoutPipe()
	if err2 != nil {
		return err2
	}
	if err2 = cmd.Start(); err2 != nil {
		return err2
	}

	if b, err3 := io.ReadAll(rc); err3 != nil {
		return err3
	} else {
		if strings.Contains(string(b), "ready") {
			h.ipfs = cmd
			cmd.Wait()
		} else {
			cmd.Process.Kill()
			return errors.New("ipfs daemon run failed")
		}
	}
	return nil

}

func (h *HostNode) keyTableDistribute(ctx context.Context, period time.Duration, errs chan<- error) {
	ticker := time.NewTicker(time.Second * period)
	for range ticker.C {

		errNum := h.Serv.KeyTableDistribute()
		if errNum >= h.Router.Sum() {
			errs <- err.KeyTableDistributeException
		}

		select {
		case <-ctx.Done():
			ticker.Stop()
			close(errs)
			return
		default:
		}
	}
}

func (h *HostNode) autoPeriod(nodeNum int) time.Duration {
	switch {
	case nodeNum >= 100:
		return 60
	case nodeNum >= 50:
		return 30
	case nodeNum >= 25:
		return 15
	default:
		return 10
	}
}

// KeyTableDistributeOn
// The unit of argument `period` is second.
// The argument `period` will be affective if `auto` is true,
// and it'll be useless if `auto` is false.
func (h *HostNode) KeyTableDistributeOn(auto bool, period int, ctx context.Context) <-chan error {
	errs := make(chan error, 15)
	if auto {
		go h.keyTableDistribute(ctx, h.autoPeriod(h.Router.Sum()), errs)
	} else {
		go h.keyTableDistribute(ctx, time.Duration(period), errs)
	}
	return errs
}
