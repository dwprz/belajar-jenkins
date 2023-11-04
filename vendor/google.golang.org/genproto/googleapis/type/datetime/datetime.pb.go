// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/type/datetime.proto

package datetime

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents civil time (or occasionally physical time).
//
// This type can represent a civil time in one of a few possible ways:
//
//   - When utc_offset is set and time_zone is unset: a civil time on a calendar
//     day with a particular offset from UTC.
//   - When time_zone is set and utc_offset is unset: a civil time on a calendar
//     day in a particular time zone.
//   - When neither time_zone nor utc_offset is set: a civil time on a calendar
//     day in local time.
//
// The date is relative to the Proleptic Gregorian Calendar.
//
// If year is 0, the DateTime is considered not to have a specific year. month
// and day must have valid, non-zero values.
//
// This type may also be used to represent a physical time if all the date and
// time fields are set and either case of the `time_offset` oneof is set.
// Consider using `Timestamp` message for physical time instead. If your use
// case also would like to store the user's timezone, that can be done in
// another field.
//
// This type is more flexible than some applications may want. Make sure to
// document and validate your application's limitations.
type DateTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. Year of date. Must be from 1 to 9999, or 0 if specifying a
	// datetime without a year.
	Year int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	// Required. Month of year. Must be from 1 to 12.
	Month int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	// Required. Day of month. Must be from 1 to 31 and valid for the year and
	// month.
	Day int32 `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
	// Required. Hours of day in 24 hour format. Should be from 0 to 23. An API
	// may choose to allow the value "24:00:00" for scenarios like business
	// closing time.
	Hours int32 `protobuf:"varint,4,opt,name=hours,proto3" json:"hours,omitempty"`
	// Required. Minutes of hour of day. Must be from 0 to 59.
	Minutes int32 `protobuf:"varint,5,opt,name=minutes,proto3" json:"minutes,omitempty"`
	// Required. Seconds of minutes of the time. Must normally be from 0 to 59. An
	// API may allow the value 60 if it allows leap-seconds.
	Seconds int32 `protobuf:"varint,6,opt,name=seconds,proto3" json:"seconds,omitempty"`
	// Required. Fractions of seconds in nanoseconds. Must be from 0 to
	// 999,999,999.
	Nanos int32 `protobuf:"varint,7,opt,name=nanos,proto3" json:"nanos,omitempty"`
	// Optional. Specifies either the UTC offset or the time zone of the DateTime.
	// Choose carefully between them, considering that time zone data may change
	// in the future (for example, a country modifies their DST start/end dates,
	// and future DateTimes in the affected range had already been stored).
	// If omitted, the DateTime is considered to be in local time.
	//
	// Types that are assignable to TimeOffset:
	//	*DateTime_UtcOffset
	//	*DateTime_TimeZone
	TimeOffset isDateTime_TimeOffset `protobuf_oneof:"time_offset"`
}

func (x *DateTime) Reset() {
	*x = DateTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_type_datetime_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DateTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateTime) ProtoMessage() {}

func (x *DateTime) ProtoReflect() protoreflect.Message {
	mi := &file_google_type_datetime_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateTime.ProtoReflect.Descriptor instead.
func (*DateTime) Descriptor() ([]byte, []int) {
	return file_google_type_datetime_proto_rawDescGZIP(), []int{0}
}

func (x *DateTime) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *DateTime) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *DateTime) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *DateTime) GetHours() int32 {
	if x != nil {
		return x.Hours
	}
	return 0
}

func (x *DateTime) GetMinutes() int32 {
	if x != nil {
		return x.Minutes
	}
	return 0
}

func (x *DateTime) GetSeconds() int32 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *DateTime) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

func (m *DateTime) GetTimeOffset() isDateTime_TimeOffset {
	if m != nil {
		return m.TimeOffset
	}
	return nil
}

func (x *DateTime) GetUtcOffset() *durationpb.Duration {
	if x, ok := x.GetTimeOffset().(*DateTime_UtcOffset); ok {
		return x.UtcOffset
	}
	return nil
}

func (x *DateTime) GetTimeZone() *TimeZone {
	if x, ok := x.GetTimeOffset().(*DateTime_TimeZone); ok {
		return x.TimeZone
	}
	return nil
}

type isDateTime_TimeOffset interface {
	isDateTime_TimeOffset()
}

type DateTime_UtcOffset struct {
	// UTC offset. Must be whole seconds, between -18 hours and +18 hours.
	// For example, a UTC offset of -4:00 would be represented as
	// { seconds: -14400 }.
	UtcOffset *durationpb.Duration `protobuf:"bytes,8,opt,name=utc_offset,json=utcOffset,proto3,oneof"`
}

type DateTime_TimeZone struct {
	// Time zone.
	TimeZone *TimeZone `protobuf:"bytes,9,opt,name=time_zone,json=timeZone,proto3,oneof"`
}

func (*DateTime_UtcOffset) isDateTime_TimeOffset() {}

func (*DateTime_TimeZone) isDateTime_TimeOffset() {}

// Represents a time zone from the
// [IANA Time Zone Database](https://www.iana.org/time-zones).
type TimeZone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// IANA Time Zone Database time zone, e.g. "America/New_York".
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Optional. IANA Time Zone Database version number, e.g. "2019a".
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *TimeZone) Reset() {
	*x = TimeZone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_type_datetime_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeZone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeZone) ProtoMessage() {}

func (x *TimeZone) ProtoReflect() protoreflect.Message {
	mi := &file_google_type_datetime_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeZone.ProtoReflect.Descriptor instead.
func (*TimeZone) Descriptor() ([]byte, []int) {
	return file_google_type_datetime_proto_rawDescGZIP(), []int{1}
}

func (x *TimeZone) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TimeZone) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_google_type_datetime_proto protoreflect.FileDescriptor

var file_google_type_datetime_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61,
	0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x02, 0x0a, 0x08, 0x44, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f,
	0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68,
	0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x64,
	0x61, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x75,
	0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x75, 0x74,
	0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x61, 0x6e,
	0x6f, 0x73, 0x12, 0x3a, 0x0a, 0x0a, 0x75, 0x74, 0x63, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x48, 0x00, 0x52, 0x09, 0x75, 0x74, 0x63, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x34,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x48, 0x00, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65,
	0x5a, 0x6f, 0x6e, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x22, 0x34, 0x0a, 0x08, 0x54, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x69, 0x0a, 0x0f, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x42, 0x0d, 0x44, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69,
	0x6d, 0x65, 0x3b, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0xf8, 0x01, 0x01, 0xa2, 0x02,
	0x03, 0x47, 0x54, 0x50, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_type_datetime_proto_rawDescOnce sync.Once
	file_google_type_datetime_proto_rawDescData = file_google_type_datetime_proto_rawDesc
)

func file_google_type_datetime_proto_rawDescGZIP() []byte {
	file_google_type_datetime_proto_rawDescOnce.Do(func() {
		file_google_type_datetime_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_type_datetime_proto_rawDescData)
	})
	return file_google_type_datetime_proto_rawDescData
}

var file_google_type_datetime_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_type_datetime_proto_goTypes = []interface{}{
	(*DateTime)(nil),            // 0: google.type.DateTime
	(*TimeZone)(nil),            // 1: google.type.TimeZone
	(*durationpb.Duration)(nil), // 2: google.protobuf.Duration
}
var file_google_type_datetime_proto_depIdxs = []int32{
	2, // 0: google.type.DateTime.utc_offset:type_name -> google.protobuf.Duration
	1, // 1: google.type.DateTime.time_zone:type_name -> google.type.TimeZone
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_type_datetime_proto_init() }
func file_google_type_datetime_proto_init() {
	if File_google_type_datetime_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_type_datetime_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DateTime); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_type_datetime_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeZone); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_google_type_datetime_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*DateTime_UtcOffset)(nil),
		(*DateTime_TimeZone)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_type_datetime_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_type_datetime_proto_goTypes,
		DependencyIndexes: file_google_type_datetime_proto_depIdxs,
		MessageInfos:      file_google_type_datetime_proto_msgTypes,
	}.Build()
	File_google_type_datetime_proto = out.File
	file_google_type_datetime_proto_rawDesc = nil
	file_google_type_datetime_proto_goTypes = nil
	file_google_type_datetime_proto_depIdxs = nil
}
