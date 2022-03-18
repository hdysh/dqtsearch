package models

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
)

type Unit struct {
	ID        int64       `json:"id"`
	Name      string      `json:"name"`
	FamilyId  int         `json:"-"`
	Family    Family      `gorm:"ForeignKey:FamilyId" json:"family"`
	RoleId    int         `json:"-"`
	Role      Role        `gorm:"ForeignKey:RoleId" json:"role"`
	RarityId  int64       `json:"-"`
	Rarity    Rarity      `gorm:"ForeignKey:RarityId" json:"rarity"`
	Icon      string      `json:"icon"`
	UnitSkill []UnitSkill `json:"skills"`
}

func (u *Unit) FindAllUnit(db *gorm.DB) (*[]Unit, error) {
	var err error
	units := []Unit{}

	db.Preload("Family").Find(&units)
	err = db.Debug().Model(&Unit{}).Limit(100).Find(&units).Error
	if err != nil {
		return &[]Unit{}, err
	}
	return &units, err
}
func (u *Unit) FindUnitWithFilter(db *gorm.DB, q map[string][]string) (*[]Unit, error) {
	var err error
	units := []Unit{}

	db = db.Set("gorm:auto_preload", true)
	skill := false
	for k, v := range q {
		if k == "name" {
			db = db.Where("name ILIKE ?", "%"+v[0]+"%")
		}
		if k == "family" {
			db = db.Where("family_id = ?", v[0])
		}
		if k == "role" {
			db = db.Where("role_id = ?", v[0])
		}
		if k == "rarity" {
			db = db.Where("rarity_id = ?", v[0])
		}
		if strings.Contains(k, "skill_") {
			if !skill {
				db = db.Debug().Joins("left join unit_skills on units.id = unit_skills.unit_id").
					Joins("inner join skills on unit_skills.skill_id = skills.id")
				skill = true
			}
			if k == "skill_name" {
				db = db.Where("skills.name ILIKE ?", "%"+v[0]+"%")
			}
			if k == "skill_desc" {
				db = db.Where("skills.description ILIKE ?", "%"+v[0]+"%")
			}
		}
	}

	err = db.Debug().Model(&Unit{}).Limit(20).Find(&units).Error
	if err != nil {
		return &[]Unit{}, err
	}
	return &units, err
}
func (u *Unit) FindUnitById(db *gorm.DB, uid uint64) (*Unit, error) {

	db = db.Set("gorm:auto_preload", true)
	err := db.Where("id = ?", uid).Find(&u).Error
	if err != nil {
		return u, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return u, errors.New("user not found")
	}
	return u, err
}
