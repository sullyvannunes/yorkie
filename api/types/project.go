/*
 * Copyright 2022 The Yorkie Authors. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package types

import "time"

// Project is a project that consists of multiple documents and clients.
// It allows developers to work on the same cluster.
type Project struct {
	// ID is the unique ID of the project.
	ID ID `json:"id"`

	// Name is the name of this project.
	Name string `json:"name"`

	// PublicKey is the API key of this project.
	PublicKey string `json:"public_key"`

	// SecretKey is the secret key of this project.
	SecretKey string `json:"secret_key"`

	// AuthWebhookURL is the url of the authorization webhook.
	AuthWebhookURL string `json:"auth_webhook_url"`

	// AuthWebhookMethods is the methods that run the authorization webhook.
	AuthWebhookMethods []string `json:"auth_webhook_methods"`

	// CreatedAt is the time when the project was created.
	CreatedAt time.Time `json:"created_at"`
}