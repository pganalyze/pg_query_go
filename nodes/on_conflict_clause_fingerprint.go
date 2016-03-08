// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node OnConflictClause) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("OnConflictClause")

	if int(node.Action) != 0 {
		ctx.WriteString("action")
		ctx.WriteString(strconv.Itoa(int(node.Action)))
	}

	if node.Infer != nil {
		subCtx := FingerprintSubContext{}
		node.Infer.Fingerprint(&subCtx, "Infer")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("infer")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
	// Intentionally ignoring node.Location for fingerprinting

	if len(node.TargetList.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.TargetList.Fingerprint(&subCtx, "TargetList")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("targetList")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.WhereClause != nil {
		subCtx := FingerprintSubContext{}
		node.WhereClause.Fingerprint(&subCtx, "WhereClause")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("whereClause")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}
