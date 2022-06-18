//++++++++++++++++++++++++++++++++++++++++
// 《Go语言高级开发与实战》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 知乎：https://www.zhihu.com/people/shirdonl
// 公众号:源码大数据
// 仓库地址：https://gitee.com/shirdonl/goAdvanced
// 仓库地址：https://github.com/shirdonl/goAdvanced
//++++++++++++++++++++++++++++++++++++++++

package cache

type LocalCache struct {
	LocalInStock     int64
	LocalSalesVolume int64
}

//本地扣库存,返回bool值
func (cache *LocalCache) LocalDeductionStock() bool {
	cache.LocalSalesVolume = cache.LocalSalesVolume + 1
	return cache.LocalSalesVolume <= cache.LocalInStock
}