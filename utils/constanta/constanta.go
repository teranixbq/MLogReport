package constanta

const (
	URL_STORAGE = "https://cimxqffotlogzqvadisz.supabase.co/storage/v1"
	URL         = URL_STORAGE + "/object/public/"
)

// error
const (
	ERROR        = "error"
	NOT_FOUND    = "record not found"
	DATA_NULL    = "no data available yet"
	EXISTS       = ERROR + ": %s data already exists"
	WRONG_PASS   = ERROR + ": password is wrong"
	INVALID_DATE = ERROR + ": date must be in 'yyyy-mm-dd' format"
	DATE_START   = ERROR + ": start date must be today or later"
	DATE_END     = ERROR + ": the end date must be after the start date"
	DATE_END_EQ  = ERROR + ": the start date must be different from the end date"
	TIME_EXISTS  = ERROR + ": there is still an active time period"
)

const (
	PENDING  = "pending"
	REJECTED = "rejected"
	APPROVE  = "approve"
)
