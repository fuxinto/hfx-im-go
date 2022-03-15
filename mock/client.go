package mock

import (
	"HFXIM/common/HIM"
	"HFXIM/service/gate/him"
	"HFXIM/service/gate/him/tcp"
	"HFXIM/service/gate/him/websocket"
	"context"
	"github.com/gobwas/ws"
	"github.com/golang/protobuf/proto"
	"github.com/zeromicro/go-zero/core/logx"
	"net"
	"time"
)

// ClientDemo Client demo
type ClientDemo struct {
}

func (c *ClientDemo) Start(userID, protocol, addr string) {
	var cli him.Client

	// step1: 初始化客户端
	if protocol == "ws" {
		cli = websocket.NewClient(userID, "client", websocket.ClientOptions{})
		// set dialer
		cli.SetDialer(&WebsocketDialer{})
	} else if protocol == "tcp" {
		cli = tcp.NewClient(userID, "client", tcp.ClientOptions{Heartbeat: time.Second * 30})
		cli.SetDialer(&TCPDialer{})
	}
	// step2: 建立连接
	err := cli.Connect(addr)

	if err != nil {
		logx.Error(err)
	}
	data, err := proto.Marshal(&HIM.LoginReq{Token: "测试服务token"})
	if err != nil {
		logx.Error(err)
	}
	pack := &HIM.Pack{
		Type: HIM.Pack_loginReq,
		Body: data,
	}
	data, err = proto.Marshal(pack)
	if err != nil {
		logx.Error(err)
	}
	cli.Send(data)

	//count := 10
	//go func() {
	//	// step3: 发送消息然后退出
	//	for i := 0; i < count; i++ {
	//		err := cli.Send([]byte(fmt.Sprintf("hello_%d", i)))
	//		if err != nil {
	//			logx.Error(err)
	//			return
	//		}
	//		time.Sleep(time.Millisecond * 10)
	//	}
	//}()

	// step4: 接收消息
	//recv := 0
	for {
		frame, err := cli.Read()
		if err != nil {
			logx.Info(err)
			break
		}

		if frame.GetOpCode() != him.OpBinary {
			continue
		}
		//recv++
		logx.Infof("%s receive message", frame.GetPayload())
		//if recv == count { // 接收完消息
		//	break
		//}
	}
	//退出
	cli.Close()
}

// WebsocketDialer WebsocketDialer
type WebsocketDialer struct {
}

// DialAndHandshake DialAndHandshake
func (d *WebsocketDialer) DialAndHandshake(ctx him.DialerContext) (net.Conn, error) {
	// 1 调用ws.Dial拨号
	conn, _, _, err := ws.Dial(context.TODO(), ctx.Address)
	if err != nil {
		return nil, err
	}
	// 2. 发送用户认证信息，示例就是userid
	////err = wsutil.WriteClientBinary(conn, []byte(ctx.Id))
	//if err != nil {
	//	return nil, err
	//}
	// 3. return conn
	return conn, nil
}

// TCPDialer TCPDialer
type TCPDialer struct {
}

// DialAndHandshake DialAndHandshake
func (d *TCPDialer) DialAndHandshake(ctx him.DialerContext) (net.Conn, error) {
	logx.Info("start dial: ", ctx.Address)
	// 1 调用net.Dial拨号
	conn, err := net.Dial("tcp", ctx.Address)
	if err != nil {
		return nil, err
	}
	// 2. 发送用户认证信息，示例就是userid
	//err = tcp.WriteFrame(conn, socket_im.OpBinary, []byte(ctx.Id))
	if err != nil {
		return nil, err
	}
	// 3. return conn
	return conn, nil
}
