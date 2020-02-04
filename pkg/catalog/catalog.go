// Package catalog takes a sysl module with attributes defined (catalogFields) and serves a webserver listing the applications and endpoints
// It also uses GRPCUI and Redoc in order to generate an interactive page to interact with all the endpoints
// GRPC currently uses server reflection TODO: Support gpcui directly from swagger files
package catalog

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/fullstorydev/grpcui"
	"github.com/fullstorydev/grpcurl"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"net/http"

	"github.com/anz-bank/sysl/pkg/exporter"
	"github.com/anz-bank/sysl/pkg/sysl"
	"github.com/fullstorydev/grpcui/standalone"
)

// Server to set context of catalog
type Server struct {
	Fs       afero.Fs
	Log      *logrus.Logger
	Modules  []*sysl.Module
	Fields   []string
	BasePath string `long:"base-path" description:"the base path to serve the spec and UI at"`
	Path     string
	Resource string
	Flavor   string `short:"F" long:"flavor" description:"the flavor of docs, can be swagger or redoc" default:"redoc" choice:"redoc,swagger"` //nolint: lll
	DocURL   string `long:"doc-url" description:"override the url which takes a url query param to render the doc ui"`
	NoOpen   bool   `long:"no-open" description:"when present won't open the the browser to show the url"`
	NoUI     bool   `long:"no-ui" description:"when present, only the swagger spec will be served"`
	Flatten  bool   `long:"flatten" description:"when present, flatten the swagger spec before serving it"`
	// Port     string `long:"port" short:"p" description:"the port to serve this site" env:"PORT"`
	Host string `long:"host" description:"the interface to serve this site, defaults to 0.0.0.0" env:"HOST"`
}

// WebService is the type which will be rendered on the home page of the html/json as a row
type WebService struct {
	App           *sysl.Application `json:"-"`
	Fields        []string          `json:"-"`
	Attrs         map[string]string
	ServiceName   string
	SwaggerUILink string
}

// ListHandlers registers handlers for both the homepage, if t is json the header will be set as json content type
func (c *Server) ListHandlers(contents []byte, t string, pattern string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if t == "json" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
		}
		_, err := w.Write(contents)
		if err != nil {
			c.Log.Errorf(err.Error())
			panic(err)
		}
	})
}

// Serve Runs the command and runs a webserver on catalogURL of a list of endpoints in the sysl file
func (c *Server) Serve() error {
	var err error
	services, err := c.BuildCatalog()
	if err != nil {
		c.Log.Errorf(err.Error())
		return err
	}
	json, err := json.Marshal(services)
	if err != nil {
		c.Log.Errorf(err.Error())
		return err
	}
	html, err := renderHTML(services)
	if err != nil {
		c.Log.Errorf(err.Error())
		return err
	}
	c.ListHandlers(json, "json", "/json")
	c.ListHandlers(html, "html", "/")
	addr := c.Host
	fmt.Println(addr)
	err = http.ListenAndServe(addr, nil)
	c.Log.Errorf(err.Error())
	return err
}

// BuildCatalog unpacks a sysl modules and registers a http handler for the grpcui or swagger ui
func (c *Server) BuildCatalog() ([]WebService, error) {
	var ser []WebService
	var h http.Handler
	var err error
	var serviceCount int
	for _, m := range c.Modules {
		for i, a := range m.GetApps() {
			var attr = make(map[string]string, 10)
			atts := a.GetAttrs()
			for key, value := range atts {
				attr[key] = value.GetS()
			}
			_, ok := attr["type"]
			if ok {
				serviceCount++
				serviceName := strings.Join(a.Name.GetPart(), "")
				var re = regexp.MustCompile(`(?m)\w*`)
				serviceName = strings.Join(re.FindAllString(serviceName, -1), "")
				serviceName = strings.ToLower(serviceName) + strconv.Itoa(serviceCount)
				// serviceName = strings.ReplaceAll(serviceName, " ", "")
				fmt.Println(serviceName)

				c.Log.Infof("eofn: %s", i)
				newService := WebService{
					App:           a,
					Fields:        c.Fields,
					Attrs:         attr,
					ServiceName:   serviceName,
					SwaggerUILink: "/" + serviceName,
				}
				c.Log.Infof("Adding %s service: %s from %s", newService.Attrs["type"], newService.ServiceName, newService.Attrs["deploy.prod.url"])
				switch strings.ToUpper(newService.Attrs["type"]) {
				case "GRPC":
					h, err = c.GrpcUIHandler(newService)
				case "REST":
					c.Log.Info("Hello")
					c.Log.Info(c)
					c.Log.Info(newService)
					h, err = c.SwaggerUIHandler(newService)
					c.Log.Info("elvns")

				}
				if err != nil {
					c.Log.Errorf(err.Error())
				}
				h = http.StripPrefix(newService.SwaggerUILink, h)
				http.Handle(newService.SwaggerUILink+"/", h)
				c.Log.Errorf("newService.SwaggerUILink" + newService.SwaggerUILink + "/")
				ser = append(ser, newService)
				c.Log.Infof("Added %s service: %s from %s", newService.Attrs["type"], newService.ServiceName, newService.Attrs["deploy.prod.url"])
			}
		}
	}
	return ser, nil
}

// GrpcUIHandler creates and returns a http handler for a grpcui server
func (c *Server) GrpcUIHandler(service WebService) (http.Handler, error) {
	ctx := context.Background()

	// Plaintext doesn't work for GRPC Server Reflection calls to Fabric
	var creds credentials.TransportCredentials
	var opts []grpc.DialOption
	var err error
	network := "tcp"
	creds, err = grpcurl.ClientTransportCredentials(false, "", "", "")
	if err != nil {
		c.Log.Errorf(err.Error())
		return nil, err
	}
	cc, err := grpcurl.BlockingDial(ctx, network, service.Attrs["deploy.prod.url"], creds, opts...)
	if err != nil {
		c.Log.Errorf(err.Error())
		return nil, err
	}
	methods, err := grpcui.AllMethodsViaReflection(ctx, cc)
	if err != nil {
		c.Log.Errorf(err.Error())
		return nil, err
	}
	file, err := c.Fs.Create("index.html")
	if err != nil {
		c.Log.Errorf(err.Error())
		return nil, err
	}
	_, err = file.Write(grpcui.WebFormContents("invoke", "metadata", methods))
	if err != nil {
		c.Log.Errorf(err.Error())
		return nil, err
	}
	file, err = c.Fs.Create("grpc-web-form.js")
	if err != nil {
		c.Log.Errorf(err.Error())
		return nil, err
	}
	_, err = file.Write(grpcui.WebFormScript())
	if err != nil {
		c.Log.Errorf(err.Error())
		return nil, err
	}
	file, err = c.Fs.Create("grpc-web-form.css")
	if err != nil {
		c.Log.Errorf(err.Error())
		return nil, err
	}
	_, err = file.Write(grpcui.WebFormSampleCSS())
	if err != nil {
		c.Log.Errorf(err.Error())
		return nil, err
	}
	h, err := standalone.HandlerViaReflection(ctx, cc, service.SwaggerUILink)
	if err != nil {
		c.Log.Errorf(err.Error())
		return nil, err
	}
	return h, nil
}

// SwaggerUIHandler creates and returns a http handler for a SwaggerUI server
func (c *Server) SwaggerUIHandler(service WebService) (http.Handler, error) {
	// basePath := "/"
	// swag := Server{BasePath: basePath, Path: "/", Resource: service.SwaggerUILink}
	c.Resource = service.SwaggerUILink
	swaggerExporter := exporter.MakeSwaggerExporter(service.App, c.Log)
	err := swaggerExporter.GenerateSwagger()
	if err != nil {
		c.Log.Errorf(err.Error())
		return nil, err
	}
	output, err := swaggerExporter.SerializeOutput("json")
	if err != nil {
		c.Log.Errorf(err.Error())
		return nil, err
	}
	return c.SwaggerUI(output), nil
}
