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

// Package fdf provides support for loading form field data from Form Field Data (FDF) files.
package fdf ;import (_f "bufio";_df "bytes";_g "encoding/hex";_a "errors";_c "fmt";_cg "github.com/unidoc/unipdf/v3/common";_fc "github.com/unidoc/unipdf/v3/core";_aa "io";_gg "os";_ee "regexp";_ge "sort";_ag "strconv";_d "strings";);func _gea (_baa string )(_fc .PdfObjectReference ,error ){_edb :=_fc .PdfObjectReference {};_fcac :=_ecc .FindStringSubmatch (_baa );if len (_fcac )< 3{_cg .Log .Debug ("\u0045\u0072\u0072or\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065");return _edb ,_a .New ("\u0075n\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061r\u0073e\u0020r\u0065\u0066\u0065\u0072\u0065\u006e\u0063e");};_cee ,_egf :=_ag .Atoi (_fcac [1]);if _egf !=nil {_cg .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0070a\u0072\u0073\u0069n\u0067\u0020\u006fb\u006a\u0065c\u0074\u0020\u006e\u0075\u006d\u0062e\u0072 '\u0025\u0073\u0027\u0020\u002d\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u0075\u006d\u0020\u003d\u0020\u0030",_fcac [1]);return _edb ,nil ;};_edb .ObjectNumber =int64 (_cee );_dgf ,_egf :=_ag .Atoi (_fcac [2]);if _egf !=nil {_cg .Log .Debug ("\u0045\u0072r\u006f\u0072\u0020\u0070\u0061r\u0073\u0069\u006e\u0067\u0020g\u0065\u006e\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0027\u0025\u0073\u0027\u0020\u002d\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u0067\u0065\u006e\u0020\u003d\u0020\u0030",_fcac [2]);return _edb ,nil ;};_edb .GenerationNumber =int64 (_dgf );return _edb ,nil ;};func (_eca *fdfParser )parseArray ()(*_fc .PdfObjectArray ,error ){_daac :=_fc .MakeArray ();_eca ._ecd .ReadByte ();for {_eca .skipSpaces ();_bf ,_acb :=_eca ._ecd .Peek (1);if _acb !=nil {return _daac ,_acb ;};if _bf [0]==']'{_eca ._ecd .ReadByte ();break ;};_eeaa ,_acb :=_eca .parseObject ();if _acb !=nil {return _daac ,_acb ;};_daac .Append (_eeaa );};return _daac ,nil ;};func (_ga *fdfParser )getFileOffset ()int64 {_de ,_ :=_ga ._dfe .Seek (0,_aa .SeekCurrent );_de -=int64 (_ga ._ecd .Buffered ());return _de ;};var _acf =_ee .MustCompile ("\u0025\u0025\u0045O\u0046");func (_eadb *fdfParser )parseFdfVersion ()(int ,int ,error ){_eadb ._dfe .Seek (0,_aa .SeekStart );_cde :=20;_bba :=make ([]byte ,_cde );_eadb ._dfe .Read (_bba );_bbc :=_fec .FindStringSubmatch (string (_bba ));if len (_bbc )< 3{_beb ,_acg ,_eed :=_eadb .seekFdfVersionTopDown ();if _eed !=nil {_cg .Log .Debug ("F\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065\u0063\u006f\u0076\u0065\u0072\u0079\u0020\u002d\u0020\u0075n\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0066\u0069nd\u0020\u0076\u0065r\u0073i\u006f\u006e");return 0,0,_eed ;};return _beb ,_acg ,nil ;};_cef ,_cca :=_ag .Atoi (_bbc [1]);if _cca !=nil {return 0,0,_cca ;};_eeed ,_cca :=_ag .Atoi (_bbc [2]);if _cca !=nil {return 0,0,_cca ;};_cg .Log .Debug ("\u0046\u0064\u0066\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020%\u0064\u002e\u0025\u0064",_cef ,_eeed );return _cef ,_eeed ,nil ;};

// FieldDictionaries returns a map of field names to field dictionaries.
func (fdf *Data )FieldDictionaries ()(map[string ]*_fc .PdfObjectDictionary ,error ){_bd :=map[string ]*_fc .PdfObjectDictionary {};for _bc :=0;_bc < fdf ._eee .Len ();_bc ++{_ff ,_cag :=_fc .GetDict (fdf ._eee .Get (_bc ));if _cag {_gcb ,_ :=_fc .GetString (_ff .Get ("\u0054"));if _gcb !=nil {_bd [_gcb .Str ()]=_ff ;};};};return _bd ,nil ;};func (_aff *fdfParser )readComment ()(string ,error ){var _ae _df .Buffer ;_ ,_da :=_aff .skipSpaces ();if _da !=nil {return _ae .String (),_da ;};_fca :=true ;for {_bga ,_cfd :=_aff ._ecd .Peek (1);if _cfd !=nil {_cg .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073",_cfd .Error ());return _ae .String (),_cfd ;};if _fca &&_bga [0]!='%'{return _ae .String (),_a .New ("c\u006f\u006d\u006d\u0065\u006e\u0074 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0073\u0074a\u0072\u0074\u0020w\u0069t\u0068\u0020\u0025");};_fca =false ;if (_bga [0]!='\r')&&(_bga [0]!='\n'){_fab ,_ :=_aff ._ecd .ReadByte ();_ae .WriteByte (_fab );}else {break ;};};return _ae .String (),nil ;};

// LoadFromPath loads FDF form data from file path `fdfPath`.
func LoadFromPath (fdfPath string )(*Data ,error ){_ca ,_eg :=_gg .Open (fdfPath );if _eg !=nil {return nil ,_eg ;};defer _ca .Close ();return Load (_ca );};func (_bfg *fdfParser )seekFdfVersionTopDown ()(int ,int ,error ){_bfg ._dfe .Seek (0,_aa .SeekStart );_bfg ._ecd =_f .NewReader (_bfg ._dfe );_dcaf :=20;_dcbg :=make ([]byte ,_dcaf );for {_adec ,_gbd :=_bfg ._ecd .ReadByte ();if _gbd !=nil {if _gbd ==_aa .EOF {break ;}else {return 0,0,_gbd ;};};if _fc .IsDecimalDigit (_adec )&&_dcbg [_dcaf -1]=='.'&&_fc .IsDecimalDigit (_dcbg [_dcaf -2])&&_dcbg [_dcaf -3]=='-'&&_dcbg [_dcaf -4]=='F'&&_dcbg [_dcaf -5]=='D'&&_dcbg [_dcaf -6]=='P'{_edee :=int (_dcbg [_dcaf -2]-'0');_cead :=int (_adec -'0');return _edee ,_cead ,nil ;};_dcbg =append (_dcbg [1:_dcaf ],_adec );};return 0,0,_a .New ("\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064");};func _gcac (_ade _aa .ReadSeeker )(*fdfParser ,error ){_cdeaa :=&fdfParser {};_cdeaa ._dfe =_ade ;_cdeaa ._dg =map[int64 ]_fc .PdfObject {};_aag ,_ffb ,_bbdg :=_cdeaa .parseFdfVersion ();if _bbdg !=nil {_cg .Log .Error ("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0076e\u0072\u0073\u0069o\u006e:\u0020\u0025\u0076",_bbdg );return nil ,_bbdg ;};_cdeaa ._feb =_aag ;_cdeaa ._cf =_ffb ;_bbdg =_cdeaa .parse ();return _cdeaa ,_bbdg ;};func (_bac *fdfParser )parseNumber ()(_fc .PdfObject ,error ){return _fc .ParseNumber (_bac ._ecd )};func (_fb *fdfParser )parseObject ()(_fc .PdfObject ,error ){_cg .Log .Trace ("\u0052e\u0061d\u0020\u0064\u0069\u0072\u0065c\u0074\u0020o\u0062\u006a\u0065\u0063\u0074");_fb .skipSpaces ();for {_agda ,_deb :=_fb ._ecd .Peek (2);if _deb !=nil {return nil ,_deb ;};_cg .Log .Trace ("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073",string (_agda ));if _agda [0]=='/'{_gad ,_efa :=_fb .parseName ();_cg .Log .Trace ("\u002d\u003e\u004ea\u006d\u0065\u003a\u0020\u0027\u0025\u0073\u0027",_gad );return &_gad ,_efa ;}else if _agda [0]=='('{_cg .Log .Trace ("\u002d>\u0053\u0074\u0072\u0069\u006e\u0067!");return _fb .parseString ();}else if _agda [0]=='['{_cg .Log .Trace ("\u002d\u003e\u0041\u0072\u0072\u0061\u0079\u0021");return _fb .parseArray ();}else if (_agda [0]=='<')&&(_agda [1]=='<'){_cg .Log .Trace ("\u002d>\u0044\u0069\u0063\u0074\u0021");return _fb .parseDict ();}else if _agda [0]=='<'{_cg .Log .Trace ("\u002d\u003e\u0048\u0065\u0078\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0021");return _fb .parseHexString ();}else if _agda [0]=='%'{_fb .readComment ();_fb .skipSpaces ();}else {_cg .Log .Trace ("\u002d\u003eN\u0075\u006d\u0062e\u0072\u0020\u006f\u0072\u0020\u0072\u0065\u0066\u003f");_agda ,_ =_fb ._ecd .Peek (15);_ffa :=string (_agda );_cg .Log .Trace ("\u0050\u0065\u0065k\u0020\u0073\u0074\u0072\u003a\u0020\u0025\u0073",_ffa );if (len (_ffa )> 3)&&(_ffa [:4]=="\u006e\u0075\u006c\u006c"){_eda ,_bag :=_fb .parseNull ();return &_eda ,_bag ;}else if (len (_ffa )> 4)&&(_ffa [:5]=="\u0066\u0061\u006cs\u0065"){_gee ,_cfg :=_fb .parseBool ();return &_gee ,_cfg ;}else if (len (_ffa )> 3)&&(_ffa [:4]=="\u0074\u0072\u0075\u0065"){_gcda ,_edbc :=_fb .parseBool ();return &_gcda ,_edbc ;};_deda :=_ecc .FindStringSubmatch (_ffa );if len (_deda )> 1{_agda ,_ =_fb ._ecd .ReadBytes ('R');_cg .Log .Trace ("\u002d\u003e\u0020\u0021\u0052\u0065\u0066\u003a\u0020\u0027\u0025\u0073\u0027",string (_agda [:]));_gde ,_dcf :=_gea (string (_agda ));return &_gde ,_dcf ;};_gdbd :=_ba .FindStringSubmatch (_ffa );if len (_gdbd )> 1{_cg .Log .Trace ("\u002d\u003e\u0020\u004e\u0075\u006d\u0062\u0065\u0072\u0021");return _fb .parseNumber ();};_gdbd =_af .FindStringSubmatch (_ffa );if len (_gdbd )> 1{_cg .Log .Trace ("\u002d\u003e\u0020\u0045xp\u006f\u006e\u0065\u006e\u0074\u0069\u0061\u006c\u0020\u004e\u0075\u006d\u0062\u0065r\u0021");_cg .Log .Trace ("\u0025\u0020\u0073",_gdbd );return _fb .parseNumber ();};_cg .Log .Debug ("\u0045R\u0052\u004f\u0052\u0020U\u006e\u006b\u006e\u006f\u0077n\u0020(\u0070e\u0065\u006b\u0020\u0022\u0025\u0073\u0022)",_ffa );return nil ,_a .New ("\u006f\u0062\u006a\u0065\u0063t\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0065\u0072\u0072\u006fr\u0020\u002d\u0020\u0075\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e");};};};func (_ed *fdfParser )readTextLine ()(string ,error ){var _ad _df .Buffer ;for {_eff ,_db :=_ed ._ecd .Peek (1);if _db !=nil {_cg .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073",_db .Error ());return _ad .String (),_db ;};if (_eff [0]!='\r')&&(_eff [0]!='\n'){_ffd ,_ :=_ed ._ecd .ReadByte ();_ad .WriteByte (_ffd );}else {break ;};};return _ad .String (),nil ;};func (_aba *fdfParser )parseNull ()(_fc .PdfObjectNull ,error ){_ ,_ebb :=_aba ._ecd .Discard (4);return _fc .PdfObjectNull {},_ebb ;};func (_cc *fdfParser )setFileOffset (_ece int64 ){_cc ._dfe .Seek (_ece ,_aa .SeekStart );_cc ._ecd =_f .NewReader (_cc ._dfe );};var _af =_ee .MustCompile ("\u005e\u005b\u005c+-\u002e\u005d\u002a\u0028\u005b\u0030\u002d\u0039\u002e]\u002b)\u0065[\u005c+\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d\u0039\u002e\u005d\u002b\u0029");

// Data represents forms data format (FDF) file data.
type Data struct{_agd *_fc .PdfObjectDictionary ;_eee *_fc .PdfObjectArray ;};func (_cab *fdfParser )skipComments ()error {if _ ,_ded :=_cab .skipSpaces ();_ded !=nil {return _ded ;};_ffg :=true ;for {_bgd ,_bb :=_cab ._ecd .Peek (1);if _bb !=nil {_cg .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073",_bb .Error ());return _bb ;};if _ffg &&_bgd [0]!='%'{return nil ;};_ffg =false ;if (_bgd [0]!='\r')&&(_bgd [0]!='\n'){_cab ._ecd .ReadByte ();}else {break ;};};return _cab .skipComments ();};func (_dfb *fdfParser )skipSpaces ()(int ,error ){_dge :=0;for {_ceb ,_cbf :=_dfb ._ecd .ReadByte ();if _cbf !=nil {return 0,_cbf ;};if _fc .IsWhiteSpace (_ceb ){_dge ++;}else {_dfb ._ecd .UnreadByte ();break ;};};return _dge ,nil ;};func _ggfcb (_cbc string )(*fdfParser ,error ){_baed :=fdfParser {};_fdf :=[]byte (_cbc );_gdd :=_df .NewReader (_fdf );_baed ._dfe =_gdd ;_baed ._dg =map[int64 ]_fc .PdfObject {};_efgg :=_f .NewReader (_gdd );_baed ._ecd =_efgg ;_baed ._fa =int64 (len (_cbc ));return &_baed ,_baed .parse ();};var _fec =_ee .MustCompile ("\u0025F\u0044F\u002d\u0028\u005c\u0064\u0029\u005c\u002e\u0028\u005c\u0064\u0029");var _ba =_ee .MustCompile ("\u005e\u005b\u005c\u002b\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d9\u002e\u005d\u002b\u0029");

// FieldValues implements interface model.FieldValueProvider.
// Returns a map of field names to values (PdfObjects).
func (fdf *Data )FieldValues ()(map[string ]_fc .PdfObject ,error ){_gd ,_cb :=fdf .FieldDictionaries ();if _cb !=nil {return nil ,_cb ;};var _bgc []string ;for _dc :=range _gd {_bgc =append (_bgc ,_dc );};_ge .Strings (_bgc );_ea :=map[string ]_fc .PdfObject {};for _ ,_abd :=range _bgc {_ce :=_gd [_abd ];_fd :=_fc .TraceToDirectObject (_ce .Get ("\u0056"));_ea [_abd ]=_fd ;};return _ea ,nil ;};func (_ffc *fdfParser )parseDict ()(*_fc .PdfObjectDictionary ,error ){_cg .Log .Trace ("\u0052\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0046\u0044\u0046\u0020D\u0069\u0063\u0074\u0021");_be :=_fc .MakeDict ();_effe ,_ :=_ffc ._ecd .ReadByte ();if _effe !='<'{return nil ,_a .New ("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074");};_effe ,_ =_ffc ._ecd .ReadByte ();if _effe !='<'{return nil ,_a .New ("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074");};for {_ffc .skipSpaces ();_ffc .skipComments ();_gdg ,_efaf :=_ffc ._ecd .Peek (2);if _efaf !=nil {return nil ,_efaf ;};_cg .Log .Trace ("D\u0069c\u0074\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021",string (_gdg ),string (_gdg ));if (_gdg [0]=='>')&&(_gdg [1]=='>'){_cg .Log .Trace ("\u0045\u004f\u0046\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079");_ffc ._ecd .ReadByte ();_ffc ._ecd .ReadByte ();break ;};_cg .Log .Trace ("\u0050a\u0072s\u0065\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0021");_fabe ,_efaf :=_ffc .parseName ();_cg .Log .Trace ("\u004be\u0079\u003a\u0020\u0025\u0073",_fabe );if _efaf !=nil {_cg .Log .Debug ("E\u0052\u0052\u004f\u0052\u0020\u0052e\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006ea\u006d\u0065\u0020e\u0072r\u0020\u0025\u0073",_efaf );return nil ,_efaf ;};if len (_fabe )> 4&&_fabe [len (_fabe )-4:]=="\u006e\u0075\u006c\u006c"{_bcb :=_fabe [0:len (_fabe )-4];_cg .Log .Debug ("\u0054\u0061\u006b\u0069n\u0067\u0020\u0063\u0061\u0072\u0065\u0020\u006f\u0066\u0020n\u0075l\u006c\u0020\u0062\u0075\u0067\u0020\u0028%\u0073\u0029",_fabe );_cg .Log .Debug ("\u004e\u0065\u0077\u0020ke\u0079\u0020\u0022\u0025\u0073\u0022\u0020\u003d\u0020\u006e\u0075\u006c\u006c",_bcb );_ffc .skipSpaces ();_geed ,_ :=_ffc ._ecd .Peek (1);if _geed [0]=='/'{_be .Set (_bcb ,_fc .MakeNull ());continue ;};};_ffc .skipSpaces ();_cd ,_efaf :=_ffc .parseObject ();if _efaf !=nil {return nil ,_efaf ;};_be .Set (_fabe ,_cd );_cg .Log .Trace ("\u0064\u0069\u0063\u0074\u005b\u0025\u0073\u005d\u0020\u003d\u0020\u0025\u0073",_fabe ,_cd .String ());};_cg .Log .Trace ("\u0072\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0046\u0044\u0046\u0020\u0044\u0069\u0063\u0074\u0021");return _be ,nil ;};func (_ggd *fdfParser )readAtLeast (_cae []byte ,_ead int )(int ,error ){_ggc :=_ead ;_bgcb :=0;_gcd :=0;for _ggc > 0{_ac ,_bdc :=_ggd ._ecd .Read (_cae [_bgcb :]);if _bdc !=nil {_cg .Log .Debug ("\u0045\u0052\u0052O\u0052\u0020\u0046\u0061i\u006c\u0065\u0064\u0020\u0072\u0065\u0061d\u0069\u006e\u0067\u0020\u0028\u0025\u0064\u003b\u0025\u0064\u0029\u0020\u0025\u0073",_ac ,_gcd ,_bdc .Error ());return _bgcb ,_a .New ("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065a\u0064\u0069\u006e\u0067");};_gcd ++;_bgcb +=_ac ;_ggc -=_ac ;};return _bgcb ,nil ;};func (_ege *fdfParser )parse ()error {_ege ._dfe .Seek (0,_aa .SeekStart );_ege ._ecd =_f .NewReader (_ege ._dfe );for {_ege .skipComments ();_egd ,_abe :=_ege ._ecd .Peek (20);if _abe !=nil {_cg .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020r\u0065a\u0064\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a");return _abe ;};if _d .HasPrefix (string (_egd ),"\u0074r\u0061\u0069\u006c\u0065\u0072"){_ege ._ecd .Discard (7);_ege .skipSpaces ();_ege .skipComments ();_dgd ,_ :=_ege .parseDict ();_ege ._cfb =_dgd ;break ;};_bgb :=_bda .FindStringSubmatchIndex (string (_egd ));if len (_bgb )< 6{_cg .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029",string (_egd ));return _a .New ("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065");};_geac ,_abe :=_ege .parseIndirectObject ();if _abe !=nil {return _abe ;};switch _egc :=_geac .(type ){case *_fc .PdfIndirectObject :_ege ._dg [_egc .ObjectNumber ]=_egc ;case *_fc .PdfObjectStream :_ege ._dg [_egc .ObjectNumber ]=_egc ;default:return _a .New ("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");};};return nil ;};type fdfParser struct{_feb int ;_cf int ;_dg map[int64 ]_fc .PdfObject ;_dfe _aa .ReadSeeker ;_ecd *_f .Reader ;_fa int64 ;_cfb *_fc .PdfObjectDictionary ;};func (_aea *fdfParser )parseString ()(*_fc .PdfObjectString ,error ){_aea ._ecd .ReadByte ();var _fed _df .Buffer ;_ecg :=1;for {_fedf ,_bae :=_aea ._ecd .Peek (1);if _bae !=nil {return _fc .MakeString (_fed .String ()),_bae ;};if _fedf [0]=='\\'{_aea ._ecd .ReadByte ();_agg ,_edeg :=_aea ._ecd .ReadByte ();if _edeg !=nil {return _fc .MakeString (_fed .String ()),_edeg ;};if _fc .IsOctalDigit (_agg ){_cbe ,_gdb :=_aea ._ecd .Peek (2);if _gdb !=nil {return _fc .MakeString (_fed .String ()),_gdb ;};var _fdc []byte ;_fdc =append (_fdc ,_agg );for _ ,_abde :=range _cbe {if _fc .IsOctalDigit (_abde ){_fdc =append (_fdc ,_abde );}else {break ;};};_aea ._ecd .Discard (len (_fdc )-1);_cg .Log .Trace ("\u004e\u0075\u006d\u0065ri\u0063\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0022\u0025\u0073\u0022",_fdc );_aad ,_gdb :=_ag .ParseUint (string (_fdc ),8,32);if _gdb !=nil {return _fc .MakeString (_fed .String ()),_gdb ;};_fed .WriteByte (byte (_aad ));continue ;};switch _agg {case 'n':_fed .WriteRune ('\n');case 'r':_fed .WriteRune ('\r');case 't':_fed .WriteRune ('\t');case 'b':_fed .WriteRune ('\b');case 'f':_fed .WriteRune ('\f');case '(':_fed .WriteRune ('(');case ')':_fed .WriteRune (')');case '\\':_fed .WriteRune ('\\');};continue ;}else if _fedf [0]=='('{_ecg ++;}else if _fedf [0]==')'{_ecg --;if _ecg ==0{_aea ._ecd .ReadByte ();break ;};};_fg ,_ :=_aea ._ecd .ReadByte ();_fed .WriteByte (_fg );};return _fc .MakeString (_fed .String ()),nil ;};func (_gaf *fdfParser )parseIndirectObject ()(_fc .PdfObject ,error ){_fbf :=_fc .PdfIndirectObject {};_cg .Log .Trace ("\u002dR\u0065a\u0064\u0020\u0069\u006e\u0064i\u0072\u0065c\u0074\u0020\u006f\u0062\u006a");_dcfg ,_eag :=_gaf ._ecd .Peek (20);if _eag !=nil {_cg .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020r\u0065a\u0064\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a");return &_fbf ,_eag ;};_cg .Log .Trace ("\u0028\u0069\u006edi\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0020\u0070\u0065\u0065\u006b\u0020\u0022\u0025\u0073\u0022",string (_dcfg ));_eaga :=_bda .FindStringSubmatchIndex (string (_dcfg ));if len (_eaga )< 6{_cg .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029",string (_dcfg ));return &_fbf ,_a .New ("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065");};_gaf ._ecd .Discard (_eaga [0]);_cg .Log .Trace ("O\u0066\u0066\u0073\u0065\u0074\u0073\u0020\u0025\u0020\u0064",_eaga );_dbe :=_eaga [1]-_eaga [0];_ged :=make ([]byte ,_dbe );_ ,_eag =_gaf .readAtLeast (_ged ,_dbe );if _eag !=nil {_cg .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0061\u0062l\u0065\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020-\u0020\u0025\u0073",_eag );return nil ,_eag ;};_cg .Log .Trace ("\u0074\u0065\u0078t\u006c\u0069\u006e\u0065\u003a\u0020\u0025\u0073",_ged );_fda :=_bda .FindStringSubmatch (string (_ged ));if len (_fda )< 3{_cg .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029",string (_ged ));return &_fbf ,_a .New ("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065");};_fdcf ,_ :=_ag .Atoi (_fda [1]);_ecb ,_ :=_ag .Atoi (_fda [2]);_fbf .ObjectNumber =int64 (_fdcf );_fbf .GenerationNumber =int64 (_ecb );for {_dfg ,_febb :=_gaf ._ecd .Peek (2);if _febb !=nil {return &_fbf ,_febb ;};_cg .Log .Trace ("I\u006ed\u002e\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021",string (_dfg ),string (_dfg ));if _fc .IsWhiteSpace (_dfg [0]){_gaf .skipSpaces ();}else if _dfg [0]=='%'{_gaf .skipComments ();}else if (_dfg [0]=='<')&&(_dfg [1]=='<'){_cg .Log .Trace ("\u0043\u0061\u006c\u006c\u0020\u0050\u0061\u0072\u0073e\u0044\u0069\u0063\u0074");_fbf .PdfObject ,_febb =_gaf .parseDict ();_cg .Log .Trace ("\u0045\u004f\u0046\u0020Ca\u006c\u006c\u0020\u0050\u0061\u0072\u0073\u0065\u0044\u0069\u0063\u0074\u003a\u0020%\u0076",_febb );if _febb !=nil {return &_fbf ,_febb ;};_cg .Log .Trace ("\u0050\u0061\u0072\u0073\u0065\u0064\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002e.\u002e\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e");}else if (_dfg [0]=='/')||(_dfg [0]=='(')||(_dfg [0]=='[')||(_dfg [0]=='<'){_fbf .PdfObject ,_febb =_gaf .parseObject ();if _febb !=nil {return &_fbf ,_febb ;};_cg .Log .Trace ("P\u0061\u0072\u0073\u0065\u0064\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u002e\u002e\u002e \u0066\u0069\u006ei\u0073h\u0065\u0064\u002e");}else {if _dfg [0]=='e'{_dfc ,_bdb :=_gaf .readTextLine ();if _bdb !=nil {return nil ,_bdb ;};if len (_dfc )>=6&&_dfc [0:6]=="\u0065\u006e\u0064\u006f\u0062\u006a"{break ;};}else if _dfg [0]=='s'{_dfg ,_ =_gaf ._ecd .Peek (10);if string (_dfg [:6])=="\u0073\u0074\u0072\u0065\u0061\u006d"{_cea :=6;if len (_dfg )> 6{if _fc .IsWhiteSpace (_dfg [_cea ])&&_dfg [_cea ]!='\r'&&_dfg [_cea ]!='\n'{_cg .Log .Debug ("\u004e\u006fn\u002d\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0074\u0020\u0046\u0044\u0046\u0020\u006e\u006f\u0074 \u0065\u006e\u0064\u0069\u006e\u0067 \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006c\u0069\u006e\u0065\u0020\u0070\u0072o\u0070\u0065r\u006c\u0079\u0020\u0077i\u0074\u0068\u0020\u0045\u004fL\u0020\u006d\u0061\u0072\u006b\u0065\u0072");_cea ++;};if _dfg [_cea ]=='\r'{_cea ++;if _dfg [_cea ]=='\n'{_cea ++;};}else if _dfg [_cea ]=='\n'{_cea ++;};};_gaf ._ecd .Discard (_cea );_acgg ,_adc :=_fbf .PdfObject .(*_fc .PdfObjectDictionary );if !_adc {return nil ,_a .New ("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006di\u0073s\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079");};_cg .Log .Trace ("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0069c\u0074\u0020\u0025\u0073",_acgg );_fabf ,_gca :=_acgg .Get ("\u004c\u0065\u006e\u0067\u0074\u0068").(*_fc .PdfObjectInteger );if !_gca {return nil ,_a .New ("\u0073\u0074re\u0061\u006d\u0020l\u0065\u006e\u0067\u0074h n\u0065ed\u0073\u0020\u0074\u006f\u0020\u0062\u0065 a\u006e\u0020\u0069\u006e\u0074\u0065\u0067e\u0072");};_eaea :=*_fabf ;if _eaea < 0{return nil ,_a .New ("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006e\u0065\u0065\u0064\u0073\u0020\u0074\u006f \u0062e\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0030");};if int64 (_eaea )> _gaf ._fa {_cg .Log .Debug ("\u0045\u0052R\u004f\u0052\u003a\u0020\u0053t\u0072\u0065\u0061\u006d\u0020l\u0065\u006e\u0067\u0074\u0068\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0069\u007a\u0065");return nil ,_a .New ("\u0069n\u0076\u0061l\u0069\u0064\u0020\u0073t\u0072\u0065\u0061m\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u002c\u0020la\u0072\u0067\u0065r\u0020\u0074h\u0061\u006e\u0020\u0066\u0069\u006ce\u0020\u0073i\u007a\u0065");};_efea :=make ([]byte ,_eaea );_ ,_febb =_gaf .readAtLeast (_efea ,int (_eaea ));if _febb !=nil {_cg .Log .Debug ("E\u0052\u0052\u004f\u0052 s\u0074r\u0065\u0061\u006d\u0020\u0028%\u0064\u0029\u003a\u0020\u0025\u0058",len (_efea ),_efea );_cg .Log .Debug ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v",_febb );return nil ,_febb ;};_ggdg :=_fc .PdfObjectStream {};_ggdg .Stream =_efea ;_ggdg .PdfObjectDictionary =_fbf .PdfObject .(*_fc .PdfObjectDictionary );_ggdg .ObjectNumber =_fbf .ObjectNumber ;_ggdg .GenerationNumber =_fbf .GenerationNumber ;_gaf .skipSpaces ();_gaf ._ecd .Discard (9);_gaf .skipSpaces ();return &_ggdg ,nil ;};};_fbf .PdfObject ,_febb =_gaf .parseObject ();return &_fbf ,_febb ;};};_cg .Log .Trace ("\u0052\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0021");return &_fbf ,nil ;};func (_bbafb *fdfParser )trace (_bgdg _fc .PdfObject )_fc .PdfObject {switch _dcb :=_bgdg .(type ){case *_fc .PdfObjectReference :_effbg ,_fabfd :=_bbafb ._dg [_dcb .ObjectNumber ].(*_fc .PdfIndirectObject );if _fabfd {return _effbg .PdfObject ;};_cg .Log .Debug ("\u0054\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");return nil ;case *_fc .PdfIndirectObject :return _dcb .PdfObject ;};return _bgdg ;};func (_bbdf *fdfParser )parseBool ()(_fc .PdfObjectBool ,error ){_dcd ,_eae :=_bbdf ._ecd .Peek (4);if _eae !=nil {return _fc .PdfObjectBool (false ),_eae ;};if (len (_dcd )>=4)&&(string (_dcd [:4])=="\u0074\u0072\u0075\u0065"){_bbdf ._ecd .Discard (4);return _fc .PdfObjectBool (true ),nil ;};_dcd ,_eae =_bbdf ._ecd .Peek (5);if _eae !=nil {return _fc .PdfObjectBool (false ),_eae ;};if (len (_dcd )>=5)&&(string (_dcd [:5])=="\u0066\u0061\u006cs\u0065"){_bbdf ._ecd .Discard (5);return _fc .PdfObjectBool (false ),nil ;};return _fc .PdfObjectBool (false ),_a .New ("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg");};

// Root returns the Root of the FDF document.
func (_bfeg *fdfParser )Root ()(*_fc .PdfObjectDictionary ,error ){if _bfeg ._cfb !=nil {if _gef ,_cdea :=_bfeg .trace (_bfeg ._cfb .Get ("\u0052\u006f\u006f\u0074")).(*_fc .PdfObjectDictionary );_cdea {if _efb ,_beg :=_bfeg .trace (_gef .Get ("\u0046\u0044\u0046")).(*_fc .PdfObjectDictionary );_beg {return _efb ,nil ;};};};var _gce []int64 ;for _dca :=range _bfeg ._dg {_gce =append (_gce ,_dca );};_ge .Slice (_gce ,func (_aac ,_fcb int )bool {return _gce [_aac ]< _gce [_fcb ]});for _ ,_ddb :=range _gce {_bfb :=_bfeg ._dg [_ddb ];if _fdd ,_eaeg :=_bfeg .trace (_bfb ).(*_fc .PdfObjectDictionary );_eaeg {if _gbc ,_effb :=_bfeg .trace (_fdd .Get ("\u0046\u0044\u0046")).(*_fc .PdfObjectDictionary );_effb {return _gbc ,nil ;};};};return nil ,_a .New ("\u0046\u0044\u0046\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064");};

// Load loads FDF form data from `r`.
func Load (r _aa .ReadSeeker )(*Data ,error ){_gf ,_gc :=_gcac (r );if _gc !=nil {return nil ,_gc ;};_ggf ,_gc :=_gf .Root ();if _gc !=nil {return nil ,_gc ;};_b ,_fe :=_fc .GetArray (_ggf .Get ("\u0046\u0069\u0065\u006c\u0064\u0073"));if !_fe {return nil ,_a .New ("\u0066\u0069\u0065\u006c\u0064\u0073\u0020\u006d\u0069s\u0073\u0069\u006e\u0067");};return &Data {_eee :_b ,_agd :_ggf },nil ;};func (_ccc *fdfParser )seekToEOFMarker (_dga int64 )error {_ced :=int64 (0);_gdbdd :=int64 (1000);for _ced < _dga {if _dga <=(_gdbdd +_ced ){_gdbdd =_dga -_ced ;};_ ,_efed :=_ccc ._dfe .Seek (-_ced -_gdbdd ,_aa .SeekEnd );if _efed !=nil {return _efed ;};_edeb :=make ([]byte ,_gdbdd );_ccc ._dfe .Read (_edeb );_cg .Log .Trace ("\u004c\u006f\u006f\u006bi\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0045\u004f\u0046 \u006da\u0072\u006b\u0065\u0072\u003a\u0020\u0022%\u0073\u0022",string (_edeb ));_ega :=_acf .FindAllStringIndex (string (_edeb ),-1);if _ega !=nil {_fcc :=_ega [len (_ega )-1];_cg .Log .Trace ("\u0049\u006e\u0064\u003a\u0020\u0025\u0020\u0064",_ega );_ccc ._dfe .Seek (-_ced -_gdbdd +int64 (_fcc [0]),_aa .SeekEnd );return nil ;};_cg .Log .Debug ("\u0057\u0061\u0072\u006e\u0069\u006eg\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064\u0021\u0020\u002d\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020s\u0065e\u006b\u0069\u006e\u0067");_ced +=_gdbdd ;};_cg .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006be\u0072 \u0077\u0061\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e");return _a .New ("\u0045\u004f\u0046\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064");};func (_daa *fdfParser )parseName ()(_fc .PdfObjectName ,error ){var _efe _df .Buffer ;_bbd :=false ;for {_effg ,_gcbe :=_daa ._ecd .Peek (1);if _gcbe ==_aa .EOF {break ;};if _gcbe !=nil {return _fc .PdfObjectName (_efe .String ()),_gcbe ;};if !_bbd {if _effg [0]=='/'{_bbd =true ;_daa ._ecd .ReadByte ();}else if _effg [0]=='%'{_daa .readComment ();_daa .skipSpaces ();}else {_cg .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020N\u0061\u006d\u0065\u0020\u0073\u0074\u0061\u0072\u0074\u0069\u006e\u0067\u0020w\u0069\u0074\u0068\u0020\u0025\u0073\u0020(\u0025\u0020\u0078\u0029",_effg ,_effg );return _fc .PdfObjectName (_efe .String ()),_c .Errorf ("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006ea\u006d\u0065:\u0020\u0028\u0025\u0063\u0029",_effg [0]);};}else {if _fc .IsWhiteSpace (_effg [0]){break ;}else if (_effg [0]=='/')||(_effg [0]=='[')||(_effg [0]=='(')||(_effg [0]==']')||(_effg [0]=='<')||(_effg [0]=='>'){break ;}else if _effg [0]=='#'{_ede ,_deg :=_daa ._ecd .Peek (3);if _deg !=nil {return _fc .PdfObjectName (_efe .String ()),_deg ;};_daa ._ecd .Discard (3);_efg ,_deg :=_g .DecodeString (string (_ede [1:3]));if _deg !=nil {return _fc .PdfObjectName (_efe .String ()),_deg ;};_efe .Write (_efg );}else {_afe ,_ :=_daa ._ecd .ReadByte ();_efe .WriteByte (_afe );};};};return _fc .PdfObjectName (_efe .String ()),nil ;};var _ecc =_ee .MustCompile ("^\u005c\u0073\u002a\u0028\\d\u002b)\u005c\u0073\u002b\u0028\u005cd\u002b\u0029\u005c\u0073\u002b\u0052");var _bda =_ee .MustCompile ("\u0028\u005c\u0064\u002b)\\\u0073\u002b\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002b\u006f\u0062\u006a");func (_gab *fdfParser )parseHexString ()(*_fc .PdfObjectString ,error ){_gab ._ecd .ReadByte ();var _fdb _df .Buffer ;for {_ecf ,_fef :=_gab ._ecd .Peek (1);if _fef !=nil {return _fc .MakeHexString (""),_fef ;};if _ecf [0]=='>'{_gab ._ecd .ReadByte ();break ;};_agge ,_ :=_gab ._ecd .ReadByte ();if !_fc .IsWhiteSpace (_agge ){_fdb .WriteByte (_agge );};};if _fdb .Len ()%2==1{_fdb .WriteRune ('0');};_dgc ,_ggfg :=_g .DecodeString (_fdb .String ());if _ggfg !=nil {_cg .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0068\u0065\u0078\u0020\u0073\u0074r\u0069\u006e\u0067\u003a\u0020\u0027\u0025\u0073\u0027 \u002d\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u0061n\u0020\u0065\u006d\u0070\u0074\u0079 \u0073\u0074\u0072i\u006e\u0067",_fdb .String ());return _fc .MakeHexString (""),nil ;};return _fc .MakeHexString (string (_dgc )),nil ;};