// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node RangeTblEntry) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "RangeTblEntry")
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

	if node.Eref != nil {
		node.Eref.Fingerprint(ctx)
	}

	for _, subNode := range node.Functions {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Joinaliasvars {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.SecurityQuals {
		subNode.Fingerprint(ctx)
	}

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
