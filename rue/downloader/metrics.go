// Copyright 2015 The go-ruereum Authors
// This file is part of the go-ruereum library.
//
// The go-ruereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ruereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ruereum library. If not, see <http://www.gnu.org/licenses/>.

// Contains the metrics collected by the downloader.

package downloader

import (
	"github.com/Rue-Foundation/go-rue/metrics"
)

var (
	headerInMeter      = metrics.NewMeter("rue/downloader/headers/in")
	headerReqTimer     = metrics.NewTimer("rue/downloader/headers/req")
	headerDropMeter    = metrics.NewMeter("rue/downloader/headers/drop")
	headerTimeoutMeter = metrics.NewMeter("rue/downloader/headers/timeout")

	bodyInMeter      = metrics.NewMeter("rue/downloader/bodies/in")
	bodyReqTimer     = metrics.NewTimer("rue/downloader/bodies/req")
	bodyDropMeter    = metrics.NewMeter("rue/downloader/bodies/drop")
	bodyTimeoutMeter = metrics.NewMeter("rue/downloader/bodies/timeout")

	receiptInMeter      = metrics.NewMeter("rue/downloader/receipts/in")
	receiptReqTimer     = metrics.NewTimer("rue/downloader/receipts/req")
	receiptDropMeter    = metrics.NewMeter("rue/downloader/receipts/drop")
	receiptTimeoutMeter = metrics.NewMeter("rue/downloader/receipts/timeout")

	stateInMeter   = metrics.NewMeter("rue/downloader/states/in")
	stateDropMeter = metrics.NewMeter("rue/downloader/states/drop")
)
