// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RangeTblEntry) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("RangeTblEntry")

	if node.Alias != nil {
		ctx.WriteString("alias")
		node.Alias.Fingerprint(ctx, "Alias")
	}

	if node.CheckAsUser != 0 {
		ctx.WriteString("checkAsUser")
		ctx.WriteString(strconv.Itoa(int(node.CheckAsUser)))
	}

	if len(node.Ctecolcollations.Items) > 0 {
		ctx.WriteString("ctecolcollations")
		node.Ctecolcollations.Fingerprint(ctx, "Ctecolcollations")
	}

	if len(node.Ctecoltypes.Items) > 0 {
		ctx.WriteString("ctecoltypes")
		node.Ctecoltypes.Fingerprint(ctx, "Ctecoltypes")
	}

	if len(node.Ctecoltypmods.Items) > 0 {
		ctx.WriteString("ctecoltypmods")
		node.Ctecoltypmods.Fingerprint(ctx, "Ctecoltypmods")
	}

	if node.Ctelevelsup != 0 {
		ctx.WriteString("ctelevelsup")
		ctx.WriteString(strconv.Itoa(int(node.Ctelevelsup)))
	}

	if node.Ctename != nil {
		ctx.WriteString("ctename")
		ctx.WriteString(*node.Ctename)
	}

	if node.Eref != nil {
		ctx.WriteString("eref")
		node.Eref.Fingerprint(ctx, "Eref")
	}

	if node.Funcordinality {
		ctx.WriteString("funcordinality")
		ctx.WriteString(strconv.FormatBool(node.Funcordinality))
	}

	if len(node.Functions.Items) > 0 {
		ctx.WriteString("functions")
		node.Functions.Fingerprint(ctx, "Functions")
	}

	if node.InFromCl {
		ctx.WriteString("inFromCl")
		ctx.WriteString(strconv.FormatBool(node.InFromCl))
	}

	if node.Inh {
		ctx.WriteString("inh")
		ctx.WriteString(strconv.FormatBool(node.Inh))
	}

	if len(node.Joinaliasvars.Items) > 0 {
		ctx.WriteString("joinaliasvars")
		node.Joinaliasvars.Fingerprint(ctx, "Joinaliasvars")
	}

	if int(node.Jointype) != 0 {
		ctx.WriteString("jointype")
		ctx.WriteString(strconv.Itoa(int(node.Jointype)))
	}

	if node.Lateral {
		ctx.WriteString("lateral")
		ctx.WriteString(strconv.FormatBool(node.Lateral))
	}

	ctx.WriteString("modifiedCols")
	for _, val := range node.ModifiedCols {
		ctx.WriteString(strconv.Itoa(int(val)))
	}

	if node.Relid != 0 {
		ctx.WriteString("relid")
		ctx.WriteString(strconv.Itoa(int(node.Relid)))
	}

	if node.Relkind != 0 {
		ctx.WriteString("relkind")
		ctx.WriteString(string(node.Relkind))

	}

	if node.RequiredPerms != 0 {
		ctx.WriteString("requiredPerms")
		ctx.WriteString(strconv.Itoa(int(node.RequiredPerms)))
	}

	if int(node.Rtekind) != 0 {
		ctx.WriteString("rtekind")
		ctx.WriteString(strconv.Itoa(int(node.Rtekind)))
	}

	if len(node.SecurityQuals.Items) > 0 {
		ctx.WriteString("securityQuals")
		node.SecurityQuals.Fingerprint(ctx, "SecurityQuals")
	}

	if node.SecurityBarrier {
		ctx.WriteString("security_barrier")
		ctx.WriteString(strconv.FormatBool(node.SecurityBarrier))
	}

	ctx.WriteString("selectedCols")
	for _, val := range node.SelectedCols {
		ctx.WriteString(strconv.Itoa(int(val)))
	}

	if node.SelfReference {
		ctx.WriteString("self_reference")
		ctx.WriteString(strconv.FormatBool(node.SelfReference))
	}

	if node.Subquery != nil {
		ctx.WriteString("subquery")
		node.Subquery.Fingerprint(ctx, "Subquery")
	}

	if len(node.ValuesCollations.Items) > 0 {
		ctx.WriteString("values_collations")
		node.ValuesCollations.Fingerprint(ctx, "ValuesCollations")
	}

	if len(node.ValuesLists.Items) > 0 {
		ctx.WriteString("values_lists")
		node.ValuesLists.Fingerprint(ctx, "ValuesLists")
	}
}
