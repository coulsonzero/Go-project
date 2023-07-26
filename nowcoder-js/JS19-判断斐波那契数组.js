/**
 * 要求以Boolean的形式返回参数数组是否为斐波那契数列。
 * 在数学上，斐波那契数列以如下方法定义：F(0)=0，F(1)=1, F(n)=F(n - 1)+F(n - 2)（n ≥ 2，n ∈ N）
 * 注意：1. [0,1,1]为最短有效斐波那契数列
 */

const _isFibonacci = array => {
    // 补全代码
    if (array.length < 3) return false

    for (let i = 2; i < array.length; i++) {
        if (array[i] !== array[i-1] + array[i-2]) {
            return false
        }
    }

    return true
}

const _isFibonacci2 = array => {
    // 补全代码
    for (let i = 2; i < array.length; i++) {
        if (array[i] !== array[i-1] + array[i-2]) {
            return false
        }
    }

    return array.length > 1
}