// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterExtensionContentsStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("AlterExtensionContentsStmt")
	ctx.WriteString(strconv.Itoa(int(node.Action)))

	if node.Extname != nil {
		ctx.WriteString(*node.Extname)
	}

	for _, subNode := range node.Objargs {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Objname {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Objtype)))
}
