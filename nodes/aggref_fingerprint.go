// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node Aggref) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("Aggref")
	ctx.WriteString(strconv.Itoa(int(node.Aggcollid)))
	node.Aggdirectargs.Fingerprint(ctx, "Aggdirectargs")
	node.Aggdistinct.Fingerprint(ctx, "Aggdistinct")

	if node.Aggfilter != nil {
		node.Aggfilter.Fingerprint(ctx, "Aggfilter")
	}

	ctx.WriteString(strconv.Itoa(int(node.Aggfnoid)))
	ctx.WriteString(string(node.Aggkind))
	ctx.WriteString(strconv.Itoa(int(node.Agglevelsup)))
	node.Aggorder.Fingerprint(ctx, "Aggorder")
	ctx.WriteString(strconv.FormatBool(node.Aggstar))
	ctx.WriteString(strconv.Itoa(int(node.Aggtype)))
	ctx.WriteString(strconv.FormatBool(node.Aggvariadic))
	node.Args.Fingerprint(ctx, "Args")
	ctx.WriteString(strconv.Itoa(int(node.Inputcollid)))
	// Intentionally ignoring node.Location for fingerprinting

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
