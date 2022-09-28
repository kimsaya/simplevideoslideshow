package player

import "watervideodisplay/src/application/cores/files"

var index int = 0

var IsTimeOff = true

func GetPlayNext() string {
	if IsTimeOff {
		return ""
	}
	list := files.ReadDirectory(files.GetResourceDirectory())
	if len(list) == 0 {
		return files.GetVideoEndedPath()
	}
	if index >= len(list) {
		index = 0
	}
	path := files.GetResourceDirectory() + list[index].Name()
	index++
	return path
}
