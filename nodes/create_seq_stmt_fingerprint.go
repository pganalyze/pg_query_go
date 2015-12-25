// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateSeqStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateSeqStmt")
	node.Options.Fingerprint(ctx, "Options")
	ctx.WriteString(strconv.Itoa(int(node.OwnerId)))

	if node.Sequence != nil {
		node.Sequence.Fingerprint(ctx, "Sequence")
	}
}
