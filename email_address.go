package valueobject

import (
	"errors"
	"regexp"
)

// Errors block
var (
	ErrInvalidEmailAddress = errors.New("Not a valid email address")
)

type EmailAddress struct {
	value string
}

// NewEmailAddress creates a new numeral
func NewEmailAddress(v interface{}) (EmailAddress, error) {
	var n EmailAddress
	switch t := v.(type) {
	case string:
		match, _ := regexp.MatchString(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, t)
		if !match {
			return n, ErrInvalidEmailAddress
		}
		n.value = t
	}

	return n, nil
}

// String returns string representation
func (n EmailAddress) String() string {
	return n.value
}

// Equals checks that two values are the same
func (n EmailAddress) Equals(value Value) bool {
	otherEmailAddress, ok := value.(EmailAddress)
	return ok && n.value == otherEmailAddress.value
}
