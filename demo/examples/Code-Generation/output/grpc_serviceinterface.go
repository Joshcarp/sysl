// Code generated by sysl DO NOT EDIT.
 package simple 
 
 import (
 "context" 
 simplegrpc "github.service.anz/anzx/bff/gen/simple/grpc" 
 
)
 
 // GrpcServiceInterface for Simple 
 type GrpcServiceInterface struct {
 	 GetStuffList func ( ctx context.Context , req *simplegrpc.GetStuffListRequest , client GetStuffListClient ) ( *simplegrpc.GetStuffListResponse , error ) 
 }

 
 