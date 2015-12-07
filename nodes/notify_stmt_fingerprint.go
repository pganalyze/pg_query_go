// Auto-generated - DO NOT EDIT

package pg_query

func (node NotifyStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("NOTIFYSTMT")

	if node.Conditionname != nil {
		ctx.WriteString(*node.Conditionname)
	}

	if node.Payload != nil {
		ctx.WriteString(*node.Payload)
	}
}
