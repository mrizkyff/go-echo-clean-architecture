package response

import (
	"github.com/google/uuid"
	"time"
)

type AccessLogResponseDto struct {
	ID           uuid.UUID `json:"id"`
	AccessTime   time.Time `json:"access_time"`
	IpAddress    string    `json:"ip_address"`
	ClientInfo   string    `json:"client_info"`
	LinkID       uuid.UUID `json:"link_id"`
	UserID       uuid.UUID `json:"user_id"`
	OriginalLink string    `json:"original_link"`
	ShortLink    string    `json:"short_link"`
	UserName     string    `json:"user_name"`
}
