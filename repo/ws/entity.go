package main

type WsResponse struct {
	Action string `json:"action"`
	Resp   string `json:"resp"`
}

type WsMessage struct {
	Action       string                 `json:"action"`
	Uuid         string                 `json:"uuid"`
	Status       string                 `json:"status"`
	Data         map[string]interface{} `json:"data"`
	SendDateTime string                 `json:"sendDate"`
}

type WriteBody struct {
	Action    string      `json:"action"`
	BizCode   string      `json:"bizCode"`
	DeviceNo  string      `json:"deviceNo"`
	Source    string      `json:"source"`
	Timestamp string      `json:"timestamp"`
	ExtraMap  interface{} `json:"extraMap,omitempty"`
	EndSign   string      `json:"endSign,omitempty"`
	Digest    string      `json:"digest"`
	Data      string      `json:"data"`
}

type WriteBodyData struct {
	AudioBase64 string `json:"audioBase64,omitempty"`
	AudioBytes  []byte `json:"audioBytes,omitempty"`
	AudioSegIdx int64  `json:"audioSegIdx,omitempty"`
	AudioText   string `json:"audioText,omitempty"`
}

// stream 需要的结构
type WriteBodyExtra struct {
	Encoding         string `json:"encoding,omitempty"`
	SampleRate       int    `json:"sampleRate,omitempty"`
	SampleBit        int    `json:"sampleBit,omitempty"`
	SampleChannels   int    `json:"sampleChannels,omitempty"`
	SampleMillSecond int    `json:"sampleMillSecond,omitempty"`
	TotalMillSecond  int    `json:"totalMillSecond,omitempty"`
	MemberId         string `json:"memberId,omitempty"`
	PageIndex        int    `json:"pageIndex,omitempty"`
	TotalPage        int    `json:"totalPage,omitempty"`
	VoiceFeatureId   string `json:"voiceFeatureId,omitempty"`
	CurrentChainId   string `json:"currentChainId,omitempty"`
	SpecialInstruct  string `json:"specialInstruct,omitempty"`
	ConversationId   string `json:"conversationId,omitempty"`
	WorkFlowType     string `json:"workFlowType,omitempty"`
	LocalAsrSign     bool   `json:"localAsrSign,omitempty"`
	ModelSwitch      string `json:"modelSwitch,omitempty"`
	AudioSegIdx      int64  `json:"audioSegIdx,omitempty"`
}

// cron ext
type CronExt struct {
	Cron   string                 `json:"cron"`   // * * * * * *
	Action string                 `json:"action"` // 具体方法名
	Params map[string]interface{} `json:"params"`
}

// wifi ext
type WifiExt struct {
	SSID     string `json:"ssid"`     // ssid
	Password string `json:"password"` // pwd
}

// bluetooth ext
type BluetoothExt struct {
}

// volume ext
type VolumeExt struct {
	Level int `json:"level"` // 音量大小
}

// mode ext
type ModeExt struct {
	Type    string `json:"type"`    // 已唤醒 、 待唤醒
	ChainId string `json:"chainId"` // 打断带上chainId
}

// system ext
type SystemExt struct {
}

// audio_stream ext
type AudioStreamExt struct {
	Speed string `json:"speed"`
	Audio string `json:"audio"`
}

// audio_file ext
type AudioFileExt struct {
	Ext   string `json:"ext"` // mp3 / wav
	Speed string `json:"speed"`
	Url   string `json:"url"`
}

// VoicePrintExt 声纹下发
type VoicePrintExt struct {
	Timestamp int64          `json:"timestamp"`
	List      []RemoteSample `json:"list"`
	Local     bool           `json:"local"`
}

type RemoteSample struct {
	RoleId    string                 `json:"roleId"`    // 角色
	UnknownId string                 `json:"unknownId"` // 陌生人临时ID
	Resources []RemoteSampleResource `json:"resources"` // 声纹地址
}

type RemoteSampleResource struct {
	Filename string `json:"filename"`
	Url      string `json:"url"`
}

type AsrModeExt struct {
	Type string `json:"type"`
}

type SceneExt struct {
	Type string `json:"type"`
}

// binding ext
type BindingExt struct {
	TimbreId string `json:"timbreId"`
}

// chat text ext
type KeywordExt struct {
	Type     string `json:"type"` // wake_word
	Content  string `json:"content"`
	TimbreId string `json:"timbreId"`
	Local    bool   `json:"local"`
}

// version ext
type VersionExt struct {
	Url          string `json:"url"`          // 最新的下载地址
	Version      string `json:"version"`      // 最新版本号
	ExpectedHash string `json:"expectedHash"` // 文件完整性hash
}

type SysParam struct {
	PrivateIp         string `json:"privateIp"`
	MacAddress        string `json:"macAddress"`
	FirmwareVersion   string `json:"firmwareVersion"`
	IsAutoUpgrade     int    `json:"isAutoUpgrade"`
	DeviceTemperature int    `json:"deviceTemperature"`
	DialogStatus      string `json:"dialogStatus"`
	LoseStatus        bool   `json:"loseStatus"`
	AsrMode           string `json:"asrMode"`
	Scene             string `json:"scene"`
}

type PowerParam struct {
	ElectricQuantity  int    `json:"electricQuantity"`
	IsCharging        bool   `json:"isCharging"`
	PowerSourceStatus string `json:"powerSourceStatus"`
}

type WifiParam struct {
	Ssid         string `json:"ssid"`
	Status       string `json:"status"`
	TotalUseFlow uint64 `json:"totalUseFlow"`
}

type SimParam struct {
	Status       string `json:"status"`
	SerialNumber string `json:"serialNumber"`
	TotalUseFlow uint64 `json:"totalUseFlow"`
}

type BluetoothParam struct {
	Status string `json:"status"`
	Name   string `json:"name"`
}

type VolumeParam struct {
	Level int `json:"level"`
}

type MicParam struct {
	Status string `json:"status"`
}

type CpuParam struct {
	ModelName string `json:"modelName"`
	Mhz       string `json:"mhz"`
	Cores     int    `json:"cores"`
}

type NetParam struct {
	Signal       int    `json:"signal"`
	SwitchStatus int    `json:"switchStatus"`
	SignalStatus string `json:"signalStatus"`
}

type MemParam struct {
	Total     uint64 `json:"total"`
	Available uint64 `json:"available"`
	Used      uint64 `json:"used"`
	Free      uint64 `json:"free"`
	UsageRate int    `json:"usageRate"`
}

type ModuleParam struct {
	Stt bool `json:"stt"`
}

// DeviceAudioExt 语音下发
type DeviceAudioExt struct {
	BeepTypeAudio  string  `json:"beepTypeAudio"`
	AudioBase64    string  `json:"audioBase64"`
	AudioText      string  `json:"audioText"`
	AudioFormat    string  `json:"audioFormat"`
	MemberId       string  `json:"memberId"`
	Speed          string  `json:"speed"`
	ChainId        string  `json:"chainId"`
	InputTokens    float64 `json:"inputTokens,omitempty"`
	OutputTokens   float64 `json:"outputTokens,omitempty"`
	AudioPriority  int     `json:"audioPriority"` //语音优先级，1-最高优，10-普通优先级（语音对话默认10）
	NeedInterrupt  int     `json:"needInterrupt"` //是否需要打断，1-打断，0-不打断  （语音对话默认1）
	NeedReceipt    int     `json:"needReceipt"`   //是否需要回执，1-回执，0-不回执（语音对话默认0）
	ConversationId string  `json:"conversationId"`
	SignStatus     string  `json:"signStatus"`
	AudioType      string  `json:"audioType"`  // 工作流程场景
	AfterEvent     string  `json:"afterEvent"` // 后置事件
}

type ModuleExt struct {
	Type    string `json:"type"` // asr
	Pattern string `json:"pattern"`
}

type ReportReceiptExt struct {
	ChainId string `json:"chainId"`
}
