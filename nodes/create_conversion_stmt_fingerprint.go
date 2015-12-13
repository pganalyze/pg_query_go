// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateConversionStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CreateConversionStmt")

	for _, subNode := range node.ConversionName {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.Def))

	if node.ForEncodingName != nil {
		ctx.WriteString(*node.ForEncodingName)
	}

	for _, subNode := range node.FuncName {
		subNode.Fingerprint(ctx)
	}

	if node.ToEncodingName != nil {
		ctx.WriteString(*node.ToEncodingName)
	}
}
