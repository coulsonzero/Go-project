/**
 * 根据输入的数字范围[start,end]和随机数个数"n"生成随机数
 * 数组不能有相同元素
 * Math.floor(): 左闭右开
 * Math.round(): 右闭
 */

const _getUniqueNums = (start,end,n) => {
    // 补全代码
    let arr = []
    while (arr.length < n) {
        let v = Math.floor(Math.random() * (end - start) + start)
        if (!arr.includes(v)) {
            arr.push(v)
        }
    }
    return arr
}

console.log(_getUniqueNums(2, 10, 6))