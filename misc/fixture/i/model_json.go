// generated by jwg -output misc/fixture/i/model_json.go misc/fixture/i; DO NOT EDIT

package i

import (
	"encoding/json"
)

// for Person
type PersonJson struct {
	Name     string `json:"name,omitempty"`
	Age      int    `json:"age,omitempty"`
	Password string `json:"password,omitempty"`
}

type PersonJsonList []*PersonJson

type PersonPropertyEncoder func(src *Person, dest *PersonJson) error

type PersonPropertyDecoder func(src *PersonJson, dest *Person) error

type PersonPropertyInfo struct {
	name    string
	Encoder PersonPropertyEncoder
	Decoder PersonPropertyDecoder
}

type PersonJsonBuilder struct {
	_properties map[string]*PersonPropertyInfo
	Name        *PersonPropertyInfo
	Age         *PersonPropertyInfo
	Password    *PersonPropertyInfo
}

func NewPersonJsonBuilder() *PersonJsonBuilder {
	return &PersonJsonBuilder{
		_properties: map[string]*PersonPropertyInfo{},
		Name: &PersonPropertyInfo{
			name: "Name",
			Encoder: func(src *Person, dest *PersonJson) error {
				if src == nil {
					return nil
				}
				dest.Name = src.Name
				return nil
			},
			Decoder: func(src *PersonJson, dest *Person) error {
				if src == nil {
					return nil
				}
				dest.Name = src.Name
				return nil
			},
		},
		Age: &PersonPropertyInfo{
			name: "Age",
			Encoder: func(src *Person, dest *PersonJson) error {
				if src == nil {
					return nil
				}
				dest.Age = src.Age
				return nil
			},
			Decoder: func(src *PersonJson, dest *Person) error {
				if src == nil {
					return nil
				}
				dest.Age = src.Age
				return nil
			},
		},
		Password: &PersonPropertyInfo{
			name: "Password",
			Encoder: func(src *Person, dest *PersonJson) error {
				if src == nil {
					return nil
				}
				dest.Password = src.Password
				return nil
			},
			Decoder: func(src *PersonJson, dest *Person) error {
				if src == nil {
					return nil
				}
				dest.Password = src.Password
				return nil
			},
		},
	}
}

func (b *PersonJsonBuilder) AddAll() *PersonJsonBuilder {
	b._properties["Name"] = b.Name
	b._properties["Age"] = b.Age
	b._properties["Password"] = b.Password
	return b
}

func (b *PersonJsonBuilder) Add(info *PersonPropertyInfo) *PersonJsonBuilder {
	b._properties[info.name] = info
	return b
}

func (b *PersonJsonBuilder) Remove(info *PersonPropertyInfo) *PersonJsonBuilder {
	delete(b._properties, info.name)
	return b
}

func (b *PersonJsonBuilder) Convert(orig *Person) (*PersonJson, error) {
	if orig == nil {
		return nil, nil
	}
	ret := &PersonJson{}

	for _, info := range b._properties {
		if err := info.Encoder(orig, ret); err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (b *PersonJsonBuilder) ConvertList(orig []*Person) (PersonJsonList, error) {
	if orig == nil {
		return nil, nil
	}

	list := make(PersonJsonList, len(orig))
	for idx, or := range orig {
		json, err := b.Convert(or)
		if err != nil {
			return nil, err
		}
		list[idx] = json
	}

	return list, nil
}

func (orig *PersonJson) Convert() (*Person, error) {
	ret := &Person{}

	b := NewPersonJsonBuilder().AddAll()
	for _, info := range b._properties {
		if err := info.Decoder(orig, ret); err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (jsonList PersonJsonList) Convert() ([]*Person, error) {
	orig := ([]*PersonJson)(jsonList)

	list := make([]*Person, len(orig))
	for idx, or := range orig {
		obj, err := or.Convert()
		if err != nil {
			return nil, err
		}
		list[idx] = obj
	}

	return list, nil
}

func (b *PersonJsonBuilder) Marshal(orig *Person) ([]byte, error) {
	ret, err := b.Convert(orig)
	if err != nil {
		return nil, err
	}
	return json.Marshal(ret)
}

// for People
type PeopleJson struct {
	ShowPrivateInfo bool          `json:"showPrivateInfo,omitempty"`
	List            []*PersonJson `json:"list,omitempty"`
}

type PeopleJsonList []*PeopleJson

type PeoplePropertyEncoder func(src *People, dest *PeopleJson) error

type PeoplePropertyDecoder func(src *PeopleJson, dest *People) error

type PeoplePropertyInfo struct {
	name    string
	Encoder PeoplePropertyEncoder
	Decoder PeoplePropertyDecoder
}

type PeopleJsonBuilder struct {
	_properties     map[string]*PeoplePropertyInfo
	ShowPrivateInfo *PeoplePropertyInfo
	List            *PeoplePropertyInfo
}

func NewPeopleJsonBuilder() *PeopleJsonBuilder {
	return &PeopleJsonBuilder{
		_properties: map[string]*PeoplePropertyInfo{},
		ShowPrivateInfo: &PeoplePropertyInfo{
			name: "ShowPrivateInfo",
			Encoder: func(src *People, dest *PeopleJson) error {
				if src == nil {
					return nil
				}
				dest.ShowPrivateInfo = src.ShowPrivateInfo
				return nil
			},
			Decoder: func(src *PeopleJson, dest *People) error {
				if src == nil {
					return nil
				}
				dest.ShowPrivateInfo = src.ShowPrivateInfo
				return nil
			},
		},
		List: &PeoplePropertyInfo{
			name: "List",
			Encoder: func(src *People, dest *PeopleJson) error {
				if src == nil {
					return nil
				}
				list, err := NewPersonJsonBuilder().AddAll().ConvertList(src.List)
				if err != nil {
					return err
				}
				dest.List = ([]*PersonJson)(list)
				return nil
			},
			Decoder: func(src *PeopleJson, dest *People) error {
				if src == nil {
					return nil
				}
				list := make([]*Person, len(src.List))
				for idx, obj := range src.List {
					if obj == nil {
						continue
					}
					d, err := obj.Convert()
					if err != nil {
						return err
					}
					list[idx] = d
				}
				dest.List = list
				return nil
			},
		},
	}
}

func (b *PeopleJsonBuilder) AddAll() *PeopleJsonBuilder {
	b._properties["ShowPrivateInfo"] = b.ShowPrivateInfo
	b._properties["List"] = b.List
	return b
}

func (b *PeopleJsonBuilder) Add(info *PeoplePropertyInfo) *PeopleJsonBuilder {
	b._properties[info.name] = info
	return b
}

func (b *PeopleJsonBuilder) Remove(info *PeoplePropertyInfo) *PeopleJsonBuilder {
	delete(b._properties, info.name)
	return b
}

func (b *PeopleJsonBuilder) Convert(orig *People) (*PeopleJson, error) {
	if orig == nil {
		return nil, nil
	}
	ret := &PeopleJson{}

	for _, info := range b._properties {
		if err := info.Encoder(orig, ret); err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (b *PeopleJsonBuilder) ConvertList(orig []*People) (PeopleJsonList, error) {
	if orig == nil {
		return nil, nil
	}

	list := make(PeopleJsonList, len(orig))
	for idx, or := range orig {
		json, err := b.Convert(or)
		if err != nil {
			return nil, err
		}
		list[idx] = json
	}

	return list, nil
}

func (orig *PeopleJson) Convert() (*People, error) {
	ret := &People{}

	b := NewPeopleJsonBuilder().AddAll()
	for _, info := range b._properties {
		if err := info.Decoder(orig, ret); err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (jsonList PeopleJsonList) Convert() ([]*People, error) {
	orig := ([]*PeopleJson)(jsonList)

	list := make([]*People, len(orig))
	for idx, or := range orig {
		obj, err := or.Convert()
		if err != nil {
			return nil, err
		}
		list[idx] = obj
	}

	return list, nil
}

func (b *PeopleJsonBuilder) Marshal(orig *People) ([]byte, error) {
	ret, err := b.Convert(orig)
	if err != nil {
		return nil, err
	}
	return json.Marshal(ret)
}