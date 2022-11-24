package helper

import (
	"github.com/rapid7/go-get-proxied/proxy"
)

func GetSystempConfig(reqUrl string) (PROXYSTR string, USERPASS string, err error) {
	p := proxy.NewProvider("")

	proxy_ := p.GetHTTPSProxy(reqUrl)
	PROXYSTR = proxy_.Host()
	username, _ := proxy_.Username()
	password, _ := proxy_.Password()
	USERPASS = username + ":" + password
	return
}
func GetIEProxyConfig(reqUrl string) (PROXYSTR string, USERPASS string, err error) {
	// http.DefaultTransport.(*http.Transport).Proxy =

	// url := ieproxy.GetProxyFunc()

	// fmt.Printf("URL IE" + url)

	return
}
