module swg-server

go 1.22

require (
	github.com/google/uuid v1.6.0
	github.com/gorilla/websocket v1.5.3
	github.com/mattn/go-sqlite3 v1.14.22
	golang.org/x/crypto v0.24.0
)

// This replace directive is an artifact of verifying this build in a
// network-restricted sandbox (golang.org itself was blocked, so the
// golang.org/x/crypto vanity import couldn't resolve; github.com/golang/crypto
// is the same canonical repo it redirects to). It's harmless to keep on a
// normal network, or safe to delete if you'd rather use the standard path.
replace golang.org/x/crypto => github.com/golang/crypto v0.24.0
