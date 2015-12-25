// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ViewStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ViewStmt")
	node.Aliases.Fingerprint(ctx, "Aliases")
	node.Options.Fingerprint(ctx, "Options")

	if node.Query != nil {
		node.Query.Fingerprint(ctx, "Query")
	}

	ctx.WriteString(strconv.FormatBool(node.Replace))

	if node.View != nil {
		node.View.Fingerprint(ctx, "View")
	}

	ctx.WriteString(strconv.Itoa(int(node.WithCheckOption)))
}
