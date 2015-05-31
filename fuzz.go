// +build gofuzz

package wd

func Fuzz(data []byte) int {
	median := len(data) / 2
	a := string(data[:median])
	b := string(data[median:])
	ColouredDiff(a, b, true)
	return 1
}
