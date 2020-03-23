package whoami

import (
	"context"

	"github.com/miekg/dns"
)

type Whoami struct{}

// Name implements the Handler interface.
func (wh Whoami) Name() string { return "whoami" }

// ServeDNS implements the plugin.Handler interface.
func (wh Whoami) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	return 0, nil
}
