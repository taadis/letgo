package etcd

import (
	"time"

	"go.etcd.io/etcd/client/v3"
)

type etcdWatcher struct {
	watchChan clientv3.WatchChan
	client    *clientv3.Client
	timeout   time.Duration
	stopChan  chan bool
}

func newEtcd

func (w *etcdWatcher) Next() *registry.Resu

func (w *etcdWatcher) Stop() {
	select {
	case <-w.stopChan:
		return
	default:
		close(w.stopChan)
	}
}
