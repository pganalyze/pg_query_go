// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RangeVar) Fingerprint(ctx FingerprintContext, parentNode Node, parentFieldName string) {
	ctx.WriteString("RangeVar")

	if node.Alias != nil {
		subCtx := FingerprintSubContext{}
		node.Alias.Fingerprint(&subCtx, node, "Alias")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("alias")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.Catalogname != nil {
		ctx.WriteString("catalogname")
		ctx.WriteString(*node.Catalogname)
	}

	if int(node.InhOpt) != 0 {
		ctx.WriteString("inhOpt")
		ctx.WriteString(strconv.Itoa(int(node.InhOpt)))
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Relname != nil {
		ctx.WriteString("relname")
		ctx.WriteString(*node.Relname)
	}

	if node.Relpersistence != 0 {
		ctx.WriteString("relpersistence")
		ctx.WriteString(string(node.Relpersistence))

	}

	if node.Schemaname != nil {
		ctx.WriteString("schemaname")
		ctx.WriteString(*node.Schemaname)
	}
}
