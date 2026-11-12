package github

// Buffer accumulates rendered HTML.
type Buffer struct {
	Data []byte
}

func (b *Buffer) Put(value string) { b.Data = append(b.Data, value...) }
func (b *Buffer) Putc(c byte)      { b.Data = append(b.Data, c) }
func (b *Buffer) String() string   { return string(b.Data) }

func writeAttributeEscape(out *Buffer, value string) {
	for i := 0; i < len(value); i++ {
		switch value[i] {
		case '&':
			out.Put("&amp;")
		case '"':
			out.Put("&quot;")
		default:
			out.Putc(value[i])
		}
	}
}
