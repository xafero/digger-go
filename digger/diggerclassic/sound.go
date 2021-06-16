package diggerclassic

type Sound struct {
	dig                  DiggerCore
	wavetype             int
	t2val                int
	t0val                int
	musvol               int
	spkrmode             int
	timerrate            int
	timercount           int
	pulsewidth           int
	volume               int
	timerclock           int
	soundflag            bool
	musicflag            bool
	sndflag              bool
	soundpausedflag      bool
	soundlevdoneflag     bool
	nljpointer           int
	nljnoteduration      int
	newlevjingle         []int
	soundfallflag        bool
	soundfallf           bool
	soundfallvalue       int
	soundfalln           int
	soundbreakflag       bool
	soundbreakduration   int
	soundbreakvalue      int
	soundwobbleflag      bool
	soundwobblen         int
	soundfireflag        bool
	soundfirevalue       int
	soundfiren           int
	soundexplodeflag     bool
	soundexplodevalue    int
	soundexplodeduration int
	soundbonusflag       bool
	soundbonusn          int
	soundemflag          bool
	soundemeraldflag     bool
	soundemeraldduration int
	emerfreq             int
	soundemeraldn        int
	soundgoldflag        bool
	soundgoldf           bool
	soundgoldvalue1      int
	soundgoldvalue2      int
	soundgoldduration    int
	soundeatmflag        bool
	soundeatmvalue       int
	soundeatmduration    int
	soundeatmn           int
	soundddieflag        bool
	soundddien           int
	soundddievalue       int
	sound1upflag         bool
	sound1upduration     int
	musicplaying         bool
	musicp               int
	tuneno               int
	noteduration         int
	notevalue            int
	musicmaxvol          int
	musicattackrate      int
	musicsustainlevel    int
	musicdecayrate       int
	musicnotewidth       int
	musicreleaserate     int
	musicstage           int
	musicn               int
	soundt0flag          bool
	int8flag             bool
}

func NewSoundObj(d DiggerCore) *Sound {
	rcvr := new(Sound)

	rcvr.newlevjingle = []int{0x8e8, 0x712, 0x5f2, 0x7f0, 0x6ac, 0x54c, 0x712, 0x5f2, 0x4b8, 0x474, 0x474} // [11]

	rcvr.wavetype = 0
	rcvr.t2val = 0
	rcvr.t0val = 0
	rcvr.musvol = 0
	rcvr.spkrmode = 0
	rcvr.timerrate = 0x7d0
	rcvr.timercount = 0
	rcvr.pulsewidth = 1
	rcvr.volume = 0
	rcvr.timerclock = 0
	rcvr.soundflag = true
	rcvr.musicflag = true
	rcvr.sndflag = false
	rcvr.soundpausedflag = false
	rcvr.soundlevdoneflag = false
	rcvr.nljpointer = 0
	rcvr.nljnoteduration = 0
	rcvr.soundfallflag = false
	rcvr.soundfallf = false
	rcvr.soundfalln = 0
	rcvr.soundbreakflag = false
	rcvr.soundbreakduration = 0
	rcvr.soundbreakvalue = 0
	rcvr.soundwobbleflag = false
	rcvr.soundwobblen = 0
	rcvr.soundfireflag = false
	rcvr.soundfiren = 0
	rcvr.soundexplodeflag = false
	rcvr.soundbonusflag = false
	rcvr.soundbonusn = 0
	rcvr.soundemflag = false
	rcvr.soundemeraldflag = false
	rcvr.soundgoldflag = false
	rcvr.soundgoldf = false
	rcvr.soundeatmflag = false
	rcvr.soundddieflag = false
	rcvr.sound1upflag = false
	rcvr.sound1upduration = 0
	rcvr.musicplaying = false
	rcvr.musicp = 0
	rcvr.tuneno = 0
	rcvr.noteduration = 0
	rcvr.notevalue = 0
	rcvr.musicmaxvol = 0
	rcvr.musicattackrate = 0
	rcvr.musicsustainlevel = 0
	rcvr.musicdecayrate = 0
	rcvr.musicnotewidth = 0
	rcvr.musicreleaserate = 0
	rcvr.musicstage = 0
	rcvr.musicn = 0
	rcvr.soundt0flag = false
	rcvr.int8flag = false
	rcvr.dig = d
	return rcvr
}

func (rcvr *Sound) initsound() {
	rcvr.wavetype = 2
	rcvr.t0val = 12000
	rcvr.musvol = 8
	rcvr.t2val = 40
	rcvr.soundt0flag = true
	rcvr.sndflag = true
	rcvr.spkrmode = 0
	rcvr.int8flag = false
	rcvr.setsoundt2()
	rcvr.soundstop()
	rcvr.startint8()
	rcvr.timerrate = 0x4000
}

func (rcvr *Sound) killsound() {
}

func (rcvr *Sound) music(tune int) {
	rcvr.tuneno = tune
	rcvr.musicp = 0
	rcvr.noteduration = 0
	switch tune {
	case 0:
		rcvr.musicmaxvol = 50
		rcvr.musicattackrate = 20
		rcvr.musicsustainlevel = 20
		rcvr.musicdecayrate = 10
		rcvr.musicreleaserate = 4
	case 1:
		rcvr.musicmaxvol = 50
		rcvr.musicattackrate = 50
		rcvr.musicsustainlevel = 8
		rcvr.musicdecayrate = 15
		rcvr.musicreleaserate = 1
	case 2:
		rcvr.musicmaxvol = 50
		rcvr.musicattackrate = 50
		rcvr.musicsustainlevel = 25
		rcvr.musicdecayrate = 5
		rcvr.musicreleaserate = 1
	}
	rcvr.musicplaying = true
	if tune == 2 {
		rcvr.soundddieoff()
	}
}
func (rcvr *Sound) musicoff() {
	rcvr.musicplaying = false
	rcvr.musicp = 0
}
func (rcvr *Sound) musicupdate() {
	if !rcvr.musicplaying {
		return
	}
	if rcvr.noteduration != 0 {
		rcvr.noteduration--
	} else {
		rcvr.musicstage = 0
		rcvr.musicn = 0
		switch rcvr.tuneno {
		case 0:
			rcvr.musicnotewidth = rcvr.noteduration - 3
			rcvr.musicp += 2
		case 1:
			rcvr.musicnotewidth = 12
			rcvr.musicp += 2
		case 2:
			rcvr.musicnotewidth = rcvr.noteduration - 10
			rcvr.musicp += 2
		}
	}
	rcvr.musicn++
	rcvr.wavetype = 1
	rcvr.t0val = rcvr.notevalue
	if rcvr.musicn >= rcvr.musicnotewidth {
		rcvr.musicstage = 2
	}
	switch rcvr.musicstage {
	case 0:
		if rcvr.musvol+rcvr.musicattackrate >= rcvr.musicmaxvol {
			rcvr.musicstage = 1
			rcvr.musvol = rcvr.musicmaxvol
			break
		}
		rcvr.musvol += rcvr.musicattackrate
	case 1:
		if rcvr.musvol-rcvr.musicdecayrate <= rcvr.musicsustainlevel {
			rcvr.musvol = rcvr.musicsustainlevel
			break
		}
		rcvr.musvol -= rcvr.musicdecayrate
	case 2:
		if rcvr.musvol-rcvr.musicreleaserate <= 1 {
			rcvr.musvol = 1
			break
		}
		rcvr.musvol -= rcvr.musicreleaserate
	}
	if rcvr.musvol == 1 {
		rcvr.t0val = 0x7d00
	}
}
func (rcvr *Sound) s0fillbuffer() {
}
func (rcvr *Sound) s0killsound() {
	rcvr.setsoundt2()
	rcvr.stopint8()
}
func (rcvr *Sound) s0setupsound() {
	rcvr.startint8()
}
func (rcvr *Sound) setsoundmode() {
	rcvr.spkrmode = rcvr.wavetype
	if !rcvr.soundt0flag && rcvr.sndflag {
		rcvr.soundt0flag = true
	}
}
func (rcvr *Sound) setsoundt2() {
	if rcvr.soundt0flag {
		rcvr.spkrmode = 0
		rcvr.soundt0flag = false
	}
}
func (rcvr *Sound) sett0() {
	if rcvr.sndflag {
		if rcvr.t0val < 1000 && (rcvr.wavetype == 1 || rcvr.wavetype == 2) {
			rcvr.t0val = 1000
		}
		rcvr.timerrate = rcvr.t0val
		if rcvr.musvol < 1 {
			rcvr.musvol = 1
		}
		if rcvr.musvol > 50 {
			rcvr.musvol = 50
		}
		rcvr.pulsewidth = rcvr.musvol * rcvr.volume
		rcvr.setsoundmode()
	}
}
func (rcvr *Sound) sett2val(t2v int) {
}
func (rcvr *Sound) setupsound() {
}
func (rcvr *Sound) sound1up() {
	rcvr.sound1upduration = 96
	rcvr.sound1upflag = true
}
func (rcvr *Sound) sound1upoff() {
	rcvr.sound1upflag = false
}
func (rcvr *Sound) sound1upupdate() {
	if rcvr.sound1upflag {
		if rcvr.sound1upduration/3%2 != 0 {
			rcvr.t2val = rcvr.sound1upduration<<2 + 600
		}
		rcvr.sound1upduration--
		if rcvr.sound1upduration < 1 {
			rcvr.sound1upflag = false
		}
	}
}
func (rcvr *Sound) soundbonus() {
	rcvr.soundbonusflag = true
}
func (rcvr *Sound) soundbonusoff() {
	rcvr.soundbonusflag = false
	rcvr.soundbonusn = 0
}
func (rcvr *Sound) soundbonusupdate() {
	if rcvr.soundbonusflag {
		rcvr.soundbonusn++
		if rcvr.soundbonusn > 15 {
			rcvr.soundbonusn = 0
		}
		if rcvr.soundbonusn >= 0 && rcvr.soundbonusn < 6 {
			rcvr.t2val = 0x4ce
		}
		if rcvr.soundbonusn >= 8 && rcvr.soundbonusn < 14 {
			rcvr.t2val = 0x5e9
		}
	}
}
func (rcvr *Sound) soundbreak() {
	rcvr.soundbreakduration = 3
	if rcvr.soundbreakvalue < 15000 {
		rcvr.soundbreakvalue = 15000
	}
	rcvr.soundbreakflag = true
}
func (rcvr *Sound) soundbreakoff() {
	rcvr.soundbreakflag = false
}
func (rcvr *Sound) soundbreakupdate() {
	if rcvr.soundbreakflag {
		if rcvr.soundbreakduration != 0 {
			rcvr.soundbreakduration--
			rcvr.t2val = rcvr.soundbreakvalue
		} else {
			rcvr.soundbreakflag = false
		}
	}
}
func (rcvr *Sound) soundddie() {
	rcvr.soundddien = 0
	rcvr.soundddievalue = 20000
	rcvr.soundddieflag = true
}
func (rcvr *Sound) soundddieoff() {
	rcvr.soundddieflag = false
}
func (rcvr *Sound) soundddieupdate() {
	if rcvr.soundddieflag {
		rcvr.soundddien++
		if rcvr.soundddien == 1 {
			rcvr.musicoff()
		}
		if rcvr.soundddien >= 1 && rcvr.soundddien <= 10 {
			rcvr.soundddievalue = 20000 - rcvr.soundddien*1000
		}
		if rcvr.soundddien > 10 {
			rcvr.soundddievalue += 500
		}
		if rcvr.soundddievalue > 30000 {
			rcvr.soundddieoff()
		}
		rcvr.t2val = rcvr.soundddievalue
	}
}
func (rcvr *Sound) soundeatm() {
	rcvr.soundeatmduration = 20
	rcvr.soundeatmn = 3
	rcvr.soundeatmvalue = 2000
	rcvr.soundeatmflag = true
}
func (rcvr *Sound) soundeatmoff() {
	rcvr.soundeatmflag = false
}
func (rcvr *Sound) soundeatmupdate() {
	if rcvr.soundeatmflag {
		if rcvr.soundeatmn != 0 {
			if rcvr.soundeatmduration != 0 {
				if rcvr.soundeatmduration%4 == 1 {
					rcvr.t2val = rcvr.soundeatmvalue
				}
				if rcvr.soundeatmduration%4 == 3 {
					rcvr.t2val = rcvr.soundeatmvalue - int(rcvr.soundeatmvalue)>>4
				}
				rcvr.soundeatmduration--
				rcvr.soundeatmvalue -= int(rcvr.soundeatmvalue) >> 4
			} else {
				rcvr.soundeatmduration = 20
				rcvr.soundeatmn--
				rcvr.soundeatmvalue = 2000
			}
		} else {
			rcvr.soundeatmflag = false
		}
	}
}
func (rcvr *Sound) soundem() {
	rcvr.soundemflag = true
}
func (rcvr *Sound) soundemerald(emocttime int) {
	if emocttime != 0 {
		switch rcvr.emerfreq {
		case 0x8e8:
			rcvr.emerfreq = 0x7f0
		case 0x7f0:
			rcvr.emerfreq = 0x712
		case 0x712:
			rcvr.emerfreq = 0x6ac
		case 0x6ac:
			rcvr.emerfreq = 0x5f2
		case 0x5f2:
			rcvr.emerfreq = 0x54c
		case 0x54c:
			rcvr.emerfreq = 0x4b8
		case 0x4b8:
			rcvr.emerfreq = 0x474
			rcvr.dig.GetScores().scoreoctave()
		case 0x474:
			rcvr.emerfreq = 0x8e8
		}
	} else {
		rcvr.emerfreq = 0x8e8
	}
	rcvr.soundemeraldduration = 7
	rcvr.soundemeraldn = 0
	rcvr.soundemeraldflag = true
}
func (rcvr *Sound) soundemeraldoff() {
	rcvr.soundemeraldflag = false
}
func (rcvr *Sound) soundemeraldupdate() {
	if rcvr.soundemeraldflag {
		if rcvr.soundemeraldduration != 0 {
			if rcvr.soundemeraldn == 0 || rcvr.soundemeraldn == 1 {
				rcvr.t2val = rcvr.emerfreq
			}
			rcvr.soundemeraldn++
			if rcvr.soundemeraldn > 7 {
				rcvr.soundemeraldn = 0
				rcvr.soundemeraldduration--
			}
		} else {
			rcvr.soundemeraldoff()
		}
	}
}
func (rcvr *Sound) soundemoff() {
	rcvr.soundemflag = false
}
func (rcvr *Sound) soundemupdate() {
	if rcvr.soundemflag {
		rcvr.t2val = 1000
		rcvr.soundemoff()
	}
}
func (rcvr *Sound) soundexplode() {
	rcvr.soundexplodevalue = 1500
	rcvr.soundexplodeduration = 10
	rcvr.soundexplodeflag = true
	rcvr.soundfireoff()
}
func (rcvr *Sound) soundexplodeoff() {
	rcvr.soundexplodeflag = false
}
func (rcvr *Sound) soundexplodeupdate() {
	if rcvr.soundexplodeflag {
		if rcvr.soundexplodeduration != 0 {
			rcvr.soundexplodevalue = rcvr.soundexplodevalue - int(rcvr.soundexplodevalue)>>3
			rcvr.t2val = rcvr.soundexplodevalue - int(rcvr.soundexplodevalue)>>3
			rcvr.soundexplodeduration--
		} else {
			rcvr.soundexplodeflag = false
		}
	}
}

func (rcvr *Sound) soundfall() {
	rcvr.soundfallvalue = 1000
	rcvr.soundfallflag = true
}

func (rcvr *Sound) soundfalloff() {
	rcvr.soundfallflag = false
	rcvr.soundfalln = 0
}

func (rcvr *Sound) soundfallupdate() {
	if rcvr.soundfallflag {
		if rcvr.soundfalln < 1 {
			rcvr.soundfalln++
			if rcvr.soundfallf {
				rcvr.t2val = rcvr.soundfallvalue
			}
		} else {
			rcvr.soundfalln = 0
			if rcvr.soundfallf {
				rcvr.soundfallvalue += 50
				rcvr.soundfallf = false
			} else {
				rcvr.soundfallf = true
			}
		}
	}
}

func (rcvr *Sound) soundfire() {
	rcvr.soundfirevalue = 500
	rcvr.soundfireflag = true
}

func (rcvr *Sound) soundfireoff() {
	rcvr.soundfireflag = false
	rcvr.soundfiren = 0
}

func (rcvr *Sound) soundfireupdate() {
	if rcvr.soundfireflag {
		if rcvr.soundfiren == 1 {
			rcvr.soundfiren = 0
			rcvr.soundfirevalue += rcvr.soundfirevalue / 55
			rcvr.t2val = rcvr.soundfirevalue + rcvr.dig.GetMain().randno(int(rcvr.soundfirevalue)>>3)
			if rcvr.soundfirevalue > 30000 {
				rcvr.soundfireoff()
			}
		} else {
			rcvr.soundfiren++
		}
	}
}

func (rcvr *Sound) soundgold() {
	rcvr.soundgoldvalue1 = 500
	rcvr.soundgoldvalue2 = 4000
	rcvr.soundgoldduration = 30
	rcvr.soundgoldf = false
	rcvr.soundgoldflag = true
}

func (rcvr *Sound) soundgoldoff() {
	rcvr.soundgoldflag = false
}

func (rcvr *Sound) soundgoldupdate() {
	if rcvr.soundgoldflag {
		if rcvr.soundgoldduration != 0 {
			rcvr.soundgoldduration--
		} else {
			rcvr.soundgoldflag = false
		}
		if rcvr.soundgoldf {
			rcvr.soundgoldf = false
			rcvr.t2val = rcvr.soundgoldvalue1
		} else {
			rcvr.soundgoldf = true
			rcvr.t2val = rcvr.soundgoldvalue2
		}
		rcvr.soundgoldvalue1 += int(rcvr.soundgoldvalue1) >> 4
		rcvr.soundgoldvalue2 -= int(rcvr.soundgoldvalue2) >> 4
	}
}

func (rcvr *Sound) soundint() {
	rcvr.timerclock++
	if rcvr.soundflag && !rcvr.sndflag {
		rcvr.sndflag = true
		rcvr.musicflag = true
	}
	if !rcvr.soundflag && rcvr.sndflag {
		rcvr.sndflag = false
		rcvr.setsoundt2()
	}
	if rcvr.sndflag && !rcvr.soundpausedflag {
		rcvr.t0val = 0x7d00
		rcvr.t2val = 40
		if rcvr.musicflag {
			rcvr.musicupdate()
		}
		rcvr.soundemeraldupdate()
		rcvr.soundwobbleupdate()
		rcvr.soundddieupdate()
		rcvr.soundbreakupdate()
		rcvr.soundgoldupdate()
		rcvr.soundemupdate()
		rcvr.soundexplodeupdate()
		rcvr.soundfireupdate()
		rcvr.soundeatmupdate()
		rcvr.soundfallupdate()
		rcvr.sound1upupdate()
		rcvr.soundbonusupdate()
		if rcvr.t0val == 0x7d00 || rcvr.t2val != 40 {
			rcvr.setsoundt2()
		} else {
			rcvr.setsoundmode()
			rcvr.sett0()
		}
		rcvr.sett2val(rcvr.t2val)
	}
}

func (rcvr *Sound) soundlevdone() {
	/* Thread.Sleep (1000) */
}

func (rcvr *Sound) soundlevdoneoff() {
	rcvr.soundlevdoneflag = false
	rcvr.soundpausedflag = false
}

func (rcvr *Sound) soundlevdoneupdate() {
	if rcvr.sndflag {
		if rcvr.nljpointer < 11 {
			rcvr.t2val = rcvr.newlevjingle[rcvr.nljpointer]
		}
		rcvr.t0val = rcvr.t2val + 35
		rcvr.musvol = 50
		rcvr.setsoundmode()
		rcvr.sett0()
		rcvr.sett2val(rcvr.t2val)
		if rcvr.nljnoteduration > 0 {
			rcvr.nljnoteduration--
		} else {
			rcvr.nljnoteduration = 20
			rcvr.nljpointer++
			if rcvr.nljpointer > 10 {
				rcvr.soundlevdoneoff()
			}
		}
	} else {
		rcvr.soundlevdoneflag = false
	}
}

func (rcvr *Sound) soundoff() {
}

func (rcvr *Sound) soundpause() {
	rcvr.soundpausedflag = true
}

func (rcvr *Sound) soundpauseoff() {
	rcvr.soundpausedflag = false
}

func (rcvr *Sound) soundstop() {
	rcvr.soundfalloff()
	rcvr.soundwobbleoff()
	rcvr.soundfireoff()
	rcvr.musicoff()
	rcvr.soundbonusoff()
	rcvr.soundexplodeoff()
	rcvr.soundbreakoff()
	rcvr.soundemoff()
	rcvr.soundemeraldoff()
	rcvr.soundgoldoff()
	rcvr.soundeatmoff()
	rcvr.soundddieoff()
	rcvr.sound1upoff()
}

func (rcvr *Sound) soundwobble() {
	rcvr.soundwobbleflag = true
}

func (rcvr *Sound) soundwobbleoff() {
	rcvr.soundwobbleflag = false
	rcvr.soundwobblen = 0
}

func (rcvr *Sound) soundwobbleupdate() {
	if rcvr.soundwobbleflag {
		rcvr.soundwobblen++
		if rcvr.soundwobblen > 63 {
			rcvr.soundwobblen = 0
		}
		switch rcvr.soundwobblen {
		case 0:
			rcvr.t2val = 0x7d0
		case 16:
			fallthrough
		case 48:
			rcvr.t2val = 0x9c4
		case 32:
			rcvr.t2val = 0xbb8
		}
	}
}

func (rcvr *Sound) startint8() {
	if !rcvr.int8flag {
		rcvr.timerrate = 0x4000
		rcvr.int8flag = true
	}
}

func (rcvr *Sound) stopint8() {
	if rcvr.int8flag {
		rcvr.int8flag = false
	}
	rcvr.sett2val(40)
}
