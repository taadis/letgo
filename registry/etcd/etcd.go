package etcd

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"hash"
	"log"
	"net"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/taadis/letgo/registry"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var prefix = "/micro/registry"

type etcdRegistry struct {
	client  *clientv3.Client
	options registry.Options

	sync.RWMutex
	register map[string]uint64
	leases   map[string]clientv3.LeaseID
}

func (r *etcdRegistry) String() string {
	return "etcd"
}

func NewRegistry(opts ...registry.Option) registry.Registry {
	r := new(etcdRegistry)
	r.options = registry.Options{}
	r.register = make(map[string]uint64)
	r.leases = make(map[string]clientv3.LeaseID)

	username := os.Getenv("ETCD_USERNAME")
	password := os.Getenv("ETCD_PASSWORD")
	address := os.Getenv("MICRO_REGISTRY_ADDRESS")
	if len(username) > 0 && len(password) > 0 {
		opts = append(opts, Auth(username, password))
	}
	if len(address) > 0 {
		opts = append(opts, registry.WithAddress(address))
	}
	configure(r, opts...)
	return r
}

func configure(r *etcdRegistry, opts ...registry.Option) error {
	config := clientv3.Config{}
	config.Endpoints = []string{"127.0.0.1:2379"}

	for _, o := range opts {
		o(&r.options)
	}

	if r.options.Timeout == 0 {
		r.options.Timeout = 5 * time.Second
	}
	config.DialTimeout = r.options.Timeout

	if r.options.Secure || r.options.TLSConfig != nil {
		tlsConfig := r.options.TLSConfig
		if tlsConfig == nil {
			tlsConfig = &tls.Config{
				InsecureSkipVerify: true,
			}
		}
		config.TLS = tlsConfig
	}

	var addrList []string
	for _, addr := range r.options.Address {
		if len(addr) == 0 {
			continue
		}
		host, port, err := net.SplitHostPort(addr)
		if ae, ok := err.(*net.AddrError); ok && ae.Err == "missing port in address" {
			host = addr
			port = "2379"
		}
		addrList = append(addrList, net.JoinHostPort(host, port))
	}

	// if we got addrList then we will update
	if len(addrList) > 0 {
		config.Endpoints = addrList
	}

	c, err := clientv3.New(config)
	if err != nil {
		return err
	}
	r.client = c
	return nil
}

func encode(s *registry.Service) string {
	bs, _ := json.Marshal(s)
	return string(bs)
}

func decode(bs []byte) *registry.Service {
	var s *registry.Service
	_ = json.Unmarshal(bs, &s)
	return s
}

func nodePath(s, id string) string {
	service := strings.Replace(s, "/", "-", -1)
	node := strings.Replace(id, "/", "-", -1)
	return path.Join(prefix, service, node)
}

func (r *etcdRegistry) Init(opts ...registry.Option) error {
	return configure(r, opts...)
}

func (r *etcdRegistry) Options() registry.Options {
	return r.options
}

func (r *etcdRegistry) registerNodes(s *registry.Service, node *registry.Node, opts ...registry.RegistryOption) error {
	if len(s.Nodes) == 0 {
		return errors.New("require at least one node")
	}

	// check existing lease cache
	r.RLock()
	leaseID, ok := r.leases[s.Name+node.Id]
	r.RUnlock()

	if !ok {
		// todo:https://github.com/asim/go-micro/blob/master/plugins/registry/etcd/etcd.go#L166
	}

	var leaseNotFound bool

	// renew the lease if it exists
	if leaseID > 0 {
		log.Printf("renewing existing lease for %s %d", s.Name, leaseID)
		_, err := r.client.KeepAliveOnce(context.TODO(), leaseID)
		if err != nil {
			log.Fatalf("keep alive once error:%+v", err)
			return err
		}
		// lease not found do register
		leaseNotFound = true
	}

	// create hash of service; unit64
	h, err := hash.Hash(node, nil)
	if err != nil {
		log.Fatalf("create service hash error:%+v", err)
		return fmt.Errorf("create service has error:%+v", err)
	}
}

func (r *etcdRegistry) Register(s *registry.Service, opts ...registry.RegistryOption) error {
	if len(s.Nodes) == 0 {
		return errors.New("require at least one node")
	}

	var lastErr error

	// register each node individually
	for _, node := range s.Nodes {
		err := r.registerNode(s, node, opts...)
		if err != nil {
			lastErr = err
		}
	}

	return lastErr
}

func (r *etcdRegistry) Deregister(s *registry.Service, opts ...registry.DeregisterOption) error {
	if len(s.Nodes) == 0 {
		return errors.New("require at least one node")
	}

	for _, node := range s.Nodes {
		r.Lock()
		// delete our hash of the service
		delete(r.register, s.Name+node.Id)
		// delete our lease of the service
		delete(r.leases, s.Name+node.Id)
		r.Unlock()

		ctx, cancel := context.WithTimeout(context.Background(), r.options.Timeout)
		defer cancel()

		key := nodePath(s.Name, node.Id)
		_, err := r.client.Delete(ctx, key)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *etcdRegistry) GetService(name string, opts ...registry.GetOption)

func (r *etcdRegistry) Watch(opts ...registry.WatchOption) (registry.Watcher, error) {
	return newEtcdWatcher(r, r.options.Timeout, opts)
}
