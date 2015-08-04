#include "pg_query.h"

const char* progname = "pg_query";

void pg_query_init(void)
{
	MemoryContextInit();
}
