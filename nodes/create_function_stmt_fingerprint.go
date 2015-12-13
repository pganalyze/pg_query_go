// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateFunctionStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CreateFunctionStmt")

	for _, subNode := range node.Funcname {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Parameters {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.Replace))

	if node.ReturnType != nil {
		node.ReturnType.Fingerprint(ctx)
	}

	for _, subNode := range node.WithClause {
		subNode.Fingerprint(ctx)
	}
}
