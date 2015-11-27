// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 * IndexOptInfo
 *		Per-index information for planning/optimization
 *
 *		indexkeys[], indexcollations[], opfamily[], and opcintype[]
 *		each have ncolumns entries.
 *
 *		sortopfamily[], reverse_sort[], and nulls_first[] likewise have
 *		ncolumns entries, if the index is ordered; but if it is unordered,
 *		those pointers are NULL.
 *
 *		Zeroes in the indexkeys[] array indicate index columns that are
 *		expressions; there is one element in indexprs for each such column.
 *
 *		For an ordered index, reverse_sort[] and nulls_first[] describe the
 *		sort ordering of a forward indexscan; we can also consider a backward
 *		indexscan, which will generate the reverse ordering.
 *
 *		The indexprs and indpred expressions have been run through
 *		prepqual.c and eval_const_expressions() for ease of matching to
 *		WHERE clauses. indpred is in implicit-AND form.
 *
 *		indextlist is a TargetEntry list representing the index columns.
 *		It provides an equivalent base-relation Var for each simple column,
 *		and links to the matching indexprs element for each expression column.
 */
type IndexOptInfo struct {
	Indexoid      Oid         `json:"indexoid"`      /* OID of the index relation */
	Reltablespace Oid         `json:"reltablespace"` /* tablespace of index (not table) */
	Rel           *RelOptInfo `json:"rel"`           /* back-link to index's table */

	/* index-size statistics (from pg_class and elsewhere) */
	Pages      BlockNumber `json:"pages"`       /* number of disk pages in index */
	Tuples     float64     `json:"tuples"`      /* number of index tuples in index */
	TreeHeight int         `json:"tree_height"` /* index tree height, or -1 if unknown */

	/* index descriptor information */
	Ncolumns        int   `json:"ncolumns"`        /* number of columns in index */
	Indexkeys       *int  `json:"indexkeys"`       /* column numbers of index's keys, or 0 */
	Indexcollations *Oid  `json:"indexcollations"` /* OIDs of collations of index columns */
	Opfamily        *Oid  `json:"opfamily"`        /* OIDs of operator families for columns */
	Opcintype       *Oid  `json:"opcintype"`       /* OIDs of opclass declared input data types */
	Sortopfamily    *Oid  `json:"sortopfamily"`    /* OIDs of btree opfamilies, if orderable */
	ReverseSort     *bool `json:"reverse_sort"`    /* is sort order descending? */
	NullsFirst      *bool `json:"nulls_first"`     /* do NULLs come first in the sort order? */
	Relam           Oid   `json:"relam"`           /* OID of the access method (in pg_am) */

	Amcostestimate RegProcedure `json:"amcostestimate"` /* OID of the access method's cost fcn */

	Indexprs []Node `json:"indexprs"` /* expressions for non-simple index columns */
	Indpred  []Node `json:"indpred"`  /* predicate if a partial index, else NIL */

	Indextlist []Node `json:"indextlist"` /* targetlist representing index columns */

	PredOk         bool `json:"predOK"`         /* true if predicate matches query */
	Unique         bool `json:"unique"`         /* true if a unique index */
	Immediate      bool `json:"immediate"`      /* is uniqueness enforced immediately? */
	Hypothetical   bool `json:"hypothetical"`   /* true if index doesn't really exist */
	Canreturn      bool `json:"canreturn"`      /* can index return IndexTuples? */
	Amcanorderbyop bool `json:"amcanorderbyop"` /* does AM support order by operator result? */
	Amoptionalkey  bool `json:"amoptionalkey"`  /* can query omit key for the first column? */
	Amsearcharray  bool `json:"amsearcharray"`  /* can AM handle ScalarArrayOpExpr quals? */
	Amsearchnulls  bool `json:"amsearchnulls"`  /* can AM search for NULL/NOT NULL entries? */
	Amhasgettuple  bool `json:"amhasgettuple"`  /* does AM have amgettuple interface? */
	Amhasgetbitmap bool `json:"amhasgetbitmap"` /* does AM have amgetbitmap interface? */
}

func (node IndexOptInfo) MarshalJSON() ([]byte, error) {
	type IndexOptInfoMarshalAlias IndexOptInfo
	return json.Marshal(map[string]interface{}{
		"INDEXOPTINFO": (*IndexOptInfoMarshalAlias)(&node),
	})
}

func (node *IndexOptInfo) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["indexoid"] != nil {
		err = json.Unmarshal(fields["indexoid"], &node.Indexoid)
		if err != nil {
			return
		}
	}

	if fields["reltablespace"] != nil {
		err = json.Unmarshal(fields["reltablespace"], &node.Reltablespace)
		if err != nil {
			return
		}
	}

	if fields["rel"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["rel"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(RelOptInfo)
			node.Rel = &val
		}
	}

	if fields["pages"] != nil {
		err = json.Unmarshal(fields["pages"], &node.Pages)
		if err != nil {
			return
		}
	}

	if fields["tuples"] != nil {
		err = json.Unmarshal(fields["tuples"], &node.Tuples)
		if err != nil {
			return
		}
	}

	if fields["tree_height"] != nil {
		err = json.Unmarshal(fields["tree_height"], &node.TreeHeight)
		if err != nil {
			return
		}
	}

	if fields["ncolumns"] != nil {
		err = json.Unmarshal(fields["ncolumns"], &node.Ncolumns)
		if err != nil {
			return
		}
	}

	if fields["indexkeys"] != nil {
		err = json.Unmarshal(fields["indexkeys"], &node.Indexkeys)
		if err != nil {
			return
		}
	}

	if fields["indexcollations"] != nil {
		err = json.Unmarshal(fields["indexcollations"], &node.Indexcollations)
		if err != nil {
			return
		}
	}

	if fields["opfamily"] != nil {
		err = json.Unmarshal(fields["opfamily"], &node.Opfamily)
		if err != nil {
			return
		}
	}

	if fields["opcintype"] != nil {
		err = json.Unmarshal(fields["opcintype"], &node.Opcintype)
		if err != nil {
			return
		}
	}

	if fields["sortopfamily"] != nil {
		err = json.Unmarshal(fields["sortopfamily"], &node.Sortopfamily)
		if err != nil {
			return
		}
	}

	if fields["reverse_sort"] != nil {
		err = json.Unmarshal(fields["reverse_sort"], &node.ReverseSort)
		if err != nil {
			return
		}
	}

	if fields["nulls_first"] != nil {
		err = json.Unmarshal(fields["nulls_first"], &node.NullsFirst)
		if err != nil {
			return
		}
	}

	if fields["relam"] != nil {
		err = json.Unmarshal(fields["relam"], &node.Relam)
		if err != nil {
			return
		}
	}

	if fields["amcostestimate"] != nil {
		err = json.Unmarshal(fields["amcostestimate"], &node.Amcostestimate)
		if err != nil {
			return
		}
	}

	if fields["indexprs"] != nil {
		node.Indexprs, err = UnmarshalNodeArrayJSON(fields["indexprs"])
		if err != nil {
			return
		}
	}

	if fields["indpred"] != nil {
		node.Indpred, err = UnmarshalNodeArrayJSON(fields["indpred"])
		if err != nil {
			return
		}
	}

	if fields["indextlist"] != nil {
		node.Indextlist, err = UnmarshalNodeArrayJSON(fields["indextlist"])
		if err != nil {
			return
		}
	}

	if fields["predOK"] != nil {
		err = json.Unmarshal(fields["predOK"], &node.PredOk)
		if err != nil {
			return
		}
	}

	if fields["unique"] != nil {
		err = json.Unmarshal(fields["unique"], &node.Unique)
		if err != nil {
			return
		}
	}

	if fields["immediate"] != nil {
		err = json.Unmarshal(fields["immediate"], &node.Immediate)
		if err != nil {
			return
		}
	}

	if fields["hypothetical"] != nil {
		err = json.Unmarshal(fields["hypothetical"], &node.Hypothetical)
		if err != nil {
			return
		}
	}

	if fields["canreturn"] != nil {
		err = json.Unmarshal(fields["canreturn"], &node.Canreturn)
		if err != nil {
			return
		}
	}

	if fields["amcanorderbyop"] != nil {
		err = json.Unmarshal(fields["amcanorderbyop"], &node.Amcanorderbyop)
		if err != nil {
			return
		}
	}

	if fields["amoptionalkey"] != nil {
		err = json.Unmarshal(fields["amoptionalkey"], &node.Amoptionalkey)
		if err != nil {
			return
		}
	}

	if fields["amsearcharray"] != nil {
		err = json.Unmarshal(fields["amsearcharray"], &node.Amsearcharray)
		if err != nil {
			return
		}
	}

	if fields["amsearchnulls"] != nil {
		err = json.Unmarshal(fields["amsearchnulls"], &node.Amsearchnulls)
		if err != nil {
			return
		}
	}

	if fields["amhasgettuple"] != nil {
		err = json.Unmarshal(fields["amhasgettuple"], &node.Amhasgettuple)
		if err != nil {
			return
		}
	}

	if fields["amhasgetbitmap"] != nil {
		err = json.Unmarshal(fields["amhasgetbitmap"], &node.Amhasgetbitmap)
		if err != nil {
			return
		}
	}

	return
}
