package ws

import (
	"fmt"
	"sync"
	"time"
)

type Manager struct {
	Clients     map[*Client]bool   // 全部的连接
	ClientsLock sync.RWMutex       // 读写锁
	Users       map[string]*Client // 登录的用户 // appId+uuid
	UserLock    sync.RWMutex       // 读写锁
	Register    chan *Client       // 连接连接处理
	Login       chan *Login        // 用户登录处理
	Unregister  chan *Client       // 断开连接处理程序
	Broadcast   chan []byte        // 广播 向全部成员发送数据
}

func NewManager() (manager *Manager) {
	manager = &Manager{
		Clients:    make(map[*Client]bool),
		Users:      make(map[string]*Client),
		Register:   make(chan *Client, 1000),
		Login:      make(chan *Login, 1000),
		Unregister: make(chan *Client, 1000),
		Broadcast:  make(chan []byte, 1000),
	}
	return
}

// 获取用户key
func GetUserKey(appId uint32, userId string) (key string) {
	key = fmt.Sprintf("%d_%s", appId, userId)
	return
}

///**************************  manager  ***************************************/

func (m *Manager) InClient(c *Client) (ok bool) {
	m.ClientsLock.RLock()
	defer m.ClientsLock.RUnlock()
	// 连接存在，在添加
	_, ok = m.Clients[c]
	return ok
}

func (m *Manager) GetClients() (clients map[*Client]bool) {
	clients = make(map[*Client]bool)
	m.ClientsRanage(func(c *Client, b bool) (result bool) {
		clients[c] = b
		return true
	})
	return
}

func (m *Manager) ClientsRanage(f func(c *Client, b bool) (result bool)) {
	for k, v := range m.Clients {
		if f(k, v) == false {
			return
		}
	}
}

func (m *Manager) GetClientsLen() int {
	return len(m.Clients)
}

func (m *Manager) AddClient(c *Client) {
	m.ClientsLock.Lock()
	defer m.ClientsLock.Unlock()
	m.Clients[c] = true
}

func (m *Manager) DelClient(c *Client) {
	m.ClientsLock.Lock()
	defer m.ClientsLock.Unlock()
	if _, ok := m.Clients[c]; ok {
		delete(m.Clients, c)
	}
}

func (m *Manager) GetUserClient(appId uint32, userId string) *Client {

	m.UserLock.RLock()
	defer m.UserLock.RUnlock()

	k := GetUserKey(appId, userId)
	if v, ok := m.Users[k]; ok {
		return v
	}
	return nil
}

// 用户数
func (m *Manager) GetUsersLen() int {
	return len(m.Users)
}

// 添加用户
func (m *Manager) AddUser(key string, c *Client) {
	m.UserLock.Lock()
	defer m.UserLock.Unlock()
	m.Users[key] = c
}

// 删除用户
func (m *Manager) DelUser(c *Client) (result bool) {
	m.UserLock.Lock()
	defer m.UserLock.Unlock()
	k := GetUserKey(c.AppId, c.UserId)
	if v, ok := m.Users[k]; ok {
		if v.Addr != c.Addr {
			return
		}
		delete(m.Users, k)
		return true
	}
	return
}

// 获取用户的key
func (m *Manager) GetUserKeys() (keys []string) {
	keys = make([]string, 0)
	m.UserLock.RLock()
	defer m.UserLock.RUnlock()
	for k := range m.Users {
		keys = append(keys, k)
	}
	return
}

func (m *Manager) GetUserList(appId uint32) (users []string) {
	users = make([]string, 0)
	m.UserLock.RLock()
	defer m.UserLock.RUnlock()

	for _, v := range m.Users {
		if v.AppId == appId {
			users = append(users, v.UserId)
		}
	}
	return
}

func (m *Manager) GetUserClients() (cs []*Client) {
	cs = make([]*Client, 0)
	m.UserLock.RLock()
	defer m.UserLock.RUnlock()
	for _, v := range m.Users {
		cs = append(cs, v)
	}
	return
}

func (m *Manager) sendAll(msg []byte, ignoreClient *Client) {
	cs := m.GetUserClients()
	for _, conn := range cs {
		if conn != ignoreClient {
			conn.SendMsg(msg)
		}
	}
}

func (m *Manager) sendAppIdAll(msg []byte, appId uint32, ignoreClient *Client) {
	cs := m.GetUserClients()
	for _, conn := range cs {
		if conn != ignoreClient && conn.AppId == appId {
			conn.SendMsg(msg)
		}
	}
}

func (m *Manager) EventRegister(c *Client) {
	m.AddClient(c)
	c.Send <- []byte("链接成功")
}

// 用户登录
func (m *Manager) EventLogin(l *Login) {
	c := l.Client
	if m.InClient(c) {
		key := l.GetKey()
		m.AddUser(key, l.Client)
	}
}

func (m *Manager) EventUnregister(c *Client) {
	m.DelClient(c)
	if !m.DelUser(c) {
		return
	}
}

func (m *Manager) start() {
	for {
		select {
		case conn := <-m.Register:
			m.EventRegister(conn)
		case login := <-m.Login:
			m.EventLogin(login)
		case conn := <-m.Unregister:
			m.EventUnregister(conn)
		case msg := <-m.Broadcast:
			cs := m.GetClients()
			for conn := range cs {
				select {
				case conn.Send <- msg:
				default:
					close(conn.Send)
				}
			}
		}
	}
}

func GetManagerInfo(isDebug string) (managerInfo map[string]interface{}) {
	managerInfo = make(map[string]interface{})
	managerInfo["clientsLen"] = ClientManager.GetClientsLen()        // 客户端连接数
	managerInfo["userLen"] = ClientManager.GetUsersLen()             // 登录用户数
	managerInfo["chanRegisterLen"] = len(ClientManager.Register)     // 未处理连接事件数
	managerInfo["chanLoginLen"] = len(ClientManager.Login)           // 未处理登录事件数
	managerInfo["chanUnregisterLen"] = len(ClientManager.Unregister) // 未处理退出登录事件数
	managerInfo["chanBroadcastLen"] = len(ClientManager.Broadcast)   // 未处理广播事件数

	if isDebug == "true" {
		addrList := make([]string, 0)
		ClientManager.ClientsRanage(func(c *Client, b bool) (result bool) {
			addrList = append(addrList, c.Addr)
			return true
		})
		users := ClientManager.GetUserKeys()
		managerInfo["clients"] = addrList // 客户端列表
		managerInfo["users"] = users      // 登录用户列表
	}
	return
}

// 获取用户所在的连接
func GetUserClient(appId uint32, userId string) *Client {
	return ClientManager.GetUserClient(appId, userId)
}

// 定时清理超时连接
func ClearTimeoutConnections() {
	currentTime := uint64(time.Now().Unix())
	cs := ClientManager.GetClients()
	for c := range cs {
		if c.IsHeartbeatTimeout(currentTime) {
			c.Socket.Close()
		}
	}
}

// 获取全部用户
func GetUserList(appId uint32) (users []string) {
	return ClientManager.GetUserList(appId)
}

// 全员广播
func AllSendMessages(appId uint32, userId string, data string) {
	ignoreClient := ClientManager.GetUserClient(appId, userId)
	ClientManager.sendAppIdAll([]byte(data), appId, ignoreClient)
}
