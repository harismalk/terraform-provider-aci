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

func CreateFvAEPg(attributes map[string]interface{}) map[string]interface{} {
	ctx := context.Background()
	var diags diag.Diagnostics
	data := &provider.FvAEPgResourceModel{}
	if v, ok := attributes["parent_dn"].(string); ok && v != "" {
		data.ParentDn = types.StringValue(v)
	}
	if v, ok := attributes["annotation"].(string); ok && v != "" {
		data.Annotation = types.StringValue(v)
	}
	if v, ok := attributes["description"].(string); ok && v != "" {
		data.Descr = types.StringValue(v)
	}
	if v, ok := attributes["contract_exception_tag"].(string); ok && v != "" {
		data.ExceptionTag = types.StringValue(v)
	}
	if v, ok := attributes["flood_in_encapsulation"].(string); ok && v != "" {
		data.FloodOnEncap = types.StringValue(v)
	}
	if v, ok := attributes["forwarding_control"].(string); ok && v != "" {
		data.FwdCtrl = types.StringValue(v)
	}
	if v, ok := attributes["has_multicast_source"].(string); ok && v != "" {
		data.HasMcastSource = types.StringValue(v)
	}
	if v, ok := attributes["useg_epg"].(string); ok && v != "" {
		data.IsAttrBasedEPg = types.StringValue(v)
	}
	if v, ok := attributes["match_criteria"].(string); ok && v != "" {
		data.MatchT = types.StringValue(v)
	}
	if v, ok := attributes["name"].(string); ok && v != "" {
		data.Name = types.StringValue(v)
	}
	if v, ok := attributes["name_alias"].(string); ok && v != "" {
		data.NameAlias = types.StringValue(v)
	}
	if v, ok := attributes["intra_epg_isolation"].(string); ok && v != "" {
		data.PcEnfPref = types.StringValue(v)
	}
	if v, ok := attributes["pc_tag"].(string); ok && v != "" {
		data.PcTag = types.StringValue(v)
	}
	if v, ok := attributes["preferred_group_member"].(string); ok && v != "" {
		data.PrefGrMemb = types.StringValue(v)
	}
	if v, ok := attributes["priority"].(string); ok && v != "" {
		data.Prio = types.StringValue(v)
	}
	if v, ok := attributes["admin_state"].(string); ok && v != "" {
		data.Shutdown = types.StringValue(v)
	}
	planFvCrtrn := convertToFvCrtrnFvAEPg(attributes["epg_useg_block_statement"])
	planFvRsAEPgMonPol := convertToFvRsAEPgMonPolFvAEPg(attributes["relation_to_application_epg_monitoring_policy"])
	planFvRsBd := convertToFvRsBdFvAEPg(attributes["relation_to_bridge_domain"])
	planFvRsCons := convertToFvRsConsFvAEPg(attributes["relation_to_consumed_contracts"])
	planFvRsConsIf := convertToFvRsConsIfFvAEPg(attributes["relation_to_imported_contracts"])
	planFvRsCustQosPol := convertToFvRsCustQosPolFvAEPg(attributes["relation_to_custom_qos_policy"])
	planFvRsDomAtt := convertToFvRsDomAttFvAEPg(attributes["relation_to_domains"])
	planFvRsDppPol := convertToFvRsDppPolFvAEPg(attributes["relation_to_data_plane_policing_policy"])
	planFvRsFcPathAtt := convertToFvRsFcPathAttFvAEPg(attributes["relation_to_fibre_channel_paths"])
	planFvRsIntraEpg := convertToFvRsIntraEpgFvAEPg(attributes["relation_to_intra_epg_contracts"])
	planFvRsNodeAtt := convertToFvRsNodeAttFvAEPg(attributes["relation_to_static_leafs"])
	planFvRsPathAtt := convertToFvRsPathAttFvAEPg(attributes["relation_to_static_paths"])
	planFvRsProtBy := convertToFvRsProtByFvAEPg(attributes["relation_to_taboo_contracts"])
	planFvRsProv := convertToFvRsProvFvAEPg(attributes["relation_to_provided_contracts"])
	planFvRsSecInherited := convertToFvRsSecInheritedFvAEPg(attributes["relation_to_contract_masters"])
	planFvRsTrustCtrl := convertToFvRsTrustCtrlFvAEPg(attributes["relation_to_trust_control_policy"])
	planTagAnnotation := convertToTagAnnotationFvAEPg(attributes["annotations"])
	planTagTag := convertToTagTagFvAEPg(attributes["tags"])

	newAciFvAEPg := provider.GetFvAEPgCreateJsonPayload(ctx, &diags, true, data, planFvCrtrn, planFvCrtrn, planFvRsAEPgMonPol, planFvRsAEPgMonPol, planFvRsBd, planFvRsBd, planFvRsCons, planFvRsCons, planFvRsConsIf, planFvRsConsIf, planFvRsCustQosPol, planFvRsCustQosPol, planFvRsDomAtt, planFvRsDomAtt, planFvRsDppPol, planFvRsDppPol, planFvRsFcPathAtt, planFvRsFcPathAtt, planFvRsIntraEpg, planFvRsIntraEpg, planFvRsNodeAtt, planFvRsNodeAtt, planFvRsPathAtt, planFvRsPathAtt, planFvRsProtBy, planFvRsProtBy, planFvRsProv, planFvRsProv, planFvRsSecInherited, planFvRsSecInherited, planFvRsTrustCtrl, planFvRsTrustCtrl, planTagAnnotation, planTagAnnotation, planTagTag, planTagTag)

	jsonPayload := newAciFvAEPg.EncodeJSON(container.EncodeOptIndent("", "  "))

	var customData map[string]interface{}
	json.Unmarshal(jsonPayload, &customData)

	payload := customData

	provider.SetFvAEPgId(ctx, data)
	attrs := payload["fvAEPg"].(map[string]interface{})["attributes"].(map[string]interface{})
	attrs["dn"] = data.Id.ValueString()

	return payload
}
func convertToFvCrtrnFvAEPg(resources interface{}) []provider.FvCrtrnFvAEPgResourceModel {
	var planResources []provider.FvCrtrnFvAEPgResourceModel
	if resources, ok := resources.([]interface{}); ok {
		for _, resource := range resources {
			resourceMap := resource.(map[string]interface{})
			planResources = append(planResources, provider.FvCrtrnFvAEPgResourceModel{
				Annotation: types.StringValue(resourceMap["annotation"].(string)),
				Descr:      types.StringValue(resourceMap["description"].(string)),
				Match:      types.StringValue(resourceMap["match"].(string)),
				Name:       types.StringValue(resourceMap["name"].(string)),
				NameAlias:  types.StringValue(resourceMap["name_alias"].(string)),
				OwnerKey:   types.StringValue(resourceMap["owner_key"].(string)),
				OwnerTag:   types.StringValue(resourceMap["owner_tag"].(string)),
				Prec:       types.StringValue(resourceMap["precedence"].(string)),
				Scope:      types.StringValue(resourceMap["scope"].(string)),
			})
		}
	}
	return planResources
}
func convertToFvRsAEPgMonPolFvAEPg(resources interface{}) []provider.FvRsAEPgMonPolFvAEPgResourceModel {
	var planResources []provider.FvRsAEPgMonPolFvAEPgResourceModel
	if resources, ok := resources.([]interface{}); ok {
		for _, resource := range resources {
			resourceMap := resource.(map[string]interface{})
			planResources = append(planResources, provider.FvRsAEPgMonPolFvAEPgResourceModel{
				Annotation:      types.StringValue(resourceMap["annotation"].(string)),
				TnMonEPGPolName: types.StringValue(resourceMap["monitoring_policy_name"].(string)),
			})
		}
	}
	return planResources
}
func convertToFvRsBdFvAEPg(resources interface{}) []provider.FvRsBdFvAEPgResourceModel {
	var planResources []provider.FvRsBdFvAEPgResourceModel
	if resources, ok := resources.([]interface{}); ok {
		for _, resource := range resources {
			resourceMap := resource.(map[string]interface{})
			planResources = append(planResources, provider.FvRsBdFvAEPgResourceModel{
				Annotation: types.StringValue(resourceMap["annotation"].(string)),
				TnFvBDName: types.StringValue(resourceMap["bridge_domain_name"].(string)),
			})
		}
	}
	return planResources
}
func convertToFvRsConsFvAEPg(resources interface{}) []provider.FvRsConsFvAEPgResourceModel {
	var planResources []provider.FvRsConsFvAEPgResourceModel
	if resources, ok := resources.([]interface{}); ok {
		for _, resource := range resources {
			resourceMap := resource.(map[string]interface{})
			planResources = append(planResources, provider.FvRsConsFvAEPgResourceModel{
				Annotation:   types.StringValue(resourceMap["annotation"].(string)),
				Prio:         types.StringValue(resourceMap["priority"].(string)),
				TnVzBrCPName: types.StringValue(resourceMap["contract_name"].(string)),
			})
		}
	}
	return planResources
}
func convertToFvRsConsIfFvAEPg(resources interface{}) []provider.FvRsConsIfFvAEPgResourceModel {
	var planResources []provider.FvRsConsIfFvAEPgResourceModel
	if resources, ok := resources.([]interface{}); ok {
		for _, resource := range resources {
			resourceMap := resource.(map[string]interface{})
			planResources = append(planResources, provider.FvRsConsIfFvAEPgResourceModel{
				Annotation:   types.StringValue(resourceMap["annotation"].(string)),
				Prio:         types.StringValue(resourceMap["priority"].(string)),
				TnVzCPIfName: types.StringValue(resourceMap["imported_contract_name"].(string)),
			})
		}
	}
	return planResources
}
func convertToFvRsCustQosPolFvAEPg(resources interface{}) []provider.FvRsCustQosPolFvAEPgResourceModel {
	var planResources []provider.FvRsCustQosPolFvAEPgResourceModel
	if resources, ok := resources.([]interface{}); ok {
		for _, resource := range resources {
			resourceMap := resource.(map[string]interface{})
			planResources = append(planResources, provider.FvRsCustQosPolFvAEPgResourceModel{
				Annotation:         types.StringValue(resourceMap["annotation"].(string)),
				TnQosCustomPolName: types.StringValue(resourceMap["custom_qos_policy_name"].(string)),
			})
		}
	}
	return planResources
}
func convertToFvRsDomAttFvAEPg(resources interface{}) []provider.FvRsDomAttFvAEPgResourceModel {
	var planResources []provider.FvRsDomAttFvAEPgResourceModel
	if resources, ok := resources.([]interface{}); ok {
		for _, resource := range resources {
			resourceMap := resource.(map[string]interface{})
			planResources = append(planResources, provider.FvRsDomAttFvAEPgResourceModel{
				Annotation:          types.StringValue(resourceMap["annotation"].(string)),
				BindingType:         types.StringValue(resourceMap["binding_type"].(string)),
				ClassPref:           types.StringValue(resourceMap["class_preference"].(string)),
				CustomEpgName:       types.StringValue(resourceMap["custom_epg_name"].(string)),
				Delimiter:           types.StringValue(resourceMap["delimiter"].(string)),
				Encap:               types.StringValue(resourceMap["encapsulation"].(string)),
				EncapMode:           types.StringValue(resourceMap["encapsulation_mode"].(string)),
				EpgCos:              types.StringValue(resourceMap["epg_cos"].(string)),
				EpgCosPref:          types.StringValue(resourceMap["epg_cos_pref"].(string)),
				InstrImedcy:         types.StringValue(resourceMap["deployment_immediacy"].(string)),
				IpamDhcpOverride:    types.StringValue(resourceMap["ipam_dhcp_override"].(string)),
				IpamEnabled:         types.StringValue(resourceMap["ipam_enabled"].(string)),
				IpamGateway:         types.StringValue(resourceMap["ipam_gateway"].(string)),
				LagPolicyName:       types.StringValue(resourceMap["lag_policy_name"].(string)),
				NetflowDir:          types.StringValue(resourceMap["netflow_direction"].(string)),
				NetflowPref:         types.StringValue(resourceMap["enable_netflow"].(string)),
				NumPorts:            types.StringValue(resourceMap["number_of_ports"].(string)),
				PortAllocation:      types.StringValue(resourceMap["port_allocation"].(string)),
				PrimaryEncap:        types.StringValue(resourceMap["primary_encapsulation"].(string)),
				PrimaryEncapInner:   types.StringValue(resourceMap["primary_encapsulation_inner"].(string)),
				ResImedcy:           types.StringValue(resourceMap["resolution_immediacy"].(string)),
				SecondaryEncapInner: types.StringValue(resourceMap["secondary_encapsulation_inner"].(string)),
				SwitchingMode:       types.StringValue(resourceMap["switching_mode"].(string)),
				TDn:                 types.StringValue(resourceMap["target_dn"].(string)),
				Untagged:            types.StringValue(resourceMap["untagged"].(string)),
				VnetOnly:            types.StringValue(resourceMap["vnet_only"].(string)),
			})
		}
	}
	return planResources
}
func convertToFvRsDppPolFvAEPg(resources interface{}) []provider.FvRsDppPolFvAEPgResourceModel {
	var planResources []provider.FvRsDppPolFvAEPgResourceModel
	if resources, ok := resources.([]interface{}); ok {
		for _, resource := range resources {
			resourceMap := resource.(map[string]interface{})
			planResources = append(planResources, provider.FvRsDppPolFvAEPgResourceModel{
				Annotation:      types.StringValue(resourceMap["annotation"].(string)),
				TnQosDppPolName: types.StringValue(resourceMap["data_plane_policing_policy_name"].(string)),
			})
		}
	}
	return planResources
}
func convertToFvRsFcPathAttFvAEPg(resources interface{}) []provider.FvRsFcPathAttFvAEPgResourceModel {
	var planResources []provider.FvRsFcPathAttFvAEPgResourceModel
	if resources, ok := resources.([]interface{}); ok {
		for _, resource := range resources {
			resourceMap := resource.(map[string]interface{})
			planResources = append(planResources, provider.FvRsFcPathAttFvAEPgResourceModel{
				Annotation: types.StringValue(resourceMap["annotation"].(string)),
				Descr:      types.StringValue(resourceMap["description"].(string)),
				TDn:        types.StringValue(resourceMap["target_dn"].(string)),
				Vsan:       types.StringValue(resourceMap["vsan"].(string)),
				VsanMode:   types.StringValue(resourceMap["vsan_mode"].(string)),
			})
		}
	}
	return planResources
}
func convertToFvRsIntraEpgFvAEPg(resources interface{}) []provider.FvRsIntraEpgFvAEPgResourceModel {
	var planResources []provider.FvRsIntraEpgFvAEPgResourceModel
	if resources, ok := resources.([]interface{}); ok {
		for _, resource := range resources {
			resourceMap := resource.(map[string]interface{})
			planResources = append(planResources, provider.FvRsIntraEpgFvAEPgResourceModel{
				Annotation:   types.StringValue(resourceMap["annotation"].(string)),
				TnVzBrCPName: types.StringValue(resourceMap["contract_name"].(string)),
			})
		}
	}
	return planResources
}
func convertToFvRsNodeAttFvAEPg(resources interface{}) []provider.FvRsNodeAttFvAEPgResourceModel {
	var planResources []provider.FvRsNodeAttFvAEPgResourceModel
	if resources, ok := resources.([]interface{}); ok {
		for _, resource := range resources {
			resourceMap := resource.(map[string]interface{})
			planResources = append(planResources, provider.FvRsNodeAttFvAEPgResourceModel{
				Annotation:  types.StringValue(resourceMap["annotation"].(string)),
				Descr:       types.StringValue(resourceMap["description"].(string)),
				Encap:       types.StringValue(resourceMap["encapsulation"].(string)),
				InstrImedcy: types.StringValue(resourceMap["deployment_immediacy"].(string)),
				Mode:        types.StringValue(resourceMap["mode"].(string)),
				TDn:         types.StringValue(resourceMap["target_dn"].(string)),
			})
		}
	}
	return planResources
}
func convertToFvRsPathAttFvAEPg(resources interface{}) []provider.FvRsPathAttFvAEPgResourceModel {
	var planResources []provider.FvRsPathAttFvAEPgResourceModel
	if resources, ok := resources.([]interface{}); ok {
		for _, resource := range resources {
			resourceMap := resource.(map[string]interface{})
			planResources = append(planResources, provider.FvRsPathAttFvAEPgResourceModel{
				Annotation:   types.StringValue(resourceMap["annotation"].(string)),
				Descr:        types.StringValue(resourceMap["description"].(string)),
				Encap:        types.StringValue(resourceMap["encapsulation"].(string)),
				InstrImedcy:  types.StringValue(resourceMap["deployment_immediacy"].(string)),
				Mode:         types.StringValue(resourceMap["mode"].(string)),
				PrimaryEncap: types.StringValue(resourceMap["primary_encapsulation"].(string)),
				TDn:          types.StringValue(resourceMap["target_dn"].(string)),
			})
		}
	}
	return planResources
}
func convertToFvRsProtByFvAEPg(resources interface{}) []provider.FvRsProtByFvAEPgResourceModel {
	var planResources []provider.FvRsProtByFvAEPgResourceModel
	if resources, ok := resources.([]interface{}); ok {
		for _, resource := range resources {
			resourceMap := resource.(map[string]interface{})
			planResources = append(planResources, provider.FvRsProtByFvAEPgResourceModel{
				Annotation:    types.StringValue(resourceMap["annotation"].(string)),
				TnVzTabooName: types.StringValue(resourceMap["taboo_contract_name"].(string)),
			})
		}
	}
	return planResources
}
func convertToFvRsProvFvAEPg(resources interface{}) []provider.FvRsProvFvAEPgResourceModel {
	var planResources []provider.FvRsProvFvAEPgResourceModel
	if resources, ok := resources.([]interface{}); ok {
		for _, resource := range resources {
			resourceMap := resource.(map[string]interface{})
			planResources = append(planResources, provider.FvRsProvFvAEPgResourceModel{
				Annotation:   types.StringValue(resourceMap["annotation"].(string)),
				MatchT:       types.StringValue(resourceMap["match_criteria"].(string)),
				Prio:         types.StringValue(resourceMap["priority"].(string)),
				TnVzBrCPName: types.StringValue(resourceMap["contract_name"].(string)),
			})
		}
	}
	return planResources
}
func convertToFvRsSecInheritedFvAEPg(resources interface{}) []provider.FvRsSecInheritedFvAEPgResourceModel {
	var planResources []provider.FvRsSecInheritedFvAEPgResourceModel
	if resources, ok := resources.([]interface{}); ok {
		for _, resource := range resources {
			resourceMap := resource.(map[string]interface{})
			planResources = append(planResources, provider.FvRsSecInheritedFvAEPgResourceModel{
				Annotation: types.StringValue(resourceMap["annotation"].(string)),
				TDn:        types.StringValue(resourceMap["target_dn"].(string)),
			})
		}
	}
	return planResources
}
func convertToFvRsTrustCtrlFvAEPg(resources interface{}) []provider.FvRsTrustCtrlFvAEPgResourceModel {
	var planResources []provider.FvRsTrustCtrlFvAEPgResourceModel
	if resources, ok := resources.([]interface{}); ok {
		for _, resource := range resources {
			resourceMap := resource.(map[string]interface{})
			planResources = append(planResources, provider.FvRsTrustCtrlFvAEPgResourceModel{
				Annotation:            types.StringValue(resourceMap["annotation"].(string)),
				TnFhsTrustCtrlPolName: types.StringValue(resourceMap["trust_control_policy_name"].(string)),
			})
		}
	}
	return planResources
}
func convertToTagAnnotationFvAEPg(resources interface{}) []provider.TagAnnotationFvAEPgResourceModel {
	var planResources []provider.TagAnnotationFvAEPgResourceModel
	if resources, ok := resources.([]interface{}); ok {
		for _, resource := range resources {
			resourceMap := resource.(map[string]interface{})
			planResources = append(planResources, provider.TagAnnotationFvAEPgResourceModel{
				Key:   types.StringValue(resourceMap["key"].(string)),
				Value: types.StringValue(resourceMap["value"].(string)),
			})
		}
	}
	return planResources
}
func convertToTagTagFvAEPg(resources interface{}) []provider.TagTagFvAEPgResourceModel {
	var planResources []provider.TagTagFvAEPgResourceModel
	if resources, ok := resources.([]interface{}); ok {
		for _, resource := range resources {
			resourceMap := resource.(map[string]interface{})
			planResources = append(planResources, provider.TagTagFvAEPgResourceModel{
				Key:   types.StringValue(resourceMap["key"].(string)),
				Value: types.StringValue(resourceMap["value"].(string)),
			})
		}
	}
	return planResources
}
