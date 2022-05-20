// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: him.proto
//
// For information on using the generated types, please see the documentation:
//   https://github.com/apple/swift-protobuf/

import Foundation
import SwiftProtobuf

// If the compiler emits an error on this type, it is because this file
// was generated by a version of the `protoc` Swift plug-in that is
// incompatible with the version of SwiftProtobuf to which you are linking.
// Please ensure that you are building against the same version of the API
// that was used to generate this file.
fileprivate struct _GeneratedWithProtocGenSwiftVersion: SwiftProtobuf.ProtobufAPIVersionCheck {
  struct _2: SwiftProtobuf.ProtobufAPIVersion_2 {}
  typealias Version = _2
}

///消息状态
enum Pb_MessageStatus: SwiftProtobuf.Enum {
  typealias RawValue = Int
  case init_ // = 0
  case sending // = 1
  case sendSucc // = 2
  case sendFail // = 3
  case hasDeleted // = 4
  case localImported // = 5
  case localRevoked // = 6
  case UNRECOGNIZED(Int)

  init() {
    self = .init_
  }

  init?(rawValue: Int) {
    switch rawValue {
    case 0: self = .init_
    case 1: self = .sending
    case 2: self = .sendSucc
    case 3: self = .sendFail
    case 4: self = .hasDeleted
    case 5: self = .localImported
    case 6: self = .localRevoked
    default: self = .UNRECOGNIZED(rawValue)
    }
  }

  var rawValue: Int {
    switch self {
    case .init_: return 0
    case .sending: return 1
    case .sendSucc: return 2
    case .sendFail: return 3
    case .hasDeleted: return 4
    case .localImported: return 5
    case .localRevoked: return 6
    case .UNRECOGNIZED(let i): return i
    }
  }

}

#if swift(>=4.2)

extension Pb_MessageStatus: CaseIterable {
  // The compiler won't synthesize support with the UNRECOGNIZED case.
  static var allCases: [Pb_MessageStatus] = [
    .init_,
    .sending,
    .sendSucc,
    .sendFail,
    .hasDeleted,
    .localImported,
    .localRevoked,
  ]
}

#endif  // swift(>=4.2)

enum Pb_PackType: SwiftProtobuf.Enum {
  typealias RawValue = Int
  case loginReq // = 0
  case loginAck // = 1

  ///消息同步当做心跳，可以避免在线情况下因为网络等原因消息未送达，不用等用户下次登录就可以重新推送消息
  case msgPullReq // = 2
  case msgPullAck // = 3
  case msgHistoryReq // = 4
  case msgHistoryAck // = 5
  case msgReq // = 6
  case msgAck // = 7
  case UNRECOGNIZED(Int)

  init() {
    self = .loginReq
  }

  init?(rawValue: Int) {
    switch rawValue {
    case 0: self = .loginReq
    case 1: self = .loginAck
    case 2: self = .msgPullReq
    case 3: self = .msgPullAck
    case 4: self = .msgHistoryReq
    case 5: self = .msgHistoryAck
    case 6: self = .msgReq
    case 7: self = .msgAck
    default: self = .UNRECOGNIZED(rawValue)
    }
  }

  var rawValue: Int {
    switch self {
    case .loginReq: return 0
    case .loginAck: return 1
    case .msgPullReq: return 2
    case .msgPullAck: return 3
    case .msgHistoryReq: return 4
    case .msgHistoryAck: return 5
    case .msgReq: return 6
    case .msgAck: return 7
    case .UNRECOGNIZED(let i): return i
    }
  }

}

#if swift(>=4.2)

extension Pb_PackType: CaseIterable {
  // The compiler won't synthesize support with the UNRECOGNIZED case.
  static var allCases: [Pb_PackType] = [
    .loginReq,
    .loginAck,
    .msgPullReq,
    .msgPullAck,
    .msgHistoryReq,
    .msgHistoryAck,
    .msgReq,
    .msgAck,
  ]
}

#endif  // swift(>=4.2)

///会话类型
enum Pb_ConversationType: SwiftProtobuf.Enum {
  typealias RawValue = Int
  case c2C // = 0
  case group // = 1
  case UNRECOGNIZED(Int)

  init() {
    self = .c2C
  }

  init?(rawValue: Int) {
    switch rawValue {
    case 0: self = .c2C
    case 1: self = .group
    default: self = .UNRECOGNIZED(rawValue)
    }
  }

  var rawValue: Int {
    switch self {
    case .c2C: return 0
    case .group: return 1
    case .UNRECOGNIZED(let i): return i
    }
  }

}

#if swift(>=4.2)

extension Pb_ConversationType: CaseIterable {
  // The compiler won't synthesize support with the UNRECOGNIZED case.
  static var allCases: [Pb_ConversationType] = [
    .c2C,
    .group,
  ]
}

#endif  // swift(>=4.2)

///消息类型
enum Pb_ElemType: SwiftProtobuf.Enum {
  typealias RawValue = Int
  case custom // = 0
  case text // = 1
  case image // = 2
  case UNRECOGNIZED(Int)

  init() {
    self = .custom
  }

  init?(rawValue: Int) {
    switch rawValue {
    case 0: self = .custom
    case 1: self = .text
    case 2: self = .image
    default: self = .UNRECOGNIZED(rawValue)
    }
  }

  var rawValue: Int {
    switch self {
    case .custom: return 0
    case .text: return 1
    case .image: return 2
    case .UNRECOGNIZED(let i): return i
    }
  }

}

#if swift(>=4.2)

extension Pb_ElemType: CaseIterable {
  // The compiler won't synthesize support with the UNRECOGNIZED case.
  static var allCases: [Pb_ElemType] = [
    .custom,
    .text,
    .image,
  ]
}

#endif  // swift(>=4.2)

///消息最上层
struct Pb_Pack {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  var type: Pb_PackType = .loginReq

  var body: Data = Data()

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}
}

///设备登录
struct Pb_LoginReq {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  var token: String = String()

  var deviceID: String = String()

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}
}

///登录回执
struct Pb_LoginAck {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  var code: Int32 = 0

  var msg: String = String()

  var userID: String = String()

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}
}

struct Pb_Message {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  ///会话类型
  var conversationType: Pb_ConversationType = .c2C

  ///消息类型
  var type: Pb_ElemType = .custom

  ///会话id
  var conversationID: String = String()

  ///app端消息id
  var msgID: String = String()

  ///全局唯一id
  var msgUid: String = String()

  ///消息状态
  var status: Pb_MessageStatus = .init_

  ///消息发送者
  var senderID: String = String()

  ///消息接收者
  var targetID: String = String()

  ///消息发送者昵称
  var nickName: String = String()

  ///消息发送者头像
  var faceURL: String = String()

  ///消息内容
  var content: String = String()

  ///消息发送时间
  var timestamp: Int64 = 0

  ///消息自定义数据
  var cloudCustomData: Data = Data()

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}
}

///消息回执
struct Pb_MessageAck {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  var msgID: String = String()

  var msgUid: String = String()

  ///消息成功200，失败，拉黑，不是好友
  var code: Int32 = 0

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}
}

///离线消息拉取
struct Pb_MessagePullReq {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  var userID: String = String()

  var timestamp: Int64 = 0

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}
}

///离线消息拉取响应
struct Pb_MessagePullAck {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  var msglist: [Pb_Message] = []

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}
}

///历史消息响应 
struct Pb_MsgHistoryAck {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  var msglist: [Pb_Message] = []

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}
}

///历史消息拉取
struct Pb_MsgHistoryReq {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  var userID: String = String()

  var conversationID: String = String()

  var count: Int64 = 0

  var timestamp: Int64 = 0

  var unknownFields = SwiftProtobuf.UnknownStorage()

  init() {}
}

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "pb"

extension Pb_MessageStatus: SwiftProtobuf._ProtoNameProviding {
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "init"),
    1: .same(proto: "sending"),
    2: .same(proto: "sendSucc"),
    3: .same(proto: "sendFail"),
    4: .same(proto: "hasDeleted"),
    5: .same(proto: "localImported"),
    6: .same(proto: "localRevoked"),
  ]
}

extension Pb_PackType: SwiftProtobuf._ProtoNameProviding {
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "loginReq"),
    1: .same(proto: "loginAck"),
    2: .same(proto: "msgPullReq"),
    3: .same(proto: "msgPullAck"),
    4: .same(proto: "msgHistoryReq"),
    5: .same(proto: "msgHistoryAck"),
    6: .same(proto: "msgReq"),
    7: .same(proto: "msgAck"),
  ]
}

extension Pb_ConversationType: SwiftProtobuf._ProtoNameProviding {
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "c2c"),
    1: .same(proto: "group"),
  ]
}

extension Pb_ElemType: SwiftProtobuf._ProtoNameProviding {
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "custom"),
    1: .same(proto: "text"),
    2: .same(proto: "image"),
  ]
}

extension Pb_Pack: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".Pack"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "type"),
    2: .same(proto: "body"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularEnumField(value: &self.type) }()
      case 2: try { try decoder.decodeSingularBytesField(value: &self.body) }()
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if self.type != .loginReq {
      try visitor.visitSingularEnumField(value: self.type, fieldNumber: 1)
    }
    if !self.body.isEmpty {
      try visitor.visitSingularBytesField(value: self.body, fieldNumber: 2)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: Pb_Pack, rhs: Pb_Pack) -> Bool {
    if lhs.type != rhs.type {return false}
    if lhs.body != rhs.body {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Pb_LoginReq: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".LoginReq"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "token"),
    2: .same(proto: "deviceId"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.token) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.deviceID) }()
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.token.isEmpty {
      try visitor.visitSingularStringField(value: self.token, fieldNumber: 1)
    }
    if !self.deviceID.isEmpty {
      try visitor.visitSingularStringField(value: self.deviceID, fieldNumber: 2)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: Pb_LoginReq, rhs: Pb_LoginReq) -> Bool {
    if lhs.token != rhs.token {return false}
    if lhs.deviceID != rhs.deviceID {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Pb_LoginAck: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".LoginAck"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "code"),
    2: .same(proto: "msg"),
    3: .same(proto: "userId"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularInt32Field(value: &self.code) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.msg) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self.userID) }()
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if self.code != 0 {
      try visitor.visitSingularInt32Field(value: self.code, fieldNumber: 1)
    }
    if !self.msg.isEmpty {
      try visitor.visitSingularStringField(value: self.msg, fieldNumber: 2)
    }
    if !self.userID.isEmpty {
      try visitor.visitSingularStringField(value: self.userID, fieldNumber: 3)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: Pb_LoginAck, rhs: Pb_LoginAck) -> Bool {
    if lhs.code != rhs.code {return false}
    if lhs.msg != rhs.msg {return false}
    if lhs.userID != rhs.userID {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Pb_Message: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".Message"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    13: .same(proto: "conversationType"),
    1: .same(proto: "type"),
    2: .same(proto: "conversationId"),
    3: .same(proto: "msgID"),
    4: .same(proto: "msgUid"),
    5: .same(proto: "status"),
    6: .same(proto: "senderId"),
    7: .same(proto: "targetId"),
    8: .same(proto: "nickName"),
    9: .same(proto: "faceUrl"),
    10: .same(proto: "content"),
    11: .same(proto: "timestamp"),
    12: .same(proto: "cloudCustomData"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularEnumField(value: &self.type) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.conversationID) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self.msgID) }()
      case 4: try { try decoder.decodeSingularStringField(value: &self.msgUid) }()
      case 5: try { try decoder.decodeSingularEnumField(value: &self.status) }()
      case 6: try { try decoder.decodeSingularStringField(value: &self.senderID) }()
      case 7: try { try decoder.decodeSingularStringField(value: &self.targetID) }()
      case 8: try { try decoder.decodeSingularStringField(value: &self.nickName) }()
      case 9: try { try decoder.decodeSingularStringField(value: &self.faceURL) }()
      case 10: try { try decoder.decodeSingularStringField(value: &self.content) }()
      case 11: try { try decoder.decodeSingularInt64Field(value: &self.timestamp) }()
      case 12: try { try decoder.decodeSingularBytesField(value: &self.cloudCustomData) }()
      case 13: try { try decoder.decodeSingularEnumField(value: &self.conversationType) }()
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if self.type != .custom {
      try visitor.visitSingularEnumField(value: self.type, fieldNumber: 1)
    }
    if !self.conversationID.isEmpty {
      try visitor.visitSingularStringField(value: self.conversationID, fieldNumber: 2)
    }
    if !self.msgID.isEmpty {
      try visitor.visitSingularStringField(value: self.msgID, fieldNumber: 3)
    }
    if !self.msgUid.isEmpty {
      try visitor.visitSingularStringField(value: self.msgUid, fieldNumber: 4)
    }
    if self.status != .init_ {
      try visitor.visitSingularEnumField(value: self.status, fieldNumber: 5)
    }
    if !self.senderID.isEmpty {
      try visitor.visitSingularStringField(value: self.senderID, fieldNumber: 6)
    }
    if !self.targetID.isEmpty {
      try visitor.visitSingularStringField(value: self.targetID, fieldNumber: 7)
    }
    if !self.nickName.isEmpty {
      try visitor.visitSingularStringField(value: self.nickName, fieldNumber: 8)
    }
    if !self.faceURL.isEmpty {
      try visitor.visitSingularStringField(value: self.faceURL, fieldNumber: 9)
    }
    if !self.content.isEmpty {
      try visitor.visitSingularStringField(value: self.content, fieldNumber: 10)
    }
    if self.timestamp != 0 {
      try visitor.visitSingularInt64Field(value: self.timestamp, fieldNumber: 11)
    }
    if !self.cloudCustomData.isEmpty {
      try visitor.visitSingularBytesField(value: self.cloudCustomData, fieldNumber: 12)
    }
    if self.conversationType != .c2C {
      try visitor.visitSingularEnumField(value: self.conversationType, fieldNumber: 13)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: Pb_Message, rhs: Pb_Message) -> Bool {
    if lhs.conversationType != rhs.conversationType {return false}
    if lhs.type != rhs.type {return false}
    if lhs.conversationID != rhs.conversationID {return false}
    if lhs.msgID != rhs.msgID {return false}
    if lhs.msgUid != rhs.msgUid {return false}
    if lhs.status != rhs.status {return false}
    if lhs.senderID != rhs.senderID {return false}
    if lhs.targetID != rhs.targetID {return false}
    if lhs.nickName != rhs.nickName {return false}
    if lhs.faceURL != rhs.faceURL {return false}
    if lhs.content != rhs.content {return false}
    if lhs.timestamp != rhs.timestamp {return false}
    if lhs.cloudCustomData != rhs.cloudCustomData {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Pb_MessageAck: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".MessageAck"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "msgId"),
    2: .same(proto: "msgUid"),
    3: .same(proto: "code"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.msgID) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.msgUid) }()
      case 3: try { try decoder.decodeSingularInt32Field(value: &self.code) }()
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.msgID.isEmpty {
      try visitor.visitSingularStringField(value: self.msgID, fieldNumber: 1)
    }
    if !self.msgUid.isEmpty {
      try visitor.visitSingularStringField(value: self.msgUid, fieldNumber: 2)
    }
    if self.code != 0 {
      try visitor.visitSingularInt32Field(value: self.code, fieldNumber: 3)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: Pb_MessageAck, rhs: Pb_MessageAck) -> Bool {
    if lhs.msgID != rhs.msgID {return false}
    if lhs.msgUid != rhs.msgUid {return false}
    if lhs.code != rhs.code {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Pb_MessagePullReq: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".MessagePullReq"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "userId"),
    2: .same(proto: "timestamp"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.userID) }()
      case 2: try { try decoder.decodeSingularInt64Field(value: &self.timestamp) }()
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.userID.isEmpty {
      try visitor.visitSingularStringField(value: self.userID, fieldNumber: 1)
    }
    if self.timestamp != 0 {
      try visitor.visitSingularInt64Field(value: self.timestamp, fieldNumber: 2)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: Pb_MessagePullReq, rhs: Pb_MessagePullReq) -> Bool {
    if lhs.userID != rhs.userID {return false}
    if lhs.timestamp != rhs.timestamp {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Pb_MessagePullAck: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".MessagePullAck"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "msglist"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeRepeatedMessageField(value: &self.msglist) }()
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.msglist.isEmpty {
      try visitor.visitRepeatedMessageField(value: self.msglist, fieldNumber: 1)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: Pb_MessagePullAck, rhs: Pb_MessagePullAck) -> Bool {
    if lhs.msglist != rhs.msglist {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Pb_MsgHistoryAck: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".MsgHistoryAck"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "msglist"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeRepeatedMessageField(value: &self.msglist) }()
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.msglist.isEmpty {
      try visitor.visitRepeatedMessageField(value: self.msglist, fieldNumber: 1)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: Pb_MsgHistoryAck, rhs: Pb_MsgHistoryAck) -> Bool {
    if lhs.msglist != rhs.msglist {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Pb_MsgHistoryReq: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  static let protoMessageName: String = _protobuf_package + ".MsgHistoryReq"
  static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "userId"),
    2: .same(proto: "conversationId"),
    3: .same(proto: "count"),
    4: .same(proto: "timestamp"),
  ]

  mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.userID) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.conversationID) }()
      case 3: try { try decoder.decodeSingularInt64Field(value: &self.count) }()
      case 4: try { try decoder.decodeSingularInt64Field(value: &self.timestamp) }()
      default: break
      }
    }
  }

  func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.userID.isEmpty {
      try visitor.visitSingularStringField(value: self.userID, fieldNumber: 1)
    }
    if !self.conversationID.isEmpty {
      try visitor.visitSingularStringField(value: self.conversationID, fieldNumber: 2)
    }
    if self.count != 0 {
      try visitor.visitSingularInt64Field(value: self.count, fieldNumber: 3)
    }
    if self.timestamp != 0 {
      try visitor.visitSingularInt64Field(value: self.timestamp, fieldNumber: 4)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  static func ==(lhs: Pb_MsgHistoryReq, rhs: Pb_MsgHistoryReq) -> Bool {
    if lhs.userID != rhs.userID {return false}
    if lhs.conversationID != rhs.conversationID {return false}
    if lhs.count != rhs.count {return false}
    if lhs.timestamp != rhs.timestamp {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}
