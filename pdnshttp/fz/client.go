/*
Copyright © 2021 Michael Bruskov <mixanemca@yandex.ru>

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

package fz

import "github.com/mixanemca/dnscli/pdnshttp"

type client struct {
	httpClient *pdnshttp.PDNSClient
}

// New creates a new ForwardZone client
func New(hc *pdnshttp.PDNSClient) Client {
	return &client{
		httpClient: hc,
	}
}
