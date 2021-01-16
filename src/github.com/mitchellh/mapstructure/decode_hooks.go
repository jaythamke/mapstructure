<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="theme-color" content="#375EAB">

  <title>src/github.com/mitchellh/mapstructure/decode_hooks.go - Go Documentation Server</title>

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
    <a href="/src">src</a>/<a href="/src/github.com">github.com</a>/<a href="/src/github.com/mitchellh">mitchellh</a>/<a href="/src/github.com/mitchellh/mapstructure">mapstructure</a>/<span class="text-muted">decode_hooks.go</span>
  </h1>





  <h2>
    Documentation: <a href="/pkg/github.com/mitchellh/mapstructure">github.com/mitchellh/mapstructure</a>
  </h2>



<div id="nav"></div>


<script type='text/javascript'>document.ANALYSIS_DATA = null;</script>
<pre><span id="L1" class="ln">     1&nbsp;&nbsp;</span>package mapstructure
<span id="L2" class="ln">     2&nbsp;&nbsp;</span>
<span id="L3" class="ln">     3&nbsp;&nbsp;</span>import (
<span id="L4" class="ln">     4&nbsp;&nbsp;</span>	&#34;encoding&#34;
<span id="L5" class="ln">     5&nbsp;&nbsp;</span>	&#34;errors&#34;
<span id="L6" class="ln">     6&nbsp;&nbsp;</span>	&#34;fmt&#34;
<span id="L7" class="ln">     7&nbsp;&nbsp;</span>	&#34;net&#34;
<span id="L8" class="ln">     8&nbsp;&nbsp;</span>	&#34;reflect&#34;
<span id="L9" class="ln">     9&nbsp;&nbsp;</span>	&#34;strconv&#34;
<span id="L10" class="ln">    10&nbsp;&nbsp;</span>	&#34;strings&#34;
<span id="L11" class="ln">    11&nbsp;&nbsp;</span>	&#34;time&#34;
<span id="L12" class="ln">    12&nbsp;&nbsp;</span>)
<span id="L13" class="ln">    13&nbsp;&nbsp;</span>
<span id="L14" class="ln">    14&nbsp;&nbsp;</span><span class="comment">// typedDecodeHook takes a raw DecodeHookFunc (an interface{}) and turns</span>
<span id="L15" class="ln">    15&nbsp;&nbsp;</span><span class="comment">// it into the proper DecodeHookFunc type, such as DecodeHookFuncType.</span>
<span id="L16" class="ln">    16&nbsp;&nbsp;</span>func typedDecodeHook(h DecodeHookFunc) DecodeHookFunc {
<span id="L17" class="ln">    17&nbsp;&nbsp;</span>	<span class="comment">// Create variables here so we can reference them with the reflect pkg</span>
<span id="L18" class="ln">    18&nbsp;&nbsp;</span>	var f1 DecodeHookFuncType
<span id="L19" class="ln">    19&nbsp;&nbsp;</span>	var f2 DecodeHookFuncKind
<span id="L20" class="ln">    20&nbsp;&nbsp;</span>	var f3 DecodeHookFuncValue
<span id="L21" class="ln">    21&nbsp;&nbsp;</span>
<span id="L22" class="ln">    22&nbsp;&nbsp;</span>	<span class="comment">// Fill in the variables into this interface and the rest is done</span>
<span id="L23" class="ln">    23&nbsp;&nbsp;</span>	<span class="comment">// automatically using the reflect package.</span>
<span id="L24" class="ln">    24&nbsp;&nbsp;</span>	potential := []interface{}{f1, f2, f3}
<span id="L25" class="ln">    25&nbsp;&nbsp;</span>
<span id="L26" class="ln">    26&nbsp;&nbsp;</span>	v := reflect.ValueOf(h)
<span id="L27" class="ln">    27&nbsp;&nbsp;</span>	vt := v.Type()
<span id="L28" class="ln">    28&nbsp;&nbsp;</span>	for _, raw := range potential {
<span id="L29" class="ln">    29&nbsp;&nbsp;</span>		pt := reflect.ValueOf(raw).Type()
<span id="L30" class="ln">    30&nbsp;&nbsp;</span>		if vt.ConvertibleTo(pt) {
<span id="L31" class="ln">    31&nbsp;&nbsp;</span>			return v.Convert(pt).Interface()
<span id="L32" class="ln">    32&nbsp;&nbsp;</span>		}
<span id="L33" class="ln">    33&nbsp;&nbsp;</span>	}
<span id="L34" class="ln">    34&nbsp;&nbsp;</span>
<span id="L35" class="ln">    35&nbsp;&nbsp;</span>	return nil
<span id="L36" class="ln">    36&nbsp;&nbsp;</span>}
<span id="L37" class="ln">    37&nbsp;&nbsp;</span>
<span id="L38" class="ln">    38&nbsp;&nbsp;</span><span class="comment">// DecodeHookExec executes the given decode hook. This should be used</span>
<span id="L39" class="ln">    39&nbsp;&nbsp;</span><span class="comment">// since it&#39;ll naturally degrade to the older backwards compatible DecodeHookFunc</span>
<span id="L40" class="ln">    40&nbsp;&nbsp;</span><span class="comment">// that took reflect.Kind instead of reflect.Type.</span>
<span id="L41" class="ln">    41&nbsp;&nbsp;</span>func DecodeHookExec(
<span id="L42" class="ln">    42&nbsp;&nbsp;</span>	raw DecodeHookFunc,
<span id="L43" class="ln">    43&nbsp;&nbsp;</span>	from reflect.Value, to reflect.Value) (interface{}, error) {
<span id="L44" class="ln">    44&nbsp;&nbsp;</span>
<span id="L45" class="ln">    45&nbsp;&nbsp;</span>	switch f := typedDecodeHook(raw).(type) {
<span id="L46" class="ln">    46&nbsp;&nbsp;</span>	case DecodeHookFuncType:
<span id="L47" class="ln">    47&nbsp;&nbsp;</span>		return f(from.Type(), to.Type(), from.Interface())
<span id="L48" class="ln">    48&nbsp;&nbsp;</span>	case DecodeHookFuncKind:
<span id="L49" class="ln">    49&nbsp;&nbsp;</span>		return f(from.Kind(), to.Kind(), from.Interface())
<span id="L50" class="ln">    50&nbsp;&nbsp;</span>	case DecodeHookFuncValue:
<span id="L51" class="ln">    51&nbsp;&nbsp;</span>		return f(from, to)
<span id="L52" class="ln">    52&nbsp;&nbsp;</span>	default:
<span id="L53" class="ln">    53&nbsp;&nbsp;</span>		return nil, errors.New(&#34;invalid decode hook signature&#34;)
<span id="L54" class="ln">    54&nbsp;&nbsp;</span>	}
<span id="L55" class="ln">    55&nbsp;&nbsp;</span>}
<span id="L56" class="ln">    56&nbsp;&nbsp;</span>
<span id="L57" class="ln">    57&nbsp;&nbsp;</span><span class="comment">// ComposeDecodeHookFunc creates a single DecodeHookFunc that</span>
<span id="L58" class="ln">    58&nbsp;&nbsp;</span><span class="comment">// automatically composes multiple DecodeHookFuncs.</span>
<span id="L59" class="ln">    59&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L60" class="ln">    60&nbsp;&nbsp;</span><span class="comment">// The composed funcs are called in order, with the result of the</span>
<span id="L61" class="ln">    61&nbsp;&nbsp;</span><span class="comment">// previous transformation.</span>
<span id="L62" class="ln">    62&nbsp;&nbsp;</span>func ComposeDecodeHookFunc(fs ...DecodeHookFunc) DecodeHookFunc {
<span id="L63" class="ln">    63&nbsp;&nbsp;</span>	return func(f reflect.Value, t reflect.Value) (interface{}, error) {
<span id="L64" class="ln">    64&nbsp;&nbsp;</span>		var err error
<span id="L65" class="ln">    65&nbsp;&nbsp;</span>		var data interface{}
<span id="L66" class="ln">    66&nbsp;&nbsp;</span>		newFrom := f
<span id="L67" class="ln">    67&nbsp;&nbsp;</span>		for _, f1 := range fs {
<span id="L68" class="ln">    68&nbsp;&nbsp;</span>			data, err = DecodeHookExec(f1, newFrom, t)
<span id="L69" class="ln">    69&nbsp;&nbsp;</span>			if err != nil {
<span id="L70" class="ln">    70&nbsp;&nbsp;</span>				return nil, err
<span id="L71" class="ln">    71&nbsp;&nbsp;</span>			}
<span id="L72" class="ln">    72&nbsp;&nbsp;</span>			newFrom = reflect.ValueOf(data)
<span id="L73" class="ln">    73&nbsp;&nbsp;</span>		}
<span id="L74" class="ln">    74&nbsp;&nbsp;</span>
<span id="L75" class="ln">    75&nbsp;&nbsp;</span>		return data, nil
<span id="L76" class="ln">    76&nbsp;&nbsp;</span>	}
<span id="L77" class="ln">    77&nbsp;&nbsp;</span>}
<span id="L78" class="ln">    78&nbsp;&nbsp;</span>
<span id="L79" class="ln">    79&nbsp;&nbsp;</span><span class="comment">// StringToSliceHookFunc returns a DecodeHookFunc that converts</span>
<span id="L80" class="ln">    80&nbsp;&nbsp;</span><span class="comment">// string to []string by splitting on the given sep.</span>
<span id="L81" class="ln">    81&nbsp;&nbsp;</span>func StringToSliceHookFunc(sep string) DecodeHookFunc {
<span id="L82" class="ln">    82&nbsp;&nbsp;</span>	return func(
<span id="L83" class="ln">    83&nbsp;&nbsp;</span>		f reflect.Kind,
<span id="L84" class="ln">    84&nbsp;&nbsp;</span>		t reflect.Kind,
<span id="L85" class="ln">    85&nbsp;&nbsp;</span>		data interface{}) (interface{}, error) {
<span id="L86" class="ln">    86&nbsp;&nbsp;</span>		if f != reflect.String || t != reflect.Slice {
<span id="L87" class="ln">    87&nbsp;&nbsp;</span>			return data, nil
<span id="L88" class="ln">    88&nbsp;&nbsp;</span>		}
<span id="L89" class="ln">    89&nbsp;&nbsp;</span>
<span id="L90" class="ln">    90&nbsp;&nbsp;</span>		raw := data.(string)
<span id="L91" class="ln">    91&nbsp;&nbsp;</span>		if raw == &#34;&#34; {
<span id="L92" class="ln">    92&nbsp;&nbsp;</span>			return []string{}, nil
<span id="L93" class="ln">    93&nbsp;&nbsp;</span>		}
<span id="L94" class="ln">    94&nbsp;&nbsp;</span>
<span id="L95" class="ln">    95&nbsp;&nbsp;</span>		return strings.Split(raw, sep), nil
<span id="L96" class="ln">    96&nbsp;&nbsp;</span>	}
<span id="L97" class="ln">    97&nbsp;&nbsp;</span>}
<span id="L98" class="ln">    98&nbsp;&nbsp;</span>
<span id="L99" class="ln">    99&nbsp;&nbsp;</span><span class="comment">// StringToTimeDurationHookFunc returns a DecodeHookFunc that converts</span>
<span id="L100" class="ln">   100&nbsp;&nbsp;</span><span class="comment">// strings to time.Duration.</span>
<span id="L101" class="ln">   101&nbsp;&nbsp;</span>func StringToTimeDurationHookFunc() DecodeHookFunc {
<span id="L102" class="ln">   102&nbsp;&nbsp;</span>	return func(
<span id="L103" class="ln">   103&nbsp;&nbsp;</span>		f reflect.Type,
<span id="L104" class="ln">   104&nbsp;&nbsp;</span>		t reflect.Type,
<span id="L105" class="ln">   105&nbsp;&nbsp;</span>		data interface{}) (interface{}, error) {
<span id="L106" class="ln">   106&nbsp;&nbsp;</span>		if f.Kind() != reflect.String {
<span id="L107" class="ln">   107&nbsp;&nbsp;</span>			return data, nil
<span id="L108" class="ln">   108&nbsp;&nbsp;</span>		}
<span id="L109" class="ln">   109&nbsp;&nbsp;</span>		if t != reflect.TypeOf(time.Duration(5)) {
<span id="L110" class="ln">   110&nbsp;&nbsp;</span>			return data, nil
<span id="L111" class="ln">   111&nbsp;&nbsp;</span>		}
<span id="L112" class="ln">   112&nbsp;&nbsp;</span>
<span id="L113" class="ln">   113&nbsp;&nbsp;</span>		<span class="comment">// Convert it by parsing</span>
<span id="L114" class="ln">   114&nbsp;&nbsp;</span>		return time.ParseDuration(data.(string))
<span id="L115" class="ln">   115&nbsp;&nbsp;</span>	}
<span id="L116" class="ln">   116&nbsp;&nbsp;</span>}
<span id="L117" class="ln">   117&nbsp;&nbsp;</span>
<span id="L118" class="ln">   118&nbsp;&nbsp;</span><span class="comment">// StringToIPHookFunc returns a DecodeHookFunc that converts</span>
<span id="L119" class="ln">   119&nbsp;&nbsp;</span><span class="comment">// strings to net.IP</span>
<span id="L120" class="ln">   120&nbsp;&nbsp;</span>func StringToIPHookFunc() DecodeHookFunc {
<span id="L121" class="ln">   121&nbsp;&nbsp;</span>	return func(
<span id="L122" class="ln">   122&nbsp;&nbsp;</span>		f reflect.Type,
<span id="L123" class="ln">   123&nbsp;&nbsp;</span>		t reflect.Type,
<span id="L124" class="ln">   124&nbsp;&nbsp;</span>		data interface{}) (interface{}, error) {
<span id="L125" class="ln">   125&nbsp;&nbsp;</span>		if f.Kind() != reflect.String {
<span id="L126" class="ln">   126&nbsp;&nbsp;</span>			return data, nil
<span id="L127" class="ln">   127&nbsp;&nbsp;</span>		}
<span id="L128" class="ln">   128&nbsp;&nbsp;</span>		if t != reflect.TypeOf(net.IP{}) {
<span id="L129" class="ln">   129&nbsp;&nbsp;</span>			return data, nil
<span id="L130" class="ln">   130&nbsp;&nbsp;</span>		}
<span id="L131" class="ln">   131&nbsp;&nbsp;</span>
<span id="L132" class="ln">   132&nbsp;&nbsp;</span>		<span class="comment">// Convert it by parsing</span>
<span id="L133" class="ln">   133&nbsp;&nbsp;</span>		ip := net.ParseIP(data.(string))
<span id="L134" class="ln">   134&nbsp;&nbsp;</span>		if ip == nil {
<span id="L135" class="ln">   135&nbsp;&nbsp;</span>			return net.IP{}, fmt.Errorf(&#34;failed parsing ip %v&#34;, data)
<span id="L136" class="ln">   136&nbsp;&nbsp;</span>		}
<span id="L137" class="ln">   137&nbsp;&nbsp;</span>
<span id="L138" class="ln">   138&nbsp;&nbsp;</span>		return ip, nil
<span id="L139" class="ln">   139&nbsp;&nbsp;</span>	}
<span id="L140" class="ln">   140&nbsp;&nbsp;</span>}
<span id="L141" class="ln">   141&nbsp;&nbsp;</span>
<span id="L142" class="ln">   142&nbsp;&nbsp;</span><span class="comment">// StringToIPNetHookFunc returns a DecodeHookFunc that converts</span>
<span id="L143" class="ln">   143&nbsp;&nbsp;</span><span class="comment">// strings to net.IPNet</span>
<span id="L144" class="ln">   144&nbsp;&nbsp;</span>func StringToIPNetHookFunc() DecodeHookFunc {
<span id="L145" class="ln">   145&nbsp;&nbsp;</span>	return func(
<span id="L146" class="ln">   146&nbsp;&nbsp;</span>		f reflect.Type,
<span id="L147" class="ln">   147&nbsp;&nbsp;</span>		t reflect.Type,
<span id="L148" class="ln">   148&nbsp;&nbsp;</span>		data interface{}) (interface{}, error) {
<span id="L149" class="ln">   149&nbsp;&nbsp;</span>		if f.Kind() != reflect.String {
<span id="L150" class="ln">   150&nbsp;&nbsp;</span>			return data, nil
<span id="L151" class="ln">   151&nbsp;&nbsp;</span>		}
<span id="L152" class="ln">   152&nbsp;&nbsp;</span>		if t != reflect.TypeOf(net.IPNet{}) {
<span id="L153" class="ln">   153&nbsp;&nbsp;</span>			return data, nil
<span id="L154" class="ln">   154&nbsp;&nbsp;</span>		}
<span id="L155" class="ln">   155&nbsp;&nbsp;</span>
<span id="L156" class="ln">   156&nbsp;&nbsp;</span>		<span class="comment">// Convert it by parsing</span>
<span id="L157" class="ln">   157&nbsp;&nbsp;</span>		_, net, err := net.ParseCIDR(data.(string))
<span id="L158" class="ln">   158&nbsp;&nbsp;</span>		return net, err
<span id="L159" class="ln">   159&nbsp;&nbsp;</span>	}
<span id="L160" class="ln">   160&nbsp;&nbsp;</span>}
<span id="L161" class="ln">   161&nbsp;&nbsp;</span>
<span id="L162" class="ln">   162&nbsp;&nbsp;</span><span class="comment">// StringToTimeHookFunc returns a DecodeHookFunc that converts</span>
<span id="L163" class="ln">   163&nbsp;&nbsp;</span><span class="comment">// strings to time.Time.</span>
<span id="L164" class="ln">   164&nbsp;&nbsp;</span>func StringToTimeHookFunc(layout string) DecodeHookFunc {
<span id="L165" class="ln">   165&nbsp;&nbsp;</span>	return func(
<span id="L166" class="ln">   166&nbsp;&nbsp;</span>		f reflect.Type,
<span id="L167" class="ln">   167&nbsp;&nbsp;</span>		t reflect.Type,
<span id="L168" class="ln">   168&nbsp;&nbsp;</span>		data interface{}) (interface{}, error) {
<span id="L169" class="ln">   169&nbsp;&nbsp;</span>		if f.Kind() != reflect.String {
<span id="L170" class="ln">   170&nbsp;&nbsp;</span>			return data, nil
<span id="L171" class="ln">   171&nbsp;&nbsp;</span>		}
<span id="L172" class="ln">   172&nbsp;&nbsp;</span>		if t != reflect.TypeOf(time.Time{}) {
<span id="L173" class="ln">   173&nbsp;&nbsp;</span>			return data, nil
<span id="L174" class="ln">   174&nbsp;&nbsp;</span>		}
<span id="L175" class="ln">   175&nbsp;&nbsp;</span>
<span id="L176" class="ln">   176&nbsp;&nbsp;</span>		<span class="comment">// Convert it by parsing</span>
<span id="L177" class="ln">   177&nbsp;&nbsp;</span>		return time.Parse(layout, data.(string))
<span id="L178" class="ln">   178&nbsp;&nbsp;</span>	}
<span id="L179" class="ln">   179&nbsp;&nbsp;</span>}
<span id="L180" class="ln">   180&nbsp;&nbsp;</span>
<span id="L181" class="ln">   181&nbsp;&nbsp;</span><span class="comment">// WeaklyTypedHook is a DecodeHookFunc which adds support for weak typing to</span>
<span id="L182" class="ln">   182&nbsp;&nbsp;</span><span class="comment">// the decoder.</span>
<span id="L183" class="ln">   183&nbsp;&nbsp;</span><span class="comment">//</span>
<span id="L184" class="ln">   184&nbsp;&nbsp;</span><span class="comment">// Note that this is significantly different from the WeaklyTypedInput option</span>
<span id="L185" class="ln">   185&nbsp;&nbsp;</span><span class="comment">// of the DecoderConfig.</span>
<span id="L186" class="ln">   186&nbsp;&nbsp;</span>func WeaklyTypedHook(
<span id="L187" class="ln">   187&nbsp;&nbsp;</span>	f reflect.Kind,
<span id="L188" class="ln">   188&nbsp;&nbsp;</span>	t reflect.Kind,
<span id="L189" class="ln">   189&nbsp;&nbsp;</span>	data interface{}) (interface{}, error) {
<span id="L190" class="ln">   190&nbsp;&nbsp;</span>	dataVal := reflect.ValueOf(data)
<span id="L191" class="ln">   191&nbsp;&nbsp;</span>	switch t {
<span id="L192" class="ln">   192&nbsp;&nbsp;</span>	case reflect.String:
<span id="L193" class="ln">   193&nbsp;&nbsp;</span>		switch f {
<span id="L194" class="ln">   194&nbsp;&nbsp;</span>		case reflect.Bool:
<span id="L195" class="ln">   195&nbsp;&nbsp;</span>			if dataVal.Bool() {
<span id="L196" class="ln">   196&nbsp;&nbsp;</span>				return &#34;1&#34;, nil
<span id="L197" class="ln">   197&nbsp;&nbsp;</span>			}
<span id="L198" class="ln">   198&nbsp;&nbsp;</span>			return &#34;0&#34;, nil
<span id="L199" class="ln">   199&nbsp;&nbsp;</span>		case reflect.Float32:
<span id="L200" class="ln">   200&nbsp;&nbsp;</span>			return strconv.FormatFloat(dataVal.Float(), &#39;f&#39;, -1, 64), nil
<span id="L201" class="ln">   201&nbsp;&nbsp;</span>		case reflect.Int:
<span id="L202" class="ln">   202&nbsp;&nbsp;</span>			return strconv.FormatInt(dataVal.Int(), 10), nil
<span id="L203" class="ln">   203&nbsp;&nbsp;</span>		case reflect.Slice:
<span id="L204" class="ln">   204&nbsp;&nbsp;</span>			dataType := dataVal.Type()
<span id="L205" class="ln">   205&nbsp;&nbsp;</span>			elemKind := dataType.Elem().Kind()
<span id="L206" class="ln">   206&nbsp;&nbsp;</span>			if elemKind == reflect.Uint8 {
<span id="L207" class="ln">   207&nbsp;&nbsp;</span>				return string(dataVal.Interface().([]uint8)), nil
<span id="L208" class="ln">   208&nbsp;&nbsp;</span>			}
<span id="L209" class="ln">   209&nbsp;&nbsp;</span>		case reflect.Uint:
<span id="L210" class="ln">   210&nbsp;&nbsp;</span>			return strconv.FormatUint(dataVal.Uint(), 10), nil
<span id="L211" class="ln">   211&nbsp;&nbsp;</span>		}
<span id="L212" class="ln">   212&nbsp;&nbsp;</span>	}
<span id="L213" class="ln">   213&nbsp;&nbsp;</span>
<span id="L214" class="ln">   214&nbsp;&nbsp;</span>	return data, nil
<span id="L215" class="ln">   215&nbsp;&nbsp;</span>}
<span id="L216" class="ln">   216&nbsp;&nbsp;</span>
<span id="L217" class="ln">   217&nbsp;&nbsp;</span>func RecursiveStructToMapHookFunc() DecodeHookFunc {
<span id="L218" class="ln">   218&nbsp;&nbsp;</span>	return func(f reflect.Value, t reflect.Value) (interface{}, error) {
<span id="L219" class="ln">   219&nbsp;&nbsp;</span>		if f.Kind() != reflect.Struct {
<span id="L220" class="ln">   220&nbsp;&nbsp;</span>			return f.Interface(), nil
<span id="L221" class="ln">   221&nbsp;&nbsp;</span>		}
<span id="L222" class="ln">   222&nbsp;&nbsp;</span>
<span id="L223" class="ln">   223&nbsp;&nbsp;</span>		var i interface{} = struct{}{}
<span id="L224" class="ln">   224&nbsp;&nbsp;</span>		if t.Type() != reflect.TypeOf(&amp;i).Elem() {
<span id="L225" class="ln">   225&nbsp;&nbsp;</span>			return f.Interface(), nil
<span id="L226" class="ln">   226&nbsp;&nbsp;</span>		}
<span id="L227" class="ln">   227&nbsp;&nbsp;</span>
<span id="L228" class="ln">   228&nbsp;&nbsp;</span>		m := make(map[string]interface{})
<span id="L229" class="ln">   229&nbsp;&nbsp;</span>		t.Set(reflect.ValueOf(m))
<span id="L230" class="ln">   230&nbsp;&nbsp;</span>
<span id="L231" class="ln">   231&nbsp;&nbsp;</span>		return f.Interface(), nil
<span id="L232" class="ln">   232&nbsp;&nbsp;</span>	}
<span id="L233" class="ln">   233&nbsp;&nbsp;</span>}
<span id="L234" class="ln">   234&nbsp;&nbsp;</span>
<span id="L235" class="ln">   235&nbsp;&nbsp;</span><span class="comment">// TextUnmarshallerHookFunc returns a DecodeHookFunc that applies</span>
<span id="L236" class="ln">   236&nbsp;&nbsp;</span><span class="comment">// strings to the UnmarshalText function, when the target type</span>
<span id="L237" class="ln">   237&nbsp;&nbsp;</span><span class="comment">// implements the encoding.TextUnmarshaler interface</span>
<span id="L238" class="ln">   238&nbsp;&nbsp;</span>func TextUnmarshallerHookFunc() DecodeHookFuncType {
<span id="L239" class="ln">   239&nbsp;&nbsp;</span>	return func(
<span id="L240" class="ln">   240&nbsp;&nbsp;</span>		f reflect.Type,
<span id="L241" class="ln">   241&nbsp;&nbsp;</span>		t reflect.Type,
<span id="L242" class="ln">   242&nbsp;&nbsp;</span>		data interface{}) (interface{}, error) {
<span id="L243" class="ln">   243&nbsp;&nbsp;</span>		if f.Kind() != reflect.String {
<span id="L244" class="ln">   244&nbsp;&nbsp;</span>			return data, nil
<span id="L245" class="ln">   245&nbsp;&nbsp;</span>		}
<span id="L246" class="ln">   246&nbsp;&nbsp;</span>		result := reflect.New(t).Interface()
<span id="L247" class="ln">   247&nbsp;&nbsp;</span>		unmarshaller, ok := result.(encoding.TextUnmarshaler)
<span id="L248" class="ln">   248&nbsp;&nbsp;</span>		if !ok {
<span id="L249" class="ln">   249&nbsp;&nbsp;</span>			return data, nil
<span id="L250" class="ln">   250&nbsp;&nbsp;</span>		}
<span id="L251" class="ln">   251&nbsp;&nbsp;</span>		if err := unmarshaller.UnmarshalText([]byte(data.(string))); err != nil {
<span id="L252" class="ln">   252&nbsp;&nbsp;</span>			return nil, err
<span id="L253" class="ln">   253&nbsp;&nbsp;</span>		}
<span id="L254" class="ln">   254&nbsp;&nbsp;</span>		return result, nil
<span id="L255" class="ln">   255&nbsp;&nbsp;</span>	}
<span id="L256" class="ln">   256&nbsp;&nbsp;</span>}
<span id="L257" class="ln">   257&nbsp;&nbsp;</span>
</pre><p><a href="/src/github.com/mitchellh/mapstructure/decode_hooks.go?m=text">View as plain text</a></p>

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
