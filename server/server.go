package server

import (
	"net"
)

//TODO: Change to use struct to use Error message
// type WhoisServer struct {
//     URL string
//     ErrorMessage string
// }
// ABOGADO = WhoisServer {
//     URL: "whois.nic.abogado"
//     ErrorMessage: "This domain name has not been registered."
// }

const (
	ABOGADO                = "whois.nic.abogado"
	AC                     = "whois.nic.ac"
	ACCOUNT                = "whois.nic.accountant"
	AE                     = "whois.aeda.net.ae"
	AERO                   = "whois.aero"
	AFILIAS                = "whois.afilias.net"
	AFILIAS_GRS            = "whois.afilias-grs.info"
	AFILIAS_GRS2           = "whois2.afilias-grs.net"
	AG                     = "whois.nic.ag"
	AI                     = "whois.nic.ai"
	ALSACE                 = "whois.nic.alsace"
	AMNIC                  = "whois.amnic.net"
	AMSTERDAM              = "whois.nic.amsterdam"
	AQUARELLE              = "whois-aquarelle.nic.fr"
	AS                     = "whois.nic.as"
	ASIA                   = "whois.nic.asia"
	AT                     = "whois.nic.at"
	AUNIC                  = "whois.aunic.net"
	FR                     = "whois.nic.fr"
	AU                     = "whois.audns.net.au"
	AW                     = "whois.nic.aw"
	AX                     = "whois.ax"
	BANK                   = "whois.nic.bank"
	BAR                    = "whois.nic.bar"
	BARCLAYCARD            = "whois.nic.barclaycard"
	BARCLAYS               = "whois.nic.barclays"
	BAYERN                 = "whois.nic.bayern"
	BE                     = "whois.dns.be"
	BEER                   = "whois.nic.beer"
	BERLIN                 = "whois.nic.berlin"
	BI                     = "whois1.nic.bi"
	BID                    = "whois.nic.bid"
	BIZ                    = "whois.biz"
	BJ                     = "whois.nic.bj"
	BLOG                   = "whois.nic.blog"
	BRUSSELS               = "whois.nic.brussels"
	BUDAPEST               = "whois.nic.budapest"
	BUILD                  = "whois.nic.build"
	BUZZ                   = "whois.nic.buzz"
	BW                     = "whois.nic.net.bw"
	BY                     = "whois.cctld.by"
	BZH                    = "whois-bzh.nic.fr"
	CA                     = "whois.cira.ca"
	CANCERRESEARCH         = "whois.nic.cancerresearch"
	CAPETOWN               = "capetown-whois.registry.net.za"
	CAREER                 = "whois.nic.career"
	CASA                   = "whois.nic.casa"
	CAT                    = "whois.cat"
	CC                     = "ccwhois.verisign-grs.com"
	CH                     = "whois.nic.ch"
	CI                     = "whois.nic.ci"
	CL                     = "whois.nic.ci"
	CLOUD                  = "whois.nic.cloud"
	CLUB                   = "whois.nic.club"
	CM                     = "whois.netcom.cm"
	CN                     = "whois.cnnic.net.cn"
	CO                     = "whois.nic.co"
	COLOGNE                = "whois.nic.cologne"
	COOKING                = "whois.nic.cooking"
	COOP                   = "whois.nic.coop"
	CR                     = "whois.nic.cr"
	CRICKET                = "whois.nic.cricket"
	CUISINELLA             = "whois.nic.cuisinella"
	CX                     = "whois.nic.cx"
	CYMRU                  = "whois.nic.cymru"
	CZ                     = "whois.nic.cz"
	DATE                   = "whois.nic.date"
	DE                     = "whois.denic.de"
	DK                     = "whois.dk-hostmaster.dk"
	DM                     = "whois.nic.dm"
	DO                     = "whois.nic.do"
	DOWNLOAD               = "whois.nic.download"
	DURBAN                 = "durban-whois.registry.net.za"
	DZ                     = "whois.nic.dz"
	EE                     = "whois.tld.ee"
	EU                     = "whois.eu"
	LV                     = "whois.biz"
	EU_ORG                 = "whois.eu.org"
	EUROVISION             = "whois.nic.eurovision"
	EUS                    = "whois.nic.eus"
	FAITH                  = "whois.nic.faith"
	FASHION                = "whois.nic.fashion"
	FI                     = "whois.fi"
	FILM                   = "whois.nic.film"
	FIRMDALE               = "whois.nic.firmdale"
	FISHING                = "whois.nic.fishing"
	FIT                    = "whois.nic.fit"
	FRL                    = "whois.nic.frl"
	FROGANS                = "whois.nic.frogans"
	GA                     = "whois.dot.ga"
	GAL                    = "whois.nic.gal"
	GAMES                  = "whois.nic.games"
	GARDEN                 = "whois.nic.garden"
	GD                     = "whois.nic.gd"
	GDN                    = "whois.nic.gdn"
	GENT                   = "whois.nic.gent"
	GG                     = "whois.gg"
	GL                     = "whois.nic.gl"
	GLOBAL                 = "whois.nic.global"
	GMX                    = "whois.nic.gmx"
	GOLD                   = "whois.nic.gold"
	GOP                    = "whois.nic.gop"
	GOV                    = "whois.nic.gov"
	GQ                     = "whois.dominio.gq"
	GY                     = "whois.registry.gy"
	HAMBURG                = "whois.nic.hamburg"
	HK                     = "whois.hkirc.hk"
	HN                     = "whois.nic.hn"
	HOURSE                 = "whois.nic.horse"
	HR                     = "whois.dns.hr"
	HT                     = "whois.nic.ht"
	HU                     = "whois.nic.hu"
	IBM                    = "whois.nic.ibm"
	ID                     = "whois.pandi.or.id"
	IE                     = "whois.domainregistry.ie"
	IFM                    = "whois.nic.ifm"
	IL                     = "whois.isoc.org.il"
	IM                     = "whois.nic.im"
	IN                     = "whois.registry.in"
	INT                    = "whois.iana.org"
	IO                     = "whois.nic.io"
	IQ                     = "whois.cmc.iq"
	IR                     = "whois.nic.ir"
	IS                     = "whois.isnic.is"
	IT                     = "whois.nic.it"
	JAVA                   = "whois.nic.java"
	JE                     = "whois.je"
	JETZT                  = "whois.nic.jetzt"
	JP                     = "whois.jprs.jp"
	JOBS                   = "whois.nic.jobs"
	JOBURG                 = "joburg-whois.registry.net.za"
	KE                     = "whois.kenic.or.ke"
	KG                     = "whois.kg"
	KI                     = "whois.nic.ki"
	KIWI                   = "whois.nic.kiwi"
	KOELN                  = "whois.nic.koeln"
	KR                     = "whois.kr"
	KR_OR                  = "whois.nic.or.kr"
	KY                     = "whois.kyregistry.ky"
	KZ                     = "whois.nic.kz"
	LA                     = "whois.nic.la"
	LACAIXA                = "whois.nic.lacaixa"
	LAT                    = "whois.nic.lat"
	LATROBE                = "whois.nic.latrobe"
	LECLERC                = "whois-leclerc.nic.fr"
	LI                     = "whois.nic.li"
	LIVE                   = "whois.nic.live"
	LOAN                   = "whois.nic.loan"
	LONDON                 = "whois.nic.london"
	LT                     = "whois.domreg.lt"
	LU                     = "whois.dns.lu"
	LUXE                   = "whois.nic.luxe"
	LUXURY                 = "whois.nic.luxury"
	MA                     = "whois.iam.net.ma"
	MADRID                 = "whois.madrid.rs.corenic.net"
	MANGO                  = "whois.nic.mango"
	MD                     = "whois.nic.md"
	ME                     = "whois.nic.me"
	MEN                    = "whois.nic.men"
	MENU                   = "whois.nic.menu"
	MG                     = "whois.nic.mg"
	MIAMI                  = "whois.nic.miami"
	MK                     = "whois.marnet.mk"
	ML                     = "whois.dot.ml"
	MO                     = "whois.monic.mo"
	MOE                    = "whois.nic.moe"
	MONASH                 = "whois.nic.monash"
	MOSCOW                 = "whois.nic.moscow"
	MS                     = "whois.nic.ms"
	MU                     = "whois.nic.mu"
	MUSEUM                 = "whois.museum"
	MX                     = "whois.mx"
	MZ                     = "whois.nic.mz"
	NA                     = "whois.na-nic.com.na"
	NAME                   = "whois.nic.name"
	NC                     = "whois.nc"
	NEWS                   = "whois.nic.news"
	NF                     = "whois.nic.nf"
	NG                     = "whois.nic.net.ng"
	PUBLICINTERESTREGISTRY = "whois.publicinterestregistry.net"
	NL                     = "whois.domain-registry.nl"
	NO                     = "whois.norid.no"
	NRW                    = "whois.nic.nrw"
	NU                     = "whois.iis.nu"
	NYC                    = "whois.nic.nyc"
	NZ                     = "whois.srs.net.nz"
	OM                     = "whois.registry.om"
	ONE                    = "whois.nic.one"
	ONLINE                 = "whois.nic.online"
	OOO                    = "whois.nic.ooo"
	OVH                    = "whois-ovh.nic.fr"
	PARIS                  = "whois-paris.nic.fr"
	PARTY                  = "whois.nic.party"
	PE                     = "kero.yachay.pe"
	PF                     = "whois.registry.pf"
	PHYSIO                 = "whois.nic.physio"
	PLUS                   = "whois.nic.plus"
	PL                     = "whois.dns.pl"
	PM                     = "whois.nic.pm"
	POST                   = "whois.dotpostregistry.net"
	PT                     = "whois.dns.pt"
	RU                     = "whois.tcinet.ru"
	QA                     = "whois.registry.qa"
	QPON                   = "whois.nic.qpon"
	QUEBEC                 = "whois.nic.quebec"
	RACING                 = "whois.nic.racing"
	RE                     = "whois.nic.re"
	REGISTRO               = "whois.registro.br"
	REISE                  = "whois.nic.reise"
	REVIEW                 = "whois.nic.review"
	RO                     = "whois.rotld.ro"
	RODEO                  = "whois.nic.rodeo"
	RS                     = "whois.rnids.rs"
	RIPE                   = "whois.ripe.net"
	RUHR                   = "whois.nic.ruhr"
	RU_NIC                 = "whois.nic.ru"
	SAMSUNG                = "whois.nic.samsung"
	SB                     = "whois.nic.sb"
	SCA                    = "whois.nic.sca"
	SCB                    = "whois.nic.scb"
	SCHMIDT                = "whois.nic.schmidt"
	SCOT                   = "whois.nic.scot"
	SCIENCE                = "whois.nic.science"
	SE                     = "whois.iis.se"
	SG                     = "whois.sgnic.sg"
	SH                     = "whois.nic.sh"
	SI                     = "whois.arnes.si"
	SK                     = "whois.sk-nic.sk"
	SKY                    = "whois.nic.sky"
	SM                     = "whois.nic.sm"
	SN                     = "whois.nic.sn"
	SO                     = "whois.nic.so"
	ST                     = "whois.nic.st"
	STREAM                 = "whois.nic.stream"
	STUDY                  = "whois.nic.study"
	SUCKS                  = "whois.nic.sucks"
	SURF                   = "whois.nic.surf"
	SX                     = "whois.sx"
	SY                     = "whois.tld.sy"
	SYDNEY                 = "whois.nic.sydney"
	TAIPEI                 = "whois.nic.taipei"
	TATAR                  = "whois.nic.tatar"
	TC                     = "whois.nic.tc"
	TEL                    = "whois.nic.tel"
	TF                     = "whois.nic.tf"
	TH                     = "whois.thnic.co.th"
	THNIC                  = "whois.thnic.net"
	TIROL                  = "whois.nic.tirol"
	TK                     = "whois.dot.tk"
	TL                     = "whois.nic.tl"
	TM                     = "whois.nic.tm"
	TN                     = "whois.ati.tn"
	TOP                    = "whois.nic.top"
	TR                     = "whois.nic.tr"
	TR_EDU                 = "whois.metu.edu.tr"
	TRADE                  = "whois.nic.trade"
	TRAVEL                 = "whois.nic.travel"
	TRUST                  = "whois.nic.trust"
	TV                     = "tvwhois.verisign-grs.com"
	TW                     = "whois.twnic.net.tw"
	TWNIC                  = "whois.twnic.net"
	TZ                     = "whois.tznic.or.tz"
	UA                     = "whois.ua"
	UA_BIZ                 = "whois.biz.ua"
	UA_CO                  = "whois.co.ua"
	UG                     = "whois.co.ug"
	UK                     = "whois.nic.uk"
	UK_JA                  = "whois.ja.net"
	UNO                    = "whois.nic.uno"
	US                     = "whois.nic.us"
	UY                     = "whois.nic.org.uy"
	UZ                     = "whois.cctld.uz"
	VERSICHERUNG           = "whois.nic.versicherung"
	VG                     = "whois.nic.vg"
	VIP                    = "whois.nic.vip"
	VLAANDEREN             = "whois.nic.vlaanderen"
	VODKA                  = "whois.nic.vodka"
	WALES                  = "whois.nic.wales"
	WANG                   = "whois.gtld.knet.cn"
	WEBCAM                 = "whois.nic.webcam"
	WED                    = "whois.nic.wed"
	WEDDING                = "whois.nic.wedding"
	WF                     = "whois.nic.wf"
	WHOSWHO                = "whois.nic.whoswho"
	WIEN                   = "whois.nic.wien"
	WIN                    = "whois.nic.win"
	WORK                   = "whois.nic.work"
	WS                     = "whois.website.ws"
	WTC                    = "whois.nic.wtc"
	PS                     = "whois.pnina.ps"
	XXX                    = "whois.nic.xxx"
	YOGA                   = "whois.nic.yoga"
	YT                     = "whois.nic.yt"
	ZA_CO                  = "whois.registry.net.za"
	ZA_NET                 = "net-whois.registry.net.za"
	ZA_ORG                 = "org-whois.registry.net.za"
	ZA_WEB                 = "web-whois.registry.net.za"
	ZM                     = "whois.nic.zm"
	XN45Q11C               = "whois.nic.xn--45q11c"
	XN80ADXHKS             = "whois.nic.xn--80adxhks"
	XN80ASEHDB             = "whois.nic.xn--80asehdb"
	XN80ASWG               = "whois.nic.xn--80aswg"
	XND1ACJ3B              = "whois.nic.xn--d1acj3b"
	XNHXT814E              = "whois.nic.xn--hxt814e"
	XNJ1AMH                = "whois.dotukr.com"
	XNMXTQ1M               = "whois.nic.xn--mxtq1m"
	XNNGBC5AZD             = "whois.nic.xn--ngbc5azd"
	XNP1ACF                = "whois.nic.xn--p1acf"
	XNTCKWE                = "whois.nic.xn--tckwe"
	CNNIC                  = "cwhois.cnnic.cn"
	CONAC                  = "whois.conac.cn"
	NGTLD                  = "whois.ngtld.cn"
	AGITSYS                = "whois.agitsys.net"
	BR_NIC                 = "whois.gtlds.nic.br"
	KNET                   = "whois.gtld.knet.cn"
	KSREGISTRY             = "whois.ksregistry.net"
	GMONIC                 = "whois.nic.gmo"
	UNIREGISTRY            = "whois.uniregistry.net"
	CENTRALNIC             = "whois.centralnic.com"
	GOOGLE                 = "whois.nic.google"
	AFILIAS_SRS            = "whois.afilias-srs.net"
	DONUTS                 = "whois.donuts.co"
	INTERNIC               = "whois.internic.net"
	VERISIGN_GRS           = "whois.verisign-grs.com"
	VOTING                 = "whois.voting.tld-box.at"
	OSAKA                  = "whois.nic.osaka"

	IP_WHOIS_SERVER      = "whois.iana.org"
	DEFAULT_WHOIS_SERVER = "whois-servers.net"
	// TODO: Add domains server list
	//PH
	//TT
	//"http://www.za.net/cgi-bin/whois.cgi?domain=
	//"http://whois.teleinfo.cn/main.action?queryType=0&tp=domain&queryData=
)

var server = map[string]string{
	".abogado":                  ABOGADO,
	".ac":                       AC,
	".ac.ac":                    AC,
	".co.ac":                    AC,
	".gv.ac":                    AC,
	".or.ac":                    AC,
	".account":                  ACCOUNT,
	".ad":                       AE,
	".xn--mgbaam7a8h":           AE,
	".aero":                     AERO,
	".ag":                       AG,
	".ai":                       AI,
	".alsace":                   ALSACE,
	".am":                       AMNIC,
	".amsterdam":                AMSTERDAM,
	".aquarelle":                AQUARELLE,
	".as":                       AS,
	".asia":                     ASIA,
	".at":                       AT,
	".ac.at":                    AT,
	".co.at":                    AT,
	".gv.at":                    AT,
	".or.at":                    AT,
	".asn.au":                   AUNIC,
	".com.au":                   AUNIC,
	".edu.au":                   AUNIC,
	".id.au":                    AUNIC,
	".net.au":                   AUNIC,
	".org.au":                   AUNIC,
	".asso.fr":                  FR,
	".fr":                       FR,
	".presse.fr":                FR,
	".tm.fr":                    FR,
	".au":                       AU,
	".aw":                       AW,
	".ax":                       AX,
	".bank":                     BANK,
	".bar":                      BAR,
	".barclaycard":              BARCLAYCARD,
	".barclays":                 BARCLAYS,
	".bayern":                   BAYERN,
	".be":                       BE,
	".ac.be":                    BE,
	".beer":                     BEER,
	".berlin":                   BERLIN,
	".bi":                       BI,
	".bid":                      BID,
	".biz":                      BIZ,
	".bj":                       BJ,
	".blog":                     BLOG,
	".br":                       REGISTRO,
	".adm.br":                   REGISTRO,
	".adv.br":                   REGISTRO,
	".am.br":                    REGISTRO,
	".arq.br":                   REGISTRO,
	".art.br":                   REGISTRO,
	".bio.br":                   REGISTRO,
	".cng.br":                   REGISTRO,
	".cnt.br":                   REGISTRO,
	".com.br":                   REGISTRO,
	".ecn.br":                   REGISTRO,
	".eng.br":                   REGISTRO,
	".esp.br":                   REGISTRO,
	".etc.br":                   REGISTRO,
	".eti.br":                   REGISTRO,
	".fm.br":                    REGISTRO,
	".fot.br":                   REGISTRO,
	".fst.br":                   REGISTRO,
	".g12.br":                   REGISTRO,
	".gov.br":                   REGISTRO,
	".ind.br":                   REGISTRO,
	".inf.br":                   REGISTRO,
	".jor.br":                   REGISTRO,
	".lel.br":                   REGISTRO,
	".med.br":                   REGISTRO,
	".mil.br":                   REGISTRO,
	".net.br":                   REGISTRO,
	".nom.br":                   REGISTRO,
	".ntr.br":                   REGISTRO,
	".odo.br":                   REGISTRO,
	".org.br":                   REGISTRO,
	".ppg.br":                   REGISTRO,
	".pro.br":                   REGISTRO,
	".psc.br":                   REGISTRO,
	".psi.br":                   REGISTRO,
	".rec.br":                   REGISTRO,
	".slg.br":                   REGISTRO,
	".tmp.br":                   REGISTRO,
	".tur.br":                   REGISTRO,
	".tv.br":                    REGISTRO,
	".vet.br":                   REGISTRO,
	".zlg.br":                   REGISTRO,
	".brussels":                 BRUSSELS,
	".budapest":                 BUDAPEST,
	".build":                    BUILD,
	".buzz":                     BUZZ,
	".bw":                       BW,
	".by":                       BY,
	".bz":                       AFILIAS_GRS,
	".mn":                       AFILIAS_GRS,
	".bzh":                      BZH,
	".ca":                       CA,
	".cancerresearch":           CANCERRESEARCH,
	".capetown":                 CAPETOWN,
	".career":                   CAREER,
	".casa":                     CASA,
	".cat":                      CAT,
	".cc":                       VERISIGN_GRS,
	".ch":                       CH,
	".ci":                       CI,
	".cl":                       CL,
	".cloud":                    CLOUD,
	".club":                     CLUB,
	".cm":                       CM,
	".cn":                       CNNIC,
	".ac.cn":                    CNNIC,
	".ah.cn":                    CNNIC,
	".bj.cn":                    CNNIC,
	".com.cn":                   CNNIC,
	".cq.cn":                    CNNIC,
	".edu.cn":                   CNNIC,
	".gd.cn":                    CNNIC,
	".gov.cn":                   CNNIC,
	".gs.cn":                    CNNIC,
	".gx.cn":                    CNNIC,
	".gz.cn":                    CNNIC,
	".hb.cn":                    CNNIC,
	".he.cn":                    CNNIC,
	".hi.cn":                    CNNIC,
	".hk.cn":                    CNNIC,
	".hl.cn":                    CNNIC,
	".hn.cn":                    CNNIC,
	".jl.cn":                    CNNIC,
	".js.cn":                    CNNIC,
	".ln.cn":                    CNNIC,
	".mo.cn":                    CNNIC,
	".net.cn":                   CNNIC,
	".nm.cn":                    CNNIC,
	".nx.cn":                    CNNIC,
	".org.cn":                   CNNIC,
	".qh.cn":                    CNNIC,
	".sc.cn":                    CNNIC,
	".sh.cn":                    CNNIC,
	".sn.cn":                    CNNIC,
	".tj.cn":                    CNNIC,
	".tw.cn":                    CNNIC,
	".xj.cn":                    CNNIC,
	".xz.cn":                    CNNIC,
	".yn.cn":                    CNNIC,
	".zj.cn":                    CNNIC,
	".co":                       CO,
	".arts.co":                  CO,
	".co.co":                    CO,
	".com.co":                   CO,
	".edu.co":                   CO,
	".firm.co":                  CO,
	".gov.co":                   CO,
	".info.co":                  CO,
	".int.co":                   CO,
	".mil.co":                   CO,
	".net.co":                   CO,
	".nom.co":                   CO,
	".org.co":                   CO,
	".rec.co":                   CO,
	".web.co":                   CO,
	".cologne":                  COLOGNE,
	".com":                      VERISIGN_GRS,
	".net":                      VERISIGN_GRS,
	".cooking":                  COOKING,
	".coop":                     COOP,
	".cr":                       CR,
	".ac.cr":                    CR,
	".co.cr":                    CR,
	".ed.cr":                    CR,
	".fi.cr":                    CR,
	".go.cr":                    CR,
	".or.cr":                    CR,
	".sa.cr":                    CR,
	".cricket":                  CRICKET,
	".cuisinella":               CUISINELLA,
	".cx":                       CX,
	".cymru":                    CYMRU,
	".cz":                       CZ,
	".date":                     DATE,
	".de":                       DE,
	".dk":                       DK,
	".dm":                       DM,
	".do":                       DO,
	".com.do":                   DO,
	".net.do":                   DO,
	".org.do":                   DO,
	".edu.do":                   DO,
	".gov.do":                   DO,
	".gob.do":                   DO,
	".download":                 DOWNLOAD,
	".durban":                   DURBAN,
	".dz":                       DZ,
	".com.dz":                   DZ,
	".net.dz":                   DZ,
	".org.dz":                   DZ,
	".gov.dz":                   DZ,
	".edu.dz":                   DZ,
	".asso.dz":                  DZ,
	".art.dz":                   DZ,
	".pol.dz":                   DZ,
	".xn--lgbbat1ad8j":          DZ,
	".edu":                      INTERNIC,
	".mil":                      INTERNIC,
	".ee":                       EE,
	".eu":                       EU,
	".eu.lv":                    BIZ,
	".eu.org":                   EU_ORG,
	".eurovision":               EUROVISION,
	".eus":                      EUS,
	".faith":                    FAITH,
	".fashion":                  FASHION,
	".fi":                       FI,
	".film":                     FILM,
	".firmdale":                 FIRMDALE,
	".fishing":                  FISHING,
	".fit":                      FIT,
	".frl":                      FRL,
	".frogans":                  FROGANS,
	".ga":                       GA,
	".gal":                      GAL,
	".games":                    GAMES,
	".garden":                   GARDEN,
	".gd":                       GD,
	".gdn":                      GDN,
	".gent":                     GENT,
	".gg":                       GG,
	".gi":                       AFILIAS_GRS2,
	".lc":                       AFILIAS_GRS2,
	".sc":                       AFILIAS_GRS2,
	".vc":                       AFILIAS_GRS2,
	".com.vc":                   AFILIAS_GRS2,
	".net.vc":                   AFILIAS_GRS2,
	".org.vc":                   AFILIAS_GRS2,
	".gov.vc":                   AFILIAS_GRS2,
	".gl":                       GL,
	".global":                   GLOBAL,
	".gmx":                      GMX,
	".gold":                     GOLD,
	".gop":                      GOP,
	".gov":                      GOV,
	".gq":                       GQ,
	".gy":                       GY,
	".hamburg":                  HAMBURG,
	".hk":                       HK,
	".xn--j6w193g":              HK,
	".com.hk":                   HK,
	".edu.hk":                   HK,
	".net.hk":                   HK,
	".org.hk":                   HK,
	".hn":                       HN,
	".hourse":                   HOURSE,
	".hr":                       HR,
	".ht":                       HT,
	".hu":                       HU,
	".ibm":                      IBM,
	".id":                       ID,
	".co.id":                    ID,
	".net.id":                   ID,
	".or.id":                    ID,
	".ac.id":                    ID,
	".go.id":                    ID,
	".mil.id":                   ID,
	".web.id":                   ID,
	".my.id":                    ID,
	".sch.id":                   ID,
	".biz.id":                   ID,
	".desa.id":                  ID,
	".ie":                       IE,
	".ifm":                      IFM,
	".il":                       IL,
	".ac.il":                    IL,
	".co.il":                    IL,
	".gov.il":                   IL,
	".k12.il":                   IL,
	".muni.il":                  IL,
	".net.il":                   IL,
	".org.il":                   IL,
	".im":                       IM,
	".in":                       IN,
	".ac.in":                    IN,
	".co.in":                    IN,
	".ernet.in":                 IN,
	".gov.in":                   IN,
	".net.in":                   IN,
	".org.in":                   IN,
	".res.in":                   IN,
	".xn--2scrj9c":              IN,
	".xn--3hcrj9c":              IN,
	".xn--45br5cyl":             IN,
	".xn--45brj9c":              IN,
	".xn--fpcrj9c3d":            IN,
	".xn--gecrj9c":              IN,
	".xn--h2breg3eve":           IN,
	".xn--h2brj9c":              IN,
	".xn--h2brj9c8c":            IN,
	".xn--rvc1e0am3e":           IN,
	".xn--s9brj9c":              IN,
	".xn--xkc2dl3a5ee0h":        IN,
	".int":                      INT,
	".io":                       IO,
	".iq":                       IQ,
	".ir":                       IR,
	".xn--mgba3a4f16a":          IR,
	".is":                       IS,
	".it":                       IT,
	".java":                     JAVA,
	".je":                       JE,
	".jetzt":                    JETZT,
	".jp":                       JP,
	".ac.jp":                    JP,
	".ad.jp":                    JP,
	".co.jp":                    JP,
	".ed.jp":                    JP,
	".go.jp":                    JP,
	".gr.jp":                    JP,
	".lg.jp":                    JP,
	".ne.jp":                    JP,
	".or.jp":                    JP,
	".hokkaido.jp":              JP,
	".aomori.jp":                JP,
	".iwate.jp":                 JP,
	".miyagi.jp":                JP,
	".akita.jp":                 JP,
	".yamagata.jp":              JP,
	".fukushima.jp":             JP,
	".ibaraki.jp":               JP,
	".tochigi.jp":               JP,
	".gunma.jp":                 JP,
	".saitama.jp":               JP,
	".chiba.jp":                 JP,
	".tokyo.jp":                 JP,
	".kanagawa.jp":              JP,
	".niigata.jp":               JP,
	".toyama.jp":                JP,
	".ishikawa.jp":              JP,
	".fukui.jp":                 JP,
	".yamanashi.jp":             JP,
	".nagano.jp":                JP,
	".gifu.jp":                  JP,
	".shizuoka.jp":              JP,
	".aichi.jp":                 JP,
	".mie.jp":                   JP,
	".shiga.jp":                 JP,
	".kyoto.jp":                 JP,
	".osaka.jp":                 JP,
	".hyogo.jp":                 JP,
	".nara.jp":                  JP,
	".wakayama.jp":              JP,
	".tottori.jp":               JP,
	".shimane.jp":               JP,
	".okayama.jp":               JP,
	".hiroshima.jp":             JP,
	".yamaguchi.jp":             JP,
	".tokushima.jp":             JP,
	".kagawa.jp":                JP,
	".ehime.jp":                 JP,
	".kochi.jp":                 JP,
	".fukuoka.jp":               JP,
	".saga.jp":                  JP,
	".nagasaki.jp":              JP,
	".kumamoto.jp":              JP,
	".oita.jp":                  JP,
	".miyazaki.jp":              JP,
	".kagoshima.jp":             JP,
	".okinawa.jp":               JP,
	".jobs":                     JOBS,
	".joburg":                   JOBURG,
	".ke":                       KE,
	".ac.ke":                    KE,
	".co.ke":                    KE,
	".kg":                       KG,
	".com.kg":                   KG,
	".edu.kg":                   KG,
	".gov.kg":                   KG,
	".mil.kg":                   KG,
	".net.kg":                   KG,
	".org.kg":                   KG,
	".ki":                       KI,
	".kiwi":                     KIWI,
	".koeln":                    KOELN,
	".kr":                       KR,
	".xn--3e0b707e":             KR,
	".xn--cg4bki":               KR,
	".ac.kr":                    KR_OR,
	".co.kr":                    KR_OR,
	".go.kr":                    KR_OR,
	".ne.kr":                    KR_OR,
	".nm.kr":                    KR_OR,
	".or.kr":                    KR_OR,
	".re.kr":                    KR_OR,
	".ky":                       KY,
	".kz":                       KZ,
	".xn--80ao21a":              KZ,
	".la":                       LA,
	".lacaixa":                  LACAIXA,
	".lat":                      LAT,
	".latrobe":                  LATROBE,
	".leclerc":                  LECLERC,
	".li":                       LI,
	".live":                     LIVE,
	".loan":                     LOAN,
	".london":                   LONDON,
	".lt":                       LT,
	".lu":                       LU,
	".luxe":                     LUXE,
	".luxury":                   LUXURY,
	".ma":                       MA,
	".madrid":                   MADRID,
	".mango":                    MANGO,
	".md":                       MD,
	".me":                       ME,
	".men":                      MEN,
	".menu":                     MENU,
	".mg":                       MG,
	".miami":                    MIAMI,
	".mk":                       MK,
	".com.mk":                   MK,
	".net.mk":                   MK,
	".inf.mk":                   MK,
	".org.mk":                   MK,
	".edu.mk":                   MK,
	".xn--d1alf":                MK,
	".ml":                       ML,
	".mo":                       MO,
	".moe":                      MOE,
	".monash":                   MONASH,
	".moscow":                   MOSCOW,
	".ms":                       MS,
	".mu":                       MU,
	".museum":                   MUSEUM,
	".mx":                       MX,
	".com.mx":                   MX,
	".edu.mx":                   MX,
	".gob.mx":                   MX,
	".gov.mx":                   MX,
	".net.mx":                   MX,
	".org.mx":                   MX,
	".mz":                       MZ,
	".gov.mz":                   MZ,
	".na":                       NA,
	".name":                     NAME,
	".nc":                       NC,
	".news":                     NEWS,
	".nf":                       NF,
	".ng":                       NG,
	".com.ng":                   NG,
	".org.ng":                   NG,
	".gov.ng":                   NG,
	".edu.ng":                   NG,
	".net.ng":                   NG,
	".sch.ng":                   NG,
	".name.ng":                  NG,
	".mobi.ng":                  NG,
	".mil.ng":                   NG,
	".i.ng":                     NG,
	".ngo":                      PUBLICINTERESTREGISTRY,
	".ong":                      PUBLICINTERESTREGISTRY,
	".org":                      PUBLICINTERESTREGISTRY,
	".xn--c1avg":                PUBLICINTERESTREGISTRY,
	".xn--i1b6b1a6a2e":          PUBLICINTERESTREGISTRY,
	".xn--nqv7f":                PUBLICINTERESTREGISTRY,
	".nl":                       NL,
	".no":                       NO,
	".nrw":                      NRW,
	".nu":                       NU,
	".nyc":                      NYC,
	".nz":                       NZ,
	".ac.nz":                    NZ,
	".co.nz":                    NZ,
	".geek.nz":                  NZ,
	".gen.nz":                   NZ,
	".iwi.nz":                   NZ,
	".kiwi.nz":                  NZ,
	".maori.nz":                 NZ,
	".net.nz":                   NZ,
	".org.nz":                   NZ,
	".school.nz":                NZ,
	".om":                       OM,
	".co.om":                    OM,
	".com.om":                   OM,
	".org.om":                   OM,
	".net.om":                   OM,
	".edu.om":                   OM,
	".gov.om":                   OM,
	".museum.om":                OM,
	".pro.om":                   OM,
	".med.om":                   OM,
	".xn--mgb9awbf":             OM,
	".one":                      ONE,
	".online":                   ONLINE,
	".ooo":                      OOO,
	".osaka":                    OSAKA,
	".ovh":                      OVH,
	".paris":                    PARIS,
	".party":                    PARTY,
	".pe":                       PE,
	".com.pe":                   PE,
	".net.pe":                   PE,
	".org.pe":                   PE,
	".nom.pe":                   PE,
	".pf":                       PF,
	".physio":                   PHYSIO,
	".plus":                     PLUS,
	".pl":                       PL,
	".agro.pl":                  PL,
	".aid.pl":                   PL,
	".atm.pl":                   PL,
	".augustow.pl":              PL,
	".auto.pl":                  PL,
	".babia-gora.pl":            PL,
	".bedzin.pl":                PL,
	".beskidy.pl":               PL,
	".bialowieza.pl":            PL,
	".bialystok.pl":             PL,
	".bielawa.pl":               PL,
	".bieszczady.pl":            PL,
	".biz.pl":                   PL,
	".boleslawiec.pl":           PL,
	".bydgoszcz.pl":             PL,
	".bytom.pl":                 PL,
	".cieszyn.pl":               PL,
	".com.pl":                   PL,
	".czeladz.pl":               PL,
	".czest.pl":                 PL,
	".dlugoleka.pl":             PL,
	".edu.pl":                   PL,
	".elblag.pl":                PL,
	".elk.pl":                   PL,
	".glogow.pl":                PL,
	".gmina.pl":                 PL,
	".gniezno.pl":               PL,
	".gorlice.pl":               PL,
	".grajewo.pl":               PL,
	".gsm.pl":                   PL,
	".ilawa.pl":                 PL,
	".info.pl":                  PL,
	".jaworzno.pl":              PL,
	".jelenia-gora.pl":          PL,
	".jgora.pl":                 PL,
	".kalisz.pl":                PL,
	".karpacz.pl":               PL,
	".kartuzy.pl":               PL,
	".kaszuby.pl":               PL,
	".katowice.pl":              PL,
	".kazimierz-dolny.pl":       PL,
	".kepno.pl":                 PL,
	".ketrzyn.pl":               PL,
	".klodzko.pl":               PL,
	".kobierzyce.pl":            PL,
	".kolobrzeg.pl":             PL,
	".konin.pl":                 PL,
	".konskowola.pl":            PL,
	".kutno.pl":                 PL,
	".lapy.pl":                  PL,
	".lebork.pl":                PL,
	".legnica.pl":               PL,
	".lezajsk.pl":               PL,
	".limanowa.pl":              PL,
	".lomza.pl":                 PL,
	".lowicz.pl":                PL,
	".lubin.pl":                 PL,
	".lukow.pl":                 PL,
	".mail.pl":                  PL,
	".malbork.pl":               PL,
	".malopolska.pl":            PL,
	".mazowsze.pl":              PL,
	".mazury.pl":                PL,
	".media.pl":                 PL,
	".miasta.pl":                PL,
	".mielec.pl":                PL,
	".mielno.pl":                PL,
	".mil.pl":                   PL,
	".mragowo.pl":               PL,
	".naklo.pl":                 PL,
	".net.pl":                   PL,
	".nieruchomosci.pl":         PL,
	".nom.pl":                   PL,
	".nowaruda.pl":              PL,
	".nysa.pl":                  PL,
	".olawa.pl":                 PL,
	".olecko.pl":                PL,
	".olkusz.pl":                PL,
	".olsztyn.pl":               PL,
	".opoczno.pl":               PL,
	".opole.pl":                 PL,
	".org.pl":                   PL,
	".ostroda.pl":               PL,
	".ostroleka.pl":             PL,
	".ostrowiec.pl":             PL,
	".ostrowwlkp.pl":            PL,
	".pc.pl":                    PL,
	".pila.pl":                  PL,
	".pisz.pl":                  PL,
	".podhale.pl":               PL,
	".podlasie.pl":              PL,
	".polkowice.pl":             PL,
	".pomorskie.pl":             PL,
	".pomorze.pl":               PL,
	".powiat.pl":                PL,
	".priv.pl":                  PL,
	".prochowice.pl":            PL,
	".pruszkow.pl":              PL,
	".przeworsk.pl":             PL,
	".pulawy.pl":                PL,
	".radom.pl":                 PL,
	".rawa-maz.pl":              PL,
	".realestate.pl":            PL,
	".rel.pl":                   PL,
	".rybnik.pl":                PL,
	".rzeszow.pl":               PL,
	".sanok.pl":                 PL,
	".sejny.pl":                 PL,
	".sex.pl":                   PL,
	".shop.pl":                  PL,
	".sklep.pl":                 PL,
	".skoczow.pl":               PL,
	".slask.pl":                 PL,
	".slupsk.pl":                PL,
	".sos.pl":                   PL,
	".sosnowiec.pl":             PL,
	".stalowa-wola.pl":          PL,
	".starachowice.pl":          PL,
	".stargard.pl":              PL,
	".suwalki.pl":               PL,
	".swidnica.pl":              PL,
	".swiebodzin.pl":            PL,
	".swinoujscie.pl":           PL,
	".szczecin.pl":              PL,
	".szczytno.pl":              PL,
	".szkola.pl":                PL,
	".targi.pl":                 PL,
	".tarnobrzeg.pl":            PL,
	".tgory.pl":                 PL,
	".tm.pl":                    PL,
	".tourism.pl":               PL,
	".travel.pl":                PL,
	".turek.pl":                 PL,
	".turystyka.pl":             PL,
	".tychy.pl":                 PL,
	".ustka.pl":                 PL,
	".walbrzych.pl":             PL,
	".warmia.pl":                PL,
	".warszawa.pl":              PL,
	".waw.pl":                   PL,
	".wegrow.pl":                PL,
	".wielun.pl":                PL,
	".wlocl.pl":                 PL,
	".wloclawek.pl":             PL,
	".wodzislaw.pl":             PL,
	".wolomin.pl":               PL,
	".wroclaw.pl":               PL,
	".zachpomor.pl":             PL,
	".zagan.pl":                 PL,
	".zarow.pl":                 PL,
	".zgora.pl":                 PL,
	".zgorzelec.pl":             PL,
	".pm":                       PM,
	".post":                     POST,
	".pt":                       PT,
	".com.pt":                   PT,
	".edu.pt":                   PT,
	".gov.pt":                   PT,
	".int.pt":                   PT,
	".net.pt":                   PT,
	".nome.pt":                  PT,
	".org.pt":                   PT,
	".publ.pt":                  PT,
	".net.ru":                   RU,
	".ru":                       RU,
	".su":                       RU,
	".xn--p1ai":                 RU,
	".qa":                       QA,
	".com.qa":                   QA,
	".net.qa":                   QA,
	".org.qa":                   QA,
	".edu.qa":                   QA,
	".sch.qa":                   QA,
	".gov.qa":                   QA,
	".xn--wgbl6a":               QA,
	".qpon":                     QPON,
	".quebec":                   QUEBEC,
	".racing":                   RACING,
	".re":                       RE,
	".reise":                    REISE,
	".review":                   REVIEW,
	".ro":                       RO,
	".arts.ro":                  RO,
	".com.ro":                   RO,
	".firm.ro":                  RO,
	".info.ro":                  RO,
	".nom.ro":                   RO,
	".nt.ro":                    RO,
	".org.ro":                   RO,
	".rec.ro":                   RO,
	".store.ro":                 RO,
	".tm.ro":                    RO,
	".www.ro":                   RO,
	".rodeo":                    RODEO,
	".rs":                       RS,
	".co.rs":                    RS,
	".edu.rs":                   RS,
	".in.rs":                    RS,
	".org.rs":                   RS,
	".com.ru":                   RU_NIC,
	".msk.ru":                   RU_NIC,
	".org.ru":                   RU_NIC,
	".pp.ru":                    RU_NIC,
	".sochi.su":                 RU_NIC,
	".spb.ru":                   RU_NIC,
	".ruhr":                     RUHR,
	".samsung":                  SAMSUNG,
	".sb":                       SB,
	".sca":                      SCA,
	".scb":                      SCB,
	".schmidt":                  SCHMIDT,
	".science":                  SCIENCE,
	".scot":                     SCOT,
	".se":                       SE,
	".sg":                       SG,
	".com.sg":                   SG,
	".edu.sg":                   SG,
	".gov.sg":                   SG,
	".net.sg":                   SG,
	".org.sg":                   SG,
	".xn--clchc0ea0b2g2a9gcd":   SG,
	".xn--yfro4i67o":            SG,
	".sh":                       SH,
	".si":                       SI,
	".sky":                      SKY,
	".sm":                       SM,
	".sn":                       SN,
	".so":                       SO,
	".st":                       ST,
	".stream":                   STREAM,
	".study":                    STUDY,
	".sucks":                    SUCKS,
	".surf":                     SURF,
	".sx":                       SX,
	".sy":                       SY,
	".xn--ogbpf8fl":             SY,
	".sydney":                   SYDNEY,
	".taipei":                   TAIPEI,
	".tatar":                    TATAR,
	".tc":                       TC,
	".tel":                      TEL,
	".tf":                       TF,
	".th":                       THNIC,
	".xn--o3cw4h":               TH,
	".ac.th":                    THNIC,
	".co.th":                    THNIC,
	".go.th":                    THNIC,
	".in.th":                    THNIC,
	".mi.th":                    THNIC,
	".net.th":                   THNIC,
	".or.th":                    THNIC,
	".tirol":                    TIROL,
	".tk":                       TK,
	".tl":                       TL,
	".tm":                       TM,
	".tn":                       TN,
	".com.tn":                   TN,
	".ens.tn":                   TN,
	".fin.tn":                   TN,
	".gov.tn":                   TN,
	".ind.tn":                   TN,
	".intl.tn":                  TN,
	".nat.tn":                   TN,
	".net.tn":                   TN,
	".org.tn":                   TN,
	".info.tn":                  TN,
	".perso.tn":                 TN,
	".tourism.tn":               TN,
	".edunet.tn":                TN,
	".rnrt.tn":                  TN,
	".rns.tn":                   TN,
	".rnu.tn":                   TN,
	".mincom.tn":                TN,
	".agrinet.tn":               TN,
	".turen.tn":                 TN,
	".top":                      TOP,
	".tr":                       TR,
	".bbs.tr":                   TR_EDU,
	".com.tr":                   TR_EDU,
	".edu.tr":                   TR_EDU,
	".gov.tr":                   TR_EDU,
	".k12.tr":                   TR_EDU,
	".mil.tr":                   TR_EDU,
	".net.tr":                   TR_EDU,
	".org.tr":                   TR_EDU,
	".trade":                    TRADE,
	".travel":                   TRAVEL,
	".trust":                    TRUST,
	".tv":                       TV,
	".tw":                       TW,
	".xn--kprw13d":              TW,
	".xn--kpry57d":              TW,
	".com.tw":                   TWNIC,
	".net.tw":                   TWNIC,
	".org.tw":                   TWNIC,
	".tz":                       TZ,
	".co.tz":                    TZ,
	".or.tz":                    TZ,
	".ac.tz":                    TZ,
	".go.tz":                    TZ,
	".hotel.tz":                 TZ,
	".info.tz":                  TZ,
	".me.tz":                    TZ,
	".mobi.tz":                  TZ,
	".ne.tz":                    TZ,
	".sc.tz":                    TZ,
	".tv.tz":                    TZ,
	".ua":                       UA,
	".com.ua":                   UA,
	".dn.ua":                    UA,
	".kh.ua":                    UA,
	".kiev.ua":                  UA,
	".lg.ua":                    UA,
	".lviv.ua":                  UA,
	".net.ua":                   UA,
	".org.ua":                   UA,
	".kharkov.ua":               UA,
	".kherson.ua":               UA,
	".edu.ua":                   UA,
	".in.ua":                    UA,
	".biz.ua":                   UA_BIZ,
	".co.ua":                    UA_CO,
	".ug":                       UG,
	".ac.ug":                    UG,
	".co.ug":                    UG,
	".go.ug":                    UG,
	".ne.ug":                    UG,
	".or.ug":                    UG,
	".org.ug":                   UG,
	".sc.ug":                    UG,
	".uk":                       UK,
	".co.uk":                    UK,
	".ltd.uk":                   UK,
	".me.uk":                    UK,
	".net.uk":                   UK,
	".org.uk":                   UK,
	".plc.uk":                   UK,
	".ac.uk":                    UK_JA,
	".gov.uk":                   UK_JA,
	".uno":                      UNO,
	".us":                       US,
	".uy":                       UY,
	".uz":                       UZ,
	".versicherung":             VERSICHERUNG,
	".vg":                       VG,
	".vip":                      VIP,
	".vlaanderen":               VLAANDEREN,
	".vodka":                    VODKA,
	".voting":                   VOTING,
	".wales":                    WALES,
	".wang":                     KNET,
	".xn--3bst00m":              KNET,
	".xn--6qq986b3xl":           KNET,
	".xn--czru2d":               KNET,
	".xn--fiq64b":               KNET,
	".webcam":                   WEBCAM,
	".wed":                      WED,
	".wedding":                  WEDDING,
	".wf":                       WF,
	".whoswho":                  WHOSWHO,
	".wien":                     WIEN,
	".win":                      WIN,
	".work":                     WORK,
	".ws":                       WS,
	".wtc":                      WTC,
	".ps":                       PS,
	".xn--ygbi2ammx":            PS,
	".xxx":                      XXX,
	".yoga":                     YOGA,
	".yt":                       YT,
	".co.za":                    ZA_CO,
	".net.za":                   ZA_NET,
	".org.za":                   ZA_ORG,
	".web.za":                   ZA_WEB,
	".zm":                       ZM,
	".xn--45q11c":               XN45Q11C,
	".xn--80adxhks":             XN80ADXHKS,
	".xn--80asehdb":             XN80ASEHDB,
	".xn--80aswg":               XN80ASWG,
	".xn--d1acj3b":              XND1ACJ3B,
	".xn--hxt814e":              XNHXT814E,
	".xn--j1amh":                XNJ1AMH,
	".xn--mxtq1m":               XNMXTQ1M,
	".xn--ngbc5azd":             XNNGBC5AZD,
	".xn--p1acf":                XNP1ACF,
	".xn--tckwe":                XNTCKWE,
	".xn--fiqs8s":               CNNIC,
	".xn--fiqz9s":               CNNIC,
	".xn--55qw42g":              CONAC,
	".xn--zfr164b":              CONAC,
	".xn--1qqw23a":              NGTLD,
	".xn--55qx5d":               NGTLD,
	".xn--io0a7i":               NGTLD,
	".xn--xhq521b":              NGTLD,
	".nowruz":                   AGITSYS,
	".pars":                     AGITSYS,
	".shia":                     AGITSYS,
	".tci":                      AGITSYS,
	".bom":                      BR_NIC,
	".final":                    BR_NIC,
	".globo":                    BR_NIC,
	".rio":                      BR_NIC,
	".uol":                      BR_NIC,
	".baidu":                    KNET,
	".xn--30rr7y":               KNET,
	".xn--9et52u":               KNET,
	".allfinanz":                KSREGISTRY,
	".bio":                      KSREGISTRY,
	".bmw":                      KSREGISTRY,
	".cam":                      KSREGISTRY,
	".desi":                     KSREGISTRY,
	".dvag":                     KSREGISTRY,
	".flsmidth":                 KSREGISTRY,
	".mini":                     KSREGISTRY,
	".pohl":                     KSREGISTRY,
	".saarland":                 KSREGISTRY,
	".spiegel":                  KSREGISTRY,
	".tui":                      KSREGISTRY,
	".xn--vermgensberater-ctb":  KSREGISTRY,
	".xn--vermgensberatung-pwb": KSREGISTRY,
	".zuerich":                  KSREGISTRY,
	".canon":                    GMONIC,
	".datsun":                   GMONIC,
	".fujitsu":                  GMONIC,
	".ggee":                     GMONIC,
	".goldpoint":                GMONIC,
	".goo":                      GMONIC,
	".hisamitsu":                GMONIC,
	".hitachi":                  GMONIC,
	".infiniti":                 GMONIC,
	".jcb":                      GMONIC,
	".kddi":                     GMONIC,
	".kyoto":                    GMONIC,
	".lotte":                    GMONIC,
	".mitsubishi":               GMONIC,
	".mtpc":                     GMONIC,
	".nagoya":                   GMONIC,
	".nico":                     GMONIC,
	".nissan":                   GMONIC,
	".okinawa":                  GMONIC,
	".panasonic":                GMONIC,
	".pioneer":                  GMONIC,
	".sharp":                    GMONIC,
	".shop":                     GMONIC,
	".tokyo":                    GMONIC,
	".toshiba":                  GMONIC,
	".yodobashi":                GMONIC,
	".yokohama":                 GMONIC,
	".archi":                    AFILIAS,
	".bet":                      AFILIAS,
	".black":                    AFILIAS,
	".blue":                     AFILIAS,
	".green":                    AFILIAS,
	".info":                     AFILIAS,
	".kim":                      AFILIAS,
	".lgbt":                     AFILIAS,
	".llc":                      AFILIAS,
	".lotto":                    AFILIAS,
	".mobi":                     AFILIAS,
	".organic":                  AFILIAS,
	".pet":                      AFILIAS,
	".pink":                     AFILIAS,
	".poker":                    AFILIAS,
	".pro":                      AFILIAS,
	".promo":                    AFILIAS,
	".red":                      AFILIAS,
	".shiksha":                  AFILIAS,
	".ski":                      AFILIAS,
	".vote":                     AFILIAS,
	".voto":                     AFILIAS,
	".xn--6frz82g":              AFILIAS,
	".audio":                    UNIREGISTRY,
	".auto":                     UNIREGISTRY,
	".blackfriday":              UNIREGISTRY,
	".car":                      UNIREGISTRY,
	".cars":                     UNIREGISTRY,
	".christmas":                UNIREGISTRY,
	".click":                    UNIREGISTRY,
	".diet":                     UNIREGISTRY,
	".flowers":                  UNIREGISTRY,
	".gift":                     UNIREGISTRY,
	".guitars":                  UNIREGISTRY,
	".help":                     UNIREGISTRY,
	".hiphop":                   UNIREGISTRY,
	".hosting":                  UNIREGISTRY,
	".juegos":                   UNIREGISTRY,
	".link":                     UNIREGISTRY,
	".lol":                      UNIREGISTRY,
	".mom":                      UNIREGISTRY,
	".photo":                    UNIREGISTRY,
	".pics":                     UNIREGISTRY,
	".property":                 UNIREGISTRY,
	".sexy":                     UNIREGISTRY,
	".tattoo":                   UNIREGISTRY,
	".ae.com":                   CENTRALNIC,
	".ae.org":                   CENTRALNIC,
	".ar.com":                   CENTRALNIC,
	".ar":                       CENTRALNIC,
	".com.ar":                   CENTRALNIC,
	".net.ar":                   CENTRALNIC,
	".org.ar":                   CENTRALNIC,
	".bet.ar":                   CENTRALNIC,
	".edu.ar":                   CENTRALNIC,
	".gob.ar":                   CENTRALNIC,
	".int.ar":                   CENTRALNIC,
	".mil.ar":                   CENTRALNIC,
	".musica.ar":                CENTRALNIC,
	".tur.ar":                   CENTRALNIC,
	".art":                      CENTRALNIC,
	".best":                     CENTRALNIC,
	".br.com":                   CENTRALNIC,
	".ceo":                      CENTRALNIC,
	".college":                  CENTRALNIC,
	".cn.com":                   CENTRALNIC,
	".cyou":                     CENTRALNIC,
	".de.com":                   CENTRALNIC,
	".com.de":                   CENTRALNIC,
	".design":                   CENTRALNIC,
	".eu.com":                   CENTRALNIC,
	".fans":                     CENTRALNIC,
	".feedback":                 CENTRALNIC,
	".fm":                       CENTRALNIC,
	".fo":                       CENTRALNIC,
	".fun":                      CENTRALNIC,
	".gb.com":                   CENTRALNIC,
	".gb.net":                   CENTRALNIC,
	".gr.com":                   CENTRALNIC,
	".hu.com":                   CENTRALNIC,
	".hu.net":                   CENTRALNIC,
	".host":                     CENTRALNIC,
	".in.net":                   CENTRALNIC,
	".icu":                      CENTRALNIC,
	".ink":                      CENTRALNIC,
	".jp.net":                   CENTRALNIC,
	".jpn.com":                  CENTRALNIC,
	".kr.com":                   CENTRALNIC,
	".love":                     CENTRALNIC,
	".mex.com":                  CENTRALNIC,
	".no.com":                   CENTRALNIC,
	".press":                    CENTRALNIC,
	".pw":                       CENTRALNIC,
	".qc.com":                   CENTRALNIC,
	".reit":                     CENTRALNIC,
	".rent":                     CENTRALNIC,
	".rest":                     CENTRALNIC,
	".ru.com":                   CENTRALNIC,
	".sa.com":                   CENTRALNIC,
	".se.com":                   CENTRALNIC,
	".se.net":                   CENTRALNIC,
	".site":                     CENTRALNIC,
	".space":                    CENTRALNIC,
	".store":                    CENTRALNIC,
	".tech":                     CENTRALNIC,
	".uk.com":                   CENTRALNIC,
	".uk.net":                   CENTRALNIC,
	".us.com":                   CENTRALNIC,
	".us.org":                   CENTRALNIC,
	".uy.com":                   CENTRALNIC,
	".website":                  CENTRALNIC,
	".web.com":                  CENTRALNIC,
	".wiki":                     CENTRALNIC,
	".wme":                      CENTRALNIC,
	".xyz":                      CENTRALNIC,
	".za.com":                   CENTRALNIC,
	".ads":                      GOOGLE,
	".android":                  GOOGLE,
	".app":                      GOOGLE,
	".boo":                      GOOGLE,
	".cal":                      GOOGLE,
	".channel":                  GOOGLE,
	".chrome":                   GOOGLE,
	".dad":                      GOOGLE,
	".day":                      GOOGLE,
	".dclk":                     GOOGLE,
	".dev":                      GOOGLE,
	".docs":                     GOOGLE,
	".drive":                    GOOGLE,
	".eat":                      GOOGLE,
	".esq":                      GOOGLE,
	".fly":                      GOOGLE,
	".foo":                      GOOGLE,
	".gbiz":                     GOOGLE,
	".gle":                      GOOGLE,
	".gmail":                    GOOGLE,
	".goog":                     GOOGLE,
	".google":                   GOOGLE,
	".guge":                     GOOGLE,
	".hangout":                  GOOGLE,
	".here":                     GOOGLE,
	".how":                      GOOGLE,
	".ing":                      GOOGLE,
	".map":                      GOOGLE,
	".meet":                     GOOGLE,
	".meme":                     GOOGLE,
	".mov":                      GOOGLE,
	".new":                      GOOGLE,
	".nexus":                    GOOGLE,
	".page":                     GOOGLE,
	".phd":                      GOOGLE,
	".play":                     GOOGLE,
	".prod":                     GOOGLE,
	".prof":                     GOOGLE,
	".rsvp":                     GOOGLE,
	".search":                   GOOGLE,
	".soy":                      GOOGLE,
	".xn--flw351e":              GOOGLE,
	".xn--q9jyb4c":              GOOGLE,
	".xn--qcka1pmc":             GOOGLE,
	".youtube":                  GOOGLE,
	".zip":                      GOOGLE,
	".abarth":                   AFILIAS_SRS,
	".abbott":                   AFILIAS_SRS,
	".abbvie":                   AFILIAS_SRS,
	".aco":                      AFILIAS_SRS,
	".adult":                    AFILIAS_SRS,
	".agakhan":                  AFILIAS_SRS,
	".aigo":                     AFILIAS_SRS,
	".akdn":                     AFILIAS_SRS,
	".alfaromeo":                AFILIAS_SRS,
	".allstate":                 AFILIAS_SRS,
	".apple":                    AFILIAS_SRS,
	".audi":                     AFILIAS_SRS,
	".autos":                    AFILIAS_SRS,
	".avianca":                  AFILIAS_SRS,
	".beats":                    AFILIAS_SRS,
	".bnpparibas":               AFILIAS_SRS,
	".boats":                    AFILIAS_SRS,
	".boehringer":               AFILIAS_SRS,
	".bugatti":                  AFILIAS_SRS,
	".buy":                      AFILIAS_SRS,
	".cbs":                      AFILIAS_SRS,
	".ceb":                      AFILIAS_SRS,
	".cern":                     AFILIAS_SRS,
	".chrysler":                 AFILIAS_SRS,
	".cipriani":                 AFILIAS_SRS,
	".creditunion":              AFILIAS_SRS,
	".dabur":                    AFILIAS_SRS,
	".dodge":                    AFILIAS_SRS,
	".eco":                      AFILIAS_SRS,
	".edeka":                    AFILIAS_SRS,
	".emerck":                   AFILIAS_SRS,
	".esurance":                 AFILIAS_SRS,
	".extraspace":               AFILIAS_SRS,
	".fage":                     AFILIAS_SRS,
	".fiat":                     AFILIAS_SRS,
	".fido":                     AFILIAS_SRS,
	".gea":                      AFILIAS_SRS,
	".godaddy":                  AFILIAS_SRS,
	".hermes":                   AFILIAS_SRS,
	".homes":                    AFILIAS_SRS,
	".imamat":                   AFILIAS_SRS,
	".ismaili":                  AFILIAS_SRS,
	".ist":                      AFILIAS_SRS,
	".istanbul":                 AFILIAS_SRS,
	".itv":                      AFILIAS_SRS,
	".jcp":                      AFILIAS_SRS,
	".jeep":                     AFILIAS_SRS,
	".jll":                      AFILIAS_SRS,
	".lamborghini":              AFILIAS_SRS,
	".lancia":                   AFILIAS_SRS,
	".lasalle":                  AFILIAS_SRS,
	".ltda":                     AFILIAS_SRS,
	".marriott":                 AFILIAS_SRS,
	".mit":                      AFILIAS_SRS,
	".mopar":                    AFILIAS_SRS,
	".motorcycles":              AFILIAS_SRS,
	".natura":                   AFILIAS_SRS,
	".nokia":                    AFILIAS_SRS,
	".nra":                      AFILIAS_SRS,
	".onl":                      AFILIAS_SRS,
	".porn":                     AFILIAS_SRS,
	".pr":                       AFILIAS_SRS,
	".progressive":              AFILIAS_SRS,
	".pwc":                      AFILIAS_SRS,
	".redumbrella":              AFILIAS_SRS,
	".rich":                     AFILIAS_SRS,
	".rogers":                   AFILIAS_SRS,
	".schaeffler":               AFILIAS_SRS,
	".sew":                      AFILIAS_SRS,
	".sex":                      AFILIAS_SRS,
	".shaw":                     AFILIAS_SRS,
	".showtime":                 AFILIAS_SRS,
	".shriram":                  AFILIAS_SRS,
	".srl":                      AFILIAS_SRS,
	".srt":                      AFILIAS_SRS,
	".stada":                    AFILIAS_SRS,
	".stockholm":                AFILIAS_SRS,
	".temasek":                  AFILIAS_SRS,
	".travelers":                AFILIAS_SRS,
	".travelersinsurance":       AFILIAS_SRS,
	".trv":                      AFILIAS_SRS,
	".uconnect":                 AFILIAS_SRS,
	".vegas":                    AFILIAS_SRS,
	".vig":                      AFILIAS_SRS,
	".viking":                   AFILIAS_SRS,
	".volkswagen":               AFILIAS_SRS,
	".xn--4gbrim":               AFILIAS_SRS,
	".xn--b4w605ferd":           AFILIAS_SRS,
	".yachts":                   AFILIAS_SRS,
	".zara":                     AFILIAS_SRS,
	".academy":                  DONUTS,
	".accountants":              DONUTS,
	".actor":                    DONUTS,
	".agency":                   DONUTS,
	".airforce":                 DONUTS,
	".apartments":               DONUTS,
	".army":                     DONUTS,
	".associates":               DONUTS,
	".attorney":                 DONUTS,
	".auction":                  DONUTS,
	".band":                     DONUTS,
	".bargains":                 DONUTS,
	".bike":                     DONUTS,
	".bingo":                    DONUTS,
	".boutique":                 DONUTS,
	".builders":                 DONUTS,
	".business":                 DONUTS,
	".cab":                      DONUTS,
	".cafe":                     DONUTS,
	".camera":                   DONUTS,
	".camp":                     DONUTS,
	".capital":                  DONUTS,
	".cards":                    DONUTS,
	".care":                     DONUTS,
	".careers":                  DONUTS,
	".cash":                     DONUTS,
	".casino":                   DONUTS,
	".catering":                 DONUTS,
	".center":                   DONUTS,
	".chat":                     DONUTS,
	".cheap":                    DONUTS,
	".church":                   DONUTS,
	".city":                     DONUTS,
	".claims":                   DONUTS,
	".cleaning":                 DONUTS,
	".clinic":                   DONUTS,
	".clothing":                 DONUTS,
	".coach":                    DONUTS,
	".codes":                    DONUTS,
	".coffee":                   DONUTS,
	".community":                DONUTS,
	".company":                  DONUTS,
	".computer":                 DONUTS,
	".condos":                   DONUTS,
	".construction":             DONUTS,
	".consulting":               DONUTS,
	".contractors":              DONUTS,
	".cool":                     DONUTS,
	".coupons":                  DONUTS,
	".credit":                   DONUTS,
	".creditcard":               DONUTS,
	".cruises":                  DONUTS,
	".dance":                    DONUTS,
	".dating":                   DONUTS,
	".deals":                    DONUTS,
	".degree":                   DONUTS,
	".delivery":                 DONUTS,
	".democrat":                 DONUTS,
	".dental":                   DONUTS,
	".dentist":                  DONUTS,
	".diamonds":                 DONUTS,
	".digital":                  DONUTS,
	".direct":                   DONUTS,
	".directory":                DONUTS,
	".discount":                 DONUTS,
	".doctor":                   DONUTS,
	".dog":                      DONUTS,
	".domains":                  DONUTS,
	".education":                DONUTS,
	".email":                    DONUTS,
	".energy":                   DONUTS,
	".engineer":                 DONUTS,
	".engineering":              DONUTS,
	".enterprises":              DONUTS,
	".equipment":                DONUTS,
	".estate":                   DONUTS,
	".events":                   DONUTS,
	".exchange":                 DONUTS,
	".expert":                   DONUTS,
	".exposed":                  DONUTS,
	".fail":                     DONUTS,
	".farm":                     DONUTS,
	".finance":                  DONUTS,
	".financial":                DONUTS,
	".fish":                     DONUTS,
	".fitness":                  DONUTS,
	".flights":                  DONUTS,
	".florist":                  DONUTS,
	".football":                 DONUTS,
	".forsale":                  DONUTS,
	".foundation":               DONUTS,
	".fund":                     DONUTS,
	".furniture":                DONUTS,
	".futbol":                   DONUTS,
	".gallery":                  DONUTS,
	".gifts":                    DONUTS,
	".gives":                    DONUTS,
	".glass":                    DONUTS,
	".gmbh":                     DONUTS,
	".graphics":                 DONUTS,
	".gratis":                   DONUTS,
	".gripe":                    DONUTS,
	".group":                    DONUTS,
	".guide":                    DONUTS,
	".guru":                     DONUTS,
	".haus":                     DONUTS,
	".healthcare":               DONUTS,
	".holdings":                 DONUTS,
	".holiday":                  DONUTS,
	".hospital":                 DONUTS,
	".house":                    DONUTS,
	".immo":                     DONUTS,
	".immobilien":               DONUTS,
	".industries":               DONUTS,
	".institute":                DONUTS,
	".insure":                   DONUTS,
	".international":            DONUTS,
	".investments":              DONUTS,
	".irish":                    DONUTS,
	".kaufen":                   DONUTS,
	".kitchen":                  DONUTS,
	".land":                     DONUTS,
	".lawyer":                   DONUTS,
	".lease":                    DONUTS,
	".legal":                    DONUTS,
	".life":                     DONUTS,
	".lighting":                 DONUTS,
	".limited":                  DONUTS,
	".limo":                     DONUTS,
	".loans":                    DONUTS,
	".ltd":                      DONUTS,
	".maison":                   DONUTS,
	".management":               DONUTS,
	".market":                   DONUTS,
	".marketing":                DONUTS,
	".media":                    DONUTS,
	".memorial":                 DONUTS,
	".moda":                     DONUTS,
	".money":                    DONUTS,
	".mortgage":                 DONUTS,
	".movie":                    DONUTS,
	".navy":                     DONUTS,
	".network":                  DONUTS,
	".ninja":                    DONUTS,
	".partners":                 DONUTS,
	".parts":                    DONUTS,
	".photography":              DONUTS,
	".photos":                   DONUTS,
	".pictures":                 DONUTS,
	".pizza":                    DONUTS,
	".place":                    DONUTS,
	".plumbing":                 DONUTS,
	".productions":              DONUTS,
	".properties":               DONUTS,
	".pub":                      DONUTS,
	".recipes":                  DONUTS,
	".rehab":                    DONUTS,
	".reisen":                   DONUTS,
	".rentals":                  DONUTS,
	".repair":                   DONUTS,
	".report":                   DONUTS,
	".republican":               DONUTS,
	".restaurant":               DONUTS,
	".reviews":                  DONUTS,
	".rip":                      DONUTS,
	".rocks":                    DONUTS,
	".run":                      DONUTS,
	".sale":                     DONUTS,
	".salon":                    DONUTS,
	".sarl":                     DONUTS,
	".school":                   DONUTS,
	".schule":                   DONUTS,
	".services":                 DONUTS,
	".shoes":                    DONUTS,
	".shopping":                 DONUTS,
	".singles":                  DONUTS,
	".social":                   DONUTS,
	".software":                 DONUTS,
	".solar":                    DONUTS,
	".solutions":                DONUTS,
	".style":                    DONUTS,
	".supplies":                 DONUTS,
	".supply":                   DONUTS,
	".support":                  DONUTS,
	".surgery":                  DONUTS,
	".systems":                  DONUTS,
	".tax":                      DONUTS,
	".taxi":                     DONUTS,
	".team":                     DONUTS,
	".technology":               DONUTS,
	".tennis":                   DONUTS,
	".theater":                  DONUTS,
	".tienda":                   DONUTS,
	".tips":                     DONUTS,
	".tires":                    DONUTS,
	".today":                    DONUTS,
	".tools":                    DONUTS,
	".tours":                    DONUTS,
	".town":                     DONUTS,
	".toys":                     DONUTS,
	".training":                 DONUTS,
	".university":               DONUTS,
	".vacations":                DONUTS,
	".ventures":                 DONUTS,
	".vet":                      DONUTS,
	".viajes":                   DONUTS,
	".video":                    DONUTS,
	".villas":                   DONUTS,
	".vin":                      DONUTS,
	".vision":                   DONUTS,
	".voyage":                   DONUTS,
	".watch":                    DONUTS,
	".works":                    DONUTS,
	".world":                    DONUTS,
	".wtf":                      DONUTS,
	".xn--czrs0t":               DONUTS,
	".xn--fjq720a":              DONUTS,
	".xn--unup4y":               DONUTS,
	".xn--vhquv":                DONUTS,
	".zone":                     DONUTS,
}

func GetWhoisServer(tld string) string {
	value, ok := server[tld]
	if !ok {
		ip := net.ParseIP(tld)
		if ip != nil {
			return IP_WHOIS_SERVER
		} else {
			return DEFAULT_WHOIS_SERVER
		}
	}
	return value
}
