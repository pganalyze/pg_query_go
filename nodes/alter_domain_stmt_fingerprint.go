// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterDomainStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "AlterDomainStmt")
	if node.Def != nil {
		node.Def.Fingerprint(ctx)
	}

	for _, subNode := range node.TypeName {
		subNode.Fingerprint(ctx)
	}
}
