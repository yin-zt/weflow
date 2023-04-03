package parser

import (
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"testing"
)

func TestParser(t *testing.T) {
	var data = "[{\"nodeModel\":1,\"nodeName\":\"发起人\",\"nodeId\":\"1640993392605401088\",\"parentId\":\"\"},{\"nodeModel\":2,\"nodeName\":\"审批\",\"nodeId\":\"1640993449224310784\",\"parentId\":\"\",\"perData\":\"{\\\\\\\"1640993433239818240\\\\\\\":\\\\\\\"2\\\\\\\",\\\\\\\"1640993434883985408\\\\\\\":\\\\\\\"2\\\\\\\"}\",\"handlerList\":\"[{\\\\\\\"handlerId\\\\\\\":\\\\\\\"547\\\\\\\",\\\\\\\"handlerName\\\\\\\":\\\\\\\"xuch01\\\\\\\",\\\\\\\"handlerType\\\\\\\":1,\\\\\\\"handlerSort\\\\\\\":1}]\"},{\"nodeModel\":6,\"nodeName\":\"分支节点\",\"nodeId\":\"1640993508049424384\",\"parentId\":\"\",\"children\":[[{\"nodeModel\":5,\"nodeName\":\"条件1\",\"nodeId\":\"1640993508049424385\",\"parentId\":\"1640993508049424384\"},{\"nodeModel\":2,\"nodeName\":\"审批\",\"nodeId\":\"1640993526328201216\",\"parentId\":\"1640993508049424384\",\"perData\":\"{\\\\\\\"1640993433239818240\\\\\\\":\\\\\\\"2\\\\\\\",\\\\\\\"1640993434883985408\\\\\\\":\\\\\\\"2\\\\\\\"}\",\"handlerList\":\"[{\\\\\\\"handlerId\\\\\\\":\\\\\\\"547\\\\\\\",\\\\\\\"handlerName\\\\\\\":\\\\\\\"xuch01\\\\\\\",\\\\\\\"handlerType\\\\\\\":1,\\\\\\\"handlerSort\\\\\\\":1}]\"}],[{\"nodeModel\":5,\"nodeName\":\"条件2\",\"nodeId\":\"1640993508049424386\",\"parentId\":\"1640993508049424384\"},{\"nodeModel\":4,\"nodeName\":\"知会\",\"nodeId\":\"1640993535555670016\",\"parentId\":\"1640993508049424384\",\"perData\":\"{\\\\\\\"1640993433239818240\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"1640993434883985408\\\\\\\":\\\\\\\"1\\\\\\\"}\",\"handlerList\":\"[{\\\\\\\"handlerId\\\\\\\":\\\\\\\"547\\\\\\\",\\\\\\\"handlerName\\\\\\\":\\\\\\\"xuch01\\\\\\\",\\\\\\\"handlerType\\\\\\\":1,\\\\\\\"handlerSort\\\\\\\":1}]\"}]]},{\"nodeModel\":7,\"nodeName\":\"分支汇聚\",\"nodeId\":\"1640993508053618688\",\"parentId\":\"\"},{\"nodeModel\":8,\"nodeName\":\"流程结束\",\"nodeId\":\"1640993392605401089\",\"parentId\":\"\"}]"
	nodes := FlowParserServiceImpl.Parser(data)
	marshal, _ := json.Marshal(nodes)

	hlog.Infof("数据为%v", string(marshal))
}

type A struct {
}

type Api interface {
	ATest()
}

type Api2 interface {
	ATest()
}

func (receiver A) ATest() {
	fmt.Println("ATest")
}

func Test(t *testing.T) {
	var a = A{}
	a.ATest()

}
