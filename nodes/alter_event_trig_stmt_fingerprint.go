// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterEventTrigStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ALTEREVENTTRIGSTMT")

	if node.Trigname != nil {
		io.WriteString(ctx.hash, *node.Trigname)
	}
}
