# distributedfs
personal tracking repo for the distributed fileserver

PHASE - I = p2p lib (PeerToPeer)
1.1 - Intialized TCP listener and checked with the telnet command (Done)
1.2 - Connections are being Accepted and Handled - minmail funcs (Done)
1.3 - HandShakeFunc Definition i.e some time we have to implement it. Minimal setup, better to have it right now. 
-- refactored func name of Accept and Handle func to make them private (Done)
-- Defined the HandShakeFunc as function callback way but not in interface (more biolerplate for current requirement), TCPTranport struct (to make part of TCPTransport obj. will not allow flexibility for changes) etc. Quick fullfilling requirement using function callback event satisfies the dependency injection mean to have a control to change the behavior i.e replacing defalut handler to real handler or fake test handler. (Done)
-- Also check if the HandShake Func is provided or not - which make us independent like either defined or not, if defined than we use and vice versa - woooooo 