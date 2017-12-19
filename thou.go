// Package thou inserts thousands-separators into int, float64, or string
// representations of numbers. The thousands-separator, radix, and decimal
// precision are configurable.
package thou

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// join concatenates and returns formatted output
func join(ints string, fracts string, prec int, sep string, rad string) string {
	// trim a trailing separator from the `ints` string
	separated := strings.Trim(ints, sep)

	// if precision is 0, don't append the radix and fractionals
	if prec == 0 {
		return separated
	}

	// otherwise, concatenate the separated ints, radix, and fractionals
	return separated + rad + fracts
}

// split splits a string representation of a number into integers and
// fractionals.
func split(str string) (int, string, error) {
	// split the string on the radix
	// QUESTION: does Go *always* use "." as the radix, regardless of
	// internationalization?
	parts := strings.Split(str, ".")

	// if there is no radix, set fractionals to empty string
	if len(parts) == 1 {
		parts = append(parts, "")
	}

	// parse the integers into an integer type (of the appropriate width for the
	// system)
	ints, err := strconv.ParseInt(parts[0], 10, strconv.IntSize)
	if err != nil {
		return 0, "", err
	}

	// return the parts
	return int(ints), parts[1], nil
}

// SepF inserts thousands-separators into float64 values.
func SepF(value float64, prec int, sep string, rad string) (string, error) {
	// format value as a string
	str := strconv.FormatFloat(value, 'f', prec, 64)

	// split value into integers and fractionals
	ints, fracts, err := split(str)
	if err != nil {
		return "", err
	}

	// insert thousands-separators into the integers
	separated := SepI(ints, sep)

	// concatenate the result
	return join(separated, fracts, prec, sep, rad), nil
}

// SepI inserts thousands-separators into int values.
func SepI(value int, sep string) string {
	// take the absolute value of `value`
	neg := false
	if value < 0 {
		neg = true
		value *= -1
	}

	// convert `value` into a string
	str := strconv.Itoa(value)
	separated := ""

	// iterate backwards across the integers
	for i, j := len(str)-1, 0; i >= 0; i-- {
		val := string(str[i])

		// append a thousands separator after every 3rd digit
		if j%3 == 0 {
			val += sep
		}
		separated = val + separated
		j++
	}

	// concatenate the result
	joined := join(separated, "", 0, sep, "")

	// prepend a minus sign if value is negative
	if neg == true {
		joined = "-" + joined
	}

	// done
	return joined
}

// SepS inserts thousands-separators into string representations of numbers.
// Note that SepS is ~15x slower than SepI and SepF.
func SepS(value string, prec int, sep string, rad string) (string, error) {
	// assert that value looks like a number
	valid, err := regexp.MatchString("^[\\-1-9]\\d*(\\.\\d+)?$", value)
	if err != nil {
		return "", err
	}
	if valid != true {
		return "", errors.New(value + " must be a number.")
	}

	// split value into integers and fractionals
	ints, fracts, err := split(value)
	if err != nil {
		return "", err
	}

	// insert thousands-separators into the integers
	separated := SepI(ints, sep)

	// concatenate the result
	return join(separated, fracts, prec, sep, rad), nil
}
