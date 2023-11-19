package dateonly

import (
	"time"
)

// DateOnly 型を定義
type DateOnly struct {
	time.Time
}

// NewDateOnly 関数を使用して DateOnly インスタンスを作成
func NewDateOnly(year, month, day int) DateOnly {
	return DateOnly{time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)}
}

// NewDateOnlyFromString 日付文字列からDateOnlyインスタンスを作成
func NewDateOnlyFromString(dateString, layout string) (DateOnly, error) {
	parsedTime, err := time.Parse(layout, dateString)
	if err != nil {
		return DateOnly{}, err
	}
	return DateOnly{parsedTime}, nil
}

// NewDateOnlyFromStringAndLocation 関数を使用して DateOnly インスタンスを文字列とロケーションから作成
func NewDateOnlyFromStringAndLocation(dateString, locationName string) (DateOnly, error) {
	// ロケーションを取得
	loc, err := time.LoadLocation(locationName)
	if err != nil {
		return DateOnly{}, err
	}

	// ロケーションを指定して日付文字列をパース
	parsedTime, err := time.ParseInLocation("2006-01-02", dateString, loc)
	if err != nil {
		return DateOnly{}, err
	}

	return DateOnly{parsedTime}, nil
}

// String レイアウト省略時は"2006-01-02"でフォーマットする
func (s *DateOnly) String(layout ...string) string {
	if len(layout) != 0 {
		return s.Format(layout[0])
	}
	return s.Format(time.DateOnly)
}
