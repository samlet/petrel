// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/samlet/petrel/alfin/modules/asset/ent/asset"
	"github.com/samlet/petrel/alfin/modules/asset/ent/schema"
	"github.com/samlet/petrel/alfin/modules/asset/ent/workloadpkg"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	assetMixin := schema.Asset{}.Mixin()
	assetMixinFields0 := assetMixin[0].Fields()
	_ = assetMixinFields0
	assetFields := schema.Asset{}.Fields()
	_ = assetFields
	// assetDescCreateTime is the schema descriptor for create_time field.
	assetDescCreateTime := assetMixinFields0[0].Descriptor()
	// asset.DefaultCreateTime holds the default value on creation for the create_time field.
	asset.DefaultCreateTime = assetDescCreateTime.Default.(func() time.Time)
	// assetDescUpdateTime is the schema descriptor for update_time field.
	assetDescUpdateTime := assetMixinFields0[1].Descriptor()
	// asset.DefaultUpdateTime holds the default value on creation for the update_time field.
	asset.DefaultUpdateTime = assetDescUpdateTime.Default.(func() time.Time)
	// asset.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	asset.UpdateDefaultUpdateTime = assetDescUpdateTime.UpdateDefault.(func() time.Time)
	workloadpkgFields := schema.WorkloadPkg{}.Fields()
	_ = workloadpkgFields
	// workloadpkgDescName is the schema descriptor for name field.
	workloadpkgDescName := workloadpkgFields[0].Descriptor()
	// workloadpkg.DefaultName holds the default value on creation for the name field.
	workloadpkg.DefaultName = workloadpkgDescName.Default.(string)
}
