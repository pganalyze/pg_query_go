// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterExtensionContentsStmt) Fingerprint(ctx FingerprintContext, parentNode Node, parentFieldName string) {
	ctx.WriteString("AlterExtensionContentsStmt")

	if node.Action != 0 {
		ctx.WriteString("action")
		ctx.WriteString(strconv.Itoa(int(node.Action)))
	}

	if node.Extname != nil {
		ctx.WriteString("extname")
		ctx.WriteString(*node.Extname)
	}

	if len(node.Objargs.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Objargs.Fingerprint(&subCtx, node, "Objargs")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("objargs")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.Objname.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Objname.Fingerprint(&subCtx, node, "Objname")

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
