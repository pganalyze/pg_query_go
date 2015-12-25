// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node WindowFunc) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("WindowFunc")

	if node.Aggfilter != nil {
		node.Aggfilter.Fingerprint(ctx, "Aggfilter")
	}

	node.Args.Fingerprint(ctx, "Args")
	ctx.WriteString(strconv.Itoa(int(node.Inputcollid)))
	// Intentionally ignoring node.Location for fingerprinting

	ctx.WriteString(strconv.FormatBool(node.Winagg))
	ctx.WriteString(strconv.Itoa(int(node.Wincollid)))
	ctx.WriteString(strconv.Itoa(int(node.Winfnoid)))
	ctx.WriteString(strconv.Itoa(int(node.Winref)))
	ctx.WriteString(strconv.FormatBool(node.Winstar))
	ctx.WriteString(strconv.Itoa(int(node.Wintype)))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
