package thou

import (
	"testing"
)

// TestJoin tests the `join` function
func TestJoin(t *testing.T) {
	cases := []struct {
		Ints     string
		Fracts   string
		Prec     int
		Sep      string
		Rad      string
		Expected string
	}{
		{"12,345,", "678", 0, ",", ".", "12,345"},
		{"12,345,", "678", 3, ",", ".", "12,345.678"},
	}

	for _, c := range cases {
		out := join(c.Ints, c.Fracts, c.Prec, c.Sep, c.Rad)
		if out != c.Expected {
			t.Errorf("join() failed to properly join: %q", out)
		}
	}
}

// TestSplit tests the `split` function
func TestSplit(t *testing.T) {
	cases := []struct {
		In     string
		Ints   int
		Fracts string
	}{
		{"0.5", 0, "5"},
		{"1234", 1234, ""},
		{"1234.56", 1234, "56"},
	}

	for _, c := range cases {
		ints, fracts, err := split(c.In)

		if err != nil {
			t.Errorf("split() threw an error: %q", err)
		}

		if ints != c.Ints {
			t.Errorf("split() returned invalid ints: %q => %q", c.In, ints)
		}
		if fracts != c.Fracts {
			t.Errorf("split() returned invalid fracts: %q => %q", c.In, fracts)
		}
	}
}

// TestSepI tests the `SepI` function
func TestSepI(t *testing.T) {
	cases := []struct {
		In       int
		Expected string
	}{
		// positive numbers
		{1, "1"},
		{10, "10"},
		{100, "100"},
		{1000, "1,000"},
		{10000, "10,000"},
		{100000, "100,000"},
		{1000000, "1,000,000"},

		// negative numbers
		{-1, "-1"},
		{-10, "-10"},
		{-100, "-100"},
		{-1000, "-1,000"},
		{-10000, "-10,000"},
		{-100000, "-100,000"},
		{-1000000, "-1,000,000"},
	}

	for _, c := range cases {
		out := SepI(c.In, ",")
		if out != c.Expected {
			t.Errorf(
				"SepI() failed to properly separate: %d => %q (expected: %q)",
				c.In,
				out,
				c.Expected,
			)
		}
	}
}

// TestSepF tests the `SepF` function
func TestSepF(t *testing.T) {
	cases := []struct {
		In       float64
		Prec     int
		Sep      string
		Rad      string
		Expected string
	}{
		// positive numbers
		{1, 0, ",", ".", "1"},
		{10, 0, ",", ".", "10"},
		{100, 0, ",", ".", "100"},
		{1000, 0, ",", ".", "1,000"},
		{10000, 0, ",", ".", "10,000"},
		{100000, 0, ",", ".", "100,000"},
		{1000000, 0, ",", ".", "1,000,000"},
		{1000000.0, 1, ",", ".", "1,000,000.0"},
		{1000000.00, 2, ",", ".", "1,000,000.00"},
		{1000000.000, 3, ",", ".", "1,000,000.000"},
		{1000000.0000, 4, ",", ".", "1,000,000.0000"},
		{1000000.0000, 4, "x", ".", "1x000x000.0000"},
		{1000000.0000, 4, ",", "x", "1,000,000x0000"},

		// negative numbers
		{-1, 0, ",", ".", "-1"},
		{-10, 0, ",", ".", "-10"},
		{-100, 0, ",", ".", "-100"},
		{-1000, 0, ",", ".", "-1,000"},
		{-10000, 0, ",", ".", "-10,000"},
		{-100000, 0, ",", ".", "-100,000"},
		{-1000000, 0, ",", ".", "-1,000,000"},
		{-1000000.0, 1, ",", ".", "-1,000,000.0"},
		{-1000000.00, 2, ",", ".", "-1,000,000.00"},
		{-1000000.000, 3, ",", ".", "-1,000,000.000"},
		{-1000000.0000, 4, ",", ".", "-1,000,000.0000"},
		{-1000000.0000, 4, "x", ".", "-1x000x000.0000"},
		{-1000000.0000, 4, ",", "x", "-1,000,000x0000"},
	}

	for _, c := range cases {
		out, err := SepF(c.In, c.Prec, c.Sep, c.Rad)
		if err != nil {
			t.Errorf("SepF() threw an error: %q", err)
		}

		if out != c.Expected {
			t.Errorf(
				"SepF() failed to properly separate: %g => %q (expected: %q)",
				c.In,
				out,
				c.Expected,
			)
		}
	}
}

// TestSepS tests the `SepS` function.
func TestSepS(t *testing.T) {
	cases := []struct {
		In       string
		Prec     int
		Sep      string
		Rad      string
		Expected string
	}{
		// positive numbers
		{"1", 0, ",", ".", "1"},
		{"10", 0, ",", ".", "10"},
		{"100", 0, ",", ".", "100"},
		{"1000", 0, ",", ".", "1,000"},
		{"10000", 0, ",", ".", "10,000"},
		{"100000", 0, ",", ".", "100,000"},
		{"1000000", 0, ",", ".", "1,000,000"},
		{"1000000.0", 1, ",", ".", "1,000,000.0"},
		{"1000000.00", 2, ",", ".", "1,000,000.00"},
		{"1000000.000", 3, ",", ".", "1,000,000.000"},
		{"1000000.0000", 4, ",", ".", "1,000,000.0000"},
		{"1000000.0000", 4, "x", ".", "1x000x000.0000"},
		{"1000000.0000", 4, ",", "x", "1,000,000x0000"},

		// negative numbers
		{"-1", 0, ",", ".", "-1"},
		{"-10", 0, ",", ".", "-10"},
		{"-100", 0, ",", ".", "-100"},
		{"-1000", 0, ",", ".", "-1,000"},
		{"-10000", 0, ",", ".", "-10,000"},
		{"-100000", 0, ",", ".", "-100,000"},
		{"-1000000", 0, ",", ".", "-1,000,000"},
		{"-1000000.0", 1, ",", ".", "-1,000,000.0"},
		{"-1000000.00", 2, ",", ".", "-1,000,000.00"},
		{"-1000000.000", 3, ",", ".", "-1,000,000.000"},
		{"-1000000.0000", 4, ",", ".", "-1,000,000.0000"},
		{"-1000000.0000", 4, "x", ".", "-1x000x000.0000"},
		{"-1000000.0000", 4, ",", "x", "-1,000,000x0000"},
	}

	for _, c := range cases {
		out, err := SepS(c.In, c.Prec, c.Sep, c.Rad)
		if err != nil {
			t.Errorf("SepS() threw an error: %q", err)
		}

		if out != c.Expected {
			t.Errorf(
				"SepS() failed to properly separate: %q => %q (expected: %q)",
				c.In,
				out,
				c.Expected,
			)
		}
	}

	// assert that errors are thrown when appropriate
	errCases := []struct {
		In   string
		Prec int
		Sep  string
		Rad  string
	}{
		{"not a number", 0, ",", "."},
		{"not a number.1234", 0, ",", "."},
		{"1234.not a number", 0, ",", "."},
		{"-not a number", 0, ",", "."},
		{"-not a number.1234", 0, ",", "."},
		{"-1234.not a number", 0, ",", "."},
		{"1,234", 0, ",", "."},
		{"1,234.00", 0, ",", "."},
	}

	for _, c := range errCases {
		_, err := SepS(c.In, c.Prec, c.Sep, c.Rad)
		if err == nil {
			t.Errorf("SepS() failed to throw an error for %q (invalid number)", c.In)
		}
	}
}

// BenchmarkSepI benchmarks SepI
func BenchmarkSepI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SepI(100000000, ",")
	}
}

// BenchmarkSepF benchmarks SepF
func BenchmarkSepF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SepF(100000000.0000, 4, ",", ".")
	}
}

// BenchmarkSepS benchmarks SepS
func BenchmarkSepS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SepS("100000000.0000", 4, ",", ".")
	}
}
