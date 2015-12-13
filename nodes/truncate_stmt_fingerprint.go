// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node TruncateStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("TruncateStmt")
	ctx.WriteString(strconv.Itoa(int(node.Behavior)))

	for _, subNode := range node.Relations {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.RestartSeqs))
}
