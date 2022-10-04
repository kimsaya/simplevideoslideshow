package player

import "watervideodisplay/src/application/cores/files"

var index int = 0
var listSize int = 0

var IsTimeOff = false

func GetPlayNext() string {
	if IsTimeOff {
		return ""
	}
	list := files.ReadDirectory(files.GetResourceDirectory())
	listSize = len(list)
	if listSize == 0 {
		return files.GetVideoEndedPath()
	}
	if index >= listSize {
		index = 0
	}
	path := files.GetResourceDirectory() + list[index].Name()
	index++
	return path
}

func GetPlay(ind int) string {
	list := files.ReadDirectory(files.GetResourceDirectory())
	listSize = len(list)
	if ind >= listSize {
		ind = 0
	}
	path := files.GetResourceDirectory() + list[ind].Name()
	return path
}
