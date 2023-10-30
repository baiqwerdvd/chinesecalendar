// Code generated by "scripts/generator"; DO NOT EDIT.
// Code generated by "scripts/generator"; DO NOT EDIT.
// Code generated by "scripts/generator"; DO NOT EDIT.
package chinesecalendar

import (
	"time"
	. "github.com/wangzeping722/chinesecalendar/internal"
)

var (
	// 节假日定义
	minDay = Date(2004, 1, 1)
	maxDay = Date(2024, 10, 12)
	// 节假日
	holidays = map[time.Time]Holiday{
		Date(2004, 1, 1):NewYearsDay,
		Date(2004, 1, 22):SpringFestival,
		Date(2004, 1, 23):SpringFestival,
		Date(2004, 1, 24):SpringFestival,
		Date(2004, 1, 25):SpringFestival,
		Date(2004, 1, 26):SpringFestival,
		Date(2004, 1, 27):SpringFestival,
		Date(2004, 1, 28):SpringFestival,
		Date(2004, 5, 1):LabourDay,
		Date(2004, 5, 2):LabourDay,
		Date(2004, 5, 3):LabourDay,
		Date(2004, 5, 4):LabourDay,
		Date(2004, 5, 5):LabourDay,
		Date(2004, 5, 6):LabourDay,
		Date(2004, 5, 7):LabourDay,
		Date(2004, 10, 1):NationalDay,
		Date(2004, 10, 2):NationalDay,
		Date(2004, 10, 3):NationalDay,
		Date(2004, 10, 4):NationalDay,
		Date(2004, 10, 5):NationalDay,
		Date(2004, 10, 6):NationalDay,
		Date(2004, 10, 7):NationalDay,
		Date(2005, 1, 1):NewYearsDay,
		Date(2005, 1, 2):NewYearsDay,
		Date(2005, 1, 3):NewYearsDay,
		Date(2005, 2, 9):SpringFestival,
		Date(2005, 2, 10):SpringFestival,
		Date(2005, 2, 11):SpringFestival,
		Date(2005, 2, 12):SpringFestival,
		Date(2005, 2, 13):SpringFestival,
		Date(2005, 2, 14):SpringFestival,
		Date(2005, 2, 15):SpringFestival,
		Date(2005, 5, 1):LabourDay,
		Date(2005, 5, 2):LabourDay,
		Date(2005, 5, 3):LabourDay,
		Date(2005, 5, 4):LabourDay,
		Date(2005, 5, 5):LabourDay,
		Date(2005, 5, 6):LabourDay,
		Date(2005, 5, 7):LabourDay,
		Date(2005, 10, 1):NationalDay,
		Date(2005, 10, 2):NationalDay,
		Date(2005, 10, 3):NationalDay,
		Date(2005, 10, 4):NationalDay,
		Date(2005, 10, 5):NationalDay,
		Date(2005, 10, 6):NationalDay,
		Date(2005, 10, 7):NationalDay,
		Date(2006, 1, 1):NewYearsDay,
		Date(2006, 1, 2):NewYearsDay,
		Date(2006, 1, 3):NewYearsDay,
		Date(2006, 1, 29):SpringFestival,
		Date(2006, 1, 30):SpringFestival,
		Date(2006, 1, 31):SpringFestival,
		Date(2006, 2, 1):SpringFestival,
		Date(2006, 2, 2):SpringFestival,
		Date(2006, 2, 3):SpringFestival,
		Date(2006, 2, 4):SpringFestival,
		Date(2006, 5, 1):LabourDay,
		Date(2006, 5, 2):LabourDay,
		Date(2006, 5, 3):LabourDay,
		Date(2006, 5, 4):LabourDay,
		Date(2006, 5, 5):LabourDay,
		Date(2006, 5, 6):LabourDay,
		Date(2006, 5, 7):LabourDay,
		Date(2006, 10, 1):NationalDay,
		Date(2006, 10, 2):NationalDay,
		Date(2006, 10, 3):NationalDay,
		Date(2006, 10, 4):NationalDay,
		Date(2006, 10, 5):NationalDay,
		Date(2006, 10, 6):NationalDay,
		Date(2006, 10, 7):NationalDay,
		Date(2007, 1, 1):NewYearsDay,
		Date(2007, 1, 2):NewYearsDay,
		Date(2007, 1, 3):NewYearsDay,
		Date(2007, 2, 18):SpringFestival,
		Date(2007, 2, 19):SpringFestival,
		Date(2007, 2, 20):SpringFestival,
		Date(2007, 2, 21):SpringFestival,
		Date(2007, 2, 22):SpringFestival,
		Date(2007, 2, 23):SpringFestival,
		Date(2007, 2, 24):SpringFestival,
		Date(2007, 5, 1):LabourDay,
		Date(2007, 5, 2):LabourDay,
		Date(2007, 5, 3):LabourDay,
		Date(2007, 5, 4):LabourDay,
		Date(2007, 5, 5):LabourDay,
		Date(2007, 5, 6):LabourDay,
		Date(2007, 5, 7):LabourDay,
		Date(2007, 10, 1):NationalDay,
		Date(2007, 10, 2):NationalDay,
		Date(2007, 10, 3):NationalDay,
		Date(2007, 10, 4):NationalDay,
		Date(2007, 10, 5):NationalDay,
		Date(2007, 10, 6):NationalDay,
		Date(2007, 10, 7):NationalDay,
		Date(2007, 12, 30):NewYearsDay,
		Date(2007, 12, 31):NewYearsDay,
		Date(2008, 1, 1):NewYearsDay,
		Date(2008, 2, 6):SpringFestival,
		Date(2008, 2, 7):SpringFestival,
		Date(2008, 2, 8):SpringFestival,
		Date(2008, 2, 9):SpringFestival,
		Date(2008, 2, 10):SpringFestival,
		Date(2008, 2, 11):SpringFestival,
		Date(2008, 2, 12):SpringFestival,
		Date(2008, 4, 4):TombSweepingDay,
		Date(2008, 4, 5):TombSweepingDay,
		Date(2008, 4, 6):TombSweepingDay,
		Date(2008, 5, 1):LabourDay,
		Date(2008, 5, 2):LabourDay,
		Date(2008, 5, 3):LabourDay,
		Date(2008, 6, 7):DragonBoatFestival,
		Date(2008, 6, 8):DragonBoatFestival,
		Date(2008, 6, 9):DragonBoatFestival,
		Date(2008, 9, 13):MidAutumnFestival,
		Date(2008, 9, 14):MidAutumnFestival,
		Date(2008, 9, 15):MidAutumnFestival,
		Date(2008, 9, 29):NationalDay,
		Date(2008, 9, 30):NationalDay,
		Date(2008, 10, 1):NationalDay,
		Date(2008, 10, 2):NationalDay,
		Date(2008, 10, 3):NationalDay,
		Date(2008, 10, 4):NationalDay,
		Date(2008, 10, 5):NationalDay,
		Date(2009, 1, 1):NewYearsDay,
		Date(2009, 1, 2):NewYearsDay,
		Date(2009, 1, 3):NewYearsDay,
		Date(2009, 1, 25):SpringFestival,
		Date(2009, 1, 26):SpringFestival,
		Date(2009, 1, 27):SpringFestival,
		Date(2009, 1, 28):SpringFestival,
		Date(2009, 1, 29):SpringFestival,
		Date(2009, 1, 30):SpringFestival,
		Date(2009, 1, 31):SpringFestival,
		Date(2009, 4, 4):TombSweepingDay,
		Date(2009, 4, 5):TombSweepingDay,
		Date(2009, 4, 6):TombSweepingDay,
		Date(2009, 5, 1):LabourDay,
		Date(2009, 5, 2):LabourDay,
		Date(2009, 5, 3):LabourDay,
		Date(2009, 5, 28):DragonBoatFestival,
		Date(2009, 5, 29):DragonBoatFestival,
		Date(2009, 5, 30):DragonBoatFestival,
		Date(2009, 10, 1):NationalDay,
		Date(2009, 10, 2):NationalDay,
		Date(2009, 10, 3):MidAutumnFestival,
		Date(2009, 10, 4):NationalDay,
		Date(2009, 10, 5):NationalDay,
		Date(2009, 10, 6):NationalDay,
		Date(2009, 10, 7):NationalDay,
		Date(2009, 10, 8):NationalDay,
		Date(2010, 1, 1):NewYearsDay,
		Date(2010, 1, 2):NewYearsDay,
		Date(2010, 1, 3):NewYearsDay,
		Date(2010, 2, 13):SpringFestival,
		Date(2010, 2, 14):SpringFestival,
		Date(2010, 2, 15):SpringFestival,
		Date(2010, 2, 16):SpringFestival,
		Date(2010, 2, 17):SpringFestival,
		Date(2010, 2, 18):SpringFestival,
		Date(2010, 2, 19):SpringFestival,
		Date(2010, 4, 3):TombSweepingDay,
		Date(2010, 4, 4):TombSweepingDay,
		Date(2010, 4, 5):TombSweepingDay,
		Date(2010, 5, 1):LabourDay,
		Date(2010, 5, 2):LabourDay,
		Date(2010, 5, 3):LabourDay,
		Date(2010, 6, 14):DragonBoatFestival,
		Date(2010, 6, 15):DragonBoatFestival,
		Date(2010, 6, 16):DragonBoatFestival,
		Date(2010, 9, 22):MidAutumnFestival,
		Date(2010, 9, 23):MidAutumnFestival,
		Date(2010, 9, 24):MidAutumnFestival,
		Date(2010, 10, 1):NationalDay,
		Date(2010, 10, 2):NationalDay,
		Date(2010, 10, 3):NationalDay,
		Date(2010, 10, 4):NationalDay,
		Date(2010, 10, 5):NationalDay,
		Date(2010, 10, 6):NationalDay,
		Date(2010, 10, 7):NationalDay,
		Date(2011, 1, 1):NewYearsDay,
		Date(2011, 1, 2):NewYearsDay,
		Date(2011, 1, 3):NewYearsDay,
		Date(2011, 2, 2):SpringFestival,
		Date(2011, 2, 3):SpringFestival,
		Date(2011, 2, 4):SpringFestival,
		Date(2011, 2, 5):SpringFestival,
		Date(2011, 2, 6):SpringFestival,
		Date(2011, 2, 7):SpringFestival,
		Date(2011, 2, 8):SpringFestival,
		Date(2011, 4, 3):TombSweepingDay,
		Date(2011, 4, 4):TombSweepingDay,
		Date(2011, 4, 5):TombSweepingDay,
		Date(2011, 4, 30):LabourDay,
		Date(2011, 5, 1):LabourDay,
		Date(2011, 5, 2):LabourDay,
		Date(2011, 6, 4):DragonBoatFestival,
		Date(2011, 6, 6):DragonBoatFestival,
		Date(2011, 9, 10):MidAutumnFestival,
		Date(2011, 9, 11):MidAutumnFestival,
		Date(2011, 9, 12):MidAutumnFestival,
		Date(2011, 10, 1):NationalDay,
		Date(2011, 10, 2):NationalDay,
		Date(2011, 10, 3):NationalDay,
		Date(2011, 10, 4):NationalDay,
		Date(2011, 10, 5):NationalDay,
		Date(2011, 10, 6):NationalDay,
		Date(2011, 10, 7):NationalDay,
		Date(2012, 1, 1):NewYearsDay,
		Date(2012, 1, 2):NewYearsDay,
		Date(2012, 1, 3):NewYearsDay,
		Date(2012, 1, 22):SpringFestival,
		Date(2012, 1, 23):SpringFestival,
		Date(2012, 1, 24):SpringFestival,
		Date(2012, 1, 25):SpringFestival,
		Date(2012, 1, 26):SpringFestival,
		Date(2012, 1, 27):SpringFestival,
		Date(2012, 1, 28):SpringFestival,
		Date(2012, 4, 2):TombSweepingDay,
		Date(2012, 4, 3):TombSweepingDay,
		Date(2012, 4, 4):TombSweepingDay,
		Date(2012, 4, 29):LabourDay,
		Date(2012, 4, 30):LabourDay,
		Date(2012, 5, 1):LabourDay,
		Date(2012, 6, 22):DragonBoatFestival,
		Date(2012, 6, 24):DragonBoatFestival,
		Date(2012, 9, 30):MidAutumnFestival,
		Date(2012, 10, 1):NationalDay,
		Date(2012, 10, 2):NationalDay,
		Date(2012, 10, 3):NationalDay,
		Date(2012, 10, 4):NationalDay,
		Date(2012, 10, 5):NationalDay,
		Date(2012, 10, 6):NationalDay,
		Date(2012, 10, 7):NationalDay,
		Date(2013, 1, 1):NewYearsDay,
		Date(2013, 1, 2):NewYearsDay,
		Date(2013, 1, 3):NewYearsDay,
		Date(2013, 2, 9):SpringFestival,
		Date(2013, 2, 10):SpringFestival,
		Date(2013, 2, 11):SpringFestival,
		Date(2013, 2, 12):SpringFestival,
		Date(2013, 2, 13):SpringFestival,
		Date(2013, 2, 14):SpringFestival,
		Date(2013, 2, 15):SpringFestival,
		Date(2013, 4, 4):TombSweepingDay,
		Date(2013, 4, 5):TombSweepingDay,
		Date(2013, 4, 6):TombSweepingDay,
		Date(2013, 4, 29):LabourDay,
		Date(2013, 4, 30):LabourDay,
		Date(2013, 5, 1):LabourDay,
		Date(2013, 6, 10):DragonBoatFestival,
		Date(2013, 6, 11):DragonBoatFestival,
		Date(2013, 6, 12):DragonBoatFestival,
		Date(2013, 9, 19):MidAutumnFestival,
		Date(2013, 9, 20):MidAutumnFestival,
		Date(2013, 9, 21):MidAutumnFestival,
		Date(2013, 10, 1):NationalDay,
		Date(2013, 10, 2):NationalDay,
		Date(2013, 10, 3):NationalDay,
		Date(2013, 10, 4):NationalDay,
		Date(2013, 10, 5):NationalDay,
		Date(2013, 10, 6):NationalDay,
		Date(2013, 10, 7):NationalDay,
		Date(2014, 1, 1):NewYearsDay,
		Date(2014, 1, 31):SpringFestival,
		Date(2014, 2, 1):SpringFestival,
		Date(2014, 2, 2):SpringFestival,
		Date(2014, 2, 3):SpringFestival,
		Date(2014, 2, 4):SpringFestival,
		Date(2014, 2, 5):SpringFestival,
		Date(2014, 2, 6):SpringFestival,
		Date(2014, 4, 5):TombSweepingDay,
		Date(2014, 4, 6):TombSweepingDay,
		Date(2014, 4, 7):TombSweepingDay,
		Date(2014, 5, 1):LabourDay,
		Date(2014, 5, 2):LabourDay,
		Date(2014, 5, 3):LabourDay,
		Date(2014, 6, 2):DragonBoatFestival,
		Date(2014, 9, 8):MidAutumnFestival,
		Date(2014, 10, 1):NationalDay,
		Date(2014, 10, 2):NationalDay,
		Date(2014, 10, 3):NationalDay,
		Date(2014, 10, 4):NationalDay,
		Date(2014, 10, 5):NationalDay,
		Date(2014, 10, 6):NationalDay,
		Date(2014, 10, 7):NationalDay,
		Date(2015, 1, 1):NewYearsDay,
		Date(2015, 1, 2):NewYearsDay,
		Date(2015, 1, 3):NewYearsDay,
		Date(2015, 2, 18):SpringFestival,
		Date(2015, 2, 19):SpringFestival,
		Date(2015, 2, 20):SpringFestival,
		Date(2015, 2, 21):SpringFestival,
		Date(2015, 2, 22):SpringFestival,
		Date(2015, 2, 23):SpringFestival,
		Date(2015, 2, 24):SpringFestival,
		Date(2015, 4, 5):TombSweepingDay,
		Date(2015, 4, 6):TombSweepingDay,
		Date(2015, 5, 1):LabourDay,
		Date(2015, 6, 20):DragonBoatFestival,
		Date(2015, 6, 22):DragonBoatFestival,
		Date(2015, 9, 3):NationalDay,
		Date(2015, 9, 4):NationalDay,
		Date(2015, 9, 27):MidAutumnFestival,
		Date(2015, 10, 1):NationalDay,
		Date(2015, 10, 2):NationalDay,
		Date(2015, 10, 3):NationalDay,
		Date(2015, 10, 4):NationalDay,
		Date(2015, 10, 5):NationalDay,
		Date(2015, 10, 6):NationalDay,
		Date(2015, 10, 7):NationalDay,
		Date(2017, 1, 1):NewYearsDay,
		Date(2017, 1, 2):NewYearsDay,
		Date(2017, 1, 27):SpringFestival,
		Date(2017, 1, 28):SpringFestival,
		Date(2017, 1, 29):SpringFestival,
		Date(2017, 1, 30):SpringFestival,
		Date(2017, 1, 31):SpringFestival,
		Date(2017, 2, 1):SpringFestival,
		Date(2017, 2, 2):SpringFestival,
		Date(2017, 2, 7):SpringFestival,
		Date(2017, 2, 8):SpringFestival,
		Date(2017, 2, 9):SpringFestival,
		Date(2017, 2, 10):SpringFestival,
		Date(2017, 2, 11):SpringFestival,
		Date(2017, 2, 12):SpringFestival,
		Date(2017, 2, 13):SpringFestival,
		Date(2017, 4, 2):TombSweepingDay,
		Date(2017, 4, 3):TombSweepingDay,
		Date(2017, 4, 4):TombSweepingDay,
		Date(2017, 5, 1):LabourDay,
		Date(2017, 5, 2):LabourDay,
		Date(2017, 5, 28):DragonBoatFestival,
		Date(2017, 5, 29):DragonBoatFestival,
		Date(2017, 5, 30):DragonBoatFestival,
		Date(2017, 6, 9):DragonBoatFestival,
		Date(2017, 6, 10):DragonBoatFestival,
		Date(2017, 6, 11):DragonBoatFestival,
		Date(2017, 9, 15):MidAutumnFestival,
		Date(2017, 9, 16):MidAutumnFestival,
		Date(2017, 9, 17):MidAutumnFestival,
		Date(2017, 10, 1):NationalDay,
		Date(2017, 10, 2):NationalDay,
		Date(2017, 10, 3):NationalDay,
		Date(2017, 10, 4):MidAutumnFestival,
		Date(2017, 10, 5):NationalDay,
		Date(2017, 10, 6):NationalDay,
		Date(2017, 10, 7):NationalDay,
		Date(2017, 10, 8):NationalDay,
		Date(2018, 1, 1):NewYearsDay,
		Date(2018, 2, 15):SpringFestival,
		Date(2018, 2, 16):SpringFestival,
		Date(2018, 2, 17):SpringFestival,
		Date(2018, 2, 18):SpringFestival,
		Date(2018, 2, 19):SpringFestival,
		Date(2018, 2, 20):SpringFestival,
		Date(2018, 2, 21):SpringFestival,
		Date(2018, 4, 5):TombSweepingDay,
		Date(2018, 4, 6):TombSweepingDay,
		Date(2018, 4, 7):TombSweepingDay,
		Date(2018, 4, 29):LabourDay,
		Date(2018, 4, 30):LabourDay,
		Date(2018, 5, 1):LabourDay,
		Date(2018, 6, 18):DragonBoatFestival,
		Date(2018, 9, 24):MidAutumnFestival,
		Date(2018, 10, 1):NationalDay,
		Date(2018, 10, 2):NationalDay,
		Date(2018, 10, 3):NationalDay,
		Date(2018, 10, 4):NationalDay,
		Date(2018, 10, 5):NationalDay,
		Date(2018, 10, 6):NationalDay,
		Date(2018, 10, 7):NationalDay,
		Date(2018, 12, 30):NewYearsDay,
		Date(2018, 12, 31):NewYearsDay,
		Date(2019, 1, 1):NewYearsDay,
		Date(2019, 2, 4):SpringFestival,
		Date(2019, 2, 5):SpringFestival,
		Date(2019, 2, 6):SpringFestival,
		Date(2019, 2, 7):SpringFestival,
		Date(2019, 2, 8):SpringFestival,
		Date(2019, 2, 9):SpringFestival,
		Date(2019, 2, 10):SpringFestival,
		Date(2019, 4, 5):TombSweepingDay,
		Date(2019, 4, 6):TombSweepingDay,
		Date(2019, 4, 7):TombSweepingDay,
		Date(2019, 5, 1):LabourDay,
		Date(2019, 5, 2):LabourDay,
		Date(2019, 5, 3):LabourDay,
		Date(2019, 5, 4):LabourDay,
		Date(2019, 6, 7):DragonBoatFestival,
		Date(2019, 6, 8):DragonBoatFestival,
		Date(2019, 6, 9):DragonBoatFestival,
		Date(2019, 9, 13):MidAutumnFestival,
		Date(2019, 9, 14):MidAutumnFestival,
		Date(2019, 9, 15):MidAutumnFestival,
		Date(2019, 10, 1):NationalDay,
		Date(2019, 10, 2):NationalDay,
		Date(2019, 10, 3):NationalDay,
		Date(2019, 10, 4):NationalDay,
		Date(2019, 10, 5):NationalDay,
		Date(2019, 10, 6):NationalDay,
		Date(2019, 10, 7):NationalDay,
		Date(2020, 1, 1):NewYearsDay,
		Date(2020, 1, 24):SpringFestival,
		Date(2020, 1, 25):SpringFestival,
		Date(2020, 1, 26):SpringFestival,
		Date(2020, 1, 27):SpringFestival,
		Date(2020, 1, 28):SpringFestival,
		Date(2020, 1, 29):SpringFestival,
		Date(2020, 1, 30):SpringFestival,
		Date(2020, 1, 31):SpringFestival,
		Date(2020, 2, 1):SpringFestival,
		Date(2020, 2, 2):SpringFestival,
		Date(2020, 4, 4):TombSweepingDay,
		Date(2020, 4, 5):TombSweepingDay,
		Date(2020, 4, 6):TombSweepingDay,
		Date(2020, 5, 1):LabourDay,
		Date(2020, 5, 2):LabourDay,
		Date(2020, 5, 3):LabourDay,
		Date(2020, 5, 4):LabourDay,
		Date(2020, 5, 5):LabourDay,
		Date(2020, 6, 25):DragonBoatFestival,
		Date(2020, 6, 26):DragonBoatFestival,
		Date(2020, 6, 27):DragonBoatFestival,
		Date(2020, 10, 1):NationalDay,
		Date(2020, 10, 2):NationalDay,
		Date(2020, 10, 3):NationalDay,
		Date(2020, 10, 4):NationalDay,
		Date(2020, 10, 5):NationalDay,
		Date(2020, 10, 6):NationalDay,
		Date(2020, 10, 7):NationalDay,
		Date(2020, 10, 8):NationalDay,
		Date(2021, 1, 1):NewYearsDay,
		Date(2021, 1, 2):NewYearsDay,
		Date(2021, 1, 3):NewYearsDay,
		Date(2021, 2, 11):SpringFestival,
		Date(2021, 2, 12):SpringFestival,
		Date(2021, 2, 13):SpringFestival,
		Date(2021, 2, 14):SpringFestival,
		Date(2021, 2, 15):SpringFestival,
		Date(2021, 2, 16):SpringFestival,
		Date(2021, 2, 17):SpringFestival,
		Date(2021, 4, 3):TombSweepingDay,
		Date(2021, 4, 4):TombSweepingDay,
		Date(2021, 4, 5):TombSweepingDay,
		Date(2021, 5, 1):LabourDay,
		Date(2021, 5, 2):LabourDay,
		Date(2021, 5, 3):LabourDay,
		Date(2021, 5, 4):LabourDay,
		Date(2021, 5, 5):LabourDay,
		Date(2021, 6, 12):DragonBoatFestival,
		Date(2021, 6, 13):DragonBoatFestival,
		Date(2021, 6, 14):DragonBoatFestival,
		Date(2021, 9, 19):MidAutumnFestival,
		Date(2021, 9, 20):MidAutumnFestival,
		Date(2021, 9, 21):MidAutumnFestival,
		Date(2021, 10, 1):NationalDay,
		Date(2021, 10, 2):NationalDay,
		Date(2021, 10, 3):NationalDay,
		Date(2021, 10, 4):NationalDay,
		Date(2021, 10, 5):NationalDay,
		Date(2021, 10, 6):NationalDay,
		Date(2021, 10, 7):NationalDay,
		Date(2022, 1, 1):NewYearsDay,
		Date(2022, 1, 2):NewYearsDay,
		Date(2022, 1, 3):NewYearsDay,
		Date(2022, 1, 31):SpringFestival,
		Date(2022, 2, 1):SpringFestival,
		Date(2022, 2, 2):SpringFestival,
		Date(2022, 2, 3):SpringFestival,
		Date(2022, 2, 4):SpringFestival,
		Date(2022, 2, 5):SpringFestival,
		Date(2022, 2, 6):SpringFestival,
		Date(2022, 4, 3):TombSweepingDay,
		Date(2022, 4, 4):TombSweepingDay,
		Date(2022, 4, 5):TombSweepingDay,
		Date(2022, 4, 30):LabourDay,
		Date(2022, 5, 1):LabourDay,
		Date(2022, 5, 2):LabourDay,
		Date(2022, 5, 3):LabourDay,
		Date(2022, 5, 4):LabourDay,
		Date(2022, 6, 3):DragonBoatFestival,
		Date(2022, 6, 4):DragonBoatFestival,
		Date(2022, 6, 5):DragonBoatFestival,
		Date(2022, 9, 10):MidAutumnFestival,
		Date(2022, 9, 11):MidAutumnFestival,
		Date(2022, 9, 12):MidAutumnFestival,
		Date(2022, 10, 1):NationalDay,
		Date(2022, 10, 2):NationalDay,
		Date(2022, 10, 3):NationalDay,
		Date(2022, 10, 4):NationalDay,
		Date(2022, 10, 5):NationalDay,
		Date(2022, 10, 6):NationalDay,
		Date(2022, 10, 7):NationalDay,
		Date(2022, 12, 31):NewYearsDay,
		Date(2023, 1, 1):NewYearsDay,
		Date(2023, 1, 2):NewYearsDay,
		Date(2023, 1, 21):SpringFestival,
		Date(2023, 1, 22):SpringFestival,
		Date(2023, 1, 23):SpringFestival,
		Date(2023, 1, 24):SpringFestival,
		Date(2023, 1, 25):SpringFestival,
		Date(2023, 1, 26):SpringFestival,
		Date(2023, 1, 27):SpringFestival,
		Date(2023, 4, 5):TombSweepingDay,
		Date(2023, 4, 29):LabourDay,
		Date(2023, 4, 30):LabourDay,
		Date(2023, 5, 1):LabourDay,
		Date(2023, 5, 2):LabourDay,
		Date(2023, 5, 3):LabourDay,
		Date(2023, 6, 22):DragonBoatFestival,
		Date(2023, 6, 23):DragonBoatFestival,
		Date(2023, 6, 24):DragonBoatFestival,
		Date(2023, 9, 29):MidAutumnFestival,
		Date(2023, 9, 30):NationalDay,
		Date(2023, 10, 1):NationalDay,
		Date(2023, 10, 2):NationalDay,
		Date(2023, 10, 3):NationalDay,
		Date(2023, 10, 4):NationalDay,
		Date(2023, 10, 5):NationalDay,
		Date(2023, 10, 6):NationalDay,
		Date(2024, 1, 1):NewYearsDay,
		Date(2024, 2, 10):SpringFestival,
		Date(2024, 2, 11):SpringFestival,
		Date(2024, 2, 12):SpringFestival,
		Date(2024, 2, 13):SpringFestival,
		Date(2024, 2, 14):SpringFestival,
		Date(2024, 2, 15):SpringFestival,
		Date(2024, 2, 16):SpringFestival,
		Date(2024, 2, 17):SpringFestival,
		Date(2024, 4, 4):TombSweepingDay,
		Date(2024, 4, 5):TombSweepingDay,
		Date(2024, 4, 6):TombSweepingDay,
		Date(2024, 5, 1):LabourDay,
		Date(2024, 5, 2):LabourDay,
		Date(2024, 5, 3):LabourDay,
		Date(2024, 5, 4):LabourDay,
		Date(2024, 5, 5):LabourDay,
		Date(2024, 6, 10):DragonBoatFestival,
		Date(2024, 9, 15):MidAutumnFestival,
		Date(2024, 9, 16):MidAutumnFestival,
		Date(2024, 9, 17):MidAutumnFestival,
		Date(2024, 10, 1):NationalDay,
		Date(2024, 10, 2):NationalDay,
		Date(2024, 10, 3):NationalDay,
		Date(2024, 10, 4):NationalDay,
		Date(2024, 10, 5):NationalDay,
		Date(2024, 10, 6):NationalDay,
		Date(2024, 10, 7):NationalDay,
		
	}

	// 工作日
	workdays = map[time.Time]Holiday{
		Date(2004, 1, 17):SpringFestival,
		Date(2004, 1, 18):SpringFestival,
		Date(2004, 5, 8):LabourDay,
		Date(2004, 5, 9):LabourDay,
		Date(2004, 10, 9):NationalDay,
		Date(2004, 10, 10):NationalDay,
		Date(2005, 2, 5):SpringFestival,
		Date(2005, 2, 6):SpringFestival,
		Date(2005, 4, 30):LabourDay,
		Date(2005, 5, 8):LabourDay,
		Date(2005, 10, 8):NationalDay,
		Date(2005, 10, 9):NationalDay,
		Date(2006, 1, 28):SpringFestival,
		Date(2006, 2, 5):SpringFestival,
		Date(2006, 4, 29):LabourDay,
		Date(2006, 4, 30):LabourDay,
		Date(2006, 9, 30):NationalDay,
		Date(2006, 10, 8):NationalDay,
		Date(2006, 12, 30):NewYearsDay,
		Date(2006, 12, 31):NewYearsDay,
		Date(2007, 2, 17):SpringFestival,
		Date(2007, 2, 25):SpringFestival,
		Date(2007, 4, 28):LabourDay,
		Date(2007, 4, 29):LabourDay,
		Date(2007, 9, 29):NationalDay,
		Date(2007, 9, 30):NationalDay,
		Date(2007, 12, 29):NewYearsDay,
		Date(2008, 2, 2):SpringFestival,
		Date(2008, 2, 3):SpringFestival,
		Date(2008, 5, 4):LabourDay,
		Date(2008, 9, 27):NationalDay,
		Date(2008, 9, 28):NationalDay,
		Date(2009, 1, 4):NewYearsDay,
		Date(2009, 1, 24):SpringFestival,
		Date(2009, 2, 1):SpringFestival,
		Date(2009, 5, 31):DragonBoatFestival,
		Date(2009, 9, 27):NationalDay,
		Date(2009, 10, 10):NationalDay,
		Date(2010, 2, 20):SpringFestival,
		Date(2010, 2, 21):SpringFestival,
		Date(2010, 6, 12):DragonBoatFestival,
		Date(2010, 6, 13):DragonBoatFestival,
		Date(2010, 9, 19):MidAutumnFestival,
		Date(2010, 9, 25):MidAutumnFestival,
		Date(2010, 9, 26):NationalDay,
		Date(2010, 10, 9):NationalDay,
		Date(2011, 1, 30):SpringFestival,
		Date(2011, 2, 12):SpringFestival,
		Date(2011, 4, 2):TombSweepingDay,
		Date(2011, 10, 8):NationalDay,
		Date(2011, 10, 9):NationalDay,
		Date(2011, 12, 31):NewYearsDay,
		Date(2012, 1, 21):SpringFestival,
		Date(2012, 1, 29):SpringFestival,
		Date(2012, 3, 31):TombSweepingDay,
		Date(2012, 4, 1):TombSweepingDay,
		Date(2012, 4, 28):LabourDay,
		Date(2012, 9, 29):NationalDay,
		Date(2013, 1, 5):NewYearsDay,
		Date(2013, 1, 6):NewYearsDay,
		Date(2013, 2, 16):SpringFestival,
		Date(2013, 2, 17):SpringFestival,
		Date(2013, 4, 7):TombSweepingDay,
		Date(2013, 4, 27):LabourDay,
		Date(2013, 4, 28):LabourDay,
		Date(2013, 6, 8):DragonBoatFestival,
		Date(2013, 6, 9):DragonBoatFestival,
		Date(2013, 9, 22):MidAutumnFestival,
		Date(2013, 9, 29):NationalDay,
		Date(2013, 10, 12):NationalDay,
		Date(2014, 1, 26):SpringFestival,
		Date(2014, 2, 8):SpringFestival,
		Date(2014, 5, 4):LabourDay,
		Date(2014, 9, 28):NationalDay,
		Date(2014, 10, 11):NationalDay,
		Date(2015, 1, 4):NewYearsDay,
		Date(2015, 2, 15):SpringFestival,
		Date(2015, 2, 28):SpringFestival,
		Date(2015, 9, 6):NationalDay,
		Date(2015, 10, 10):NationalDay,
		Date(2017, 1, 22):SpringFestival,
		Date(2017, 2, 4):SpringFestival,
		Date(2017, 2, 6):SpringFestival,
		Date(2017, 2, 14):SpringFestival,
		Date(2017, 4, 1):TombSweepingDay,
		Date(2017, 5, 27):DragonBoatFestival,
		Date(2017, 6, 12):DragonBoatFestival,
		Date(2017, 9, 18):MidAutumnFestival,
		Date(2017, 9, 30):NationalDay,
		Date(2017, 10, 8):NationalDay,
		Date(2017, 10, 9):NationalDay,
		Date(2018, 2, 11):SpringFestival,
		Date(2018, 2, 24):SpringFestival,
		Date(2018, 4, 8):TombSweepingDay,
		Date(2018, 4, 28):LabourDay,
		Date(2018, 9, 29):NationalDay,
		Date(2018, 9, 30):NationalDay,
		Date(2018, 12, 29):NewYearsDay,
		Date(2019, 2, 2):SpringFestival,
		Date(2019, 2, 3):SpringFestival,
		Date(2019, 4, 28):LabourDay,
		Date(2019, 5, 5):LabourDay,
		Date(2019, 9, 29):NationalDay,
		Date(2019, 10, 12):NationalDay,
		Date(2020, 1, 19):SpringFestival,
		Date(2020, 4, 26):LabourDay,
		Date(2020, 5, 9):LabourDay,
		Date(2020, 6, 28):DragonBoatFestival,
		Date(2020, 9, 27):NationalDay,
		Date(2020, 10, 10):NationalDay,
		Date(2021, 2, 7):SpringFestival,
		Date(2021, 2, 20):SpringFestival,
		Date(2021, 4, 25):LabourDay,
		Date(2021, 5, 8):LabourDay,
		Date(2021, 9, 18):MidAutumnFestival,
		Date(2021, 9, 26):NationalDay,
		Date(2021, 10, 9):NationalDay,
		Date(2022, 1, 29):SpringFestival,
		Date(2022, 1, 30):SpringFestival,
		Date(2022, 4, 2):TombSweepingDay,
		Date(2022, 4, 24):LabourDay,
		Date(2022, 5, 7):LabourDay,
		Date(2022, 10, 8):NationalDay,
		Date(2022, 10, 9):NationalDay,
		Date(2023, 1, 28):SpringFestival,
		Date(2023, 1, 29):SpringFestival,
		Date(2023, 4, 23):LabourDay,
		Date(2023, 5, 6):LabourDay,
		Date(2023, 6, 25):DragonBoatFestival,
		Date(2023, 10, 7):NationalDay,
		Date(2023, 10, 8):NationalDay,
		Date(2024, 2, 4):SpringFestival,
		Date(2024, 2, 18):SpringFestival,
		Date(2024, 4, 7):TombSweepingDay,
		Date(2024, 4, 28):LabourDay,
		Date(2024, 5, 11):LabourDay,
		Date(2024, 9, 14):MidAutumnFestival,
		Date(2024, 9, 29):NationalDay,
		Date(2024, 10, 12):NationalDay,
		
	}

	// 替代日
	inLieuDays = map[time.Time]Holiday{
		Date(2004, 1, 27):SpringFestival,
		Date(2004, 1, 28):SpringFestival,
		Date(2004, 5, 6):LabourDay,
		Date(2004, 5, 7):LabourDay,
		Date(2004, 10, 6):NationalDay,
		Date(2004, 10, 7):NationalDay,
		Date(2005, 2, 14):SpringFestival,
		Date(2005, 2, 15):SpringFestival,
		Date(2005, 5, 5):LabourDay,
		Date(2005, 5, 6):LabourDay,
		Date(2005, 10, 6):NationalDay,
		Date(2005, 10, 7):NationalDay,
		Date(2006, 2, 2):SpringFestival,
		Date(2006, 2, 3):SpringFestival,
		Date(2006, 5, 4):LabourDay,
		Date(2006, 5, 5):LabourDay,
		Date(2006, 10, 5):NationalDay,
		Date(2006, 10, 6):NationalDay,
		Date(2007, 1, 2):NewYearsDay,
		Date(2007, 1, 3):NewYearsDay,
		Date(2007, 2, 22):SpringFestival,
		Date(2007, 2, 23):SpringFestival,
		Date(2007, 5, 4):LabourDay,
		Date(2007, 5, 7):LabourDay,
		Date(2007, 10, 4):NationalDay,
		Date(2007, 10, 5):NationalDay,
		Date(2007, 12, 31):NewYearsDay,
		Date(2008, 2, 11):SpringFestival,
		Date(2008, 2, 12):SpringFestival,
		Date(2008, 5, 2):LabourDay,
		Date(2008, 9, 29):NationalDay,
		Date(2008, 9, 30):NationalDay,
		Date(2009, 1, 2):NewYearsDay,
		Date(2009, 1, 29):SpringFestival,
		Date(2009, 1, 30):SpringFestival,
		Date(2009, 5, 29):DragonBoatFestival,
		Date(2009, 10, 7):NationalDay,
		Date(2009, 10, 8):NationalDay,
		Date(2010, 2, 18):SpringFestival,
		Date(2010, 2, 19):SpringFestival,
		Date(2010, 6, 14):DragonBoatFestival,
		Date(2010, 6, 15):DragonBoatFestival,
		Date(2010, 9, 23):MidAutumnFestival,
		Date(2010, 9, 24):MidAutumnFestival,
		Date(2010, 10, 6):NationalDay,
		Date(2010, 10, 7):NationalDay,
		Date(2011, 2, 7):SpringFestival,
		Date(2011, 2, 8):SpringFestival,
		Date(2011, 4, 4):TombSweepingDay,
		Date(2011, 10, 6):NationalDay,
		Date(2011, 10, 7):NationalDay,
		Date(2012, 1, 3):NewYearsDay,
		Date(2012, 1, 26):SpringFestival,
		Date(2012, 1, 27):SpringFestival,
		Date(2012, 4, 2):TombSweepingDay,
		Date(2012, 4, 3):TombSweepingDay,
		Date(2012, 4, 30):LabourDay,
		Date(2012, 10, 5):NationalDay,
		Date(2013, 1, 2):NewYearsDay,
		Date(2013, 1, 3):NewYearsDay,
		Date(2013, 2, 14):SpringFestival,
		Date(2013, 2, 15):SpringFestival,
		Date(2013, 4, 5):TombSweepingDay,
		Date(2013, 4, 29):LabourDay,
		Date(2013, 4, 30):LabourDay,
		Date(2013, 6, 10):DragonBoatFestival,
		Date(2013, 6, 11):DragonBoatFestival,
		Date(2013, 9, 20):MidAutumnFestival,
		Date(2013, 10, 4):NationalDay,
		Date(2013, 10, 7):NationalDay,
		Date(2014, 2, 5):SpringFestival,
		Date(2014, 2, 6):SpringFestival,
		Date(2014, 5, 2):LabourDay,
		Date(2014, 10, 6):NationalDay,
		Date(2014, 10, 7):NationalDay,
		Date(2015, 1, 2):NewYearsDay,
		Date(2015, 2, 23):SpringFestival,
		Date(2015, 2, 24):SpringFestival,
		Date(2015, 9, 4):NationalDay,
		Date(2015, 10, 7):NationalDay,
		Date(2017, 2, 1):SpringFestival,
		Date(2017, 2, 2):SpringFestival,
		Date(2017, 2, 11):SpringFestival,
		Date(2017, 2, 12):SpringFestival,
		Date(2017, 4, 3):TombSweepingDay,
		Date(2017, 5, 29):DragonBoatFestival,
		Date(2017, 6, 10):DragonBoatFestival,
		Date(2017, 9, 16):MidAutumnFestival,
		Date(2017, 10, 6):NationalDay,
		Date(2017, 10, 7):NationalDay,
		Date(2018, 2, 19):SpringFestival,
		Date(2018, 2, 20):SpringFestival,
		Date(2018, 2, 21):SpringFestival,
		Date(2018, 4, 6):TombSweepingDay,
		Date(2018, 4, 30):LabourDay,
		Date(2018, 10, 4):NationalDay,
		Date(2018, 10, 5):NationalDay,
		Date(2018, 12, 31):NewYearsDay,
		Date(2019, 2, 4):SpringFestival,
		Date(2019, 2, 8):SpringFestival,
		Date(2019, 5, 2):LabourDay,
		Date(2019, 5, 3):LabourDay,
		Date(2019, 10, 4):NationalDay,
		Date(2019, 10, 7):NationalDay,
		Date(2020, 1, 29):SpringFestival,
		Date(2020, 5, 4):LabourDay,
		Date(2020, 5, 5):LabourDay,
		Date(2020, 6, 26):DragonBoatFestival,
		Date(2020, 10, 7):NationalDay,
		Date(2020, 10, 8):NationalDay,
		Date(2021, 2, 16):SpringFestival,
		Date(2021, 2, 17):SpringFestival,
		Date(2021, 5, 4):LabourDay,
		Date(2021, 5, 5):LabourDay,
		Date(2021, 9, 20):MidAutumnFestival,
		Date(2021, 10, 6):NationalDay,
		Date(2021, 10, 7):NationalDay,
		Date(2022, 2, 3):SpringFestival,
		Date(2022, 2, 4):SpringFestival,
		Date(2022, 4, 4):TombSweepingDay,
		Date(2022, 5, 3):LabourDay,
		Date(2022, 5, 4):LabourDay,
		Date(2022, 10, 6):NationalDay,
		Date(2022, 10, 7):NationalDay,
		Date(2023, 1, 26):SpringFestival,
		Date(2023, 1, 27):SpringFestival,
		Date(2023, 5, 2):LabourDay,
		Date(2023, 5, 3):LabourDay,
		Date(2023, 6, 23):DragonBoatFestival,
		Date(2023, 10, 5):NationalDay,
		Date(2023, 10, 6):NationalDay,
		Date(2024, 2, 15):SpringFestival,
		Date(2024, 2, 16):SpringFestival,
		Date(2024, 4, 5):TombSweepingDay,
		Date(2024, 5, 2):LabourDay,
		Date(2024, 5, 3):LabourDay,
		Date(2024, 9, 17):MidAutumnFestival,
		Date(2024, 10, 4):NationalDay,
		Date(2024, 10, 7):NationalDay,
		
	}
)
