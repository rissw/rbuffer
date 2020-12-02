package rbuffer

// RBuf is the main struct in rbuffer package
type RBuf struct {
	B        []byte
	i        int
	capacity int
	length   int
}

// New create buff with i byte size capacity , zero length
func New(i int) *RBuf {
	return &RBuf{
		B:        make([]byte, 0, i),
		i:        i,
		capacity: i,
		length:   0,
	}
}

func (r *RBuf) expand() {
	newBuffer := make([]byte, r.length, r.capacity+r.i)
	copy(newBuffer, r.B)
	r.B = newBuffer
	r.capacity = cap(r.B)
}

// Reset will zero-ed the data, and set the initial buffer as created
func (r *RBuf) Reset() {
	r.B = make([]byte, 0, r.i)
	r.length = 0
	r.capacity = r.i
}

// Cap returns the buffer capacity
func (r *RBuf) Cap() int {
	return r.capacity
}

// Len returns the buffer data length
func (r *RBuf) Len() int {
	return r.length
}

// Push puts b data in front of the slice
func (r *RBuf) Push(b []byte) {
	lenData := len(b)
	if r.capacity-r.length < lenData {
		for { // expand until it fits the data
			r.expand()
			if r.capacity-r.length >= lenData {
				break
			}
		}
	}
	r.B = append(r.B, b...)
	copy(r.B[lenData:], r.B[:r.length])
	copy(r.B, b)
	r.length = len(r.B)
}

// PushByte put b byte in front of the slice
func (r *RBuf) PushByte(b byte) {
	if r.capacity == r.length {
		r.expand()
	}
	r.B = append(r.B, b)
	copy(r.B[1:], r.B[:r.length])
	r.B[0] = b
	r.length++
}

// Append adds b data after the slice
func (r *RBuf) Append(b []byte) {
	lenData := len(b)
	if r.capacity-r.length < lenData {
		for { // expand until it fits the data
			r.expand()
			if r.capacity-r.length >= lenData {
				break
			}
		}
	}
	r.B = append(r.B, b...)
	r.length = len(r.B)
}

// AppendByte put b byte in front of the slice
func (r *RBuf) AppendByte(b byte) {
	if r.capacity == r.length {
		r.expand()
	}
	r.B = append(r.B, b)
	r.length++
}
