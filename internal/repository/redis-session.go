package repository

import (
	"mini-api/internal/api/constant"
	"mini-api/internal/api/dto"
)

func (h Handler) SetStaffSession(session dto.Session) error {
	key := GetKeyStaffSession(session.StaffID)
	return h.redis.JSONSet(key, session, constant.TTLJWTExpires)
}

func (h Handler) GetStaffSession(staffId uint32) (dto.Session, error) {
	key := GetKeyStaffSession(staffId)
	session := dto.Session{}
	err := h.redis.JSONGet(key, &session)
	return session, err
}

func (h Handler) DelStaffSession(staffId uint32) {
	h.Del(GetKeyStaffSession(staffId))
}
