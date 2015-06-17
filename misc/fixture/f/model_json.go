// generated by jwg -type SampleF -output misc/fixture/f/model_json.go misc/fixture/f; DO NOT EDIT

package f

import (
	"encoding/json"
	"github.com/favclip/jwg/misc/fixture/a"
	bravo "github.com/favclip/jwg/misc/fixture/b"
)

// for SampleF
type SampleFJson struct {
	A *a.Sample     `json:"a,omitempty"`
	B *bravo.Sample `json:"b,omitempty"`
}

type SampleFJsonList []*SampleFJson

type SampleFPropertyEncoder func(src *SampleF, dest *SampleFJson) error

type SampleFPropertyDecoder func(src *SampleFJson, dest *SampleF) error

type SampleFPropertyInfo struct {
	name    string
	Encoder SampleFPropertyEncoder
	Decoder SampleFPropertyDecoder
}

type SampleFJsonBuilder struct {
	_properties map[string]*SampleFPropertyInfo
	A           *SampleFPropertyInfo
	B           *SampleFPropertyInfo
}

func NewSampleFJsonBuilder() *SampleFJsonBuilder {
	return &SampleFJsonBuilder{
		_properties: map[string]*SampleFPropertyInfo{},
		A: &SampleFPropertyInfo{
			name: "A",
			Encoder: func(src *SampleF, dest *SampleFJson) error {
				if src == nil {
					return nil
				}
				dest.A = src.A
				return nil
			},
			Decoder: func(src *SampleFJson, dest *SampleF) error {
				if src == nil {
					return nil
				}
				dest.A = src.A
				return nil
			},
		},
		B: &SampleFPropertyInfo{
			name: "B",
			Encoder: func(src *SampleF, dest *SampleFJson) error {
				if src == nil {
					return nil
				}
				dest.B = src.B
				return nil
			},
			Decoder: func(src *SampleFJson, dest *SampleF) error {
				if src == nil {
					return nil
				}
				dest.B = src.B
				return nil
			},
		},
	}
}

func (b *SampleFJsonBuilder) AddAll() *SampleFJsonBuilder {
	b._properties["A"] = b.A
	b._properties["B"] = b.B
	return b
}

func (b *SampleFJsonBuilder) Add(info *SampleFPropertyInfo) *SampleFJsonBuilder {
	b._properties[info.name] = info
	return b
}

func (b *SampleFJsonBuilder) Remove(info *SampleFPropertyInfo) *SampleFJsonBuilder {
	delete(b._properties, info.name)
	return b
}

func (b *SampleFJsonBuilder) Convert(orig *SampleF) (*SampleFJson, error) {
	if orig == nil {
		return nil, nil
	}
	ret := &SampleFJson{}

	for _, info := range b._properties {
		if err := info.Encoder(orig, ret); err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (b *SampleFJsonBuilder) ConvertList(orig []*SampleF) (SampleFJsonList, error) {
	if orig == nil {
		return nil, nil
	}

	list := make(SampleFJsonList, len(orig))
	for idx, or := range orig {
		json, err := b.Convert(or)
		if err != nil {
			return nil, err
		}
		list[idx] = json
	}

	return list, nil
}

func (orig *SampleFJson) Convert() (*SampleF, error) {
	ret := &SampleF{}

	b := NewSampleFJsonBuilder().AddAll()
	for _, info := range b._properties {
		if err := info.Decoder(orig, ret); err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (jsonList SampleFJsonList) Convert() ([]*SampleF, error) {
	orig := ([]*SampleFJson)(jsonList)

	list := make([]*SampleF, len(orig))
	for idx, or := range orig {
		obj, err := or.Convert()
		if err != nil {
			return nil, err
		}
		list[idx] = obj
	}

	return list, nil
}

func (b *SampleFJsonBuilder) Marshal(orig *SampleF) ([]byte, error) {
	ret, err := b.Convert(orig)
	if err != nil {
		return nil, err
	}
	return json.Marshal(ret)
}