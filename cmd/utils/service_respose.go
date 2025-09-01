package utils

type ServiceStatus int

const (
	Success ServiceStatus = iota
	ValidationError
	ProcessingError
	DataNotFound
	DuplicateEntry
	InsufficientPermissions
	ResourceLocked
	ConfigurationError
	ExternalServiceError
	TimeoutError
	UnknownError
)

type statusMeta struct {
	Code    string
	IsError bool
}

var serviceStatusRegistry = map[ServiceStatus]statusMeta{
	Success:                 {"SUCCESS", false},
	ValidationError:         {"VALIDATION_ERROR", true},
	ProcessingError:         {"PROCESSING_ERROR", true},
	DataNotFound:            {"DATA_NOT_FOUND", true},
	DuplicateEntry:          {"DUPLICATE_ENTRY", true},
	InsufficientPermissions: {"INSUFFICIENT_PERMISSIONS", true},
	ResourceLocked:          {"RESOURCE_LOCKED", true},
	ConfigurationError:      {"CONFIGURATION_ERROR", true},
	ExternalServiceError:    {"EXTERNAL_SERVICE_ERROR", true},
	TimeoutError:            {"TIMEOUT_ERROR", true},
	UnknownError:            {"UNKNOWN_ERROR", true},
}

func (s ServiceStatus) IsError() bool {
	if meta, ok := serviceStatusRegistry[s]; ok {
		return meta.IsError
	}
	return true
}

func (s ServiceStatus) ErrorCode() string {
	if meta, ok := serviceStatusRegistry[s]; ok {
		return meta.Code
	}

	return "UNDEFINED"
}
