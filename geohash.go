package geohash
// geohash.go
// Geohash library for Golong
// (c) 2014 Codefor
// Distributed under the MIT License

import(
    "strings"
)

type T struct {
    Odd string
    Even string
}

var(
    BITS	[5]int
    //BASE32	[]byte
    BASE32	string
    NEIGHBORS	map[string]T
    BORDERS	map[string]T
)

const(
    PRECISION = 12
)

func init(){
    BITS = [5]int{16, 8, 4, 2, 1}
    //BASE32 = []byte("0123456789bcdefghjkmnpqrstuvwxyz")
    BASE32 = "0123456789bcdefghjkmnpqrstuvwxyz"

    /**
    NEIGHBORS.bottom.odd = NEIGHBORS.left.even;
    NEIGHBORS.top.odd = NEIGHBORS.right.even;
    NEIGHBORS.left.odd = NEIGHBORS.bottom.even;
    NEIGHBORS.right.odd = NEIGHBORS.top.even;
    */
    NEIGHBORS = map[string]T{
	"right":T{
	    Odd	    :	"p0r21436x8zb9dcf5h7kjnmqesgutwvy",
	    Even    :	"bc01fg45238967deuvhjyznpkmstqrwx",
	},
	"left":T{
	    Odd	    :	"14365h7k9dcfesgujnmqp0r2twvyx8zb",
	    Even    :	"238967debc01fg45kmstqrwxuvhjyznp",
	},
	"top":T{
	    Odd	    :	"bc01fg45238967deuvhjyznpkmstqrwx",
	    Even    :	"p0r21436x8zb9dcf5h7kjnmqesgutwvy",
	},
	"bottom":T{
	    Odd	    :	"238967debc01fg45kmstqrwxuvhjyznp",
	    Even    :	"14365h7k9dcfesgujnmqp0r2twvyx8zb",
	},
    };

    /**
    BORDERS.bottom.odd = BORDERS.left.even;
    BORDERS.top.odd = BORDERS.right.even;
    BORDERS.left.odd = BORDERS.bottom.even;
    BORDERS.right.odd = BORDERS.top.even;
    */
    BORDERS   = map[string]T{
	"right":T{
	    Odd	    :	"prxz",
	    Even    :	"bcfguvyz",
	},
	"left":T{
	    Odd	    :	"028b",
	    Even    :	"0145hjnp",
	},
	"top":T{
	    Odd	    :	"bcfguvyz",
	    Even    :	"prxz",
	},
	"bottom":T{
	    Odd	    :	"0145hjnp",
	    Even    :	"028b",
	},
    };
}

func Encode(latitude, longitude float64)string{
    is_even := true

    lat := [2]float64{-90.0,90.0}
    lon := [2]float64{-180.0,180.0}
    var mid float64

    var bit,ch int
    var geohash []byte

    for len(geohash) < PRECISION {
	for i:=0;i< 5;i++{
	    if (is_even) {
		mid = (lon[0] + lon[1]) / 2;
		if longitude > mid {
		    ch |= BITS[bit];
		    lon[0] = mid;
		} else{
		    lon[1] = mid;
		}
	    } else {
		mid = (lat[0] + lat[1]) / 2;
		if (latitude > mid) {
		    ch |= BITS[bit];
		    lat[0] = mid;
		} else{
		    lat[1] = mid;
		}
	    }
	    bit ++
	    is_even = !is_even;
	}

	geohash = append(geohash,BASE32[ch])
	bit = 0
	ch = 0
    }
    return string(geohash)
}

func Decode(geohash string) ([3]float64,[3]float64){
    is_even := true

    lat := [3]float64{-90.0,90.0,0}
    lon := [3]float64{-180.0,180.0,0}
    //lat_err,lon_err
    err := [2]float64{90.0,180.0}

    for i:=0; i< len(geohash); i++ {
	c := geohash[i];
	cd := strings.IndexByte(BASE32,c)
	for j:=0; j<5; j++ {
	    mask := BITS[j];
	    if (is_even) {
		err[1] /= 2;
		if cd & mask != 0 {
		    lon[0] = (lon[0] + lon[1]) / 2
		}else{
		    lon[1] = (lon[0] + lon[1]) / 2
		}
	    } else {
		err[0] /= 2;
		if cd & mask != 0 {
		    lat[0] = (lat[0] + lat[1]) / 2
		}else{
		    lat[1] = (lat[0] + lat[1]) / 2
		}
	    }
	    is_even = !is_even;
	}
    }
    lat[2] = (lat[0] + lat[1])/2;
    lon[2] = (lon[0] + lon[1])/2;

    return lat,lon
}

func Adjacent(srcgeohash, dir string) string{
    lastChr := string(srcgeohash[len(srcgeohash) -1])
    base := srcgeohash[:len(srcgeohash)-1]

    var border,neighbor string
    if len(srcgeohash) % 2 == 1{
	border = BORDERS[dir].Odd
	neighbor = NEIGHBORS[dir].Odd
    }else{
	border = BORDERS[dir].Even
	neighbor = NEIGHBORS[dir].Even
    }

    if strings.Index(border,lastChr) != -1{
	base = Adjacent(base, dir)
    }

    return base + string(BASE32[strings.Index(neighbor,lastChr)])
}


