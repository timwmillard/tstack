package handler

import (
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	AppEnv       string
	Port         int32
	Host         url.URL
	CookieSecure bool

	Stripe StripeConfig
}

type StripeConfig struct {
	APIKey        string
	WebhookSecret string
}

func LoadConfig() Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	var port int32 = 8080
	if p, _ := strconv.Atoi(os.Getenv("PORT")); p != 0 {
		port = int32(p)
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = fmt.Sprintf("http://localhost:%d", port)
	}
	hostUrl, _ := url.Parse(host)

	cookieSecure, _ := strconv.ParseBool(os.Getenv("COOKIE_SECURE"))

	return Config{
		AppEnv:       env,
		Port:         port,
		Host:         *hostUrl,
		CookieSecure: cookieSecure,
		Stripe: StripeConfig{
			APIKey:        os.Getenv("STRIPE_API_KEY"),
			WebhookSecret: os.Getenv("STRIPE_WEBHOOK_SECRET"),
		},
	}
}

func LogConfig(config Config) {
	slog.Info("Config:",
		"app_env", config.AppEnv,
		"port", config.Port,
		"host", config.Host.String(),
	)
	slog.Info("Config: Stripe",
		"stripe_api_key", safeStripe(config.Stripe.APIKey),
		"stripe_webhook_secret", safeStripe(config.Stripe.WebhookSecret),
	)
}

func safeString(s string) string {
	var safe string
	for range s {
		safe += "x"
	}
	return safe
}

func safeStripe(s string) string {
	if s == "" {
		return ""
	}
	parts := strings.Split(s, "_")

	var safe string
	for i := 0; i < len(parts)-1; i++ {
		safe += parts[i] + "_"
	}

	for range parts[len(parts)-1] {
		safe += "x"
	}
	return safe
}
