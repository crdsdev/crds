/*
Copyright 2019 The CRDS Authors.

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

package crder

import (
	"encoding/json"
	"errors"

	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	servervalidation "k8s.io/apiextensions-apiserver/pkg/apiserver/validation"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/validation"
	"sigs.k8s.io/yaml"
)

const (
	getStoredGVKErr          = "unable to determind stored gvk"
	createSchemaValidatorErr = "could not create schema validator"
	yamlToJSONErr            = "could not convert yaml to json"
	getTypeMetaErr           = "could not get type metadata for crd instance"
	wrongGVKErr              = "crd instance was not of correct group version kind"
	instanceConversionErr    = "could not convert crd instance json to instance"
)

// CRDer generates instances of a CustomResourceDefinition.
type CRDer struct {
	crd    *apiextensions.CustomResourceDefinition
	gvk    *schema.GroupVersionKind
	isBeta bool
}

// NewCRDer returns a new CRDer type.
func NewCRDer(data []byte, isBeta bool) (*CRDer, error) {
	internal := &apiextensions.CustomResourceDefinition{}
	if isBeta {
		if err := convertV1Beta1ToInternal(data, internal); err != nil {
			return nil, err
		}
	} else {
		if err := convertV1ToInternal(data, internal); err != nil {
			return nil, err
		}
	}

	gvk := getStoredGVK(internal)
	if gvk == nil {
		return nil, errors.New(getStoredGVKErr)
	}

	return &CRDer{crd: internal, gvk: gvk, isBeta: isBeta}, nil
}

// Validate returns true if CRD instance is valid.
func (c *CRDer) Validate(data []byte) error {
	sv := getStoredSchema(c.crd.Spec)

	s, _, err := servervalidation.NewSchemaValidator(sv)
	if err != nil {
		return errors.New(createSchemaValidatorErr)
	}

	j, err := yaml.YAMLToJSONStrict(data)
	if err != nil {
		return errors.New(yamlToJSONErr)
	}

	meta := &metav1.TypeMeta{}
	if err := json.Unmarshal(j, meta); err != nil {
		return errors.New(getTypeMetaErr)
	}

	if !isStoredGVK(meta, c.gvk) {
		return errors.New(wrongGVKErr)
	}

	var instance interface{}
	if err := json.Unmarshal(j, &instance); err != nil {
		return errors.New(instanceConversionErr)
	}

	res := servervalidation.ValidateCustomResource(nil, instance, s)
	if len(res) > 0 {
		return errors.New(res.ToAggregate().Error())
	}
	return nil
}

func convertV1ToInternal(data []byte, internal *apiextensions.CustomResourceDefinition) error {
	crd := &v1.CustomResourceDefinition{}
	if err := yaml.Unmarshal(data, crd); err != nil {
		return err
	}
	v1.SetDefaults_CustomResourceDefinition(crd)
	if err := v1.Convert_v1_CustomResourceDefinition_To_apiextensions_CustomResourceDefinition(crd, internal, nil); err != nil {
		return err
	}
	errList := validation.ValidateCustomResourceDefinition(internal, v1.SchemeGroupVersion)
	if len(errList) > 0 {
		return errors.New(errList.ToAggregate().Error())
	}

	return nil
}

func convertV1Beta1ToInternal(data []byte, internal *apiextensions.CustomResourceDefinition) error {
	crd := &v1beta1.CustomResourceDefinition{}
	if err := yaml.Unmarshal(data, crd); err != nil {
		return err
	}
	v1beta1.SetObjectDefaults_CustomResourceDefinition(crd)
	if err := v1beta1.Convert_v1beta1_CustomResourceDefinition_To_apiextensions_CustomResourceDefinition(crd, internal, nil); err != nil {
		return err
	}
	errList := validation.ValidateCustomResourceDefinition(internal, v1beta1.SchemeGroupVersion)
	if len(errList) > 0 {
		return errors.New(errList.ToAggregate().Error())
	}

	return nil
}

func getStoredSchema(spec apiextensions.CustomResourceDefinitionSpec) *apiextensions.CustomResourceValidation {
	if spec.Validation != nil {
		return spec.Validation
	}
	for _, v := range spec.Versions {
		if v.Storage {
			return v.Schema
		}
	}
	return nil
}

func getStoredGVK(crd *apiextensions.CustomResourceDefinition) *schema.GroupVersionKind {
	for _, v := range crd.Spec.Versions {
		if v.Storage {
			return &schema.GroupVersionKind{
				Group:   crd.Spec.Group,
				Version: v.Name,
				Kind:    crd.Spec.Names.Kind,
			}
		}
	}

	return nil
}

func isStoredGVK(meta *metav1.TypeMeta, gvk *schema.GroupVersionKind) bool {
	if meta.GroupVersionKind() == *gvk {
		return true
	}

	return false
}
