// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ViewStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ViewStmt")
	if len(node.Aliases.Items) > 0 {
		ctx.WriteString("aliases")
		node.Aliases.Fingerprint(ctx, "Aliases")
	}

	if len(node.Options.Items) > 0 {
		ctx.WriteString("options")
		node.Options.Fingerprint(ctx, "Options")
	}

	if node.Query != nil {
		ctx.WriteString("query")
		node.Query.Fingerprint(ctx, "Query")
	}

	if node.Replace {
		ctx.WriteString("replace")
		ctx.WriteString(strconv.FormatBool(node.Replace))
	}

	if node.View != nil {
		ctx.WriteString("view")
		node.View.Fingerprint(ctx, "View")
	}

	if int(node.WithCheckOption) != 0 {
		ctx.WriteString("withCheckOption")
		ctx.WriteString(strconv.Itoa(int(node.WithCheckOption)))
	}
}
