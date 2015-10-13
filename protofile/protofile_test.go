package protofile

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	oproto "github.com/dparrish/openinstrument/proto"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestWriteFile(c *C) {
	filename := filepath.Join(c.MkDir(), "protofile_testwrite.dat")

	{
		// Write a temporary file containing two labels
		writer, err := Write(filename)
		c.Assert(err, IsNil)
		defer writer.Close()
		c.Assert(writer.Tell(), Equals, int64(0))

		msg := &oproto.Label{
			Label: "greeting",
			Value: "Hello world!",
		}
		i, err := writer.Write(msg)
		c.Assert(err, IsNil)
		c.Assert(i, Equals, int64(32))
		j := i
		c.Assert(writer.Tell(), Equals, j)

		msg = &oproto.Label{
			Label: "greeting",
			Value: "Hola!",
		}
		i, err = writer.Write(msg)
		j += i
		c.Assert(i, Equals, int64(25))
		c.Assert(err, IsNil)
		c.Assert(writer.Tell(), Equals, j)

		// Write to a specific place in the file
		msg = &oproto.Label{
			Label: "greeting",
			Value: "Far out man",
		}
		i, err = writer.WriteAt(60, msg)
		c.Assert(i, Equals, int64(31))
		c.Assert(writer.Tell(), Equals, i+60)
		c.Assert(err, IsNil)

	}

	{
		// Read back the contents and ensure they are the same
		reader, err := Read(filename)
		c.Assert(err, IsNil)
		defer reader.Close()
		c.Assert(reader.Tell(), Equals, int64(0))

		msg := &oproto.Label{}
		i, err := reader.Read(msg)
		c.Assert(err, IsNil)
		c.Assert(i, Equals, int64(32))
		c.Assert(msg.Label, Equals, "greeting")
		c.Assert(msg.Value, Equals, "Hello world!")
		j := i
		c.Assert(reader.Tell(), Equals, j)

		msg = &oproto.Label{}
		i, err = reader.Read(msg)
		j += i
		c.Assert(err, IsNil)
		c.Assert(i, Equals, int64(25))
		c.Assert(msg.Label, Equals, "greeting")
		c.Assert(msg.Value, Equals, "Hola!")
		c.Assert(reader.Tell(), Equals, j)

		// Read the next message, which is after a few random bytes
		msg = &oproto.Label{}
		i, err = reader.Read(msg)
		j = 60 + i
		c.Assert(err, IsNil)
		c.Assert(i, Equals, int64(31))
		c.Assert(msg.Label, Equals, "greeting")
		c.Assert(msg.Value, Equals, "Far out man")
		c.Assert(reader.Tell(), Equals, j)

		// Read from a specific place in the file
		msg = &oproto.Label{}
		i, err = reader.ReadAt(60, msg)
		c.Assert(err, IsNil)
		c.Assert(i, Equals, int64(31))
		c.Assert(msg.Label, Equals, "greeting")
		c.Assert(msg.Value, Equals, "Far out man")
		c.Assert(reader.Tell(), Equals, j)

		// An attempt to read past the end of the file should return an error
		msg = &oproto.Label{}
		i, err = reader.Read(msg)
		c.Assert(err, ErrorMatches, "EOF")
		c.Assert(i, Equals, int64(0))
	}
}

func (s *MySuite) TestValueStreamReader(c *C) {
	filename := filepath.Join(c.MkDir(), "protofile_testvar.dat")

	{
		// Write a temporary file containing two value streams
		writer, err := Write(filename)
		c.Assert(err, IsNil)
		defer writer.Close()

		vs := &oproto.ValueStream{
			Variable: &oproto.StreamVariable{Name: "/test/bar"},
			Value: []*oproto.Value{
				{Timestamp: uint64(1), DoubleValue: 1.1},
				{Timestamp: uint64(2), DoubleValue: 1.2},
				{Timestamp: uint64(3), DoubleValue: 1.3},
			},
		}
		writer.Write(vs)

		vs = &oproto.ValueStream{
			Variable: &oproto.StreamVariable{Name: "/test/foo"},
			Value: []*oproto.Value{
				{Timestamp: uint64(1), DoubleValue: 1.1},
				{Timestamp: uint64(2), DoubleValue: 1.2},
				{Timestamp: uint64(3), DoubleValue: 1.3},
			},
		}
		writer.Write(vs)
	}

	{
		// Read back the contents and check
		file, err := Read(filename)
		c.Assert(err, IsNil)
		defer file.Close()
		reader := file.ValueStreamReader(500)
		vs := <-reader
		c.Check(vs.Variable.Name, Equals, "/test/bar")
		c.Check(vs.Value[0].DoubleValue, Equals, 1.1)
		c.Check(vs.Value[1].DoubleValue, Equals, 1.2)
		c.Check(vs.Value[2].DoubleValue, Equals, 1.3)

		vs = <-reader
		c.Check(vs.Variable.Name, Equals, "/test/foo")
		c.Check(vs.Value[0].DoubleValue, Equals, 1.1)
		c.Check(vs.Value[1].DoubleValue, Equals, 1.2)
		c.Check(vs.Value[2].DoubleValue, Equals, 1.3)

		for range reader {
			log.Printf("Got unexpected value")
			c.Fail()
		}
	}
}

func (s *MySuite) TestValueStreamWriter(c *C) {
	filename := filepath.Join(c.MkDir(), "protofile_testvar.dat")

	{
		// Write a temporary file containing two value streams
		file, err := Write(filename)
		c.Assert(err, IsNil)
		defer file.Close()
		writer, done := file.ValueStreamWriter(10)

		vs := &oproto.ValueStream{
			Variable: &oproto.StreamVariable{Name: "/test/bar"},
			Value: []*oproto.Value{
				{Timestamp: uint64(1), DoubleValue: 1.1},
				{Timestamp: uint64(2), DoubleValue: 1.2},
				{Timestamp: uint64(3), DoubleValue: 1.3},
			},
		}
		writer <- vs

		vs = &oproto.ValueStream{
			Variable: &oproto.StreamVariable{Name: "/test/foo"},
			Value: []*oproto.Value{
				{Timestamp: uint64(1), DoubleValue: 1.1},
				{Timestamp: uint64(2), DoubleValue: 1.2},
				{Timestamp: uint64(3), DoubleValue: 1.3},
			},
		}
		writer <- vs
		close(writer)
		<-done
	}

	{
		// Read back the contents and check
		file, err := Read(filename)
		c.Assert(err, IsNil)
		defer file.Close()
		reader := file.ValueStreamReader(500)
		vs := <-reader
		c.Check(vs.Variable.Name, Equals, "/test/bar")
		c.Check(vs.Value[0].DoubleValue, Equals, 1.1)
		c.Check(vs.Value[1].DoubleValue, Equals, 1.2)
		c.Check(vs.Value[2].DoubleValue, Equals, 1.3)

		vs = <-reader
		c.Check(vs.Variable.Name, Equals, "/test/foo")
		c.Check(vs.Value[0].DoubleValue, Equals, 1.1)
		c.Check(vs.Value[1].DoubleValue, Equals, 1.2)
		c.Check(vs.Value[2].DoubleValue, Equals, 1.3)

		for range reader {
			log.Printf("Got unexpected value")
			c.Fail()
		}
	}
}

/*
func (s *MySuite) TestValueStreamWriterMemoryLeak(c *C) {
	filename := filepath.Join(c.MkDir(), "protofile_testvar.dat")
	memstats := &runtime.MemStats{}

	for i := 0; i < 5; i++ {
		vs := &oproto.ValueStream{
			Variable: &oproto.StreamVariable{Name: "/test/bar"},
			Value:    []*oproto.Value{},
		}
		for i := 0; i < 600000; i++ {
			vs.Value = append(vs.Value, &oproto.Value{Timestamp: uint64(openinstrument.NowMs()), DoubleValue: 1.1})
		}
		// Write a temporary file containing two value streams
		file, err := Write(filename)
		c.Assert(err, IsNil)
		defer file.Close()
		writer, done := file.ValueStreamWriter(10)

		for j := 0; j < 100; j++ {
			writer <- vs
		}

		close(writer)
		<-done

		runtime.ReadMemStats(memstats)
		log.Printf("After iteration %d: %d Alloc, %d Sys, %d HeapAlloc, %d HeapSys", i, memstats.Alloc, memstats.Sys, memstats.HeapAlloc, memstats.HeapSys)
	}
	c.Fail()
}
*/

func (s *MySuite) BenchmarkReader(c *C) {
	filename := filepath.Join(c.MkDir(), "protofile_testvar.dat")

	{
		// Write a temporary file containing lots of data
		file, err := Write(filename)
		c.Assert(err, IsNil)
		writer, done := file.ValueStreamWriter(10)

		for i := 0; i < 10000; i++ {
			vs := &oproto.ValueStream{
				Variable: &oproto.StreamVariable{Name: "/test/bar"},
				Value: []*oproto.Value{
					{Timestamp: uint64(i), DoubleValue: float64(i)},
				},
			}
			writer <- vs
		}
		close(writer)
		<-done
		file.Close()
	}

	for run := 0; run < c.N; run++ {
		// Read back the contents
		file, err := Read(filename)
		c.Assert(err, IsNil)
		defer file.Close()
		reader := file.ValueStreamReader(500)
		for range reader {
		}
	}
}

// Write 100000 ValueStreams containing 1 value each
func (s *MySuite) BenchmarkWriterManyStreams(c *C) {
	filename := filepath.Join(c.MkDir(), "protofile_testvar.dat")

	vs := &oproto.ValueStream{
		Variable: &oproto.StreamVariable{Name: "/test/bar"},
		Value: []*oproto.Value{
			{Timestamp: uint64(1), DoubleValue: float64(1.1)},
		},
	}

	for run := 0; run < c.N; run++ {
		os.Remove(filename)
		file, err := Write(filename)
		c.Assert(err, IsNil)
		defer file.Close()
		writer, done := file.ValueStreamWriter(100000)

		for i := 0; i < 100000; i++ {
			writer <- vs
		}
		close(writer)
		<-done
	}

	file, _ := Read(filename)
	defer file.Close()
	stat, _ := file.Stat()
	log.Printf("BenchmarkWriterManyStreams wrote %d kB", stat.Size()/1024)
}

// Write 1 ValueStream with 100000 values
func (s *MySuite) BenchmarkWriterManyValues(c *C) {
	filename := filepath.Join(c.MkDir(), "protofile_testvar.dat")

	vs := &oproto.ValueStream{
		Variable: &oproto.StreamVariable{Name: "/test/bar"},
		Value:    []*oproto.Value{},
	}
	for j := 0; j < 100000; j++ {
		vs.Value = append(vs.Value, &oproto.Value{Timestamp: uint64(j), DoubleValue: float64(1.1)})
	}

	for run := 0; run < c.N; run++ {
		os.Remove(filename)
		file, err := Write(filename)
		c.Assert(err, IsNil)
		defer file.Close()
		writer, done := file.ValueStreamWriter(10)
		writer <- vs
		close(writer)
		<-done
	}

	file, _ := Read(filename)
	defer file.Close()
	stat, _ := file.Stat()
	log.Printf("BenchmarkWriterManyValues wrote %d kB", stat.Size()/1024)
}
