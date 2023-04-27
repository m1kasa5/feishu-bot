package service

import (
	"feishu-bot/service/openai"
	"github.com/patrickmn/go-cache"
	"time"
)

type SessionMode string
type SessionService struct {
	cache *cache.Cache
}
type PicSetting struct {
	resolution Resolution
}
type Resolution string
type SessionMeta struct {
	Mode       SessionMode       `json:"mode"`
	Msg        []openai.Messages `json:"msg,omitempty"`
	PicSetting PicSetting        `json:"pic_setting,omitempty"`
	AIMode     openai.AIMode     `json:"ai_mode,omitempty"`
}

const (
	Resolution256  Resolution = "256x256"
	Resolution512  Resolution = "512x512"
	Resolution1024 Resolution = "1024x1024"
)
const (
	ModePicCreate SessionMode = "pic_create"
	ModePicVary   SessionMode = "pic_vary"
	ModeGPT       SessionMode = "gpt"
)

type SessionServiceCacheInterface interface {
	Get(sessionId string) *SessionMeta
	Set(sessionId string, sessionMeta *SessionMeta)
	GetMsg(sessionId string) []openai.Messages
	SetMsg(sessionId string, msg []openai.Messages)
	SetMode(sessionId string, mode SessionMode)
	GetMode(sessionId string) SessionMode
	GetAIMode(sessionId string) openai.AIMode
	SetAIMode(sessionId string, aiMode openai.AIMode)
	SetPicResolution(sessionId string, resolution Resolution)
	GetPicResolution(sessionId string) string
	Clear(sessionId string)
}

var sessionServices *SessionService

func (s *SessionService) Get(sessionId string) *SessionMeta {
	sessonCtx, ok := s.cache.Get(sessionId)

	if !ok {
		return nil
	}
	sessionMeta := sessonCtx.(*SessionMeta)
	return sessionMeta
}

func (s *SessionService) Set(sessionId string, sessionMeta *SessionMeta) {
	maxCacheTime := 12 * time.Hour
	s.cache.Set(sessionId, sessionMeta, maxCacheTime)
}

func (s *SessionService) GetMsg(sessionId string) []openai.Messages {
	sessionCtx, ok := s.cache.Get(sessionId)
	if !ok {
		return nil
	}
	sessionMeta := sessionCtx.(*SessionMeta)
	return sessionMeta.Msg
}

func (s *SessionService) SetMsg(sessionId string, msg []openai.Messages) {
	//TODO implement me
	panic("implement me")
}

func (s *SessionService) SetMode(sessionId string, mode SessionMode) {
	//TODO implement me
	panic("implement me")
}

func (s *SessionService) GetMode(sessionId string) SessionMode {
	//TODO implement me
	panic("implement me")
}

func (s *SessionService) GetAIMode(sessionId string) openai.AIMode {
	//TODO implement me
	panic("implement me")
}

func (s *SessionService) SetAIMode(sessionId string, aiMode openai.AIMode) {
	//TODO implement me
	panic("implement me")
}

func (s *SessionService) SetPicResolution(sessionId string, resolution Resolution) {
	//TODO implement me
	panic("implement me")
}

func (s *SessionService) GetPicResolution(sessionId string) string {
	//TODO implement me
	panic("implement me")
}

func (s *SessionService) Clear(sessionId string) {
	//TODO implement me
	panic("implement me")
}
