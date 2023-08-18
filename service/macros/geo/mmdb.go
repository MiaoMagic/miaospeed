package geo

import (
	"github.com/moshaoli688/miaospeed/interfaces"
	"github.com/moshaoli688/miaospeed/utils"
)

func RunMMDBCheck(rawIp string) *interfaces.GeoInfo {
	if record := utils.QueryMaxMindDB(rawIp); record != nil {
		return &interfaces.GeoInfo{
			ASN:    record.ASN,
			ASNOrg: record.ASNOrg,
			Org:    record.ASNOrg,
			ISP:    record.ASNOrg, // inaccurate, just fallback
			IP:     rawIp,

			Country:       record.Country.Names.EN,
			CountryCode:   record.Country.ISOCode,
			ContinentCode: record.Continent.Code,
			TimeZone:      record.Location.TimeZone,
			Lat:           record.Location.Latitude,
			Lon:           record.Location.Longitude,
		}
	}

	return nil
}
