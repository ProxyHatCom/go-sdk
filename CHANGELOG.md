# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2026-02-14

### Added

- Initial release of ProxyHat Go SDK
- Client with functional options (`WithBaseURL`, `WithHTTPClient`, `WithTimeout`)
- Authentication service (register, login, logout, OAuth, social accounts)
- Sub-users service (CRUD, bulk operations, usage reset)
- Sub-user groups service (CRUD)
- Locations service (countries, regions, cities, ISPs, zipcodes with pagination)
- Analytics service (traffic, requests, domain breakdown)
- Proxy presets service (CRUD)
- Profile service (preferences, API keys)
- Two-factor authentication service (enable, disable, recovery codes, password change)
- Email service (request/confirm/cancel email change)
- Coupons service (validate, apply, redeem)
- Plans service (regular and subscription plans, pricing)
- Payments service (CRUD, invoice download, cryptocurrencies)
- Structured error types with status code helpers
- Zero external dependencies — stdlib `net/http` only

[0.1.0]: https://github.com/ProxyHatCom/go-sdk/releases/tag/v0.1.0
