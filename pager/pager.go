package pager

import "os"

const PageSize = 4096

const MagicNumber uint32 = 0xDEADB0DB

type PageType uint8

const InternalNode PageType = 0x01
const LeafNode PageType = 0x02
const OverFlow PageType = 0x03
const OffsetMagic = 0
const OffsetPageID = 4
const OffsetPageType = 8

type Page [PageSize]byte
type Pager struct {
	file        *os.File
	pagesNumber int
	pagecache   map[uint32]*Page
}

func CreatePager(Path string) (*Pager, error) {

}
