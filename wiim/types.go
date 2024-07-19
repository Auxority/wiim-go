package wiim

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

type Channel int
type HexString string
type IsMuted bool
type LoopMode int
type PlaybackMode int
type Integer int
type Status string

type StatusResponse struct {
	Album             HexString    `json:"Album"`
	Artist            HexString    `json:"Artist"`
	Status            Status       `json:"status"`
	Title             HexString    `json:"Title"`
	Vendor            string       `json:"vendor"`
	Channel           Channel      `json:"ch"`
	EqPreset          Integer      `json:"eq"`
	IsMuted           IsMuted      `json:"mute"`
	LoopMode          LoopMode     `json:"loop"`
	PlaybackMode      PlaybackMode `json:"mode"`
	TotalTracks       Integer      `json:"plicount"`
	TrackDurationMs   Integer      `json:"totlen"`
	CurrentTrackIndex Integer      `json:"plicurrenttrackindex"`
	TrackProgressMs   Integer      `json:"curpos"`
	Volume            Integer      `json:"vol"`
}

const (
	Stereo Channel = iota
	Left
	Right
)

const (
	Playing Status = "play"
	Paused  Status = "pause"
	Stopped Status = "stop"
	Loading Status = "loading"
)

const (
	RepeatAll LoopMode = iota
	RepeatOne
	ShuffleRepeat
	ShuffleOnce
	Sequential
)

const (
	None            PlaybackMode = 0
	Airplay         PlaybackMode = 1
	ThirdPartyDLNA  PlaybackMode = 2
	Default         PlaybackMode = 10
	USBPlaylist     PlaybackMode = 11
	TFCardPlaylist  PlaybackMode = 16
	SpotifyConnect  PlaybackMode = 31
	TidalConnect    PlaybackMode = 32
	AuxIn           PlaybackMode = 40
	Bluetooth       PlaybackMode = 41
	ExternalStorage PlaybackMode = 42
	OpticalIn       PlaybackMode = 43
	Mirror          PlaybackMode = 44
	VoiceMail       PlaybackMode = 60
	Slave           PlaybackMode = 99
)

func (m *IsMuted) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = str[1 : len(str)-1]

	switch str {
	case "0":
		*m = false
	case "1":
		*m = true
	default:
		return fmt.Errorf("unknown IsMuted: %s", str)
	}

	return nil
}

func (m IsMuted) String() string {
	if m {
		return "Muted"
	}

	return "Unmuted"
}

func (l LoopMode) String() string {
	switch l {
	case RepeatAll:
		return "RepeatAll"
	case RepeatOne:
		return "RepeatOne"
	case ShuffleRepeat:
		return "ShuffleRepeat"
	case ShuffleOnce:
		return "ShuffleOnce"
	case Sequential:
		return "Sequential"
	default:
		return ""
	}
}

func (p *PlaybackMode) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = str[1 : len(str)-1]

	switch str {
	case "0":
		*p = None
	case "1":
		*p = Airplay
	case "2":
		*p = ThirdPartyDLNA
	case "10":
		*p = Default
	case "11":
		*p = USBPlaylist
	case "16":
		*p = TFCardPlaylist
	case "31":
		*p = SpotifyConnect
	case "32":
		*p = TidalConnect
	case "40":
		*p = AuxIn
	case "41":
		*p = Bluetooth
	case "42":
		*p = ExternalStorage
	case "43":
		*p = OpticalIn
	case "44":
		*p = Mirror
	case "60":
		*p = VoiceMail
	case "99":
		*p = Slave
	default:
		return fmt.Errorf("unknown PlaybackMode: %s", str)
	}

	return nil
}

func (p PlaybackMode) String() string {
	switch p {
	case None:
		return "None"
	case Airplay:
		return "Airplay"
	case ThirdPartyDLNA:
		return "ThirdPartyDLNA"
	case Default:
		return "Default"
	case USBPlaylist:
		return "USBPlaylist"
	case TFCardPlaylist:
		return "TFCardPlaylist"
	case SpotifyConnect:
		return "SpotifyConnect"
	case TidalConnect:
		return "TidalConnect"
	case AuxIn:
		return "AuxIn"
	case Bluetooth:
		return "Bluetooth"
	case ExternalStorage:
		return "ExternalStorage"
	case OpticalIn:
		return "OpticalIn"
	case Mirror:
		return "Mirror"
	case VoiceMail:
		return "VoiceMail"
	case Slave:
		return "Slave"
	default:
		return ""
	}
}

func (l *LoopMode) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = str[1 : len(str)-1]

	switch str {
	case "0":
		*l = RepeatAll
	case "1":
		*l = RepeatOne
	case "2":
		*l = ShuffleRepeat
	case "3":
		*l = ShuffleOnce
	case "4":
		*l = Sequential
	default:
		return fmt.Errorf("unknown LoopMode: %s", str)
	}

	return nil
}

func (s Status) String() string {
	switch s {
	case Playing:
		return "Playing"
	case Paused:
		return "Paused"
	case Stopped:
		return "Stopped"
	case Loading:
		return "Loading"
	default:
		return ""
	}
}

func (v *Status) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = str[1 : len(str)-1]

	switch str {
	case string(Playing):
		*v = Playing
	case string(Paused):
		*v = Paused
	case string(Stopped):
		*v = Stopped
	case string(Loading):
		*v = Loading
	default:
		return fmt.Errorf("unknown Status: %s", str)
	}

	return nil
}

func (v *HexString) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = str[1 : len(str)-1]

	decoded, err := hex.DecodeString(str)
	if err != nil {
		return failedToDecodeHexString(err)
	}

	*v = HexString(decoded)

	return nil
}

func (v *HexString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%x\"", *v)), nil
}

func (c Channel) String() string {
	switch c {
	case Stereo:
		return "Stereo"
	case Left:
		return "Left"
	case Right:
		return "Right"
	default:
		return ""
	}
}

func (v *Channel) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = str[1 : len(str)-1]

	switch str {
	case "0":
		*v = Stereo
	case "1":
		*v = Left
	case "2":
		*v = Right
	default:
		return fmt.Errorf("unknown Channel: %s", str)
	}

	return nil
}

func (v *Integer) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = str[1 : len(str)-1]

	value, err := strconv.Atoi(str)
	if err != nil {
		return fmt.Errorf("failed to unmarshal Volume: %w", err)
	}

	*v = Integer(value)

	return nil
}
