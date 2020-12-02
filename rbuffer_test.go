package rbuffer

import "testing"

func TestAppend(t *testing.T) {
	rb := New(100)
	for i := byte(0); i < 101; i++ {
		rb.AppendByte(i)
	}
	if rb.Cap() != 200 || rb.Len() != 101 {
		t.Fail()
	}
}

func TestReset(t *testing.T) {
	rb := New(100)
	for i := byte(0); i < 101; i++ {
		rb.AppendByte(i)
	}
	if rb.Cap() != 200 || rb.Len() != 101 {
		t.FailNow()
	}
	rb.Reset()
	if rb.Cap() != 100 || rb.Len() != 0 {
		t.Fail()
	}
}

func TestAppendData(t *testing.T) {
	rb := New(5)
	rb.AppendByte('A')
	rb.Append([]byte("BCDEFGH"))
	if rb.capacity != 10 || rb.length != 8 || string(rb.B) != "ABCDEFGH" {
		t.Fail()
	}
}

func TestPushData(t *testing.T) {
	rb := New(5)
	rb.AppendByte('A')
	rb.Append([]byte("BCDEFGH"))
	rb.PushByte('9')
	rb.Push([]byte("012345678"))
	if rb.capacity != 20 || rb.length != 18 || string(rb.B) != "0123456789ABCDEFGH" {
		t.Fail()
	}
}
