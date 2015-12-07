// Auto-generated - DO NOT EDIT

package pg_query

func (node Limit) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("LIMIT")

	if node.LimitCount != nil {
		node.LimitCount.Fingerprint(ctx)
	}

	if node.LimitOffset != nil {
		node.LimitOffset.Fingerprint(ctx)
	}
}
