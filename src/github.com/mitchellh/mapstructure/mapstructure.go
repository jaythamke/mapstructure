<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="theme-color" content="#375EAB">

  <title>src/github.com/mitchellh/mapstructure/mapstructure.go - Go Documentation Server</title>

<link type="text/css" rel="stylesheet" href="/lib/godoc/style.css">

<script>window.initFuncs = [];</script>
<script src="/lib/godoc/jquery.js" defer></script>



<script>var goVersion = "go1.14.2";</script>
<script src="/lib/godoc/godocs.js" defer></script>
</head>
<body>

<div id='lowframe' style="position: fixed; bottom: 0; left: 0; height: 0; width: 100%; border-top: thin solid grey; background-color: white; overflow: auto;">
...
</div><!-- #lowframe -->

<div id="topbar" class="wide"><div class="container">
<div class="top-heading" id="heading-wide"><a href="/pkg/">Go Documentation Server</a></div>
<div class="top-heading" id="heading-narrow"><a href="/pkg/">GoDoc</a></div>
<a href="#" id="menu-button"><span id="menu-button-arrow">&#9661;</span></a>
<form method="GET" action="/search">
<div id="menu">

<span class="search-box"><input type="search" id="search" name="q" placeholder="Search" aria-label="Search" required><button type="submit"><span><!-- magnifying glass: --><svg width="24" height="24" viewBox="0 0 24 24"><title>submit search</title><path d="M15.5 14h-.79l-.28-.27C15.41 12.59 16 11.11 16 9.5 16 5.91 13.09 3 9.5 3S3 5.91 3 9.5 5.91 16 9.5 16c1.61 0 3.09-.59 4.23-1.57l.27.28v.79l5 4.99L20.49 19l-4.99-5zm-6 0C7.01 14 5 11.99 5 9.5S7.01 5 9.5 5 14 7.01 14 9.5 11.99 14 9.5 14z"/><path d="M0 0h24v24H0z" fill="none"/></svg></span></button></span>
</div>
</form>

</div></div>



<div id="page" class="wide">
<div class="container">


  <h1>
    Source file
    <a href="/src">src</a>/<a href="/src/github.com">github.com</a>/<a href="/src/github.com/mitchellh">mitchellh</a>/<a href="/src/github.com/mitchellh/mapstructure">mapstructure</a>/<span class="text-muted">mapstructure.go</span>
  </h1>





  <h2>
    Documentation: <a href="/pkg/github.com/mitchellh/mapstructure">github.com/mitchellh/mapstructure</a>
  </h2>



<div id="nav"></div>


<script type='text/javascript'>document.ANALYSIS_DATA = null;</script>
<pre><span id="L1" class="ln">     1&nbsp;&nbsp;</span><span class="comment">// Package mapstructure exposes functionality to convert one arbitrary</span>
<span id="L2" class="ln">     2&nbsp;&nbsp;</span><span class="comment">// Go type into another, typically to convert a map[string]interface{}</span>
<span id="L3" class="ln">     3&nbsp;&nbsp;</span><span class="comment">// into a native Go structure.</span>
<span id="L4" class="ln">     4&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L5" class="ln">     5&nbsp;&nbsp;</span><span class="comment">// The Go structure can be arbitrarily complex, containing slices,</span>
<span id="L6" class="ln">     6&nbsp;&nbsp;</span><span class="comment">// other structs, etc. and the decoder will properly decode nested</span>
<span id="L7" class="ln">     7&nbsp;&nbsp;</span><span class="comment">// maps and so on into the proper structures in the native Go struct.</span>
<span id="L8" class="ln">     8&nbsp;&nbsp;</span><span class="comment">// See the examples to see what the decoder is capable of.</span>
<span id="L9" class="ln">     9&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L10" class="ln">    10&nbsp;&nbsp;</span><span class="comment">// The simplest function to start with is Decode.</span>
<span id="L11" class="ln">    11&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L12" class="ln">    12&nbsp;&nbsp;</span><span class="comment">// Field Tags</span>
<span id="L13" class="ln">    13&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L14" class="ln">    14&nbsp;&nbsp;</span><span class="comment">// When decoding to a struct, mapstructure will use the field name by</span>
<span id="L15" class="ln">    15&nbsp;&nbsp;</span><span class="comment">// default to perform the mapping. For example, if a struct has a field</span>
<span id="L16" class="ln">    16&nbsp;&nbsp;</span><span class="comment">// &#34;Username&#34; then mapstructure will look for a key in the source value</span>
<span id="L17" class="ln">    17&nbsp;&nbsp;</span><span class="comment">// of &#34;username&#34; (case insensitive).</span>
<span id="L18" class="ln">    18&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L19" class="ln">    19&nbsp;&nbsp;</span><span class="comment">//     type User struct {</span>
<span id="L20" class="ln">    20&nbsp;&nbsp;</span><span class="comment">//         Username string</span>
<span id="L21" class="ln">    21&nbsp;&nbsp;</span><span class="comment">//     }</span>
<span id="L22" class="ln">    22&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L23" class="ln">    23&nbsp;&nbsp;</span><span class="comment">// You can change the behavior of mapstructure by using struct tags.</span>
<span id="L24" class="ln">    24&nbsp;&nbsp;</span><span class="comment">// The default struct tag that mapstructure looks for is &#34;mapstructure&#34;</span>
<span id="L25" class="ln">    25&nbsp;&nbsp;</span><span class="comment">// but you can customize it using DecoderConfig.</span>
<span id="L26" class="ln">    26&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L27" class="ln">    27&nbsp;&nbsp;</span><span class="comment">// Renaming Fields</span>
<span id="L28" class="ln">    28&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L29" class="ln">    29&nbsp;&nbsp;</span><span class="comment">// To rename the key that mapstructure looks for, use the &#34;mapstructure&#34;</span>
<span id="L30" class="ln">    30&nbsp;&nbsp;</span><span class="comment">// tag and set a value directly. For example, to change the &#34;username&#34; example</span>
<span id="L31" class="ln">    31&nbsp;&nbsp;</span><span class="comment">// above to &#34;user&#34;:</span>
<span id="L32" class="ln">    32&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L33" class="ln">    33&nbsp;&nbsp;</span><span class="comment">//     type User struct {</span>
<span id="L34" class="ln">    34&nbsp;&nbsp;</span><span class="comment">//         Username string `mapstructure:&#34;user&#34;`</span>
<span id="L35" class="ln">    35&nbsp;&nbsp;</span><span class="comment">//     }</span>
<span id="L36" class="ln">    36&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L37" class="ln">    37&nbsp;&nbsp;</span><span class="comment">// Embedded Structs and Squashing</span>
<span id="L38" class="ln">    38&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L39" class="ln">    39&nbsp;&nbsp;</span><span class="comment">// Embedded structs are treated as if they&#39;re another field with that name.</span>
<span id="L40" class="ln">    40&nbsp;&nbsp;</span><span class="comment">// By default, the two structs below are equivalent when decoding with</span>
<span id="L41" class="ln">    41&nbsp;&nbsp;</span><span class="comment">// mapstructure:</span>
<span id="L42" class="ln">    42&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L43" class="ln">    43&nbsp;&nbsp;</span><span class="comment">//     type Person struct {</span>
<span id="L44" class="ln">    44&nbsp;&nbsp;</span><span class="comment">//         Name string</span>
<span id="L45" class="ln">    45&nbsp;&nbsp;</span><span class="comment">//     }</span>
<span id="L46" class="ln">    46&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L47" class="ln">    47&nbsp;&nbsp;</span><span class="comment">//     type Friend struct {</span>
<span id="L48" class="ln">    48&nbsp;&nbsp;</span><span class="comment">//         Person</span>
<span id="L49" class="ln">    49&nbsp;&nbsp;</span><span class="comment">//     }</span>
<span id="L50" class="ln">    50&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L51" class="ln">    51&nbsp;&nbsp;</span><span class="comment">//     type Friend struct {</span>
<span id="L52" class="ln">    52&nbsp;&nbsp;</span><span class="comment">//         Person Person</span>
<span id="L53" class="ln">    53&nbsp;&nbsp;</span><span class="comment">//     }</span>
<span id="L54" class="ln">    54&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L55" class="ln">    55&nbsp;&nbsp;</span><span class="comment">// This would require an input that looks like below:</span>
<span id="L56" class="ln">    56&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L57" class="ln">    57&nbsp;&nbsp;</span><span class="comment">//     map[string]interface{}{</span>
<span id="L58" class="ln">    58&nbsp;&nbsp;</span><span class="comment">//         &#34;person&#34;: map[string]interface{}{&#34;name&#34;: &#34;alice&#34;},</span>
<span id="L59" class="ln">    59&nbsp;&nbsp;</span><span class="comment">//     }</span>
<span id="L60" class="ln">    60&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L61" class="ln">    61&nbsp;&nbsp;</span><span class="comment">// If your &#34;person&#34; value is NOT nested, then you can append &#34;,squash&#34; to</span>
<span id="L62" class="ln">    62&nbsp;&nbsp;</span><span class="comment">// your tag value and mapstructure will treat it as if the embedded struct</span>
<span id="L63" class="ln">    63&nbsp;&nbsp;</span><span class="comment">// were part of the struct directly. Example:</span>
<span id="L64" class="ln">    64&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L65" class="ln">    65&nbsp;&nbsp;</span><span class="comment">//     type Friend struct {</span>
<span id="L66" class="ln">    66&nbsp;&nbsp;</span><span class="comment">//         Person `mapstructure:&#34;,squash&#34;`</span>
<span id="L67" class="ln">    67&nbsp;&nbsp;</span><span class="comment">//     }</span>
<span id="L68" class="ln">    68&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L69" class="ln">    69&nbsp;&nbsp;</span><span class="comment">// Now the following input would be accepted:</span>
<span id="L70" class="ln">    70&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L71" class="ln">    71&nbsp;&nbsp;</span><span class="comment">//     map[string]interface{}{</span>
<span id="L72" class="ln">    72&nbsp;&nbsp;</span><span class="comment">//         &#34;name&#34;: &#34;alice&#34;,</span>
<span id="L73" class="ln">    73&nbsp;&nbsp;</span><span class="comment">//     }</span>
<span id="L74" class="ln">    74&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L75" class="ln">    75&nbsp;&nbsp;</span><span class="comment">// When decoding from a struct to a map, the squash tag squashes the struct</span>
<span id="L76" class="ln">    76&nbsp;&nbsp;</span><span class="comment">// fields into a single map. Using the example structs from above:</span>
<span id="L77" class="ln">    77&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L78" class="ln">    78&nbsp;&nbsp;</span><span class="comment">//     Friend{Person: Person{Name: &#34;alice&#34;}}</span>
<span id="L79" class="ln">    79&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L80" class="ln">    80&nbsp;&nbsp;</span><span class="comment">// Will be decoded into a map:</span>
<span id="L81" class="ln">    81&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L82" class="ln">    82&nbsp;&nbsp;</span><span class="comment">//     map[string]interface{}{</span>
<span id="L83" class="ln">    83&nbsp;&nbsp;</span><span class="comment">//         &#34;name&#34;: &#34;alice&#34;,</span>
<span id="L84" class="ln">    84&nbsp;&nbsp;</span><span class="comment">//     }</span>
<span id="L85" class="ln">    85&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L86" class="ln">    86&nbsp;&nbsp;</span><span class="comment">// DecoderConfig has a field that changes the behavior of mapstructure</span>
<span id="L87" class="ln">    87&nbsp;&nbsp;</span><span class="comment">// to always squash embedded structs.</span>
<span id="L88" class="ln">    88&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L89" class="ln">    89&nbsp;&nbsp;</span><span class="comment">// Remainder Values</span>
<span id="L90" class="ln">    90&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L91" class="ln">    91&nbsp;&nbsp;</span><span class="comment">// If there are any unmapped keys in the source value, mapstructure by</span>
<span id="L92" class="ln">    92&nbsp;&nbsp;</span><span class="comment">// default will silently ignore them. You can error by setting ErrorUnused</span>
<span id="L93" class="ln">    93&nbsp;&nbsp;</span><span class="comment">// in DecoderConfig. If you&#39;re using Metadata you can also maintain a slice</span>
<span id="L94" class="ln">    94&nbsp;&nbsp;</span><span class="comment">// of the unused keys.</span>
<span id="L95" class="ln">    95&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L96" class="ln">    96&nbsp;&nbsp;</span><span class="comment">// You can also use the &#34;,remain&#34; suffix on your tag to collect all unused</span>
<span id="L97" class="ln">    97&nbsp;&nbsp;</span><span class="comment">// values in a map. The field with this tag MUST be a map type and should</span>
<span id="L98" class="ln">    98&nbsp;&nbsp;</span><span class="comment">// probably be a &#34;map[string]interface{}&#34; or &#34;map[interface{}]interface{}&#34;.</span>
<span id="L99" class="ln">    99&nbsp;&nbsp;</span><span class="comment">// See example below:</span>
<span id="L100" class="ln">   100&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L101" class="ln">   101&nbsp;&nbsp;</span><span class="comment">//     type Friend struct {</span>
<span id="L102" class="ln">   102&nbsp;&nbsp;</span><span class="comment">//         Name  string</span>
<span id="L103" class="ln">   103&nbsp;&nbsp;</span><span class="comment">//         Other map[string]interface{} `mapstructure:&#34;,remain&#34;`</span>
<span id="L104" class="ln">   104&nbsp;&nbsp;</span><span class="comment">//     }</span>
<span id="L105" class="ln">   105&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L106" class="ln">   106&nbsp;&nbsp;</span><span class="comment">// Given the input below, Other would be populated with the other</span>
<span id="L107" class="ln">   107&nbsp;&nbsp;</span><span class="comment">// values that weren&#39;t used (everything but &#34;name&#34;):</span>
<span id="L108" class="ln">   108&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L109" class="ln">   109&nbsp;&nbsp;</span><span class="comment">//     map[string]interface{}{</span>
<span id="L110" class="ln">   110&nbsp;&nbsp;</span><span class="comment">//         &#34;name&#34;:    &#34;bob&#34;,</span>
<span id="L111" class="ln">   111&nbsp;&nbsp;</span><span class="comment">//         &#34;address&#34;: &#34;123 Maple St.&#34;,</span>
<span id="L112" class="ln">   112&nbsp;&nbsp;</span><span class="comment">//     }</span>
<span id="L113" class="ln">   113&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L114" class="ln">   114&nbsp;&nbsp;</span><span class="comment">// Omit Empty Values</span>
<span id="L115" class="ln">   115&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L116" class="ln">   116&nbsp;&nbsp;</span><span class="comment">// When decoding from a struct to any other value, you may use the</span>
<span id="L117" class="ln">   117&nbsp;&nbsp;</span><span class="comment">// &#34;,omitempty&#34; suffix on your tag to omit that value if it equates to</span>
<span id="L118" class="ln">   118&nbsp;&nbsp;</span><span class="comment">// the zero value. The zero value of all types is specified in the Go</span>
<span id="L119" class="ln">   119&nbsp;&nbsp;</span><span class="comment">// specification.</span>
<span id="L120" class="ln">   120&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L121" class="ln">   121&nbsp;&nbsp;</span><span class="comment">// For example, the zero type of a numeric type is zero (&#34;0&#34;). If the struct</span>
<span id="L122" class="ln">   122&nbsp;&nbsp;</span><span class="comment">// field value is zero and a numeric type, the field is empty, and it won&#39;t</span>
<span id="L123" class="ln">   123&nbsp;&nbsp;</span><span class="comment">// be encoded into the destination type.</span>
<span id="L124" class="ln">   124&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L125" class="ln">   125&nbsp;&nbsp;</span><span class="comment">//     type Source {</span>
<span id="L126" class="ln">   126&nbsp;&nbsp;</span><span class="comment">//         Age int `mapstructure:&#34;,omitempty&#34;`</span>
<span id="L127" class="ln">   127&nbsp;&nbsp;</span><span class="comment">//     }</span>
<span id="L128" class="ln">   128&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L129" class="ln">   129&nbsp;&nbsp;</span><span class="comment">// Unexported fields</span>
<span id="L130" class="ln">   130&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L131" class="ln">   131&nbsp;&nbsp;</span><span class="comment">// Since unexported (private) struct fields cannot be set outside the package</span>
<span id="L132" class="ln">   132&nbsp;&nbsp;</span><span class="comment">// where they are defined, the decoder will simply skip them.</span>
<span id="L133" class="ln">   133&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L134" class="ln">   134&nbsp;&nbsp;</span><span class="comment">// For this output type definition:</span>
<span id="L135" class="ln">   135&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L136" class="ln">   136&nbsp;&nbsp;</span><span class="comment">//     type Exported struct {</span>
<span id="L137" class="ln">   137&nbsp;&nbsp;</span><span class="comment">//         private string // this unexported field will be skipped</span>
<span id="L138" class="ln">   138&nbsp;&nbsp;</span><span class="comment">//         Public string</span>
<span id="L139" class="ln">   139&nbsp;&nbsp;</span><span class="comment">//     }</span>
<span id="L140" class="ln">   140&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L141" class="ln">   141&nbsp;&nbsp;</span><span class="comment">// Using this map as input:</span>
<span id="L142" class="ln">   142&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L143" class="ln">   143&nbsp;&nbsp;</span><span class="comment">//     map[string]interface{}{</span>
<span id="L144" class="ln">   144&nbsp;&nbsp;</span><span class="comment">//         &#34;private&#34;: &#34;I will be ignored&#34;,</span>
<span id="L145" class="ln">   145&nbsp;&nbsp;</span><span class="comment">//         &#34;Public&#34;:  &#34;I made it through!&#34;,</span>
<span id="L146" class="ln">   146&nbsp;&nbsp;</span><span class="comment">//     }</span>
<span id="L147" class="ln">   147&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L148" class="ln">   148&nbsp;&nbsp;</span><span class="comment">// The following struct will be decoded:</span>
<span id="L149" class="ln">   149&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L150" class="ln">   150&nbsp;&nbsp;</span><span class="comment">//     type Exported struct {</span>
<span id="L151" class="ln">   151&nbsp;&nbsp;</span><span class="comment">//         private: &#34;&#34; // field is left with an empty string (zero value)</span>
<span id="L152" class="ln">   152&nbsp;&nbsp;</span><span class="comment">//         Public: &#34;I made it through!&#34;</span>
<span id="L153" class="ln">   153&nbsp;&nbsp;</span><span class="comment">//     }</span>
<span id="L154" class="ln">   154&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L155" class="ln">   155&nbsp;&nbsp;</span><span class="comment">// Other Configuration</span>
<span id="L156" class="ln">   156&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L157" class="ln">   157&nbsp;&nbsp;</span><span class="comment">// mapstructure is highly configurable. See the DecoderConfig struct</span>
<span id="L158" class="ln">   158&nbsp;&nbsp;</span><span class="comment">// for other features and options that are supported.</span>
<span id="L159" class="ln">   159&nbsp;&nbsp;</span>package mapstructure
<span id="L160" class="ln">   160&nbsp;&nbsp;</span>
<span id="L161" class="ln">   161&nbsp;&nbsp;</span>import (
<span id="L162" class="ln">   162&nbsp;&nbsp;</span>	&#34;encoding/json&#34;
<span id="L163" class="ln">   163&nbsp;&nbsp;</span>	&#34;errors&#34;
<span id="L164" class="ln">   164&nbsp;&nbsp;</span>	&#34;fmt&#34;
<span id="L165" class="ln">   165&nbsp;&nbsp;</span>	&#34;reflect&#34;
<span id="L166" class="ln">   166&nbsp;&nbsp;</span>	&#34;sort&#34;
<span id="L167" class="ln">   167&nbsp;&nbsp;</span>	&#34;strconv&#34;
<span id="L168" class="ln">   168&nbsp;&nbsp;</span>	&#34;strings&#34;
<span id="L169" class="ln">   169&nbsp;&nbsp;</span>)
<span id="L170" class="ln">   170&nbsp;&nbsp;</span>
<span id="L171" class="ln">   171&nbsp;&nbsp;</span><span class="comment">// DecodeHookFunc is the callback function that can be used for</span>
<span id="L172" class="ln">   172&nbsp;&nbsp;</span><span class="comment">// data transformations. See &#34;DecodeHook&#34; in the DecoderConfig</span>
<span id="L173" class="ln">   173&nbsp;&nbsp;</span><span class="comment">// struct.</span>
<span id="L174" class="ln">   174&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L175" class="ln">   175&nbsp;&nbsp;</span><span class="comment">// The type must be one of DecodeHookFuncType, DecodeHookFuncKind, or</span>
<span id="L176" class="ln">   176&nbsp;&nbsp;</span><span class="comment">// DecodeHookFuncValue.</span>
<span id="L177" class="ln">   177&nbsp;&nbsp;</span><span class="comment">// Values are a superset of Types (Values can return types), and Types are a</span>
<span id="L178" class="ln">   178&nbsp;&nbsp;</span><span class="comment">// superset of Kinds (Types can return Kinds) and are generally a richer thing</span>
<span id="L179" class="ln">   179&nbsp;&nbsp;</span><span class="comment">// to use, but Kinds are simpler if you only need those.</span>
<span id="L180" class="ln">   180&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L181" class="ln">   181&nbsp;&nbsp;</span><span class="comment">// The reason DecodeHookFunc is multi-typed is for backwards compatibility:</span>
<span id="L182" class="ln">   182&nbsp;&nbsp;</span><span class="comment">// we started with Kinds and then realized Types were the better solution,</span>
<span id="L183" class="ln">   183&nbsp;&nbsp;</span><span class="comment">// but have a promise to not break backwards compat so we now support</span>
<span id="L184" class="ln">   184&nbsp;&nbsp;</span><span class="comment">// both.</span>
<span id="L185" class="ln">   185&nbsp;&nbsp;</span>type DecodeHookFunc interface{}
<span id="L186" class="ln">   186&nbsp;&nbsp;</span>
<span id="L187" class="ln">   187&nbsp;&nbsp;</span><span class="comment">// DecodeHookFuncType is a DecodeHookFunc which has complete information about</span>
<span id="L188" class="ln">   188&nbsp;&nbsp;</span><span class="comment">// the source and target types.</span>
<span id="L189" class="ln">   189&nbsp;&nbsp;</span>type DecodeHookFuncType func(reflect.Type, reflect.Type, interface{}) (interface{}, error)
<span id="L190" class="ln">   190&nbsp;&nbsp;</span>
<span id="L191" class="ln">   191&nbsp;&nbsp;</span><span class="comment">// DecodeHookFuncKind is a DecodeHookFunc which knows only the Kinds of the</span>
<span id="L192" class="ln">   192&nbsp;&nbsp;</span><span class="comment">// source and target types.</span>
<span id="L193" class="ln">   193&nbsp;&nbsp;</span>type DecodeHookFuncKind func(reflect.Kind, reflect.Kind, interface{}) (interface{}, error)
<span id="L194" class="ln">   194&nbsp;&nbsp;</span>
<span id="L195" class="ln">   195&nbsp;&nbsp;</span><span class="comment">// DecodeHookFuncRaw is a DecodeHookFunc which has complete access to both the source and target</span>
<span id="L196" class="ln">   196&nbsp;&nbsp;</span><span class="comment">// values.</span>
<span id="L197" class="ln">   197&nbsp;&nbsp;</span>type DecodeHookFuncValue func(from reflect.Value, to reflect.Value) (interface{}, error)
<span id="L198" class="ln">   198&nbsp;&nbsp;</span>
<span id="L199" class="ln">   199&nbsp;&nbsp;</span><span class="comment">// DecoderConfig is the configuration that is used to create a new decoder</span>
<span id="L200" class="ln">   200&nbsp;&nbsp;</span><span class="comment">// and allows customization of various aspects of decoding.</span>
<span id="L201" class="ln">   201&nbsp;&nbsp;</span>type DecoderConfig struct {
<span id="L202" class="ln">   202&nbsp;&nbsp;</span>	<span class="comment">// DecodeHook, if set, will be called before any decoding and any</span>
<span id="L203" class="ln">   203&nbsp;&nbsp;</span>	<span class="comment">// type conversion (if WeaklyTypedInput is on). This lets you modify</span>
<span id="L204" class="ln">   204&nbsp;&nbsp;</span>	<span class="comment">// the values before they&#39;re set down onto the resulting struct. The</span>
<span id="L205" class="ln">   205&nbsp;&nbsp;</span>	<span class="comment">// DecodeHook is called for every map and value in the input. This means</span>
<span id="L206" class="ln">   206&nbsp;&nbsp;</span>	<span class="comment">// that if a struct has embedded fields with squash tags the decode hook</span>
<span id="L207" class="ln">   207&nbsp;&nbsp;</span>	<span class="comment">// is called only once with all of the input data, not once for each</span>
<span id="L208" class="ln">   208&nbsp;&nbsp;</span>	<span class="comment">// embedded struct.</span>
<span id="L209" class="ln">   209&nbsp;&nbsp;</span>	<span class="comment">//</span>
<span id="L210" class="ln">   210&nbsp;&nbsp;</span>	<span class="comment">// If an error is returned, the entire decode will fail with that error.</span>
<span id="L211" class="ln">   211&nbsp;&nbsp;</span>	DecodeHook DecodeHookFunc
<span id="L212" class="ln">   212&nbsp;&nbsp;</span>
<span id="L213" class="ln">   213&nbsp;&nbsp;</span>	<span class="comment">// If ErrorUnused is true, then it is an error for there to exist</span>
<span id="L214" class="ln">   214&nbsp;&nbsp;</span>	<span class="comment">// keys in the original map that were unused in the decoding process</span>
<span id="L215" class="ln">   215&nbsp;&nbsp;</span>	<span class="comment">// (extra keys).</span>
<span id="L216" class="ln">   216&nbsp;&nbsp;</span>	ErrorUnused bool
<span id="L217" class="ln">   217&nbsp;&nbsp;</span>
<span id="L218" class="ln">   218&nbsp;&nbsp;</span>	<span class="comment">// ZeroFields, if set to true, will zero fields before writing them.</span>
<span id="L219" class="ln">   219&nbsp;&nbsp;</span>	<span class="comment">// For example, a map will be emptied before decoded values are put in</span>
<span id="L220" class="ln">   220&nbsp;&nbsp;</span>	<span class="comment">// it. If this is false, a map will be merged.</span>
<span id="L221" class="ln">   221&nbsp;&nbsp;</span>	ZeroFields bool
<span id="L222" class="ln">   222&nbsp;&nbsp;</span>
<span id="L223" class="ln">   223&nbsp;&nbsp;</span>	<span class="comment">// If WeaklyTypedInput is true, the decoder will make the following</span>
<span id="L224" class="ln">   224&nbsp;&nbsp;</span>	<span class="comment">// &#34;weak&#34; conversions:</span>
<span id="L225" class="ln">   225&nbsp;&nbsp;</span>	<span class="comment">//</span>
<span id="L226" class="ln">   226&nbsp;&nbsp;</span>	<span class="comment">//   - bools to string (true = &#34;1&#34;, false = &#34;0&#34;)</span>
<span id="L227" class="ln">   227&nbsp;&nbsp;</span>	<span class="comment">//   - numbers to string (base 10)</span>
<span id="L228" class="ln">   228&nbsp;&nbsp;</span>	<span class="comment">//   - bools to int/uint (true = 1, false = 0)</span>
<span id="L229" class="ln">   229&nbsp;&nbsp;</span>	<span class="comment">//   - strings to int/uint (base implied by prefix)</span>
<span id="L230" class="ln">   230&nbsp;&nbsp;</span>	<span class="comment">//   - int to bool (true if value != 0)</span>
<span id="L231" class="ln">   231&nbsp;&nbsp;</span>	<span class="comment">//   - string to bool (accepts: 1, t, T, TRUE, true, True, 0, f, F,</span>
<span id="L232" class="ln">   232&nbsp;&nbsp;</span>	<span class="comment">//     FALSE, false, False. Anything else is an error)</span>
<span id="L233" class="ln">   233&nbsp;&nbsp;</span>	<span class="comment">//   - empty array = empty map and vice versa</span>
<span id="L234" class="ln">   234&nbsp;&nbsp;</span>	<span class="comment">//   - negative numbers to overflowed uint values (base 10)</span>
<span id="L235" class="ln">   235&nbsp;&nbsp;</span>	<span class="comment">//   - slice of maps to a merged map</span>
<span id="L236" class="ln">   236&nbsp;&nbsp;</span>	<span class="comment">//   - single values are converted to slices if required. Each</span>
<span id="L237" class="ln">   237&nbsp;&nbsp;</span>	<span class="comment">//     element is weakly decoded. For example: &#34;4&#34; can become []int{4}</span>
<span id="L238" class="ln">   238&nbsp;&nbsp;</span>	<span class="comment">//     if the target type is an int slice.</span>
<span id="L239" class="ln">   239&nbsp;&nbsp;</span>	<span class="comment">//</span>
<span id="L240" class="ln">   240&nbsp;&nbsp;</span>	WeaklyTypedInput bool
<span id="L241" class="ln">   241&nbsp;&nbsp;</span>
<span id="L242" class="ln">   242&nbsp;&nbsp;</span>	<span class="comment">// Squash will squash embedded structs.  A squash tag may also be</span>
<span id="L243" class="ln">   243&nbsp;&nbsp;</span>	<span class="comment">// added to an individual struct field using a tag.  For example:</span>
<span id="L244" class="ln">   244&nbsp;&nbsp;</span>	<span class="comment">//</span>
<span id="L245" class="ln">   245&nbsp;&nbsp;</span>	<span class="comment">//  type Parent struct {</span>
<span id="L246" class="ln">   246&nbsp;&nbsp;</span>	<span class="comment">//      Child `mapstructure:&#34;,squash&#34;`</span>
<span id="L247" class="ln">   247&nbsp;&nbsp;</span>	<span class="comment">//  }</span>
<span id="L248" class="ln">   248&nbsp;&nbsp;</span>	Squash bool
<span id="L249" class="ln">   249&nbsp;&nbsp;</span>
<span id="L250" class="ln">   250&nbsp;&nbsp;</span>	<span class="comment">// Metadata is the struct that will contain extra metadata about</span>
<span id="L251" class="ln">   251&nbsp;&nbsp;</span>	<span class="comment">// the decoding. If this is nil, then no metadata will be tracked.</span>
<span id="L252" class="ln">   252&nbsp;&nbsp;</span>	Metadata *Metadata
<span id="L253" class="ln">   253&nbsp;&nbsp;</span>
<span id="L254" class="ln">   254&nbsp;&nbsp;</span>	<span class="comment">// Result is a pointer to the struct that will contain the decoded</span>
<span id="L255" class="ln">   255&nbsp;&nbsp;</span>	<span class="comment">// value.</span>
<span id="L256" class="ln">   256&nbsp;&nbsp;</span>	Result interface{}
<span id="L257" class="ln">   257&nbsp;&nbsp;</span>
<span id="L258" class="ln">   258&nbsp;&nbsp;</span>	<span class="comment">// The tag name that mapstructure reads for field names. This</span>
<span id="L259" class="ln">   259&nbsp;&nbsp;</span>	<span class="comment">// defaults to &#34;mapstructure&#34;</span>
<span id="L260" class="ln">   260&nbsp;&nbsp;</span>	TagName string
<span id="L261" class="ln">   261&nbsp;&nbsp;</span>}
<span id="L262" class="ln">   262&nbsp;&nbsp;</span>
<span id="L263" class="ln">   263&nbsp;&nbsp;</span><span class="comment">// A Decoder takes a raw interface value and turns it into structured</span>
<span id="L264" class="ln">   264&nbsp;&nbsp;</span><span class="comment">// data, keeping track of rich error information along the way in case</span>
<span id="L265" class="ln">   265&nbsp;&nbsp;</span><span class="comment">// anything goes wrong. Unlike the basic top-level Decode method, you can</span>
<span id="L266" class="ln">   266&nbsp;&nbsp;</span><span class="comment">// more finely control how the Decoder behaves using the DecoderConfig</span>
<span id="L267" class="ln">   267&nbsp;&nbsp;</span><span class="comment">// structure. The top-level Decode method is just a convenience that sets</span>
<span id="L268" class="ln">   268&nbsp;&nbsp;</span><span class="comment">// up the most basic Decoder.</span>
<span id="L269" class="ln">   269&nbsp;&nbsp;</span>type Decoder struct {
<span id="L270" class="ln">   270&nbsp;&nbsp;</span>	config *DecoderConfig
<span id="L271" class="ln">   271&nbsp;&nbsp;</span>}
<span id="L272" class="ln">   272&nbsp;&nbsp;</span>
<span id="L273" class="ln">   273&nbsp;&nbsp;</span><span class="comment">// Metadata contains information about decoding a structure that</span>
<span id="L274" class="ln">   274&nbsp;&nbsp;</span><span class="comment">// is tedious or difficult to get otherwise.</span>
<span id="L275" class="ln">   275&nbsp;&nbsp;</span>type Metadata struct {
<span id="L276" class="ln">   276&nbsp;&nbsp;</span>	<span class="comment">// Keys are the keys of the structure which were successfully decoded</span>
<span id="L277" class="ln">   277&nbsp;&nbsp;</span>	Keys []string
<span id="L278" class="ln">   278&nbsp;&nbsp;</span>
<span id="L279" class="ln">   279&nbsp;&nbsp;</span>	<span class="comment">// Unused is a slice of keys that were found in the raw value but</span>
<span id="L280" class="ln">   280&nbsp;&nbsp;</span>	<span class="comment">// weren&#39;t decoded since there was no matching field in the result interface</span>
<span id="L281" class="ln">   281&nbsp;&nbsp;</span>	Unused []string
<span id="L282" class="ln">   282&nbsp;&nbsp;</span>}
<span id="L283" class="ln">   283&nbsp;&nbsp;</span>
<span id="L284" class="ln">   284&nbsp;&nbsp;</span><span class="comment">// Decode takes an input structure and uses reflection to translate it to</span>
<span id="L285" class="ln">   285&nbsp;&nbsp;</span><span class="comment">// the output structure. output must be a pointer to a map or struct.</span>
<span id="L286" class="ln">   286&nbsp;&nbsp;</span>func Decode(input interface{}, output interface{}) error {
<span id="L287" class="ln">   287&nbsp;&nbsp;</span>	config := &amp;DecoderConfig{
<span id="L288" class="ln">   288&nbsp;&nbsp;</span>		Metadata: nil,
<span id="L289" class="ln">   289&nbsp;&nbsp;</span>		Result:   output,
<span id="L290" class="ln">   290&nbsp;&nbsp;</span>	}
<span id="L291" class="ln">   291&nbsp;&nbsp;</span>
<span id="L292" class="ln">   292&nbsp;&nbsp;</span>	decoder, err := NewDecoder(config)
<span id="L293" class="ln">   293&nbsp;&nbsp;</span>	if err != nil {
<span id="L294" class="ln">   294&nbsp;&nbsp;</span>		return err
<span id="L295" class="ln">   295&nbsp;&nbsp;</span>	}
<span id="L296" class="ln">   296&nbsp;&nbsp;</span>
<span id="L297" class="ln">   297&nbsp;&nbsp;</span>	return decoder.Decode(input)
<span id="L298" class="ln">   298&nbsp;&nbsp;</span>}
<span id="L299" class="ln">   299&nbsp;&nbsp;</span>
<span id="L300" class="ln">   300&nbsp;&nbsp;</span><span class="comment">// WeakDecode is the same as Decode but is shorthand to enable</span>
<span id="L301" class="ln">   301&nbsp;&nbsp;</span><span class="comment">// WeaklyTypedInput. See DecoderConfig for more info.</span>
<span id="L302" class="ln">   302&nbsp;&nbsp;</span>func WeakDecode(input, output interface{}) error {
<span id="L303" class="ln">   303&nbsp;&nbsp;</span>	config := &amp;DecoderConfig{
<span id="L304" class="ln">   304&nbsp;&nbsp;</span>		Metadata:         nil,
<span id="L305" class="ln">   305&nbsp;&nbsp;</span>		Result:           output,
<span id="L306" class="ln">   306&nbsp;&nbsp;</span>		WeaklyTypedInput: true,
<span id="L307" class="ln">   307&nbsp;&nbsp;</span>	}
<span id="L308" class="ln">   308&nbsp;&nbsp;</span>
<span id="L309" class="ln">   309&nbsp;&nbsp;</span>	decoder, err := NewDecoder(config)
<span id="L310" class="ln">   310&nbsp;&nbsp;</span>	if err != nil {
<span id="L311" class="ln">   311&nbsp;&nbsp;</span>		return err
<span id="L312" class="ln">   312&nbsp;&nbsp;</span>	}
<span id="L313" class="ln">   313&nbsp;&nbsp;</span>
<span id="L314" class="ln">   314&nbsp;&nbsp;</span>	return decoder.Decode(input)
<span id="L315" class="ln">   315&nbsp;&nbsp;</span>}
<span id="L316" class="ln">   316&nbsp;&nbsp;</span>
<span id="L317" class="ln">   317&nbsp;&nbsp;</span><span class="comment">// DecodeMetadata is the same as Decode, but is shorthand to</span>
<span id="L318" class="ln">   318&nbsp;&nbsp;</span><span class="comment">// enable metadata collection. See DecoderConfig for more info.</span>
<span id="L319" class="ln">   319&nbsp;&nbsp;</span>func DecodeMetadata(input interface{}, output interface{}, metadata *Metadata) error {
<span id="L320" class="ln">   320&nbsp;&nbsp;</span>	config := &amp;DecoderConfig{
<span id="L321" class="ln">   321&nbsp;&nbsp;</span>		Metadata: metadata,
<span id="L322" class="ln">   322&nbsp;&nbsp;</span>		Result:   output,
<span id="L323" class="ln">   323&nbsp;&nbsp;</span>	}
<span id="L324" class="ln">   324&nbsp;&nbsp;</span>
<span id="L325" class="ln">   325&nbsp;&nbsp;</span>	decoder, err := NewDecoder(config)
<span id="L326" class="ln">   326&nbsp;&nbsp;</span>	if err != nil {
<span id="L327" class="ln">   327&nbsp;&nbsp;</span>		return err
<span id="L328" class="ln">   328&nbsp;&nbsp;</span>	}
<span id="L329" class="ln">   329&nbsp;&nbsp;</span>
<span id="L330" class="ln">   330&nbsp;&nbsp;</span>	return decoder.Decode(input)
<span id="L331" class="ln">   331&nbsp;&nbsp;</span>}
<span id="L332" class="ln">   332&nbsp;&nbsp;</span>
<span id="L333" class="ln">   333&nbsp;&nbsp;</span><span class="comment">// WeakDecodeMetadata is the same as Decode, but is shorthand to</span>
<span id="L334" class="ln">   334&nbsp;&nbsp;</span><span class="comment">// enable both WeaklyTypedInput and metadata collection. See</span>
<span id="L335" class="ln">   335&nbsp;&nbsp;</span><span class="comment">// DecoderConfig for more info.</span>
<span id="L336" class="ln">   336&nbsp;&nbsp;</span>func WeakDecodeMetadata(input interface{}, output interface{}, metadata *Metadata) error {
<span id="L337" class="ln">   337&nbsp;&nbsp;</span>	config := &amp;DecoderConfig{
<span id="L338" class="ln">   338&nbsp;&nbsp;</span>		Metadata:         metadata,
<span id="L339" class="ln">   339&nbsp;&nbsp;</span>		Result:           output,
<span id="L340" class="ln">   340&nbsp;&nbsp;</span>		WeaklyTypedInput: true,
<span id="L341" class="ln">   341&nbsp;&nbsp;</span>	}
<span id="L342" class="ln">   342&nbsp;&nbsp;</span>
<span id="L343" class="ln">   343&nbsp;&nbsp;</span>	decoder, err := NewDecoder(config)
<span id="L344" class="ln">   344&nbsp;&nbsp;</span>	if err != nil {
<span id="L345" class="ln">   345&nbsp;&nbsp;</span>		return err
<span id="L346" class="ln">   346&nbsp;&nbsp;</span>	}
<span id="L347" class="ln">   347&nbsp;&nbsp;</span>
<span id="L348" class="ln">   348&nbsp;&nbsp;</span>	return decoder.Decode(input)
<span id="L349" class="ln">   349&nbsp;&nbsp;</span>}
<span id="L350" class="ln">   350&nbsp;&nbsp;</span>
<span id="L351" class="ln">   351&nbsp;&nbsp;</span><span class="comment">// NewDecoder returns a new decoder for the given configuration. Once</span>
<span id="L352" class="ln">   352&nbsp;&nbsp;</span><span class="comment">// a decoder has been returned, the same configuration must not be used</span>
<span id="L353" class="ln">   353&nbsp;&nbsp;</span><span class="comment">// again.</span>
<span id="L354" class="ln">   354&nbsp;&nbsp;</span>func NewDecoder(config *DecoderConfig) (*Decoder, error) {
<span id="L355" class="ln">   355&nbsp;&nbsp;</span>	val := reflect.ValueOf(config.Result)
<span id="L356" class="ln">   356&nbsp;&nbsp;</span>	if val.Kind() != reflect.Ptr {
<span id="L357" class="ln">   357&nbsp;&nbsp;</span>		return nil, errors.New(&#34;result must be a pointer&#34;)
<span id="L358" class="ln">   358&nbsp;&nbsp;</span>	}
<span id="L359" class="ln">   359&nbsp;&nbsp;</span>
<span id="L360" class="ln">   360&nbsp;&nbsp;</span>	val = val.Elem()
<span id="L361" class="ln">   361&nbsp;&nbsp;</span>	if !val.CanAddr() {
<span id="L362" class="ln">   362&nbsp;&nbsp;</span>		return nil, errors.New(&#34;result must be addressable (a pointer)&#34;)
<span id="L363" class="ln">   363&nbsp;&nbsp;</span>	}
<span id="L364" class="ln">   364&nbsp;&nbsp;</span>
<span id="L365" class="ln">   365&nbsp;&nbsp;</span>	if config.Metadata != nil {
<span id="L366" class="ln">   366&nbsp;&nbsp;</span>		if config.Metadata.Keys == nil {
<span id="L367" class="ln">   367&nbsp;&nbsp;</span>			config.Metadata.Keys = make([]string, 0)
<span id="L368" class="ln">   368&nbsp;&nbsp;</span>		}
<span id="L369" class="ln">   369&nbsp;&nbsp;</span>
<span id="L370" class="ln">   370&nbsp;&nbsp;</span>		if config.Metadata.Unused == nil {
<span id="L371" class="ln">   371&nbsp;&nbsp;</span>			config.Metadata.Unused = make([]string, 0)
<span id="L372" class="ln">   372&nbsp;&nbsp;</span>		}
<span id="L373" class="ln">   373&nbsp;&nbsp;</span>	}
<span id="L374" class="ln">   374&nbsp;&nbsp;</span>
<span id="L375" class="ln">   375&nbsp;&nbsp;</span>	if config.TagName == &#34;&#34; {
<span id="L376" class="ln">   376&nbsp;&nbsp;</span>		config.TagName = &#34;mapstructure&#34;
<span id="L377" class="ln">   377&nbsp;&nbsp;</span>	}
<span id="L378" class="ln">   378&nbsp;&nbsp;</span>
<span id="L379" class="ln">   379&nbsp;&nbsp;</span>	result := &amp;Decoder{
<span id="L380" class="ln">   380&nbsp;&nbsp;</span>		config: config,
<span id="L381" class="ln">   381&nbsp;&nbsp;</span>	}
<span id="L382" class="ln">   382&nbsp;&nbsp;</span>
<span id="L383" class="ln">   383&nbsp;&nbsp;</span>	return result, nil
<span id="L384" class="ln">   384&nbsp;&nbsp;</span>}
<span id="L385" class="ln">   385&nbsp;&nbsp;</span>
<span id="L386" class="ln">   386&nbsp;&nbsp;</span><span class="comment">// Decode decodes the given raw interface to the target pointer specified</span>
<span id="L387" class="ln">   387&nbsp;&nbsp;</span><span class="comment">// by the configuration.</span>
<span id="L388" class="ln">   388&nbsp;&nbsp;</span>func (d *Decoder) Decode(input interface{}) error {
<span id="L389" class="ln">   389&nbsp;&nbsp;</span>	return d.decode(&#34;&#34;, input, reflect.ValueOf(d.config.Result).Elem())
<span id="L390" class="ln">   390&nbsp;&nbsp;</span>}
<span id="L391" class="ln">   391&nbsp;&nbsp;</span>
<span id="L392" class="ln">   392&nbsp;&nbsp;</span><span class="comment">// Decodes an unknown data type into a specific reflection value.</span>
<span id="L393" class="ln">   393&nbsp;&nbsp;</span>func (d *Decoder) decode(name string, input interface{}, outVal reflect.Value) error {
<span id="L394" class="ln">   394&nbsp;&nbsp;</span>	var inputVal reflect.Value
<span id="L395" class="ln">   395&nbsp;&nbsp;</span>	if input != nil {
<span id="L396" class="ln">   396&nbsp;&nbsp;</span>		inputVal = reflect.ValueOf(input)
<span id="L397" class="ln">   397&nbsp;&nbsp;</span>
<span id="L398" class="ln">   398&nbsp;&nbsp;</span>		<span class="comment">// We need to check here if input is a typed nil. Typed nils won&#39;t</span>
<span id="L399" class="ln">   399&nbsp;&nbsp;</span>		<span class="comment">// match the &#34;input == nil&#34; below so we check that here.</span>
<span id="L400" class="ln">   400&nbsp;&nbsp;</span>		if inputVal.Kind() == reflect.Ptr &amp;&amp; inputVal.IsNil() {
<span id="L401" class="ln">   401&nbsp;&nbsp;</span>			input = nil
<span id="L402" class="ln">   402&nbsp;&nbsp;</span>		}
<span id="L403" class="ln">   403&nbsp;&nbsp;</span>	}
<span id="L404" class="ln">   404&nbsp;&nbsp;</span>
<span id="L405" class="ln">   405&nbsp;&nbsp;</span>	if input == nil {
<span id="L406" class="ln">   406&nbsp;&nbsp;</span>		<span class="comment">// If the data is nil, then we don&#39;t set anything, unless ZeroFields is set</span>
<span id="L407" class="ln">   407&nbsp;&nbsp;</span>		<span class="comment">// to true.</span>
<span id="L408" class="ln">   408&nbsp;&nbsp;</span>		if d.config.ZeroFields {
<span id="L409" class="ln">   409&nbsp;&nbsp;</span>			outVal.Set(reflect.Zero(outVal.Type()))
<span id="L410" class="ln">   410&nbsp;&nbsp;</span>
<span id="L411" class="ln">   411&nbsp;&nbsp;</span>			if d.config.Metadata != nil &amp;&amp; name != &#34;&#34; {
<span id="L412" class="ln">   412&nbsp;&nbsp;</span>				d.config.Metadata.Keys = append(d.config.Metadata.Keys, name)
<span id="L413" class="ln">   413&nbsp;&nbsp;</span>			}
<span id="L414" class="ln">   414&nbsp;&nbsp;</span>		}
<span id="L415" class="ln">   415&nbsp;&nbsp;</span>		return nil
<span id="L416" class="ln">   416&nbsp;&nbsp;</span>	}
<span id="L417" class="ln">   417&nbsp;&nbsp;</span>
<span id="L418" class="ln">   418&nbsp;&nbsp;</span>	if !inputVal.IsValid() {
<span id="L419" class="ln">   419&nbsp;&nbsp;</span>		<span class="comment">// If the input value is invalid, then we just set the value</span>
<span id="L420" class="ln">   420&nbsp;&nbsp;</span>		<span class="comment">// to be the zero value.</span>
<span id="L421" class="ln">   421&nbsp;&nbsp;</span>		outVal.Set(reflect.Zero(outVal.Type()))
<span id="L422" class="ln">   422&nbsp;&nbsp;</span>		if d.config.Metadata != nil &amp;&amp; name != &#34;&#34; {
<span id="L423" class="ln">   423&nbsp;&nbsp;</span>			d.config.Metadata.Keys = append(d.config.Metadata.Keys, name)
<span id="L424" class="ln">   424&nbsp;&nbsp;</span>		}
<span id="L425" class="ln">   425&nbsp;&nbsp;</span>		return nil
<span id="L426" class="ln">   426&nbsp;&nbsp;</span>	}
<span id="L427" class="ln">   427&nbsp;&nbsp;</span>
<span id="L428" class="ln">   428&nbsp;&nbsp;</span>	if d.config.DecodeHook != nil {
<span id="L429" class="ln">   429&nbsp;&nbsp;</span>		<span class="comment">// We have a DecodeHook, so let&#39;s pre-process the input.</span>
<span id="L430" class="ln">   430&nbsp;&nbsp;</span>		var err error
<span id="L431" class="ln">   431&nbsp;&nbsp;</span>		input, err = DecodeHookExec(d.config.DecodeHook, inputVal, outVal)
<span id="L432" class="ln">   432&nbsp;&nbsp;</span>		if err != nil {
<span id="L433" class="ln">   433&nbsp;&nbsp;</span>			return fmt.Errorf(&#34;error decoding &#39;%s&#39;: %s&#34;, name, err)
<span id="L434" class="ln">   434&nbsp;&nbsp;</span>		}
<span id="L435" class="ln">   435&nbsp;&nbsp;</span>	}
<span id="L436" class="ln">   436&nbsp;&nbsp;</span>
<span id="L437" class="ln">   437&nbsp;&nbsp;</span>	var err error
<span id="L438" class="ln">   438&nbsp;&nbsp;</span>	outputKind := getKind(outVal)
<span id="L439" class="ln">   439&nbsp;&nbsp;</span>	addMetaKey := true
<span id="L440" class="ln">   440&nbsp;&nbsp;</span>	switch outputKind {
<span id="L441" class="ln">   441&nbsp;&nbsp;</span>	case reflect.Bool:
<span id="L442" class="ln">   442&nbsp;&nbsp;</span>		err = d.decodeBool(name, input, outVal)
<span id="L443" class="ln">   443&nbsp;&nbsp;</span>	case reflect.Interface:
<span id="L444" class="ln">   444&nbsp;&nbsp;</span>		err = d.decodeBasic(name, input, outVal)
<span id="L445" class="ln">   445&nbsp;&nbsp;</span>	case reflect.String:
<span id="L446" class="ln">   446&nbsp;&nbsp;</span>		err = d.decodeString(name, input, outVal)
<span id="L447" class="ln">   447&nbsp;&nbsp;</span>	case reflect.Int:
<span id="L448" class="ln">   448&nbsp;&nbsp;</span>		err = d.decodeInt(name, input, outVal)
<span id="L449" class="ln">   449&nbsp;&nbsp;</span>	case reflect.Uint:
<span id="L450" class="ln">   450&nbsp;&nbsp;</span>		err = d.decodeUint(name, input, outVal)
<span id="L451" class="ln">   451&nbsp;&nbsp;</span>	case reflect.Float32:
<span id="L452" class="ln">   452&nbsp;&nbsp;</span>		err = d.decodeFloat(name, input, outVal)
<span id="L453" class="ln">   453&nbsp;&nbsp;</span>	case reflect.Struct:
<span id="L454" class="ln">   454&nbsp;&nbsp;</span>		err = d.decodeStruct(name, input, outVal)
<span id="L455" class="ln">   455&nbsp;&nbsp;</span>	case reflect.Map:
<span id="L456" class="ln">   456&nbsp;&nbsp;</span>		err = d.decodeMap(name, input, outVal)
<span id="L457" class="ln">   457&nbsp;&nbsp;</span>	case reflect.Ptr:
<span id="L458" class="ln">   458&nbsp;&nbsp;</span>		addMetaKey, err = d.decodePtr(name, input, outVal)
<span id="L459" class="ln">   459&nbsp;&nbsp;</span>	case reflect.Slice:
<span id="L460" class="ln">   460&nbsp;&nbsp;</span>		err = d.decodeSlice(name, input, outVal)
<span id="L461" class="ln">   461&nbsp;&nbsp;</span>	case reflect.Array:
<span id="L462" class="ln">   462&nbsp;&nbsp;</span>		err = d.decodeArray(name, input, outVal)
<span id="L463" class="ln">   463&nbsp;&nbsp;</span>	case reflect.Func:
<span id="L464" class="ln">   464&nbsp;&nbsp;</span>		err = d.decodeFunc(name, input, outVal)
<span id="L465" class="ln">   465&nbsp;&nbsp;</span>	default:
<span id="L466" class="ln">   466&nbsp;&nbsp;</span>		<span class="comment">// If we reached this point then we weren&#39;t able to decode it</span>
<span id="L467" class="ln">   467&nbsp;&nbsp;</span>		return fmt.Errorf(&#34;%s: unsupported type: %s&#34;, name, outputKind)
<span id="L468" class="ln">   468&nbsp;&nbsp;</span>	}
<span id="L469" class="ln">   469&nbsp;&nbsp;</span>
<span id="L470" class="ln">   470&nbsp;&nbsp;</span>	<span class="comment">// If we reached here, then we successfully decoded SOMETHING, so</span>
<span id="L471" class="ln">   471&nbsp;&nbsp;</span>	<span class="comment">// mark the key as used if we&#39;re tracking metainput.</span>
<span id="L472" class="ln">   472&nbsp;&nbsp;</span>	if addMetaKey &amp;&amp; d.config.Metadata != nil &amp;&amp; name != &#34;&#34; {
<span id="L473" class="ln">   473&nbsp;&nbsp;</span>		d.config.Metadata.Keys = append(d.config.Metadata.Keys, name)
<span id="L474" class="ln">   474&nbsp;&nbsp;</span>	}
<span id="L475" class="ln">   475&nbsp;&nbsp;</span>
<span id="L476" class="ln">   476&nbsp;&nbsp;</span>	return err
<span id="L477" class="ln">   477&nbsp;&nbsp;</span>}
<span id="L478" class="ln">   478&nbsp;&nbsp;</span>
<span id="L479" class="ln">   479&nbsp;&nbsp;</span><span class="comment">// This decodes a basic type (bool, int, string, etc.) and sets the</span>
<span id="L480" class="ln">   480&nbsp;&nbsp;</span><span class="comment">// value to &#34;data&#34; of that type.</span>
<span id="L481" class="ln">   481&nbsp;&nbsp;</span>func (d *Decoder) decodeBasic(name string, data interface{}, val reflect.Value) error {
<span id="L482" class="ln">   482&nbsp;&nbsp;</span>	if val.IsValid() &amp;&amp; val.Elem().IsValid() {
<span id="L483" class="ln">   483&nbsp;&nbsp;</span>		elem := val.Elem()
<span id="L484" class="ln">   484&nbsp;&nbsp;</span>
<span id="L485" class="ln">   485&nbsp;&nbsp;</span>		<span class="comment">// If we can&#39;t address this element, then its not writable. Instead,</span>
<span id="L486" class="ln">   486&nbsp;&nbsp;</span>		<span class="comment">// we make a copy of the value (which is a pointer and therefore</span>
<span id="L487" class="ln">   487&nbsp;&nbsp;</span>		<span class="comment">// writable), decode into that, and replace the whole value.</span>
<span id="L488" class="ln">   488&nbsp;&nbsp;</span>		copied := false
<span id="L489" class="ln">   489&nbsp;&nbsp;</span>		if !elem.CanAddr() {
<span id="L490" class="ln">   490&nbsp;&nbsp;</span>			copied = true
<span id="L491" class="ln">   491&nbsp;&nbsp;</span>
<span id="L492" class="ln">   492&nbsp;&nbsp;</span>			<span class="comment">// Make *T</span>
<span id="L493" class="ln">   493&nbsp;&nbsp;</span>			copy := reflect.New(elem.Type())
<span id="L494" class="ln">   494&nbsp;&nbsp;</span>
<span id="L495" class="ln">   495&nbsp;&nbsp;</span>			<span class="comment">// *T = elem</span>
<span id="L496" class="ln">   496&nbsp;&nbsp;</span>			copy.Elem().Set(elem)
<span id="L497" class="ln">   497&nbsp;&nbsp;</span>
<span id="L498" class="ln">   498&nbsp;&nbsp;</span>			<span class="comment">// Set elem so we decode into it</span>
<span id="L499" class="ln">   499&nbsp;&nbsp;</span>			elem = copy
<span id="L500" class="ln">   500&nbsp;&nbsp;</span>		}
<span id="L501" class="ln">   501&nbsp;&nbsp;</span>
<span id="L502" class="ln">   502&nbsp;&nbsp;</span>		<span class="comment">// Decode. If we have an error then return. We also return right</span>
<span id="L503" class="ln">   503&nbsp;&nbsp;</span>		<span class="comment">// away if we&#39;re not a copy because that means we decoded directly.</span>
<span id="L504" class="ln">   504&nbsp;&nbsp;</span>		if err := d.decode(name, data, elem); err != nil || !copied {
<span id="L505" class="ln">   505&nbsp;&nbsp;</span>			return err
<span id="L506" class="ln">   506&nbsp;&nbsp;</span>		}
<span id="L507" class="ln">   507&nbsp;&nbsp;</span>
<span id="L508" class="ln">   508&nbsp;&nbsp;</span>		<span class="comment">// If we&#39;re a copy, we need to set te final result</span>
<span id="L509" class="ln">   509&nbsp;&nbsp;</span>		val.Set(elem.Elem())
<span id="L510" class="ln">   510&nbsp;&nbsp;</span>		return nil
<span id="L511" class="ln">   511&nbsp;&nbsp;</span>	}
<span id="L512" class="ln">   512&nbsp;&nbsp;</span>
<span id="L513" class="ln">   513&nbsp;&nbsp;</span>	dataVal := reflect.ValueOf(data)
<span id="L514" class="ln">   514&nbsp;&nbsp;</span>
<span id="L515" class="ln">   515&nbsp;&nbsp;</span>	<span class="comment">// If the input data is a pointer, and the assigned type is the dereference</span>
<span id="L516" class="ln">   516&nbsp;&nbsp;</span>	<span class="comment">// of that exact pointer, then indirect it so that we can assign it.</span>
<span id="L517" class="ln">   517&nbsp;&nbsp;</span>	<span class="comment">// Example: *string to string</span>
<span id="L518" class="ln">   518&nbsp;&nbsp;</span>	if dataVal.Kind() == reflect.Ptr &amp;&amp; dataVal.Type().Elem() == val.Type() {
<span id="L519" class="ln">   519&nbsp;&nbsp;</span>		dataVal = reflect.Indirect(dataVal)
<span id="L520" class="ln">   520&nbsp;&nbsp;</span>	}
<span id="L521" class="ln">   521&nbsp;&nbsp;</span>
<span id="L522" class="ln">   522&nbsp;&nbsp;</span>	if !dataVal.IsValid() {
<span id="L523" class="ln">   523&nbsp;&nbsp;</span>		dataVal = reflect.Zero(val.Type())
<span id="L524" class="ln">   524&nbsp;&nbsp;</span>	}
<span id="L525" class="ln">   525&nbsp;&nbsp;</span>
<span id="L526" class="ln">   526&nbsp;&nbsp;</span>	dataValType := dataVal.Type()
<span id="L527" class="ln">   527&nbsp;&nbsp;</span>	if !dataValType.AssignableTo(val.Type()) {
<span id="L528" class="ln">   528&nbsp;&nbsp;</span>		return fmt.Errorf(
<span id="L529" class="ln">   529&nbsp;&nbsp;</span>			&#34;&#39;%s&#39; expected type &#39;%s&#39;, got &#39;%s&#39;&#34;,
<span id="L530" class="ln">   530&nbsp;&nbsp;</span>			name, val.Type(), dataValType)
<span id="L531" class="ln">   531&nbsp;&nbsp;</span>	}
<span id="L532" class="ln">   532&nbsp;&nbsp;</span>
<span id="L533" class="ln">   533&nbsp;&nbsp;</span>	val.Set(dataVal)
<span id="L534" class="ln">   534&nbsp;&nbsp;</span>	return nil
<span id="L535" class="ln">   535&nbsp;&nbsp;</span>}
<span id="L536" class="ln">   536&nbsp;&nbsp;</span>
<span id="L537" class="ln">   537&nbsp;&nbsp;</span>func (d *Decoder) decodeString(name string, data interface{}, val reflect.Value) error {
<span id="L538" class="ln">   538&nbsp;&nbsp;</span>	dataVal := reflect.Indirect(reflect.ValueOf(data))
<span id="L539" class="ln">   539&nbsp;&nbsp;</span>	dataKind := getKind(dataVal)
<span id="L540" class="ln">   540&nbsp;&nbsp;</span>
<span id="L541" class="ln">   541&nbsp;&nbsp;</span>	converted := true
<span id="L542" class="ln">   542&nbsp;&nbsp;</span>	switch {
<span id="L543" class="ln">   543&nbsp;&nbsp;</span>	case dataKind == reflect.String:
<span id="L544" class="ln">   544&nbsp;&nbsp;</span>		val.SetString(dataVal.String())
<span id="L545" class="ln">   545&nbsp;&nbsp;</span>	case dataKind == reflect.Bool &amp;&amp; d.config.WeaklyTypedInput:
<span id="L546" class="ln">   546&nbsp;&nbsp;</span>		if dataVal.Bool() {
<span id="L547" class="ln">   547&nbsp;&nbsp;</span>			val.SetString(&#34;1&#34;)
<span id="L548" class="ln">   548&nbsp;&nbsp;</span>		} else {
<span id="L549" class="ln">   549&nbsp;&nbsp;</span>			val.SetString(&#34;0&#34;)
<span id="L550" class="ln">   550&nbsp;&nbsp;</span>		}
<span id="L551" class="ln">   551&nbsp;&nbsp;</span>	case dataKind == reflect.Int &amp;&amp; d.config.WeaklyTypedInput:
<span id="L552" class="ln">   552&nbsp;&nbsp;</span>		val.SetString(strconv.FormatInt(dataVal.Int(), 10))
<span id="L553" class="ln">   553&nbsp;&nbsp;</span>	case dataKind == reflect.Uint &amp;&amp; d.config.WeaklyTypedInput:
<span id="L554" class="ln">   554&nbsp;&nbsp;</span>		val.SetString(strconv.FormatUint(dataVal.Uint(), 10))
<span id="L555" class="ln">   555&nbsp;&nbsp;</span>	case dataKind == reflect.Float32 &amp;&amp; d.config.WeaklyTypedInput:
<span id="L556" class="ln">   556&nbsp;&nbsp;</span>		val.SetString(strconv.FormatFloat(dataVal.Float(), &#39;f&#39;, -1, 64))
<span id="L557" class="ln">   557&nbsp;&nbsp;</span>	case dataKind == reflect.Slice &amp;&amp; d.config.WeaklyTypedInput,
<span id="L558" class="ln">   558&nbsp;&nbsp;</span>		dataKind == reflect.Array &amp;&amp; d.config.WeaklyTypedInput:
<span id="L559" class="ln">   559&nbsp;&nbsp;</span>		dataType := dataVal.Type()
<span id="L560" class="ln">   560&nbsp;&nbsp;</span>		elemKind := dataType.Elem().Kind()
<span id="L561" class="ln">   561&nbsp;&nbsp;</span>		switch elemKind {
<span id="L562" class="ln">   562&nbsp;&nbsp;</span>		case reflect.Uint8:
<span id="L563" class="ln">   563&nbsp;&nbsp;</span>			var uints []uint8
<span id="L564" class="ln">   564&nbsp;&nbsp;</span>			if dataKind == reflect.Array {
<span id="L565" class="ln">   565&nbsp;&nbsp;</span>				uints = make([]uint8, dataVal.Len(), dataVal.Len())
<span id="L566" class="ln">   566&nbsp;&nbsp;</span>				for i := range uints {
<span id="L567" class="ln">   567&nbsp;&nbsp;</span>					uints[i] = dataVal.Index(i).Interface().(uint8)
<span id="L568" class="ln">   568&nbsp;&nbsp;</span>				}
<span id="L569" class="ln">   569&nbsp;&nbsp;</span>			} else {
<span id="L570" class="ln">   570&nbsp;&nbsp;</span>				uints = dataVal.Interface().([]uint8)
<span id="L571" class="ln">   571&nbsp;&nbsp;</span>			}
<span id="L572" class="ln">   572&nbsp;&nbsp;</span>			val.SetString(string(uints))
<span id="L573" class="ln">   573&nbsp;&nbsp;</span>		default:
<span id="L574" class="ln">   574&nbsp;&nbsp;</span>			converted = false
<span id="L575" class="ln">   575&nbsp;&nbsp;</span>		}
<span id="L576" class="ln">   576&nbsp;&nbsp;</span>	default:
<span id="L577" class="ln">   577&nbsp;&nbsp;</span>		converted = false
<span id="L578" class="ln">   578&nbsp;&nbsp;</span>	}
<span id="L579" class="ln">   579&nbsp;&nbsp;</span>
<span id="L580" class="ln">   580&nbsp;&nbsp;</span>	if !converted {
<span id="L581" class="ln">   581&nbsp;&nbsp;</span>		return fmt.Errorf(
<span id="L582" class="ln">   582&nbsp;&nbsp;</span>			&#34;&#39;%s&#39; expected type &#39;%s&#39;, got unconvertible type &#39;%s&#39;, value: &#39;%v&#39;&#34;,
<span id="L583" class="ln">   583&nbsp;&nbsp;</span>			name, val.Type(), dataVal.Type(), data)
<span id="L584" class="ln">   584&nbsp;&nbsp;</span>	}
<span id="L585" class="ln">   585&nbsp;&nbsp;</span>
<span id="L586" class="ln">   586&nbsp;&nbsp;</span>	return nil
<span id="L587" class="ln">   587&nbsp;&nbsp;</span>}
<span id="L588" class="ln">   588&nbsp;&nbsp;</span>
<span id="L589" class="ln">   589&nbsp;&nbsp;</span>func (d *Decoder) decodeInt(name string, data interface{}, val reflect.Value) error {
<span id="L590" class="ln">   590&nbsp;&nbsp;</span>	dataVal := reflect.Indirect(reflect.ValueOf(data))
<span id="L591" class="ln">   591&nbsp;&nbsp;</span>	dataKind := getKind(dataVal)
<span id="L592" class="ln">   592&nbsp;&nbsp;</span>	dataType := dataVal.Type()
<span id="L593" class="ln">   593&nbsp;&nbsp;</span>
<span id="L594" class="ln">   594&nbsp;&nbsp;</span>	switch {
<span id="L595" class="ln">   595&nbsp;&nbsp;</span>	case dataKind == reflect.Int:
<span id="L596" class="ln">   596&nbsp;&nbsp;</span>		val.SetInt(dataVal.Int())
<span id="L597" class="ln">   597&nbsp;&nbsp;</span>	case dataKind == reflect.Uint:
<span id="L598" class="ln">   598&nbsp;&nbsp;</span>		val.SetInt(int64(dataVal.Uint()))
<span id="L599" class="ln">   599&nbsp;&nbsp;</span>	case dataKind == reflect.Float32:
<span id="L600" class="ln">   600&nbsp;&nbsp;</span>		val.SetInt(int64(dataVal.Float()))
<span id="L601" class="ln">   601&nbsp;&nbsp;</span>	case dataKind == reflect.Bool &amp;&amp; d.config.WeaklyTypedInput:
<span id="L602" class="ln">   602&nbsp;&nbsp;</span>		if dataVal.Bool() {
<span id="L603" class="ln">   603&nbsp;&nbsp;</span>			val.SetInt(1)
<span id="L604" class="ln">   604&nbsp;&nbsp;</span>		} else {
<span id="L605" class="ln">   605&nbsp;&nbsp;</span>			val.SetInt(0)
<span id="L606" class="ln">   606&nbsp;&nbsp;</span>		}
<span id="L607" class="ln">   607&nbsp;&nbsp;</span>	case dataKind == reflect.String &amp;&amp; d.config.WeaklyTypedInput:
<span id="L608" class="ln">   608&nbsp;&nbsp;</span>		str := dataVal.String()
<span id="L609" class="ln">   609&nbsp;&nbsp;</span>		if str == &#34;&#34; {
<span id="L610" class="ln">   610&nbsp;&nbsp;</span>			str = &#34;0&#34;
<span id="L611" class="ln">   611&nbsp;&nbsp;</span>		}
<span id="L612" class="ln">   612&nbsp;&nbsp;</span>
<span id="L613" class="ln">   613&nbsp;&nbsp;</span>		i, err := strconv.ParseInt(str, 0, val.Type().Bits())
<span id="L614" class="ln">   614&nbsp;&nbsp;</span>		if err == nil {
<span id="L615" class="ln">   615&nbsp;&nbsp;</span>			val.SetInt(i)
<span id="L616" class="ln">   616&nbsp;&nbsp;</span>		} else {
<span id="L617" class="ln">   617&nbsp;&nbsp;</span>			return fmt.Errorf(&#34;cannot parse &#39;%s&#39; as int: %s&#34;, name, err)
<span id="L618" class="ln">   618&nbsp;&nbsp;</span>		}
<span id="L619" class="ln">   619&nbsp;&nbsp;</span>	case dataType.PkgPath() == &#34;encoding/json&#34; &amp;&amp; dataType.Name() == &#34;Number&#34;:
<span id="L620" class="ln">   620&nbsp;&nbsp;</span>		jn := data.(json.Number)
<span id="L621" class="ln">   621&nbsp;&nbsp;</span>		i, err := jn.Int64()
<span id="L622" class="ln">   622&nbsp;&nbsp;</span>		if err != nil {
<span id="L623" class="ln">   623&nbsp;&nbsp;</span>			return fmt.Errorf(
<span id="L624" class="ln">   624&nbsp;&nbsp;</span>				&#34;error decoding json.Number into %s: %s&#34;, name, err)
<span id="L625" class="ln">   625&nbsp;&nbsp;</span>		}
<span id="L626" class="ln">   626&nbsp;&nbsp;</span>		val.SetInt(i)
<span id="L627" class="ln">   627&nbsp;&nbsp;</span>	default:
<span id="L628" class="ln">   628&nbsp;&nbsp;</span>		return fmt.Errorf(
<span id="L629" class="ln">   629&nbsp;&nbsp;</span>			&#34;&#39;%s&#39; expected type &#39;%s&#39;, got unconvertible type &#39;%s&#39;, value: &#39;%v&#39;&#34;,
<span id="L630" class="ln">   630&nbsp;&nbsp;</span>			name, val.Type(), dataVal.Type(), data)
<span id="L631" class="ln">   631&nbsp;&nbsp;</span>	}
<span id="L632" class="ln">   632&nbsp;&nbsp;</span>
<span id="L633" class="ln">   633&nbsp;&nbsp;</span>	return nil
<span id="L634" class="ln">   634&nbsp;&nbsp;</span>}
<span id="L635" class="ln">   635&nbsp;&nbsp;</span>
<span id="L636" class="ln">   636&nbsp;&nbsp;</span>func (d *Decoder) decodeUint(name string, data interface{}, val reflect.Value) error {
<span id="L637" class="ln">   637&nbsp;&nbsp;</span>	dataVal := reflect.Indirect(reflect.ValueOf(data))
<span id="L638" class="ln">   638&nbsp;&nbsp;</span>	dataKind := getKind(dataVal)
<span id="L639" class="ln">   639&nbsp;&nbsp;</span>	dataType := dataVal.Type()
<span id="L640" class="ln">   640&nbsp;&nbsp;</span>
<span id="L641" class="ln">   641&nbsp;&nbsp;</span>	switch {
<span id="L642" class="ln">   642&nbsp;&nbsp;</span>	case dataKind == reflect.Int:
<span id="L643" class="ln">   643&nbsp;&nbsp;</span>		i := dataVal.Int()
<span id="L644" class="ln">   644&nbsp;&nbsp;</span>		if i &lt; 0 &amp;&amp; !d.config.WeaklyTypedInput {
<span id="L645" class="ln">   645&nbsp;&nbsp;</span>			return fmt.Errorf(&#34;cannot parse &#39;%s&#39;, %d overflows uint&#34;,
<span id="L646" class="ln">   646&nbsp;&nbsp;</span>				name, i)
<span id="L647" class="ln">   647&nbsp;&nbsp;</span>		}
<span id="L648" class="ln">   648&nbsp;&nbsp;</span>		val.SetUint(uint64(i))
<span id="L649" class="ln">   649&nbsp;&nbsp;</span>	case dataKind == reflect.Uint:
<span id="L650" class="ln">   650&nbsp;&nbsp;</span>		val.SetUint(dataVal.Uint())
<span id="L651" class="ln">   651&nbsp;&nbsp;</span>	case dataKind == reflect.Float32:
<span id="L652" class="ln">   652&nbsp;&nbsp;</span>		f := dataVal.Float()
<span id="L653" class="ln">   653&nbsp;&nbsp;</span>		if f &lt; 0 &amp;&amp; !d.config.WeaklyTypedInput {
<span id="L654" class="ln">   654&nbsp;&nbsp;</span>			return fmt.Errorf(&#34;cannot parse &#39;%s&#39;, %f overflows uint&#34;,
<span id="L655" class="ln">   655&nbsp;&nbsp;</span>				name, f)
<span id="L656" class="ln">   656&nbsp;&nbsp;</span>		}
<span id="L657" class="ln">   657&nbsp;&nbsp;</span>		val.SetUint(uint64(f))
<span id="L658" class="ln">   658&nbsp;&nbsp;</span>	case dataKind == reflect.Bool &amp;&amp; d.config.WeaklyTypedInput:
<span id="L659" class="ln">   659&nbsp;&nbsp;</span>		if dataVal.Bool() {
<span id="L660" class="ln">   660&nbsp;&nbsp;</span>			val.SetUint(1)
<span id="L661" class="ln">   661&nbsp;&nbsp;</span>		} else {
<span id="L662" class="ln">   662&nbsp;&nbsp;</span>			val.SetUint(0)
<span id="L663" class="ln">   663&nbsp;&nbsp;</span>		}
<span id="L664" class="ln">   664&nbsp;&nbsp;</span>	case dataKind == reflect.String &amp;&amp; d.config.WeaklyTypedInput:
<span id="L665" class="ln">   665&nbsp;&nbsp;</span>		str := dataVal.String()
<span id="L666" class="ln">   666&nbsp;&nbsp;</span>		if str == &#34;&#34; {
<span id="L667" class="ln">   667&nbsp;&nbsp;</span>			str = &#34;0&#34;
<span id="L668" class="ln">   668&nbsp;&nbsp;</span>		}
<span id="L669" class="ln">   669&nbsp;&nbsp;</span>
<span id="L670" class="ln">   670&nbsp;&nbsp;</span>		i, err := strconv.ParseUint(str, 0, val.Type().Bits())
<span id="L671" class="ln">   671&nbsp;&nbsp;</span>		if err == nil {
<span id="L672" class="ln">   672&nbsp;&nbsp;</span>			val.SetUint(i)
<span id="L673" class="ln">   673&nbsp;&nbsp;</span>		} else {
<span id="L674" class="ln">   674&nbsp;&nbsp;</span>			return fmt.Errorf(&#34;cannot parse &#39;%s&#39; as uint: %s&#34;, name, err)
<span id="L675" class="ln">   675&nbsp;&nbsp;</span>		}
<span id="L676" class="ln">   676&nbsp;&nbsp;</span>	case dataType.PkgPath() == &#34;encoding/json&#34; &amp;&amp; dataType.Name() == &#34;Number&#34;:
<span id="L677" class="ln">   677&nbsp;&nbsp;</span>		jn := data.(json.Number)
<span id="L678" class="ln">   678&nbsp;&nbsp;</span>		i, err := jn.Int64()
<span id="L679" class="ln">   679&nbsp;&nbsp;</span>		if err != nil {
<span id="L680" class="ln">   680&nbsp;&nbsp;</span>			return fmt.Errorf(
<span id="L681" class="ln">   681&nbsp;&nbsp;</span>				&#34;error decoding json.Number into %s: %s&#34;, name, err)
<span id="L682" class="ln">   682&nbsp;&nbsp;</span>		}
<span id="L683" class="ln">   683&nbsp;&nbsp;</span>		if i &lt; 0 &amp;&amp; !d.config.WeaklyTypedInput {
<span id="L684" class="ln">   684&nbsp;&nbsp;</span>			return fmt.Errorf(&#34;cannot parse &#39;%s&#39;, %d overflows uint&#34;,
<span id="L685" class="ln">   685&nbsp;&nbsp;</span>				name, i)
<span id="L686" class="ln">   686&nbsp;&nbsp;</span>		}
<span id="L687" class="ln">   687&nbsp;&nbsp;</span>		val.SetUint(uint64(i))
<span id="L688" class="ln">   688&nbsp;&nbsp;</span>	default:
<span id="L689" class="ln">   689&nbsp;&nbsp;</span>		return fmt.Errorf(
<span id="L690" class="ln">   690&nbsp;&nbsp;</span>			&#34;&#39;%s&#39; expected type &#39;%s&#39;, got unconvertible type &#39;%s&#39;, value: &#39;%v&#39;&#34;,
<span id="L691" class="ln">   691&nbsp;&nbsp;</span>			name, val.Type(), dataVal.Type(), data)
<span id="L692" class="ln">   692&nbsp;&nbsp;</span>	}
<span id="L693" class="ln">   693&nbsp;&nbsp;</span>
<span id="L694" class="ln">   694&nbsp;&nbsp;</span>	return nil
<span id="L695" class="ln">   695&nbsp;&nbsp;</span>}
<span id="L696" class="ln">   696&nbsp;&nbsp;</span>
<span id="L697" class="ln">   697&nbsp;&nbsp;</span>func (d *Decoder) decodeBool(name string, data interface{}, val reflect.Value) error {
<span id="L698" class="ln">   698&nbsp;&nbsp;</span>	dataVal := reflect.Indirect(reflect.ValueOf(data))
<span id="L699" class="ln">   699&nbsp;&nbsp;</span>	dataKind := getKind(dataVal)
<span id="L700" class="ln">   700&nbsp;&nbsp;</span>
<span id="L701" class="ln">   701&nbsp;&nbsp;</span>	switch {
<span id="L702" class="ln">   702&nbsp;&nbsp;</span>	case dataKind == reflect.Bool:
<span id="L703" class="ln">   703&nbsp;&nbsp;</span>		val.SetBool(dataVal.Bool())
<span id="L704" class="ln">   704&nbsp;&nbsp;</span>	case dataKind == reflect.Int &amp;&amp; d.config.WeaklyTypedInput:
<span id="L705" class="ln">   705&nbsp;&nbsp;</span>		val.SetBool(dataVal.Int() != 0)
<span id="L706" class="ln">   706&nbsp;&nbsp;</span>	case dataKind == reflect.Uint &amp;&amp; d.config.WeaklyTypedInput:
<span id="L707" class="ln">   707&nbsp;&nbsp;</span>		val.SetBool(dataVal.Uint() != 0)
<span id="L708" class="ln">   708&nbsp;&nbsp;</span>	case dataKind == reflect.Float32 &amp;&amp; d.config.WeaklyTypedInput:
<span id="L709" class="ln">   709&nbsp;&nbsp;</span>		val.SetBool(dataVal.Float() != 0)
<span id="L710" class="ln">   710&nbsp;&nbsp;</span>	case dataKind == reflect.String &amp;&amp; d.config.WeaklyTypedInput:
<span id="L711" class="ln">   711&nbsp;&nbsp;</span>		b, err := strconv.ParseBool(dataVal.String())
<span id="L712" class="ln">   712&nbsp;&nbsp;</span>		if err == nil {
<span id="L713" class="ln">   713&nbsp;&nbsp;</span>			val.SetBool(b)
<span id="L714" class="ln">   714&nbsp;&nbsp;</span>		} else if dataVal.String() == &#34;&#34; {
<span id="L715" class="ln">   715&nbsp;&nbsp;</span>			val.SetBool(false)
<span id="L716" class="ln">   716&nbsp;&nbsp;</span>		} else {
<span id="L717" class="ln">   717&nbsp;&nbsp;</span>			return fmt.Errorf(&#34;cannot parse &#39;%s&#39; as bool: %s&#34;, name, err)
<span id="L718" class="ln">   718&nbsp;&nbsp;</span>		}
<span id="L719" class="ln">   719&nbsp;&nbsp;</span>	default:
<span id="L720" class="ln">   720&nbsp;&nbsp;</span>		return fmt.Errorf(
<span id="L721" class="ln">   721&nbsp;&nbsp;</span>			&#34;&#39;%s&#39; expected type &#39;%s&#39;, got unconvertible type &#39;%s&#39;, value: &#39;%v&#39;&#34;,
<span id="L722" class="ln">   722&nbsp;&nbsp;</span>			name, val.Type(), dataVal.Type(), data)
<span id="L723" class="ln">   723&nbsp;&nbsp;</span>	}
<span id="L724" class="ln">   724&nbsp;&nbsp;</span>
<span id="L725" class="ln">   725&nbsp;&nbsp;</span>	return nil
<span id="L726" class="ln">   726&nbsp;&nbsp;</span>}
<span id="L727" class="ln">   727&nbsp;&nbsp;</span>
<span id="L728" class="ln">   728&nbsp;&nbsp;</span>func (d *Decoder) decodeFloat(name string, data interface{}, val reflect.Value) error {
<span id="L729" class="ln">   729&nbsp;&nbsp;</span>	dataVal := reflect.Indirect(reflect.ValueOf(data))
<span id="L730" class="ln">   730&nbsp;&nbsp;</span>	dataKind := getKind(dataVal)
<span id="L731" class="ln">   731&nbsp;&nbsp;</span>	dataType := dataVal.Type()
<span id="L732" class="ln">   732&nbsp;&nbsp;</span>
<span id="L733" class="ln">   733&nbsp;&nbsp;</span>	switch {
<span id="L734" class="ln">   734&nbsp;&nbsp;</span>	case dataKind == reflect.Int:
<span id="L735" class="ln">   735&nbsp;&nbsp;</span>		val.SetFloat(float64(dataVal.Int()))
<span id="L736" class="ln">   736&nbsp;&nbsp;</span>	case dataKind == reflect.Uint:
<span id="L737" class="ln">   737&nbsp;&nbsp;</span>		val.SetFloat(float64(dataVal.Uint()))
<span id="L738" class="ln">   738&nbsp;&nbsp;</span>	case dataKind == reflect.Float32:
<span id="L739" class="ln">   739&nbsp;&nbsp;</span>		val.SetFloat(dataVal.Float())
<span id="L740" class="ln">   740&nbsp;&nbsp;</span>	case dataKind == reflect.Bool &amp;&amp; d.config.WeaklyTypedInput:
<span id="L741" class="ln">   741&nbsp;&nbsp;</span>		if dataVal.Bool() {
<span id="L742" class="ln">   742&nbsp;&nbsp;</span>			val.SetFloat(1)
<span id="L743" class="ln">   743&nbsp;&nbsp;</span>		} else {
<span id="L744" class="ln">   744&nbsp;&nbsp;</span>			val.SetFloat(0)
<span id="L745" class="ln">   745&nbsp;&nbsp;</span>		}
<span id="L746" class="ln">   746&nbsp;&nbsp;</span>	case dataKind == reflect.String &amp;&amp; d.config.WeaklyTypedInput:
<span id="L747" class="ln">   747&nbsp;&nbsp;</span>		str := dataVal.String()
<span id="L748" class="ln">   748&nbsp;&nbsp;</span>		if str == &#34;&#34; {
<span id="L749" class="ln">   749&nbsp;&nbsp;</span>			str = &#34;0&#34;
<span id="L750" class="ln">   750&nbsp;&nbsp;</span>		}
<span id="L751" class="ln">   751&nbsp;&nbsp;</span>
<span id="L752" class="ln">   752&nbsp;&nbsp;</span>		f, err := strconv.ParseFloat(str, val.Type().Bits())
<span id="L753" class="ln">   753&nbsp;&nbsp;</span>		if err == nil {
<span id="L754" class="ln">   754&nbsp;&nbsp;</span>			val.SetFloat(f)
<span id="L755" class="ln">   755&nbsp;&nbsp;</span>		} else {
<span id="L756" class="ln">   756&nbsp;&nbsp;</span>			return fmt.Errorf(&#34;cannot parse &#39;%s&#39; as float: %s&#34;, name, err)
<span id="L757" class="ln">   757&nbsp;&nbsp;</span>		}
<span id="L758" class="ln">   758&nbsp;&nbsp;</span>	case dataType.PkgPath() == &#34;encoding/json&#34; &amp;&amp; dataType.Name() == &#34;Number&#34;:
<span id="L759" class="ln">   759&nbsp;&nbsp;</span>		jn := data.(json.Number)
<span id="L760" class="ln">   760&nbsp;&nbsp;</span>		i, err := jn.Float64()
<span id="L761" class="ln">   761&nbsp;&nbsp;</span>		if err != nil {
<span id="L762" class="ln">   762&nbsp;&nbsp;</span>			return fmt.Errorf(
<span id="L763" class="ln">   763&nbsp;&nbsp;</span>				&#34;error decoding json.Number into %s: %s&#34;, name, err)
<span id="L764" class="ln">   764&nbsp;&nbsp;</span>		}
<span id="L765" class="ln">   765&nbsp;&nbsp;</span>		val.SetFloat(i)
<span id="L766" class="ln">   766&nbsp;&nbsp;</span>	default:
<span id="L767" class="ln">   767&nbsp;&nbsp;</span>		return fmt.Errorf(
<span id="L768" class="ln">   768&nbsp;&nbsp;</span>			&#34;&#39;%s&#39; expected type &#39;%s&#39;, got unconvertible type &#39;%s&#39;, value: &#39;%v&#39;&#34;,
<span id="L769" class="ln">   769&nbsp;&nbsp;</span>			name, val.Type(), dataVal.Type(), data)
<span id="L770" class="ln">   770&nbsp;&nbsp;</span>	}
<span id="L771" class="ln">   771&nbsp;&nbsp;</span>
<span id="L772" class="ln">   772&nbsp;&nbsp;</span>	return nil
<span id="L773" class="ln">   773&nbsp;&nbsp;</span>}
<span id="L774" class="ln">   774&nbsp;&nbsp;</span>
<span id="L775" class="ln">   775&nbsp;&nbsp;</span>func (d *Decoder) decodeMap(name string, data interface{}, val reflect.Value) error {
<span id="L776" class="ln">   776&nbsp;&nbsp;</span>	valType := val.Type()
<span id="L777" class="ln">   777&nbsp;&nbsp;</span>	valKeyType := valType.Key()
<span id="L778" class="ln">   778&nbsp;&nbsp;</span>	valElemType := valType.Elem()
<span id="L779" class="ln">   779&nbsp;&nbsp;</span>
<span id="L780" class="ln">   780&nbsp;&nbsp;</span>	<span class="comment">// By default we overwrite keys in the current map</span>
<span id="L781" class="ln">   781&nbsp;&nbsp;</span>	valMap := val
<span id="L782" class="ln">   782&nbsp;&nbsp;</span>
<span id="L783" class="ln">   783&nbsp;&nbsp;</span>	<span class="comment">// If the map is nil or we&#39;re purposely zeroing fields, make a new map</span>
<span id="L784" class="ln">   784&nbsp;&nbsp;</span>	if valMap.IsNil() || d.config.ZeroFields {
<span id="L785" class="ln">   785&nbsp;&nbsp;</span>		<span class="comment">// Make a new map to hold our result</span>
<span id="L786" class="ln">   786&nbsp;&nbsp;</span>		mapType := reflect.MapOf(valKeyType, valElemType)
<span id="L787" class="ln">   787&nbsp;&nbsp;</span>		valMap = reflect.MakeMap(mapType)
<span id="L788" class="ln">   788&nbsp;&nbsp;</span>	}
<span id="L789" class="ln">   789&nbsp;&nbsp;</span>
<span id="L790" class="ln">   790&nbsp;&nbsp;</span>	<span class="comment">// Check input type and based on the input type jump to the proper func</span>
<span id="L791" class="ln">   791&nbsp;&nbsp;</span>	dataVal := reflect.Indirect(reflect.ValueOf(data))
<span id="L792" class="ln">   792&nbsp;&nbsp;</span>	switch dataVal.Kind() {
<span id="L793" class="ln">   793&nbsp;&nbsp;</span>	case reflect.Map:
<span id="L794" class="ln">   794&nbsp;&nbsp;</span>		return d.decodeMapFromMap(name, dataVal, val, valMap)
<span id="L795" class="ln">   795&nbsp;&nbsp;</span>
<span id="L796" class="ln">   796&nbsp;&nbsp;</span>	case reflect.Struct:
<span id="L797" class="ln">   797&nbsp;&nbsp;</span>		return d.decodeMapFromStruct(name, dataVal, val, valMap)
<span id="L798" class="ln">   798&nbsp;&nbsp;</span>
<span id="L799" class="ln">   799&nbsp;&nbsp;</span>	case reflect.Array, reflect.Slice:
<span id="L800" class="ln">   800&nbsp;&nbsp;</span>		if d.config.WeaklyTypedInput {
<span id="L801" class="ln">   801&nbsp;&nbsp;</span>			return d.decodeMapFromSlice(name, dataVal, val, valMap)
<span id="L802" class="ln">   802&nbsp;&nbsp;</span>		}
<span id="L803" class="ln">   803&nbsp;&nbsp;</span>
<span id="L804" class="ln">   804&nbsp;&nbsp;</span>		fallthrough
<span id="L805" class="ln">   805&nbsp;&nbsp;</span>
<span id="L806" class="ln">   806&nbsp;&nbsp;</span>	default:
<span id="L807" class="ln">   807&nbsp;&nbsp;</span>		return fmt.Errorf(&#34;&#39;%s&#39; expected a map, got &#39;%s&#39;&#34;, name, dataVal.Kind())
<span id="L808" class="ln">   808&nbsp;&nbsp;</span>	}
<span id="L809" class="ln">   809&nbsp;&nbsp;</span>}
<span id="L810" class="ln">   810&nbsp;&nbsp;</span>
<span id="L811" class="ln">   811&nbsp;&nbsp;</span>func (d *Decoder) decodeMapFromSlice(name string, dataVal reflect.Value, val reflect.Value, valMap reflect.Value) error {
<span id="L812" class="ln">   812&nbsp;&nbsp;</span>	<span class="comment">// Special case for BC reasons (covered by tests)</span>
<span id="L813" class="ln">   813&nbsp;&nbsp;</span>	if dataVal.Len() == 0 {
<span id="L814" class="ln">   814&nbsp;&nbsp;</span>		val.Set(valMap)
<span id="L815" class="ln">   815&nbsp;&nbsp;</span>		return nil
<span id="L816" class="ln">   816&nbsp;&nbsp;</span>	}
<span id="L817" class="ln">   817&nbsp;&nbsp;</span>
<span id="L818" class="ln">   818&nbsp;&nbsp;</span>	for i := 0; i &lt; dataVal.Len(); i++ {
<span id="L819" class="ln">   819&nbsp;&nbsp;</span>		err := d.decode(
<span id="L820" class="ln">   820&nbsp;&nbsp;</span>			name+&#34;[&#34;+strconv.Itoa(i)+&#34;]&#34;,
<span id="L821" class="ln">   821&nbsp;&nbsp;</span>			dataVal.Index(i).Interface(), val)
<span id="L822" class="ln">   822&nbsp;&nbsp;</span>		if err != nil {
<span id="L823" class="ln">   823&nbsp;&nbsp;</span>			return err
<span id="L824" class="ln">   824&nbsp;&nbsp;</span>		}
<span id="L825" class="ln">   825&nbsp;&nbsp;</span>	}
<span id="L826" class="ln">   826&nbsp;&nbsp;</span>
<span id="L827" class="ln">   827&nbsp;&nbsp;</span>	return nil
<span id="L828" class="ln">   828&nbsp;&nbsp;</span>}
<span id="L829" class="ln">   829&nbsp;&nbsp;</span>
<span id="L830" class="ln">   830&nbsp;&nbsp;</span>func (d *Decoder) decodeMapFromMap(name string, dataVal reflect.Value, val reflect.Value, valMap reflect.Value) error {
<span id="L831" class="ln">   831&nbsp;&nbsp;</span>	valType := val.Type()
<span id="L832" class="ln">   832&nbsp;&nbsp;</span>	valKeyType := valType.Key()
<span id="L833" class="ln">   833&nbsp;&nbsp;</span>	valElemType := valType.Elem()
<span id="L834" class="ln">   834&nbsp;&nbsp;</span>
<span id="L835" class="ln">   835&nbsp;&nbsp;</span>	<span class="comment">// Accumulate errors</span>
<span id="L836" class="ln">   836&nbsp;&nbsp;</span>	errors := make([]string, 0)
<span id="L837" class="ln">   837&nbsp;&nbsp;</span>
<span id="L838" class="ln">   838&nbsp;&nbsp;</span>	<span class="comment">// If the input data is empty, then we just match what the input data is.</span>
<span id="L839" class="ln">   839&nbsp;&nbsp;</span>	if dataVal.Len() == 0 {
<span id="L840" class="ln">   840&nbsp;&nbsp;</span>		if dataVal.IsNil() {
<span id="L841" class="ln">   841&nbsp;&nbsp;</span>			if !val.IsNil() {
<span id="L842" class="ln">   842&nbsp;&nbsp;</span>				val.Set(dataVal)
<span id="L843" class="ln">   843&nbsp;&nbsp;</span>			}
<span id="L844" class="ln">   844&nbsp;&nbsp;</span>		} else {
<span id="L845" class="ln">   845&nbsp;&nbsp;</span>			<span class="comment">// Set to empty allocated value</span>
<span id="L846" class="ln">   846&nbsp;&nbsp;</span>			val.Set(valMap)
<span id="L847" class="ln">   847&nbsp;&nbsp;</span>		}
<span id="L848" class="ln">   848&nbsp;&nbsp;</span>
<span id="L849" class="ln">   849&nbsp;&nbsp;</span>		return nil
<span id="L850" class="ln">   850&nbsp;&nbsp;</span>	}
<span id="L851" class="ln">   851&nbsp;&nbsp;</span>
<span id="L852" class="ln">   852&nbsp;&nbsp;</span>	for _, k := range dataVal.MapKeys() {
<span id="L853" class="ln">   853&nbsp;&nbsp;</span>		fieldName := name + &#34;[&#34; + k.String() + &#34;]&#34;
<span id="L854" class="ln">   854&nbsp;&nbsp;</span>
<span id="L855" class="ln">   855&nbsp;&nbsp;</span>		<span class="comment">// First decode the key into the proper type</span>
<span id="L856" class="ln">   856&nbsp;&nbsp;</span>		currentKey := reflect.Indirect(reflect.New(valKeyType))
<span id="L857" class="ln">   857&nbsp;&nbsp;</span>		if err := d.decode(fieldName, k.Interface(), currentKey); err != nil {
<span id="L858" class="ln">   858&nbsp;&nbsp;</span>			errors = appendErrors(errors, err)
<span id="L859" class="ln">   859&nbsp;&nbsp;</span>			continue
<span id="L860" class="ln">   860&nbsp;&nbsp;</span>		}
<span id="L861" class="ln">   861&nbsp;&nbsp;</span>
<span id="L862" class="ln">   862&nbsp;&nbsp;</span>		<span class="comment">// Next decode the data into the proper type</span>
<span id="L863" class="ln">   863&nbsp;&nbsp;</span>		v := dataVal.MapIndex(k).Interface()
<span id="L864" class="ln">   864&nbsp;&nbsp;</span>		currentVal := reflect.Indirect(reflect.New(valElemType))
<span id="L865" class="ln">   865&nbsp;&nbsp;</span>		if err := d.decode(fieldName, v, currentVal); err != nil {
<span id="L866" class="ln">   866&nbsp;&nbsp;</span>			errors = appendErrors(errors, err)
<span id="L867" class="ln">   867&nbsp;&nbsp;</span>			continue
<span id="L868" class="ln">   868&nbsp;&nbsp;</span>		}
<span id="L869" class="ln">   869&nbsp;&nbsp;</span>
<span id="L870" class="ln">   870&nbsp;&nbsp;</span>		valMap.SetMapIndex(currentKey, currentVal)
<span id="L871" class="ln">   871&nbsp;&nbsp;</span>	}
<span id="L872" class="ln">   872&nbsp;&nbsp;</span>
<span id="L873" class="ln">   873&nbsp;&nbsp;</span>	<span class="comment">// Set the built up map to the value</span>
<span id="L874" class="ln">   874&nbsp;&nbsp;</span>	val.Set(valMap)
<span id="L875" class="ln">   875&nbsp;&nbsp;</span>
<span id="L876" class="ln">   876&nbsp;&nbsp;</span>	<span class="comment">// If we had errors, return those</span>
<span id="L877" class="ln">   877&nbsp;&nbsp;</span>	if len(errors) &gt; 0 {
<span id="L878" class="ln">   878&nbsp;&nbsp;</span>		return &amp;Error{errors}
<span id="L879" class="ln">   879&nbsp;&nbsp;</span>	}
<span id="L880" class="ln">   880&nbsp;&nbsp;</span>
<span id="L881" class="ln">   881&nbsp;&nbsp;</span>	return nil
<span id="L882" class="ln">   882&nbsp;&nbsp;</span>}
<span id="L883" class="ln">   883&nbsp;&nbsp;</span>
<span id="L884" class="ln">   884&nbsp;&nbsp;</span>func (d *Decoder) decodeMapFromStruct(name string, dataVal reflect.Value, val reflect.Value, valMap reflect.Value) error {
<span id="L885" class="ln">   885&nbsp;&nbsp;</span>	typ := dataVal.Type()
<span id="L886" class="ln">   886&nbsp;&nbsp;</span>	for i := 0; i &lt; typ.NumField(); i++ {
<span id="L887" class="ln">   887&nbsp;&nbsp;</span>		<span class="comment">// Get the StructField first since this is a cheap operation. If the</span>
<span id="L888" class="ln">   888&nbsp;&nbsp;</span>		<span class="comment">// field is unexported, then ignore it.</span>
<span id="L889" class="ln">   889&nbsp;&nbsp;</span>		f := typ.Field(i)
<span id="L890" class="ln">   890&nbsp;&nbsp;</span>		if f.PkgPath != &#34;&#34; {
<span id="L891" class="ln">   891&nbsp;&nbsp;</span>			continue
<span id="L892" class="ln">   892&nbsp;&nbsp;</span>		}
<span id="L893" class="ln">   893&nbsp;&nbsp;</span>
<span id="L894" class="ln">   894&nbsp;&nbsp;</span>		<span class="comment">// Next get the actual value of this field and verify it is assignable</span>
<span id="L895" class="ln">   895&nbsp;&nbsp;</span>		<span class="comment">// to the map value.</span>
<span id="L896" class="ln">   896&nbsp;&nbsp;</span>		v := dataVal.Field(i)
<span id="L897" class="ln">   897&nbsp;&nbsp;</span>		if !v.Type().AssignableTo(valMap.Type().Elem()) {
<span id="L898" class="ln">   898&nbsp;&nbsp;</span>			return fmt.Errorf(&#34;cannot assign type &#39;%s&#39; to map value field of type &#39;%s&#39;&#34;, v.Type(), valMap.Type().Elem())
<span id="L899" class="ln">   899&nbsp;&nbsp;</span>		}
<span id="L900" class="ln">   900&nbsp;&nbsp;</span>
<span id="L901" class="ln">   901&nbsp;&nbsp;</span>		tagValue := f.Tag.Get(d.config.TagName)
<span id="L902" class="ln">   902&nbsp;&nbsp;</span>		keyName := f.Name
<span id="L903" class="ln">   903&nbsp;&nbsp;</span>
<span id="L904" class="ln">   904&nbsp;&nbsp;</span>		<span class="comment">// If Squash is set in the config, we squash the field down.</span>
<span id="L905" class="ln">   905&nbsp;&nbsp;</span>		squash := d.config.Squash &amp;&amp; v.Kind() == reflect.Struct &amp;&amp; f.Anonymous
<span id="L906" class="ln">   906&nbsp;&nbsp;</span>
<span id="L907" class="ln">   907&nbsp;&nbsp;</span>		<span class="comment">// Determine the name of the key in the map</span>
<span id="L908" class="ln">   908&nbsp;&nbsp;</span>		if index := strings.Index(tagValue, &#34;,&#34;); index != -1 {
<span id="L909" class="ln">   909&nbsp;&nbsp;</span>			if tagValue[:index] == &#34;-&#34; {
<span id="L910" class="ln">   910&nbsp;&nbsp;</span>				continue
<span id="L911" class="ln">   911&nbsp;&nbsp;</span>			}
<span id="L912" class="ln">   912&nbsp;&nbsp;</span>			<span class="comment">// If &#34;omitempty&#34; is specified in the tag, it ignores empty values.</span>
<span id="L913" class="ln">   913&nbsp;&nbsp;</span>			if strings.Index(tagValue[index+1:], &#34;omitempty&#34;) != -1 &amp;&amp; isEmptyValue(v) {
<span id="L914" class="ln">   914&nbsp;&nbsp;</span>				continue
<span id="L915" class="ln">   915&nbsp;&nbsp;</span>			}
<span id="L916" class="ln">   916&nbsp;&nbsp;</span>
<span id="L917" class="ln">   917&nbsp;&nbsp;</span>			<span class="comment">// If &#34;squash&#34; is specified in the tag, we squash the field down.</span>
<span id="L918" class="ln">   918&nbsp;&nbsp;</span>			squash = !squash &amp;&amp; strings.Index(tagValue[index+1:], &#34;squash&#34;) != -1
<span id="L919" class="ln">   919&nbsp;&nbsp;</span>			if squash {
<span id="L920" class="ln">   920&nbsp;&nbsp;</span>				<span class="comment">// When squashing, the embedded type can be a pointer to a struct.</span>
<span id="L921" class="ln">   921&nbsp;&nbsp;</span>				if v.Kind() == reflect.Ptr &amp;&amp; v.Elem().Kind() == reflect.Struct {
<span id="L922" class="ln">   922&nbsp;&nbsp;</span>					v = v.Elem()
<span id="L923" class="ln">   923&nbsp;&nbsp;</span>				}
<span id="L924" class="ln">   924&nbsp;&nbsp;</span>
<span id="L925" class="ln">   925&nbsp;&nbsp;</span>				<span class="comment">// The final type must be a struct</span>
<span id="L926" class="ln">   926&nbsp;&nbsp;</span>				if v.Kind() != reflect.Struct {
<span id="L927" class="ln">   927&nbsp;&nbsp;</span>					return fmt.Errorf(&#34;cannot squash non-struct type &#39;%s&#39;&#34;, v.Type())
<span id="L928" class="ln">   928&nbsp;&nbsp;</span>				}
<span id="L929" class="ln">   929&nbsp;&nbsp;</span>			}
<span id="L930" class="ln">   930&nbsp;&nbsp;</span>			keyName = tagValue[:index]
<span id="L931" class="ln">   931&nbsp;&nbsp;</span>		} else if len(tagValue) &gt; 0 {
<span id="L932" class="ln">   932&nbsp;&nbsp;</span>			if tagValue == &#34;-&#34; {
<span id="L933" class="ln">   933&nbsp;&nbsp;</span>				continue
<span id="L934" class="ln">   934&nbsp;&nbsp;</span>			}
<span id="L935" class="ln">   935&nbsp;&nbsp;</span>			keyName = tagValue
<span id="L936" class="ln">   936&nbsp;&nbsp;</span>		}
<span id="L937" class="ln">   937&nbsp;&nbsp;</span>
<span id="L938" class="ln">   938&nbsp;&nbsp;</span>		switch v.Kind() {
<span id="L939" class="ln">   939&nbsp;&nbsp;</span>		<span class="comment">// this is an embedded struct, so handle it differently</span>
<span id="L940" class="ln">   940&nbsp;&nbsp;</span>		case reflect.Struct:
<span id="L941" class="ln">   941&nbsp;&nbsp;</span>			x := reflect.New(v.Type())
<span id="L942" class="ln">   942&nbsp;&nbsp;</span>			x.Elem().Set(v)
<span id="L943" class="ln">   943&nbsp;&nbsp;</span>
<span id="L944" class="ln">   944&nbsp;&nbsp;</span>			vType := valMap.Type()
<span id="L945" class="ln">   945&nbsp;&nbsp;</span>			vKeyType := vType.Key()
<span id="L946" class="ln">   946&nbsp;&nbsp;</span>			vElemType := vType.Elem()
<span id="L947" class="ln">   947&nbsp;&nbsp;</span>			mType := reflect.MapOf(vKeyType, vElemType)
<span id="L948" class="ln">   948&nbsp;&nbsp;</span>			vMap := reflect.MakeMap(mType)
<span id="L949" class="ln">   949&nbsp;&nbsp;</span>
<span id="L950" class="ln">   950&nbsp;&nbsp;</span>			<span class="comment">// Creating a pointer to a map so that other methods can completely</span>
<span id="L951" class="ln">   951&nbsp;&nbsp;</span>			<span class="comment">// overwrite the map if need be (looking at you decodeMapFromMap). The</span>
<span id="L952" class="ln">   952&nbsp;&nbsp;</span>			<span class="comment">// indirection allows the underlying map to be settable (CanSet() == true)</span>
<span id="L953" class="ln">   953&nbsp;&nbsp;</span>			<span class="comment">// where as reflect.MakeMap returns an unsettable map.</span>
<span id="L954" class="ln">   954&nbsp;&nbsp;</span>			addrVal := reflect.New(vMap.Type())
<span id="L955" class="ln">   955&nbsp;&nbsp;</span>			reflect.Indirect(addrVal).Set(vMap)
<span id="L956" class="ln">   956&nbsp;&nbsp;</span>
<span id="L957" class="ln">   957&nbsp;&nbsp;</span>			err := d.decode(keyName, x.Interface(), reflect.Indirect(addrVal))
<span id="L958" class="ln">   958&nbsp;&nbsp;</span>			if err != nil {
<span id="L959" class="ln">   959&nbsp;&nbsp;</span>				return err
<span id="L960" class="ln">   960&nbsp;&nbsp;</span>			}
<span id="L961" class="ln">   961&nbsp;&nbsp;</span>
<span id="L962" class="ln">   962&nbsp;&nbsp;</span>			<span class="comment">// the underlying map may have been completely overwritten so pull</span>
<span id="L963" class="ln">   963&nbsp;&nbsp;</span>			<span class="comment">// it indirectly out of the enclosing value.</span>
<span id="L964" class="ln">   964&nbsp;&nbsp;</span>			vMap = reflect.Indirect(addrVal)
<span id="L965" class="ln">   965&nbsp;&nbsp;</span>
<span id="L966" class="ln">   966&nbsp;&nbsp;</span>			if squash {
<span id="L967" class="ln">   967&nbsp;&nbsp;</span>				for _, k := range vMap.MapKeys() {
<span id="L968" class="ln">   968&nbsp;&nbsp;</span>					valMap.SetMapIndex(k, vMap.MapIndex(k))
<span id="L969" class="ln">   969&nbsp;&nbsp;</span>				}
<span id="L970" class="ln">   970&nbsp;&nbsp;</span>			} else {
<span id="L971" class="ln">   971&nbsp;&nbsp;</span>				valMap.SetMapIndex(reflect.ValueOf(keyName), vMap)
<span id="L972" class="ln">   972&nbsp;&nbsp;</span>			}
<span id="L973" class="ln">   973&nbsp;&nbsp;</span>
<span id="L974" class="ln">   974&nbsp;&nbsp;</span>		default:
<span id="L975" class="ln">   975&nbsp;&nbsp;</span>			valMap.SetMapIndex(reflect.ValueOf(keyName), v)
<span id="L976" class="ln">   976&nbsp;&nbsp;</span>		}
<span id="L977" class="ln">   977&nbsp;&nbsp;</span>	}
<span id="L978" class="ln">   978&nbsp;&nbsp;</span>
<span id="L979" class="ln">   979&nbsp;&nbsp;</span>	if val.CanAddr() {
<span id="L980" class="ln">   980&nbsp;&nbsp;</span>		val.Set(valMap)
<span id="L981" class="ln">   981&nbsp;&nbsp;</span>	}
<span id="L982" class="ln">   982&nbsp;&nbsp;</span>
<span id="L983" class="ln">   983&nbsp;&nbsp;</span>	return nil
<span id="L984" class="ln">   984&nbsp;&nbsp;</span>}
<span id="L985" class="ln">   985&nbsp;&nbsp;</span>
<span id="L986" class="ln">   986&nbsp;&nbsp;</span>func (d *Decoder) decodePtr(name string, data interface{}, val reflect.Value) (bool, error) {
<span id="L987" class="ln">   987&nbsp;&nbsp;</span>	<span class="comment">// If the input data is nil, then we want to just set the output</span>
<span id="L988" class="ln">   988&nbsp;&nbsp;</span>	<span class="comment">// pointer to be nil as well.</span>
<span id="L989" class="ln">   989&nbsp;&nbsp;</span>	isNil := data == nil
<span id="L990" class="ln">   990&nbsp;&nbsp;</span>	if !isNil {
<span id="L991" class="ln">   991&nbsp;&nbsp;</span>		switch v := reflect.Indirect(reflect.ValueOf(data)); v.Kind() {
<span id="L992" class="ln">   992&nbsp;&nbsp;</span>		case reflect.Chan,
<span id="L993" class="ln">   993&nbsp;&nbsp;</span>			reflect.Func,
<span id="L994" class="ln">   994&nbsp;&nbsp;</span>			reflect.Interface,
<span id="L995" class="ln">   995&nbsp;&nbsp;</span>			reflect.Map,
<span id="L996" class="ln">   996&nbsp;&nbsp;</span>			reflect.Ptr,
<span id="L997" class="ln">   997&nbsp;&nbsp;</span>			reflect.Slice:
<span id="L998" class="ln">   998&nbsp;&nbsp;</span>			isNil = v.IsNil()
<span id="L999" class="ln">   999&nbsp;&nbsp;</span>		}
<span id="L1000" class="ln">  1000&nbsp;&nbsp;</span>	}
<span id="L1001" class="ln">  1001&nbsp;&nbsp;</span>	if isNil {
<span id="L1002" class="ln">  1002&nbsp;&nbsp;</span>		if !val.IsNil() &amp;&amp; val.CanSet() {
<span id="L1003" class="ln">  1003&nbsp;&nbsp;</span>			nilValue := reflect.New(val.Type()).Elem()
<span id="L1004" class="ln">  1004&nbsp;&nbsp;</span>			val.Set(nilValue)
<span id="L1005" class="ln">  1005&nbsp;&nbsp;</span>		}
<span id="L1006" class="ln">  1006&nbsp;&nbsp;</span>
<span id="L1007" class="ln">  1007&nbsp;&nbsp;</span>		return true, nil
<span id="L1008" class="ln">  1008&nbsp;&nbsp;</span>	}
<span id="L1009" class="ln">  1009&nbsp;&nbsp;</span>
<span id="L1010" class="ln">  1010&nbsp;&nbsp;</span>	<span class="comment">// Create an element of the concrete (non pointer) type and decode</span>
<span id="L1011" class="ln">  1011&nbsp;&nbsp;</span>	<span class="comment">// into that. Then set the value of the pointer to this type.</span>
<span id="L1012" class="ln">  1012&nbsp;&nbsp;</span>	valType := val.Type()
<span id="L1013" class="ln">  1013&nbsp;&nbsp;</span>	valElemType := valType.Elem()
<span id="L1014" class="ln">  1014&nbsp;&nbsp;</span>	if val.CanSet() {
<span id="L1015" class="ln">  1015&nbsp;&nbsp;</span>		realVal := val
<span id="L1016" class="ln">  1016&nbsp;&nbsp;</span>		if realVal.IsNil() || d.config.ZeroFields {
<span id="L1017" class="ln">  1017&nbsp;&nbsp;</span>			realVal = reflect.New(valElemType)
<span id="L1018" class="ln">  1018&nbsp;&nbsp;</span>		}
<span id="L1019" class="ln">  1019&nbsp;&nbsp;</span>
<span id="L1020" class="ln">  1020&nbsp;&nbsp;</span>		if err := d.decode(name, data, reflect.Indirect(realVal)); err != nil {
<span id="L1021" class="ln">  1021&nbsp;&nbsp;</span>			return false, err
<span id="L1022" class="ln">  1022&nbsp;&nbsp;</span>		}
<span id="L1023" class="ln">  1023&nbsp;&nbsp;</span>
<span id="L1024" class="ln">  1024&nbsp;&nbsp;</span>		val.Set(realVal)
<span id="L1025" class="ln">  1025&nbsp;&nbsp;</span>	} else {
<span id="L1026" class="ln">  1026&nbsp;&nbsp;</span>		if err := d.decode(name, data, reflect.Indirect(val)); err != nil {
<span id="L1027" class="ln">  1027&nbsp;&nbsp;</span>			return false, err
<span id="L1028" class="ln">  1028&nbsp;&nbsp;</span>		}
<span id="L1029" class="ln">  1029&nbsp;&nbsp;</span>	}
<span id="L1030" class="ln">  1030&nbsp;&nbsp;</span>	return false, nil
<span id="L1031" class="ln">  1031&nbsp;&nbsp;</span>}
<span id="L1032" class="ln">  1032&nbsp;&nbsp;</span>
<span id="L1033" class="ln">  1033&nbsp;&nbsp;</span>func (d *Decoder) decodeFunc(name string, data interface{}, val reflect.Value) error {
<span id="L1034" class="ln">  1034&nbsp;&nbsp;</span>	<span class="comment">// Create an element of the concrete (non pointer) type and decode</span>
<span id="L1035" class="ln">  1035&nbsp;&nbsp;</span>	<span class="comment">// into that. Then set the value of the pointer to this type.</span>
<span id="L1036" class="ln">  1036&nbsp;&nbsp;</span>	dataVal := reflect.Indirect(reflect.ValueOf(data))
<span id="L1037" class="ln">  1037&nbsp;&nbsp;</span>	if val.Type() != dataVal.Type() {
<span id="L1038" class="ln">  1038&nbsp;&nbsp;</span>		return fmt.Errorf(
<span id="L1039" class="ln">  1039&nbsp;&nbsp;</span>			&#34;&#39;%s&#39; expected type &#39;%s&#39;, got unconvertible type &#39;%s&#39;, value: &#39;%v&#39;&#34;,
<span id="L1040" class="ln">  1040&nbsp;&nbsp;</span>			name, val.Type(), dataVal.Type(), data)
<span id="L1041" class="ln">  1041&nbsp;&nbsp;</span>	}
<span id="L1042" class="ln">  1042&nbsp;&nbsp;</span>	val.Set(dataVal)
<span id="L1043" class="ln">  1043&nbsp;&nbsp;</span>	return nil
<span id="L1044" class="ln">  1044&nbsp;&nbsp;</span>}
<span id="L1045" class="ln">  1045&nbsp;&nbsp;</span>
<span id="L1046" class="ln">  1046&nbsp;&nbsp;</span>func (d *Decoder) decodeSlice(name string, data interface{}, val reflect.Value) error {
<span id="L1047" class="ln">  1047&nbsp;&nbsp;</span>	dataVal := reflect.Indirect(reflect.ValueOf(data))
<span id="L1048" class="ln">  1048&nbsp;&nbsp;</span>	dataValKind := dataVal.Kind()
<span id="L1049" class="ln">  1049&nbsp;&nbsp;</span>	valType := val.Type()
<span id="L1050" class="ln">  1050&nbsp;&nbsp;</span>	valElemType := valType.Elem()
<span id="L1051" class="ln">  1051&nbsp;&nbsp;</span>	sliceType := reflect.SliceOf(valElemType)
<span id="L1052" class="ln">  1052&nbsp;&nbsp;</span>
<span id="L1053" class="ln">  1053&nbsp;&nbsp;</span>	<span class="comment">// If we have a non array/slice type then we first attempt to convert.</span>
<span id="L1054" class="ln">  1054&nbsp;&nbsp;</span>	if dataValKind != reflect.Array &amp;&amp; dataValKind != reflect.Slice {
<span id="L1055" class="ln">  1055&nbsp;&nbsp;</span>		if d.config.WeaklyTypedInput {
<span id="L1056" class="ln">  1056&nbsp;&nbsp;</span>			switch {
<span id="L1057" class="ln">  1057&nbsp;&nbsp;</span>			<span class="comment">// Slice and array we use the normal logic</span>
<span id="L1058" class="ln">  1058&nbsp;&nbsp;</span>			case dataValKind == reflect.Slice, dataValKind == reflect.Array:
<span id="L1059" class="ln">  1059&nbsp;&nbsp;</span>				break
<span id="L1060" class="ln">  1060&nbsp;&nbsp;</span>
<span id="L1061" class="ln">  1061&nbsp;&nbsp;</span>			<span class="comment">// Empty maps turn into empty slices</span>
<span id="L1062" class="ln">  1062&nbsp;&nbsp;</span>			case dataValKind == reflect.Map:
<span id="L1063" class="ln">  1063&nbsp;&nbsp;</span>				if dataVal.Len() == 0 {
<span id="L1064" class="ln">  1064&nbsp;&nbsp;</span>					val.Set(reflect.MakeSlice(sliceType, 0, 0))
<span id="L1065" class="ln">  1065&nbsp;&nbsp;</span>					return nil
<span id="L1066" class="ln">  1066&nbsp;&nbsp;</span>				}
<span id="L1067" class="ln">  1067&nbsp;&nbsp;</span>				<span class="comment">// Create slice of maps of other sizes</span>
<span id="L1068" class="ln">  1068&nbsp;&nbsp;</span>				return d.decodeSlice(name, []interface{}{data}, val)
<span id="L1069" class="ln">  1069&nbsp;&nbsp;</span>
<span id="L1070" class="ln">  1070&nbsp;&nbsp;</span>			case dataValKind == reflect.String &amp;&amp; valElemType.Kind() == reflect.Uint8:
<span id="L1071" class="ln">  1071&nbsp;&nbsp;</span>				return d.decodeSlice(name, []byte(dataVal.String()), val)
<span id="L1072" class="ln">  1072&nbsp;&nbsp;</span>
<span id="L1073" class="ln">  1073&nbsp;&nbsp;</span>			<span class="comment">// All other types we try to convert to the slice type</span>
<span id="L1074" class="ln">  1074&nbsp;&nbsp;</span>			<span class="comment">// and &#34;lift&#34; it into it. i.e. a string becomes a string slice.</span>
<span id="L1075" class="ln">  1075&nbsp;&nbsp;</span>			default:
<span id="L1076" class="ln">  1076&nbsp;&nbsp;</span>				<span class="comment">// Just re-try this function with data as a slice.</span>
<span id="L1077" class="ln">  1077&nbsp;&nbsp;</span>				return d.decodeSlice(name, []interface{}{data}, val)
<span id="L1078" class="ln">  1078&nbsp;&nbsp;</span>			}
<span id="L1079" class="ln">  1079&nbsp;&nbsp;</span>		}
<span id="L1080" class="ln">  1080&nbsp;&nbsp;</span>
<span id="L1081" class="ln">  1081&nbsp;&nbsp;</span>		return fmt.Errorf(
<span id="L1082" class="ln">  1082&nbsp;&nbsp;</span>			&#34;&#39;%s&#39;: source data must be an array or slice, got %s&#34;, name, dataValKind)
<span id="L1083" class="ln">  1083&nbsp;&nbsp;</span>	}
<span id="L1084" class="ln">  1084&nbsp;&nbsp;</span>
<span id="L1085" class="ln">  1085&nbsp;&nbsp;</span>	<span class="comment">// If the input value is nil, then don&#39;t allocate since empty != nil</span>
<span id="L1086" class="ln">  1086&nbsp;&nbsp;</span>	if dataVal.IsNil() {
<span id="L1087" class="ln">  1087&nbsp;&nbsp;</span>		return nil
<span id="L1088" class="ln">  1088&nbsp;&nbsp;</span>	}
<span id="L1089" class="ln">  1089&nbsp;&nbsp;</span>
<span id="L1090" class="ln">  1090&nbsp;&nbsp;</span>	valSlice := val
<span id="L1091" class="ln">  1091&nbsp;&nbsp;</span>	if valSlice.IsNil() || d.config.ZeroFields {
<span id="L1092" class="ln">  1092&nbsp;&nbsp;</span>		<span class="comment">// Make a new slice to hold our result, same size as the original data.</span>
<span id="L1093" class="ln">  1093&nbsp;&nbsp;</span>		valSlice = reflect.MakeSlice(sliceType, dataVal.Len(), dataVal.Len())
<span id="L1094" class="ln">  1094&nbsp;&nbsp;</span>	}
<span id="L1095" class="ln">  1095&nbsp;&nbsp;</span>
<span id="L1096" class="ln">  1096&nbsp;&nbsp;</span>	<span class="comment">// Accumulate any errors</span>
<span id="L1097" class="ln">  1097&nbsp;&nbsp;</span>	errors := make([]string, 0)
<span id="L1098" class="ln">  1098&nbsp;&nbsp;</span>
<span id="L1099" class="ln">  1099&nbsp;&nbsp;</span>	for i := 0; i &lt; dataVal.Len(); i++ {
<span id="L1100" class="ln">  1100&nbsp;&nbsp;</span>		currentData := dataVal.Index(i).Interface()
<span id="L1101" class="ln">  1101&nbsp;&nbsp;</span>		for valSlice.Len() &lt;= i {
<span id="L1102" class="ln">  1102&nbsp;&nbsp;</span>			valSlice = reflect.Append(valSlice, reflect.Zero(valElemType))
<span id="L1103" class="ln">  1103&nbsp;&nbsp;</span>		}
<span id="L1104" class="ln">  1104&nbsp;&nbsp;</span>		currentField := valSlice.Index(i)
<span id="L1105" class="ln">  1105&nbsp;&nbsp;</span>
<span id="L1106" class="ln">  1106&nbsp;&nbsp;</span>		fieldName := name + &#34;[&#34; + strconv.Itoa(i) + &#34;]&#34;
<span id="L1107" class="ln">  1107&nbsp;&nbsp;</span>		if err := d.decode(fieldName, currentData, currentField); err != nil {
<span id="L1108" class="ln">  1108&nbsp;&nbsp;</span>			errors = appendErrors(errors, err)
<span id="L1109" class="ln">  1109&nbsp;&nbsp;</span>		}
<span id="L1110" class="ln">  1110&nbsp;&nbsp;</span>	}
<span id="L1111" class="ln">  1111&nbsp;&nbsp;</span>
<span id="L1112" class="ln">  1112&nbsp;&nbsp;</span>	<span class="comment">// Finally, set the value to the slice we built up</span>
<span id="L1113" class="ln">  1113&nbsp;&nbsp;</span>	val.Set(valSlice)
<span id="L1114" class="ln">  1114&nbsp;&nbsp;</span>
<span id="L1115" class="ln">  1115&nbsp;&nbsp;</span>	<span class="comment">// If there were errors, we return those</span>
<span id="L1116" class="ln">  1116&nbsp;&nbsp;</span>	if len(errors) &gt; 0 {
<span id="L1117" class="ln">  1117&nbsp;&nbsp;</span>		return &amp;Error{errors}
<span id="L1118" class="ln">  1118&nbsp;&nbsp;</span>	}
<span id="L1119" class="ln">  1119&nbsp;&nbsp;</span>
<span id="L1120" class="ln">  1120&nbsp;&nbsp;</span>	return nil
<span id="L1121" class="ln">  1121&nbsp;&nbsp;</span>}
<span id="L1122" class="ln">  1122&nbsp;&nbsp;</span>
<span id="L1123" class="ln">  1123&nbsp;&nbsp;</span>func (d *Decoder) decodeArray(name string, data interface{}, val reflect.Value) error {
<span id="L1124" class="ln">  1124&nbsp;&nbsp;</span>	dataVal := reflect.Indirect(reflect.ValueOf(data))
<span id="L1125" class="ln">  1125&nbsp;&nbsp;</span>	dataValKind := dataVal.Kind()
<span id="L1126" class="ln">  1126&nbsp;&nbsp;</span>	valType := val.Type()
<span id="L1127" class="ln">  1127&nbsp;&nbsp;</span>	valElemType := valType.Elem()
<span id="L1128" class="ln">  1128&nbsp;&nbsp;</span>	arrayType := reflect.ArrayOf(valType.Len(), valElemType)
<span id="L1129" class="ln">  1129&nbsp;&nbsp;</span>
<span id="L1130" class="ln">  1130&nbsp;&nbsp;</span>	valArray := val
<span id="L1131" class="ln">  1131&nbsp;&nbsp;</span>
<span id="L1132" class="ln">  1132&nbsp;&nbsp;</span>	if valArray.Interface() == reflect.Zero(valArray.Type()).Interface() || d.config.ZeroFields {
<span id="L1133" class="ln">  1133&nbsp;&nbsp;</span>		<span class="comment">// Check input type</span>
<span id="L1134" class="ln">  1134&nbsp;&nbsp;</span>		if dataValKind != reflect.Array &amp;&amp; dataValKind != reflect.Slice {
<span id="L1135" class="ln">  1135&nbsp;&nbsp;</span>			if d.config.WeaklyTypedInput {
<span id="L1136" class="ln">  1136&nbsp;&nbsp;</span>				switch {
<span id="L1137" class="ln">  1137&nbsp;&nbsp;</span>				<span class="comment">// Empty maps turn into empty arrays</span>
<span id="L1138" class="ln">  1138&nbsp;&nbsp;</span>				case dataValKind == reflect.Map:
<span id="L1139" class="ln">  1139&nbsp;&nbsp;</span>					if dataVal.Len() == 0 {
<span id="L1140" class="ln">  1140&nbsp;&nbsp;</span>						val.Set(reflect.Zero(arrayType))
<span id="L1141" class="ln">  1141&nbsp;&nbsp;</span>						return nil
<span id="L1142" class="ln">  1142&nbsp;&nbsp;</span>					}
<span id="L1143" class="ln">  1143&nbsp;&nbsp;</span>
<span id="L1144" class="ln">  1144&nbsp;&nbsp;</span>				<span class="comment">// All other types we try to convert to the array type</span>
<span id="L1145" class="ln">  1145&nbsp;&nbsp;</span>				<span class="comment">// and &#34;lift&#34; it into it. i.e. a string becomes a string array.</span>
<span id="L1146" class="ln">  1146&nbsp;&nbsp;</span>				default:
<span id="L1147" class="ln">  1147&nbsp;&nbsp;</span>					<span class="comment">// Just re-try this function with data as a slice.</span>
<span id="L1148" class="ln">  1148&nbsp;&nbsp;</span>					return d.decodeArray(name, []interface{}{data}, val)
<span id="L1149" class="ln">  1149&nbsp;&nbsp;</span>				}
<span id="L1150" class="ln">  1150&nbsp;&nbsp;</span>			}
<span id="L1151" class="ln">  1151&nbsp;&nbsp;</span>
<span id="L1152" class="ln">  1152&nbsp;&nbsp;</span>			return fmt.Errorf(
<span id="L1153" class="ln">  1153&nbsp;&nbsp;</span>				&#34;&#39;%s&#39;: source data must be an array or slice, got %s&#34;, name, dataValKind)
<span id="L1154" class="ln">  1154&nbsp;&nbsp;</span>
<span id="L1155" class="ln">  1155&nbsp;&nbsp;</span>		}
<span id="L1156" class="ln">  1156&nbsp;&nbsp;</span>		if dataVal.Len() &gt; arrayType.Len() {
<span id="L1157" class="ln">  1157&nbsp;&nbsp;</span>			return fmt.Errorf(
<span id="L1158" class="ln">  1158&nbsp;&nbsp;</span>				&#34;&#39;%s&#39;: expected source data to have length less or equal to %d, got %d&#34;, name, arrayType.Len(), dataVal.Len())
<span id="L1159" class="ln">  1159&nbsp;&nbsp;</span>
<span id="L1160" class="ln">  1160&nbsp;&nbsp;</span>		}
<span id="L1161" class="ln">  1161&nbsp;&nbsp;</span>
<span id="L1162" class="ln">  1162&nbsp;&nbsp;</span>		<span class="comment">// Make a new array to hold our result, same size as the original data.</span>
<span id="L1163" class="ln">  1163&nbsp;&nbsp;</span>		valArray = reflect.New(arrayType).Elem()
<span id="L1164" class="ln">  1164&nbsp;&nbsp;</span>	}
<span id="L1165" class="ln">  1165&nbsp;&nbsp;</span>
<span id="L1166" class="ln">  1166&nbsp;&nbsp;</span>	<span class="comment">// Accumulate any errors</span>
<span id="L1167" class="ln">  1167&nbsp;&nbsp;</span>	errors := make([]string, 0)
<span id="L1168" class="ln">  1168&nbsp;&nbsp;</span>
<span id="L1169" class="ln">  1169&nbsp;&nbsp;</span>	for i := 0; i &lt; dataVal.Len(); i++ {
<span id="L1170" class="ln">  1170&nbsp;&nbsp;</span>		currentData := dataVal.Index(i).Interface()
<span id="L1171" class="ln">  1171&nbsp;&nbsp;</span>		currentField := valArray.Index(i)
<span id="L1172" class="ln">  1172&nbsp;&nbsp;</span>
<span id="L1173" class="ln">  1173&nbsp;&nbsp;</span>		fieldName := name + &#34;[&#34; + strconv.Itoa(i) + &#34;]&#34;
<span id="L1174" class="ln">  1174&nbsp;&nbsp;</span>		if err := d.decode(fieldName, currentData, currentField); err != nil {
<span id="L1175" class="ln">  1175&nbsp;&nbsp;</span>			errors = appendErrors(errors, err)
<span id="L1176" class="ln">  1176&nbsp;&nbsp;</span>		}
<span id="L1177" class="ln">  1177&nbsp;&nbsp;</span>	}
<span id="L1178" class="ln">  1178&nbsp;&nbsp;</span>
<span id="L1179" class="ln">  1179&nbsp;&nbsp;</span>	<span class="comment">// Finally, set the value to the array we built up</span>
<span id="L1180" class="ln">  1180&nbsp;&nbsp;</span>	val.Set(valArray)
<span id="L1181" class="ln">  1181&nbsp;&nbsp;</span>
<span id="L1182" class="ln">  1182&nbsp;&nbsp;</span>	<span class="comment">// If there were errors, we return those</span>
<span id="L1183" class="ln">  1183&nbsp;&nbsp;</span>	if len(errors) &gt; 0 {
<span id="L1184" class="ln">  1184&nbsp;&nbsp;</span>		return &amp;Error{errors}
<span id="L1185" class="ln">  1185&nbsp;&nbsp;</span>	}
<span id="L1186" class="ln">  1186&nbsp;&nbsp;</span>
<span id="L1187" class="ln">  1187&nbsp;&nbsp;</span>	return nil
<span id="L1188" class="ln">  1188&nbsp;&nbsp;</span>}
<span id="L1189" class="ln">  1189&nbsp;&nbsp;</span>
<span id="L1190" class="ln">  1190&nbsp;&nbsp;</span>func (d *Decoder) decodeStruct(name string, data interface{}, val reflect.Value) error {
<span id="L1191" class="ln">  1191&nbsp;&nbsp;</span>	dataVal := reflect.Indirect(reflect.ValueOf(data))
<span id="L1192" class="ln">  1192&nbsp;&nbsp;</span>
<span id="L1193" class="ln">  1193&nbsp;&nbsp;</span>	<span class="comment">// If the type of the value to write to and the data match directly,</span>
<span id="L1194" class="ln">  1194&nbsp;&nbsp;</span>	<span class="comment">// then we just set it directly instead of recursing into the structure.</span>
<span id="L1195" class="ln">  1195&nbsp;&nbsp;</span>	if dataVal.Type() == val.Type() {
<span id="L1196" class="ln">  1196&nbsp;&nbsp;</span>		val.Set(dataVal)
<span id="L1197" class="ln">  1197&nbsp;&nbsp;</span>		return nil
<span id="L1198" class="ln">  1198&nbsp;&nbsp;</span>	}
<span id="L1199" class="ln">  1199&nbsp;&nbsp;</span>
<span id="L1200" class="ln">  1200&nbsp;&nbsp;</span>	dataValKind := dataVal.Kind()
<span id="L1201" class="ln">  1201&nbsp;&nbsp;</span>	switch dataValKind {
<span id="L1202" class="ln">  1202&nbsp;&nbsp;</span>	case reflect.Map:
<span id="L1203" class="ln">  1203&nbsp;&nbsp;</span>		return d.decodeStructFromMap(name, dataVal, val)
<span id="L1204" class="ln">  1204&nbsp;&nbsp;</span>
<span id="L1205" class="ln">  1205&nbsp;&nbsp;</span>	case reflect.Struct:
<span id="L1206" class="ln">  1206&nbsp;&nbsp;</span>		<span class="comment">// Not the most efficient way to do this but we can optimize later if</span>
<span id="L1207" class="ln">  1207&nbsp;&nbsp;</span>		<span class="comment">// we want to. To convert from struct to struct we go to map first</span>
<span id="L1208" class="ln">  1208&nbsp;&nbsp;</span>		<span class="comment">// as an intermediary.</span>
<span id="L1209" class="ln">  1209&nbsp;&nbsp;</span>
<span id="L1210" class="ln">  1210&nbsp;&nbsp;</span>		<span class="comment">// Make a new map to hold our result</span>
<span id="L1211" class="ln">  1211&nbsp;&nbsp;</span>		mapType := reflect.TypeOf((map[string]interface{})(nil))
<span id="L1212" class="ln">  1212&nbsp;&nbsp;</span>		mval := reflect.MakeMap(mapType)
<span id="L1213" class="ln">  1213&nbsp;&nbsp;</span>
<span id="L1214" class="ln">  1214&nbsp;&nbsp;</span>		<span class="comment">// Creating a pointer to a map so that other methods can completely</span>
<span id="L1215" class="ln">  1215&nbsp;&nbsp;</span>		<span class="comment">// overwrite the map if need be (looking at you decodeMapFromMap). The</span>
<span id="L1216" class="ln">  1216&nbsp;&nbsp;</span>		<span class="comment">// indirection allows the underlying map to be settable (CanSet() == true)</span>
<span id="L1217" class="ln">  1217&nbsp;&nbsp;</span>		<span class="comment">// where as reflect.MakeMap returns an unsettable map.</span>
<span id="L1218" class="ln">  1218&nbsp;&nbsp;</span>		addrVal := reflect.New(mval.Type())
<span id="L1219" class="ln">  1219&nbsp;&nbsp;</span>
<span id="L1220" class="ln">  1220&nbsp;&nbsp;</span>		reflect.Indirect(addrVal).Set(mval)
<span id="L1221" class="ln">  1221&nbsp;&nbsp;</span>		if err := d.decodeMapFromStruct(name, dataVal, reflect.Indirect(addrVal), mval); err != nil {
<span id="L1222" class="ln">  1222&nbsp;&nbsp;</span>			return err
<span id="L1223" class="ln">  1223&nbsp;&nbsp;</span>		}
<span id="L1224" class="ln">  1224&nbsp;&nbsp;</span>
<span id="L1225" class="ln">  1225&nbsp;&nbsp;</span>		result := d.decodeStructFromMap(name, reflect.Indirect(addrVal), val)
<span id="L1226" class="ln">  1226&nbsp;&nbsp;</span>		return result
<span id="L1227" class="ln">  1227&nbsp;&nbsp;</span>
<span id="L1228" class="ln">  1228&nbsp;&nbsp;</span>	default:
<span id="L1229" class="ln">  1229&nbsp;&nbsp;</span>		return fmt.Errorf(&#34;&#39;%s&#39; expected a map, got &#39;%s&#39;&#34;, name, dataVal.Kind())
<span id="L1230" class="ln">  1230&nbsp;&nbsp;</span>	}
<span id="L1231" class="ln">  1231&nbsp;&nbsp;</span>}
<span id="L1232" class="ln">  1232&nbsp;&nbsp;</span>
<span id="L1233" class="ln">  1233&nbsp;&nbsp;</span>func (d *Decoder) decodeStructFromMap(name string, dataVal, val reflect.Value) error {
<span id="L1234" class="ln">  1234&nbsp;&nbsp;</span>	dataValType := dataVal.Type()
<span id="L1235" class="ln">  1235&nbsp;&nbsp;</span>	if kind := dataValType.Key().Kind(); kind != reflect.String &amp;&amp; kind != reflect.Interface {
<span id="L1236" class="ln">  1236&nbsp;&nbsp;</span>		return fmt.Errorf(
<span id="L1237" class="ln">  1237&nbsp;&nbsp;</span>			&#34;&#39;%s&#39; needs a map with string keys, has &#39;%s&#39; keys&#34;,
<span id="L1238" class="ln">  1238&nbsp;&nbsp;</span>			name, dataValType.Key().Kind())
<span id="L1239" class="ln">  1239&nbsp;&nbsp;</span>	}
<span id="L1240" class="ln">  1240&nbsp;&nbsp;</span>
<span id="L1241" class="ln">  1241&nbsp;&nbsp;</span>	dataValKeys := make(map[reflect.Value]struct{})
<span id="L1242" class="ln">  1242&nbsp;&nbsp;</span>	dataValKeysUnused := make(map[interface{}]struct{})
<span id="L1243" class="ln">  1243&nbsp;&nbsp;</span>	for _, dataValKey := range dataVal.MapKeys() {
<span id="L1244" class="ln">  1244&nbsp;&nbsp;</span>		dataValKeys[dataValKey] = struct{}{}
<span id="L1245" class="ln">  1245&nbsp;&nbsp;</span>		dataValKeysUnused[dataValKey.Interface()] = struct{}{}
<span id="L1246" class="ln">  1246&nbsp;&nbsp;</span>	}
<span id="L1247" class="ln">  1247&nbsp;&nbsp;</span>
<span id="L1248" class="ln">  1248&nbsp;&nbsp;</span>	errors := make([]string, 0)
<span id="L1249" class="ln">  1249&nbsp;&nbsp;</span>
<span id="L1250" class="ln">  1250&nbsp;&nbsp;</span>	<span class="comment">// This slice will keep track of all the structs we&#39;ll be decoding.</span>
<span id="L1251" class="ln">  1251&nbsp;&nbsp;</span>	<span class="comment">// There can be more than one struct if there are embedded structs</span>
<span id="L1252" class="ln">  1252&nbsp;&nbsp;</span>	<span class="comment">// that are squashed.</span>
<span id="L1253" class="ln">  1253&nbsp;&nbsp;</span>	structs := make([]reflect.Value, 1, 5)
<span id="L1254" class="ln">  1254&nbsp;&nbsp;</span>	structs[0] = val
<span id="L1255" class="ln">  1255&nbsp;&nbsp;</span>
<span id="L1256" class="ln">  1256&nbsp;&nbsp;</span>	<span class="comment">// Compile the list of all the fields that we&#39;re going to be decoding</span>
<span id="L1257" class="ln">  1257&nbsp;&nbsp;</span>	<span class="comment">// from all the structs.</span>
<span id="L1258" class="ln">  1258&nbsp;&nbsp;</span>	type field struct {
<span id="L1259" class="ln">  1259&nbsp;&nbsp;</span>		field reflect.StructField
<span id="L1260" class="ln">  1260&nbsp;&nbsp;</span>		val   reflect.Value
<span id="L1261" class="ln">  1261&nbsp;&nbsp;</span>	}
<span id="L1262" class="ln">  1262&nbsp;&nbsp;</span>
<span id="L1263" class="ln">  1263&nbsp;&nbsp;</span>	<span class="comment">// remainField is set to a valid field set with the &#34;remain&#34; tag if</span>
<span id="L1264" class="ln">  1264&nbsp;&nbsp;</span>	<span class="comment">// we are keeping track of remaining values.</span>
<span id="L1265" class="ln">  1265&nbsp;&nbsp;</span>	var remainField *field
<span id="L1266" class="ln">  1266&nbsp;&nbsp;</span>
<span id="L1267" class="ln">  1267&nbsp;&nbsp;</span>	fields := []field{}
<span id="L1268" class="ln">  1268&nbsp;&nbsp;</span>	for len(structs) &gt; 0 {
<span id="L1269" class="ln">  1269&nbsp;&nbsp;</span>		structVal := structs[0]
<span id="L1270" class="ln">  1270&nbsp;&nbsp;</span>		structs = structs[1:]
<span id="L1271" class="ln">  1271&nbsp;&nbsp;</span>
<span id="L1272" class="ln">  1272&nbsp;&nbsp;</span>		structType := structVal.Type()
<span id="L1273" class="ln">  1273&nbsp;&nbsp;</span>
<span id="L1274" class="ln">  1274&nbsp;&nbsp;</span>		for i := 0; i &lt; structType.NumField(); i++ {
<span id="L1275" class="ln">  1275&nbsp;&nbsp;</span>			fieldType := structType.Field(i)
<span id="L1276" class="ln">  1276&nbsp;&nbsp;</span>			fieldVal := structVal.Field(i)
<span id="L1277" class="ln">  1277&nbsp;&nbsp;</span>			if fieldVal.Kind() == reflect.Ptr &amp;&amp; fieldVal.Elem().Kind() == reflect.Struct {
<span id="L1278" class="ln">  1278&nbsp;&nbsp;</span>				<span class="comment">// Handle embedded struct pointers as embedded structs.</span>
<span id="L1279" class="ln">  1279&nbsp;&nbsp;</span>				fieldVal = fieldVal.Elem()
<span id="L1280" class="ln">  1280&nbsp;&nbsp;</span>			}
<span id="L1281" class="ln">  1281&nbsp;&nbsp;</span>
<span id="L1282" class="ln">  1282&nbsp;&nbsp;</span>			<span class="comment">// If &#34;squash&#34; is specified in the tag, we squash the field down.</span>
<span id="L1283" class="ln">  1283&nbsp;&nbsp;</span>			squash := d.config.Squash &amp;&amp; fieldVal.Kind() == reflect.Struct &amp;&amp; fieldType.Anonymous
<span id="L1284" class="ln">  1284&nbsp;&nbsp;</span>			remain := false
<span id="L1285" class="ln">  1285&nbsp;&nbsp;</span>
<span id="L1286" class="ln">  1286&nbsp;&nbsp;</span>			<span class="comment">// We always parse the tags cause we&#39;re looking for other tags too</span>
<span id="L1287" class="ln">  1287&nbsp;&nbsp;</span>			tagParts := strings.Split(fieldType.Tag.Get(d.config.TagName), &#34;,&#34;)
<span id="L1288" class="ln">  1288&nbsp;&nbsp;</span>			for _, tag := range tagParts[1:] {
<span id="L1289" class="ln">  1289&nbsp;&nbsp;</span>				if tag == &#34;squash&#34; {
<span id="L1290" class="ln">  1290&nbsp;&nbsp;</span>					squash = true
<span id="L1291" class="ln">  1291&nbsp;&nbsp;</span>					break
<span id="L1292" class="ln">  1292&nbsp;&nbsp;</span>				}
<span id="L1293" class="ln">  1293&nbsp;&nbsp;</span>
<span id="L1294" class="ln">  1294&nbsp;&nbsp;</span>				if tag == &#34;remain&#34; {
<span id="L1295" class="ln">  1295&nbsp;&nbsp;</span>					remain = true
<span id="L1296" class="ln">  1296&nbsp;&nbsp;</span>					break
<span id="L1297" class="ln">  1297&nbsp;&nbsp;</span>				}
<span id="L1298" class="ln">  1298&nbsp;&nbsp;</span>			}
<span id="L1299" class="ln">  1299&nbsp;&nbsp;</span>
<span id="L1300" class="ln">  1300&nbsp;&nbsp;</span>			if squash {
<span id="L1301" class="ln">  1301&nbsp;&nbsp;</span>				if fieldVal.Kind() != reflect.Struct {
<span id="L1302" class="ln">  1302&nbsp;&nbsp;</span>					errors = appendErrors(errors,
<span id="L1303" class="ln">  1303&nbsp;&nbsp;</span>						fmt.Errorf(&#34;%s: unsupported type for squash: %s&#34;, fieldType.Name, fieldVal.Kind()))
<span id="L1304" class="ln">  1304&nbsp;&nbsp;</span>				} else {
<span id="L1305" class="ln">  1305&nbsp;&nbsp;</span>					structs = append(structs, fieldVal)
<span id="L1306" class="ln">  1306&nbsp;&nbsp;</span>				}
<span id="L1307" class="ln">  1307&nbsp;&nbsp;</span>				continue
<span id="L1308" class="ln">  1308&nbsp;&nbsp;</span>			}
<span id="L1309" class="ln">  1309&nbsp;&nbsp;</span>
<span id="L1310" class="ln">  1310&nbsp;&nbsp;</span>			<span class="comment">// Build our field</span>
<span id="L1311" class="ln">  1311&nbsp;&nbsp;</span>			if remain {
<span id="L1312" class="ln">  1312&nbsp;&nbsp;</span>				remainField = &amp;field{fieldType, fieldVal}
<span id="L1313" class="ln">  1313&nbsp;&nbsp;</span>			} else {
<span id="L1314" class="ln">  1314&nbsp;&nbsp;</span>				<span class="comment">// Normal struct field, store it away</span>
<span id="L1315" class="ln">  1315&nbsp;&nbsp;</span>				fields = append(fields, field{fieldType, fieldVal})
<span id="L1316" class="ln">  1316&nbsp;&nbsp;</span>			}
<span id="L1317" class="ln">  1317&nbsp;&nbsp;</span>		}
<span id="L1318" class="ln">  1318&nbsp;&nbsp;</span>	}
<span id="L1319" class="ln">  1319&nbsp;&nbsp;</span>
<span id="L1320" class="ln">  1320&nbsp;&nbsp;</span>	<span class="comment">// for fieldType, field := range fields {</span>
<span id="L1321" class="ln">  1321&nbsp;&nbsp;</span>	for _, f := range fields {
<span id="L1322" class="ln">  1322&nbsp;&nbsp;</span>		field, fieldValue := f.field, f.val
<span id="L1323" class="ln">  1323&nbsp;&nbsp;</span>		fieldName := field.Name
<span id="L1324" class="ln">  1324&nbsp;&nbsp;</span>
<span id="L1325" class="ln">  1325&nbsp;&nbsp;</span>		tagValue := field.Tag.Get(d.config.TagName)
<span id="L1326" class="ln">  1326&nbsp;&nbsp;</span>		tagValue = strings.SplitN(tagValue, &#34;,&#34;, 2)[0]
<span id="L1327" class="ln">  1327&nbsp;&nbsp;</span>		if tagValue != &#34;&#34; {
<span id="L1328" class="ln">  1328&nbsp;&nbsp;</span>			fieldName = tagValue
<span id="L1329" class="ln">  1329&nbsp;&nbsp;</span>		}
<span id="L1330" class="ln">  1330&nbsp;&nbsp;</span>
<span id="L1331" class="ln">  1331&nbsp;&nbsp;</span>		rawMapKey := reflect.ValueOf(fieldName)
<span id="L1332" class="ln">  1332&nbsp;&nbsp;</span>		rawMapVal := dataVal.MapIndex(rawMapKey)
<span id="L1333" class="ln">  1333&nbsp;&nbsp;</span>		if !rawMapVal.IsValid() {
<span id="L1334" class="ln">  1334&nbsp;&nbsp;</span>			<span class="comment">// Do a slower search by iterating over each key and</span>
<span id="L1335" class="ln">  1335&nbsp;&nbsp;</span>			<span class="comment">// doing case-insensitive search.</span>
<span id="L1336" class="ln">  1336&nbsp;&nbsp;</span>			for dataValKey := range dataValKeys {
<span id="L1337" class="ln">  1337&nbsp;&nbsp;</span>				mK, ok := dataValKey.Interface().(string)
<span id="L1338" class="ln">  1338&nbsp;&nbsp;</span>				if !ok {
<span id="L1339" class="ln">  1339&nbsp;&nbsp;</span>					<span class="comment">// Not a string key</span>
<span id="L1340" class="ln">  1340&nbsp;&nbsp;</span>					continue
<span id="L1341" class="ln">  1341&nbsp;&nbsp;</span>				}
<span id="L1342" class="ln">  1342&nbsp;&nbsp;</span>
<span id="L1343" class="ln">  1343&nbsp;&nbsp;</span>				if strings.EqualFold(mK, fieldName) {
<span id="L1344" class="ln">  1344&nbsp;&nbsp;</span>					rawMapKey = dataValKey
<span id="L1345" class="ln">  1345&nbsp;&nbsp;</span>					rawMapVal = dataVal.MapIndex(dataValKey)
<span id="L1346" class="ln">  1346&nbsp;&nbsp;</span>					break
<span id="L1347" class="ln">  1347&nbsp;&nbsp;</span>				}
<span id="L1348" class="ln">  1348&nbsp;&nbsp;</span>			}
<span id="L1349" class="ln">  1349&nbsp;&nbsp;</span>
<span id="L1350" class="ln">  1350&nbsp;&nbsp;</span>			if !rawMapVal.IsValid() {
<span id="L1351" class="ln">  1351&nbsp;&nbsp;</span>				<span class="comment">// There was no matching key in the map for the value in</span>
<span id="L1352" class="ln">  1352&nbsp;&nbsp;</span>				<span class="comment">// the struct. Just ignore.</span>
<span id="L1353" class="ln">  1353&nbsp;&nbsp;</span>				continue
<span id="L1354" class="ln">  1354&nbsp;&nbsp;</span>			}
<span id="L1355" class="ln">  1355&nbsp;&nbsp;</span>		}
<span id="L1356" class="ln">  1356&nbsp;&nbsp;</span>
<span id="L1357" class="ln">  1357&nbsp;&nbsp;</span>		if !fieldValue.IsValid() {
<span id="L1358" class="ln">  1358&nbsp;&nbsp;</span>			<span class="comment">// This should never happen</span>
<span id="L1359" class="ln">  1359&nbsp;&nbsp;</span>			panic(&#34;field is not valid&#34;)
<span id="L1360" class="ln">  1360&nbsp;&nbsp;</span>		}
<span id="L1361" class="ln">  1361&nbsp;&nbsp;</span>
<span id="L1362" class="ln">  1362&nbsp;&nbsp;</span>		<span class="comment">// If we can&#39;t set the field, then it is unexported or something,</span>
<span id="L1363" class="ln">  1363&nbsp;&nbsp;</span>		<span class="comment">// and we just continue onwards.</span>
<span id="L1364" class="ln">  1364&nbsp;&nbsp;</span>		if !fieldValue.CanSet() {
<span id="L1365" class="ln">  1365&nbsp;&nbsp;</span>			continue
<span id="L1366" class="ln">  1366&nbsp;&nbsp;</span>		}
<span id="L1367" class="ln">  1367&nbsp;&nbsp;</span>
<span id="L1368" class="ln">  1368&nbsp;&nbsp;</span>		<span class="comment">// Delete the key we&#39;re using from the unused map so we stop tracking</span>
<span id="L1369" class="ln">  1369&nbsp;&nbsp;</span>		delete(dataValKeysUnused, rawMapKey.Interface())
<span id="L1370" class="ln">  1370&nbsp;&nbsp;</span>
<span id="L1371" class="ln">  1371&nbsp;&nbsp;</span>		<span class="comment">// If the name is empty string, then we&#39;re at the root, and we</span>
<span id="L1372" class="ln">  1372&nbsp;&nbsp;</span>		<span class="comment">// don&#39;t dot-join the fields.</span>
<span id="L1373" class="ln">  1373&nbsp;&nbsp;</span>		if name != &#34;&#34; {
<span id="L1374" class="ln">  1374&nbsp;&nbsp;</span>			fieldName = name + &#34;.&#34; + fieldName
<span id="L1375" class="ln">  1375&nbsp;&nbsp;</span>		}
<span id="L1376" class="ln">  1376&nbsp;&nbsp;</span>
<span id="L1377" class="ln">  1377&nbsp;&nbsp;</span>		if err := d.decode(fieldName, rawMapVal.Interface(), fieldValue); err != nil {
<span id="L1378" class="ln">  1378&nbsp;&nbsp;</span>			errors = appendErrors(errors, err)
<span id="L1379" class="ln">  1379&nbsp;&nbsp;</span>		}
<span id="L1380" class="ln">  1380&nbsp;&nbsp;</span>	}
<span id="L1381" class="ln">  1381&nbsp;&nbsp;</span>
<span id="L1382" class="ln">  1382&nbsp;&nbsp;</span>	<span class="comment">// If we have a &#34;remain&#34;-tagged field and we have unused keys then</span>
<span id="L1383" class="ln">  1383&nbsp;&nbsp;</span>	<span class="comment">// we put the unused keys directly into the remain field.</span>
<span id="L1384" class="ln">  1384&nbsp;&nbsp;</span>	if remainField != nil &amp;&amp; len(dataValKeysUnused) &gt; 0 {
<span id="L1385" class="ln">  1385&nbsp;&nbsp;</span>		<span class="comment">// Build a map of only the unused values</span>
<span id="L1386" class="ln">  1386&nbsp;&nbsp;</span>		remain := map[interface{}]interface{}{}
<span id="L1387" class="ln">  1387&nbsp;&nbsp;</span>		for key := range dataValKeysUnused {
<span id="L1388" class="ln">  1388&nbsp;&nbsp;</span>			remain[key] = dataVal.MapIndex(reflect.ValueOf(key)).Interface()
<span id="L1389" class="ln">  1389&nbsp;&nbsp;</span>		}
<span id="L1390" class="ln">  1390&nbsp;&nbsp;</span>
<span id="L1391" class="ln">  1391&nbsp;&nbsp;</span>		<span class="comment">// Decode it as-if we were just decoding this map onto our map.</span>
<span id="L1392" class="ln">  1392&nbsp;&nbsp;</span>		if err := d.decodeMap(name, remain, remainField.val); err != nil {
<span id="L1393" class="ln">  1393&nbsp;&nbsp;</span>			errors = appendErrors(errors, err)
<span id="L1394" class="ln">  1394&nbsp;&nbsp;</span>		}
<span id="L1395" class="ln">  1395&nbsp;&nbsp;</span>
<span id="L1396" class="ln">  1396&nbsp;&nbsp;</span>		<span class="comment">// Set the map to nil so we have none so that the next check will</span>
<span id="L1397" class="ln">  1397&nbsp;&nbsp;</span>		<span class="comment">// not error (ErrorUnused)</span>
<span id="L1398" class="ln">  1398&nbsp;&nbsp;</span>		dataValKeysUnused = nil
<span id="L1399" class="ln">  1399&nbsp;&nbsp;</span>	}
<span id="L1400" class="ln">  1400&nbsp;&nbsp;</span>
<span id="L1401" class="ln">  1401&nbsp;&nbsp;</span>	if d.config.ErrorUnused &amp;&amp; len(dataValKeysUnused) &gt; 0 {
<span id="L1402" class="ln">  1402&nbsp;&nbsp;</span>		keys := make([]string, 0, len(dataValKeysUnused))
<span id="L1403" class="ln">  1403&nbsp;&nbsp;</span>		for rawKey := range dataValKeysUnused {
<span id="L1404" class="ln">  1404&nbsp;&nbsp;</span>			keys = append(keys, rawKey.(string))
<span id="L1405" class="ln">  1405&nbsp;&nbsp;</span>		}
<span id="L1406" class="ln">  1406&nbsp;&nbsp;</span>		sort.Strings(keys)
<span id="L1407" class="ln">  1407&nbsp;&nbsp;</span>
<span id="L1408" class="ln">  1408&nbsp;&nbsp;</span>		err := fmt.Errorf(&#34;&#39;%s&#39; has invalid keys: %s&#34;, name, strings.Join(keys, &#34;, &#34;))
<span id="L1409" class="ln">  1409&nbsp;&nbsp;</span>		errors = appendErrors(errors, err)
<span id="L1410" class="ln">  1410&nbsp;&nbsp;</span>	}
<span id="L1411" class="ln">  1411&nbsp;&nbsp;</span>
<span id="L1412" class="ln">  1412&nbsp;&nbsp;</span>	if len(errors) &gt; 0 {
<span id="L1413" class="ln">  1413&nbsp;&nbsp;</span>		return &amp;Error{errors}
<span id="L1414" class="ln">  1414&nbsp;&nbsp;</span>	}
<span id="L1415" class="ln">  1415&nbsp;&nbsp;</span>
<span id="L1416" class="ln">  1416&nbsp;&nbsp;</span>	<span class="comment">// Add the unused keys to the list of unused keys if we&#39;re tracking metadata</span>
<span id="L1417" class="ln">  1417&nbsp;&nbsp;</span>	if d.config.Metadata != nil {
<span id="L1418" class="ln">  1418&nbsp;&nbsp;</span>		for rawKey := range dataValKeysUnused {
<span id="L1419" class="ln">  1419&nbsp;&nbsp;</span>			key := rawKey.(string)
<span id="L1420" class="ln">  1420&nbsp;&nbsp;</span>			if name != &#34;&#34; {
<span id="L1421" class="ln">  1421&nbsp;&nbsp;</span>				key = name + &#34;.&#34; + key
<span id="L1422" class="ln">  1422&nbsp;&nbsp;</span>			}
<span id="L1423" class="ln">  1423&nbsp;&nbsp;</span>
<span id="L1424" class="ln">  1424&nbsp;&nbsp;</span>			d.config.Metadata.Unused = append(d.config.Metadata.Unused, key)
<span id="L1425" class="ln">  1425&nbsp;&nbsp;</span>		}
<span id="L1426" class="ln">  1426&nbsp;&nbsp;</span>	}
<span id="L1427" class="ln">  1427&nbsp;&nbsp;</span>
<span id="L1428" class="ln">  1428&nbsp;&nbsp;</span>	return nil
<span id="L1429" class="ln">  1429&nbsp;&nbsp;</span>}
<span id="L1430" class="ln">  1430&nbsp;&nbsp;</span>
<span id="L1431" class="ln">  1431&nbsp;&nbsp;</span>func isEmptyValue(v reflect.Value) bool {
<span id="L1432" class="ln">  1432&nbsp;&nbsp;</span>	switch getKind(v) {
<span id="L1433" class="ln">  1433&nbsp;&nbsp;</span>	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
<span id="L1434" class="ln">  1434&nbsp;&nbsp;</span>		return v.Len() == 0
<span id="L1435" class="ln">  1435&nbsp;&nbsp;</span>	case reflect.Bool:
<span id="L1436" class="ln">  1436&nbsp;&nbsp;</span>		return !v.Bool()
<span id="L1437" class="ln">  1437&nbsp;&nbsp;</span>	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
<span id="L1438" class="ln">  1438&nbsp;&nbsp;</span>		return v.Int() == 0
<span id="L1439" class="ln">  1439&nbsp;&nbsp;</span>	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
<span id="L1440" class="ln">  1440&nbsp;&nbsp;</span>		return v.Uint() == 0
<span id="L1441" class="ln">  1441&nbsp;&nbsp;</span>	case reflect.Float32, reflect.Float64:
<span id="L1442" class="ln">  1442&nbsp;&nbsp;</span>		return v.Float() == 0
<span id="L1443" class="ln">  1443&nbsp;&nbsp;</span>	case reflect.Interface, reflect.Ptr:
<span id="L1444" class="ln">  1444&nbsp;&nbsp;</span>		return v.IsNil()
<span id="L1445" class="ln">  1445&nbsp;&nbsp;</span>	}
<span id="L1446" class="ln">  1446&nbsp;&nbsp;</span>	return false
<span id="L1447" class="ln">  1447&nbsp;&nbsp;</span>}
<span id="L1448" class="ln">  1448&nbsp;&nbsp;</span>
<span id="L1449" class="ln">  1449&nbsp;&nbsp;</span>func getKind(val reflect.Value) reflect.Kind {
<span id="L1450" class="ln">  1450&nbsp;&nbsp;</span>	kind := val.Kind()
<span id="L1451" class="ln">  1451&nbsp;&nbsp;</span>
<span id="L1452" class="ln">  1452&nbsp;&nbsp;</span>	switch {
<span id="L1453" class="ln">  1453&nbsp;&nbsp;</span>	case kind &gt;= reflect.Int &amp;&amp; kind &lt;= reflect.Int64:
<span id="L1454" class="ln">  1454&nbsp;&nbsp;</span>		return reflect.Int
<span id="L1455" class="ln">  1455&nbsp;&nbsp;</span>	case kind &gt;= reflect.Uint &amp;&amp; kind &lt;= reflect.Uint64:
<span id="L1456" class="ln">  1456&nbsp;&nbsp;</span>		return reflect.Uint
<span id="L1457" class="ln">  1457&nbsp;&nbsp;</span>	case kind &gt;= reflect.Float32 &amp;&amp; kind &lt;= reflect.Float64:
<span id="L1458" class="ln">  1458&nbsp;&nbsp;</span>		return reflect.Float32
<span id="L1459" class="ln">  1459&nbsp;&nbsp;</span>	default:
<span id="L1460" class="ln">  1460&nbsp;&nbsp;</span>		return kind
<span id="L1461" class="ln">  1461&nbsp;&nbsp;</span>	}
<span id="L1462" class="ln">  1462&nbsp;&nbsp;</span>}
<span id="L1463" class="ln">  1463&nbsp;&nbsp;</span>
</pre><p><a href="/src/github.com/mitchellh/mapstructure/mapstructure.go?m=text">View as plain text</a></p>

<div id="footer">
Build version go1.14.2.<br>
Except as <a href="https://developers.google.com/site-policies#restrictions">noted</a>,
the content of this page is licensed under the
Creative Commons Attribution 3.0 License,
and code is licensed under a <a href="/LICENSE">BSD license</a>.<br>
<a href="/doc/tos.html">Terms of Service</a> |
<a href="http://www.google.com/intl/en/policies/privacy/">Privacy Policy</a>
</div>

</div><!-- .container -->
</div><!-- #page -->
</body>
</html>
