package tunnel

import (
    "net"
    "io"
    "fmt"
    snappy "github.com/mreiferson/go-snappystream"
)

type Conn struct {
    reader   io.Reader
    writer   io.WriteCloser
    cipher *Cipher
    pool *recycler
}

func NewConn(conn net.Conn, compress bool, cipher *Cipher, pool *recycler) *Conn {
    var reader io.Reader= conn
    var writer io.WriteCloser= conn
    if compress {
        reader = snappy.NewReader(conn, false)
        writer = snappy.NewBufferedWriter(conn)
    }
    return &Conn{
        reader: reader,
        writer: writer,
        cipher: cipher,
        pool: pool,
    }
}

func (c *Conn) Read(b []byte) (int, error) {
    if c.cipher == nil {
        return c.reader.Read(b)
    }
    n, err := c.reader.Read(b)
    fmt.Printf("here5.5 %v\n", err)
    if n > 0 {
        fmt.Println("here6")
        c.cipher.decrypt(b[0:n], b[0:n])
        fmt.Println("here7")
    }
    return n, err
}

func (c *Conn) Write(b []byte) (int, error) {
    if c.cipher == nil {
        return c.writer.Write(b)
    }
    c.cipher.encrypt(b, b)
    return c.writer.Write(b)
}

func (c *Conn) Close() error {
    return c.writer.Close()
}
