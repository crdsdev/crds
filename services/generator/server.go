package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/gobuffalo/flect"
	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
	apiext "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	pb "github.com/crdsdev/crds/services/generator/genproto"
)

const listenPort = "8000"

type echoService struct{}

func main() {
	port := listenPort
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}

	svc := new(echoService)

	srv := grpc.NewServer()
	log.Infof("starting echo service on tcp: %q", lis.Addr().String())
	pb.RegisterEchoServiceServer(srv, svc)
	err = srv.Serve(lis)
	log.Fatal(err)
}

func (es *echoService) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	groupKind := schema.GroupKind{
		Group: "test.crds.dev",
		Kind:  "Runner",
	}
	defaultPlural := flect.Pluralize(strings.ToLower(groupKind.Kind))
	crd := apiext.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiext.SchemeGroupVersion.String(),
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: defaultPlural + "." + groupKind.Group,
		},
		Spec: apiext.CustomResourceDefinitionSpec{
			Group: groupKind.Group,
			Names: apiext.CustomResourceDefinitionNames{
				Kind:     groupKind.Kind,
				ListKind: groupKind.Kind + "List",
				Plural:   defaultPlural,
				Singular: strings.ToLower(groupKind.Kind),
			},
			Scope: apiext.NamespaceScoped,
			Versions: []apiext.CustomResourceDefinitionVersion{
				{
					Name:   "v1",
					Served: true,
					Schema: &apiext.CustomResourceValidation{
						OpenAPIV3Schema: &apiext.JSONSchemaProps{
							Properties: map[string]apiext.JSONSchemaProps{},
						},
					},
				},
			},
		},
	}

	log.Info(crd)
	return &pb.EchoResponse{
		Output: req.GetInput(),
	}, nil
}
