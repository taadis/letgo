package netx

import (
	"net"
	"testing"
)

func TestNetDNS(t *testing.T) {
	// 通过域名查询IPv4/IPv6信息
	t.Run("LookupIP", func(t *testing.T) {
		ips, err := net.LookupIP("baidu.com")
		if err != nil {
			t.Fatal(err)
		}
		for _, ip := range ips {
			t.Log(ip.String())
		}
	})

	// 查询CNAME信息
	t.Run("LookupCNAME", func(t *testing.T) {
		cname, err := net.LookupCNAME("m.baidu.com")
		if err != nil {
			t.Fatal(err)
		}
		t.Log(cname)
	})

	// 查询NameServer(NS)
	t.Run("LookupNS", func(t *testing.T) {
		ns, err := net.LookupNS("baidu.com")
		if err != nil {
			t.Fatal(err)
		}
		for _, n := range ns {
			t.Log(n.Host)
		}
		// output:
		// ns4.baidu.com.
		// dns.baidu.com.
		// ns3.baidu.com.
		// ns7.baidu.com.
		// ns2.baidu.com.
	})

	t.Run("LookupMX", func(t *testing.T) {
		mxs, err := net.LookupMX("baidu.com")
		if err != nil {
			t.Fatal(err)
		}
		for _, mx := range mxs {
			t.Log(mx.Host, mx.Pref)
		}
	})

	t.Run("LookupTXT", func(t *testing.T) {
		txtRecords, err := net.LookupTXT("baidu.com")
		if err != nil {
			t.Fatal(err)
		}
		for _, txt := range txtRecords {
			t.Log(txt)
		}
	})

	// 上面是使用标准库net获取DNS信息,如果想要得到更多数据,可以借助三方库
	// 如:https://github.com/miekg/dns
}
