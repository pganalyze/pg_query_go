// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node TruncateStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("TruncateStmt")

	if int(node.Behavior) != 0 {
		ctx.WriteString("behavior")
		ctx.WriteString(strconv.Itoa(int(node.Behavior)))
	}

	if len(node.Relations.Items) > 0 {
		ctx.WriteString("relations")
		node.Relations.Fingerprint(ctx, "Relations")
	}

	if node.RestartSeqs {
		ctx.WriteString("restart_seqs")
		ctx.WriteString(strconv.FormatBool(node.RestartSeqs))
	}
}
