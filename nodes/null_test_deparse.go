// Auto-generated - DO NOT EDIT

package pg_query

import (
	"errors"
	"fmt"
	"strings"
)

func (node NullTest) Deparse(ctx DeparseContext) (string, error) {
	out := make([]string, 0)
	if node.Arg == nil {
		return "", errors.New("argument cannot be null for null test (ironically)")
	}

	if str, err := node.Arg.Deparse(DeparseContextNone); err != nil {
		return "", err
	} else {
		out = append(out, str)
	}

	switch node.Nulltesttype {
	case IS_NULL:
		out = append(out, "IS NULL")
	case IS_NOT_NULL:
		out = append(out, "IS NOT NULL")
	default:
		return "", errors.New(fmt.Sprintf("could not parse null test type (%d)", node.Nulltesttype))
	}

	return strings.Join(out, " "), nil
}
