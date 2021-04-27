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

// IsValidRegion checks if r is an Ali supported region.
func IsValidRegion(r string) bool {
	for _, v := range ValidRegions {
		if r == string(v) {
			return true
		}
	}
	return false
}
