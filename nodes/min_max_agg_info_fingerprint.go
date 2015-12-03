// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node MinMaxAggInfo) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "MinMaxAggInfo")
	if node.Param != nil {
		node.Param.Fingerprint(ctx)
	}

	if node.Path != nil {
		node.Path.Fingerprint(ctx)
	}

	if node.Subroot != nil {
		node.Subroot.Fingerprint(ctx)
	}

	if node.Target != nil {
		node.Target.Fingerprint(ctx)
	}
}
