package Core

import (
	"apiGateway/Config"
	"fmt"
	"github.com/micro/go-micro/registry"
	"math/rand"
	"sort"
	"time"
)

type IpTable struct {
	Addr           string
	Weight         int
	CurrentWeight  int
	CurrentConnect int
}

func NewIpTable(addr string, weight int, currentWeight int, connect int) *IpTable {
	return &IpTable{
		Addr:           addr,
		Weight:         weight,
		CurrentWeight:  currentWeight,
		CurrentConnect: connect,
	}
}

func Test() {
	ipTables := make([]*IpTable, 0)
	ipTables = append(ipTables, NewIpTable("192.168.1.1:8888", 1, 0, 0))
	ipTables = append(ipTables, NewIpTable("192.168.1.2:8888", 2, 0, 0))
	ipTables = append(ipTables, NewIpTable("192.168.1.3:8888", 7, 0, 0))
	for {
		fmt.Println(WeightedLeastConnection(ipTables))
	}
}

func GetAddr(services []*registry.Service, strategy string) string {
	var ips []*IpTable
	var nodes []*registry.Node
	for _, service := range services {
		nodes = append(nodes, service.Nodes...)
	}
	if len(nodes) == 0 {
		return ""
	}
	for _, node := range nodes {
		ips = append(ips, NewIpTable(node.Address, 1, 0, 0))
	}
	// 策略匹配选择服务地址
	switch strategy {
	case Config.Random:
		return Random(ips).Addr
	case Config.WeightedRandom:
		return WeightedRandom(ips).Addr
	case Config.RoundRobin:
		return RoundRobin(ips).Addr
	case Config.WeightedRoundRobin:
		return WeightedRoundRobin(ips).Addr
	case Config.SmoothlyWeightedRoundRobin:
		return SmoothlyWeightedRoundRobin(ips).Addr
	case Config.LeastConnection:
		return LeastConnection(ips).Addr
	case Config.WeightedLeastConnection:
		return WeightedLeastConnection(ips).Addr
	default:
		return Random(ips).Addr
	}
}

// 标记轮询的地址，使用当前轮询地址
var roundRobinIndex = 0

// 轮询
func RoundRobin(ips []*IpTable) *IpTable {
	roundRobinIndex++
	if roundRobinIndex >= len(ips) {
		roundRobinIndex = roundRobinIndex % len(ips)
	}
	return ips[roundRobinIndex]
}

// 标记轮询的地址，使用当前轮询地址
var weightRoundRobinIndex = 0

// 加权轮询
func WeightedRoundRobin(ips []*IpTable) *IpTable {
	var totalWeight = 0
	for _, Ip := range ips {
		totalWeight += Ip.Weight
	}
	nums := weightRoundRobinIndex % totalWeight
	weightRoundRobinIndex++

	// 根据ip切片顺序来排权重区间
	for _, Ip := range ips {
		// 落在当前区间返回
		if Ip.Weight >= nums {
			return Ip
		}
		// 如果未落在当前区间，则减掉当前区间
		nums -= Ip.Weight
	}
	return nil
}

// 平滑加权轮询
func SmoothlyWeightedRoundRobin(ips []*IpTable) *IpTable {
	var index = -1
	var totalWeight = 0

	for i, Ip := range ips {
		Ip.CurrentWeight += Ip.Weight
		totalWeight += Ip.Weight

		if index == -1 || ips[index].CurrentWeight < Ip.CurrentWeight {
			index = i
		}
	}

	ips[index].CurrentWeight -= totalWeight
	return ips[index]
}

// 完全随机
func Random(ips []*IpTable) *IpTable {
	// 使用系统时间作为种子
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(ips))
	return ips[index]
}

// 加权随机
func WeightedRandom(ips []*IpTable) *IpTable {
	var totalWeight = 0
	for _, Ip := range ips {
		totalWeight += Ip.Weight
	}
	rand.Seed(time.Now().UnixNano())
	nums := rand.Intn(totalWeight)
	// 根据ip切片顺序来排权重区间
	for _, Ip := range ips {
		// 落在当前区间返回
		if Ip.Weight >= nums {
			return Ip
		}
		// 如果未落在当前区间，则减掉当前区间
		nums -= Ip.Weight
	}
	return nil
}

func Hash(ips []*IpTable) {

}

// 最少连接
func LeastConnection(ips []*IpTable) *IpTable {
	// 根据连接数升序排序
	sort.Slice(ips, func(i, j int) bool {
		return ips[i].CurrentConnect < ips[j].CurrentConnect
	})
	// 连接加一并返回当前最少连接的ip
	ips[0].CurrentConnect++
	return ips[0]
}

// 加权最少连接
func WeightedLeastConnection(ips []*IpTable) *IpTable {
	var index = 0
	for i, Ip := range ips {
		// 优先级相关性:连接负相关，权重正相关 if W(i)/C(i) > W(j)/C(j)  return ip[i]
		if Ip.Weight*ips[index].CurrentConnect >= Ip.CurrentConnect*ips[index].Weight {
			index = i
		}
	}
	// 连接加一并返回当前最少连接的ip
	ips[index].CurrentConnect++
	return ips[index]
}

func LeastResponseTime(ips []*IpTable) {

}
