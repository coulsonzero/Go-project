/**
 * 如果第二个参数 bUnicode255For1 === true，则所有字符长度为 1
 * 否则如果字符 Unicode 编码 > 255 则长度为 2
 * input: 'hello world, 牛客', false
 * output: 17
 */

function strLength(s, bUnicode255For1) {
    if (bUnicode255For1) {
        return s.length
    }

    let length = 0
    for (let i = 0; i < s.length; i++) {
        length += s.charCodeAt(i) > 255 ? 2 : 1
    }
    return length
}

function strLength2(s, bUnicode255For1) {
    let res = s.length
    if(bUnicode255For1 !== true) {
        for(let i in s){
            if(s.charCodeAt(i) > 255) res++
        }
    }
    return res
}