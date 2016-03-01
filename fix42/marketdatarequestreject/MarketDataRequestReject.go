//Package marketdatarequestreject msg type = Y.
package marketdatarequestreject

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
)

//Message is a MarketDataRequestReject FIX Message
type Message struct {
	FIXMsgType string `fix:"Y"`
	Header     fix42.Header
	//MDReqID is a required field for MarketDataRequestReject.
	MDReqID string `fix:"262"`
	//MDReqRejReason is a non-required field for MarketDataRequestReject.
	MDReqRejReason *string `fix:"281"`
	//Text is a non-required field for MarketDataRequestReject.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for MarketDataRequestReject.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for MarketDataRequestReject.
	EncodedText *string `fix:"355"`
	Trailer     fix42.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.BeginStringFIX42, "Y", r
}
