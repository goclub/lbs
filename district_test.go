package lbs_test

import (
	"github.com/goclub/lbs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDistrict_DistrictInfo(t *testing.T) {
	d, err  := lbs.NewDistrict(lbs.DistrictDersion20211103) ; assert.NoError(t, err)
	{
		info, has, err := d.DistrictInfo("310101") ; assert.NoError(t, err)
		assert.Equal(t,has, true)
		assert.Equal(t, info, lbs.DistrictInfo{
			ProvinceFullName:"上海市",
			ProvinceADCode: "310000",
			CityFullName: "上海市",
			CityADCode: "310000",
			DistrictFullName: "黄浦区",
			DistrictADCode: "310101",
		})
	}
	{
		info, has, err := d.DistrictInfo("340123") ; assert.NoError(t, err)
		assert.Equal(t,has, true)
		assert.Equal(t, info, lbs.DistrictInfo{
			ProvinceFullName:"安徽省",
			ProvinceADCode: "340000",
			CityFullName: "合肥市",
			CityADCode: "340100",
			DistrictFullName: "肥西县",
			DistrictADCode: "340123",
		})
	}
	{
		// 340100 合肥市
		info, has, err := d.DistrictInfo("340100") ; assert.NoError(t, err)
		assert.Equal(t,has, false)
		assert.Equal(t, info, lbs.DistrictInfo{
			ProvinceFullName:"",
			ProvinceADCode: "",
			CityFullName: "",
			CityADCode: "",
			DistrictFullName: "",
			DistrictADCode: "",
		})
	}
}

func TestCheckDistricts(t *testing.T) {
	d, err  := lbs.NewDistrict(lbs.DistrictDersion20211103) ; assert.NoError(t, err)
	for _, item := range d.Districts {
		provincePart := item.DistrictADCode[0:2]
		cityPart := item.DistrictADCode[2:4]
		districtPart := item.DistrictADCode[4:6]
		assert.Equal(t,item.ProvinceADCode, provincePart+"0000", item.DistrictADCode)
		// 不是直辖市的区才需要判断 city ADCode
		if item.ProvinceADCode != item.CityADCode {
			assert.Equal(t,item.CityADCode, provincePart + cityPart + "00", item.DistrictADCode)
			continue
		}
		assert.NotEqualf(t, t, districtPart, "0000")
	}
}