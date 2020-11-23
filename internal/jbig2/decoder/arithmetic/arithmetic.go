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

package arithmetic ;import (_b "fmt";_c "github.com/unidoc/unipdf/v3/common";_df "github.com/unidoc/unipdf/v3/internal/bitwise";_g "github.com/unidoc/unipdf/v3/internal/jbig2/internal";_e "io";_d "strings";);func NewStats (contextSize int32 ,index int32 )*DecoderStats {return &DecoderStats {_cbf :index ,_eca :contextSize ,_ffa :make ([]byte ,contextSize ),_dd :make ([]byte ,contextSize )};};func (_ecac *DecoderStats )Reset (){for _fbb :=0;_fbb < len (_ecac ._ffa );_fbb ++{_ecac ._ffa [_fbb ]=0;_ecac ._dd [_fbb ]=0;};};func (_da *Decoder )lpsExchange (_fab *DecoderStats ,_faf int32 ,_fca uint32 )int {_dcf :=_fab .getMps ();if _da ._bd < _fca {_fab .setEntry (int (_ea [_faf ][1]));_da ._bd =_fca ;return int (_dcf );};if _ea [_faf ][3]==1{_fab .toggleMps ();};_fab .setEntry (int (_ea [_faf ][2]));_da ._bd =_fca ;return int (1-_dcf );};func (_ab *Decoder )DecodeInt (stats *DecoderStats )(int32 ,error ){var (_eac ,_bcc int32 ;_bb ,_f ,_de int ;_fc error ;);if stats ==nil {stats =NewStats (512,1);};_ab ._bc =1;_f ,_fc =_ab .decodeIntBit (stats );if _fc !=nil {return 0,_fc ;};_bb ,_fc =_ab .decodeIntBit (stats );if _fc !=nil {return 0,_fc ;};if _bb ==1{_bb ,_fc =_ab .decodeIntBit (stats );if _fc !=nil {return 0,_fc ;};if _bb ==1{_bb ,_fc =_ab .decodeIntBit (stats );if _fc !=nil {return 0,_fc ;};if _bb ==1{_bb ,_fc =_ab .decodeIntBit (stats );if _fc !=nil {return 0,_fc ;};if _bb ==1{_bb ,_fc =_ab .decodeIntBit (stats );if _fc !=nil {return 0,_fc ;};if _bb ==1{_de =32;_bcc =4436;}else {_de =12;_bcc =340;};}else {_de =8;_bcc =84;};}else {_de =6;_bcc =20;};}else {_de =4;_bcc =4;};}else {_de =2;_bcc =0;};for _ebe :=0;_ebe < _de ;_ebe ++{_bb ,_fc =_ab .decodeIntBit (stats );if _fc !=nil {return 0,_fc ;};_eac =(_eac <<1)|int32 (_bb );};_eac +=_bcc ;if _f ==0{return _eac ,nil ;}else if _f ==1&&_eac > 0{return -_eac ,nil ;};return 0,_g .ErrOOB ;};func (_bdge *DecoderStats )toggleMps (){_bdge ._dd [_bdge ._cbf ]^=1};var (_ea =[][4]uint32 {{0x5601,1,1,1},{0x3401,2,6,0},{0x1801,3,9,0},{0x0AC1,4,12,0},{0x0521,5,29,0},{0x0221,38,33,0},{0x5601,7,6,1},{0x5401,8,14,0},{0x4801,9,14,0},{0x3801,10,14,0},{0x3001,11,17,0},{0x2401,12,18,0},{0x1C01,13,20,0},{0x1601,29,21,0},{0x5601,15,14,1},{0x5401,16,14,0},{0x5101,17,15,0},{0x4801,18,16,0},{0x3801,19,17,0},{0x3401,20,18,0},{0x3001,21,19,0},{0x2801,22,19,0},{0x2401,23,20,0},{0x2201,24,21,0},{0x1C01,25,22,0},{0x1801,26,23,0},{0x1601,27,24,0},{0x1401,28,25,0},{0x1201,29,26,0},{0x1101,30,27,0},{0x0AC1,31,28,0},{0x09C1,32,29,0},{0x08A1,33,30,0},{0x0521,34,31,0},{0x0441,35,32,0},{0x02A1,36,33,0},{0x0221,37,34,0},{0x0141,38,35,0},{0x0111,39,36,0},{0x0085,40,37,0},{0x0049,41,38,0},{0x0025,42,39,0},{0x0015,43,40,0},{0x0009,44,41,0},{0x0005,45,42,0},{0x0001,45,43,0},{0x5601,46,46,0}};);func (_fe *Decoder )init ()error {_fe ._eb =_fe ._dc .StreamPosition ();_ac ,_ba :=_fe ._dc .ReadByte ();if _ba !=nil {_c .Log .Debug ("B\u0075\u0066\u0066\u0065\u0072\u0030 \u0072\u0065\u0061\u0064\u0042\u0079\u0074\u0065\u0020f\u0061\u0069\u006ce\u0064.\u0020\u0025\u0076",_ba );return _ba ;};_fe ._eab =_ac ;_fe ._cc =uint64 (_ac )<<16;if _ba =_fe .readByte ();_ba !=nil {return _ba ;};_fe ._cc <<=7;_fe ._ec -=7;_fe ._bd =0x8000;_fe ._bf ++;return nil ;};func New (r _df .StreamReader )(*Decoder ,error ){_cce :=&Decoder {_dc :r ,ContextSize :[]uint32 {16,13,10,10},ReferedToContextSize :[]uint32 {13,10}};if _dfc :=_cce .init ();_dfc !=nil {return nil ,_dfc ;};return _cce ,nil ;};func (_dfg *DecoderStats )cx ()byte {return _dfg ._ffa [_dfg ._cbf ]};func (_acc *DecoderStats )Overwrite (dNew *DecoderStats ){for _ed :=0;_ed < len (_acc ._ffa );_ed ++{_acc ._ffa [_ed ]=dNew ._ffa [_ed ];_acc ._dd [_ed ]=dNew ._dd [_ed ];};};func (_gb *Decoder )DecodeBit (stats *DecoderStats )(int ,error ){var (_ga int ;_bcd =_ea [stats .cx ()][0];_dfcc =int32 (stats .cx ()););defer func (){_gb ._bf ++}();_gb ._bd -=_bcd ;if (_gb ._cc >>16)< uint64 (_bcd ){_ga =_gb .lpsExchange (stats ,_dfcc ,_bcd );if _dce :=_gb .renormalize ();_dce !=nil {return 0,_dce ;};}else {_gb ._cc -=uint64 (_bcd )<<16;if (_gb ._bd &0x8000)==0{_ga =_gb .mpsExchange (stats ,_dfcc );if _gag :=_gb .renormalize ();_gag !=nil {return 0,_gag ;};}else {_ga =int (stats .getMps ());};};return _ga ,nil ;};type Decoder struct{ContextSize []uint32 ;ReferedToContextSize []uint32 ;_dc _df .StreamReader ;_eab uint8 ;_cc uint64 ;_bd uint32 ;_bc int64 ;_ec int32 ;_bf int32 ;_eb int64 ;};func (_gc *Decoder )DecodeIAID (codeLen uint64 ,stats *DecoderStats )(int64 ,error ){_gc ._bc =1;var _fd uint64 ;for _fd =0;_fd < codeLen ;_fd ++{stats .SetIndex (int32 (_gc ._bc ));_gf ,_ccg :=_gc .DecodeBit (stats );if _ccg !=nil {return 0,_ccg ;};_gc ._bc =(_gc ._bc <<1)|int64 (_gf );};_fa :=_gc ._bc -(1<<codeLen );return _fa ,nil ;};func (_bdgb *DecoderStats )SetIndex (index int32 ){_bdgb ._cbf =index };func (_dgf *Decoder )decodeIntBit (_fbf *DecoderStats )(int ,error ){_fbf .SetIndex (int32 (_dgf ._bc ));_dcb ,_ff :=_dgf .DecodeBit (_fbf );if _ff !=nil {_c .Log .Debug ("\u0041\u0072\u0069\u0074\u0068\u006d\u0065t\u0069\u0063\u0044e\u0063\u006f\u0064e\u0072\u0020'\u0064\u0065\u0063\u006f\u0064\u0065I\u006etB\u0069\u0074\u0027\u002d\u003e\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0042\u0069\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0076",_ff );return _dcb ,_ff ;};if _dgf ._bc < 256{_dgf ._bc =((_dgf ._bc <<uint64 (1))|int64 (_dcb ))&0x1ff;}else {_dgf ._bc =(((_dgf ._bc <<uint64 (1)|int64 (_dcb ))&511)|256)&0x1ff;};return _dcb ,nil ;};func (_gg *Decoder )renormalize ()error {for {if _gg ._ec ==0{if _bdg :=_gg .readByte ();_bdg !=nil {return _bdg ;};};_gg ._bd <<=1;_gg ._cc <<=1;_gg ._ec --;if (_gg ._bd &0x8000)!=0{break ;};};_gg ._cc &=0xffffffff;return nil ;};func (_gd *DecoderStats )getMps ()byte {return _gd ._dd [_gd ._cbf ]};func (_gagf *Decoder )readByte ()error {if _gagf ._dc .StreamPosition ()> _gagf ._eb {if _ ,_ag :=_gagf ._dc .Seek (-1,_e .SeekCurrent );_ag !=nil {return _ag ;};};_bbf ,_dg :=_gagf ._dc .ReadByte ();if _dg !=nil {return _dg ;};_gagf ._eab =_bbf ;if _gagf ._eab ==0xFF{_fb ,_dge :=_gagf ._dc .ReadByte ();if _dge !=nil {return _dge ;};if _fb > 0x8F{_gagf ._cc +=0xFF00;_gagf ._ec =8;if _ ,_ef :=_gagf ._dc .Seek (-2,_e .SeekCurrent );_ef !=nil {return _ef ;};}else {_gagf ._cc +=uint64 (_fb )<<9;_gagf ._ec =7;};}else {_bbf ,_dg =_gagf ._dc .ReadByte ();if _dg !=nil {return _dg ;};_gagf ._eab =_bbf ;_gagf ._cc +=uint64 (_gagf ._eab )<<8;_gagf ._ec =8;};_gagf ._cc &=0xFFFFFFFFFF;return nil ;};func (_gac *DecoderStats )Copy ()*DecoderStats {_gbf :=&DecoderStats {_eca :_gac ._eca ,_ffa :make ([]byte ,_gac ._eca )};for _bbb :=0;_bbb < len (_gac ._ffa );_bbb ++{_gbf ._ffa [_bbb ]=_gac ._ffa [_bbb ];};return _gbf ;};type DecoderStats struct{_cbf int32 ;_eca int32 ;_ffa []byte ;_dd []byte ;};func (_db *DecoderStats )setEntry (_ecg int ){_bbc :=byte (_ecg &0x7f);_db ._ffa [_db ._cbf ]=_bbc };func (_fce *Decoder )mpsExchange (_bcce *DecoderStats ,_cf int32 )int {_fec :=_bcce ._dd [_bcce ._cbf ];if _fce ._bd < _ea [_cf ][0]{if _ea [_cf ][3]==1{_bcce .toggleMps ();};_bcce .setEntry (int (_ea [_cf ][2]));return int (1-_fec );};_bcce .setEntry (int (_ea [_cf ][1]));return int (_fec );};func (_bcb *DecoderStats )String ()string {_edf :=&_d .Builder {};_edf .WriteString (_b .Sprintf ("S\u0074\u0061\u0074\u0073\u003a\u0020\u0020\u0025\u0064\u000a",len (_bcb ._ffa )));for _dgfg ,_edb :=range _bcb ._ffa {if _edb !=0{_edf .WriteString (_b .Sprintf ("N\u006f\u0074\u0020\u007aer\u006f \u0061\u0074\u003a\u0020\u0025d\u0020\u002d\u0020\u0025\u0064\u000a",_dgfg ,_edb ));};};return _edf .String ();};