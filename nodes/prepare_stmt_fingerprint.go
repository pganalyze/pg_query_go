// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node PrepareStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "PREPARESTMT")

	for _, subNode := range node.Argtypes {
		subNode.Fingerprint(ctx)
	}

	if node.Name != nil {
		io.WriteString(ctx.hash, *node.Name)
	}

	if node.Query != nil {
		node.Query.Fingerprint(ctx)
	}
}
