在 IDEA 中下载 leetcode plugins。 
而后在 setting-Tool 中编辑 leetcode plugins
选择对应的变成语言，并且打开 custom Template，将模板改为下面的值。


Code File Name：
```
$!{question.frontendQuestionId}.${question.titleSlug}/solution_test
```

Code Template：
```
package leetcode
import(
"testing"
)

${question.content}

${question.code}

func Test$!velocityTool.camelCaseName(${question.titleSlug})(t *testing.T){

}
```


