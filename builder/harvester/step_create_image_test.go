// Copyright IBM Corp. 2020, 2026
// SPDX-License-Identifier: MPL-2.0

package harvester

import (
	"reflect"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
)

func TestOutputImageStorageClassParameters(t *testing.T) {
	srcParams := map[string]string{"diskSelector": "vm-data", "numberOfReplicas": "1"}

	t.Run("nil_without_config_or_source", func(t *testing.T) {
		state := new(multistep.BasicStateBag)
		if got := outputImageStorageClassParameters(&Config{}, state); got != nil {
			t.Fatalf("expected nil, got %v", got)
		}
	})

	t.Run("copies_source_image_parameters", func(t *testing.T) {
		state := new(multistep.BasicStateBag)
		state.Put(sourceImageSCParamsStateKey, srcParams)

		got := outputImageStorageClassParameters(&Config{}, state)
		if !reflect.DeepEqual(got, srcParams) {
			t.Fatalf("expected %v, got %v", srcParams, got)
		}
		got["diskSelector"] = "changed"
		if srcParams["diskSelector"] != "vm-data" {
			t.Fatal("expected a copy, source parameters were modified")
		}
	})

	t.Run("config_overrides_source", func(t *testing.T) {
		state := new(multistep.BasicStateBag)
		state.Put(sourceImageSCParamsStateKey, srcParams)
		cfg := &Config{OutputImageStorageClassParameters: map[string]string{"numberOfReplicas": "3"}}

		got := outputImageStorageClassParameters(cfg, state)
		if want := map[string]string{"numberOfReplicas": "3"}; !reflect.DeepEqual(got, want) {
			t.Fatalf("expected %v, got %v", want, got)
		}
	})

	t.Run("empty_source_parameters", func(t *testing.T) {
		state := new(multistep.BasicStateBag)
		state.Put(sourceImageSCParamsStateKey, map[string]string(nil))
		if got := outputImageStorageClassParameters(&Config{}, state); got != nil {
			t.Fatalf("expected nil, got %v", got)
		}
	})
}
