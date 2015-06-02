package canopus

import (
	"log"
	"net"
)

type CoapClient struct {
	localAddr  *net.UDPAddr
	remoteAddr *net.UDPAddr
	conn       *net.UDPConn

	evtOnError   EventHandler
	evtOnStartup EventHandler
	evtOnClose   EventHandler
}

func (c CoapClient) Dial(host string) {
	remAddr, err := net.ResolveUDPAddr("udp", host)
	IfErr(err)
	c.remoteAddr = remAddr

	conn, err := net.DialUDP("udp", c.localAddr, c.remoteAddr)
	IfErr(err)
	c.conn = conn
}

func (c *CoapClient) doSend(req *CoapRequest, conn *net.UDPConn) (*CoapResponse, error) {
	log.Println(req, conn)
	resp, err := SendMessage(req.GetMessage(), conn)

	return resp, err
}

func (c *CoapClient) Send(req *CoapRequest) (*CoapResponse, error) {
	log.Println("@@", req, c.conn)
	return c.doSend(req, c.conn)
}

func (c *CoapClient) SendTo(req *CoapRequest, conn *net.UDPConn) (*CoapResponse, error) {
	return c.doSend(req, conn)
}

func (c *CoapClient) SendAsync(req *CoapRequest, fn ResponseHandler) {

}

func (c *CoapClient) Close() {
	c.conn.Close()
}

func (c *CoapClient) callEvent(eh EventHandler) {

}

func (c *CoapClient) OnStartup(eh EventHandler) {
	c.evtOnStartup = eh
}

func (c *CoapClient) OnError(eh EventHandler) {
	c.evtOnError = eh
}

func (c *CoapClient) OnClose(eh EventHandler) {
	c.evtOnClose = eh
}
