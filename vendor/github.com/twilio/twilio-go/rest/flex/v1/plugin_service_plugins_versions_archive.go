/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Flex
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"net/url"
	"strings"
)

// Optional parameters for the method 'UpdatePluginVersionArchive'
type UpdatePluginVersionArchiveParams struct {
	// The Flex-Metadata HTTP request header
	FlexMetadata *string `json:"Flex-Metadata,omitempty"`
}

func (params *UpdatePluginVersionArchiveParams) SetFlexMetadata(FlexMetadata string) *UpdatePluginVersionArchiveParams {
	params.FlexMetadata = &FlexMetadata
	return params
}

//
func (c *ApiService) UpdatePluginVersionArchive(PluginSid string, Sid string, params *UpdatePluginVersionArchiveParams) (*FlexV1PluginVersionArchive, error) {
	path := "/v1/PluginService/Plugins/{PluginSid}/Versions/{Sid}/Archive"
	path = strings.Replace(path, "{"+"PluginSid"+"}", PluginSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FlexMetadata != nil {
		headers["Flex-Metadata"] = *params.FlexMetadata
	}
	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1PluginVersionArchive{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}