package appdata

import (
	"time"
)

// Global variables
const (
	InfoColorFloat = "\033[1;34m%.0f\033[0m\t"
	InfoColorUint  = "\033[1;34m%d\033[0m\t"
	InfoColor      = "\033[1;34m%20s\033[0m\t"
	SuccessColor   = "\033[1;32m%20s\033[0m\t"
	ErrorColor     = "\033[1;31m%s\033[0m"
	DebugColor     = "\033[1;35m%s\033[0m"

	ColorSuccess = "\033[32m"
	ColorBanner  = "\033[34;4m"
	ColorInfo    = "\033[37m"
	ColorDimmed  = "\033[37m"
	ColorReset   = "\033[0m"
	ColorError   = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorWarning = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorPurple  = "\033[35m"
	ColorCyan    = "\033[36m"
	ColorWhite   = "\033[37m\033[47m"
)

// Foreground text colors
// const (
// 	FgBlack Attribute = iota + 30
// 	FgRed 31
// 	FgGreen 32
// 	FgYellow 33
// 	FgBlue 34
// 	FgMagenta 35
// 	FgCyan 36
// 	FgWhite 37
// )

// --------------------------------- TICK CHANNEL ---------------------------------

var (
	ChNseTicks chan TickData
	ChStkTick  chan TickData
)

type TickData struct {
	Timestamp       time.Time
	LastTradedPrice float64
	Symbol          string
	LastPrice       float64
	Buy_Demand      uint32
	Sell_Demand     uint32
	TradesTillNow   uint32
	OpenInterest    uint32
}

// --------------------------------- USER STRATEGIES ---------------------------------
type UserStrategies_S struct {
	Strategy     string
	Enabled      bool
	Engine       string
	Trigger_time time.Time
	Trigger_days string
	Cdl_size     int
	Instruments  string
	Parameters   Parameters_S
}

type Parameters_S struct {
	Kite_Setting    KiteSetting_S    `json:"kite_setting"`
	Controls        Controls_S       `json:"controls"`
	Option_setting  OptionsSetting_S `json:"options_setting"`
	Futures_Setting FuturesSetting_S `json:"futures_setting"`
}

type KiteSetting_S struct {
	Products     string `json:"products"`
	Varieties    string `json:"varieties"`
	OrderType    string `json:"order_type"`
	Validities   string `json:"validities"`
	PositionType string `json:"position_type"`
}

type Controls_S struct {
	TradeSimulate    bool      `json:"trade_simulate"`
	TargetPer        float64   `json:"target_per"`              // "target": 1,
	SlPer            float64   `json:"stoploss_per"`            // "sl": 1,
	DeepSlPer        float64   `json:"deep_stoploss_per"`       // "deepsl": 1
	DelayedSlMin     time.Time `json:"delayed_stoploss_min"`    // 	"delayed_stoploss_min": "00:30:00",
	StallDetectMin   time.Time `json:"stall_detect_period_min"` // 	"stall_detect_period_min": "00:30:00"
	MaxBudget        float64   `json:"budget_max_per"`          // "limit_budget": 50%,
	LimitAmount      float64   `json:"limit_amount"`
	TrailTarget      bool      `json:"trail_target_en"`      // 	"trail_target_en": true,
	PositionReversal bool      `json:"position_reversal_en"` // 	"position_reversal_en": true,
	WinningRatio     float64   `json:"winning_ratio"`        // "winning_ratio": 80%,
}

type OptionsSetting_S struct {
	OrderRoute       string `json:"order_route"`
	OptionLevel      int    `json:"option_level"`
	OptionExpiryWeek int    `json:"option_expiry_week"`
}

type FuturesSetting_S struct {
	FuturesExpiryMonth    int  `json:"futures_expiry_month"`
	SkipExipryWeekFutures bool `json:"skip_exipry_week"`
}

// --------------------------------- ORDER BOOK  ---------------------------------
type OrderBook_S struct {
	Id            uint16
	Date          time.Time
	Instr         string
	Strategy      string
	Status        string
	Dir           string
	Exit_reason   string
	Info          Info_S
	Targets       Targets_S
	Orders_entr   []Trade
	Orders_exit   []Trade
	Post_analysis string
}

// Trade represents an individual trade response.
type Trade struct {
	AveragePrice      float64 `json:"average_price"`
	Quantity          float64 `json:"quantity"`
	TradeID           string  `json:"trade_id"`
	Product           string  `json:"product"`
	FillTimestamp     string  `json:"fill_timestamp"`
	ExchangeTimestamp string  `json:"exchange_timestamp"`
	ExchangeOrderID   string  `json:"exchange_order_id"`
	OrderID           string  `json:"order_id"`
	TransactionType   string  `json:"transaction_type"`
	TradingSymbol     string  `json:"tradingsymbol"`
	Exchange          string  `json:"exchange"`
	InstrumentToken   uint32  `json:"instrument_token"`
}

type Targets_S struct {
	Entry    float64 `json:"entry"`
	Target   float64 `json:"target"`
	Stoploss float64 `json:"stoploss"`
}

type Info_S struct {
	TradingSymbol     string  `json:"trading_symbol"`
	Order_simulation  bool    `json:"order_simulation"`
	Exchange          string  `json:"exchange"`
	OrderIdEntr       uint64  `json:"order_id_entr"`
	OrderIdExit       uint64  `json:"order_id_exit"`
	QtyReq            float64 `json:"qty_req"`
	QtyFilled         float64 `json:"qty_filled"`
	UserExitRequested bool    `json:"user_exit_requested"`
	AvgPriceEnter     float64 `json:"avg_price_entr"`
	AvgPriceExit      float64 `json:"avg_price_exit"`
	ErrorCount        uint64  `json:"error_count"`
}

// --------------------------------- API SIGNAL  ---------------------------------
type ApiSignal struct {
	Status   string    `json:"status"`
	Id       uint16    `json:"id"`
	Date     time.Time `json:"date"`
	Instr    string    `json:"instr"`
	Strategy string    `json:"strategy"`
	Dir      string    `json:"dir"`
	Entry    float64   `json:"entry"`
	Target   float64   `json:"target"`
	Stoploss float64   `json:"stoploss"`
}

// --------------------------------- ENV VARIABLES ---------------------------------

var UserSettings = []string{
	"APP_LIVE_TRADING_MODE",
	"ZERODHA_USER_ID",
	"ZERODHA_PASSWORD",
	"ZERODHA_API_KEY",
	"ZERODHA_PIN",
	"ZERODHA_API_SECRET",
	"ZERODHA_TOTP_SECRET_KEY",
	"ZERODHA_REQ_TOKEN_URL",
	"TIMESCALEDB_ADDRESS",
	"TIMESCALEDB_USERNAME",
	"TIMESCALEDB_PASSWORD",
	"TIMESCALEDB_PORT",
	"ALGO_ANALYSIS_ADDRESS",
	"DB_TBL_INSTRUMENTS",
	"DB_TBL_TICK_NSEFUT",
	"DB_TBL_TICK_NSESTK",
	"DB_TBL_USER_SYMBOLS",
	"DB_TBL_USER_SETTING",
	"DB_TBL_USER_STRATEGIES",
	"DB_TBL_ORDER_BOOK",
	"DB_TEST_PREFIX",
	"DB_TBL_PREFIX_USER_ID",
}

var (
	Env = make(map[string]string)
)
