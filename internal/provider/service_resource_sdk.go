// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"clickhouse/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"math/big"
	"time"
)

func (r *ServiceResourceModel) ToCreateSDKType() *shared.ServicePostRequest {
	backupID := new(string)
	if !r.BackupID.IsUnknown() && !r.BackupID.IsNull() {
		*backupID = r.BackupID.ValueString()
	} else {
		backupID = nil
	}
	cloudProvider := new(shared.ServicePostRequestCloudProvider)
	if !r.CloudProvider.IsUnknown() && !r.CloudProvider.IsNull() {
		*cloudProvider = shared.ServicePostRequestCloudProvider(r.CloudProvider.ValueString())
	} else {
		cloudProvider = nil
	}
	idleScaling := r.IdleScaling.ValueBool()
	idleTimeoutMinutes, _ := r.IdleTimeoutMinutes.ValueBigFloat().Float64()
	ipAccessList := make([]shared.IPAccessListEntry, 0)
	for _, ipAccessListItem := range r.IPAccessList {
		description := ipAccessListItem.Description.ValueString()
		source := ipAccessListItem.Source.ValueString()
		ipAccessList = append(ipAccessList, shared.IPAccessListEntry{
			Description: description,
			Source:      source,
		})
	}
	maxTotalMemoryGb := new(float64)
	if !r.MaxTotalMemoryGb.IsUnknown() && !r.MaxTotalMemoryGb.IsNull() {
		*maxTotalMemoryGb, _ = r.MaxTotalMemoryGb.ValueBigFloat().Float64()
	} else {
		maxTotalMemoryGb = nil
	}
	minTotalMemoryGb := new(float64)
	if !r.MinTotalMemoryGb.IsUnknown() && !r.MinTotalMemoryGb.IsNull() {
		*minTotalMemoryGb, _ = r.MinTotalMemoryGb.ValueBigFloat().Float64()
	} else {
		minTotalMemoryGb = nil
	}
	name := r.Name.ValueString()
	region := shared.ServicePostRequestRegion(r.Region.ValueString())
	tier := shared.ServicePostRequestTier(r.Tier.ValueString())
	out := shared.ServicePostRequest{
		BackupID:           backupID,
		CloudProvider:      cloudProvider,
		IdleScaling:        idleScaling,
		IdleTimeoutMinutes: idleTimeoutMinutes,
		IPAccessList:       ipAccessList,
		MaxTotalMemoryGb:   maxTotalMemoryGb,
		MinTotalMemoryGb:   minTotalMemoryGb,
		Name:               name,
		Region:             region,
		Tier:               tier,
	}
	return &out
}

func (r *ServiceResourceModel) ToGetSDKType() *shared.ServicePostRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *ServiceResourceModel) ToUpdateSDKType() *shared.ServicePatchRequest {
	var ipAccessList *shared.IPAccessListPatch
	if r.IPAccessList != nil {
		ipAccessList = &shared.IPAccessListPatch{}
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}
	out := shared.ServicePatchRequest{
		IPAccessList: ipAccessList,
		Name:         name,
	}
	return &out
}

func (r *ServiceResourceModel) ToDeleteSDKType() *shared.ServicePostRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *ServiceResourceModel) RefreshFromGetResponse(resp *shared.Service) {
	if resp.CloudProvider != nil {
		r.CloudProvider = types.StringValue(string(*resp.CloudProvider))
	} else {
		r.CloudProvider = types.StringNull()
	}
	if resp.CreatedAt != nil {
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339))
	} else {
		r.CreatedAt = types.StringNull()
	}
	r.Endpoints = nil
	for _, endpointsItem := range resp.Endpoints {
		var endpoints1 ServiceEndpoint
		if endpointsItem.Host != nil {
			endpoints1.Host = types.StringValue(*endpointsItem.Host)
		} else {
			endpoints1.Host = types.StringNull()
		}
		if endpointsItem.Port != nil {
			endpoints1.Port = types.NumberValue(big.NewFloat(*endpointsItem.Port))
		} else {
			endpoints1.Port = types.NumberNull()
		}
		if endpointsItem.Protocol != nil {
			endpoints1.Protocol = types.StringValue(string(*endpointsItem.Protocol))
		} else {
			endpoints1.Protocol = types.StringNull()
		}
		r.Endpoints = append(r.Endpoints, endpoints1)
	}
	if resp.ID != nil {
		r.ID = types.StringValue(*resp.ID)
	} else {
		r.ID = types.StringNull()
	}
	if resp.IdleScaling != nil {
		r.IdleScaling = types.BoolValue(*resp.IdleScaling)
	} else {
		r.IdleScaling = types.BoolNull()
	}
	if resp.IdleTimeoutMinutes != nil {
		r.IdleTimeoutMinutes = types.NumberValue(big.NewFloat(*resp.IdleTimeoutMinutes))
	} else {
		r.IdleTimeoutMinutes = types.NumberNull()
	}
	r.IPAccessList = nil
	for _, ipAccessListItem := range resp.IPAccessList {
		var ipAccessList1 IPAccessListEntry
		ipAccessList1.Description = types.StringValue(ipAccessListItem.Description)
		ipAccessList1.Source = types.StringValue(ipAccessListItem.Source)
		r.IPAccessList = append(r.IPAccessList, ipAccessList1)
	}
	if resp.MaxTotalMemoryGb != nil {
		r.MaxTotalMemoryGb = types.NumberValue(big.NewFloat(*resp.MaxTotalMemoryGb))
	} else {
		r.MaxTotalMemoryGb = types.NumberNull()
	}
	if resp.MinTotalMemoryGb != nil {
		r.MinTotalMemoryGb = types.NumberValue(big.NewFloat(*resp.MinTotalMemoryGb))
	} else {
		r.MinTotalMemoryGb = types.NumberNull()
	}
	if resp.Name != nil {
		r.Name = types.StringValue(*resp.Name)
	} else {
		r.Name = types.StringNull()
	}
	if resp.Region != nil {
		r.Region = types.StringValue(string(*resp.Region))
	} else {
		r.Region = types.StringNull()
	}
	if resp.State != nil {
		r.State = types.StringValue(string(*resp.State))
	} else {
		r.State = types.StringNull()
	}
	if resp.Tier != nil {
		r.Tier = types.StringValue(string(*resp.Tier))
	} else {
		r.Tier = types.StringNull()
	}
}

func (r *ServiceResourceModel) RefreshFromCreateResponse(resp *shared.Service) {
	r.RefreshFromGetResponse(resp)
}

func (r *ServiceResourceModel) RefreshFromUpdateResponse(resp *shared.Service) {
	r.RefreshFromGetResponse(resp)
}
