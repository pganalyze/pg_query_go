// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateConversionStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateConversionStmt")
	node.ConversionName.Fingerprint(ctx, "ConversionName")
	ctx.WriteString(strconv.FormatBool(node.Def))

	if node.ForEncodingName != nil {
		ctx.WriteString(*node.ForEncodingName)
	}

	node.FuncName.Fingerprint(ctx, "FuncName")

	if node.ToEncodingName != nil {
		ctx.WriteString(*node.ToEncodingName)
	}
}
