/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * See package.json for author and maintainer info
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */

package models

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SubscriptionsGetNotificationsResponse subscriptions get notifications response
// swagger:model SubscriptionsGetNotificationsResponse
type SubscriptionsGetNotificationsResponse struct {

	// Identifies what kind of resource this is. Value: the fixed string "weave#subscriptionsGetNotificationsResponse".
	Kind *string `json:"kind,omitempty"`

	// Past client notifications.
	Notifications []*ClientNotification `json:"notifications"`
}

// Validate validates this subscriptions get notifications response
func (m *SubscriptionsGetNotificationsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNotifications(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubscriptionsGetNotificationsResponse) validateNotifications(formats strfmt.Registry) error {

	if swag.IsZero(m.Notifications) { // not required
		return nil
	}

	for i := 0; i < len(m.Notifications); i++ {

		if swag.IsZero(m.Notifications[i]) { // not required
			continue
		}

		if m.Notifications[i] != nil {

			if err := m.Notifications[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}