package conf

// server config
const (
	ServerAddress = "127.0.0.1"
	ServerHost    = "api"
	ServerPort    = "8080"
)

// database config
const (
	DBDriver   = "mysql"
	DBAddress  = "127.0.0.1"
	DBHost     = "db" // containerの関係で
	DBPort     = "3306"
	DBUser     = "worker"
	DBPassword = "password"
	DBName     = "homepage"
)

// client serverの情報
const (
	// ClientServerAddr = "*"
	ClientServerAddr = "http://web:80"
)

// ディレクトリの場所とか
const (
	// file
	FileDir = "."
)

// cookieのこと
const (
	CookieName   = "_l_semi_homepage_session"
	CookieMaxAge = 3600 * 1
)

// formatまわり
const (
	DatetimeFormat    = "2006/01/02 15:04:05"
	DateFormat        = "2006/01/02"
	LogDatetimeFormat = "2006/01/02 15:04:05"
)

const (
	// DefaultTagID 備品を登録する時、タグの入力がなかったら、その他に誘導したい...
	DefaultTagID = 6
)
