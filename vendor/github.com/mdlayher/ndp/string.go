// Code generated by "stringer -type=Preference -output=string.go"; DO NOT EDIT.

package ndp

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Medium-0]
	_ = x[High-1]
	_ = x[prfReserved-2]
	_ = x[Low-3]
}

const _Preference_name = "MediumHighprfReservedLow"

var _Preference_index = [...]uint8{0, 6, 10, 21, 24}

func (i Preference) String() string {
	if i < 0 || i >= Preference(len(_Preference_index)-1) {
		return "Preference(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Preference_name[_Preference_index[i]:_Preference_index[i+1]]
}
