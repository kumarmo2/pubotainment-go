package constants

// TODO: this should not be in the code.
// should something like AWS secret manager.
var JWT_SECRET_ADMIN []byte = []byte("kumarmo2")
var JWT_SECRET_USER []byte = []byte("kumarmo2user")

const USER_AUTH_COOKIE_NAME = "userAuth"
const ADMIN_AUTH_COOKIE_NAME = "auth"
