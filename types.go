package proxyhat

// Auth types

type RegisterResponse struct {
	Message     string `json:"message"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Requires2FA bool   `json:"requires_2fa"`
}

type TrafficInfo struct {
	Subscription          *string `json:"subscription"`
	SubscriptionStartsAt  *string `json:"subscription_starts_at"`
	SubscriptionExpiresAt *string `json:"subscription_expires_at"`
	RegularBytes          int     `json:"regular_bytes"`
	RegularHuman          string  `json:"regular_human"`
	SubscriptionBytes     int     `json:"subscription_bytes"`
	SubscriptionHuman     string  `json:"subscription_human"`
	TotalBytes            int     `json:"total_bytes"`
	TotalHuman            string  `json:"total_human"`
}

type User struct {
	UUID    string      `json:"uuid"`
	Name    string      `json:"name"`
	Email   string      `json:"email"`
	Traffic TrafficInfo `json:"traffic"`
}

type SupportedProvider struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type SocialAccount struct {
	Provider    string  `json:"provider"`
	Email       *string `json:"email"`
	ConnectedAt *string `json:"connected_at"`
}

// Sub-user types

type SubUser struct {
	UUID             string  `json:"uuid"`
	ProxyUsername    string  `json:"proxy_username"`
	IsDefaultUser    bool    `json:"is_default_user"`
	IsTrafficLimited bool    `json:"is_traffic_limited"`
	UsedTraffic      int     `json:"used_traffic"`
	TrafficLimit     int     `json:"traffic_limit"`
	LifecycleStatus  string  `json:"lifecycle_status"`
	Name             *string `json:"name"`
	Notes            *string `json:"notes"`
	SubUserGroupID   *string `json:"sub_user_group_id"`
	CreatedAt        string  `json:"created_at"`
}

type ResetUsageResponse struct {
	Reset int `json:"reset"`
}

type BulkDeleteResponse struct {
	Requested int `json:"requested"`
	Deleted   int `json:"deleted"`
	Skipped   int `json:"skipped"`
	NotFound  int `json:"not_found"`
	Failed    int `json:"failed"`
}

// Sub-user group types

type SubUserGroup struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Description   *string `json:"description"`
	SubUsersCount int     `json:"sub_users_count"`
	CreatedAt     string  `json:"created_at"`
	SubUsers      []any   `json:"sub_users"`
}

// Location types

type Country struct {
	Code           string `json:"code"`
	Name           string `json:"name"`
	Availability   string `json:"availability"`
	ConnectionType string `json:"connection_type"`
}

type Region struct {
	Code           string  `json:"code"`
	Name           string  `json:"name"`
	CountryCode    string  `json:"country_code"`
	Availability   *string `json:"availability"`
	ConnectionType *string `json:"connection_type"`
}

type City struct {
	Code           string  `json:"code"`
	Name           string  `json:"name"`
	CountryCode    string  `json:"country_code"`
	RegionCode     *string `json:"region_code"`
	Availability   *string `json:"availability"`
	ConnectionType *string `json:"connection_type"`
}

type ISP struct {
	Code           string  `json:"code"`
	Name           string  `json:"name"`
	CountryCode    string  `json:"country_code"`
	Availability   *string `json:"availability"`
	ConnectionType *string `json:"connection_type"`
}

type Zipcode struct {
	Code           string  `json:"code"`
	Name           string  `json:"name"`
	CountryCode    string  `json:"country_code"`
	CityCode       *string `json:"city_code"`
	Availability   *string `json:"availability"`
	ConnectionType *string `json:"connection_type"`
}

// Analytics types

type TimeSeriesResponse struct {
	Labels []string `json:"labels"`
	Data   []int    `json:"data"`
}

type TotalResponse struct {
	Total int `json:"total"`
}

type DomainBreakdownItem struct {
	Domain    string `json:"domain"`
	Bandwidth int    `json:"bandwidth"`
	Requests  int    `json:"requests"`
}

type DomainBreakdownResponse struct {
	Items []DomainBreakdownItem `json:"items"`
}

// Profile types

type Preferences struct {
	Data map[string]any `json:"data"`
}

type APIKey struct {
	ID             string  `json:"id"`
	Name           *string `json:"name"`
	PlainTextToken *string `json:"plain_text_token"`
	CreatedAt      *string `json:"created_at"`
}

// Two-factor types

type TwoFactorStatus struct {
	Enabled bool `json:"enabled"`
}

type TwoFactorEnableResponse struct {
	QR            string   `json:"qr"`
	Secret        string   `json:"secret"`
	RecoveryCodes []string `json:"recovery_codes"`
}

type RecoveryCodes struct {
	Codes []string `json:"codes"`
}

// Email types

type EmailChangeResponse struct {
	Message string `json:"message"`
}

// Coupon types

type Coupon struct {
	ID          string   `json:"id"`
	Code        string   `json:"code"`
	Type        string   `json:"type"`
	Data        any      `json:"data"`
	Discount    *float64 `json:"discount"`
	FinalAmount *float64 `json:"final_amount"`
}

type CouponResponse struct {
	Success bool    `json:"success"`
	Coupon  *Coupon `json:"coupon"`
}

// Plan types

type RegularPlan struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	GB         int     `json:"gb"`
	PricePerGB float64 `json:"price_per_gb"`
	PriceTotal float64 `json:"price_total"`
	Currency   string  `json:"currency"`
}

type SubscriptionPlan struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	GB              int     `json:"gb"`
	PricePerGB      float64 `json:"price_per_gb"`
	PriceTotal      float64 `json:"price_total"`
	Period          string  `json:"period"`
	RolloverEnabled bool    `json:"rollover_enabled"`
}

// Payment types

type PaymentCreateResponse struct {
	Success   bool   `json:"success"`
	PaymentID string `json:"payment_id"`
}

type CryptoInfo struct {
	Code     string  `json:"code"`
	Currency string  `json:"currency"`
	Network  string  `json:"network"`
	Icon     *string `json:"icon"`
	Label    *string `json:"label"`
}

type PaymentDetails struct {
	PayAddress  string     `json:"pay_address"`
	CryptoAmount float64   `json:"crypto_amount"`
	AmountUSD   float64    `json:"amount_usd"`
	Crypto      CryptoInfo `json:"crypto"`
	Status      string     `json:"status"`
	TxHash      *string    `json:"tx_hash"`
	ExpiresAt   *string    `json:"expires_at"`
	CompletedAt *string    `json:"completed_at"`
}

type Payment struct {
	ID        string   `json:"id"`
	Type      string   `json:"type"`
	Status    string   `json:"status"`
	Amount    *float64 `json:"amount"`
	Currency  *string  `json:"currency"`
	CreatedAt *string  `json:"created_at"`
}

type Cryptocurrency struct {
	Code     string  `json:"code"`
	Currency string  `json:"currency"`
	Network  string  `json:"network"`
	Icon     *string `json:"icon"`
	Label    *string `json:"label"`
}

// Proxy preset types

type ProxyPreset struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Data      map[string]any `json:"data"`
	CreatedAt *string        `json:"created_at"`
}
