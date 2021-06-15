package diggerclassic

type NewSound struct {
	dig *digger
}

func NewNewSound(e *digger) *NewSound {
	d := new(NewSound)
	d.dig = e
	return d
}

func (n NewSound) fallEnd() {

}

func (n NewSound) fireStart() {

}

func (n NewSound) startBonusPulse() {

}

func (n NewSound) endBonusPulse() {

}

func (n NewSound) startBonusBackgroundMusic() {

}

func (n NewSound) fireEnd() {

}

func (n NewSound) stopNormalBackgroundMusic() {

}

func (n NewSound) startNormalBackgroundMusic() {

}

func (n NewSound) stopBonusBackgroundMusic() {

}

func (n NewSound) playDeath() {

}

func (n NewSound) killAll() {

}

func (n NewSound) playEatEmerald(full bool) {

}

func (n NewSound) playExplode() {

}

func (n NewSound) fallStart() {

}

func (n NewSound) playLooseLevel(lvl int) {

}

func (n NewSound) playBagBreak() {

}
