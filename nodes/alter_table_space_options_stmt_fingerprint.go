// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterTableSpaceOptionsStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterTableSpaceOptionsStmt")
	ctx.WriteString(strconv.FormatBool(node.IsReset))

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx, "Options")
	}

	if node.Tablespacename != nil {
		ctx.WriteString(*node.Tablespacename)
	}
}
