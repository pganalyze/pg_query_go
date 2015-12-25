// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterSeqStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterSeqStmt")
	ctx.WriteString(strconv.FormatBool(node.MissingOk))
	node.Options.Fingerprint(ctx, "Options")

	if node.Sequence != nil {
		node.Sequence.Fingerprint(ctx, "Sequence")
	}
}
