// Auto-generated from postgres/src/include/postgres.h - DO NOT EDIT

package pg_query

import "encoding/json"

type varatt_external struct {
	VaRawsize    int32 `json:"va_rawsize"`    /* Original data size (includes header) */
	VaExtsize    int32 `json:"va_extsize"`    /* External saved size (doesn't) */
	VaValueid    Oid   `json:"va_valueid"`    /* Unique ID of value within TOAST table */
	VaToastrelid Oid   `json:"va_toastrelid"` /* RelID of TOAST table containing it */
}

func (node varatt_external) MarshalJSON() ([]byte, error) {
	type varatt_externalMarshalAlias varatt_external
	return json.Marshal(map[string]interface{}{
		"VARATT_EXTERNAL": (*varatt_externalMarshalAlias)(&node),
	})
}

func (node *varatt_external) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["va_rawsize"] != nil {
		err = json.Unmarshal(fields["va_rawsize"], &node.VaRawsize)
		if err != nil {
			return
		}
	}

	if fields["va_extsize"] != nil {
		err = json.Unmarshal(fields["va_extsize"], &node.VaExtsize)
		if err != nil {
			return
		}
	}

	if fields["va_valueid"] != nil {
		err = json.Unmarshal(fields["va_valueid"], &node.VaValueid)
		if err != nil {
			return
		}
	}

	if fields["va_toastrelid"] != nil {
		err = json.Unmarshal(fields["va_toastrelid"], &node.VaToastrelid)
		if err != nil {
			return
		}
	}

	return
}
