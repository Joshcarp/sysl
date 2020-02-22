package importer

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/emicklei/proto"
	"github.com/sirupsen/logrus"
)

type protoLoader struct {
	logger    *logrus.Logger
	types     TypeList
	info      SyslInfo
	endpoints []MethodEndpoints
}

func (p *protoLoader) handleService(s *proto.Service) {
	p.endpoints = append(p.endpoints, MethodEndpoints{
		Method: s.Name,
		Endpoints: []Endpoint{Endpoint{
			Path:        s.Name,
			Description: `strings.Join(s.Comment.Lines, "\n")`,
			Params:      Parameters{},
			Responses:   nil,
		}},
	})
}

func (p *protoLoader) handleMessage(m *proto.Message) {
	p.types.types = append(p.types.types, &StandardType{name: m.Name})

	fmt.Println(m.Name)
}

func LoadDescriptorSet(input string) (*proto.Proto, error) {
	reader := strings.NewReader(input)
	parser := proto.NewParser(reader)
	definition, err := parser.Parse()
	if err != nil {
		return nil, err
	}
	return definition, nil
}

func LoadProtoText(args OutputData, text string, logger *logrus.Logger) (out string, err error) {
	fmt.Println(text)

	definition, err := LoadDescriptorSet(text)
	if err != nil {
		panic(err)
		return "", err
	}

	pLoader := &protoLoader{info: SyslInfo{
		OutputData: OutputData{
			AppName:     definition.Filename,
			Package:     "",
			SwaggerRoot: "",
			Mode:        "",
		},
		Title:       "",
		Description: "",
		OtherFields: nil,
	}}

	proto.Walk(definition,
		proto.WithService(pLoader.handleService),
		proto.WithMessage(pLoader.handleMessage))

	result := &bytes.Buffer{}
	w := newWriter(result, logger)
	if err := w.Write(pLoader.info, pLoader.types, pLoader.endpoints...); err != nil {
		panic(err)
		return "", err
	}
	return result.String(), nil
}

//func importProto(args OutputData,
//	protoc *descriptor.FileDescriptorProto,
//	logger *logrus.Logger, basepath string) (out string, err error) {
//	l := &loader{
//		logger:        logger,
//		externalSpecs: make(map[string]*loader),
//		spec:          protoc,
//		types:         TypeList{},
//		mode:          args.Mode,
//		swaggerRoot:   args.SwaggerRoot,
//	}
//	l.initTypes()
//	endpoints := l.initEndpoints()
//
//	result := &bytes.Buffer{}
//	w := newWriter(result, logger)
//	if err := w.Write(l.initInfo(args, basepath), l.types, endpoints...); err != nil {
//		return "", err
//	}
//	return result.String(), nil
//}
