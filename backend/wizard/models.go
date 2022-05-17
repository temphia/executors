package wizard

type Wizard struct {
	Title       string             `json:"title,omitempty"`
	Splash      Splash             `json:"splash,omitempty"`
	Dev         bool               `json:"dev,omitempty"`
	StageGroups []StageGroup       `json:"stage_groups,omitempty"`
	Stages      map[string]*Stage  `json:"stages,omitempty"`
	Sources     map[string]*Source `json:"sources,omitempty"`
}

type StageGroup struct {
	Name        string   `json:"name,omitempty"`
	Stages      []string `json:"stages,omitempty"`
	LastMessage string   `json:"last_message,omitempty"`
	PreventBack bool     `json:"prevent_back,omitempty"`
	OnNext      string   `json:"on_next,omitempty"`
	BeforeStart string   `json:"before_start,omitempty"`
	BeforeEnd   string   `json:"before_end,omitempty"`
	AfterEnd    string   `json:"after_end,omitempty"`
}

type Splash struct {
	Message        string   `json:"message,omitempty"`
	Fields         []*Field `json:"fields,omitempty"`
	Skip           bool     `json:"skip,omitempty"`
	BeforeGenerate string   `json:"before_generate,omitempty"`
	BeforeValidate string   `json:"before_validate,omitempty"`
}

type Stage struct {
	Name           string   `json:"name,omitempty"`
	Message        string   `json:"message,omitempty"`
	Fields         []*Field `json:"fields,omitempty"`
	BeforeGenerate string   `json:"before_generate,omitempty"`
	BeforeValidate string   `json:"before_validate,omitempty"`
}

type Field struct {
	Name     string                 `json:"name,omitempty"`
	Info     string                 `json:"info,omitempty"`
	Type     string                 `json:"type,omitempty"`
	Pattern  string                 `json:"pattern,omitempty"`
	Optional bool                   `json:"optional,omitempty"`
	Options  interface{}            `json:"options,omitempty"`
	Attrs    map[string]interface{} `json:"attrs,omitempty"`
	Source   string                 `json:"source,omitempty"`
}

type Source struct {
	Name      string                 `json:"name,omitempty"`
	Type      string                 `json:"type,omitempty"`
	Value     string                 `json:"value,omitempty"`
	Data      interface{}            `json:"data,omitempty"`
	ExtraMeta map[string]interface{} `json:"extra_meta,omitempty"`

	// auto_cache ? => by_submission by_plug by_user by_stage_group by_stage by_field by_ctx_var_key
}

// this will be send as opaque data
type Submission struct {
	Id               string                            `json:"id,omitempty"`
	StageGroup       string                            `json:"stage_group,omitempty"`
	CurrentStage     string                            `json:"curr_stage,omitempty"`
	Data             map[string]map[string]interface{} `json:"data,omitempty"`
	SharedVars       map[string]interface{}            `json:"shared_vars,omitempty"`
	ParentStageGroup string                            `json:"parent_stage_group,omitempty"` // incase of nested group
	ParentStage      string                            `json:"parent_stage,omitempty"`
	PrevStages       []string                          `json:"prev_stages,omitempty"`
}

type RequestSplash struct {
	HasExecData bool `json:"has_exec_data,omitempty"`
}
type ResponseSplash struct {
	WizardTitle string                 `json:"wizard_title,omitempty"`
	Message     string                 `json:"message,omitempty"`
	Fields      []*Field               `json:"fields,omitempty"`
	DataSources map[string]interface{} `json:"data_sources,omitempty"`
	SkipSplash  bool                   `json:"skip_splash,omitempty"`
}

type RequestStart struct {
	Data        map[string]interface{} `json:"data,omitempty"`
	ExecOptions interface{}            `json:"exec_options,omitempty"`
}

type ResponseStart struct {
	StartStage  bool                   `json:"stage_started,omitempty"`
	StageTitle  string                 `json:"stage_title,omitempty"`
	Message     string                 `json:"message,omitempty"`
	Fields      []*Field               `json:"fields,omitempty"`
	DataSources map[string]interface{} `json:"data_sources,omitempty"`
	OpaqueData  []byte                 `json:"odata,omitempty"`
	Ok          bool                   `json:"ok"`
}

type RequestNext struct {
	Data       map[string]interface{} `json:"data,omitempty"`
	OpaqueData []byte                 `json:"odata,omitempty"`
}

type ResponseNext struct {
	StageTitle  string                 `json:"stage_title,omitempty"`
	Message     string                 `json:"message,omitempty"`
	Fields      []*Field               `json:"fields,omitempty"`
	DataSources map[string]interface{} `json:"data_sources,omitempty"`
	OpaqueData  []byte                 `json:"odata,omitempty"`
	Ok          bool                   `json:"ok"`
	Final       bool                   `json:"final"`
	Errors      map[string]string      `json:"errors,omitempty"`
}

type ResponseFinal struct {
	LastMessage string `json:"last_message,omitempty"`
	Ok          bool   `json:"ok"`
	Final       bool   `json:"final"`
}
