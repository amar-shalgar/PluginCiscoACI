//(C) Copyright [2020] Hewlett Packard Enterprise Development LP
//
//Licensed under the Apache License, Version 2.0 (the "License"); you may
//not use this file except in compliance with the License. You may obtain
//a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//License for the specific language governing permissions and limitations
// under the License.

//Package capmodel ...
package capmodel

//FabricHelath ...
type FabricHelath struct {
	TotalCount string               `json:"totalCount"`
	IMData     []FabricHelathIMData `json:"imdata"`
}

// FabricHelathIMData ...
type FabricHelathIMData struct {
	FabricHelathData FabricHelathData `json:"fabricHealthTotal"`
}

// FabricHelathData ...
type FabricHelathData struct {
	Attributes map[string]interface{} `json:"attributes"`
}
