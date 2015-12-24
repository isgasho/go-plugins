/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

// DO NOT EDIT. THIS FILE IS AUTO-GENERATED BY $KUBEROOT/hack/update-generated-conversions.sh

package v1alpha1

import (
	reflect "reflect"

	api "github.com/micro/go-plugins/Godeps/_workspace/src/k8s.io/kubernetes/pkg/api"
	metrics "github.com/micro/go-plugins/Godeps/_workspace/src/k8s.io/kubernetes/pkg/apis/metrics"
	conversion "github.com/micro/go-plugins/Godeps/_workspace/src/k8s.io/kubernetes/pkg/conversion"
)

func autoconvert_metrics_RawNode_To_v1alpha1_RawNode(in *metrics.RawNode, out *RawNode, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*metrics.RawNode))(in)
	}
	if err := s.Convert(&in.TypeMeta, &out.TypeMeta, 0); err != nil {
		return err
	}
	return nil
}

func convert_metrics_RawNode_To_v1alpha1_RawNode(in *metrics.RawNode, out *RawNode, s conversion.Scope) error {
	return autoconvert_metrics_RawNode_To_v1alpha1_RawNode(in, out, s)
}

func autoconvert_metrics_RawPod_To_v1alpha1_RawPod(in *metrics.RawPod, out *RawPod, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*metrics.RawPod))(in)
	}
	if err := s.Convert(&in.TypeMeta, &out.TypeMeta, 0); err != nil {
		return err
	}
	return nil
}

func convert_metrics_RawPod_To_v1alpha1_RawPod(in *metrics.RawPod, out *RawPod, s conversion.Scope) error {
	return autoconvert_metrics_RawPod_To_v1alpha1_RawPod(in, out, s)
}

func autoconvert_v1alpha1_RawNode_To_metrics_RawNode(in *RawNode, out *metrics.RawNode, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*RawNode))(in)
	}
	if err := s.Convert(&in.TypeMeta, &out.TypeMeta, 0); err != nil {
		return err
	}
	return nil
}

func convert_v1alpha1_RawNode_To_metrics_RawNode(in *RawNode, out *metrics.RawNode, s conversion.Scope) error {
	return autoconvert_v1alpha1_RawNode_To_metrics_RawNode(in, out, s)
}

func autoconvert_v1alpha1_RawPod_To_metrics_RawPod(in *RawPod, out *metrics.RawPod, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*RawPod))(in)
	}
	if err := s.Convert(&in.TypeMeta, &out.TypeMeta, 0); err != nil {
		return err
	}
	return nil
}

func convert_v1alpha1_RawPod_To_metrics_RawPod(in *RawPod, out *metrics.RawPod, s conversion.Scope) error {
	return autoconvert_v1alpha1_RawPod_To_metrics_RawPod(in, out, s)
}

func init() {
	err := api.Scheme.AddGeneratedConversionFuncs(
		autoconvert_metrics_RawNode_To_v1alpha1_RawNode,
		autoconvert_metrics_RawPod_To_v1alpha1_RawPod,
		autoconvert_v1alpha1_RawNode_To_metrics_RawNode,
		autoconvert_v1alpha1_RawPod_To_metrics_RawPod,
	)
	if err != nil {
		// If one of the conversion functions is malformed, detect it immediately.
		panic(err)
	}
}
