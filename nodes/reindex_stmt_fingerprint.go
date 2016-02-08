// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ReindexStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ReindexStmt")

	if node.DoSystem {
		ctx.WriteString("do_system")
		ctx.WriteString(strconv.FormatBool(node.DoSystem))
	}

	if node.DoUser {
		ctx.WriteString("do_user")
		ctx.WriteString(strconv.FormatBool(node.DoUser))
	}

	if int(node.Kind) != 0 {
		ctx.WriteString("kind")
		ctx.WriteString(strconv.Itoa(int(node.Kind)))
	}

	if node.Name != nil {
		ctx.WriteString("name")
		ctx.WriteString(*node.Name)
	}

	if node.Relation != nil {
		subCtx := FingerprintSubContext{}
		node.Relation.Fingerprint(&subCtx, "Relation")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("relation")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}
