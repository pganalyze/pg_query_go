// Auto-generated - DO NOT EDIT

package pg_query

func (node ClosePortalStmt) Fingerprint(ctx FingerprintContext, parentNode Node, parentFieldName string) {
	ctx.WriteString("ClosePortalStmt")

	if node.Portalname != nil {
		ctx.WriteString("portalname")
		ctx.WriteString(*node.Portalname)
	}
}
