// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node NotifyStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "NOTIFYSTMT")

	if node.Conditionname != nil {
		io.WriteString(ctx.hash, *node.Conditionname)
	}

	if node.Payload != nil {
		io.WriteString(ctx.hash, *node.Payload)
	}
}
