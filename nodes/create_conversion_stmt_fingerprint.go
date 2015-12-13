// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateConversionStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateConversionStmt")

	for _, subNode := range node.ConversionName {
		subNode.Fingerprint(ctx, "ConversionName")
	}

	ctx.WriteString(strconv.FormatBool(node.Def))

	if node.ForEncodingName != nil {
		ctx.WriteString(*node.ForEncodingName)
	}

	for _, subNode := range node.FuncName {
		subNode.Fingerprint(ctx, "FuncName")
	}

	if node.ToEncodingName != nil {
		ctx.WriteString(*node.ToEncodingName)
	}
}
