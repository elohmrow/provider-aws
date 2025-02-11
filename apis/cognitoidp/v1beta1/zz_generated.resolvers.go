/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta11 "github.com/upbound/provider-aws/apis/acm/v1beta1"
	v1beta1 "github.com/upbound/provider-aws/apis/iam/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this IdentityProvider.
func (mg *IdentityProvider) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.UserPoolID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.UserPoolIDRef,
		Selector:     mg.Spec.ForProvider.UserPoolIDSelector,
		To: reference.To{
			List:    &UserPoolList{},
			Managed: &UserPool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.UserPoolID")
	}
	mg.Spec.ForProvider.UserPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserPoolIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ResourceServer.
func (mg *ResourceServer) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.UserPoolID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.UserPoolIDRef,
		Selector:     mg.Spec.ForProvider.UserPoolIDSelector,
		To: reference.To{
			List:    &UserPoolList{},
			Managed: &UserPool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.UserPoolID")
	}
	mg.Spec.ForProvider.UserPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserPoolIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this User.
func (mg *User) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.UserPoolID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.UserPoolIDRef,
		Selector:     mg.Spec.ForProvider.UserPoolIDSelector,
		To: reference.To{
			List:    &UserPoolList{},
			Managed: &UserPool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.UserPoolID")
	}
	mg.Spec.ForProvider.UserPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserPoolIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this UserPool.
func (mg *UserPool) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.SMSConfiguration); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SMSConfiguration[i3].SnsCallerArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.SMSConfiguration[i3].SnsCallerArnRef,
			Selector:     mg.Spec.ForProvider.SMSConfiguration[i3].SnsCallerArnSelector,
			To: reference.To{
				List:    &v1beta1.RoleList{},
				Managed: &v1beta1.Role{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.SMSConfiguration[i3].SnsCallerArn")
		}
		mg.Spec.ForProvider.SMSConfiguration[i3].SnsCallerArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.SMSConfiguration[i3].SnsCallerArnRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this UserPoolClient.
func (mg *UserPoolClient) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.AnalyticsConfiguration); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AnalyticsConfiguration[i3].RoleArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.AnalyticsConfiguration[i3].RoleArnRef,
			Selector:     mg.Spec.ForProvider.AnalyticsConfiguration[i3].RoleArnSelector,
			To: reference.To{
				List:    &v1beta1.RoleList{},
				Managed: &v1beta1.Role{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.AnalyticsConfiguration[i3].RoleArn")
		}
		mg.Spec.ForProvider.AnalyticsConfiguration[i3].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.AnalyticsConfiguration[i3].RoleArnRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.UserPoolID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.UserPoolIDRef,
		Selector:     mg.Spec.ForProvider.UserPoolIDSelector,
		To: reference.To{
			List:    &UserPoolList{},
			Managed: &UserPool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.UserPoolID")
	}
	mg.Spec.ForProvider.UserPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserPoolIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this UserPoolDomain.
func (mg *UserPoolDomain) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CertificateArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.CertificateArnRef,
		Selector:     mg.Spec.ForProvider.CertificateArnSelector,
		To: reference.To{
			List:    &v1beta11.CertificateList{},
			Managed: &v1beta11.Certificate{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CertificateArn")
	}
	mg.Spec.ForProvider.CertificateArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CertificateArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.UserPoolID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.UserPoolIDRef,
		Selector:     mg.Spec.ForProvider.UserPoolIDSelector,
		To: reference.To{
			List:    &UserPoolList{},
			Managed: &UserPool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.UserPoolID")
	}
	mg.Spec.ForProvider.UserPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserPoolIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this UserPoolUICustomization.
func (mg *UserPoolUICustomization) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClientID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClientIDRef,
		Selector:     mg.Spec.ForProvider.ClientIDSelector,
		To: reference.To{
			List:    &UserPoolClientList{},
			Managed: &UserPoolClient{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClientID")
	}
	mg.Spec.ForProvider.ClientID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClientIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.UserPoolID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.UserPoolIDRef,
		Selector:     mg.Spec.ForProvider.UserPoolIDSelector,
		To: reference.To{
			List:    &UserPoolList{},
			Managed: &UserPool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.UserPoolID")
	}
	mg.Spec.ForProvider.UserPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserPoolIDRef = rsp.ResolvedReference

	return nil
}
