// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node XmlExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "XmlExpr")

	for _, subNode := range node.ArgNames {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.NamedArgs {
		subNode.Fingerprint(ctx)
	}
}
