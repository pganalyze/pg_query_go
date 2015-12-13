// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterTableSpaceOptionsStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("AlterTableSpaceOptionsStmt")
	ctx.WriteString(strconv.FormatBool(node.IsReset))

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Tablespacename != nil {
		ctx.WriteString(*node.Tablespacename)
	}
}
