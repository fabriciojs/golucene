package tokenattributes

import (
	"github.com/balzaczyy/golucene/core/util"
)

/* The term text of a Token. */
type CharTermAttribute interface {
	// Copies the contents of buffer into the termBuffer array
	CopyBuffer(buffer []rune)
	// Returns the internal termBuffer rune slice which you can then
	// directly alter. If the slice is too small for your token, use
	// ResizeBuffer(int) to increase it. After altering the buffer, be
	// sure to call SetLength() to record the number of valid runes
	// that were placed into the termBuffer.
	//
	// NOTE: the returned buffer may be larger than the valid Length().
	Buffer() []rune
	Length() int
}

const MIN_BUFFER_SIZE = 10

/* Default implementation of CharTermAttribute. */
type CharTermAttributeImpl struct {
	termBuffer []rune
	termLength int
	bytes      *util.BytesRef
}

func newCharTermAttributeImpl() *util.AttributeImpl {
	ans := &CharTermAttributeImpl{
		termBuffer: make([]rune, util.Oversize(MIN_BUFFER_SIZE, util.NUM_BYTES_CHAR)),
		bytes:      util.NewBytesRef(make([]byte, 0, MIN_BUFFER_SIZE)),
	}
	return util.NewAttributeImpl(ans)
}

func (a *CharTermAttributeImpl) Interfaces() []string {
	return []string{"CharTermAttribute", "TermToBytesRefAttribute"}
}

func (a *CharTermAttributeImpl) CopyBuffer(buffer []rune) {
	a.growTermBuffer(len(buffer))
	copy(a.termBuffer, buffer)
	a.termLength = len(buffer)
}

func (a *CharTermAttributeImpl) Buffer() []rune {
	return a.termBuffer
}

func (a *CharTermAttributeImpl) growTermBuffer(newSize int) {
	if len(a.termBuffer) < newSize {
		// not big enough: create a new slice with slight over allocation:
		a.termBuffer = make([]rune, util.Oversize(newSize, util.NUM_BYTES_CHAR))
	}
}

func (a *CharTermAttributeImpl) FillBytesRef() int {
	s := string(a.termBuffer[:a.termLength])
	hash := hashstr(s)
	a.bytes.Value = []byte(s)
	return hash
}

const primeRK = 16777619

/* simple string hash used by Go strings package */
func hashstr(sep string) int {
	hash := uint32(0)
	for i := 0; i < len(sep); i++ {
		hash = hash*primeRK + uint32(sep[i])
	}
	return int(hash)
}

func (a *CharTermAttributeImpl) BytesRef() *util.BytesRef {
	return a.bytes
}

func (a *CharTermAttributeImpl) Length() int {
	return a.termLength
}

func (a *CharTermAttributeImpl) Clear() {
	a.termLength = 0
}

func (a *CharTermAttributeImpl) String() string {
	return string(a.termBuffer[:a.termLength])
}