// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterTSConfigurationStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("AlterTSConfigurationStmt")

	for _, subNode := range node.Cfgname {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Dicts {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.MissingOk))
	ctx.WriteString(strconv.FormatBool(node.Override))
	ctx.WriteString(strconv.FormatBool(node.Replace))

	for _, subNode := range node.Tokentype {
		subNode.Fingerprint(ctx)
	}
}
