package lbs

import (
	xjson "github.com/goclub/json"
)

type District struct {
	RelationshipList []Relationship
	WithoutDistrictCity map[/*cityADCode*/string]Relationship
}

func NewDistrict(data AreaData) (d District, err error) {
	d.WithoutDistrictCity = map[string]Relationship{}
	var relationships []Relationship
	var provinces []AreaDataItem
	err = xjson.Unmarshal(data.Province, &provinces) ; if err != nil {
	    return
	}
	var cities []AreaDataItem
	err = xjson.Unmarshal(data.City, &cities) ; if err != nil {
		return
	}
	var districts []AreaDataItem
	err = xjson.Unmarshal(data.District, &districts) ; if err != nil {
		return
	}
	for _, province := range provinces {
		start := province.ChildrenIndex[0]
		end := province.ChildrenIndex[1]+1
		subcities := cities[start:end]
		for _, city := range subcities {
			if len(city.ChildrenIndex) == 0 {
				// 直辖市的区在腾讯的lbs数据中只在 cities 存在
				relationships = append(relationships, Relationship{
					Level: LevelDistrict,
					// 上海市-上海市-黄浦区
					ProvinceFullName: province.FullName, // Province 上海市
					ProvinceADCode:   province.ADCode,
					CityFullName:     province.FullName, // city 上海市
					CityADCode:       province.ADCode,
					DistrictFullName: city.FullName, // district 黄浦区
					DistrictADCode:   city.ADCode,
				})
				continue
			}
			start := city.ChildrenIndex[0]
			end := city.ChildrenIndex[1]+1
			subDistricts := districts[start:end]
			for _, district := range subDistricts {
				isWithoutDistrictCity := len(subDistricts) == 1
				item := Relationship{
					Level: LevelDistrict,
					IsWithoutDistrictCity: isWithoutDistrictCity,
					ProvinceFullName: province.FullName,
					ProvinceADCode:   province.ADCode,
					CityFullName:     city.FullName,
					CityADCode:       city.ADCode,
					DistrictFullName: district.FullName,
					DistrictADCode:   district.ADCode,
				}
				relationships = append(relationships, item)
				if isWithoutDistrictCity {
					d.WithoutDistrictCity[city.ADCode] = item
				}
			}
		}
	}
	d.RelationshipList = relationships
	return
}
type Level uint8
const LevelProvince Level = 1
const LevelCity Level = 2
const LevelDistrict Level = 3
type Relationship struct {
	Level Level
	IsWithoutDistrictCity bool `note:"不设区的地级市"`
	ProvinceFullName string
	ProvinceADCode string
	CityFullName string
	CityADCode string
	DistrictFullName string
	DistrictADCode string
}
func (d District) Relationship(adcode string) (data Relationship, has bool, err error) {
	for _, item := range d.RelationshipList {
		// 边界情况: (4个不设区的地级市) 行政区划如下: 广东省 440000 东莞市 441900 东莞市 441999, (注意我没有打错,就是东莞市下面有个东莞市但是adcode不一样) 有时候腾讯lbs返回的DistrictADCode 是441900
		if item.IsWithoutDistrictCity {
			// 入参东莞市 441900
			if item.CityADCode == adcode {
				target, hasTarget := d.WithoutDistrictCity[item.CityADCode]
				if hasTarget {
					data = target
					has = true
					return
				}
			}
		}
		if item.DistrictADCode == adcode {
			has = true
			data = item
			return
		}
	}
	return
}