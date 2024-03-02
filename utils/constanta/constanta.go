package constanta

const (
	URL_STORAGE = "https://cimxqffotlogzqvadisz.supabase.co/storage/v1"
	URL         = URL_STORAGE + "/object/public/"
)

// error
const (
	ERROR      = "error"
	NOT_FOUND  = "record not found"
	EXISTS     = ERROR + ": %s data already exists"
	WRONG_PASS = ERROR + ": password is wrong"
)

const (
	PENDING  = "pending"
	REJECTED = "rejected"
	APPROVE  = "approve"
)
