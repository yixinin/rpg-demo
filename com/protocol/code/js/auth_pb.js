// source: auth.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var header_pb = require('./header_pb.js');
goog.object.extend(proto, header_pb);
goog.exportSymbol('proto.protocol.LoginAck', null, global);
goog.exportSymbol('proto.protocol.LoginReq', null, global);
goog.exportSymbol('proto.protocol.LogoutAck', null, global);
goog.exportSymbol('proto.protocol.LogoutReq', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.protocol.LoginReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.protocol.LoginReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protocol.LoginReq.displayName = 'proto.protocol.LoginReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.protocol.LoginAck = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.protocol.LoginAck, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protocol.LoginAck.displayName = 'proto.protocol.LoginAck';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.protocol.LogoutReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.protocol.LogoutReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protocol.LogoutReq.displayName = 'proto.protocol.LogoutReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.protocol.LogoutAck = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.protocol.LogoutAck, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protocol.LogoutAck.displayName = 'proto.protocol.LogoutAck';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.protocol.LoginReq.prototype.toObject = function(opt_includeInstance) {
  return proto.protocol.LoginReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protocol.LoginReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protocol.LoginReq.toObject = function(includeInstance, msg) {
  var f, obj = {
    header: (f = msg.getHeader()) && header_pb.ReqHeader.toObject(includeInstance, f),
    username: jspb.Message.getFieldWithDefault(msg, 2, ""),
    password: jspb.Message.getFieldWithDefault(msg, 3, ""),
    validcode: jspb.Message.getFieldWithDefault(msg, 4, ""),
    validid: jspb.Message.getFieldWithDefault(msg, 5, ""),
    phonenumber: jspb.Message.getFieldWithDefault(msg, 6, ""),
    devicecode: jspb.Message.getFieldWithDefault(msg, 7, ""),
    invitecode: jspb.Message.getFieldWithDefault(msg, 8, ""),
    logintype: jspb.Message.getFieldWithDefault(msg, 9, 0),
    register: jspb.Message.getBooleanFieldWithDefault(msg, 10, false)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.protocol.LoginReq}
 */
proto.protocol.LoginReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protocol.LoginReq;
  return proto.protocol.LoginReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protocol.LoginReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protocol.LoginReq}
 */
proto.protocol.LoginReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new header_pb.ReqHeader;
      reader.readMessage(value,header_pb.ReqHeader.deserializeBinaryFromReader);
      msg.setHeader(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setUsername(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setPassword(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setValidcode(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setValidid(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setPhonenumber(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setDevicecode(value);
      break;
    case 8:
      var value = /** @type {string} */ (reader.readString());
      msg.setInvitecode(value);
      break;
    case 9:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setLogintype(value);
      break;
    case 10:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setRegister(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.protocol.LoginReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protocol.LoginReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protocol.LoginReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protocol.LoginReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getHeader();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      header_pb.ReqHeader.serializeBinaryToWriter
    );
  }
  f = message.getUsername();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getPassword();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getValidcode();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getValidid();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getPhonenumber();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getDevicecode();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getInvitecode();
  if (f.length > 0) {
    writer.writeString(
      8,
      f
    );
  }
  f = message.getLogintype();
  if (f !== 0) {
    writer.writeInt32(
      9,
      f
    );
  }
  f = message.getRegister();
  if (f) {
    writer.writeBool(
      10,
      f
    );
  }
};


/**
 * optional ReqHeader header = 1;
 * @return {?proto.protocol.ReqHeader}
 */
proto.protocol.LoginReq.prototype.getHeader = function() {
  return /** @type{?proto.protocol.ReqHeader} */ (
    jspb.Message.getWrapperField(this, header_pb.ReqHeader, 1));
};


/**
 * @param {?proto.protocol.ReqHeader|undefined} value
 * @return {!proto.protocol.LoginReq} returns this
*/
proto.protocol.LoginReq.prototype.setHeader = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.protocol.LoginReq} returns this
 */
proto.protocol.LoginReq.prototype.clearHeader = function() {
  return this.setHeader(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.protocol.LoginReq.prototype.hasHeader = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional string userName = 2;
 * @return {string}
 */
proto.protocol.LoginReq.prototype.getUsername = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.protocol.LoginReq} returns this
 */
proto.protocol.LoginReq.prototype.setUsername = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string password = 3;
 * @return {string}
 */
proto.protocol.LoginReq.prototype.getPassword = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.protocol.LoginReq} returns this
 */
proto.protocol.LoginReq.prototype.setPassword = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string validCode = 4;
 * @return {string}
 */
proto.protocol.LoginReq.prototype.getValidcode = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.protocol.LoginReq} returns this
 */
proto.protocol.LoginReq.prototype.setValidcode = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string validId = 5;
 * @return {string}
 */
proto.protocol.LoginReq.prototype.getValidid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.protocol.LoginReq} returns this
 */
proto.protocol.LoginReq.prototype.setValidid = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string phoneNumber = 6;
 * @return {string}
 */
proto.protocol.LoginReq.prototype.getPhonenumber = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * @param {string} value
 * @return {!proto.protocol.LoginReq} returns this
 */
proto.protocol.LoginReq.prototype.setPhonenumber = function(value) {
  return jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * optional string DeviceCode = 7;
 * @return {string}
 */
proto.protocol.LoginReq.prototype.getDevicecode = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/**
 * @param {string} value
 * @return {!proto.protocol.LoginReq} returns this
 */
proto.protocol.LoginReq.prototype.setDevicecode = function(value) {
  return jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * optional string InviteCode = 8;
 * @return {string}
 */
proto.protocol.LoginReq.prototype.getInvitecode = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/**
 * @param {string} value
 * @return {!proto.protocol.LoginReq} returns this
 */
proto.protocol.LoginReq.prototype.setInvitecode = function(value) {
  return jspb.Message.setProto3StringField(this, 8, value);
};


/**
 * optional int32 loginType = 9;
 * @return {number}
 */
proto.protocol.LoginReq.prototype.getLogintype = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 9, 0));
};


/**
 * @param {number} value
 * @return {!proto.protocol.LoginReq} returns this
 */
proto.protocol.LoginReq.prototype.setLogintype = function(value) {
  return jspb.Message.setProto3IntField(this, 9, value);
};


/**
 * optional bool register = 10;
 * @return {boolean}
 */
proto.protocol.LoginReq.prototype.getRegister = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 10, false));
};


/**
 * @param {boolean} value
 * @return {!proto.protocol.LoginReq} returns this
 */
proto.protocol.LoginReq.prototype.setRegister = function(value) {
  return jspb.Message.setProto3BooleanField(this, 10, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.protocol.LoginAck.prototype.toObject = function(opt_includeInstance) {
  return proto.protocol.LoginAck.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protocol.LoginAck} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protocol.LoginAck.toObject = function(includeInstance, msg) {
  var f, obj = {
    header: (f = msg.getHeader()) && header_pb.AckHeader.toObject(includeInstance, f),
    token: jspb.Message.getFieldWithDefault(msg, 2, ""),
    devicecode: jspb.Message.getFieldWithDefault(msg, 3, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.protocol.LoginAck}
 */
proto.protocol.LoginAck.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protocol.LoginAck;
  return proto.protocol.LoginAck.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protocol.LoginAck} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protocol.LoginAck}
 */
proto.protocol.LoginAck.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new header_pb.AckHeader;
      reader.readMessage(value,header_pb.AckHeader.deserializeBinaryFromReader);
      msg.setHeader(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setToken(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setDevicecode(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.protocol.LoginAck.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protocol.LoginAck.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protocol.LoginAck} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protocol.LoginAck.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getHeader();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      header_pb.AckHeader.serializeBinaryToWriter
    );
  }
  f = message.getToken();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getDevicecode();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
};


/**
 * optional AckHeader header = 1;
 * @return {?proto.protocol.AckHeader}
 */
proto.protocol.LoginAck.prototype.getHeader = function() {
  return /** @type{?proto.protocol.AckHeader} */ (
    jspb.Message.getWrapperField(this, header_pb.AckHeader, 1));
};


/**
 * @param {?proto.protocol.AckHeader|undefined} value
 * @return {!proto.protocol.LoginAck} returns this
*/
proto.protocol.LoginAck.prototype.setHeader = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.protocol.LoginAck} returns this
 */
proto.protocol.LoginAck.prototype.clearHeader = function() {
  return this.setHeader(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.protocol.LoginAck.prototype.hasHeader = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional string token = 2;
 * @return {string}
 */
proto.protocol.LoginAck.prototype.getToken = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.protocol.LoginAck} returns this
 */
proto.protocol.LoginAck.prototype.setToken = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string DeviceCode = 3;
 * @return {string}
 */
proto.protocol.LoginAck.prototype.getDevicecode = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.protocol.LoginAck} returns this
 */
proto.protocol.LoginAck.prototype.setDevicecode = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.protocol.LogoutReq.prototype.toObject = function(opt_includeInstance) {
  return proto.protocol.LogoutReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protocol.LogoutReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protocol.LogoutReq.toObject = function(includeInstance, msg) {
  var f, obj = {
    header: (f = msg.getHeader()) && header_pb.ReqHeader.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.protocol.LogoutReq}
 */
proto.protocol.LogoutReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protocol.LogoutReq;
  return proto.protocol.LogoutReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protocol.LogoutReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protocol.LogoutReq}
 */
proto.protocol.LogoutReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new header_pb.ReqHeader;
      reader.readMessage(value,header_pb.ReqHeader.deserializeBinaryFromReader);
      msg.setHeader(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.protocol.LogoutReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protocol.LogoutReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protocol.LogoutReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protocol.LogoutReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getHeader();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      header_pb.ReqHeader.serializeBinaryToWriter
    );
  }
};


/**
 * optional ReqHeader header = 1;
 * @return {?proto.protocol.ReqHeader}
 */
proto.protocol.LogoutReq.prototype.getHeader = function() {
  return /** @type{?proto.protocol.ReqHeader} */ (
    jspb.Message.getWrapperField(this, header_pb.ReqHeader, 1));
};


/**
 * @param {?proto.protocol.ReqHeader|undefined} value
 * @return {!proto.protocol.LogoutReq} returns this
*/
proto.protocol.LogoutReq.prototype.setHeader = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.protocol.LogoutReq} returns this
 */
proto.protocol.LogoutReq.prototype.clearHeader = function() {
  return this.setHeader(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.protocol.LogoutReq.prototype.hasHeader = function() {
  return jspb.Message.getField(this, 1) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.protocol.LogoutAck.prototype.toObject = function(opt_includeInstance) {
  return proto.protocol.LogoutAck.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protocol.LogoutAck} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protocol.LogoutAck.toObject = function(includeInstance, msg) {
  var f, obj = {
    header: (f = msg.getHeader()) && header_pb.AckHeader.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.protocol.LogoutAck}
 */
proto.protocol.LogoutAck.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protocol.LogoutAck;
  return proto.protocol.LogoutAck.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protocol.LogoutAck} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protocol.LogoutAck}
 */
proto.protocol.LogoutAck.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new header_pb.AckHeader;
      reader.readMessage(value,header_pb.AckHeader.deserializeBinaryFromReader);
      msg.setHeader(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.protocol.LogoutAck.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protocol.LogoutAck.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protocol.LogoutAck} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protocol.LogoutAck.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getHeader();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      header_pb.AckHeader.serializeBinaryToWriter
    );
  }
};


/**
 * optional AckHeader header = 1;
 * @return {?proto.protocol.AckHeader}
 */
proto.protocol.LogoutAck.prototype.getHeader = function() {
  return /** @type{?proto.protocol.AckHeader} */ (
    jspb.Message.getWrapperField(this, header_pb.AckHeader, 1));
};


/**
 * @param {?proto.protocol.AckHeader|undefined} value
 * @return {!proto.protocol.LogoutAck} returns this
*/
proto.protocol.LogoutAck.prototype.setHeader = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.protocol.LogoutAck} returns this
 */
proto.protocol.LogoutAck.prototype.clearHeader = function() {
  return this.setHeader(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.protocol.LogoutAck.prototype.hasHeader = function() {
  return jspb.Message.getField(this, 1) != null;
};


goog.object.extend(exports, proto.protocol);
