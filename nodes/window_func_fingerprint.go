// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node WindowFunc) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("WindowFunc")

	if node.Aggfilter != nil {
		ctx.WriteString("aggfilter")
		node.Aggfilter.Fingerprint(ctx, "Aggfilter")
	}

	if len(node.Args.Items) > 0 {
		ctx.WriteString("args")
		node.Args.Fingerprint(ctx, "Args")
	}

	if node.Inputcollid != 0 {
		ctx.WriteString("inputcollid")
		ctx.WriteString(strconv.Itoa(int(node.Inputcollid)))
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Winagg {
		ctx.WriteString("winagg")
		ctx.WriteString(strconv.FormatBool(node.Winagg))
	}

	if node.Wincollid != 0 {
		ctx.WriteString("wincollid")
		ctx.WriteString(strconv.Itoa(int(node.Wincollid)))
	}

	if node.Winfnoid != 0 {
		ctx.WriteString("winfnoid")
		ctx.WriteString(strconv.Itoa(int(node.Winfnoid)))
	}

	if node.Winref != 0 {
		ctx.WriteString("winref")
		ctx.WriteString(strconv.Itoa(int(node.Winref)))
	}

	if node.Winstar {
		ctx.WriteString("winstar")
		ctx.WriteString(strconv.FormatBool(node.Winstar))
	}

	if node.Wintype != 0 {
		ctx.WriteString("wintype")
		ctx.WriteString(strconv.Itoa(int(node.Wintype)))
	}

	if node.Xpr != nil {
		ctx.WriteString("xpr")
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
