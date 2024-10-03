package response

import "net/http"

type Status struct {
	Text string `json:"text"`
	Code int    `json:"code"`
}

func (s Status) Equal(other Status) bool {
	return s.Code == other.Code
}

var (
	InternalServerError = Status{
		Text: "Internal server error",
		Code: http.StatusInternalServerError,
	}
	NotFound = Status{
		Text: "Not found",
		Code: http.StatusNotFound,
	}
	BadRequest = Status{
		Text: "Bad request",
		Code: http.StatusBadRequest,
	}
	Conflict = Status{
		Text: "Conflict",
		Code: http.StatusConflict,
	}
	Unknown = Status{
		Text: "Unknown",
		Code: http.StatusNotImplemented,
	}

	InvalidEmail = Status{
		Text: "Invalid email",
		Code: http.StatusBadRequest,
	}
	InvalidPassword = Status{
		Text: "Invalid password",
		Code: http.StatusBadRequest,
	}
	InvalidUsername = Status{
		Text: "Invalid username",
		Code: http.StatusBadRequest,
	}
	EmailAlreadyExists = Status{
		Text: "Email already exists",
		Code: http.StatusConflict,
	}
	UsernameAlreadyExists = Status{
		Text: "Username already exists",
		Code: http.StatusConflict,
	}
	Unauthorized = Status{
		Text: "Unauthorized",
		Code: http.StatusUnauthorized,
	}
	Forbidden = Status{
		Text: "Forbidden",
		Code: http.StatusForbidden,
	}

	DBQueryError = Status{
		Text: "Database query error",
		Code: http.StatusInternalServerError,
	}
	DBExecutionError = Status{
		Text: "Database execution error",
		Code: http.StatusInternalServerError,
	}
	DBRowsError = Status{
		Text: "Database rows error",
		Code: http.StatusInternalServerError,
	}
	DBLastRowIdError = Status{
		Text: "Database last row ID error",
		Code: http.StatusInternalServerError,
	}
	DBScanError = Status{
		Text: "Database scan error",
		Code: http.StatusInternalServerError,
	}
	DBTransactionError = Status{
		Text: "Database transaction error",
		Code: http.StatusInternalServerError,
	}
	DBTransactionClosed = Status{
		Text: "Database transaction closed",
		Code: http.StatusInternalServerError,
	}
	DBCommitError = Status{
		Text: "Database commit error",
		Code: http.StatusInternalServerError,
	}
	DBItemAlreadyExists = Status{
		Text: "Database item already exists",
		Code: http.StatusConflict,
	}

	JsonDecodingError = Status{
		Text: "JSON decoding error",
		Code: http.StatusInternalServerError,
	}
	JsonEncodingError = Status{
		Text: "JSON encoding error",
		Code: http.StatusInternalServerError,
	}

	SuccessfulCreation = Status{
		Text: "Successful creation",
		Code: http.StatusCreated,
	}
	SuccessfulDeletion = Status{
		Text: "Successful deletion",
		Code: http.StatusOK,
	}
	SuccessfulUpdate = Status{
		Text: "Successful update",
		Code: http.StatusOK,
	}
	SuccessfulSearch = Status{
		Text: "Successful search",
		Code: http.StatusOK,
	}

	FailedCreation = Status{
		Text: "Failed creation",
		Code: http.StatusConflict,
	}
	FailedDeletion = Status{
		Text: "Failed deletion",
		Code: http.StatusConflict,
	}
	FailedUpdation = Status{
		Text: "Failed update",
		Code: http.StatusConflict,
	}
	FailedSearch = Status{
		Text: "Failed search",
		Code: http.StatusConflict,
	}

	EncryptionError = Status{
		Text: "Encryption error",
		Code: http.StatusInternalServerError,
	}
	DecryptionError = Status{
		Text: "Decryption error",
		Code: http.StatusInternalServerError,
	}
)
