// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterSeqStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterSeqStmt")

	if node.MissingOk {
		ctx.WriteString("missing_ok")
		ctx.WriteString(strconv.FormatBool(node.MissingOk))
	}

	if len(node.Options.Items) > 0 {
		ctx.WriteString("options")
		node.Options.Fingerprint(ctx, "Options")
	}

	if node.Sequence != nil {
		ctx.WriteString("sequence")
		node.Sequence.Fingerprint(ctx, "Sequence")
	}
}
