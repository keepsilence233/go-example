package obj_pool

import (
	"errors"
	"time"
)

// Connection 连接信息
type Connection struct {
}

// ConnectionPool 连接对象池
type ConnectionPool struct {
	bufChan chan *Connection
}

// NewConnectionPool 创建连接池
// numOfObj channel 大小
func NewConnectionPool(numOfObj int) *ConnectionPool {
	connPool := ConnectionPool{}
	connPool.bufChan = make(chan *Connection, numOfObj)
	for i := 0; i < numOfObj; i++ {
		connPool.bufChan <- &Connection{}
	}
	return &connPool
}

// GetConnection 获取连接
func (p *ConnectionPool) GetConnection(timeout time.Duration) (*Connection, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("time out") //超时控制
	}
}

// ReleaseConnection 释放连接
func (p *ConnectionPool) ReleaseConnection(conn *Connection) error {
	select {
	case p.bufChan <- conn:
		return nil
	default:
		return errors.New("overflow")
	}
}
