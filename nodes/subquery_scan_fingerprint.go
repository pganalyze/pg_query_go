// Auto-generated - DO NOT EDIT

package pg_query

func (node SubqueryScan) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("SUBQUERYSCAN")

	if node.Subplan != nil {
		node.Subplan.Fingerprint(ctx)
	}
}
