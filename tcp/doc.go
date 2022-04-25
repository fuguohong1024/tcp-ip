package main

/*
# https://www.howtouselinux.com/post/tcpdump-capture-packets-with-tcp-flags
TCP Flags List
SYN (synchronize) – The synchronization flag is used as a first step in establishing a three-way handshake between two hosts. Only the first packet from both the sender and receiver should have this flag set.
同步: 同步标志用于在两台主机间建立三次握手的第一步。只有发送方和接受方第一个包才应该设置此标志。
ACK (acknowledgment) – The acknowledgment flag is used to acknowledge the successful receipt of a packet.
确认： 确认标志用于确认成功收到了数据包
FIN (finish) – The finished flag means there is no more data from the sender. Therefore, it is used in the last packet sent from the sender.
完成： 完成标志意味着没有更多的数据发送，即数据发送已结束，用于发送方发送最后一个包。
URG (urgent) – The urgent flag is used to notify the receiver to process the urgent packets before processing all other packets. The receiver will be notified when all known urgent data has been received.
紧急： 紧急标志用于通知接收方在处理所有其他数据包之前处理该紧急数据包。 当接收到所有已知的紧急数据时，将通知接收方。即需先接受所有含紧急标志的包之后再处理其他包。
PSH (push) – The push flag is somewhat similar to the URG flag and tells the receiver to process these packets as they are received instead of buffering them.
推送： push 标志有点类似于 URG 标志，它告诉接收器在接收到这些数据包时对其进行处理，而不是对它们进行缓冲。
RST (reset) – The reset flag gets sent from the receiver to the sender when a packet is sent to a particular host that was not expecting it.
重置： 当数据包被发送到不期望它的特定主机时，重置标志从接收器发送到发送器。重置该链接，因为发错了，或者不希望目标主机链接。
ECE – This flag is responsible for indicating if the TCP peer is ECN capable. See RFC 3168 for more details.
CWR – The congestion window reduced flag is used by the sending host to indicate it received a packet with the ECE flag set. See RFC 3168 for more details.
NS (experimental) – The nonce sum flag is still an experimental flag used to help protect against accidental malicious concealment of packets from the sender. See RFC 3540 for more details.
*/
