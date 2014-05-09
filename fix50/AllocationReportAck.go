package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//AllocationReportAck msg type = AT.
type AllocationReportAck struct {
	message.Message
}

//AllocationReportAckBuilder builds AllocationReportAck messages.
type AllocationReportAckBuilder struct {
	message.MessageBuilder
}

//CreateAllocationReportAckBuilder returns an initialized AllocationReportAckBuilder with specified required fields.
func CreateAllocationReportAckBuilder(
	allocreportid *field.AllocReportIDField,
	allocid *field.AllocIDField) AllocationReportAckBuilder {
	var builder AllocationReportAckBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50))
	builder.Header.Set(field.NewMsgType("AT"))
	builder.Body.Set(allocreportid)
	builder.Body.Set(allocid)
	return builder
}

//AllocReportID is a required field for AllocationReportAck.
func (m AllocationReportAck) AllocReportID() (*field.AllocReportIDField, errors.MessageRejectError) {
	f := &field.AllocReportIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocReportID reads a AllocReportID from AllocationReportAck.
func (m AllocationReportAck) GetAllocReportID(f *field.AllocReportIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocID is a required field for AllocationReportAck.
func (m AllocationReportAck) AllocID() (*field.AllocIDField, errors.MessageRejectError) {
	f := &field.AllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from AllocationReportAck.
func (m AllocationReportAck) GetAllocID(f *field.AllocIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for AllocationReportAck.
func (m AllocationReportAck) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from AllocationReportAck.
func (m AllocationReportAck) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryAllocID is a non-required field for AllocationReportAck.
func (m AllocationReportAck) SecondaryAllocID() (*field.SecondaryAllocIDField, errors.MessageRejectError) {
	f := &field.SecondaryAllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryAllocID reads a SecondaryAllocID from AllocationReportAck.
func (m AllocationReportAck) GetSecondaryAllocID(f *field.SecondaryAllocIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for AllocationReportAck.
func (m AllocationReportAck) TradeDate() (*field.TradeDateField, errors.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from AllocationReportAck.
func (m AllocationReportAck) GetTradeDate(f *field.TradeDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for AllocationReportAck.
func (m AllocationReportAck) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from AllocationReportAck.
func (m AllocationReportAck) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocStatus is a non-required field for AllocationReportAck.
func (m AllocationReportAck) AllocStatus() (*field.AllocStatusField, errors.MessageRejectError) {
	f := &field.AllocStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocStatus reads a AllocStatus from AllocationReportAck.
func (m AllocationReportAck) GetAllocStatus(f *field.AllocStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocRejCode is a non-required field for AllocationReportAck.
func (m AllocationReportAck) AllocRejCode() (*field.AllocRejCodeField, errors.MessageRejectError) {
	f := &field.AllocRejCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocRejCode reads a AllocRejCode from AllocationReportAck.
func (m AllocationReportAck) GetAllocRejCode(f *field.AllocRejCodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocReportType is a non-required field for AllocationReportAck.
func (m AllocationReportAck) AllocReportType() (*field.AllocReportTypeField, errors.MessageRejectError) {
	f := &field.AllocReportTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocReportType reads a AllocReportType from AllocationReportAck.
func (m AllocationReportAck) GetAllocReportType(f *field.AllocReportTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocIntermedReqType is a non-required field for AllocationReportAck.
func (m AllocationReportAck) AllocIntermedReqType() (*field.AllocIntermedReqTypeField, errors.MessageRejectError) {
	f := &field.AllocIntermedReqTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocIntermedReqType reads a AllocIntermedReqType from AllocationReportAck.
func (m AllocationReportAck) GetAllocIntermedReqType(f *field.AllocIntermedReqTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchStatus is a non-required field for AllocationReportAck.
func (m AllocationReportAck) MatchStatus() (*field.MatchStatusField, errors.MessageRejectError) {
	f := &field.MatchStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMatchStatus reads a MatchStatus from AllocationReportAck.
func (m AllocationReportAck) GetMatchStatus(f *field.MatchStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for AllocationReportAck.
func (m AllocationReportAck) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from AllocationReportAck.
func (m AllocationReportAck) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for AllocationReportAck.
func (m AllocationReportAck) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from AllocationReportAck.
func (m AllocationReportAck) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for AllocationReportAck.
func (m AllocationReportAck) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from AllocationReportAck.
func (m AllocationReportAck) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for AllocationReportAck.
func (m AllocationReportAck) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from AllocationReportAck.
func (m AllocationReportAck) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for AllocationReportAck.
func (m AllocationReportAck) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from AllocationReportAck.
func (m AllocationReportAck) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAllocs is a non-required field for AllocationReportAck.
func (m AllocationReportAck) NoAllocs() (*field.NoAllocsField, errors.MessageRejectError) {
	f := &field.NoAllocsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoAllocs reads a NoAllocs from AllocationReportAck.
func (m AllocationReportAck) GetNoAllocs(f *field.NoAllocsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a non-required field for AllocationReportAck.
func (m AllocationReportAck) ClearingBusinessDate() (*field.ClearingBusinessDateField, errors.MessageRejectError) {
	f := &field.ClearingBusinessDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from AllocationReportAck.
func (m AllocationReportAck) GetClearingBusinessDate(f *field.ClearingBusinessDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPxIndicator is a non-required field for AllocationReportAck.
func (m AllocationReportAck) AvgPxIndicator() (*field.AvgPxIndicatorField, errors.MessageRejectError) {
	f := &field.AvgPxIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPxIndicator reads a AvgPxIndicator from AllocationReportAck.
func (m AllocationReportAck) GetAvgPxIndicator(f *field.AvgPxIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Quantity is a non-required field for AllocationReportAck.
func (m AllocationReportAck) Quantity() (*field.QuantityField, errors.MessageRejectError) {
	f := &field.QuantityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuantity reads a Quantity from AllocationReportAck.
func (m AllocationReportAck) GetQuantity(f *field.QuantityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocTransType is a non-required field for AllocationReportAck.
func (m AllocationReportAck) AllocTransType() (*field.AllocTransTypeField, errors.MessageRejectError) {
	f := &field.AllocTransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocTransType reads a AllocTransType from AllocationReportAck.
func (m AllocationReportAck) GetAllocTransType(f *field.AllocTransTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}
