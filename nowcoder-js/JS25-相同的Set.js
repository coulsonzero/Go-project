const _isSameSet1 = (s1, s2) => {
    // 补全代码
    // 1. 判断长度是否一致
    if(s1.size !== s2.size) {
        return false
    }
    // 2. ES6[...]扩展字符串：将set对象的伪数组转换为数组，再调用every进而判断
    // 该数组中的每一项是否存在于另一个set对象中
    return [...s1].every(i => s2.has(i))
}

const _isSameSet2 = (s1, s2) => {
    return s1.size === s2.size && Array.from(s1).every(item => s2.has(item))
}

const _isSameSet = (s1, s2) => {
    return JSON.stringify(s1) === JSON.stringify(s2)
}