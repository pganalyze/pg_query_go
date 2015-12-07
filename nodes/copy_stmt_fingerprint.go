// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CopyStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("COPY")

	for _, subNode := range node.Attlist {
		subNode.Fingerprint(ctx)
	}

	if node.Filename != nil {
		ctx.WriteString(*node.Filename)
	}

	ctx.WriteString(strconv.FormatBool(node.IsFrom))
	ctx.WriteString(strconv.FormatBool(node.IsProgram))

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Query != nil {
		node.Query.Fingerprint(ctx)
	}

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}
}
