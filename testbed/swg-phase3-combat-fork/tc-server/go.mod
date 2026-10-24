module tc-server

go 1.25.0

require (
	github.com/google/uuid v1.6.0
	github.com/gorilla/websocket v1.5.3
	golang.org/x/crypto v0.24.0
	modernc.org/sqlite v1.59.0
)

require (
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/mattn/go-isatty v0.0.24 // indirect
	github.com/ncruces/go-strftime v1.0.0 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	golang.org/x/sys v0.47.0 // indirect
	modernc.org/libc v1.75.7 // indirect
	modernc.org/mathutil v1.7.1 // indirect
	modernc.org/memory v1.12.1 // indirect
)

// This replace directive is an artifact of verifying this build in a
// network-restricted sandbox (golang.org itself was blocked, so the
// golang.org/x/crypto vanity import couldn't resolve; github.com/golang/crypto
// is the same canonical repo it redirects to). It's harmless to keep on a
// normal network, or safe to delete if you'd rather use the standard path.
replace golang.org/x/crypto => github.com/golang/crypto v0.24.0
