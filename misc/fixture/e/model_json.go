// generated by jwg -type Sample -output misc/fixture/e/model_json.go misc/fixture/e; DO NOT EDIT

package e

import (
	"encoding/json"
)

// for Sample
type SampleJson struct {
	Str *string `json:"str,omitempty"`
	Foo *Foo    `json:"foo,omitempty"`
}

type SampleJsonList []*SampleJson

type SamplePropertyEncoder func(src *Sample, dest *SampleJson) error

type SamplePropertyDecoder func(src *SampleJson, dest *Sample) error

type SamplePropertyInfo struct {
	name    string
	Encoder SamplePropertyEncoder
	Decoder SamplePropertyDecoder
}

type SampleJsonBuilder struct {
	_properties map[string]*SamplePropertyInfo
	Str         *SamplePropertyInfo
	Foo         *SamplePropertyInfo
}

func NewSampleJsonBuilder() *SampleJsonBuilder {
	return &SampleJsonBuilder{
		_properties: map[string]*SamplePropertyInfo{},
		Str: &SamplePropertyInfo{
			name: "Str",
			Encoder: func(src *Sample, dest *SampleJson) error {
				if src == nil {
					return nil
				}
				dest.Str = src.Str
				return nil
			},
			Decoder: func(src *SampleJson, dest *Sample) error {
				if src == nil {
					return nil
				}
				dest.Str = src.Str
				return nil
			},
		},
		Foo: &SamplePropertyInfo{
			name: "Foo",
			Encoder: func(src *Sample, dest *SampleJson) error {
				if src == nil {
					return nil
				}
				dest.Foo = src.Foo
				return nil
			},
			Decoder: func(src *SampleJson, dest *Sample) error {
				if src == nil {
					return nil
				}
				dest.Foo = src.Foo
				return nil
			},
		},
	}
}

func (b *SampleJsonBuilder) AddAll() *SampleJsonBuilder {
	b._properties["Str"] = b.Str
	b._properties["Foo"] = b.Foo
	return b
}

func (b *SampleJsonBuilder) Add(info *SamplePropertyInfo) *SampleJsonBuilder {
	b._properties[info.name] = info
	return b
}

func (b *SampleJsonBuilder) Remove(info *SamplePropertyInfo) *SampleJsonBuilder {
	delete(b._properties, info.name)
	return b
}

func (b *SampleJsonBuilder) Convert(orig *Sample) (*SampleJson, error) {
	if orig == nil {
		return nil, nil
	}
	ret := &SampleJson{}

	for _, info := range b._properties {
		if err := info.Encoder(orig, ret); err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (b *SampleJsonBuilder) ConvertList(orig []*Sample) (SampleJsonList, error) {
	if orig == nil {
		return nil, nil
	}

	list := make(SampleJsonList, len(orig))
	for idx, or := range orig {
		json, err := b.Convert(or)
		if err != nil {
			return nil, err
		}
		list[idx] = json
	}

	return list, nil
}

func (orig *SampleJson) Convert() (*Sample, error) {
	ret := &Sample{}

	b := NewSampleJsonBuilder().AddAll()
	for _, info := range b._properties {
		if err := info.Decoder(orig, ret); err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (jsonList SampleJsonList) Convert() ([]*Sample, error) {
	orig := ([]*SampleJson)(jsonList)

	list := make([]*Sample, len(orig))
	for idx, or := range orig {
		obj, err := or.Convert()
		if err != nil {
			return nil, err
		}
		list[idx] = obj
	}

	return list, nil
}

func (b *SampleJsonBuilder) Marshal(orig *Sample) ([]byte, error) {
	ret, err := b.Convert(orig)
	if err != nil {
		return nil, err
	}
	return json.Marshal(ret)
}
