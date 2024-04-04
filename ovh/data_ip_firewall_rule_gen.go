// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package ovh

import (
	"context"

	ovhtypes "github.com/ovh/terraform-provider-ovh/ovh/types"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func IpFirewallRuleDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"action": schema.StringAttribute{
				CustomType:          ovhtypes.TfStringType{},
				Computed:            true,
				Description:         "Possible values for action",
				MarkdownDescription: "Possible values for action",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"deny",
						"permit",
					),
				},
			},
			"creation_date": schema.StringAttribute{
				CustomType: ovhtypes.TfStringType{},
				Computed:   true,
			},
			"destination": schema.StringAttribute{
				CustomType:          ovhtypes.TfStringType{},
				Computed:            true,
				Description:         "Destination ip for your rule",
				MarkdownDescription: "Destination ip for your rule",
			},
			"destination_port": schema.Int64Attribute{
				CustomType:          ovhtypes.TfInt64Type{},
				Computed:            true,
				Description:         "Destination port for your rule. Only with TCP/UDP protocol",
				MarkdownDescription: "Destination port for your rule. Only with TCP/UDP protocol",
			},
			"fragments": schema.BoolAttribute{
				CustomType:          ovhtypes.TfBoolType{},
				Computed:            true,
				Description:         "Fragments option",
				MarkdownDescription: "Fragments option",
			},
			"ip": schema.StringAttribute{
				CustomType:          ovhtypes.TfStringType{},
				Required:            true,
				Description:         "IP (v4 or v6) CIDR notation (e.g., 192.0.2.0/24)",
				MarkdownDescription: "IP (v4 or v6) CIDR notation (e.g., 192.0.2.0/24)",
			},
			"ip_on_firewall": schema.StringAttribute{
				CustomType:          ovhtypes.TfStringType{},
				Required:            true,
				Description:         "IPv4 address (e.g., 192.0.2.0)",
				MarkdownDescription: "IPv4 address (e.g., 192.0.2.0)",
			},
			"protocol": schema.StringAttribute{
				CustomType:          ovhtypes.TfStringType{},
				Computed:            true,
				Description:         "Possible values for protocol",
				MarkdownDescription: "Possible values for protocol",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"ah",
						"esp",
						"gre",
						"icmp",
						"ipv4",
						"tcp",
						"udp",
					),
				},
			},
			"rule": schema.StringAttribute{
				CustomType: ovhtypes.TfStringType{},
				Computed:   true,
			},
			"sequence": schema.Int64Attribute{
				CustomType:          ovhtypes.TfInt64Type{},
				Required:            true,
				Description:         "Possible values for action",
				MarkdownDescription: "Possible values for action",
				Validators: []validator.Int64{
					int64validator.OneOf(
						0,
						1,
						2,
						3,
						4,
						5,
						6,
						7,
						8,
						9,
						10,
						11,
						12,
						13,
						14,
						15,
						16,
						17,
						18,
						19,
					),
				},
			},
			"source": schema.StringAttribute{
				CustomType:          ovhtypes.TfStringType{},
				Computed:            true,
				Description:         "IPv4 CIDR notation (e.g., 192.0.2.0/24)",
				MarkdownDescription: "IPv4 CIDR notation (e.g., 192.0.2.0/24)",
			},
			"source_port": schema.Int64Attribute{
				CustomType:          ovhtypes.TfInt64Type{},
				Computed:            true,
				Description:         "Source port for your rule. Only with TCP/UDP protocol",
				MarkdownDescription: "Source port for your rule. Only with TCP/UDP protocol",
			},
			"state": schema.StringAttribute{
				CustomType:          ovhtypes.TfStringType{},
				Computed:            true,
				Description:         "Current state of your rule",
				MarkdownDescription: "Current state of your rule",
			},
			"tcp_option": schema.StringAttribute{
				CustomType:          ovhtypes.TfStringType{},
				Computed:            true,
				Description:         "TCP option on your rule",
				MarkdownDescription: "TCP option on your rule",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"established",
						"syn",
					),
				},
			},
		},
	}
}