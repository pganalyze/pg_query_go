// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node SubqueryScan) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "SUBQUERYSCAN")

	if node.Subplan != nil {
		node.Subplan.Fingerprint(ctx)
	}
}
