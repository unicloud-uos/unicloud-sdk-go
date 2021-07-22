package common

// Region represents ECS region
type Region string

// Constants of region definition
const (
	Beijing3 = Region("HB1-BJMYB")
)

var ValidRegions = []Region{
	Beijing3,
}

var RegionEndpoint = map[Region]string{
	Beijing3: "api.unicloud.com",
}

// IsValidRegion checks if r is an Ali supported region.
func IsValidRegion(r string) bool {
	for _, v := range ValidRegions {
		if r == string(v) {
			return true
		}
	}
	return false
}

func GetRegionEndpoint(region string) string {
	return RegionEndpoint[Region(region)]
}
