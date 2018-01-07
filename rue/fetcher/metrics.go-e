// Copyright 2015 The go-rueereum Authors
// This file is part of the go-rueereum library.
//
// The go-rueereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-rueereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-rueereum library. If not, see <http://www.gnu.org/licenses/>.

// Contains the metrics collected by the fetcher.

package fetcher

import (
	"github.com/Rue-Foundation/go-rue/metrics"
)

var (
	propAnnounceInMeter   = metrics.NewMeter("rue/fetcher/prop/announces/in")
	propAnnounceOutTimer  = metrics.NewTimer("rue/fetcher/prop/announces/out")
	propAnnounceDropMeter = metrics.NewMeter("rue/fetcher/prop/announces/drop")
	propAnnounceDOSMeter  = metrics.NewMeter("rue/fetcher/prop/announces/dos")

	propBroadcastInMeter   = metrics.NewMeter("rue/fetcher/prop/broadcasts/in")
	propBroadcastOutTimer  = metrics.NewTimer("rue/fetcher/prop/broadcasts/out")
	propBroadcastDropMeter = metrics.NewMeter("rue/fetcher/prop/broadcasts/drop")
	propBroadcastDOSMeter  = metrics.NewMeter("rue/fetcher/prop/broadcasts/dos")

	headerFetchMeter = metrics.NewMeter("rue/fetcher/fetch/headers")
	bodyFetchMeter   = metrics.NewMeter("rue/fetcher/fetch/bodies")

	headerFilterInMeter  = metrics.NewMeter("rue/fetcher/filter/headers/in")
	headerFilterOutMeter = metrics.NewMeter("rue/fetcher/filter/headers/out")
	bodyFilterInMeter    = metrics.NewMeter("rue/fetcher/filter/bodies/in")
	bodyFilterOutMeter   = metrics.NewMeter("rue/fetcher/filter/bodies/out")
)
