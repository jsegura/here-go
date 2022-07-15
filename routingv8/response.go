package routingv8

import (
	"encoding/json"
	"time"
)

type ErrorCodes []ErrorCode

func (e *ErrorCodes) UnmarshalJSON(b []byte) error {
	errorCodes := make([]int, 0)
	if err := json.Unmarshal(b, &errorCodes); err != nil {
		return err
	}
	codes := make([]ErrorCode, 0, len(errorCodes))
	for _, errorCode := range errorCodes {
		codes = append(codes, ErrorCode(errorCode))
	}
	*e = codes
	return nil
}

type ErrorCode int

// See https://developer.here.com/documentation/matrix-routing-api/8.6.0/api-reference-swagger.html
// for detailed explanations of each error.
const (
	ErrorCodeSuccess            = 0
	ErrorCodeDisconnected       = 1
	ErrorCodeMatchingFailed     = 2
	ErrorCodeParameterViolation = 3
	ErrorCodeUnknown            = 99
)

// MatrixResponse contains the calculated route matrix.
type MatrixResponse struct {
	// NumOrigins used to calculate matrix
	NumOrigins int `json:"numOrigins"`
	// NumDestinations used to calculate matrix
	NumDestinations int `json:"numDestinations"`
	// TravelTimes calculated using origins and destinations. Nil if not requested in MatrixAttributes.
	TravelTimes []int32 `json:"travelTimes"`
	// Distances calculated using origins and destinations. Nil if not requested in MatrixAttributes.
	Distances []int32 `json:"distances"`
	// ErrorCodes contains potential route errors. Nil if no errors occurred.
	ErrorCodes ErrorCodes `json:"errorCodes"`
}

// CalculateMatrixResponse is used to provide results of a matrix calculation.
type CalculateMatrixResponse struct {
	// MatrixID is unique identifier of the matrix
	MatrixID string `json:"matrixId"`
	// Matrix contains the calculated matrix data.
	Matrix MatrixResponse `json:"matrix"`
	// RegionDefinition to be used to calculate matrix.
	RegionDefinition RegionDefinition `json:"regionDefinition"`
}

// RoutesResponse contains the possible routes.
type RoutesResponse struct {
	// Routes in the possible routes between the origin and target.
	Routes []Route `json:"routes"`
	// ErrorCodes contains potential route errors. Nil if no errors occurred.
	ErrorCodes ErrorCodes `json:"errorCodes"`
}

// Route contains all the sections of a route.
type Route struct {
	// Id of the route
	ID string `json:"id"`
	// Sections in the route
	Sections []Section `json:"sections"`
}

// Section with the information of the departure, arrival location and summary.
type Section struct {
	// Id of the section
	ID string `json:"id"`
	// The type used in the section
	Type string `json:"type"`
	// Departure is the location of the departure
	Departure RoutePlace `json:"departure"`
	// Arrival is the location of the arrival
	Arrival RoutePlace `json:"arrival"`
	// Summary contain info on the duration and length of section
	Summary       Summary      `json:"summary"`
	TravelSummary Summary      `json:"travelSummary"`
	Polyline      string       `json:"polyline"`
	Spans         []Span       `json:"spans"`
	Notices       []Notice     `json:"notices"`
	Language      string       `json:"language"`
	Transport     Transport    `json:"transport"`
	Incidents     []Incident   `json:"incidents"`
	Tolls         []Toll       `json:"tolls"`
	TollSystems   []TollSystem `json:"tollSystems"`
}

type Toll struct {
	CountryCode             string                   `json:"countryCode"`
	TollSystemRef           int                      `json:"tollSystemRef"`
	TollSystem              string                   `json:"tollSystem"`
	Fares                   []Fare                   `json:"fares"`
	TollCollectionLocations []TollCollectionLocation `json:"tollCollectionLocations"`
}
type Fare struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Price          Price    `json:"price"`
	ConvertedPrice Price    `json:"convertedPrice"`
	Reason         string   `json:"reason"`
	PaymentMethods []string `json:"paymentMethods"`
}
type Price struct {
	Type     string  `json:"type"`
	Currency string  `json:"currency"`
	Value    float64 `json:"value"`
}
type TollCollectionLocation struct {
	Name     string      `json:"name"`
	Location GeoWaypoint `json:"location"`
}
type TollSystem struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	LanguageCode string `json:"languageCode"`
}

type Notice struct {
	Title    string         `json:"title"`
	Code     string         `json:"code"`
	Severity string         `json:"severity"`
	Details  []NoticeDetail `jsopn:"details"`
}

type Transport struct {
	Mode string `json:"mode"`
}
type NoticeDetail struct {
	Type           string `json:"type"`
	Cause          string `json:"cause"`
	MaxGrossWeight int    `json:"maxGrossWeight"`
}

type Incident struct {
	Type        string    `json:"type"`
	Criticality string    `json:"criticality"`
	ValidFrom   time.Time `json:"validFrom"`
	ValidUntil  time.Time `json:"validUntil"`
	Description string    `json:"description"`
}

type RoutePlace struct {
	Time  time.Time `json:"time"`
	Place Place     `json:"place"`
}

type Span struct {
	Offset     int     `json:"offset"`
	Length     int     `json:"length"`
	Duration   int     `json:"duration"`
	SpeedLimit float64 `json:"speedLimit,omitempty"`
	// MaxSpeed   float64 `json:"maxSpeed,omitempty"`
	// Incidents  []int   `json:"incidents,omitempty"`
	// Notices    []int   `json:"notices,omitempty"`
}

// Place with lat and long info on where the place is.
type Place struct {
	// Type is the struct
	Type string `json:"type"`
	// Location in lat and long
	Location GeoWaypoint `json:"location"`
	// OriginalLocation in lat and long
	OriginalLocation GeoWaypoint `json:"originalLocation"`
}

// Summary contains the duration and length info.
type Summary struct {
	// Duration is the total duration of the action, section etc
	Duration int32 `json:"duration"`
	// Length is the total length
	Length int32 `json:"length"`
	// BaseDuration is the duration without dynamic traffic information
	BaseDuration int32 `json:"baseDuration"`
}

// HereErrorResponse is returned when an error is returned from the Here Maps API.
type HereErrorResponse struct {
	// Title of the error
	Title string `json:"title"`
	// Http status code
	Status int `json:"status"`
	// Here Maps API error code
	Code string `json:"code"`
	// Cause of the error
	Cause string `json:"cause"`
	// Action Suggested to fix error
	Action string `json:"action"`
}
