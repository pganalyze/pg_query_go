// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RuleStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("RuleStmt")
	if len(node.Actions.Items) > 0 {
		ctx.WriteString("actions")
		node.Actions.Fingerprint(ctx, "Actions")
	}

	if int(node.Event) != 0 {
		ctx.WriteString("event")
		ctx.WriteString(strconv.Itoa(int(node.Event)))
	}

	if node.Instead {
		ctx.WriteString("instead")
		ctx.WriteString(strconv.FormatBool(node.Instead))
	}

	if node.Relation != nil {
		ctx.WriteString("relation")
		node.Relation.Fingerprint(ctx, "Relation")
	}

	if node.Replace {
		ctx.WriteString("replace")
		ctx.WriteString(strconv.FormatBool(node.Replace))
	}

	if node.Rulename != nil {
		ctx.WriteString("rulename")
		ctx.WriteString(*node.Rulename)
	}

	if node.WhereClause != nil {
		ctx.WriteString("whereClause")
		node.WhereClause.Fingerprint(ctx, "WhereClause")
	}
}
