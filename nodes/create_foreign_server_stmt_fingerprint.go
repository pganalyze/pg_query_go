// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateForeignServerStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CREATEFOREIGNSERVERSTMT")

	if node.Fdwname != nil {
		ctx.WriteString(*node.Fdwname)
	}

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Servername != nil {
		ctx.WriteString(*node.Servername)
	}

	if node.Servertype != nil {
		ctx.WriteString(*node.Servertype)
	}

	if node.Version != nil {
		ctx.WriteString(*node.Version)
	}
}
