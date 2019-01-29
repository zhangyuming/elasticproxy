package runner

import "elasticproxy/proxy"

func init() {
	proxy.RegistryRequestModifyer(&kibana{})
}
