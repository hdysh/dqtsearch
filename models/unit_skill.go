package models

type UnitSkill struct {
	UnitID  int64 `json:"-"`
	SkillID int64 `json:"-"`
	Skill   Skill `json:"skill"`
	Level   int   `json:"level"`
}
