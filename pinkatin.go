package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	rl "github.com/lachee/raylib-goplus/raylib"
)

var ( // MARK: var

	// center line
	centerlineson bool
	// select screen
	purpiltinlocked    bool
	selectscreennumber = 1
	drawselectscreenon bool
	// intro start screens
	introcirlcesv2layout                                                             = make([]circlev2, 200)
	nicholasimontextfade                                                             = float32(0.0)
	introcircleson, nicholasimontextfadeon                                           bool
	intronicholasimontexton2, intronicholasimontexton, intropinkbackon, rayliblogoon bool
	rayliblogov2, intropinkbackv2                                                    rl.Vector2
	gologov2                                                                         rl.Vector2
	introtimer1on, introtimer2on, introtimer3on, introtimer4on, introtimer5on        bool
	introtimer1, introtimer3, introtimer4                                            = 2, 2, 2
	introtimer5                                                                      = 3
	introtimer1frame, introtimer3frame, introtimer4frame, introtimer5frame           int
	startlogofade                                                                    = float32(0.0)
	startlogofadeon, startlogofadeon2                                                bool
	drawstarton                                                                      bool
	gologo                                                                           = rl.NewRectangle(1736, 344, 284, 382)
	rayliblogo                                                                       = rl.NewRectangle(1778, 764, 256, 256)
	introenemycolor                                                                  = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
	introenemy                                                                       = 1
	chooseintroweapon                                                                = 1
	introlr                                                                          bool
	introv2                                                                          rl.Vector2
	introenemyv2                                                                     rl.Vector2
	introscreenon                                                                    bool
	introtextcolor                                                                   = rl.Magenta
	introtextcolor2                                                                  = rl.White
	// on/off timers
	onoff15, onoff30, onoff3, onoff10, onoff6 bool
	// inventory
	inventoryweaponselect         bool
	inventoryfade                 = float32(0.5)
	inventoryfadeon               bool
	inventoryselect               = 1
	inventorycount                int
	inventorypause                bool
	inventoryxswitch, inventoryon bool
	inventoryx                    = -600
	inventoryy                    = monh / 3
	// options
	controlleron, soundon, inputlayouton      bool
	optionselecty                             int
	optionson, fullscreenon, fullscreenactive bool
	optionselect                              = 1
	resolutionselect                          = 1
	// bookcase
	shieldon, flipscreen, bookcaseactive bool
	shieldfade                           = float32(0.4)
	shieldfadeon                         bool
	shieldnumber                         = 100
	// end key
	endkeycollected           bool
	endkeyarrowrotation       = float32(0)
	endkeyactive, placekey    bool
	endkeyblock, endkeynumber int
	// text
	shopflicker bool
	// tools
	pickaxeactive, spadeactive bool
	// weapons
	icefreezev2                                                                           = make([]rl.Vector2, 100)
	icefreezefade                                                                         = float32(0.5)
	icefreezefadeon                                                                       bool
	icefreezeframe, icefreezetimer                                                        int
	icefreezeactive                                                                       bool
	uzziangle                                                                             = float32(0.0)
	icewandactive, bigswordactive, uzziactive, axeactive, pumpactionactive, bazookaactive bool
	swingrighton, swinglefton                                                             bool
	weaponrotationswitch                                                                  bool
	weaponrotation                                                                        = float32(0.0)
	equippedweapon                                                                        bool
	bombtotal                                                                             = 3
	bombblock, explodeblock, bombdirection                                                int
	bombon, bombdown                                                                      bool
	pumpactiontimer, pumpactionframe                                                      int
	// shop
	shopitemprice1, shopitemprice2, shopitemprice3                                 int
	shopitemrotation1                                                              = float32(0)
	shopitemrotation2                                                              = float32(45)
	shopitemrotation3                                                              = float32(90)
	shopkeepertalkon                                                               bool
	shopkeepertalk, shopkeepertalktimer, shopkeepertalktimer2, shopkeepertalkframe int
	// loot
	chesttotal int
	// cameras
	cameraspecialtextzoom   = float32(1.0)
	cameraspecialtextzoomon bool
	// update area
	updateblock, updateblockend, updatew, updateh, updatea int
	// screen shake
	fallscreenshake   = true
	screenshakeactive bool
	timer1            = 3
	timer2            = 6
	timer3            = 3
	// quests
	questtotal                int
	questarrowrotation        = float32(0.0)
	questdetails, questactive bool
	activequest               = questitem{}
	// power ups
	mapzoomtimer                  int
	mapcollected                  bool
	meteorv2                      rl.Vector2
	createmeteor                  bool
	meteorpoint, meteorx, meteory int

	pegasustimer, meteortimer                                                                                                                                         int
	ninjastaractive, tntactive, flameactive, moneybagactive, arrowactive, meteoractive, fireballactive, colorizeactive, bulletactive, grenadactive, boxinggloveactive bool
	coffeecount, arrowcount, ninjastarcount                                                                                                                           int
	powerupcount                                                                                                                                                      int
	powerupblockmovetimer                                                                                                                                             = 3
	powerupblockmoveon                                                                                                                                                bool
	powerupsxy                                                                                                                                                        rl.Vector2
	fireballfade                                                                                                                                                      = float32(1.0)
	fireballfadeon                                                                                                                                                    bool
	// enemies
	rabbitactive, beeactive, slimeactive, skullactive, birdactive, plantactive, mushroomactive, spikesactive, radishactive, snailactive, rockactive, bigbatactive, ghostactive, pigactive, chameleonactive bool
	enemiesoverlapcount                                                                                                                                                                                    int
	enemypositions                                                                                                                                                                                         = make([]int, 100)
	enemypositionscount                                                                                                                                                                                    int
	// conversations
	conversationon, conversation1, conversation2, conversation3, conversation4, conversation5, conversation6 bool
	conversationframe                                                                                        int
	// colors fades
	lightbrown                                                                                                                                                                                                                                                                                                                                                                                 = rl.NewColor(198, 125, 67, 255)
	ghostfade                                                                                                                                                                                                                                                                                                                                                                                  = float32(1.0)
	ghostfadeon                                                                                                                                                                                                                                                                                                                                                                                bool
	chestcolor                                                                                                                                                                                                                                                                                                                                                                                 = rl.NewColor(228, 155, 88, 255)
	lightcolor1, lightcolor2, lightcolor3, lightcolor4, lightcolor5, lightcolor6, othercolor1, othercolor2, othercolor3, othercolor4, tilecolor, enemycolor1, enemycolor2, enemycolor3, enemycolor4, enemycolor5, enemycolor6, enemycolor7, enemycolor8, scenerycolor1, scenerycolor2, scenerycolor3, scenerycolor4, scenerycolor5, scenerycolor6, scenerycolor7, scenerycolor8, platformcolor rl.Color
	gemcolor                                                                                                                                                                                                                                                                                                                                                                                   = rl.Red
	charactercolor                                                                                                                                                                                                                                                                                                                                                                             = rl.NewColor(255, 219, 172, 255)
	doorfadeon                                                                                                                                                                                                                                                                                                                                                                                 bool
	doorfade                                                                                                                                                                                                                                                                                                                                                                                   = float32(0.8)
	// background
	backychange int
	// shop
	shopkeeperblock1 = 1000
	shopkeeperlr     bool
	// discoveries
	discoveryinfox, discoveryinfoy int
	discoveryinfoon                bool
	discoverytimer                 = 4
	batdiscovered                  bool
	discoveryimgx                  = -30
	discoverytext2shadowx          = -730
	discoverytext2x                = -731
	discoverytext1x                = -770
	discoverytext1shadowx          = -771
	discoveryrecx                  = -800
	discoveryon                    bool
	// environment
	underwateron, underwatertimercreated, gravityoff bool
	underwatertimer                                  int

	// player
	playerrolldegrees                                       = float32(0)
	playerrollactive                                        bool
	playerlron                                              bool
	playerattack                                            = 2
	jumppause                                               = 2
	coinstotal, gemstotal                                   int
	hpfadeon                                                bool
	hpfade                                                  = float32(0.7)
	hp                                                      = 5
	emoteon, emoteswitch                                    bool
	idletimer, chooseemote                                  int
	flossendblock, flossx, flossy, flossdirection           int
	flosson                                                 bool
	moveoff                                                 bool
	playerrunon                                             bool
	playerdirection                                         = 2
	playerychange                                           int
	emotefade                                               = float32(0.5)
	playerfade                                              = float32(1.0)
	playerfadeon, emotefadeon                               bool
	jumpheightspring                                        = 20
	jumpheight                                              = 5
	jumpcount                                               = 0
	jumpactive, fallactive, jumpactivespring                bool
	player, playerx, playery, playerh, playerv, playerstart int
	// icons
	tickicon = rl.NewRectangle(609, 46, 21, 17)
	// animations
	springon bool
	// animals
	batrotation = float32(0.0)
	// extras
	campfireflame = rl.NewRectangle(908, 253, 32, 46)
	campfire      = rl.NewRectangle(957, 221, 14, 16)
	bookcase      = rl.NewRectangle(701, 1, 16, 31)
	grave1        = rl.NewRectangle(723, 3, 16, 16)
	grave2        = rl.NewRectangle(722, 26, 18, 17)
	grave3        = rl.NewRectangle(746, 3, 12, 17)
	grave4        = rl.NewRectangle(745, 27, 15, 16)
	// traps
	spike = rl.NewRectangle(689, 5, 8, 33)
	// fx
	freezeimg                   = rl.NewRectangle(916, 184, 16, 14)
	bubbley, bubbley2, bubbley3 int
	scanlineson                 = true
	backdriftx, backdrifty      int
	backdrifton                 bool
	// images
	tileblock rl.Rectangle
	// player
	hpimg       = rl.NewRectangle(53, 156, 16, 14)
	hpimg2      = rl.NewRectangle(71, 158, 12, 10)
	playerright = rl.NewRectangle(0, 63, 16, 16)
	playerleft  = rl.NewRectangle(81, 80, 16, 16)
	player2left = rl.NewRectangle(184, 80, 16, 16)
	emote1      = rl.NewRectangle(0, 101, 32, 32)
	emote2      = rl.NewRectangle(96, 101, 32, 32)
	emote3      = rl.NewRectangle(194, 101, 32, 32)
	emote4      = rl.NewRectangle(290, 103, 32, 32)
	emote5      = rl.NewRectangle(388, 103, 32, 32)
	emote6      = rl.NewRectangle(486, 104, 32, 32)
	// options
	ctrlkey        = rl.NewRectangle(1930, 1, 95, 68)
	tabkey         = rl.NewRectangle(1830, 1, 95, 68)
	altkey         = rl.NewRectangle(1730, 2, 95, 68)
	escapekey      = rl.NewRectangle(1610, 73, 80, 76)
	spacekey       = rl.NewRectangle(1587, 1, 133, 71)
	arrowkeysimg   = rl.NewRectangle(1338, 0, 236, 150)
	controllerimg  = rl.NewRectangle(1040, 7, 120, 78)
	controllericon = rl.NewRectangle(1257, 8, 40, 28)
	keyboardicon   = rl.NewRectangle(1257, 41, 40, 25)
	musiconicon    = rl.NewRectangle(1174, 7, 31, 35)
	musicofficon   = rl.NewRectangle(1213, 8, 32, 34)
	// locked
	lockedicon = rl.NewRectangle(1023, 195, 16, 17)
	// railway
	track = rl.NewRectangle(865, 233, 16, 13)
	cart  = rl.NewRectangle(885, 232, 16, 16)
	// keys
	key1 = rl.NewRectangle(578, 307, 14, 17)
	key2 = rl.NewRectangle(602, 307, 12, 18)
	key3 = rl.NewRectangle(627, 307, 10, 18)
	key4 = rl.NewRectangle(647, 307, 14, 16)
	key5 = rl.NewRectangle(673, 306, 12, 18)
	key6 = rl.NewRectangle(696, 307, 12, 16)
	key7 = rl.NewRectangle(579, 331, 13, 18)
	key8 = rl.NewRectangle(603, 332, 10, 17)
	key9 = rl.NewRectangle(628, 332, 10, 17)
	// ufos
	ufo1  = rl.NewRectangle(922, 218, 16, 12)
	ufo2  = rl.NewRectangle(905, 233, 16, 14)
	ufo2r = rl.NewRectangle(922, 233, 16, 14)
	// powerups
	mapzoom           = rl.NewRectangle(1002, 134, 14, 14)
	flame             = rl.NewRectangle(531, 50, 20, 18)
	healthpotion      = rl.NewRectangle(1003, 106, 17, 17)
	boxingglove       = rl.NewRectangle(980, 106, 18, 17)
	boxingglover      = rl.NewRectangle(953, 133, 18, 17)
	warp              = rl.NewRectangle(957, 106, 18, 18)
	grenade           = rl.NewRectangle(933, 106, 16, 19)
	bulletrandom      = rl.NewRectangle(910, 107, 13, 18)
	colorize          = rl.NewRectangle(881, 105, 20, 19)
	spanner           = rl.NewRectangle(859, 106, 17, 18)
	gemgenerate       = rl.NewRectangle(773, 88, 16, 13)
	lightning         = rl.NewRectangle(793, 85, 19, 18)
	arrowrandom       = rl.NewRectangle(814, 83, 18, 18)
	arrowrandoml      = rl.NewRectangle(976, 132, 18, 18)
	pegasusflying     = rl.NewRectangle(837, 82, 19, 19)
	coolnoburn        = rl.NewRectangle(862, 83, 18, 18)
	zippoburning      = rl.NewRectangle(884, 80, 20, 20)
	moneybag          = rl.NewRectangle(910, 83, 17, 18)
	riprebirth        = rl.NewRectangle(933, 83, 18, 18)
	tntenemiesexplode = rl.NewRectangle(957, 82, 18, 19)
	caffeinejump      = rl.NewRectangle(980, 80, 19, 19)
	random7           = rl.NewRectangle(1002, 80, 20, 20)
	ninjastar         = rl.NewRectangle(835, 108, 16, 16)
	meteorlrg         = rl.NewRectangle(765, 4, 33, 32)
	meteorsml         = rl.NewRectangle(812, 108, 18, 18)
	starspeed         = rl.NewRectangle(786, 108, 19, 18)
	fireball          = rl.NewRectangle(609, 0, 15, 10)
	fireballl         = rl.NewRectangle(594, 0, 15, 10)
	// quest
	fox          = rl.NewRectangle(998, 442, 19, 18)
	koala        = rl.NewRectangle(975, 442, 20, 18)
	octopus      = rl.NewRectangle(954, 443, 18, 18)
	blob         = rl.NewRectangle(930, 441, 20, 18)
	pigeon       = rl.NewRectangle(906, 440, 19, 18)
	ladybug      = rl.NewRectangle(882, 439, 18, 18)
	snail        = rl.NewRectangle(857, 441, 17, 15)
	dog          = rl.NewRectangle(835, 441, 18, 17)
	budgie       = rl.NewRectangle(813, 441, 17, 18)
	snake        = rl.NewRectangle(861, 465, 19, 18)
	cat          = rl.NewRectangle(886, 465, 18, 18)
	blob2        = rl.NewRectangle(911, 467, 18, 17)
	fly          = rl.NewRectangle(933, 466, 19, 19)
	mouse        = rl.NewRectangle(958, 467, 17, 16)
	fly2         = rl.NewRectangle(979, 465, 19, 19)
	santa        = rl.NewRectangle(1000, 465, 19, 19)
	chefhat      = rl.NewRectangle(997, 491, 19, 19)
	ovenmit      = rl.NewRectangle(973, 489, 19, 19)
	icecream     = rl.NewRectangle(950, 491, 19, 19)
	gingerbread  = rl.NewRectangle(925, 491, 19, 19)
	cowboyhat    = rl.NewRectangle(902, 492, 20, 16)
	disk         = rl.NewRectangle(873, 492, 20, 20)
	socks        = rl.NewRectangle(848, 495, 19, 20)
	puzzlepiece  = rl.NewRectangle(1001, 544, 20, 20)
	legopiece    = rl.NewRectangle(1001, 568, 19, 20)
	slipper      = rl.NewRectangle(974, 543, 19, 15)
	anchor       = rl.NewRectangle(872, 521, 19, 20)
	xmasstocking = rl.NewRectangle(898, 520, 18, 18)
	pants        = rl.NewRectangle(920, 519, 17, 19)
	santahat     = rl.NewRectangle(945, 520, 20, 16)
	sunglasses   = rl.NewRectangle(971, 518, 20, 17)
	coathanger   = rl.NewRectangle(996, 517, 19, 16)
	crown        = rl.NewRectangle(865, 545, 18, 17)
	ribbon       = rl.NewRectangle(892, 545, 20, 18)
	tshirt       = rl.NewRectangle(920, 544, 19, 18)
	tophat       = rl.NewRectangle(947, 541, 18, 19)
	guitar       = rl.NewRectangle(871, 570, 20, 20)
	ring         = rl.NewRectangle(901, 574, 14, 18)
	alien        = rl.NewRectangle(921, 573, 16, 18)
	teddybear    = rl.NewRectangle(948, 571, 19, 20)
	aladdinslamp = rl.NewRectangle(977, 575, 20, 15)
	arrowup      = rl.NewRectangle(4, 520, 12, 14)
	arrowright   = rl.NewRectangle(19, 521, 14, 12)
	arrowdown    = rl.NewRectangle(36, 520, 12, 14)
	arrowleft    = rl.NewRectangle(51, 521, 14, 12)
	// doors
	bigdoor      = rl.NewRectangle(286, 0, 36, 34)
	doorsign     = rl.NewRectangle(17, 213, 16, 16)
	door7        = rl.NewRectangle(441, 39, 16, 16)
	door6        = rl.NewRectangle(425, 39, 16, 16)
	door5        = rl.NewRectangle(409, 39, 16, 16)
	door4        = rl.NewRectangle(457, 23, 16, 16)
	door3        = rl.NewRectangle(441, 23, 16, 16)
	door2        = rl.NewRectangle(425, 23, 16, 16)
	door1        = rl.NewRectangle(409, 23, 16, 16)
	exitdoor     = rl.NewRectangle(356, 22, 18, 17)
	exitdooropen = rl.NewRectangle(380, 22, 18, 17)
	// enemies images
	enemybulletr  = rl.NewRectangle(981, 42, 14, 7)
	enemybulletl  = rl.NewRectangle(964, 42, 14, 7)
	chameleonr    = rl.NewRectangle(0, 943, 84, 38)
	chameleonl    = rl.NewRectangle(0, 986, 84, 38)
	pigr          = rl.NewRectangle(440, 908, 36, 30)
	pigl          = rl.NewRectangle(0, 910, 36, 30)
	ghostl        = rl.NewRectangle(442, 875, 44, 30)
	ghostr        = rl.NewRectangle(0, 874, 44, 30)
	bigbatr       = rl.NewRectangle(544, 837, 46, 30)
	bigbatl       = rl.NewRectangle(549, 801, 46, 30)
	rockr         = rl.NewRectangle(0, 833, 38, 34)
	rockl         = rl.NewRectangle(0, 793, 38, 34)
	snailr        = rl.NewRectangle(385, 758, 38, 24)
	snaill        = rl.NewRectangle(0, 757, 38, 24)
	radishr       = rl.NewRectangle(195, 712, 30, 38)
	radishl       = rl.NewRectangle(6, 712, 30, 38)
	spikesroof    = rl.NewRectangle(368, 681, 44, 26)
	spikesfloor   = rl.NewRectangle(2, 678, 44, 26)
	spikesleft    = rl.NewRectangle(958, 669, 26, 44)
	spikesright   = rl.NewRectangle(993, 669, 26, 44)
	plantr        = rl.NewRectangle(492, 619, 44, 42)
	plantl        = rl.NewRectangle(4, 616, 44, 42)
	slimerroof    = rl.NewRectangle(1041, 981, 44, 30)
	slimelroof    = rl.NewRectangle(1036, 940, 44, 30)
	slimer        = rl.NewRectangle(102, 411, 44, 30)
	slimel        = rl.NewRectangle(97, 452, 44, 30)
	beesting      = rl.NewRectangle(115, 359, 36, 34)
	rabbitr       = rl.NewRectangle(114, 256, 34, 44)
	rabbitl       = rl.NewRectangle(110, 307, 34, 44)
	birdr         = rl.NewRectangle(523, 535, 32, 32)
	birdl         = rl.NewRectangle(229, 537, 32, 32)
	floatingskull = rl.NewRectangle(90, 490, 44, 38)
	skull         = rl.NewRectangle(55, 37, 16, 16)
	mushroom      = rl.NewRectangle(0, 568, 32, 32)
	// animals
	bat = rl.NewRectangle(0, 196, 18, 11)
	// loot
	chestopen = rl.NewRectangle(926, 134, 17, 16)
	chesthalf = rl.NewRectangle(902, 134, 17, 16)
	chest     = rl.NewRectangle(878, 134, 17, 16)
	gem       = rl.NewRectangle(105, 175, 16, 16)
	coin      = rl.NewRectangle(0, 174, 16, 16)
	// MARK: tools images
	pickaraxer = rl.NewRectangle(1105, 162, 14, 14)
	pickaraxel = rl.NewRectangle(1105, 176, 14, 14)
	dust       = rl.NewRectangle(1005, 62, 16, 16)
	spader     = rl.NewRectangle(1087, 160, 16, 16)
	spadel     = rl.NewRectangle(1087, 176, 16, 16)
	umbrella   = rl.NewRectangle(82, 174, 18, 17)
	hook       = rl.NewRectangle(0, 153, 8, 15)
	// characters
	character1 = rl.NewRectangle(1, 277, 16, 15)
	character2 = rl.NewRectangle(25, 276, 16, 16)
	character3 = rl.NewRectangle(1, 300, 16, 16)
	character4 = rl.NewRectangle(25, 300, 16, 16)
	character5 = rl.NewRectangle(1, 324, 16, 16)
	character6 = rl.NewRectangle(25, 324, 16, 16)
	// shop
	shopermarkerr = rl.NewRectangle(577, 50, 12, 12)
	shopermarker  = rl.NewRectangle(562, 50, 12, 12)
	shopkeeper    = rl.NewRectangle(0, 546, 21, 21)
	shopkeeperl   = rl.NewRectangle(112, 544, 21, 21)
	// MARK:  weapons
	icewandl           = rl.NewRectangle(1049, 159, 16, 16)
	icewandr           = rl.NewRectangle(1065, 159, 16, 16)
	bigswordr          = rl.NewRectangle(934, 163, 16, 16)
	bigswordl          = rl.NewRectangle(882, 180, 16, 16)
	uzzibullet         = rl.NewRectangle(1023, 151, 5, 5)
	uzzil              = rl.NewRectangle(1024, 174, 14, 12)
	uzzir              = rl.NewRectangle(1024, 162, 14, 12)
	rocketl            = rl.NewRectangle(878, 163, 15, 15)
	rocketr            = rl.NewRectangle(893, 163, 15, 15)
	bazooka            = rl.NewRectangle(958, 165, 16, 10)
	bazookal           = rl.NewRectangle(958, 175, 16, 10)
	axe                = rl.NewRectangle(1000, 159, 16, 16)
	axel               = rl.NewRectangle(1000, 175, 16, 16)
	shotgunbulletl     = rl.NewRectangle(1001, 271, 17, 13)
	shotgunbulletldown = rl.NewRectangle(957, 270, 15, 15)
	shotgunbulletlup   = rl.NewRectangle(980, 270, 15, 15)
	shotgunbulletr     = rl.NewRectangle(955, 249, 17, 13)
	shotgunbulletrup   = rl.NewRectangle(978, 248, 15, 15)
	shotgunbulletrdown = rl.NewRectangle(1001, 248, 15, 15)
	blast              = rl.NewRectangle(931, 359, 80, 57)
	pumpactionl        = rl.NewRectangle(978, 176, 16, 8)
	pumpaction         = rl.NewRectangle(978, 166, 16, 8)
	bomb1              = rl.NewRectangle(37, 155, 14, 13)
	sword1             = rl.NewRectangle(0, 138, 12, 12)
	sword1l            = rl.NewRectangle(12, 138, 12, 12)
	// scenery
	scenery1  = rl.NewRectangle(320, 152, 16, 17)
	scenery2  = rl.NewRectangle(342, 153, 17, 16)
	scenery3  = rl.NewRectangle(366, 152, 16, 17)
	scenery4  = rl.NewRectangle(413, 153, 16, 16)
	scenery5  = rl.NewRectangle(438, 153, 16, 16)
	scenery6  = rl.NewRectangle(461, 153, 16, 16)
	scenery7  = rl.NewRectangle(485, 153, 16, 16)
	scenery8  = rl.NewRectangle(509, 153, 16, 16)
	scenery9  = rl.NewRectangle(533, 153, 16, 16)
	scenery10 = rl.NewRectangle(564, 152, 16, 16)
	scenery11 = rl.NewRectangle(580, 152, 16, 16)
	scenery12 = rl.NewRectangle(596, 152, 16, 16)
	scenery13 = rl.NewRectangle(612, 152, 16, 16)
	scenery14 = rl.NewRectangle(628, 152, 16, 16)
	scenery15 = rl.NewRectangle(644, 152, 16, 16)
	scenery16 = rl.NewRectangle(660, 152, 16, 16)
	scenery17 = rl.NewRectangle(681, 155, 17, 17)
	scenery18 = rl.NewRectangle(702, 151, 16, 16)
	scenery19 = rl.NewRectangle(703, 177, 16, 16)
	scenery20 = rl.NewRectangle(681, 179, 18, 17)
	scenery21 = rl.NewRectangle(649, 180, 17, 17)
	scenery22 = rl.NewRectangle(625, 180, 18, 17)
	scenery23 = rl.NewRectangle(603, 179, 16, 17)
	scenery24 = rl.NewRectangle(584, 177, 16, 16)
	scenery25 = rl.NewRectangle(559, 177, 16, 17)
	scenery26 = rl.NewRectangle(535, 177, 16, 16)
	scenery27 = rl.NewRectangle(511, 178, 16, 16)
	scenery28 = rl.NewRectangle(487, 178, 16, 16)
	scenery29 = rl.NewRectangle(465, 177, 16, 16)
	scenery30 = rl.NewRectangle(441, 177, 16, 16)
	scenery31 = rl.NewRectangle(417, 177, 16, 16)
	scenery32 = rl.NewRectangle(393, 177, 16, 16)
	scenery33 = rl.NewRectangle(369, 177, 16, 16)
	scenery34 = rl.NewRectangle(345, 177, 16, 16)
	scenery35 = rl.NewRectangle(320, 176, 18, 17)
	// blocks
	marioblock     = rl.NewRectangle(129, 32, 16, 16)
	plat1          = rl.NewRectangle(353, 0, 16, 16)
	plat2          = rl.NewRectangle(369, 0, 16, 16)
	plat3          = rl.NewRectangle(385, 0, 16, 16)
	plat4          = rl.NewRectangle(401, 0, 16, 16)
	plat5          = rl.NewRectangle(417, 0, 16, 16)
	plat6          = rl.NewRectangle(433, 0, 16, 16)
	plat7          = rl.NewRectangle(449, 0, 16, 16)
	plat8          = rl.NewRectangle(465, 0, 16, 16)
	block1         = rl.NewRectangle(0, 0, 16, 16)
	block2         = rl.NewRectangle(16, 0, 16, 16)
	block3         = rl.NewRectangle(32, 0, 16, 16)
	block4         = rl.NewRectangle(48, 0, 16, 16)
	block5         = rl.NewRectangle(64, 0, 16, 16)
	block6         = rl.NewRectangle(80, 0, 16, 16)
	block7         = rl.NewRectangle(96, 0, 16, 16)
	block8         = rl.NewRectangle(112, 0, 16, 16)
	block9         = rl.NewRectangle(128, 0, 16, 16)
	block10        = rl.NewRectangle(144, 0, 16, 16)
	block11        = rl.NewRectangle(160, 0, 16, 16)
	block12        = rl.NewRectangle(176, 0, 16, 16)
	block13        = rl.NewRectangle(192, 0, 16, 16)
	block14        = rl.NewRectangle(208, 0, 16, 16)
	block15        = rl.NewRectangle(224, 0, 16, 16)
	block16        = rl.NewRectangle(240, 0, 16, 16)
	block17        = rl.NewRectangle(256, 0, 16, 16)
	block18        = rl.NewRectangle(240, 16, 16, 16)
	block19        = rl.NewRectangle(256, 16, 16, 16)
	block20        = rl.NewRectangle(240, 32, 16, 16)
	block21        = rl.NewRectangle(256, 32, 16, 16)
	platformleft   = rl.NewRectangle(1, 47, 16, 16)
	platformmiddle = rl.NewRectangle(17, 47, 16, 16)
	platformright  = rl.NewRectangle(33, 47, 16, 16)
	// back
	lamp       = rl.NewRectangle(857, 137, 15, 45)
	extraback7 = rl.NewRectangle(800, 0, 32, 32)
	extraback6 = rl.NewRectangle(832, 0, 32, 32)
	extraback5 = rl.NewRectangle(864, 0, 32, 32)
	extraback4 = rl.NewRectangle(896, 0, 32, 32)
	extraback3 = rl.NewRectangle(928, 0, 32, 32)
	extraback2 = rl.NewRectangle(960, 0, 32, 32)
	extraback1 = rl.NewRectangle(992, 0, 32, 32)
	shopback1  = rl.NewRectangle(113, 32, 16, 16)
	shopback2  = rl.NewRectangle(1006, 40, 16, 16)
	rope1      = rl.NewRectangle(647, 1, 16, 16)
	vine2      = rl.NewRectangle(669, 23, 16, 16)
	vine1      = rl.NewRectangle(644, 23, 16, 16)
	chain4     = rl.NewRectangle(619, 23, 16, 16)
	chain3     = rl.NewRectangle(596, 23, 16, 16)
	chain2     = rl.NewRectangle(572, 23, 16, 16)
	chain1     = rl.NewRectangle(548, 23, 16, 16)
	banner1    = rl.NewRectangle(0, 211, 16, 16)
	banner2    = rl.NewRectangle(0, 227, 16, 16)
	banner3    = rl.NewRectangle(0, 243, 16, 16)
	torch2     = rl.NewRectangle(26, 153, 8, 16)
	torch1     = rl.NewRectangle(10, 153, 8, 16)
	back1      = rl.NewRectangle(0, 16, 16, 16)
	back2      = rl.NewRectangle(16, 16, 16, 16)
	back3      = rl.NewRectangle(32, 16, 16, 16)
	back4      = rl.NewRectangle(48, 16, 16, 16)
	back5      = rl.NewRectangle(64, 16, 16, 16)
	back6      = rl.NewRectangle(80, 16, 16, 16)
	back7      = rl.NewRectangle(96, 16, 16, 16)
	back8      = rl.NewRectangle(112, 16, 16, 16)
	back9      = rl.NewRectangle(128, 16, 16, 16)
	back10     = rl.NewRectangle(144, 16, 16, 16)
	// interactables
	spring = rl.NewRectangle(1, 32, 16, 16)
	lever  = rl.NewRectangle(514, 0, 16, 16)
	lever2 = rl.NewRectangle(530, 0, 16, 16)
	// level
	currentlevel                  = 1
	blocktype, plattype, backtype int
	blockvert, blockhoriz         int
	blockw                        = 16
	blockh                        = 16
	drawblock, drawblocknext      int
	draww                         = 100
	drawh                         = 40
	drawarea                      = draww * drawh

	levelw = 2000 // number of 16X16 blocks wide
	levelh = 800  // number of 16X16 blocks high
	levela = levelw * levelh

	vertbelowtunnellayout = make([]tunnel, 100)
	vertbelowtunnelcount  int

	railwaylayout         = make([]string, levelw*levelh)
	trapslayout           = make([]string, levelw*levelh)
	extraslayout          = make([]string, levelw*levelh)
	shopitemsstructs      = make([]shopitem, levelw*levelh)
	shopitemslayout       = make([]string, levelw*levelh)
	textmarkerslayout     = make([]string, levelw*levelh)
	powerupscollected     = make([]powerup, 100)
	powerupsstructslayout = make([]powerup, levelw*levelh)
	powerupslayout        = make([]string, levelw*levelh)
	scenerylayout         = make([]string, levelw*levelh)
	deadenemieslayout     = make([]string, levelw*levelh)
	enemybulletslayout    = make([]string, levelw*levelh)
	enemiesstructslayout  = make([]enemy, levelw*levelh)
	enemieslayout         = make([]string, levelw*levelh)
	doorpositionslayout   = make([]int, levelw*levelh)
	doorslayout           = make([]string, levelw*levelh)
	waterlayout           = make([]string, levelw*levelh)
	animalstructslayout   = make([]animal, levelw*levelh)
	animalslayout         = make([]string, levelw*levelh)
	lootlayout            = make([]string, levelw*levelh)
	weaponslayout         = make([]string, levelw*levelh)
	grenadeslayout        = make([]grenadestruct, levelw*levelh)
	decorationslayout     = make([]string, levelw*levelh)
	leveltileslayout      = make([]string, levelw*levelh)
	levellayout           = make([]string, levelw*levelh)
	back3layout           = make([]string, levelw*levelh)
	back2layout           = make([]string, levelw*levelh)
	backlayout            = make([]string, levelw*levelh)
	interactableslayout   = make([]string, levelw*levelh)
	characterslayout      = make([]string, levelw*levelh)
	// core
	gridon                                                  bool
	fps                                                     = 30
	pauseon                                                 bool
	mouseblock                                              int
	framecount                                              int
	debugon                                                 bool
	monh, monw, monhorig, monworig                          int
	monh32, monw32                                          int32
	mousepos                                                rl.Vector2
	imgs                                                    rl.Texture2D
	camera, camera4x, camera8x, camera2x, cameraspecialtext rl.Camera2D
)

type circlev2 struct {
	center rl.Vector2
	radius float32
	color  rl.Color
}

type shopitem struct {
	item            string
	rotation, price int
}
type powerup struct {
	name   string
	active bool
	number int
	img    rl.Rectangle
}
type questitem struct {
	questtype, reward, item string
	block                   int
	img                     rl.Rectangle
}
type enemy struct {
	hp, movedir, movecount, jumpcount int
	name                              string
	img                               rl.Rectangle
	jump, fall, updown                bool
}
type animal struct {
	movement int
	name     string
	active   bool
	rotation float32
}
type tunnel struct {
	startblock, width, height int
	created                   bool
}
type grenadestruct struct {
	direction, height, distance int
}

func raylib() { // MARK: raylib
	rl.InitWindow(monw, monh, "pinkatin")
	rl.SetExitKey(rl.KeyEnd)          // key to end the game and close window
	imgs = rl.LoadTexture("imgs.png") // load images
	rl.SetTargetFPS(fps)
	// rl.HideCursor()
	// 	rl.ToggleFullscreen()
	for !rl.WindowShouldClose() {
		mousepos = rl.GetMousePosition()
		framecount++

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)

		// back layer level layout check
		drawx := 0
		drawy := 0
		count := 0
		drawblock = drawblocknext
		for a := 0; a < drawarea; a++ {
			checklevel := levellayout[drawblock]
			checkback := backlayout[drawblock]
			checkback2 := back2layout[drawblock]

			switch checklevel {
			case ".":
				if rolldice() == 6 {
					enemypositions[enemypositionscount] = drawblock
					enemypositionscount++
				}
				if enemypositionscount >= 98 {
					enemypositionscount = 0
				}

			}

			switch checkback {
			case ".":
				switch backtype {
				case 1:
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, back1, drawv2, rl.Fade(rl.DarkGray, 0.2))
					rl.DrawTextureRec(imgs, back1, drawv2, rl.Fade(rl.Black, 0.2))
					drawv2 = rl.NewVector2(float32(drawx+backdriftx), float32(drawy+backdrifty))
					rl.DrawTextureRec(imgs, back1, drawv2, rl.Fade(rl.DarkGray, 0.2))
					rl.DrawTextureRec(imgs, back1, drawv2, rl.Fade(rl.Black, 0.2))
				case 2:
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, back2, drawv2, rl.Fade(rl.DarkGray, 0.2))
					rl.DrawTextureRec(imgs, back2, drawv2, rl.Fade(rl.Black, 0.2))
					drawv2 = rl.NewVector2(float32(drawx+backdriftx), float32(drawy+backdrifty))
					rl.DrawTextureRec(imgs, back2, drawv2, rl.Fade(rl.DarkGray, 0.2))
					rl.DrawTextureRec(imgs, back2, drawv2, rl.Fade(rl.Black, 0.2))
				case 3:
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, back3, drawv2, rl.Fade(rl.DarkGray, 0.2))
					rl.DrawTextureRec(imgs, back3, drawv2, rl.Fade(rl.Black, 0.2))
					drawv2 = rl.NewVector2(float32(drawx+backdriftx), float32(drawy+backdrifty))
					rl.DrawTextureRec(imgs, back3, drawv2, rl.Fade(rl.DarkGray, 0.2))
					rl.DrawTextureRec(imgs, back3, drawv2, rl.Fade(rl.Black, 0.2))
				case 4:
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, back4, drawv2, rl.Fade(rl.DarkGray, 0.2))
					rl.DrawTextureRec(imgs, back4, drawv2, rl.Fade(rl.Black, 0.2))
					drawv2 = rl.NewVector2(float32(drawx+backdriftx), float32(drawy+backdrifty))
					rl.DrawTextureRec(imgs, back4, drawv2, rl.Fade(rl.DarkGray, 0.2))
					rl.DrawTextureRec(imgs, back4, drawv2, rl.Fade(rl.Black, 0.2))
				case 5:
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, back5, drawv2, rl.Fade(rl.DarkGray, 0.2))
					rl.DrawTextureRec(imgs, back5, drawv2, rl.Fade(rl.Black, 0.2))
					drawv2 = rl.NewVector2(float32(drawx+backdriftx), float32(drawy+backdrifty))
					rl.DrawTextureRec(imgs, back5, drawv2, rl.Fade(rl.DarkGray, 0.2))
					rl.DrawTextureRec(imgs, back5, drawv2, rl.Fade(rl.Black, 0.2))
				case 6:
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, back6, drawv2, rl.Fade(rl.DarkGray, 0.2))
					rl.DrawTextureRec(imgs, back6, drawv2, rl.Fade(rl.Black, 0.2))
					drawv2 = rl.NewVector2(float32(drawx+backdriftx), float32(drawy+backdrifty))
					rl.DrawTextureRec(imgs, back6, drawv2, rl.Fade(rl.DarkGray, 0.2))
					rl.DrawTextureRec(imgs, back6, drawv2, rl.Fade(rl.Black, 0.2))
				case 7:
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, back7, drawv2, rl.Fade(rl.DarkGray, 0.2))
					rl.DrawTextureRec(imgs, back7, drawv2, rl.Fade(rl.Black, 0.2))
					drawv2 = rl.NewVector2(float32(drawx+backdriftx), float32(drawy+backdrifty))
					rl.DrawTextureRec(imgs, back7, drawv2, rl.Fade(rl.DarkGray, 0.2))
					rl.DrawTextureRec(imgs, back7, drawv2, rl.Fade(rl.Black, 0.2))
				case 8:
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, back8, drawv2, rl.Fade(rl.DarkGray, 0.2))
					rl.DrawTextureRec(imgs, back8, drawv2, rl.Fade(rl.Black, 0.2))
					drawv2 = rl.NewVector2(float32(drawx+backdriftx), float32(drawy+backdrifty))
					rl.DrawTextureRec(imgs, back8, drawv2, rl.Fade(rl.DarkGray, 0.2))
					rl.DrawTextureRec(imgs, back8, drawv2, rl.Fade(rl.Black, 0.2))
				case 9:
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, back9, drawv2, rl.Fade(rl.DarkGray, 0.2))
					rl.DrawTextureRec(imgs, back9, drawv2, rl.Fade(rl.Black, 0.2))
					drawv2 = rl.NewVector2(float32(drawx+backdriftx), float32(drawy+backdrifty))
					rl.DrawTextureRec(imgs, back9, drawv2, rl.Fade(rl.DarkGray, 0.2))
					rl.DrawTextureRec(imgs, back9, drawv2, rl.Fade(rl.Black, 0.2))
				case 10:
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, back10, drawv2, rl.Fade(rl.DarkGray, 0.2))
					rl.DrawTextureRec(imgs, back10, drawv2, rl.Fade(rl.Black, 0.2))
					drawv2 = rl.NewVector2(float32(drawx+backdriftx), float32(drawy+backdrifty))
					rl.DrawTextureRec(imgs, back10, drawv2, rl.Fade(rl.DarkGray, 0.2))
					rl.DrawTextureRec(imgs, back10, drawv2, rl.Fade(rl.Black, 0.2))
				}
			}

			switch checkback2 {
			case "chain1":
				drawv2 := rl.NewVector2(float32(drawx), float32(drawy-backychange))
				drawv2shadow := rl.NewVector2(float32(drawx+1), float32((drawy+1)-backychange))
				rl.DrawTextureRec(imgs, chain1, drawv2shadow, rl.Fade(rl.Black, 0.5))
				rl.DrawTextureRec(imgs, chain1, drawv2, rl.Fade(rl.DarkGray, 0.2))
			case "chain2":
				drawv2 := rl.NewVector2(float32(drawx), float32(drawy-backychange))
				drawv2shadow := rl.NewVector2(float32(drawx+1), float32((drawy+1)-backychange))
				rl.DrawTextureRec(imgs, chain2, drawv2shadow, rl.Fade(rl.Black, 0.5))
				rl.DrawTextureRec(imgs, chain2, drawv2, rl.Fade(rl.DarkGray, 0.2))
			case "chain3":
				drawv2 := rl.NewVector2(float32(drawx), float32(drawy-backychange))
				drawv2shadow := rl.NewVector2(float32(drawx+1), float32((drawy+1)-backychange))
				rl.DrawTextureRec(imgs, chain3, drawv2shadow, rl.Fade(rl.Black, 0.5))
				rl.DrawTextureRec(imgs, chain3, drawv2, rl.Fade(rl.DarkGray, 0.2))
			case "chain4":
				drawv2 := rl.NewVector2(float32(drawx), float32(drawy-backychange))
				drawv2shadow := rl.NewVector2(float32(drawx+1), float32((drawy+1)-backychange))
				rl.DrawTextureRec(imgs, chain4, drawv2shadow, rl.Fade(rl.Black, 0.5))
				rl.DrawTextureRec(imgs, chain4, drawv2, rl.Fade(rl.DarkGray, 0.2))
			case "vine1":
				drawv2 := rl.NewVector2(float32(drawx), float32(drawy-backychange))
				drawv2shadow := rl.NewVector2(float32(drawx+1), float32((drawy+1)-backychange))
				rl.DrawTextureRec(imgs, vine1, drawv2shadow, rl.Fade(rl.Black, 0.5))
				rl.DrawTextureRec(imgs, vine1, drawv2, rl.Fade(rl.DarkGray, 0.2))
			case "vine2":
				drawv2 := rl.NewVector2(float32(drawx), float32(drawy-backychange))
				drawv2shadow := rl.NewVector2(float32(drawx+1), float32((drawy+1)-backychange))
				rl.DrawTextureRec(imgs, vine2, drawv2shadow, rl.Fade(rl.Black, 0.5))
				rl.DrawTextureRec(imgs, vine2, drawv2, rl.Fade(rl.DarkGray, 0.2))
			case "rope1":
				drawv2 := rl.NewVector2(float32(drawx), float32(drawy-backychange))
				drawv2shadow := rl.NewVector2(float32(drawx+1), float32((drawy+1)-backychange))
				rl.DrawTextureRec(imgs, rope1, drawv2shadow, rl.Fade(rl.Black, 0.5))
				rl.DrawTextureRec(imgs, rope1, drawv2, rl.Fade(rl.DarkGray, 0.2))
			}

			drawx += blockw
			count++
			drawblock++
			if count == draww {
				count = 0
				drawblock += levelw - draww
				drawy += blockh
				drawx = 0
			}
		}
		if !pauseon {
			//  platforms decorations scenery layer traps
			drawx = 0
			drawy = 0
			count = 0
			drawblock = drawblocknext
			for a := 0; a < drawarea; a++ {
				checkblock := levellayout[drawblock]
				checktiles := leveltileslayout[drawblock]
				checkinteractables := interactableslayout[drawblock]
				checkdecorations := decorationslayout[drawblock]
				checktraps := trapslayout[drawblock]
				checkwater := waterlayout[drawblock]
				checkdoors := doorslayout[drawblock]
				checkscenery := scenerylayout[drawblock]
				checkback3 := back3layout[drawblock]
				checktext := textmarkerslayout[drawblock]

				if createmeteor == false {
					if meteorpoint == drawblock {
						meteorx = drawx
						meteory = drawy
					}
				}

				// extra back
				switch checkback3 {
				case "extraback7":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					drawv2shadow := rl.NewVector2(float32(drawx+2), float32(drawy+2))
					rl.DrawTextureRec(imgs, extraback7, drawv2shadow, rl.Fade(rl.Black, 0.3))
					rl.DrawTextureRec(imgs, extraback7, drawv2, rl.Fade(rl.DarkGray, 0.3))
				case "extraback6":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					drawv2shadow := rl.NewVector2(float32(drawx+2), float32(drawy+2))
					rl.DrawTextureRec(imgs, extraback6, drawv2shadow, rl.Fade(rl.Black, 0.3))
					rl.DrawTextureRec(imgs, extraback6, drawv2, rl.Fade(rl.DarkGray, 0.3))
				case "extraback5":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					drawv2shadow := rl.NewVector2(float32(drawx+2), float32(drawy+2))
					rl.DrawTextureRec(imgs, extraback5, drawv2shadow, rl.Fade(rl.Black, 0.3))
					rl.DrawTextureRec(imgs, extraback5, drawv2, rl.Fade(rl.DarkGray, 0.3))
				case "extraback4":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					drawv2shadow := rl.NewVector2(float32(drawx+2), float32(drawy+2))
					rl.DrawTextureRec(imgs, extraback4, drawv2shadow, rl.Fade(rl.Black, 0.3))
					rl.DrawTextureRec(imgs, extraback4, drawv2, rl.Fade(rl.DarkGray, 0.3))
				case "extraback3":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					drawv2shadow := rl.NewVector2(float32(drawx+2), float32(drawy+2))
					rl.DrawTextureRec(imgs, extraback3, drawv2shadow, rl.Fade(rl.Black, 0.3))
					rl.DrawTextureRec(imgs, extraback3, drawv2, rl.Fade(rl.DarkGray, 0.3))
				case "extraback2":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					drawv2shadow := rl.NewVector2(float32(drawx+2), float32(drawy+2))
					rl.DrawTextureRec(imgs, extraback2, drawv2shadow, rl.Fade(rl.Black, 0.3))
					rl.DrawTextureRec(imgs, extraback2, drawv2, rl.Fade(rl.DarkGray, 0.3))
				case "extraback1":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					drawv2shadow := rl.NewVector2(float32(drawx+2), float32(drawy+2))
					rl.DrawTextureRec(imgs, extraback1, drawv2shadow, rl.Fade(rl.Black, 0.3))
					rl.DrawTextureRec(imgs, extraback1, drawv2, rl.Fade(rl.DarkGray, 0.3))

				}
				// scenery
				switch checkscenery {
				case "scenery1":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery1, drawv2, rl.Fade(scenerycolor1, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery1, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery2":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery2, drawv2, rl.Fade(scenerycolor2, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery2, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery3":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery3, drawv2, rl.Fade(scenerycolor3, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery3, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery4":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery4, drawv2, rl.Fade(scenerycolor4, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery4, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery5":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery5, drawv2, rl.Fade(scenerycolor5, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery5, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery6":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery6, drawv2, rl.Fade(scenerycolor6, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery6, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery7":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery7, drawv2, rl.Fade(scenerycolor7, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery7, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery8":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery8, drawv2, rl.Fade(scenerycolor8, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery8, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery9":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery9, drawv2, rl.Fade(scenerycolor1, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery9, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery10":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery10, drawv2, rl.Fade(scenerycolor2, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery10, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery11":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery11, drawv2, rl.Fade(scenerycolor3, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery11, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery12":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery12, drawv2, rl.Fade(scenerycolor4, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery12, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery13":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery13, drawv2, rl.Fade(scenerycolor5, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery13, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery14":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery14, drawv2, rl.Fade(scenerycolor6, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery14, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery15":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery15, drawv2, rl.Fade(scenerycolor7, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery15, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery16":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery16, drawv2, rl.Fade(scenerycolor8, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery16, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery17":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery17, drawv2, rl.Fade(scenerycolor1, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery17, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery18":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery18, drawv2, rl.Fade(scenerycolor2, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery18, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery19":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery19, drawv2, rl.Fade(scenerycolor3, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery19, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery20":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery20, drawv2, rl.Fade(scenerycolor4, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery20, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery21":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery21, drawv2, rl.Fade(scenerycolor5, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery21, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery22":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery22, drawv2, rl.Fade(scenerycolor6, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery22, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery23":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery23, drawv2, rl.Fade(scenerycolor7, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery23, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery24":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery24, drawv2, rl.Fade(scenerycolor8, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery24, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery25":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery25, drawv2, rl.Fade(scenerycolor1, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery25, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery26":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery26, drawv2, rl.Fade(scenerycolor2, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery26, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery27":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery27, drawv2, rl.Fade(scenerycolor3, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery27, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery28":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery28, drawv2, rl.Fade(scenerycolor4, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery28, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery29":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery29, drawv2, rl.Fade(scenerycolor5, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery29, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery30":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery30, drawv2, rl.Fade(scenerycolor6, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery30, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery31":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery31, drawv2, rl.Fade(scenerycolor7, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery31, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery32":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery32, drawv2, rl.Fade(scenerycolor8, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery32, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery33":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery33, drawv2, rl.Fade(scenerycolor1, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery33, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery34":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery34, drawv2, rl.Fade(scenerycolor2, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery34, drawv2, rl.Fade(rl.White, 0.3))
					}
				case "scenery35":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, scenery35, drawv2, rl.Fade(scenerycolor3, 0.3))
					} else {
						rl.DrawTextureRec(imgs, scenery35, drawv2, rl.Fade(rl.White, 0.3))
					}
				}

				// MARK: draw traps
				switch checktraps {
				case "spike3", "spike2", "spike1", "spike4", "spike5", "spike6":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, spike, drawv2, rl.White)
				}

				// MARK: draw level tiles
				switch checktiles {
				case "platl":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					drawv2shadow := rl.NewVector2(float32(drawx+2), float32(drawy+2))
					rl.DrawTextureRec(imgs, platformleft, drawv2shadow, rl.Black)
					if colorizeactive {
						rl.DrawTextureRec(imgs, platformleft, drawv2, platformcolor)
					} else {
						rl.DrawTextureRec(imgs, platformleft, drawv2, rl.White)
					}
				case "platm":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					drawv2shadow := rl.NewVector2(float32(drawx+2), float32(drawy+2))
					rl.DrawTextureRec(imgs, platformmiddle, drawv2shadow, rl.Black)
					if colorizeactive {
						rl.DrawTextureRec(imgs, platformmiddle, drawv2, platformcolor)
					} else {
						rl.DrawTextureRec(imgs, platformmiddle, drawv2, rl.White)
					}
				case "platr":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					drawv2shadow := rl.NewVector2(float32(drawx+2), float32(drawy+2))
					rl.DrawTextureRec(imgs, platformright, drawv2shadow, rl.Black)
					if colorizeactive {
						rl.DrawTextureRec(imgs, platformright, drawv2, platformcolor)
					} else {
						rl.DrawTextureRec(imgs, platformright, drawv2, rl.White)
					}
				case "<":
					rl.DrawRectangle(drawx, drawy, 15, 15, rl.Blue)
				case ">":
					rl.DrawRectangle(drawx, drawy, 15, 15, rl.Magenta)
				case "shopback1":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					drawv2shadow := rl.NewVector2(float32(drawx+2), float32(drawy+2))
					rl.DrawTextureRec(imgs, shopback1, drawv2shadow, rl.Fade(rl.Black, 0.3))
					rl.DrawTextureRec(imgs, shopback1, drawv2, rl.Fade(rl.White, 0.2))
				case "shopback2":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					drawv2shadow := rl.NewVector2(float32(drawx+2), float32(drawy+2))
					rl.DrawTextureRec(imgs, shopback2, drawv2shadow, rl.Fade(rl.Black, 0.3))
					rl.DrawTextureRec(imgs, shopback2, drawv2, rl.Fade(rl.White, 0.2))
				case "smoke3":
					rl.DrawCircle(drawx, drawy, rFloat32(7, 12), rl.Fade(rl.White, ((rFloat32(3, 9))/10)))
					leveltileslayout[drawblock] = "smoke2"
				case "smoke2":
					rl.DrawCircle(drawx, drawy, rFloat32(5, 10), rl.Fade(rl.White, ((rFloat32(3, 9))/10)))
					leveltileslayout[drawblock] = "smoke1"
				case "smoke1":
					rl.DrawCircle(drawx, drawy, rFloat32(4, 8), rl.Fade(rl.White, ((rFloat32(3, 9))/10)))
					leveltileslayout[drawblock] = "smoke"
				case "smoke":
					rl.DrawCircle(drawx, drawy, rFloat32(3, 6), rl.Fade(rl.White, ((rFloat32(3, 9))/10)))
					leveltileslayout[drawblock] = "."
				case "&":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					drawv2shadow := rl.NewVector2(float32(drawx+2), float32(drawy+2))
					rl.DrawTextureRec(imgs, marioblock, drawv2shadow, rl.Black)
					rl.DrawTextureRec(imgs, marioblock, drawv2, rl.White)
				case "%":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					drawv2shadow := rl.NewVector2(float32(drawx+2), float32(drawy+2))
					if powerupblockmoveon {
						drawv2.Y -= 8
						drawv2shadow.Y -= 8
					}
					rl.DrawTextureRec(imgs, marioblock, drawv2shadow, rl.Black)
					rl.DrawTextureRec(imgs, marioblock, drawv2, rl.Gold)

				case "_", "#":
					//	rl.DrawRectangle(drawx, drawy, 15, 15, rl.Blue)
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					drawv2shadow := rl.NewVector2(float32(drawx+2), float32(drawy+2))

					rl.DrawTextureRec(imgs, tileblock, drawv2shadow, rl.Black)
					if colorizeactive {
						rl.DrawTextureRec(imgs, tileblock, drawv2, tilecolor)
					} else {
						rl.DrawTextureRec(imgs, tileblock, drawv2, rl.White)
					}

				case "^":
					//	rl.DrawRectangle(drawx, drawy, 15, 15, rl.Blue)
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					drawv2shadow := rl.NewVector2(float32(drawx+2), float32(drawy+2))
					switch plattype {
					case 1:
						rl.DrawTextureRec(imgs, plat1, drawv2shadow, rl.Black)
						if colorizeactive {
							rl.DrawTextureRec(imgs, plat1, drawv2, platformcolor)
						} else {
							rl.DrawTextureRec(imgs, plat1, drawv2, rl.White)
						}
					case 2:
						rl.DrawTextureRec(imgs, plat2, drawv2shadow, rl.Black)
						if colorizeactive {
							rl.DrawTextureRec(imgs, plat2, drawv2, platformcolor)
						} else {
							rl.DrawTextureRec(imgs, plat2, drawv2, rl.White)
						}
					case 3:
						rl.DrawTextureRec(imgs, plat3, drawv2shadow, rl.Black)
						if colorizeactive {
							rl.DrawTextureRec(imgs, plat3, drawv2, platformcolor)
						} else {
							rl.DrawTextureRec(imgs, plat3, drawv2, rl.White)
						}
					case 4:
						rl.DrawTextureRec(imgs, plat4, drawv2shadow, rl.Black)
						if colorizeactive {
							rl.DrawTextureRec(imgs, plat4, drawv2, platformcolor)
						} else {
							rl.DrawTextureRec(imgs, plat4, drawv2, rl.White)
						}
					case 5:
						rl.DrawTextureRec(imgs, plat5, drawv2shadow, rl.Black)
						if colorizeactive {
							rl.DrawTextureRec(imgs, plat5, drawv2, platformcolor)
						} else {
							rl.DrawTextureRec(imgs, plat5, drawv2, rl.White)
						}
					case 6:
						rl.DrawTextureRec(imgs, plat6, drawv2shadow, rl.Black)
						if colorizeactive {
							rl.DrawTextureRec(imgs, plat6, drawv2, platformcolor)
						} else {
							rl.DrawTextureRec(imgs, plat6, drawv2, rl.White)
						}
					case 7:
						rl.DrawTextureRec(imgs, plat7, drawv2shadow, rl.Black)
						if colorizeactive {
							rl.DrawTextureRec(imgs, plat7, drawv2, platformcolor)
						} else {
							rl.DrawTextureRec(imgs, plat7, drawv2, rl.White)
						}
					case 8:
						rl.DrawTextureRec(imgs, plat8, drawv2shadow, rl.Black)
						if colorizeactive {
							rl.DrawTextureRec(imgs, plat8, drawv2, platformcolor)
						} else {
							rl.DrawTextureRec(imgs, plat8, drawv2, rl.White)
						}
					}

				}

				// draw text markers
				switch checktext {
				case "shopmarker":
					drawv2 := rl.NewVector2(float32(drawx-8), float32(drawy+4))
					drawv2shadow := rl.NewVector2(float32(drawx-7), float32(drawy+5))

					rl.DrawTextureRec(imgs, shopermarker, drawv2shadow, rl.Black)
					rl.DrawTextureRec(imgs, shopermarker, drawv2, rl.White)

					if rolldice() < 3 {
						rl.DrawTextureRec(imgs, shopermarker, drawv2, rl.Fade(rl.Magenta, 0.7))
					} else if rolldice() > 4 {
						rl.DrawTextureRec(imgs, shopermarker, drawv2, rl.Fade(rl.Purple, 0.7))
					}

				case "shop":
					if shopflicker {
						rl.DrawText("the", drawx+24, drawy-89, 20, rl.Black)
						rl.DrawText("shop", drawx-1, drawy-79, 30, rl.Black)
					} else {
						rl.DrawText("the", drawx+23, drawy-88, 20, rl.Black)
						rl.DrawText("shop", drawx-2, drawy-78, 30, rl.Black)
					}
					rl.DrawText("the", drawx+25, drawy-90, 20, rl.White)
					rl.DrawText("shop", drawx, drawy-80, 30, rl.White)
					if rolldice() < 3 {
						rl.DrawText("the", drawx+25, drawy-90, 20, rl.Fade(rl.Magenta, 0.7))
						rl.DrawText("shop", drawx, drawy-80, 30, rl.Fade(rl.Magenta, 0.7))
					} else if rolldice() > 4 {
						rl.DrawText("the", drawx+25, drawy-90, 20, rl.Fade(rl.Purple, 0.7))
						rl.DrawText("shop", drawx, drawy-80, 30, rl.Fade(rl.Purple, 0.7))
					}
				}

				// water
				switch checkwater {
				case "watertop":
					waveh := rInt(2, 6)
					rl.DrawRectangle(drawx, drawy+waveh, 4, 16-waveh, rl.Fade(rl.Aqua, 0.2))
					waveh = rInt(2, 6)
					rl.DrawRectangle(drawx+4, drawy+waveh, 4, 16-waveh, rl.Fade(rl.Aqua, 0.2))
					waveh = rInt(2, 6)
					rl.DrawRectangle(drawx+8, drawy+waveh, 4, 16-waveh, rl.Fade(rl.Aqua, 0.2))
					waveh = rInt(2, 6)
					rl.DrawRectangle(drawx+12, drawy+waveh, 4, 16-waveh, rl.Fade(rl.Aqua, 0.2))
				case "water":
					rl.DrawRectangle(drawx, drawy, 16, 16, rl.Fade(rl.Aqua, 0.2))
				}

				// decorations
				switch checkdecorations {
				case "lamp1":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy-29))
					torchflicker := rFloat32(2, 6)
					torchflicker = torchflicker / 10
					rl.DrawCircleGradient(int(drawv2.X+7), int(drawv2.Y+4), 40, rl.Fade(lightcolor1, torchflicker), rl.Transparent)
					rl.DrawTextureRec(imgs, lamp, drawv2, rl.Fade(rl.White, 0.8))
				case "lamp2":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy-29))
					torchflicker := rFloat32(2, 6)
					torchflicker = torchflicker / 10
					rl.DrawCircleGradient(int(drawv2.X+7), int(drawv2.Y+4), 40, rl.Fade(lightcolor2, torchflicker), rl.Transparent)
					rl.DrawTextureRec(imgs, lamp, drawv2, rl.Fade(rl.White, 0.8))
				case "lamp3":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy-29))
					torchflicker := rFloat32(2, 6)
					torchflicker = torchflicker / 10
					rl.DrawCircleGradient(int(drawv2.X+7), int(drawv2.Y+4), 40, rl.Fade(lightcolor3, torchflicker), rl.Transparent)
					rl.DrawTextureRec(imgs, lamp, drawv2, rl.Fade(rl.White, 0.8))
				case "lamp45":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy-29))
					torchflicker := rFloat32(2, 6)
					torchflicker = torchflicker / 10
					rl.DrawCircleGradient(int(drawv2.X+7), int(drawv2.Y+4), 40, rl.Fade(lightcolor4, torchflicker), rl.Transparent)
					rl.DrawTextureRec(imgs, lamp, drawv2, rl.Fade(rl.White, 0.8))
				case "lamp5":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy-29))
					torchflicker := rFloat32(2, 6)
					torchflicker = torchflicker / 10
					rl.DrawCircleGradient(int(drawv2.X+7), int(drawv2.Y+4), 40, rl.Fade(lightcolor5, torchflicker), rl.Transparent)
					rl.DrawTextureRec(imgs, lamp, drawv2, rl.Fade(rl.White, 0.8))
				case "lamp6":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy-29))
					torchflicker := rFloat32(2, 6)
					torchflicker = torchflicker / 10
					rl.DrawCircleGradient(int(drawv2.X+7), int(drawv2.Y+4), 40, rl.Fade(lightcolor6, torchflicker), rl.Transparent)
					rl.DrawTextureRec(imgs, lamp, drawv2, rl.Fade(rl.White, 0.8))
				case "torch":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					torchflicker := rFloat32(2, 6)
					torchflicker = torchflicker / 10
					rl.DrawCircleGradient(int(drawv2.X+4), int(drawv2.Y+4), 40, rl.Fade(rl.Orange, torchflicker), rl.Transparent)
					rl.DrawTextureRec(imgs, torch1, drawv2, rl.Fade(rl.White, 0.8))

				case "banner1":

					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, banner1, drawv2, rl.White)

					if rolldice() < 3 {
						rl.DrawTextureRec(imgs, banner1, drawv2, rl.Fade(rl.Magenta, 0.7))
					} else if rolldice() > 4 {
						rl.DrawTextureRec(imgs, banner1, drawv2, rl.Fade(rl.Purple, 0.7))
					}

					if rolldice() > 3 {
						rl.DrawTextureRec(imgs, banner1, drawv2, rl.Fade(rl.Magenta, 0.5))
					}
					if player == drawblock || player+1 == drawblock || player-1 == drawblock || player+2 == drawblock || player-2 == drawblock {
						tri1 := rl.NewVector2(float32(drawx+15), float32(drawy-32))
						tri2 := rl.NewVector2(float32(drawx+8), float32(drawy-32))
						tri3 := rl.NewVector2(float32(drawx+8), float32(drawy))
						rl.DrawRectangle(drawx-10, drawy-50, 82, 18, rl.Fade(rl.White, 0.8))
						rl.DrawTriangle(tri1, tri2, tri3, rl.Fade(rl.White, 0.8))
						rl.DrawText("shop nearby", drawx-2, drawy-46, 10, rl.Black)
					}

				}
				switch checkdoors {
				case "bigdoor":
					drawv2 := rl.NewVector2(float32(drawx-10), float32(drawy-18))
					rl.DrawTextureRec(imgs, bigdoor, drawv2, rl.White)
					if player == drawblock || player+1 == drawblock || player-1 == drawblock || player+2 == drawblock || player-2 == drawblock {
						tri1 := rl.NewVector2(float32(drawx+15), float32(drawy-32))
						tri2 := rl.NewVector2(float32(drawx+8), float32(drawy-32))
						tri3 := rl.NewVector2(float32(drawx+8), float32(drawy))
						rl.DrawTriangle(tri1, tri2, tri3, rl.Fade(rl.White, 0.8))
						rl.DrawRectangle(drawx-10, drawy-50, 125, 18, rl.Fade(rl.White, 0.8))
						rl.DrawText("where does this lead ?", drawx-2, drawy-46, 10, rl.Black)
					}
				}

				// interactables
				switch checkinteractables {
				case "z":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, spring, drawv2, rl.Fade(rl.White, 0.8))
				}

				// level layout
				if gridon {
					switch checkblock {
					case ".":
						rl.DrawRectangleLines(drawx, drawy, 15, 15, rl.Fade(rl.Green, 0.1))
					case "_":
						rl.DrawRectangleLines(drawx, drawy, 15, 15, rl.Fade(rl.Blue, 0.1))
					}
				}

				drawx += blockw
				count++
				drawblock++

				if count == draww {
					count = 0
					drawblock += levelw - draww
					drawy += blockh
					drawx = 0
				}
			}

			// weapons loot animals enemies layer characters shopkeeper
			drawx = 0
			drawy = 0
			count = 0
			drawblock = drawblocknext
			for a := 0; a < drawarea; a++ {

				checkweapons := weaponslayout[drawblock]
				checkloot := lootlayout[drawblock]
				checkanimals := animalslayout[drawblock]
				checkenemies := enemieslayout[drawblock]
				checkdeadenemies := deadenemieslayout[drawblock]
				checkpowerups := powerupslayout[drawblock]
				checkenemybullets := enemybulletslayout[drawblock]
				checkshopitems := shopitemslayout[drawblock]
				checkcharacters := characterslayout[drawblock]
				checkextras := extraslayout[drawblock]
				checkrailway := railwaylayout[drawblock]

				// MARK: draw animals
				if checkanimals != "" {
					checkanimaltype := animalstructslayout[drawblock]

					switch checkanimaltype.name {
					case "bat":
						if !batdiscovered {
							batdiscovered = true
							discoveryon = true
						}
						originv2 := rl.NewVector2(0, 0)
						destrec := rl.NewRectangle(float32(drawx), float32(drawy), 18, 11)
						rl.DrawTexturePro(imgs, bat, destrec, originv2, checkanimaltype.rotation, rl.Fade(rl.White, 0.3))
					}

				}

				// MARK: draw shopkeeper characters
				switch checkcharacters {
				case "character5":
					if rolldice()+rolldice() > 10 {
						drawv2 := rl.NewVector2(float32(drawx), float32(drawy-1))
						rl.DrawTextureRec(imgs, character5, drawv2, charactercolor)
					} else {
						drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
						rl.DrawTextureRec(imgs, character5, drawv2, charactercolor)
					}

					if player == drawblock || player+1 == drawblock || player-1 == drawblock || player+2 == drawblock || player-2 == drawblock {
						tri1 := rl.NewVector2(float32(drawx+15), float32(drawy-32))
						tri2 := rl.NewVector2(float32(drawx+8), float32(drawy-32))
						tri3 := rl.NewVector2(float32(drawx+8), float32(drawy))

						conversationon = true

						if questactive {
							if conversation1 {
								rl.DrawRectangle(drawx-10, drawy-50, 95, 18, rl.Fade(rl.White, 0.8))
								rl.DrawTriangle(tri1, tri2, tri3, rl.Fade(rl.White, 0.8))
								rl.DrawText("you will need to", drawx-2, drawy-46, 10, rl.Black)
							} else if conversation2 {
								rl.DrawRectangle(drawx-10, drawy-50, 132, 18, rl.Fade(rl.White, 0.8))
								rl.DrawTriangle(tri1, tri2, tri3, rl.Fade(rl.White, 0.8))
								rl.DrawText("complete your current ", drawx-2, drawy-46, 10, rl.Black)
							} else if conversation3 {
								rl.DrawRectangle(drawx-10, drawy-50, 109, 18, rl.Fade(rl.White, 0.8))
								rl.DrawTriangle(tri1, tri2, tri3, rl.Fade(rl.White, 0.8))
								rl.DrawText("quest or cancel it", drawx-2, drawy-46, 10, rl.Black)
							} else if conversation4 {
								rl.DrawRectangle(drawx-10, drawy-50, 112, 18, rl.Fade(rl.White, 0.8))
								rl.DrawTriangle(tri1, tri2, tri3, rl.Fade(rl.White, 0.8))
								rl.DrawText("to get a new quest", drawx-2, drawy-46, 10, rl.Black)
							} else if conversation5 {
								rl.DrawRectangle(drawx-10, drawy-50, 94, 18, rl.Fade(rl.White, 0.8))
								rl.DrawTriangle(tri1, tri2, tri3, rl.Fade(rl.White, 0.8))
								rl.DrawText("from me. do you", drawx-2, drawy-46, 10, rl.Black)
							} else if conversation6 {
								rl.DrawRectangle(drawx-10, drawy-50, 130, 18, rl.Fade(rl.White, 0.8))
								rl.DrawTriangle(tri1, tri2, tri3, rl.Fade(rl.White, 0.8))
								rl.DrawText("want to cancel quest?", drawx-2, drawy-46, 10, rl.Black)
							}
						} else {

							v2 := rl.NewVector2(float32(drawx-12), float32(drawy+18))
							rl.DrawTextureRec(imgs, tickicon, v2, rl.Magenta)
							rl.DrawText("yes", drawx+7, drawy+25, 10, rl.Black)
							rl.DrawText("yes", drawx+8, drawy+24, 10, rl.Magenta)

							if rl.IsKeyPressed(rl.KeyLeftControl) {
								createquest()
								questactive = true
								questdetails = true
							}

							if conversation1 {
								rl.DrawRectangle(drawx-10, drawy-50, 84, 18, rl.Fade(rl.White, 0.8))
								rl.DrawTriangle(tri1, tri2, tri3, rl.Fade(rl.White, 0.8))
								rl.DrawText("hello pinkatin", drawx-2, drawy-46, 10, rl.Black)
							} else if conversation2 {
								rl.DrawRectangle(drawx-10, drawy-50, 95, 18, rl.Fade(rl.White, 0.8))
								rl.DrawTriangle(tri1, tri2, tri3, rl.Fade(rl.White, 0.8))
								rl.DrawText("want to go on a", drawx-2, drawy-46, 10, rl.Black)
							} else if conversation3 {
								rl.DrawRectangle(drawx-10, drawy-50, 90, 18, rl.Fade(rl.White, 0.8))
								rl.DrawTriangle(tri1, tri2, tri3, rl.Fade(rl.White, 0.8))
								rl.DrawText("quest for me?", drawx-2, drawy-46, 10, rl.Black)
							} else if conversation4 {
								rl.DrawRectangle(drawx-10, drawy-50, 84, 18, rl.Fade(rl.White, 0.8))
								rl.DrawTriangle(tri1, tri2, tri3, rl.Fade(rl.White, 0.8))
								rl.DrawText("hello pinkatin", drawx-2, drawy-46, 10, rl.Black)
							} else if conversation5 {
								rl.DrawRectangle(drawx-10, drawy-50, 95, 18, rl.Fade(rl.White, 0.8))
								rl.DrawTriangle(tri1, tri2, tri3, rl.Fade(rl.White, 0.8))
								rl.DrawText("want to go on a", drawx-2, drawy-46, 10, rl.Black)
							} else if conversation6 {
								rl.DrawRectangle(drawx-10, drawy-50, 90, 18, rl.Fade(rl.White, 0.8))
								rl.DrawTriangle(tri1, tri2, tri3, rl.Fade(rl.White, 0.8))
								rl.DrawText("quest for me?", drawx-2, drawy-46, 10, rl.Black)
							}
						}
					}
				case "character1":
					if rolldice()+rolldice() > 10 {
						drawv2 := rl.NewVector2(float32(drawx), float32(drawy-1))
						rl.DrawTextureRec(imgs, character1, drawv2, charactercolor)
					} else {
						drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
						rl.DrawTextureRec(imgs, character1, drawv2, charactercolor)
					}
					if player == drawblock || player+1 == drawblock || player-1 == drawblock || player+2 == drawblock || player-2 == drawblock {
						tri1 := rl.NewVector2(float32(drawx+15), float32(drawy-32))
						tri2 := rl.NewVector2(float32(drawx+8), float32(drawy-32))
						tri3 := rl.NewVector2(float32(drawx+8), float32(drawy))

						conversationon = true

						if conversation1 {
							rl.DrawRectangle(drawx-10, drawy-50, 130, 18, rl.Fade(rl.White, 0.8))
							rl.DrawTriangle(tri1, tri2, tri3, rl.Fade(rl.White, 0.8))
							rl.DrawText("hello i am here to help", drawx-2, drawy-46, 10, rl.Black)
						} else if conversation2 {
							rl.DrawRectangle(drawx-10, drawy-50, 120, 18, rl.Fade(rl.White, 0.8))
							rl.DrawTriangle(tri1, tri2, tri3, rl.Fade(rl.White, 0.8))
							rl.DrawText("find a way to escape", drawx-2, drawy-46, 10, rl.Black)
						} else if conversation3 {
							rl.DrawRectangle(drawx-10, drawy-50, 135, 18, rl.Fade(rl.White, 0.8))
							rl.DrawTriangle(tri1, tri2, tri3, rl.Fade(rl.White, 0.8))
							rl.DrawText("the dungeons of burgle", drawx-2, drawy-46, 10, rl.Black)
						} else if conversation4 {
							rl.DrawRectangle(drawx-10, drawy-50, 130, 18, rl.Fade(rl.White, 0.8))
							rl.DrawTriangle(tri1, tri2, tri3, rl.Fade(rl.White, 0.8))
							rl.DrawText("navigate flurgle woods", drawx-2, drawy-46, 10, rl.Black)
						} else if conversation5 {
							rl.DrawRectangle(drawx-10, drawy-50, 120, 18, rl.Fade(rl.White, 0.8))
							rl.DrawTriangle(tri1, tri2, tri3, rl.Fade(rl.White, 0.8))
							rl.DrawText("climb castle crangle", drawx-2, drawy-46, 10, rl.Black)
						} else if conversation6 {
							rl.DrawRectangle(drawx-10, drawy-50, 140, 18, rl.Fade(rl.White, 0.8))
							rl.DrawTriangle(tri1, tri2, tri3, rl.Fade(rl.White, 0.8))
							rl.DrawText("rescue princess blungle", drawx-2, drawy-46, 10, rl.Black)
						}
					}

				case "shopkeeper":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy-4))
					if shopkeeperlr {
						rl.DrawTextureRec(imgs, shopkeeper, drawv2, rl.White)
						if shopkeepertalkon {
							tri1 := rl.NewVector2(float32(drawx+15), float32(drawy-32))
							tri2 := rl.NewVector2(float32(drawx+8), float32(drawy-32))
							tri3 := rl.NewVector2(float32(drawx+10), float32(drawy-8))
							rl.DrawTriangle(tri1, tri2, tri3, rl.Fade(rl.White, 0.8))
							switch shopkeepertalk {
							case 0:
								rl.DrawRectangle(drawx-10, drawy-50, 110, 18, rl.Fade(rl.White, 0.8))
								rl.DrawText("rats, rats, rats...", drawx-2, drawy-46, 10, rl.Black)
							case 1:
								rl.DrawRectangle(drawx-10, drawy-50, 100, 18, rl.Fade(rl.White, 0.8))
								rl.DrawText("mice, mice, mice...", drawx-2, drawy-46, 10, rl.Black)
							case 2:
								rl.DrawRectangle(drawx-10, drawy-50, 115, 18, rl.Fade(rl.White, 0.8))
								rl.DrawText("get out of my shop!", drawx-2, drawy-46, 10, rl.Black)
							case 3:
								rl.DrawRectangle(drawx-10, drawy-50, 110, 18, rl.Fade(rl.White, 0.8))
								rl.DrawText("another slow day...", drawx-2, drawy-46, 10, rl.Black)
							case 4:
								rl.DrawRectangle(drawx-10, drawy-50, 130, 18, rl.Fade(rl.White, 0.8))
								rl.DrawText("roaches everywhere...", drawx-2, drawy-46, 10, rl.Black)
							case 5:
								rl.DrawRectangle(drawx-10, drawy-50, 110, 18, rl.Fade(rl.White, 0.8))
								rl.DrawText("i need a vacation...", drawx-2, drawy-46, 10, rl.Black)
							case 6:
								rl.DrawRectangle(drawx-10, drawy-50, 118, 18, rl.Fade(rl.White, 0.8))
								rl.DrawText("almost closing time...", drawx-2, drawy-46, 10, rl.Black)
							case 7:
								rl.DrawRectangle(drawx-10, drawy-50, 135, 18, rl.Fade(rl.White, 0.8))
								rl.DrawText("more haste less speed", drawx-2, drawy-46, 10, rl.Black)
							case 8:
								rl.DrawRectangle(drawx-10, drawy-50, 130, 18, rl.Fade(rl.White, 0.8))
								rl.DrawText("blame it on the boogie", drawx-2, drawy-46, 10, rl.Black)
							case 9:
								rl.DrawRectangle(drawx-10, drawy-50, 110, 18, rl.Fade(rl.White, 0.8))
								rl.DrawText("i need new shoes...", drawx-2, drawy-46, 10, rl.Black)
							case 10:
								rl.DrawRectangle(drawx-10, drawy-50, 135, 18, rl.Fade(rl.White, 0.8))
								rl.DrawText("tango classes anyone?", drawx-2, drawy-46, 10, rl.Black)
							case 11:
								rl.DrawRectangle(drawx-10, drawy-50, 130, 18, rl.Fade(rl.White, 0.8))
								rl.DrawText("limbo contest anyone?", drawx-2, drawy-46, 10, rl.Black)
							case 12:
								rl.DrawRectangle(drawx-10, drawy-50, 135, 18, rl.Fade(rl.White, 0.8))
								rl.DrawText("la la bamba, la la bamba", drawx-2, drawy-46, 10, rl.Black)
							}
						}
					} else {
						rl.DrawTextureRec(imgs, shopkeeperl, drawv2, rl.White)

						if shopkeepertalkon {
							tri1 := rl.NewVector2(float32(drawx+15), float32(drawy-32))
							tri2 := rl.NewVector2(float32(drawx+8), float32(drawy-32))
							tri3 := rl.NewVector2(float32(drawx+10), float32(drawy-8))
							rl.DrawTriangle(tri1, tri2, tri3, rl.Fade(rl.White, 0.8))
							switch shopkeepertalk {
							case 0:
								rl.DrawRectangle(drawx-10, drawy-50, 110, 18, rl.Fade(rl.White, 0.8))
								rl.DrawText("rats, rats, rats...", drawx-2, drawy-46, 10, rl.Black)
							case 1:
								rl.DrawRectangle(drawx-10, drawy-50, 100, 18, rl.Fade(rl.White, 0.8))
								rl.DrawText("mice, mice, mice...", drawx-2, drawy-46, 10, rl.Black)
							case 2:
								rl.DrawRectangle(drawx-10, drawy-50, 115, 18, rl.Fade(rl.White, 0.8))
								rl.DrawText("get out of my shop!", drawx-2, drawy-46, 10, rl.Black)
							case 3:
								rl.DrawRectangle(drawx-10, drawy-50, 110, 18, rl.Fade(rl.White, 0.8))
								rl.DrawText("another slow day...", drawx-2, drawy-46, 10, rl.Black)
							case 4:
								rl.DrawRectangle(drawx-10, drawy-50, 130, 18, rl.Fade(rl.White, 0.8))
								rl.DrawText("roaches everywhere...", drawx-2, drawy-46, 10, rl.Black)
							case 5:
								rl.DrawRectangle(drawx-10, drawy-50, 110, 18, rl.Fade(rl.White, 0.8))
								rl.DrawText("i need a vacation...", drawx-2, drawy-46, 10, rl.Black)
							case 6:
								rl.DrawRectangle(drawx-10, drawy-50, 118, 18, rl.Fade(rl.White, 0.8))
								rl.DrawText("almost closing time...", drawx-2, drawy-46, 10, rl.Black)
							case 7:
								rl.DrawRectangle(drawx-10, drawy-50, 135, 18, rl.Fade(rl.White, 0.8))
								rl.DrawText("more haste less speed", drawx-2, drawy-46, 10, rl.Black)
							case 8:
								rl.DrawRectangle(drawx-10, drawy-50, 130, 18, rl.Fade(rl.White, 0.8))
								rl.DrawText("blame it on the boogie", drawx-2, drawy-46, 10, rl.Black)
							case 9:
								rl.DrawRectangle(drawx-10, drawy-50, 110, 18, rl.Fade(rl.White, 0.8))
								rl.DrawText("i need new shoes...", drawx-2, drawy-46, 10, rl.Black)
							case 10:
								rl.DrawRectangle(drawx-10, drawy-50, 135, 18, rl.Fade(rl.White, 0.8))
								rl.DrawText("tango classes anyone?", drawx-2, drawy-46, 10, rl.Black)
							case 11:
								rl.DrawRectangle(drawx-10, drawy-50, 130, 18, rl.Fade(rl.White, 0.8))
								rl.DrawText("limbo contest anyone?", drawx-2, drawy-46, 10, rl.Black)
							case 12:
								rl.DrawRectangle(drawx-10, drawy-50, 135, 18, rl.Fade(rl.White, 0.8))
								rl.DrawText("la la bamba, la la bamba", drawx-2, drawy-46, 10, rl.Black)
							}
						}
					}
					shopkeeperblock1 = drawblock
				}

				switch checkdeadenemies {
				case "skull":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, skull, drawv2, charactercolor)
				case "deadenemy":
					rl.DrawCircle(drawx, drawy, rFloat32(12, 18), rl.Fade(rl.Maroon, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx-16, drawy-16, rFloat32(12, 18), rl.Fade(rl.Maroon, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx+16, drawy, rFloat32(12, 18), rl.Fade(rl.Maroon, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx+16, drawy+16, rFloat32(12, 18), rl.Fade(rl.Maroon, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx, drawy+12, rFloat32(12, 18), rl.Fade(rl.Maroon, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx+32, drawy+32, rFloat32(12, 18), rl.Fade(rl.Maroon, ((rFloat32(3, 9))/10)))
					deadenemieslayout[drawblock] = "deadenemy2"
				case "deadenemy2":
					rl.DrawCircle(drawx, drawy, rFloat32(8, 15), rl.Fade(rl.Maroon, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx-8, drawy-8, rFloat32(8, 15), rl.Fade(rl.Maroon, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx+16, drawy-16, rFloat32(8, 15), rl.Fade(rl.Maroon, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx+16, drawy+16, rFloat32(8, 15), rl.Fade(rl.Maroon, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx, drawy+12, rFloat32(8, 15), rl.Fade(rl.Maroon, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx+32, drawy+32, rFloat32(8, 15), rl.Fade(rl.Maroon, ((rFloat32(3, 9))/10)))
					deadenemieslayout[drawblock] = "deadenemy3"
				case "deadenemy3":
					rl.DrawCircle(drawx, drawy, rFloat32(3, 8), rl.Fade(rl.Maroon, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx-8, drawy-8, rFloat32(3, 8), rl.Fade(rl.Maroon, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx+16, drawy-16, rFloat32(3, 8), rl.Fade(rl.Maroon, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx+16, drawy+16, rFloat32(3, 8), rl.Fade(rl.Maroon, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx, drawy+12, rFloat32(3, 8), rl.Fade(rl.Maroon, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx+32, drawy+32, rFloat32(3, 8), rl.Fade(rl.Maroon, ((rFloat32(3, 9))/10)))
					deadenemieslayout[drawblock] = "deadenemy4"
				case "deadenemy4":
					rl.DrawCircle(drawx, drawy, rFloat32(2, 5), rl.Fade(rl.Maroon, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx-8, drawy-8, rFloat32(2, 5), rl.Fade(rl.Maroon, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx+16, drawy-16, rFloat32(2, 5), rl.Fade(rl.Maroon, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx+16, drawy+16, rFloat32(2, 5), rl.Fade(rl.Maroon, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx, drawy+12, rFloat32(2, 5), rl.Fade(rl.Maroon, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx+32, drawy+32, rFloat32(2, 5), rl.Fade(rl.Maroon, ((rFloat32(3, 9))/10)))
					deadenemieslayout[drawblock] = "skull"
				}

				// MARK: draw extras
				switch checkextras {
				case "campfire":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, campfire, drawv2, rl.White)
					drawv2 = rl.NewVector2(float32(drawx-9), float32(drawy-28))
					rl.DrawTextureRec(imgs, campfireflame, drawv2, rl.Fade(rl.White, rF32(0.1, 0.6)))

				case "grave1":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, grave1, drawv2, rl.White)
				case "grave2":
					drawv2 := rl.NewVector2(float32(drawx-1), float32(drawy-1))
					rl.DrawTextureRec(imgs, grave2, drawv2, rl.White)
				case "grave3":
					drawv2 := rl.NewVector2(float32(drawx+2), float32(drawy))
					rl.DrawTextureRec(imgs, grave3, drawv2, rl.White)
				case "grave4":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, grave4, drawv2, rl.White)
				case "bookcase":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy-15))
					rl.DrawTextureRec(imgs, bookcase, drawv2, lightbrown)
					if bookcaseactive {
						rl.DrawRectangle(drawx-12, drawy-34, 40, 16, rl.Fade(rl.White, 0.7))
						rl.DrawText("read?", drawx-6, drawy-30, 10, rl.Black)

						if rl.IsKeyPressed(rl.KeyLeftControl) {
							readbook()
						}
					}
				case "ufo1", "ufo1r":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, ufo1, drawv2, rl.Fade(othercolor1, 0.8))
					} else {
						rl.DrawTextureRec(imgs, ufo1, drawv2, rl.White)
					}
				case "ufo2":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, ufo2, drawv2, othercolor2)
					} else {
						rl.DrawTextureRec(imgs, ufo2, drawv2, rl.White)
					}
				case "ufo2r":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					if colorizeactive {
						rl.DrawTextureRec(imgs, ufo2r, drawv2, othercolor2)
					} else {
						rl.DrawTextureRec(imgs, ufo2r, drawv2, rl.White)
					}
				}

				// MARK: draw railway
				switch checkrailway {
				case "track":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy-4))
					rl.DrawTextureRec(imgs, track, drawv2, rl.Brown)
				}

				// MARK: draw enemy bullets
				switch checkenemybullets {
				case "bulletr":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, enemybulletr, drawv2, rl.Green)
				case "bulletl":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, enemybulletl, drawv2, rl.Green)
				}

				// MARK: draw enemies
				switch checkenemies {
				case "pig":
					check := enemiesstructslayout[drawblock]
					drawv2 := rl.NewVector2(float32(drawx-10), float32(drawy-14))
					if check.movedir == 4 {
						if colorizeactive {
							rl.DrawTextureRec(imgs, pigl, drawv2, enemycolor5)
						} else {
							rl.DrawTextureRec(imgs, pigl, drawv2, rl.White)
						}
					} else {
						if colorizeactive {
							rl.DrawTextureRec(imgs, pigr, drawv2, enemycolor5)
						} else {
							rl.DrawTextureRec(imgs, pigr, drawv2, rl.White)
						}
					}
				case "ghost":
					check := enemiesstructslayout[drawblock]
					drawv2 := rl.NewVector2(float32(drawx-14), float32(drawy-14))
					switch check.movedir {
					case 4, 2, 3, 5:
						rl.DrawTextureRec(imgs, ghostr, drawv2, rl.Fade(rl.White, ghostfade))
					case 8, 6, 7, 1:
						rl.DrawTextureRec(imgs, ghostl, drawv2, rl.Fade(rl.White, ghostfade))
					}
				case "bat":
					check := enemiesstructslayout[drawblock]
					drawv2 := rl.NewVector2(float32(drawx-15), float32(drawy-14))
					switch check.movedir {
					case 4, 2, 3, 5:
						if colorizeactive {
							rl.DrawTextureRec(imgs, bigbatr, drawv2, enemycolor4)
						} else {
							rl.DrawTextureRec(imgs, bigbatr, drawv2, rl.White)
						}
					case 8, 6, 7, 1:
						if colorizeactive {
							rl.DrawTextureRec(imgs, bigbatl, drawv2, enemycolor4)
						} else {
							rl.DrawTextureRec(imgs, bigbatl, drawv2, rl.White)
						}
					}
				case "rock":
					check := enemiesstructslayout[drawblock]
					drawv2 := rl.NewVector2(float32(drawx-11), float32(drawy-18))
					if check.movedir == 4 {
						if colorizeactive {
							rl.DrawTextureRec(imgs, rockr, drawv2, enemycolor3)
						} else {
							rl.DrawTextureRec(imgs, rockr, drawv2, rl.White)
						}
					} else {
						if colorizeactive {
							rl.DrawTextureRec(imgs, rockl, drawv2, enemycolor3)
						} else {
							rl.DrawTextureRec(imgs, rockl, drawv2, rl.White)
						}
					}
				case "snail":
					check := enemiesstructslayout[drawblock]
					drawv2 := rl.NewVector2(float32(drawx-11), float32(drawy-8))
					if check.movedir == 4 {
						if colorizeactive {
							rl.DrawTextureRec(imgs, snailr, drawv2, enemycolor2)
						} else {
							rl.DrawTextureRec(imgs, snailr, drawv2, rl.White)
						}
					} else {
						if colorizeactive {
							rl.DrawTextureRec(imgs, snaill, drawv2, enemycolor2)
						} else {
							rl.DrawTextureRec(imgs, snaill, drawv2, rl.White)
						}
					}
				case "radish":
					check := enemiesstructslayout[drawblock]
					drawv2 := rl.NewVector2(float32(drawx-8), float32(drawy-22))
					if check.movedir == 4 || check.movedir == 2 {
						if colorizeactive {
							rl.DrawTextureRec(imgs, radishr, drawv2, enemycolor1)
						} else {
							rl.DrawTextureRec(imgs, radishr, drawv2, rl.White)
						}
					} else if check.movedir == 8 || check.movedir == 6 {
						if colorizeactive {
							rl.DrawTextureRec(imgs, radishl, drawv2, enemycolor1)
						} else {
							rl.DrawTextureRec(imgs, radishl, drawv2, rl.White)
						}
					}
				case "spike":
					check := enemiesstructslayout[drawblock]
					if check.updown {
						if check.movedir == 2 {
							if levellayout[drawblock+1] == "." {
								drawv2 := rl.NewVector2(float32(drawx), float32(drawy-14))
								if colorizeactive {
									rl.DrawTextureRec(imgs, spikesleft, drawv2, enemycolor8)
								} else {
									rl.DrawTextureRec(imgs, spikesleft, drawv2, rl.White)
								}
							} else if levellayout[drawblock-1] == "." {
								drawv2 := rl.NewVector2(float32(drawx-10), float32(drawy-14))
								if colorizeactive {
									rl.DrawTextureRec(imgs, spikesright, drawv2, enemycolor8)
								} else {
									rl.DrawTextureRec(imgs, spikesright, drawv2, rl.White)
								}
							}
						} else if check.movedir == 4 {
							drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
							if colorizeactive {
								rl.DrawTextureRec(imgs, spikesroof, drawv2, enemycolor8)
							} else {
								rl.DrawTextureRec(imgs, spikesroof, drawv2, rl.White)
							}
						} else if check.movedir == 8 {
							drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
							if colorizeactive {
								rl.DrawTextureRec(imgs, spikesroof, drawv2, enemycolor8)
							} else {
								rl.DrawTextureRec(imgs, spikesroof, drawv2, rl.White)
							}
						} else if check.movedir == 6 {
							if levellayout[drawblock+1] == "." {
								drawv2 := rl.NewVector2(float32(drawx), float32(drawy-14))
								if colorizeactive {
									rl.DrawTextureRec(imgs, spikesleft, drawv2, enemycolor8)
								} else {
									rl.DrawTextureRec(imgs, spikesleft, drawv2, rl.White)
								}
							} else if levellayout[drawblock-1] == "." {
								drawv2 := rl.NewVector2(float32(drawx-10), float32(drawy-14))
								if colorizeactive {
									rl.DrawTextureRec(imgs, spikesright, drawv2, enemycolor8)
								} else {
									rl.DrawTextureRec(imgs, spikesright, drawv2, rl.White)
								}
							}
						}
					} else {
						drawv2 := rl.NewVector2(float32(drawx), float32(drawy-10))
						if colorizeactive {
							rl.DrawTextureRec(imgs, spikesfloor, drawv2, enemycolor8)
						} else {
							rl.DrawTextureRec(imgs, spikesfloor, drawv2, rl.White)
						}
					}

				case "plant":
					check := enemiesstructslayout[drawblock]
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy-26))
					if check.movedir == 4 {
						if colorizeactive {
							rl.DrawTextureRec(imgs, plantr, drawv2, enemycolor7)
						} else {
							rl.DrawTextureRec(imgs, plantr, drawv2, rl.White)
						}
					} else {
						if colorizeactive {
							rl.DrawTextureRec(imgs, plantl, drawv2, enemycolor7)
						} else {
							rl.DrawTextureRec(imgs, plantl, drawv2, rl.White)
						}
					}
				case "slime":
					check := enemiesstructslayout[drawblock]
					if check.updown {
						drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
						if check.movedir == 4 {
							if colorizeactive {
								rl.DrawTextureRec(imgs, slimerroof, drawv2, enemycolor6)
							} else {
								rl.DrawTextureRec(imgs, slimerroof, drawv2, rl.White)
							}
						} else {
							if colorizeactive {
								rl.DrawTextureRec(imgs, slimelroof, drawv2, enemycolor6)
							} else {
								rl.DrawTextureRec(imgs, slimelroof, drawv2, rl.White)
							}
						}
					} else {
						drawv2 := rl.NewVector2(float32(drawx), float32(drawy-14))
						if check.movedir == 4 {
							rl.DrawTextureRec(imgs, slimer, drawv2, rl.White)
						} else {
							rl.DrawTextureRec(imgs, slimel, drawv2, rl.White)
						}
					}

				case "rabbit":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy-28))
					check := enemiesstructslayout[drawblock]
					switch check.movedir {
					case 4, 3:
						if colorizeactive {
							rl.DrawTextureRec(imgs, rabbitr, drawv2, enemycolor5)
						} else {
							rl.DrawTextureRec(imgs, rabbitr, drawv2, rl.White)
						}
					case 8, 1:
						if colorizeactive {
							rl.DrawTextureRec(imgs, rabbitl, drawv2, enemycolor4)
						} else {
							rl.DrawTextureRec(imgs, rabbitl, drawv2, rl.White)
						}
					}
				case "bee":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy-8))
					if colorizeactive {
						rl.DrawTextureRec(imgs, beesting, drawv2, enemycolor4)
					} else {
						rl.DrawTextureRec(imgs, beesting, drawv2, rl.White)
					}
				case "bird":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy-12))
					checkenemy := enemiesstructslayout[drawblock]
					if checkenemy.movedir == 2 {
						if colorizeactive {
							rl.DrawTextureRec(imgs, birdr, drawv2, enemycolor1)
						} else {
							rl.DrawTextureRec(imgs, birdr, drawv2, rl.White)
						}
					} else {
						if colorizeactive {
							rl.DrawTextureRec(imgs, birdl, drawv2, enemycolor1)
						} else {
							rl.DrawTextureRec(imgs, birdl, drawv2, rl.White)
						}
					}
				case "skull":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy-12))
					if colorizeactive {
						rl.DrawTextureRec(imgs, floatingskull, drawv2, enemycolor2)
					} else {
						rl.DrawTextureRec(imgs, floatingskull, drawv2, rl.White)
					}
				case "warpmushroom":
					rl.DrawCircle(drawx+8, drawy+8, 64, rl.Black)
					rl.DrawCircleLines(drawx+8, drawy+8, 64, rl.Fade(rl.White, 0.7))
					enemieslayout[drawblock] = "warpmushroom1"
				case "warpmushroom1":
					rl.DrawCircle(drawx+8, drawy+8, 48, rl.Black)
					rl.DrawCircleLines(drawx+8, drawy+8, 48, rl.Fade(rl.White, 0.7))
					enemieslayout[drawblock] = "warpmushroom2"
				case "warpmushroom2":
					rl.DrawCircle(drawx+8, drawy+8, 32, rl.Black)
					rl.DrawCircleLines(drawx+8, drawy+8, 32, rl.Fade(rl.White, 0.7))
					enemieslayout[drawblock] = "warpmushroom3"
				case "warpmushroom3":
					rl.DrawCircle(drawx+8, drawy+8, 16, rl.Black)
					rl.DrawCircleLines(drawx+8, drawy+8, 16, rl.Fade(rl.White, 0.7))
					enemieslayout[drawblock] = "warpmushroom4"
				case "warpmushroom4":
					rl.DrawCircle(drawx+8, drawy+8, 8, rl.Black)
					rl.DrawCircleLines(drawx+8, drawy+8, 8, rl.Fade(rl.White, 0.7))
					enemieslayout[drawblock] = "mushroom"
				case "mushroom":
					drawv2 := rl.NewVector2(float32(drawx-8), float32(drawy-16))
					if colorizeactive {
						rl.DrawTextureRec(imgs, mushroom, drawv2, enemycolor3)
					} else {
						rl.DrawTextureRec(imgs, mushroom, drawv2, rl.White)
					}
				}

				// MARK: draw weapons
				switch checkweapons {
				case "uzzibulletup", "uzzibulletrup", "uzzibulletr", "uzzibulletlup", "uzzibulletl":
					drawv2 := rl.NewVector2(float32(drawx+5), float32(drawy+5))
					rl.DrawTextureRec(imgs, uzzibullet, drawv2, rl.White)
				case "rocketl":
					originv2 := rl.NewVector2(0, 0)
					destrec := rl.NewRectangle(float32(drawx), float32(drawy), 16, 16)
					rl.DrawTexturePro(imgs, rocketl, destrec, originv2, -45, rl.White)
				case "rocketr":
					originv2 := rl.NewVector2(0, 0)
					destrec := rl.NewRectangle(float32(drawx), float32(drawy), 16, 16)
					rl.DrawTexturePro(imgs, rocketr, destrec, originv2, 45, rl.White)
				case "shotgunbulletright":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, shotgunbulletr, drawv2, rl.White)
				case "shotgunbulletleft":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, shotgunbulletl, drawv2, rl.White)
				case "shotgunbulletdownright":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, shotgunbulletrdown, drawv2, rl.White)
				case "shotgunbulletupright":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, shotgunbulletrup, drawv2, rl.White)
				case "shotgunbulletupleft":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, shotgunbulletlup, drawv2, rl.White)
				case "shotgunbulletdownleft":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, shotgunbulletldown, drawv2, rl.White)
				case "boxingglover":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, boxingglover, drawv2, rl.White)
				case "boxingglove":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, boxingglove, drawv2, rl.White)
				case "grenade":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, grenade, drawv2, rl.White)
				case "bullet":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, bulletrandom, drawv2, rl.White)
				case "ninjastar1", "ninjastar2", "ninjastar3", "ninjastar4", "ninjastar5", "ninjastar6", "ninjastar7", "ninjastar8":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, ninjastar, drawv2, rl.White)
				case "meteortl":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, meteorlrg, drawv2, rl.White)
				case "arrowl":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, arrowrandoml, drawv2, rl.White)
				case "arrow":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, arrowrandom, drawv2, rl.White)
				case "flame":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, flame, drawv2, rl.Orange)
					if framecount%6 == 0 {
						weaponslayout[drawblock] = ""
					}
				case "flame1":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, flame, drawv2, rl.Fade(rl.Orange, 0.8))
					if framecount%6 == 0 {
						weaponslayout[drawblock] = ""
						weaponslayout[drawblock+1] = "flame"
					}
				case "flame2":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, flame, drawv2, rl.Fade(rl.Orange, 0.6))
					if framecount%6 == 0 {
						weaponslayout[drawblock] = ""
						weaponslayout[drawblock+1] = "flame1"
					}
				case "flame3":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, flame, drawv2, rl.Fade(rl.Orange, 0.4))
					if framecount%6 == 0 {
						weaponslayout[drawblock] = ""
						weaponslayout[drawblock+1] = "flame2"
					}
				case "flame4":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, flame, drawv2, rl.Fade(rl.Orange, 0.2))
					if framecount%6 == 0 {
						weaponslayout[drawblock] = ""
						weaponslayout[drawblock+1] = "flame3"
					}
				case "flame8":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, flame, drawv2, rl.Fade(rl.Orange, 0.8))
					weaponslayout[drawblock] = ""
					weaponslayout[drawblock-1] = "flame"
				case "flame7":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, flame, drawv2, rl.Fade(rl.Orange, 0.6))
					weaponslayout[drawblock] = ""
					weaponslayout[drawblock-1] = "flame8"
				case "flame6":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, flame, drawv2, rl.Fade(rl.Orange, 0.4))
					weaponslayout[drawblock] = ""
					weaponslayout[drawblock-1] = "flame7"

				case "flame5":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, flame, drawv2, rl.Fade(rl.Orange, 0.2))
					weaponslayout[drawblock] = ""
					weaponslayout[drawblock-1] = "flame6"
				case "deadfireball":
					rl.DrawCircle(drawx, drawy, rFloat32(12, 18), rl.Fade(rl.Orange, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx-16, drawy-16, rFloat32(12, 18), rl.Fade(rl.Orange, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx+16, drawy, rFloat32(12, 18), rl.Fade(rl.Orange, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx+16, drawy+16, rFloat32(12, 18), rl.Fade(rl.Orange, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx, drawy+12, rFloat32(12, 18), rl.Fade(rl.Orange, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx+32, drawy+32, rFloat32(12, 18), rl.Fade(rl.Orange, ((rFloat32(3, 9))/10)))
					weaponslayout[drawblock] = "deadfireball2"
				case "deadfireball2":
					rl.DrawCircle(drawx, drawy, rFloat32(8, 15), rl.Fade(rl.Orange, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx-8, drawy-8, rFloat32(8, 15), rl.Fade(rl.Orange, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx+16, drawy-16, rFloat32(8, 15), rl.Fade(rl.Orange, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx+16, drawy+16, rFloat32(8, 15), rl.Fade(rl.Orange, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx, drawy+12, rFloat32(8, 15), rl.Fade(rl.Orange, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx+32, drawy+32, rFloat32(8, 15), rl.Fade(rl.Orange, ((rFloat32(3, 9))/10)))
					weaponslayout[drawblock] = "deadfireball3"
				case "deadfireball3":
					rl.DrawCircle(drawx, drawy, rFloat32(3, 8), rl.Fade(rl.Orange, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx-8, drawy-8, rFloat32(3, 8), rl.Fade(rl.Orange, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx+16, drawy-16, rFloat32(3, 8), rl.Fade(rl.Orange, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx+16, drawy+16, rFloat32(3, 8), rl.Fade(rl.Orange, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx, drawy+12, rFloat32(3, 8), rl.Fade(rl.Orange, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx+32, drawy+32, rFloat32(3, 8), rl.Fade(rl.Orange, ((rFloat32(3, 9))/10)))
					weaponslayout[drawblock] = "deadfireball4"
				case "deadfireball4":
					rl.DrawCircle(drawx, drawy, rFloat32(2, 5), rl.Fade(rl.Orange, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx-8, drawy-8, rFloat32(2, 5), rl.Fade(rl.Orange, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx+16, drawy-16, rFloat32(2, 5), rl.Fade(rl.Orange, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx+16, drawy+16, rFloat32(2, 5), rl.Fade(rl.Orange, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx, drawy+12, rFloat32(2, 5), rl.Fade(rl.Orange, ((rFloat32(3, 9))/10)))
					rl.DrawCircle(drawx+32, drawy+32, rFloat32(2, 5), rl.Fade(rl.Orange, ((rFloat32(3, 9))/10)))
					weaponslayout[drawblock] = ""
				case "fireball":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, fireball, drawv2, rl.Fade(rl.Orange, fireballfade))
				case "fireballl":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, fireballl, drawv2, rl.Fade(rl.Orange, fireballfade))
				case "bomb":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, bomb1, drawv2, rl.White)
				}

				// MARK: draw powerups
				if checkpowerups != "" && checkpowerups != "map" {
					checkpowerup := powerupsstructslayout[drawblock]
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy-4))
					drawv2shadow := rl.NewVector2(float32(drawx-1), float32(drawy-3))
					rl.DrawTextureRec(imgs, checkpowerup.img, drawv2shadow, rl.Black)
					rl.DrawTextureRec(imgs, checkpowerup.img, drawv2, rl.White)

					arrowv2 := rl.NewVector2(float32(drawx+2), float32(drawy-24-playerychange))
					rl.DrawTextureRec(imgs, arrowdown, arrowv2, rl.Magenta)
					arrowv2.X++
					arrowv2.Y--
					rl.DrawTextureRec(imgs, arrowdown, arrowv2, rl.White)
					rl.DrawText("powerup", drawx-12, drawy-38, 10, rl.Black)
					rl.DrawText("powerup", drawx-11, drawy-39, 10, rl.Magenta)
					rl.DrawText("powerup", drawx-10, drawy-40, 10, rl.White)

					if player == drawblock {
						collectpowerup(drawblock)
					}
				} else if checkpowerups == "map" {
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, mapzoom, drawv2, rl.White)
					arrowv2 := rl.NewVector2(float32(drawx), float32(drawy-18-playerychange))
					rl.DrawTextureRec(imgs, arrowdown, arrowv2, rl.Magenta)
					arrowv2.X++
					arrowv2.Y--
					rl.DrawTextureRec(imgs, arrowdown, arrowv2, rl.White)
					rl.DrawText("map", drawx-2, drawy-32, 10, rl.Black)
					rl.DrawText("map", drawx-1, drawy-33, 10, rl.Magenta)
					rl.DrawText("map", drawx, drawy-34, 10, rl.White)
					if player == drawblock {
						mapcollected = true
						mapzoomtimer = 5
						powerupslayout[drawblock] = ""
					}
				}

				// MARK: draw loot
				switch checkloot {

				case "bazooka":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, bazooka, drawv2, rl.White)
				case "uzzi":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, uzzir, drawv2, rl.White)
				case "sword":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, bigswordr, drawv2, rl.White)
				case "icewand":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, icewandr, drawv2, rl.White)
				case "pickaxe":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, pickaraxer, drawv2, rl.White)
				case "spade":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, spader, drawv2, rl.White)

				case "axe":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, axe, drawv2, rl.White)
				case "pumpaction":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, pumpaction, drawv2, rl.White)
				case "bombcollect":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, bomb1, drawv2, rl.White)
				case "chest":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, chest, drawv2, chestcolor)
					if player == drawblock || player+1 == drawblock || player-1 == drawblock || player+2 == drawblock || player-2 == drawblock {
						tri1 := rl.NewVector2(float32(drawx+15), float32(drawy-32))
						tri2 := rl.NewVector2(float32(drawx+8), float32(drawy-32))
						tri3 := rl.NewVector2(float32(drawx+8), float32(drawy))

						rl.DrawRectangle(drawx-10, drawy-50, 56, 18, rl.Fade(rl.White, 0.8))
						rl.DrawTriangle(tri1, tri2, tri3, rl.Fade(rl.White, 0.8))
						rl.DrawText("open me", drawx-2, drawy-46, 10, rl.Black)
					}

				case "chesthalf":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, chesthalf, drawv2, chestcolor)
					if framecount%6 == 0 {
						lootlayout[drawblock] = "chestopen"
					}
				case "chestopen":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					rl.DrawTextureRec(imgs, chestopen, drawv2, chestcolor)
				case "coin":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					drawv2shadow := rl.NewVector2(float32(drawx+1), float32(drawy+2))
					rl.DrawTextureRec(imgs, coin, drawv2shadow, rl.Black)
					rl.DrawTextureRec(imgs, coin, drawv2, rl.Fade(rl.Yellow, 0.5))
				case "gem":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy))
					drawv2shadow := rl.NewVector2(float32(drawx+1), float32(drawy+2))
					rl.DrawTextureRec(imgs, gem, drawv2shadow, rl.Black)
					rl.DrawTextureRec(imgs, gem, drawv2, rl.Fade(gemcolor, 0.7))
				}

				// MARK: draw shop items
				switch checkshopitems {
				case "random7":
					// shop item
					checkitem := shopitemsstructs[drawblock]
					originv2 := rl.NewVector2(9, 9)
					destrec := rl.NewRectangle(float32(drawx), float32((drawy)), 18, 18)
					switch checkitem.rotation {
					case 1:
						rl.DrawTexturePro(imgs, random7, destrec, originv2, shopitemrotation1, rl.White)
						if rolldice() < 3 {
							rl.DrawTexturePro(imgs, random7, destrec, originv2, shopitemrotation1, rl.Fade(rl.Magenta, 0.7))
						} else if rolldice() > 4 {
							rl.DrawTexturePro(imgs, random7, destrec, originv2, shopitemrotation1, rl.Fade(rl.Purple, 0.7))
						}
					case 2:
						rl.DrawTexturePro(imgs, random7, destrec, originv2, shopitemrotation2, rl.White)
						if rolldice() < 3 {
							rl.DrawTexturePro(imgs, random7, destrec, originv2, shopitemrotation2, rl.Fade(rl.Magenta, 0.7))
						} else if rolldice() > 4 {
							rl.DrawTexturePro(imgs, random7, destrec, originv2, shopitemrotation2, rl.Fade(rl.Purple, 0.7))
						}
					case 3:
						rl.DrawTexturePro(imgs, random7, destrec, originv2, shopitemrotation3, rl.White)
						if rolldice() < 3 {
							rl.DrawTexturePro(imgs, random7, destrec, originv2, shopitemrotation3, rl.Fade(rl.Magenta, 0.7))
						} else if rolldice() > 4 {
							rl.DrawTexturePro(imgs, random7, destrec, originv2, shopitemrotation3, rl.Fade(rl.Purple, 0.7))
						}
					}
					// item price
					shopitemprice1TEXT := strconv.Itoa(checkitem.price)
					rl.DrawText(shopitemprice1TEXT, drawx-6, drawy-21, 10, rl.Black)
					rl.DrawText(shopitemprice1TEXT, drawx-5, drawy-22, 10, rl.White)
					if rolldice() < 3 {
						rl.DrawText(shopitemprice1TEXT, drawx-5, drawy-22, 10, rl.Fade(rl.Magenta, 0.7))
					} else if rolldice() > 4 {
						rl.DrawText(shopitemprice1TEXT, drawx-5, drawy-22, 10, rl.Fade(rl.Purple, 0.7))
					}
					// item coin
					originv2 = rl.NewVector2(0, 0)
					destrec = rl.NewRectangle(float32(drawx+6), float32((drawy - 22)), 10, 10)
					rl.DrawTexturePro(imgs, coin, destrec, originv2, 0, rl.Fade(rl.Yellow, 0.5))
					// item marker arrow
					drawv2 := rl.NewVector2(float32(drawx-21), float32(drawy-23))
					drawv2shadow := rl.NewVector2(float32(drawx-20), float32(drawy-22))
					rl.DrawTextureRec(imgs, shopermarkerr, drawv2shadow, rl.Black)
					rl.DrawTextureRec(imgs, shopermarkerr, drawv2, rl.White)
					if rolldice() < 3 {
						rl.DrawTextureRec(imgs, shopermarkerr, drawv2, rl.Fade(rl.Magenta, 0.7))
					} else if rolldice() > 4 {
						rl.DrawTextureRec(imgs, shopermarkerr, drawv2, rl.Fade(rl.Purple, 0.7))
					}
					if player == drawblock {
						rl.DrawRectangle(drawx-24, drawy-40, 44, 14, rl.Fade(rl.White, 0.8))
						rl.DrawText("random", drawx-20, drawy-38, 10, rl.Black)
						if rl.IsKeyPressed(rl.KeyLeftControl) {
							newpowerup := powerup{}
							newpowerup.active = true
							newpowerup.name = "random7"
							newpowerup.img = random7
							newpowerup.number = 1
							powerupslayout[drawblock] = "random7"
							powerupsstructslayout[drawblock] = newpowerup
							collectpowerup(drawblock)
							shopitemslayout[drawblock] = ""
							coinstotal -= checkitem.price
						}
					}
				case "warp":
					// shop item
					checkitem := shopitemsstructs[drawblock]
					originv2 := rl.NewVector2(9, 9)
					destrec := rl.NewRectangle(float32(drawx), float32((drawy)), 18, 18)
					switch checkitem.rotation {
					case 1:
						rl.DrawTexturePro(imgs, warp, destrec, originv2, shopitemrotation1, rl.White)
						if rolldice() < 3 {
							rl.DrawTexturePro(imgs, warp, destrec, originv2, shopitemrotation1, rl.Fade(rl.Magenta, 0.7))
						} else if rolldice() > 4 {
							rl.DrawTexturePro(imgs, warp, destrec, originv2, shopitemrotation1, rl.Fade(rl.Purple, 0.7))
						}
					case 2:
						rl.DrawTexturePro(imgs, warp, destrec, originv2, shopitemrotation2, rl.White)
						if rolldice() < 3 {
							rl.DrawTexturePro(imgs, warp, destrec, originv2, shopitemrotation2, rl.Fade(rl.Magenta, 0.7))
						} else if rolldice() > 4 {
							rl.DrawTexturePro(imgs, warp, destrec, originv2, shopitemrotation2, rl.Fade(rl.Purple, 0.7))
						}
					case 3:
						rl.DrawTexturePro(imgs, warp, destrec, originv2, shopitemrotation3, rl.White)
						if rolldice() < 3 {
							rl.DrawTexturePro(imgs, warp, destrec, originv2, shopitemrotation3, rl.Fade(rl.Magenta, 0.7))
						} else if rolldice() > 4 {
							rl.DrawTexturePro(imgs, warp, destrec, originv2, shopitemrotation3, rl.Fade(rl.Purple, 0.7))
						}
					}
					// item price
					shopitemprice1TEXT := strconv.Itoa(checkitem.price)
					rl.DrawText(shopitemprice1TEXT, drawx-6, drawy-21, 10, rl.Black)
					rl.DrawText(shopitemprice1TEXT, drawx-5, drawy-22, 10, rl.White)
					if rolldice() < 3 {
						rl.DrawText(shopitemprice1TEXT, drawx-5, drawy-22, 10, rl.Fade(rl.Magenta, 0.7))
					} else if rolldice() > 4 {
						rl.DrawText(shopitemprice1TEXT, drawx-5, drawy-22, 10, rl.Fade(rl.Purple, 0.7))
					}
					// item coin
					originv2 = rl.NewVector2(0, 0)
					destrec = rl.NewRectangle(float32(drawx+6), float32((drawy - 22)), 10, 10)
					rl.DrawTexturePro(imgs, coin, destrec, originv2, 0, rl.Fade(rl.Yellow, 0.5))
					// item marker arrow
					drawv2 := rl.NewVector2(float32(drawx-21), float32(drawy-23))
					drawv2shadow := rl.NewVector2(float32(drawx-20), float32(drawy-22))
					rl.DrawTextureRec(imgs, shopermarkerr, drawv2shadow, rl.Black)
					rl.DrawTextureRec(imgs, shopermarkerr, drawv2, rl.White)
					if rolldice() < 3 {
						rl.DrawTextureRec(imgs, shopermarkerr, drawv2, rl.Fade(rl.Magenta, 0.7))
					} else if rolldice() > 4 {
						rl.DrawTextureRec(imgs, shopermarkerr, drawv2, rl.Fade(rl.Purple, 0.7))
					}
					if player == drawblock {
						rl.DrawRectangle(drawx-34, drawy-40, 50, 14, rl.Fade(rl.White, 0.8))
						rl.DrawText("teleport", drawx-30, drawy-38, 10, rl.Black)
						if rl.IsKeyPressed(rl.KeyLeftControl) {
							newpowerup := powerup{}
							newpowerup.active = true
							newpowerup.name = "warp"
							newpowerup.img = warp
							newpowerup.number = 1
							powerupslayout[drawblock] = "warp"
							powerupsstructslayout[drawblock] = newpowerup
							collectpowerup(drawblock)
							shopitemslayout[drawblock] = ""
							coinstotal -= checkitem.price
						}
					}
				case "ninjastar":
					// shop item
					checkitem := shopitemsstructs[drawblock]
					originv2 := rl.NewVector2(9, 9)
					destrec := rl.NewRectangle(float32(drawx), float32((drawy)), 18, 18)
					switch checkitem.rotation {
					case 1:
						rl.DrawTexturePro(imgs, ninjastar, destrec, originv2, shopitemrotation1, rl.White)
						if rolldice() < 3 {
							rl.DrawTexturePro(imgs, ninjastar, destrec, originv2, shopitemrotation1, rl.Fade(rl.Magenta, 0.7))
						} else if rolldice() > 4 {
							rl.DrawTexturePro(imgs, ninjastar, destrec, originv2, shopitemrotation1, rl.Fade(rl.Purple, 0.7))
						}
					case 2:
						rl.DrawTexturePro(imgs, ninjastar, destrec, originv2, shopitemrotation2, rl.White)
						if rolldice() < 3 {
							rl.DrawTexturePro(imgs, ninjastar, destrec, originv2, shopitemrotation2, rl.Fade(rl.Magenta, 0.7))
						} else if rolldice() > 4 {
							rl.DrawTexturePro(imgs, ninjastar, destrec, originv2, shopitemrotation2, rl.Fade(rl.Purple, 0.7))
						}
					case 3:
						rl.DrawTexturePro(imgs, ninjastar, destrec, originv2, shopitemrotation3, rl.White)
						if rolldice() < 3 {
							rl.DrawTexturePro(imgs, ninjastar, destrec, originv2, shopitemrotation3, rl.Fade(rl.Magenta, 0.7))
						} else if rolldice() > 4 {
							rl.DrawTexturePro(imgs, ninjastar, destrec, originv2, shopitemrotation3, rl.Fade(rl.Purple, 0.7))
						}
					}
					// item price
					shopitemprice1TEXT := strconv.Itoa(checkitem.price)
					rl.DrawText(shopitemprice1TEXT, drawx-6, drawy-21, 10, rl.Black)
					rl.DrawText(shopitemprice1TEXT, drawx-5, drawy-22, 10, rl.White)
					if rolldice() < 3 {
						rl.DrawText(shopitemprice1TEXT, drawx-5, drawy-22, 10, rl.Fade(rl.Magenta, 0.7))
					} else if rolldice() > 4 {
						rl.DrawText(shopitemprice1TEXT, drawx-5, drawy-22, 10, rl.Fade(rl.Purple, 0.7))
					}
					// item coin
					originv2 = rl.NewVector2(0, 0)
					destrec = rl.NewRectangle(float32(drawx+6), float32((drawy - 22)), 10, 10)
					rl.DrawTexturePro(imgs, coin, destrec, originv2, 0, rl.Fade(rl.Yellow, 0.5))
					// item marker arrow
					drawv2 := rl.NewVector2(float32(drawx-21), float32(drawy-23))
					drawv2shadow := rl.NewVector2(float32(drawx-20), float32(drawy-22))
					rl.DrawTextureRec(imgs, shopermarkerr, drawv2shadow, rl.Black)
					rl.DrawTextureRec(imgs, shopermarkerr, drawv2, rl.White)
					if rolldice() < 3 {
						rl.DrawTextureRec(imgs, shopermarkerr, drawv2, rl.Fade(rl.Magenta, 0.7))
					} else if rolldice() > 4 {
						rl.DrawTextureRec(imgs, shopermarkerr, drawv2, rl.Fade(rl.Purple, 0.7))
					}
					if player == drawblock {
						rl.DrawRectangle(drawx-34, drawy-40, 70, 14, rl.Fade(rl.White, 0.8))
						rl.DrawText("like a ninja", drawx-30, drawy-38, 10, rl.Black)
						if rl.IsKeyPressed(rl.KeyLeftControl) {
							newpowerup := powerup{}
							newpowerup.active = true
							newpowerup.name = "ninjastar"
							newpowerup.img = ninjastar
							newpowerup.number = 1
							powerupslayout[drawblock] = "ninjastar"
							powerupsstructslayout[drawblock] = newpowerup
							collectpowerup(drawblock)
							shopitemslayout[drawblock] = ""
							coinstotal -= checkitem.price
						}
					}
				case "arrow":
					// shop item
					checkitem := shopitemsstructs[drawblock]
					originv2 := rl.NewVector2(9, 9)
					destrec := rl.NewRectangle(float32(drawx), float32((drawy)), 18, 18)
					switch checkitem.rotation {
					case 1:
						rl.DrawTexturePro(imgs, arrowrandom, destrec, originv2, shopitemrotation1, rl.White)
						if rolldice() < 3 {
							rl.DrawTexturePro(imgs, arrowrandom, destrec, originv2, shopitemrotation1, rl.Fade(rl.Magenta, 0.7))
						} else if rolldice() > 4 {
							rl.DrawTexturePro(imgs, arrowrandom, destrec, originv2, shopitemrotation1, rl.Fade(rl.Purple, 0.7))
						}
					case 2:
						rl.DrawTexturePro(imgs, arrowrandom, destrec, originv2, shopitemrotation2, rl.White)
						if rolldice() < 3 {
							rl.DrawTexturePro(imgs, arrowrandom, destrec, originv2, shopitemrotation2, rl.Fade(rl.Magenta, 0.7))
						} else if rolldice() > 4 {
							rl.DrawTexturePro(imgs, arrowrandom, destrec, originv2, shopitemrotation2, rl.Fade(rl.Purple, 0.7))
						}
					case 3:
						rl.DrawTexturePro(imgs, arrowrandom, destrec, originv2, shopitemrotation3, rl.White)
						if rolldice() < 3 {
							rl.DrawTexturePro(imgs, arrowrandom, destrec, originv2, shopitemrotation3, rl.Fade(rl.Magenta, 0.7))
						} else if rolldice() > 4 {
							rl.DrawTexturePro(imgs, arrowrandom, destrec, originv2, shopitemrotation3, rl.Fade(rl.Purple, 0.7))
						}
					}
					// item price
					shopitemprice1TEXT := strconv.Itoa(checkitem.price)
					rl.DrawText(shopitemprice1TEXT, drawx-6, drawy-21, 10, rl.Black)
					rl.DrawText(shopitemprice1TEXT, drawx-5, drawy-22, 10, rl.White)
					if rolldice() < 3 {
						rl.DrawText(shopitemprice1TEXT, drawx-5, drawy-22, 10, rl.Fade(rl.Magenta, 0.7))
					} else if rolldice() > 4 {
						rl.DrawText(shopitemprice1TEXT, drawx-5, drawy-22, 10, rl.Fade(rl.Purple, 0.7))
					}
					// item coin
					originv2 = rl.NewVector2(0, 0)
					destrec = rl.NewRectangle(float32(drawx+6), float32((drawy - 22)), 10, 10)
					rl.DrawTexturePro(imgs, coin, destrec, originv2, 0, rl.Fade(rl.Yellow, 0.5))
					// item marker arrow
					drawv2 := rl.NewVector2(float32(drawx-21), float32(drawy-23))
					drawv2shadow := rl.NewVector2(float32(drawx-20), float32(drawy-22))
					rl.DrawTextureRec(imgs, shopermarkerr, drawv2shadow, rl.Black)
					rl.DrawTextureRec(imgs, shopermarkerr, drawv2, rl.White)
					if rolldice() < 3 {
						rl.DrawTextureRec(imgs, shopermarkerr, drawv2, rl.Fade(rl.Magenta, 0.7))
					} else if rolldice() > 4 {
						rl.DrawTextureRec(imgs, shopermarkerr, drawv2, rl.Fade(rl.Purple, 0.7))
					}
					if player == drawblock {
						rl.DrawRectangle(drawx-34, drawy-40, 62, 14, rl.Fade(rl.White, 0.8))
						rl.DrawText("robin hood", drawx-30, drawy-38, 10, rl.Black)
						if rl.IsKeyPressed(rl.KeyLeftControl) {
							newpowerup := powerup{}
							newpowerup.active = true
							newpowerup.name = "arrow"
							newpowerup.img = arrowrandom
							newpowerup.number = 1
							powerupslayout[drawblock] = "arrow"
							powerupsstructslayout[drawblock] = newpowerup
							collectpowerup(drawblock)
							shopitemslayout[drawblock] = ""
							coinstotal -= checkitem.price
						}
					}
				case "zippo":
					// shop item
					checkitem := shopitemsstructs[drawblock]
					originv2 := rl.NewVector2(9, 9)
					destrec := rl.NewRectangle(float32(drawx), float32((drawy)), 18, 18)
					switch checkitem.rotation {
					case 1:
						rl.DrawTexturePro(imgs, zippoburning, destrec, originv2, shopitemrotation1, rl.White)
						if rolldice() < 3 {
							rl.DrawTexturePro(imgs, zippoburning, destrec, originv2, shopitemrotation1, rl.Fade(rl.Magenta, 0.7))
						} else if rolldice() > 4 {
							rl.DrawTexturePro(imgs, zippoburning, destrec, originv2, shopitemrotation1, rl.Fade(rl.Purple, 0.7))
						}
					case 2:
						rl.DrawTexturePro(imgs, zippoburning, destrec, originv2, shopitemrotation2, rl.White)
						if rolldice() < 3 {
							rl.DrawTexturePro(imgs, zippoburning, destrec, originv2, shopitemrotation2, rl.Fade(rl.Magenta, 0.7))
						} else if rolldice() > 4 {
							rl.DrawTexturePro(imgs, zippoburning, destrec, originv2, shopitemrotation2, rl.Fade(rl.Purple, 0.7))
						}
					case 3:
						rl.DrawTexturePro(imgs, zippoburning, destrec, originv2, shopitemrotation3, rl.White)
						if rolldice() < 3 {
							rl.DrawTexturePro(imgs, zippoburning, destrec, originv2, shopitemrotation3, rl.Fade(rl.Magenta, 0.7))
						} else if rolldice() > 4 {
							rl.DrawTexturePro(imgs, zippoburning, destrec, originv2, shopitemrotation3, rl.Fade(rl.Purple, 0.7))
						}
					}
					// item price
					shopitemprice1TEXT := strconv.Itoa(checkitem.price)
					rl.DrawText(shopitemprice1TEXT, drawx-6, drawy-21, 10, rl.Black)
					rl.DrawText(shopitemprice1TEXT, drawx-5, drawy-22, 10, rl.White)
					if rolldice() < 3 {
						rl.DrawText(shopitemprice1TEXT, drawx-5, drawy-22, 10, rl.Fade(rl.Magenta, 0.7))
					} else if rolldice() > 4 {
						rl.DrawText(shopitemprice1TEXT, drawx-5, drawy-22, 10, rl.Fade(rl.Purple, 0.7))
					}
					// item coin
					originv2 = rl.NewVector2(0, 0)
					destrec = rl.NewRectangle(float32(drawx+6), float32((drawy - 22)), 10, 10)
					rl.DrawTexturePro(imgs, coin, destrec, originv2, 0, rl.Fade(rl.Yellow, 0.5))
					// item marker arrow
					drawv2 := rl.NewVector2(float32(drawx-21), float32(drawy-23))
					drawv2shadow := rl.NewVector2(float32(drawx-20), float32(drawy-22))
					rl.DrawTextureRec(imgs, shopermarkerr, drawv2shadow, rl.Black)
					rl.DrawTextureRec(imgs, shopermarkerr, drawv2, rl.White)
					if rolldice() < 3 {
						rl.DrawTextureRec(imgs, shopermarkerr, drawv2, rl.Fade(rl.Magenta, 0.7))
					} else if rolldice() > 4 {
						rl.DrawTextureRec(imgs, shopermarkerr, drawv2, rl.Fade(rl.Purple, 0.7))
					}
					if player == drawblock {
						rl.DrawRectangle(drawx-24, drawy-40, 40, 14, rl.Fade(rl.White, 0.8))
						rl.DrawText("flames", drawx-20, drawy-38, 10, rl.Black)
						if rl.IsKeyPressed(rl.KeyLeftControl) {
							newpowerup := powerup{}
							newpowerup.active = true
							newpowerup.name = "zippo"
							newpowerup.img = zippoburning
							newpowerup.number = 1
							powerupslayout[drawblock] = "zippo"
							powerupsstructslayout[drawblock] = newpowerup
							flameactive = true
							collectpowerup(drawblock)
							shopitemslayout[drawblock] = ""
							coinstotal -= checkitem.price
						}
					}
				case "tnt":
					// shop item
					checkitem := shopitemsstructs[drawblock]
					originv2 := rl.NewVector2(9, 9)
					destrec := rl.NewRectangle(float32(drawx), float32((drawy)), 18, 18)
					switch checkitem.rotation {
					case 1:
						rl.DrawTexturePro(imgs, tntenemiesexplode, destrec, originv2, shopitemrotation1, rl.White)
						if rolldice() < 3 {
							rl.DrawTexturePro(imgs, tntenemiesexplode, destrec, originv2, shopitemrotation1, rl.Fade(rl.Magenta, 0.7))
						} else if rolldice() > 4 {
							rl.DrawTexturePro(imgs, tntenemiesexplode, destrec, originv2, shopitemrotation1, rl.Fade(rl.Purple, 0.7))
						}
					case 2:
						rl.DrawTexturePro(imgs, tntenemiesexplode, destrec, originv2, shopitemrotation2, rl.White)
						if rolldice() < 3 {
							rl.DrawTexturePro(imgs, tntenemiesexplode, destrec, originv2, shopitemrotation2, rl.Fade(rl.Magenta, 0.7))
						} else if rolldice() > 4 {
							rl.DrawTexturePro(imgs, tntenemiesexplode, destrec, originv2, shopitemrotation2, rl.Fade(rl.Purple, 0.7))
						}
					case 3:
						rl.DrawTexturePro(imgs, tntenemiesexplode, destrec, originv2, shopitemrotation3, rl.White)
						if rolldice() < 3 {
							rl.DrawTexturePro(imgs, tntenemiesexplode, destrec, originv2, shopitemrotation3, rl.Fade(rl.Magenta, 0.7))
						} else if rolldice() > 4 {
							rl.DrawTexturePro(imgs, tntenemiesexplode, destrec, originv2, shopitemrotation3, rl.Fade(rl.Purple, 0.7))
						}
					}
					// item price
					shopitemprice1TEXT := strconv.Itoa(checkitem.price)
					rl.DrawText(shopitemprice1TEXT, drawx-6, drawy-21, 10, rl.Black)
					rl.DrawText(shopitemprice1TEXT, drawx-5, drawy-22, 10, rl.White)
					if rolldice() < 3 {
						rl.DrawText(shopitemprice1TEXT, drawx-5, drawy-22, 10, rl.Fade(rl.Magenta, 0.7))
					} else if rolldice() > 4 {
						rl.DrawText(shopitemprice1TEXT, drawx-5, drawy-22, 10, rl.Fade(rl.Purple, 0.7))
					}
					// item coin
					originv2 = rl.NewVector2(0, 0)
					destrec = rl.NewRectangle(float32(drawx+6), float32((drawy - 22)), 10, 10)
					rl.DrawTexturePro(imgs, coin, destrec, originv2, 0, rl.Fade(rl.Yellow, 0.5))
					// item marker arrow
					drawv2 := rl.NewVector2(float32(drawx-21), float32(drawy-23))
					drawv2shadow := rl.NewVector2(float32(drawx-20), float32(drawy-22))
					rl.DrawTextureRec(imgs, shopermarkerr, drawv2shadow, rl.Black)
					rl.DrawTextureRec(imgs, shopermarkerr, drawv2, rl.White)
					if rolldice() < 3 {
						rl.DrawTextureRec(imgs, shopermarkerr, drawv2, rl.Fade(rl.Magenta, 0.7))
					} else if rolldice() > 4 {
						rl.DrawTextureRec(imgs, shopermarkerr, drawv2, rl.Fade(rl.Purple, 0.7))
					}
					if player == drawblock {
						rl.DrawRectangle(drawx-44, drawy-40, 90, 14, rl.Fade(rl.White, 0.8))
						rl.DrawText("enemies explode", drawx-40, drawy-38, 10, rl.Black)
						if rl.IsKeyPressed(rl.KeyLeftControl) {
							newpowerup := powerup{}
							newpowerup.active = true
							newpowerup.name = "tnt"
							newpowerup.img = tntenemiesexplode
							newpowerup.number = 1
							powerupslayout[drawblock] = "tnt"
							powerupsstructslayout[drawblock] = newpowerup
							tntactive = true
							collectpowerup(drawblock)
							shopitemslayout[drawblock] = ""
							coinstotal -= checkitem.price
						}
					}
				case "coffee":
					// shop item
					checkitem := shopitemsstructs[drawblock]
					originv2 := rl.NewVector2(9, 9)
					destrec := rl.NewRectangle(float32(drawx), float32((drawy)), 18, 18)
					switch checkitem.rotation {
					case 1:
						rl.DrawTexturePro(imgs, caffeinejump, destrec, originv2, shopitemrotation1, rl.White)
						if rolldice() < 3 {
							rl.DrawTexturePro(imgs, caffeinejump, destrec, originv2, shopitemrotation1, rl.Fade(rl.Magenta, 0.7))
						} else if rolldice() > 4 {
							rl.DrawTexturePro(imgs, caffeinejump, destrec, originv2, shopitemrotation1, rl.Fade(rl.Purple, 0.7))
						}
					case 2:
						rl.DrawTexturePro(imgs, caffeinejump, destrec, originv2, shopitemrotation2, rl.White)
						if rolldice() < 3 {
							rl.DrawTexturePro(imgs, caffeinejump, destrec, originv2, shopitemrotation2, rl.Fade(rl.Magenta, 0.7))
						} else if rolldice() > 4 {
							rl.DrawTexturePro(imgs, caffeinejump, destrec, originv2, shopitemrotation2, rl.Fade(rl.Purple, 0.7))
						}
					case 3:
						rl.DrawTexturePro(imgs, caffeinejump, destrec, originv2, shopitemrotation3, rl.White)
						if rolldice() < 3 {
							rl.DrawTexturePro(imgs, caffeinejump, destrec, originv2, shopitemrotation3, rl.Fade(rl.Magenta, 0.7))
						} else if rolldice() > 4 {
							rl.DrawTexturePro(imgs, caffeinejump, destrec, originv2, shopitemrotation3, rl.Fade(rl.Purple, 0.7))
						}
					}
					// item price
					shopitemprice1TEXT := strconv.Itoa(checkitem.price)
					rl.DrawText(shopitemprice1TEXT, drawx-6, drawy-21, 10, rl.Black)
					rl.DrawText(shopitemprice1TEXT, drawx-5, drawy-22, 10, rl.White)
					if rolldice() < 3 {
						rl.DrawText(shopitemprice1TEXT, drawx-5, drawy-22, 10, rl.Fade(rl.Magenta, 0.7))
					} else if rolldice() > 4 {
						rl.DrawText(shopitemprice1TEXT, drawx-5, drawy-22, 10, rl.Fade(rl.Purple, 0.7))
					}
					// item coin
					originv2 = rl.NewVector2(0, 0)
					destrec = rl.NewRectangle(float32(drawx+6), float32((drawy - 22)), 10, 10)
					rl.DrawTexturePro(imgs, coin, destrec, originv2, 0, rl.Fade(rl.Yellow, 0.5))
					// item marker arrow
					drawv2 := rl.NewVector2(float32(drawx-21), float32(drawy-23))
					drawv2shadow := rl.NewVector2(float32(drawx-20), float32(drawy-22))
					rl.DrawTextureRec(imgs, shopermarkerr, drawv2shadow, rl.Black)
					rl.DrawTextureRec(imgs, shopermarkerr, drawv2, rl.White)
					if rolldice() < 3 {
						rl.DrawTextureRec(imgs, shopermarkerr, drawv2, rl.Fade(rl.Magenta, 0.7))
					} else if rolldice() > 4 {
						rl.DrawTextureRec(imgs, shopermarkerr, drawv2, rl.Fade(rl.Purple, 0.7))
					}
					if player == drawblock {
						rl.DrawRectangle(drawx-24, drawy-40, 42, 14, rl.Fade(rl.White, 0.8))
						rl.DrawText("jump +1", drawx-20, drawy-38, 10, rl.Black)
						if rl.IsKeyPressed(rl.KeyLeftControl) {
							newpowerup := powerup{}
							newpowerup.active = true
							newpowerup.name = "coffee"
							newpowerup.img = caffeinejump
							newpowerup.number = 1
							powerupslayout[drawblock] = "coffee"
							powerupsstructslayout[drawblock] = newpowerup
							collectpowerup(drawblock)
							shopitemslayout[drawblock] = ""
							coinstotal -= checkitem.price
						}
					}
				case "colorize":
					// shop item
					checkitem := shopitemsstructs[drawblock]
					originv2 := rl.NewVector2(9, 9)
					destrec := rl.NewRectangle(float32(drawx), float32((drawy)), 18, 18)
					switch checkitem.rotation {
					case 1:
						rl.DrawTexturePro(imgs, colorize, destrec, originv2, shopitemrotation1, rl.White)
						if rolldice() < 3 {
							rl.DrawTexturePro(imgs, colorize, destrec, originv2, shopitemrotation1, rl.Fade(rl.Magenta, 0.7))
						} else if rolldice() > 4 {
							rl.DrawTexturePro(imgs, colorize, destrec, originv2, shopitemrotation1, rl.Fade(rl.Purple, 0.7))
						}
					case 2:
						rl.DrawTexturePro(imgs, colorize, destrec, originv2, shopitemrotation2, rl.White)
						if rolldice() < 3 {
							rl.DrawTexturePro(imgs, colorize, destrec, originv2, shopitemrotation2, rl.Fade(rl.Magenta, 0.7))
						} else if rolldice() > 4 {
							rl.DrawTexturePro(imgs, colorize, destrec, originv2, shopitemrotation2, rl.Fade(rl.Purple, 0.7))
						}
					case 3:
						rl.DrawTexturePro(imgs, colorize, destrec, originv2, shopitemrotation3, rl.White)
						if rolldice() < 3 {
							rl.DrawTexturePro(imgs, colorize, destrec, originv2, shopitemrotation3, rl.Fade(rl.Magenta, 0.7))
						} else if rolldice() > 4 {
							rl.DrawTexturePro(imgs, colorize, destrec, originv2, shopitemrotation3, rl.Fade(rl.Purple, 0.7))
						}
					}
					// item price
					shopitemprice1TEXT := strconv.Itoa(checkitem.price)
					rl.DrawText(shopitemprice1TEXT, drawx-6, drawy-21, 10, rl.Black)
					rl.DrawText(shopitemprice1TEXT, drawx-5, drawy-22, 10, rl.White)
					if rolldice() < 3 {
						rl.DrawText(shopitemprice1TEXT, drawx-5, drawy-22, 10, rl.Fade(rl.Magenta, 0.7))
					} else if rolldice() > 4 {
						rl.DrawText(shopitemprice1TEXT, drawx-5, drawy-22, 10, rl.Fade(rl.Purple, 0.7))
					}
					// item coin
					originv2 = rl.NewVector2(0, 0)
					destrec = rl.NewRectangle(float32(drawx+6), float32((drawy - 22)), 10, 10)
					rl.DrawTexturePro(imgs, coin, destrec, originv2, 0, rl.Fade(rl.Yellow, 0.5))
					// item marker arrow
					drawv2 := rl.NewVector2(float32(drawx-21), float32(drawy-23))
					drawv2shadow := rl.NewVector2(float32(drawx-20), float32(drawy-22))
					rl.DrawTextureRec(imgs, shopermarkerr, drawv2shadow, rl.Black)
					rl.DrawTextureRec(imgs, shopermarkerr, drawv2, rl.White)
					if rolldice() < 3 {
						rl.DrawTextureRec(imgs, shopermarkerr, drawv2, rl.Fade(rl.Magenta, 0.7))
					} else if rolldice() > 4 {
						rl.DrawTextureRec(imgs, shopermarkerr, drawv2, rl.Fade(rl.Purple, 0.7))
					}
					if player == drawblock {
						rl.DrawRectangle(drawx-44, drawy-40, 87, 14, rl.Fade(rl.White, 0.8))
						rl.DrawText("color the world", drawx-40, drawy-38, 10, rl.Black)
						if rl.IsKeyPressed(rl.KeyLeftControl) {
							newpowerup := powerup{}
							newpowerup.active = true
							newpowerup.name = "colorize"
							newpowerup.img = colorize
							newpowerup.number = 1
							powerupslayout[drawblock] = "colorize"
							powerupsstructslayout[drawblock] = newpowerup
							collectpowerup(drawblock)
							shopitemslayout[drawblock] = ""
							coinstotal -= checkitem.price
						}
					}
				}

				// MARK: draw active quest
				if activequest.block == drawblock {
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy-4))
					rl.DrawTextureRec(imgs, activequest.img, drawv2, rl.White)
					rl.DrawTextureRec(imgs, activequest.img, drawv2, rl.Fade(rl.Magenta, hpfade))
					arrowv2 := rl.NewVector2(float32(drawx+4), float32(drawy-24-playerychange))
					rl.DrawTextureRec(imgs, arrowdown, arrowv2, rl.White)
					rl.DrawText("quest", drawx-4, drawy-40, 10, rl.White)
					if player == drawblock {
						if rl.IsKeyPressed(rl.KeyLeftControl) {
							questactive = false
							questdetails = false
							activequest = questitem{}
							questtotal++
						}
					}
				}

				// MARK: draw end key
				if endkeyblock == drawblock {
					switch endkeynumber {
					case 1:
						drawv2 := rl.NewVector2(float32(drawx), float32(drawy-3))
						rl.DrawTextureRec(imgs, key1, drawv2, rl.White)
					case 2:
						drawv2 := rl.NewVector2(float32(drawx), float32(drawy-3))
						rl.DrawTextureRec(imgs, key2, drawv2, rl.White)
					case 3:
						drawv2 := rl.NewVector2(float32(drawx), float32(drawy-3))
						rl.DrawTextureRec(imgs, key3, drawv2, rl.White)
					case 4:
						drawv2 := rl.NewVector2(float32(drawx), float32(drawy-3))
						rl.DrawTextureRec(imgs, key4, drawv2, rl.White)
					case 5:
						drawv2 := rl.NewVector2(float32(drawx), float32(drawy-3))
						rl.DrawTextureRec(imgs, key5, drawv2, rl.White)
					case 6:
						drawv2 := rl.NewVector2(float32(drawx), float32(drawy-3))
						rl.DrawTextureRec(imgs, key6, drawv2, rl.White)
					case 7:
						drawv2 := rl.NewVector2(float32(drawx), float32(drawy-3))
						rl.DrawTextureRec(imgs, key7, drawv2, rl.White)
					case 8:
						drawv2 := rl.NewVector2(float32(drawx), float32(drawy-3))
						rl.DrawTextureRec(imgs, key8, drawv2, rl.White)
					case 9:
						drawv2 := rl.NewVector2(float32(drawx), float32(drawy-3))
						rl.DrawTextureRec(imgs, key9, drawv2, rl.White)
					}

					if rl.IsKeyPressed(rl.KeySpace) {
						endkeyactive = false
						endkeycollected = true
						endkeyblock = 0
					}

				}

				drawx += blockw
				count++
				drawblock++

				if count == draww {
					count = 0
					drawblock += levelw - draww
					drawy += blockh
					drawx = 0
				}
			}

			// MARK: draw player layer
			drawy = 0
			count = 0
			drawblock = drawblocknext
			for a := 0; a < drawarea; a++ {

				// MARK: draw railway 2
				checkrailway := railwaylayout[drawblock]
				switch checkrailway {
				case "cartr", "cartl", "cart":
					drawv2 := rl.NewVector2(float32(drawx), float32(drawy+2))
					rl.DrawTextureRec(imgs, cart, drawv2, rl.Gold)
				}

				if flossendblock == drawblock {
					flossx = drawx
					flossy = drawy
				}

				if drawblock == player {

					// water bubbles
					if waterlayout[player] != "" {
						if !underwateron {
							underwateron = true
						}
						rl.DrawCircle(drawx+rInt(4, 13), (drawy)-bubbley, rFloat32(1, 4), rl.Fade(rl.Aqua, 0.3))
						rl.DrawCircle(drawx+rInt(4, 13), (drawy-10)-bubbley, rFloat32(3, 6), rl.Fade(rl.Aqua, 0.3))
						if framecount%2 == 0 {
							bubbley++
						}
						if bubbley > 12 {
							bubbley = 0
						}
					}
					// draw umbrella
					if fallactive && jumppause == 0 && !flosson && levellayout[player+levelw] == "." {
						drawv2 := rl.NewVector2(float32(drawx), float32(drawy-16))
						rl.DrawTextureRec(imgs, umbrella, drawv2, rl.White)
						rl.DrawTextureRec(imgs, umbrella, drawv2, rl.Fade(rl.Magenta, playerfade))
					}
					// MARK: draw player weapon
					equippedweapon = true
					if equippedweapon {

						if pickaxeactive {
							if playerdirection == 2 {
								originv2 := rl.NewVector2(-13, 5)
								destrec := rl.NewRectangle(float32(drawx), float32((drawy)-playerychange), 16, 16)
								if swingrighton {
									originv2 = rl.NewVector2(-16, 18)
									// dust
									drawv2 := rl.NewVector2(float32(drawx-16), float32(drawy))
									rl.DrawTextureRec(imgs, dust, drawv2, rl.Fade(rl.White, 0.5))
									drawv2 = rl.NewVector2(float32(drawx-32), float32(drawy))
									rl.DrawTextureRec(imgs, dust, drawv2, rl.Fade(rl.White, 0.5))
								}
								rl.DrawTexturePro(imgs, pickaraxer, destrec, originv2, weaponrotation, rl.Green)
							}
							if playerdirection == 4 {
								originv2 := rl.NewVector2(13, 5)
								destrec := rl.NewRectangle(float32(drawx), float32((drawy)-playerychange), 16, 16)
								if swinglefton {
									originv2 = rl.NewVector2(22, 7)
									// dust
									drawv2 := rl.NewVector2(float32(drawx+16), float32(drawy))
									rl.DrawTextureRec(imgs, dust, drawv2, rl.Fade(rl.White, 0.5))
									drawv2 = rl.NewVector2(float32(drawx+32), float32(drawy))
									rl.DrawTextureRec(imgs, dust, drawv2, rl.Fade(rl.White, 0.5))
								}
								rl.DrawTexturePro(imgs, pickaraxel, destrec, originv2, weaponrotation, rl.Green)
							}

						} else if spadeactive {

							if playerdirection == 2 {
								originv2 := rl.NewVector2(-12, -6)
								destrec := rl.NewRectangle(float32(drawx), float32((drawy)-playerychange), 12, 12)
								rotation := float32(0)
								if swingrighton {
									originv2 = rl.NewVector2(10, -12)
									rotation = -75
								}
								rl.DrawTexturePro(imgs, spader, destrec, originv2, rotation, rl.Green)
							}
							if playerdirection == 4 {
								originv2 := rl.NewVector2(8, -7)
								destrec := rl.NewRectangle(float32(drawx), float32((drawy)-playerychange), 12, 12)
								rotation := float32(0)
								if swinglefton {
									rotation = 75
									originv2 = rl.NewVector2(0, 0)
								}
								rl.DrawTexturePro(imgs, spadel, destrec, originv2, rotation, rl.Green)
							}

							// dust
							if swingrighton || swinglefton {
								drawv2 := rl.NewVector2(float32(drawx), float32(drawy-16))
								rl.DrawTextureRec(imgs, dust, drawv2, rl.Fade(rl.White, 0.5))
								drawv2 = rl.NewVector2(float32(drawx), float32(drawy-32))
								rl.DrawTextureRec(imgs, dust, drawv2, rl.Fade(rl.White, 0.5))
							}
						} else if icewandactive {

							if playerdirection == 2 {
								originv2 := rl.NewVector2(-13, 5)
								destrec := rl.NewRectangle(float32(drawx), float32((drawy)-playerychange), 16, 16)
								if swingrighton {
									originv2 = rl.NewVector2(-16, 18)
								}
								rl.DrawTexturePro(imgs, icewandr, destrec, originv2, weaponrotation, rl.White)
							}
							if playerdirection == 4 {
								originv2 := rl.NewVector2(13, 5)
								destrec := rl.NewRectangle(float32(drawx), float32((drawy)-playerychange), 16, 16)
								if swinglefton {
									originv2 = rl.NewVector2(22, 7)
								}
								rl.DrawTexturePro(imgs, icewandl, destrec, originv2, weaponrotation, rl.White)
							}

						} else if bigswordactive {

							if playerdirection == 2 {
								originv2 := rl.NewVector2(-13, 5)
								destrec := rl.NewRectangle(float32(drawx), float32((drawy)-playerychange), 16, 16)
								if swingrighton {
									originv2 = rl.NewVector2(-16, 18)
								}
								rl.DrawTexturePro(imgs, bigswordr, destrec, originv2, weaponrotation, rl.White)
							}
							if playerdirection == 4 {
								originv2 := rl.NewVector2(13, 5)
								destrec := rl.NewRectangle(float32(drawx), float32((drawy)-playerychange), 16, 16)
								if swinglefton {
									originv2 = rl.NewVector2(22, 7)
								}
								rl.DrawTexturePro(imgs, bigswordl, destrec, originv2, weaponrotation, rl.White)
							}

						} else if uzziactive { // uzzi
							if playerdirection == 2 {
								originv2 := rl.NewVector2(-11, -2)
								destrec := rl.NewRectangle(float32(drawx), float32((drawy)-playerychange), 14, 12)
								rl.DrawTexturePro(imgs, uzzir, destrec, originv2, uzziangle, rl.White)
							}
							if playerdirection == 4 {
								originv2 := rl.NewVector2(10, -2)
								destrec := rl.NewRectangle(float32(drawx), float32((drawy)-playerychange), 14, 12)

								rl.DrawTexturePro(imgs, uzzil, destrec, originv2, uzziangle, rl.White)
							}
						} else if bazookaactive { // bazooka
							if playerdirection == 2 {
								drawv2 := rl.NewVector2(float32(drawx+11), float32((drawy+2)-playerychange))
								rl.DrawTextureRec(imgs, bazookal, drawv2, rl.White)
								// bazooka blast
								if rl.IsKeyPressed(rl.KeySpace) {
									blastv2 := rl.NewVector2(float32(drawx+15), float32(drawy-25))
									rl.DrawTextureRec(imgs, blast, blastv2, rl.Fade(rl.White, 0.5))
								}
							} else if playerdirection == 4 {
								drawv2 := rl.NewVector2(float32(drawx-11), float32((drawy+2)-playerychange))
								rl.DrawTextureRec(imgs, bazooka, drawv2, rl.White)
								// bazooka blast
								if rl.IsKeyPressed(rl.KeySpace) {
									blastv2 := rl.NewVector2(float32(drawx-80), float32(drawy-25))
									rl.DrawTextureRec(imgs, blast, blastv2, rl.Fade(rl.White, 0.5))
								}
							}

						} else if axeactive { // axe
							if playerdirection == 2 {
								originv2 := rl.NewVector2(-14, 5)
								destrec := rl.NewRectangle(float32(drawx), float32((drawy)-playerychange), 16, 16)
								if swingrighton {
									originv2 = rl.NewVector2(-18, 18)
									if flameactive {
										weaponslayout[player+1] = "flame4"
									}
								}
								rl.DrawTexturePro(imgs, axe, destrec, originv2, weaponrotation, rl.White)
							}
							if playerdirection == 4 {
								originv2 := rl.NewVector2(14, 5)
								destrec := rl.NewRectangle(float32(drawx), float32((drawy)-playerychange), 16, 16)
								if swinglefton {
									originv2 = rl.NewVector2(24, 7)
									if flameactive {
										weaponslayout[player-1] = "flame5"
									}
								}
								rl.DrawTexturePro(imgs, axel, destrec, originv2, weaponrotation, rl.White)
							}
						} else if pumpactionactive { // pump action
							if playerdirection == 2 {
								drawv2 := rl.NewVector2(float32(drawx+14), float32((drawy+4)-playerychange))
								rl.DrawTextureRec(imgs, pumpaction, drawv2, rl.White)

							}
							if playerdirection == 4 {
								drawv2 := rl.NewVector2(float32(drawx-14), float32((drawy+4)-playerychange))
								rl.DrawTextureRec(imgs, pumpactionl, drawv2, rl.White)

							}

						} else { // default sword
							if playerdirection == 2 {
								originv2 := rl.NewVector2(-15, 2)
								destrec := rl.NewRectangle(float32(drawx), float32((drawy)-playerychange), 12, 12)
								if swingrighton {
									originv2 = rl.NewVector2(-16, 16)
									if flameactive {
										weaponslayout[player+1] = "flame4"
									}
								}
								rl.DrawTexturePro(imgs, sword1, destrec, originv2, weaponrotation, rl.White)
							} else if playerdirection == 4 {
								originv2 := rl.NewVector2(11, 2)
								destrec := rl.NewRectangle(float32(drawx), float32((drawy)-playerychange), 12, 12)
								if swinglefton {
									originv2 = rl.NewVector2(18, 4)
									if flameactive {
										weaponslayout[player-1] = "flame5"
									}
								}
								rl.DrawTexturePro(imgs, sword1l, destrec, originv2, weaponrotation, rl.White)
							}
						}
					}
					// draw emote
					if emoteswitch && !flosson && !fallactive {
						emotev2 := rl.NewVector2(float32(drawx-8), float32((drawy-30)-playerychange))
						switch chooseemote {
						case 1:
							rl.DrawTextureRec(imgs, emote1, emotev2, rl.Fade(rl.Violet, emotefade))
						case 2:
							rl.DrawTextureRec(imgs, emote2, emotev2, rl.Fade(rl.Violet, emotefade))
						case 3:
							rl.DrawTextureRec(imgs, emote3, emotev2, rl.Fade(rl.Violet, emotefade))
						case 4:
							rl.DrawTextureRec(imgs, emote4, emotev2, rl.Fade(rl.Violet, emotefade))
						case 5:
							rl.DrawTextureRec(imgs, emote5, emotev2, rl.Fade(rl.Violet, emotefade))
						case 6:
							rl.DrawTextureRec(imgs, emote6, emotev2, rl.Fade(rl.Violet, emotefade))

						}
					}
					//	draw player
					if playerrollactive {

						originv2 := rl.NewVector2(8, 8)
						destrec := rl.NewRectangle(float32(drawx+8), float32(drawy+8), 16, 16)
						if playerdirection == 2 {
							rl.DrawTexturePro(imgs, playerright, destrec, originv2, playerrolldegrees, rl.White)
							rl.DrawTexturePro(imgs, playerright, destrec, originv2, playerrolldegrees, rl.Fade(rl.Magenta, playerfade))
						} else if playerdirection == 4 {
							rl.DrawTexturePro(imgs, playerleft, destrec, originv2, playerrolldegrees, rl.White)
							rl.DrawTexturePro(imgs, playerleft, destrec, originv2, playerrolldegrees, rl.Fade(rl.Magenta, playerfade))
						}

					} else {
						drawv2 := rl.NewVector2(float32(drawx), float32(drawy-playerychange))
						if playerdirection == 2 {
							rl.DrawTextureRec(imgs, playerright, drawv2, rl.White)
							rl.DrawTextureRec(imgs, playerright, drawv2, rl.Fade(rl.Magenta, playerfade))
							if playerrunon {
								rl.DrawPixel(drawx-rInt(3, 11), drawy+rInt(8, 16), rl.Magenta)
								rl.DrawPixel(drawx-rInt(3, 11), drawy+rInt(8, 16), rl.Magenta)
							}
						} else if playerdirection == 4 {
							rl.DrawTextureRec(imgs, playerleft, drawv2, rl.White)
							rl.DrawTextureRec(imgs, playerleft, drawv2, rl.Fade(rl.Magenta, playerfade))
							if playerrunon {
								rl.DrawPixel(drawx+rInt(19, 27), drawy+rInt(8, 16), rl.Magenta)
								rl.DrawPixel(drawx+rInt(19, 27), drawy+rInt(8, 16), rl.Magenta)
							}
						}
					}

					playerx = drawx
					playery = drawy
					if flosson {
						rl.DrawLine(playerx+9, playery+4, flossx+9, flossy+28, rl.Purple)
						rl.DrawLine(playerx+8, playery+4, flossx+8, flossy+28, rl.Pink)
						hookv2 := rl.NewVector2(float32(flossx+5), float32(flossy+12))
						rl.DrawTextureRec(imgs, hook, hookv2, rl.Magenta)
					}

					// draw shield
					if shieldon {
						rl.DrawCircleGradient(drawx+8, drawy+6, 24, rl.Transparent, rl.Fade(rl.Magenta, shieldfade))
					}
				}

				if count == draww {
					count = 0
					drawblock += levelw - draww
					drawy += blockh
					drawx = 0
				}

				drawx += blockw
				count++
				drawblock++

			}
		}

		rl.EndMode2D() // MARK: draw no camera

		if !inventoryon && !optionson && !introscreenon && !drawstarton {
			// inventory tab
			rl.DrawRectangle(inventoryx+600, inventoryy, 60, 40, rl.Fade(rl.White, 0.8))
			rl.DrawText("inv", inventoryx+607, inventoryy+5, 30, rl.Black)
			rl.DrawText("inv", inventoryx+608, inventoryy+4, 30, rl.Magenta)
			// inventory tab shadow
			rl.DrawRectangle(inventoryx+660, inventoryy+4, 4, 40, rl.Fade(rl.Black, 0.8))
			rl.DrawRectangle(inventoryx+600, inventoryy+40, 60, 4, rl.Fade(rl.Black, 0.8))
		}
		if icefreezeactive {

			switch rolldice() {
			case 1:
				rl.DrawRectangle(0, 0, monw, monh, rl.Fade(rl.DarkBlue, 0.1))
			case 2:
				rl.DrawRectangle(0, 0, monw, monh, rl.Fade(rl.Blue, 0.1))
			case 5:
				rl.DrawRectangle(0, 0, monw, monh, rl.Fade(rl.SkyBlue, 0.2))
			case 6:
				rl.DrawRectangle(0, 0, monw, monh, rl.Fade(rl.SkyBlue, 0.1))
			}

			if icefreezefadeon {
				icefreezefade += 0.05
				if icefreezefade >= 0.5 {
					icefreezefadeon = false
				}
			} else {
				icefreezefade -= 0.05
				if icefreezefade <= 0 {
					icefreezefadeon = true
				}
			}

			rl.BeginMode2D(camera4x)

			for a := 0; a < len(icefreezev2); a++ {
				rl.DrawTextureRec(imgs, freezeimg, icefreezev2[a], rl.Fade(rl.SkyBlue, icefreezefade))

				check := icefreezev2[a]
				check.X += rFloat32(-2, 3)
				check.Y += rFloat32(-2, 3)
				icefreezev2[a] = check
			}
			rl.EndMode2D()
		}

		if pumpactionactive || icefreezeactive {
			drawspecialtext()
		}

		update()

		if centerlineson {
			rl.DrawLine(monw/2, 0, monw/2, monh, rl.Fade(rl.Green, 0.2))
			rl.DrawLine(0, monh/2, monw, monh/2, rl.Fade(rl.Green, 0.2))
		}

		rl.EndDrawing()

	}

	rl.CloseWindow()

}
func drawselectscreen() { // MARK: drawselectscreen █████████████████████████████████████
	// background rec
	rl.DrawRectangle(0, 0, monw, monh, rl.Black)

	rl.DrawText("start a game", 47, 33, 40, rl.Magenta)
	rl.DrawText("start a game", 49, 31, 40, rl.Black)
	rl.DrawText("start a game", 50, 30, 40, rl.White)

	rl.DrawText("more options unlock", 50, 80, 30, rl.White)
	rl.DrawText("the more you play", 50, 110, 30, rl.White)

	// unlock boxes
	recx := (monw / 2) - 416
	lockedv2 := rl.NewVector2(float32(((monw/2)-416)/4), (350 / 4))
	for a := 0; a < 10; a++ {
		rl.DrawRectangle(recx, 350, 64, 64, rl.Fade(rl.Magenta, 0.2))
		rl.DrawRectangleLines(recx, 350, 64, 64, rl.Purple)
		if selectscreennumber-3 == a {

			rl.DrawText("locked", recx-18, 442, 30, rl.Magenta)
			rl.DrawText("locked", recx-17, 441, 30, rl.Black)
			rl.DrawText("locked", recx-16, 440, 30, rl.White)

			if onoff15 {
				rl.DrawRectangle(recx-16, 334, 96, 96, rl.Fade(rl.Magenta, 0.1))
				rl.DrawRectangleLines(recx-16, 334, 96, 96, rl.Fade(rl.Purple, 0.1))
				lockedv2.Y--
			} else {
				rl.DrawRectangle(recx-16, 334, 96, 96, rl.Fade(rl.Magenta, 0.2))
				rl.DrawRectangleLines(recx-16, 334, 96, 96, rl.Fade(rl.Purple, 0.2))
				lockedv2.Y++
			}
		}
		recx += 96
	}
	rl.BeginMode2D(camera4x)
	for a := 0; a < 10; a++ {
		rl.DrawTextureRec(imgs, lockedicon, lockedv2, rl.White)
		lockedv2.X += 24
	}
	rl.EndMode2D()

	// select keys
	if rl.IsKeyPressed(rl.KeyUp) {
		if selectscreennumber > 2 && selectscreennumber < 13 {
			selectscreennumber = 1
		} else if selectscreennumber == 13 {
			selectscreennumber = 3
		}
	}
	if rl.IsKeyPressed(rl.KeyDown) {
		if selectscreennumber < 3 {
			selectscreennumber = 3
		} else if selectscreennumber > 2 && selectscreennumber < 13 {
			selectscreennumber = 13
		}
	}
	if rl.IsKeyPressed(rl.KeyLeft) {
		selectscreennumber--
		if selectscreennumber < 1 {
			selectscreennumber = 1
		}
	}
	if rl.IsKeyPressed(rl.KeyRight) {
		selectscreennumber++
		if selectscreennumber > 13 {
			selectscreennumber = 13
		}
	}

	// start button
	if selectscreennumber == 13 {
		rl.DrawRectangle((monw/2)-150, 510, 300, 120, rl.Fade(rl.Magenta, 0.2))
	}
	rl.DrawRectangle((monw/2)-130, 530, 260, 80, rl.Magenta)
	rl.DrawText("start", (monw/2)-80, 540, 60, rl.Black)

	// draw select box locked text
	switch selectscreennumber {
	case 3:
		rl.DrawText("locked", monw-303, 33, 40, rl.Magenta)
		rl.DrawText("locked", monw-301, 31, 40, rl.Black)
		rl.DrawText("locked", monw-300, 30, 40, rl.White)
		rl.DrawText("some text here", monw-350, 80, 30, rl.White)
		rl.DrawText("some text here", monw-350, 110, 30, rl.White)
	case 4:
		rl.DrawText("locked", monw-303, 33, 40, rl.Magenta)
		rl.DrawText("locked", monw-301, 31, 40, rl.Black)
		rl.DrawText("locked", monw-300, 30, 40, rl.White)
		rl.DrawText("some text here", monw-350, 80, 30, rl.White)
		rl.DrawText("some text here", monw-350, 110, 30, rl.White)
	case 5:
		rl.DrawText("locked", monw-303, 33, 40, rl.Magenta)
		rl.DrawText("locked", monw-301, 31, 40, rl.Black)
		rl.DrawText("locked", monw-300, 30, 40, rl.White)
		rl.DrawText("some text here", monw-350, 80, 30, rl.White)
		rl.DrawText("some text here", monw-350, 110, 30, rl.White)
	case 6:
		rl.DrawText("locked", monw-303, 33, 40, rl.Magenta)
		rl.DrawText("locked", monw-301, 31, 40, rl.Black)
		rl.DrawText("locked", monw-300, 30, 40, rl.White)
		rl.DrawText("some text here", monw-350, 80, 30, rl.White)
		rl.DrawText("some text here", monw-350, 110, 30, rl.White)
	case 7:
		rl.DrawText("locked", monw-303, 33, 40, rl.Magenta)
		rl.DrawText("locked", monw-301, 31, 40, rl.Black)
		rl.DrawText("locked", monw-300, 30, 40, rl.White)
		rl.DrawText("some text here", monw-350, 80, 30, rl.White)
		rl.DrawText("some text here", monw-350, 110, 30, rl.White)
	case 8:
		rl.DrawText("locked", monw-303, 33, 40, rl.Magenta)
		rl.DrawText("locked", monw-301, 31, 40, rl.Black)
		rl.DrawText("locked", monw-300, 30, 40, rl.White)
		rl.DrawText("some text here", monw-350, 80, 30, rl.White)
		rl.DrawText("some text here", monw-350, 110, 30, rl.White)
	case 9:
		rl.DrawText("locked", monw-303, 33, 40, rl.Magenta)
		rl.DrawText("locked", monw-301, 31, 40, rl.Black)
		rl.DrawText("locked", monw-300, 30, 40, rl.White)
		rl.DrawText("some text here", monw-350, 80, 30, rl.White)
		rl.DrawText("some text here", monw-350, 110, 30, rl.White)
	case 10:
		rl.DrawText("locked", monw-303, 33, 40, rl.Magenta)
		rl.DrawText("locked", monw-301, 31, 40, rl.Black)
		rl.DrawText("locked", monw-300, 30, 40, rl.White)
		rl.DrawText("some text here", monw-350, 80, 30, rl.White)
		rl.DrawText("some text here", monw-350, 110, 30, rl.White)
	case 11:
		rl.DrawText("locked", monw-303, 33, 40, rl.Magenta)
		rl.DrawText("locked", monw-301, 31, 40, rl.Black)
		rl.DrawText("locked", monw-300, 30, 40, rl.White)
		rl.DrawText("some text here", monw-350, 80, 30, rl.White)
		rl.DrawText("some text here", monw-350, 110, 30, rl.White)
	case 12:
		rl.DrawText("locked", monw-303, 33, 40, rl.Magenta)
		rl.DrawText("locked", monw-301, 31, 40, rl.Black)
		rl.DrawText("locked", monw-300, 30, 40, rl.White)
		rl.DrawText("some text here", monw-350, 80, 30, rl.White)
		rl.DrawText("some text here", monw-350, 110, 30, rl.White)

	case 1:
		if onoff15 {
			rl.DrawRectangle((monw/2)-220, 50, 200, 200, rl.Fade(rl.Magenta, 0.1))
			rl.DrawRectangleLines((monw/2)-220, 50, 200, 200, rl.Fade(rl.Purple, 0.1))
		} else {
			rl.DrawRectangle((monw/2)-220, 50, 200, 200, rl.Fade(rl.Magenta, 0.2))
			rl.DrawRectangleLines((monw/2)-220, 50, 200, 200, rl.Fade(rl.Purple, 0.2))
		}
		if framecount%3 == 0 {
			playerleft.X -= 16
			if playerleft.X < 16 {
				playerleft.X = 81
			}
		}
		rl.DrawText("pinkatin", (monw/2)-176, 262, 30, rl.Magenta)
		rl.DrawText("pinkatin", (monw/2)-175, 261, 30, rl.Black)
		rl.DrawText("pinkatin", (monw/2)-174, 260, 30, rl.White)
	case 2:
		rl.DrawText("locked", monw-303, 33, 40, rl.Magenta)
		rl.DrawText("locked", monw-301, 31, 40, rl.Black)
		rl.DrawText("locked", monw-300, 30, 40, rl.White)
		rl.DrawText("another character", monw-350, 80, 30, rl.White)
		rl.DrawText("different gameplay", monw-350, 110, 30, rl.White)
		if onoff15 {
			rl.DrawRectangle((monw/2)+20, 50, 200, 200, rl.Fade(rl.Magenta, 0.1))
			rl.DrawRectangleLines((monw/2)+20, 50, 200, 200, rl.Fade(rl.Purple, 0.1))
		} else {
			rl.DrawRectangle((monw/2)+20, 50, 200, 200, rl.Fade(rl.Magenta, 0.2))
			rl.DrawRectangleLines((monw/2)+20, 50, 200, 200, rl.Fade(rl.Purple, 0.2))
		}
		if framecount%3 == 0 {
			player2left.X -= 16
			if player2left.X < 119 {
				player2left.X = 184
			}
		}
		if purpiltinlocked {
			rl.DrawText("locked", (monw/2)+68, 262, 30, rl.Purple)
			rl.DrawText("locked", (monw/2)+69, 261, 30, rl.Black)
			rl.DrawText("locked", (monw/2)+70, 260, 30, rl.White)
		} else {
			rl.DrawText("purpiltin", (monw/2)+53, 262, 30, rl.Purple)
			rl.DrawText("purpiltin", (monw/2)+54, 261, 30, rl.Black)
			rl.DrawText("purpiltin", (monw/2)+55, 260, 30, rl.White)
		}
	}

	// draw pinkatin
	pinkatinv2 := rl.NewVector2(float32(((monw/2)-180)/8), 10)
	rl.BeginMode2D(camera8x)
	rl.DrawTextureRec(imgs, playerleft, pinkatinv2, rl.White)
	rl.DrawTextureRec(imgs, playerleft, pinkatinv2, rl.Fade(rl.Magenta, playerfade))
	// draw purpiltin
	pinkatinv2.X += 30
	rl.DrawTextureRec(imgs, player2left, pinkatinv2, rl.White)
	rl.DrawTextureRec(imgs, player2left, pinkatinv2, rl.Fade(rl.Purple, playerfade))

	purpiltinlocked = true
	if purpiltinlocked {
		rl.DrawTextureRec(imgs, lockedicon, pinkatinv2, rl.White)
	}
	rl.EndMode2D()

}
func drawstartscreen() { // MARK: drawstartscreen

	// background rec
	rl.DrawRectangle(0, 0, monw, monh, rl.Black)
	// go logo
	if gologov2.X < float32((monw/2)-193) {
		gologov2.X += 20
	} else {
		rl.DrawText("built with", (monw/2)-90, (monh/2)-330, 40, rl.White)
		introtimer1on = true
	}
	if introtimer2on {
		gologov2.X += 20
		if gologov2.X > float32(monw+100) {
			introtimer2on = false
		}
	}
	if rolldice() == 6 {
		rl.DrawTextureRec(imgs, gologo, gologov2, rl.Fade(rl.Magenta, startlogofade))
	} else {
		rl.DrawTextureRec(imgs, gologo, gologov2, rl.Fade(rl.White, startlogofade))
	}
	// raylib logo
	if rayliblogoon {
		if rayliblogov2.Y > float32((monh/3)-80) {
			rayliblogov2.Y -= 20
		} else {
			introtimer3on = true
		}
		rl.DrawTextureRec(imgs, rayliblogo, rayliblogov2, rl.White)
	}
	// pink back
	if intropinkbackon {
		recsize := rl.NewVector2(float32(monw), float32(monh))
		rl.DrawRectangleV(intropinkbackv2, recsize, rl.Magenta)
		if intropinkbackv2.Y < 0 {
			intropinkbackv2.Y += 20
		} else {
			intronicholasimontexton = true
		}
	}
	// nicholasimon text
	if !intronicholasimontexton2 {
		if intronicholasimontexton {
			rl.DrawText("a game by nicholasimon", (monw/2)-230, (monh/3)*2, 40, rl.Fade(rl.Black, nicholasimontextfade))
			if nicholasimontextfade < 1.0 {
				nicholasimontextfade += 0.05
			} else {
				introtimer4on = true
			}
		}
	}
	if introcircleson {
		for a := 0; a < len(introcirlcesv2layout); a++ {
			circle := introcirlcesv2layout[a]
			rl.DrawCircleV(circle.center, circle.radius, circle.color)
			if circle.radius < 400 {
				circle.radius += rFloat32(1, 5)
			}
			introcirlcesv2layout[a] = circle
		}
	}

}
func drawinventory() { // MARK: drawinventory

	// background rectangle
	rl.DrawRectangle(inventoryx, inventoryy, 800, 90, rl.Fade(rl.White, 0.6))
	//background rectangle shadow
	rl.DrawRectangle(inventoryx, inventoryy+90, 804, 4, rl.Fade(rl.Black, 0.8))
	rl.DrawRectangle(inventoryx+800, inventoryy+44, 4, 46, rl.Fade(rl.Black, 0.8))

	v2 := rl.NewVector2(float32((inventoryx+16)/4), float32((inventoryy+16)/4))

	// select rectangles
	switch inventoryselect {
	case 1:
		rl.DrawRectangleLines((inventoryx + 12), inventoryy+13, 64, 64, rl.Magenta)
		rl.DrawRectangleLines((inventoryx + 11), inventoryy+12, 66, 66, rl.Magenta)
		rl.DrawRectangle((inventoryx + 11), inventoryy+12, 66, 66, rl.Fade(rl.Magenta, inventoryfade))
	case 2:
		rl.DrawRectangleLines((inventoryx + 97), inventoryy+13, 64, 64, rl.Magenta)
		rl.DrawRectangleLines((inventoryx + 96), inventoryy+12, 66, 66, rl.Magenta)
		rl.DrawRectangle((inventoryx + 96), inventoryy+12, 66, 66, rl.Fade(rl.Magenta, inventoryfade))
	case 3:
		rl.DrawRectangleLines((inventoryx + 177), inventoryy+14, 64, 64, rl.Magenta)
		rl.DrawRectangleLines((inventoryx + 176), inventoryy+13, 66, 66, rl.Magenta)
		rl.DrawRectangle((inventoryx + 176), inventoryy+13, 66, 66, rl.Fade(rl.Magenta, inventoryfade))
	case 4:
		rl.DrawRectangleLines((inventoryx + 257), inventoryy+14, 64, 64, rl.Magenta)
		rl.DrawRectangleLines((inventoryx + 256), inventoryy+13, 66, 66, rl.Magenta)
		rl.DrawRectangle((inventoryx + 256), inventoryy+13, 66, 66, rl.Fade(rl.Magenta, inventoryfade))
	case 5:
		rl.DrawRectangleLines((inventoryx + 337), inventoryy+14, 64, 64, rl.Magenta)
		rl.DrawRectangleLines((inventoryx + 336), inventoryy+13, 66, 66, rl.Magenta)
		rl.DrawRectangle((inventoryx + 336), inventoryy+13, 66, 66, rl.Fade(rl.Magenta, inventoryfade))
	case 6:
		rl.DrawRectangleLines((inventoryx + 417), inventoryy+14, 64, 64, rl.Magenta)
		rl.DrawRectangleLines((inventoryx + 416), inventoryy+13, 66, 66, rl.Magenta)
		rl.DrawRectangle((inventoryx + 416), inventoryy+13, 66, 66, rl.Fade(rl.Magenta, inventoryfade))
	case 7:
		rl.DrawRectangleLines((inventoryx + 495), inventoryy+13, 64, 64, rl.Magenta)
		rl.DrawRectangleLines((inventoryx + 494), inventoryy+12, 66, 66, rl.Magenta)
		rl.DrawRectangle((inventoryx + 494), inventoryy+12, 66, 66, rl.Fade(rl.Magenta, inventoryfade))
	case 8:
		rl.DrawRectangleLines((inventoryx + 575), inventoryy+13, 64, 64, rl.Magenta)
		rl.DrawRectangleLines((inventoryx + 574), inventoryy+12, 66, 66, rl.Magenta)
		rl.DrawRectangle((inventoryx + 574), inventoryy+12, 66, 66, rl.Fade(rl.Magenta, inventoryfade))
	case 9:
		rl.DrawRectangleLines((inventoryx + 652), inventoryy+13, 64, 64, rl.Magenta)
		rl.DrawRectangleLines((inventoryx + 651), inventoryy+12, 66, 66, rl.Magenta)
		rl.DrawRectangle((inventoryx + 651), inventoryy+12, 66, 66, rl.Fade(rl.Magenta, inventoryfade))
	}

	// inventory item slots
	rl.BeginMode2D(camera4x)
	rl.DrawTextureRec(imgs, pickaraxel, v2, rl.White)
	v2.X += 20
	rl.DrawTextureRec(imgs, spadel, v2, rl.White)
	v2.X += 20
	rl.DrawTextureRec(imgs, icewandl, v2, rl.White)
	v2.X += 22
	v2.Y += 2
	rl.DrawTextureRec(imgs, sword1l, v2, rl.White)
	v2.X += 18
	v2.Y -= 2
	rl.DrawTextureRec(imgs, bigswordl, v2, rl.White)
	v2.X += 20
	rl.DrawTextureRec(imgs, axel, v2, rl.White)
	v2.X += 20
	v2.Y += 4
	rl.DrawTextureRec(imgs, pumpactionl, v2, rl.White)
	v2.X += 20
	v2.Y--
	rl.DrawTextureRec(imgs, bazookal, v2, rl.White)
	v2.X += 20
	v2.Y--
	rl.DrawTextureRec(imgs, uzzil, v2, rl.White)
	rl.EndMode2D()

	if inventoryxswitch {
		if inventoryx > -600 {
			inventoryx -= 40
		} else {
			inventoryon = false
			inventoryxswitch = false
		}
	} else {
		if inventoryx < 0 {
			inventoryx += 40
		} else {
			// inventory tab
			rl.DrawRectangle(inventoryx+800, inventoryy, 60, 40, rl.Fade(rl.White, 0.6))
			rl.DrawText("inv", inventoryx+807, inventoryy+5, 30, rl.Black)
			rl.DrawText("inv", inventoryx+808, inventoryy+4, 30, rl.Magenta)
			// inventory tab shadow
			rl.DrawRectangle(inventoryx+860, inventoryy+4, 4, 40, rl.Fade(rl.Black, 0.8))
			rl.DrawRectangle(inventoryx+800, inventoryy+40, 60, 4, rl.Fade(rl.Black, 0.8))
		}
	}
	// select item
	if rl.IsKeyPressed(rl.KeySpace) {
		switch inventoryselect {
		case 1:
			pickaxeactive = true
			icewandactive = false
			spadeactive = false
			bigswordactive = false
			axeactive = false
			pumpactionactive = false
			bazookaactive = false
			uzziactive = false

			inventoryxswitch = true
		case 2:
			spadeactive = true
			icewandactive = false
			pickaxeactive = false
			bigswordactive = false
			axeactive = false
			pumpactionactive = false
			bazookaactive = false
			uzziactive = false

			inventoryxswitch = true
		case 3:
			icewandactive = true
			spadeactive = false
			pickaxeactive = false
			bigswordactive = false
			axeactive = false
			pumpactionactive = false
			bazookaactive = false
			uzziactive = false

			inventoryxswitch = true
		case 4:
			icewandactive = false
			spadeactive = false
			pickaxeactive = false
			bigswordactive = false
			axeactive = false
			pumpactionactive = false
			bazookaactive = false
			uzziactive = false

			inventoryxswitch = true
		case 5:
			bigswordactive = true
			icewandactive = false
			spadeactive = false
			pickaxeactive = false
			axeactive = false
			pumpactionactive = false
			bazookaactive = false
			uzziactive = false

			inventoryxswitch = true
		case 6:
			axeactive = true
			bigswordactive = false
			icewandactive = false
			spadeactive = false
			pickaxeactive = false
			pumpactionactive = false
			bazookaactive = false
			uzziactive = false

			inventoryxswitch = true
		case 7:
			pumpactionactive = true
			axeactive = false
			bigswordactive = false
			icewandactive = false
			spadeactive = false
			pickaxeactive = false
			bazookaactive = false
			uzziactive = false

			inventoryxswitch = true
		case 8:
			bazookaactive = true
			pumpactionactive = false
			axeactive = false
			bigswordactive = false
			icewandactive = false
			spadeactive = false
			pickaxeactive = false
			uzziactive = false

			inventoryxswitch = true
		case 9:
			uzziactive = true
			bazookaactive = false
			pumpactionactive = false
			axeactive = false
			bigswordactive = false
			icewandactive = false
			spadeactive = false
			pickaxeactive = false

			inventoryxswitch = true
		}
	}
	// close inventory
	if rl.IsKeyPressed(rl.KeyF8) && !inventorypause {
		inventoryxswitch = true
	}
	// move select box
	if rl.IsKeyPressed(rl.KeyRight) {
		inventoryselect++
		if inventoryselect > 9 {
			inventoryselect = 1
		}
	} else if rl.IsKeyPressed(rl.KeyLeft) {
		inventoryselect--
		if inventoryselect < 1 {
			inventoryselect = 9
		}
	}

}
func drawoptions() { // MARK: drawoptions

	if inputlayouton {

		v2 := rl.NewVector2(float32((monw/2-236)/4), 12)
		rl.BeginMode2D(camera4x)
		// controller
		rl.DrawTextureRec(imgs, controllerimg, v2, rl.Magenta)
		rl.EndMode2D()
		// alt
		v2 = rl.NewVector2(float32((monw / 2)), 550)
		rl.DrawTextureRec(imgs, altkey, v2, rl.Magenta)
		rl.DrawText("bomb", (monw/2)+112, 571, 40, rl.Magenta)
		rl.DrawText("bomb", (monw/2)+114, 569, 40, rl.Black)
		rl.DrawText("bomb", (monw/2)+115, 568, 40, rl.White)
		// control
		v2 = rl.NewVector2(float32((monw/2)-305), 550)
		rl.DrawTextureRec(imgs, ctrlkey, v2, rl.Magenta)
		rl.DrawText("bungee", (monw/2)-193, 571, 40, rl.Magenta)
		rl.DrawText("bungee", (monw/2)-191, 569, 40, rl.Black)
		rl.DrawText("bungee", (monw/2)-190, 568, 40, rl.White)
		// tab
		v2 = rl.NewVector2(float32((monw/2)-550), 450)
		rl.DrawTextureRec(imgs, tabkey, v2, rl.Magenta)
		rl.DrawText("inventory", (monw/2)-437, 471, 40, rl.Magenta)
		rl.DrawText("inventory", (monw/2)-435, 469, 40, rl.Black)
		rl.DrawText("inventory", (monw/2)-434, 468, 40, rl.White)
		// escape
		v2 = rl.NewVector2(float32((monw/2)-550), 550)
		rl.DrawTextureRec(imgs, escapekey, v2, rl.Magenta)
		rl.DrawText("menu", (monw/2)-448, 571, 40, rl.Magenta)
		rl.DrawText("menu", (monw/2)-446, 569, 40, rl.Black)
		rl.DrawText("menu", (monw/2)-445, 568, 40, rl.White)
		// space
		v2 = rl.NewVector2(float32((monw/2)-190), 450)
		rl.DrawTextureRec(imgs, spacekey, v2, rl.Magenta)
		rl.DrawText("attack | open", (monw/2)-38, 471, 40, rl.Magenta)
		rl.DrawText("attack | open", (monw/2)-36, 469, 40, rl.Black)
		rl.DrawText("attack | open", (monw/2)-35, 468, 40, rl.White)
		// arrows
		v2 = rl.NewVector2(float32((monw/2 + 250)), 465)
		rl.DrawTextureRec(imgs, arrowkeysimg, v2, rl.Magenta)
		rl.DrawText("move", monw/2+438, 471, 40, rl.Magenta)
		rl.DrawText("move", monw/2+436, 469, 40, rl.Black)
		rl.DrawText("move", monw/2+435, 468, 40, rl.White)
		// left controller
		rl.DrawRectangle(monw/2-333, 208, 182, 4, rl.Pink)
		rl.DrawRectangle(monw/2-332, 207, 182, 4, rl.Black)
		rl.DrawRectangle(monw/2-330, 205, 182, 4, rl.White)
		rl.DrawText("move", monw/2-448, 188, 40, rl.Magenta)
		rl.DrawText("move", monw/2-446, 186, 40, rl.Black)
		rl.DrawText("move", monw/2-445, 185, 40, rl.White)
		// right controller
		// B
		rl.DrawRectangle(monw/2+207, 193, 185, 4, rl.Pink)
		rl.DrawRectangle(monw/2+206, 192, 185, 4, rl.Black)
		rl.DrawRectangle(monw/2+204, 190, 185, 4, rl.White)
		rl.DrawText("b text", monw/2+413, 175, 40, rl.Magenta)
		rl.DrawText("b text", monw/2+415, 173, 40, rl.Black)
		rl.DrawText("b text", monw/2+416, 172, 40, rl.White)
		// A
		rl.DrawRectangle(monw/2+161, 243, 4, 20, rl.Pink)
		rl.DrawRectangle(monw/2+160, 242, 4, 20, rl.Black)
		rl.DrawRectangle(monw/2+158, 240, 4, 20, rl.White)
		rl.DrawRectangle(monw/2+161, 263, 185, 4, rl.Pink)
		rl.DrawRectangle(monw/2+160, 262, 185, 4, rl.Black)
		rl.DrawRectangle(monw/2+158, 260, 185, 4, rl.White)
		rl.DrawText("a text", monw/2+369, 243, 40, rl.Magenta)
		rl.DrawText("a text", monw/2+371, 241, 40, rl.Black)
		rl.DrawText("a text", monw/2+372, 240, 40, rl.White)
		// Y
		rl.DrawRectangle(monw/2+158, 123, 4, 22, rl.Pink)
		rl.DrawRectangle(monw/2+157, 122, 4, 24, rl.Black)
		rl.DrawRectangle(monw/2+155, 120, 4, 26, rl.White)
		rl.DrawRectangle(monw/2+161, 123, 182, 4, rl.Pink)
		rl.DrawRectangle(monw/2+159, 122, 183, 4, rl.Black)
		rl.DrawRectangle(monw/2+155, 120, 185, 4, rl.White)
		rl.DrawText("y text", monw/2+362, 103, 40, rl.Magenta)
		rl.DrawText("y text", monw/2+364, 101, 40, rl.Black)
		rl.DrawText("y text", monw/2+365, 100, 40, rl.White)
		// X
		rl.DrawRectangle(monw/2+121, 205, 4, 120, rl.Pink)
		rl.DrawRectangle(monw/2+120, 204, 4, 122, rl.Black)
		rl.DrawRectangle(monw/2+118, 202, 4, 122, rl.White)
		rl.DrawRectangle(monw/2+121, 323, 185, 4, rl.Pink)
		rl.DrawRectangle(monw/2+120, 322, 185, 4, rl.Black)
		rl.DrawRectangle(monw/2+118, 320, 185, 4, rl.White)
		rl.DrawText("x text", monw/2+327, 303, 40, rl.Magenta)
		rl.DrawText("x text", monw/2+329, 301, 40, rl.Black)
		rl.DrawText("x text", monw/2+330, 300, 40, rl.White)

	} else {
		rl.DrawRectangle(monw/2-240, optionselecty, 590, 60, rl.Fade(rl.Purple, 0.1))

		rl.DrawText("options", monw/2-103, 53, 80, rl.Magenta)
		rl.DrawText("options", monw/2-101, 51, 80, rl.Black)
		rl.DrawText("options", monw/2-100, 50, 80, rl.White)

		rl.DrawText("resolution", monw/2-183, 183, 40, rl.Magenta)
		rl.DrawText("resolution", monw/2-181, 181, 40, rl.Black)
		rl.DrawText("resolution", monw/2-180, 180, 40, rl.White)
		if fullscreenon {
			rl.DrawText("fullscreen", monw/2-43, 243, 40, rl.Magenta)
			rl.DrawText("fullscreen", monw/2-41, 241, 40, rl.Black)
			rl.DrawText("fullscreen", monw/2-40, 240, 40, rl.White)
		} else {
			rl.DrawText("window", monw/2-23, 243, 40, rl.Magenta)
			rl.DrawText("window", monw/2-21, 241, 40, rl.Black)
			rl.DrawText("window", monw/2-20, 240, 40, rl.White)
		}
		if soundon {
			rl.DrawText("sound on", monw/2-38, 303, 40, rl.Magenta)
			rl.DrawText("sound on", monw/2-36, 301, 40, rl.Black)
			rl.DrawText("sound on", monw/2-35, 300, 40, rl.White)

		} else {
			rl.DrawText("sound off", monw/2-38, 303, 40, rl.Magenta)
			rl.DrawText("sound off", monw/2-36, 301, 40, rl.Black)
			rl.DrawText("sound off", monw/2-35, 300, 40, rl.White)
		}
		if controlleron {
			rl.DrawText("controller on", monw/2-68, 363, 40, rl.Magenta)
			rl.DrawText("controller on", monw/2-66, 361, 40, rl.Black)
			rl.DrawText("controller on", monw/2-65, 360, 40, rl.White)

		} else {
			rl.DrawText("controller off", monw/2-68, 363, 40, rl.Magenta)
			rl.DrawText("controller off", monw/2-66, 361, 40, rl.Black)
			rl.DrawText("controller off", monw/2-65, 360, 40, rl.White)
		}
		rl.DrawText("input layout", monw/2-68, 423, 40, rl.Magenta)
		rl.DrawText("input layout", monw/2-66, 421, 40, rl.Black)
		rl.DrawText("input layout", monw/2-65, 420, 40, rl.White)

		v4 := rl.NewVector2(float32(monw/2-123), 370)
		v3 := rl.NewVector2(float32(monw/2-121), 368)
		v2 := rl.NewVector2(float32(monw/2-120), 367)
		if controlleron {
			rl.DrawTextureRec(imgs, controllericon, v4, rl.Magenta)
			rl.DrawTextureRec(imgs, controllericon, v3, rl.Black)
			rl.DrawTextureRec(imgs, controllericon, v2, rl.White)
		} else {
			rl.DrawTextureRec(imgs, keyboardicon, v4, rl.Magenta)
			rl.DrawTextureRec(imgs, keyboardicon, v3, rl.Black)
			rl.DrawTextureRec(imgs, keyboardicon, v2, rl.White)
		}

		v4 = rl.NewVector2(float32(monw/2-88), 307)
		v3 = rl.NewVector2(float32(monw/2-86), 305)
		v2 = rl.NewVector2(float32(monw/2-85), 304)
		if soundon {
			rl.DrawTextureRec(imgs, musiconicon, v4, rl.Magenta)
			rl.DrawTextureRec(imgs, musiconicon, v3, rl.Black)
			rl.DrawTextureRec(imgs, musiconicon, v2, rl.White)
		} else {
			rl.DrawTextureRec(imgs, musicofficon, v4, rl.Magenta)
			rl.DrawTextureRec(imgs, musicofficon, v3, rl.Black)
			rl.DrawTextureRec(imgs, musicofficon, v2, rl.White)
		}

		switch resolutionselect {
		case 1:
			rl.DrawText("1280 X 720", monw/2+67, 183, 40, rl.Magenta)
			rl.DrawText("1280 X 720", monw/2+69, 181, 40, rl.Black)
			rl.DrawText("1280 X 720", monw/2+70, 180, 40, rl.White)
		case 2:
			rl.DrawText("1366 X 768", monw/2+67, 183, 40, rl.Magenta)
			rl.DrawText("1366 X 768", monw/2+69, 181, 40, rl.Black)
			rl.DrawText("1366 X 768", monw/2+70, 180, 40, rl.White)
		case 3:
			rl.DrawText("1440 X 900", monw/2+67, 183, 40, rl.Magenta)
			rl.DrawText("1440 X 900", monw/2+69, 181, 40, rl.Black)
			rl.DrawText("1440 X 900", monw/2+70, 180, 40, rl.White)
		case 4:
			rl.DrawText("1536 X 864", monw/2+67, 183, 40, rl.Magenta)
			rl.DrawText("1536 X 864", monw/2+69, 181, 40, rl.Black)
			rl.DrawText("1536 X 864", monw/2+70, 180, 40, rl.White)
		case 5:
			rl.DrawText("1600 X 900", monw/2+67, 183, 40, rl.Magenta)
			rl.DrawText("1600 X 900", monw/2+69, 181, 40, rl.Black)
			rl.DrawText("1600 X 900", monw/2+70, 180, 40, rl.White)
		case 6:
			rl.DrawText("1920 X 1080", monw/2+67, 183, 40, rl.Magenta)
			rl.DrawText("1920 X 1080", monw/2+69, 181, 40, rl.Black)
			rl.DrawText("1920 X 1080", monw/2+70, 180, 40, rl.White)
		case 7:
			rl.DrawText("3840 X 2160", monw/2+67, 183, 40, rl.Magenta)
			rl.DrawText("3840 X 2160", monw/2+69, 181, 40, rl.Black)
			rl.DrawText("3840 X 2160", monw/2+70, 180, 40, rl.White)
		}
	}
	switch optionselect {
	case 1:
		if rl.IsKeyPressed(rl.KeyRight) {
			resolutionselect++
			if resolutionselect > 7 {
				resolutionselect = 1
			}
		}
		if rl.IsKeyPressed(rl.KeyLeft) {
			resolutionselect--
			if resolutionselect < 1 {
				resolutionselect = 7
			}
		}
		if rl.IsKeyPressed(rl.KeySpace) {
			setnewscreen()
		}
	case 2:
		if rl.IsKeyPressed(rl.KeyRight) {
			if fullscreenon {
				fullscreenon = false
			} else {
				fullscreenon = true
			}
		}
		if rl.IsKeyPressed(rl.KeyLeft) {
			if fullscreenon {
				fullscreenon = false
			} else {
				fullscreenon = true
			}
		}
		if rl.IsKeyPressed(rl.KeySpace) {
			setnewscreen()
		}
	case 3:
		if rl.IsKeyPressed(rl.KeyRight) {
			if soundon {
				soundon = false
			} else {
				soundon = true
			}
		}
		if rl.IsKeyPressed(rl.KeyLeft) {
			if soundon {
				soundon = false
			} else {
				soundon = true
			}
		}
	case 4:
		if rl.IsKeyPressed(rl.KeyRight) {
			if controlleron {
				controlleron = false
			} else {
				controlleron = true
			}
		}
		if rl.IsKeyPressed(rl.KeyLeft) {
			if controlleron {
				controlleron = false
			} else {
				controlleron = true
			}
		}
	case 5:
		if rl.IsKeyPressed(rl.KeySpace) {
			if inputlayouton {
				inputlayouton = false
			} else {
				inputlayouton = true
			}
		}
	}

	if rl.IsKeyPressed(rl.KeyDown) && !inputlayouton {
		optionselect++
		optionselecty += 60
		if optionselect > 5 {
			optionselect = 1
			optionselecty = 170
		}
	} else if rl.IsKeyPressed(rl.KeyUp) && !inputlayouton {
		optionselect--
		optionselecty -= 60
		if optionselect < 1 {
			optionselect = 5
			optionselecty = 410
		}
	}

}
func drawquestdetails() { // MARK: drawquestdetails

	rl.DrawRectangle(monw-250, 200, 230, 130, rl.Fade(rl.White, 0.7))
	rl.DrawText("quest:", monw-230, 210, 30, rl.Black)
	rl.DrawText(activequest.questtype, monw-230, 250, 30, rl.Black)
	rl.DrawText(activequest.item, monw-230, 290, 30, rl.Black)
}
func drawdiscoveryinfobox() { // drawdiscoveryinfobox
	rl.DrawRectangle(200, 200, monw-400, monh-400, rl.Fade(rl.Black, 0.9))
	if batdiscovered {
		rl.DrawText("bat", discoveryinfox, discoveryinfoy, 40, rl.White)
		rl.DrawText("harmless", discoveryinfox, discoveryinfoy+42, 20, rl.White)
		batv2 := rl.NewVector2(float32(discoveryinfox/4+55), float32(discoveryinfoy/4+2))
		rl.BeginMode2D(camera4x)
		rl.DrawTextureRec(imgs, bat, batv2, rl.Fade(rl.DarkBlue, 0.8))
		rl.EndMode2D()
	}
}
func drawdiscoveries() { // MARK: drawdiscoveries
	if !pauseon {
		if batdiscovered && discoverytimer != 0 {
			rl.DrawRectangle(discoveryrecx, 80, 800, 90, rl.Fade(rl.White, 0.8))
			rl.DrawRectangle(discoveryrecx, 170, 796, 4, rl.Fade(rl.Black, 0.8))
			rl.DrawText("you have discovered a bat", discoverytext1x, 91, 40, rl.Magenta)
			rl.DrawText("you have discovered a bat", discoverytext1shadowx, 90, 40, rl.Black)
			rl.DrawText(" >> harmless background critter", discoverytext2x, 136, 20, rl.Magenta)
			rl.DrawText(" >> harmless background critter", discoverytext2shadowx, 135, 20, rl.Black)
			batv2 := rl.NewVector2(float32(discoveryimgx), 25)
			rl.BeginMode2D(camera4x)

			rl.DrawTextureRec(imgs, bat, batv2, rl.DarkBlue)
			rl.EndMode2D()
		}

	}
}
func drawinfobar() { // MARK: drawinfobar

	if !pauseon {
		rl.BeginMode2D(camera4x)
		// hp
		hpimgv2 := rl.NewVector2(4, 2)
		for a := 0; a < hp; a++ {
			hpimgv2shadow := rl.NewVector2(hpimgv2.X+1, hpimgv2.Y+1)
			rl.DrawTextureRec(imgs, hpimg, hpimgv2shadow, rl.Black)
			rl.DrawTextureRec(imgs, hpimg, hpimgv2, rl.Pink)
			rl.DrawTextureRec(imgs, hpimg, hpimgv2, rl.Fade(rl.Magenta, hpfade))
			hpimgv2.X += 16
		}
		// quest arrow
		if questactive {
			originv2 := rl.NewVector2(6, 7)
			destrec := rl.NewRectangle(14, 30, 12, 14)
			rl.DrawTexturePro(imgs, arrowup, destrec, originv2, questarrowrotation, rl.White)
			rl.DrawTexturePro(imgs, arrowup, destrec, originv2, questarrowrotation, rl.Fade(rl.Magenta, hpfade))
		}
		// end key arrow
		if endkeyactive {
			originv2 := rl.NewVector2(6, 7)
			destrec := rl.NewRectangle(14, 50, 12, 14)
			rl.DrawTexturePro(imgs, arrowup, destrec, originv2, endkeyarrowrotation, rl.White)
			rl.DrawTexturePro(imgs, arrowup, destrec, originv2, endkeyarrowrotation, rl.Fade(rl.Magenta, hpfade))
		}

		coinv2 := rl.NewVector2(float32(monw/4-50), 3)
		rl.DrawTextureRec(imgs, coin, coinv2, rl.Fade(rl.Yellow, 0.8))
		gemv2 := rl.NewVector2(float32(monw/4-90), 3)
		rl.DrawTextureRec(imgs, gem, gemv2, rl.Fade(gemcolor, 0.9))
		bombv2 := rl.NewVector2(float32(monw/4-132), 5)
		rl.DrawCircle(int(bombv2.X+6), int(bombv2.Y+6), 8, rl.White)
		rl.DrawTextureRec(imgs, bomb1, bombv2, rl.Black)
		rl.DrawTextureRec(imgs, bomb1, bombv2, rl.Fade(rl.Magenta, hpfade))

		// end key
		if endkeycollected {
			drawv2 := rl.NewVector2(float32(monw-100), float32(150))
			switch endkeynumber {
			case 1:
				rl.DrawTextureRec(imgs, key1, drawv2, rl.White)
			case 2:
				rl.DrawTextureRec(imgs, key2, drawv2, rl.White)
			case 3:
				rl.DrawTextureRec(imgs, key3, drawv2, rl.White)
			case 4:
				rl.DrawTextureRec(imgs, key4, drawv2, rl.White)
			case 5:
				rl.DrawTextureRec(imgs, key5, drawv2, rl.White)
			case 6:
				rl.DrawTextureRec(imgs, key6, drawv2, rl.White)
			case 7:
				rl.DrawTextureRec(imgs, key7, drawv2, rl.White)
			case 8:
				rl.DrawTextureRec(imgs, key8, drawv2, rl.White)
			case 9:
				rl.DrawTextureRec(imgs, key9, drawv2, rl.White)
			}
		}
		rl.EndMode2D()

		// alert level
		rl.DrawRectangle(monw-110, 105, 32, 32, rl.DarkGreen)
		rl.DrawText("green", monw-250, 100, 40, rl.DarkGreen)
		// shield bar
		shieldnumber = 60
		if shieldon {
			rl.DrawRectangle(monw-100, 302, 42, 302, rl.Fade(rl.Black, 0.5))
			ychange := 0
			for a := 0; a < shieldnumber; a++ {
				rl.DrawRectangle(monw-99, 600-ychange, 40, 3, rl.Magenta)
				ychange += 3
			}
			rl.DrawRectangleLines(monw-100, 302, 42, 302, rl.Purple)

			rl.DrawText("S", monw-47, 308, 40, rl.Black)
			rl.DrawText("S", monw-50, 305, 40, rl.Magenta)
			rl.DrawText("H", monw-47, 358, 40, rl.Black)
			rl.DrawText("H", monw-50, 355, 40, rl.Magenta)
			rl.DrawText("I", monw-47, 408, 40, rl.Black)
			rl.DrawText("I", monw-50, 405, 40, rl.Magenta)
			rl.DrawText("E", monw-47, 458, 40, rl.Black)
			rl.DrawText("E", monw-50, 455, 40, rl.Magenta)
			rl.DrawText("L", monw-47, 508, 40, rl.Black)
			rl.DrawText("L", monw-50, 505, 40, rl.Magenta)
			rl.DrawText("D", monw-47, 558, 40, rl.Black)
			rl.DrawText("D", monw-50, 555, 40, rl.Magenta)
		}

		if questactive {
			rl.DrawText("quest", 98, 92, 60, rl.Black)
			rl.DrawText("quest", 100, 90, 60, rl.Pink)
			rl.DrawText("quest", 100, 90, 60, rl.Fade(rl.Magenta, hpfade))
		}
		if endkeyactive {
			rl.DrawText("exit key", 98, 166, 60, rl.Black)
			rl.DrawText("exit key", 100, 164, 60, rl.Pink)
			rl.DrawText("exit key", 100, 164, 60, rl.Fade(rl.Magenta, hpfade))
		}

		gemstotalTEXT := strconv.Itoa(gemstotal)
		rl.DrawText(gemstotalTEXT, monw-288, 10, 80, rl.Black)
		rl.DrawText(gemstotalTEXT, monw-290, 8, 80, rl.Pink)
		rl.DrawText(gemstotalTEXT, monw-290, 8, 80, rl.Fade(rl.Magenta, hpfade))

		bombtotalTEXT := strconv.Itoa(bombtotal)
		rl.DrawText(bombtotalTEXT, monw-448, 10, 80, rl.Black)
		rl.DrawText(bombtotalTEXT, monw-450, 8, 80, rl.Pink)
		rl.DrawText(bombtotalTEXT, monw-450, 8, 80, rl.Fade(rl.Magenta, hpfade))

		coinstotalTEXT := strconv.Itoa(coinstotal)
		rl.DrawText(coinstotalTEXT, monw-118, 10, 80, rl.Black)
		rl.DrawText(coinstotalTEXT, monw-120, 8, 80, rl.Pink)
		rl.DrawText(coinstotalTEXT, monw-120, 8, 80, rl.Fade(rl.Magenta, hpfade))

		infobarpowerupsv2 := rl.NewVector2(4, float32(monh/4-24))
		poweruptextx := 70
		poweruptexty := monh - 74
		for a := 0; a < len(powerupscollected); a++ {
			checkpowerup := powerupscollected[a]
			rl.BeginMode2D(camera4x)
			rl.DrawTextureRec(imgs, checkpowerup.img, infobarpowerupsv2, rl.White)
			rl.DrawTextureRec(imgs, checkpowerup.img, infobarpowerupsv2, rl.Fade(rl.Brown, fireballfade))
			rl.EndMode2D()
			if checkpowerup.number > 1 {
				numberTEXT := strconv.Itoa(checkpowerup.number)
				rl.DrawText(numberTEXT, poweruptextx-2, poweruptexty+2, 20, rl.Black)
				rl.DrawText(numberTEXT, poweruptextx, poweruptexty, 20, rl.Magenta)
			}

			infobarpowerupsv2.X += 18

		}

	}

}
func drawintro() { // MARK: drawintro

	if controlleron {
		rl.DrawText("press any button to start", 37, monh-77, 40, introtextcolor2)
		rl.DrawText("press any button to start", 39, monh-79, 40, rl.Black)
		rl.DrawText("press any button to start", 40, monh-80, 40, rl.White)
		rl.DrawText("press any button to start", 40, monh-80, 40, introtextcolor)
	} else {
		rl.DrawText("press space to start", 37, monh-77, 40, introtextcolor2)
		rl.DrawText("press space to start", 39, monh-79, 40, rl.Black)
		rl.DrawText("press space to start", 40, monh-80, 40, rl.White)
		rl.DrawText("press space to start", 40, monh-80, 40, introtextcolor)
	}

	if framecount%15 == 0 {
		if introtextcolor == rl.Magenta {
			introtextcolor = rl.White
			introtextcolor2 = rl.Magenta
		} else {
			introtextcolor = rl.Magenta
			introtextcolor2 = rl.White
		}
	}

	rl.DrawText("pinkatin", monw-503, monh-137, 120, rl.Black)
	rl.DrawText("pinkatin", monw-499, monh-141, 120, rl.White)
	rl.DrawText("pinkatin", monw-498, monh-142, 120, rl.Black)
	rl.DrawText("pinkatin", monw-497, monh-143, 120, rl.White)
	if onoff3 {
		rl.DrawText("pinkatin", monw-497, monh-143, 120, rl.Fade(rl.Magenta, playerfade))
	} else {
		rl.DrawText("pinkatin", monw-495, monh-145, 120, rl.Fade(rl.Magenta, playerfade))
	}

	birdactive = true
	weaponv2 := rl.NewVector2(introv2.X+16, introv2.Y+8)
	if introlr {
		// left
		playerdirection = 4
		introv2.X--
		introenemyv2.X++
		if introv2.X <= -32 {
			introlr = false
			introenemy = rInt(1, 5)
			chooseintroweapon = rInt(1, 5)
			introenemycolor = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)

			introv2.Y = rFloat32(10, (monh/2)/8)
			introenemyv2.Y = rFloat32(10, (monh/2)/8)
		}
		weaponv2 = rl.NewVector2(introv2.X-12, introv2.Y-4)
	} else {
		// right
		playerdirection = 2
		introv2.X++
		introenemyv2.X--
		if introv2.X >= float32((monw/8)+48) {
			introlr = true
			introenemy = rInt(1, 5)
			chooseintroweapon = rInt(1, 5)
			introenemycolor = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
			introv2.Y = rFloat32(10, (monh/2)/8)
			introenemyv2.Y = rFloat32(10, (monh/2)/8)
		}
		weaponv2 = rl.NewVector2(introv2.X+12, introv2.Y-4)
	}

	// weapon bob
	if onoff6 {
		weaponv2.Y--
	} else {
		weaponv2.Y++
	}

	switch introenemy {
	case 1:
		birdactive = true
		ghostactive = false
		spikesactive = false
		radishactive = false
	case 2:
		birdactive = false
		ghostactive = true
		spikesactive = false
		radishactive = false
	case 3:
		birdactive = false
		ghostactive = false
		spikesactive = true
		radishactive = false
	case 4:
		birdactive = false
		ghostactive = false
		spikesactive = false
		radishactive = true
	}

	// draw intro characters
	rl.BeginMode2D(camera8x)
	if introlr {
		rl.DrawTextureRec(imgs, playerleft, introv2, rl.Fade(rl.Magenta, playerfade))
		switch chooseintroweapon {
		case 1:
			weaponv2.Y += 7
			weaponv2.X += 2
			rl.DrawTextureRec(imgs, uzzil, weaponv2, rl.White)
		case 2:
			weaponv2.Y += 7
			rl.DrawTextureRec(imgs, pumpactionl, weaponv2, rl.White)
		case 3:
			rl.DrawTextureRec(imgs, bigswordl, weaponv2, rl.White)
		case 4:
			rl.DrawTextureRec(imgs, axel, weaponv2, rl.White)
		}

		switch introenemy {
		case 1:
			rl.DrawTextureRec(imgs, birdr, introenemyv2, introenemycolor)
		case 2:
			rl.DrawTextureRec(imgs, ghostr, introenemyv2, introenemycolor)
		case 3:
			rl.DrawTextureRec(imgs, spikesfloor, introenemyv2, introenemycolor)
		case 4:
			rl.DrawTextureRec(imgs, radishr, introenemyv2, introenemycolor)
		}
	} else {
		rl.DrawTextureRec(imgs, playerright, introv2, rl.Fade(rl.Magenta, playerfade))
		switch chooseintroweapon {
		case 1:
			weaponv2.Y += 8
			rl.DrawTextureRec(imgs, uzzir, weaponv2, rl.White)
		case 2:
			weaponv2.Y += 8
			rl.DrawTextureRec(imgs, pumpaction, weaponv2, rl.White)
		case 3:
			rl.DrawTextureRec(imgs, bigswordr, weaponv2, rl.White)
		case 4:
			rl.DrawTextureRec(imgs, axe, weaponv2, rl.White)
		}
		switch introenemy {
		case 1:
			rl.DrawTextureRec(imgs, birdl, introenemyv2, introenemycolor)
		case 2:
			rl.DrawTextureRec(imgs, ghostl, introenemyv2, introenemycolor)
		case 3:
			rl.DrawTextureRec(imgs, spikesfloor, introenemyv2, introenemycolor)
		case 4:
			rl.DrawTextureRec(imgs, radishl, introenemyv2, introenemycolor)
		}
	}
	rl.EndMode2D()

}
func drawspecialtext() { // MARK: drawspecialtext

	if cameraspecialtextzoomon {
		cameraspecialtextzoom -= 0.01
		cameraspecialtext.Target.Y -= 10
		if cameraspecialtextzoom < 1.0 {
			cameraspecialtextzoomon = false
		}
	} else {
		cameraspecialtextzoom += 0.01
		cameraspecialtext.Target.Y += 10
		if cameraspecialtextzoom > 1.1 {
			cameraspecialtextzoomon = true
		}
	}

	rl.BeginMode2D(cameraspecialtext)
	if pumpactionactive {

		pumpactiontimerTEXT := strconv.Itoa(pumpactiontimer)

		rl.DrawText("pump action ends in", (monw/7)-6, ((monh/4)*3)+6, 100, rl.Black)
		rl.DrawText("pump action ends in", (monw/7)-3, ((monh/4)*3)+3, 100, rl.Purple)
		rl.DrawText("pump action ends in", (monw/7)-1, ((monh/4)*3)+1, 100, rl.Black)
		rl.DrawText("pump action ends in", monw/7, (monh/4)*3, 100, rl.Magenta)

		rl.DrawText(pumpactiontimerTEXT, (monw-(monw/3-30))-6, ((monh/4)*3)+6, 120, rl.Black)
		rl.DrawText(pumpactiontimerTEXT, (monw-(monw/3-30))-3, ((monh/4)*3)+3, 120, rl.Purple)
		rl.DrawText(pumpactiontimerTEXT, (monw-(monw/3-30))-1, ((monh/4)*3)+1, 120, rl.Black)
		rl.DrawText(pumpactiontimerTEXT, monw-(monw/3-30), (monh/4)*3, 120, rl.Magenta)
	} else if icefreezeactive {

		icefreezetimerTEXT := strconv.Itoa(icefreezetimer)

		rl.DrawText("enemy freeze ends in", (monw/7)-6, ((monh/4)*3)+6, 100, rl.Black)
		rl.DrawText("enemy freeze ends in", (monw/7)-3, ((monh/4)*3)+3, 100, rl.Purple)
		rl.DrawText("enemy freeze ends in", (monw/7)-1, ((monh/4)*3)+1, 100, rl.Black)
		rl.DrawText("enemy freeze ends in", monw/7, (monh/4)*3, 100, rl.Magenta)

		rl.DrawText(icefreezetimerTEXT, (monw-(monw/4+30))-6, ((monh/4)*3)+6, 120, rl.Black)
		rl.DrawText(icefreezetimerTEXT, (monw-(monw/4+30))-3, ((monh/4)*3)+3, 120, rl.Purple)
		rl.DrawText(icefreezetimerTEXT, (monw-(monw/4+30))-1, ((monh/4)*3)+1, 120, rl.Black)
		rl.DrawText(icefreezetimerTEXT, monw-(monw/4+30), (monh/4)*3, 120, rl.Magenta)

	}
	rl.EndMode2D()

	cameraspecialtext.Zoom = cameraspecialtextzoom

}
func update() { // MARK: update
	timers()
	input()
	environment()
	horizvert()
	jumpplayer()
	collect()
	updatedrawblock()
	drawinfobar()
	moveanimals()
	movewater()
	createenemies()
	moveenemies()
	moveinteractables()
	moveprojectiles()
	moveupdatearea()
	updatearea()

	if inventoryon && !optionson && !introscreenon {
		drawinventory()
	}

	if optionson {
		drawoptions()
	}

	if questtotal >= 3 && !placekey {
		endkeyactive = true
	}
	if endkeyactive && !placekey {
		createendkey()
		placekey = true
	}
	if endkeyactive {
		updateendkey()
	}

	if levellayout[player+levelw] != "." {
		playerrollactive = false
	}

	if questdetails {
		drawquestdetails()
	}
	if mapcollected {
		mapzoomactive()
	}
	if screenshakeactive {
		screenshake()
	}

	if meteoractive {
		movemeteor()
	}
	if questactive {
		updatequest()
	}

	if fireballactive {
		movefireball()
	}

	if discoveryinfoon {
		drawdiscoveryinfobox()
	}
	if discoveryon {
		drawdiscoveries()
	}
	if flosson {
		flossmovement()
	}

	if introscreenon {
		drawintro()
		pauseon = true
		if rl.IsKeyPressed(rl.KeySpace) {
			introscreenon = false
			pauseon = false
		}
	}

	animations()
	if !moveoff && !flipscreen {
		if camera.Zoom == 2.0 {
			camera.Target.X = float32(playerx - (monw / 4))
			camera.Target.Y = float32(playery - (monh / 4))
		} else if camera.Zoom == 3.0 {
			camera.Target.X = float32(playerx - (monw / 6))
			camera.Target.Y = float32(playery - (monh / 6))
		} else if camera.Zoom == 4.0 {
			camera.Target.X = float32(playerx - (monw / 8))
			camera.Target.Y = float32(playery - (monh / 8))
		}
	}

	if drawstarton {
		drawstartscreen()
		if rl.IsKeyPressed(rl.KeySpace) {
			drawstarton = false
			introscreenon = true
		}
	}
	if drawselectscreenon {
		drawselectscreen()
	}

	fx()

	if debugon {
		debug()
	}

}
func updatearea() { // MARK: updatearea
	updatew = draww * 2
	updateh = drawh * 2
	updatea = updateh * updatew
	block := player
	block -= draww
	block -= drawh * levelw
	updateblock = block
	block = player
	block += draww
	block += drawh * levelw
	updateblockend = block
}
func updateendkey() { // MARK: updateendkey

	questh, questv := endkeyblock/levelw, endkeyblock%levelw
	endkeydirection := 0

	if playerh-questh > 0 && playerh-questh < 15 && questv > playerv {
		endkeydirection = 4
	} else if questh-playerh > 0 && questh-playerh < 15 && questv > playerv {
		endkeydirection = 4
	}

	if playerh-questh > 0 && playerh-questh < 15 && questv < playerv {
		endkeydirection = 8
	} else if questh-playerh > 0 && questh-playerh < 15 && questv < playerv {
		endkeydirection = 8
	}

	if playerh-questh > 15 {
		if playerv-questv > 0 && playerv-questv < 15 {
			endkeydirection = 2
		} else if questv-playerv > 0 && questv-playerv < 15 {
			endkeydirection = 2
		} else if playerv > questv && playerv-questv > 15 {
			endkeydirection = 1
		} else if questv > playerv && questv-playerv > 15 {
			endkeydirection = 3
		}
	} else if questh-playerh > 15 {
		if playerv-questv > 0 && playerv-questv < 15 {
			endkeydirection = 6
		} else if questv-playerv > 0 && questv-playerv < 15 {
			endkeydirection = 6
		} else if questv > playerv && questv-playerv > 15 {
			endkeydirection = 5
		} else if playerv > questv && playerv-questv > 15 {
			endkeydirection = 7
		}
	}

	switch endkeydirection {
	case 1:
		endkeyarrowrotation = 315.0
	case 2:
		endkeyarrowrotation = 0.0
	case 3:
		endkeyarrowrotation = 45.0
	case 4:
		endkeyarrowrotation = 90.0
	case 5:
		endkeyarrowrotation = 135.0
	case 6:
		endkeyarrowrotation = 180.0
	case 7:
		endkeyarrowrotation = 225.0
	case 8:
		endkeyarrowrotation = 270.0
	}

}
func updatequest() { // MARK: updatequest

	questh, questv := activequest.block/levelw, activequest.block%levelw
	questdirection := 0

	if playerh-questh > 0 && playerh-questh < 15 && questv > playerv {
		questdirection = 4
	} else if questh-playerh > 0 && questh-playerh < 15 && questv > playerv {
		questdirection = 4
	}

	if playerh-questh > 0 && playerh-questh < 15 && questv < playerv {
		questdirection = 8
	} else if questh-playerh > 0 && questh-playerh < 15 && questv < playerv {
		questdirection = 8
	}

	if playerh-questh > 15 {
		if playerv-questv > 0 && playerv-questv < 15 {
			questdirection = 2
		} else if questv-playerv > 0 && questv-playerv < 15 {
			questdirection = 2
		} else if playerv > questv && playerv-questv > 15 {
			questdirection = 1
		} else if questv > playerv && questv-playerv > 15 {
			questdirection = 3
		}
	} else if questh-playerh > 15 {
		if playerv-questv > 0 && playerv-questv < 15 {
			questdirection = 6
		} else if questv-playerv > 0 && questv-playerv < 15 {
			questdirection = 6
		} else if questv > playerv && questv-playerv > 15 {
			questdirection = 5
		} else if playerv > questv && playerv-questv > 15 {
			questdirection = 7
		}
	}

	switch questdirection {
	case 1:
		questarrowrotation = 315.0
	case 2:
		questarrowrotation = 0.0
	case 3:
		questarrowrotation = 45.0
	case 4:
		questarrowrotation = 90.0
	case 5:
		questarrowrotation = 135.0
	case 6:
		questarrowrotation = 180.0
	case 7:
		questarrowrotation = 225.0
	case 8:
		questarrowrotation = 270.0
	}

}
func timers() { // MARK: timers

	if introtimer5on {
		introtimer5frame++
		if introtimer5frame%30 == 0 {
			introtimer5--
		}
		if introtimer5 <= 0 {
			drawstarton = false
			introscreenon = true
			introtimer5on = false
		}
	}

	if introtimer4on {
		introtimer4frame++
		if introtimer4frame%30 == 0 {
			introtimer4--
		}
		if introtimer4 <= 0 {
			introcircleson = true
			for a := 0; a < len(introcirlcesv2layout); a++ {
				x := rFloat32(0, monw)
				y := rFloat32(0, monh)

				xy := rl.NewVector2(x, y)

				newcircle := circlev2{}

				newcircle.center = xy
				newcircle.radius = 0
				newcircle.color = rl.Black
				introcirlcesv2layout[a] = newcircle
			}
			intronicholasimontexton2 = true
			introtimer4on = false
			introtimer5on = true
		}
	}

	if introtimer1on {
		introtimer1frame++
		if introtimer1frame%30 == 0 {
			introtimer1--
		}
		if introtimer1 <= 0 {
			rayliblogoon = true
			introtimer2on = true
			introtimer1on = false
			startlogofadeon2 = true
		}
	}
	if introtimer3on {
		introtimer3frame++
		if introtimer3frame%30 == 0 {
			introtimer3--
		}
		if introtimer3 <= 0 {
			intropinkbackon = true
		}
	}

	if drawstarton {
		if startlogofadeon {
			startlogofade += 0.02
			if startlogofade >= 1.0 {
				startlogofadeon = false
			}
		}
		if startlogofadeon2 {
			startlogofade -= 0.02
			if startlogofade <= 0.0 {
				startlogofadeon2 = false
			}
		}
	}

	if framecount%6 == 0 {
		if onoff6 {
			onoff6 = false
		} else {
			onoff6 = true
		}
	}
	if framecount%15 == 0 {
		if onoff15 {
			onoff15 = false
		} else {
			onoff15 = true
		}
	}
	if framecount%30 == 0 {
		if onoff30 {
			onoff30 = false
		} else {
			onoff30 = true
		}
	}
	if framecount%3 == 0 {
		if onoff3 {
			onoff3 = false
		} else {
			onoff3 = true
		}
	}
	if framecount%10 == 0 {
		if onoff10 {
			onoff10 = false
		} else {
			onoff10 = true
		}
	}

	if inventorypause {
		inventorycount++
		if inventorycount >= 30 {
			inventorypause = false
		}

	}

	if framecount%10 == 0 {
		if shopflicker {
			shopflicker = false
		} else {
			shopflicker = true
		}
	}

	if icefreezeactive {
		icefreezeframe++
		if icefreezeframe%30 == 0 {
			icefreezetimer--
			if icefreezetimer <= 0 {
				icefreezeactive = false
			}
		}
	}

	if shopkeepertalkon {
		shopkeepertalkframe++
		if shopkeepertalkframe%30 == 0 {
			shopkeepertalktimer--
			if shopkeepertalktimer <= 0 {
				shopkeepertalkon = false
				shopkeepertalktimer2 = rInt(5, 15)
				shopkeepertalkframe = 0
			}
		}
	} else {
		shopkeepertalkframe++
		if shopkeepertalkframe%30 == 0 {
			shopkeepertalktimer2--
			if shopkeepertalktimer2 <= 0 {
				shopkeepertalkon = true
				shopkeepertalktimer = rInt(5, 15)
				shopkeepertalkframe = 0
				shopkeepertalk = rInt(0, 13)
			}
		}
	}

	if pumpactionactive {
		pumpactionframe++
		if pumpactionframe%30 == 0 {
			pumpactiontimer--
			if pumpactiontimer <= 0 {
				pumpactionactive = false
			}
		}
	}

	if underwateron {
		if !underwatertimercreated {
			underwatertimer = 3
			underwatertimercreated = true
		}
		if framecount%30 == 0 {
			underwatertimer--
			if underwatertimer <= 0 {
				hp--
				underwatertimer = 3
			}
		}
		if waterlayout[player] == "" {
			underwateron = false
			underwatertimercreated = false
		}
	}

	if meteoractive {
		if createmeteor == false {
			for {
				meteorpoint = player + rInt(-50, 50)
				meteorpoint += rInt(-30, 30) * levelw

				if levellayout[meteorpoint] == "." && levellayout[meteorpoint+levelw] == "." {
					if meteorpoint < levela-(levelw*50) {
						meteorpoint += levelw
					}
				}
				if levellayout[meteorpoint] == "_" && levellayout[meteorpoint+levelw] == "_" && levellayout[meteorpoint-levelw] == "_" {
					if meteorpoint > levelw*50 {
						meteorpoint -= levelw
					} else {
						meteorpoint = player + rInt(-50, 50)
						meteorpoint += rInt(-30, 30) * levelw
					}
				}

				if levellayout[meteorpoint] == "." && levellayout[meteorpoint+levelw] == "_" {
					break
				}
			}

			block := meteorpoint
			blockh := block / levelw
			block -= levelw * (blockh - 10)

			weaponslayout[block] = "meteortl"
			weaponslayout[block+1] = "meteor"
			weaponslayout[block+levelw] = "meteor"
			weaponslayout[(block+1)+levelw] = "meteor"

			meteortimer = rInt(10, 50)
			createmeteor = true
		}
	}

	if createmeteor == true {
		if framecount%30 == 0 {
			meteortimer--
			if meteortimer <= 0 {
				createmeteor = false
			}
		}
	}

	if pegasustimer != 0 {
		if framecount%30 == 0 {
			pegasustimer--
			if pegasustimer <= 0 {
				gravityoff = false
				for a := 0; a < len(powerupscollected); a++ {
					checkpowerup := powerupscollected[a]
					if checkpowerup.name == "pegasus" {
						powerupscollected[a] = powerup{}
						powerupcount--
					}
				}
			}
		}

	}

	if powerupblockmovetimer > 0 {
		if framecount%3 == 0 {
			powerupblockmovetimer--
			if powerupblockmovetimer <= 0 {
				powerupblockmoveon = false
			}
		}
	}

	if framecount%150 == 0 {
		resets(2)
	}

	if conversationon {
		conversationframe++

		if conversationframe < 60 {
			conversation1 = true
			conversation2 = false
			conversation3 = false
			conversation4 = false
			conversation5 = false
			conversation6 = false
		} else if conversationframe > 60 && conversationframe < 120 {
			conversation1 = false
			conversation2 = true
			conversation3 = false
			conversation4 = false
			conversation5 = false
			conversation6 = false
		} else if conversationframe > 120 && conversationframe < 180 {
			conversation1 = false
			conversation2 = true
			conversation3 = false
			conversation4 = false
			conversation5 = false
			conversation6 = false
		} else if conversationframe > 180 && conversationframe < 240 {
			conversation1 = false
			conversation2 = false
			conversation3 = true
			conversation4 = false
			conversation5 = false
			conversation6 = false
		} else if conversationframe > 240 && conversationframe < 300 {
			conversation1 = false
			conversation2 = false
			conversation3 = false
			conversation4 = true
			conversation5 = false
			conversation6 = false
		} else if conversationframe > 300 && conversationframe < 360 {
			conversation1 = false
			conversation2 = false
			conversation3 = false
			conversation4 = false
			conversation5 = true
			conversation6 = false
		} else if conversationframe > 360 && conversationframe < 420 {
			conversation1 = false
			conversation2 = false
			conversation3 = false
			conversation4 = false
			conversation5 = false
			conversation6 = true
		} else if conversationframe > 420 {
			conversationframe = 0
			conversationon = false
			conversation1, conversation2, conversation3, conversation4, conversation5, conversation6 = false, false, false, false, false, false

		}

	}

	if discoveryon {
		if framecount%30 == 0 {
			discoverytimer--
		}
		if discoverytimer == 0 {
			discoveryon = false
			discoverytimer = 4
		}
	}

	if playerrunon {
		emoteswitch = false
		idletimer = 0
		emoteon = false
	} else {
		if framecount%30 == 0 {
			idletimer++
		}
	}

	if idletimer > 3 {
		playerlron = true
		emoteon = true
		if !emoteswitch {
			chooseemote = rInt(1, 7)
			emoteswitch = true
		}
	} else {
		playerlron = false
	}

	if jumppause != 0 {
		if framecount%30 == 0 {
			jumppause--
		}
	}

}
func animations() { // MARK: animations

	// inventory fade
	if inventoryfadeon {
		inventoryfade += 0.03
		if inventoryfade >= 0.5 {
			inventoryfadeon = false
		}
	} else {
		inventoryfade -= 0.03
		if inventoryfade <= 0.2 {
			inventoryfadeon = true
		}
	}

	// shield fade
	if framecount%3 == 0 {
		if shieldfadeon {
			shieldfade += 0.05
			if shieldfade >= 0.4 {
				shieldfadeon = false
			}
		} else {
			shieldfade -= 0.05
			if shieldfade <= 0.1 {
				shieldfadeon = true
			}
		}
	}
	// shop items rotation
	shopitemrotation1 += 2
	shopitemrotation2 += 2
	shopitemrotation3 += 2

	// player roll
	if playerrollactive {
		if playerdirection == 2 {
			playerrolldegrees += 12.0
		} else if playerdirection == 4 {
			playerrolldegrees -= 12.0
		}
	}

	// powerups fades
	if fireballactive {
		if fireballfadeon {
			fireballfade += 0.08
			if fireballfade >= 1.0 {
				fireballfadeon = false
			}
		} else {
			fireballfade -= 0.08
			if fireballfade < 0.6 {
				fireballfadeon = true
			}
		}
	}
	// enemies
	if pigactive {
		pigl.X += 36
		pigr.X += 36

		if pigr.X > 396 {
			pigr.X = 0
		}
		if pigl.X > 840 {
			pigl.X = 440
		}
	}
	if ghostactive {

		ghostl.X += 44
		ghostr.X += 44

		if ghostl.X > 846 {
			ghostl.X = 442
		}
		if ghostr.X > 402 {
			ghostr.X = 0
		}
		if ghostfadeon {
			ghostfade += 0.01
			if ghostfade >= 1.0 {
				ghostfadeon = false
			}
		} else {
			ghostfade -= 0.01
			if ghostfade <= 0 {
				ghostfadeon = true
			}
		}
	}

	if bigbatactive {
		bigbatl.X += 46
		bigbatr.X += 46
		if bigbatl.X > 828 {
			bigbatl.X = 549
		}
		if bigbatr.X > 822 {
			bigbatr.X = 544
		}
	}

	if rockactive {
		rockl.X += 38
		rockr.X += 38
		if rockr.X > 496 {
			rockr.X = 0
		}
		if rockl.X > 496 {
			rockl.X = 0
		}
	}

	if radishactive {
		radishl.X += 30
		radishr.X += 30
		if radishl.X > 158 {
			radishl.X = 6
		}
		if radishr.X > 348 {
			radishr.X = 195
		}
	}

	if plantactive {
		plantl.X += 44
		plantr.X += 44
		if plantl.X > 455 {
			plantl.X = 4
		}
		if plantr.X > 936 {
			plantr.X = 492
		}
	}

	if slimeactive {
		slimel.X += 44
		slimer.X += 44
		slimelroof.X += 44
		slimerroof.X += 44
		if slimel.X > 498 {
			slimel.X = 97
		}
		if slimer.X > 500 {
			slimer.X = 102
		}
		if slimelroof.X > 1438 {
			slimelroof.X = 1036
		}
		if slimerroof.X > 1440 {
			slimerroof.X = 1041
		}
	}
	if rabbitactive {
		rabbitl.X += 34
		rabbitr.X += 34
		if rabbitl.X > 486 {
			rabbitl.X = 110
		}
		if rabbitr.X > 492 {
			rabbitr.X = 114
		}
	}

	if framecount%3 == 0 {
		if snailactive {
			snaill.X += 38
			snailr.X += 38
			if snaill.X > 346 {
				snaill.X = 0
			}
			if snailr.X > 728 {
				snailr.X = 385
			}
		}

		if birdactive {
			birdl.X += 32
			if birdl.X > 490 {
				birdl.X = 229
			}
			birdr.X += 32
			if birdr.X > 790 {
				birdr.X = 523
			}
		}
		if spikesactive {
			spikesfloor.X += 44
			spikesroof.X += 44
			spikesleft.Y += 44
			spikesright.Y += 44

			if spikesleft.Y > 950 {
				spikesleft.Y = 669
			}
			if spikesright.Y > 950 {
				spikesright.Y = 669
			}
			if spikesroof.X > 682 {
				spikesroof.X = 368
			}
			if spikesfloor.X > 316 {
				spikesfloor.X = 2
			}
		}
		if beeactive {
			beesting.X += 36
			if beesting.X > 370 {
				beesting.X = 115
			}
		}
		if skullactive {
			floatingskull.X += 52
			if floatingskull.X > 460 {
				floatingskull.X = 90
			}
		}
		if mushroomactive {
			mushroom.X += 32
			if mushroom.X > 482 {
				mushroom.X = 0
			}
		}
	}

	// door fade
	if doorfadeon {
		doorfade += 0.08
		if doorfade > 0.8 {
			doorfadeon = false
		}
	} else {
		doorfade -= 0.08
		if doorfade < 0.4 {
			doorfadeon = true
		}
	}

	// move background vines chains
	backychange++
	if backychange == 16 {
		backychange = 0
	}
	// move shopkeeper
	if framecount%5 == 0 {
		if shopkeeperlr {
			if levellayout[shopkeeperblock1+1] == "." {
				characterslayout[shopkeeperblock1] = ""
				shopkeeperblock1++
				characterslayout[shopkeeperblock1] = "shopkeeper"
			} else {
				shopkeeperlr = false
			}
		} else {
			if levellayout[shopkeeperblock1-1] == "." {
				characterslayout[shopkeeperblock1] = ""
				shopkeeperblock1--
				characterslayout[shopkeeperblock1] = "shopkeeper"
			} else {
				shopkeeperlr = true
			}
		}
	}
	// discovery box
	if discoveryon {
		if discoveryrecx != 0 {
			discoveryimgx += 5
			discoveryrecx += 20
			discoverytext1shadowx += 20
			discoverytext1x += 20
			discoverytext2shadowx += 20
			discoverytext2x += 20
		}
	}
	// bat rotation
	if framecount%2 == 0 {
		if flipcoin() {
			batrotation += 3.0
			if batrotation >= 30.0 {
				batrotation -= 6.0
			}
		} else {
			batrotation -= 3.0
			if batrotation <= -30.0 {
				batrotation += 6.0
			}
		}
	}

	// hp fade
	if hpfadeon {
		hpfade += 0.05
		if hpfade >= 0.7 {
			hpfadeon = false
		}
	} else {
		hpfade -= 0.05
		if hpfade <= 0.0 {
			hpfadeon = true
		}
	}

	// weapon movement idle
	if playerrunon {
		if weaponrotationswitch {
			weaponrotation -= 2.0
			if weaponrotation < 0.0 {
				weaponrotationswitch = false
			}
		} else {
			weaponrotation += 2.0
			if weaponrotation > 10.0 {
				weaponrotationswitch = true
			}
		}
	} else if swinglefton {
		weaponrotation = -45.0
	} else if swingrighton {
		weaponrotation = 45.0
	} else {
		weaponrotation = 0.0
	}

	if framecount%6 == 0 {

		// gemcolors

		if gemcolor == rl.Red {
			gemcolor = rl.Green
		} else if gemcolor == rl.Green {
			gemcolor = rl.Yellow
		} else if gemcolor == rl.Yellow {
			gemcolor = rl.Blue
		} else if gemcolor == rl.Blue {
			gemcolor = rl.Aqua
		} else if gemcolor == rl.Aqua {
			gemcolor = rl.Orange
		} else if gemcolor == rl.Orange {
			gemcolor = rl.Red
		}

		// spring
		if springon {
			spring.X -= 16
			if spring.X == 1 {
				springon = false
			}
		} else {
			spring.X += 16
			if spring.X == 33 {
				springon = true
			}
		}

		// emotes
		if emoteon {
			switch chooseemote {
			case 1:
				emote1.X += 32
				if emote1.X > 64 {
					emote1.X = 0
				}
			case 2:
				emote2.X += 32
				if emote2.X > 160 {
					emote2.X = 96
				}
			case 3:
				emote3.X += 32
				if emote3.X > 258 {
					emote3.X = 194
				}
			case 4:
				emote4.X += 32
				if emote4.X > 354 {
					emote4.X = 290
				}
			case 5:
				emote5.X += 32
				if emote5.X > 452 {
					emote5.X = 388
				}
			case 6:
				emote6.X += 32
				if emote6.X > 550 {
					emote6.X = 486
				}

			}

		}

	}

	if framecount%3 == 0 {
		// shopkeeper
		shopkeeper.X += 21
		shopkeeperl.X += 21
		if shopkeeper.X > 86 {
			shopkeeper.X = 0
		}
		if shopkeeperl.X > 198 {
			shopkeeperl.X = 112
		}
		// coin gem
		coin.X += 16
		if coin.X > 64 {
			coin.X = 0
		}
		gem.X += 16
		if gem.X > 155 {
			gem.X = 105
		}
		// player run
		if playerrunon || introscreenon {
			if playerdirection == 2 { // right
				playerright.X += 16
				if playerright.X > 64 {
					playerright.X = 0
				}

			} else if playerdirection == 4 { // left
				playerleft.X -= 16
				if playerleft.X < 16 {
					playerleft.X = 81
				}
			}
		}

		// emote fade
		if emotefadeon {
			emotefade += 0.05
			if emotefade > 0.5 {
				emotefadeon = false
			}
		} else {
			emotefade -= 0.05
			if emotefade < 0.3 {
				emotefadeon = true
			}
		}

		// player color fade
		if playerfadeon {
			playerfade += 0.1
			if playerfade >= 1.0 {
				playerfadeon = false
			}
		} else {
			playerfade -= 0.1
			if playerfade <= 0.7 {
				playerfadeon = true
			}
		}
	}

	if framecount%8 == 0 {
		// player bob
		if playerychange == 0 {
			playerychange = 1
		} else {
			playerychange = 0
		}

	}
	// player idle left right
	if !playerrunon && !introscreenon && playerlron {
		if framecount%15 == 0 {
			if flipcoin() {
				playerdirection = 4
			} else {
				playerdirection = 2
			}
		}
	}

}
func fx() { // MARK: fx
	// back drift
	if framecount%3 == 0 {
		if backdrifton {
			backdriftx--
			backdrifty--
			if backdriftx == -4 {
				backdrifton = false
			}
		} else {
			backdriftx++
			backdrifty++
			if backdriftx == 4 {
				backdrifton = true
			}
		}
	}
	// scan lines
	if scanlineson {
		linex := 0
		liney := 2
		for a := 0; a < monh; a++ {
			rl.DrawLine(linex, liney, monw, liney, rl.Fade(rl.Black, 0.3))
			liney += 4
			a += 2
		}
	}
	// pixel noise
	for a := 0; a < 100; a++ {
		rl.DrawRectangle(rInt(10, monw-10), rInt(10, monh-10), 2, 2, rl.Black)
	}

}
func mapzoomactive() { // MARK: screenshake

	camera.Zoom = 2.0

	if framecount%30 == 0 {
		mapzoomtimer--
		if mapzoomtimer <= 0 {
			camera.Zoom = 4.0
			mapcollected = false
		}
	}

}
func screenshake() { // MARK: screenshake
	if timer1 > 0 {
		camera.Rotation -= 0.2
		timer1--
	} else if timer2 > 0 {
		camera.Rotation += 0.2
		timer2--
	} else if timer3 > 0 {
		camera.Rotation -= 0.2
		timer3--
	} else {
		timer1 = 2
		timer2 = 4
		timer3 = 2
		screenshakeactive = false
	}
}
func moveplayer(direction int) { // MARK: moveplayer

	switch direction {
	case 2:
		if playerv < levelw-32 {
			if levellayout[player+1] == "." {
				player++
			}
		}
	case 4:
		if playerv > 32 {
			if levellayout[player-1] == "." {
				player--
			}
		}
	}

}
func jumpplayer() { // MARK: jumpplayer
	if jumpactive {
		if playerh > 2 {
			if levellayout[player-levelw] == "." {
				player -= levelw
				jumpcount++
			} else if levellayout[player-levelw] == "%" {
				powerupblock(player - levelw)
				jumpcount = 0
				fallactive = true
				jumpactive = false
			} else {
				jumpcount = 0
				fallactive = true
				jumpactive = false
			}
		}
		if jumpcount >= jumpheight {
			jumpcount = 0
			fallactive = true
			jumpactive = false
		}
	}

	if jumpactivespring {
		if playerh > 2 {
			if levellayout[player-levelw] == "." {
				player -= levelw
				jumpcount++
			} else if levellayout[player-levelw] == "%" {
				powerupblock(player - levelw)
				jumpcount = 0
				fallactive = true
				jumpactive = false
			} else {
				jumpcount = 0
				fallactive = true
				jumpactivespring = false
			}
		}
		if jumpcount >= jumpheightspring {
			jumpcount = 0
			fallactive = true
			jumpactivespring = false
		}

	}

}
func moveupdatearea() { // MARK: moveupdatearea
	count := 0
	block := updateblock
	// up left
	for a := 0; a < updatea; a++ {
		checkextras := extraslayout[block]
		checktraps := trapslayout[block]
		checkrailway := railwaylayout[block]

		if checkrailway != "" {
			if checkrailway == "cartl" {
				if framecount%3 == 0 {
					if railwaylayout[(block-1)+levelw] == "track" {
						railwaylayout[block] = ""
						railwaylayout[block-1] = "cartl"
					} else {
						railwaylayout[block] = "cartr"
					}
				}
			}
		}
		if checktraps != "" {
			if framecount%3 == 0 {
				if checktraps == "spike4" {
					trapslayout[block] = ""
					trapslayout[block-levelw] = "spike5"
				} else if checktraps == "spike5" {
					trapslayout[block] = ""
					trapslayout[block-levelw] = "spike6"
				} else if checktraps == "spike6" {
					trapslayout[block] = "spike3"
				}
			}
		}
		if checkextras != "" {
			if extraslayout[block] == "bookcase" {
				if player == block || player == block+1 || player == block-1 {
					bookcaseactive = true
				} else {
					bookcaseactive = false
				}
			}
			if extraslayout[block] == "ufo1" {
				if rolldice()+rolldice() == 12 {
					extraslayout[block] = "ufo1r"
				} else if rolldice() == 6 {
					if flipcoin() {
						if levellayout[block-levelw] == "." {
							extraslayout[block] = ""
							extraslayout[block-levelw] = "ufo1"
						}
					} else {
						if levellayout[block+levelw] == "." {
							extraslayout[block] = ""
							extraslayout[block+levelw] = "ufo1"
						}
					}
				} else {
					if levellayout[block-1] == "." {
						extraslayout[block] = ""
						extraslayout[block-1] = "ufo1"
					} else {
						extraslayout[block] = "ufo1r"
					}
				}
			}
			if extraslayout[block] == "ufo2" {
				if rolldice()+rolldice() == 12 {
					extraslayout[block] = "ufo2r"
				} else if rolldice() == 6 {
					if flipcoin() {
						if levellayout[block-levelw] == "." {
							extraslayout[block] = ""
							extraslayout[block-levelw] = "ufo2"
						}
					} else {
						if levellayout[block+levelw] == "." {
							extraslayout[block] = ""
							extraslayout[block+levelw] = "ufo2"
						}
					}
				} else {
					if levellayout[block-1] == "." {
						extraslayout[block] = ""
						extraslayout[block-1] = "ufo2"
					} else {
						extraslayout[block] = "ufo2r"
					}
				}
			}
		}
		count++
		block++
		if count == updatew {
			count = 0
			block -= updatew
			block += levelw
		}

	}
	// down right
	count = 0
	block = updateblockend
	for a := 0; a < updatea; a++ {
		checkextras := extraslayout[block]
		checktraps := trapslayout[block]
		checkrailway := railwaylayout[block]

		if checkrailway != "" {
			if checkrailway == "cartr" {
				if framecount%3 == 0 {
					if railwaylayout[(block+1)+levelw] == "track" {
						railwaylayout[block] = ""
						railwaylayout[block+1] = "cartr"
					} else {
						railwaylayout[block] = "cartl"
					}
				}
			}
		}

		if checktraps != "" {
			if framecount%3 == 0 {
				if checktraps == "spike3" {
					trapslayout[block] = ""
					trapslayout[block+levelw] = "spike2"
				} else if checktraps == "spike2" {
					trapslayout[block] = ""
					trapslayout[block+levelw] = "spike1"
				} else if checktraps == "spike1" {
					trapslayout[block] = "spike4"
				}
			}
		}
		if checkextras != "" {
			if extraslayout[block] == "ufo1r" {
				if rolldice()+rolldice() == 12 {
					extraslayout[block] = "ufo1"
				} else if rolldice() == 6 {
					if flipcoin() {
						if levellayout[block-levelw] == "." {
							extraslayout[block] = ""
							extraslayout[block-levelw] = "ufo1r"
						}
					} else {
						if levellayout[block+levelw] == "." {
							extraslayout[block] = ""
							extraslayout[block+levelw] = "ufo1r"
						}
					}
				} else {
					if levellayout[block+1] == "." {
						extraslayout[block] = ""
						extraslayout[block+1] = "ufo1r"
					} else {
						extraslayout[block] = "ufo1"
					}
				}
			}
			if extraslayout[block] == "ufo2r" {
				if rolldice()+rolldice() == 12 {
					extraslayout[block] = "ufo2"
				} else if rolldice() == 6 {
					if flipcoin() {
						if levellayout[block-levelw] == "." {
							extraslayout[block] = ""
							extraslayout[block-levelw] = "ufo2r"
						}
					} else {
						if levellayout[block+levelw] == "." {
							extraslayout[block] = ""
							extraslayout[block+levelw] = "ufo2r"
						}
					}
				} else {
					if levellayout[block+1] == "." {
						extraslayout[block] = ""
						extraslayout[block+1] = "ufo2r"
					} else {
						extraslayout[block] = "ufo2"
					}
				}
			}
		}

		count++
		block--
		if count == updatew {
			count = 0
			block += updatew
			block -= levelw
		}
	}

}
func movemeteor() { // MARK: movemeteor

	count := 0
	block := updateblockend
	for a := 0; a < updatea; a++ {

		if weaponslayout[block] == "meteortl" {

			meteorh := block / levelw
			meteorpointh := meteorpoint / levelw

			if meteorpointh > meteorh {
				weaponslayout[block] = ""
				weaponslayout[block+1] = ""
				weaponslayout[block+levelw] = ""
				weaponslayout[(block + 1 + levelw)] = ""

				weaponslayout[block+levelw] = "meteortl"
				weaponslayout[(block+levelw)+1] = "meteor"
				weaponslayout[block+(levelw*2)] = "meteor"
				weaponslayout[block+((levelw*2)+1)] = "meteor"

				levellayout[block] = "."
				leveltileslayout[block] = "."
				levellayout[block+1] = "."
				leveltileslayout[block+1] = "."

			} else {
				levellayout[block] = "_"
				levellayout[block+1] = "_"
			}
		}

		count++
		block--
		if count == updatew {
			count = 0
			block += updatew
			block -= levelw
		}
	}

}
func movewater() { // MARK: movewater

	count := 0
	block := updateblock

	for a := 0; a < updatea; a++ {

		if waterlayout[block] == "water" && waterlayout[block+levelw] == "watertop" {
			waterlayout[block] = "watertop"
			waterlayout[block+levelw] = "water"
		}
		if waterlayout[block] == "watertop" && waterlayout[block+levelw] == "watertop" {
			waterlayout[block] = "watertop"
			waterlayout[block+levelw] = "water"
		}
		if waterlayout[block] == "water" && waterlayout[block-levelw] == "" {
			waterlayout[block] = "watertop"
		}

		count++
		block++
		if count == updatew {
			count = 0
			block -= updatew
			block += levelw
		}

	}

	count = 0
	block = updateblockend
	for a := 0; a < updatea; a++ {

		// water gravity down
		if waterlayout[block] == "water" && waterlayout[block+levelw] == "" && levellayout[block+levelw] == "." {
			waterlayout[block] = ""
			waterlayout[block+levelw] = "water"
		}
		if waterlayout[block] == "watertop" && waterlayout[block+levelw] == "" && levellayout[block+levelw] == "." {
			waterlayout[block] = ""
			waterlayout[block+levelw] = "watertop"
		}

		// water left right both sides empty
		if waterlayout[block] == "water" && waterlayout[block+1] == "" && waterlayout[block-1] == "" && levellayout[block+1] == "." && levellayout[block-1] == "." {
			if flipcoin() {
				waterlayout[block] = ""
				waterlayout[block+1] = "water"
			} else {
				waterlayout[block] = ""
				waterlayout[block-1] = "water"
			}
		}
		if waterlayout[block] == "watertop" && waterlayout[block+1] == "" && waterlayout[block-1] == "" && levellayout[block+1] == "." && levellayout[block-1] == "." {
			if flipcoin() {
				waterlayout[block] = ""
				waterlayout[block+1] = "watertop"
			} else {
				waterlayout[block] = ""
				waterlayout[block-1] = "watertop"
			}
		}

		// water right
		if waterlayout[block] == "watertop" && waterlayout[block+1] == "" && waterlayout[block-1] == "watertop" && levellayout[block+1] == "." {
			waterlayout[block] = ""
			waterlayout[block+1] = "watertop"
		}
		if waterlayout[block] == "water" && waterlayout[block+1] == "" && waterlayout[block-1] == "water" && levellayout[block+1] == "." {
			waterlayout[block] = ""
			waterlayout[block+1] = "water"
		}
		if waterlayout[block] == "water" && waterlayout[block+1] == "" && levellayout[block+1] == "." && levellayout[block-1] == "_" {
			waterlayout[block] = ""
			waterlayout[block+1] = "water"
		}
		if waterlayout[block] == "watertop" && waterlayout[block+1] == "" && levellayout[block+1] == "." && levellayout[block-1] == "_" {
			waterlayout[block] = ""
			waterlayout[block+1] = "watertop"
		}

		// water left
		if waterlayout[block] == "watertop" && waterlayout[block-1] == "" && waterlayout[block+1] == "watertop" && levellayout[block-1] == "." {
			waterlayout[block] = ""
			waterlayout[block-1] = "watertop"
		}
		if waterlayout[block] == "water" && waterlayout[block-1] == "" && waterlayout[block+1] == "water" && levellayout[block-1] == "." {
			waterlayout[block] = ""
			waterlayout[block-1] = "water"
		}
		if waterlayout[block] == "water" && waterlayout[block-1] == "" && levellayout[block-1] == "." && levellayout[block+1] == "_" {
			waterlayout[block] = ""
			waterlayout[block-1] = "water"
		}
		if waterlayout[block] == "watertop" && waterlayout[block-1] == "" && levellayout[block-1] == "." && levellayout[block+1] == "_" {
			waterlayout[block] = ""
			waterlayout[block-1] = "watertop"
		}

		count++
		block--
		if count == updatew {
			count = 0
			block += updatew
			block -= levelw
		}

	}

}
func movefireball() {

	count := 0
	block := drawblocknext

	// move fireball left
	for a := 0; a < drawarea; a++ {
		if weaponslayout[block] == "fireballl" {
			fireballv := block % levelw
			if playerv-fireballv < 30 {
				if enemieslayout[block-1] != "" {
					checkenemy := enemiesstructslayout[block-1]
					checkenemy.hp -= 2
					if checkenemy.hp <= 0 {
						killenemy(block - 1)
						weaponslayout[block] = ""
					} else {
						enemiesstructslayout[block-1] = checkenemy
						weaponslayout[block] = ""
					}
				}
				if levellayout[block-1] == "." {
					weaponslayout[block] = ""
					weaponslayout[block-1] = "fireballl"
				} else {
					weaponslayout[block] = ""
					weaponslayout[block+1] = "deadfireball"
				}
			} else {
				weaponslayout[block] = ""
			}
		}
		count++
		block++
		if count == draww {
			count = 0
			block += levelw
			block -= draww
		}
	}

	count = 0
	block = drawblocknext
	block += levelw * drawh
	block += draww
	// move fireball right
	for a := drawarea; a > 0; a-- {
		if weaponslayout[block] == "fireball" {
			fireballv := block % levelw
			if fireballv-playerv < 30 {
				if enemieslayout[block+1] != "" {
					checkenemy := enemiesstructslayout[block-1]
					checkenemy.hp -= 2
					if checkenemy.hp <= 0 {
						killenemy(block + 1)
						weaponslayout[block] = ""

					} else {
						enemiesstructslayout[block-1] = checkenemy
						weaponslayout[block] = ""
					}
				}
				if levellayout[block+1] == "." {
					weaponslayout[block] = ""
					weaponslayout[block+1] = "fireball"
				} else {
					weaponslayout[block] = ""
					weaponslayout[block+1] = "deadfireball"
				}
			} else {
				weaponslayout[block] = ""
			}
		}
		count++
		block--
		if count == draww {
			count = 0
			block -= levelw
			block += draww
		}
	}

}
func moveinteractables() { // MARK: moveinteractables
	if interactableslayout[player] != "" {
		if interactableslayout[player] == "z" {
			jumpactivespring = true
			jumppause = 2
		}
	}
}
func moveenemies() { // MARK: moveenemies

	if framecount%3 == 0 {

		if !icefreezeactive {

			// left up movements
			count := 0
			block := updateblock
			for a := 0; a < updatea; a++ {

				// pig left up
				if enemieslayout[block] == "pig" {
					if levellayout[block+levelw] != "." {
						holderstruct := enemiesstructslayout[block]
						if holderstruct.movedir == 8 {
							check, nextblock := checknextblock(block, 8)
							if check && moveenemycollisions(block, 8) {
								enemieslayout[block] = ""
								enemieslayout[nextblock] = holderstruct.name
								if rolldice()+rolldice() > 8 {
									holderstruct.movedir = 4
								}
								enemiesstructslayout[block] = enemy{}
								enemiesstructslayout[nextblock] = holderstruct
							} else {
								holderstruct.movedir = 4
								enemiesstructslayout[block] = holderstruct
							}
						}
					}
				}
				// ghost left up
				if enemieslayout[block] == "ghost" {
					holderstruct := enemiesstructslayout[block]
					if holderstruct.movedir == 8 {
						check, nextblock := checknextblock(block, 8)
						if check && moveenemycollisions(block, 8) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							if rolldice() == 6 {
								choose := rInt(1, 9)
								switch choose {
								case 1:
									holderstruct.movedir = 1
								case 2:
									holderstruct.movedir = 2
								case 3:
									holderstruct.movedir = 3
								case 4:
									holderstruct.movedir = 4
								case 5:
									holderstruct.movedir = 5
								case 6:
									holderstruct.movedir = 6
								case 7:
									holderstruct.movedir = 7
								case 8:
									holderstruct.movedir = 8
								}
							}
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							choose := rInt(1, 9)
							switch choose {
							case 1:
								holderstruct.movedir = 1
							case 2:
								holderstruct.movedir = 2
							case 3:
								holderstruct.movedir = 3
							case 4:
								holderstruct.movedir = 4
							case 5:
								holderstruct.movedir = 5
							case 6:
								holderstruct.movedir = 6
							case 7:
								holderstruct.movedir = 7
							case 8:
								holderstruct.movedir = 8
							}
							enemiesstructslayout[block] = holderstruct
						}
					} else if holderstruct.movedir == 1 {
						check, nextblock := checknextblock(block, 1)
						if check && moveenemycollisions(block, 1) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							if rolldice() == 6 {
								choose := rInt(1, 9)
								switch choose {
								case 1:
									holderstruct.movedir = 1
								case 2:
									holderstruct.movedir = 2
								case 3:
									holderstruct.movedir = 3
								case 4:
									holderstruct.movedir = 4
								case 5:
									holderstruct.movedir = 5
								case 6:
									holderstruct.movedir = 6
								case 7:
									holderstruct.movedir = 7
								case 8:
									holderstruct.movedir = 8
								}
							}
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							choose := rInt(1, 9)
							switch choose {
							case 1:
								holderstruct.movedir = 1
							case 2:
								holderstruct.movedir = 2
							case 3:
								holderstruct.movedir = 3
							case 4:
								holderstruct.movedir = 4
							case 5:
								holderstruct.movedir = 5
							case 6:
								holderstruct.movedir = 6
							case 7:
								holderstruct.movedir = 7
							case 8:
								holderstruct.movedir = 8
							}
							enemiesstructslayout[block] = holderstruct
						}
					} else if holderstruct.movedir == 2 {
						check, nextblock := checknextblock(block, 2)
						if check && moveenemycollisions(block, 2) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							if rolldice() == 6 {
								choose := rInt(1, 9)
								switch choose {
								case 1:
									holderstruct.movedir = 1
								case 2:
									holderstruct.movedir = 2
								case 3:
									holderstruct.movedir = 3
								case 4:
									holderstruct.movedir = 4
								case 5:
									holderstruct.movedir = 5
								case 6:
									holderstruct.movedir = 6
								case 7:
									holderstruct.movedir = 7
								case 8:
									holderstruct.movedir = 8
								}
							}
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							choose := rInt(1, 9)
							switch choose {
							case 1:
								holderstruct.movedir = 1
							case 2:
								holderstruct.movedir = 2
							case 3:
								holderstruct.movedir = 3
							case 4:
								holderstruct.movedir = 4
							case 5:
								holderstruct.movedir = 5
							case 6:
								holderstruct.movedir = 6
							case 7:
								holderstruct.movedir = 7
							case 8:
								holderstruct.movedir = 8
							}
							enemiesstructslayout[block] = holderstruct
						}
					} else if holderstruct.movedir == 3 {
						check, nextblock := checknextblock(block, 3)
						if check && moveenemycollisions(block, 3) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							if rolldice() == 6 {
								choose := rInt(1, 9)
								switch choose {
								case 1:
									holderstruct.movedir = 1
								case 2:
									holderstruct.movedir = 2
								case 3:
									holderstruct.movedir = 3
								case 4:
									holderstruct.movedir = 4
								case 5:
									holderstruct.movedir = 5
								case 6:
									holderstruct.movedir = 6
								case 7:
									holderstruct.movedir = 7
								case 8:
									holderstruct.movedir = 8
								}
							}
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							choose := rInt(1, 9)
							switch choose {
							case 1:
								holderstruct.movedir = 1
							case 2:
								holderstruct.movedir = 2
							case 3:
								holderstruct.movedir = 3
							case 4:
								holderstruct.movedir = 4
							case 5:
								holderstruct.movedir = 5
							case 6:
								holderstruct.movedir = 6
							case 7:
								holderstruct.movedir = 7
							case 8:
								holderstruct.movedir = 8
							}
							enemiesstructslayout[block] = holderstruct
						}
					}
				}
				// bat left up
				if enemieslayout[block] == "bat" {
					holderstruct := enemiesstructslayout[block]
					if holderstruct.movedir == 8 {
						check, nextblock := checknextblock(block, 8)
						if check && moveenemycollisions(block, 8) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							if rolldice() == 6 {
								choose := rInt(1, 9)
								switch choose {
								case 1:
									holderstruct.movedir = 1
								case 2:
									holderstruct.movedir = 2
								case 3:
									holderstruct.movedir = 3
								case 4:
									holderstruct.movedir = 4
								case 5:
									holderstruct.movedir = 5
								case 6:
									holderstruct.movedir = 6
								case 7:
									holderstruct.movedir = 7
								case 8:
									holderstruct.movedir = 8
								}
							}
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							choose := rInt(1, 9)
							switch choose {
							case 1:
								holderstruct.movedir = 1
							case 2:
								holderstruct.movedir = 2
							case 3:
								holderstruct.movedir = 3
							case 4:
								holderstruct.movedir = 4
							case 5:
								holderstruct.movedir = 5
							case 6:
								holderstruct.movedir = 6
							case 7:
								holderstruct.movedir = 7
							case 8:
								holderstruct.movedir = 8
							}
							enemiesstructslayout[block] = holderstruct
						}
					} else if holderstruct.movedir == 1 {
						check, nextblock := checknextblock(block, 1)
						if check && moveenemycollisions(block, 1) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							if rolldice() == 6 {
								choose := rInt(1, 9)
								switch choose {
								case 1:
									holderstruct.movedir = 1
								case 2:
									holderstruct.movedir = 2
								case 3:
									holderstruct.movedir = 3
								case 4:
									holderstruct.movedir = 4
								case 5:
									holderstruct.movedir = 5
								case 6:
									holderstruct.movedir = 6
								case 7:
									holderstruct.movedir = 7
								case 8:
									holderstruct.movedir = 8
								}
							}
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							choose := rInt(1, 9)
							switch choose {
							case 1:
								holderstruct.movedir = 1
							case 2:
								holderstruct.movedir = 2
							case 3:
								holderstruct.movedir = 3
							case 4:
								holderstruct.movedir = 4
							case 5:
								holderstruct.movedir = 5
							case 6:
								holderstruct.movedir = 6
							case 7:
								holderstruct.movedir = 7
							case 8:
								holderstruct.movedir = 8
							}
							enemiesstructslayout[block] = holderstruct
						}
					} else if holderstruct.movedir == 2 {
						check, nextblock := checknextblock(block, 2)
						if check && moveenemycollisions(block, 2) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							if rolldice() == 6 {
								choose := rInt(1, 9)
								switch choose {
								case 1:
									holderstruct.movedir = 1
								case 2:
									holderstruct.movedir = 2
								case 3:
									holderstruct.movedir = 3
								case 4:
									holderstruct.movedir = 4
								case 5:
									holderstruct.movedir = 5
								case 6:
									holderstruct.movedir = 6
								case 7:
									holderstruct.movedir = 7
								case 8:
									holderstruct.movedir = 8
								}
							}
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							choose := rInt(1, 9)
							switch choose {
							case 1:
								holderstruct.movedir = 1
							case 2:
								holderstruct.movedir = 2
							case 3:
								holderstruct.movedir = 3
							case 4:
								holderstruct.movedir = 4
							case 5:
								holderstruct.movedir = 5
							case 6:
								holderstruct.movedir = 6
							case 7:
								holderstruct.movedir = 7
							case 8:
								holderstruct.movedir = 8
							}
							enemiesstructslayout[block] = holderstruct
						}
					} else if holderstruct.movedir == 3 {
						check, nextblock := checknextblock(block, 3)
						if check && moveenemycollisions(block, 3) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							if rolldice() == 6 {
								choose := rInt(1, 9)
								switch choose {
								case 1:
									holderstruct.movedir = 1
								case 2:
									holderstruct.movedir = 2
								case 3:
									holderstruct.movedir = 3
								case 4:
									holderstruct.movedir = 4
								case 5:
									holderstruct.movedir = 5
								case 6:
									holderstruct.movedir = 6
								case 7:
									holderstruct.movedir = 7
								case 8:
									holderstruct.movedir = 8
								}
							}
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							choose := rInt(1, 9)
							switch choose {
							case 1:
								holderstruct.movedir = 1
							case 2:
								holderstruct.movedir = 2
							case 3:
								holderstruct.movedir = 3
							case 4:
								holderstruct.movedir = 4
							case 5:
								holderstruct.movedir = 5
							case 6:
								holderstruct.movedir = 6
							case 7:
								holderstruct.movedir = 7
							case 8:
								holderstruct.movedir = 8
							}
							enemiesstructslayout[block] = holderstruct
						}
					}
				}

				// rock left
				if enemieslayout[block] == "rock" {
					holderstruct := enemiesstructslayout[block]
					if holderstruct.movedir == 8 {
						check, nextblock := checknextblock(block, 8)
						if check && moveenemycollisions(block, 8) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							if rolldice()+rolldice() > 10 {
								holderstruct.movedir = 4
							}
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							holderstruct.movedir = 4
							enemiesstructslayout[nextblock] = holderstruct
						}
					}
				}
				// snail left
				if enemieslayout[block] == "snail" {
					if framecount%4 == 0 {
						holderstruct := enemiesstructslayout[block]
						if holderstruct.movedir == 8 {
							check, nextblock := checknextblock(block, 8)
							if check && moveenemycollisions(block, 8) {
								enemieslayout[block] = ""
								enemieslayout[nextblock] = holderstruct.name
								enemiesstructslayout[block] = enemy{}
								enemiesstructslayout[nextblock] = holderstruct

							} else {
								holderstruct.movedir = 4
								enemiesstructslayout[nextblock] = holderstruct
							}
						}
					}
				}

				// radish left
				if enemieslayout[block] == "radish" {
					holderstruct := enemiesstructslayout[block]
					if holderstruct.movedir == 8 {
						check, nextblock := checknextblock(block, 8)
						if check && moveenemycollisions(block, 8) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							if rolldice()+rolldice() > 10 {
								if flipcoin() {
									holderstruct.movedir = 2
								} else {
									holderstruct.movedir = 6
								}
							}
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							holderstruct.movedir = 4
							enemiesstructslayout[nextblock] = holderstruct
						}
					}
				}

				// radish up
				if enemieslayout[block] == "radish" {
					holderstruct := enemiesstructslayout[block]
					if holderstruct.movedir == 2 {
						check, nextblock := checknextblock(block, 2)
						if check && moveenemycollisions(block, 2) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							if flipcoin() {
								holderstruct.movedir = 4
							} else {
								holderstruct.movedir = 8
							}
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							holderstruct.movedir = 6
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						}
					}
				}

				// spikes left
				if enemieslayout[block] == "spike" {
					holderstruct := enemiesstructslayout[block]
					if holderstruct.movedir == 8 && holderstruct.fall == false {
						check, nextblock := checknextblock(block, 8)
						if check && moveenemycollisions(block, 8) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							if holderstruct.updown {
								if levellayout[nextblock-levelw] == "." {
									holderstruct.fall = true
									holderstruct.updown = false
								}
							} else {
								if levellayout[nextblock+levelw] == "." {
									holderstruct.fall = true
									holderstruct.updown = false
								}
							}
							if rolldice()+rolldice() == 12 {
								holderstruct.movedir = 4
							}
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							if levellayout[block-levelw] == "." {
								holderstruct.movedir = 2
								holderstruct.updown = true
								enemiesstructslayout[nextblock] = holderstruct
							} else if levellayout[block+levelw] == "." {
								holderstruct.movedir = 6
								holderstruct.updown = true
								enemiesstructslayout[nextblock] = holderstruct
							}
						}
					}
				}

				// spike up
				if enemieslayout[block] == "spike" {
					holderstruct := enemiesstructslayout[block]
					if holderstruct.movedir == 2 {
						if levellayout[block+1] == "." && levellayout[block-1] == "." {
							if levellayout[(block+1)+levelw] != "." {
								holderstruct.updown = false
								holderstruct.movedir = 4
								check, nextblock := checknextblock(block, 4)
								if check && moveenemycollisions(block, 4) {
									enemieslayout[block] = ""
									enemieslayout[nextblock] = holderstruct.name
									enemiesstructslayout[block] = enemy{}
									enemiesstructslayout[nextblock] = holderstruct
								}
							} else if levellayout[(block-1)+levelw] != "." {
								holderstruct.updown = false
								holderstruct.movedir = 8
								check, nextblock := checknextblock(block, 8)
								if check && moveenemycollisions(block, 8) {
									enemieslayout[block] = ""
									enemieslayout[nextblock] = holderstruct.name
									enemiesstructslayout[block] = enemy{}
									enemiesstructslayout[nextblock] = holderstruct
								}
							} else {
								holderstruct.fall = true
								holderstruct.updown = false
								if flipcoin() {
									holderstruct.movedir = 4
								} else {
									holderstruct.movedir = 8
								}
							}
						} else {
							check, nextblock := checknextblock(block, 2)
							if check && moveenemycollisions(block, 2) {
								enemieslayout[block] = ""
								enemieslayout[nextblock] = holderstruct.name
								enemiesstructslayout[block] = enemy{}
								enemiesstructslayout[nextblock] = holderstruct
							} else {
								if levellayout[block+1] == "." {
									holderstruct.movedir = 4
									enemiesstructslayout[nextblock] = holderstruct
								} else if levellayout[block-1] == "." {
									holderstruct.movedir = 8
									enemiesstructslayout[nextblock] = holderstruct
								}
							}
						}
					}
				}
				// plant left
				if enemieslayout[block] == "plant" {
					holderstruct := enemiesstructslayout[block]
					if holderstruct.movedir == 8 {
						check, nextblock := checknextblock(block, 8)
						if check && moveenemycollisions(block, 8) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							holderstruct.movedir = 4
							enemiesstructslayout[nextblock] = holderstruct
						}
					}

					if rolldice() == 6 {
						enemybulletslayout[block-4] = "bulletl"
					}
				}

				// slime up
				if enemieslayout[block] == "slime" {
					holderstruct := enemiesstructslayout[block]

					if framecount%(rInt(75, 120)) == 0 {
						if holderstruct.updown == true {
							holderstruct.updown = false
							enemiesstructslayout[block] = holderstruct
						} else if holderstruct.updown == false {
							holderstruct.updown = true
							enemiesstructslayout[block] = holderstruct
						}
					}

					if rolldice()+rolldice() == 12 {
						if holderstruct.movedir == 4 {
							holderstruct.movedir = 8
							enemiesstructslayout[block] = holderstruct
						} else if holderstruct.movedir == 8 {
							holderstruct.movedir = 4
							enemiesstructslayout[block] = holderstruct
						}
					}

					if holderstruct.updown == true && levellayout[block-levelw] == "." {
						if moveenemycollisions(block, 2) {
							holderstruct.jump = true
							enemieslayout[block] = ""
							enemieslayout[block-levelw] = "slime"
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[block-levelw] = holderstruct
						}
					}
				}

				// slime left
				if enemieslayout[block] == "slime" {
					holderstruct := enemiesstructslayout[block]
					if holderstruct.movedir == 8 {
						check, nextblock := checknextblock(block, 8)
						if check {
							if moveenemycollisions(block, 8) {
								enemieslayout[block] = ""
								enemieslayout[nextblock] = "slime"
								enemiesstructslayout[block] = enemy{}
								enemiesstructslayout[nextblock] = holderstruct
							}
						} else {
							holderstruct.movedir = 4
							enemiesstructslayout[block] = holderstruct
						}
					}
				}

				// rabbit left up
				if enemieslayout[block] == "rabbit" {
					holderstruct := enemiesstructslayout[block]
					if levellayout[block+levelw] != "." {
						holderstruct.fall = false
					}
					if holderstruct.jump == false && holderstruct.fall == false {
						if holderstruct.movedir == 8 {
							check, nextblock := checknextblock(block, 8)
							if check && moveenemycollisions(block, 8) {
								enemieslayout[block] = ""
								enemieslayout[nextblock] = "rabbit"
								if rolldice() == 6 {
									holderstruct.movedir = 4
								}
								enemiesstructslayout[block] = enemy{}
								enemiesstructslayout[nextblock] = holderstruct
							} else {
								holderstruct.movedir = 4
								enemiesstructslayout[block] = enemy{}
								enemiesstructslayout[nextblock] = holderstruct
							}
						}
					}

					if holderstruct.jump == true {
						if holderstruct.movedir == 3 {
							check, nextblock := checknextblock(block, 3)
							if check {
								enemieslayout[block] = ""
								enemieslayout[nextblock] = "rabbit"
								holderstruct.jumpcount++
								enemiesstructslayout[block] = enemy{}
								enemiesstructslayout[nextblock] = holderstruct
							} else {
								holderstruct.jump = false
								holderstruct.jumpcount = 0
								holderstruct.movedir = 4
								enemiesstructslayout[nextblock] = holderstruct
							}
							if holderstruct.jumpcount >= 4 {
								holderstruct.jump = false
								holderstruct.jumpcount = 0
								holderstruct.movedir = 4
								enemiesstructslayout[nextblock] = holderstruct
							}
						}
						if holderstruct.movedir == 1 {
							check, nextblock := checknextblock(block, 1)
							if check {
								enemieslayout[block] = ""
								enemieslayout[nextblock] = "rabbit"
								holderstruct.jumpcount++
								enemiesstructslayout[block] = enemy{}
								enemiesstructslayout[nextblock] = holderstruct
							} else {
								holderstruct.jump = false
								holderstruct.jumpcount = 0
								holderstruct.movedir = 8
								enemiesstructslayout[nextblock] = holderstruct
							}
							if holderstruct.jumpcount >= 4 {
								holderstruct.jump = false
								holderstruct.jumpcount = 0
								holderstruct.movedir = 8
								enemiesstructslayout[nextblock] = holderstruct
							}
						}

					}

				} // end rabbit

				if enemieslayout[block] == "bee" {
					holderstruct := enemiesstructslayout[block]

					switch holderstruct.movedir {
					case 1:
						check, nextblock := checknextblock(block, 1)
						if check && moveenemycollisions(block, 1) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = "bee"
							enemiesstructslayout[block] = enemy{}
							if rolldice() == 6 {
								holderstruct.movedir = 6
							} else {
								choose := rInt(1, 6)
								switch choose {
								case 1:
									holderstruct.movedir = 1
								case 2:
									holderstruct.movedir = 2
								case 3:
									holderstruct.movedir = 3
								case 4:
									holderstruct.movedir = 4
								case 5:
									holderstruct.movedir = 8
								}
							}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							if rolldice() == 6 {
								holderstruct.movedir = 6
							} else {
								choose := rInt(1, 6)
								switch choose {
								case 1:
									holderstruct.movedir = 1
								case 2:
									holderstruct.movedir = 2
								case 3:
									holderstruct.movedir = 3
								case 4:
									holderstruct.movedir = 4
								case 5:
									holderstruct.movedir = 8
								}
							}
							enemiesstructslayout[nextblock] = holderstruct
						}
					case 2:
						check, nextblock := checknextblock(block, 2)
						if check && moveenemycollisions(block, 2) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = "bee"
							enemiesstructslayout[block] = enemy{}
							if rolldice() == 6 {
								holderstruct.movedir = 6
							} else {
								choose := rInt(1, 6)
								switch choose {
								case 1:
									holderstruct.movedir = 1
								case 2:
									holderstruct.movedir = 2
								case 3:
									holderstruct.movedir = 3
								case 4:
									holderstruct.movedir = 4
								case 5:
									holderstruct.movedir = 8
								}
							}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							if rolldice() == 6 {
								holderstruct.movedir = 6
							} else {
								choose := rInt(1, 6)
								switch choose {
								case 1:
									holderstruct.movedir = 1
								case 2:
									holderstruct.movedir = 2
								case 3:
									holderstruct.movedir = 3
								case 4:
									holderstruct.movedir = 4
								case 5:
									holderstruct.movedir = 8
								}
							}
							enemiesstructslayout[nextblock] = holderstruct
						}
					case 3:
						check, nextblock := checknextblock(block, 3)
						if check && moveenemycollisions(block, 3) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = "bee"
							enemiesstructslayout[block] = enemy{}
							if rolldice() == 6 {
								holderstruct.movedir = 6
							} else {
								choose := rInt(1, 6)
								switch choose {
								case 1:
									holderstruct.movedir = 1
								case 2:
									holderstruct.movedir = 2
								case 3:
									holderstruct.movedir = 3
								case 4:
									holderstruct.movedir = 4
								case 5:
									holderstruct.movedir = 8
								}
							}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							if rolldice() == 6 {
								holderstruct.movedir = 6
							} else {
								choose := rInt(1, 6)
								switch choose {
								case 1:
									holderstruct.movedir = 1
								case 2:
									holderstruct.movedir = 2
								case 3:
									holderstruct.movedir = 3
								case 4:
									holderstruct.movedir = 4
								case 5:
									holderstruct.movedir = 8
								}
							}
							enemiesstructslayout[nextblock] = holderstruct
						}
					case 8:
						check, nextblock := checknextblock(block, 8)
						if check && moveenemycollisions(block, 8) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = "bee"
							enemiesstructslayout[block] = enemy{}
							if rolldice() == 6 {
								holderstruct.movedir = 6
							} else {
								choose := rInt(1, 6)
								switch choose {
								case 1:
									holderstruct.movedir = 1
								case 2:
									holderstruct.movedir = 2
								case 3:
									holderstruct.movedir = 3
								case 4:
									holderstruct.movedir = 4
								case 5:
									holderstruct.movedir = 8
								}
							}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							if rolldice() == 6 {
								holderstruct.movedir = 6
							} else {
								choose := rInt(1, 6)
								switch choose {
								case 1:
									holderstruct.movedir = 1
								case 2:
									holderstruct.movedir = 2
								case 3:
									holderstruct.movedir = 3
								case 4:
									holderstruct.movedir = 4
								case 5:
									holderstruct.movedir = 8
								}
							}
							enemiesstructslayout[nextblock] = holderstruct
						}
					}
				} // end bee

				if enemieslayout[block] == "bird" {

					if moveenemycollisions(block, 8) {
						holderstruct := enemiesstructslayout[block]
						if holderstruct.movedir == 4 {
							if levellayout[block-1] == "." {
								enemieslayout[block] = ""
								enemieslayout[block-1] = "bird"
								enemiesstructslayout[block-1] = holderstruct
							} else if levellayout[block-1] == "_" || levellayout[block-1] == "#" || levellayout[block-1] == "%" {
								holderstruct.movedir = 2
								enemiesstructslayout[block] = holderstruct
							}
						}
					} else {
						holderstruct := enemiesstructslayout[block]
						holderstruct.movedir = 2
						enemiesstructslayout[block] = holderstruct
					}

				} // end bird

				if enemieslayout[block] == "skull" {
					choose := rolldice()
					choose2 := rolldice() + rolldice()
					if flipcoin() {
						if choose > 3 {
							if moveenemycollisions(block, 2) {
								if levellayout[block-levelw] == "." {
									enemieslayout[block] = ""
									holderstruct := enemiesstructslayout[block]
									enemiesstructslayout[block] = enemy{}
									enemieslayout[block-levelw] = "skull"
									enemiesstructslayout[block-levelw] = holderstruct
								}
							}
						}
					} else {
						if choose2 > 10 {
							if flipcoin() {
								if moveenemycollisions(block, 8) {
									if levellayout[block-1] == "." {
										enemieslayout[block] = ""
										holderstruct := enemiesstructslayout[block]
										enemiesstructslayout[block] = enemy{}
										enemieslayout[block-1] = "skull"
										enemiesstructslayout[block-1] = holderstruct
									}
								}
							}
						}
					}
				} // end skull

				if enemieslayout[block] == "mushroom" {
					checkenemy := enemiesstructslayout[block]
					switch checkenemy.movedir {
					case 2:
						if levellayout[block-levelw] == "." {
							if levellayout[block+levelw] == "_" || levellayout[block+levelw] == "#" {
								if enemieslayout[block-levelw] == "" {
									holder := enemieslayout[block]
									enemieslayout[block] = ""
									holderstruct := enemiesstructslayout[block]
									enemiesstructslayout[block] = enemy{}
									enemiesstructslayout[block-levelw] = holderstruct
									enemieslayout[block-levelw] = holder
								} else {
									movenemycollisionsmushrooms(block-levelw, block, 2)
								}
							} else if levellayout[block+levelw] == "." && levellayout[block+(levelw*2)] == "_" {
								holder := enemieslayout[block]
								enemieslayout[block] = ""
								holderstruct := enemiesstructslayout[block]
								enemiesstructslayout[block] = enemy{}
								enemiesstructslayout[block-levelw] = holderstruct
								enemieslayout[block-levelw] = holder
							} else {
								if flipcoin() {
									checkenemy.movedir = 4
									enemiesstructslayout[block] = checkenemy
								} else {
									checkenemy.movedir = 8
									enemiesstructslayout[block] = checkenemy
								}
							}
						} else {
							if flipcoin() {
								checkenemy.movedir = 4
								enemiesstructslayout[block] = checkenemy
							} else {
								checkenemy.movedir = 8
								enemiesstructslayout[block] = checkenemy
							}
						}
					case 8:
						if rolldice() == 6 {
							checkenemy.movedir = 4
							enemiesstructslayout[block] = checkenemy
						}
						if rolldice()+rolldice() > 10 {
							checkenemy.movedir = 2
							enemiesstructslayout[block] = checkenemy
						}
						if levellayout[block-1] == "." {
							if enemieslayout[block-1] == "" {
								enemyvert := block % levelw
								if enemyvert > 30 {
									holder := enemieslayout[block]
									holderstruct := enemiesstructslayout[block]
									enemieslayout[block] = ""
									enemiesstructslayout[block] = enemy{}
									enemiesstructslayout[block-1] = holderstruct
									enemieslayout[block-1] = holder
								} else {
									checkenemy.movedir = 4
									enemiesstructslayout[block] = checkenemy
								}
							} else {
								movenemycollisionsmushrooms(block, block-1, 8)
							}
						} else {
							checkenemy.movedir = 4
							enemiesstructslayout[block] = checkenemy
						}
					}
				} // end mushroom
				count++
				block++
				if count == updatew {
					count = 0
					block -= updatew
					block += levelw
				}
			}
			// right down movements
			count = 0
			block = updateblockend
			for a := 0; a < updatea; a++ {

				//  pig right down
				if enemieslayout[block] == "pig" {
					if levellayout[block+levelw] == "." {
						holderstruct := enemiesstructslayout[block]
						check, nextblock := checknextblock(block, 6)
						if check && moveenemycollisions(block, 6) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						}
					} else {
						holderstruct := enemiesstructslayout[block]
						if holderstruct.movedir == 4 {
							check, nextblock := checknextblock(block, 4)
							if check && moveenemycollisions(block, 4) {
								enemieslayout[block] = ""
								enemieslayout[nextblock] = holderstruct.name
								if rolldice()+rolldice() > 8 {
									holderstruct.movedir = 8
								}
								enemiesstructslayout[block] = enemy{}
								enemiesstructslayout[nextblock] = holderstruct
							} else {
								holderstruct.movedir = 8
								enemiesstructslayout[block] = holderstruct
							}
						}
					}
				}

				// ghost right down
				if enemieslayout[block] == "ghost" {
					holderstruct := enemiesstructslayout[block]
					if holderstruct.movedir == 4 {
						check, nextblock := checknextblock(block, 4)
						if check && moveenemycollisions(block, 4) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							if rolldice() == 6 {
								choose := rInt(1, 9)
								switch choose {
								case 1:
									holderstruct.movedir = 1
								case 2:
									holderstruct.movedir = 2
								case 3:
									holderstruct.movedir = 3
								case 4:
									holderstruct.movedir = 4
								case 5:
									holderstruct.movedir = 5
								case 6:
									holderstruct.movedir = 6
								case 7:
									holderstruct.movedir = 7
								case 8:
									holderstruct.movedir = 8
								}
							}
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							choose := rInt(1, 9)
							switch choose {
							case 1:
								holderstruct.movedir = 1
							case 2:
								holderstruct.movedir = 2
							case 3:
								holderstruct.movedir = 3
							case 4:
								holderstruct.movedir = 4
							case 5:
								holderstruct.movedir = 5
							case 6:
								holderstruct.movedir = 6
							case 7:
								holderstruct.movedir = 7
							case 8:
								holderstruct.movedir = 8
							}
							enemiesstructslayout[block] = holderstruct
						}
					} else if holderstruct.movedir == 5 {
						check, nextblock := checknextblock(block, 5)
						if check && moveenemycollisions(block, 5) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							if rolldice() == 6 {
								choose := rInt(1, 9)
								switch choose {
								case 1:
									holderstruct.movedir = 1
								case 2:
									holderstruct.movedir = 2
								case 3:
									holderstruct.movedir = 3
								case 4:
									holderstruct.movedir = 4
								case 5:
									holderstruct.movedir = 5
								case 6:
									holderstruct.movedir = 6
								case 7:
									holderstruct.movedir = 7
								case 8:
									holderstruct.movedir = 8
								}
							}
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							choose := rInt(1, 9)
							switch choose {
							case 1:
								holderstruct.movedir = 1
							case 2:
								holderstruct.movedir = 2
							case 3:
								holderstruct.movedir = 3
							case 4:
								holderstruct.movedir = 4
							case 5:
								holderstruct.movedir = 5
							case 6:
								holderstruct.movedir = 6
							case 7:
								holderstruct.movedir = 7
							case 8:
								holderstruct.movedir = 8
							}
							enemiesstructslayout[block] = holderstruct
						}
					} else if holderstruct.movedir == 6 {
						check, nextblock := checknextblock(block, 6)
						if check && moveenemycollisions(block, 6) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							if rolldice() == 6 {
								choose := rInt(1, 9)
								switch choose {
								case 1:
									holderstruct.movedir = 1
								case 2:
									holderstruct.movedir = 2
								case 3:
									holderstruct.movedir = 3
								case 4:
									holderstruct.movedir = 4
								case 5:
									holderstruct.movedir = 5
								case 6:
									holderstruct.movedir = 6
								case 7:
									holderstruct.movedir = 7
								case 8:
									holderstruct.movedir = 8
								}
							}
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							choose := rInt(1, 9)
							switch choose {
							case 1:
								holderstruct.movedir = 1
							case 2:
								holderstruct.movedir = 2
							case 3:
								holderstruct.movedir = 3
							case 4:
								holderstruct.movedir = 4
							case 5:
								holderstruct.movedir = 5
							case 6:
								holderstruct.movedir = 6
							case 7:
								holderstruct.movedir = 7
							case 8:
								holderstruct.movedir = 8
							}
							enemiesstructslayout[block] = holderstruct
						}
					} else if holderstruct.movedir == 7 {
						check, nextblock := checknextblock(block, 7)
						if check && moveenemycollisions(block, 7) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							if rolldice() == 6 {
								choose := rInt(1, 9)
								switch choose {
								case 1:
									holderstruct.movedir = 1
								case 2:
									holderstruct.movedir = 2
								case 3:
									holderstruct.movedir = 3
								case 4:
									holderstruct.movedir = 4
								case 5:
									holderstruct.movedir = 5
								case 6:
									holderstruct.movedir = 6
								case 7:
									holderstruct.movedir = 7
								case 8:
									holderstruct.movedir = 8
								}
							}
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							choose := rInt(1, 9)
							switch choose {
							case 1:
								holderstruct.movedir = 1
							case 2:
								holderstruct.movedir = 2
							case 3:
								holderstruct.movedir = 3
							case 4:
								holderstruct.movedir = 4
							case 5:
								holderstruct.movedir = 5
							case 6:
								holderstruct.movedir = 6
							case 7:
								holderstruct.movedir = 7
							case 8:
								holderstruct.movedir = 8
							}
							enemiesstructslayout[block] = holderstruct
						}
					}
				}
				// bat right down
				if enemieslayout[block] == "bat" {
					holderstruct := enemiesstructslayout[block]
					if holderstruct.movedir == 4 {
						check, nextblock := checknextblock(block, 4)
						if check && moveenemycollisions(block, 4) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							if rolldice() == 6 {
								choose := rInt(1, 9)
								switch choose {
								case 1:
									holderstruct.movedir = 1
								case 2:
									holderstruct.movedir = 2
								case 3:
									holderstruct.movedir = 3
								case 4:
									holderstruct.movedir = 4
								case 5:
									holderstruct.movedir = 5
								case 6:
									holderstruct.movedir = 6
								case 7:
									holderstruct.movedir = 7
								case 8:
									holderstruct.movedir = 8
								}
							}
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							choose := rInt(1, 9)
							switch choose {
							case 1:
								holderstruct.movedir = 1
							case 2:
								holderstruct.movedir = 2
							case 3:
								holderstruct.movedir = 3
							case 4:
								holderstruct.movedir = 4
							case 5:
								holderstruct.movedir = 5
							case 6:
								holderstruct.movedir = 6
							case 7:
								holderstruct.movedir = 7
							case 8:
								holderstruct.movedir = 8
							}
							enemiesstructslayout[block] = holderstruct
						}
					} else if holderstruct.movedir == 5 {
						check, nextblock := checknextblock(block, 5)
						if check && moveenemycollisions(block, 5) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							if rolldice() == 6 {
								choose := rInt(1, 9)
								switch choose {
								case 1:
									holderstruct.movedir = 1
								case 2:
									holderstruct.movedir = 2
								case 3:
									holderstruct.movedir = 3
								case 4:
									holderstruct.movedir = 4
								case 5:
									holderstruct.movedir = 5
								case 6:
									holderstruct.movedir = 6
								case 7:
									holderstruct.movedir = 7
								case 8:
									holderstruct.movedir = 8
								}
							}
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							choose := rInt(1, 9)
							switch choose {
							case 1:
								holderstruct.movedir = 1
							case 2:
								holderstruct.movedir = 2
							case 3:
								holderstruct.movedir = 3
							case 4:
								holderstruct.movedir = 4
							case 5:
								holderstruct.movedir = 5
							case 6:
								holderstruct.movedir = 6
							case 7:
								holderstruct.movedir = 7
							case 8:
								holderstruct.movedir = 8
							}
							enemiesstructslayout[block] = holderstruct
						}
					} else if holderstruct.movedir == 6 {
						check, nextblock := checknextblock(block, 6)
						if check && moveenemycollisions(block, 6) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							if rolldice() == 6 {
								choose := rInt(1, 9)
								switch choose {
								case 1:
									holderstruct.movedir = 1
								case 2:
									holderstruct.movedir = 2
								case 3:
									holderstruct.movedir = 3
								case 4:
									holderstruct.movedir = 4
								case 5:
									holderstruct.movedir = 5
								case 6:
									holderstruct.movedir = 6
								case 7:
									holderstruct.movedir = 7
								case 8:
									holderstruct.movedir = 8
								}
							}
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							choose := rInt(1, 9)
							switch choose {
							case 1:
								holderstruct.movedir = 1
							case 2:
								holderstruct.movedir = 2
							case 3:
								holderstruct.movedir = 3
							case 4:
								holderstruct.movedir = 4
							case 5:
								holderstruct.movedir = 5
							case 6:
								holderstruct.movedir = 6
							case 7:
								holderstruct.movedir = 7
							case 8:
								holderstruct.movedir = 8
							}
							enemiesstructslayout[block] = holderstruct
						}
					} else if holderstruct.movedir == 7 {
						check, nextblock := checknextblock(block, 7)
						if check && moveenemycollisions(block, 7) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							if rolldice() == 6 {
								choose := rInt(1, 9)
								switch choose {
								case 1:
									holderstruct.movedir = 1
								case 2:
									holderstruct.movedir = 2
								case 3:
									holderstruct.movedir = 3
								case 4:
									holderstruct.movedir = 4
								case 5:
									holderstruct.movedir = 5
								case 6:
									holderstruct.movedir = 6
								case 7:
									holderstruct.movedir = 7
								case 8:
									holderstruct.movedir = 8
								}
							}
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							choose := rInt(1, 9)
							switch choose {
							case 1:
								holderstruct.movedir = 1
							case 2:
								holderstruct.movedir = 2
							case 3:
								holderstruct.movedir = 3
							case 4:
								holderstruct.movedir = 4
							case 5:
								holderstruct.movedir = 5
							case 6:
								holderstruct.movedir = 6
							case 7:
								holderstruct.movedir = 7
							case 8:
								holderstruct.movedir = 8
							}
							enemiesstructslayout[block] = holderstruct
						}
					}
				}

				// rock right
				if enemieslayout[block] == "rock" {
					holderstruct := enemiesstructslayout[block]

					if levellayout[block+levelw] == "." {
						check, nextblock := checknextblock(block, 6)
						if check && moveenemycollisions(block, 6) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						}
					} else {
						if holderstruct.movedir == 4 {
							check, nextblock := checknextblock(block, 4)
							if check && moveenemycollisions(block, 4) {
								enemieslayout[block] = ""
								enemieslayout[nextblock] = holderstruct.name
								if rolldice()+rolldice() > 10 {
									holderstruct.movedir = 8
								}
								enemiesstructslayout[block] = enemy{}
								enemiesstructslayout[nextblock] = holderstruct
							} else {
								holderstruct.movedir = 8
								enemiesstructslayout[nextblock] = holderstruct
							}
						}
					}
				}

				if enemieslayout[block] == "rock" {
					holderstruct := enemiesstructslayout[block]
					if rolldice()+rolldice() == 12 {
						if flipcoin() {
							holderstruct.movedir = 4
							enemiesstructslayout[block] = holderstruct
							if flipcoin() {
								bombon = false
								bombdown = false
								explodeblock = block + 4
								explodeblock -= levelw * rInt(12, 15)
								bombexplode()
							} else {
								bombon = false
								bombdown = false
								explodeblock = block + rInt(12, 20)
								explodeblock -= levelw * 5
								bombexplode()
							}
						} else {
							holderstruct.movedir = 8
							enemiesstructslayout[block] = holderstruct
							if flipcoin() {
								bombon = false
								bombdown = false
								explodeblock = block - 4
								explodeblock -= levelw * rInt(12, 15)
								bombexplode()
							} else {
								bombon = false
								bombdown = false
								explodeblock = block - rInt(12, 20)
								explodeblock -= levelw * 5
								bombexplode()
							}
						}

					}
				}

				// snail right
				if enemieslayout[block] == "snail" {

					if levellayout[block+levelw] == "." {
						holderstruct := enemiesstructslayout[block]
						check, nextblock := checknextblock(block, 6)
						if check && moveenemycollisions(block, 6) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						}
					} else {
						if framecount%4 == 0 {
							holderstruct := enemiesstructslayout[block]
							if holderstruct.movedir == 4 {
								check, nextblock := checknextblock(block, 4)
								if check && moveenemycollisions(block, 4) {
									enemieslayout[block] = ""
									enemieslayout[nextblock] = holderstruct.name
									enemiesstructslayout[block] = enemy{}
									enemiesstructslayout[nextblock] = holderstruct
								} else {
									holderstruct.movedir = 8
									enemiesstructslayout[nextblock] = holderstruct
								}
							}
						}
					}
				}

				// radish right
				if enemieslayout[block] == "radish" {
					holderstruct := enemiesstructslayout[block]
					if holderstruct.movedir == 4 {
						check, nextblock := checknextblock(block, 4)
						if check && moveenemycollisions(block, 4) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							if rolldice()+rolldice() > 10 {
								if flipcoin() {
									holderstruct.movedir = 2
								} else {
									holderstruct.movedir = 6
								}
							}
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							holderstruct.movedir = 8
							enemiesstructslayout[nextblock] = holderstruct
						}
					}
				}

				// radish down
				if enemieslayout[block] == "radish" {
					holderstruct := enemiesstructslayout[block]
					if holderstruct.movedir == 6 {
						check, nextblock := checknextblock(block, 6)
						if check && moveenemycollisions(block, 6) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							if flipcoin() {
								holderstruct.movedir = 4
							} else {
								holderstruct.movedir = 8
							}
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							holderstruct.movedir = 2
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						}
					}
				}

				// spikes right
				if enemieslayout[block] == "spike" {
					holderstruct := enemiesstructslayout[block]
					if holderstruct.movedir == 4 && holderstruct.fall == false {
						check, nextblock := checknextblock(block, 4)
						if check && moveenemycollisions(block, 4) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							if holderstruct.updown {
								if levellayout[nextblock-levelw] == "." {
									holderstruct.fall = true
									holderstruct.updown = false
								}
							} else {
								if levellayout[nextblock+levelw] == "." {
									holderstruct.fall = true
									holderstruct.updown = false
								}
							}
							if rolldice()+rolldice() == 12 {
								holderstruct.movedir = 8
							}
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							if levellayout[block-levelw] == "." {
								holderstruct.movedir = 2
								holderstruct.updown = true
								enemiesstructslayout[nextblock] = holderstruct
							} else if levellayout[block+levelw] == "." {
								holderstruct.movedir = 6
								holderstruct.updown = true
								enemiesstructslayout[nextblock] = holderstruct
							}
						}
					}
				}

				// spikes gravity
				if enemieslayout[block] == "spike" {
					holderstruct := enemiesstructslayout[block]
					if holderstruct.fall {
						check, nextblock := checknextblock(block, 6)
						if check && moveenemycollisions(block, 6) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							if levellayout[nextblock+levelw] != "." {
								holderstruct.fall = false
								holderstruct.updown = false
								if flipcoin() {
									holderstruct.movedir = 4
								} else {
									holderstruct.movedir = 8
								}
							}
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						}
					}
				}
				// spikes down
				if enemieslayout[block] == "spike" {
					holderstruct := enemiesstructslayout[block]
					if holderstruct.movedir == 6 {
						check, nextblock := checknextblock(block, 6)
						if check && moveenemycollisions(block, 6) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							if levellayout[block+1] == "." {
								holderstruct.movedir = 4
								holderstruct.updown = false
								enemiesstructslayout[nextblock] = holderstruct
							} else if levellayout[block-1] == "." {
								holderstruct.movedir = 8
								holderstruct.updown = false
								enemiesstructslayout[nextblock] = holderstruct
							}
						}
					}
				}

				// plant down
				if enemieslayout[block] == "plant" {
					holderstruct := enemiesstructslayout[block]

					if levellayout[block+levelw] == "." {
						if moveenemycollisions(block, 6) {
							enemieslayout[block] = ""
							enemiesstructslayout[block] = enemy{}
							enemieslayout[block+levelw] = holderstruct.name
							enemiesstructslayout[block+levelw] = holderstruct
						}

					}

				}

				// plant right
				if enemieslayout[block] == "plant" {
					holderstruct := enemiesstructslayout[block]
					if holderstruct.movedir == 4 {
						check, nextblock := checknextblock(block, 4)
						if check && moveenemycollisions(block, 4) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = holderstruct.name
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							holderstruct.movedir = 8
							enemiesstructslayout[nextblock] = holderstruct
						}
					}

					if rolldice() == 6 {
						enemybulletslayout[block+2] = "bulletr"
					}
				}

				// slime down
				if enemieslayout[block] == "slime" {
					holderstruct := enemiesstructslayout[block]
					if moveenemycollisions(block, 6) {
						if holderstruct.updown == false && levellayout[block+levelw] == "." {
							holderstruct.jump = true
							enemieslayout[block] = ""
							enemieslayout[block+levelw] = "slime"
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[block+levelw] = holderstruct
						} else if holderstruct.updown == false && levellayout[block+levelw] != "." {
							holderstruct.jump = false
						}
					}
				}
				// slime right
				if enemieslayout[block] == "slime" {
					holderstruct := enemiesstructslayout[block]
					if holderstruct.movedir == 4 {
						check, nextblock := checknextblock(block, 4)
						if check {
							if moveenemycollisions(block, 4) {
								enemieslayout[block] = ""
								enemieslayout[nextblock] = "slime"
								enemiesstructslayout[block] = enemy{}
								enemiesstructslayout[nextblock] = holderstruct
							}
						} else {
							holderstruct.movedir = 8
							enemiesstructslayout[block] = holderstruct
						}
					}
				}

				// rabbit right down
				if enemieslayout[block] == "rabbit" {
					holderstruct := enemiesstructslayout[block]
					if holderstruct.jump == false {
						if levellayout[block+levelw] == "." {
							holderstruct.fall = true
							if moveenemycollisions(block, 6) {
								enemieslayout[block] = ""
								enemieslayout[block+levelw] = "rabbit"
								enemiesstructslayout[block] = enemy{}
								enemiesstructslayout[block+levelw] = holderstruct
							}
						} else if levellayout[block+levelw] != "." {
							holderstruct.fall = false
						}
						if holderstruct.fall == false {
							if holderstruct.movedir == 4 {
								check, nextblock := checknextblock(block, 4)
								if check && moveenemycollisions(block, 4) {
									enemieslayout[block] = ""
									enemieslayout[nextblock] = "rabbit"
									if rolldice() == 6 {
										holderstruct.movedir = 8
									}
									enemiesstructslayout[block] = enemy{}
									enemiesstructslayout[nextblock] = holderstruct
								} else {
									holderstruct.movedir = 8
									enemiesstructslayout[block] = enemy{}
									enemiesstructslayout[nextblock] = holderstruct
								}
								block = nextblock
							}
						}
						if rolldice()+rolldice() == 12 {
							holderstruct.jump = true
							if flipcoin() {
								holderstruct.movedir = 1
							} else {
								holderstruct.movedir = 3
							}
							enemiesstructslayout[block] = holderstruct
						}
					}
				} // end rabbit

				if enemieslayout[block] == "bee" {
					holderstruct := enemiesstructslayout[block]

					switch holderstruct.movedir {

					case 4:
						check, nextblock := checknextblock(block, 4)
						if check && moveenemycollisions(block, 4) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = "bee"
							enemiesstructslayout[block] = enemy{}
							if rolldice() == 6 {
								holderstruct.movedir = 6
							} else {
								choose := rInt(1, 6)
								switch choose {
								case 1:
									holderstruct.movedir = 1
								case 2:
									holderstruct.movedir = 2
								case 3:
									holderstruct.movedir = 3
								case 4:
									holderstruct.movedir = 4
								case 5:
									holderstruct.movedir = 8
								}
							}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							if rolldice() == 6 {
								holderstruct.movedir = 6
							} else {
								choose := rInt(1, 6)
								switch choose {
								case 1:
									holderstruct.movedir = 1
								case 2:
									holderstruct.movedir = 2
								case 3:
									holderstruct.movedir = 3
								case 4:
									holderstruct.movedir = 4
								case 5:
									holderstruct.movedir = 8
								}
							}
							enemiesstructslayout[nextblock] = holderstruct
						}
					case 6:
						check, nextblock := checknextblock(block, 6)
						if check && moveenemycollisions(block, 6) {
							enemieslayout[block] = ""
							enemieslayout[nextblock] = "bee"
							enemiesstructslayout[block] = enemy{}
							enemiesstructslayout[nextblock] = holderstruct
						} else {
							choose := rInt(1, 6)
							switch choose {
							case 1:
								holderstruct.movedir = 1
							case 2:
								holderstruct.movedir = 2
							case 3:
								holderstruct.movedir = 3
							case 4:
								holderstruct.movedir = 4
							case 5:
								holderstruct.movedir = 8
							}
							enemiesstructslayout[nextblock] = holderstruct
						}
					}
				} // end bee

				if enemieslayout[block] == "bird" {
					if moveenemycollisions(block, 4) {
						holderstruct := enemiesstructslayout[block]
						if holderstruct.movedir == 2 {

							if levellayout[block+1] == "." {
								enemieslayout[block] = ""
								enemieslayout[block+1] = "bird"
								enemiesstructslayout[block+1] = holderstruct
							} else if levellayout[block+1] == "_" || levellayout[block+1] == "#" || levellayout[block+1] == "%" {
								holderstruct.movedir = 4
								enemiesstructslayout[block] = holderstruct
							}
						}
					}
				} // end bird

				if enemieslayout[block] == "skull" {

					choose := rolldice()
					choose2 := rolldice() + rolldice()

					if flipcoin() {
						if choose < 4 {
							if moveenemycollisions(block, 6) {
								if levellayout[block+levelw] == "." {
									enemieslayout[block] = ""
									holderstruct := enemiesstructslayout[block]
									enemiesstructslayout[block] = enemy{}
									enemieslayout[block+levelw] = "skull"
									enemiesstructslayout[block+levelw] = holderstruct
								}
							}
						}
					} else {
						if choose2 > 10 {
							if flipcoin() {
								if moveenemycollisions(block, 4) {
									if levellayout[block+1] == "." {
										enemieslayout[block] = ""
										holderstruct := enemiesstructslayout[block]
										enemiesstructslayout[block] = enemy{}
										enemieslayout[block+1] = "skull"
										enemiesstructslayout[block+1] = holderstruct
									}
								}
							}
						}
					}
				} // end skull

				if enemieslayout[block] == "mushroom" {
					checkenemy := enemiesstructslayout[block]
					switch checkenemy.movedir {
					case 4:
						if rolldice() == 6 {
							checkenemy.movedir = 8
							enemiesstructslayout[block] = checkenemy
						}
						if rolldice()+rolldice() > 10 {
							checkenemy.movedir = 2
							enemiesstructslayout[block] = checkenemy
						}
						if levellayout[block+1] == "." {
							if enemieslayout[block+1] == "" {
								enemyvert := block % levelw
								if enemyvert < levelw-30 {
									holder := enemieslayout[block]
									enemieslayout[block] = ""
									holderstruct := enemiesstructslayout[block]
									enemiesstructslayout[block] = enemy{}
									enemiesstructslayout[block+1] = holderstruct
									enemieslayout[block+1] = holder
								} else {
									movenemycollisionsmushrooms(block+1, block, 4)
								}
							} else {
								checkenemy.movedir = 8
								enemiesstructslayout[block] = checkenemy
							}
						} else {
							checkenemy.movedir = 8
							enemiesstructslayout[block] = checkenemy
						}

					}
				} // end mushroom
				count++
				block--
				if count == updatew {
					count = 0
					block += updatew
					block -= levelw
				}
			}

			// mushroom gravity
			count = 0
			block = updateblockend
			for a := 0; a < updatea; a++ {
				if enemieslayout[block] == "mushroom" {
					if levellayout[block+levelw] != "_" && levellayout[block+levelw] != "#" {
						holder := enemieslayout[block]
						holderstruct := enemiesstructslayout[block]
						if holderstruct.movedir != 2 {
							if enemieslayout[block+levelw] == "" {
								enemieslayout[block] = ""
								enemiesstructslayout[block] = enemy{}
								enemiesstructslayout[block+levelw] = holderstruct
								enemieslayout[block+levelw] = holder
							} else {
								movenemycollisionsmushrooms(block, block+levelw, 6)
							}
						}
					}
				}
				count++
				block--
				if count == updatew {
					count = 0
					block += updatew
					block -= levelw
				}
			}
		}

	}

}
func moveenemycollisions(block, dir int) bool { // MARK: moveenemycollisions

	nocollision := true

	switch dir {
	case 1:
		if enemieslayout[(block-1)-levelw] != "" {
			nocollision = false
		}
	case 2:
		if enemieslayout[(block)-levelw] != "" {
			nocollision = false
		}
	case 3:
		if enemieslayout[(block+1)-levelw] != "" {
			nocollision = false
		}
	case 4:
		if enemieslayout[(block+1)] != "" {
			nocollision = false
		}
	case 5:
		if enemieslayout[(block+1)+levelw] != "" {
			nocollision = false
		}
	case 6:
		if enemieslayout[(block)+levelw] != "" {
			nocollision = false
		}
	case 7:
		if enemieslayout[(block-1)+levelw] != "" {
			nocollision = false
		}
	case 8:
		if enemieslayout[(block-1)] != "" {
			nocollision = false
		}
	}

	return nocollision

}
func movenemycollisionsmushrooms(block1, block2, direction int) { // MARK: movenemycollisionsmushrooms

	switch direction {
	case 8, 4:
		rightenemy := enemiesstructslayout[block1]
		leftenemy := enemiesstructslayout[block2]
		rightenemy.movedir = 4
		leftenemy.movedir = 8
		enemiesstructslayout[block1] = rightenemy
		enemiesstructslayout[block2] = leftenemy
	case 6, 2:
		holder := enemieslayout[block1]
		holderstruct := enemiesstructslayout[block1]
		if levellayout[block1+1] == "." {
			enemieslayout[block1] = ""
			enemiesstructslayout[block1] = enemy{}
			enemiesstructslayout[block1+1] = holderstruct
			enemieslayout[block1+1] = holder
		} else if levellayout[block1-1] == "." {
			enemieslayout[block1] = ""
			enemiesstructslayout[block1] = enemy{}
			enemiesstructslayout[block1-1] = holderstruct
			enemieslayout[block1-1] = holder
		} else if levellayout[block1-2] == "." {
			enemieslayout[block1] = ""
			enemiesstructslayout[block1] = enemy{}
			enemiesstructslayout[block1-2] = holderstruct
			enemieslayout[block1-2] = holder
		} else if levellayout[block1+2] == "." {
			enemieslayout[block1] = ""
			enemiesstructslayout[block1] = enemy{}
			enemiesstructslayout[block1+2] = holderstruct
			enemieslayout[block1+2] = holder
		} else {
			enemieslayout[block1] = ""
			enemiesstructslayout[block1] = enemy{}
			enemiesstructslayout[block1-levelw] = holderstruct
			enemieslayout[block1-levelw] = holder
		}
	}

}
func moveprojectiles() { // MARK: moveprojectiles

	// move bomb
	if bombon {
		bombh := bombblock / levelw
		switch bombdirection {
		case 2:
			if playerh-bombh < 4 && !bombdown {
				if levellayout[(bombblock-levelw)+1] == "." {
					weaponslayout[bombblock] = ""
					bombblock -= levelw
					bombblock++
					weaponslayout[bombblock] = "bomb"
				} else {
					bombdown = true
				}
			} else {
				bombdown = true
			}
			if bombdown {
				if levellayout[(bombblock+levelw)+1] == "." {
					weaponslayout[bombblock] = ""
					bombblock += levelw
					bombblock++
					weaponslayout[bombblock] = "bomb"
				} else {
					bombon = false
					bombdown = false
					explodeblock = bombblock - 5
					explodeblock -= levelw * 5
					bombexplode()
				}
			}
		case 4:
			if playerh-bombh < 4 && !bombdown {
				if levellayout[(bombblock-levelw)-1] == "." {
					weaponslayout[bombblock] = ""
					bombblock -= levelw
					bombblock--
					weaponslayout[bombblock] = "bomb"
				} else {
					bombdown = true
				}
			} else {
				bombdown = true
			}
			if bombdown {
				if levellayout[(bombblock+levelw)+1] == "." {
					weaponslayout[bombblock] = ""
					bombblock += levelw
					bombblock--
					weaponslayout[bombblock] = "bomb"
				} else {
					bombon = false
					bombdown = false
					explodeblock = bombblock - 5
					explodeblock -= levelw * 5
					bombexplode()
				}
			}
		}
	}

	// enemy bullets right down
	if plantactive {
		count := 0
		block := updateblockend
		for a := 0; a < updatea; a++ {
			if enemybulletslayout[block] == "bulletr" {
				check, nextblock := checknextblock(block, 4)
				if check {
					enemybulletslayout[block] = ""
					enemybulletslayout[nextblock] = "bulletr"
				} else {
					enemybulletslayout[block] = ""
				}
			}
			count++
			block--
			if count == updatew {
				count = 0
				block += updatew
				block -= levelw
			}
		}
	}

	// right down projectiles
	if boxinggloveactive || ninjastaractive || grenadactive || pumpactionactive || bazookaactive || uzziactive {
		count := 0
		block := updateblockend

		for a := 0; a < updatea; a++ {

			if uzziactive {
				if weaponslayout[block] == "uzzibulletr" {
					if levellayout[block+1] == "." {
						weaponslayout[block] = ""
						weaponslayout[block+1] = "uzzibulletr"
					} else if levellayout[block+1] != "." {
						weaponslayout[block] = ""
					}
				}
			}

			if bazookaactive {
				if weaponslayout[block] == "rocketr" {
					if levellayout[block+1] == "." && enemieslayout[block+1] == "" {
						weaponslayout[block] = ""
						weaponslayout[block+1] = "rocketr"
					} else if enemieslayout[block+1] != "" {
						killenemy(block + 1)
						weaponslayout[block] = ""
						explodeblock = block - 5
						explodeblock -= levelw * 5
						bombexplode()
					} else if levellayout[block+1] != "." {
						weaponslayout[block] = ""
						explodeblock = block - 5
						explodeblock -= levelw * 5
						bombexplode()
					}
				}
			}

			if pumpactionactive {
				if weaponslayout[block] == "shotgunbulletright" {
					if enemieslayout[block] != "" {
						checkenemy := enemiesstructslayout[block]
						checkenemy.hp -= 2
						if checkenemy.hp <= 0 {
							killenemy(block)
							weaponslayout[block] = ""
							if tntactive {
								explodeblock = block - 5
								explodeblock -= levelw * 5
								bombexplode()
							}
						} else {
							enemiesstructslayout[block] = checkenemy
							weaponslayout[block] = ""
						}
					}
					check, nextblock := checknextblock(block, 4)
					if check {
						weaponslayout[block] = ""
						weaponslayout[nextblock] = "shotgunbulletright"
					} else {
						weaponslayout[block] = ""
					}
				}
				if weaponslayout[block] == "shotgunbulletdownright" {
					if enemieslayout[block] != "" {
						checkenemy := enemiesstructslayout[block]
						checkenemy.hp -= 2
						if checkenemy.hp <= 0 {
							killenemy(block)
							weaponslayout[block] = ""
							if tntactive {
								explodeblock = block - 5
								explodeblock -= levelw * 5
								bombexplode()
							}
						} else {
							enemiesstructslayout[block] = checkenemy
							weaponslayout[block] = ""
						}
					}
					check, nextblock := checknextblock(block, 5)
					if check {
						weaponslayout[block] = ""
						weaponslayout[nextblock] = "shotgunbulletdownright"
					} else {
						weaponslayout[block] = ""
					}
				}
				if weaponslayout[block] == "shotgunbulletdownleft" {
					if enemieslayout[block] != "" {
						checkenemy := enemiesstructslayout[block]
						checkenemy.hp -= 2
						if checkenemy.hp <= 0 {
							killenemy(block)
							weaponslayout[block] = ""
							if tntactive {
								explodeblock = block - 5
								explodeblock -= levelw * 5
								bombexplode()
							}
						} else {
							enemiesstructslayout[block] = checkenemy
							weaponslayout[block] = ""
						}
					}
					check, nextblock := checknextblock(block, 7)
					if check {
						weaponslayout[block] = ""
						weaponslayout[nextblock] = "shotgunbulletdownleft"
					} else {
						weaponslayout[block] = ""
					}
				}
			}

			if grenadactive {
				if weaponslayout[block] == "grenade" {
					checkgrenade := grenadeslayout[block]
					if checkgrenade.direction == 2 {
						if levellayout[(block-1)+levelw] == "." {
							weaponslayout[block] = ""
							grenadeslayout[block] = grenadestruct{}
							checkgrenade.distance--
							weaponslayout[(block-1)+levelw] = "grenade"
							grenadeslayout[(block-1)+levelw] = checkgrenade
						} else {
							weaponslayout[block] = ""
							grenadeslayout[block] = grenadestruct{}
							explodeblock = block - 5
							explodeblock -= levelw * 5
							bombexplode()
						}
					}
					if checkgrenade.direction == 3 {
						if levellayout[(block+1)+levelw] == "." {
							weaponslayout[block] = ""
							grenadeslayout[block] = grenadestruct{}
							checkgrenade.distance--
							weaponslayout[(block+1)+levelw] = "grenade"
							grenadeslayout[(block+1)+levelw] = checkgrenade
						} else {
							weaponslayout[block] = ""
							grenadeslayout[block] = grenadestruct{}
							explodeblock = block - 5
							explodeblock -= levelw * 5
							bombexplode()
						}
					}
				}
			}

			if boxinggloveactive {
				if weaponslayout[block] == "boxingglover" || weaponslayout[block] == "boxingglove" {
					if enemieslayout[block] != "" {
						checkenemy := enemiesstructslayout[block]
						checkenemy.hp--
						if checkenemy.hp <= 0 {
							killenemy(block)
							weaponslayout[block] = ""
							if tntactive {
								explodeblock = block - 5
								explodeblock -= levelw * 5
								bombexplode()
							}
						} else {
							enemiesstructslayout[block] = checkenemy
							weaponslayout[block] = ""
						}
					} else if weaponslayout[block] == "boxingglover" {
						if levellayout[(block+1)+levelw] == "." {
							weaponslayout[block] = ""
							weaponslayout[(block+1)+levelw] = "boxingglover"
						}
					} else if weaponslayout[block] == "boxingglove" {
						if levellayout[(block-1)+levelw] == "." {
							weaponslayout[block] = ""
							weaponslayout[(block-1)+levelw] = "boxingglove"
						}
					}
				}
			}
			if ninjastaractive {
				if weaponslayout[block] == "ninjastar4" || weaponslayout[block] == "ninjastar5" || weaponslayout[block] == "ninjastar6" || weaponslayout[block] == "ninjastar7" {

					if enemieslayout[block] != "" {
						checkenemy := enemiesstructslayout[block]
						checkenemy.hp--
						if checkenemy.hp <= 0 {
							killenemy(block)
							weaponslayout[block] = ""
							if tntactive {
								explodeblock = block - 5
								explodeblock -= levelw * 5
								bombexplode()
							}
						} else {
							enemiesstructslayout[block] = checkenemy
							weaponslayout[block] = ""
						}
					} else {

						if weaponslayout[block] == "ninjastar4" && enemieslayout[block] == "" {
							if levellayout[(block+1)] == "." {
								weaponslayout[block] = ""
								weaponslayout[(block + 1)] = "ninjastar4"
							}
						}
						if weaponslayout[block] == "ninjastar5" && enemieslayout[block] == "" {
							if levellayout[(block+1)+levelw] == "_" && levellayout[(block+1)+(levelw*2)] == "." {
								weaponslayout[block] = ""
								weaponslayout[(block+1)+(levelw*2)] = "ninjastar5"
							}
						}
						if weaponslayout[block] == "ninjastar6" && enemieslayout[block] == "" {
							if levellayout[(block)+levelw] == "_" && levellayout[(block)+(levelw*2)] == "." {
								weaponslayout[block] = ""
								weaponslayout[(block)+(levelw*2)] = "ninjastar6"
							}
						}
						if weaponslayout[block] == "ninjastar7" && enemieslayout[block] == "" {
							if levellayout[(block-1)+levelw] == "_" && levellayout[(block-1)+(levelw*2)] == "." {
								weaponslayout[block] = ""
								weaponslayout[(block-1)+(levelw*2)] = "ninjastar7"
							}
						}
					}
				}
			}
			count++
			block--
			if count == updatew {
				count = 0
				block += updatew
				block -= levelw
			}

		}
	}

	// enemy bullets left up
	if plantactive {
		count := 0
		block := updateblock
		for a := 0; a < updatea; a++ {
			if enemybulletslayout[block] == "bulletl" {
				check, nextblock := checknextblock(block, 8)
				if check {
					enemybulletslayout[block] = ""
					enemybulletslayout[nextblock] = "bulletl"
				} else {
					enemybulletslayout[block] = ""
				}
			}

			count++
			block++
			if count == updatew {
				count = 0
				block -= updatew
				block += levelw
			}

		}
	}

	// left up projectiles
	if arrowactive || bulletactive || ninjastaractive || flameactive || grenadactive || pumpactionactive || bazookaactive || uzziactive {
		count := 0
		block := updateblock
		for a := 0; a < updatea; a++ {

			if uzziactive {
				if weaponslayout[block] == "uzzibulletl" {
					if levellayout[(block-1)] == "." {
						weaponslayout[block] = ""
						weaponslayout[(block - 1)] = "uzzibulletl"
					} else if levellayout[(block-1)] != "." {
						weaponslayout[block] = ""
					}
				}
				if weaponslayout[block] == "uzzibulletlup" {
					if levellayout[(block-1)-levelw] == "." {
						weaponslayout[block] = ""
						weaponslayout[(block-1)-levelw] = "uzzibulletlup"
					} else if levellayout[(block-1)-levelw] != "." {
						weaponslayout[block] = ""
					}
				}
				if weaponslayout[block] == "uzzibulletrup" {
					if levellayout[(block+1)-levelw] == "." {
						weaponslayout[block] = ""
						weaponslayout[(block+1)-levelw] = "uzzibulletrup"
					} else if levellayout[(block+1)-levelw] != "." {
						weaponslayout[block] = ""
					}
				}
				if weaponslayout[block] == "uzzibulletup" {
					if levellayout[(block)-levelw] == "." {
						weaponslayout[block] = ""
						weaponslayout[(block)-levelw] = "uzzibulletup"
					} else if levellayout[(block)-levelw] != "." {
						weaponslayout[block] = ""
					}
				}
			}

			if bazookaactive {
				if weaponslayout[block] == "rocketl" {
					if levellayout[block-1] == "." {
						weaponslayout[block] = ""
						weaponslayout[block-1] = "rocketl"
					} else if enemieslayout[block-1] != "" {
						killenemy(block - 1)
						weaponslayout[block] = ""
						explodeblock = block - 5
						explodeblock -= levelw * 5
						bombexplode()
					} else if levellayout[block-1] != "." {
						weaponslayout[block] = ""
						explodeblock = block - 5
						explodeblock -= levelw * 5
						bombexplode()
					}
				}
			}

			if pumpactionactive {
				if weaponslayout[block] == "shotgunbulletleft" {
					if enemieslayout[block] != "" {
						checkenemy := enemiesstructslayout[block]
						checkenemy.hp -= 2
						if checkenemy.hp <= 0 {
							killenemy(block)
							weaponslayout[block] = ""
							if tntactive {
								explodeblock = block - 5
								explodeblock -= levelw * 5
								bombexplode()
							}
						} else {
							enemiesstructslayout[block] = checkenemy
							weaponslayout[block] = ""
						}
					}
					check, nextblock := checknextblock(block, 8)
					if check {
						weaponslayout[block] = ""
						weaponslayout[nextblock] = "shotgunbulletleft"
					} else {
						weaponslayout[block] = ""
					}
				}
				if weaponslayout[block] == "shotgunbulletupleft" {
					if enemieslayout[block] != "" {
						checkenemy := enemiesstructslayout[block]
						checkenemy.hp -= 2
						if checkenemy.hp <= 0 {
							killenemy(block)
							weaponslayout[block] = ""
							if tntactive {
								explodeblock = block - 5
								explodeblock -= levelw * 5
								bombexplode()
							}
						} else {
							enemiesstructslayout[block] = checkenemy
							weaponslayout[block] = ""
						}
					}
					check, nextblock := checknextblock(block, 1)
					if check {
						weaponslayout[block] = ""
						weaponslayout[nextblock] = "shotgunbulletupleft"
					} else {
						weaponslayout[block] = ""
					}
				}
				if weaponslayout[block] == "shotgunbulletupright" {
					if enemieslayout[block] != "" {
						checkenemy := enemiesstructslayout[block]
						checkenemy.hp -= 2
						if checkenemy.hp <= 0 {
							killenemy(block)
							weaponslayout[block] = ""
							if tntactive {
								explodeblock = block - 5
								explodeblock -= levelw * 5
								bombexplode()
							}
						} else {
							enemiesstructslayout[block] = checkenemy
							weaponslayout[block] = ""
						}
					}
					check, nextblock := checknextblock(block, 3)
					if check {
						weaponslayout[block] = ""
						weaponslayout[nextblock] = "shotgunbulletupright"
					} else {
						weaponslayout[block] = ""
					}
				}
			}

			if grenadactive {
				if weaponslayout[block] == "grenade" {
					checkgrenade := grenadeslayout[block]
					if checkgrenade.direction == 0 || checkgrenade.direction == 1 {
						if checkgrenade.distance > 0 {
							if checkgrenade.direction == 0 {
								if levellayout[(block-1)-levelw] == "." {
									weaponslayout[block] = ""
									grenadeslayout[block] = grenadestruct{}
									checkgrenade.distance--
									weaponslayout[(block-1)-levelw] = "grenade"
									grenadeslayout[(block-1)-levelw] = checkgrenade
								} else {
									checkgrenade.direction = 2
									grenadeslayout[block] = checkgrenade
								}
							}
							if checkgrenade.direction == 1 {
								if levellayout[(block+1)-levelw] == "." {
									weaponslayout[block] = ""
									grenadeslayout[block] = grenadestruct{}
									checkgrenade.distance--
									weaponslayout[(block+1)-levelw] = "grenade"
									grenadeslayout[(block+1)-levelw] = checkgrenade
								} else {
									checkgrenade.direction = 3
									grenadeslayout[block] = checkgrenade
								}
							}
						} else {
							if checkgrenade.direction == 0 {
								checkgrenade.direction = 2
								grenadeslayout[block] = checkgrenade
							}
							if checkgrenade.direction == 1 {
								checkgrenade.direction = 3
								grenadeslayout[block] = checkgrenade
							}
						}
					}
				}
			}

			if bulletactive {
				if weaponslayout[block] == "bullet" {

					if enemieslayout[block] != "" {
						checkenemy := enemiesstructslayout[block]
						checkenemy.hp--
						if checkenemy.hp <= 0 {
							killenemy(block)
							weaponslayout[block] = ""
							if tntactive {
								explodeblock = block - 5
								explodeblock -= levelw * 5
								bombexplode()
							}
						} else {
							enemiesstructslayout[block] = checkenemy
							weaponslayout[block] = ""
						}
					} else {
						if levellayout[block-levelw] == "." {
							weaponslayout[block] = ""
							weaponslayout[block-levelw] = "bullet"
						}

					}
				}
			}
			if ninjastaractive {

				if weaponslayout[block] == "ninjastar1" || weaponslayout[block] == "ninjastar2" || weaponslayout[block] == "ninjastar3" || weaponslayout[block] == "ninjastar8" {

					if enemieslayout[block] != "" {
						checkenemy := enemiesstructslayout[block]
						checkenemy.hp--
						if checkenemy.hp <= 0 {
							killenemy(block)
							weaponslayout[block] = ""
							if tntactive {
								explodeblock = block - 5
								explodeblock -= levelw * 5
								bombexplode()
							}
						} else {
							enemiesstructslayout[block] = checkenemy
							weaponslayout[block] = ""
						}
					} else {

						if weaponslayout[block] == "ninjastar1" && enemieslayout[block] == "" {
							if levellayout[(block-1)-levelw] == "." {
								weaponslayout[block] = ""
								weaponslayout[(block-1)-levelw] = "ninjastar1"
							}
						}
						if weaponslayout[block] == "ninjastar2" && enemieslayout[block] == "" {
							if levellayout[(block)-levelw] == "." {
								weaponslayout[block] = ""
								weaponslayout[(block)-levelw] = "ninjastar2"
							}
						}
						if weaponslayout[block] == "ninjastar3" && enemieslayout[block] == "" {
							if levellayout[(block+1)-levelw] == "." {
								weaponslayout[block] = ""
								weaponslayout[(block+1)-levelw] = "ninjastar3"
							}
						}
						if weaponslayout[block] == "ninjastar8" && enemieslayout[block] == "" {
							if levellayout[(block-1)] == "." {
								weaponslayout[block] = ""
								weaponslayout[(block - 1)] = "ninjastar8"
							}
						}
					}
				}
			}
			if flameactive {

				if weaponslayout[block] == "flame" || weaponslayout[block] == "flame2" || weaponslayout[block] == "flame3" || weaponslayout[block] == "flame4" || weaponslayout[block] == "flame5" || weaponslayout[block] == "flame6" || weaponslayout[block] == "flame7" || weaponslayout[block] == "flame8" {

					if enemieslayout[block] != "" {
						checkenemy := enemiesstructslayout[block]
						checkenemy.hp--
						if checkenemy.hp <= 0 {
							killenemy(block)
							weaponslayout[block] = ""
							if tntactive {
								explodeblock = block - 5
								explodeblock -= levelw * 5
								bombexplode()
							}
						} else {
							enemiesstructslayout[block] = checkenemy
							weaponslayout[block] = ""
						}
					}

				}
			}
			if arrowactive {
				if weaponslayout[block] == "arrow" {
					if enemieslayout[block] != "" {
						checkenemy := enemiesstructslayout[block]
						checkenemy.hp--
						if checkenemy.hp <= 0 {
							killenemy(block)
							weaponslayout[block] = ""
							if tntactive {
								explodeblock = block - 5
								explodeblock -= levelw * 5
								bombexplode()
							}
						} else {
							enemiesstructslayout[block] = checkenemy
							weaponslayout[block] = ""
						}

					} else {
						if levellayout[(block+1)-levelw] == "." {
							weaponslayout[block] = ""
							weaponslayout[(block+1)-levelw] = "arrow"
						}
					}
				}
				if weaponslayout[block] == "arrowl" {
					if enemieslayout[block] != "" {
						checkenemy := enemiesstructslayout[block]
						checkenemy.hp--
						if checkenemy.hp <= 0 {
							killenemy(block)
							weaponslayout[block] = ""
							if tntactive {
								explodeblock = block - 5
								explodeblock -= levelw * 5
								bombexplode()
							}
						} else {
							enemiesstructslayout[block] = checkenemy
							weaponslayout[block] = ""
						}
					} else {
						if levellayout[(block-1)-levelw] == "." {
							weaponslayout[block] = ""
							weaponslayout[(block-1)-levelw] = "arrowl"
						}
					}
				}
			}

			count++
			block++
			if count == updatew {
				count = 0
				block -= updatew
				block += levelw
			}
		}
	}
}
func moveanimals() { // MARK: moveanimals
	if framecount%3 == 0 {
		// left up
		count := 0
		block := updateblock
		for a := 0; a < updatea; a++ {
			if animalslayout[block] != "" {
				checkanimal := animalstructslayout[block]
				switch checkanimal.movement {
				case 1:
					check, nextblock := checknextblock(block, 1)
					if check {
						if animalslayout[nextblock] == "" {
							animalslayout[block] = ""
							animalslayout[nextblock] = "bat"
							animalstructslayout[block] = animal{}
							animalstructslayout[nextblock] = checkanimal
						}
					}
				case 2:
					check, nextblock := checknextblock(block, 2)
					if check {
						if animalslayout[nextblock] == "" {
							animalslayout[block] = ""
							animalslayout[nextblock] = "bat"
							animalstructslayout[block] = animal{}
							animalstructslayout[nextblock] = checkanimal
						}
					}
				case 3:
					check, nextblock := checknextblock(block, 3)
					if check {
						if animalslayout[nextblock] == "" {
							animalslayout[block] = ""
							animalslayout[nextblock] = "bat"
							animalstructslayout[block] = animal{}
							animalstructslayout[nextblock] = checkanimal
						}
					}
				case 8:
					check, nextblock := checknextblock(block, 8)
					if check {
						if animalslayout[nextblock] == "" {
							animalslayout[block] = ""
							animalslayout[nextblock] = "bat"
							animalstructslayout[block] = animal{}
							animalstructslayout[nextblock] = checkanimal
						}
					}
				}
			}
			count++
			block++
			if count == updatew {
				count = 0
				block -= updatew
				block += levelw
			}
		}

		// right down
		count = 0
		block = updateblockend
		for a := 0; a < updatea; a++ {
			if animalslayout[block] != "" {
				checkanimal := animalstructslayout[block]
				switch checkanimal.movement {
				case 4:
					check, nextblock := checknextblock(block, 4)
					if check {
						if animalslayout[nextblock] == "" {
							animalslayout[block] = ""
							animalslayout[nextblock] = "bat"
							animalstructslayout[block] = animal{}
							animalstructslayout[nextblock] = checkanimal
						}
					}
				case 5:
					check, nextblock := checknextblock(block, 5)
					if check {
						if animalslayout[nextblock] == "" {
							animalslayout[block] = ""
							animalslayout[nextblock] = "bat"
							animalstructslayout[block] = animal{}
							animalstructslayout[nextblock] = checkanimal
						}
					}
				case 6:
					check, nextblock := checknextblock(block, 6)
					if check {
						if animalslayout[nextblock] == "" {
							animalslayout[block] = ""
							animalslayout[nextblock] = "bat"
							animalstructslayout[block] = animal{}
							animalstructslayout[nextblock] = checkanimal
						}
					}
				case 7:
					check, nextblock := checknextblock(block, 7)
					if check {
						if animalslayout[nextblock] == "" {
							animalslayout[block] = ""
							animalslayout[nextblock] = "bat"
							animalstructslayout[block] = animal{}
							animalstructslayout[nextblock] = checkanimal
						}
					}
				}

			}
			count++
			block--
			if count == updatew {
				count = 0
				block += updatew
				block -= levelw
			}
		}

		// change bat movement rotation
		count = 0
		block = updateblock
		for a := 0; a < updatea; a++ {

			if animalslayout[block] != "" {
				checkanimal := animalstructslayout[block]
				checkanimal.movement = rInt(1, 9)
				checkanimal.rotation = rFloat32(-45, 45)
				animalstructslayout[block] = checkanimal
			}
			count++
			block++
			if count == updatew {
				count = 0
				block -= updatew
				block += levelw
			}
		}
	}
}
func checknextblock(block, dir int) (bool, int) { // MARK: checknextblock

	check := true
	nextblock := block
	switch dir {
	case 1:
		if levellayout[(block-1)-levelw] == "." {
			nextblock = (block - 1) - levelw
		} else {
			check = false
		}
	case 2:
		if levellayout[(block)-levelw] == "." {
			nextblock = (block) - levelw
		} else {
			check = false
		}
	case 3:
		if levellayout[(block+1)-levelw] == "." {
			nextblock = (block + 1) - levelw
		} else {
			check = false
		}
	case 4:
		if levellayout[block+1] == "." {
			nextblock = block + 1
		} else {
			check = false
		}
	case 5:
		if levellayout[(block+1)+levelw] == "." {
			nextblock = (block + 1) + levelw
		} else {
			check = false
		}
	case 6:
		if levellayout[(block)+levelw] == "." {
			nextblock = (block) + levelw
		} else {
			check = false
		}
	case 7:
		if levellayout[(block-1)+levelw] == "." {
			nextblock = (block - 1) + levelw
		} else {
			check = false
		}
	case 8:
		if levellayout[block-1] == "." {
			nextblock = block - 1
		} else {
			check = false
		}

	}

	return check, nextblock

}
func bomb() { // MARK: bomb
	bombon = true
	bombblock = player

	switch playerdirection {
	case 2:
		bombdirection = 2
		if levellayout[(player-levelw)+1] != "_" && levellayout[(player-levelw)+1] != "#" {
			bombblock -= levelw
			bombblock++
			weaponslayout[bombblock] = "bomb"
		} else {
			if levellayout[(player-levelw)] != "_" && levellayout[(player-levelw)] != "#" {
				bombblock -= levelw
				weaponslayout[bombblock] = "bomb"
			} else {
				weaponslayout[player] = "bomb"
			}
		}
	case 4:
		bombdirection = 4
		if levellayout[(player-levelw)-1] != "_" && levellayout[(player-levelw)-1] != "#" {
			bombblock -= levelw
			bombblock--
			weaponslayout[bombblock] = "bomb"
		} else {
			if levellayout[(player-levelw)] != "_" && levellayout[(player-levelw)] != "#" {
				bombblock -= levelw
				weaponslayout[bombblock] = "bomb"
			} else {
				weaponslayout[player] = "bomb"
			}
		}

	}

}
func bombexplode() { // MARK: bombexplode
	explodeblockholder := explodeblock
	for a := 0; a < 10; a++ {
		length := rInt(8, 12)
		for b := 0; b < length; b++ {
			if levellayout[explodeblock] != "#" {
				levellayout[explodeblock] = "."
				if enemieslayout[explodeblock] != "rock" {
					enemieslayout[explodeblock] = ""
					enemiesstructslayout[explodeblock] = enemy{}
				}
			}
			if leveltileslayout[explodeblock] != "shopback1" && leveltileslayout[explodeblock] != "shopback2" && levellayout[explodeblock] != "#" {
				leveltileslayout[explodeblock] = "smoke3"
				scenerylayout[explodeblock] = ""
			}
			explodeblock += levelw
		}
		explodeblock = explodeblockholder
		explodeblock++
		explodeblockholder = explodeblock
		explodeblock += rInt(-2, 3) * levelw
	}
	weaponslayout[bombblock] = ""
	if rolldice() == 6 {
		lootlayout[bombblock] = "coin"
	}
	bombblock = 0
}
func floss(direction int) { // MARK: floss

	switch direction {
	case 1:

	case 2:
		flossdirection = 2
		check := player
		for {
			check -= levelw
			if levellayout[check] == "_" || levellayout[check] == "#" || levellayout[check] == "%" {
				flossendblock = check
				flossendblockh := flossendblock / levelw
				if playerh-flossendblockh > 18 {
					flossx = playerx
				}
				break
			}
		}
	case 3:

	}

}
func flossmovement() { // MARK: flossmovement

	switch flossdirection {
	case 2:
		if player != flossendblock+(levelw*2) {
			if levellayout[player-levelw] != "_" && levellayout[player-levelw] != "%" && levellayout[player-levelw] != "#" {
				player -= levelw
			} else if levellayout[player-levelw] == "%" {
				powerupblock(player - levelw)
				jumppause = 2
				flosson = false
			} else {
				jumppause = 2
				flosson = false
			}
		} else {
			flosson = false
		}
	}

}
func environment() { // MARK: environment
	if !gravityoff {
		// gravity
		if levellayout[player+levelw] != "_" && levellayout[player+levelw] != "#" && levellayout[player+levelw] != "%" {
			if jumpactive == false && flosson == false && jumpactivespring == false {
				player += levelw
				fallactive = true
				fallscreenshake = false
			}
		} else if levellayout[player+levelw] == "_" || levellayout[player+levelw] == "#" || levellayout[player+levelw] == "%" && fallactive == true {
			fallactive = false
			if fallscreenshake == false {
				screenshakeactive = true
				fallscreenshake = true
			}
		}
	}
}
func killenemy(block int) { // MARK: killenemy
	enemieslayout[block] = ""
	enemiesstructslayout[block] = enemy{}
	deadenemieslayout[block] = "deadenemy"
	if moneybagactive {
		if rolldice() == 6 {
			lootlayout[block+1] = "coin"
		}
	}
	if tntactive {
		explodeblock = block - 5
		explodeblock -= levelw * 5
		bombexplode()
	}

}
func killallenemies() { // MARK: killallenemies

	count := 0
	block := updateblock
	for a := 0; a < updatea; a++ {

		if enemieslayout[block] != "" {
			killenemy(block)
		}

		count++
		block++
		if count == updatew {
			count = 0
			block -= updatew
			block += levelw
		}
	}

}
func doorjump() { // MARK: doorjump

	if doorslayout[player] != "" {
		player = doorpositionslayout[player]

	} else if doorslayout[player+1] != "" {
		player = doorpositionslayout[player+1]
	} else if doorslayout[player-1] != "" {
		player = doorpositionslayout[player-1]
	}

}
func powerupblock(block int) { // MARK: powerupblock

	powerupblockmoveon = true
	powerupblockmovetimer = 2

	choose := rInt(1, 18)

	switch choose {
	case 1:
		newpowerup := powerup{}
		newpowerup.active = true
		newpowerup.name = "coffee"
		newpowerup.img = caffeinejump
		newpowerup.number = 1
		powerupslayout[block-levelw] = "coffee"
		powerupsstructslayout[block-levelw] = newpowerup
	case 2:
		newpowerup := powerup{}
		newpowerup.active = true
		newpowerup.name = "tnt"
		newpowerup.img = tntenemiesexplode
		newpowerup.number = 1
		powerupslayout[block-levelw] = "tnt"
		powerupsstructslayout[block-levelw] = newpowerup
	case 3:
		newpowerup := powerup{}
		newpowerup.active = true
		newpowerup.name = "rip"
		newpowerup.img = riprebirth
		newpowerup.number = 1
		powerupslayout[block-levelw] = "rip"
		powerupsstructslayout[block-levelw] = newpowerup
	case 4:
		newpowerup := powerup{}
		newpowerup.active = true
		newpowerup.name = "moneybag"
		newpowerup.img = moneybag
		newpowerup.number = 1
		powerupslayout[block-levelw] = "moneybag"
		powerupsstructslayout[block-levelw] = newpowerup
	case 5:
		newpowerup := powerup{}
		newpowerup.active = true
		newpowerup.name = "zippo"
		newpowerup.img = zippoburning
		newpowerup.number = 1
		powerupslayout[block-levelw] = "zippo"
		powerupsstructslayout[block-levelw] = newpowerup
	case 6:
		newpowerup := powerup{}
		newpowerup.active = true
		newpowerup.name = "pegasus"
		newpowerup.img = pegasusflying
		newpowerup.number = 1
		powerupslayout[block-levelw] = "pegasus"
		powerupsstructslayout[block-levelw] = newpowerup
	case 7:
		newpowerup := powerup{}
		newpowerup.active = true
		newpowerup.name = "arrow"
		newpowerup.img = arrowrandom
		newpowerup.number = 1
		powerupslayout[block-levelw] = "arrow"
		powerupsstructslayout[block-levelw] = newpowerup
	case 8:
		newpowerup := powerup{}
		newpowerup.active = true
		newpowerup.name = "meteor"
		newpowerup.img = meteorsml
		newpowerup.number = 1
		powerupslayout[block-levelw] = "meteor"
		powerupsstructslayout[block-levelw] = newpowerup
	case 9:
		newpowerup := powerup{}
		newpowerup.active = true
		newpowerup.name = "ninjastar"
		newpowerup.img = ninjastar
		newpowerup.number = 1
		powerupslayout[block-levelw] = "ninjastar"
		powerupsstructslayout[block-levelw] = newpowerup
	case 10:
		newpowerup := powerup{}
		newpowerup.active = true
		newpowerup.name = "colorize"
		newpowerup.img = colorize
		newpowerup.number = 1
		powerupslayout[block-levelw] = "colorize"
		powerupsstructslayout[block-levelw] = newpowerup
	case 11:
		newpowerup := powerup{}
		newpowerup.active = true
		newpowerup.name = "bullet"
		newpowerup.img = bulletrandom
		newpowerup.number = 1
		powerupslayout[block-levelw] = "bullet"
		powerupsstructslayout[block-levelw] = newpowerup
	case 12:
		newpowerup := powerup{}
		newpowerup.active = true
		newpowerup.name = "grenade"
		newpowerup.img = grenade
		newpowerup.number = 1
		powerupslayout[block-levelw] = "grenade"
		powerupsstructslayout[block-levelw] = newpowerup
	case 13:
		newpowerup := powerup{}
		newpowerup.active = true
		newpowerup.name = "warp"
		newpowerup.img = warp
		newpowerup.number = 1
		powerupslayout[block-levelw] = "warp"
		powerupsstructslayout[block-levelw] = newpowerup
	case 14:
		newpowerup := powerup{}
		newpowerup.active = true
		newpowerup.name = "boxingglove"
		newpowerup.img = boxingglove
		newpowerup.number = 1
		powerupslayout[block-levelw] = "boxingglove"
		powerupsstructslayout[block-levelw] = newpowerup
	case 15:
		newpowerup := powerup{}
		newpowerup.active = true
		newpowerup.name = "fireball"
		newpowerup.img = fireball
		newpowerup.number = 1
		powerupslayout[block-levelw] = "fireball"
		powerupsstructslayout[block-levelw] = newpowerup
	case 16:
		newpowerup := powerup{}
		newpowerup.active = true
		newpowerup.name = "healthpotion"
		newpowerup.img = healthpotion
		newpowerup.number = 1
		powerupslayout[block-levelw] = "healthpotion"
		powerupsstructslayout[block-levelw] = newpowerup
	case 17:
		newpowerup := powerup{}
		newpowerup.active = true
		newpowerup.name = "random7"
		newpowerup.img = random7
		newpowerup.number = 1
		powerupslayout[block-levelw] = "random7"
		powerupsstructslayout[block-levelw] = newpowerup
	}

	levellayout[block] = "#"
	leveltileslayout[block] = "&"

}
func collectpowerup(block int) { // MARK: collectpowerup

	checkpowerup := powerupsstructslayout[block]

	switch checkpowerup.name {
	case "fireball":
		fireballactive = true
	case "boxingglove":
		boxinggloveactive = true
	case "warp":
		for {
			choose := rInt(levelw*50, levela-(levelw*50))
			if levellayout[choose] == "." {
				player = choose
				break
			}
		}
	case "grenade":
		grenadactive = true
	case "bullet":
		bulletactive = true
	case "colorize":
		colorizeactive = true
		othercolor1 = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
		othercolor2 = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
		othercolor3 = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
		othercolor4 = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
		tilecolor = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
		enemycolor1 = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
		enemycolor2 = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
		enemycolor3 = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
		enemycolor4 = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
		enemycolor5 = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
		enemycolor6 = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
		enemycolor7 = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
		enemycolor8 = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
		scenerycolor1 = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
		scenerycolor2 = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
		scenerycolor3 = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
		scenerycolor4 = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
		scenerycolor5 = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
		scenerycolor6 = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
		scenerycolor7 = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
		scenerycolor8 = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
		platformcolor = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
	case "ninjastar":
		ninjastaractive = true
		ninjastarcount++
	case "meteor":
		meteoractive = true
	case "arrow":
		arrowactive = true
		arrowcount++
	case "pegasus":
		gravityoff = true
		pegasustimer = 10
	case "zippo":
		flameactive = true
	case "moneybag":
		moneybagactive = true
	case "tnt":
		tntactive = true
	case "coffee":
		jumpheight += 2
		coffeecount++
	case "healthpotion":
		hp++
	case "random7":
		choose := rInt(1, 21)
		switch choose {
		case 1:
			checkpowerup.name = "coffee"
			checkpowerup.img = caffeinejump
			checkpowerup.number = 1
		case 2:
			checkpowerup.name = "tnt"
			checkpowerup.img = tntenemiesexplode
			checkpowerup.number = 1
		case 3:
			checkpowerup.name = "rip"
			checkpowerup.img = riprebirth
			checkpowerup.number = 1
		case 4:
			checkpowerup.name = "moneybag"
			checkpowerup.img = moneybag
			checkpowerup.number = 1
		case 5:
			checkpowerup.name = "zippo"
			checkpowerup.img = zippoburning
			checkpowerup.number = 1
		case 6:
			checkpowerup.name = "pegasus"
			checkpowerup.img = pegasusflying
			checkpowerup.number = 1
		case 7:
			checkpowerup.name = "arrow"
			checkpowerup.img = arrowrandom
			checkpowerup.number = 1
		case 8:
			checkpowerup.name = "meteor"
			checkpowerup.img = meteorsml
			checkpowerup.number = 1
		case 9:
			checkpowerup.name = "ninjastar"
			checkpowerup.img = ninjastar
			checkpowerup.number = 1
		case 10:
			checkpowerup.name = "colorize"
			checkpowerup.img = colorize
			checkpowerup.number = 1
		case 11:
			checkpowerup.name = "bullet"
			checkpowerup.img = bulletrandom
			checkpowerup.number = 1
		case 12:
			checkpowerup.name = "grenade"
			checkpowerup.img = grenade
			checkpowerup.number = 1
		case 13:
			checkpowerup.name = "warp"
			checkpowerup.img = warp
			checkpowerup.number = 1
		case 14:
			checkpowerup.name = "boxingglove"
			checkpowerup.img = boxingglove
			checkpowerup.number = 1
		case 15:
			checkpowerup.name = "fireball"
			checkpowerup.img = fireball
			checkpowerup.number = 1
		}

	}

	collected := false
	for a := 0; a < len(powerupscollected); a++ {
		collectedpowerup := powerupscollected[a]
		if collectedpowerup.name == checkpowerup.name {
			if collectedpowerup.name != "colorize" && collectedpowerup.name != "meteor" && collectedpowerup.name != "grenade" && collectedpowerup.name != "pegasus" && collectedpowerup.name != "tnt" && collectedpowerup.name != "boxinglove" && collectedpowerup.name != "bullet" && collectedpowerup.name != "moneybag" && collectedpowerup.name != "warp" {
				collectedpowerup.number++
				powerupscollected[a] = collectedpowerup
			}
			powerupsstructslayout[block] = powerup{}
			powerupslayout[block] = ""
			collected = true
		}
	}

	if !collected {
		if checkpowerup.name != "healthpotion" && checkpowerup.name != "random7" && checkpowerup.name != "warp" {
			powerupscollected[powerupcount] = checkpowerup
			powerupsstructslayout[block] = powerup{}
			powerupslayout[block] = ""
			powerupcount++
		} else {
			powerupsstructslayout[block] = powerup{}
			powerupslayout[block] = ""
		}
	}

}
func collect() { // MARK: collect

	if lootlayout[player] == "spade" {
		lootlayout[player] = ""
		spadeactive = true
	} else if lootlayout[player] == "pickaxe" {
		lootlayout[player] = ""
		pickaxeactive = true
	} else if lootlayout[player] == "icewand" {
		lootlayout[player] = ""
		icewandactive = true
	} else if lootlayout[player] == "sword" {
		lootlayout[player] = ""
		bigswordactive = true
	} else if lootlayout[player] == "uzzi" {
		lootlayout[player] = ""
		uzziactive = true
	} else if lootlayout[player] == "bazooka" {
		lootlayout[player] = ""
		bazookaactive = true
	} else if lootlayout[player] == "axe" {
		lootlayout[player] = ""
		axeactive = true
	} else if lootlayout[player] == "pumpaction" {
		lootlayout[player] = ""
		pumpactiontimer = 10
		pumpactionframe = 0
		pumpactionactive = true
	} else if lootlayout[player] == "coin" {
		coinstotal++
		lootlayout[player] = ""
	} else if lootlayout[player] == "gem" {
		gemstotal++
		lootlayout[player] = ""
	} else if lootlayout[player] == "bombcollect" {
		bombtotal++
		lootlayout[player] = ""
	}
}
func openchest(block int) { // MARK: openchest

	if rolldice() < 4 {
		lootlayout[block-levelw] = "bombcollect"
	} else {
		choose := rInt(1, 9)
		switch choose {
		case 1:
			lootlayout[block-levelw] = "pumpaction"
		case 2:
			lootlayout[block-levelw] = "axe"
		case 3:
			lootlayout[block-levelw] = "bazooka"
		case 4:
			lootlayout[block-levelw] = "uzzi"
		case 5:
			lootlayout[block-levelw] = "sword"
		case 6:
			lootlayout[block-levelw] = "icewand"
		case 7:
			lootlayout[block-levelw] = "pickaxe"
		case 8:
			lootlayout[block-levelw] = "spade"
		}
	}

}
func readbook() { // MARK: readbook

	choose := rInt(1, 11)

	choose = 3

	switch choose {
	case 3:
		extraslayout[player+2] = "campfire"
	case 1:
		// flip camera
		flipscreen = true
		camera.Zoom = 4.0
		camera.Rotation = 180.0
		camera.Target.X = float32(playerx + (monw / 8))
		camera.Target.Y = float32(playery + (monh / 8))
	case 2:
		// shield
		shieldon = true
	}

}
func createspikes() { // MARK: createspikes

	number := 200

	for {
		block := rInt(levelw*150, levela-(levelw*100))
		if levellayout[block] == "." {
			if levellayout[block+levelw] == "_" {
				trapslayout[block-levelw] = "spike3"
				number--
			}
		}
		if number <= 0 {
			break
		}
	}

}
func createendkey() { // MARK: createendkey

	for {
		block := rInt(levelw*100, (levela - (levelw * 100)))
		if levellayout[block] == "." {
			if levellayout[block+levelw] == "_" || levellayout[block+levelw] == "#" {
				endkeyblock = block
				endkeynumber = rInt(1, 10)
				break
			}
		}
	}

}
func createshopitems(block int) { // MARK: createshopitems

	newitem := shopitem{}
	newitem.price = rInt(8, 21)
	newitem.rotation = rInt(1, 4)

	choose := rInt(1, 9)

	switch choose {
	case 1:
		newitem.item = "colorize"
	case 2:
		newitem.item = "coffee"
	case 3:
		newitem.item = "tnt"
	case 4:
		newitem.item = "zippo"
	case 5:
		newitem.item = "arrow"
	case 6:
		newitem.item = "ninjastar"
	case 7:
		newitem.item = "warp"
	case 8:
		newitem.item = "random7"
	}
	shopitemsstructs[block] = newitem
	shopitemslayout[block] = newitem.item

}
func createcharacters() { // MARK: createcharacters

	characterslayout[player+10] = "character1"
	block := (player + levelw) - 5
	for a := 0; a < 20; a++ {
		levellayout[block] = "_"
		leveltileslayout[block] = "_"
		if rolldice() < 3 {
			levellayout[block+levelw] = "_"
			leveltileslayout[block+levelw] = "_"
		} else if rolldice() > 4 {
			levellayout[block+levelw] = "_"
			leveltileslayout[block+levelw] = "_"
			levellayout[block+(levelw*2)] = "_"
			leveltileslayout[block+(levelw*2)] = "_"
		}
		block++
	}

}
func createinteractables() { // MARK: createinteractables

	for a := 0; a < 200; a++ {
		for {
			block := rInt(levelw*50, levelw*(levelh-50))
			if levellayout[block] == "." && levellayout[block+levelw] == "_" {
				interactableslayout[block] = "z"
				break
			}
		}
	}

}
func createquest() { // MARK: createquest

	choose := rInt(1, 4)

	switch choose {
	case 1:
		activequest.questtype = "find"
	case 2:
		activequest.questtype = "rescue"
	}

	for {
		block := rInt(levelw*50, levelw*(levelh-50))
		if levellayout[block] == "." && levellayout[block+levelw] == "_" {
			activequest.block = block
			break
		}
	}

	switch activequest.questtype {
	case "rescue":
		choose = rInt(1, 17)
		switch choose {
		case 1:
			activequest.item = "the budgie"
			activequest.img = budgie
		case 2:
			activequest.item = "the dog"
			activequest.img = dog
		case 3:
			activequest.item = "the snail"
			activequest.img = snail
		case 4:
			activequest.item = "the ladybug"
			activequest.img = ladybug
		case 5:
			activequest.item = "the pigeon"
			activequest.img = pigeon
		case 6:
			activequest.item = "the blob"
			activequest.img = blob
		case 7:
			activequest.item = "the octopus"
			activequest.img = octopus
		case 8:
			activequest.item = "the koala"
			activequest.img = koala
		case 9:
			activequest.item = "the fox"
			activequest.img = fox
		case 10:
			activequest.item = "the santa"
			activequest.img = santa
		case 11:
			activequest.item = "the fly"
			activequest.img = fly
		case 12:
			activequest.item = "the mouse"
			activequest.img = mouse
		case 13:
			activequest.item = "the fly"
			activequest.img = fly2
		case 14:
			activequest.item = "the blob"
			activequest.img = blob2
		case 15:
			activequest.item = "the cat"
			activequest.img = cat
		case 16:
			activequest.item = "the snake"
			activequest.img = snake
		}
	case "find":
		choose = rInt(1, 26)
		switch choose {
		case 1:
			activequest.item = "socks"
			activequest.img = socks
		case 2:
			activequest.item = "disk"
			activequest.img = disk
		case 3:
			activequest.item = "cowboy hat"
			activequest.img = cowboyhat
		case 4:
			activequest.item = "gingerbread man"
			activequest.img = gingerbread
		case 5:
			activequest.item = "ice cream"
			activequest.img = icecream
		case 6:
			activequest.item = "oven glove"
			activequest.img = ovenmit
		case 7:
			activequest.item = "chef hat"
			activequest.img = chefhat
		case 8:
			activequest.item = "anchor"
			activequest.img = anchor
		case 9:
			activequest.item = "xmas stocking"
			activequest.img = xmasstocking
		case 10:
			activequest.item = "pants"
			activequest.img = pants
		case 11:
			activequest.item = "santa's hat"
			activequest.img = santahat
		case 12:
			activequest.item = "sunglasses"
			activequest.img = sunglasses
		case 13:
			activequest.item = "coat hanger"
			activequest.img = coathanger
		case 14:
			activequest.item = "socks"
			activequest.img = socks
		case 15:
			activequest.item = "ribbon"
			activequest.img = ribbon
		case 16:
			activequest.item = "t-shirt"
			activequest.img = tshirt
		case 17:
			activequest.item = "top hat"
			activequest.img = socks
		case 18:
			activequest.item = "slipper"
			activequest.img = slipper
		case 19:
			activequest.item = "puzzle piece"
			activequest.img = puzzlepiece
		case 20:
			activequest.item = "guitar"
			activequest.img = guitar
		case 21:
			activequest.item = "ring"
			activequest.img = ring
		case 22:
			activequest.item = "teddy bear"
			activequest.img = teddybear
		case 23:
			activequest.item = "aladdin's lamp"
			activequest.img = aladdinslamp
		case 24:
			activequest.item = "lego piece"
			activequest.img = legopiece
		case 25:
			activequest.item = "alien"
			activequest.img = alien
		}

	}

}
func createninjastar() { // MARK: createninjastar

	for a := 0; a < ninjastarcount; a++ {
		choose := rInt(1, 9)
		switch choose {
		case 1:
			if levellayout[(player-1)-levelw] == "." {
				weaponslayout[(player-1)-levelw] = "ninjastar1"
			}
		case 2:
			if levellayout[(player)-levelw] == "." {
				weaponslayout[(player)-levelw] = "ninjastar2"
			}
		case 3:
			if levellayout[(player+1)-levelw] == "." {
				weaponslayout[(player+1)-levelw] = "ninjastar3"
			}
		case 4:
			if levellayout[(player+1)] == "." {
				weaponslayout[(player + 1)] = "ninjastar4"
			}
		case 5:
			if levellayout[(player+1)+levelw] == "." {
				weaponslayout[(player+1)+levelw] = "ninjastar5"
			}
		case 6:
			if levellayout[(player)+levelw] == "." {
				weaponslayout[(player)+levelw] = "ninjastar6"
			}
		case 7:
			if levellayout[(player-1)+levelw] == "." {
				weaponslayout[(player-1)+levelw] = "ninjastar7"
			}
		case 8:
			if levellayout[(player-1)] == "." {
				weaponslayout[(player - 1)] = "ninjastar8"
			}
		}
	}

}
func createarrow() { // MARK: createarrow
	if playerdirection == 2 {
		if levellayout[(player+1)-levelw] == "." {
			weaponslayout[(player+1)-levelw] = "arrow"
		}
	} else {
		if levellayout[(player-1)-levelw] == "." {
			weaponslayout[(player-1)-levelw] = "arrowl"
		}
	}
	if arrowcount > 1 {
		if arrowcount == 2 {
			if playerdirection == 2 {
				if levellayout[(player+1)-(levelw*2)] == "." {
					weaponslayout[(player+1)-(levelw*2)] = "arrow"
				}
			} else {
				if levellayout[(player-1)-(levelw*2)] == "." {
					weaponslayout[(player-1)-(levelw*2)] = "arrowl"
				}
			}
		} else if arrowcount == 3 {
			if playerdirection == 2 {
				if levellayout[(player+1)-(levelw*2)] == "." {
					weaponslayout[(player+1)-(levelw*2)] = "arrow"
				}
				if levellayout[(player+1)-(levelw*3)] == "." {
					weaponslayout[(player+1)-(levelw*3)] = "arrow"
				}
			} else {
				if levellayout[(player-1)-(levelw*2)] == "." {
					weaponslayout[(player-1)-(levelw*2)] = "arrowl"
				}
				if levellayout[(player-1)-(levelw*3)] == "." {
					weaponslayout[(player-1)-(levelw*3)] = "arrowl"
				}
			}
		} else if arrowcount == 4 {
			if playerdirection == 2 {
				if levellayout[(player+1)-(levelw*2)] == "." {
					weaponslayout[(player+1)-(levelw*2)] = "arrow"
				}
				if levellayout[(player+1)-(levelw*3)] == "." {
					weaponslayout[(player+1)-(levelw*3)] = "arrow"
				}
				if levellayout[(player+1)-(levelw*4)] == "." {
					weaponslayout[(player+1)-(levelw*4)] = "arrow"
				}
			} else {
				if levellayout[(player-1)-(levelw*2)] == "." {
					weaponslayout[(player-1)-(levelw*2)] = "arrowl"
				}
				if levellayout[(player-1)-(levelw*3)] == "." {
					weaponslayout[(player-1)-(levelw*3)] = "arrowl"
				}
				if levellayout[(player-1)-(levelw*4)] == "." {
					weaponslayout[(player-1)-(levelw*4)] = "arrowl"
				}
			}
		} else if arrowcount >= 5 {
			if playerdirection == 2 {
				if levellayout[(player+1)-(levelw*2)] == "." {
					weaponslayout[(player+1)-(levelw*2)] = "arrow"
				}
				if levellayout[(player+1)-(levelw*3)] == "." {
					weaponslayout[(player+1)-(levelw*3)] = "arrow"
				}
				if levellayout[(player+1)-(levelw*4)] == "." {
					weaponslayout[(player+1)-(levelw*4)] = "arrow"
				}
				if levellayout[(player+1)-(levelw*5)] == "." {
					weaponslayout[(player+1)-(levelw*5)] = "arrow"
				}
			} else {
				if levellayout[(player-1)-(levelw*2)] == "." {
					weaponslayout[(player-1)-(levelw*2)] = "arrowl"
				}
				if levellayout[(player-1)-(levelw*3)] == "." {
					weaponslayout[(player-1)-(levelw*3)] = "arrowl"
				}
				if levellayout[(player-1)-(levelw*4)] == "." {
					weaponslayout[(player-1)-(levelw*4)] = "arrowl"
				}
				if levellayout[(player-1)-(levelw*5)] == "." {
					weaponslayout[(player-1)-(levelw*5)] = "arrowl"
				}
			}
		}

	}
}
func creategrenade() { // MARK: creategrenade

	grenadeblock := player - levelw

	newgrenade := grenadestruct{}

	newgrenade.direction = rInt(0, 2)
	newgrenade.distance = rInt(4, 12)
	if newgrenade.distance%2 != 0 {
		newgrenade.distance++
	}

	if newgrenade.direction == 0 {
		grenadeblock -= 2
	}

	if newgrenade.direction == 1 {
		grenadeblock += 2
	}

	grenadeslayout[grenadeblock] = newgrenade
	weaponslayout[grenadeblock] = "grenade"

}
func createfireball() { // MARK: createfireball
	if playerdirection == 2 {
		if levellayout[player+2] == "." {
			weaponslayout[player+2] = "fireball"
		}
	}

	if playerdirection == 4 {
		if levellayout[player-2] == "." {
			weaponslayout[player-2] = "fireballl"
		}
	}

}
func createcoinsgems() { // MARK: createcoins

	for a := 0; a < 100; a++ {
		for {
			choose := rInt(levelw*20, levela-(levelw*5))
			if levellayout[choose] == "." {
				lootlayout[choose] = "coin"
				break
			}
		}
	}
	for a := 0; a < 20; a++ {
		for {
			choose := rInt(levelw*20, levela-(levelw*5))
			if levellayout[choose] == "." && lootlayout[choose] == "" {
				lootlayout[choose] = "gem"
				break
			}
		}
	}

}
func createstartenemies() { // MARK: createstartenemies

	mushroomactive = true

	enemy1 := rInt(1, 14)
	enemy2 := rInt(1, 14)
	enemy3 := rInt(1, 14)

	enemytype1 := "floatingskull"
	enemytype2 := "floatingskull"
	enemytype3 := "floatingskull"

	switch enemy1 {
	case 1:
		enemytype1 = "skull"
		skullactive = true
	case 2:
		enemytype1 = "bird"
		birdactive = true
	case 3:
		enemytype1 = "rabbit"
		rabbitactive = true
	case 4:
		enemytype1 = "bee"
		beeactive = true
	case 5:
		enemytype1 = "slime"
		slimeactive = true
	case 6:
		enemytype1 = "plant"
		plantactive = true
	case 7:
		enemytype1 = "spikes"
		spikesactive = true
	case 8:
		enemytype1 = "radish"
		radishactive = true
	case 9:
		enemytype1 = "snail"
		snailactive = true
	case 10:
		enemytype1 = "rock"
		rockactive = true
	case 11:
		enemytype1 = "bat"
		bigbatactive = true
	case 12:
		enemytype1 = "ghost"
		ghostactive = true
	case 13:
		enemytype1 = "pig"
		pigactive = true
	}
	switch enemy2 {
	case 1:
		enemytype2 = "skull"
		skullactive = true
	case 2:
		enemytype2 = "bird"
		birdactive = true
	case 3:
		enemytype2 = "rabbit"
		rabbitactive = true
	case 4:
		enemytype2 = "bee"
		beeactive = true
	case 5:
		enemytype2 = "slime"
		slimeactive = true
	case 6:
		enemytype2 = "plant"
		plantactive = true
	case 7:
		enemytype2 = "spikes"
		spikesactive = true
	case 8:
		enemytype2 = "radish"
		radishactive = true
	case 9:
		enemytype2 = "snail"
		snailactive = true
	case 10:
		enemytype2 = "rock"
		rockactive = true
	case 11:
		enemytype2 = "bat"
		bigbatactive = true
	case 12:
		enemytype2 = "ghost"
		ghostactive = true
	case 13:
		enemytype2 = "pig"
		pigactive = true
	}
	switch enemy3 {
	case 1:
		enemytype3 = "skull"
		skullactive = true
	case 2:
		enemytype3 = "bird"
		birdactive = true
	case 3:
		enemytype3 = "rabbit"
		rabbitactive = true
	case 4:
		enemytype3 = "bee"
		beeactive = true
	case 5:
		enemytype3 = "slime"
		slimeactive = true
	case 6:
		enemytype3 = "plant"
		plantactive = true
	case 7:
		enemytype3 = "spikes"
		spikesactive = true
	case 8:
		enemytype3 = "radish"
		radishactive = true
	case 9:
		enemytype3 = "snail"
		snailactive = true
	case 10:
		enemytype3 = "rock"
		rockactive = true
	case 11:
		enemytype3 = "bat"
		bigbatactive = true
	case 12:
		enemytype3 = "ghost"
		ghostactive = true
	case 13:
		enemytype3 = "pig"
		pigactive = true
	}

	// create enemy1
	for a := 0; a < 200; a++ {
		for {
			block := rInt(levelw*50, levelw*(levelh-50))
			if levellayout[block] == "." && levellayout[block+levelw] == "." && enemieslayout[block] == "" {
				enemieslayout[block] = enemytype1
				newenemy := enemy{}
				newenemy.hp = 2
				if flipcoin() {
					newenemy.movedir = 2
				} else {
					newenemy.movedir = 4
				}
				enemiesstructslayout[block] = newenemy
				break
			}
		}
	}

	// create enemy2
	for a := 0; a < 200; a++ {
		for {
			block := rInt(levelw*50, levelw*(levelh-50))
			if levellayout[block] == "." && levellayout[block+levelw] == "." && enemieslayout[block] == "" {
				enemieslayout[block] = enemytype2
				newenemy := enemy{}
				newenemy.hp = 2
				if flipcoin() {
					newenemy.movedir = 2
				} else {
					newenemy.movedir = 4
				}
				enemiesstructslayout[block] = newenemy
				break
			}
		}
	}
	// create enemy3
	for a := 0; a < 200; a++ {
		for {
			block := rInt(levelw*50, levelw*(levelh-50))
			if levellayout[block] == "." && levellayout[block+levelw] == "." && enemieslayout[block] == "" {
				enemieslayout[block] = enemytype3
				newenemy := enemy{}
				newenemy.hp = 2
				if flipcoin() {
					newenemy.movedir = 2
				} else {
					newenemy.movedir = 4
				}
				enemiesstructslayout[block] = newenemy
				break
			}
		}
	}

}
func createenemies() { // MARK: createenemies

	currentenemytype := "warpmushroom"
	enemyblock := enemypositions[rInt(0, len(enemypositions))]

	switch currentlevel {

	case 1:
		switch currentenemytype {
		case "warpmushroom":
			if framecount%120 == 0 {
				enemieslayout[enemyblock] = currentenemytype
				newenemy := enemy{}
				newenemy.hp = 2
				newenemy.img = mushroom
				newenemy.name = "mushroom"
				if flipcoin() {
					newenemy.movedir = 4
				} else {
					newenemy.movedir = 8
				}
				enemiesstructslayout[enemyblock] = newenemy
			}

		}
	}

}
func createdoors() { // MARK: createdoors

	for a := 0; a < 20; a++ {

		start := levelw * 100
		end := levela - (levelw * 100)
		block := 0
		block2 := 0
		for {
			choose := rInt(start, end)
			if levellayout[choose] == "." {
				block = choose
				break
			}
		}
		if levellayout[block+levelw] == "." {
			for {
				block += levelw
				if levellayout[block+levelw] == "_" {
					break
				}
			}
		}
		doorslayout[block] = "bigdoor"
		levellayout[(block+levelw)-3] = "#"
		levellayout[(block+levelw)-2] = "#"
		levellayout[(block+levelw)-1] = "#"
		levellayout[(block + levelw)] = "#"
		levellayout[(block+levelw)+1] = "#"
		levellayout[(block+levelw)+2] = "#"
		levellayout[(block+levelw)+3] = "#"
		leveltileslayout[(block+levelw)-3] = "#"
		leveltileslayout[(block+levelw)-2] = "#"
		leveltileslayout[(block+levelw)-1] = "#"
		leveltileslayout[(block + levelw)] = "#"
		leveltileslayout[(block+levelw)+1] = "#"
		leveltileslayout[(block+levelw)+2] = "#"
		leveltileslayout[(block+levelw)+3] = "#"

		for {
			choose := rInt(start, end)
			if levellayout[choose] == "." {
				block2 = choose
				break
			}
		}
		if levellayout[block2+levelw] == "." {
			for {
				block2 += levelw
				if levellayout[block2+levelw] == "_" {
					break
				}
			}
		}

		doorslayout[block2] = "bigdoor"

		doorpositionslayout[block] = block2
		doorpositionslayout[block2] = block

	}

}
func createshops() { // MARK: createshops

	shopnumber := rInt(4, 9)

	for {

		roomblock := 0
		for {
			roomblock = rInt(levelw*(levelh/2), levela-(levelw*50))

			if levellayout[roomblock] != "." {
				if levellayout[roomblock-13] != "." && levellayout[roomblock+13] != "." {
					if levellayout[roomblock+16] != "." {
						break
					}
				}
			}
		}

		roomblockholder := roomblock

		count := 0
		rooml := rInt(18, 26)
		roomh := rInt(11, 16)
		rooma := roomh * rooml

		// shop sign text
		textmarkerblock := roomblock + 4
		textmarkerblock += levelw * 7
		textmarkerslayout[textmarkerblock] = "shop"

		shoplevelblock := textmarkerblock

		// shop border blocks
		bordera := (rooml + 8) * (roomh * 2)
		borderblock := roomblock - 4

		for a := 0; a < bordera; a++ {
			levellayout[borderblock] = "#"
			leveltileslayout[borderblock] = "#"
			count++
			borderblock++
			if count == rooml+8 {
				count = 0
				borderblock += levelw
				borderblock -= rooml + 8
			}
		}

		// shop background
		count = 0
		shopbackgroundtile := "shopback1"

		if flipcoin() {
			shopbackgroundtile = "shopback2"
		}

		for a := 0; a < rooma; a++ {
			levellayout[roomblock] = "."
			leveltileslayout[roomblock] = shopbackgroundtile
			count++
			roomblock++
			if count == rooml {
				if flipcoin() {
					levellayout[roomblock] = "."
					leveltileslayout[roomblock] = shopbackgroundtile
					if flipcoin() {
						levellayout[roomblock+1] = "."
						leveltileslayout[roomblock+1] = shopbackgroundtile
					}
				}
				roomblock += levelw
				roomblock -= rooml
				count = 0
			}
		}

		roomblock = roomblockholder

		for a := 0; a < roomh; a++ {
			if flipcoin() {
				leveltileslayout[roomblock-1] = shopbackgroundtile
				if flipcoin() {
					leveltileslayout[roomblock-2] = shopbackgroundtile
				}
			}
			roomblock += levelw
		}

		// entry passage
		passageblock := roomblockholder
		passageblock += (rooml - 4)
		passageblock -= levelw
		passageblockholder := passageblock
		passagecount := 0

		for {
			levellayout[passageblock] = "."
			leveltileslayout[passageblock] = "."
			levellayout[passageblock-levelw] = "."
			leveltileslayout[passageblock-levelw] = "."
			levellayout[(passageblock-levelw)+1] = "."
			leveltileslayout[(passageblock-levelw)+1] = "."
			passageblock -= levelw * 2
			passageblock++
			passagecount++
			if levellayout[passageblock] == "." || passagecount > 50 {
				break
			}
		}

		passageblock = passageblockholder
		passageblock--
		passageblockholder = passageblock

		for {
			levellayout[passageblock] = "."
			leveltileslayout[passageblock] = "."
			levellayout[passageblock-levelw] = "."
			leveltileslayout[passageblock-levelw] = "."
			levellayout[(passageblock-levelw)+1] = "."
			leveltileslayout[(passageblock-levelw)+1] = "."

			textmarkerslayout[(passageblock-levelw)+1] = "shopmarker"

			passageblock -= levelw * 2
			passageblock++
			passagecount++
			if levellayout[passageblock] == "." || passagecount > 50 {
				break
			}
		}

		passageblock = passageblockholder
		passageblock--
		stairsblock := passageblock

		for {
			levellayout[passageblock] = "."
			leveltileslayout[passageblock] = "."
			levellayout[passageblock-levelw] = "."
			leveltileslayout[passageblock-levelw] = "."
			levellayout[(passageblock-levelw)+1] = "."
			leveltileslayout[(passageblock-levelw)+1] = "."

			passageblock -= levelw * 2
			passageblock++
			passagecount++
			if levellayout[passageblock] == "." || passagecount > 50 {
				break
			}
		}

		// stairs
		stairsblock += levelw * 4
		stairscount := 0
		for {
			for a := 0; a < 3; a++ {
				levellayout[stairsblock] = "#"
				leveltileslayout[stairsblock] = "^"
				stairsblock++
			}
			stairsblock += levelw * 3
			stairsblock -= 6
			stairscount++

			if levellayout[stairsblock] != "." || stairscount > 20 {
				break
			}
		}

		// shop items level
		shoplevelblock -= 3
		shoplevelblock += levelw
		itemsblock := shoplevelblock

		for a := 0; a < 14; a++ {
			levellayout[shoplevelblock] = "#"
			leveltileslayout[shoplevelblock] = "^"
			shoplevelblock++
		}

		// shop items
		itemsblock -= levelw
		itemsblock += 2
		createshopitems(itemsblock)
		itemsblock += 4
		createshopitems(itemsblock)
		itemsblock += 4
		createshopitems(itemsblock)

		// place shop banner
		bannerblock := roomblockholder
		bannerblock += rInt(-4, 5)
		bannercount := 0
		for {
			bannerblock -= levelw
			if levellayout[bannerblock] == "." {
				if levellayout[bannerblock+levelw] == "_" || levellayout[bannerblock+levelw] == "#" {

					decorationslayout[bannerblock] = "banner1"

					levellayout[bannerblock+levelw] = "#"
					levellayout[(bannerblock+levelw)-1] = "#"
					levellayout[(bannerblock+levelw)+1] = "#"

					leveltileslayout[bannerblock+levelw] = "#"
					leveltileslayout[(bannerblock+levelw)-1] = "#"
					leveltileslayout[(bannerblock+levelw)+1] = "#"

					break
				}
			}
			bannerblock += rInt(-4, 5)
			bannercount++
			if bannercount > 100 {
				break
			}
		}

		// place shopkeeper
		shopkeeperblock := roomblockholder
		shopkeeperblock += 5
		shopkeeperblock += (levelw * (roomh - 1))
		characterslayout[shopkeeperblock] = "shopkeeper"

		shopnumber--
		if shopnumber <= 0 {
			break
		}

	}
}
func createanimals() { // MARK: createanimals

	// add bats
	for a := 0; a < 200; a++ {

		blocknumber := 0
		for {
			choose := rInt(levela/2, levela)
			if levellayout[choose] == "." && levellayout[choose+levelw] == "." {
				blocknumber = choose
				break
			}
		}
		newanimal := animal{}
		newanimal.name = "bat"
		newanimal.active = true
		newanimal.movement = rInt(1, 9)
		newanimal.rotation = rFloat32(-45, 45)
		animalstructslayout[blocknumber] = newanimal
		animalslayout[blocknumber] = newanimal.name

	}

}
func createlevel() { // MARK: createlevel

	blocktype = rInt(1, 22)
	switch blocktype {
	case 1:
		tileblock = block1
	case 2:
		tileblock = block2
	case 3:
		tileblock = block3
	case 4:
		tileblock = block4
	case 5:
		tileblock = block5
	case 6:
		tileblock = block6
	case 7:
		tileblock = block7
	case 8:
		tileblock = block8
	case 9:
		tileblock = block9
	case 10:
		tileblock = block10
	case 11:
		tileblock = block11
	case 12:
		tileblock = block12
	case 13:
		tileblock = block13
	case 14:
		tileblock = block14
	case 15:
		tileblock = block15
	case 16:
		tileblock = block16
	case 17:
		tileblock = block17
	case 18:
		tileblock = block18
	case 19:
		tileblock = block19
	case 20:
		tileblock = block20
	case 21:
		tileblock = block21
	}

	plattype = rInt(1, 9)
	backtype = rInt(1, 11)
	// clear maps
	for a := 0; a < len(levellayout); a++ {
		levellayout[a] = "_"
		leveltileslayout[a] = "_"
		backlayout[a] = "."
		back2layout[a] = "."
		scenerylayout[a] = "."
	}
	// background detail items
	for a := 0; a < levelw; a++ {

		if rolldice()+rolldice()+rolldice() > 16 {
			block := a
			choose := rInt(1, 8)
			for b := 0; b < levelh; b++ {
				switch choose {
				case 1:
					back2layout[block] = "chain1"
					block += levelw
				case 2:
					back2layout[block] = "chain2"
					block += levelw
				case 3:
					back2layout[block] = "chain3"
					block += levelw
				case 4:
					back2layout[block] = "chain4"
					block += levelw
				case 5:
					back2layout[block] = "vine1"
					block += levelw
				case 6:
					back2layout[block] = "vine2"
					block += levelw
				case 7:
					back2layout[block] = "rope1"
					block += levelw
				}
			}
		}

	}
	// extra back
	extrabackground := "extraback1"
	choose := rInt(1, 8)
	switch choose {
	case 1:
		extrabackground = "extraback1"
	case 2:
		extrabackground = "extraback2"
	case 3:
		extrabackground = "extraback3"
	case 4:
		extrabackground = "extraback4"
	case 5:
		extrabackground = "extraback5"
	case 6:
		extrabackground = "extraback6"
	case 7:
		extrabackground = "extraback7"
	}
	for a := 30; a < levelw-30; a++ {
		if rolldice()+rolldice()+rolldice() > 16 {
			block := a

			for b := 0; b < levelh/2; b++ {
				number := rInt(3, 11)
				block += rInt(-10, 6)
				for c := 0; c < number; c++ {
					back3layout[block+(2*c)] = extrabackground
				}
				block += levelw * 2
			}
		}
	}
	// clear extra back duplicates
	for a := 0; a < len(back3layout); a++ {
		if back3layout[a] != "" {
			if back3layout[a] != "" && back3layout[a+1] != "" {
				back3layout[a+1] = ""
			}
			if a > levelw*2 {
				if back3layout[a] != "" && back3layout[(a+1)-levelw] != "" {
					back3layout[(a+1)-levelw] = ""
				}
			}
			if a < levela-(levelw*2) {
				if back3layout[a] != "" && back3layout[(a+1)+levelw] != "" {
					back3layout[(a+1)+levelw] = ""
				}
			}
		}
	}

	// main passage above
	for a := levelw * 100; a < levelw*150; a++ {
		levellayout[a] = "."
		leveltileslayout[a] = "."
	}
	createleveldetailabove()
	// secondary passages above
	passageblock := levelw * 180
	passageblock += 50

	passagedirection := rInt(1, 9)
	totallength := 0

	for {
		passagel := rInt(3, 25)
		passagew := rInt(1, 6)
		for a := 0; a < passagel; a++ {
			switch passagedirection {
			case 1:
				levellayout[passageblock] = "."
				leveltileslayout[passageblock] = "."
				createtunnel(passageblock, passagew, 2)
				passageblock -= levelw
				levellayout[passageblock] = "."
				leveltileslayout[passageblock] = "."
				createtunnel(passageblock, passagew, 5)
				passageblock--
				levellayout[passageblock] = "."
				leveltileslayout[passageblock] = "."
				createtunnel(passageblock, passagew, 1)
				end, endirection := checkblockhv(passageblock)
				if end {
					passagedirection = endirection
				}
			case 2:
				levellayout[passageblock] = "."
				leveltileslayout[passageblock] = "."
				createtunnel(passageblock, passagew, 2)
				passageblock -= levelw
				end, endirection := checkblockhv(passageblock)
				if end {
					passagedirection = endirection
				}
			case 3:
				levellayout[passageblock] = "."
				leveltileslayout[passageblock] = "."
				createtunnel(passageblock, passagew, 3)
				passageblock -= levelw
				levellayout[passageblock] = "."
				leveltileslayout[passageblock] = "."
				createtunnel(passageblock, passagew, 6)
				passageblock++
				levellayout[passageblock] = "."
				leveltileslayout[passageblock] = "."
				createtunnel(passageblock, passagew, 1)
				end, endirection := checkblockhv(passageblock)
				if end {
					passagedirection = endirection
				}
			case 4:
				levellayout[passageblock] = "."
				leveltileslayout[passageblock] = "."
				createtunnel(passageblock, passagew, 1)
				passageblock++
				end, endirection := checkblockhv(passageblock)
				if end {
					passagedirection = endirection
				}
			case 5:
				levellayout[passageblock] = "."
				leveltileslayout[passageblock] = "."
				createtunnel(passageblock, passagew, 3)
				passageblock += levelw
				levellayout[passageblock] = "."
				leveltileslayout[passageblock] = "."
				createtunnel(passageblock, passagew, 7)
				passageblock++
				levellayout[passageblock] = "."
				leveltileslayout[passageblock] = "."
				createtunnel(passageblock, passagew, 4)
				end, endirection := checkblockhv(passageblock)
				if end {
					passagedirection = endirection
				}
			case 6:
				levellayout[passageblock] = "."
				leveltileslayout[passageblock] = "."
				createtunnel(passageblock, passagew, 2)
				passageblock += levelw
				end, endirection := checkblockhv(passageblock)
				if end {
					passagedirection = endirection
				}
			case 7:
				levellayout[passageblock] = "."
				leveltileslayout[passageblock] = "."
				createtunnel(passageblock, passagew, 2)
				passageblock += levelw
				levellayout[passageblock] = "."
				leveltileslayout[passageblock] = "."
				createtunnel(passageblock, passagew, 8)
				passageblock--
				levellayout[passageblock] = "."
				leveltileslayout[passageblock] = "."
				createtunnel(passageblock, passagew, 4)
				end, endirection := checkblockhv(passageblock)
				if end {
					passagedirection = endirection
				}
			case 8:
				levellayout[passageblock] = "."
				leveltileslayout[passageblock] = "."
				createtunnel(passageblock, passagew, 1)
				passageblock--
				end, endirection := checkblockhv(passageblock)
				if end {
					passagedirection = endirection
				}
			}
		}
		end, endirection := checkblockhv(passageblock)
		if end {
			passagedirection = endirection
		} else {
			passagedirection = rInt(1, 9)
		}

		totallength += passagel
		if totallength > 10000 {
			break
		}
	}

	// main passage
	for a := levelw * ((levelh / 2) - 8); a < levela/2; a++ {
		levellayout[a] = "."
		leveltileslayout[a] = "."

		if a > levela/2-(levelw*7) {
			if rolldice()+rolldice()+rolldice() == 18 {
				createwater(a)
			}
		}
	}
	// roof detail
	for a := levelw * ((levelh / 2) - 8); a < levelw*((levelh/2)-7); a++ {
		roofblock := a
		number := rInt(0, 7)
		for b := 0; b < number; b++ {
			levellayout[roofblock] = "."
			leveltileslayout[roofblock] = "."
			roofblock -= levelw
		}
	}

	// vertical tunnels below
	startblock := (player + 20) + levelw
	startblockvert := startblock % levelw
	for a := startblockvert; a < levelw-32; a++ {
		startblock++
		if rolldice()+rolldice() > 11 {

			block := startblock
			tunnelw := rInt(6, 15)
			tunnelh := rInt(30, 116)
			count := 0
			platformcount := 0
			leftop := block
			interiortunnelplatforms := true
			righttop := block + tunnelw
			maxplatform := (tunnelw / 2) + 1

			newtunnel := tunnel{}

			newtunnel.height = tunnelh
			newtunnel.width = tunnelw
			newtunnel.startblock = block
			newtunnel.created = true

			vertbelowtunnellayout[vertbelowtunnelcount] = newtunnel
			vertbelowtunnelcount++

			for a := 0; a < (tunnelw * tunnelh); a++ {
				levellayout[block] = "."
				leveltileslayout[block] = "."
				block++
				count++
				if count == tunnelw {
					block += levelw
					block -= tunnelw
					if checkhorizvert(block, 3) { // check bottom border
						break
					}
					count = 0
					platformcount++
				}
			}
			for a := 0; a < tunnelh; a++ {
				blockchange := leftop
				number := rInt(0, 4)
				for b := 0; b < number; b++ {
					levellayout[blockchange] = "."
					leveltileslayout[blockchange] = "."
					blockchange--
				}
				blockchange = righttop
				number = rInt(0, 4)
				for b := 0; b < number; b++ {
					levellayout[blockchange] = "."
					leveltileslayout[blockchange] = "."
					blockchange++
				}

				// interior tunnel platforms
				if interiortunnelplatforms {
					if rolldice() == 6 {
						blockchange := leftop
						blockchange += rInt(1, tunnelw/2)
						platformlength := rInt(1, maxplatform)

						for b := 0; b < platformlength; b++ {
							levellayout[blockchange] = "_"
							leveltileslayout[blockchange] = "^"
							blockchange++
						}
					}
				}
				leftop += levelw
				righttop += levelw
			}
			startblock += (tunnelw * 2)
			a += (tunnelw * 2)
		}
	} // end vertical tunnels below

	// horizontal tunnels below
	for d := 0; d < len(vertbelowtunnellayout); d++ {
		checktunnel := vertbelowtunnellayout[d]
		if checktunnel.created {
			startblock = checktunnel.startblock
			startblock += checktunnel.width
			maxheight := 20
			maxnumber := checktunnel.height / maxheight
			for a := 0; a < maxnumber; a++ {
				if flipcoin() {
					updown := rInt(1, 4)
					tunnelh := rInt(8, 15)
					tunnelw := rInt(40, 150)
					startblock += levelw * ((a + 1) * maxheight)
					block := startblock + (tunnelh * levelw)
					chestblock := (block + 5) - levelw
					for b := 0; b < tunnelw; b++ {
						levellayout[block] = "."
						leveltileslayout[block] = "."
						hblock := block
						newtunnelh := tunnelh - rInt(0, 3)
						for c := 0; c < newtunnelh; c++ {
							levellayout[hblock] = "."
							leveltileslayout[hblock] = "."
							hblock -= levelw
						}
						if updown == 1 {
							if rolldice() == 6 {
								block -= levelw
								levellayout[block] = "."
								leveltileslayout[block] = "."
							}
						} else if updown == 3 {
							if rolldice() == 6 {
								if !checkhorizvert(block, 3) {
									block += levelw
									levellayout[block] = "."
									leveltileslayout[block] = "."
								}
							}
						}
						block++
					}

					deleteblock := chestblock
					deleteblock -= 4
					deleteblock -= levelw * 8
					count := 0
					for d := 0; d < 64; d++ {
						leveltileslayout[deleteblock] = "."
						levellayout[deleteblock] = "."
						deleteblock++
						count++
						if count == 8 {
							count = 0
							deleteblock += levelw
							deleteblock -= 8
						}
					}
					chestlevelblock := chestblock
					chestlevelblock += levelw
					chestlevelblock -= 4
					for d := 0; d < 8; d++ {
						levellayout[chestlevelblock] = "#"
						leveltileslayout[chestlevelblock] = "#"
						chestlevelblock++
					}
					lootlayout[chestblock] = "chest"
					chesttotal++
				}
			}
		}
	}
	// border ends
	block := 0
	block2 := levelw - 40
	holder1 := block
	holder2 := block2
	for a := 0; a < 40; a++ {
		for b := 0; b < levelh; b++ {
			levellayout[block] = "#"
			leveltileslayout[block] = "#"
			waterlayout[block] = ""
			levellayout[block2] = "#"
			leveltileslayout[block2] = "#"
			waterlayout[block2] = ""
			block += levelw
			block2 += levelw
		}
		block = holder1
		block2 = holder2
		block++
		block2++
		holder1 = block
		holder2 = block2
	}
	for a := 0; a < levelw*50; a++ {
		levellayout[a] = "#"
		leveltileslayout[a] = "#"
	}
	for a := levela - (levelw * 50); a < levela; a++ {
		levellayout[a] = "#"
		leveltileslayout[a] = "#"
	}

	// power up blocks
	for a := 0; a < 200; a++ {
		for {
			block := rInt(levelw*120, levelw*(levelh-50))
			if levellayout[block] == "." && levellayout[block+levelw] == "." && levellayout[block-levelw] == "." {
				levellayout[block] = "%"
				leveltileslayout[block] = "%"
				break
			}
		}
	}

	// map zoom blocks
	for a := 0; a < 50; a++ {
		for {
			block := rInt(levelw*120, levelw*(levelh-50))
			if levellayout[block] == "." && levellayout[block+levelw] == "_" {
				powerupslayout[block] = "map"
				break
			}
		}
	}
	// torches
	for a := levela/2 - (levelw * 4); a < levela/2-(levelw*3); a++ {
		if rolldice()+rolldice() == 12 {
			if levellayout[a] == "." {
				decorationslayout[a] = "torch"
				a += 6
			}
		}
	}
	// scenery
	for a := 0; a < 300; a++ {

		for {
			selected := false
			choose := rInt(levelw*20, levela-(levelw*5))

			if levellayout[choose] == "." && levellayout[choose+levelw] == "_" {
				choose2 := rInt(1, 36)

				switch choose2 {
				case 1:
					scenerylayout[choose] = "scenery1"
				case 2:
					scenerylayout[choose] = "scenery2"
				case 3:
					scenerylayout[choose] = "scenery3"
				case 4:
					scenerylayout[choose] = "scenery4"
				case 5:
					scenerylayout[choose] = "scenery5"
				case 6:
					scenerylayout[choose] = "scenery6"
				case 7:
					scenerylayout[choose] = "scenery7"
				case 8:
					scenerylayout[choose] = "scenery8"
				case 9:
					scenerylayout[choose] = "scenery9"
				case 10:
					scenerylayout[choose] = "scenery10"
				case 11:
					scenerylayout[choose] = "scenery11"
				case 12:
					scenerylayout[choose] = "scenery12"
				case 13:
					scenerylayout[choose] = "scenery13"
				case 14:
					scenerylayout[choose] = "scenery14"
				case 15:
					scenerylayout[choose] = "scenery15"
				case 16:
					scenerylayout[choose] = "scenery16"
				case 17:
					scenerylayout[choose] = "scenery17"
				case 18:
					scenerylayout[choose] = "scenery18"
				case 19:
					scenerylayout[choose] = "scenery19"
				case 20:
					scenerylayout[choose] = "scenery20"
				case 21:
					scenerylayout[choose] = "scenery21"
				case 22:
					scenerylayout[choose] = "scenery22"
				case 23:
					scenerylayout[choose] = "scenery23"
				case 24:
					scenerylayout[choose] = "scenery24"
				case 25:
					scenerylayout[choose] = "scenery25"
				case 26:
					scenerylayout[choose] = "scenery26"
				case 27:
					scenerylayout[choose] = "scenery27"
				case 28:
					scenerylayout[choose] = "scenery28"
				case 29:
					scenerylayout[choose] = "scenery29"
				case 30:
					scenerylayout[choose] = "scenery30"
				case 31:
					scenerylayout[choose] = "scenery31"
				case 32:
					scenerylayout[choose] = "scenery32"
				case 33:
					scenerylayout[choose] = "scenery33"
				case 34:
					scenerylayout[choose] = "scenery34"
				case 35:
					scenerylayout[choose] = "scenery35"
				}

				selected = true
			}
			if selected {
				break
			}

		}

	}

}
func createleveldetailabove() { // MARK: createleveldetailabove

	lightcolor1 = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
	lightcolor2 = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
	lightcolor3 = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
	lightcolor4 = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
	lightcolor5 = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)
	lightcolor6 = rl.NewColor(uint8(rInt(1, 255)), uint8(rInt(1, 255)), uint8(rInt(1, 255)), 255)

	block := levelw * 145
	block += 50
	blockholder := block

	for {
		choose := rInt(1, 7)
		length1 := rInt(10, 25)
		//	length2 := rInt(10, 25)
		length1holder := length1
		//	length2holder := length1
		//	area := length1 * length2
		switch choose {
		case 6:
			for {
				plat := rInt(4, 9)
				blockholder2 := block
				for a := 0; a < plat+1; a++ {
					levellayout[block] = "_"
					leveltileslayout[block] = "^"
					levellayout[block-(levelw*plat)] = "_"
					leveltileslayout[block-(levelw*plat)] = "^"
					block++

					if rolldice()+rolldice() == 12 {
						if decorationslayout[(block-1)-levelw] == "" {
							switch rolldice() {
							case 1:
								decorationslayout[block-levelw] = "lamp1"
								levellayout[block] = "#"
								leveltileslayout[block] = "^"
							case 2:
								decorationslayout[block-levelw] = "lamp2"
								levellayout[block] = "#"
								leveltileslayout[block] = "^"
							case 3:
								decorationslayout[block-levelw] = "lamp3"
								levellayout[block] = "#"
								leveltileslayout[block] = "^"
							case 4:
								decorationslayout[block-levelw] = "lamp4"
								levellayout[block] = "#"
								leveltileslayout[block] = "^"
							case 5:
								decorationslayout[block-levelw] = "lamp5"
								levellayout[block] = "#"
								leveltileslayout[block] = "^"
							case 6:
								decorationslayout[block-levelw] = "lamp6"
								levellayout[block] = "#"
								leveltileslayout[block] = "^"
							}
						}
					}
				}
				block = blockholder2
				for a := 0; a < plat; a++ {
					levellayout[block] = "_"
					leveltileslayout[block] = "^"
					levellayout[block+plat] = "_"
					leveltileslayout[block+plat] = "^"
					block -= levelw
				}
				block -= levelw * rInt(2, 6)
				block += rInt(-5, 6)
				if checkh(block - (levelw * 8)) {
					break
				}
			}
			block = blockholder
			block += 20
			blockholder = block
		case 5:
			for {
				plat := rInt(2, 5)
				if rolldice() < 5 {
					if flipcoin() {
						for a := 0; a < plat; a++ {
							levellayout[block] = "_"
							leveltileslayout[block] = "^"
							block++
							if rolldice()+rolldice() == 12 {
								if decorationslayout[(block-1)-levelw] == "" {
									switch rolldice() {
									case 1:
										decorationslayout[block-levelw] = "lamp1"
										levellayout[block] = "#"
										leveltileslayout[block] = "^"
									case 2:
										decorationslayout[block-levelw] = "lamp2"
										levellayout[block] = "#"
										leveltileslayout[block] = "^"
									case 3:
										decorationslayout[block-levelw] = "lamp3"
										levellayout[block] = "#"
										leveltileslayout[block] = "^"
									case 4:
										decorationslayout[block-levelw] = "lamp4"
										levellayout[block] = "#"
										leveltileslayout[block] = "^"
									case 5:
										decorationslayout[block-levelw] = "lamp5"
										levellayout[block] = "#"
										leveltileslayout[block] = "^"
									case 6:
										decorationslayout[block-levelw] = "lamp6"
										levellayout[block] = "#"
										leveltileslayout[block] = "^"
									}
								}
							}
							if checkh(block) {
								break
							}
						}
					} else {
						for a := 0; a < plat; a++ {
							levellayout[block] = "_"
							leveltileslayout[block] = "^"
							block--
							if rolldice()+rolldice() == 12 {
								if decorationslayout[(block+1)-levelw] == "" {
									switch rolldice() {
									case 1:
										decorationslayout[block-levelw] = "lamp1"
										levellayout[block] = "#"
										leveltileslayout[block] = "^"
									case 2:
										decorationslayout[block-levelw] = "lamp2"
										levellayout[block] = "#"
										leveltileslayout[block] = "^"
									case 3:
										decorationslayout[block-levelw] = "lamp3"
										levellayout[block] = "#"
										leveltileslayout[block] = "^"
									case 4:
										decorationslayout[block-levelw] = "lamp4"
										levellayout[block] = "#"
										leveltileslayout[block] = "^"
									case 5:
										decorationslayout[block-levelw] = "lamp5"
										levellayout[block] = "#"
										leveltileslayout[block] = "^"
									case 6:
										decorationslayout[block-levelw] = "lamp6"
										levellayout[block] = "#"
										leveltileslayout[block] = "^"
									}
								}
							}
						}
					}
				} else {
					for a := 0; a < plat; a++ {
						levellayout[block] = "_"
						leveltileslayout[block] = "^"
						block -= levelw
						if checkv(block) {
							break
						}
					}
				}
				if checkh(block) {
					break
				}
				if checkv(block) {
					break
				}
			}
			block = blockholder
			block += 15
			blockholder = block
		case 4:
			for {
				blockholder2 := block
				for a := 0; a < length1+1; a++ {
					levellayout[block] = "_"
					leveltileslayout[block] = "^"
					levellayout[block-(levelw*length1)] = "_"
					leveltileslayout[block-(levelw*length1)] = "^"
					block++
					if rolldice()+rolldice() == 12 {
						if decorationslayout[(block-1)-levelw] == "" {
							switch rolldice() {
							case 1:
								decorationslayout[block-levelw] = "lamp1"
								levellayout[block] = "#"
								leveltileslayout[block] = "^"
							case 2:
								decorationslayout[block-levelw] = "lamp2"
								levellayout[block] = "#"
								leveltileslayout[block] = "^"
							case 3:
								decorationslayout[block-levelw] = "lamp3"
								levellayout[block] = "#"
								leveltileslayout[block] = "^"
							case 4:
								decorationslayout[block-levelw] = "lamp4"
								levellayout[block] = "#"
								leveltileslayout[block] = "^"
							case 5:
								decorationslayout[block-levelw] = "lamp5"
								levellayout[block] = "#"
								leveltileslayout[block] = "^"
							case 6:
								decorationslayout[block-levelw] = "lamp6"
								levellayout[block] = "#"
								leveltileslayout[block] = "^"
							}
						}
					}
				}
				block = blockholder2
				for a := 0; a < length1; a++ {
					levellayout[block] = "_"
					leveltileslayout[block] = "^"
					levellayout[block+length1] = "_"
					leveltileslayout[block+length1] = "^"
					block -= levelw
					if checkh(block) {
						break
					}
				}
				block = blockholder2
				blockchange := rInt(-5, 6)

				if blockchange == -1 || blockchange == 0 || blockchange == 1 {
					if blockchange == -1 {
						blockchange = -2
					}
					if blockchange == 0 {
						if flipcoin() {
							blockchange = -2
						} else {
							blockchange = 2
						}
					}
					if blockchange == 1 {
						blockchange = 2
					}
				}

				block += blockchange
				block -= levelw * rInt(5, 15)
				blockholder2 = block
				if checkh(block - (levelw * length1)) {
					break
				}
			}
			block = blockholder
			block += length1 + 10
			blockholder = block
		case 3:
			block -= levelw * rInt(0, 6)
			blockholder2 := block

			for {
				plat := rInt(3, 7)
				for a := 0; a < plat; a++ {
					levellayout[block] = "_"
					leveltileslayout[block] = "^"
					block++
					block -= levelw
				}
				block--
				block += levelw
				if rolldice()+rolldice() == 12 {
					if decorationslayout[(block-1)-levelw] == "" {
						switch rolldice() {
						case 1:
							decorationslayout[block-levelw] = "lamp1"
							levellayout[block] = "#"
							leveltileslayout[block] = "^"
						case 2:
							decorationslayout[block-levelw] = "lamp2"
							levellayout[block] = "#"
							leveltileslayout[block] = "^"
						case 3:
							decorationslayout[block-levelw] = "lamp3"
							levellayout[block] = "#"
							leveltileslayout[block] = "^"
						case 4:
							decorationslayout[block-levelw] = "lamp4"
							levellayout[block] = "#"
							leveltileslayout[block] = "^"
						case 5:
							decorationslayout[block-levelw] = "lamp5"
							levellayout[block] = "#"
							leveltileslayout[block] = "^"
						case 6:
							decorationslayout[block-levelw] = "lamp6"
							levellayout[block] = "#"
							leveltileslayout[block] = "^"
						}
					}
				}
				for a := 0; a < plat; a++ {
					levellayout[block] = "_"
					leveltileslayout[block] = "^"
					block++
					block += levelw
				}

				block = blockholder2
				block -= levelw * (plat + 1)
				block += rInt(-3, 4)
				block -= levelw * rInt(0, 4)
				blockholder2 = block
				if checkh(block) {
					break
				}
			}
			block = blockholder
			block += 20
			blockholder = block
		case 2:
			block -= levelw * rInt(0, 6)
			for {
				plat := rInt(3, 9)
				blockholder2 := block
				for a := 0; a < plat; a++ {
					levellayout[block] = "_"
					leveltileslayout[block] = "^"
					block++
					if rolldice()+rolldice() == 12 {
						if decorationslayout[(block-1)-levelw] == "" {
							switch rolldice() {
							case 1:
								decorationslayout[block-levelw] = "lamp1"
								levellayout[block] = "#"
								leveltileslayout[block] = "^"
							case 2:
								decorationslayout[block-levelw] = "lamp2"
								levellayout[block] = "#"
								leveltileslayout[block] = "^"
							case 3:
								decorationslayout[block-levelw] = "lamp3"
								levellayout[block] = "#"
								leveltileslayout[block] = "^"
							case 4:
								decorationslayout[block-levelw] = "lamp4"
								levellayout[block] = "#"
								leveltileslayout[block] = "^"
							case 5:
								decorationslayout[block-levelw] = "lamp5"
								levellayout[block] = "#"
								leveltileslayout[block] = "^"
							case 6:
								decorationslayout[block-levelw] = "lamp6"
								levellayout[block] = "#"
								leveltileslayout[block] = "^"
							}
						}
					}
				}
				block = blockholder2
				block += rInt(-3, 4)
				block -= levelw * rInt(3, 9)
				if checkh(block) {
					break
				}
			}
			block = blockholder
			block += 20
			blockholder = block
		case 1:
			for {
				for a := 0; a < length1holder; a++ {
					levellayout[block] = "_"
					leveltileslayout[block] = "^"
					block++
					if checkv(block) {
						break
					}
				}
				length1holder -= 2
				for a := 0; a < length1holder; a++ {
					levellayout[block] = "_"
					leveltileslayout[block] = "^"
					block -= levelw
					if checkh(block) {
						break
					}
				}
				length1holder -= 2
				for a := 0; a < length1holder; a++ {
					levellayout[block] = "_"
					leveltileslayout[block] = "^"
					block--
				}
				length1holder -= 2
				for a := 0; a < length1holder; a++ {
					levellayout[block] = "_"
					leveltileslayout[block] = "^"
					block += levelw
					if checkh(block) {
						break
					}
				}
				length1holder -= 2

				if length1holder <= 2 {
					break
				}
			}
			block = blockholder
			block += length1 + 15
			blockholder = block
			length1holder = length1
		}
		if checkv(block) {
			break
		}
	}

	number := rInt(50, 100)
	for {
		block = rInt(levelw*105, levelw*145)
		plat := rInt(3, 9)
		if leveltileslayout[block] == "." {
			for a := 0; a < plat; a++ {
				if a == 0 {
					levellayout[block] = "#"
					leveltileslayout[block] = "platl"
					levellayout[block-levelw] = "."
					leveltileslayout[block-levelw] = "."
					levellayout[block-(levelw*2)] = "."
					leveltileslayout[block-(levelw*2)] = "."
					levellayout[block+levelw] = "."
					leveltileslayout[block+levelw] = "."
					// clear end space
					levellayout[block-1] = "."
					leveltileslayout[block-1] = "."
					levellayout[(block-1)-levelw] = "."
					leveltileslayout[(block-1)-levelw] = "."
					levellayout[(block-1)-(levelw*2)] = "."
					leveltileslayout[(block-1)-(levelw*2)] = "."
					levellayout[(block-1)+levelw] = "."
					leveltileslayout[(block-1)+levelw] = "."
				} else if a == plat-1 {
					levellayout[block] = "#"
					leveltileslayout[block] = "platr"
					leveltileslayout[block-levelw] = "."
					levellayout[block-(levelw*2)] = "."
					leveltileslayout[block-(levelw*2)] = "."
					levellayout[block+levelw] = "."
					leveltileslayout[block+levelw] = "."
					// clear end space
					levellayout[block+1] = "."
					leveltileslayout[block+1] = "."
					levellayout[(block+1)-levelw] = "."
					leveltileslayout[(block+1)-levelw] = "."
					levellayout[(block+1)-(levelw*2)] = "."
					leveltileslayout[(block+1)-(levelw*2)] = "."
					levellayout[(block+1)+levelw] = "."
					leveltileslayout[(block+1)+levelw] = "."
				} else {
					levellayout[block] = "#"
					leveltileslayout[block] = "platm"
					leveltileslayout[block-levelw] = "."
					levellayout[block-(levelw*2)] = "."
					leveltileslayout[block-(levelw*2)] = "."
					levellayout[block+levelw] = "."
					leveltileslayout[block+levelw] = "."
				}
				block++
			}
		}

		number--
		if number <= 0 {
			break
		}
	}
	if flipcoin() {
		extraslayout[(levelw*120)+60] = "ufo1"
	} else {
		extraslayout[(levelw*120)+65] = "ufo2"
	}

}
func createtunnel(block, width, direction int) { // MARK: createtunnel

	switch direction {

	case 1:
		for a := 0; a < width; a++ {
			levellayout[block] = "."
			leveltileslayout[block] = "."
			block -= levelw
		}
	case 2:
		for a := 0; a < width; a++ {
			levellayout[block] = "."
			leveltileslayout[block] = "."
			block++
		}
	case 3:
		for a := 0; a < width; a++ {
			levellayout[block] = "."
			leveltileslayout[block] = "."
			block--
		}
	case 4:
		for a := 0; a < width; a++ {
			levellayout[block] = "."
			leveltileslayout[block] = "."
			block += levelw
		}

	case 5:
		area := width * width
		count := 0
		for a := 0; a < area; a++ {
			levellayout[block] = "."
			leveltileslayout[block] = "."
			block++
			count++
			if count == width {
				count = 0
				block -= levelw
				block -= width
			}
		}
	case 6:
		area := width * width
		count := 0
		for a := 0; a < area; a++ {
			levellayout[block] = "."
			leveltileslayout[block] = "."
			block--
			count++
			if count == width {
				count = 0
				block -= levelw
				block += width
			}
		}
	case 7:
		area := width * width
		count := 0
		for a := 0; a < area; a++ {
			levellayout[block] = "."
			leveltileslayout[block] = "."
			block++
			count++
			if count == width {
				count = 0
				block += levelw
				block -= width
			}
		}
	case 8:
		area := width * width
		count := 0
		for a := 0; a < area; a++ {
			levellayout[block] = "."
			leveltileslayout[block] = "."
			block--
			count++
			if count == width {
				count = 0
				block += levelw
				block += width
			}
		}
	}

}
func createwater(block int) { // MARK: createwater

	waterw := rInt(3, 15)
	waterh := rInt(3, 8)
	watera := waterw * waterh

	block -= waterw
	block += levelw

	count := 0
	holder := block
	for a := 0; a < watera; a++ {
		leveltileslayout[block] = "."
		levellayout[block] = "."
		waterlayout[block] = "water"
		block++
		count++
		if count == waterw {
			count = 0
			if flipcoin() {
				number := rInt(1, 5)
				for b := 0; b < number; b++ {
					leveltileslayout[block+b] = "."
					levellayout[block+b] = "."
					waterlayout[block+b] = "water"
				}
			}
			block = holder
			if flipcoin() {
				number := rInt(1, 5)
				for b := 0; b < number; b++ {
					leveltileslayout[block-b] = "."
					levellayout[block-b] = "."
					waterlayout[block-b] = "water"
				}
			}
			block = holder
			block += levelw
			holder = block

		}
	}

}
func checkv(block int) bool {
	end := false
	blockv := block % levelw
	if blockv > levelw-70 {
		end = true
	}
	return end
}
func checkh(block int) bool {
	end := false
	blockh := block / levelw
	if blockh < 105 {
		end = true
	}
	return end

}
func checkblockhv(block int) (bool, int) {
	end, endirection := false, 0
	blockv, blockh := block%levelw, block/levelw

	if blockv < 40 {
		end = true
		endirection = 4
	} else if blockv > levelw-40 {
		end = true
		endirection = 8
	}

	if blockh < 150 {
		end = true
		endirection = 6
	} else if blockh > levelh/2 {
		end = true
		endirection = 2
	}

	return end, endirection

}
func input() { // MARK: input

	if rl.IsKeyPressed(rl.KeyKpMultiply) {
		if centerlineson {
			centerlineson = false
		} else {
			centerlineson = true
		}
	}

	if rl.IsKeyPressed(rl.KeyEscape) {
		if optionson {
			optionson = false
			pauseon = false
		} else {
			optionson = true
			pauseon = true
			optionselecty = 170
		}
	}

	if rl.IsKeyReleased(rl.KeySpace) {
		swingrighton = false
		swinglefton = false
		uzziangle = 0
	}

	if rl.IsKeyPressed(rl.KeySpace) {

		if lootlayout[player] == "chest" || lootlayout[player+1] == "chest" || lootlayout[player-1] == "chest" {

			if lootlayout[player] == "chest" {
				openchest(player)
				lootlayout[player] = "chesthalf"
			}
			if lootlayout[player+1] == "chest" {
				openchest(player + 1)
				lootlayout[player+1] = "chesthalf"
			}
			if lootlayout[player-1] == "chest" {
				openchest(player - 1)
				lootlayout[player-1] = "chesthalf"
			}

		}

		if doorslayout[player] != "" || doorslayout[player+1] != "" || doorslayout[player-1] != "" {
			doorjump()
		} else {
			if playerdirection == 2 {
				swingrighton = true
				if enemieslayout[player] != "" {
					checkenemy := enemiesstructslayout[player]
					checkenemy.hp -= playerattack
					if checkenemy.hp <= 0 {
						killenemy(player)
					} else {
						enemiesstructslayout[player] = checkenemy
					}
				}
				if enemieslayout[player+1] != "" {
					if enemieslayout[player+1] != "" {
						checkenemy := enemiesstructslayout[player+1]
						checkenemy.hp -= playerattack
						if checkenemy.hp <= 0 {
							killenemy(player + 1)
						} else {
							enemiesstructslayout[player+1] = checkenemy
						}
					}
				}
				if enemieslayout[player+2] != "" {
					if enemieslayout[player+2] != "" {
						checkenemy := enemiesstructslayout[player+2]
						checkenemy.hp -= playerattack
						if checkenemy.hp <= 0 {
							killenemy(player + 2)
						} else {
							enemiesstructslayout[player+2] = checkenemy
						}
					}
				}
				if enemieslayout[player+3] != "" {
					if enemieslayout[player+3] != "" {
						checkenemy := enemiesstructslayout[player+2]
						checkenemy.hp -= playerattack
						if checkenemy.hp <= 0 {
							killenemy(player + 3)
						} else {
							enemiesstructslayout[player+3] = checkenemy
						}
					}
				}
			} else if playerdirection == 4 {
				swinglefton = true
				if enemieslayout[player] != "" {
					checkenemy := enemiesstructslayout[player]
					checkenemy.hp -= playerattack
					if checkenemy.hp <= 0 {
						killenemy(player)
					} else {
						enemiesstructslayout[player] = checkenemy
					}
				}
				if enemieslayout[player-1] != "" {
					if enemieslayout[player-1] != "" {
						checkenemy := enemiesstructslayout[player-1]
						checkenemy.hp -= playerattack
						if checkenemy.hp <= 0 {
							killenemy(player - 1)
						} else {
							enemiesstructslayout[player-1] = checkenemy
						}
					}
				}
				if enemieslayout[player-2] != "" {
					if enemieslayout[player-2] != "" {
						checkenemy := enemiesstructslayout[player-2]
						checkenemy.hp -= playerattack
						if checkenemy.hp <= 0 {
							killenemy(player - 2)
						} else {
							enemiesstructslayout[player-2] = checkenemy
						}
					}
				}
				if enemieslayout[player-3] != "" {
					if enemieslayout[player-3] != "" {
						checkenemy := enemiesstructslayout[player-2]
						checkenemy.hp -= playerattack
						if checkenemy.hp <= 0 {
							killenemy(player - 3)
						} else {
							enemiesstructslayout[player-3] = checkenemy
						}
					}
				}
			}
		}

		if pickaxeactive {
			emoteon = false
			idletimer = 0
			if playerdirection == 2 {
				if levellayout[player+1] == "_" {
					levellayout[player+1] = "."
					leveltileslayout[player+1] = "."
				}
			}
			if playerdirection == 4 {
				if levellayout[player-1] == "_" {
					levellayout[player-1] = "."
					leveltileslayout[player-1] = "."
				}
			}
		} else if spadeactive {
			emoteon = false
			idletimer = 0
			if levellayout[player+levelw] == "_" {
				levellayout[player+levelw] = "."
				leveltileslayout[player+levelw] = "."
			}
		} else if icewandactive {
			if !icefreezeactive {
				for a := 0; a < len(icefreezev2); a++ {
					x := rFloat32(3, (monw / 4))
					y := rFloat32(3, (monh / 4))
					icefreezev2[a] = rl.NewVector2(x, y)
				}
				icefreezeactive = true
				icefreezetimer = 5
				icefreezeframe = 0
			}
		} else if uzziactive {

			if playerdirection == 2 {

				uzziangle = rFloat32(-25, 15)
				number := rInt(4, 8)

				switch number {
				case 4:
					if flipcoin() {
						weaponslayout[(player+2)-(levelw)] = "uzzibulletup"
					} else {
						weaponslayout[(player+2)-(levelw)] = "uzzibulletrup"
					}
					if flipcoin() {
						weaponslayout[(player)-(levelw*2)] = "uzzibulletrup"
					} else {
						weaponslayout[(player)-(levelw*2)] = "uzzibulletr"
					}
					if flipcoin() {
						weaponslayout[(player + 1)] = "uzzibulletrup"
					} else {
						weaponslayout[(player + 1)] = "uzzibulletup"
					}
					if flipcoin() {
						weaponslayout[(player + 2)] = "uzzibulletr"
					} else {
						weaponslayout[(player + 2)] = "uzzibulletrup"
					}
				case 5:
					if flipcoin() {
						weaponslayout[(player+2)-(levelw)] = "uzzibulletup"
					} else {
						weaponslayout[(player+2)-(levelw)] = "uzzibulletrup"
					}
					if flipcoin() {
						weaponslayout[(player)-(levelw*2)] = "uzzibulletrup"
					} else {
						weaponslayout[(player)-(levelw*2)] = "uzzibulletr"
					}
					if flipcoin() {
						weaponslayout[(player + 1)] = "uzzibulletrup"
					} else {
						weaponslayout[(player + 1)] = "uzzibulletup"
					}
					if flipcoin() {
						weaponslayout[(player + 2)] = "uzzibulletr"
					} else {
						weaponslayout[(player + 2)] = "uzzibulletrup"
					}
					if flipcoin() {
						weaponslayout[(player+1)-(levelw)] = "uzzibulletr"
					} else {
						weaponslayout[(player+1)-(levelw)] = "uzzibulletrup"
					}
				case 6:
					if flipcoin() {
						weaponslayout[(player+2)-(levelw)] = "uzzibulletup"
					} else {
						weaponslayout[(player+2)-(levelw)] = "uzzibulletrup"
					}
					if flipcoin() {
						weaponslayout[(player)-(levelw*2)] = "uzzibulletrup"
					} else {
						weaponslayout[(player)-(levelw*2)] = "uzzibulletr"
					}
					if flipcoin() {
						weaponslayout[(player + 1)] = "uzzibulletrup"
					} else {
						weaponslayout[(player + 1)] = "uzzibulletup"
					}
					if flipcoin() {
						weaponslayout[(player + 2)] = "uzzibulletr"
					} else {
						weaponslayout[(player + 2)] = "uzzibulletrup"
					}
					if flipcoin() {
						weaponslayout[(player+1)-(levelw)] = "uzzibulletr"
					} else {
						weaponslayout[(player+1)-(levelw)] = "uzzibulletrup"
					}
					if flipcoin() {
						weaponslayout[(player+4)-(levelw*2)] = "uzzibulletup"
					} else {
						weaponslayout[(player+4)-(levelw*2)] = "uzzibulletr"
					}
				case 7:
					if flipcoin() {
						weaponslayout[(player+2)-(levelw)] = "uzzibulletup"
					} else {
						weaponslayout[(player+2)-(levelw)] = "uzzibulletrup"
					}
					if flipcoin() {
						weaponslayout[(player)-(levelw*2)] = "uzzibulletrup"
					} else {
						weaponslayout[(player)-(levelw*2)] = "uzzibulletr"
					}
					if flipcoin() {
						weaponslayout[(player + 1)] = "uzzibulletrup"
					} else {
						weaponslayout[(player + 1)] = "uzzibulletup"
					}
					if flipcoin() {
						weaponslayout[(player + 2)] = "uzzibulletr"
					} else {
						weaponslayout[(player + 2)] = "uzzibulletrup"
					}
					if flipcoin() {
						weaponslayout[(player+1)-(levelw)] = "uzzibulletr"
					} else {
						weaponslayout[(player+1)-(levelw)] = "uzzibulletrup"
					}
					if flipcoin() {
						weaponslayout[(player+4)-(levelw*2)] = "uzzibulletup"
					} else {
						weaponslayout[(player+4)-(levelw*2)] = "uzzibulletr"
					}
					if flipcoin() {
						weaponslayout[(player+3)-(levelw*2)] = "uzzibulletup"
					} else {
						weaponslayout[(player+3)-(levelw*2)] = "uzzibulletrup"
					}
				}
			} else if playerdirection == 4 {

				uzziangle = rFloat32(-15, 25)
				number := rInt(4, 8)

				switch number {
				case 4:
					if flipcoin() {
						weaponslayout[(player-2)-(levelw)] = "uzzibulletup"
					} else {
						weaponslayout[(player-2)-(levelw)] = "uzzibulletlup"
					}
					if flipcoin() {
						weaponslayout[(player)-(levelw*2)] = "uzzibulletlup"
					} else {
						weaponslayout[(player)-(levelw*2)] = "uzzibulletl"
					}
					if flipcoin() {
						weaponslayout[(player - 1)] = "uzzibulletlup"
					} else {
						weaponslayout[(player - 1)] = "uzzibulletup"
					}
					if flipcoin() {
						weaponslayout[(player - 2)] = "uzzibulletl"
					} else {
						weaponslayout[(player - 2)] = "uzzibulletlup"
					}
				case 5:
					if flipcoin() {
						weaponslayout[(player-2)-(levelw)] = "uzzibulletup"
					} else {
						weaponslayout[(player-2)-(levelw)] = "uzzibulletlup"
					}
					if flipcoin() {
						weaponslayout[(player)-(levelw*2)] = "uzzibulletlup"
					} else {
						weaponslayout[(player)-(levelw*2)] = "uzzibulletl"
					}
					if flipcoin() {
						weaponslayout[(player - 1)] = "uzzibulletlup"
					} else {
						weaponslayout[(player - 1)] = "uzzibulletup"
					}
					if flipcoin() {
						weaponslayout[(player - 2)] = "uzzibulletl"
					} else {
						weaponslayout[(player - 2)] = "uzzibulletlup"
					}
					if flipcoin() {
						weaponslayout[(player-1)-(levelw)] = "uzzibulletl"
					} else {
						weaponslayout[(player-1)-(levelw)] = "uzzibulletlup"
					}
				case 6:
					if flipcoin() {
						weaponslayout[(player-2)-(levelw)] = "uzzibulletup"
					} else {
						weaponslayout[(player-2)-(levelw)] = "uzzibulletlup"
					}
					if flipcoin() {
						weaponslayout[(player)-(levelw*2)] = "uzzibulletlup"
					} else {
						weaponslayout[(player)-(levelw*2)] = "uzzibulletl"
					}
					if flipcoin() {
						weaponslayout[(player - 1)] = "uzzibulletlup"
					} else {
						weaponslayout[(player - 1)] = "uzzibulletup"
					}
					if flipcoin() {
						weaponslayout[(player - 2)] = "uzzibulletl"
					} else {
						weaponslayout[(player - 2)] = "uzzibulletlup"
					}
					if flipcoin() {
						weaponslayout[(player-1)-(levelw)] = "uzzibulletl"
					} else {
						weaponslayout[(player-1)-(levelw)] = "uzzibulletlup"
					}
					if flipcoin() {
						weaponslayout[(player-4)-(levelw*2)] = "uzzibulletup"
					} else {
						weaponslayout[(player-4)-(levelw*2)] = "uzzibulletl"
					}
				case 7:
					if flipcoin() {
						weaponslayout[(player-2)-(levelw)] = "uzzibulletup"
					} else {
						weaponslayout[(player-2)-(levelw)] = "uzzibulletlup"
					}
					if flipcoin() {
						weaponslayout[(player)-(levelw*2)] = "uzzibulletlup"
					} else {
						weaponslayout[(player)-(levelw*2)] = "uzzibulletl"
					}
					if flipcoin() {
						weaponslayout[(player - 1)] = "uzzibulletlup"
					} else {
						weaponslayout[(player - 1)] = "uzzibulletup"
					}
					if flipcoin() {
						weaponslayout[(player - 2)] = "uzzibulletl"
					} else {
						weaponslayout[(player - 2)] = "uzzibulletlup"
					}
					if flipcoin() {
						weaponslayout[(player-1)-(levelw)] = "uzzibulletl"
					} else {
						weaponslayout[(player-1)-(levelw)] = "uzzibulletlup"
					}
					if flipcoin() {
						weaponslayout[(player-4)-(levelw*2)] = "uzzibulletup"
					} else {
						weaponslayout[(player-4)-(levelw*2)] = "uzzibulletl"
					}
					if flipcoin() {
						weaponslayout[(player-3)-(levelw*2)] = "uzzibulletup"
					} else {
						weaponslayout[(player-3)-(levelw*2)] = "uzzibulletlup"
					}
				}
			}

		} else if pumpactionactive {
			if playerdirection == 2 {
				if levellayout[player+1] == "." {
					weaponslayout[player+1] = "shotgunbulletright"
				}
				if levellayout[(player+1)-levelw] == "." {
					weaponslayout[(player+1)-levelw] = "shotgunbulletupright"
				}
				if levellayout[(player+1)+levelw] == "." {
					weaponslayout[(player+1)+levelw] = "shotgunbulletdownright"
				}
				if levellayout[(player+3)-(levelw)] == "." {
					weaponslayout[(player+3)-(levelw)] = "shotgunbulletupright"
				}
				if levellayout[(player+3)+(levelw)] == "." {
					weaponslayout[(player+3)+(levelw)] = "shotgunbulletdownright"
				}
			}
			if playerdirection == 4 {
				if levellayout[player-1] == "." {
					weaponslayout[player-1] = "shotgunbulletleft"
				}
				if levellayout[(player-1)-levelw] == "." {
					weaponslayout[(player-1)-levelw] = "shotgunbulletupleft"
				}
				if levellayout[(player-1)+levelw] == "." {
					weaponslayout[(player-1)+levelw] = "shotgunbulletdownleft"
				}
				if levellayout[(player-3)-(levelw)] == "." {
					weaponslayout[(player-3)-(levelw)] = "shotgunbulletupleft"
				}
				if levellayout[(player-3)+(levelw)] == "." {
					weaponslayout[(player-3)+(levelw)] = "shotgunbulletdownleft"
				}
			}
		} else if bazookaactive {
			if !playerrollactive {
				playerrollactive = true
			}
			if playerdirection == 2 {
				if levellayout[player+2] == "." {
					weaponslayout[player+2] = "rocketr"
				}
				if levellayout[player-1] == "." {
					player--
				}
				if levellayout[player-1] == "." {
					player--
				}
				if levellayout[player-1] == "." {
					player--
				}
			} else if playerdirection == 4 {
				if levellayout[player-2] == "." {
					weaponslayout[player-2] = "rocketl"
				}
				if levellayout[player+1] == "." {
					player++
				}
				if levellayout[player+1] == "." {
					player++
				}
				if levellayout[player+1] == "." {
					player++
				}
			}
		} else {
			if boxinggloveactive {
				if levellayout[player-(levelw*4)] == "." {
					if flipcoin() {
						weaponslayout[player-(levelw*4)] = "boxingglove"
					} else {
						weaponslayout[player-(levelw*4)] = "boxingglover"
					}
				} else if levellayout[player-(levelw*3)] == "." {
					if flipcoin() {
						weaponslayout[player-(levelw*3)] = "boxingglove"
					} else {
						weaponslayout[player-(levelw*3)] = "boxingglover"
					}

				} else if levellayout[player-(levelw*2)] == "." {
					if flipcoin() {
						weaponslayout[player-(levelw*2)] = "boxingglove"
					} else {
						weaponslayout[player-(levelw*2)] = "boxingglover"
					}
				}
			}
			if grenadactive {
				if rolldice() > 4 {
					creategrenade()
				}
			}
			if bulletactive {
				if levellayout[player-levelw] == "." {
					weaponslayout[player-levelw] = "bullet"
				}
			}
			if ninjastaractive {
				createninjastar()
			}
			if fireballactive {
				createfireball()
			}
			if arrowactive {
				createarrow()
			}
		}
	}

	if rl.IsKeyPressed(rl.KeyLeftAlt) {
		bomb()
	}
	if rl.IsKeyPressed(rl.KeyLeftControl) {
		if flosson {
			flosson = false
		} else {
			floss(2)
			flosson = true
		}
	}
	if rl.IsKeyPressed(rl.KeyKp4) {
		if drawselectscreenon {
			drawselectscreenon = false
		} else {
			drawselectscreenon = true
		}
	}
	if rl.IsKeyPressed(rl.KeyKp3) {
		if drawstarton {
			drawstarton = false
			pauseon = false
		} else {
			drawstarton = true
			intropinkbackon = false
			introtimer4on, rayliblogoon, introtimer1on, introtimer2on, introtimer3on = false, false, false, false, false
			gologov2 = rl.NewVector2(-300, float32((monh/2)-262))
			rayliblogov2 = rl.NewVector2(float32((monw/2)-128), float32(monh+10))
			intropinkbackv2 = rl.NewVector2(0, float32(0-monh))
			introtimer1frame, introtimer3frame, introtimer4frame, introtimer5frame = 0, 0, 0, 0
			introtimer3, introtimer1, introtimer4 = 2, 2, 2
			startlogofadeon, startlogofadeon2 = true, false
			startlogofade = float32(0.0)
			pauseon = true
			introcircleson = false
			intronicholasimontexton2 = false
			nicholasimontextfade = 0.0
			for a := 0; a < len(introcirlcesv2layout); a++ {
				x := rFloat32(0, monw)
				y := rFloat32(0, monh)

				xy := rl.NewVector2(x, y)

				newcircle := circlev2{}

				newcircle.center = xy
				newcircle.radius = 0
				newcircle.color = rl.Black
				introcirlcesv2layout[a] = newcircle
			}
		}

	}
	if rl.IsKeyPressed(rl.KeyKp2) {
		if discoveryon {
			discoveryon = false
		} else {
			discoveryon = true
		}
	}
	if rl.IsKeyPressed(rl.KeyKp1) {

		if introscreenon {
			introscreenon = false
			pauseon = false
		} else {
			introscreenon = true
			optionson = false
		}
	}

	if rl.IsKeyPressed(rl.KeyPause) {
		if pauseon {
			pauseon = false
		} else {
			pauseon = true
		}
	}
	if rl.IsKeyPressed(rl.KeyTab) {
		if !inventoryon {
			inventoryon = true
			inventorypause = true
			inventorycount = 0
		}
	}
	if rl.IsKeyPressed(rl.KeyF7) {
		screenshakeactive = true
	}
	if rl.IsKeyPressed(rl.KeyF6) {
		createquest()
		questactive = true
	}
	if rl.IsKeyPressed(rl.KeyF5) {
		if discoveryinfoon {
			discoveryinfoon = false
		} else {
			discoveryinfoon = true
		}
	}
	if rl.IsKeyPressed(rl.KeyF4) {
		if moveoff {
			moveoff = false
		} else {
			moveoff = true
		}
	}
	if rl.IsKeyPressed(rl.KeyF3) {
		if gravityoff {
			gravityoff = false
		} else {
			gravityoff = true
		}
	}
	if rl.IsKeyPressed(rl.KeyF2) {
		player = playerstart
	}
	if rl.IsKeyPressed(rl.KeyF1) {
		if gridon {
			gridon = false
		} else {
			gridon = true
		}
	}

	if !pauseon && !inventoryon && !introscreenon && !drawstarton && !drawselectscreenon {
		if rl.IsKeyDown(rl.KeyRight) {
			if moveoff {
				drawblocknext += 2
			} else {
				moveplayer(2)
				playerdirection = 2
				playerrunon = true
			}
		} else if rl.IsKeyDown(rl.KeyLeft) {
			if moveoff {
				drawblocknext -= 2
			} else {
				moveplayer(4)
				playerdirection = 4
				playerrunon = true
			}
		} else {
			playerrunon = false
		}
		if rl.IsKeyDown(rl.KeyUp) {
			if moveoff {
				drawblocknext -= (levelw * 2)
			} else if gravityoff {
				if levellayout[player-levelw] == "." {
					player -= levelw
				}
			} else if levellayout[player-levelw] == "_" && levellayout[player-(levelw*2)] == "." {
				player -= levelw * 2
			} else if levellayout[player-levelw] == "#" && levellayout[player-(levelw*2)] == "." {
				player -= levelw * 2
			} else {
				if jumpactive == false && fallactive == false {
					jumpactive = true
					playerrollactive = true
					jumppause = 2
				}
			}
		}
		if rl.IsKeyDown(rl.KeyDown) {
			if moveoff {
				drawblocknext += (levelw * 2)
			} else if gravityoff {
				if levellayout[player+levelw] == "." {
					player += levelw
				}
			} else {
				if levellayout[player+levelw] == "_" && levellayout[player+(levelw*2)] == "." {
					player += levelw * 2
				}
			}
		}
	}

	if rl.IsKeyPressed(rl.KeyKpDecimal) {
		if debugon {
			debugon = false
		} else {
			debugon = true
		}
	}

	if rl.IsKeyPressed(rl.KeyKpAdd) {
		if camera.Zoom == 1.0 {
			camera.Zoom = 2.0
		} else if camera.Zoom == 2.0 {
			camera.Zoom = 3.0
		} else if camera.Zoom == 3.0 {
			camera.Zoom = 4.0
		}
	}
	if rl.IsKeyPressed(rl.KeyKpSubtract) {
		if camera.Zoom == 2.0 {
			camera.Zoom = 1.0
		} else if camera.Zoom == 3.0 {
			camera.Zoom = 2.0
		} else if camera.Zoom == 4.0 {
			camera.Zoom = 3.0
		}
	}

}
func debug() { // MARK: debug
	rl.DrawRectangle(monw-300, 0, 500, monw, rl.Fade(rl.Black, 0.8))
	rl.DrawFPS(monw-290, monh-100)

	mousexTEXT := fmt.Sprintf("%g", mousepos.X)
	mouseyTEXT := fmt.Sprintf("%g", mousepos.Y)
	camerazoomTEXT := fmt.Sprintf("%g", camera.Zoom)
	playerTEXT := strconv.Itoa(player)
	playerhTEXT := strconv.Itoa(playerh)
	playervTEXT := strconv.Itoa(playerv)
	jumpactiveTEXT := strconv.FormatBool(jumpactive)
	drawblockTEXT := strconv.Itoa(drawblock)
	vertbelowtunnelcountTEXT := strconv.Itoa(vertbelowtunnelcount)
	bombblockTEXT := strconv.Itoa(bombblock)
	activequestblockTEXT := strconv.Itoa(activequest.block)
	powerupcountTEXT := strconv.Itoa(powerupcount)
	arrowcountTEXT := strconv.Itoa(arrowcount)
	meteortimerTEXT := strconv.Itoa(meteortimer)
	arrowactiveTEXT := strconv.FormatBool(arrowactive)
	chesttotalTEXT := strconv.Itoa(chesttotal)
	introcirclesonTEXT := strconv.FormatBool(introcircleson)

	rl.DrawText(mousexTEXT, monw-290, 10, 10, rl.White)
	rl.DrawText("mouseX", monw-150, 10, 10, rl.White)
	rl.DrawText(mouseyTEXT, monw-290, 20, 10, rl.White)
	rl.DrawText("mouseY", monw-150, 20, 10, rl.White)
	rl.DrawText(camerazoomTEXT, monw-290, 30, 10, rl.White)
	rl.DrawText("camerazoom", monw-150, 30, 10, rl.White)
	rl.DrawText(playerTEXT, monw-290, 40, 10, rl.White)
	rl.DrawText("player", monw-150, 40, 10, rl.White)
	rl.DrawText(playerhTEXT, monw-290, 50, 10, rl.White)
	rl.DrawText("playerh", monw-150, 50, 10, rl.White)
	rl.DrawText(playervTEXT, monw-290, 60, 10, rl.White)
	rl.DrawText("playerv", monw-150, 60, 10, rl.White)
	rl.DrawText(jumpactiveTEXT, monw-290, 70, 10, rl.White)
	rl.DrawText("jumpactive", monw-150, 70, 10, rl.White)
	rl.DrawText(drawblockTEXT, monw-290, 80, 10, rl.White)
	rl.DrawText("drawblock", monw-150, 80, 10, rl.White)
	rl.DrawText(vertbelowtunnelcountTEXT, monw-290, 90, 10, rl.White)
	rl.DrawText("vertbelowtunnelcount", monw-150, 90, 10, rl.White)
	rl.DrawText(bombblockTEXT, monw-290, 100, 10, rl.White)
	rl.DrawText("bombblock", monw-150, 100, 10, rl.White)
	rl.DrawText(activequestblockTEXT, monw-290, 110, 10, rl.White)
	rl.DrawText("activequestblock", monw-150, 110, 10, rl.White)
	rl.DrawText(powerupcountTEXT, monw-290, 120, 10, rl.White)
	rl.DrawText("powerupcount", monw-150, 120, 10, rl.White)
	rl.DrawText(arrowcountTEXT, monw-290, 130, 10, rl.White)
	rl.DrawText("arrowcount", monw-150, 130, 10, rl.White)
	rl.DrawText(meteortimerTEXT, monw-290, 140, 10, rl.White)
	rl.DrawText("meteortimer", monw-150, 140, 10, rl.White)
	rl.DrawText(arrowactiveTEXT, monw-290, 150, 10, rl.White)
	rl.DrawText("arrowactive", monw-150, 150, 10, rl.White)
	rl.DrawText(chesttotalTEXT, monw-290, 160, 10, rl.White)
	rl.DrawText("chesttotal", monw-150, 160, 10, rl.White)
	rl.DrawText(introcirclesonTEXT, monw-290, 170, 10, rl.White)
	rl.DrawText("introcircleson", monw-150, 170, 10, rl.White)

}
func initialize() { // MARK: initialize
	player = levelw * 148
	player += 50
	playerstart = player
	introv2 = rl.NewVector2(float32((monw/2)/8), float32((monh/2)/8))
	introenemyv2 = rl.NewVector2(float32((monw/2)/8), float32((monh/4)/8))
	gologov2 = rl.NewVector2(-300, float32((monh/2)-262))
	rayliblogov2 = rl.NewVector2(float32((monw/2)-128), float32(monh+10))
	intropinkbackv2 = rl.NewVector2(0, float32(0-monh))
	discoveryinfox = 240
	discoveryinfoy = 240

	inventoryy = monh - (monh / 3)

	createlevel()
	createshops()
	createanimals()
	createspikes()
	createdoors()
	createcoinsgems()
	createinteractables()
	createstartenemies()
	createcharacters()

	extraslayout[player-5] = "bookcase"

	railwaylayout[(player + 3)] = "cartr"
	railwaylayout[(player+2)+levelw] = "track"
	railwaylayout[(player+3)+levelw] = "track"
	railwaylayout[(player+4)+levelw] = "track"
	railwaylayout[(player+5)+levelw] = "track"

}
func resets(choose int) { // MARK: resets

	switch choose {
	case 1:
		vertbelowtunnelcount = 0
		for a := 0; a < len(vertbelowtunnellayout); a++ {
			vertbelowtunnellayout[a] = tunnel{}
		}
	case 2:
		count := 0
		block := updateblock
		for a := 0; a < updatea; a++ {
			deadenemieslayout[block] = ""
			count++
			block++
			if count == updatew {
				count = 0
				block -= updatew
				block += levelw
			}
		}

	}

}
func main() { // MARK: main
	rand.Seed(time.Now().UnixNano()) // random numbers
	rl.SetTraceLogLevel(rl.LogError) // hides info window
	rl.InitWindow(monw, monh, "setscreen")
	setscreen()
	rl.CloseWindow()
	initialize()
	raylib()
}
func setscreen() { // MARK: setscreen
	monhorig = rl.GetScreenHeight()
	monworig = rl.GetScreenWidth()

	monh = monhorig
	monw = monworig
	monh32 = int32(monh)
	monw32 = int32(monw)

	monh = 720
	monw = 1280
	rl.SetWindowSize(monw, monh)
	camera.Zoom = 4.0
	camera.Target.X = 0
	camera.Target.Y = 0

	camera8x.Zoom = 8.0
	camera2x.Zoom = 2.0
	camera4x.Zoom = 4.0
	powerupsxy = rl.NewVector2(5, float32(monh/4-14))

}
func setnewscreen() { // MARK: setnewscreen

	switch resolutionselect {
	case 1:
		monw = 1280
		monh = 720
		rl.SetWindowSize(monw, monh)
		if fullscreenon && !fullscreenactive {
			rl.ToggleFullscreen()
			fullscreenactive = true
		} else if !fullscreenon && fullscreenactive {
			fullscreenactive = false
			rl.ToggleFullscreen()
			rl.SetWindowPosition((monworig-monw)/2, (monhorig-monh)/2)
		}
	case 2:
		monw = 1366
		monh = 768
		rl.SetWindowSize(monw, monh)
		if fullscreenon && !fullscreenactive {
			rl.ToggleFullscreen()
			fullscreenactive = true
		} else if !fullscreenon && fullscreenactive {
			fullscreenactive = false
			rl.ToggleFullscreen()
			rl.SetWindowPosition((monworig-monw)/2, (monhorig-monh)/2)
		}
	case 3:
		monw = 1440
		monh = 900
		rl.SetWindowSize(monw, monh)
		if fullscreenon && !fullscreenactive {
			rl.ToggleFullscreen()
			fullscreenactive = true
		} else if !fullscreenon && fullscreenactive {
			fullscreenactive = false
			rl.ToggleFullscreen()
			rl.SetWindowPosition((monworig-monw)/2, (monhorig-monh)/2)
		}
	case 4:
		monw = 1536
		monh = 864
		rl.SetWindowSize(monw, monh)
		if fullscreenon && !fullscreenactive {
			rl.ToggleFullscreen()
			fullscreenactive = true
		} else if !fullscreenon && fullscreenactive {
			fullscreenactive = false
			rl.ToggleFullscreen()
			rl.SetWindowPosition((monworig-monw)/2, (monhorig-monh)/2)
		}
	case 5:
		monw = 1600
		monh = 900
		rl.SetWindowSize(monw, monh)
		if fullscreenon && !fullscreenactive {
			rl.ToggleFullscreen()
			fullscreenactive = true
		} else if !fullscreenon && fullscreenactive {
			fullscreenactive = false
			rl.ToggleFullscreen()
			rl.SetWindowPosition((monworig-monw)/2, (monhorig-monh)/2)
		}
	case 6:
		monw = 1920
		monh = 1080
		rl.SetWindowSize(monw, monh)
		if fullscreenon && !fullscreenactive {
			rl.ToggleFullscreen()
			fullscreenactive = true
		} else if !fullscreenon && fullscreenactive {
			fullscreenactive = false
			rl.ToggleFullscreen()
			rl.SetWindowPosition((monworig-monw)/2, (monhorig-monh)/2)
		}
	case 7:
		monw = 3840
		monh = 2160
		rl.SetWindowSize(monw, monh)
		if fullscreenon && !fullscreenactive {
			rl.ToggleFullscreen()
			fullscreenactive = true
		} else if !fullscreenon && fullscreenactive {
			fullscreenactive = false
			rl.ToggleFullscreen()
			rl.SetWindowPosition((monworig-monw)/2, (monhorig-monh)/2)
		}
	}

	updatescreenchanges()
}
func updatescreenchanges() { // MARK: updatescreenchanges

	introv2 = rl.NewVector2(float32((monw/2)/8), float32((monh/2)/8))

}
func updatedrawblock() { // MARK: updatedrawblock

	if !moveoff {
		if playerv >= 31 {
			drawblocknext = (player - (levelw * 20)) - draww/3
		}
	}
}
func checkhorizvert(blocknumber, direction int) bool { // MARK: checkhorizvert

	stop := false

	blockvert, blockhoriz = blocknumber%levelw, blocknumber/levelw

	switch direction {

	case 3:
		if blockhoriz > levelh-42 {
			stop = true
		}
	}

	return stop

}
func horizvert() { // MARK: horizvert
	playerh, playerv = player/levelw, player%levelw
}

// random numbers
func rF32(min, max float32) float32 {
	return (rand.Float32() * (max - min)) + min

}
func rInt(min, max int) int {
	return rand.Intn(max-min) + min
}
func rInt32(min, max int) int32 {
	a := int32(rand.Intn(max-min) + min)
	return a
}
func rFloat32(min, max int) float32 {
	a := float32(rand.Intn(max-min) + min)
	return a
}
func flipcoin() bool {
	var b bool
	a := rInt(0, 10001)
	if a < 5000 {
		b = true
	}
	return b
}
func rolldice() int {
	a := rInt(1, 7)
	return a
}
