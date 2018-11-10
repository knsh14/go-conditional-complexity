package testdata

func simpleSelect(ch chan int) {
	select {
	case <-ch:
		println("foo")
	}
}
