// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterTableSpaceOptionsStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterTableSpaceOptionsStmt")

	if node.IsReset {
		ctx.WriteString("isReset")
		ctx.WriteString(strconv.FormatBool(node.IsReset))
	}

	if len(node.Options.Items) > 0 {
		ctx.WriteString("options")
		node.Options.Fingerprint(ctx, "Options")
	}

	if node.Tablespacename != nil {
		ctx.WriteString("tablespacename")
		ctx.WriteString(*node.Tablespacename)
	}
}
