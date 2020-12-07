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

package huffman ;import (_ac "errors";_a "fmt";_fb "github.com/unidoc/unipdf/v3/internal/bitwise";_c "github.com/unidoc/unipdf/v3/internal/jbig2/internal";_gd "math";_g "strings";);func (_aca *FixedSizeTable )RootNode ()*InternalNode {return _aca ._ce };type Node interface{Decode (_cfg _fb .StreamReader )(int64 ,error );String ()string ;};var _ Tabler =&EncodedTable {};type InternalNode struct{_cfb int32 ;_fbb Node ;_cbgc Node ;};func NewEncodedTable (table BasicTabler )(*EncodedTable ,error ){_ab :=&EncodedTable {_b :&InternalNode {},BasicTabler :table };if _gb :=_ab .parseTable ();_gb !=nil {return nil ,_gb ;};return _ab ,nil ;};type Code struct{_dg int32 ;_bef int32 ;_gad int32 ;_cgd bool ;_caf int32 ;};func _age (_agb int32 )*InternalNode {return &InternalNode {_cfb :_agb }};func (_da *EncodedTable )InitTree (codeTable []*Code )error {_fca (codeTable );for _ ,_dc :=range codeTable {if _db :=_da ._b .append (_dc );_db !=nil {return _db ;};};return nil ;};func (_gdg *EncodedTable )Decode (r _fb .StreamReader )(int64 ,error ){return _gdg ._b .Decode (r )};func NewFixedSizeTable (codeTable []*Code )(*FixedSizeTable ,error ){_bg :=&FixedSizeTable {_ce :&InternalNode {}};if _afd :=_bg .InitTree (codeTable );_afd !=nil {return nil ,_afd ;};return _bg ,nil ;};func (_df *InternalNode )Decode (r _fb .StreamReader )(int64 ,error ){_cef ,_abe :=r .ReadBit ();if _abe !=nil {return 0,_abe ;};if _cef ==0{return _df ._fbb .Decode (r );};return _df ._cbgc .Decode (r );};func (_ba *OutOfBandNode )Decode (r _fb .StreamReader )(int64 ,error ){return 0,_c .ErrOOB };func (_ad *StandardTable )String ()string {return _ad ._bf .String ()+"\u000a"};func (_cd *OutOfBandNode )String ()string {return _a .Sprintf ("\u0025\u0030\u00364\u0062",int64 (_gd .MaxInt64 ));};func (_cba *ValueNode )String ()string {return _a .Sprintf ("\u0025\u0064\u002f%\u0064",_cba ._afc ,_cba ._ggbf );};func (_cfge *InternalNode )pad (_gda *_g .Builder ){for _fcg :=int32 (0);_fcg < _cfge ._cfb ;_fcg ++{_gda .WriteString ("\u0020\u0020\u0020");};};func (_cg *StandardTable )InitTree (codeTable []*Code )error {_fca (codeTable );for _ ,_bgf :=range codeTable {if _gf :=_cg ._bf .append (_bgf );_gf !=nil {return _gf ;};};return nil ;};func (_dbg *InternalNode )append (_bb *Code )(_de error ){if _bb ._dg ==0{return nil ;};_dec :=_bb ._dg -1-_dbg ._cfb ;if _dec < 0{return _ac .New ("\u004e\u0065\u0067\u0061\u0074\u0069\u0076\u0065\u0020\u0073\u0068\u0069\u0066\u0074\u0069n\u0067 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006c\u006c\u006f\u0077\u0065\u0064");};_ddd :=(_bb ._caf >>uint (_dec ))&0x1;if _dec ==0{if _bb ._bef ==-1{if _ddd ==1{if _dbg ._cbgc !=nil {return _a .Errorf ("O\u004f\u0042\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0073\u0065\u0074\u0020\u0066o\u0072\u0020\u0063o\u0064e\u0020\u0025\u0073",_bb );};_dbg ._cbgc =_cda (_bb );}else {if _dbg ._fbb !=nil {return _a .Errorf ("O\u004f\u0042\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0073\u0065\u0074\u0020\u0066o\u0072\u0020\u0063o\u0064e\u0020\u0025\u0073",_bb );};_dbg ._fbb =_cda (_bb );};}else {if _ddd ==1{if _dbg ._cbgc !=nil {return _a .Errorf ("\u0056\u0061\u006cue\u0020\u004e\u006f\u0064\u0065\u0020\u0061\u006c\u0072e\u0061d\u0079 \u0073e\u0074\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0025\u0073",_bb );};_dbg ._cbgc =_ca (_bb );}else {if _dbg ._fbb !=nil {return _a .Errorf ("\u0056\u0061\u006cue\u0020\u004e\u006f\u0064\u0065\u0020\u0061\u006c\u0072e\u0061d\u0079 \u0073e\u0074\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0025\u0073",_bb );};_dbg ._fbb =_ca (_bb );};};}else {if _ddd ==1{if _dbg ._cbgc ==nil {_dbg ._cbgc =_age (_dbg ._cfb +1);};if _de =_dbg ._cbgc .(*InternalNode ).append (_bb );_de !=nil {return _de ;};}else {if _dbg ._fbb ==nil {_dbg ._fbb =_age (_dbg ._cfb +1);};if _de =_dbg ._fbb .(*InternalNode ).append (_bb );_de !=nil {return _de ;};};};return nil ;};func (_gg *EncodedTable )parseTable ()error {var (_cf []*Code ;_ggf ,_fg ,_cbg int32 ;_ef uint64 ;_af error ;);_eb :=_gg .StreamReader ();_ebb :=_gg .HtLow ();for _ebb < _gg .HtHigh (){_ef ,_af =_eb .ReadBits (byte (_gg .HtPS ()));if _af !=nil {return _af ;};_ggf =int32 (_ef );_ef ,_af =_eb .ReadBits (byte (_gg .HtRS ()));if _af !=nil {return _af ;};_fg =int32 (_ef );_cf =append (_cf ,NewCode (_ggf ,_fg ,_cbg ,false ));_ebb +=1<<uint (_fg );};_ef ,_af =_eb .ReadBits (byte (_gg .HtPS ()));if _af !=nil {return _af ;};_ggf =int32 (_ef );_fg =32;_cbg =_gg .HtLow ()-1;_cf =append (_cf ,NewCode (_ggf ,_fg ,_cbg ,true ));_ef ,_af =_eb .ReadBits (byte (_gg .HtPS ()));if _af !=nil {return _af ;};_ggf =int32 (_ef );_fg =32;_cbg =_gg .HtHigh ();_cf =append (_cf ,NewCode (_ggf ,_fg ,_cbg ,false ));if _gg .HtOOB ()==1{_ef ,_af =_eb .ReadBits (byte (_gg .HtPS ()));if _af !=nil {return _af ;};_ggf =int32 (_ef );_cf =append (_cf ,NewCode (_ggf ,-1,-1,false ));};if _af =_gg .InitTree (_cf );_af !=nil {return _af ;};return nil ;};type FixedSizeTable struct{_ce *InternalNode };func (_dbb *InternalNode )String ()string {_dd :=&_g .Builder {};_dd .WriteString ("\u000a");_dbb .pad (_dd );_dd .WriteString ("\u0030\u003a\u0020");_dd .WriteString (_dbb ._fbb .String ()+"\u000a");_dbb .pad (_dd );_dd .WriteString ("\u0031\u003a\u0020");_dd .WriteString (_dbb ._cbgc .String ()+"\u000a");return _dd .String ();};func _fca (_dfg []*Code ){var _fd int32 ;for _ ,_bc :=range _dfg {_fd =_dcb (_fd ,_bc ._dg );};_adf :=make ([]int32 ,_fd +1);for _ ,_bag :=range _dfg {_adf [_bag ._dg ]++;};var _fba int32 ;_gac :=make ([]int32 ,len (_adf )+1);_adf [0]=0;for _fde :=int32 (1);_fde <=int32 (len (_adf ));_fde ++{_gac [_fde ]=(_gac [_fde -1]+(_adf [_fde -1]))<<1;_fba =_gac [_fde ];for _ ,_dgf :=range _dfg {if _dgf ._dg ==_fde {_dgf ._caf =_fba ;_fba ++;};};};};type EncodedTable struct{BasicTabler ;_b *InternalNode ;};type Tabler interface{Decode (_ed _fb .StreamReader )(int64 ,error );InitTree (_fgg []*Code )error ;String ()string ;RootNode ()*InternalNode ;};func GetStandardTable (number int )(Tabler ,error ){if number <=0||number > len (_cec ){return nil ,_ac .New ("\u0049n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065");};_ace :=_cec [number -1];if _ace ==nil {var _gdc error ;_ace ,_gdc =_ecf (_fe [number -1]);if _gdc !=nil {return nil ,_gdc ;};_cec [number -1]=_ace ;};return _ace ,nil ;};var _ Node =&OutOfBandNode {};type StandardTable struct{_bf *InternalNode };var _cec =make ([]Tabler ,len (_fe ));func (_ag *FixedSizeTable )InitTree (codeTable []*Code )error {_fca (codeTable );for _ ,_fgf :=range codeTable {_eg :=_ag ._ce .append (_fgf );if _eg !=nil {return _eg ;};};return nil ;};func _dcb (_bde ,_dge int32 )int32 {if _bde > _dge {return _bde ;};return _dge ;};func NewCode (prefixLength ,rangeLength ,rangeLow int32 ,isLowerRange bool )*Code {return &Code {_dg :prefixLength ,_bef :rangeLength ,_gad :rangeLow ,_cgd :isLowerRange ,_caf :-1};};func (_e *EncodedTable )String ()string {return _e ._b .String ()+"\u000a"};func (_afb *ValueNode )Decode (r _fb .StreamReader )(int64 ,error ){_ec ,_bae :=r .ReadBits (byte (_afb ._afc ));if _bae !=nil {return 0,_bae ;};if _afb ._bd {_ec =-_ec ;};return int64 (_afb ._ggbf )+int64 (_ec ),nil ;};func (_faa *StandardTable )Decode (r _fb .StreamReader )(int64 ,error ){return _faa ._bf .Decode (r )};func (_bgcg *Code )String ()string {var _fcf string ;if _bgcg ._caf !=-1{_fcf =_dae (_bgcg ._caf ,_bgcg ._dg );}else {_fcf ="\u003f";};return _a .Sprintf ("%\u0073\u002f\u0025\u0064\u002f\u0025\u0064\u002f\u0025\u0064",_fcf ,_bgcg ._dg ,_bgcg ._bef ,_bgcg ._gad );};type BasicTabler interface{HtHigh ()int32 ;HtLow ()int32 ;StreamReader ()_fb .StreamReader ;HtPS ()int32 ;HtRS ()int32 ;HtOOB ()int32 ;};func _dae (_cae ,_bab int32 )string {var _ee int32 ;_ff :=make ([]rune ,_bab );for _egg :=int32 (1);_egg <=_bab ;_egg ++{_ee =_cae >>uint (_bab -_egg )&1;if _ee !=0{_ff [_egg -1]='1';}else {_ff [_egg -1]='0';};};return string (_ff );};var _ Node =&ValueNode {};func _ecf (_cea [][]int32 )(*StandardTable ,error ){var _gfg []*Code ;for _ebf :=0;_ebf < len (_cea );_ebf ++{_bgc :=_cea [_ebf ][0];_gde :=_cea [_ebf ][1];_efc :=_cea [_ebf ][2];var _be bool ;if len (_cea [_ebf ])> 3{_be =true ;};_gfg =append (_gfg ,NewCode (_bgc ,_gde ,_efc ,_be ));};_ae :=&StandardTable {_bf :_age (0)};if _bgb :=_ae .InitTree (_gfg );_bgb !=nil {return nil ,_bgb ;};return _ae ,nil ;};var _fe =[][][]int32 {{{1,4,0},{2,8,16},{3,16,272},{3,32,65808}},{{1,0,0},{2,0,1},{3,0,2},{4,3,3},{5,6,11},{6,32,75},{6,-1,0}},{{8,8,-256},{1,0,0},{2,0,1},{3,0,2},{4,3,3},{5,6,11},{8,32,-257,999},{7,32,75},{6,-1,0}},{{1,0,1},{2,0,2},{3,0,3},{4,3,4},{5,6,12},{5,32,76}},{{7,8,-255},{1,0,1},{2,0,2},{3,0,3},{4,3,4},{5,6,12},{7,32,-256,999},{6,32,76}},{{5,10,-2048},{4,9,-1024},{4,8,-512},{4,7,-256},{5,6,-128},{5,5,-64},{4,5,-32},{2,7,0},{3,7,128},{3,8,256},{4,9,512},{4,10,1024},{6,32,-2049,999},{6,32,2048}},{{4,9,-1024},{3,8,-512},{4,7,-256},{5,6,-128},{5,5,-64},{4,5,-32},{4,5,0},{5,5,32},{5,6,64},{4,7,128},{3,8,256},{3,9,512},{3,10,1024},{5,32,-1025,999},{5,32,2048}},{{8,3,-15},{9,1,-7},{8,1,-5},{9,0,-3},{7,0,-2},{4,0,-1},{2,1,0},{5,0,2},{6,0,3},{3,4,4},{6,1,20},{4,4,22},{4,5,38},{5,6,70},{5,7,134},{6,7,262},{7,8,390},{6,10,646},{9,32,-16,999},{9,32,1670},{2,-1,0}},{{8,4,-31},{9,2,-15},{8,2,-11},{9,1,-7},{7,1,-5},{4,1,-3},{3,1,-1},{3,1,1},{5,1,3},{6,1,5},{3,5,7},{6,2,39},{4,5,43},{4,6,75},{5,7,139},{5,8,267},{6,8,523},{7,9,779},{6,11,1291},{9,32,-32,999},{9,32,3339},{2,-1,0}},{{7,4,-21},{8,0,-5},{7,0,-4},{5,0,-3},{2,2,-2},{5,0,2},{6,0,3},{7,0,4},{8,0,5},{2,6,6},{5,5,70},{6,5,102},{6,6,134},{6,7,198},{6,8,326},{6,9,582},{6,10,1094},{7,11,2118},{8,32,-22,999},{8,32,4166},{2,-1,0}},{{1,0,1},{2,1,2},{4,0,4},{4,1,5},{5,1,7},{5,2,9},{6,2,13},{7,2,17},{7,3,21},{7,4,29},{7,5,45},{7,6,77},{7,32,141}},{{1,0,1},{2,0,2},{3,1,3},{5,0,5},{5,1,6},{6,1,8},{7,0,10},{7,1,11},{7,2,13},{7,3,17},{7,4,25},{8,5,41},{8,32,73}},{{1,0,1},{3,0,2},{4,0,3},{5,0,4},{4,1,5},{3,3,7},{6,1,15},{6,2,17},{6,3,21},{6,4,29},{6,5,45},{7,6,77},{7,32,141}},{{3,0,-2},{3,0,-1},{1,0,0},{3,0,1},{3,0,2}},{{7,4,-24},{6,2,-8},{5,1,-4},{4,0,-2},{3,0,-1},{1,0,0},{3,0,1},{4,0,2},{5,1,3},{6,2,5},{7,4,9},{7,32,-25,999},{7,32,25}}};func _ca (_ga *Code )*ValueNode {return &ValueNode {_afc :_ga ._bef ,_ggbf :_ga ._gad ,_bd :_ga ._cgd }};type ValueNode struct{_afc int32 ;_ggbf int32 ;_bd bool ;};func (_fc *FixedSizeTable )String ()string {return _fc ._ce .String ()+"\u000a"};func (_fae *StandardTable )RootNode ()*InternalNode {return _fae ._bf };func _cda (_ggb *Code )*OutOfBandNode {return &OutOfBandNode {}};func (_gc *FixedSizeTable )Decode (r _fb .StreamReader )(int64 ,error ){return _gc ._ce .Decode (r )};var _ Node =&InternalNode {};type OutOfBandNode struct{};func (_cb *EncodedTable )RootNode ()*InternalNode {return _cb ._b };