// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterEventTrigStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("ALTEREVENTTRIGSTMT")

	if node.Trigname != nil {
		ctx.WriteString(*node.Trigname)
	}
}
