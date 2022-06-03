package main

// got the data by using the website https://mholt.github.io/json-to-go/ and pasting the json the binary spits out

type ItermPlist struct {
	AppleAntiAliasingThreshold         int            `plist:"AppleAntiAliasingThreshold"`
	ApplePressAndHoldEnabled           bool           `plist:"ApplePressAndHoldEnabled"`
	AppleScrollAnimationEnabled        int            `plist:"AppleScrollAnimationEnabled"`
	AppleSmoothFixedFontsSizeThreshold int            `plist:"AppleSmoothFixedFontsSizeThreshold"`
	AppleWindowTabbingMode             string         `plist:"AppleWindowTabbingMode"`
	DefaultBookmarkGUID                string         `plist:"Default Bookmark Guid"`
	HapticFeedbackForEsc               bool           `plist:"HapticFeedbackForEsc"`
	HotkeyMigratedFromSingleToMulti    bool           `plist:"HotkeyMigratedFromSingleToMulti"`
	NewBookmarks                       []NewBookmarks `plist:"New Bookmarks"`
	SoundForEsc                        bool           `plist:"SoundForEsc"`
	VisualIndicatorForEsc              bool           `plist:"VisualIndicatorForEsc"`
}
type Ansi0Color struct {
	BlueComponent  int `plist:"Blue Component"`
	GreenComponent int `plist:"Green Component"`
	RedComponent   int `plist:"Red Component"`
}
type Ansi1Color struct {
	BlueComponent  int     `plist:"Blue Component"`
	GreenComponent int     `plist:"Green Component"`
	RedComponent   float64 `plist:"Red Component"`
}
type Ansi10Color struct {
	BlueComponent  float64 `plist:"Blue Component"`
	GreenComponent int     `plist:"Green Component"`
	RedComponent   float64 `plist:"Red Component"`
}
type Ansi11Color struct {
	BlueComponent  float64 `plist:"Blue Component"`
	GreenComponent int     `plist:"Green Component"`
	RedComponent   int     `plist:"Red Component"`
}
type Ansi12Color struct {
	BlueComponent  int     `plist:"Blue Component"`
	GreenComponent float64 `plist:"Green Component"`
	RedComponent   float64 `plist:"Red Component"`
}
type Ansi13Color struct {
	BlueComponent  int     `plist:"Blue Component"`
	GreenComponent float64 `plist:"Green Component"`
	RedComponent   int     `plist:"Red Component"`
}
type Ansi14Color struct {
	BlueComponent  int     `plist:"Blue Component"`
	GreenComponent int     `plist:"Green Component"`
	RedComponent   float64 `plist:"Red Component"`
}
type Ansi15Color struct {
	BlueComponent  int `plist:"Blue Component"`
	GreenComponent int `plist:"Green Component"`
	RedComponent   int `plist:"Red Component"`
}
type Ansi2Color struct {
	BlueComponent  int     `plist:"Blue Component"`
	GreenComponent float64 `plist:"Green Component"`
	RedComponent   int     `plist:"Red Component"`
}
type Ansi3Color struct {
	BlueComponent  int     `plist:"Blue Component"`
	GreenComponent float64 `plist:"Green Component"`
	RedComponent   float64 `plist:"Red Component"`
}
type Ansi4Color struct {
	BlueComponent  float64 `plist:"Blue Component"`
	GreenComponent int     `plist:"Green Component"`
	RedComponent   int     `plist:"Red Component"`
}
type Ansi5Color struct {
	BlueComponent  float64 `plist:"Blue Component"`
	GreenComponent int     `plist:"Green Component"`
	RedComponent   float64 `plist:"Red Component"`
}
type Ansi6Color struct {
	BlueComponent  float64 `plist:"Blue Component"`
	GreenComponent float64 `plist:"Green Component"`
	RedComponent   int     `plist:"Red Component"`
}
type Ansi7Color struct {
	BlueComponent  float64 `plist:"Blue Component"`
	GreenComponent float64 `plist:"Green Component"`
	RedComponent   float64 `plist:"Red Component"`
}
type Ansi8Color struct {
	BlueComponent  float64 `plist:"Blue Component"`
	GreenComponent float64 `plist:"Green Component"`
	RedComponent   float64 `plist:"Red Component"`
}
type Ansi9Color struct {
	BlueComponent  float64 `plist:"Blue Component"`
	GreenComponent float64 `plist:"Green Component"`
	RedComponent   int     `plist:"Red Component"`
}
type BackgroundColor struct {
	BlueComponent  int `plist:"Blue Component"`
	GreenComponent int `plist:"Green Component"`
	RedComponent   int `plist:"Red Component"`
}
type BoldColor struct {
	BlueComponent  int `plist:"Blue Component"`
	GreenComponent int `plist:"Green Component"`
	RedComponent   int `plist:"Red Component"`
}
type CursorColor struct {
	BlueComponent  float64 `plist:"Blue Component"`
	GreenComponent float64 `plist:"Green Component"`
	RedComponent   float64 `plist:"Red Component"`
}
type CursorTextColor struct {
	BlueComponent  int `plist:"Blue Component"`
	GreenComponent int `plist:"Green Component"`
	RedComponent   int `plist:"Red Component"`
}
type ForegroundColor struct {
	BlueComponent  float64 `plist:"Blue Component"`
	GreenComponent float64 `plist:"Green Component"`
	RedComponent   float64 `plist:"Red Component"`
}
type ZeroX2D0X40000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroX320X40000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroX330X40000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroX340X40000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroX350X40000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroX360X40000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroX370X40000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroX380X40000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf7000X220000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf7000X240000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf7000X260000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf7000X280000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf7010X220000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf7010X240000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf7010X260000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf7010X280000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf7020X220000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf7020X240000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf7020X260000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf7020X280000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf7030X220000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf7030X240000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf7030X260000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf7030X280000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf7040X20000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf7050X20000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf7060X20000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf7070X20000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf7080X20000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf7090X20000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf70A0X20000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf70B0X20000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf70C0X20000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf70D0X20000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf70E0X20000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf70F0X20000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf7290X20000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf7290X40000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf72B0X20000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type ZeroXf72B0X40000 struct {
	Action int    `plist:"Action"`
	Text   string `plist:"Text"`
}
type KeyboardMap struct {
	ZeroX2D0X40000    ZeroX2D0X40000    `plist:"0x2d-0x40000"`
	ZeroX320X40000    ZeroX320X40000    `plist:"0x32-0x40000"`
	ZeroX330X40000    ZeroX330X40000    `plist:"0x33-0x40000"`
	ZeroX340X40000    ZeroX340X40000    `plist:"0x34-0x40000"`
	ZeroX350X40000    ZeroX350X40000    `plist:"0x35-0x40000"`
	ZeroX360X40000    ZeroX360X40000    `plist:"0x36-0x40000"`
	ZeroX370X40000    ZeroX370X40000    `plist:"0x37-0x40000"`
	ZeroX380X40000    ZeroX380X40000    `plist:"0x38-0x40000"`
	ZeroXf7000X220000 ZeroXf7000X220000 `plist:"0xf700-0x220000"`
	ZeroXf7000X240000 ZeroXf7000X240000 `plist:"0xf700-0x240000"`
	ZeroXf7000X260000 ZeroXf7000X260000 `plist:"0xf700-0x260000"`
	ZeroXf7000X280000 ZeroXf7000X280000 `plist:"0xf700-0x280000"`
	ZeroXf7010X220000 ZeroXf7010X220000 `plist:"0xf701-0x220000"`
	ZeroXf7010X240000 ZeroXf7010X240000 `plist:"0xf701-0x240000"`
	ZeroXf7010X260000 ZeroXf7010X260000 `plist:"0xf701-0x260000"`
	ZeroXf7010X280000 ZeroXf7010X280000 `plist:"0xf701-0x280000"`
	ZeroXf7020X220000 ZeroXf7020X220000 `plist:"0xf702-0x220000"`
	ZeroXf7020X240000 ZeroXf7020X240000 `plist:"0xf702-0x240000"`
	ZeroXf7020X260000 ZeroXf7020X260000 `plist:"0xf702-0x260000"`
	ZeroXf7020X280000 ZeroXf7020X280000 `plist:"0xf702-0x280000"`
	ZeroXf7030X220000 ZeroXf7030X220000 `plist:"0xf703-0x220000"`
	ZeroXf7030X240000 ZeroXf7030X240000 `plist:"0xf703-0x240000"`
	ZeroXf7030X260000 ZeroXf7030X260000 `plist:"0xf703-0x260000"`
	ZeroXf7030X280000 ZeroXf7030X280000 `plist:"0xf703-0x280000"`
	ZeroXf7040X20000  ZeroXf7040X20000  `plist:"0xf704-0x20000"`
	ZeroXf7050X20000  ZeroXf7050X20000  `plist:"0xf705-0x20000"`
	ZeroXf7060X20000  ZeroXf7060X20000  `plist:"0xf706-0x20000"`
	ZeroXf7070X20000  ZeroXf7070X20000  `plist:"0xf707-0x20000"`
	ZeroXf7080X20000  ZeroXf7080X20000  `plist:"0xf708-0x20000"`
	ZeroXf7090X20000  ZeroXf7090X20000  `plist:"0xf709-0x20000"`
	ZeroXf70A0X20000  ZeroXf70A0X20000  `plist:"0xf70a-0x20000"`
	ZeroXf70B0X20000  ZeroXf70B0X20000  `plist:"0xf70b-0x20000"`
	ZeroXf70C0X20000  ZeroXf70C0X20000  `plist:"0xf70c-0x20000"`
	ZeroXf70D0X20000  ZeroXf70D0X20000  `plist:"0xf70d-0x20000"`
	ZeroXf70E0X20000  ZeroXf70E0X20000  `plist:"0xf70e-0x20000"`
	ZeroXf70F0X20000  ZeroXf70F0X20000  `plist:"0xf70f-0x20000"`
	ZeroXf7290X20000  ZeroXf7290X20000  `plist:"0xf729-0x20000"`
	ZeroXf7290X40000  ZeroXf7290X40000  `plist:"0xf729-0x40000"`
	ZeroXf72B0X20000  ZeroXf72B0X20000  `plist:"0xf72b-0x20000"`
	ZeroXf72B0X40000  ZeroXf72B0X40000  `plist:"0xf72b-0x40000"`
}
type SelectedTextColor struct {
	BlueComponent  int `plist:"Blue Component"`
	GreenComponent int `plist:"Green Component"`
	RedComponent   int `plist:"Red Component"`
}
type SelectionColor struct {
	BlueComponent  int     `plist:"Blue Component"`
	GreenComponent float64 `plist:"Green Component"`
	RedComponent   float64 `plist:"Red Component"`
}
type NewBookmarks struct {
	ASCIIAntiAliased        bool              `plist:"ASCII Anti Aliased"`
	AmbiguousDoubleWidth    bool              `plist:"Ambiguous Double Width"`
	Ansi0Color              Ansi0Color        `plist:"Ansi 0 Color"`
	Ansi1Color              Ansi1Color        `plist:"Ansi 1 Color"`
	Ansi10Color             Ansi10Color       `plist:"Ansi 10 Color"`
	Ansi11Color             Ansi11Color       `plist:"Ansi 11 Color"`
	Ansi12Color             Ansi12Color       `plist:"Ansi 12 Color"`
	Ansi13Color             Ansi13Color       `plist:"Ansi 13 Color"`
	Ansi14Color             Ansi14Color       `plist:"Ansi 14 Color"`
	Ansi15Color             Ansi15Color       `plist:"Ansi 15 Color"`
	Ansi2Color              Ansi2Color        `plist:"Ansi 2 Color"`
	Ansi3Color              Ansi3Color        `plist:"Ansi 3 Color"`
	Ansi4Color              Ansi4Color        `plist:"Ansi 4 Color"`
	Ansi5Color              Ansi5Color        `plist:"Ansi 5 Color"`
	Ansi6Color              Ansi6Color        `plist:"Ansi 6 Color"`
	Ansi7Color              Ansi7Color        `plist:"Ansi 7 Color"`
	Ansi8Color              Ansi8Color        `plist:"Ansi 8 Color"`
	Ansi9Color              Ansi9Color        `plist:"Ansi 9 Color"`
	BMGrowl                 bool              `plist:"BM Growl"`
	BackgroundColor         BackgroundColor   `plist:"Background Color"`
	BackgroundImageLocation string            `plist:"Background Image Location"`
	BlinkingCursor          bool              `plist:"Blinking Cursor"`
	Blur                    bool              `plist:"Blur"`
	BoldColor               BoldColor         `plist:"Bold Color"`
	CharacterEncoding       int               `plist:"Character Encoding"`
	CloseSessionsOnEnd      bool              `plist:"Close Sessions On End"`
	Columns                 int               `plist:"Columns"`
	Command                 string            `plist:"Command"`
	CursorColor             CursorColor       `plist:"Cursor Color"`
	CursorTextColor         CursorTextColor   `plist:"Cursor Text Color"`
	CustomCommand           string            `plist:"Custom Command"`
	CustomDirectory         string            `plist:"Custom Directory"`
	DefaultBookmark         string            `plist:"Default Bookmark"`
	Description             string            `plist:"Description"`
	DisableWindowResizing   bool              `plist:"Disable Window Resizing"`
	FlashingBell            bool              `plist:"Flashing Bell"`
	ForegroundColor         ForegroundColor   `plist:"Foreground Color"`
	GUID                    string            `plist:"Guid"`
	HorizontalSpacing       int               `plist:"Horizontal Spacing"`
	IdleCode                int               `plist:"Idle Code"`
	JobsToIgnore            []string          `plist:"Jobs to Ignore"`
	KeyboardMap             KeyboardMap       `plist:"Keyboard Map"`
	MouseReporting          bool              `plist:"Mouse Reporting"`
	Name                    string            `plist:"Name"`
	NonASCIIFont            string            `plist:"Non Ascii Font"`
	NonASCIIAntiAliased     bool              `plist:"Non-ASCII Anti Aliased"`
	NormalFont              string            `plist:"Normal Font"`
	OptionKeySends          int               `plist:"Option Key Sends"`
	PromptBeforeClosing2    bool              `plist:"Prompt Before Closing 2"`
	RightOptionKeySends     int               `plist:"Right Option Key Sends"`
	Rows                    int               `plist:"Rows"`
	Screen                  int               `plist:"Screen"`
	ScrollbackLines         int               `plist:"Scrollback Lines"`
	SelectedTextColor       SelectedTextColor `plist:"Selected Text Color"`
	SelectionColor          SelectionColor    `plist:"Selection Color"`
	SendCodeWhenIdle        bool              `plist:"Send Code When Idle"`
	Shortcut                string            `plist:"Shortcut"`
	SilenceBell             bool              `plist:"Silence Bell"`
	SyncTitle               bool              `plist:"Sync Title"`
	Tags                    []interface{}     `plist:"Tags"`
	TerminalType            string            `plist:"Terminal Type"`
	Transparency            int               `plist:"Transparency"`
	UnlimitedScrollback     bool              `plist:"Unlimited Scrollback"`
	UseBoldFont             bool              `plist:"Use Bold Font"`
	UseBrightBold           bool              `plist:"Use Bright Bold"`
	UseItalicFont           bool              `plist:"Use Italic Font"`
	UseNonASCIIFont         bool              `plist:"Use Non-ASCII Font"`
	VerticalSpacing         int               `plist:"Vertical Spacing"`
	VisualBell              bool              `plist:"Visual Bell"`
	WindowType              int               `plist:"Window Type"`
	WorkingDirectory        string            `plist:"Working Directory"`
}
