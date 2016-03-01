//Package massquoteacknowledgement msg type = b.
package massquoteacknowledgement

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/parties"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
	"time"
)

//NoQuoteSets is a repeating group in MassQuoteAcknowledgement
type NoQuoteSets struct {
	//QuoteSetID is a non-required field for NoQuoteSets.
	QuoteSetID *string `fix:"302"`
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
	//TotNoQuoteEntries is a non-required field for NoQuoteSets.
	TotNoQuoteEntries *int `fix:"304"`
	//LastFragment is a non-required field for NoQuoteSets.
	LastFragment *bool `fix:"893"`
	//NoQuoteEntries is a non-required field for NoQuoteSets.
	NoQuoteEntries []NoQuoteEntries `fix:"295,omitempty"`
}

//NoQuoteEntries is a repeating group in NoQuoteSets
type NoQuoteEntries struct {
	//QuoteEntryID is a non-required field for NoQuoteEntries.
	QuoteEntryID *string `fix:"299"`
	//Instrument Component
	Instrument instrument.Component
	//NoLegs is a non-required field for NoQuoteEntries.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//BidPx is a non-required field for NoQuoteEntries.
	BidPx *float64 `fix:"132"`
	//OfferPx is a non-required field for NoQuoteEntries.
	OfferPx *float64 `fix:"133"`
	//BidSize is a non-required field for NoQuoteEntries.
	BidSize *float64 `fix:"134"`
	//OfferSize is a non-required field for NoQuoteEntries.
	OfferSize *float64 `fix:"135"`
	//ValidUntilTime is a non-required field for NoQuoteEntries.
	ValidUntilTime *time.Time `fix:"62"`
	//BidSpotRate is a non-required field for NoQuoteEntries.
	BidSpotRate *float64 `fix:"188"`
	//OfferSpotRate is a non-required field for NoQuoteEntries.
	OfferSpotRate *float64 `fix:"190"`
	//BidForwardPoints is a non-required field for NoQuoteEntries.
	BidForwardPoints *float64 `fix:"189"`
	//OfferForwardPoints is a non-required field for NoQuoteEntries.
	OfferForwardPoints *float64 `fix:"191"`
	//MidPx is a non-required field for NoQuoteEntries.
	MidPx *float64 `fix:"631"`
	//BidYield is a non-required field for NoQuoteEntries.
	BidYield *float64 `fix:"632"`
	//MidYield is a non-required field for NoQuoteEntries.
	MidYield *float64 `fix:"633"`
	//OfferYield is a non-required field for NoQuoteEntries.
	OfferYield *float64 `fix:"634"`
	//TransactTime is a non-required field for NoQuoteEntries.
	TransactTime *time.Time `fix:"60"`
	//TradingSessionID is a non-required field for NoQuoteEntries.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoQuoteEntries.
	TradingSessionSubID *string `fix:"625"`
	//SettlDate is a non-required field for NoQuoteEntries.
	SettlDate *string `fix:"64"`
	//OrdType is a non-required field for NoQuoteEntries.
	OrdType *string `fix:"40"`
	//SettlDate2 is a non-required field for NoQuoteEntries.
	SettlDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for NoQuoteEntries.
	OrderQty2 *float64 `fix:"192"`
	//BidForwardPoints2 is a non-required field for NoQuoteEntries.
	BidForwardPoints2 *float64 `fix:"642"`
	//OfferForwardPoints2 is a non-required field for NoQuoteEntries.
	OfferForwardPoints2 *float64 `fix:"643"`
	//Currency is a non-required field for NoQuoteEntries.
	Currency *string `fix:"15"`
	//QuoteEntryRejectReason is a non-required field for NoQuoteEntries.
	QuoteEntryRejectReason *int `fix:"368"`
}

//NoLegs is a repeating group in NoQuoteEntries
type NoLegs struct {
	//InstrumentLeg Component
	InstrumentLeg instrumentleg.Component
}

//Message is a MassQuoteAcknowledgement FIX Message
type Message struct {
	FIXMsgType string `fix:"b"`
	Header     fix44.Header
	//QuoteReqID is a non-required field for MassQuoteAcknowledgement.
	QuoteReqID *string `fix:"131"`
	//QuoteID is a non-required field for MassQuoteAcknowledgement.
	QuoteID *string `fix:"117"`
	//QuoteStatus is a required field for MassQuoteAcknowledgement.
	QuoteStatus int `fix:"297"`
	//QuoteRejectReason is a non-required field for MassQuoteAcknowledgement.
	QuoteRejectReason *int `fix:"300"`
	//QuoteResponseLevel is a non-required field for MassQuoteAcknowledgement.
	QuoteResponseLevel *int `fix:"301"`
	//QuoteType is a non-required field for MassQuoteAcknowledgement.
	QuoteType *int `fix:"537"`
	//Parties Component
	Parties parties.Component
	//Account is a non-required field for MassQuoteAcknowledgement.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for MassQuoteAcknowledgement.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for MassQuoteAcknowledgement.
	AccountType *int `fix:"581"`
	//Text is a non-required field for MassQuoteAcknowledgement.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for MassQuoteAcknowledgement.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for MassQuoteAcknowledgement.
	EncodedText *string `fix:"355"`
	//NoQuoteSets is a non-required field for MassQuoteAcknowledgement.
	NoQuoteSets []NoQuoteSets `fix:"296,omitempty"`
	Trailer     fix44.Trailer
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
	return enum.BeginStringFIX44, "b", r
}
