#ifndef PG_QUERY_H
#define PG_QUERY_H

#include "postgres.h"
#include "utils/memutils.h"

#define STDERR_BUFFER_LEN 4096
//#define DEBUG

//VALUE new_parse_error(ErrorData* error);

void pg_query_init(void);
char* pg_query_normalize(char* input);
char* pg_query_raw_parse(char* input);

#endif
