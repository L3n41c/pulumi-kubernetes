// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha2

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "kubernetes:resource.k8s.io/v1alpha2:PodSchedulingContext":
		r = &PodSchedulingContext{}
	case "kubernetes:resource.k8s.io/v1alpha2:PodSchedulingContextList":
		r = &PodSchedulingContextList{}
	case "kubernetes:resource.k8s.io/v1alpha2:PodSchedulingContextPatch":
		r = &PodSchedulingContextPatch{}
	case "kubernetes:resource.k8s.io/v1alpha2:ResourceClaim":
		r = &ResourceClaim{}
	case "kubernetes:resource.k8s.io/v1alpha2:ResourceClaimList":
		r = &ResourceClaimList{}
	case "kubernetes:resource.k8s.io/v1alpha2:ResourceClaimPatch":
		r = &ResourceClaimPatch{}
	case "kubernetes:resource.k8s.io/v1alpha2:ResourceClaimTemplate":
		r = &ResourceClaimTemplate{}
	case "kubernetes:resource.k8s.io/v1alpha2:ResourceClaimTemplateList":
		r = &ResourceClaimTemplateList{}
	case "kubernetes:resource.k8s.io/v1alpha2:ResourceClaimTemplatePatch":
		r = &ResourceClaimTemplatePatch{}
	case "kubernetes:resource.k8s.io/v1alpha2:ResourceClass":
		r = &ResourceClass{}
	case "kubernetes:resource.k8s.io/v1alpha2:ResourceClassList":
		r = &ResourceClassList{}
	case "kubernetes:resource.k8s.io/v1alpha2:ResourceClassPatch":
		r = &ResourceClassPatch{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := kubernetes.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"kubernetes",
		"resource.k8s.io/v1alpha2",
		&module{version},
	)
}
