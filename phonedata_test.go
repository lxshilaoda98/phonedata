package phonedata

import (
	"fmt"
	"testing"
	"time"
)

func since(t time.Time) int64 {
	return time.Since(t).Nanoseconds()
}

func init() {

	Debug()
}

func BenchmarkFindPhone(b *testing.B) {

	b.RunParallel(func(p *testing.PB) {

		var i = 0
		for p.Next() {
			i++
			_, err := Find(fmt.Sprintf("%s%d%s", "1897", i&10000, "45"), nil)
			if err != nil {
				b.Fatal(err)
			}
		}

	})

}

func TestFindPhone1(t *testing.T) {

	_, err := Find("13580198235123123213213", nil)
	if err == nil {
		t.Fatal("错误的结果")
	}
	t.Log(err)
}
func TestFindPhone2(t *testing.T) {

	_, err := Find("1300", nil)
	if err == nil {
		t.Fatal("错误的结果")
	}
	t.Log(err)
}
func TestFindPhone3(t *testing.T) {

	pr, err := Find("1703576", nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pr)
}

func TestFindPhone4(t *testing.T) {

	_, err := Find("10074872323", nil)
	if err == nil {
		t.Fatal("错误的结果")
	}
	t.Log(err)
}

func TestFindPhone5(t *testing.T) {

	_, err := Find("afsd32323", nil)
	if err == nil {
		t.Fatal("错误的结果")
	}
	t.Log(err)
}

func TestFindPhone1952947(t *testing.T) {

	//1952947,广西,玉林,移动
	info, err := Find("1952947", nil)
	if err != nil {
		t.Fatal("错误的结果")
	}
	t.Log(info)
	str := fmt.Sprintf("%s,%s,%s,%s", info.PhoneNum, info.Province, info.City, info.CardType)
	if str != "1952947,广西,玉林,中国移动" {
		t.Fatal("验证失败")
	}

}

func TestFindPhone1669981(t *testing.T) {

	//1669981,新疆,乌鲁木齐,联通
	info, err := Find("1669981", nil)
	if err != nil {
		t.Fatal("错误的结果")
	}
	t.Log(info)
	str := fmt.Sprintf("%s,%s,%s,%s", info.PhoneNum, info.Province, info.City, info.CardType)
	if str != "1669981,新疆,乌鲁木齐,中国联通" {
		t.Fatal("验证失败")
	}

}
