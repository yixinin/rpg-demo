// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: server-center.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Protocol {

  /// <summary>Holder for reflection information generated from server-center.proto</summary>
  public static partial class ServerCenterReflection {

    #region Descriptor
    /// <summary>File descriptor for server-center.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static ServerCenterReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "ChNzZXJ2ZXItY2VudGVyLnByb3RvEghwcm90b2NvbBoUY2xpZW50LWFjY291",
            "bnQucHJvdG8aEWNsaWVudC1jaGF0LnByb3RvGhFjbGllbnQtZ2FtZS5wcm90",
            "bxoMaGVhZGVyLnByb3RvGgt1bml0eS5wcm90byJLCgxFbnRlckdhbWVSZXAS",
            "CwoDdWlkGAEgASgDEg4KBnJvb21JZBgCIAEoAxIOCgZnYW1lSWQYAyABKAMS",
            "DgoGZ2FtZUlwGAQgASgJIhoKC1JlQ29ubmVjUmVxEgsKA3VpZBgBIAEoAyIc",
            "CgpFbmRHYW1lUmVxEg4KBnJvb21JZBgBIAEoAyILCglTZXR0bGVSZXEiYAoN",
            "U2V0dGxlQ29pblJlcRIOCgZnYW1lSWQYASABKAUSDgoGcm9vbUlkGAIgASgD",
            "Eg8KB3JvdW5kTm8YAyABKAkSDAoEY29pbhgEIAEoAxIQCghnYW1lVHlwZRgF",
            "IAEoBTL9AwoSQ2VudGVyNEdhdGVTZXJ2aWNlEjEKBUxvZ2luEhIucHJvdG9j",
            "b2wuTG9naW5SZXEaEi5wcm90b2NvbC5Mb2dpbkFjayIAEjQKBkxvZ291dBIT",
            "LnByb3RvY29sLkxvZ291dFJlcRoTLnByb3RvY29sLkxvZ291dEFjayIAEjcK",
            "CVJlQ29ubmVjdBIVLnByb3RvY29sLlJlQ29ubmVjUmVxGhEucHJvdG9jb2wu",
            "Q2FsbEFjayIAEjYKCEtpY2tVc2VyEhUucHJvdG9jb2wuS2lja1VzZXJSZXEa",
            "ES5wcm90b2NvbC5DYWxsQWNrIgASWwoTR2V0R2FtZVJvb21UeXBlTGlzdBIg",
            "LnByb3RvY29sLkdldEdhbWVSb29tVHlwZUxpc3RSZXEaIC5wcm90b2NvbC5H",
            "ZXRHYW1lUm9vbVR5cGVMaXN0QWNrIgASOgoKU2VuZFRvVXNlchIXLnByb3Rv",
            "Y29sLlNlbmRUb1VzZXJSZXEaES5wcm90b2NvbC5DYWxsQWNrIgASOgoKU2Vu",
            "ZFRvUm9vbRIXLnByb3RvY29sLlNlbmRUb1Jvb21SZXEaES5wcm90b2NvbC5D",
            "YWxsQWNrIgASOAoJQm9hcmRjYXN0EhYucHJvdG9jb2wuQm9hcmRjYXN0UmVx",
            "GhEucHJvdG9jb2wuQ2FsbEFjayIAMvIBChJDZW50ZXI0R2FtZVNlcnZpY2US",
            "NgoIS2lja1VzZXISFS5wcm90b2NvbC5LaWNrVXNlclJlcRoRLnByb3RvY29s",
            "LkNhbGxBY2siABI0CgdFbmRHYW1lEhQucHJvdG9jb2wuRW5kR2FtZVJlcRoR",
            "LnByb3RvY29sLkNhbGxBY2siABIyCgZTZXR0bGUSEy5wcm90b2NvbC5TZXR0",
            "bGVSZXEaES5wcm90b2NvbC5DYWxsQWNrIgASOgoKU2V0dGxlQ29pbhIXLnBy",
            "b3RvY29sLlNldHRsZUNvaW5SZXEaES5wcm90b2NvbC5DYWxsQWNrIgBiBnBy",
            "b3RvMw=="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Protocol.ClientAccountReflection.Descriptor, global::Protocol.ClientChatReflection.Descriptor, global::Protocol.ClientGameReflection.Descriptor, global::Protocol.HeaderReflection.Descriptor, global::Protocol.UnityReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Protocol.EnterGameRep), global::Protocol.EnterGameRep.Parser, new[]{ "Uid", "RoomId", "GameId", "GameIp" }, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Protocol.ReConnecReq), global::Protocol.ReConnecReq.Parser, new[]{ "Uid" }, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Protocol.EndGameReq), global::Protocol.EndGameReq.Parser, new[]{ "RoomId" }, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Protocol.SettleReq), global::Protocol.SettleReq.Parser, null, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Protocol.SettleCoinReq), global::Protocol.SettleCoinReq.Parser, new[]{ "GameId", "RoomId", "RoundNo", "Coin", "GameType" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class EnterGameRep : pb::IMessage<EnterGameRep> {
    private static readonly pb::MessageParser<EnterGameRep> _parser = new pb::MessageParser<EnterGameRep>(() => new EnterGameRep());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<EnterGameRep> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Protocol.ServerCenterReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public EnterGameRep() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public EnterGameRep(EnterGameRep other) : this() {
      uid_ = other.uid_;
      roomId_ = other.roomId_;
      gameId_ = other.gameId_;
      gameIp_ = other.gameIp_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public EnterGameRep Clone() {
      return new EnterGameRep(this);
    }

    /// <summary>Field number for the "uid" field.</summary>
    public const int UidFieldNumber = 1;
    private long uid_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public long Uid {
      get { return uid_; }
      set {
        uid_ = value;
      }
    }

    /// <summary>Field number for the "roomId" field.</summary>
    public const int RoomIdFieldNumber = 2;
    private long roomId_;
    /// <summary>
    ///房间号
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public long RoomId {
      get { return roomId_; }
      set {
        roomId_ = value;
      }
    }

    /// <summary>Field number for the "gameId" field.</summary>
    public const int GameIdFieldNumber = 3;
    private long gameId_;
    /// <summary>
    ///游戏id
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public long GameId {
      get { return gameId_; }
      set {
        gameId_ = value;
      }
    }

    /// <summary>Field number for the "gameIp" field.</summary>
    public const int GameIpFieldNumber = 4;
    private string gameIp_ = "";
    /// <summary>
    ///game服务器
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string GameIp {
      get { return gameIp_; }
      set {
        gameIp_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as EnterGameRep);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(EnterGameRep other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Uid != other.Uid) return false;
      if (RoomId != other.RoomId) return false;
      if (GameId != other.GameId) return false;
      if (GameIp != other.GameIp) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Uid != 0L) hash ^= Uid.GetHashCode();
      if (RoomId != 0L) hash ^= RoomId.GetHashCode();
      if (GameId != 0L) hash ^= GameId.GetHashCode();
      if (GameIp.Length != 0) hash ^= GameIp.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (Uid != 0L) {
        output.WriteRawTag(8);
        output.WriteInt64(Uid);
      }
      if (RoomId != 0L) {
        output.WriteRawTag(16);
        output.WriteInt64(RoomId);
      }
      if (GameId != 0L) {
        output.WriteRawTag(24);
        output.WriteInt64(GameId);
      }
      if (GameIp.Length != 0) {
        output.WriteRawTag(34);
        output.WriteString(GameIp);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (Uid != 0L) {
        size += 1 + pb::CodedOutputStream.ComputeInt64Size(Uid);
      }
      if (RoomId != 0L) {
        size += 1 + pb::CodedOutputStream.ComputeInt64Size(RoomId);
      }
      if (GameId != 0L) {
        size += 1 + pb::CodedOutputStream.ComputeInt64Size(GameId);
      }
      if (GameIp.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(GameIp);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(EnterGameRep other) {
      if (other == null) {
        return;
      }
      if (other.Uid != 0L) {
        Uid = other.Uid;
      }
      if (other.RoomId != 0L) {
        RoomId = other.RoomId;
      }
      if (other.GameId != 0L) {
        GameId = other.GameId;
      }
      if (other.GameIp.Length != 0) {
        GameIp = other.GameIp;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 8: {
            Uid = input.ReadInt64();
            break;
          }
          case 16: {
            RoomId = input.ReadInt64();
            break;
          }
          case 24: {
            GameId = input.ReadInt64();
            break;
          }
          case 34: {
            GameIp = input.ReadString();
            break;
          }
        }
      }
    }

  }

  public sealed partial class ReConnecReq : pb::IMessage<ReConnecReq> {
    private static readonly pb::MessageParser<ReConnecReq> _parser = new pb::MessageParser<ReConnecReq>(() => new ReConnecReq());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<ReConnecReq> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Protocol.ServerCenterReflection.Descriptor.MessageTypes[1]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ReConnecReq() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ReConnecReq(ReConnecReq other) : this() {
      uid_ = other.uid_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ReConnecReq Clone() {
      return new ReConnecReq(this);
    }

    /// <summary>Field number for the "uid" field.</summary>
    public const int UidFieldNumber = 1;
    private long uid_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public long Uid {
      get { return uid_; }
      set {
        uid_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as ReConnecReq);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(ReConnecReq other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Uid != other.Uid) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Uid != 0L) hash ^= Uid.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (Uid != 0L) {
        output.WriteRawTag(8);
        output.WriteInt64(Uid);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (Uid != 0L) {
        size += 1 + pb::CodedOutputStream.ComputeInt64Size(Uid);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(ReConnecReq other) {
      if (other == null) {
        return;
      }
      if (other.Uid != 0L) {
        Uid = other.Uid;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 8: {
            Uid = input.ReadInt64();
            break;
          }
        }
      }
    }

  }

  public sealed partial class EndGameReq : pb::IMessage<EndGameReq> {
    private static readonly pb::MessageParser<EndGameReq> _parser = new pb::MessageParser<EndGameReq>(() => new EndGameReq());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<EndGameReq> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Protocol.ServerCenterReflection.Descriptor.MessageTypes[2]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public EndGameReq() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public EndGameReq(EndGameReq other) : this() {
      roomId_ = other.roomId_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public EndGameReq Clone() {
      return new EndGameReq(this);
    }

    /// <summary>Field number for the "roomId" field.</summary>
    public const int RoomIdFieldNumber = 1;
    private long roomId_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public long RoomId {
      get { return roomId_; }
      set {
        roomId_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as EndGameReq);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(EndGameReq other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (RoomId != other.RoomId) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (RoomId != 0L) hash ^= RoomId.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (RoomId != 0L) {
        output.WriteRawTag(8);
        output.WriteInt64(RoomId);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (RoomId != 0L) {
        size += 1 + pb::CodedOutputStream.ComputeInt64Size(RoomId);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(EndGameReq other) {
      if (other == null) {
        return;
      }
      if (other.RoomId != 0L) {
        RoomId = other.RoomId;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 8: {
            RoomId = input.ReadInt64();
            break;
          }
        }
      }
    }

  }

  public sealed partial class SettleReq : pb::IMessage<SettleReq> {
    private static readonly pb::MessageParser<SettleReq> _parser = new pb::MessageParser<SettleReq>(() => new SettleReq());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<SettleReq> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Protocol.ServerCenterReflection.Descriptor.MessageTypes[3]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public SettleReq() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public SettleReq(SettleReq other) : this() {
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public SettleReq Clone() {
      return new SettleReq(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as SettleReq);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(SettleReq other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(SettleReq other) {
      if (other == null) {
        return;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
        }
      }
    }

  }

  public sealed partial class SettleCoinReq : pb::IMessage<SettleCoinReq> {
    private static readonly pb::MessageParser<SettleCoinReq> _parser = new pb::MessageParser<SettleCoinReq>(() => new SettleCoinReq());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<SettleCoinReq> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Protocol.ServerCenterReflection.Descriptor.MessageTypes[4]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public SettleCoinReq() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public SettleCoinReq(SettleCoinReq other) : this() {
      gameId_ = other.gameId_;
      roomId_ = other.roomId_;
      roundNo_ = other.roundNo_;
      coin_ = other.coin_;
      gameType_ = other.gameType_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public SettleCoinReq Clone() {
      return new SettleCoinReq(this);
    }

    /// <summary>Field number for the "gameId" field.</summary>
    public const int GameIdFieldNumber = 1;
    private int gameId_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int GameId {
      get { return gameId_; }
      set {
        gameId_ = value;
      }
    }

    /// <summary>Field number for the "roomId" field.</summary>
    public const int RoomIdFieldNumber = 2;
    private long roomId_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public long RoomId {
      get { return roomId_; }
      set {
        roomId_ = value;
      }
    }

    /// <summary>Field number for the "roundNo" field.</summary>
    public const int RoundNoFieldNumber = 3;
    private string roundNo_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string RoundNo {
      get { return roundNo_; }
      set {
        roundNo_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "coin" field.</summary>
    public const int CoinFieldNumber = 4;
    private long coin_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public long Coin {
      get { return coin_; }
      set {
        coin_ = value;
      }
    }

    /// <summary>Field number for the "gameType" field.</summary>
    public const int GameTypeFieldNumber = 5;
    private int gameType_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int GameType {
      get { return gameType_; }
      set {
        gameType_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as SettleCoinReq);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(SettleCoinReq other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (GameId != other.GameId) return false;
      if (RoomId != other.RoomId) return false;
      if (RoundNo != other.RoundNo) return false;
      if (Coin != other.Coin) return false;
      if (GameType != other.GameType) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (GameId != 0) hash ^= GameId.GetHashCode();
      if (RoomId != 0L) hash ^= RoomId.GetHashCode();
      if (RoundNo.Length != 0) hash ^= RoundNo.GetHashCode();
      if (Coin != 0L) hash ^= Coin.GetHashCode();
      if (GameType != 0) hash ^= GameType.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (GameId != 0) {
        output.WriteRawTag(8);
        output.WriteInt32(GameId);
      }
      if (RoomId != 0L) {
        output.WriteRawTag(16);
        output.WriteInt64(RoomId);
      }
      if (RoundNo.Length != 0) {
        output.WriteRawTag(26);
        output.WriteString(RoundNo);
      }
      if (Coin != 0L) {
        output.WriteRawTag(32);
        output.WriteInt64(Coin);
      }
      if (GameType != 0) {
        output.WriteRawTag(40);
        output.WriteInt32(GameType);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (GameId != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(GameId);
      }
      if (RoomId != 0L) {
        size += 1 + pb::CodedOutputStream.ComputeInt64Size(RoomId);
      }
      if (RoundNo.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(RoundNo);
      }
      if (Coin != 0L) {
        size += 1 + pb::CodedOutputStream.ComputeInt64Size(Coin);
      }
      if (GameType != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(GameType);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(SettleCoinReq other) {
      if (other == null) {
        return;
      }
      if (other.GameId != 0) {
        GameId = other.GameId;
      }
      if (other.RoomId != 0L) {
        RoomId = other.RoomId;
      }
      if (other.RoundNo.Length != 0) {
        RoundNo = other.RoundNo;
      }
      if (other.Coin != 0L) {
        Coin = other.Coin;
      }
      if (other.GameType != 0) {
        GameType = other.GameType;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 8: {
            GameId = input.ReadInt32();
            break;
          }
          case 16: {
            RoomId = input.ReadInt64();
            break;
          }
          case 26: {
            RoundNo = input.ReadString();
            break;
          }
          case 32: {
            Coin = input.ReadInt64();
            break;
          }
          case 40: {
            GameType = input.ReadInt32();
            break;
          }
        }
      }
    }

  }

  #endregion

}

#endregion Designer generated code
