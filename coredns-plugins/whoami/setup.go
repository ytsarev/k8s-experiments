package whoami

import "github.com/coredns/coredns/plugin"

func init() {
	plugin.Register("whoami", setup)
}
