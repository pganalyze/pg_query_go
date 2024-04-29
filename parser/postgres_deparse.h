#ifndef POSTGRES_DEPARSE_H
#define POSTGRES_DEPARSE_H

#include "lib/stringinfo.h"
#include "nodes/parsenodes.h"

// We cannot define static functions in header file so we have to take out the static keyword.
extern void deparseRawStmt(StringInfo str, RawStmt *raw_stmt);
extern void deparseAnyOperator(StringInfo str, List *op);
extern void deparseExpr(StringInfo str, Node *node);
extern void deparseIndexElem(StringInfo str, IndexElem* index_elem);
extern void deparseTypeName(StringInfo str, TypeName *type_name);

#endif
