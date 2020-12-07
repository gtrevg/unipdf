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

package sampling ;import (_g "github.com/unidoc/unipdf/v3/internal/bitwise";_ae "github.com/unidoc/unipdf/v3/internal/imageutil";_a "io";);func (_ad *Reader )ReadSample ()(uint32 ,error ){if _ad ._dc ==_ad ._d .Height {return 0,_a .EOF ;};_bc ,_ge :=_ad ._b .ReadBits (byte (_ad ._d .BitsPerComponent ));if _ge !=nil {return 0,_ge ;};_ad ._ec --;if _ad ._ec ==0{_ad ._ec =_ad ._d .ColorComponents ;_ad ._de ++;};if _ad ._de ==_ad ._d .Width {if _ad ._c {_ad ._b .ConsumeRemainingBits ();};_ad ._de =0;_ad ._dc ++;};return uint32 (_bc ),nil ;};type Reader struct{_d _ae .ImageBase ;_b *_g .Reader ;_de ,_dc ,_ec int ;_c bool ;};type SampleReader interface{ReadSample ()(uint32 ,error );ReadSamples (_aed []uint32 )error ;};func ResampleBytes (data []byte ,bitsPerSample int )[]uint32 {var _dea []uint32 ;_cfg :=bitsPerSample ;var _cd uint32 ;var _gf byte ;_be :=0;_ab :=0;_fd :=0;for _fd < len (data ){if _be > 0{_gd :=_be ;if _cfg < _gd {_gd =_cfg ;};_cd =(_cd <<uint (_gd ))|uint32 (_gf >>uint (8-_gd ));_be -=_gd ;if _be > 0{_gf =_gf <<uint (_gd );}else {_gf =0;};_cfg -=_gd ;if _cfg ==0{_dea =append (_dea ,_cd );_cfg =bitsPerSample ;_cd =0;_ab ++;};}else {_bca :=data [_fd ];_fd ++;_aa :=8;if _cfg < _aa {_aa =_cfg ;};_be =8-_aa ;_cd =(_cd <<uint (_aa ))|uint32 (_bca >>uint (_be ));if _aa < 8{_gf =_bca <<uint (_aa );};_cfg -=_aa ;if _cfg ==0{_dea =append (_dea ,_cd );_cfg =bitsPerSample ;_cd =0;_ab ++;};};};for _be >=bitsPerSample {_bcg :=_be ;if _cfg < _bcg {_bcg =_cfg ;};_cd =(_cd <<uint (_bcg ))|uint32 (_gf >>uint (8-_bcg ));_be -=_bcg ;if _be > 0{_gf =_gf <<uint (_bcg );}else {_gf =0;};_cfg -=_bcg ;if _cfg ==0{_dea =append (_dea ,_cd );_cfg =bitsPerSample ;_cd =0;_ab ++;};};return _dea ;};func (_f *Reader )ReadSamples (samples []uint32 )(_bb error ){for _cf :=0;_cf < len (samples );_cf ++{samples [_cf ],_bb =_f .ReadSample ();if _bb !=nil {return _bb ;};};return nil ;};func NewReader (img _ae .ImageBase )*Reader {return &Reader {_b :_g .NewReader (img .Data ),_d :img ,_ec :img .ColorComponents ,_c :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };};type SampleWriter interface{WriteSample (_bbd uint32 )error ;WriteSamples (_bf []uint32 )error ;};func ResampleUint32 (data []uint32 ,bitsPerInputSample int ,bitsPerOutputSample int )[]uint32 {var _geb []uint32 ;_ea :=bitsPerOutputSample ;var _fc uint32 ;var _cg uint32 ;_dec :=0;_fe :=0;_af :=0;for _af < len (data ){if _dec > 0{_gc :=_dec ;if _ea < _gc {_gc =_ea ;};_fc =(_fc <<uint (_gc ))|(_cg >>uint (bitsPerInputSample -_gc ));_dec -=_gc ;if _dec > 0{_cg =_cg <<uint (_gc );}else {_cg =0;};_ea -=_gc ;if _ea ==0{_geb =append (_geb ,_fc );_ea =bitsPerOutputSample ;_fc =0;_fe ++;};}else {_gebf :=data [_af ];_af ++;_ef :=bitsPerInputSample ;if _ea < _ef {_ef =_ea ;};_dec =bitsPerInputSample -_ef ;_fc =(_fc <<uint (_ef ))|(_gebf >>uint (_dec ));if _ef < bitsPerInputSample {_cg =_gebf <<uint (_ef );};_ea -=_ef ;if _ea ==0{_geb =append (_geb ,_fc );_ea =bitsPerOutputSample ;_fc =0;_fe ++;};};};for _dec >=bitsPerOutputSample {_ag :=_dec ;if _ea < _ag {_ag =_ea ;};_fc =(_fc <<uint (_ag ))|(_cg >>uint (bitsPerInputSample -_ag ));_dec -=_ag ;if _dec > 0{_cg =_cg <<uint (_ag );}else {_cg =0;};_ea -=_ag ;if _ea ==0{_geb =append (_geb ,_fc );_ea =bitsPerOutputSample ;_fc =0;_fe ++;};};if _ea > 0&&_ea < bitsPerOutputSample {_fc <<=uint (_ea );_geb =append (_geb ,_fc );};return _geb ;};type Writer struct{_bfc _ae .ImageBase ;_gb *_g .Writer ;_dg ,_efd int ;_ed bool ;};func NewWriter (img _ae .ImageBase )*Writer {return &Writer {_gb :_g .NewWriterMSB (img .Data ),_bfc :img ,_efd :img .ColorComponents ,_ed :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };};func (_cfga *Writer )WriteSample (sample uint32 )error {if _ ,_bd :=_cfga ._gb .WriteBits (uint64 (sample ),_cfga ._bfc .BitsPerComponent );_bd !=nil {return _bd ;};_cfga ._efd --;if _cfga ._efd ==0{_cfga ._efd =_cfga ._bfc .ColorComponents ;_cfga ._dg ++;};if _cfga ._dg ==_cfga ._bfc .Width {if _cfga ._ed {_cfga ._gb .FinishByte ();};_cfga ._dg =0;};return nil ;};func (_cgg *Writer )WriteSamples (samples []uint32 )error {for _eab :=0;_eab < len (samples );_eab ++{if _ee :=_cgg .WriteSample (samples [_eab ]);_ee !=nil {return _ee ;};};return nil ;};