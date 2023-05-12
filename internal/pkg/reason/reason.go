package reason

var (
	InvalidRequest      = "request format invalid"
	InternalServerError = "interval server error"
)

var (
	CurrencyFound         = "currency found"
	CurrencyNotFound      = "currency not found"
	CurrencySuccessBrowse = "success browse currency"
	CurrencySuccessCreate = "success to create currency"
	CurrencySuccessUpdate = "success to update currency"
	CurrencySuccessDelete = "success to delete currency"
	CurrencyFailedBrowse  = "can't browse currency"
	CurrencyFailedCreate  = "failed to create currency"
	CurrencyFailedUpdate  = "failed to update currency"
	CurrencyFailedDelete  = "failed to delete currency"
)

var (
	TransactionTypeFound         = "transaction type found"
	TransactionTypeNotFound      = "transaction type not found"
	TransactionTypeSuccessBrowse = "success browse transaction type"
	TransactionTypeSuccessCreate = "success to create transaction type"
	TransactionTypeSuccessUpdate = "success to update transaction type"
	TransactionTypeSuccessDelete = "success to delete transaction type"
	TransactionTypeFailedBrowse  = "can't browse transaction type"
	TransactionTypeFailedCreate  = "failed to create transaction type"
	TransactionTypeFailedUpdate  = "failed to update transaction type"
	TransactionTypeFailedDelete  = "failed to delete transaction type"
)

var (
	TransactionCategoryFound         = "transaction category found"
	TransactionCategoryNotFound      = "transaction category not found"
	TransactionCategorySuccessBrowse = "success browse transaction category"
	TransactionCategorySuccessCreate = "success to create transaction category"
	TransactionCategorySuccessUpdate = "success to update transaction category"
	TransactionCategorySuccessDelete = "success to delete transaction category"
	TransactionCategoryFailedBrowse  = "can't browse transaction category"
	TransactionCategoryFailedCreate  = "failed to create transaction category"
	TransactionCategoryFailedUpdate  = "failed to update transaction category"
	TransactionCategoryFailedDelete  = "failed to delete transaction category"
)

var (
	UserFound         = "user found"
	UserNotFound      = "user not found"
	UserSuccessBrowse = "success browse user"
	UserSuccessDelete = "success to delete user"
	UserFailedBrowse  = "can't browse user"
	UserFailedCreate  = "failed to create user"
	UserFailedDelete  = "failed to delete user"

	UserPasswordChangeSuccess = "success to change user password"
	UserPasswordChangeFailed  = "failed to change user password"
)

var (
	RegisterSuccess   = "success registered"
	RegisterFailed    = "failed to register"
	RegisterUserExist = "failed to register, username already taken"

	PasswordChangeSuccess = "succes to change password"
	PasswordChangeFailed  = "failed to change password"

	LoginFailed   = "failed to login"
	LoginSuccess  = "success to login"
	LogoutFailed  = "failed to logout"
	LogoutSuccess = "success to logout"

	RefreshToken        = "success to refresh token"
	RefreshTokenInvalid = "failed to refresh token"

	Unauthenticated = "unauthenticated, please use your credentials to login"
	Unauthorized    = "unauthorized, you don't have access to this resource"
)
