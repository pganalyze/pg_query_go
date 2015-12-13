// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RuleStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("RuleStmt")

	for _, subNode := range node.Actions {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Event)))
	ctx.WriteString(strconv.FormatBool(node.Instead))

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.Replace))

	if node.Rulename != nil {
		ctx.WriteString(*node.Rulename)
	}

	if node.WhereClause != nil {
		node.WhereClause.Fingerprint(ctx)
	}
}
