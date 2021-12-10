package lbs

import xjson "github.com/goclub/json"

type District struct {
	Districts []DistrictInfo
}

func NewDistrict(data AreaData) (d District, err error) {
	var districtInfos []DistrictInfo
	var provinces []AreaDataItem
	err = xjson.Unmarshal(data.Province, &provinces) ; if err != nil {
	    return
	}
	var citys []AreaDataItem
	err = xjson.Unmarshal(data.City, &citys) ; if err != nil {
		return
	}
	var districts []AreaDataItem
	err = xjson.Unmarshal(data.District, &districts) ; if err != nil {
		return
	}
	for _, province := range provinces {
		start := province.ChildrenIndex[0]
		end := province.ChildrenIndex[1]+1
		subCitys := citys[start:end]
		for _, city := range subCitys {
			if len(city.ChildrenIndex) == 0 {
				// 直辖市的区在腾讯的lbs数据中只在 citys 存在
				districtInfos = append(districtInfos, DistrictInfo{
					ProvinceFullName: province.FullName,
					ProvinceADCode:   province.ADCode,
					CityFullName:     province.FullName,
					CityADCode:       province.ADCode,
					// 上海市-上海市-黄浦区
					DistrictFullName: city.FullName,
					DistrictADCode:   city.ADCode,
				})
				continue
			}
			start := city.ChildrenIndex[0]
			end := city.ChildrenIndex[1]+1
			subDistricts := districts[start:end]
			for _, district := range subDistricts {
				districtInfos = append(districtInfos, DistrictInfo{
					ProvinceFullName: province.FullName,
					ProvinceADCode:   province.ADCode,
					CityFullName:     city.FullName,
					CityADCode:       city.ADCode,
					DistrictFullName: district.FullName,
					DistrictADCode:   district.ADCode,
				})
			}
		}
	}
	d.Districts = districtInfos
	return
}
type DistrictInfo struct {
	ProvinceFullName string
	ProvinceADCode string
	CityFullName string
	CityADCode string
	DistrictFullName string
	DistrictADCode string
}
func (d District) DistrictInfo(districtADCode string) (info DistrictInfo, has bool, err error) {
	for _, district := range d.Districts {
		if district.DistrictADCode == districtADCode {
			has = true
			info = district
			return
		}
	}
	return
}