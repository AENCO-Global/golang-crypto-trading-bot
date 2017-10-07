package exchangeWrappers
import (
	"github.com/Alexbrem/golang-crypto-trading-bot/environment"

	yobitAPI "github.com/Alexbrem/golang-crypto-trading-bot/exchangeWrappers/yobitapiwrap"
)

// YobitWrapper provides a Generic wrapper of the Yobit API.
type YobitWrapper struct {
	yobitAPI *yobitAPI.Yobit //Represents the helper of the Yobit API.
}

// NewYobitWrapper creates a generic wrapper of the Yobit API.
func NewYobitWrapper(apiKey string, apiSecret string) ExchangeWrapper {
	return YobitWrapper{
		yobitAPI: yobitAPI.New(apiKey, apiSecret),
	}
}
