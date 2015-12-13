// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterEventTrigStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterEventTrigStmt")
	ctx.WriteString(string(node.Tgenabled))

	if node.Trigname != nil {
		ctx.WriteString(*node.Trigname)
	}
}
