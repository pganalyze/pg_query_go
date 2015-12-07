// Auto-generated - DO NOT EDIT

package pg_query

func (node FieldSelect) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("FIELDSELECT")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}
}
