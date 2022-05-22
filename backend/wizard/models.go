package wizard

import "github.com/rs/xid"

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
	BeforeGenerate string   `json:"before_generate,omitempty"` // fixme => allow only "determisnistic" kinds of bindings
	BeforeValidate string   `json:"before_validate,omitempty"`
	BeforeBack     string   `json:"before_back,omitempty"`
}

func (s *Stage) getField(name string) *Field {
	for _, f := range s.Fields {
		if f.Name == name {
			return f
		}
	}
	return nil
}

type Field struct {
	Name     string                 `json:"name,omitempty"`
	Info     string                 `json:"info,omitempty"`
	Type     string                 `json:"type,omitempty"`
	Pattern  string                 `json:"pattern,omitempty"`
	Optional bool                   `json:"optional,omitempty"`
	Attrs    map[string]interface{} `json:"attrs,omitempty"`
	Source   string                 `json:"source,omitempty"`
}

type Source struct {
	Name      string                 `json:"name,omitempty"`
	Type      string                 `json:"type,omitempty"`
	Target    string                 `json:"target,omitempty"`
	Data      interface{}            `json:"data,omitempty"`
	ExtraMeta map[string]interface{} `json:"extra_meta,omitempty"`

	// auto_cache ? => by_submission by_plug by_user by_stage_group by_stage by_field by_ctx_var_key
}

// this will be send as opaque data
type Submission struct {
	Id               string                            `json:"id,omitempty"`
	StageGroup       string                            `json:"stage_group,omitempty"`
	CurrentStage     string                            `json:"curr_stage,omitempty"`
	Data             map[string]map[string]interface{} `json:"data,omitempty"` // all stage data
	SharedVars       map[string]interface{}            `json:"shared_vars,omitempty"`
	ParentStageGroup string                            `json:"parent_stage_group,omitempty"` // incase of nested group
	ParentStage      string                            `json:"parent_stage,omitempty"`
	VisitedStages    []string                          `json:"visited_stages,omitempty"`
}

func newSub(group, stage string) Submission {
	return Submission{
		Id:               xid.New().String(),
		StageGroup:       group,
		CurrentStage:     stage,
		Data:             make(map[string]map[string]interface{}),
		SharedVars:       make(map[string]interface{}),
		ParentStageGroup: "",
		ParentStage:      "",
		VisitedStages:    []string{},
	}
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
	SplashData   map[string]interface{} `json:"splash_data,omitempty"`
	StartRawData interface{}            `json:"start_raw_data,omitempty"`
}

type ResponseStart struct {
	StartStage  bool                   `json:"stage_started,omitempty"`
	StageTitle  string                 `json:"stage_title,omitempty"`
	Message     string                 `json:"message,omitempty"`
	Fields      []*Field               `json:"fields,omitempty"`
	DataSources map[string]interface{} `json:"data_sources,omitempty"`
	PrevData    map[string]interface{} `json:"prev_data,omitempty"`
	OpaqueData  []byte                 `json:"odata,omitempty"`
	Ok          bool                   `json:"ok"`
}

type RequestStartNested struct {
	ParentOpaqueData []byte      `json:"parent_odata,omitempty"`
	Field            string      `json:"field,omitempty"`
	StartRawData     interface{} `json:"start_raw_data,omitempty"`
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
	PrevData    map[string]interface{} `json:"prev_data,omitempty"`
	OpaqueData  []byte                 `json:"odata,omitempty"`
	Ok          bool                   `json:"ok"`
	Final       bool                   `json:"final"`
	Errors      map[string]string      `json:"errors,omitempty"`
}

type ResponseFinal struct {
	LastMessage string      `json:"last_message,omitempty"`
	Ok          bool        `json:"ok"`
	Final       bool        `json:"final"`
	FinalData   interface{} `json:"final_data"`
}

type RequestBack struct {
	OpaqueData []byte `json:"odata,omitempty"`
}

type ResponseBack struct {
	StageTitle  string                 `json:"stage_title,omitempty"`
	Message     string                 `json:"message,omitempty"`
	Fields      []*Field               `json:"fields,omitempty"`
	DataSources map[string]interface{} `json:"data_sources,omitempty"`
	PrevData    map[string]interface{} `json:"prev_data,omitempty"`
	OpaqueData  []byte                 `json:"odata,omitempty"`
	Ok          bool                   `json:"ok"`
	Final       bool                   `json:"final"`
}

// elements types
const (
	BASIC_SHORTTEXT    = "basic.shorttext"
	BASIC_LONGTEXT     = "basic.longtext"
	BASIC_RANGE        = "basic.range"
	BASIC_SELECT       = "basic.select"
	BASIC_MULTI_SELECT = "basic.multiselect"
	BASIC_PHONE        = "basic.phone"
	BASIC_CHECKBOX     = "basic.checkbox"
	BASIC_COLOR        = "basic.color"
	BASIC_DATE         = "basic.date"
	BASIC_DATETIME     = "basic.datetime"
	BASIC_EMAIL        = "basic.email"
	BASIC_NUMBER       = "basic.number"
	BASIC_PARAGRAPH    = "basic.paragraph"

	SELECT_MONTH  = "select.month"
	SELECT_WEEK   = "select.week"
	SELECT_NESTED = "select.nested"

	IMAGE            = "imaget"
	IMAGE_INLINE     = "image.inline"
	FILE             = "file"
	FILE_INLINE      = "file.inline"
	MARKDOWN         = "markdown"
	MARKDOWN_PREVIEW = "markdown.preview"
	SECRET           = "secret"
	QUESTION         = "question"
	FULLNAME         = "fullname"

	JSON_MULTI_SELECT  = "json.select"
	JSON_MULTI_INLINE  = "json.inline"
	JSON_MULTI_NESTED  = "json.nested"
	JSON_SINGLE_SELECT = "json.select"
	JSON_SINGLE_INLINE = "json.inline"
	JSON_SINGLE_NESTED = "json.nested"

	LOCAT         = "locat"
	LOCAT_CIRCLE  = "locat.circle"
	LOCAT_AREA    = "locat.area"
	LOCAT_ADDRESS = "locat.addr" // https://github.com/kelvins/geocoder
	HTML          = "html"

	VIEW_IMAGE         = "view.image"
	VIEW_FILE          = "view.file"
	VIEW_CARDS         = "view.cards"
	VIEW_ALBUM         = "view.album"
	VIEW_CHARTJS       = "view.chartjs"
	VIEW_AUTOTABLE     = "view.autotable"
	VIEW_METRICS_CARD  = "view.metrics_card"
	VIEW_METRICS_TABLE = "view.metrics_table"
)
