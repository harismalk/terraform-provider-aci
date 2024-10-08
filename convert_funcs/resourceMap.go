package convert_funcs

type createFunc func(map[string]interface{}) map[string]interface{}

var ResourceMap = map[string]createFunc{
	"aci_relation_to_taboo_contract":                   CreateFvRsProtBy,
	"aci_pim_route_map_entry":                          CreatePimRouteMapEntry,
	"aci_endpoint_tag_ip":                              CreateFvEpIpTag,
	"aci_epg_useg_ad_group_attribute":                  CreateFvIdGroupAttr,
	"aci_epg_useg_ip_attribute":                        CreateFvIpAttr,
	"aci_epg_useg_mac_attribute":                       CreateFvMacAttr,
	"aci_relation_to_consumed_contract":                CreateFvRsCons,
	"aci_relation_to_imported_contract":                CreateFvRsConsIf,
	"aci_trust_control_policy":                         CreateFhsTrustCtrlPol,
	"aci_endpoint_tag_mac":                             CreateFvEpMacTag,
	"aci_vrf_fallback_route_group_member":              CreateFvFBRMember,
	"aci_custom_qos_policy":                            CreateQosCustomPol,
	"aci_relation_to_static_path":                      CreateFvRsPathAtt,
	"aci_epg_useg_sub_block_statement":                 CreateFvSCrtrn,
	"aci_relation_to_vrf_fallback_route_group":         CreateL3extRsOutToFBRGroup,
	"aci_relation_to_consumed_out_of_band_contract":    CreateMgmtRsOoBCons,
	"aci_netflow_monitor_policy":                       CreateNetflowMonitorPol,
	"aci_relation_to_domain":                           CreateFvRsDomAtt,
	"aci_relation_to_intra_epg_contract":               CreateFvRsIntraEpg,
	"aci_pim_route_map_policy":                         CreatePimRouteMapPol,
	"aci_route_control_profile":                        CreateRtctrlProfile,
	"aci_l3out_node_sid_profile":                       CreateMplsNodeSidP,
	"aci_netflow_record_policy":                        CreateNetflowRecordPol,
	"aci_out_of_band_contract":                         CreateVzOOBBrCP,
	"aci_relation_to_fibre_channel_path":               CreateFvRsFcPathAtt,
	"aci_epg_useg_vm_attribute":                        CreateFvVmAttr,
	"aci_relation_to_static_leaf":                      CreateFvRsNodeAtt,
	"aci_annotation":                                   CreateTagAnnotation,
	"aci_vrf_fallback_route_group":                     CreateFvFBRGroup,
	"aci_vrf_fallback_route":                           CreateFvFBRoute,
	"aci_tag":                                          CreateTagTag,
	"aci_external_management_network_instance_profile": CreateMgmtInstP,
	"aci_relation_to_netflow_exporter":                 CreateNetflowRsMonitorToExporter,
	"aci_data_plane_policing_policy":                   CreateQosDppPol,
	"aci_relation_to_provided_contract":                CreateFvRsProv,
	"aci_l3out_provider_label":                         CreateL3extProvLbl,
	"aci_application_epg":                              CreateFvAEPg,
	"aci_epg_useg_block_statement":                     CreateFvCrtrn,
	"aci_endpoint_security_group":                      CreateFvESg,
	"aci_l3out_redistribute_policy":                    CreateL3extRsRedistributePol,
	"aci_external_management_network_subnet":           CreateMgmtSubnet,
	"aci_epg_useg_dns_attribute":                       CreateFvDnsAttr,
	"aci_relation_to_contract_master":                  CreateFvRsSecInherited,
	"aci_l3out_consumer_label":                         CreateL3extConsLbl,
}
