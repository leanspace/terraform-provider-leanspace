package connections

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type LeafSpaceConnectionTF struct {
	general_objects.AuditModelTF
	Name                types.String `tfsdk:"name"`
	DomainUrl           types.String `tfsdk:"domain_url"`
	AuthenticationToken types.String `tfsdk:"authentication_token"`
	Password            types.String `tfsdk:"password"`
	Username            types.String `tfsdk:"username"`
	Status              types.String `tfsdk:"status"`
}

func (s *LeafSpaceConnection) ToTF() any {
	return &LeafSpaceConnectionTF{
		AuditModelTF:        general_objects.AuditModelToTF(&s.AuditModel),
		Name:                helper.TFStringValue(s.Name),
		DomainUrl:           helper.TFStringValue(s.DomainUrl),
		AuthenticationToken: helper.TFStringPtrValue(s.AuthenticationToken),
		Password:            helper.TFStringValue(s.Password),
		Username:            helper.TFStringValue(s.Username),
		Status:              helper.TFStringValue(s.Status),
	}
}

func (tf *LeafSpaceConnectionTF) ToAPI() any {
	return &LeafSpaceConnection{
		AuditModel:          general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:                helper.FromTFString(tf.Name),
		DomainUrl:           helper.FromTFString(tf.DomainUrl),
		AuthenticationToken: helper.FromTFStringPtr(tf.AuthenticationToken),
		Password:            helper.FromTFString(tf.Password),
		Username:            helper.FromTFString(tf.Username),
		Status:              helper.FromTFString(tf.Status),
	}
}

// LeafSpaceConnectionDSTF is the data-source-specific TF model for unique reads.
// It matches the FilterSchema fields (id, created_at, created_by, domain_url, etc.).
type LeafSpaceConnectionDSTF struct {
	ID             types.String `tfsdk:"id"`
	CreatedAt      types.String `tfsdk:"created_at"`
	CreatedBy      types.String `tfsdk:"created_by"`
	DomainURL      types.String `tfsdk:"domain_url"`
	LastModifiedAt types.String `tfsdk:"last_modified_at"`
	LastModifiedBy types.String `tfsdk:"last_modified_by"`
	Status         types.String `tfsdk:"status"`
}

func (s *LeafSpaceConnection) ToDSTF() any {
	return &LeafSpaceConnectionDSTF{
		ID:             helper.TFStringValue(s.ID),
		CreatedAt:      helper.TFStringValue(s.CreatedAt),
		CreatedBy:      helper.TFStringValue(s.CreatedBy),
		DomainURL:      helper.TFStringValue(s.DomainUrl),
		LastModifiedAt: helper.TFStringValue(s.LastModifiedAt),
		LastModifiedBy: helper.TFStringValue(s.LastModifiedBy),
		Status:         helper.TFStringValue(s.Status),
	}
}
