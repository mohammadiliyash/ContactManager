package models

import (
	"strings"

	"github.com/miliyash/ms-contact-manager/app/utils"
)

//PostalCode struct
type PostalCode struct {
	utils.ServiceResult
	ID               string  `json:"pid"`
	Value            string  `json:"value"`
	City             *string `json:"city" logging:"-"`
	State            *string `json:"state" logging:"-"`
	CityAbbreviation *string `json:"cityAbbreviation,omitempty"`
}

// PostalCodeDict dic object for json data
var PostalCodeDict map[string]*PostalCode

// init is creating a dictionary which helps in finding elements at good speed
// we can take a chance at this time as it is one time job and once data is loaded in memory
// Reading data will be fast -- This is a sample approach , this can be done in better way
func init() {
	PostalCodeDict = make(map[string]*PostalCode)

	for _, r := range dataArray {

		pc, found := PostalCodeDict[r.Zip]

		if found {
			updatePostalCode(&r, pc)
		} else {
			PostalCodeDict[r.Zip] = createPostalCode(&r)
		}
	}
}

func createPostalCode(r *GeoDataRecord) *PostalCode {

	var pcc *string
	var pcs *string
	var pcca *string
	if len(r.LL) > 0 {
		city := titleCase(r.City)
		state := strings.ToUpper(r.State)
		if len(r.CityAbb) > 0 {
			cabb := titleCase(r.CityAbb)
			pcca = &cabb
		}
		pcc = &city
		pcs = &state
	}

	rtn := PostalCode{
		Value:            r.Zip,
		City:             pcc,
		CityAbbreviation: pcca,
		State:            pcs,
	}
	return &rtn
}

func updatePostalCode(r *GeoDataRecord, pc *PostalCode) {

	if len(r.LL) > 0 {
		city := titleCase(r.City)
		state := strings.ToUpper(r.State)
		var pcca *string
		if len(r.CityAbb) > 0 {
			cabb := titleCase(r.CityAbb)
			pcca = &cabb
		}
		pc.City = &city
		pc.State = &state
		pc.CityAbbreviation = pcca
	}
}

func titleCase(t string) string {
	return strings.Title(strings.ToLower(t))
}
