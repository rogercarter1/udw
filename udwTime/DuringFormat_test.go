package udwTime

import (
	"github.com/tachyon-protocol/udw/udwTest"
	"testing"
	"time"
)

func TestDurationFormat(ot *testing.T) {
	udwTest.Equal(DurationFormat(time.Second), "1.00s")
	udwTest.Equal(DurationFormat(1000*time.Second), "16.67min")
	udwTest.Equal(DurationFormat(12345*time.Millisecond), "12.35s")
	udwTest.Equal(DurationFormat(1234*time.Microsecond), "1.23ms")
	udwTest.Equal(DurationFormat(1234*time.Nanosecond), "1.23µs")
	udwTest.Equal(DurationFormat(Day), "1.00day")
	udwTest.Equal(DurationFormat(time.Hour), "1.00h")
	udwTest.Equal(DurationFormat(Day*30), "30.00day")
	udwTest.Equal(DurationFormat(Day*400), "1.10year")
	udwTest.Equal(DurationFormatPadding(1234*time.Nanosecond), "    1.23µs")
	udwTest.Equal(DurationFormatPadding(1234*time.Millisecond), "     1.23s")
}

func TestDurationFormatTimeMysql(t *testing.T) {
	udwTest.Equal(DurationFormatTimeMysql(time.Second), "00:00:01")
	udwTest.Equal(DurationFormatTimeMysql(time.Minute), "00:01:00")
	udwTest.Equal(DurationFormatTimeMysql(time.Hour), "01:00:00")
	udwTest.Equal(DurationFormatTimeMysql(99*time.Hour), "99:00:00")
	udwTest.Equal(DurationFormatTimeMysql(999*time.Hour), "999:00:00")

	udwTest.Equal(DurationFormatTimeMysql(-time.Minute), "-00:01:00")
	udwTest.Equal(DurationFormatTimeMysql(-time.Second), "-00:00:01")
	udwTest.Equal(DurationFormatTimeMysql(-time.Millisecond), "00:00:00")
	udwTest.Equal(DurationFormatTimeMysql(time.Millisecond), "00:00:00")

}

func TestDurationFormatBefore(t *testing.T) {
	udwTest.Equal(DurationFormatBefore(time.Second), "1 minute ago")
	udwTest.Equal(DurationFormatBefore(119*time.Second), "1 minute ago")
	udwTest.Equal(DurationFormatBefore(120*time.Second), "2 minutes ago")
	udwTest.Equal(DurationFormatBefore(3600*time.Second), "1 hour ago")
	udwTest.Equal(DurationFormatBefore((2*3600-1)*time.Second), "1 hour ago")
	udwTest.Equal(DurationFormatBefore(2*3600*time.Second), "2 hours ago")
	udwTest.Equal(DurationFormatBefore(24*3600*time.Second), "1 day ago")
	udwTest.Equal(DurationFormatBefore(25*3600*time.Second), "1 day ago")
	udwTest.Equal(DurationFormatBefore(48*3600*time.Second), "2 days ago")
}

func TestDurationFormatByHourMin(t *testing.T) {
	udwTest.Equal(DurationFormatByHourMin(time.Hour), "01:00")

	udwTest.Equal(DurationFormatByHourMin(time.Hour*24), "24:00")
	udwTest.Equal(DurationFormatByHourMin(time.Hour*25), "25:00")
	udwTest.Equal(DurationFormatByHourMin(time.Hour+time.Minute*59), "01:59")

}
