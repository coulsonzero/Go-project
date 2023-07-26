/**
 * 要求将数组参数中的多维数组扩展为一维数组并返回该数组。
 * input: [1,[2,[3,[4]]]]
 * output: [1,2,3,4]
 */

const _flatten = arr => {
    // 补全代码
    return arr.toString().split(',').map(item => Number(item));
}

const _flatten2 = arr => {
    // 补全代码
    return arr.reduce((pre, item) => {
        return pre.concat((item instanceof Array) ? _flatten(item) : item)
    },[])
}

const _flatten3 = arr => {
    // 补全代码
    return arr.reduce((pre, item) => pre.concat((item instanceof Array) ? _flatten(item) : item),[])
}