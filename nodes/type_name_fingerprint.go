// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node TypeName) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("TypeName")
	if len(node.ArrayBounds.Items) > 0 {
		ctx.WriteString("arrayBounds")
		node.ArrayBounds.Fingerprint(ctx, "ArrayBounds")
	}

	// Intentionally ignoring node.Location for fingerprinting

	if len(node.Names.Items) > 0 {
		ctx.WriteString("names")
		node.Names.Fingerprint(ctx, "Names")
	}

	if node.PctType {
		ctx.WriteString("pct_type")
		ctx.WriteString(strconv.FormatBool(node.PctType))
	}

	if node.Setof {
		ctx.WriteString("setof")
		ctx.WriteString(strconv.FormatBool(node.Setof))
	}

	if node.TypeOid != 0 {
		ctx.WriteString("typeOid")
		ctx.WriteString(strconv.Itoa(int(node.TypeOid)))
	}

	if node.Typemod != 0 {
		ctx.WriteString("typemod")
		ctx.WriteString(strconv.Itoa(int(node.Typemod)))
	}

	if len(node.Typmods.Items) > 0 {
		ctx.WriteString("typmods")
		node.Typmods.Fingerprint(ctx, "Typmods")
	}
}
