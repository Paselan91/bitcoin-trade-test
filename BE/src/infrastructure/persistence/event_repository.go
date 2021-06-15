package persistence

import (
	"app/src/domain"
	"app/src/domain/repository"
	"github.com/jinzhu/gorm"
	"log"
)

type EventRepository struct {
	Conn *gorm.DB
}

func NewEventRepository(conn *gorm.DB) repository.EventRepository {
	return &EventRepository{Conn: conn}
}

func (s *SignalEvent) Save() bool {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	repo := persistence.EventRepositoryImpl(conn)
	return repo.Save(s)
}

// func GetSignalEventsByCount(loadEvents int) *SignalEvents {
// 	cmd := fmt.Sprintf(`SELECT * FROM (
//         SELECT time, product_code, side, price, size FROM %s WHERE product_code = ? ORDER BY time DESC LIMIT ? )
//         ORDER BY time ASC;`, tableNameSignalEvents)
// 	rows, err := DbConnection.Query(cmd, config.Config.ProductCode, loadEvents)
// 	if err != nil {
// 		return nil
// 	}
// 	defer rows.Close()

// 	var signalEvents SignalEvents
// 	for rows.Next() {
// 		var signalEvent SignalEvent
// 		rows.Scan(&signalEvent.Time, &signalEvent.ProductCode, &signalEvent.Side, &signalEvent.Price, &signalEvent.Size)
// 		signalEvents.Signals = append(signalEvents.Signals, signalEvent)
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		return nil
// 	}
// 	return &signalEvents
// }

// func GetSignalEventsAfterTime(timeTime time.Time) *SignalEvents {
// 	cmd := fmt.Sprintf(`SELECT * FROM (
//                 SELECT time, product_code, side, price, size FROM %s
//                 WHERE DATETIME(time) >= DATETIME(?)
//                 ORDER BY time DESC
//             ) ORDER BY time ASC;`, tableNameSignalEvents)
// 	rows, err := DbConnection.Query(cmd, timeTime.Format(time.RFC3339))
// 	if err != nil {
// 		return nil
// 	}
// 	defer rows.Close()

// 	var signalEvents SignalEvents
// 	for rows.Next() {
// 		var signalEvent SignalEvent
// 		rows.Scan(&signalEvent.Time, &signalEvent.ProductCode, &signalEvent.Side, &signalEvent.Price, &signalEvent.Size)
// 		signalEvents.Signals = append(signalEvents.Signals, signalEvent)
// 	}
// 	return &signalEvents
// }

// func (s *SignalEvents) CanBuy(time time.Time) bool {
// 	lenSignals := len(s.Signals)
// 	if lenSignals == 0 {
// 		return true
// 	}

// 	lastSignal := s.Signals[lenSignals-1]
// 	if lastSignal.Side == "SELL" && lastSignal.Time.Before(time) {
// 		return true
// 	}
// 	return false
// }

// func (s *SignalEvents) CanSell(time time.Time) bool {
// 	lenSignals := len(s.Signals)
// 	if lenSignals == 0 {
// 		return false
// 	}

// 	lastSignal := s.Signals[lenSignals-1]
// 	if lastSignal.Side == "BUY" && lastSignal.Time.Before(time) {
// 		return true
// 	}
// 	return false
// }

func (s *SignalEvents) Buy(
	ProductCode string,
	time time.Time,
	price, float64,
	size float64,
	save bool,
) bool {

	// if !s.CanBuy(time) {
	// 	return false
	// }
	signalEvent := SignalEvent{
		ProductCode: ProductCode,
		Time:        time,
		Side:        "BUY",
		Price:       price,
		Size:        size,
	}
	if save {
		signalEvent.Save()
	}
	s.Signals = append(s.Signals, signalEvent)
	return true
}

// func (s *SignalEvents) Sell(productCode string, time time.Time, price, size float64, save bool) bool {

// 	if !s.CanSell(time) {
// 		return false
// 	}

// 	signalEvent := SignalEvent{
// 		ProductCode: productCode,
// 		Time:        time,
// 		Side:        "SELL",
// 		Price:       price,
// 		Size:        size,
// 	}
// 	if save {
// 		signalEvent.Save()
// 	}
// 	s.Signals = append(s.Signals, signalEvent)
// 	return true
// }

// func (s *SignalEvents) Profit() float64 {
// 	total := 0.0
// 	beforeSell := 0.0
// 	isHolding := false
// 	for i, signalEvent := range s.Signals {
// 		if i == 0 && signalEvent.Side == "SELL" {
// 			continue
// 		}
// 		if signalEvent.Side == "BUY" {
// 			total -= signalEvent.Price * signalEvent.Size
// 			isHolding = true
// 		}
// 		if signalEvent.Side == "SEL" {
// 			total += signalEvent.Price * signalEvent.Size
// 			isHolding = false
// 			beforeSell = total
// 		}
// 	}
// 	if isHolding == true {
// 		return beforeSell
// 	}
// 	return total
// }

// func (s SignalEvents) MarshalJSON() ([]byte, error) {
// 	value, err := json.Marshal(&struct {
// 		Signals []SignalEvent `json:"signals,omitempty"`
// 		Profit  float64       `json:"profit,omitempty"`
// 	}{
// 		Signals: s.Signals,
// 		Profit:  s.Profit(),
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return value, err
// }

// func (s *SignalEvents) CollectAfter(time time.Time) *SignalEvents {
// 	for i, signal := range s.Signals {
// 		if time.After(signal.Time) {
// 			continue
// 		}
// 		return &SignalEvents{Signals: s.Signals[i:]}
// 	}
// 	return nil
// }
