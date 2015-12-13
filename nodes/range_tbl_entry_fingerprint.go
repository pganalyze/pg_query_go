// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RangeTblEntry) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("RangeTblEntry")

	if node.Alias != nil {
		node.Alias.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.CheckAsUser)))

	for _, subNode := range node.Ctecolcollations {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Ctecoltypes {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Ctecoltypmods {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Ctelevelsup)))

	if node.Ctename != nil {
		ctx.WriteString(*node.Ctename)
	}

	if node.Eref != nil {
		node.Eref.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.Funcordinality))

	for _, subNode := range node.Functions {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.InFromCl))
	ctx.WriteString(strconv.FormatBool(node.Inh))

	for _, subNode := range node.Joinaliasvars {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Jointype)))
	ctx.WriteString(strconv.FormatBool(node.Lateral))

	for _, val := range node.ModifiedCols {
		ctx.WriteString(strconv.Itoa(int(val)))
	}

	ctx.WriteString(strconv.Itoa(int(node.Relid)))
	ctx.WriteString(string(node.Relkind))
	ctx.WriteString(strconv.Itoa(int(node.RequiredPerms)))
	ctx.WriteString(strconv.Itoa(int(node.Rtekind)))

	for _, subNode := range node.SecurityQuals {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.SecurityBarrier))

	for _, val := range node.SelectedCols {
		ctx.WriteString(strconv.Itoa(int(val)))
	}

	ctx.WriteString(strconv.FormatBool(node.SelfReference))

	if node.Subquery != nil {
		node.Subquery.Fingerprint(ctx)
	}

	for _, subNode := range node.ValuesCollations {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.ValuesLists {
		subNode.Fingerprint(ctx)
	}
}
