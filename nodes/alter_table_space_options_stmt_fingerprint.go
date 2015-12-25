// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterTableSpaceOptionsStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterTableSpaceOptionsStmt")
	ctx.WriteString(strconv.FormatBool(node.IsReset))
	node.Options.Fingerprint(ctx, "Options")

	if node.Tablespacename != nil {
		ctx.WriteString(*node.Tablespacename)
	}
}
