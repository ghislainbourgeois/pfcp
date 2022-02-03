// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package pfcpType

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalVolumeMeasurement(t *testing.T) {
	testData := VolumeMeasurement{
		Dlvol:          true,
		Ulvol:          true,
		Tovol:          true,
		TotalVolume:    987654321987654321,
		UplinkVolume:   123456789123456789,
		DownlinkVolume: 864197532864197532,
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	assert.Equal(t, []byte{7, 13, 180, 218, 95, 126, 244, 18, 177, 1, 182, 155, 75, 172, 208, 95, 21, 11, 254, 63, 19, 210, 35, 179, 156}, buf)
}

func TestUnmarshalVolumeMeasurement(t *testing.T) {
	buf := []byte{7, 13, 180, 218, 95, 126, 244, 18, 177, 1, 182, 155, 75, 172, 208, 95, 21, 11, 254, 63, 19, 210, 35, 179, 156}
	var testData VolumeMeasurement
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := VolumeMeasurement{
		Dlvol:          true,
		Ulvol:          true,
		Tovol:          true,
		TotalVolume:    987654321987654321,
		UplinkVolume:   123456789123456789,
		DownlinkVolume: 864197532864197532,
	}
	assert.Equal(t, expectData, testData)
}
