package variable

import (
  "code.google.com/p/goprotobuf/proto"
  openinstrument_proto "github.com/dparrish/openinstrument/proto"
  . "gopkg.in/check.v1"
  "testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestVariableCreation(c *C) {
  newvar := NewFromString("/openinstrument/test{label1=value1,label2=spaced value}")
  c.Check(newvar.String(), Equals, "/openinstrument/test{label1=value1,label2=spaced value}")

  labels := make([]*openinstrument_proto.Label, 2)
  labels[0] = &openinstrument_proto.Label{Label: proto.String("label1"), Value: proto.String("value1")}
  labels[1] = &openinstrument_proto.Label{Label: proto.String("label2"), Value: proto.String("spaced value")}
  newvar = NewFromProto(&openinstrument_proto.StreamVariable{
    Name:  proto.String("/openinstrument/test"),
    Label: labels,
  })
  c.Check(newvar.String(), Equals, "/openinstrument/test{label1=value1,label2=spaced value}")

  p := newvar.AsProto()
  newvar.ParseFromProto(p)
  c.Check(newvar.String(), Equals, "/openinstrument/test{label1=value1,label2=spaced value}")
}

func (s *MySuite) TestToProto(c *C) {
  newvar := NewFromString("/openinstrument/test{label1=value1,label2=spaced value}")
  p := newvar.AsProto()
  newvar.ParseFromProto(p)
  c.Check(newvar.String(), Equals, "/openinstrument/test{label1=value1,label2=spaced value}")
}

func (s *MySuite) TestMatch(c *C) {
  var1 := NewFromString("/openinstrument/test{label1=value1,label2=spaced value}")
  c.Check(var1.Match(NewFromString("/openinstrument/*")), Equals, true)
  c.Check(var1.Match(NewFromString("/openinstrx/*")), Equals, false)
  c.Check(var1.Match(NewFromString("/openinstrument/")), Equals, false)
  c.Check(var1.Match(NewFromString("/openinstrument/test")), Equals, true)
  c.Check(var1.Match(NewFromString("/openinstrument/test{label1=value1}")), Equals, true)
  c.Check(var1.Match(NewFromString("/openinstrument/test{label1=value2}")), Equals, false)
  c.Check(var1.Match(NewFromString("/openinstrument/test{label1=*}")), Equals, true)
  c.Check(var1.Match(NewFromString("/openinstrument/test{label1=}")), Equals, false)
  c.Check(var1.Match(NewFromString("/openinstrument/test{label3=*}")), Equals, false)
}
