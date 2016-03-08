// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateSchemaStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateSchemaStmt")

	if node.Authrole != nil {
		subCtx := FingerprintSubContext{}
		node.Authrole.Fingerprint(&subCtx, "Authrole")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("authrole")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.IfNotExists {
		ctx.WriteString("if_not_exists")
		ctx.WriteString(strconv.FormatBool(node.IfNotExists))
	}

	if len(node.SchemaElts.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.SchemaElts.Fingerprint(&subCtx, "SchemaElts")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("schemaElts")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.Schemaname != nil {
		ctx.WriteString("schemaname")
		ctx.WriteString(*node.Schemaname)
	}
}
