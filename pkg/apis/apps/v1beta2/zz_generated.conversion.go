// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1beta2

import (
	v1beta2 "k8s.io/api/apps/v1beta2"
	core_v1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api "k8s.io/kubernetes/pkg/api"
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	apps "k8s.io/kubernetes/pkg/apis/apps"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1beta2_RollingUpdateStatefulSetStrategy_To_apps_RollingUpdateStatefulSetStrategy,
		Convert_apps_RollingUpdateStatefulSetStrategy_To_v1beta2_RollingUpdateStatefulSetStrategy,
		Convert_v1beta2_StatefulSet_To_apps_StatefulSet,
		Convert_apps_StatefulSet_To_v1beta2_StatefulSet,
		Convert_v1beta2_StatefulSetList_To_apps_StatefulSetList,
		Convert_apps_StatefulSetList_To_v1beta2_StatefulSetList,
		Convert_v1beta2_StatefulSetSpec_To_apps_StatefulSetSpec,
		Convert_apps_StatefulSetSpec_To_v1beta2_StatefulSetSpec,
		Convert_v1beta2_StatefulSetStatus_To_apps_StatefulSetStatus,
		Convert_apps_StatefulSetStatus_To_v1beta2_StatefulSetStatus,
		Convert_v1beta2_StatefulSetUpdateStrategy_To_apps_StatefulSetUpdateStrategy,
		Convert_apps_StatefulSetUpdateStrategy_To_v1beta2_StatefulSetUpdateStrategy,
	)
}

func autoConvert_v1beta2_RollingUpdateStatefulSetStrategy_To_apps_RollingUpdateStatefulSetStrategy(in *v1beta2.RollingUpdateStatefulSetStrategy, out *apps.RollingUpdateStatefulSetStrategy, s conversion.Scope) error {
	if err := v1.Convert_Pointer_int32_To_int32(&in.Partition, &out.Partition, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta2_RollingUpdateStatefulSetStrategy_To_apps_RollingUpdateStatefulSetStrategy is an autogenerated conversion function.
func Convert_v1beta2_RollingUpdateStatefulSetStrategy_To_apps_RollingUpdateStatefulSetStrategy(in *v1beta2.RollingUpdateStatefulSetStrategy, out *apps.RollingUpdateStatefulSetStrategy, s conversion.Scope) error {
	return autoConvert_v1beta2_RollingUpdateStatefulSetStrategy_To_apps_RollingUpdateStatefulSetStrategy(in, out, s)
}

func autoConvert_apps_RollingUpdateStatefulSetStrategy_To_v1beta2_RollingUpdateStatefulSetStrategy(in *apps.RollingUpdateStatefulSetStrategy, out *v1beta2.RollingUpdateStatefulSetStrategy, s conversion.Scope) error {
	if err := v1.Convert_int32_To_Pointer_int32(&in.Partition, &out.Partition, s); err != nil {
		return err
	}
	return nil
}

// Convert_apps_RollingUpdateStatefulSetStrategy_To_v1beta2_RollingUpdateStatefulSetStrategy is an autogenerated conversion function.
func Convert_apps_RollingUpdateStatefulSetStrategy_To_v1beta2_RollingUpdateStatefulSetStrategy(in *apps.RollingUpdateStatefulSetStrategy, out *v1beta2.RollingUpdateStatefulSetStrategy, s conversion.Scope) error {
	return autoConvert_apps_RollingUpdateStatefulSetStrategy_To_v1beta2_RollingUpdateStatefulSetStrategy(in, out, s)
}

func autoConvert_v1beta2_StatefulSet_To_apps_StatefulSet(in *v1beta2.StatefulSet, out *apps.StatefulSet, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta2_StatefulSetSpec_To_apps_StatefulSetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta2_StatefulSetStatus_To_apps_StatefulSetStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta2_StatefulSet_To_apps_StatefulSet is an autogenerated conversion function.
func Convert_v1beta2_StatefulSet_To_apps_StatefulSet(in *v1beta2.StatefulSet, out *apps.StatefulSet, s conversion.Scope) error {
	return autoConvert_v1beta2_StatefulSet_To_apps_StatefulSet(in, out, s)
}

func autoConvert_apps_StatefulSet_To_v1beta2_StatefulSet(in *apps.StatefulSet, out *v1beta2.StatefulSet, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_apps_StatefulSetSpec_To_v1beta2_StatefulSetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_apps_StatefulSetStatus_To_v1beta2_StatefulSetStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_apps_StatefulSet_To_v1beta2_StatefulSet is an autogenerated conversion function.
func Convert_apps_StatefulSet_To_v1beta2_StatefulSet(in *apps.StatefulSet, out *v1beta2.StatefulSet, s conversion.Scope) error {
	return autoConvert_apps_StatefulSet_To_v1beta2_StatefulSet(in, out, s)
}

func autoConvert_v1beta2_StatefulSetList_To_apps_StatefulSetList(in *v1beta2.StatefulSetList, out *apps.StatefulSetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]apps.StatefulSet, len(*in))
		for i := range *in {
			if err := Convert_v1beta2_StatefulSet_To_apps_StatefulSet(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta2_StatefulSetList_To_apps_StatefulSetList is an autogenerated conversion function.
func Convert_v1beta2_StatefulSetList_To_apps_StatefulSetList(in *v1beta2.StatefulSetList, out *apps.StatefulSetList, s conversion.Scope) error {
	return autoConvert_v1beta2_StatefulSetList_To_apps_StatefulSetList(in, out, s)
}

func autoConvert_apps_StatefulSetList_To_v1beta2_StatefulSetList(in *apps.StatefulSetList, out *v1beta2.StatefulSetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1beta2.StatefulSet, len(*in))
		for i := range *in {
			if err := Convert_apps_StatefulSet_To_v1beta2_StatefulSet(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = make([]v1beta2.StatefulSet, 0)
	}
	return nil
}

// Convert_apps_StatefulSetList_To_v1beta2_StatefulSetList is an autogenerated conversion function.
func Convert_apps_StatefulSetList_To_v1beta2_StatefulSetList(in *apps.StatefulSetList, out *v1beta2.StatefulSetList, s conversion.Scope) error {
	return autoConvert_apps_StatefulSetList_To_v1beta2_StatefulSetList(in, out, s)
}

func autoConvert_v1beta2_StatefulSetSpec_To_apps_StatefulSetSpec(in *v1beta2.StatefulSetSpec, out *apps.StatefulSetSpec, s conversion.Scope) error {
	if err := v1.Convert_Pointer_int32_To_int32(&in.Replicas, &out.Replicas, s); err != nil {
		return err
	}
	out.Selector = (*v1.LabelSelector)(unsafe.Pointer(in.Selector))
	if err := api_v1.Convert_v1_PodTemplateSpec_To_api_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	out.VolumeClaimTemplates = *(*[]api.PersistentVolumeClaim)(unsafe.Pointer(&in.VolumeClaimTemplates))
	out.ServiceName = in.ServiceName
	out.PodManagementPolicy = apps.PodManagementPolicyType(in.PodManagementPolicy)
	if err := Convert_v1beta2_StatefulSetUpdateStrategy_To_apps_StatefulSetUpdateStrategy(&in.UpdateStrategy, &out.UpdateStrategy, s); err != nil {
		return err
	}
	out.RevisionHistoryLimit = (*int32)(unsafe.Pointer(in.RevisionHistoryLimit))
	return nil
}

func autoConvert_apps_StatefulSetSpec_To_v1beta2_StatefulSetSpec(in *apps.StatefulSetSpec, out *v1beta2.StatefulSetSpec, s conversion.Scope) error {
	if err := v1.Convert_int32_To_Pointer_int32(&in.Replicas, &out.Replicas, s); err != nil {
		return err
	}
	out.Selector = (*v1.LabelSelector)(unsafe.Pointer(in.Selector))
	if err := api_v1.Convert_api_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	out.VolumeClaimTemplates = *(*[]core_v1.PersistentVolumeClaim)(unsafe.Pointer(&in.VolumeClaimTemplates))
	out.ServiceName = in.ServiceName
	out.PodManagementPolicy = v1beta2.PodManagementPolicyType(in.PodManagementPolicy)
	if err := Convert_apps_StatefulSetUpdateStrategy_To_v1beta2_StatefulSetUpdateStrategy(&in.UpdateStrategy, &out.UpdateStrategy, s); err != nil {
		return err
	}
	out.RevisionHistoryLimit = (*int32)(unsafe.Pointer(in.RevisionHistoryLimit))
	return nil
}

func autoConvert_v1beta2_StatefulSetStatus_To_apps_StatefulSetStatus(in *v1beta2.StatefulSetStatus, out *apps.StatefulSetStatus, s conversion.Scope) error {
	out.ObservedGeneration = (*int64)(unsafe.Pointer(in.ObservedGeneration))
	out.Replicas = in.Replicas
	out.ReadyReplicas = in.ReadyReplicas
	out.CurrentReplicas = in.CurrentReplicas
	out.UpdatedReplicas = in.UpdatedReplicas
	out.CurrentRevision = in.CurrentRevision
	out.UpdateRevision = in.UpdateRevision
	return nil
}

// Convert_v1beta2_StatefulSetStatus_To_apps_StatefulSetStatus is an autogenerated conversion function.
func Convert_v1beta2_StatefulSetStatus_To_apps_StatefulSetStatus(in *v1beta2.StatefulSetStatus, out *apps.StatefulSetStatus, s conversion.Scope) error {
	return autoConvert_v1beta2_StatefulSetStatus_To_apps_StatefulSetStatus(in, out, s)
}

func autoConvert_apps_StatefulSetStatus_To_v1beta2_StatefulSetStatus(in *apps.StatefulSetStatus, out *v1beta2.StatefulSetStatus, s conversion.Scope) error {
	out.ObservedGeneration = (*int64)(unsafe.Pointer(in.ObservedGeneration))
	out.Replicas = in.Replicas
	out.ReadyReplicas = in.ReadyReplicas
	out.CurrentReplicas = in.CurrentReplicas
	out.UpdatedReplicas = in.UpdatedReplicas
	out.CurrentRevision = in.CurrentRevision
	out.UpdateRevision = in.UpdateRevision
	return nil
}

// Convert_apps_StatefulSetStatus_To_v1beta2_StatefulSetStatus is an autogenerated conversion function.
func Convert_apps_StatefulSetStatus_To_v1beta2_StatefulSetStatus(in *apps.StatefulSetStatus, out *v1beta2.StatefulSetStatus, s conversion.Scope) error {
	return autoConvert_apps_StatefulSetStatus_To_v1beta2_StatefulSetStatus(in, out, s)
}

func autoConvert_v1beta2_StatefulSetUpdateStrategy_To_apps_StatefulSetUpdateStrategy(in *v1beta2.StatefulSetUpdateStrategy, out *apps.StatefulSetUpdateStrategy, s conversion.Scope) error {
	out.Type = apps.StatefulSetUpdateStrategyType(in.Type)
	if in.RollingUpdate != nil {
		in, out := &in.RollingUpdate, &out.RollingUpdate
		*out = new(apps.RollingUpdateStatefulSetStrategy)
		if err := Convert_v1beta2_RollingUpdateStatefulSetStrategy_To_apps_RollingUpdateStatefulSetStrategy(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.RollingUpdate = nil
	}
	return nil
}

func autoConvert_apps_StatefulSetUpdateStrategy_To_v1beta2_StatefulSetUpdateStrategy(in *apps.StatefulSetUpdateStrategy, out *v1beta2.StatefulSetUpdateStrategy, s conversion.Scope) error {
	out.Type = v1beta2.StatefulSetUpdateStrategyType(in.Type)
	if in.RollingUpdate != nil {
		in, out := &in.RollingUpdate, &out.RollingUpdate
		*out = new(v1beta2.RollingUpdateStatefulSetStrategy)
		if err := Convert_apps_RollingUpdateStatefulSetStrategy_To_v1beta2_RollingUpdateStatefulSetStrategy(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.RollingUpdate = nil
	}
	return nil
}
