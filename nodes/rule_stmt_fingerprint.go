// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RuleStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("RuleStmt")
	node.Actions.Fingerprint(ctx, "Actions")
	ctx.WriteString(strconv.Itoa(int(node.Event)))
	ctx.WriteString(strconv.FormatBool(node.Instead))

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx, "Relation")
	}

	ctx.WriteString(strconv.FormatBool(node.Replace))

	if node.Rulename != nil {
		ctx.WriteString(*node.Rulename)
	}

	if node.WhereClause != nil {
		node.WhereClause.Fingerprint(ctx, "WhereClause")
	}
}
