// Copyright 2022-2026, Offchain Labs, Inc.
// For license information, see https://github.com/OffchainLabs/nitro/blob/master/LICENSE.md

package anytrust

import (
	"time"

	"github.com/tenderly/net-nitro/go-ethereum/log"

	anytrustutil "github.com/tenderly/net-nitro/daprovider/anytrust/util"
	"github.com/tenderly/net-nitro/util/pretty"
)

func logPut(store string, data []byte, timeout uint64, reader anytrustutil.Reader, more ...interface{}) {
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
