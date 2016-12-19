package constant

// Error constant
const (
	ErrAuthorizationError      = "Authorization Error"
	ErrInsufficientPermissions = "Insufficient Permissions"
)

// log type
const (
	ERROR  = -1
	INFO   = 0
	PROFIT = 1
	BUY    = 2
	SELL   = 3
	CANCEL = 4
)

// delete log time type
const (
	LastTime = "0"
	Day      = "1"
	Week     = "2"
	Month    = "3"
)

// exchange type
const (
	OkCoinCn     = "okcoin.cn"
	Huobi        = "huobi"
	Poloniex     = "poloniex"
	Btcc         = "btcc"
	Chbtc        = "chbtc"
	OkcoinFuture = "okcoin.future"
	OandaV20     = "oanda.v20"
)

// trade type
const (
	TradeTypeBuy        = "BUY"
	TradeTypeSell       = "SELL"
	TradeTypeLong       = "LONG"
	TradeTypeShort      = "SHORT"
	TradeTypeLongClose  = "LONG_CLOSE"
	TradeTypeShortClose = "SHORT_CLOSE"
)

// stock type (will useless)
const (
	BTC = "BTC"
	LTC = "LTC"
)

// CONSTS : Javascript Global Constants
var (
	CONSTS        = []string{"BTC", "LTC", "M", "M5", "M15", "M30", "H", "D", "W"}
	ExchangeTypes = []string{"okcoin.cn", "huobi", "poloniex", "chbtc", "okcoin.future", "oanda.v20"}
)
