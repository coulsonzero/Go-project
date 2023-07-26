const _getExFilename = (filename) => {
    // 补全代码
    return filename.slice(filename.lastIndexOf('.'))
}

const _getExFilename2 = (filename) => {
    return "." + filename.split(".")[1]
}


console.log(_getExFilename("problem.xml"))
