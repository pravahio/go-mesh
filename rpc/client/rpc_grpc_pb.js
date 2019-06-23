// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var rpc_pb = require('./rpc_pb.js');

function serialize_rpc_Data(arg) {
  if (!(arg instanceof rpc_pb.Data)) {
    throw new Error('Expected argument of type rpc.Data');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rpc_Data(buffer_arg) {
  return rpc_pb.Data.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rpc_PeerTopicInfo(arg) {
  if (!(arg instanceof rpc_pb.PeerTopicInfo)) {
    throw new Error('Expected argument of type rpc.PeerTopicInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rpc_PeerTopicInfo(buffer_arg) {
  return rpc_pb.PeerTopicInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rpc_PublishData(arg) {
  if (!(arg instanceof rpc_pb.PublishData)) {
    throw new Error('Expected argument of type rpc.PublishData');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rpc_PublishData(buffer_arg) {
  return rpc_pb.PublishData.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_rpc_Response(arg) {
  if (!(arg instanceof rpc_pb.Response)) {
    throw new Error('Expected argument of type rpc.Response');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_rpc_Response(buffer_arg) {
  return rpc_pb.Response.deserializeBinary(new Uint8Array(buffer_arg));
}


var MeshService = exports.MeshService = {
  registerToPublish: {
    path: '/rpc.Mesh/RegisterToPublish',
    requestStream: false,
    responseStream: false,
    requestType: rpc_pb.PeerTopicInfo,
    responseType: rpc_pb.Response,
    requestSerialize: serialize_rpc_PeerTopicInfo,
    requestDeserialize: deserialize_rpc_PeerTopicInfo,
    responseSerialize: serialize_rpc_Response,
    responseDeserialize: deserialize_rpc_Response,
  },
  publish: {
    path: '/rpc.Mesh/Publish',
    requestStream: true,
    responseStream: false,
    requestType: rpc_pb.PublishData,
    responseType: rpc_pb.Response,
    requestSerialize: serialize_rpc_PublishData,
    requestDeserialize: deserialize_rpc_PublishData,
    responseSerialize: serialize_rpc_Response,
    responseDeserialize: deserialize_rpc_Response,
  },
  subscribe: {
    path: '/rpc.Mesh/Subscribe',
    requestStream: false,
    responseStream: true,
    requestType: rpc_pb.PeerTopicInfo,
    responseType: rpc_pb.Data,
    requestSerialize: serialize_rpc_PeerTopicInfo,
    requestDeserialize: deserialize_rpc_PeerTopicInfo,
    responseSerialize: serialize_rpc_Data,
    responseDeserialize: deserialize_rpc_Data,
  },
};

exports.MeshClient = grpc.makeGenericClientConstructor(MeshService);
