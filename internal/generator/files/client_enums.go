package files

import (
	"bytes"
	"fmt"
	"github.com/eden-framework/eden-framework/internal/generator/importer"
	"github.com/eden-framework/eden-framework/internal/generator/scanner"
	"github.com/eden-framework/enumeration"
	str "github.com/eden-framework/strings"
	"github.com/go-courier/oas"
	"github.com/sirupsen/logrus"
	"io"
	"sort"
	"strings"
)

var serviceEnumMap = map[string]map[string]enumeration.Enum{}

func RegisterEnum(serviceName string, tpe string, options ...enumeration.EnumOption) {
	serviceName = strings.ToLower(str.ToUpperCamelCase(serviceName))
	if serviceEnumMap[serviceName] == nil {
		serviceEnumMap[serviceName] = map[string]enumeration.Enum{}
	}
	serviceEnumMap[serviceName][tpe] = options
}

type ClientEnumsFile struct {
	PackageName string
	Importer    *importer.PackageImporter
	a           *oas.OpenAPI
	serviceName string
}

func NewClientEnumsFile(outputPath, serviceName string, a *oas.OpenAPI) *ClientEnumsFile {
	pkgName := str.ToLowerSnakeCase("client-" + serviceName)

	f := &ClientEnumsFile{
		PackageName: pkgName,
		Importer:    importer.NewPackageImporter(""),
		a:           a,
		serviceName: serviceName,
	}

	return f
}

func (f *ClientEnumsFile) WritePackage(w io.Writer) (err error) {
	_, err = io.WriteString(w, fmt.Sprintf("package %s\n\n", f.PackageName))
	return
}

func (f *ClientEnumsFile) WriteImports(w io.Writer) (err error) {
	_, err = io.WriteString(w, f.Importer.String())
	return
}

func (f *ClientEnumsFile) WriteAll() string {
	names := make([]string, 0)
	for name := range serviceEnumMap[f.serviceName] {
		names = append(names, name)
	}
	sort.Strings(names)

	buf := bytes.NewBuffer([]byte{})

	for _, name := range names {
		if name == "Bool" {
			continue
		}
		enum := serviceEnumMap[f.serviceName][name]
		buf.Write(ToEnumDefines(name, enum))
	}

	for _, name := range names {
		if name == "Bool" {
			continue
		}
		enum := serviceEnumMap[f.serviceName][name]
		generatedEnum := NewEnum("", f.PackageName, name, scanner.Enum(enum), false)
		buf.WriteString(generatedEnum.WriteAll())
		f.Importer.Merge(generatedEnum.Importer)
	}

	return buf.String()
}

func (f *ClientEnumsFile) String() string {
	buf := bytes.NewBuffer([]byte{})

	content := f.WriteAll()

	err := f.WritePackage(buf)
	if err != nil {
		logrus.Panic(err)
	}

	err = f.WriteImports(buf)
	if err != nil {
		logrus.Panic(err)
	}

	_, err = io.WriteString(buf, content)
	if err != nil {
		logrus.Panic(err)
	}

	return buf.String()
}

func ToEnumDefines(name string, enum enumeration.Enum) []byte {
	buf := &bytes.Buffer{}

	buf.WriteString(`
// api:enum
type ` + name + ` uint

const (
`)

	buf.WriteString(str.ToUpperSnakeCase(name) + `_UNKNOWN ` + name + ` = iota
`)

	sort.Slice(enum, func(i, j int) bool {
		return enum[i].Val < enum[j].Val
	})

	index := 1
	for _, item := range enum {
		v := item.Val
		if v > index {
			buf.WriteString(`)

const (
`)
			buf.WriteString(str.ToUpperSnakeCase(name) + `__` + item.Value.(string) + fmt.Sprintf(" %s = iota + %d", name, v) + `// ` + item.Label + `
`)
			index = v + 1
			continue
		}
		index++
		buf.WriteString(str.ToUpperSnakeCase(name) + `__` + item.Value.(string) + `// ` + item.Label + `
`)
	}

	buf.WriteString(`)
`)

	return buf.Bytes()
}
