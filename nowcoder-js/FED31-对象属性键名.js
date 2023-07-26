function _keys(object) {
    // 补全代码
    return Object.keys(object)
}

function _keys2(object) {
    let arr = []
    for(let i in object) {
        arr.push(i)
    }
    return arr
}