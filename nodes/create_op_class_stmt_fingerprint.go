// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CreateOpClassStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CreateOpClassStmt")
	if node.Datatype != nil {
		node.Datatype.Fingerprint(ctx)
	}

	for _, subNode := range node.Items {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Opclassname {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Opfamilyname {
		subNode.Fingerprint(ctx)
	}
}
