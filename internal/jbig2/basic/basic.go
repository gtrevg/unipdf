//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package basic ;import _c "github.com/unidoc/unipdf/v3/internal/jbig2/errors";func (_cca *Stack )Push (v interface{}){_cca .Data =append (_cca .Data ,v )};func Abs (v int )int {if v > 0{return v ;};return -v ;};func (_eg *NumSlice )Add (v float32 ){*_eg =append (*_eg ,v )};type IntSlice []int ;func Ceil (numerator ,denominator int )int {if numerator %denominator ==0{return numerator /denominator ;};return (numerator /denominator )+1;};type Stack struct{Data []interface{};Aux *Stack ;};func (_cb *Stack )Pop ()(_edd interface{},_ag bool ){_edd ,_ag =_cb .peek ();if !_ag {return nil ,_ag ;};_cb .Data =_cb .Data [:_cb .top ()];return _edd ,true ;};func (_gf IntsMap )Delete (key uint64 ){delete (_gf ,key )};type NumSlice []float32 ;func (_cc IntsMap )GetSlice (key uint64 )([]int ,bool ){_b ,_ca :=_cc [key ];if !_ca {return nil ,false ;};return _b ,true ;};func (_dg IntSlice )Get (index int )(int ,error ){if index > len (_dg )-1{return 0,_c .Errorf ("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074","\u0069\u006e\u0064\u0065x:\u0020\u0025\u0064\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006eg\u0065",index );};return _dg [index ],nil ;};func Min (x ,y int )int {if x < y {return x ;};return y ;};func Max (x ,y int )int {if x > y {return x ;};return y ;};func (_fc *Stack )top ()int {return len (_fc .Data )-1};func (_f IntsMap )Get (key uint64 )(int ,bool ){_g ,_a :=_f [key ];if !_a {return 0,false ;};if len (_g )==0{return 0,false ;};return _g [0],true ;};func (_cg NumSlice )GetInt (i int )(int ,error ){const _ac ="\u0047\u0065\u0074\u0049\u006e\u0074";if i < 0||i > len (_cg )-1{return 0,_c .Errorf (_ac ,"\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065",i );};_de :=_cg [i ];return int (_de +Sign (_de )*0.5),nil ;};func NewIntSlice (i int )*IntSlice {_gb :=IntSlice (make ([]int ,i ));return &_gb };func (_db NumSlice )Get (i int )(float32 ,error ){if i < 0||i > len (_db )-1{return 0,_c .Errorf ("\u004e\u0075\u006dS\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074","\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065",i );};return _db [i ],nil ;};func (_ed NumSlice )GetIntSlice ()[]int {_bg :=make ([]int ,len (_ed ));for _ab ,_deb :=range _ed {_bg [_ab ]=int (_deb );};return _bg ;};func (_fa *IntSlice )Copy ()*IntSlice {_fb :=IntSlice (make ([]int ,len (*_fa )));copy (_fb ,*_fa );return &_fb ;};func (_cd *IntSlice )Add (v int )error {if _cd ==nil {return _c .Error ("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0041\u0064\u0064","\u0073\u006c\u0069\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064");};*_cd =append (*_cd ,v );return nil ;};func Sign (v float32 )float32 {if v >=0.0{return 1.0;};return -1.0;};func (_gd IntSlice )Size ()int {return len (_gd )};func (_bb *NumSlice )AddInt (v int ){*_bb =append (*_bb ,float32 (v ))};func (_ef *Stack )peek ()(interface{},bool ){_cdb :=_ef .top ();if _cdb ==-1{return nil ,false ;};return _ef .Data [_cdb ],true ;};func NewNumSlice (i int )*NumSlice {_gbd :=NumSlice (make ([]float32 ,i ));return &_gbd };func (_cf *Stack )Peek ()(_gc interface{},_bba bool ){return _cf .peek ()};func (_dbf *Stack )Len ()int {return len (_dbf .Data )};type IntsMap map[uint64 ][]int ;func (_d IntsMap )Add (key uint64 ,value int ){_d [key ]=append (_d [key ],value )};