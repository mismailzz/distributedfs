# distributedfs
personal tracking repo for the distributed fileserver

PHASE - I = p2p lib (PeerToPeer)

1.1 - Intialized TCP listener and checked with the telnet command (Done)

1.2 - Connections are being Accepted and Handled - minmail funcs (Done)

1.3 - HandShakeFunc Definition i.e some time we have to implement it. Minimal setup, better to have it right now. (Done)

-- refactored func name of Accept and Handle func to make them private (Done)

-- Defined the HandShakeFunc as function callback way but not in interface (more biolerplate for current requirement), TCPTranport struct (to make part of TCPTransport obj. will not allow flexibility for changes) etc. Quick fullfilling requirement using function callback event satisfies the dependency injection mean to have a control to change the behavior i.e replacing defalut handler to real handler or fake test handler. (Done)

-- Also check if the HandShake Func is provided or not - which make us independent like either defined or not, if defined than we use and vice versa - woooooo 

1.4 - Formatting README

1.5 - Reading the Message (RPC) from the connection in the HandleConnection func (Done)

1.6 - Implementing the Simple Decoder (Done)

-- BUG (Not Fixed): After breaking the connection from the client terminal, this server side for that connection doesnt close the conenction and stuck in loop i.e Error decoding message from [::1]:56450: EOF

-- In this we covered second example of dependency injection like using interface ( and declaring the defaultObj) - i think it can be done by func fallback approach as like HandShakeFunc for now but we opted this method because interface allows more room of change - for now we dont need but better to have seen this as an example usage. Also refer, if an opbject implement interface then it can ack like interface i.e conn in broder term and can be used as io.Reader etc

1.7 - BUG Mitigation - As connetion get error during decoding then we just close the connection (Done)

-- Prevented infinite loop 

-- BUG - but this fix doesnt ensure that the error happens due to decoding logic (for which we can retry) or does it happens due to the coonection problem. Its not a good fix. :(