package api

// import (
// 	"encoding/json"
// 	"net/http"
// )

// Coin Balance Params
type CoinBalanceParams struct {
	Username string
}

// Coin Balance Response
type CoinBalanceResponse struct {
	// Success code, usually 200
	Code int

	// Account balance
	Balance int64
}

// Error Response
type Error struct {
	// Error code
	Code int

	// Error message
	Message string
}