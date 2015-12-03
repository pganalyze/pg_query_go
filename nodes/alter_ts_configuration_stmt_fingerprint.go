// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterTSConfigurationStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "AlterTSConfigurationStmt")

	for _, subNode := range node.Cfgname {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Dicts {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Tokentype {
		subNode.Fingerprint(ctx)
	}
}
