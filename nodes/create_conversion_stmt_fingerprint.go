// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateConversionStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateConversionStmt")
	if len(node.ConversionName.Items) > 0 {
		ctx.WriteString("conversion_name")
		node.ConversionName.Fingerprint(ctx, "ConversionName")
	}

	if node.Def {
		ctx.WriteString("def")
		ctx.WriteString(strconv.FormatBool(node.Def))
	}

	if node.ForEncodingName != nil {
		ctx.WriteString("for_encoding_name")
		ctx.WriteString(*node.ForEncodingName)
	}

	if len(node.FuncName.Items) > 0 {
		ctx.WriteString("func_name")
		node.FuncName.Fingerprint(ctx, "FuncName")
	}

	if node.ToEncodingName != nil {
		ctx.WriteString("to_encoding_name")
		ctx.WriteString(*node.ToEncodingName)
	}
}
