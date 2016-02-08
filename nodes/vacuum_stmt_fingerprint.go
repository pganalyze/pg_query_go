// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node VacuumStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("VacuumStmt")

	if node.FreezeMinAge != 0 {
		ctx.WriteString("freeze_min_age")
		ctx.WriteString(strconv.Itoa(int(node.FreezeMinAge)))
	}

	if node.FreezeTableAge != 0 {
		ctx.WriteString("freeze_table_age")
		ctx.WriteString(strconv.Itoa(int(node.FreezeTableAge)))
	}

	if node.MultixactFreezeMinAge != 0 {
		ctx.WriteString("multixact_freeze_min_age")
		ctx.WriteString(strconv.Itoa(int(node.MultixactFreezeMinAge)))
	}

	if node.MultixactFreezeTableAge != 0 {
		ctx.WriteString("multixact_freeze_table_age")
		ctx.WriteString(strconv.Itoa(int(node.MultixactFreezeTableAge)))
	}

	if node.Options != 0 {
		ctx.WriteString("options")
		ctx.WriteString(strconv.Itoa(int(node.Options)))
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

	if len(node.VaCols.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.VaCols.Fingerprint(&subCtx, "VaCols")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("va_cols")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}
