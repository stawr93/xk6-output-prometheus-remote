// Package remotewrite registers the xk6-output-prometheus-remote extension
package remotewrite

import (
	"github.com/stawr93/xk6-output-prometheus-remote/pkg/remotewrite"
	"go.k6.io/k6/output"
)

func init() {
	output.RegisterExtension("xk6-prometheus-rw-forked", func(p output.Params) (output.Output, error) {
		return remotewrite.New(p)
	})
}
