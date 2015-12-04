// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node RangeTblEntry) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "RANGETBLENTRY")

	if node.Alias != nil {
		node.Alias.Fingerprint(ctx)
	}

	for _, subNode := range node.Ctecolcollations {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Ctecoltypes {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Ctecoltypmods {
		subNode.Fingerprint(ctx)
	}

	if node.Ctename != nil {
		io.WriteString(ctx.hash, *node.Ctename)
	}

	if node.Eref != nil {
		node.Eref.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.Funcordinality))

	for _, subNode := range node.Functions {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.InFromCl))
	io.WriteString(ctx.hash, strconv.FormatBool(node.Inh))

	for _, subNode := range node.Joinaliasvars {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Jointype)))
	io.WriteString(ctx.hash, strconv.FormatBool(node.Lateral))
	io.WriteString(ctx.hash, strconv.Itoa(int(node.Rtekind)))

	for _, subNode := range node.SecurityQuals {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.SecurityBarrier))
	io.WriteString(ctx.hash, strconv.FormatBool(node.SelfReference))

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
