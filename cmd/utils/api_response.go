package utils

// HttpStatus representa un estado HTTP
type HttpStatus int

const (
	Ok HttpStatus = iota
	Created
	NoContent
	BadRequest
	Unauthorized
	Forbidden
	NotFound
	Conflict
	Unprocessable
	TooManyRequests
	InternalServerError
	NotImplemented
	BadGateway
	ServiceUnavailable
	GatewayTimeout
	MethodNotAllowed
)

// statusMeta: solo código y flag de error
type statusApiMeta struct {
	Code    int
	IsError bool
}

// Registro central de HttpStatus
var httpStatusRegistry = map[HttpStatus]statusApiMeta{
	Ok:                  {200, false},
	Created:             {201, false},
	NoContent:           {204, false},
	BadRequest:          {400, true},
	Unauthorized:        {401, true},
	Forbidden:           {403, true},
	NotFound:            {404, true},
	MethodNotAllowed:    {405, true},
	Conflict:            {409, true},
	Unprocessable:       {422, true},
	TooManyRequests:     {429, true},
	InternalServerError: {500, true},
	NotImplemented:      {501, true},
	BadGateway:          {502, true},
	ServiceUnavailable:  {503, true},
	GatewayTimeout:      {504, true},
}

// Métodos sobre HttpStatus

func (s HttpStatus) Code() int {
	if meta, ok := httpStatusRegistry[s]; ok {
		return meta.Code
	}
	return 0
}

func (s HttpStatus) IsError() bool {
	if meta, ok := httpStatusRegistry[s]; ok {
		return meta.IsError
	}
	return true
}

type ApiResponse[T any] struct {
	Success bool       `json:"success"`
	Status  HttpStatus `json:"status"`
	Message string     `json:"message"`
	Data    *T         `json:"data,omitempty"`
}

// Constructor base
func NewResponse[T any](status HttpStatus, message string, data *T) ApiResponse[T] {
	return ApiResponse[T]{
		Success: !status.IsError(),
		Status:  status,
		Message: message,
		Data:    data,
	}
}

// Helpers
func SuccessResponse[T any](data T, message string) ApiResponse[T] {
	return NewResponse(Ok, message, &data)
}

func CreatedResponse[T any](data T, message string) ApiResponse[T] {
	return NewResponse(Created, message, &data)
}

func NoContentResponse(message string) ApiResponse[any] {
	return NewResponse[any](NoContent, message, nil)
}

// Error response helpers

// BadRequestResponse - 400: Solicitud malformada, parámetros inválidos o datos faltantes
func BadRequestResponse(message string) ApiResponse[any] {
	return NewResponse[any](BadRequest, message, nil)
}

// UnauthorizedResponse - 401: Falta autenticación o token inválido
func UnauthorizedResponse(message string) ApiResponse[any] {
	return NewResponse[any](Unauthorized, message, nil)
}

// ForbiddenResponse - 403: Usuario autenticado pero sin permisos para el recurso
func ForbiddenResponse(message string) ApiResponse[any] {
	return NewResponse[any](Forbidden, message, nil)
}

// NotFoundResponse - 404: Recurso no encontrado
func NotFoundResponse(message string) ApiResponse[any] {
	return NewResponse[any](NotFound, message, nil)
}

// MethodNotAllowedResponse - 405: Método HTTP no permitido para este endpoint
func MethodNotAllowedResponse(message string) ApiResponse[any] {
	return NewResponse[any](MethodNotAllowed, message, nil)
}

// ConflictResponse - 409: Conflicto con el estado actual del recurso (ej: email ya existe)
func ConflictResponse(message string) ApiResponse[any] {
	return NewResponse[any](Conflict, message, nil)
}

// UnprocessableResponse - 422: Datos válidos pero no procesables por reglas de negocio
func UnprocessableResponse(message string) ApiResponse[any] {
	return NewResponse[any](Unprocessable, message, nil)
}

// TooManyRequestsResponse - 429: Rate limit excedido, demasiadas solicitudes
func TooManyRequestsResponse(message string) ApiResponse[any] {
	return NewResponse[any](TooManyRequests, message, nil)
}

// InternalServerErrorResponse - 500: Error interno del servidor
func InternalServerErrorResponse(message string) ApiResponse[any] {
	return NewResponse[any](InternalServerError, message, nil)
}

// NotImplementedResponse - 501: Funcionalidad no implementada
func NotImplementedResponse(message string) ApiResponse[any] {
	return NewResponse[any](NotImplemented, message, nil)
}

// BadGatewayResponse - 502: Error en gateway o proxy upstream
func BadGatewayResponse(message string) ApiResponse[any] {
	return NewResponse[any](BadGateway, message, nil)
}

// ServiceUnavailableResponse - 503: Servicio temporalmente no disponible
func ServiceUnavailableResponse(message string) ApiResponse[any] {
	return NewResponse[any](ServiceUnavailable, message, nil)
}

// GatewayTimeoutResponse - 504: Timeout en gateway o proxy
func GatewayTimeoutResponse(message string) ApiResponse[any] {
	return NewResponse[any](GatewayTimeout, message, nil)
}

// ErrorResponse - Helper genérico para cualquier estado de error personalizado
func ErrorResponse(message string, status HttpStatus) ApiResponse[any] {
	return NewResponse[any](status, message, nil)
}
