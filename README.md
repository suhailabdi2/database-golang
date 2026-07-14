#Project : A key Value storage Engine 
Goal is to createa real on-disk database engine, that covers page-based storage, B-trees, write-ahead logs and transactions. 
It is structured in layers as each is its own package
Pager : disk I/O and page management 
Btree : indexing 
wal : crash indexing (write ahead log)

-Why a key-value storage first : A relational engine is just a storage engine pllus a query engine. Building a solid KV storage is way more educational 

Page Size : 4096 bytes to match Os and  hardware block size, one page read equals one I/O op. Larger pages mean shallower B-Trees but more cache pressure and higher town-write risk 

Page Header : every page starts with a fixed header at known bytes : a magic number that starts a byte 0(4 bytes), pageID at byte 4(4 bytes) and pagetype at bytes 8(1 byte). The magic number is the same at every page and detects corrupt and uninitialized pages. the page ID lets the pager verify it read the right page. The page Type tells the B-tree how to interpret the rest of the bytes

Slotted PAge layout : records at page are of variable length. The slot array grows from the header, each slot holding an (offset uint16, length uint16) pair pointing to a record. Records grow backwards from the end of the array, leaving free space in the middle. The free space is the gap in the middle. 

Page type in Go is a 4096 byte fixed array, so the compiler enforces thhe size on compile time. PageType is uint8 not string as it is written on a single byte

Pager Struct holds three unexported fields , file *os.File,the open db file, pagesNumber int ,the conunt of existing pages derived from file size / page size, and pagecache, which is an in memory cache of recently used pages keyed by pageID to prevent unnecessary disk reads
