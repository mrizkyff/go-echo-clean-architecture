package mappers

import (
	"go-echo-clean-architecture/internal/dto/response"
	"go-echo-clean-architecture/internal/models"
)

func MapAccessLogToDto(accessLog *models.AccessLog) *response.AccessLogResponseDto {
	return &response.AccessLogResponseDto{
		ID:           accessLog.ID,
		AccessTime:   accessLog.AccessTime,
		IpAddress:    accessLog.IpAddress,
		ClientInfo:   accessLog.ClientInfo,
		LinkID:       accessLog.LinkID,
		UserID:       accessLog.UserID,
		OriginalLink: accessLog.Link.OriginalLink,
		ShortLink:    accessLog.Link.ShortenLink,
		UserName:     accessLog.User.FullName,
	}
}
