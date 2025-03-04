package routingv8

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// Routes returns all possible routes between origin and destination.
// See https://developer.here.com/documentation/routing-api/dev_guide/topics/send-request.html#send-a-request
// for details about other parameters.
func (s *RoutingService) Routes(
	ctx context.Context,
	req *RoutesRequest,
) (_ *RoutesResponse, err error) {
	tm := req.TransportMode.String()
	if tm == invalid || tm == unspecified {
		return nil, fmt.Errorf("invalid transportmode")
	}

	u, err := s.URL.Parse("routes")
	if err != nil {
		return nil, err
	}

	values := make(url.Values)
	values.Add("return", "summary,polyline,elevation,actions,instructions,travelSummary,tolls,incidents")
	values.Add("transportMode", tm)
	values.Add("origin", fmt.Sprintf("%v,%v", req.Origin.Lat, req.Origin.Long))
	values.Add("destination", fmt.Sprintf("%v,%v", req.Destination.Lat, req.Destination.Long))
	values.Add("spans", "length,duration,maxSpeed,speedLimit,incidents,notices")
	values.Add("alternatives", "6")
	values.Add("currency", "EUR")

	r, err := s.Client.NewRequest(ctx, u, http.MethodGet, values.Encode(), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create get request: %v", err)
	}
	var resp RoutesResponse
	if err := s.Client.Do(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
