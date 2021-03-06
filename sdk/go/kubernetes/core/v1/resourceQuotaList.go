// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ResourceQuotaList is a list of ResourceQuota items.
type ResourceQuotaList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Items is a list of ResourceQuota objects. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
	Items ResourceQuotaTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewResourceQuotaList registers a new resource with the given unique name, arguments, and options.
func NewResourceQuotaList(ctx *pulumi.Context,
	name string, args *ResourceQuotaListArgs, opts ...pulumi.ResourceOption) (*ResourceQuotaList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("ResourceQuotaList")
	var resource ResourceQuotaList
	err := ctx.RegisterResource("kubernetes:core/v1:ResourceQuotaList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceQuotaList gets an existing ResourceQuotaList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceQuotaList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceQuotaListState, opts ...pulumi.ResourceOption) (*ResourceQuotaList, error) {
	var resource ResourceQuotaList
	err := ctx.ReadResource("kubernetes:core/v1:ResourceQuotaList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceQuotaList resources.
type resourceQuotaListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is a list of ResourceQuota objects. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
	Items []ResourceQuotaType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

type ResourceQuotaListState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is a list of ResourceQuota objects. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
	Items ResourceQuotaTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (ResourceQuotaListState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceQuotaListState)(nil)).Elem()
}

type resourceQuotaListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is a list of ResourceQuota objects. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
	Items []ResourceQuotaType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a ResourceQuotaList resource.
type ResourceQuotaListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is a list of ResourceQuota objects. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
	Items ResourceQuotaTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata metav1.ListMetaPtrInput
}

func (ResourceQuotaListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceQuotaListArgs)(nil)).Elem()
}

type ResourceQuotaListInput interface {
	pulumi.Input

	ToResourceQuotaListOutput() ResourceQuotaListOutput
	ToResourceQuotaListOutputWithContext(ctx context.Context) ResourceQuotaListOutput
}

func (ResourceQuotaList) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceQuotaList)(nil)).Elem()
}

func (i ResourceQuotaList) ToResourceQuotaListOutput() ResourceQuotaListOutput {
	return i.ToResourceQuotaListOutputWithContext(context.Background())
}

func (i ResourceQuotaList) ToResourceQuotaListOutputWithContext(ctx context.Context) ResourceQuotaListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceQuotaListOutput)
}

type ResourceQuotaListOutput struct {
	*pulumi.OutputState
}

func (ResourceQuotaListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceQuotaListOutput)(nil)).Elem()
}

func (o ResourceQuotaListOutput) ToResourceQuotaListOutput() ResourceQuotaListOutput {
	return o
}

func (o ResourceQuotaListOutput) ToResourceQuotaListOutputWithContext(ctx context.Context) ResourceQuotaListOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ResourceQuotaListOutput{})
}
