function _getday(value) {
    // 补全代码
    let week = ['天', '一', '二', '三', '四', '五', '六'] ;
    return `星期${week[value % 7]}`
}

