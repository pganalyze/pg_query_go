// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CreateForeignServerStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CREATEFOREIGNSERVERSTMT")

	if node.Fdwname != nil {
		io.WriteString(ctx.hash, *node.Fdwname)
	}

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Servername != nil {
		io.WriteString(ctx.hash, *node.Servername)
	}

	if node.Servertype != nil {
		io.WriteString(ctx.hash, *node.Servertype)
	}

	if node.Version != nil {
		io.WriteString(ctx.hash, *node.Version)
	}
}
