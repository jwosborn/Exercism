//Package leap provides func IsLeapYear which will tell whether a year (int) is a leap year
package leap

// IsLeapYear takes a year (int) and returns a boolean true if the year is a leap year and false if not
func IsLeapYear(a int) bool {
	if (a%100) == 0 && (a%400) == 0 {
		return true
	}
	if (a%4) == 0 && (a%100) != 0 {
		return true
	}
	if (a % 100) == 0 {
		return false
	}
	return false
}
