// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateCastStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateCastStmt")

	if int(node.Context) != 0 {
		ctx.WriteString("context")
		ctx.WriteString(strconv.Itoa(int(node.Context)))
	}

	if node.Func != nil {
		subCtx := FingerprintSubContext{}
		node.Func.Fingerprint(&subCtx, "Func")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("func")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.Inout {
		ctx.WriteString("inout")
		ctx.WriteString(strconv.FormatBool(node.Inout))
	}

	if node.Sourcetype != nil {
		subCtx := FingerprintSubContext{}
		node.Sourcetype.Fingerprint(&subCtx, "Sourcetype")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("sourcetype")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.Targettype != nil {
		subCtx := FingerprintSubContext{}
		node.Targettype.Fingerprint(&subCtx, "Targettype")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("targettype")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}
