# goclub/lbs


数据基于: [lbs.qq.com](https://lbs.qq.com/service/webService/webServiceGuide/webServiceDistrict)

```go
package main

import lbs "github.com/goclub/lbs"

func main () {
	// NewDistrict 有性能消耗,不要每次都调用
	d, err  := lbs.NewDistrict(lbs.DistrictDersion20211103) ; if err != nil {
		return
	}
	relationship, has, err := d.Relationship("310101"); if err != nil {
		return
	}
	/*
	assert.Equal(t,has, true)
	assert.Equal(t, relationship, lbs.Relationship{
		Level: lbs.LevelDistrict,
		ProvinceFullName:"上海市",
		ProvinceADCode: "310000",
		CityFullName: "上海市",
		CityADCode: "310000",
		DistrictFullName: "黄浦区",
		DistrictADCode: "310101",
	})
	*/
	
}

```