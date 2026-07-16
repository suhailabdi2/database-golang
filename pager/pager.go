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
	file, err := os.OpenFile(Path, os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		return nil, err
	}
	FileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	Pages := FileInfo.Size() / PageSize
	return &Pager{
		file:        file,
		pagesNumber: int(Pages),
		pagecache:   make(map[uint32]*Page),
	}, nil
}
func (p *Pager) ReadPage(pageID uint32) (Page, error) {
	if PageValue, ok := p.pagecache[pageID]; ok {
		return *PageValue, nil
	}
	ByteOffset := PageSize * int(pageID)
	page := make([]byte, PageSize)
	_, err := p.file.ReadAt(page, int64(ByteOffset))
	if err != nil {
		return Page{}, err
	}
	var P Page
	copy(P[:], page)
	p.pagecache[pageID] = &P
	return P, nil
}
