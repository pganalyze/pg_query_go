// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterTSConfigurationStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterTSConfigurationStmt")
	if len(node.Cfgname.Items) > 0 {
		ctx.WriteString("cfgname")
		node.Cfgname.Fingerprint(ctx, "Cfgname")
	}

	if len(node.Dicts.Items) > 0 {
		ctx.WriteString("dicts")
		node.Dicts.Fingerprint(ctx, "Dicts")
	}

	if node.MissingOk {
		ctx.WriteString("missing_ok")
		ctx.WriteString(strconv.FormatBool(node.MissingOk))
	}

	if node.Override {
		ctx.WriteString("override")
		ctx.WriteString(strconv.FormatBool(node.Override))
	}

	if node.Replace {
		ctx.WriteString("replace")
		ctx.WriteString(strconv.FormatBool(node.Replace))
	}

	if len(node.Tokentype.Items) > 0 {
		ctx.WriteString("tokentype")
		node.Tokentype.Fingerprint(ctx, "Tokentype")
	}
}
