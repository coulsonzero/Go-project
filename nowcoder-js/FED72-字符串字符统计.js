/**
 * input: 'hello world'
 * output: {h: 1, e: 1, l: 3, o: 2, w: 1, r: 1, d: 1}
 */

function count(str) {
    let m = {}
    for (let i = 0; i < str.length; i++) {
        m[str[i]] = m[str[i]] ? m[str[i]] + 1 : 1
    }

    return m
}