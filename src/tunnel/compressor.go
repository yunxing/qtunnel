package tunnel

import (
    "compress/zlib"
    "io"
    "bytes"
)

// Compressor encapsulates the logic of compressing and decompressing
// a stream
type Compressor struct {
    wrapped io.ReadWriter
}

// NewCompressor returns a new compressor
func NewCompressor(rw io.ReadWriter) *Compressor {
	return &Compressor{rw}
}

func (c *Compressor) Read(b []byte) (int, error) {
	buffer := make([]byte, len(b))
	n, err := c.wrapped.Read(buffer)
	if err != nil {
		return n, err
	}
	reader, err := zlib.NewReader(bytes.NewBuffer(buffer[:n]))
	if err != nil {
		return 0, err
	}
	defer func(){
		reader.Close()
	}()
	return reader.Read(b)
}

func (c *Compressor) Write(b []byte) (int, error) {
	writer := zlib.NewWriter(c.wrapped)
	defer func(){
		writer.Close()
	}()
	return writer.Write(b)
}
