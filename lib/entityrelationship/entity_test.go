// Copyright (c) Outernet Council and Contributors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package entityrelationship_test

import (
	"testing"

	"github.com/GabrielFrota/nmts/lib/entityrelationship"
	npb "github.com/GabrielFrota/nmts/proto"
	"google.golang.org/protobuf/encoding/prototext"
)

func TestEntityKindStringExamples(t *testing.T) {
	text := `ek_sdn_agent{}`

	parsed := new(npb.Entity)
	err := prototext.Unmarshal([]byte(text), parsed)
	if err != nil {
		t.Fatalf("failed to parse %s: %q", text, err)
	}
	wanted := "EK_SDN_AGENT"
	got := entityrelationship.EntityKindStringFromProto(parsed)
	if got != wanted {
		t.Fatalf("wanted: %q, got: %q", wanted, got)
	}
}
