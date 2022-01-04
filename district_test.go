package lbs_test

import (
	"github.com/goclub/lbs"
	"github.com/stretchr/testify/assert"
	"testing"
)
func TestDistrict_RelationshipLevelProvince(t *testing.T) {
	d, err  := lbs.NewDistrict(lbs.DistrictDersion20211103) ; assert.NoError(t, err)
	{
		info, has, err := d.Relationship("110000") ; assert.NoError(t, err)
		assert.Equal(t,has, true)
		assert.Equal(t, info, lbs.Relationship{
			Level: lbs.LevelProvince,
			ProvinceFullName:"北京市",
			ProvinceADCode: "110000",
			CityFullName: "",
			CityADCode: "",
			DistrictFullName: "",
			DistrictADCode: "",
		})
	}
	{
		info, has, err := d.Relationship("340000") ; assert.NoError(t, err)
		assert.Equal(t,has, true)
		assert.Equal(t, info, lbs.Relationship{
			Level: lbs.LevelProvince,
			ProvinceFullName:"安徽省",
			ProvinceADCode: "340000",
			CityFullName: "",
			CityADCode: "",
			DistrictFullName: "",
			DistrictADCode: "",
		})
	}
	{
		info, has, err := d.Relationship("310000") ; assert.NoError(t, err)
		assert.Equal(t,has, true)
		assert.Equal(t, info, lbs.Relationship{
			Level: lbs.LevelProvince,
			ProvinceFullName:"上海市",
			ProvinceADCode: "310000",
			CityFullName: "",
			CityADCode: "",
			DistrictFullName: "",
			DistrictADCode: "",
		})
	}
}
func TestDistrict_RelationshipLevelCity(t *testing.T) {
	d, err  := lbs.NewDistrict(lbs.DistrictDersion20211103) ; assert.NoError(t, err)
	{
		info, has, err := d.Relationship("440100") ; assert.NoError(t, err)
		assert.Equal(t,has, true)
		assert.Equal(t, info, lbs.Relationship{
			Level: lbs.LevelCity,
			ProvinceFullName:"广东省",
			ProvinceADCode: "440000",
			CityFullName: "广州市",
			CityADCode: "440100",
			DistrictFullName: "",
			DistrictADCode: "",
		})
	}
	{
		info, has, err := d.Relationship("440200") ; assert.NoError(t, err)
		assert.Equal(t,has, true)
		assert.Equal(t, info, lbs.Relationship{
			Level: lbs.LevelCity,
			ProvinceFullName:"广东省",
			ProvinceADCode: "440000",
			CityFullName: "韶关市",
			CityADCode: "440200",
			DistrictFullName: "",
			DistrictADCode: "",
		})
	}
	{
		info, has, err := d.Relationship("540100") ; assert.NoError(t, err)
		assert.Equal(t,has, true)
		assert.Equal(t, info, lbs.Relationship{
			Level: lbs.LevelCity,
			ProvinceFullName:"西藏自治区",
			ProvinceADCode: "540000",
			CityFullName: "拉萨市",
			CityADCode: "540100",
			DistrictFullName: "",
			DistrictADCode: "",
		})
	}
}
func TestDistrict_RelationshipLevelDistrict(t *testing.T) {
	d, err  := lbs.NewDistrict(lbs.DistrictDersion20211103) ; assert.NoError(t, err)
	{
		info, has, err := d.Relationship("310101") ; assert.NoError(t, err)
		assert.Equal(t,has, true)
		assert.Equal(t, info, lbs.Relationship{
			Level: lbs.LevelDistrict,
			ProvinceFullName:"上海市",
			ProvinceADCode: "310000",
			CityFullName: "上海市",
			CityADCode: "310000",
			DistrictFullName: "黄浦区",
			DistrictADCode: "310101",
		})
	}
	{
		info, has, err := d.Relationship("340123") ; assert.NoError(t, err)
		assert.Equal(t,has, true)
		assert.Equal(t, info, lbs.Relationship{
			Level: lbs.LevelDistrict,
			ProvinceFullName:"安徽省",
			ProvinceADCode: "340000",
			CityFullName: "合肥市",
			CityADCode: "340100",
			DistrictFullName: "肥西县",
			DistrictADCode: "340123",
		})
	}
	{
		info, has, err := d.Relationship("440311") ; assert.NoError(t, err)
		assert.Equal(t,has, true)
		assert.Equal(t, info, lbs.Relationship{
			Level: lbs.LevelDistrict,
			ProvinceFullName:"广东省",
			ProvinceADCode: "440000",
			CityFullName: "深圳市",
			CityADCode: "440300",
			DistrictFullName: "光明区",
			DistrictADCode: "440311",
		})
	}
	{
		info, has, err := d.Relationship("441900") ; assert.NoError(t, err)
		assert.Equal(t,has, true)
		assert.Equal(t, info, lbs.Relationship{
			Level: lbs.LevelDistrict,
			IsWithoutDistrictCity: true,
			ProvinceFullName:"广东省",
			ProvinceADCode: "440000",
			CityFullName: "东莞市",
			CityADCode: "441900",
			DistrictFullName: "东莞市",
			DistrictADCode: "441999",
		})
	}
	{
		info, has, err := d.Relationship("441999") ; assert.NoError(t, err)
		assert.Equal(t,has, true)
		assert.Equal(t, info, lbs.Relationship{
			Level: lbs.LevelDistrict,
			IsWithoutDistrictCity: true,
			ProvinceFullName:"广东省",
			ProvinceADCode: "440000",
			CityFullName: "东莞市",
			CityADCode: "441900",
			DistrictFullName: "东莞市",
			DistrictADCode: "441999",
		})
	}
}

func TestCheckDistricts(t *testing.T) {
	d, err  := lbs.NewDistrict(lbs.DistrictDersion20211103) ; assert.NoError(t, err)
	for _, item := range d.RelationshipList {
		if item.Level == lbs.LevelDistrict {
			provincePart := item.DistrictADCode[0:2]
			cityPart := item.DistrictADCode[2:4]
			districtPart := item.DistrictADCode[4:6]
			assert.Equal(t,item.ProvinceADCode, provincePart+"0000", item.DistrictADCode)
			// 不是直辖市的区才需要判断 city ADCode
			if item.ProvinceADCode != item.CityADCode {
				assert.Equal(t,item.CityADCode, provincePart + cityPart + "00", item.DistrictADCode)
				continue
			}
			if item.IsWithoutDistrictCity {
				assert.Equal(t, item.DistrictADCode[4:6], "99")
			}
			assert.NotEqualf(t, t, districtPart, "0000")
		}
	}
	assert.Equal(t, d.WithoutDistrictCity, map[string]lbs.Relationship{
		"441900": lbs.Relationship{
			Level:                 lbs.LevelDistrict,
			IsWithoutDistrictCity: true,
			ProvinceFullName:"广东省",
			ProvinceADCode: "440000",
			CityFullName: "东莞市",
			CityADCode: "441900",
			DistrictFullName: "东莞市",
			DistrictADCode: "441999",
		},
		"442000": lbs.Relationship{
			Level:                 lbs.LevelDistrict,
			IsWithoutDistrictCity: true,
			ProvinceFullName:"广东省",
			ProvinceADCode: "440000",
			CityFullName: "中山市",
			CityADCode: "442000",
			DistrictFullName: "中山市",
			DistrictADCode: "442099",
		},
		"460400": lbs.Relationship{
			Level:                 lbs.LevelDistrict,
			IsWithoutDistrictCity: true,
			ProvinceFullName:"海南省",
			ProvinceADCode: "460000",
			CityFullName: "儋州市",
			CityADCode: "460400",
			DistrictFullName: "儋州市",
			DistrictADCode: "460499",
		},
		"620200": lbs.Relationship{
			Level:                 lbs.LevelDistrict,
			IsWithoutDistrictCity: true,
			ProvinceFullName:"甘肃省",
			ProvinceADCode: "620000",
			CityFullName: "嘉峪关市",
			CityADCode: "620200",
			DistrictFullName: "嘉峪关市",
			DistrictADCode: "620299",
		},
	})
}