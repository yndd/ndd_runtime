package model

import (
	"errors"
	"fmt"
	"reflect"
	"sort"

	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"

	"github.com/openconfig/gnmi/proto/gnmi"
	pb "github.com/openconfig/gnmi/proto/gnmi"
)

// JSONUnmarshaler is the signature of the Unmarshal() function in the GoStruct code generated by openconfig ygot library.
type JSONUnmarshaler func([]byte, ygot.ValidatedGoStruct, ...ytypes.UnmarshalOpt) error

// GoStructEnumData is the data type to maintain GoStruct enum type.
type GoStructEnumData map[string]map[int64]ygot.EnumDefinition

// Model contains the model data and GoStruct information for the device to config.
type Model struct {
	ModelData       []*pb.ModelData
	StructRootType  reflect.Type
	SchemaTreeRoot  *yang.Entry
	JsonUnmarshaler JSONUnmarshaler
	EnumData        GoStructEnumData
}

func (m *Model) NewRootValue() interface{} {
	return reflect.New(m.StructRootType.Elem()).Interface()
}

// NewConfigStruct creates a ValidatedGoStruct of this model from jsonConfig. If jsonConfig is nil, creates an empty GoStruct.
func (m *Model) NewConfigStruct(jsonConfig []byte, validate bool) (ygot.ValidatedGoStruct, error) {
	rootStruct, ok := m.NewRootValue().(ygot.ValidatedGoStruct)
	if !ok {
		return nil, errors.New("root node is not a ygot.ValidatedGoStruct")
	}
	if jsonConfig != nil {
		if err := m.JsonUnmarshaler(jsonConfig, rootStruct); err != nil {
			return nil, err
		}
		if validate {
			if err := rootStruct.Validate(); err != nil {
				return nil, err
			}
		}
	}
	return rootStruct, nil
}

// SupportedModels returns a list of supported models.
func (m *Model) SupportedModels() []string {
	mDesc := make([]string, len(m.ModelData))
	for i, m := range m.ModelData {
		mDesc[i] = fmt.Sprintf("%s %s", m.Name, m.Version)
	}
	sort.Strings(mDesc)
	return mDesc
}

func GetModelData(cap []*gnmi.ModelData) []*gnmi.ModelData {
	modelData := make([]*gnmi.ModelData, 0)
	for _, c := range cap {
		modelData = append(modelData, &gnmi.ModelData{
			Name:         c.GetName(),
			Organization: c.GetOrganization(),
			Version:      c.GetVersion(),
		})
	}
	return modelData
}