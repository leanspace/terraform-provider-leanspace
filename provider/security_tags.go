package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// SecurityTagsDefinition opts a parent into complete security-tag collection ownership.
type SecurityTagsDefinition struct {
	ResourceIdentifier string
	TargetIDAttribute  string
}

type securityTagsResource struct {
	basePath   string
	definition SecurityTagsDefinition
}

// Its state ID is the parent's UUID, saved with SetId after a successful PUT.
// On first creation, the state ID is still empty, so use the parent UUID
func (resource securityTagsResource) resolveTargetID(resourceData *schema.ResourceData) string {
	targetID := resourceData.Id()
	if targetID == "" {
		targetID = resourceData.Get(resource.definition.TargetIDAttribute).(string)
	}
	return targetID
}

func (resource securityTagsResource) sendSecurityTagRequest(
	requestContext context.Context,
	client *Client,
	targetID string,
	httpMethod string,
	securityTags []string,
) ([]byte, error, int) {
	var requestBody []byte
	if httpMethod == http.MethodPut {
		var err error
		requestBody, err = json.Marshal(struct {
			SecurityTags []string `json:"securityTags"`
		}{securityTags})
		if err != nil {
			return nil, err, 0
		}
	}
	request, err := http.NewRequestWithContext(requestContext, httpMethod, fmt.Sprintf("%s/%s/%s/security-tags", client.HostURL, resource.basePath, targetID), bytes.NewReader(requestBody))
	if err != nil {
		return nil, err, 0
	}
	if httpMethod == http.MethodPut {
		request.Header.Set("Content-Type", "application/json")
	}
	return client.DoRequest(request, &client.Token)
}

func (resource securityTagsResource) readSecurityTags(requestContext context.Context, resourceData *schema.ResourceData, providerMeta any) diag.Diagnostics {
	targetIDAttribute := resource.definition.TargetIDAttribute
	client := providerMeta.(*Client)
	targetID := resource.resolveTargetID(resourceData)
	responseBody, err, statusCode := resource.sendSecurityTagRequest(requestContext, client, targetID, http.MethodGet, nil)
	if statusCode == http.StatusNotFound {
		resourceData.SetId("")
		return nil
	}
	if err != nil {
		return diag.FromErr(err)
	}
	var securityTags []string
	if err = json.Unmarshal(responseBody, &securityTags); err != nil {
		return diag.FromErr(err)
	}
	if securityTags == nil {
		return diag.Errorf("expected a JSON string array for security tags")
	}
	if err = resourceData.Set(targetIDAttribute, resourceData.Id()); err != nil {
		return diag.FromErr(err)
	}
	return diag.FromErr(resourceData.Set("security_tags", securityTags))
}

func (resource securityTagsResource) writeSecurityTags(requestContext context.Context, resourceData *schema.ResourceData, providerMeta any) diag.Diagnostics {
	targetIDAttribute := resource.definition.TargetIDAttribute
	client := providerMeta.(*Client)
	targetID := resource.resolveTargetID(resourceData)
	tagValues := resourceData.Get("security_tags").(*schema.Set).List()
	securityTags := make([]string, len(tagValues))
	for index, value := range tagValues {
		securityTags[index] = value.(string)
	}
	_, err, _ := resource.sendSecurityTagRequest(requestContext, client, targetID, http.MethodPut, securityTags)
	if err != nil {
		return diag.FromErr(err)
	}
	// Identify the association by its parent's UUID, not by a tag ID.
	// Save it before GET so the association remains tracked if that read fails.
	resourceData.SetId(resourceData.Get(targetIDAttribute).(string))
	return resource.readSecurityTags(requestContext, resourceData, providerMeta)
}

func (resource securityTagsResource) toResource() *schema.Resource {
	targetIDAttribute := resource.definition.TargetIDAttribute
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			targetIDAttribute: {Type: schema.TypeString, Required: true, ForceNew: true, ValidateFunc: validation.IsUUID},
			"security_tags":   {Type: schema.TypeSet, Required: true, Elem: &schema.Schema{Type: schema.TypeString}},
		},
		CreateContext: resource.writeSecurityTags,
		UpdateContext: resource.writeSecurityTags,
		ReadContext:   resource.readSecurityTags,
		DeleteContext: resource.clearSecurityTags,
		Importer: &schema.ResourceImporter{
			StateContext: resource.importSecurityTags,
		},
	}
}

func (resource securityTagsResource) clearSecurityTags(requestContext context.Context, resourceData *schema.ResourceData, providerMeta any) diag.Diagnostics {
	client := providerMeta.(*Client)
	targetID := resource.resolveTargetID(resourceData)
	_, err, statusCode := resource.sendSecurityTagRequest(requestContext, client, targetID, http.MethodPut, []string{})
	if err != nil && statusCode != http.StatusNotFound {
		return diag.FromErr(err)
	}
	resourceData.SetId("")
	return nil
}

func (resource securityTagsResource) importSecurityTags(requestContext context.Context, resourceData *schema.ResourceData, providerMeta any) ([]*schema.ResourceData, error) {
	targetIDAttribute := resource.definition.TargetIDAttribute
	_, errs := validation.IsUUID(resourceData.Id(), targetIDAttribute)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if err := resourceData.Set(targetIDAttribute, resourceData.Id()); err != nil {
		return nil, err
	}
	if diags := resource.readSecurityTags(requestContext, resourceData, providerMeta); diags.HasError() {
		return nil, fmt.Errorf("reading imported security tags: %v", diags)
	}
	return []*schema.ResourceData{resourceData}, nil
}
