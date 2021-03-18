package oceanengine

import "bytes"

func GeneralScheduleTime(start, end int) string {
	buf := bytes.Buffer{}
	buf.Grow(336)
	for i := 0; i < start; i++ {
		buf.WriteByte('0')
		buf.WriteByte('0')
	}
	for i := 0; i < end-start; i++ {
		buf.WriteByte('1')
		buf.WriteByte('1')
	}
	for i := 0; i < 24-end; i++ {
		buf.WriteByte('0')
		buf.WriteByte('0')
	}
	dayRange := buf.String()
	buf.Reset()
	for i := 0; i < 7; i++ {
		buf.WriteString(dayRange)
	}
	return buf.String()
}
