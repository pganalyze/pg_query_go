// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ViewStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("VIEWSTMT")

	for _, subNode := range node.Aliases {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Query != nil {
		node.Query.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.Replace))

	if node.View != nil {
		node.View.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.WithCheckOption)))
}
