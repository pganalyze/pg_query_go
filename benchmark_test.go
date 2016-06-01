package pg_query_test

import (
	"testing"

	"github.com/lfittl/pg_query_go"
)

func benchmarkParse(input string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := pg_query.Parse(input)

		if err != nil {
			b.Errorf("Benchmark produced error %s\n\n", err)
		}
	}
}

func benchmarkRawParse(input string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := pg_query.ParseToJSON(input)

		if err != nil {
			b.Errorf("Benchmark produced error %s\n\n", err)
		}
	}
}

func benchmarkFingerprint(input string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree, err := pg_query.Parse(input)
		if err != nil {
			b.Errorf("Benchmark produced error %s\n\n", err)
		}

		tree.Fingerprint()
	}
}

func benchmarkFastFingerprint(input string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := pg_query.FastFingerprint(input)
		if err != nil {
			b.Errorf("Benchmark produced error %s\n\n", err)
		}
	}
}

func benchmarkNormalize(input string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := pg_query.Normalize(input)
		if err != nil {
			b.Errorf("Benchmark produced error %s\n\n", err)
		}
	}
}

func BenchmarkParseSelect1(b *testing.B) {
	benchmarkParse("SELECT 1", b)
}
func BenchmarkParseSelect2(b *testing.B) {
	benchmarkParse("SELECT 1 FROM x WHERE y IN ('a', 'b', 'c')", b)
}
func BenchmarkParseCreateTable(b *testing.B) {
	benchmarkParse("CREATE TABLE types (a float(2), b float(49), c NUMERIC(2, 3), d character(4), e char(5), f varchar(6), g character varying(7))", b)
}

func BenchmarkRawParseSelect1(b *testing.B) {
	benchmarkRawParse("SELECT 1", b)
}
func BenchmarkRawParseSelect2(b *testing.B) {
	benchmarkRawParse("SELECT 1 FROM x WHERE y IN ('a', 'b', 'c')", b)
}
func BenchmarkRawParseCreateTable(b *testing.B) {
	benchmarkRawParse("CREATE TABLE types (a float(2), b float(49), c NUMERIC(2, 3), d character(4), e char(5), f varchar(6), g character varying(7))", b)
}

func BenchmarkFingerprintSelect1(b *testing.B) {
	benchmarkFingerprint("SELECT 1", b)
}
func BenchmarkFingerprintSelect2(b *testing.B) {
	benchmarkFingerprint("SELECT 1 FROM x WHERE y IN ('a', 'b', 'c')", b)
}
func BenchmarkFingerprintCreateTable(b *testing.B) {
	benchmarkFingerprint("CREATE TABLE types (a float(2), b float(49), c NUMERIC(2, 3), d character(4), e char(5), f varchar(6), g character varying(7))", b)
}

func BenchmarkFastFingerprintSelect1(b *testing.B) {
	benchmarkFastFingerprint("SELECT 1", b)
}
func BenchmarkFastFingerprintSelect2(b *testing.B) {
	benchmarkFastFingerprint("SELECT 1 FROM x WHERE y IN ('a', 'b', 'c')", b)
}
func BenchmarkFastFingerprintCreateTable(b *testing.B) {
	benchmarkFastFingerprint("CREATE TABLE types (a float(2), b float(49), c NUMERIC(2, 3), d character(4), e char(5), f varchar(6), g character varying(7))", b)
}

func BenchmarkNormalizeSelect1(b *testing.B) {
	benchmarkNormalize("SELECT 1", b)
}
func BenchmarkNormalizeSelect2(b *testing.B) {
	benchmarkNormalize("SELECT 1 FROM x WHERE y IN ('a', 'b', 'c')", b)
}
func BenchmarkNormalizeCreateTable(b *testing.B) {
	benchmarkNormalize("CREATE TABLE types (a float(2), b float(49), c NUMERIC(2, 3), d character(4), e char(5), f varchar(6), g character varying(7))", b)
}
