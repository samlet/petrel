// Code generated by capnpc-go. DO NOT EDIT.

package hashes

import (
	context "golang.org/x/net/context"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
)

type HashFactory struct{ Client capnp.Client }

// HashFactory_TypeID is the unique identifier for the type HashFactory.
const HashFactory_TypeID = 0xaead580f97fddabc

func (c HashFactory) NewSha1(ctx context.Context, params func(HashFactory_newSha1_Params) error, opts ...capnp.CallOption) HashFactory_newSha1_Results_Promise {
	if c.Client == nil {
		return HashFactory_newSha1_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xaead580f97fddabc,
			MethodID:      0,
			InterfaceName: "hashes.capnp:HashFactory",
			MethodName:    "newSha1",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(HashFactory_newSha1_Params{Struct: s}) }
	}
	return HashFactory_newSha1_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type HashFactory_Server interface {
	NewSha1(HashFactory_newSha1) error
}

func HashFactory_ServerToClient(s HashFactory_Server) HashFactory {
	c, _ := s.(server.Closer)
	return HashFactory{Client: server.New(HashFactory_Methods(nil, s), c)}
}

func HashFactory_Methods(methods []server.Method, s HashFactory_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xaead580f97fddabc,
			MethodID:      0,
			InterfaceName: "hashes.capnp:HashFactory",
			MethodName:    "newSha1",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := HashFactory_newSha1{c, opts, HashFactory_newSha1_Params{Struct: p}, HashFactory_newSha1_Results{Struct: r}}
			return s.NewSha1(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// HashFactory_newSha1 holds the arguments for a server call to HashFactory.newSha1.
type HashFactory_newSha1 struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  HashFactory_newSha1_Params
	Results HashFactory_newSha1_Results
}

type HashFactory_newSha1_Params struct{ capnp.Struct }

// HashFactory_newSha1_Params_TypeID is the unique identifier for the type HashFactory_newSha1_Params.
const HashFactory_newSha1_Params_TypeID = 0x92b20ad1a58ca0ca

func NewHashFactory_newSha1_Params(s *capnp.Segment) (HashFactory_newSha1_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return HashFactory_newSha1_Params{st}, err
}

func NewRootHashFactory_newSha1_Params(s *capnp.Segment) (HashFactory_newSha1_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return HashFactory_newSha1_Params{st}, err
}

func ReadRootHashFactory_newSha1_Params(msg *capnp.Message) (HashFactory_newSha1_Params, error) {
	root, err := msg.RootPtr()
	return HashFactory_newSha1_Params{root.Struct()}, err
}

func (s HashFactory_newSha1_Params) String() string {
	str, _ := text.Marshal(0x92b20ad1a58ca0ca, s.Struct)
	return str
}

// HashFactory_newSha1_Params_List is a list of HashFactory_newSha1_Params.
type HashFactory_newSha1_Params_List struct{ capnp.List }

// NewHashFactory_newSha1_Params creates a new list of HashFactory_newSha1_Params.
func NewHashFactory_newSha1_Params_List(s *capnp.Segment, sz int32) (HashFactory_newSha1_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return HashFactory_newSha1_Params_List{l}, err
}

func (s HashFactory_newSha1_Params_List) At(i int) HashFactory_newSha1_Params {
	return HashFactory_newSha1_Params{s.List.Struct(i)}
}

func (s HashFactory_newSha1_Params_List) Set(i int, v HashFactory_newSha1_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s HashFactory_newSha1_Params_List) String() string {
	str, _ := text.MarshalList(0x92b20ad1a58ca0ca, s.List)
	return str
}

// HashFactory_newSha1_Params_Promise is a wrapper for a HashFactory_newSha1_Params promised by a client call.
type HashFactory_newSha1_Params_Promise struct{ *capnp.Pipeline }

func (p HashFactory_newSha1_Params_Promise) Struct() (HashFactory_newSha1_Params, error) {
	s, err := p.Pipeline.Struct()
	return HashFactory_newSha1_Params{s}, err
}

type HashFactory_newSha1_Results struct{ capnp.Struct }

// HashFactory_newSha1_Results_TypeID is the unique identifier for the type HashFactory_newSha1_Results.
const HashFactory_newSha1_Results_TypeID = 0xea3e50f7663f7bdf

func NewHashFactory_newSha1_Results(s *capnp.Segment) (HashFactory_newSha1_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HashFactory_newSha1_Results{st}, err
}

func NewRootHashFactory_newSha1_Results(s *capnp.Segment) (HashFactory_newSha1_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HashFactory_newSha1_Results{st}, err
}

func ReadRootHashFactory_newSha1_Results(msg *capnp.Message) (HashFactory_newSha1_Results, error) {
	root, err := msg.RootPtr()
	return HashFactory_newSha1_Results{root.Struct()}, err
}

func (s HashFactory_newSha1_Results) String() string {
	str, _ := text.Marshal(0xea3e50f7663f7bdf, s.Struct)
	return str
}

func (s HashFactory_newSha1_Results) Hash() Hash {
	p, _ := s.Struct.Ptr(0)
	return Hash{Client: p.Interface().Client()}
}

func (s HashFactory_newSha1_Results) HasHash() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HashFactory_newSha1_Results) SetHash(v Hash) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// HashFactory_newSha1_Results_List is a list of HashFactory_newSha1_Results.
type HashFactory_newSha1_Results_List struct{ capnp.List }

// NewHashFactory_newSha1_Results creates a new list of HashFactory_newSha1_Results.
func NewHashFactory_newSha1_Results_List(s *capnp.Segment, sz int32) (HashFactory_newSha1_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HashFactory_newSha1_Results_List{l}, err
}

func (s HashFactory_newSha1_Results_List) At(i int) HashFactory_newSha1_Results {
	return HashFactory_newSha1_Results{s.List.Struct(i)}
}

func (s HashFactory_newSha1_Results_List) Set(i int, v HashFactory_newSha1_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s HashFactory_newSha1_Results_List) String() string {
	str, _ := text.MarshalList(0xea3e50f7663f7bdf, s.List)
	return str
}

// HashFactory_newSha1_Results_Promise is a wrapper for a HashFactory_newSha1_Results promised by a client call.
type HashFactory_newSha1_Results_Promise struct{ *capnp.Pipeline }

func (p HashFactory_newSha1_Results_Promise) Struct() (HashFactory_newSha1_Results, error) {
	s, err := p.Pipeline.Struct()
	return HashFactory_newSha1_Results{s}, err
}

func (p HashFactory_newSha1_Results_Promise) Hash() Hash {
	return Hash{Client: p.Pipeline.GetPipeline(0).Client()}
}

type Hash struct{ Client capnp.Client }

// Hash_TypeID is the unique identifier for the type Hash.
const Hash_TypeID = 0xf29f97dd675a9431

func (c Hash) Write(ctx context.Context, params func(Hash_write_Params) error, opts ...capnp.CallOption) Hash_write_Results_Promise {
	if c.Client == nil {
		return Hash_write_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xf29f97dd675a9431,
			MethodID:      0,
			InterfaceName: "hashes.capnp:Hash",
			MethodName:    "write",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Hash_write_Params{Struct: s}) }
	}
	return Hash_write_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Hash) Sum(ctx context.Context, params func(Hash_sum_Params) error, opts ...capnp.CallOption) Hash_sum_Results_Promise {
	if c.Client == nil {
		return Hash_sum_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xf29f97dd675a9431,
			MethodID:      1,
			InterfaceName: "hashes.capnp:Hash",
			MethodName:    "sum",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Hash_sum_Params{Struct: s}) }
	}
	return Hash_sum_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Hash_Server interface {
	Write(Hash_write) error

	Sum(Hash_sum) error
}

func Hash_ServerToClient(s Hash_Server) Hash {
	c, _ := s.(server.Closer)
	return Hash{Client: server.New(Hash_Methods(nil, s), c)}
}

func Hash_Methods(methods []server.Method, s Hash_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xf29f97dd675a9431,
			MethodID:      0,
			InterfaceName: "hashes.capnp:Hash",
			MethodName:    "write",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Hash_write{c, opts, Hash_write_Params{Struct: p}, Hash_write_Results{Struct: r}}
			return s.Write(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xf29f97dd675a9431,
			MethodID:      1,
			InterfaceName: "hashes.capnp:Hash",
			MethodName:    "sum",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Hash_sum{c, opts, Hash_sum_Params{Struct: p}, Hash_sum_Results{Struct: r}}
			return s.Sum(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// Hash_write holds the arguments for a server call to Hash.write.
type Hash_write struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Hash_write_Params
	Results Hash_write_Results
}

// Hash_sum holds the arguments for a server call to Hash.sum.
type Hash_sum struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Hash_sum_Params
	Results Hash_sum_Results
}

type Hash_write_Params struct{ capnp.Struct }

// Hash_write_Params_TypeID is the unique identifier for the type Hash_write_Params.
const Hash_write_Params_TypeID = 0xdffe94ae546cdee3

func NewHash_write_Params(s *capnp.Segment) (Hash_write_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Hash_write_Params{st}, err
}

func NewRootHash_write_Params(s *capnp.Segment) (Hash_write_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Hash_write_Params{st}, err
}

func ReadRootHash_write_Params(msg *capnp.Message) (Hash_write_Params, error) {
	root, err := msg.RootPtr()
	return Hash_write_Params{root.Struct()}, err
}

func (s Hash_write_Params) String() string {
	str, _ := text.Marshal(0xdffe94ae546cdee3, s.Struct)
	return str
}

func (s Hash_write_Params) Data() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s Hash_write_Params) HasData() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Hash_write_Params) SetData(v []byte) error {
	return s.Struct.SetData(0, v)
}

// Hash_write_Params_List is a list of Hash_write_Params.
type Hash_write_Params_List struct{ capnp.List }

// NewHash_write_Params creates a new list of Hash_write_Params.
func NewHash_write_Params_List(s *capnp.Segment, sz int32) (Hash_write_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Hash_write_Params_List{l}, err
}

func (s Hash_write_Params_List) At(i int) Hash_write_Params {
	return Hash_write_Params{s.List.Struct(i)}
}

func (s Hash_write_Params_List) Set(i int, v Hash_write_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s Hash_write_Params_List) String() string {
	str, _ := text.MarshalList(0xdffe94ae546cdee3, s.List)
	return str
}

// Hash_write_Params_Promise is a wrapper for a Hash_write_Params promised by a client call.
type Hash_write_Params_Promise struct{ *capnp.Pipeline }

func (p Hash_write_Params_Promise) Struct() (Hash_write_Params, error) {
	s, err := p.Pipeline.Struct()
	return Hash_write_Params{s}, err
}

type Hash_write_Results struct{ capnp.Struct }

// Hash_write_Results_TypeID is the unique identifier for the type Hash_write_Results.
const Hash_write_Results_TypeID = 0x80ac741ec7fb8f65

func NewHash_write_Results(s *capnp.Segment) (Hash_write_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Hash_write_Results{st}, err
}

func NewRootHash_write_Results(s *capnp.Segment) (Hash_write_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Hash_write_Results{st}, err
}

func ReadRootHash_write_Results(msg *capnp.Message) (Hash_write_Results, error) {
	root, err := msg.RootPtr()
	return Hash_write_Results{root.Struct()}, err
}

func (s Hash_write_Results) String() string {
	str, _ := text.Marshal(0x80ac741ec7fb8f65, s.Struct)
	return str
}

// Hash_write_Results_List is a list of Hash_write_Results.
type Hash_write_Results_List struct{ capnp.List }

// NewHash_write_Results creates a new list of Hash_write_Results.
func NewHash_write_Results_List(s *capnp.Segment, sz int32) (Hash_write_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Hash_write_Results_List{l}, err
}

func (s Hash_write_Results_List) At(i int) Hash_write_Results {
	return Hash_write_Results{s.List.Struct(i)}
}

func (s Hash_write_Results_List) Set(i int, v Hash_write_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s Hash_write_Results_List) String() string {
	str, _ := text.MarshalList(0x80ac741ec7fb8f65, s.List)
	return str
}

// Hash_write_Results_Promise is a wrapper for a Hash_write_Results promised by a client call.
type Hash_write_Results_Promise struct{ *capnp.Pipeline }

func (p Hash_write_Results_Promise) Struct() (Hash_write_Results, error) {
	s, err := p.Pipeline.Struct()
	return Hash_write_Results{s}, err
}

type Hash_sum_Params struct{ capnp.Struct }

// Hash_sum_Params_TypeID is the unique identifier for the type Hash_sum_Params.
const Hash_sum_Params_TypeID = 0xe74bb2d0190cf89c

func NewHash_sum_Params(s *capnp.Segment) (Hash_sum_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Hash_sum_Params{st}, err
}

func NewRootHash_sum_Params(s *capnp.Segment) (Hash_sum_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Hash_sum_Params{st}, err
}

func ReadRootHash_sum_Params(msg *capnp.Message) (Hash_sum_Params, error) {
	root, err := msg.RootPtr()
	return Hash_sum_Params{root.Struct()}, err
}

func (s Hash_sum_Params) String() string {
	str, _ := text.Marshal(0xe74bb2d0190cf89c, s.Struct)
	return str
}

// Hash_sum_Params_List is a list of Hash_sum_Params.
type Hash_sum_Params_List struct{ capnp.List }

// NewHash_sum_Params creates a new list of Hash_sum_Params.
func NewHash_sum_Params_List(s *capnp.Segment, sz int32) (Hash_sum_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Hash_sum_Params_List{l}, err
}

func (s Hash_sum_Params_List) At(i int) Hash_sum_Params { return Hash_sum_Params{s.List.Struct(i)} }

func (s Hash_sum_Params_List) Set(i int, v Hash_sum_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s Hash_sum_Params_List) String() string {
	str, _ := text.MarshalList(0xe74bb2d0190cf89c, s.List)
	return str
}

// Hash_sum_Params_Promise is a wrapper for a Hash_sum_Params promised by a client call.
type Hash_sum_Params_Promise struct{ *capnp.Pipeline }

func (p Hash_sum_Params_Promise) Struct() (Hash_sum_Params, error) {
	s, err := p.Pipeline.Struct()
	return Hash_sum_Params{s}, err
}

type Hash_sum_Results struct{ capnp.Struct }

// Hash_sum_Results_TypeID is the unique identifier for the type Hash_sum_Results.
const Hash_sum_Results_TypeID = 0xd093963b95a4e107

func NewHash_sum_Results(s *capnp.Segment) (Hash_sum_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Hash_sum_Results{st}, err
}

func NewRootHash_sum_Results(s *capnp.Segment) (Hash_sum_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Hash_sum_Results{st}, err
}

func ReadRootHash_sum_Results(msg *capnp.Message) (Hash_sum_Results, error) {
	root, err := msg.RootPtr()
	return Hash_sum_Results{root.Struct()}, err
}

func (s Hash_sum_Results) String() string {
	str, _ := text.Marshal(0xd093963b95a4e107, s.Struct)
	return str
}

func (s Hash_sum_Results) Hash() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s Hash_sum_Results) HasHash() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Hash_sum_Results) SetHash(v []byte) error {
	return s.Struct.SetData(0, v)
}

// Hash_sum_Results_List is a list of Hash_sum_Results.
type Hash_sum_Results_List struct{ capnp.List }

// NewHash_sum_Results creates a new list of Hash_sum_Results.
func NewHash_sum_Results_List(s *capnp.Segment, sz int32) (Hash_sum_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Hash_sum_Results_List{l}, err
}

func (s Hash_sum_Results_List) At(i int) Hash_sum_Results { return Hash_sum_Results{s.List.Struct(i)} }

func (s Hash_sum_Results_List) Set(i int, v Hash_sum_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s Hash_sum_Results_List) String() string {
	str, _ := text.MarshalList(0xd093963b95a4e107, s.List)
	return str
}

// Hash_sum_Results_Promise is a wrapper for a Hash_sum_Results promised by a client call.
type Hash_sum_Results_Promise struct{ *capnp.Pipeline }

func (p Hash_sum_Results_Promise) Struct() (Hash_sum_Results, error) {
	s, err := p.Pipeline.Struct()
	return Hash_sum_Results{s}, err
}

const schema_db8274f9144abc7e = "x\xda|\x92?h\x14A\x18\xc5\xdf\x9b\xdbu\x83\xba" +
	"\x9c\xe3\xc4\xe2\x0a\x8d\xca\xd9\x089X\xd3)xG\xc0" +
	"\x7f\xb1\xd9\x8d\x0aj7\xc4\xd5\x15r1\xdc\xee\x11D" +
	"P\x11D\x04E\x88\xd14\x82\x16\xdai\x8a\xd4\xe9\x03" +
	"\xa9\xa2\x9d\x8a\x06\xb1\x10R\xa6\xf1\x1f\xba\xb2s\xeem" +
	"@/\xdd\xc2\x9b\xef}\xbf\xef\xed\xdb\xf6\xa6!<{" +
	"Q\x00\xc1n{S\x1a>\xf8\xb9\xb8+yy\x13r" +
	";\x01\xcb\x01\x86NP\x10V\xba\xf4\xec\xde\x8b\xd7\x9b" +
	"\xe7\xa7!+\xb92\xc8\xe1LYx\xf7k\xb6|\xf6" +
	"\xd5\x1c\xa4[J\xaf/\x8c\xf4\x7fOn\xbd\x07\xa8v" +
	"pI\xed\xa1\x03\xa8\x9d<\xa6\x8ed_\xa9\xf3\xe9\xf9" +
	"\xa3C\x8f\x1f.w\x16\xd8F\x1d\xe4\x1a\xa8<\xd6\xc1" +
	"\xf4\xf3\xc7\xf1\xd3s3\xbfW\xd6\xeb\x01\x7f\x80\xea\x8c" +
	"\xd1\x9f|\xdbZY\x9e?\xf9\xa5\x00Tm\xae\xc2J" +
	"W\xae\xd5/~\xf5\x0f\xafv\xf8\xcc\xe0\xd09\x8e\x10" +
	"T\xa1\x99\xf4f\xce_\xfa0\xfbt\xed\x1f\xcc\xdb\x9c" +
	"V\xf7\xcd\xa2\xbb\xbc\xa3\xde\x1a\xccH\xc7Q\x18\xd7\xc6" +
	"\xb6\xe8\xc9\x89\xc9\x83\xc7u\x1c\xd5\xa6Z\x97\x93\xb0:" +
	"\x1a\xc6\xed\xf1$F\xfe\xa0\xab\x1f\xd5c\xc9\x95\xd6\xd5" +
	"\xdaD8u*\xd2^\xd5\xd7-\xdd\xec\xfd\x0e\xf0\xc9" +
	"\xc0*\xd9@7Z\xe67H9\x0c!m\xe7\xc6_" +
	"\xb3\x06}\xb2\xcb$\x0a\xa6\xb8\xdd\xac\x8e\x86\x03\x06)" +
	"\xb0J\x16`\x11\x90\xee~ \xe8+1\xe8\x17,g" +
	"ct!\xe8\xe2\xbf\x1e\x9d\xbb\xea\x1d\xde^&\x17t" +
	"\xa272\xc9@|]\xce,\xba\xb2\xd53\x9b<D" +
	"l\xc8,\x8b_\x06R\xae[\xcc\xdc\x99Q\x96b\x9f" +
	"I1o\x0e\xf3\x0eK\xef\x00\x84\xdc\xe7\xb0h\x0d\xf3" +
	"\xfa\xc9\xca^\x08\xe9:\x03\xe6\xfa\x06\x9d\xb8\xdd4)" +
	"\xff\x09\x00\x00\xff\xff\xef|\xe4\x91"

func init() {
	schemas.Register(schema_db8274f9144abc7e,
		0x80ac741ec7fb8f65,
		0x92b20ad1a58ca0ca,
		0xaead580f97fddabc,
		0xd093963b95a4e107,
		0xdffe94ae546cdee3,
		0xe74bb2d0190cf89c,
		0xea3e50f7663f7bdf,
		0xf29f97dd675a9431)
}
