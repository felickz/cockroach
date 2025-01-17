// Copyright 2021 The Cockroach Authors.
//
// Licensed as a CockroachDB Enterprise file under the Cockroach Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//     https://github.com/cockroachdb/cockroach/blob/master/licenses/CCL.txt

package kvfeed

import (
	"github.com/cockroachdb/cockroach/pkg/kv"
	"github.com/cockroachdb/cockroach/pkg/kv/kvclient/kvcoord"
	"github.com/cockroachdb/cockroach/pkg/roachpb"
)

// TestingKnobs are the testing knobs for kvfeed.
type TestingKnobs struct {
	// BeforeScanRequest is a callback invoked before issuing Scan request.
	BeforeScanRequest func(b *kv.Batch) error
	// OnRangeFeedValue invoked when rangefeed receives a value.
	OnRangeFeedValue func(kv roachpb.KeyValue) error
	// ShouldSkipCheckpoint invoked when rangefed receives a checkpoint.
	// Returns true if checkpoint should be skipped.
	ShouldSkipCheckpoint func(*roachpb.RangeFeedCheckpoint) bool
	// OnRangeFeedStart invoked when rangefeed starts.  It is given
	// the list of SpanTimePairs.
	OnRangeFeedStart func(spans []kvcoord.SpanTimePair)
}

// ModuleTestingKnobs is part of the base.ModuleTestingKnobs interface.
func (*TestingKnobs) ModuleTestingKnobs() {}
