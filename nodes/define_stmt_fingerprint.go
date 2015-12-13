// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node DefineStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("DefineStmt")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx, "Args")
	}

	for _, subNode := range node.Definition {
		subNode.Fingerprint(ctx, "Definition")
	}

	for _, subNode := range node.Defnames {
		subNode.Fingerprint(ctx, "Defnames")
	}

	ctx.WriteString(strconv.Itoa(int(node.Kind)))
	ctx.WriteString(strconv.FormatBool(node.Oldstyle))
}
