package balance

// BalanceConstant 定义了负载均衡相关的常量
type BalanceConstant struct{}

// 默认轮询策略
const DefaultStrategy = "com.rrtv.rpc.core.balancer.FullRoundBalance"