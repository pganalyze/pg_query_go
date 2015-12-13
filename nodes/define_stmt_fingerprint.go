// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node DefineStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("DefineStmt")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Definition {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Defnames {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Kind)))
	ctx.WriteString(strconv.FormatBool(node.Oldstyle))
}
