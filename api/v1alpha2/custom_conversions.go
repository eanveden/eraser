package v1alpha2

import (
	unversioned "github.com/eraser-dev/eraser/api/unversioned"
	conversion "k8s.io/apimachinery/pkg/conversion"
)

//nolint:revive
func Convert_v1alpha2_ManagerConfig_To_unversioned_ManagerConfig(in *ManagerConfig, out *unversioned.ManagerConfig, s conversion.Scope) error {
	return autoConvert_v1alpha2_ManagerConfig_To_unversioned_ManagerConfig(in, out, s)
}

//nolint:revive
func manualConvert_v1alpha2_Runtime_To_unversioned_RuntimeSpec(in *Runtime, out *unversioned.RuntimeSpec, _ conversion.Scope) error {
	out.Name = unversioned.Runtime(string(*in))
	out.Address = ""
	return nil
}

//nolint:revive
func Convert_v1alpha2_Runtime_To_unversioned_RuntimeSpec(in *Runtime, out *unversioned.RuntimeSpec, s conversion.Scope) error {
	return manualConvert_v1alpha2_Runtime_To_unversioned_RuntimeSpec(in, out, s)
}

//nolint:revive
func Convert_unversioned_ManagerConfig_To_v1alpha2_ManagerConfig(in *unversioned.ManagerConfig, out *ManagerConfig, s conversion.Scope) error {
	return autoConvert_unversioned_ManagerConfig_To_v1alpha2_ManagerConfig(in, out, s)
}

//nolint:revive
func manualConvert_unversioned_RuntimeSpec_To_v1alpha2_Runtime(in *unversioned.RuntimeSpec, out *Runtime, _ conversion.Scope) error {
	*out = Runtime(in.Name)
	return nil
}

//nolint:revive
func Convert_unversioned_RuntimeSpec_To_v1alpha2_Runtime(in *unversioned.RuntimeSpec, out *Runtime, s conversion.Scope) error {
	return manualConvert_unversioned_RuntimeSpec_To_v1alpha2_Runtime(in, out, s)
}
