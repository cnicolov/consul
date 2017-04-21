package retry

import "time"

type Fataler interface {
	Fatal(args ...interface{})
}

func Fatal(t Fataler, f func() error) {
	if err := Func(f); err != nil {
		t.Fatal(err)
	}
}

func Func(f func() error) (err error) {
	stop := time.Now().Add(2 * time.Second)
	for time.Now().Before(stop) {
		err = f()
		if err == nil {
			return
		}
		time.Sleep(25 * time.Millisecond)
	}
	return
}
