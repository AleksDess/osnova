package mess

type Messedg struct {
	InfoCh chan string
	BotCh  chan string
}

var MESSEDG Messedg

// mess.MESSEDG.
// MESSEDG mess.Messedg
// mess.MESSEDG
