package valuestream

import (
	"sort"

	oproto "github.com/dparrish/openinstrument/proto"
	"github.com/dparrish/openinstrument/variable"
)

type By func(p1, p2 *oproto.ValueStream) bool

func (by By) Sort(values []*oproto.ValueStream) {
	sfs := &valuesSorter{
		values: values,
		by:     by,
	}
	sort.Sort(sfs)
}

type valuesSorter struct {
	values []*oproto.ValueStream
	by     By
}

func (vs *valuesSorter) Len() int {
	return len(vs.values)
}

func (vs *valuesSorter) Swap(i, j int) {
	vs.values[i], vs.values[j] = vs.values[j], vs.values[i]
}

func (vs *valuesSorter) Less(i, j int) bool {
	return vs.by(vs.values[i], vs.values[j])
}

// Merge merges multiple ValueStreams, returning a channel producing sorted Values.
func Merge(streams []*oproto.ValueStream) <-chan *oproto.Value {
	c := make(chan *oproto.Value)
	go func() {
		n := len(streams)
		indexes := make([]int, n)
		var minTimestamp uint64
		var minStream int
		var minValue *oproto.Value
		for {
			minTimestamp = 0
			minStream = -1
			minValue = nil
			for i := 0; i < n; i++ {
				if indexes[i] >= len(streams[i].Value) {
					continue
				}
				v := streams[i].Value[indexes[i]]
				if minStream == -1 || v.Timestamp < minTimestamp {
					minTimestamp = v.Timestamp
					minStream = i
					minValue = v
				}
			}
			if minValue == nil {
				break
			}
			indexes[minStream]++
			c <- minValue
		}
		close(c)
	}()
	return c
}

func MergeBy(streams []*oproto.ValueStream, by string) <-chan []*oproto.ValueStream {
	c := make(chan []*oproto.ValueStream)
	go func() {
		uniqueVars := make(map[string]bool)
		uniqueLabels := make(map[string]bool)
		for _, stream := range streams {
			v := variable.NewFromProto(stream.Variable)
			uniqueVars[v.Variable] = true
			labelValue, ok := v.Labels[by]
			if !ok {
				uniqueLabels[""] = true
			} else {
				uniqueLabels[labelValue] = true
			}
		}
		for varname := range uniqueVars {
			v := variable.NewFromString(varname)
			if by == "" {
				var output []*oproto.ValueStream
				for _, stream := range streams {
					testvar := variable.NewFromProto(stream.Variable)
					if testvar.Variable != v.Variable {
						continue
					}
					output = append(output, stream)
				}
				if len(output) > 0 {
					c <- output
				}
			} else {
				for labelvalue := range uniqueLabels {
					var output []*oproto.ValueStream
					for _, stream := range streams {
						testvar := variable.NewFromProto(stream.Variable)
						if testvar.Variable != v.Variable {
							continue
						}
						value, ok := testvar.Labels[by]
						if ok && value == labelvalue {
							output = append(output, stream)
						}
					}
					if len(output) > 0 {
						c <- output
					}
				}
			}
		}
		close(c)
	}()
	return c
}

// FromChan takes a ValueStream and writes all the values individually to a new channel.
func ToChan(input *oproto.ValueStream) <-chan *oproto.Value {
	output := make(chan *oproto.Value)
	go func() {
		for _, v := range input.Value {
			output <- v
		}
		close(output)
	}()
	return output
}

// ToStream takes a channel of Values and appends each one to the supplied ValueStream.
func ToStream(input <-chan *oproto.Value, output *oproto.ValueStream) *oproto.ValueStream {
	for v := range input {
		output.Value = append(output.Value, v)
	}
	return output
}
