/*
Copyright 2021 NDD.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package parser

import (
	"bufio"
	"context"
	"fmt"
	"io"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/apimachinery/pkg/util/yaml"
)

// AnnotatedReadCloser is a wrapper around io.ReadCloser that allows
// implementations to supply additional information about data that is read.
type AnnotatedReadCloser interface {
	io.ReadCloser
	Annotate() interface{}
}

// ObjectCreaterTyper know how to create and determine the type of objects.
type ObjectCreaterTyper interface {
	runtime.ObjectCreater
	runtime.ObjectTyper
}

// Package is the set of metadata and objects in a package.
type Package struct {
	meta    []runtime.Object
	objects []runtime.Object
}

// NewPackage creates a new Package.
func NewPackage() *Package {
	return &Package{}
}

// GetMeta gets metadata from the package.
func (p *Package) GetMeta() []runtime.Object {
	return p.meta
}

// GetObjects gets objects from the package.
func (p *Package) GetObjects() []runtime.Object {
	return p.objects
}

// PackageParser is a Parser implementation for parsing packages.
type PackageParser struct {
	metaScheme ObjectCreaterTyper
	objScheme  ObjectCreaterTyper
}

// New returns a new PackageParser.
func New(meta, obj ObjectCreaterTyper) *PackageParser {
	return &PackageParser{
		metaScheme: meta,
		objScheme:  obj,
	}
}

// Parser is a package parser.
type Parser interface {
	Parse(context.Context, io.ReadCloser) (*Package, error)
}

// Parse is the underlying logic for parsing packages. It first attempts to
// decode objects recognized by the meta scheme, then attempts to decode objects
// recognized by the object scheme. Objects not recognized by either scheme
// return an error rather than being skipped.
func (p *PackageParser) Parse(ctx context.Context, reader io.ReadCloser) (*Package, error) { //nolint:gocyclo
	pkg := NewPackage()
	if reader == nil {
		return pkg, nil
	}
	defer func() { _ = reader.Close() }()
	yr := yaml.NewYAMLReader(bufio.NewReader(reader))
	dm := json.NewSerializerWithOptions(json.DefaultMetaFactory, p.metaScheme, p.metaScheme, json.SerializerOptions{Yaml: true})
	do := json.NewSerializerWithOptions(json.DefaultMetaFactory, p.objScheme, p.objScheme, json.SerializerOptions{Yaml: true})
	for {
		bytes, err := yr.Read()
		if err != nil && err != io.EOF {
			return pkg, err
		}
		if err == io.EOF {
			break
		}
		if len(bytes) == 0 {
			continue
		}
		m, _, err := dm.Decode(bytes, nil, nil)
		if err != nil {
			o, _, err := do.Decode(bytes, nil, nil)
			if err != nil {
				if anno, ok := reader.(AnnotatedReadCloser); ok {
					return pkg, errors.Wrap(err, fmt.Sprintf("%+v", anno.Annotate()))
				}
				return pkg, err
			}
			pkg.objects = append(pkg.objects, o)
			continue
		}
		pkg.meta = append(pkg.meta, m)
	}
	return pkg, nil
}

// BackendOption modifies the parser backend. Backends may accept options at
// creation time, but must accept them at initialization.
type BackendOption func(Backend)

// Backend provides a source for a parser.
type Backend interface {
	Init(context.Context, ...BackendOption) (io.ReadCloser, error)
}

// NopBackend is a parser backend with empty source.
type NopBackend struct{}

// NewNopBackend returns a new NopBackend.
func NewNopBackend(...BackendOption) *NopBackend {
	return &NopBackend{}
}

// Init initializes a NopBackend.
func (p *NopBackend) Init(ctx context.Context, bo ...BackendOption) (io.ReadCloser, error) {
	return nil, nil
}

// FsBackend is a parser backend that uses a filestystem as source.
type FsBackend struct {
	fs    afero.Fs
	dir   string
	skips []FilterFn
}

// NewFsBackend returns an FsBackend.
func NewFsBackend(fs afero.Fs, bo ...BackendOption) *FsBackend {
	f := &FsBackend{
		fs: fs,
	}
	for _, o := range bo {
		o(f)
	}
	return f
}

// Init initializes an FsBackend.
func (p *FsBackend) Init(ctx context.Context, bo ...BackendOption) (io.ReadCloser, error) {
	for _, o := range bo {
		o(p)
	}
	return NewFsReadCloser(p.fs, p.dir, p.skips...)
}

// FsDir sets the directory of an FsBackend.
func FsDir(dir string) BackendOption {
	return func(p Backend) {
		f, ok := p.(*FsBackend)
		if !ok {
			return
		}
		f.dir = dir
	}
}

// FsFilters adds FilterFns to an FsBackend.
func FsFilters(skips ...FilterFn) BackendOption {
	return func(p Backend) {
		f, ok := p.(*FsBackend)
		if !ok {
			return
		}
		f.skips = skips
	}
}
