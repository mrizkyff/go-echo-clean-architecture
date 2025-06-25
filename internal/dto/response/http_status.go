package response

import (
	"fmt"
	"net/http"
)

// HTTPStatus represents HTTP status codes with their corresponding messages
type HTTPStatus struct {
	Code    int
	Message string
}

// HTTP Status Code constants
var (
	// 1xx Informational
	Continue           = HTTPStatus{Code: http.StatusContinue, Message: "Continue"}
	SwitchingProtocols = HTTPStatus{Code: http.StatusSwitchingProtocols, Message: "Switching Protocols"}
	Processing         = HTTPStatus{Code: http.StatusProcessing, Message: "Processing"}

	// 2xx Success
	OK                   = HTTPStatus{Code: http.StatusOK, Message: "OK"}
	Created              = HTTPStatus{Code: http.StatusCreated, Message: "Created"}
	Accepted             = HTTPStatus{Code: http.StatusAccepted, Message: "Accepted"}
	NonAuthoritativeInfo = HTTPStatus{Code: http.StatusNonAuthoritativeInfo, Message: "Non-Authoritative Information"}
	NoContent            = HTTPStatus{Code: http.StatusNoContent, Message: "No Content"}
	ResetContent         = HTTPStatus{Code: http.StatusResetContent, Message: "Reset Content"}
	PartialContent       = HTTPStatus{Code: http.StatusPartialContent, Message: "Partial Content"}

	// 3xx Redirection
	MultipleChoices   = HTTPStatus{Code: http.StatusMultipleChoices, Message: "Multiple Choices"}
	MovedPermanently  = HTTPStatus{Code: http.StatusMovedPermanently, Message: "Moved Permanently"}
	Found             = HTTPStatus{Code: http.StatusFound, Message: "Found"}
	SeeOther          = HTTPStatus{Code: http.StatusSeeOther, Message: "See Other"}
	NotModified       = HTTPStatus{Code: http.StatusNotModified, Message: "Not Modified"}
	UseProxy          = HTTPStatus{Code: http.StatusUseProxy, Message: "Use Proxy"}
	TemporaryRedirect = HTTPStatus{Code: http.StatusTemporaryRedirect, Message: "Temporary Redirect"}
	PermanentRedirect = HTTPStatus{Code: http.StatusPermanentRedirect, Message: "Permanent Redirect"}

	// 4xx Client Error
	BadRequest                   = HTTPStatus{Code: http.StatusBadRequest, Message: "Bad Request"}
	Unauthorized                 = HTTPStatus{Code: http.StatusUnauthorized, Message: "Unauthorized"}
	PaymentRequired              = HTTPStatus{Code: http.StatusPaymentRequired, Message: "Payment Required"}
	Forbidden                    = HTTPStatus{Code: http.StatusForbidden, Message: "Forbidden"}
	NotFound                     = HTTPStatus{Code: http.StatusNotFound, Message: "Not Found"}
	MethodNotAllowed             = HTTPStatus{Code: http.StatusMethodNotAllowed, Message: "Method Not Allowed"}
	NotAcceptable                = HTTPStatus{Code: http.StatusNotAcceptable, Message: "Not Acceptable"}
	ProxyAuthRequired            = HTTPStatus{Code: http.StatusProxyAuthRequired, Message: "Proxy Authentication Required"}
	RequestTimeout               = HTTPStatus{Code: http.StatusRequestTimeout, Message: "Request Timeout"}
	Conflict                     = HTTPStatus{Code: http.StatusConflict, Message: "Conflict"}
	Gone                         = HTTPStatus{Code: http.StatusGone, Message: "Gone"}
	LengthRequired               = HTTPStatus{Code: http.StatusLengthRequired, Message: "Length Required"}
	PreconditionFailed           = HTTPStatus{Code: http.StatusPreconditionFailed, Message: "Precondition Failed"}
	RequestEntityTooLarge        = HTTPStatus{Code: http.StatusRequestEntityTooLarge, Message: "Request Entity Too Large"}
	RequestURITooLong            = HTTPStatus{Code: http.StatusRequestURITooLong, Message: "Request-URI Too Long"}
	UnsupportedMediaType         = HTTPStatus{Code: http.StatusUnsupportedMediaType, Message: "Unsupported Media Type"}
	RequestedRangeNotSatisfiable = HTTPStatus{Code: http.StatusRequestedRangeNotSatisfiable, Message: "Requested Range Not Satisfiable"}
	ExpectationFailed            = HTTPStatus{Code: http.StatusExpectationFailed, Message: "Expectation Failed"}
	Teapot                       = HTTPStatus{Code: http.StatusTeapot, Message: "I'm a teapot"}
	MisdirectedRequest           = HTTPStatus{Code: http.StatusMisdirectedRequest, Message: "Misdirected Request"}
	UnprocessableEntity          = HTTPStatus{Code: http.StatusUnprocessableEntity, Message: "Unprocessable Entity"}
	Locked                       = HTTPStatus{Code: http.StatusLocked, Message: "Locked"}
	FailedDependency             = HTTPStatus{Code: http.StatusFailedDependency, Message: "Failed Dependency"}
	TooEarly                     = HTTPStatus{Code: http.StatusTooEarly, Message: "Too Early"}
	UpgradeRequired              = HTTPStatus{Code: http.StatusUpgradeRequired, Message: "Upgrade Required"}
	PreconditionRequired         = HTTPStatus{Code: http.StatusPreconditionRequired, Message: "Precondition Required"}
	TooManyRequests              = HTTPStatus{Code: http.StatusTooManyRequests, Message: "Too Many Requests"}
	RequestHeaderFieldsTooLarge  = HTTPStatus{Code: http.StatusRequestHeaderFieldsTooLarge, Message: "Request Header Fields Too Large"}
	UnavailableForLegalReasons   = HTTPStatus{Code: http.StatusUnavailableForLegalReasons, Message: "Unavailable For Legal Reasons"}

	// 5xx Server Error
	InternalServerError           = HTTPStatus{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	NotImplemented                = HTTPStatus{Code: http.StatusNotImplemented, Message: "Not Implemented"}
	BadGateway                    = HTTPStatus{Code: http.StatusBadGateway, Message: "Bad Gateway"}
	ServiceUnavailable            = HTTPStatus{Code: http.StatusServiceUnavailable, Message: "Service Unavailable"}
	GatewayTimeout                = HTTPStatus{Code: http.StatusGatewayTimeout, Message: "Gateway Timeout"}
	HTTPVersionNotSupported       = HTTPStatus{Code: http.StatusHTTPVersionNotSupported, Message: "HTTP Version Not Supported"}
	VariantAlsoNegotiates         = HTTPStatus{Code: http.StatusVariantAlsoNegotiates, Message: "Variant Also Negotiates"}
	InsufficientStorage           = HTTPStatus{Code: http.StatusInsufficientStorage, Message: "Insufficient Storage"}
	LoopDetected                  = HTTPStatus{Code: http.StatusLoopDetected, Message: "Loop Detected"}
	NotExtended                   = HTTPStatus{Code: http.StatusNotExtended, Message: "Not Extended"}
	NetworkAuthenticationRequired = HTTPStatus{Code: http.StatusNetworkAuthenticationRequired, Message: "Network Authentication Required"}
)

// GetStatusByCode returns HTTPStatus by status code
func GetStatusByCode(code int) HTTPStatus {
	statusMap := map[int]HTTPStatus{
		// 1xx
		100: Continue,
		101: SwitchingProtocols,
		102: Processing,

		// 2xx
		200: OK,
		201: Created,
		202: Accepted,
		203: NonAuthoritativeInfo,
		204: NoContent,
		205: ResetContent,
		206: PartialContent,

		// 3xx
		300: MultipleChoices,
		301: MovedPermanently,
		302: Found,
		303: SeeOther,
		304: NotModified,
		305: UseProxy,
		307: TemporaryRedirect,
		308: PermanentRedirect,

		// 4xx
		400: BadRequest,
		401: Unauthorized,
		402: PaymentRequired,
		403: Forbidden,
		404: NotFound,
		405: MethodNotAllowed,
		406: NotAcceptable,
		407: ProxyAuthRequired,
		408: RequestTimeout,
		409: Conflict,
		410: Gone,
		411: LengthRequired,
		412: PreconditionFailed,
		413: RequestEntityTooLarge,
		414: RequestURITooLong,
		415: UnsupportedMediaType,
		416: RequestedRangeNotSatisfiable,
		417: ExpectationFailed,
		418: Teapot,
		421: MisdirectedRequest,
		422: UnprocessableEntity,
		423: Locked,
		424: FailedDependency,
		425: TooEarly,
		426: UpgradeRequired,
		428: PreconditionRequired,
		429: TooManyRequests,
		431: RequestHeaderFieldsTooLarge,
		451: UnavailableForLegalReasons,

		// 5xx
		500: InternalServerError,
		501: NotImplemented,
		502: BadGateway,
		503: ServiceUnavailable,
		504: GatewayTimeout,
		505: HTTPVersionNotSupported,
		506: VariantAlsoNegotiates,
		507: InsufficientStorage,
		508: LoopDetected,
		510: NotExtended,
		511: NetworkAuthenticationRequired,
	}

	if status, exists := statusMap[code]; exists {
		return status
	}

	// Return unknown status if not found
	return HTTPStatus{Code: code, Message: "Unknown Status"}
}

// String returns the formatted status string
func (h HTTPStatus) String() string {
	return fmt.Sprintf("%d %s", h.Code, h.Message)
}

// IsSuccess returns true if status code is 2xx
func (h HTTPStatus) IsSuccess() bool {
	return h.Code >= 200 && h.Code < 300
}

// IsRedirection returns true if status code is 3xx
func (h HTTPStatus) IsRedirection() bool {
	return h.Code >= 300 && h.Code < 400
}

// IsClientError returns true if status code is 4xx
func (h HTTPStatus) IsClientError() bool {
	return h.Code >= 400 && h.Code < 500
}

// IsServerError returns true if status code is 5xx
func (h HTTPStatus) IsServerError() bool {
	return h.Code >= 500 && h.Code < 600
}

// IsError returns true if status code is 4xx or 5xx
func (h HTTPStatus) IsError() bool {
	return h.IsClientError() || h.IsServerError()
}
