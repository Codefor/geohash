<?php
class Geohash{
    public $PRECISION = 12;
    public $BASE32 = "0123456789bcdefghjkmnpqrstuvwxyz";

    function __construct(){
        $this->BITS = array(16,8,4,2,1);
    }

    public function encode($latitude, $longitude ){
        $latitude = (double)$latitude;
        $longitude = (double)$longitude;

        $is_even = True;

        $lat = array(-90.0,90.0);
        $lon = array(-180.0,180.0);
        $mid = 0;

        $geohash = array();
        $ch = 0;
        while(count($geohash) < $this->PRECISION ){
            foreach($this->BITS as $bit){
                if ($is_even) {
                    $mid = ($lon[0] + $lon[1]) / 2;
                    if( $longitude > $mid ){
                        $ch |= $bit;
                        $lon[0] = $mid;
                    } else{
                        $lon[1] = $mid;
                    }
                } else {
                    $mid = ($lat[0] + $lat[1]) / 2;
                    if ($latitude > $mid) {
                        $ch |= $bit;
                        $lat[0] = $mid;
                    } else{
                        $lat[1] = $mid;
                    }
                }
                $is_even = !$is_even;
            }
            $geohash[] = $this->BASE32[$ch];
            $ch = 0;
        }
        return "".join($geohash);
    }


    public function decode($geohash){
        $is_even = True;

        $lat = array(-90.0,90.0,0);
        $lon = array(-180.0,180.0,0);
        //lat_err,lon_err
        $err = array(90.0,180.0);

        $len = strlen($geohash);
        for($i=0;$i<$len;$i++){
            $cd = strpos($this->BASE32,$geohash[$i]);
            foreach($this->BITS as $mask){
                if ($is_even) {
                    $err[1] /= 2;
                    if( ($cd & $mask) != 0){
                        $lon[0] = ($lon[0] + $lon[1]) / 2;
                    }else{
                        $lon[1] = ($lon[0] + $lon[1]) / 2;
                    }
                } else {
                    $err[0] /= 2;
                    if(($cd & $mask) != 0){
                        $lat[0] = ($lat[0] + $lat[1]) / 2;
                    }else{
                        $lat[1] = ($lat[0] + $lat[1]) / 2;
                    }
                }
                $is_even = !$is_even;
            }
        }
        $lat[2] = ($lat[0] + $lat[1])/2;
        $lon[2] = ($lon[0] + $lon[1])/2;

        return array($lat,$lon);
    }
}
