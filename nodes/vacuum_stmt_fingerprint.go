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
		ctx.WriteString("relation")
		node.Relation.Fingerprint(ctx, "Relation")
	}

	if len(node.VaCols.Items) > 0 {
		ctx.WriteString("va_cols")
		node.VaCols.Fingerprint(ctx, "VaCols")
	}
}
