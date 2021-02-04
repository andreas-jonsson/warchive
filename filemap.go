/*
Copyright (C) 2016-2021 Andreas T Jonsson

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package warchive

var fileMap = []string{
	"MUSIC01.XMI",
	"MUSIC02.XMI",
	"MUSIC03.XMI",
	"MUSIC04.XMI",
	"MUSIC05.XMI",
	"MUSIC06.XMI",
	"MUSIC07.XMI",
	"MUSIC08.XMI",
	"MUSIC09.XMI",
	"MUSIC10.XMI",
	"MUSIC11.XMI",
	"MUSIC12.XMI",
	"MUSIC13.XMI",
	"MUSIC14.XMI",
	"MUSIC15.XMI",
	"MUSIC16.XMI",
	"MUSIC17.XMI",
	"MUSIC18.XMI",
	"MUSIC19.XMI",
	"MUSIC20.XMI",
	"MUSIC21.XMI",
	"MUSIC22.XMI",
	"MUSIC23.XMI",
	"MUSIC24.XMI",
	"MUSIC25.XMI",
	"MUSIC26.XMI",
	"MUSIC27.XMI",
	"MUSIC28.XMI",
	"MUSIC29.XMI",
	"MUSIC30.XMI",
	"MUSIC31.XMI",
	"MUSIC32.XMI",
	"MUSIC33.XMI",
	"MUSIC34.XMI",
	"MUSIC35.XMI",
	"MUSIC36.XMI",
	"MUSIC37.XMI",
	"MUSIC38.XMI",
	"MUSIC39.XMI",
	"MUSIC30.XMI",
	"MUSIC41.XMI",
	"MUSIC42.XMI",
	"MUSIC43.XMI",
	"MUSIC44.XMI",
	"MUSIC45.XMI",
	"ORC11.TER",
	"ORC11.FOG",
	"HUMAN06.TER",
	"HUMAN06.FOG",
	"ORC03.TER",
	"ORC03.FOG",
	"CUSTOMF1.TER",
	"CUSTOMF1.FOG",
	"ORC10.TER",
	"ORC10.FOG",
	"HUMAN02.TER",
	"HUMAN02.FOG",
	"HUMAN05.TER",
	"HUMAN05.FOG",
	"ORC12.TER",
	"ORC12.FOG",
	"CUSTOMF2.TER",
	"CUSTOMF2.FOG",
	"HUMAN01.TER",
	"HUMAN01.FOG",
	"ORC06.TER",
	"ORC06.FOG",
	"HUMAN07.TER",
	"HUMAN07.FOG",
	"HUMAN03.TER",
	"HUMAN03.FOG",
	"HUMAN09.TER",
	"HUMAN09.FOG",
	"HUMAN10.TER",
	"HUMAN10.FOG",
	"HUMAN11.TER",
	"HUMAN11.FOG",
	"HUMAN12.TER",
	"HUMAN12.FOG",
	"ORC01.TER",
	"ORC01.FOG",
	"ORC02.TER",
	"ORC02.FOG",
	"ORC05.TER",
	"ORC05.FOG",
	"ORC07.TER",
	"ORC07.FOG",
	"ORC09.TER",
	"ORC09.FOG",
	"CUSTOMS1.TER",
	"CUSTOMS1.FOG",
	"CUSTOMS2.TER",
	"CUSTOMS2.FOG",
	"ORC04.TER",
	"ORC04.FOG",
	"HUMAN08.TER",
	"HUMAN08.FOG",
	"HUMAN04.TER",
	"HUMAN04.FOG",
	"ORC08.TER",
	"ORC08.FOG",
	"CUSTOMD1.TER",
	"CUSTOMD1.FOG",
	"CUSTOMD2.TER",
	"CUSTOMD2.FOG",
	"CUSTOMD3.TER",
	"CUSTOMD3.FOG",
	"CUSTOMD4.TER",
	"CUSTOMD4.FOG",
	"CUSTOMD5.TER",
	"CUSTOMD5.FOG",
	"CUSTOMD6.TER",
	"CUSTOMD6.FOG",
	"CUSTOMD7.TER",
	"CUSTOMD7.FOG",
	"CUSTOMD8.TER",
	"CUSTOMD8.FOG",
	"HUMAN01.MIS",
	"ORC01.MIS",
	"HUMAN02.MIS",
	"ORC02.MIS",
	"HUMAN03.MIS",
	"ORC03.MIS",
	"HUMAN04.MIS",
	"ORC04.MIS",
	"HUMAN05.MIS",
	"ORC05.MIS",
	"HUMAN06.MIS",
	"ORC06.MIS",
	"HUMAN07.MIS",
	"ORC07.MIS",
	"HUMAN08.MIS",
	"ORC08.MIS",
	"HUMAN09.MIS",
	"ORC09.MIS",
	"HUMAN10.MIS",
	"ORC10.MIS",
	"HUMAN11.MIS",
	"ORC11.MIS",
	"HUMAN12.MIS",
	"ORC12.MIS",
	"DEMO01.MIS",
	"DEMO02.MIS",
	"DEMO03.MIS",
	"DEMO04.MIS",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"FOREST.PTR",
	"FOREST.TIL",
	"FOREST.PAL",
	"SWAMP.PTR",
	"SWAMP.TIL",
	"SWAMP.PAL",
	"DUNGEON.PTR",
	"DUNGEON.TIL",
	"DUNGEON.PAL",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"SPRITE0.PAL",
	"SPRITE1.PAL",
	"",
	"",
	"",
	"",
	"BLIZZARD.IMG",
	"BLIZZARD.PAL",
	"IHRESBAR.IMG",
	"IORESBAR.IMG",
	"IHRIGBAR.IMG",
	"IORIGBAR.IMG",
	"IHBOTBAR.IMG",
	"IOBOTBAR.IMG",
	"IHMMAP01.IMG",
	"IOMMAP01.IMG",
	"IHLPANEL.IMG",
	"IOLPANEL.IMG",
	"IHMMAP02.IMG",
	"IOMMAP02.IMG",
	"INSERTCD.DLG",
	"BUY01.DLG",
	"BUY02.DLG",
	"HPOPUP.IMG",
	"OPOPUP.IMG",
	"HBIGBAR.IMG",
	"OBIGBAR.IMG",
	"BUTTON11.IMG",
	"BUTTON12.IMG",
	"BUTTON21.IMG",
	"BUTTON22.IMG",
	"BUTTON31.IMG",
	"BUTTON42.IMG",
	"TITLEBOT.IMG",
	"BTNLEFT1.IMG",
	"BTNLEFT2.IMG",
	"BTNRIGH1.IMG",
	"BTNRIGH2.IMG",
	"PANEL.IMG",
	"HSAVMENU.IMG",
	"OSAVMENU.IMG",
	"SAVSLOT1.IMG",
	"SAVSLOT2.IMG",
	"",
	"HELP.IMG",
	"HELP.PAL",
	"BNKOK01.IMG",
	"BTNOK02.IMG",
	"TITLETOP.IMG",
	"",
	"TITLE.PAL",
	"TITLE.IMG",
	"CURSORS.PAL",
	"NORMAL.CUR",
	"NOPE.CUR",
	"TARGET01.CUR",
	"TARGET02.CUR",
	"TARGET03.CUR",
	"INSPECT.CUR",
	"CROSHAIR.CUR",
	"TIME.CUR",
	"SCROLLT.CUR",
	"SCROLLTR.CUR",
	"SCROLLR.CUR",
	"SCROLLBR.CUR",
	"SCROLLB.CUR",
	"SCROLLBL.CUR",
	"SCROLLL.CUR",
	"SCROLLTL.CUR",
	"FOOTMAN.SPR",
	"GRUNT.SPR",
	"PEASANT.SPR",
	"PEON.SPR",
	"HCATAPUL.SPR",
	"OCATAPUL.SPR",
	"KNIGHT.SPR",
	"RAIDER.SPR",
	"ARCHER.SPR",
	"SPEARMAN.SPR",
	"CONJURER.SPR",
	"WARLOCK.SPR",
	"CLERIC.SPR",
	"NECROLYT.SPR",
	"MEDIVH.SPR",
	"LOTHAR.SPR",
	"WOUNDED.SPR",
	"GARONA.SPR",
	"OGRE.SPR",
	"SPIDER.SPR",
	"SLIME.SPR",
	"FIREELEM.SPR",
	"SCORPION.SPR",
	"BRIGAND.SPR",
	"THEDEAD.SPR",
	"SKELETON.SPR",
	"DAEMON.SPR",
	"WATERELE.SPR",
	"HFARM.SPR",
	"OFARM.SPR",
	"HBARRACK.SPR",
	"OBARRACK.SPR",
	"HCHURCH.SPR",
	"OTEMPLE.SPR",
	"HTOWER.SPR",
	"OTOWER.SPR",
	"HTOWNHAL.SPR",
	"OTOWNHAL.SPR",
	"HLUMBERM.SPR",
	"OLUMBERM.SPR",
	"HSTABLES.SPR",
	"OKENNELS.SPR",
	"HSMITH.SPR",
	"OSMITH.SPR",
	"HSTORMWI.SPR",
	"OBLACKRO.SPR",
	"GOLDMINE.SPR",
	"CORSPES.SPR",
	"PEASANTW.SPR",
	"PEONW.SPR",
	"PEASANTG.SPR",
	"PEONG.SPR",
	"HFARMCO.SPR",
	"OFARMCO.SPR",
	"HBARRACO.SPR",
	"OBARRACO.SPR",
	"HCHURCCO.SPR",
	"OTEMPLCO.SPR",
	"HTOWERCO.SPR",
	"OTOWERCO.SPR",
	"HTHALLCO.SPR",
	"OTHALLCO.SPR",
	"HLMILLCO.SPR",
	"OLMILLCO.SPR",
	"HSTABLCO.SPR",
	"OKENNECO.SPR",
	"HSMITHCO.SPR",
	"OSMITHCO.SPR",
	"FIREBAL1.SPR",
	"BOULDER.SPR",
	"ARROW.SPR",
	"POICLOUD.SPR",
	"FIRERAIN.SPR",
	"FLAMES.SPR",
	"LARGEFLA.SPR",
	"EXPLOSI.SPR",
	"HEALING.SPR",
	"COLLAPSE.SPR",
	"WATERBUL.SPR",
	"FIREBAL2.SPR",
	"IOPORTRA.SPR",
	"IHPORTRA.SPR",
	"PORTRAIT.SPR",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"HBRIEFING.IMG",
	"OBRIEFING.IMG",
	"HBRIEFING.PAL",
	"OBRIEFING.PAL",
	"OBRIEF01.SPR",
	"OBRIEF02.SPR",
	"OBRIEF03.SPR",
	"HBRIEF01.SPR",
	"HBRIEF02.SPR",
	"HBRIEF03.SPR",
	"HBRIEF04.SPR",
	"ORC01.MTX",
	"ORC02.MTX",
	"ORC03.MTX",
	"ORC04.MTX",
	"ORC05.MTX",
	"ORC06.MTX",
	"ORC07.MTX",
	"ORC08.MTX",
	"ORC09.MTX",
	"ORC10.MTX",
	"ORC11.MTX",
	"ORC12.MTX",
	"HUMAN01.MTX",
	"HUMAN02.MTX",
	"HUMAN03.MTX",
	"HUMAN04.MTX",
	"HUMAN05.MTX",
	"HUMAN06.MTX",
	"HUMAN07.MTX",
	"HUMAN08.MTX",
	"HUMAN09.MTX",
	"HUMAN10.MTX",
	"HUMAN11.MTX",
	"HUMAN12.MTX",
	"HKING.IMG",
	"HKING.PAL",
	"OKING.IMG",
	"OKING.PAL",
	"OKING01.SPR",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"HKINGANI.IMG",
	"OKINGANI.IMG",
	"BLIZZARD.WAV",
	"GATE.WAV",
	"CONSTRCT.VOC",
	"EXPLODE.VOC",
	"CATAPULT.VOC",
	"AXE1.VOC",
	"AXE2.VOC",
	"AXE3.VOC",
	"AXE4.VOC",
	"BLDEXPL1.VOC",
	"BLDEXPL2.VOC",
	"BLDEXPL3.VOC",
	"MENU.VOC",
	"BUTTON.WAV",
	"ERROR.VOC",
	"SWORD1.VOC",
	"SWORD2.VOC",
	"SWORD3.VOC",
	"CLUB.VOC",
	"FIREHIT.VOC",
	"FIREBALL.VOC",
	"BOWFIRE.VOC",
	"BOWHIT.VOC",
	"ONEAR1.VOC",
	"ONEAR2.VOC",
	"HNEAR1.WAV",
	"HNEAR2.WAV",
	"ODEAD.VOC",
	"HDEAD.VOC",
	"OWRKDONE.VOC",
	"HWRKDONE.WAV",
	"OHELP1.VOC",
	"OHELP2.WAV",
	"HHELP1.WAV",
	"HHELP2.WAV",
	"OREADY.VOC",
	"HREADY.WAV",
	"OYESSIR1.VOC",
	"OYESSIR2.VOC",
	"OYESSIR3.VOC",
	"OYESSIR4.VOC",
	"HYESSIR1.WAV",
	"HYESSIR2.WAV",
	"OWHAT1.VOC",
	"OWHAT2.VOC",
	"OWHAT3.VOC",
	"OWHAT4.VOC",
	"OWHAT5.VOC",
	"HWHAT1.WAV",
	"HWHAT2.WAV",
	"HWHAT3.WAV",
	"HWHAT4.WAV",
	"HWHAT5.WAV",
	"OPISSED1.VOC",
	"OPISSED2.VOC",
	"OPISSED3.WAV",
	"HPISSED1.WAV",
	"HPISSED2.WAV",
	"HPISSED3.WAV",
	"SLIME.WAV",
	"SPELL.WAV",
	"THUNK.WAV",
	"OCHANT.WAV",
	"HCHANT.WAV",
	"KENNELS.WAV",
	"STABLES.WAV",
	"SMITH.WAV",
	"BONFIRE8.WAV",
	"FIREWRK1.WAV",
	"FIREWRK2.WAV",
	"INTRO08.WAV",
	"INTRO18.WAV",
	"INTRO28.WAV",
	"INTRO38.WAV",
	"INTRO48.WAV",
	"HUMAN01.WAV",
	"HUMAN02.WAV",
	"HUMAN03.WAV",
	"HUMAN04.WAV",
	"HUMAN05.WAV",
	"HUMAN06.WAV",
	"HUMAN07.WAV",
	"HUMAN08.WAV",
	"HUMAN09.WAV",
	"HUMAN10.WAV",
	"HUMAN11.WAV",
	"HUMAN12.WAV",
	"ORC01.WAV",
	"ORC02.WAV",
	"ORC03.WAV",
	"ORC04.WAV",
	"ORC05.WAV",
	"ORC06.WAV",
	"ORC07.WAV",
	"ORC08.WAV",
	"ORC09.WAV",
	"ORC10.WAV",
	"ORC11.WAV",
	"ORC12.WAV",
	"ORCSV18.WAV",
	"ORCSV28.WAV",
	"ORCSV38.WAV",
	"HUMSV18.WAV",
	"HUMSV28.WAV",
	"HUMSV38.WAV",
	"HUMLOSE8.WAV",
	"ORCLOSE8.WAV",
	"HOUTRO18.WAV",
	"HOUTRO28.WAV",
	"OOUTRO18.WAV",
	"OOUTRO28.WAV",
}
