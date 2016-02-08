// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CommentStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CommentStmt")

	if node.Comment != nil {
		ctx.WriteString("comment")
		ctx.WriteString(*node.Comment)
	}

	if len(node.Objargs.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Objargs.Fingerprint(&subCtx, "Objargs")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("objargs")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.Objname.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Objname.Fingerprint(&subCtx, "Objname")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("objname")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if int(node.Objtype) != 0 {
		ctx.WriteString("objtype")
		ctx.WriteString(strconv.Itoa(int(node.Objtype)))
	}
}
