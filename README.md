# distributedfs
personal tracking repo for the distributed fileserver

## PHASE - I = p2p lib (PeerToPeer)

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

-- BUG (NOT FIXED) - but this fix doesnt ensure that the error happens due to decoding logic (for which we can retry) or does it happens due to the coonection problem. Its not a good fix. :( 

1.8 - Peer has been been implmeneted for TCP (Done)

-- Peer is nothing just the second name of the conn (connection) with more option as prevous testing was only based on the connection - we just change the name as our p2p architechure and to be intutive like it hard to mention the connection wording rather then peer

-- Right now, we converted or called the connection a peer when the connection being able to established or accepted successfully

-- Improvement (NOT FIXED) : 

--- TCPPeer is in the separate file peer.go -> it should be in the tcp_transport file but i think it should be fine for now - im considering because other files are kinda generic for other protocols not tightly bound to TCP only 

--- handleNewConnection can be refactor for variable naming convention - can be made better 

1.9 - Replaced the printing RPC Message option to the Channel (Done)

--- Channel would be the right way to read the message of the Peer, as this channel will be associated to every Peer 

--- Consume() will only help to provide the Channel outside the transport lib, so that we can read from the channel. This would be interface, as other protocol should have this too

--- added the Learning.md file for some observation and learning like importance of initialization 

1.10 - Notify Peer logic just added but not defined, and some refactor 

---  OnPeer func - for just in case we need to take any action for particular or overall peer when its added 

--- Created the Peer Interface, as other protocol will also define its own Peer. As the Current TCPPeer is only have Close() func, so we in other function we are just using Peer as a type because interface Peer does satisfy our TCP peer - concept of polymorphism 

## PHASE - II = Store

2.1 - Write a Simple writeStream - which create a file on the disk 