// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateSeqStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateSeqStmt")
	if len(node.Options.Items) > 0 {
		ctx.WriteString("options")
		node.Options.Fingerprint(ctx, "Options")
	}

	if node.OwnerId != 0 {
		ctx.WriteString("ownerId")
		ctx.WriteString(strconv.Itoa(int(node.OwnerId)))
	}

	if node.Sequence != nil {
		ctx.WriteString("sequence")
		node.Sequence.Fingerprint(ctx, "Sequence")
	}
}
