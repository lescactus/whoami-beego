package models

import (
	"net/url"
)

// Browser is a structure containing browser infos
type Browser struct {
	IP        string
	Port      int
	Host      string
	Headers   map[string][]string
	URL       *url.URL
	Lang      string
	UserAgent string
	Location  *Location
	JSON      string
	YAML      string
}

// Location provide browser Geo IP location informations
type Location struct {
	IP          string  `json:"ip"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
	RegionCode  string  `json:"region_code"`
	RegionName  string  `json:"region_name"`
	City        string  `json:"city"`
	ZipCode     string  `json:"zip_code"`
	Timezone    string  `json:"time_zone"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	MetroCode   int     `json:"metro_code"`
}
