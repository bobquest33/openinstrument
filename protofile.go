package openinstrument

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"

	"code.google.com/p/goprotobuf/proto"
	openinstrument_proto "github.com/dparrish/openinstrument/proto"
	"github.com/joaojeronimo/go-crc16"
)

const protoMagic uint16 = 0xDEAD

type ProtoFileReader struct {
	filename string
	file     *os.File
	stat     os.FileInfo
}

func ReadProtoFile(filename string) (*ProtoFileReader, error) {
	reader := new(ProtoFileReader)
	reader.filename = filename
	var err error
	reader.file, err = os.Open(filename)
	if err != nil {
		return nil, err
	}
	reader.stat, err = reader.file.Stat()
	if err != nil {
		reader.file.Close()
		return nil, err
	}
	return reader, nil
}

func (pfr *ProtoFileReader) Close() error {
	return pfr.file.Close()
}

func (pfr *ProtoFileReader) Tell() int64 {
	pos, _ := pfr.file.Seek(0, os.SEEK_CUR)
	return pos
}

func (pfr *ProtoFileReader) Seek(pos int64) int64 {
	npos, _ := pfr.file.Seek(pos, os.SEEK_SET)
	return npos
}

func (pfr *ProtoFileReader) Stat() (os.FileInfo, error) {
	return pfr.file.Stat()
}

func (pfr *ProtoFileReader) ReadAt(pos int64, message proto.Message) (int, error) {
	pfr.Seek(pos)
	return pfr.Read(message)
}

func (pfr *ProtoFileReader) ValueStreamReader(chanSize int) chan *openinstrument_proto.ValueStream {
	c := make(chan *openinstrument_proto.ValueStream, chanSize)
	go func() {
		for {
			value := new(openinstrument_proto.ValueStream)
			_, err := pfr.Read(value)
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println(err)
				break
			}
			c <- value
		}
		close(c)
	}()
	return c
}

func (pfr *ProtoFileReader) ValueStreamReaderUntil(maxPos uint64, chanSize int) chan *openinstrument_proto.ValueStream {
	c := make(chan *openinstrument_proto.ValueStream, chanSize)
	go func() {
		for uint64(pfr.Tell()) < maxPos {
			value := new(openinstrument_proto.ValueStream)
			_, err := pfr.Read(value)
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println(err)
				break
			}
			c <- value
		}
		close(c)
	}()
	return c
}

func (pfr *ProtoFileReader) Read(message proto.Message) (int, error) {
	for {
		pos := pfr.Tell()
		type header struct {
			Magic  uint16
			Length uint32
		}
		var h header

		err := binary.Read(pfr.file, binary.LittleEndian, &h)
		if err != nil {
			if err == io.EOF {
				return 0, io.EOF
			}
			log.Printf("Error reading record header from recordlog: %s", err)
			return 0, err
		}

		// Read Magic header
		if h.Magic != protoMagic {
			log.Printf("Protobuf delimeter at %s:%x does not match %#x", pfr.filename, pos, protoMagic)
			continue
		}
		if int64(h.Length) >= pfr.stat.Size() {
			log.Printf("Chunk length %d at %s:%x is greater than file size %d", h.Length, pfr.filename, pos, pfr.stat.Size())
			continue
		}

		// Read Proto
		buf := make([]byte, h.Length)
		n, err := pfr.file.Read(buf)
		if err != nil || uint32(n) != h.Length {
			log.Printf("Could not read %d bytes from file: %s", h.Length, err)
			return 0, io.EOF
		}

		// Read CRC
		var crc uint16
		err = binary.Read(pfr.file, binary.LittleEndian, &crc)
		if err != nil {
			log.Printf("Error reading CRC from recordlog: %s", err)
			continue
		}
		checkcrc := crc16.Crc16(buf)
		if checkcrc != crc {
			//log.Printf("CRC %x does not match %x", crc, checkcrc)
		}

		// Decode and add proto
		if err = proto.Unmarshal(buf, message); err != nil {
			return 0, fmt.Errorf("Error decoding protobuf at %s:%x: %s", pfr.filename, pos, err)
		}
		break
	}
	return 1, nil
}

func WriteProtoFile(filename string) (*ProtoFileWriter, error) {
	reader := new(ProtoFileWriter)
	reader.filename = filename
	var err error
	reader.file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		return nil, err
	}
	reader.stat, err = reader.file.Stat()
	if err != nil {
		reader.file.Close()
		return nil, err
	}
	reader.file.Seek(0, os.SEEK_END)
	return reader, nil
}

type ProtoFileWriter struct {
	filename string
	file     *os.File
	stat     os.FileInfo
}

func (pfw *ProtoFileWriter) Close() error {
	return pfw.file.Close()
}

func (pfw *ProtoFileWriter) Tell() int64 {
	pos, _ := pfw.file.Seek(0, os.SEEK_CUR)
	return pos
}

func (pfw *ProtoFileWriter) Stat() (os.FileInfo, error) {
	return pfw.file.Stat()
}

func (pfw *ProtoFileWriter) Sync() error {
	return pfw.file.Sync()
}

func (pfw *ProtoFileWriter) WriteAt(pos int64, message proto.Message) (int, error) {
	pfw.file.Seek(pos, os.SEEK_SET)
	return pfw.Write(message)
}

func (pfw *ProtoFileWriter) Write(message proto.Message) (int, error) {
	data, err := proto.Marshal(message)
	if err != nil {
		return 0, fmt.Errorf("Marshaling error: %s", err)
	}
	var buf = []interface{}{
		uint16(protoMagic),
		uint32(len(data)),
		data,
		uint16(0),
	}
	for _, v := range buf {
		err = binary.Write(pfw.file, binary.LittleEndian, v)
		if err != nil {
			return 0, fmt.Errorf("Error writing entry to protofile: %s", err)
		}
	}
	return 1, nil
}