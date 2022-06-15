package noop

// This is where you can add intents and more possible utterances for intents

var eyeColorList = []string{"eye color", "colo", "i call her", "i foller", "icolor", "ecce", "erior", "ichor", "agricola",
	"change"}
var howOldList = []string{"older", "how old", "old are you", "old or yo"}
var exploreStartList = []string{"start", "plor", "owing", "tailoring", "oding", "oring", "pling"}
var chargerList = []string{"charge", "home", "go to your", "church"}
var sleepList = []string{"flee", "sleep", "sheep"}
var morningList = []string{"morning", "mourning", "mooning", "it bore"}
var nightList = []string{"night", "might"}
var timeList = []string{"time is it", "the time", "what time", "time of"}
var byeList = []string{"good bye", "good by", "good buy", "goodbye"}
var newYearList = []string{"fireworks", "new year", "happy new", "happy to", "have been", "i now you", "no year", "enee",
	"i never", "knew her"}
var holidaysList = []string{"he holds", "christmas", "behold", "holiday"}
var helloList = []string{"hello", "are you", "high", "below", "little", "follow"}
var signInAlexaList = []string{"in intellect", "fine in electa", "in alex", "ing alex", "in an elect", "to alex",
	"in angelica"}
var signOutAlexaList = []string{"in outlet", "i now of elea", "out alexa", "out of ale"}
var loveList = []string{"love", "dove"}
var forwardList = []string{"forward", "for ward", "for word"}
var turnAroundList = []string{"around", "one eighty", "one ate he"}
var turnLeftList = []string{"rn left", "go left", "e left", "ed left", "ernest"}
var turnRightList = []string{"rn right", "go right", "e right", "ernie", "credit", "ed right"}
var rollCubeList = []string{"roll cu", "roll your cu", "all your cu", "roll human", "yorke", "old your he"}
var wheelieList = []string{"pop a w", "polwhele", "olwen", "i wieland", "do a wheel", "doorstone", "thibetan", "powell",
	"welst", "a wheel"}
var fistbumpList = []string{"this pomp", "this pump", "bump", "fistb", "fistf", "this book", "pisto", "with pomp",
	"fison", "first", "fifth", "were fifteen", "if bump", "wisdom", "this bu"}
var blackjackList = []string{"black", "cards", "game"}
var affirmativeList = []string{"yes", "correct", "hit"}
var negativeList = []string{"no", "dont", "stand"}
var nameAskList = []string{"s my name", "t my name"}
var photoList = []string{"photo", "foto", "selby"}
var praiseList = []string{"good", "awesome", "also", "as some", "of them", "battle", "t rob", "the ro"}
var abuseList = []string{"bad", "that ro", "ad ro", "a root"}

// make sure intentsList perfectly matches up with matchListList

var intentsList = []string{"intent_imperative_eyecolor",
	"intent_character_age", "intent_explore_start", "intent_system_charger", "intent_system_sleep",
	"intent_greeting_goodmorning", "intent_greeting_goodnight", "intent_clock_time",
	"intent_greeting_goodbye", "intent_seasonal_happynewyear", "intent_seasonal_happyholidays",
	"intent_greeting_hello", "intent_amazon_signin", "intent_amazon_signin", "intent_imperative_love",
	"intent_imperative_forward", "intent_imperative_turnaround", "intent_imperative_turnleft",
	"intent_imperative_turnright", "intent_play_rollcube", "intent_play_popawheelie", "intent_play_fistbump",
	"intent_play_blackjack", "intent_imperative_affirmative", "intent_imperative_negative", "intent_names_ask",
	"intent_photo_take_extend", "intent_imperative_praise", "intent_imperative_abuse"}

var matchListList = [][]string{eyeColorList, howOldList, exploreStartList,
	chargerList, sleepList, morningList, nightList, timeList, byeList, newYearList, holidaysList,
	helloList, signInAlexaList, signOutAlexaList, loveList, forwardList, turnAroundList, turnLeftList,
	turnRightList, rollCubeList, wheelieList, fistbumpList, blackjackList, affirmativeList, negativeList,
	nameAskList, photoList, praiseList, abuseList}
