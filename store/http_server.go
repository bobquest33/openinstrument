package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/dparrish/openinstrument"
	"github.com/dparrish/openinstrument/datastore"
	"github.com/dparrish/openinstrument/mutations"
	oproto "github.com/dparrish/openinstrument/proto"
	"github.com/dparrish/openinstrument/store_config"
	"github.com/dparrish/openinstrument/valuestream"
	"github.com/dparrish/openinstrument/variable"
	"github.com/golang/protobuf/proto"
)

var (
	port         = flag.Int("port", 8020, "HTTP Port to listen on")
	templatePath = flag.String("templates", "/html", "Path to HTML template files")
)

func parseRequest(w http.ResponseWriter, req *http.Request, request proto.Message) error {
	body, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	encodedBody, err := base64.StdEncoding.DecodeString(string(body))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid body: %s", err)
		return errors.New("Invalid body")
	}
	if err = proto.Unmarshal(encodedBody, request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid GetRequest: %s", err)
		return errors.New("Invalid GetRequest")
	}
	//log.Printf("Request: %s", request)
	return nil
}

func returnResponse(w http.ResponseWriter, req *http.Request, response proto.Message) error {
	//log.Printf("Response: %s", response)
	data, err := proto.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error encoding response: %s", err)
		return errors.New("Error encoding response")
	}
	//encoded_data := base64.StdEncoding.EncodeToString(data)
	//w.Write([]byte(encoded_data))
	encoder := base64.NewEncoder(base64.StdEncoding, w)
	encoder.Write(data)
	encoder.Close()
	return nil
}

func Get(w http.ResponseWriter, req *http.Request) {
	var request oproto.GetRequest
	var response oproto.GetResponse
	if parseRequest(w, req, &request) != nil {
		return
	}
	if request.GetVariable() == nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Success = false
		response.Errormessage = "No variable specified"
		returnResponse(w, req, &response)
		return
	}
	requestVariable := variable.NewFromProto(request.Variable)
	if len(requestVariable.Variable) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		response.Success = false
		response.Errormessage = "No variable specified"
		returnResponse(w, req, &response)
		return
	}
	fmt.Println(openinstrument.ProtoText(&request))
	streamChan := ds.Reader(requestVariable, request.MinTimestamp, request.MaxTimestamp)
	streams := make([]*oproto.ValueStream, 0)
	for stream := range streamChan {
		streams = append(streams, stream)
	}
	mergeBy := ""
	if len(request.Aggregation) > 0 {
		mergeBy = request.Aggregation[0].Label[0]
	}
	sc := valuestream.MergeBy(streams, mergeBy)
	response.Stream = make([]*oproto.ValueStream, 0)
	for streams := range sc {
		output := valuestream.Merge(streams)

		if request.Mutation != nil && len(request.Mutation) > 0 {
			for _, mut := range request.Mutation {
				switch mut.SampleType {
				case oproto.StreamMutation_AVERAGE:
					output = mutations.Mean(uint64(mut.SampleFrequency), output)
				case oproto.StreamMutation_MIN:
					output = mutations.Min(uint64(mut.SampleFrequency), output)
				case oproto.StreamMutation_MAX:
					output = mutations.Max(uint64(mut.SampleFrequency), output)
				case oproto.StreamMutation_RATE:
					output = mutations.Rate(uint64(mut.SampleFrequency), output)
				case oproto.StreamMutation_RATE_SIGNED:
					output = mutations.SignedRate(uint64(mut.SampleFrequency), output)
				}
			}
		}

		newstream := new(oproto.ValueStream)
		newstream.Variable = variable.NewFromProto(streams[0].Variable).AsProto()
		var valueCount uint32
		for value := range output {
			if request.MinTimestamp != 0 && value.Timestamp < request.MinTimestamp {
				// Too old
				continue
			}
			if request.MaxTimestamp != 0 && value.Timestamp > request.MaxTimestamp {
				// Too new
				continue
			}
			newstream.Value = append(newstream.Value, value)
			valueCount++
		}

		if request.MaxValues != 0 && valueCount >= request.MaxValues {
			newstream.Value = newstream.Value[uint32(len(newstream.Value))-request.MaxValues:]
		}

		response.Stream = append(response.Stream, newstream)
	}
	response.Success = true
	returnResponse(w, req, &response)
}

func Add(w http.ResponseWriter, req *http.Request) {
	var request oproto.AddRequest
	var response oproto.AddResponse
	if parseRequest(w, req, &request) != nil {
		return
	}

	c := ds.Writer()
	for _, stream := range request.Stream {
		c <- stream
	}
	close(c)

	response.Success = true
	returnResponse(w, req, &response)
}

func ListResponseAddTimer(name string, response *oproto.ListResponse) *openinstrument.Timer {
	response.Timer = append(response.Timer, &oproto.LogMessage{})
	return openinstrument.NewTimer(name, response.Timer[len(response.Timer)-1])
}

func List(w http.ResponseWriter, req *http.Request) {
	var request oproto.ListRequest
	var response oproto.ListResponse
	response.Timer = make([]*oproto.LogMessage, 0)
	if parseRequest(w, req, &request) != nil {
		return
	}
	fmt.Println(openinstrument.ProtoText(&request))

	requestVariable := variable.NewFromProto(request.Prefix)
	if len(requestVariable.Variable) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		response.Success = false
		response.Errormessage = "No variable specified"
		returnResponse(w, req, &response)
		return
	}

	// Retrieve all variables and store the names in a map for uniqueness
	timer := ListResponseAddTimer("retrieve variables", &response)
	vars := make(map[string]*oproto.StreamVariable)
	minTimestamp := time.Now().Add(time.Duration(-request.MaxAge) * time.Millisecond)
	unix := uint64(minTimestamp.Unix()) * 1000
	streamChan := ds.Reader(requestVariable, unix, 0)
	for stream := range streamChan {
		vars[variable.NewFromProto(stream.Variable).String()] = stream.Variable
		if request.MaxVariables > 0 && len(vars) == int(request.MaxVariables) {
			break
		}
	}
	timer.Stop()

	// Build the response out of the map
	timer = ListResponseAddTimer("construct response", &response)
	response.Variable = make([]*oproto.StreamVariable, 0)
	for varname := range vars {
		response.Variable = append(response.Variable, variable.NewFromString(varname).AsProto())
	}
	response.Success = true
	timer.Stop()
	returnResponse(w, req, &response)
	log.Printf("Timers: %s", response.Timer)
}

// Argument server.
func Args(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, os.Args[1:])
}

func GetConfig(w http.ResponseWriter, req *http.Request) {
	returnResponse(w, req, store_config.Config().Config)
}

func StoreStatus(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles(fmt.Sprintf("%s/store_status.html", *templatePath))
	if err != nil {
		log.Printf("Couldn't find template file: %s", err)
		return
	}
	b := []*oproto.Block{}
	for _, block := range ds.Blocks {
		b = append(b, block.ToProto())
	}
	datastore.ProtoBlockBy(func(a, b *oproto.Block) bool { return a.EndKey < b.EndKey }).Sort(b)
	p := struct {
		Title  string
		Blocks []*oproto.Block
	}{
		Title:  "Store Status",
		Blocks: b,
	}
	err = t.Execute(w, p)
	if err != nil {
		log.Println(err)
	}
}

func CompactBlock(w http.ResponseWriter, req *http.Request) {
	for _, block := range ds.Blocks {
		if !strings.HasPrefix(strings.ToLower(block.ID), strings.ToLower(req.FormValue("block"))) {
			continue
		}
		log.Printf("User request compaction of block %s", req.FormValue("block"))
		block.RequestCompact()
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "compaction requested\n")
	}
}

func InspectVariable(w http.ResponseWriter, req *http.Request) {
	//w.WriteHeader(http.StatusOK)
	t, err := template.ParseFiles(fmt.Sprintf("%s/inspect_variable.html", *templatePath))
	if err != nil {
		log.Printf("Couldn't find template file: %s", err)
		return
	}
	type varInfo struct {
		Name           string
		FirstTimestamp time.Time
		LastTimestamp  time.Time
	}
	p := struct {
		Title     string
		Query     string
		Variables []varInfo
	}{
		Title:     "Inspect Variable",
		Query:     req.FormValue("q"),
		Variables: make([]varInfo, 0),
	}

	if p.Query == "" {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Specify q=")
		return
	}

	v := variable.NewFromString(p.Query)
	c := ds.Reader(v, 0, 0)
	for stream := range c {
		lt := stream.Value[len(stream.Value)-1].EndTimestamp
		if lt == 0 {
			lt = stream.Value[len(stream.Value)-1].Timestamp
		}
		p.Variables = append(p.Variables, varInfo{
			Name:           variable.NewFromProto(stream.Variable).String(),
			FirstTimestamp: time.Unix(int64(stream.Value[0].Timestamp/1000), 0),
			LastTimestamp:  time.Unix(int64(lt/1000), 0),
		})
	}

	err = t.Execute(w, p)
	if err != nil {
		log.Println(err)
	}
}

func Query(w http.ResponseWriter, req *http.Request) {
	query := req.FormValue("q")
	showValues := req.FormValue("v") == "1"
	type Result struct {
		Variable string          `json:"name"`
		Values   [][]interface{} `json:"values"`
	}

	var minTimestamp uint64
	var maxTimestamp uint64
	var duration *time.Duration
	if req.FormValue("d") != "" {
		d, err := time.ParseDuration(req.FormValue("d"))
		if err != nil {
			w.WriteHeader(401)
			fmt.Fprintf(w, "Invalid duration")
			return
		}
		duration = &d
		t := uint64(time.Now().UnixNano()-d.Nanoseconds()) / 1000000
		minTimestamp = t
	}

	if query == "" {
		w.WriteHeader(401)
		fmt.Fprintf(w, "Specify q=")
		return
	}

	results := make([]Result, 0)

	for stream := range ds.Reader(variable.NewFromString(query), minTimestamp, maxTimestamp) {
		r := Result{
			Variable: variable.NewFromProto(stream.Variable).String(),
		}
		if !showValues {
			results = append(results, r)
			continue
		}
		r.Values = make([][]interface{}, 0)

		if duration == nil {
			// Latest value only
			if len(stream.Value) > 0 {
				v := stream.Value[len(stream.Value)-1]
				r.Values = append(r.Values, []interface{}{v.Timestamp, v.DoubleValue})
				//r.Values = append(r.Values, v)
			}
		} else {
			// All values over a specific time period
			for _, v := range stream.Value {
				if minTimestamp == 0 || minTimestamp > v.Timestamp {
					r.Values = append(r.Values, []interface{}{v.Timestamp, v.DoubleValue})
				}
			}
		}
		results = append(results, r)
	}

	b, err := json.Marshal(results)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Couldn't marshal: %s", err)
		return
	}
	w.WriteHeader(200)
	w.Write(b)
}

func serveHTTP() {
	sock, e := net.ListenTCP("tcp", &net.TCPAddr{IP: net.ParseIP(*address), Port: *port})
	if e != nil {
		log.Fatalf("Can't listen on %s: %s", net.JoinHostPort(*address, strconv.Itoa(*port)), e)
	}
	log.Printf("Serving HTTP on %v", sock.Addr().String())

	http.Handle("/list", http.HandlerFunc(List))
	http.Handle("/get", http.HandlerFunc(Get))
	http.Handle("/add", http.HandlerFunc(Add))
	http.Handle("/args", http.HandlerFunc(Args))
	http.Handle("/config", http.HandlerFunc(GetConfig))
	http.Handle("/status", http.HandlerFunc(StoreStatus))
	http.Handle("/compact", http.HandlerFunc(CompactBlock))
	http.Handle("/inspect", http.HandlerFunc(InspectVariable))
	http.Handle("/query", http.HandlerFunc(Query))
	http.Serve(sock, nil)
}
