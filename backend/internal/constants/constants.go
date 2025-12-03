package constants

import "time"

const AccessTokenTTL = 1 * time.Hour       // 1 hour
const RefreshTokenTTL = 7 * 24 * time.Hour // 7 days
const ExchangeTokenTTL = 5 * time.Minute   // 5 minutes
