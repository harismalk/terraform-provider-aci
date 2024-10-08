package convert_funcs

import (
	"context"
	"encoding/json"

	//"github.com/CiscoDevNet/terraform-provider-aci/v2/convert_funcs"
	//"github.com/CiscoDevNet/terraform-provider-aci/v2/convert_funcs"

	"github.com/CiscoDevNet/terraform-provider-aci/v2/internal/provider"
	"github.com/ciscoecosystem/aci-go-client/v2/container"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func CreateQosDppPol(attributes map[string]interface{}) map[string]interface{} {
	ctx := context.Background()
	var diags diag.Diagnostics
	data := &provider.QosDppPolResourceModel{}
	if v, ok := attributes["parent_dn"].(string); ok && v != "" {
		data.ParentDn = types.StringValue(v)
	}
	if v, ok := attributes["admin_state"].(string); ok && v != "" {
		data.AdminSt = types.StringValue(v)
	}
	if v, ok := attributes["annotation"].(string); ok && v != "" {
		data.Annotation = types.StringValue(v)
	}
	if v, ok := attributes["excessive_burst"].(string); ok && v != "" {
		data.Be = types.StringValue(v)
	}
	if v, ok := attributes["excessive_burst_unit"].(string); ok && v != "" {
		data.BeUnit = types.StringValue(v)
	}
	if v, ok := attributes["burst"].(string); ok && v != "" {
		data.Burst = types.StringValue(v)
	}
	if v, ok := attributes["burst_unit"].(string); ok && v != "" {
		data.BurstUnit = types.StringValue(v)
	}
	if v, ok := attributes["conform_action"].(string); ok && v != "" {
		data.ConformAction = types.StringValue(v)
	}
	if v, ok := attributes["conform_mark_cos"].(string); ok && v != "" {
		data.ConformMarkCos = types.StringValue(v)
	}
	if v, ok := attributes["conform_mark_dscp"].(string); ok && v != "" {
		data.ConformMarkDscp = types.StringValue(v)
	}
	if v, ok := attributes["description"].(string); ok && v != "" {
		data.Descr = types.StringValue(v)
	}
	if v, ok := attributes["exceed_action"].(string); ok && v != "" {
		data.ExceedAction = types.StringValue(v)
	}
	if v, ok := attributes["exceed_mark_cos"].(string); ok && v != "" {
		data.ExceedMarkCos = types.StringValue(v)
	}
	if v, ok := attributes["exceed_mark_dscp"].(string); ok && v != "" {
		data.ExceedMarkDscp = types.StringValue(v)
	}
	if v, ok := attributes["mode"].(string); ok && v != "" {
		data.Mode = types.StringValue(v)
	}
	if v, ok := attributes["name"].(string); ok && v != "" {
		data.Name = types.StringValue(v)
	}
	if v, ok := attributes["name_alias"].(string); ok && v != "" {
		data.NameAlias = types.StringValue(v)
	}
	if v, ok := attributes["owner_key"].(string); ok && v != "" {
		data.OwnerKey = types.StringValue(v)
	}
	if v, ok := attributes["owner_tag"].(string); ok && v != "" {
		data.OwnerTag = types.StringValue(v)
	}
	if v, ok := attributes["peak_rate"].(string); ok && v != "" {
		data.Pir = types.StringValue(v)
	}
	if v, ok := attributes["peak_rate_unit"].(string); ok && v != "" {
		data.PirUnit = types.StringValue(v)
	}
	if v, ok := attributes["rate"].(string); ok && v != "" {
		data.Rate = types.StringValue(v)
	}
	if v, ok := attributes["rate_unit"].(string); ok && v != "" {
		data.RateUnit = types.StringValue(v)
	}
	if v, ok := attributes["sharing_mode"].(string); ok && v != "" {
		data.SharingMode = types.StringValue(v)
	}
	if v, ok := attributes["type"].(string); ok && v != "" {
		data.Type = types.StringValue(v)
	}
	if v, ok := attributes["violate_action"].(string); ok && v != "" {
		data.ViolateAction = types.StringValue(v)
	}
	if v, ok := attributes["violate_mark_cos"].(string); ok && v != "" {
		data.ViolateMarkCos = types.StringValue(v)
	}
	if v, ok := attributes["violate_mark_dscp"].(string); ok && v != "" {
		data.ViolateMarkDscp = types.StringValue(v)
	}
	planTagAnnotation := convertToTagAnnotationQosDppPol(attributes["annotations"])
	planTagTag := convertToTagTagQosDppPol(attributes["tags"])

	newAciQosDppPol := provider.GetQosDppPolCreateJsonPayload(ctx, &diags, true, data, planTagAnnotation, planTagAnnotation, planTagTag, planTagTag)

	jsonPayload := newAciQosDppPol.EncodeJSON(container.EncodeOptIndent("", "  "))

	var customData map[string]interface{}
	json.Unmarshal(jsonPayload, &customData)

	payload := customData

	provider.SetQosDppPolId(ctx, data)
	attrs := payload["qosDppPol"].(map[string]interface{})["attributes"].(map[string]interface{})
	attrs["dn"] = data.Id.ValueString()

	return payload
}
func convertToTagAnnotationQosDppPol(resources interface{}) []provider.TagAnnotationQosDppPolResourceModel {
	var planResources []provider.TagAnnotationQosDppPolResourceModel
	if resources, ok := resources.([]interface{}); ok {
		for _, resource := range resources {
			resourceMap := resource.(map[string]interface{})
			planResources = append(planResources, provider.TagAnnotationQosDppPolResourceModel{
				Key:   types.StringValue(resourceMap["key"].(string)),
				Value: types.StringValue(resourceMap["value"].(string)),
			})
		}
	}
	return planResources
}
func convertToTagTagQosDppPol(resources interface{}) []provider.TagTagQosDppPolResourceModel {
	var planResources []provider.TagTagQosDppPolResourceModel
	if resources, ok := resources.([]interface{}); ok {
		for _, resource := range resources {
			resourceMap := resource.(map[string]interface{})
			planResources = append(planResources, provider.TagTagQosDppPolResourceModel{
				Key:   types.StringValue(resourceMap["key"].(string)),
				Value: types.StringValue(resourceMap["value"].(string)),
			})
		}
	}
	return planResources
}
