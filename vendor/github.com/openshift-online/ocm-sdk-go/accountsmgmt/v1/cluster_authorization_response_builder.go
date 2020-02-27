/*
Copyright (c) 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

// ClusterAuthorizationResponseBuilder contains the data and logic needed to build 'cluster_authorization_response' objects.
//
//
type ClusterAuthorizationResponseBuilder struct {
	allowed         *bool
	excessResources []*ReservedResourceBuilder
	subscription    *SubscriptionBuilder
}

// NewClusterAuthorizationResponse creates a new builder of 'cluster_authorization_response' objects.
func NewClusterAuthorizationResponse() *ClusterAuthorizationResponseBuilder {
	return new(ClusterAuthorizationResponseBuilder)
}

// Allowed sets the value of the 'allowed' attribute to the given value.
//
//
func (b *ClusterAuthorizationResponseBuilder) Allowed(value bool) *ClusterAuthorizationResponseBuilder {
	b.allowed = &value
	return b
}

// ExcessResources sets the value of the 'excess_resources' attribute to the given values.
//
//
func (b *ClusterAuthorizationResponseBuilder) ExcessResources(values ...*ReservedResourceBuilder) *ClusterAuthorizationResponseBuilder {
	b.excessResources = make([]*ReservedResourceBuilder, len(values))
	copy(b.excessResources, values)
	return b
}

// Subscription sets the value of the 'subscription' attribute to the given value.
//
//
func (b *ClusterAuthorizationResponseBuilder) Subscription(value *SubscriptionBuilder) *ClusterAuthorizationResponseBuilder {
	b.subscription = value
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *ClusterAuthorizationResponseBuilder) Copy(object *ClusterAuthorizationResponse) *ClusterAuthorizationResponseBuilder {
	if object == nil {
		return b
	}
	b.allowed = object.allowed
	if object.excessResources != nil {
		b.excessResources = make([]*ReservedResourceBuilder, len(object.excessResources))
		for i, v := range object.excessResources {
			b.excessResources[i] = NewReservedResource().Copy(v)
		}
	} else {
		b.excessResources = nil
	}
	if object.subscription != nil {
		b.subscription = NewSubscription().Copy(object.subscription)
	} else {
		b.subscription = nil
	}
	return b
}

// Build creates a 'cluster_authorization_response' object using the configuration stored in the builder.
func (b *ClusterAuthorizationResponseBuilder) Build() (object *ClusterAuthorizationResponse, err error) {
	object = new(ClusterAuthorizationResponse)
	object.allowed = b.allowed
	if b.excessResources != nil {
		object.excessResources = make([]*ReservedResource, len(b.excessResources))
		for i, v := range b.excessResources {
			object.excessResources[i], err = v.Build()
			if err != nil {
				return
			}
		}
	}
	if b.subscription != nil {
		object.subscription, err = b.subscription.Build()
		if err != nil {
			return
		}
	}
	return
}
