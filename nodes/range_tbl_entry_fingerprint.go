// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RangeTblEntry) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("RangeTblEntry")

	if node.Alias != nil {
		node.Alias.Fingerprint(ctx, "Alias")
	}

	ctx.WriteString(strconv.Itoa(int(node.CheckAsUser)))
	node.Ctecolcollations.Fingerprint(ctx, "Ctecolcollations")
	node.Ctecoltypes.Fingerprint(ctx, "Ctecoltypes")
	node.Ctecoltypmods.Fingerprint(ctx, "Ctecoltypmods")
	ctx.WriteString(strconv.Itoa(int(node.Ctelevelsup)))

	if node.Ctename != nil {
		ctx.WriteString(*node.Ctename)
	}

	if node.Eref != nil {
		node.Eref.Fingerprint(ctx, "Eref")
	}

	ctx.WriteString(strconv.FormatBool(node.Funcordinality))
	node.Functions.Fingerprint(ctx, "Functions")
	ctx.WriteString(strconv.FormatBool(node.InFromCl))
	ctx.WriteString(strconv.FormatBool(node.Inh))
	node.Joinaliasvars.Fingerprint(ctx, "Joinaliasvars")
	ctx.WriteString(strconv.Itoa(int(node.Jointype)))
	ctx.WriteString(strconv.FormatBool(node.Lateral))

	for _, val := range node.ModifiedCols {
		ctx.WriteString(strconv.Itoa(int(val)))
	}

	ctx.WriteString(strconv.Itoa(int(node.Relid)))
	ctx.WriteString(string(node.Relkind))
	ctx.WriteString(strconv.Itoa(int(node.RequiredPerms)))
	ctx.WriteString(strconv.Itoa(int(node.Rtekind)))
	node.SecurityQuals.Fingerprint(ctx, "SecurityQuals")
	ctx.WriteString(strconv.FormatBool(node.SecurityBarrier))

	for _, val := range node.SelectedCols {
		ctx.WriteString(strconv.Itoa(int(val)))
	}

	ctx.WriteString(strconv.FormatBool(node.SelfReference))

	if node.Subquery != nil {
		node.Subquery.Fingerprint(ctx, "Subquery")
	}

	node.ValuesCollations.Fingerprint(ctx, "ValuesCollations")
	node.ValuesLists.Fingerprint(ctx, "ValuesLists")
}
