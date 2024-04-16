const fs = require('fs');
const jsonData = fs.readFileSync('./build/result.json', 'utf8');
const data = JSON.parse(jsonData);
const intermediateResult = {};
const mapFunction = function (key) {
    this[key] 
    for (let key in this) {
        if (typeof this[key] === 'string') {
            const content = this[key];
            const words = content.match(/[\u4e00-\u9fa5]+/g);
            if (words) {
                words.forEach(word => {
                    if (intermediateResult[word]) {
                        intermediateResult[word]++;
                    } else {
                        intermediateResult[word] = 1;
                    }
                });
            }
        } else if (Array.isArray(this[key])) {
            this[key].forEach(element => {
                if (typeof element === 'string') {
                    const words = element.match(/[\u4e00-\u9fa5]+/g);
                    if (words) {
                        words.forEach(word => {
                            if (intermediateResult[word]) {
                                intermediateResult[word]++;
                            } else {
                                intermediateResult[word] = 1;
                            }
                        });
                    }
                }
            });
        }
    }
};
for (let movie of data) {
    mapFunction.apply(movie);
}
const resultArray = Object.entries(intermediateResult).map(([word, count]) => ({ name:word, value:count }));
resultArray.sort((a, b) => b.count - a.count);

const resultJSON = resultArray.map(item => JSON.stringify(item)).join(',\n');
const finalJSON = `[\n${resultJSON}\n]`;
fs.writeFileSync('anaylyse-result.json', finalJSON, 'utf8');

// 提取需要分析的数据字段
const scores = data.map(item => parseFloat(item.Score)); // 提取评分数据
// 计算均值
const mean = scores.reduce((acc, val) => acc + val, 0) / scores.length;
// 计算方差
const variance = scores.reduce((acc, val) => acc + Math.pow(val - mean, 2), 0) / scores.length;
// 计算标准差
const standardDeviation = Math.sqrt(variance);
// 输出结果
const result = {
  mean: mean,
  variance: variance,
  standardDeviation: standardDeviation
};
console.log("结果：", result);

