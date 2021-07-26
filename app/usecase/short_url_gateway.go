//go:generate mockgen -package gatewaymock -destination gatewaymock/short_url_gateway.go . ShortURLGateway
package usecase

type ShortURLGateway interface {
	CreateShortURL(string) (string, error)
}
