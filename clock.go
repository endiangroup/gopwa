package gopwa

import "time"

var Now = func() time.Time { return time.Now() }

func NowForce(t time.Time) {
	Now = func() time.Time { return t }
}

func NowReset() {
	Now = func() time.Time { return time.Now() }
}
