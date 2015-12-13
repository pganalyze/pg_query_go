// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node DropStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("DropStmt")

	for _, subNode := range node.Arguments {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Behavior)))
	ctx.WriteString(strconv.FormatBool(node.Concurrent))
	ctx.WriteString(strconv.FormatBool(node.MissingOk))

	for _, subNode := range node.Objects {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.RemoveType)))
}
