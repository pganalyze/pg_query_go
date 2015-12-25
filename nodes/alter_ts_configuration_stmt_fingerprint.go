// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterTSConfigurationStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterTSConfigurationStmt")
	node.Cfgname.Fingerprint(ctx, "Cfgname")
	node.Dicts.Fingerprint(ctx, "Dicts")
	ctx.WriteString(strconv.FormatBool(node.MissingOk))
	ctx.WriteString(strconv.FormatBool(node.Override))
	ctx.WriteString(strconv.FormatBool(node.Replace))
	node.Tokentype.Fingerprint(ctx, "Tokentype")
}
