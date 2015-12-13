// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node DropStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("DropStmt")

	for _, subNode := range node.Arguments {
		subNode.Fingerprint(ctx, "Arguments")
	}

	ctx.WriteString(strconv.Itoa(int(node.Behavior)))
	ctx.WriteString(strconv.FormatBool(node.Concurrent))
	ctx.WriteString(strconv.FormatBool(node.MissingOk))

	for _, subNode := range node.Objects {
		subNode.Fingerprint(ctx, "Objects")
	}

	ctx.WriteString(strconv.Itoa(int(node.RemoveType)))
}
