package Controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"reflect"
	"strings"
	"web-marisa/server/Middlewares/pkg"
	"web-marisa/server/Middlewares/segment"
	"web-marisa/server/Models"
	"web-marisa/server/Services"
)

func Add(ctx iris.Context) {
	memory := Models.Memorise{}

	err := ctx.ReadForm(&memory)

	if err != nil {
		ctx.JSON(context.Map{
			"code": 400,
			"data": err.Error(),
		})
		fmt.Errorf("Controller Add() error: %s", err)
	} else {
		toPpl := segment.Init().Cut(memory.Keyword)
		memorise := Services.FetchAllMemory()
		var real string

		if len(memorise) == 0 {
			real = pkg.Join(toPpl, ",")
			goto DATA
		}

		for _, v := range memorise {
			ratio := 0
			keywords := strings.Split(v.Keyword, ",")
			for _, keyword := range keywords {
				for _, ppl := range toPpl {
					if keyword == ppl {
						ratio++
					}
				}
				if float32(ratio) / float32(len(keywords)) >= 0.6 {
					keywords = append(keywords, toPpl...)
					real = pkg.Join(keywords, ",")
					goto DATA
				} else {
					real = pkg.Join(toPpl, ",")
					goto DATA
				}
			}
		}
		DATA:
		data := make(map[string]interface{})
		data["ip"] = memory.Ip
		data["keyword"] = real
		data["answer"] = memory.Answer
		if Services.AddMemory(data) {
			ctx.JSON(context.Map{
				"code": 200,
				"data": data,
			})
		}
	}
}

func Reply(ctx iris.Context) {
	memory := Models.Memorise{}

	err := ctx.ReadForm(&memory)
	if err != nil {
		ctx.JSON(context.Map{
			"code": 400,
			"data": err.Error(),
		})
		fmt.Errorf("Controller Reply() error: %s", err)
	} else {
		data := make(map[string]interface{})
		toPpl := segment.Init().Cut(memory.Keyword)
		memorise := Services.FetchAllMemory()
		var answer string

		if len(memorise) == 0 {
			data["answer"] = "唔嗯...不懂你在说什么呢...教教我吧~"
			ctx.JSON(context.Map{
				"code": 10001,
				"data": data,
			})
			goto END
		}

		for _, v := range memorise {
			fmt.Println(v)
			ratio := 0
			keywords := strings.Split(v.Keyword, ",")
			for _, keyword := range keywords {
				for _, ppl := range toPpl {
					if keyword == ppl {
						ratio++
					}
				}
				if float32(ratio) / float32(len(keywords)) >= 0.6 {
					answer = v.Answer
					goto DATA
				}
			}
		}
		if answer == "" {
			data["answer"] = "唔嗯...不懂你在说什么呢...教教我吧~"
			ctx.JSON(context.Map{
				"code": 10001,
				"data": data,
			})
			goto END
		}
		DATA:
		temp := Services.FetchMemory(answer)
		data["answer"] = temp.Answer
		ctx.JSON(context.Map{
			"code": 200,
			"data": data,
		})

	}
	END:
}

func Forget(ctx iris.Context) {
	memory := Models.Memorise{}

	err := ctx.ReadForm(&memory)
	if err != nil {
		ctx.JSON(context.Map{
			"code": 400,
			"data": err.Error(),
		})
		fmt.Errorf("Controller Forget() error: %s", err)
	} else {
		if Services.DeleteMemoryByAnswer(memory.Answer) {
			ctx.JSON(context.Map{
				"code": 200,
				"data": "success",
			})
		}
	}
}

func Test(ctx iris.Context) {
	all := Services.FetchAllMemory()
	fmt.Println(reflect.TypeOf(all))
	str := []string{"你好", "再见", "再见"}
	str1 := []string{"fuck", "hello", "你好"}
	for _, v := range all {
		fmt.Println(v.Ip)
		fmt.Println(v.Answer)
		fmt.Println(v.Keyword)
		fmt.Printf("%s\n", strings.Split(v.Keyword, ","))
		fmt.Println(reflect.TypeOf(strings.Split(v.Keyword, ",")))
		fmt.Printf("%s\n", pkg.DuplicateRemove(str))
	}
	fmt.Printf("%s\n", append(str, str1...))
	fmt.Printf("%s\n", pkg.DuplicateRemove(append(str, str1...)))
	fmt.Println("Join: ", pkg.Join(str, ","))
	ctx.JSON(context.Map{
		"msg": reflect.TypeOf(all),
		"data": all,
	})
}
