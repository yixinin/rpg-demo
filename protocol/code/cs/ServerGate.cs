// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: server-gate.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Protocol {

  /// <summary>Holder for reflection information generated from server-gate.proto</summary>
  public static partial class ServerGateReflection {

    #region Descriptor
    /// <summary>File descriptor for server-gate.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static ServerGateReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "ChFzZXJ2ZXItZ2F0ZS5wcm90bxIIcHJvdG9jb2waC3VuaXR5LnByb3RvGgxo",
            "ZWFkZXIucHJvdG8iLwoQTGlua0dhbWU0R2F0ZVJlcRILCgN1aWQYASADKAMS",
            "DgoGc2VydmVyGAIgASgJMooBCgtHYXRlU2VydmljZRI2CghLaWNrVXNlchIV",
            "LnByb3RvY29sLktpY2tVc2VyUmVxGhEucHJvdG9jb2wuQ2FsbEFjayIAEkMK",
            "DkxpbmtHYW1lU2VydmVyEhoucHJvdG9jb2wuTGlua0dhbWU0R2F0ZVJlcRoR",
            "LnByb3RvY29sLkNhbGxBY2siACgBYgZwcm90bzM="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Protocol.UnityReflection.Descriptor, global::Protocol.HeaderReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Protocol.LinkGame4GateReq), global::Protocol.LinkGame4GateReq.Parser, new[]{ "Uid", "Server" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class LinkGame4GateReq : pb::IMessage<LinkGame4GateReq> {
    private static readonly pb::MessageParser<LinkGame4GateReq> _parser = new pb::MessageParser<LinkGame4GateReq>(() => new LinkGame4GateReq());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<LinkGame4GateReq> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Protocol.ServerGateReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public LinkGame4GateReq() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public LinkGame4GateReq(LinkGame4GateReq other) : this() {
      uid_ = other.uid_.Clone();
      server_ = other.server_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public LinkGame4GateReq Clone() {
      return new LinkGame4GateReq(this);
    }

    /// <summary>Field number for the "uid" field.</summary>
    public const int UidFieldNumber = 1;
    private static readonly pb::FieldCodec<long> _repeated_uid_codec
        = pb::FieldCodec.ForInt64(10);
    private readonly pbc::RepeatedField<long> uid_ = new pbc::RepeatedField<long>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<long> Uid {
      get { return uid_; }
    }

    /// <summary>Field number for the "server" field.</summary>
    public const int ServerFieldNumber = 2;
    private string server_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Server {
      get { return server_; }
      set {
        server_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as LinkGame4GateReq);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(LinkGame4GateReq other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if(!uid_.Equals(other.uid_)) return false;
      if (Server != other.Server) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      hash ^= uid_.GetHashCode();
      if (Server.Length != 0) hash ^= Server.GetHashCode();
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
      uid_.WriteTo(output, _repeated_uid_codec);
      if (Server.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(Server);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      size += uid_.CalculateSize(_repeated_uid_codec);
      if (Server.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Server);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(LinkGame4GateReq other) {
      if (other == null) {
        return;
      }
      uid_.Add(other.uid_);
      if (other.Server.Length != 0) {
        Server = other.Server;
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
          case 10:
          case 8: {
            uid_.AddEntriesFrom(input, _repeated_uid_codec);
            break;
          }
          case 18: {
            Server = input.ReadString();
            break;
          }
        }
      }
    }

  }

  #endregion

}

#endregion Designer generated code
