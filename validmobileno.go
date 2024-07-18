// Description: Package validmobileno provides a function to validate a mobile number.
package validmobileno

// imported regexp package
import "regexp"

// IsValidMobileNumber validates a mobile number.
func IsValidMobileNumber(number string) bool {
	// Regular expression for validating a mobile number
	re := regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
	return re.MatchString(number)
}
