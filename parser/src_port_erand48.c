/*--------------------------------------------------------------------
 * Symbols referenced in this file:
 * - pg_lrand48
 * - _dorand48
 * - _rand48_mult
 * - _rand48_add
 * - _rand48_seed
 *--------------------------------------------------------------------
 */

/*-------------------------------------------------------------------------
 *
 * erand48.c
 *
 * This file supplies pg_erand48() and related functions, which except
 * for the names are just like the POSIX-standard erand48() family.
 * (We don't supply the full set though, only the ones we have found use
 * for in Postgres.  In particular, we do *not* implement lcong48(), so
 * that there is no need for the multiplier and addend to be variable.)
 *
 * We used to test for an operating system version rather than
 * unconditionally using our own, but (1) some versions of Cygwin have a
 * buggy erand48() that always returns zero and (2) as of 2011, glibc's
 * erand48() is strangely coded to be almost-but-not-quite thread-safe,
 * which doesn't matter for the backend but is important for pgbench.
 *
 * Portions Copyright (c) 1996-2018, PostgreSQL Global Development Group
 *
 * Portions Copyright (c) 1993 Martin Birgmeier
 * All rights reserved.
 *
 * You may redistribute unmodified or modified versions of this source
 * code provided that the above copyright notice and this and the
 * following conditions are retained.
 *
 * This software is provided ``as is'', and comes with no warranties
 * of any kind. I shall in no event be liable for anything that happens
 * to anyone/anything when using this software.
 *
 * IDENTIFICATION
 *	  src/port/erand48.c
 *
 *-------------------------------------------------------------------------
 */

#include "c.h"

#include <math.h>

#define RAND48_SEED_0	(0x330e)
#define RAND48_SEED_1	(0xabcd)
#define RAND48_SEED_2	(0x1234)
#define RAND48_MULT_0	(0xe66d)
#define RAND48_MULT_1	(0xdeec)
#define RAND48_MULT_2	(0x0005)
#define RAND48_ADD		(0x000b)

static __thread unsigned short _rand48_seed[3] = {
	RAND48_SEED_0,
	RAND48_SEED_1,
	RAND48_SEED_2
};

static __thread unsigned short _rand48_mult[3] = {
	RAND48_MULT_0,
	RAND48_MULT_1,
	RAND48_MULT_2
};

static __thread unsigned short _rand48_add = RAND48_ADD;



/*
 * Advance the 48-bit value stored in xseed[] to the next "random" number.
 */
static void
_dorand48(unsigned short xseed[3])
{
	unsigned long accu;
	unsigned short temp[2];

	accu = (unsigned long) _rand48_mult[0] * (unsigned long) xseed[0] +
		(unsigned long) _rand48_add;
	temp[0] = (unsigned short) accu;	/* lower 16 bits */
	accu >>= sizeof(unsigned short) * 8;
	accu += (unsigned long) _rand48_mult[0] * (unsigned long) xseed[1] +
		(unsigned long) _rand48_mult[1] * (unsigned long) xseed[0];
	temp[1] = (unsigned short) accu;	/* middle 16 bits */
	accu >>= sizeof(unsigned short) * 8;
	accu += _rand48_mult[0] * xseed[2] + _rand48_mult[1] * xseed[1] + _rand48_mult[2] * xseed[0];
	xseed[0] = temp[0];
	xseed[1] = temp[1];
	xseed[2] = (unsigned short) accu;
}


/*
 * Generate a random floating-point value using caller-supplied state.
 * Values are uniformly distributed over the interval [0.0, 1.0).
 */


/*
 * Generate a random non-negative integral value using internal state.
 * Values are uniformly distributed over the interval [0, 2^31).
 */
long
pg_lrand48(void)
{
	_dorand48(_rand48_seed);
	return ((long) _rand48_seed[2] << 15) + ((long) _rand48_seed[1] >> 1);
}

/*
 * Generate a random signed integral value using caller-supplied state.
 * Values are uniformly distributed over the interval [-2^31, 2^31).
 */


/*
 * Initialize the internal state using the given seed.
 *
 * Per POSIX, this uses only 32 bits from "seed" even if "long" is wider.
 * Hence, the set of possible seed values is smaller than it could be.
 * Better practice is to use caller-supplied state and initialize it with
 * random bits obtained from a high-quality source of random bits.
 *
 * Note: POSIX specifies a function seed48() that allows all 48 bits
 * of the internal state to be set, but we don't currently support that.
 */

