package kingpin

import (
	"net"
	"time"
)

// This file is autogenerated by "go generate .". Do not modify.

// Strings accumulates string values into a slice.
func (p *parserMixin) Strings() (target *[]string) {
	target = new([]string)
	p.StringsVar(target)
	return
}

func (p *parserMixin) StringsVar(target *[]string) {
	p.SetValue(newAccumulator(target, func(v interface{}) Value { return newStringValue(v.(*string)) }))
}

// Uint64List accumulates uint64 values into a slice.
func (p *parserMixin) Uint64List() (target *[]uint64) {
	target = new([]uint64)
	p.Uint64ListVar(target)
	return
}

func (p *parserMixin) Uint64ListVar(target *[]uint64) {
	p.SetValue(newAccumulator(target, func(v interface{}) Value { return newUint64Value(v.(*uint64)) }))
}

// Int64List accumulates int64 values into a slice.
func (p *parserMixin) Int64List() (target *[]int64) {
	target = new([]int64)
	p.Int64ListVar(target)
	return
}

func (p *parserMixin) Int64ListVar(target *[]int64) {
	p.SetValue(newAccumulator(target, func(v interface{}) Value { return newInt64Value(v.(*int64)) }))
}

// DurationList accumulates time.Duration values into a slice.
func (p *parserMixin) DurationList() (target *[]time.Duration) {
	target = new([]time.Duration)
	p.DurationListVar(target)
	return
}

func (p *parserMixin) DurationListVar(target *[]time.Duration) {
	p.SetValue(newAccumulator(target, func(v interface{}) Value { return newDurationValue(v.(*time.Duration)) }))
}

// IPList accumulates net.IP values into a slice.
func (p *parserMixin) IPList() (target *[]net.IP) {
	target = new([]net.IP)
	p.IPListVar(target)
	return
}

func (p *parserMixin) IPListVar(target *[]net.IP) {
	p.SetValue(newAccumulator(target, func(v interface{}) Value { return newIPValue(v.(*net.IP)) }))
}

// TCPList accumulates *net.TCPAddr values into a slice.
func (p *parserMixin) TCPList() (target *[]*net.TCPAddr) {
	target = new([]*net.TCPAddr)
	p.TCPListVar(target)
	return
}

func (p *parserMixin) TCPListVar(target *[]*net.TCPAddr) {
	p.SetValue(newAccumulator(target, func(v interface{}) Value { return newTCPAddrValue(v.(**net.TCPAddr)) }))
}