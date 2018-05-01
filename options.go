package wkhtmltopdfgo

//Type for page's orientation
type Orientation string

// Orientations
const (
	Landscape Orientation = "Landscape" // Landscape mode
	Portrait  Orientation = "Portrait"  // Portrait mode
)

// Type for page sizes
type PageSize string

// http://doc.qt.io/archives/qt-4.8/qprinter.html#PaperSize-enum (QPrinter::(.+?)\t(.+?)\t(.+?)\n => $1 PageSize = "$1" //$3\n)
const (
	A0        PageSize = "A0"        //841 x 1189 mm
	A1        PageSize = "A1"        //594 x 841 mm
	A2        PageSize = "A2"        //420 x 594 mm
	A3        PageSize = "A3"        //297 x 420 mm
	A4        PageSize = "A4"        //210 x 297 mm, 8.26 x 11.69 inches
	A5        PageSize = "A5"        //148 x 210 mm
	A6        PageSize = "A6"        //105 x 148 mm
	A7        PageSize = "A7"        //74 x 105 mm
	A8        PageSize = "A8"        //52 x 74 mm
	A9        PageSize = "A9"        //37 x 52 mm
	B0        PageSize = "B0"        //1000 x 1414 mm
	B1        PageSize = "B1"        //707 x 1000 mm
	B2        PageSize = "B2"        //500 x 707 mm
	B3        PageSize = "B3"        //353 x 500 mm
	B4        PageSize = "B4"        //250 x 353 mm
	B5        PageSize = "B5"        //176 x 250 mm, 6.93 x 9.84 inches
	B6        PageSize = "B6"        //125 x 176 mm
	B7        PageSize = "B7"        //88 x 125 mm
	B8        PageSize = "B8"        //62 x 88 mm
	B9        PageSize = "B9"        //33 x 62 mm
	B10       PageSize = "B10"       //31 x 44 mm
	C5E       PageSize = "C5E"       //163 x 229 mm
	Comm10E   PageSize = "Comm10E"   //105 x 241 mm, U.S. Common 10 Envelope
	DLE       PageSize = "DLE"       //110 x 220 mm
	Executive PageSize = "Executive" //7.5 x 10 inches, 190.5 x 254 mm
	Folio     PageSize = "Folio"     //210 x 330 mm
	Ledger    PageSize = "Ledger"    //431.8 x 279.4 mm
	Legal     PageSize = "Legal"     //8.5 x 14 inches, 215.9 x 355.6 mm
	Letter    PageSize = "Letter"    //8.5 x 11 inches, 215.9 x 279.4 mm
	Tabloid   PageSize = "Tabloid"   //279.4 x 431.8 mm
	Custom    PageSize = "Custom"    //Unknown, or a user defined size.
)

//Options for command-line arguments.
//See more: https://wkhtmltopdf.org/usage/wkhtmltopdf.txt
type WkhtmlOptions struct {
	//Global Options
	Collate           WkhtmlParam `json:"collate"`
	NoCollate         WkhtmlParam `json:"no_collate"`
	CookieJar         WkhtmlParam `json:"cookie_jar"`
	Copies            WkhtmlParam `json:"copies"`
	Dpi               WkhtmlParam `json:"dpi"`
	Grayscale         WkhtmlParam `json:"grayscale"`
	Htmldoc           WkhtmlParam `json:"htmldoc"`
	ImageDpi          WkhtmlParam `json:"image_dpi"`
	ImageQuality      WkhtmlParam `json:"image_quality"`
	Lowquality        WkhtmlParam `json:"lowquality"`
	MarginBottom      WkhtmlParam `json:"margin_bottom"`
	MarginLeft        WkhtmlParam `json:"margin_left"`
	MarginRight       WkhtmlParam `json:"margin_right"`
	MarginTop         WkhtmlParam `json:"margin_top"`
	Orientation       WkhtmlParam `json:"orientation"`
	PageHeight        WkhtmlParam `json:"page_height"`
	PageSize          WkhtmlParam `json:"page_size"`
	PageWidth         WkhtmlParam `json:"page_width"`
	NoPdfCompression  WkhtmlParam `json:"no_pdf_compression"`
	Quiet             WkhtmlParam `json:"quiet"`
	ReadArgsFromStdin WkhtmlParam `json:"read_args_from_stdin"`
	Title             WkhtmlParam `json:"title"`
	UseXServer        WkhtmlParam `json:"use_xserver"`
	//Outline Options
	DumpDefaultTocXsl WkhtmlParam `json:"dump_default_toc_xsl"`
	DumpOutline       WkhtmlParam `json:"dump_outline"`
	Outline           WkhtmlParam `json:"outline"`
	NoOutline         WkhtmlParam `json:"no_outline"`
	OutlineDepth      WkhtmlParam `json:"outline_depth"`
	//Page Options
	Allow                     WkhtmlParam `json:"allow"`
	Background                WkhtmlParam `json:"background"`
	NoBackground              WkhtmlParam `json:"no_background"`
	BypassProxyFor            WkhtmlParam `json:"bypass_proxy_for"`
	CacheDir                  WkhtmlParam `json:"cache_dir"`
	CheckboxCheckedCvg        WkhtmlParam `json:"checkbox_checked_cvg"`
	CheckboxSvg               WkhtmlParam `json:"checkbox_svg"`
	Cookie                    WkhtmlParam `json:"cookie"`
	CustomHeader              WkhtmlParam `json:"custom_header"`
	CustomHeaderPropagation   WkhtmlParam `json:"custom_header_propagation"`
	NoCustomHeaderPropagation WkhtmlParam `json:"no_custom_header_propagation"`
	DebugJavascript           WkhtmlParam `json:"debug_javascript"`
	NoDebugJavascript         WkhtmlParam `json:"no_debug_javascript"`
	DefaultHeader             WkhtmlParam `json:"default_header"`
	Encoding                  WkhtmlParam `json:"encoding"`
	DisableExternalLinks      WkhtmlParam `json:"disable_external_links"`
	EnableExternalLinks       WkhtmlParam `json:"enable_external_links"`
	DisableForms              WkhtmlParam `json:"disable_forms"`
	EnableForms               WkhtmlParam `json:"enable_forms"`
	Images                    WkhtmlParam `json:"images"`
	NoImages                  WkhtmlParam `json:"no_images"`
	DisableInternalLinks      WkhtmlParam `json:"disable_internal_links"`
	EnableInternalLinks       WkhtmlParam `json:"enable_internal_links"`
	DisableJavascript         WkhtmlParam `json:"disable_javascript"`
	EnableJavascript          WkhtmlParam `json:"enable_javascript"`
	JavascriptDelay           WkhtmlParam `json:"javascript_delay"`
	KeepRelativeLinks         WkhtmlParam `json:"keep_relative_links"`
	LoadErrorHandling         WkhtmlParam `json:"load_error_handling"`
	LoadMediaErrorHandling    WkhtmlParam `json:"load_media_error_handling"`
	DisableLocalFileAccess    WkhtmlParam `json:"disable_local_file_access"`
	EnableLocalFileAccess     WkhtmlParam `json:"enable_local_file_access"`
	MinimumFontSize           WkhtmlParam `json:"minimum_font_size"`
	ExcludeFromOutline        WkhtmlParam `json:"exclude_from_outline"`
	IncludeInOutline          WkhtmlParam `json:"include_in_outline"`
	PageOffset                WkhtmlParam `json:"page_offset"`
	Password                  WkhtmlParam `json:"password"`
	DisablePlugins            WkhtmlParam `json:"disable_plugins"`
	EnablePlugins             WkhtmlParam `json:"enable_plugins"`
	Post                      WkhtmlParam `json:"post"`
	PostFile                  WkhtmlParam `json:"post_file"`
	PrintMediaType            WkhtmlParam `json:"print_media_type"`
	NoPrintMediaType          WkhtmlParam `json:"no_print_media_type"`
	Proxy                     WkhtmlParam `json:"proxy"`
	RadiobuttonCheckedSvg     WkhtmlParam `json:"radiobutton_checked_svg"`
	RadiobuttonSvg            WkhtmlParam `json:"radiobutton_svg"`
	ResolveRelativeLinks      WkhtmlParam `json:"resolve_relative_links"`
	RunScript                 WkhtmlParam `json:"run_script"`
	DisableSmartShrinking     WkhtmlParam `json:"disable_smart_shrinking"`
	EnableSmartShrinking      WkhtmlParam `json:"enable_smart_shrinking"`
	StopSlowScripts           WkhtmlParam `json:"stop_slow_scripts"`
	NoStopSlowScripts         WkhtmlParam `json:"no_stop_slow_scripts"`
	DisableTocBackLinks       WkhtmlParam `json:"disable_toc_back_links"`
	EnableTocBackLinks        WkhtmlParam `json:"enable_toc_back_links"`
	UserStyleSheet            WkhtmlParam `json:"user_style_sheet"`
	Username                  WkhtmlParam `json:"username"`
	ViewportSize              WkhtmlParam `json:"viewport_size"`
	WindowStatus              WkhtmlParam `json:"window_status"`
	Zoom                      WkhtmlParam `json:"zoom"`
	//Headers And Footer Options
	FoterCenter    WkhtmlParam `json:"foter_center"`
	FooterFontName WkhtmlParam `json:"footer_font_name"`
	FooterFontSize WkhtmlParam `json:"footer_font_size"`
	FooterHtml     WkhtmlParam `json:"footer_html"`
	FooterLeft     WkhtmlParam `json:"footer_left"`
	FooterLine     WkhtmlParam `json:"footer_line"`
	NoFooterLine   WkhtmlParam `json:"no_footer_line"`
	FooterRight    WkhtmlParam `json:"footer_right"`
	FooterSpacing  WkhtmlParam `json:"footer_spacing"`
	HeaderCenter   WkhtmlParam `json:"header_center"`
	HeaderFontName WkhtmlParam `json:"header_font_name"`
	HeaderFontSize WkhtmlParam `json:"header_font_size"`
	HeaderHtml     WkhtmlParam `json:"header_html"`
	HeaderLeft     WkhtmlParam `json:"header_left"`
	HeaderLine     WkhtmlParam `json:"header_line"`
	NoHeaderLine   WkhtmlParam `json:"no_header_line"`
	HeaderRight    WkhtmlParam `json:"header_right"`
	HeaderSpacing  WkhtmlParam `json:"header_spacing"`
	Replace        WkhtmlParam `json:"replace"`
	//TOC Options
	DisableDottedLines  WkhtmlParam `json:"disable_dotted_lines"`
	TocHeaderText       WkhtmlParam `json:"toc_header_text"`
	TocLevelIndentation WkhtmlParam `json:"toc_level_indentation"`
	DisableTocLinks     WkhtmlParam `json:"disable_toc_links"`
	TocTextSizeShrink   WkhtmlParam `json:"toc_text_size_shrink"`
	XslStyleSheet       WkhtmlParam `json:"xsl_style_sheet"`
}

type WkhtmlParam struct {
	name  string
	value interface{}
}

func (w *WkhtmlParam) Set(v interface{}) {
	w.value = v
}

//Set default names and values for options.
//See more: https://wkhtmltopdf.org/usage/wkhtmltopdf.txt
func defaultOptions() WkhtmlOptions {
	return WkhtmlOptions{
		Collate:           WkhtmlParam{name: "--collate"},
		NoCollate:         WkhtmlParam{name: "--no-collate"},
		CookieJar:         WkhtmlParam{name: "--cookie-jar"},
		Copies:            WkhtmlParam{name: "--copies"},
		Dpi:               WkhtmlParam{name: "--dpi"},
		Grayscale:         WkhtmlParam{name: "--grayscale"},
		Htmldoc:           WkhtmlParam{name: "--htmldoc"},
		ImageDpi:          WkhtmlParam{name: "--image-dpi"},
		ImageQuality:      WkhtmlParam{name: "--image-quality"},
		Lowquality:        WkhtmlParam{name: "--lowquality"},
		MarginBottom:      WkhtmlParam{name: "--margin-bottom"},
		MarginLeft:        WkhtmlParam{name: "--margin-left"},
		MarginRight:       WkhtmlParam{name: "--margin-right"},
		MarginTop:         WkhtmlParam{name: "--margin-top"},
		Orientation:       WkhtmlParam{name: "--orientation"},
		PageHeight:        WkhtmlParam{name: "--page-height"},
		PageSize:          WkhtmlParam{name: "--page-size"},
		PageWidth:         WkhtmlParam{name: "--page-width"},
		NoPdfCompression:  WkhtmlParam{name: "--no-pdf-compression"},
		Quiet:             WkhtmlParam{name: "--quiet"},
		ReadArgsFromStdin: WkhtmlParam{name: "--read-args-from-stdin"},
		Title:             WkhtmlParam{name: "--title"},
		UseXServer:        WkhtmlParam{name: "--use-xserver"},
		//Outline Options
		DumpDefaultTocXsl: WkhtmlParam{name: "--dump-default-toc-xsl"},
		DumpOutline:       WkhtmlParam{name: "--dump-outline"},
		Outline:           WkhtmlParam{name: "--outline"},
		NoOutline:         WkhtmlParam{name: "--no-outline"},
		OutlineDepth:      WkhtmlParam{name: "--outline-depth"},
		//Page Options
		Allow:                     WkhtmlParam{name: "--allow"},
		Background:                WkhtmlParam{name: "--background"},
		NoBackground:              WkhtmlParam{name: "--no-background"},
		BypassProxyFor:            WkhtmlParam{name: "--bypass-proxy-for"},
		CacheDir:                  WkhtmlParam{name: "--cache-dir"},
		CheckboxCheckedCvg:        WkhtmlParam{name: "--checkbox-checked-svg"},
		CheckboxSvg:               WkhtmlParam{name: "--checkbox-svg"},
		Cookie:                    WkhtmlParam{name: "--cookie"},
		CustomHeader:              WkhtmlParam{name: "--custom-header"}, //string `<name> <value>`
		CustomHeaderPropagation:   WkhtmlParam{name: "--custom-header-propagation"},
		NoCustomHeaderPropagation: WkhtmlParam{name: "--no-custom-header-propagation"},
		DebugJavascript:           WkhtmlParam{name: "--debug-javascript"},
		NoDebugJavascript:         WkhtmlParam{name: "--no-debug-javascript"},
		DefaultHeader:             WkhtmlParam{name: "--default-header"},
		Encoding:                  WkhtmlParam{name: "--encoding"},
		DisableExternalLinks:      WkhtmlParam{name: "--disable-external-links"},
		EnableExternalLinks:       WkhtmlParam{name: "--enable-external-links"},
		DisableForms:              WkhtmlParam{name: "--disable-forms"},
		EnableForms:               WkhtmlParam{name: "--enable-forms"},
		Images:                    WkhtmlParam{name: "--images"},
		NoImages:                  WkhtmlParam{name: "--no-images"},
		DisableInternalLinks:      WkhtmlParam{name: "--disable-internal-links"},
		EnableInternalLinks:       WkhtmlParam{name: "--enable-internal-links"},
		DisableJavascript:         WkhtmlParam{name: "--disable-javascript"},
		EnableJavascript:          WkhtmlParam{name: "--enable-javascript"},
		JavascriptDelay:           WkhtmlParam{name: "--javascript-delay"},
		KeepRelativeLinks:         WkhtmlParam{name: "--keep-relative-links"},
		LoadErrorHandling:         WkhtmlParam{name: "--load-error-handling"},
		LoadMediaErrorHandling:    WkhtmlParam{name: "--load-media-error-handling"},
		DisableLocalFileAccess:    WkhtmlParam{name: "--disable-local-file-access"},
		EnableLocalFileAccess:     WkhtmlParam{name: "--enable-local-file-access"},
		MinimumFontSize:           WkhtmlParam{name: "--minimum-font-size"},
		ExcludeFromOutline:        WkhtmlParam{name: "--exclude-from-outline"},
		IncludeInOutline:          WkhtmlParam{name: "--include-in-outline"},
		PageOffset:                WkhtmlParam{name: "--page-offset"},
		Password:                  WkhtmlParam{name: "--password"},
		DisablePlugins:            WkhtmlParam{name: "--disable-plugins"},
		EnablePlugins:             WkhtmlParam{name: "--enable-plugins"},
		Post:                      WkhtmlParam{name: "--post"},
		PostFile:                  WkhtmlParam{name: "--post-file"},
		PrintMediaType:            WkhtmlParam{name: "--print-media-type"},
		NoPrintMediaType:          WkhtmlParam{name: "--no-print-media-type"},
		Proxy:                     WkhtmlParam{name: "--proxy"},
		RadiobuttonCheckedSvg:     WkhtmlParam{name: "--radiobutton-checked-svg"},
		RadiobuttonSvg:            WkhtmlParam{name: "--radiobutton-svg"},
		ResolveRelativeLinks:      WkhtmlParam{name: "--resolve-relative-links"},
		RunScript:                 WkhtmlParam{name: "--run-script"},
		DisableSmartShrinking:     WkhtmlParam{name: "--disable-smart-shrinking"},
		EnableSmartShrinking:      WkhtmlParam{name: "--enable-smart-shrinking"},
		StopSlowScripts:           WkhtmlParam{name: "--stop-slow-scripts"},
		NoStopSlowScripts:         WkhtmlParam{name: "--no-stop-slow-scripts"},
		DisableTocBackLinks:       WkhtmlParam{name: "--disable-toc-back-links"},
		EnableTocBackLinks:        WkhtmlParam{name: "--enable-toc-back-links"},
		UserStyleSheet:            WkhtmlParam{name: "--user-style-sheet"},
		Username:                  WkhtmlParam{name: "--username"},
		ViewportSize:              WkhtmlParam{name: "--viewport-size"},
		WindowStatus:              WkhtmlParam{name: "--window-status"},
		Zoom:                      WkhtmlParam{name: "--zoom"},
		//Headers And Footer Options
		FoterCenter:    WkhtmlParam{name: "--footer-center"},
		FooterFontName: WkhtmlParam{name: "--footer-font-name"},
		FooterFontSize: WkhtmlParam{name: "--footer-font-size"},
		FooterHtml:     WkhtmlParam{name: "--footer-html"},
		FooterLeft:     WkhtmlParam{name: "--footer-left"},
		FooterLine:     WkhtmlParam{name: "--footer-line"},
		NoFooterLine:   WkhtmlParam{name: "--no-footer-line"},
		FooterRight:    WkhtmlParam{name: "--footer-right"},
		FooterSpacing:  WkhtmlParam{name: "--footer-spacing"},
		HeaderCenter:   WkhtmlParam{name: "--header-center"},
		HeaderFontName: WkhtmlParam{name: "--header-font-name"},
		HeaderFontSize: WkhtmlParam{name: "--header-font-size"},
		HeaderHtml:     WkhtmlParam{name: "--header-html"},
		HeaderLeft:     WkhtmlParam{name: "--header-left"},
		HeaderLine:     WkhtmlParam{name: "--header-line"},
		NoHeaderLine:   WkhtmlParam{name: "--no-header-line"},
		HeaderRight:    WkhtmlParam{name: "--header-right"},
		HeaderSpacing:  WkhtmlParam{name: "--header-spacing"},
		Replace:        WkhtmlParam{name: "--replace"},
		//TOC Options
		DisableDottedLines:  WkhtmlParam{name: "--disable-dotted-lines"},
		TocHeaderText:       WkhtmlParam{name: "--toc-header-text"},
		TocLevelIndentation: WkhtmlParam{name: "--toc-level-indentation"},
		DisableTocLinks:     WkhtmlParam{name: "--disable-toc-links"},
		TocTextSizeShrink:   WkhtmlParam{name: "--toc-text-size-shrink"},
		XslStyleSheet:       WkhtmlParam{name: "--xsl-style-sheet"},
	}
}
