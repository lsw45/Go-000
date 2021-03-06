// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package rpc

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type RpcService interface {
	// Parameters:
	//  - Req
	FunCall(req *EchoReq) (r []string, err error)
}

type RpcServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewRpcServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *RpcServiceClient {
	return &RpcServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewRpcServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *RpcServiceClient {
	return &RpcServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - Req
func (p *RpcServiceClient) FunCall(req *EchoReq) (r []string, err error) {
	if err = p.sendFunCall(req); err != nil {
		return
	}
	return p.recvFunCall()
}

func (p *RpcServiceClient) sendFunCall(req *EchoReq) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("funCall", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := RpcServiceFunCallArgs{
		Req: req,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *RpcServiceClient) recvFunCall() (value []string, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "funCall" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "funCall failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "funCall failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error1 error
		error1, err = error0.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error1
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "funCall failed: invalid message type")
		return
	}
	result := RpcServiceFunCallResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type RpcServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      RpcService
}

func (p *RpcServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *RpcServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *RpcServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewRpcServiceProcessor(handler RpcService) *RpcServiceProcessor {

	self2 := &RpcServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self2.processorMap["funCall"] = &rpcServiceProcessorFunCall{handler: handler}
	return self2
}

func (p *RpcServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x3.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x3

}

type rpcServiceProcessorFunCall struct {
	handler RpcService
}

func (p *rpcServiceProcessorFunCall) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := RpcServiceFunCallArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("funCall", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := RpcServiceFunCallResult{}
	var retval []string
	var err2 error
	if retval, err2 = p.handler.FunCall(args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing funCall: "+err2.Error())
		oprot.WriteMessageBegin("funCall", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("funCall", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Req
type RpcServiceFunCallArgs struct {
	Req *EchoReq `thrift:"req,1" json:"req"`
}

func NewRpcServiceFunCallArgs() *RpcServiceFunCallArgs {
	return &RpcServiceFunCallArgs{}
}

var RpcServiceFunCallArgs_Req_DEFAULT *EchoReq

func (p *RpcServiceFunCallArgs) GetReq() *EchoReq {
	if !p.IsSetReq() {
		return RpcServiceFunCallArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *RpcServiceFunCallArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *RpcServiceFunCallArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RpcServiceFunCallArgs) readField1(iprot thrift.TProtocol) error {
	p.Req = &EchoReq{}
	if err := p.Req.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Req), err)
	}
	return nil
}

func (p *RpcServiceFunCallArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("funCall_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RpcServiceFunCallArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:req: ", p), err)
	}
	if err := p.Req.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Req), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:req: ", p), err)
	}
	return err
}

func (p *RpcServiceFunCallArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RpcServiceFunCallArgs(%+v)", *p)
}

// Attributes:
//  - Success
type RpcServiceFunCallResult struct {
	Success []string `thrift:"success,0" json:"success,omitempty"`
}

func NewRpcServiceFunCallResult() *RpcServiceFunCallResult {
	return &RpcServiceFunCallResult{}
}

var RpcServiceFunCallResult_Success_DEFAULT []string

func (p *RpcServiceFunCallResult) GetSuccess() []string {
	return p.Success
}
func (p *RpcServiceFunCallResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RpcServiceFunCallResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RpcServiceFunCallResult) readField0(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]string, 0, size)
	p.Success = tSlice
	for i := 0; i < size; i++ {
		var _elem4 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem4 = v
		}
		p.Success = append(p.Success, _elem4)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *RpcServiceFunCallResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("funCall_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RpcServiceFunCallResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.LIST, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRING, len(p.Success)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Success {
			if err := oprot.WriteString(string(v)); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *RpcServiceFunCallResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RpcServiceFunCallResult(%+v)", *p)
}
