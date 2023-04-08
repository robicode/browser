package browser

import (
	"regexp"
	"strings"

	"github.com/robicode/browser/device"
)

type Device struct {
	id        string
	name      string
	subject   DeviceMatcher
	matchers  []DeviceMatcher
	userAgent string
}

type DeviceMatcher interface {
	ID() string
	Name() string
	Matches() bool
}

func NewDevice(dataPath, userAgent string) *Device {
	var (
		id      string
		name    string
		subject DeviceMatcher
	)
	samsung := device.NewSamsung(userAgent, dataPath)
	xboxOne := device.NewXboxOne(userAgent)
	xbox360 := device.NewXbox360(userAgent)
	surface := device.NewSurface(userAgent)
	tv := device.NewTV(userAgent)
	playbook := device.NewBlackberryPlaybook(userAgent)
	wiiU := device.NewWiiU(userAgent)
	wii := device.NewWii(userAgent)
	nsw := device.NewSwitch(userAgent)
	kf := device.NewKindleFire(userAgent)
	kindle := device.NewKindle(userAgent)
	ps5 := device.NewPlaystation5(userAgent)
	ps4 := device.NewPlaystation4(userAgent)
	ps3 := device.NewPlaystation3(userAgent)
	vita := device.NewPSVita(userAgent)
	psp := device.NewPSP(userAgent)
	ipad := device.NewIPad(userAgent)
	iphone := device.NewIPhone(userAgent)
	ipod := device.NewIPodTouch(userAgent)
	android := device.NewAndroid(userAgent)
	unknown := device.NewUnknown(userAgent)

	defaultMatchers := []DeviceMatcher{
		samsung,
		xboxOne,
		xbox360,
		surface,
		tv,
		playbook,
		wiiU,
		wii,
		nsw,
		kf,
		kindle,
		ps5,
		ps4,
		ps3,
		vita,
		psp,
		ipad,
		iphone,
		ipod,
		android,
		unknown,
	}

	for _, matcher := range defaultMatchers {
		if matcher.Matches() {
			id = matcher.ID()
			name = matcher.Name()
			subject = matcher
		}
	}

	return &Device{
		id:        id,
		name:      name,
		subject:   subject,
		matchers:  defaultMatchers,
		userAgent: userAgent,
	}
}

func (d *Device) ID() string {
	return d.id
}

func (d *Device) Name() string {
	return d.name
}

func (d *Device) IsUnknown() bool {
	return d.id == "unknown_device"
}

func (d *Device) IsBlackberryPlaybook() bool {
	return d.id == "playbook"
}

func (d *Device) IsConsole() bool {
	return d.IsXbox() || d.IsPlaystation() || d.IsNintendo()
}

func (d *Device) IsIpad() bool {
	return d.id == "ipad"
}

func (d *Device) IsIphone() bool {
	return d.id == "iphone"
}

func (d *Device) IsIpodTouch() bool {
	return d.id == "ipod_touch"
}

func (d *Device) IsIpod() bool {
	return d.IsIpodTouch()
}

func (d *Device) IsKindle() bool {
	return d.id == "kindle" || d.IsKindleFire()
}

func (d *Device) IsKindleFire() bool {
	return d.id == "kindle_fire"
}

func (d *Device) IsMobile() bool {
	return d.detectMobile() && !d.IsTablet()
}

func (d *Device) IsNintendo() bool {
	return d.IsWii() || d.IsWiiU() || d.IsWiiU()
}

func (d *Device) IsPlaystation() bool {
	return d.IsPS5() || d.IsPS4() || d.IsPS3()
}

func (d *Device) IsPS3() bool {
	return d.id == "ps3"
}

func (d *Device) IsPS4() bool {
	return d.id == "ps4"
}

func (d *Device) IsPS5() bool {
	return d.id == "ps5"
}

func (d *Device) IsPSP() bool {
	return d.id == "psp"
}

func (d *Device) IsSilk() bool {
	return strings.Contains(d.userAgent, "Silk")
}

func (d *Device) IsSurface() bool {
	return d.id == "surface"
}

func (d *Device) IsTablet() bool {
	return false
}

func (d *Device) IsTV() bool {
	return d.id == "tv"
}

func (d *Device) IsVita() bool {
	return d.id == "psvita"
}

func (d *Device) IsPlaystationVita() bool {
	return d.IsVita()
}

func (d *Device) IsPSPVita() bool {
	return d.IsVita()
}

func (d *Device) IsPSVita() bool {
	return d.IsVita()
}

func (d *Device) IsWii() bool {
	return d.id == "wii"
}

func (d *Device) IsWiiU() bool {
	return d.id == "wiiu"
}

func (d *Device) IsSamsung() bool {
	return d.id == "samsung"
}

func (d *Device) IsSwitch() bool {
	return d.id == "switch"
}

func (d *Device) IsXbox() bool {
	return strings.Contains(d.userAgent, "Xbox")
}

func (d *Device) IsXbox360() bool {
	return d.id == "xbox_360"
}

func (d *Device) IsXboxOne() bool {
	return d.id == "xbox_one"
}

func (d *Device) detectMobile() bool {
	if d.IsPSP() {
		return true
	}

	zune, _ := regexp.MatchString(strings.ToLower(`zunewp7`), strings.ToLower(d.userAgent))
	if zune {
		re := regexp.MustCompile(strings.ToLower(`(android|bb\d+|meego).+mobile|avantgo|bada/|blackberry|blazer|compal|elaine|fennec|hiptop|iemobile|ip(hone|od)|iris|kindle|lge |maemo|midp|mmp|mobile.+firefox|netfront|opera m(ob|in)i|palm( os)?|phone|p(ixi|re)/|plucker|pocket|psp|series(4|6)0|symbian|treo|up\.(browser|link)|vodafone|wap|windows ce|xda|xiino`))
		re2 := regexp.MustCompile(strings.ToLower(`{1207|6310|6590|3gso|4thp|50[1-6]i|770s|802s|a wa|abac|ac(er|oo|s-)|ai(ko|rn)|al(av|ca|co)|amoi|an(ex|ny|yw)|aptu|ar(ch|go)|as(te|us)|attw|au(di|-m|r |s )|avan|be(ck|ll|nq)|bi(lb|rd)|bl(ac|az)|br(e|v)w|bumb|bw-(n|u)|c55/|capi|ccwa|cdm-|cell|chtm|cldc|cmd-|co(mp|nd)|craw|da(it|ll|ng)|dbte|dc-s|devi|dica|dmob|do(c|p)o|ds(12|-d)|el(49|ai)|em(l2|ul)|er(ic|k0)|esl8|ez([4-7]0|os|wa|ze)|fetc|fly(-|_)|g1 u|g560|gene|gf-5|g-mo|go(\.w|od)|gr(ad|un)|haie|hcit|hd-(m|p|t)|hei-|hi(pt|ta)|hp( i|ip)|hs-c|ht(c(-| |_|a|g|p|s|t)|tp)|hu(aw|tc)|i-(20|go|ma)|i230|iac( |-|/)|ibro|idea|ig01|ikom|im1k|inno|ipaq|iris|ja(t|v)a|jbro|jemu|jigs|kddi|keji|kgt( |/)|klon|kpt |kwc-|kyo(c|k)|le(no|xi)|lg( g|/(k|l|u)|50|54|-[a-w])|libw|lynx|m1-w|m3ga|m50/|ma(te|ui|xo)|mc(01|21|ca)|m-cr|me(rc|ri)|mi(o8|oa|ts)|mmef|mo(01|02|bi|de|do|t(-| |o|v)|zz)|mt(50|p1|v )|mwbp|mywa|n10[0-2]|n20[2-3]|n30(0|2)|n50(0|2|5)|n7(0(0|1)|10)|ne((c|m)-|on|tf|wf|wg|wt)|nok(6|i)|nzph|o2im|op(ti|wv)|oran|owg1|p800|pan(a|d|t)|pdxg|pg(13|-([1-8]|c))|phil|pire|pl(ay|uc)|pn-2|po(ck|rt|se)|prox|psio|pt-g|qa-a|qc(07|12|21|32|60|-[2-7]|i-)|qtek|r380|r600|raks|rim9|ro(ve|zo)|s55/|sa(ge|ma|mm|ms|ny|va)|sc(01|h-|oo|p-)|sdk/|se(c(-|0|1)|47|mc|nd|ri)|sgh-|shar|sie(-|m)|sk-0|sl(45|id)|sm(al|ar|b3|it|t5)|so(ft|ny)|sp(01|h-|v-|v )|sy(01|mb)|t2(18|50)|t6(00|10|18)|ta(gt|lk)|tcl-|tdg-|tel(i|m)|tim-|t-mo|to(pl|sh)|ts(70|m-|m3|m5)|tx-9|up(\.b|g1|si)|utst|v400|v750|veri|vi(rg|te)|vk(40|5[0-3]|-v)|vm40|voda|vulc|vx(52|53|60|61|70|80|81|83|85|98)|w3c(-| )|webc|whit|wi(g |nc|nw)|wmlb|wonu|x700|yas-|your|zeto|zte-`))

		if re.MatchString(strings.ToLower(d.userAgent)) || re2.MatchString(strings.ToLower(d.userAgent[0:3])) {
			return true
		}
	}

	return false
}
