package plugs

import (
	"strings"
)

var plug_maps = make(map[string]Plug)

type Plug struct {
	Plug_type string
	Volts     string
	Hertz     string
}

// This will display the following properties of a specified country:
// Plug_type - string
// Volts - string
// Hertz - string
func Display(country string) Plug {
	return plug_maps[country]
}

func plugType(country string) []string {
	return strings.Split(plug_maps[country].Plug_type, "/")
}

// DoINeedAConverter will take two countries as parameters, countryFrom and CountryTo
// It will then check if the two countries have any similar plug types, and if so
// It will return "No", otherwise, it will return "Yes" as the two countries don't have similar plug types.
func DoINeedAConverter(countryFrom string, countryTo string) string {
	plugTypeCountryFrom := plugType(countryFrom)
	plugTypeCountryTo := plugType(countryTo)
	for plug := range plugTypeCountryFrom {
		for plugTo := range plugTypeCountryTo {
			if plugTypeCountryFrom[plug] == plugTypeCountryTo[plugTo] {
				return "No"
			}
		}
	}
	return "Yes"
}

func init() {
	stringOfCsv :=
		`Abu Dhabi 	G	230 V	50 Hz
	Afghanistan	C / F	220 V	50 Hz
	Albania	C / F	230 V	50 Hz
	Algeria	C / F	230 V	50 Hz
	American Samoa	A / B / F / I	120 V	60 Hz
	Andorra	C / F	230 V	50 Hz
	Angola	C / F	220 V	50 Hz
	Anguilla	A / B	110 V	60 Hz
	Antigua and Barbuda	A / B	230 V	60 Hz
	Argentina	C / I	220 V	50 Hz
	Armenia	C / F	230 V	50 Hz
	Aruba	A / B / F	120 V	60 Hz
	Australia	I	230 V 	50 Hz
	Austria	C / F	230 V	50 Hz
	Azerbaijan	C / F	220 V	50 Hz
	Azores	A / B / C / F	230 V	50 Hz
	Bahamas	A / B	120 V	60 Hz
	Bahrain	G	230 V	50 Hz
	Balearic Islands	C / F	230 V	50 Hz
	Bangladesh	A / C / D / G	220 V	50 Hz
	Barbados	A / B	115 V	50 Hz
	Belarus	C / F	220 V	50 Hz
	Belgium	C / E	230 V	50 Hz
	Belize	A / B / G	110 V / 220 V	60 Hz
	Benin	C / E	220 V	50 Hz
	Bermuda	A / B	120 V	60 Hz
	Bhutan	C / D / G	230 V	50 Hz
	Bolivia	A / C	230 V	50 Hz
	Bonaire	A / C	127 V	50 Hz
	Bosnia & Herzegovina	C / F	230 V	50 Hz
	Botswana	D / G	230 V	50 Hz
	Brazil	C / N	127 V / 220 V	60 Hz
	British Virgin Islands	A / B	110 V	60 Hz
	Brunei	G	240 V	50 Hz
	Bulgaria	C / F	230 V	50 Hz
	Burkina Faso	C / E	220 V	50 Hz
	Burma	A / C / D / G / I	230 V	50 Hz
	Burundi	C / E	220 V	50 Hz
	Cambodia	A / C / G	230 V	50 Hz
	Cameroon	C / E	220 V	50 Hz
	Canada	A / B	120 V	60 Hz
	Canary Islands	C / E / F	230 V	50 Hz
	Cape Verde	C / F	230 V	50 Hz
	Cayman Islands	A / B	120 V	60 Hz
	Central African Republic	C / E	220 V	50 Hz
	Chad	C / E / F	220 V	50 Hz
	Channel Islands	C / G	230 V	50 Hz
	Chile	C / L	220 V	50 Hz
	China	A / C / I	220 V	50 Hz
	Christmas Island	I	230 V	50 Hz
	Cocos (Keeling) Islands	I	230 V	50 Hz
	Colombia	A / B	110 V	60 Hz
	Comoros	C / E	220 V	50 Hz
	Congo-Brazzaville	C / E	230 V	50 Hz
	Congo-Kinshasa	C / E	220 V	50 Hz
	Cook Islands	I	240 V	50 Hz
	Costa Rica	A / B	120 V	60 Hz
	Ivory Coast	C / E	220 V	50 Hz
	Croatia	C / F	230 V	50 Hz
	Cuba	A / B / C / L	110 V / 220 V	60 Hz
	Curaçao	A / B	127 V	50 Hz
	Cyprus	G	230 V	50 Hz
	North Cyprus  	G	230 V	50 Hz
	Czech Republic	C / E	230 V	50 Hz
	Denmark	C / E / F / K	230 V	50 Hz
	Djibouti	C / E	220 V	50 Hz
	Dominica	D / G	230 V	50 Hz
	Dominican Republic	A / B / C	120 V	60 Hz
	Dubai 	G	230 V	50 Hz
	East Timor	C / E / F / I	220 V	50 Hz
	Ecuador	A / B	120 V	60 Hz
	Egypt	C / F	220 V	50 Hz
	El Salvador	A / B	120 V	60 Hz
	England	G	230 V	50 Hz
	Equatorial Guinea	C / E	220 V	50 Hz
	Eritrea	C / L	230 V	50 Hz
	Estonia	C / F	230 V	50 Hz
	Ethiopia	C / F / G	220 V	50 Hz
	Faeroe Islands	C / E / F / K	230 V	50 Hz
	Falkland Islands	G	240 V	50 Hz
	Fiji	I	240 V	50 Hz
	Finland	C / F	230 V	50 Hz
	France	C / E	230 V	50 Hz
	French Guiana	C / E	230 V	50 Hz
	French Polynesia	C / E	220 V	60 Hz
	Gabon	C / E	220 V	50 Hz
	Gambia	G	230 V	50 Hz
	Gaza Strip (Gaza)	C / H	230 V	50 Hz
	Georgia	C / F	220 V	50 Hz
	Germany	C / F	230 V	50 Hz
	Ghana	D / G	230 V	50 Hz
	Gibraltar	G	230 V	50 Hz
	Greece	C / F	230 V	50 Hz
	Greenland	C / E / F / K	230 V	50 Hz
	Grenada	G	230 V	50 Hz
	Guadeloupe	C / E	230 V	50 Hz
	Guam	A / B	110 V	60 Hz
	Guatemala	A / B	120 V	60 Hz
	Guinea	C / F	220 V	50 Hz
	Guinea-Bissau	C / E / F	220 V	50 Hz
	Guyana	A / B / D / G	120 V / 240 V	60 Hz
	Haiti	A / B	110 V	60 Hz
	Holland	C / F	230 V	50 Hz
	Honduras	A / B	120 V	60 Hz
	Hong Kong	G	220 V	50 Hz
	Hungary	C / F	230 V	50 Hz
	Iceland	C / F	230 V	50 Hz
	India	C / D / M	230 V	50 Hz
	Indonesia	C / F	230 V	50 Hz
	Iran	C / F	230 V	50 Hz
	Iraq	C / D / G	230 V	50 Hz
	Ireland	G	230 V	50 Hz
	Northern Ireland 	G	230 V	50 Hz
	Isle of Man	C / G	230 V	50 Hz
	Israel	C / H	230 V	50 Hz
	Italy	C / F / L	230 V	50 Hz
	Jamaica	A / B	110 V	50 Hz
	Japan	A / B	100 V	50 Hz / 60 Hz
	Jordan	C / D / F / G / J	230 V	50 Hz
	Kazakhstan	C / F	220 V	50 Hz
	Kenya	G	240 V	50 Hz
	Kiribati	I	240 V	50 Hz
	North Korea 	C / F	220 V	50 Hz
	South Korea 	C / F	220 V	60 Hz
	Kosovo	C / F	230 V	50 Hz
	Kuwait	G	240 V	50 Hz
	Kyrgyzstan	C / F	220 V	50 Hz
	Laos	A / B / C / E / F	230 V	50 Hz
	Latvia	C / F	230 V	50 Hz
	Lebanon	C / D / G	230 V	50 Hz
	Lesotho	M	220 V	50 Hz
	Liberia	A / B / C / F	120 V / 220 V	60 Hz
	Libya	C / L	230 V	50 Hz
	Liechtenstein	C / J	230 V	50 Hz
	Lithuania	C / F	230 V	50 Hz
	Luxembourg	C / F	230 V	50 Hz
	Macau	G	220 V	50 Hz
	North Macedonia	C / F	230 V	50 Hz
	Madagascar	C / E	220 V	50 Hz
	Madeira	C / F	230 V	50 Hz
	Malawi	G	230 V	50 Hz
	Malaysia	G	230 V 	50 Hz
	Maldives	C / D / G / L	230 V	50 Hz
	Mali	C / E	220 V	50 Hz
	Malta	G	230 V	50 Hz
	Marshall Islands	A / B	120 V	60 Hz
	Martinique	C / E	230 V	50 Hz
	Mauritania	C / E / F	220 V	50 Hz
	Mauritius	C / G	230 V	50 Hz
	Mayotte	C / E	230 V	50 Hz
	Mexico	A / B	127 V	60 Hz
	Micronesia	A / B	120 V	60 Hz
	Moldova	C / F	230 V	50 Hz
	Monaco	C / E / F	230 V	50 Hz
	Mongolia	C / F 	230 V	50 Hz
	Montenegro	C / F	230 V	50 Hz
	Montserrat	A / B	230 V	60 Hz
	Morocco	C / E	220 V	50 Hz
	Mozambique	C / F / M	220 V	50 Hz
	Myanmar	A / C / D / G / I	230 V	50 Hz
	Namibia	D / M	220 V	50 Hz
	Nauru	I	240 V	50 Hz
	Nepal	C / D / M	230 V	50 Hz
	Netherlands	C / F	230 V	50 Hz
	New Caledonia	C / E	220 V	50 Hz
	New Zealand	I	230 V	50 Hz
	Nicaragua	A / B	120 V	60 Hz
	Niger	C / D / E	220 V	50 Hz
	Nigeria	D / G	230 V	50 Hz
	Niue	I	230 V	50 Hz
	Norfolk Island	I	230 V	50 Hz
	North Cyprus	G	230 V	50 Hz
	North Korea	C / F	220 V	50 Hz
	North Macedonia	C / F	230 V	50 Hz
	Northern Ireland	G	230 V	50 Hz
	Norway	C / F	230 V	50 Hz
	Oman	G	240 V	50 Hz
	Pakistan	C / D	230 V	50 Hz
	Palau	A / B	120 V	60 Hz
	Palestine	C / H	230 V	50 Hz
	Panama	A / B	120 V	60 Hz
	Papua New Guinea	I	240 V	50 Hz
	Paraguay	A / C	220 V	50 Hz
	Peru	A / B / C	220 V	60 Hz
	Philippines	A / B / C	220 V	60 Hz
	Pitcairn Islands	I	230 V	50 Hz
	Poland	C / E	230 V	50 Hz
	Portugal	C / F	230 V	50 Hz
	Puerto Rico	A / B	120 V	60 Hz
	Qatar	G	240 V	50 Hz
	Réunion (French overseas department)	C / E	230 V	50 Hz
	Romania	C / F	230 V	50 Hz
	Russia (officially the Russian Federation)	C / F	220 V	50 Hz
	Rwanda	C / E / F / G	230 V	50 Hz
	Saba	A / B	110 V	60 Hz
	Saint Barthélemy	C / E	230 V	60 Hz
	Saint Helena	G	230 V	50 Hz
	Saint Kitts and Nevis	D / G	230 V	60 Hz
	Saint Lucia	G	230 V	50 Hz
	Saint Martin	C / E	220 V	60 Hz
	Saint Pierre and Miquelon	C / E	230 V	50 Hz
	Saint Vincent and the Grenadines	A / B / G	110 V / 230 V	50 Hz
	Samoa	I	230 V	50 Hz
	San Marino	C / F / L	230 V	50 Hz
	São Tomé and Príncipe	C / F	230 V	50 Hz
	Saudi Arabia	G	220 V	60 Hz
	Scotland	G	230 V	50 Hz
	Senegal	C / D / E	230 V	50 Hz
	Serbia	C / F	230 V	50 Hz
	Seychelles	G	240 V	50 Hz
	Sierra Leone	D / G	230 V	50 Hz
	Singapore	G	230 V	50 Hz
	Sint Eustatius	A / B / C / F	110 V / 220 V	60 Hz
	Sint Maarten	A / B	110 V	60 Hz
	Slovakia	C / E	230 V	50 Hz
	Slovenia	C / F	230 V	50 Hz
	Solomon Islands	G / I	230 V	50 Hz
	Somalia	G	220 V	50 Hz
	Somaliland	G	220 V	50 Hz
	South Africa	C / M / N	230 V	50 Hz
	South Korea	C / F	220 V	60 Hz
	South Sudan	C / D	230 V	50 Hz
	Spain	C / F	230 V	50 Hz
	Sri Lanka	G	230 V	50 Hz
	Sudan	C / D	230 V	50 Hz
	Suriname	A / B / C / F	127 V / 220 V	60 Hz
	Swaziland	M	230 V	50 Hz
	Sweden	C / F	230 V	50 Hz
	Switzerland	C / J	230 V	50 Hz
	Syria	C / E / L	220 V	50 Hz
	Tahiti	C / E	220 V	60 Hz
	Taiwan	A / B	110 V	60 Hz
	Tajikistan	C / F	220 V	50 Hz
	Tanzania	D / G	230 V	50 Hz
	Thailand	A / B / C / O	230 V	50 Hz
	Togo	C / E	220 V	50 Hz
	Tokelau	I	230 V	50 Hz
	Tonga	I	240 V	50 Hz
	Trinidad & Tobago	A / B	115 V	60 Hz
	Tunisia	C / E	230 V	50 Hz
	Turkey	C / F	230 V	50 Hz
	Turkmenistan	C / F	220 V	50 Hz
	Turks and Caicos Islands	A / B	120 V	60 Hz
	Tuvalu	I	230 V	50 Hz
	Uganda	G	240 V	50 Hz
	Ukraine	C / F	230 V	50 Hz
	UAE	G	230 V	50 Hz
	UK	G	230 V	50 Hz
	USA	A / B	120 V	60 Hz
	United States Virgin Islands	A / B	110 V	60 Hz
	Uruguay	C / F / L	220 V	50 Hz
	Uzbekistan	C / F	220 V	50 Hz
	Vanuatu	I	230 V	50 Hz
	Vatican City	C / F / L	230 V	50 Hz
	Venezuela	A / B	120 V	60 Hz
	Vietnam	A / B / C	220 V	50 Hz
	Virgin Islands	A / B	110 V	60 Hz
	Wales	G	230 V	50 Hz
	Wallis and Futuna	C / E	220 V	50 Hz
	West Bank	C / H	230 V	50 Hz
	Western Sahara	C / E	220 V	50 Hz
	Yemen	A / D / G	230 V	50 Hz
	Zambia	C / D / G	230 V	50 Hz
	Zimbabwe	D / G	230 V	50 Hz`
	initaliseMapping(stringOfCsv)
}

func initaliseMapping(stringOfCsv string) {

	multiline_arr := strings.Split(stringOfCsv, "\n")

	for i, _ := range multiline_arr {
		line := strings.Split(multiline_arr[i], "\t")
		starting := 0
		name := strings.TrimLeft(line[starting], "\t")
		if name == "" {
			starting += 1
		}
		plug_maps[line[starting]] = Plug{string(line[starting+1]), string(line[starting+2]), string(line[starting+3])}
	}
}
