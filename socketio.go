package common

import (

	"io"
	"net"
    "time"
	"fmt"
)

//////////////////////////////////////////////////////////////////////
// blocking mode socket read/write
//////////////////////////////////////////////////////////////////////
func MyRead(conn net.Conn, packLen int64, Buffer []byte, timeout int64) (err error) {
	CheckParam(packLen <= int64(cap(Buffer)))
	var receiveLen int64
	if timeout <= 0 {
		conn.SetReadDeadline(time.Time{})
	} else {
		conn.SetReadDeadline(time.Now().Add(time.Duration(timeout)))
	}

	for receiveLen < packLen {
		tempLen, err := conn.Read(Buffer[receiveLen:packLen])
		if err == io.EOF {
			fmt.Printf("peer socket[%s] exit", conn.RemoteAddr())
			return err
		} else if err != nil {
			fmt.Printf("read socket[%s] connection error[%v]", conn.RemoteAddr(), err)
			//Assert(err != os.EAGAIN, "socket not blocking mode")
			return err
		} else {
			receiveLen += int64(tempLen)
		}
	}
	return err
}

func MyWrite(conn net.Conn, packLen int64, Buffer []byte, timeout int64) (err error) {
	CheckParam(packLen <= int64(len(Buffer)) && timeout > 0)

	//conn.SetWriteDeadline(time.Now().UTC().Add(time.Duration(timeout)))

	writtenLen := int64(0)
	for writtenLen < packLen {
		n, err := conn.Write(Buffer[writtenLen:packLen])
		// 这里将所有写入错误均认定为错误；并不区分错误如 os.EAGAIN 等。
		// Note:
		//  https://linux.die.net/man/2/write
		if err != nil {
			fmt.Printf("write socket[%s] connection error[%v]", conn.RemoteAddr(), err)
			return err
		} else if n == 0 {
			// 一般不会出现 n == 0 的情况，如果出现，则可认定为不正常状态，打印告警日志，且返回 EOF 即可；
			//  Reference: http://stackoverflow.com/questions/2176443/is-a-return-value-of-0-from-write2-in-c-an-error
			fmt.Printf("write socket[%s] zero bytes!", conn.RemoteAddr())
			return io.EOF
		} else {
			writtenLen += int64(n)
		}
	}
	return nil
}
