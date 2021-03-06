// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: market_data.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type OrderBook struct {
	// / 卖价.
	Ask float64 `protobuf:"fixed64,1,opt,name=ask,proto3" json:"ask"`
	// / 卖量.
	AskVolume float64 `protobuf:"fixed64,2,opt,name=ask_volume,json=askVolume,proto3" json:"askVolume"`
	// / 买价.
	Bid float64 `protobuf:"fixed64,3,opt,name=bid,proto3" json:"bid"`
	// / 买量.
	BidVolume float64 `protobuf:"fixed64,4,opt,name=bid_volume,json=bidVolume,proto3" json:"bidVolume"`
}

func (m *OrderBook) Reset()                    { *m = OrderBook{} }
func (m *OrderBook) String() string            { return proto.CompactTextString(m) }
func (*OrderBook) ProtoMessage()               {}
func (*OrderBook) Descriptor() ([]byte, []int) { return fileDescriptorMarketData, []int{0} }

func (m *OrderBook) GetAsk() float64 {
	if m != nil {
		return m.Ask
	}
	return 0
}

func (m *OrderBook) GetAskVolume() float64 {
	if m != nil {
		return m.AskVolume
	}
	return 0
}

func (m *OrderBook) GetBid() float64 {
	if m != nil {
		return m.Bid
	}
	return 0
}

func (m *OrderBook) GetBidVolume() float64 {
	if m != nil {
		return m.BidVolume
	}
	return 0
}

// MarketDataSnapshotLevel2 level2 tick
type MarketDataSnapshot struct {
	// 合约
	Symbol Symbol `protobuf:"bytes,1,opt,name=symbol" json:"symbol"`
	// 时间
	Time int64 `protobuf:"varint,2,opt,name=time,proto3" json:"time"`
	// 毫秒
	Milliseconds int32 `protobuf:"varint,3,opt,name=milliseconds,proto3" json:"milliseconds"`
	// 开盘
	Open float64 `protobuf:"fixed64,4,opt,name=open,proto3" json:"open"`
	// 最高
	High float64 `protobuf:"fixed64,5,opt,name=high,proto3" json:"high"`
	// 最低
	Low float64 `protobuf:"fixed64,6,opt,name=low,proto3" json:"low"`
	// 收盘
	Close float64 `protobuf:"fixed64,7,opt,name=close,proto3" json:"close"`
	// 成交量
	Volume float64 `protobuf:"fixed64,8,opt,name=volume,proto3" json:"volume"`
	// 成交金额
	Amount float64 `protobuf:"fixed64,9,opt,name=amount,proto3" json:"amount"`
	// 持仓
	Position float64 `protobuf:"fixed64,10,opt,name=position,proto3" json:"position"`
	// 最新价格
	Price float64 `protobuf:"fixed64,11,opt,name=price,proto3" json:"price"`
	// 昨收
	PreClose float64 `protobuf:"fixed64,12,opt,name=pre_close,json=preClose,proto3" json:"preClose"`
	// 昨结
	PreSettlementPrice float64 `protobuf:"fixed64,13,opt,name=pre_settlement_price,json=preSettlementPrice,proto3" json:"preSettlementPrice"`
	// 昨持仓
	PrePosition float64 `protobuf:"fixed64,14,opt,name=pre_position,json=prePosition,proto3" json:"prePosition"`
	// 结算价
	SettlementPrice float64 `protobuf:"fixed64,15,opt,name=settlement_price,json=settlementPrice,proto3" json:"settlementPrice"`
	// 涨停
	UpperLimitPrice float64 `protobuf:"fixed64,16,opt,name=upper_limit_price,json=upperLimitPrice,proto3" json:"upperLimitPrice"`
	// 跌停
	LowerLimitPrice float64 `protobuf:"fixed64,17,opt,name=lower_limit_price,json=lowerLimitPrice,proto3" json:"lowerLimitPrice"`
	// 昨虚实
	PreDelta float64 `protobuf:"fixed64,18,opt,name=pre_delta,json=preDelta,proto3" json:"preDelta"`
	// 今虚实
	Delta float64 `protobuf:"fixed64,19,opt,name=delta,proto3" json:"delta"`
	// 均价
	AveragePrice float64 `protobuf:"fixed64,20,opt,name=average_price,json=averagePrice,proto3" json:"averagePrice"`
	// 交易日
	TradingDay int32 `protobuf:"varint,21,opt,name=trading_day,json=tradingDay,proto3" json:"tradingDay"`
	// 盘口
	OrderBookList []OrderBook `protobuf:"bytes,22,rep,name=order_book_list,json=orderBookList" json:"orderBookList"`
}

func (m *MarketDataSnapshot) Reset()                    { *m = MarketDataSnapshot{} }
func (m *MarketDataSnapshot) String() string            { return proto.CompactTextString(m) }
func (*MarketDataSnapshot) ProtoMessage()               {}
func (*MarketDataSnapshot) Descriptor() ([]byte, []int) { return fileDescriptorMarketData, []int{1} }

func (m *MarketDataSnapshot) GetSymbol() Symbol {
	if m != nil {
		return m.Symbol
	}
	return Symbol{}
}

func (m *MarketDataSnapshot) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *MarketDataSnapshot) GetMilliseconds() int32 {
	if m != nil {
		return m.Milliseconds
	}
	return 0
}

func (m *MarketDataSnapshot) GetOpen() float64 {
	if m != nil {
		return m.Open
	}
	return 0
}

func (m *MarketDataSnapshot) GetHigh() float64 {
	if m != nil {
		return m.High
	}
	return 0
}

func (m *MarketDataSnapshot) GetLow() float64 {
	if m != nil {
		return m.Low
	}
	return 0
}

func (m *MarketDataSnapshot) GetClose() float64 {
	if m != nil {
		return m.Close
	}
	return 0
}

func (m *MarketDataSnapshot) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *MarketDataSnapshot) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *MarketDataSnapshot) GetPosition() float64 {
	if m != nil {
		return m.Position
	}
	return 0
}

func (m *MarketDataSnapshot) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *MarketDataSnapshot) GetPreClose() float64 {
	if m != nil {
		return m.PreClose
	}
	return 0
}

func (m *MarketDataSnapshot) GetPreSettlementPrice() float64 {
	if m != nil {
		return m.PreSettlementPrice
	}
	return 0
}

func (m *MarketDataSnapshot) GetPrePosition() float64 {
	if m != nil {
		return m.PrePosition
	}
	return 0
}

func (m *MarketDataSnapshot) GetSettlementPrice() float64 {
	if m != nil {
		return m.SettlementPrice
	}
	return 0
}

func (m *MarketDataSnapshot) GetUpperLimitPrice() float64 {
	if m != nil {
		return m.UpperLimitPrice
	}
	return 0
}

func (m *MarketDataSnapshot) GetLowerLimitPrice() float64 {
	if m != nil {
		return m.LowerLimitPrice
	}
	return 0
}

func (m *MarketDataSnapshot) GetPreDelta() float64 {
	if m != nil {
		return m.PreDelta
	}
	return 0
}

func (m *MarketDataSnapshot) GetDelta() float64 {
	if m != nil {
		return m.Delta
	}
	return 0
}

func (m *MarketDataSnapshot) GetAveragePrice() float64 {
	if m != nil {
		return m.AveragePrice
	}
	return 0
}

func (m *MarketDataSnapshot) GetTradingDay() int32 {
	if m != nil {
		return m.TradingDay
	}
	return 0
}

func (m *MarketDataSnapshot) GetOrderBookList() []OrderBook {
	if m != nil {
		return m.OrderBookList
	}
	return nil
}

// MdsList 行情列表
type MdsList struct {
	List []MarketDataSnapshot `protobuf:"bytes,1,rep,name=list" json:"list"`
}

func (m *MdsList) Reset()                    { *m = MdsList{} }
func (m *MdsList) String() string            { return proto.CompactTextString(m) }
func (*MdsList) ProtoMessage()               {}
func (*MdsList) Descriptor() ([]byte, []int) { return fileDescriptorMarketData, []int{2} }

func (m *MdsList) GetList() []MarketDataSnapshot {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*OrderBook)(nil), "pb.OrderBook")
	proto.RegisterType((*MarketDataSnapshot)(nil), "pb.MarketDataSnapshot")
	proto.RegisterType((*MdsList)(nil), "pb.MdsList")
}
func (m *OrderBook) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrderBook) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Ask != 0 {
		dAtA[i] = 0x9
		i++
		i = encodeFixed64MarketData(dAtA, i, uint64(math.Float64bits(float64(m.Ask))))
	}
	if m.AskVolume != 0 {
		dAtA[i] = 0x11
		i++
		i = encodeFixed64MarketData(dAtA, i, uint64(math.Float64bits(float64(m.AskVolume))))
	}
	if m.Bid != 0 {
		dAtA[i] = 0x19
		i++
		i = encodeFixed64MarketData(dAtA, i, uint64(math.Float64bits(float64(m.Bid))))
	}
	if m.BidVolume != 0 {
		dAtA[i] = 0x21
		i++
		i = encodeFixed64MarketData(dAtA, i, uint64(math.Float64bits(float64(m.BidVolume))))
	}
	return i, nil
}

func (m *MarketDataSnapshot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MarketDataSnapshot) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintMarketData(dAtA, i, uint64(m.Symbol.Size()))
	n1, err := m.Symbol.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.Time != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMarketData(dAtA, i, uint64(m.Time))
	}
	if m.Milliseconds != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMarketData(dAtA, i, uint64(m.Milliseconds))
	}
	if m.Open != 0 {
		dAtA[i] = 0x21
		i++
		i = encodeFixed64MarketData(dAtA, i, uint64(math.Float64bits(float64(m.Open))))
	}
	if m.High != 0 {
		dAtA[i] = 0x29
		i++
		i = encodeFixed64MarketData(dAtA, i, uint64(math.Float64bits(float64(m.High))))
	}
	if m.Low != 0 {
		dAtA[i] = 0x31
		i++
		i = encodeFixed64MarketData(dAtA, i, uint64(math.Float64bits(float64(m.Low))))
	}
	if m.Close != 0 {
		dAtA[i] = 0x39
		i++
		i = encodeFixed64MarketData(dAtA, i, uint64(math.Float64bits(float64(m.Close))))
	}
	if m.Volume != 0 {
		dAtA[i] = 0x41
		i++
		i = encodeFixed64MarketData(dAtA, i, uint64(math.Float64bits(float64(m.Volume))))
	}
	if m.Amount != 0 {
		dAtA[i] = 0x49
		i++
		i = encodeFixed64MarketData(dAtA, i, uint64(math.Float64bits(float64(m.Amount))))
	}
	if m.Position != 0 {
		dAtA[i] = 0x51
		i++
		i = encodeFixed64MarketData(dAtA, i, uint64(math.Float64bits(float64(m.Position))))
	}
	if m.Price != 0 {
		dAtA[i] = 0x59
		i++
		i = encodeFixed64MarketData(dAtA, i, uint64(math.Float64bits(float64(m.Price))))
	}
	if m.PreClose != 0 {
		dAtA[i] = 0x61
		i++
		i = encodeFixed64MarketData(dAtA, i, uint64(math.Float64bits(float64(m.PreClose))))
	}
	if m.PreSettlementPrice != 0 {
		dAtA[i] = 0x69
		i++
		i = encodeFixed64MarketData(dAtA, i, uint64(math.Float64bits(float64(m.PreSettlementPrice))))
	}
	if m.PrePosition != 0 {
		dAtA[i] = 0x71
		i++
		i = encodeFixed64MarketData(dAtA, i, uint64(math.Float64bits(float64(m.PrePosition))))
	}
	if m.SettlementPrice != 0 {
		dAtA[i] = 0x79
		i++
		i = encodeFixed64MarketData(dAtA, i, uint64(math.Float64bits(float64(m.SettlementPrice))))
	}
	if m.UpperLimitPrice != 0 {
		dAtA[i] = 0x81
		i++
		dAtA[i] = 0x1
		i++
		i = encodeFixed64MarketData(dAtA, i, uint64(math.Float64bits(float64(m.UpperLimitPrice))))
	}
	if m.LowerLimitPrice != 0 {
		dAtA[i] = 0x89
		i++
		dAtA[i] = 0x1
		i++
		i = encodeFixed64MarketData(dAtA, i, uint64(math.Float64bits(float64(m.LowerLimitPrice))))
	}
	if m.PreDelta != 0 {
		dAtA[i] = 0x91
		i++
		dAtA[i] = 0x1
		i++
		i = encodeFixed64MarketData(dAtA, i, uint64(math.Float64bits(float64(m.PreDelta))))
	}
	if m.Delta != 0 {
		dAtA[i] = 0x99
		i++
		dAtA[i] = 0x1
		i++
		i = encodeFixed64MarketData(dAtA, i, uint64(math.Float64bits(float64(m.Delta))))
	}
	if m.AveragePrice != 0 {
		dAtA[i] = 0xa1
		i++
		dAtA[i] = 0x1
		i++
		i = encodeFixed64MarketData(dAtA, i, uint64(math.Float64bits(float64(m.AveragePrice))))
	}
	if m.TradingDay != 0 {
		dAtA[i] = 0xa8
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintMarketData(dAtA, i, uint64(m.TradingDay))
	}
	if len(m.OrderBookList) > 0 {
		for _, msg := range m.OrderBookList {
			dAtA[i] = 0xb2
			i++
			dAtA[i] = 0x1
			i++
			i = encodeVarintMarketData(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *MdsList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MdsList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.List) > 0 {
		for _, msg := range m.List {
			dAtA[i] = 0xa
			i++
			i = encodeVarintMarketData(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64MarketData(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32MarketData(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMarketData(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *OrderBook) Size() (n int) {
	var l int
	_ = l
	if m.Ask != 0 {
		n += 9
	}
	if m.AskVolume != 0 {
		n += 9
	}
	if m.Bid != 0 {
		n += 9
	}
	if m.BidVolume != 0 {
		n += 9
	}
	return n
}

func (m *MarketDataSnapshot) Size() (n int) {
	var l int
	_ = l
	l = m.Symbol.Size()
	n += 1 + l + sovMarketData(uint64(l))
	if m.Time != 0 {
		n += 1 + sovMarketData(uint64(m.Time))
	}
	if m.Milliseconds != 0 {
		n += 1 + sovMarketData(uint64(m.Milliseconds))
	}
	if m.Open != 0 {
		n += 9
	}
	if m.High != 0 {
		n += 9
	}
	if m.Low != 0 {
		n += 9
	}
	if m.Close != 0 {
		n += 9
	}
	if m.Volume != 0 {
		n += 9
	}
	if m.Amount != 0 {
		n += 9
	}
	if m.Position != 0 {
		n += 9
	}
	if m.Price != 0 {
		n += 9
	}
	if m.PreClose != 0 {
		n += 9
	}
	if m.PreSettlementPrice != 0 {
		n += 9
	}
	if m.PrePosition != 0 {
		n += 9
	}
	if m.SettlementPrice != 0 {
		n += 9
	}
	if m.UpperLimitPrice != 0 {
		n += 10
	}
	if m.LowerLimitPrice != 0 {
		n += 10
	}
	if m.PreDelta != 0 {
		n += 10
	}
	if m.Delta != 0 {
		n += 10
	}
	if m.AveragePrice != 0 {
		n += 10
	}
	if m.TradingDay != 0 {
		n += 2 + sovMarketData(uint64(m.TradingDay))
	}
	if len(m.OrderBookList) > 0 {
		for _, e := range m.OrderBookList {
			l = e.Size()
			n += 2 + l + sovMarketData(uint64(l))
		}
	}
	return n
}

func (m *MdsList) Size() (n int) {
	var l int
	_ = l
	if len(m.List) > 0 {
		for _, e := range m.List {
			l = e.Size()
			n += 1 + l + sovMarketData(uint64(l))
		}
	}
	return n
}

func sovMarketData(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMarketData(x uint64) (n int) {
	return sovMarketData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OrderBook) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMarketData
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OrderBook: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrderBook: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ask", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.Ask = float64(math.Float64frombits(v))
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field AskVolume", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.AskVolume = float64(math.Float64frombits(v))
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bid", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.Bid = float64(math.Float64frombits(v))
		case 4:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field BidVolume", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.BidVolume = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipMarketData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMarketData
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MarketDataSnapshot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMarketData
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MarketDataSnapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MarketDataSnapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMarketData
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Symbol.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Milliseconds", wireType)
			}
			m.Milliseconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Milliseconds |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Open", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.Open = float64(math.Float64frombits(v))
		case 5:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field High", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.High = float64(math.Float64frombits(v))
		case 6:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Low", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.Low = float64(math.Float64frombits(v))
		case 7:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Close", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.Close = float64(math.Float64frombits(v))
		case 8:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Volume", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.Volume = float64(math.Float64frombits(v))
		case 9:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.Amount = float64(math.Float64frombits(v))
		case 10:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Position", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.Position = float64(math.Float64frombits(v))
		case 11:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.Price = float64(math.Float64frombits(v))
		case 12:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreClose", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.PreClose = float64(math.Float64frombits(v))
		case 13:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreSettlementPrice", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.PreSettlementPrice = float64(math.Float64frombits(v))
		case 14:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrePosition", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.PrePosition = float64(math.Float64frombits(v))
		case 15:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field SettlementPrice", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.SettlementPrice = float64(math.Float64frombits(v))
		case 16:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpperLimitPrice", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.UpperLimitPrice = float64(math.Float64frombits(v))
		case 17:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field LowerLimitPrice", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.LowerLimitPrice = float64(math.Float64frombits(v))
		case 18:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreDelta", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.PreDelta = float64(math.Float64frombits(v))
		case 19:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delta", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.Delta = float64(math.Float64frombits(v))
		case 20:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field AveragePrice", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.AveragePrice = float64(math.Float64frombits(v))
		case 21:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradingDay", wireType)
			}
			m.TradingDay = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TradingDay |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 22:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderBookList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMarketData
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderBookList = append(m.OrderBookList, OrderBook{})
			if err := m.OrderBookList[len(m.OrderBookList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMarketData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMarketData
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MdsList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMarketData
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MdsList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MdsList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMarketData
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.List = append(m.List, MarketDataSnapshot{})
			if err := m.List[len(m.List)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMarketData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMarketData
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMarketData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMarketData
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMarketData
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMarketData
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMarketData
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMarketData
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMarketData(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMarketData = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMarketData   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("market_data.proto", fileDescriptorMarketData) }

var fileDescriptorMarketData = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x53, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0x66, 0x9b, 0x9f, 0x36, 0x93, 0x84, 0x24, 0x26, 0x44, 0x56, 0x11, 0xa1, 0x84, 0x4b, 0xe9,
	0x21, 0x42, 0x45, 0x7d, 0x81, 0x92, 0x63, 0x2b, 0xaa, 0x8d, 0xc4, 0x75, 0xe5, 0xcd, 0x5a, 0x89,
	0xb5, 0xde, 0xb5, 0x65, 0x3b, 0x8d, 0xf2, 0x16, 0xbc, 0x15, 0x1c, 0x79, 0x04, 0x14, 0x5e, 0x04,
	0x79, 0xbc, 0x59, 0x41, 0x7b, 0x9b, 0xf9, 0x7e, 0xe6, 0xc7, 0xf2, 0xc0, 0xa8, 0x60, 0x26, 0xe7,
	0x2e, 0xc9, 0x98, 0x63, 0x73, 0x6d, 0x94, 0x53, 0xe4, 0x44, 0xa7, 0xe7, 0xbd, 0x95, 0x2a, 0x0a,
	0x55, 0x06, 0x64, 0x56, 0x40, 0xe7, 0xab, 0xc9, 0xb8, 0xb9, 0x55, 0x2a, 0x27, 0x43, 0x68, 0x30,
	0x9b, 0xd3, 0xe8, 0x22, 0xba, 0x8c, 0x62, 0x1f, 0x92, 0xb7, 0x00, 0xcc, 0xe6, 0xc9, 0xa3, 0x92,
	0xdb, 0x82, 0xd3, 0x13, 0x24, 0x3a, 0xcc, 0xe6, 0xdf, 0x10, 0xf0, 0x86, 0x54, 0x64, 0xb4, 0x11,
	0x0c, 0xa9, 0xc8, 0xbc, 0x21, 0x15, 0xd9, 0xd1, 0xd0, 0x0c, 0x86, 0x54, 0x64, 0xc1, 0x30, 0xfb,
	0xd1, 0x02, 0x72, 0x8f, 0x63, 0x2d, 0x98, 0x63, 0xcb, 0x92, 0x69, 0xbb, 0x51, 0x8e, 0xcc, 0xa0,
	0x6d, 0xf7, 0x45, 0xaa, 0x24, 0xf6, 0xee, 0x5e, 0xc3, 0x5c, 0xa7, 0xf3, 0x25, 0x22, 0x71, 0xc5,
	0x10, 0x02, 0x4d, 0x27, 0xaa, 0x21, 0x1a, 0x31, 0xc6, 0x64, 0x06, 0xbd, 0x42, 0x48, 0x29, 0x2c,
	0x5f, 0xa9, 0x32, 0xb3, 0x38, 0x48, 0x2b, 0xfe, 0x0f, 0xf3, 0x3e, 0xa5, 0x79, 0x59, 0xcd, 0x82,
	0xb1, 0xc7, 0x36, 0x62, 0xbd, 0xa1, 0xad, 0x80, 0xf9, 0xd8, 0xef, 0x22, 0xd5, 0x8e, 0xb6, 0xc3,
	0x2e, 0x52, 0xed, 0xc8, 0x18, 0x5a, 0x2b, 0xa9, 0x2c, 0xa7, 0xa7, 0x88, 0x85, 0x84, 0x4c, 0xa0,
	0x5d, 0x6d, 0x77, 0x86, 0x70, 0x95, 0x79, 0x9c, 0x15, 0x6a, 0x5b, 0x3a, 0xda, 0x09, 0x78, 0xc8,
	0xc8, 0x39, 0x9c, 0x69, 0x65, 0x85, 0x13, 0xaa, 0xa4, 0x80, 0x4c, 0x9d, 0xfb, 0x0e, 0xda, 0x88,
	0x15, 0xa7, 0xdd, 0xd0, 0x01, 0x13, 0xf2, 0x06, 0x3a, 0xda, 0xf0, 0x24, 0xf4, 0xee, 0x55, 0x16,
	0xc3, 0xbf, 0x60, 0xfb, 0x4f, 0x30, 0xf6, 0xa4, 0xe5, 0xce, 0x49, 0x5e, 0xf0, 0xd2, 0x25, 0xa1,
	0x42, 0x1f, 0x75, 0x44, 0x1b, 0xbe, 0xac, 0xa9, 0x07, 0x2c, 0xf7, 0x1e, 0x7a, 0xde, 0x51, 0x0f,
	0xf1, 0x12, 0x95, 0x5d, 0x6d, 0xf8, 0xc3, 0x71, 0x8e, 0x8f, 0x30, 0x7c, 0x56, 0x70, 0x80, 0xb2,
	0x81, 0x7d, 0x52, 0xed, 0x0a, 0x46, 0x5b, 0xad, 0xb9, 0x49, 0xa4, 0x28, 0xc4, 0x51, 0x3b, 0x0c,
	0x5a, 0x24, 0xee, 0x3c, 0x5e, 0x6b, 0xa5, 0xda, 0x3d, 0xd1, 0x8e, 0x82, 0x16, 0x89, 0x7f, 0xb4,
	0xd5, 0xd2, 0x19, 0x97, 0x8e, 0x51, 0x52, 0x2f, 0xbd, 0xf0, 0xb9, 0x7f, 0xa7, 0x40, 0xbc, 0x0a,
	0xef, 0x84, 0x09, 0xf9, 0x00, 0x7d, 0xf6, 0xc8, 0x0d, 0x5b, 0xf3, 0xaa, 0xf4, 0x18, 0xd9, 0x5e,
	0x05, 0x86, 0xba, 0xef, 0xa0, 0xeb, 0x0c, 0xcb, 0x44, 0xb9, 0x4e, 0x32, 0xb6, 0xa7, 0xaf, 0xf1,
	0x87, 0x40, 0x05, 0x2d, 0xd8, 0x9e, 0xdc, 0xc0, 0x40, 0xf9, 0x0b, 0x48, 0x52, 0xa5, 0xf2, 0x44,
	0x0a, 0xeb, 0xe8, 0xe4, 0xa2, 0x71, 0xd9, 0xbd, 0xee, 0xfb, 0x4f, 0x58, 0x1f, 0x47, 0xdc, 0x57,
	0xc7, 0xf0, 0x4e, 0x58, 0x37, 0xbb, 0x81, 0xd3, 0xfb, 0xcc, 0xfa, 0x90, 0x5c, 0x41, 0x13, 0x6d,
	0x11, 0xda, 0x26, 0xde, 0xf6, 0xfc, 0x8f, 0xc7, 0xa8, 0xb9, 0x1d, 0xfe, 0x3c, 0x4c, 0xa3, 0x5f,
	0x87, 0x69, 0xf4, 0xfb, 0x30, 0x8d, 0xbe, 0xff, 0x99, 0xbe, 0x48, 0xdb, 0x78, 0x88, 0x9f, 0xff,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x8c, 0xe1, 0xb7, 0xb6, 0xaf, 0x03, 0x00, 0x00,
}
