// please refer to the blog: https://blog.jse.li/posts/torrent/
package torrentfile

import (
	"github.com/jackpal/bencode-go"
)

// look at sample.torrent file to check bencode format
type bencodeInfo struct{
	Pieces string `bencode:"pieces"`
	PieceLength int `bencode:"piece length"` 
	Length int `bencde:"length"`
	Name string `bencode:"name"`
}

type bencodeTorrent struct{
	Announce string `bencode:"announce"`
	Info bencodeInfo `bencode:"info"`
}

// Open.. parse a torrent file
// Bencode example: d8:announce41:http://bttracker.debian.org:6969/announce7:comment35:"Debian CD from cdimage.debian.org"13:creation datei1573903810e9:httpseedsl145:https://cdimage.debian.org/cdimage/release/10.2.0//srv/cdbuilder.debian.org/dst/deb-cd/weekly-builds/amd64/iso-cd/debian-10.2.0-amd64-netinst.iso145:https://cdimage.debian.org/cdimage/archive/10.2.0//srv/cdbuilder.debian.org/dst/deb-cd/weekly-builds/amd64/iso-cd/debian-10.2.0-amd64-netinst.isoe4:infod6:lengthi351272960e4:name31:debian-10.2.0-amd64-netinst.iso12:piece lengthi262144e6:pieces26800:�����PS�^�� (binary blob of the hashes of each piece)ee
func Open(r io.Reader) (*bencodeTorrent error){
	bto := bencodeTorrent{}
	err := bencode.Unmarshal(r, &bto)
	if err != nil{
		return nil, err
	}
	return &bto, nil
}