/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// ListUsageTriggerResponse struct for ListUsageTriggerResponse
type ListUsageTriggerResponse struct {
	UsageTriggers   []ApiV2010UsageTrigger `json:"usage_triggers,omitempty"`
	End             int                    `json:"end,omitempty"`
	FirstPageUri    string                 `json:"first_page_uri,omitempty"`
	NextPageUri     *string                `json:"next_page_uri,omitempty"`
	Page            int                    `json:"page,omitempty"`
	PageSize        int                    `json:"page_size,omitempty"`
	PreviousPageUri *string                `json:"previous_page_uri,omitempty"`
	Start           int                    `json:"start,omitempty"`
	Uri             string                 `json:"uri,omitempty"`
}
