package context

type Context struct {
	traceId string //traceID 请求的上下文
	Scene   string //场景 ocr wr correct ifpr 等
	Ability string //引擎能力 math_application en_filling_blank等
	//BaseParam         *api_entity.BaseRequestParam //base中的参数
	//EngineConfig      *engine_entity.EngineConfig  //当前请求的redisConfig
	EngineBestHashKey string //当前使用的RedisHashKey
	EngineBestIp      string //引擎的IP 如果获取到了就设置进去
	stopFlag          bool   //关闭标准
}

func NewContext(traceId string) *Context {
	return &Context{
		traceId: traceId,
	}
}

func NewWholeContext(traceId, scene, ability string) *Context {
	return &Context{
		traceId: traceId,
		Scene:   scene,
		Ability: ability,
	}
}

func (c *Context) GetTraceId() string {
	return c.traceId
}

func (c *Context) SetScene(scene string) *Context {
	c.Scene = scene
	return c
}

func (c *Context) SetAbility(ability string) *Context {
	c.Ability = ability
	return c
}

func (c *Context) SetStopFlag(stopFlag bool) {
	c.stopFlag = stopFlag
}

func (c *Context) GetStopFlag() bool {
	return c.stopFlag
}
