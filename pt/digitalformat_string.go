// Code generated by "stringer -type DigitalFormat"; DO NOT EDIT.

package pt

import "strconv"

const _DigitalFormat_name = "BluerayHDTVWebDLUHDTVUnknownDigitalFormat"

var _DigitalFormat_index = [...]uint8{0, 7, 11, 16, 21, 41}

func (i DigitalFormat) String() string {
	if i >= DigitalFormat(len(_DigitalFormat_index)-1) {
		return "DigitalFormat(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DigitalFormat_name[_DigitalFormat_index[i]:_DigitalFormat_index[i+1]]
}
