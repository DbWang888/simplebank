package util


//检验货币类型

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

 //检验是否为支持的货币类型
 func IsSupportedCurrency (currency string) bool {
	switch currency{
	case USD,EUR,CAD:
		return true
	}
	return false
 }