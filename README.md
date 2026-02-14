# ProxyHat Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/ProxyHatCom/go-sdk.svg)](https://pkg.go.dev/github.com/ProxyHatCom/go-sdk?utm_source=github&utm_medium=readme&utm_campaign=sdk-go)
[![Go Version](https://img.shields.io/github/go-mod/go-version/ProxyHatCom/go-sdk)](https://github.com/ProxyHatCom/go-sdk/blob/main/go.mod)
[![CI](https://github.com/ProxyHatCom/go-sdk/actions/workflows/ci.yml/badge.svg)](https://github.com/ProxyHatCom/go-sdk/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/ProxyHatCom/go-sdk/blob/main/LICENSE)

The official Go client library for the [ProxyHat](https://proxyhat.com?utm_source=github&utm_medium=readme&utm_campaign=sdk-go) API.

## Installation

```bash
go get github.com/ProxyHatCom/go-sdk
```

**Requirements:** Go 1.21+

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"

	proxyhat "github.com/ProxyHatCom/go-sdk"
)

func main() {
	client := proxyhat.NewClient("your-api-key")

	user, err := client.Auth.User(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Hello, %s!\n", user.Name)
}
```

## Usage

### Client Configuration

```go
// Default client
client := proxyhat.NewClient("your-api-key")

// With options
client := proxyhat.NewClient("your-api-key",
	proxyhat.WithBaseURL("https://custom-api.example.com/v1"),
	proxyhat.WithTimeout(60 * time.Second),
	proxyhat.WithHTTPClient(&http.Client{
		Transport: customTransport,
	}),
)
```

### Authentication

```go
// Register a new account
resp, err := client.Auth.Register(ctx, proxyhat.RegisterParams{
	Name:                 "John Doe",
	Email:                "john@example.com",
	Password:             "secure-password",
	PasswordConfirmation: "secure-password",
})

// Login
login, err := client.Auth.Login(ctx, proxyhat.LoginParams{
	Email:    "john@example.com",
	Password: "secure-password",
})

// Get current user
user, err := client.Auth.User(ctx)
```

### Sub-Users

```go
// List all sub-users
users, err := client.SubUsers.List(ctx)

// Create a sub-user
user, err := client.SubUsers.Create(ctx, proxyhat.CreateSubUserParams{
	ProxyPassword: "proxy-pass",
	Name:          proxyhat.String("My Sub-User"),
})

// Update a sub-user
updated, err := client.SubUsers.Update(ctx, "user-id", proxyhat.UpdateSubUserParams{
	Name: proxyhat.String("New Name"),
})

// Bulk delete
resp, err := client.SubUsers.BulkDelete(ctx, []string{"id-1", "id-2"})
```

### Locations

```go
// List countries with filters
countries, err := client.Locations.Countries(ctx, &proxyhat.LocationParams{
	Limit: proxyhat.Int(10),
	Name:  proxyhat.String("United"),
})

// List cities
cities, err := client.Locations.Cities(ctx, &proxyhat.CityParams{
	RegionParams: proxyhat.RegionParams{
		CountryCode: proxyhat.String("US"),
	},
})
```

### Analytics

```go
// Get traffic data
traffic, err := client.Analytics.Traffic(ctx, &proxyhat.AnalyticsParams{
	Period: "7d",
})

// Get total traffic
total, err := client.Analytics.TrafficTotal(ctx, nil) // defaults to 24h
```

### Payments

```go
// Create a payment
payment, err := client.Payments.Create(ctx, proxyhat.CreatePaymentParams{
	Type:               "regular",
	PlanID:             "plan-id",
	CryptocurrencyCode: "BTC",
})

// Download invoice
resp, err := client.Payments.Invoice(ctx, "payment-id", "pdf")
if err != nil {
	log.Fatal(err)
}
defer resp.Body.Close()
// Write resp.Body to file...
```

### Error Handling

```go
user, err := client.Auth.User(ctx)
if err != nil {
	if proxyhat.IsAuthenticationError(err) {
		log.Fatal("Invalid API key")
	}
	if proxyhat.IsRateLimitError(err) {
		if rle, ok := proxyhat.AsRateLimitError(err); ok {
			log.Printf("Rate limited. Retry after %d seconds", rle.RetryAfter)
		}
	}
	if proxyhat.IsNotFoundError(err) {
		log.Println("Resource not found")
	}
	log.Fatal(err)
}
```

## Available Services

| Service | Description |
|---------|-------------|
| `client.Auth` | Authentication (register, login, logout, OAuth) |
| `client.SubUsers` | Sub-user management (CRUD, bulk ops) |
| `client.SubUserGroups` | Sub-user group management |
| `client.Locations` | Proxy locations (countries, regions, cities, ISPs, zipcodes) |
| `client.Analytics` | Traffic and request analytics |
| `client.ProxyPresets` | Proxy preset management |
| `client.Profile` | User preferences and API keys |
| `client.TwoFactor` | Two-factor authentication |
| `client.Email` | Email change management |
| `client.Coupons` | Coupon validation and redemption |
| `client.Plans` | Plan listing and pricing |
| `client.Payments` | Payment processing and invoices |

## Documentation

- [Getting Started](https://docs.proxyhat.com?utm_source=github&utm_medium=readme&utm_campaign=sdk-go&utm_content=getting-started)
- [API Reference](https://docs.proxyhat.com/api/auth?utm_source=github&utm_medium=readme&utm_campaign=sdk-go&utm_content=api-reference)
- [Full Documentation](https://docs.proxyhat.com?utm_source=github&utm_medium=readme&utm_campaign=sdk-go&utm_content=docs-home)

## Links

- [ProxyHat](https://proxyhat.com?utm_source=github&utm_medium=readme&utm_campaign=sdk-go&utm_content=links) — Residential & mobile proxy network
- [Dashboard](https://dashboard.proxyhat.com?utm_source=github&utm_medium=readme&utm_campaign=sdk-go&utm_content=links) — Manage proxies, sub-users, and API keys
- [GitHub](https://github.com/ProxyHatCom/go-sdk) — Source code & issue tracker

## License

MIT — see [LICENSE](LICENSE) for details.
