// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ViewStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ViewStmt")

	for _, subNode := range node.Aliases {
		subNode.Fingerprint(ctx, "Aliases")
	}

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx, "Options")
	}

	if node.Query != nil {
		node.Query.Fingerprint(ctx, "Query")
	}

	ctx.WriteString(strconv.FormatBool(node.Replace))

	if node.View != nil {
		node.View.Fingerprint(ctx, "View")
	}

	ctx.WriteString(strconv.Itoa(int(node.WithCheckOption)))
}
