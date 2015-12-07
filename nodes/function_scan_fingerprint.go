// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node FunctionScan) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("FUNCTIONSCAN")
	ctx.WriteString(strconv.FormatBool(node.Funcordinality))

	for _, subNode := range node.Functions {
		subNode.Fingerprint(ctx)
	}
}
