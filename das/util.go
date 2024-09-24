// Copyright 2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE

package das

import (
	"time"

	"github.com/tenderly/net-nitro/go-ethereum/log"
	"github.com/tenderly/net-nitro/arbstate/daprovider"
	"github.com/tenderly/net-nitro/util/pretty"
)

func logPut(store string, data []byte, timeout uint64, reader daprovider.DASReader, more ...interface{}) {
	if len(more) == 0 {
		// #nosec G115
		log.Trace(
			store, "message", pretty.FirstFewBytes(data), "timeout", time.Unix(int64(timeout), 0),
			"this", reader,
		)
	} else {
		// #nosec G115
		log.Trace(
			store, "message", pretty.FirstFewBytes(data), "timeout", time.Unix(int64(timeout), 0),
			"this", reader, more,
		)
	}
}
