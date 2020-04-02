// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package version

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Info contains versioning information. how we'll want to distribute that information.
type Info struct {
	BuildDate    *string `pulumi:"buildDate"`
	Compiler     *string `pulumi:"compiler"`
	GitCommit    *string `pulumi:"gitCommit"`
	GitTreeState *string `pulumi:"gitTreeState"`
	GitVersion   *string `pulumi:"gitVersion"`
	GoVersion    *string `pulumi:"goVersion"`
	Major        *string `pulumi:"major"`
	Minor        *string `pulumi:"minor"`
	Platform     *string `pulumi:"platform"`
}

type InfoInput interface {
	pulumi.Input

	ToInfoOutput() InfoOutput
	ToInfoOutputWithContext(context.Context) InfoOutput
}

// Info contains versioning information. how we'll want to distribute that information.
type InfoArgs struct {
	BuildDate    pulumi.StringPtrInput `pulumi:"buildDate"`
	Compiler     pulumi.StringPtrInput `pulumi:"compiler"`
	GitCommit    pulumi.StringPtrInput `pulumi:"gitCommit"`
	GitTreeState pulumi.StringPtrInput `pulumi:"gitTreeState"`
	GitVersion   pulumi.StringPtrInput `pulumi:"gitVersion"`
	GoVersion    pulumi.StringPtrInput `pulumi:"goVersion"`
	Major        pulumi.StringPtrInput `pulumi:"major"`
	Minor        pulumi.StringPtrInput `pulumi:"minor"`
	Platform     pulumi.StringPtrInput `pulumi:"platform"`
}

func (InfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Info)(nil)).Elem()
}

func (i InfoArgs) ToInfoOutput() InfoOutput {
	return i.ToInfoOutputWithContext(context.Background())
}

func (i InfoArgs) ToInfoOutputWithContext(ctx context.Context) InfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InfoOutput)
}

// Info contains versioning information. how we'll want to distribute that information.
type InfoOutput struct{ *pulumi.OutputState }

func (InfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Info)(nil)).Elem()
}

func (o InfoOutput) ToInfoOutput() InfoOutput {
	return o
}

func (o InfoOutput) ToInfoOutputWithContext(ctx context.Context) InfoOutput {
	return o
}

func (o InfoOutput) BuildDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Info) *string { return v.BuildDate }).(pulumi.StringPtrOutput)
}

func (o InfoOutput) Compiler() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Info) *string { return v.Compiler }).(pulumi.StringPtrOutput)
}

func (o InfoOutput) GitCommit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Info) *string { return v.GitCommit }).(pulumi.StringPtrOutput)
}

func (o InfoOutput) GitTreeState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Info) *string { return v.GitTreeState }).(pulumi.StringPtrOutput)
}

func (o InfoOutput) GitVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Info) *string { return v.GitVersion }).(pulumi.StringPtrOutput)
}

func (o InfoOutput) GoVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Info) *string { return v.GoVersion }).(pulumi.StringPtrOutput)
}

func (o InfoOutput) Major() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Info) *string { return v.Major }).(pulumi.StringPtrOutput)
}

func (o InfoOutput) Minor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Info) *string { return v.Minor }).(pulumi.StringPtrOutput)
}

func (o InfoOutput) Platform() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Info) *string { return v.Platform }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(InfoOutput{})
}