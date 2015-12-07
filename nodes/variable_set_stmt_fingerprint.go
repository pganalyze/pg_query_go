// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node VariableSetStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("SET")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.IsLocal))
	ctx.WriteString(strconv.Itoa(int(node.Kind)))

	if node.Name != nil {
		ctx.WriteString(*node.Name)
	}
}
