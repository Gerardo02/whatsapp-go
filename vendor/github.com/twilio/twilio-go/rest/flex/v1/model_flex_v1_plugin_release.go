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
	"time"
)

// FlexV1PluginRelease struct for FlexV1PluginRelease
type FlexV1PluginRelease struct {
	// The unique string that we created to identify the Plugin Release resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Plugin Release resource and owns this resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Plugin Configuration resource to release.
	ConfigurationSid *string `json:"configuration_sid,omitempty"`
	// The date and time in GMT when the Flex Plugin Release was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The absolute URL of the Plugin Release resource.
	Url *string `json:"url,omitempty"`
}
