/*Package store is download cache file struct
 */
package store

import (
	"sync"
	"time"
)

// FileExt is filename extension (without dot)
const FileExt = "FictionDown"

// Store is store yaml data file format
type Store struct {
	BookURL     string
	BookName    string
	Author      string    // 作者
	CoverURL    string    // 封面链接
	Description string    // 介绍
	LastUpdate  time.Time `yaml:",omitempty"` // 数据更新时间
	Tmap        []string  //盗版源
	Volumes     []Volume
}

func (store Store) Total() (Done, Example, ExampleDone, AllChaper int) {

	for _, v := range store.Volumes {
		AllChaper += len(v.Chapters)
		for _, v2 := range v.Chapters {
			if len(v2.Text) != 0 {
				Done++
			}
			if len(v2.Example) != 0 {
				Example++
			}
			if (len(v2.Example) != 0) && (len(v2.Text) != 0) {
				ExampleDone++
			}
		}
	}
	return
}

// Volume 卷
type Volume struct {
	Name     string
	IsVIP    bool
	Chapters []Chapter
}

// Chapter 章节
type Chapter struct {
	Name    string
	URL     string
	IsVIP   bool `yaml:",omitempty"`
	TURL    []string
	Text    []string
	Example []string
	Alias   []string   `yaml:"-"`
	MuxLock sync.Mutex `yaml:"-"`
}

// DiffStoreVolume 对比两个卷信息，得到减少的部分和增加的部分
func DiffStoreVolume(oldVolumes, newVolumes Volume) (sub, add []Volume) {
	// 对比卷

	// 对比章节
	return
}
