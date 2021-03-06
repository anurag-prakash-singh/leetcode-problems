package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func minDeletionSize(A []string) int {
	// Get runes
	n := len(A[0])
	m := len(A)
	aa := make([][]rune, m)
	mem := make([]int, n)
	lis := 0

	for i := 0; i < m; i++ {
		aa[i] = []rune(A[i])
	}

	for i := 0; i < n; i++ {
		mem[i] = 1
	}

	for i := 0; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			var k int

			for k = 0; k < m; k++ {
				if aa[k][j] > aa[k][i] {
					break
				}
			}

			if k == m {
				mem[i] = max(mem[i], mem[j]+1)
			}
		}

		lis = max(lis, mem[i])
	}

	return n - lis
}

func test1() {
	strs := []string{"babca", "bbazb"}

	fmt.Printf("%d\n", minDeletionSize(strs))
}

func test2() {
	strs := []string{"edcba"}

	fmt.Printf("%d\n", minDeletionSize(strs))
}

func test3() {
	strs := []string{"ghi", "def", "abc"}

	fmt.Printf("%d\n", minDeletionSize(strs))
}

func test4() {
	strs := []string{
		"hcjfndramrebpigadjalqfirbkejhqpqsjddpcfbmmmplechmhnfogaoclhdqcodirtfsjmhodsoftahlirphpfmbm", "agtiehijkliafrdskeoqnhahtgjqksprqeigmhsrpcqhqtgjbtmqqhdfhapbmnpsbpmbcmkggjcdjghfocrjfarqfd", "kdflfmfogdsirhmiotjncerjeberihkllblgbgegkkqbcmhtjdhtdfgbjlkocseottnfqaejfsjgrnjhklddpbfqbi", "fhhbcsjcmmfptbjreqhkrtjqcknktakoejjjjccabaqflrenbnkaathitkncpfctacsrsgmtjirilojfehmcnegnsc", "aljbbogenfotjphqriakrbsrjniiethobjkgjetlsjemasmadlqpsgoidljnbaohqkntapsrocfjsgkfdpgiotadms", "ichcepfekmfrpaeddcjsdobdgsiiafdrkiqeqiisjgifhtqaqelkclodehjftbamokhhkbcrltomjscjpfpoearpff", "qciateapardlstofomrgpdreollksmaleipefqbkmegrakedaihscriksefajrrsagrarqdmfllpnnrkiqmotpoern", "jchgaqnqeqlmmtrksbqnlrtelogkghcheatacpfpsqhilgholtbaaelbikqaignmiljsccpkobcbnbakpcmiplkklp", "ldnkdgdnnmqodjfsdknipsrfbqqcmpqmnkijlqeqmcqbjmgtsesmisrmjfaslqarkkrfjpqqnmlnoheiirhmdleqdj", "rcbkctqjqnfkrlehebjlcgplfqqlipdgjsbetbnrcfqmftfgpkgpbanjcnqrthslelhcrtaccgrildfbiordpsodjr", "sgbhpaiiiqcnjsitateipreptapamkgcgdqqsspftpnijpmifpcalqcqpnfdfnbpjrsrtqaoicsoladnfdtrnegobg", "rgaitbjdsooopddbknlbtpdhssofmhjpmbrttkekalamnriaanjgsmpladfbnfairaigresmjqqtrpamqaoqicradn", "gsmlbsimogchnlhcsgpbcisanghjgrrljdnsmefacsmbhgqmtjpitakqoablegbjdqeeogfmqtfmtfeokqflrmrmbk", "ngkmhhpgspfftmihjdfrnflacgsrcjpnqrjjacfhgklilftiosffjhkenrrkssoqeqcphrnbecltccmrsbertglmoi", "oddpqllessqqmpirgkgptsarjbhjhageoqhqfgmaaelpfcdngsrnfaajqgbnjbmmdornpslmfljtqjabrngoaocncd", "orfsaaienhmofipgrhbmtocjoekipbsibbqhcnkoqfhkrrqjtlfpnokkhclsgcdfoffmhssinpqllohedtlgmdpchj", "jilladsfmgjlrtikapplbmsidkgqldopghpmgqecnegmcknboigemitqjogggescjkeodcmpqgobnrnmojhfcmokog", "nhoffnscqdofcmkqlhtqchelspsnddnpjnfdcfrrhnlpgggptblqsdfilrjjdmgqmpfqralrpddgbhejrhjejsikti", "rgbcnifffocepnsgefcjbfnsfijkfrikehralknakrkmrnpfajreeaireaqtedbrnjhttdspshenrkkhgmdcjamqqf", "flochimpajnnokeceoalptlekfeplhdkiefjdtcadebterefcojqlgoamiblngrhsrpichodqharfnsiliannctmfh", "riqomakgatddqrthfqtrlmgrgsfgcbhsleccckiqchdliifsqcajoplpbamclcqklncggiecdtqphprofbtfpphbho", "isldtcooqtqstaghqiekhsnjrssaehejppriqchajlbipcepcilrlcheeheeqrclbiromdsbqiapoqpecjqngtpmkm", "absahsqbqefhttbcfgatjrnpabtabiojcgdtknslegbigjaparhonkomgbsbgajdkljklaoihmmkaanjjqbbdtjsck", "lcfafbrrsrsaikbtqticjdbabransimkdigbelsogjetmmajpnontblidfpdglftnkngorbihbamtdoqahlmlqeiqb", "irajdrfitlnmldprcembcqhncekckapnlaaerkdclbttcjffhqheahfllppntjqgngkdiimfflgkqmrcfrfeqdqskp", "knoodoafdckccmrrrjetltkgqpptkbhttidhbdciocjamhehrrtciebtohqqkahhdmdjfjlelototgjrifdkhjhfdt", "lebmrmmmbgpkftgkirrsiifncoepolntrjgdrrqmnidecpqgoqibfhijdghranctbcbdqclbhorojlthllhbborkls", "nlbljahtgiblhnlngqpajdodlcdgbdicnjpkcjpjfmefogcprhongeqcbgtnhaippmqarontprfdhedgtkajqkhggp", "rjdftaaldqllknpgmngjorsnqdndfkfmksjttfssqlhbbdbdktmpqjiqoiqrbddsrksncdpshsrtggoiooljemqpls", "glaespcpjttjcbnsknsirkafbdkmcpkifraqanrisfstabboaitioefqcsoqdhlchedgihlnfqhjhjtgncptbkmpme", "bjgdkellqhtohkhtdtroedabrgsfbdjtrnliidbotfdenidgjskggqkqcskhahedicortckfrhklfsjffjidpfkerk", "mortfcsllcrffjrnlceijjpafagoemgteofpfkfndhcikedlhnnlerqaakiietfhpnpqdcpcjbmifjlgijqcrnhebt", "tfoigqlfdbtolagsfetnsninnnefqeirbppmcbkbgorknkiaemhrpqsjnmjsdastcscnmgintdrdkfjpsnbcomdrht", "rdbecoeehseqpbcbemfhirfbjltlnkphdanriqrphttgnktcrbokrihgmjtlkhbppjlsahqkiahefhobmljbrttdik", "janeolfhldpdrjigsgliatkrprlsnjeiralfrakfsrakirsaqjmsagopcelkogbodncpkmrhpppcoblgpisataafia", "nnraaqmsqanbkbsjbrgrkbbphbhpqdrmggkgctpladpljealisijlljkpqotsodbqsrboceatjhattdkoobgokpoms", "lclfifnclefbpcphfkpgbactdkdflnqbcggehcntsojanikgnefjflrqaldnojgiigiamlkiagaiqghrmssfipjjlb", "balkspphlamjofffjlkfsjbbenaoiorojmgsqdbdnqogettjihllmqiljnlorqiqtnahrskqfqsmmfefgmjoqarcmf", "msbcoomrimmoirpmasllhboasjsqjjssighidmtggojorlenafirbrmbfjcjhrflrhgknpatsdbteeoglhliacaolr", "rfhnhijlrrksqmajdebgdhmsmskhdtcmcobntniirmnfgbjlhrbmdbdnmfcmqrafnhorchesetkiflajroshmlfbnj", "andtiboifbkfdhnsjemfbqjhsfmbjajmrtmrhgcogsjdqbiophadnlqbnaoltaaeanitddaijpsflpbsdhlrssrprh", "gflfmarlookeeolakambrkfemmgfejmsfeonlaajojtjcqtcddeihgkelhbbgifkglggirsamacldtjplcgclpmfio", "jfgkmcfkbcpjblnhafskbfjrcbrqroqsqjibsocktendpcpcitdbaegfafcralcnrbaejsjcgrriiciococajkijbq", "npgrsrarcedltgqlpfgfmksmiiarjriakaslnjjoheaqdsqmdpdrnhntmmjaqbjekglfbsrtsecnfsotkfsbjdcdom", "nhokaatqalaldjtipocddnttiktqjksihhjsgnltscslhohikmnmkaostgafcknflkdiksbjdejfsamnlkgrhotnnh", "cogogrfgmfilcqotpkkkmgtqflkgoebbteasroodklmertpphabmptqkceinalposseccerptahnrtksbclddnqfog", "hfqhiapeohsrkibdqdngasfqskkkdmafnflohsdotamhonlbtrgoaoeoefhitlmtkjpceijrcjaflosqrlcgggjnai", "kogsdcmngdpkjthlkmpdqkigeddteacimnpomrbcttrjpkoetcnbecsnmtnmcthqdesodaapomeiqoackjopojcmse", "ojbtrebpegeqjmebhllrlipqirgrndhojohsoiokhoiljjqctcmrctoarqjmbkkpliacfbmnblbdhlsifspfkgmpqh", "creopjdbkklrbrkdnitbcjcorsiethkcrjnsokiphphhabermebpatjsacjmopjjcobcmnrsstjirrrkannbbspcst", "itgdleksltmlpipabfaseqdcmmeakddestjmpmqdcmooepsotdcsrrmljpohqmbnnlbbkmmbnbpjhcmlifsljatrro", "roefoplprksiggjcrqifqjpimnscncgflrglrjlhoprndqihcdlnklecjcomqqondrdelhcejrsfsojnrlrndjmjeq", "flclqbrbokcjcnhlggdjmlhtbrrpfmbsrolnargjdaamnchntaotjjstpbokjlqeqsbqhbjnohisceoltjeofqqlik", "sfqlfofjshlocrqrgjsrgkjtmjddnioclaehffqnncsprgtedbhlfisjontsdfdldlhphqgrksqelpgqoitckqqlge", "jbmggfrdiphdjkgndjkrgttmriglrrbasmppqcnhlqgthkmfnlasihcafdhpiekifaarmqhlirtsadigbbjinlqmis", "dgchrjoidpromqesgkoqkcthtsadanjbbgbskfrrbrcidiodtpdgktareqcpgoemgdhosshnnpqlmjkbspclfmtelp", "baolbeehpcqhpktcbsegfkmqhnkjetqrlrbarabascibnhdbebfagemrdhnplsaihofchqqjkssccegqfejgqlggab", "dafbomnhjhdrglbopqqohtmqtpgfdlqmlkafolsnscjbjeghsgbdolclikrmjjtdjncnepbdfaoolcomtfebipktpa", "lpctlcrdpjnpsntgnhtadfohcithqeidpdkncnecbbdtrtlnobgkhtqobgjqjahknspqhjqnooaihgkanlkbramfhr", "bedfircchocbdejdibgboqlmnakgkjpmmrmgmaicbjqjjtsankjmjlodghqrtaklqollkmnofjlrihcfnhfestbnkg", "rpbmcdpqbbgimpmqfocledtqltctnqlkqkgebkjnrtbaoammopgtmjlsihtdptlgsnrnckmhkfbrknheemaicdkqth", "taagjgfiljohsjlhmrmobarrbngrplnqjrpnfcdpssqtomnobgiaaemhhcnqkrmmbsnjadtlalpiqssertkskbbbit", "sdpicbqhijjqclijonijntdbleosfpmqkoldckbjojabglhscstiothcggapknjsdjqsiqmkkcrgbitinmaqnsjilm", "qfsmbipfjsjknmfkbeofcscnpetgrhdthgbddgbnlnidnprhlrgcdjnlihhcebaseifolrjjpfjglnmjtedttlskqi", "rgnaishdebsmafssaahibsfgdeqoceoeolqthrnrbjthaolssdbpjekejoqjehgjeebgtnniieogalhtjgkkskjlom", "gfesrsjrnfcflhcbtbraiqibeooqsgactiornhasdftjhofnknibqramrcqgrgffrmtnmacqtroorfpoibchodkibo", "ertfqdcmojbcgrgjthsfqgijgjkntmosleprgssrlrfietlblebjsicsqljdcekpilnjncpncjsceobkbrcrlbltnk", "cfmmmiddjstrdgijsmoqtdrtkkocgemqnllbffecmcpbqdacaplhgimrehdelffdscfopsqonhojnfspnkoomsrpdp", "pcskeorjfdelosmqcfdstiafdratqrsfineqocdblnhrgtjfrcaejsbegfgbrmooalpcfrpqmoqhqefdkkkgomqaif", "gtesnmhojapbsatmisrelrotmnfsbmafjscokchfstaroeaosqagdpdidhpanasfsjsrjesfakcsfipnohbprrpors", "ccdqdrfleallhibjtckrfdqcqjoajijnqbhrfolpracrkmaqcpnbforjcpntsstiiabcgtkopmqqqqftledctrirjn", "okqqtalpcpfbesirntlnblpghehrjfmlbrfogqplbmhojopbthrnchrhjrbbmfespnqrpqcpgkpmdsnoatgkopfjbt", "gasrprproqkimdrefgmcrtplaljaplqbnrfodmffbdpnlllgfapdlqahbrbetcethfmkoernpttobhdjdhmidoatao", "ingmlggomeahmspdrhimfhactbnqjqrmrlllbbmpcjltmsokrmlfkodrjtksfffseilkaeaahgfiitotlghkebiifm", "nmtkkegsesklcrkefmbjfdmesqrjhdtgslemcpgambsollqkarjocahklgqirlqesstrnoicsjmgcjlnqgsqsrqekm", "gafhodtnehnjajcrjodtrtlltgpohhlsoidmpfpilnamchbompcjrsaarfhqngdjhjaidmciiafelkkkblieabqssh", "njnfrlinkfqeklitqigontiiidaijaigkhhmnsgoloatbhkkoepfoqcjogpibobrighmdkddststkqknnmhnbsjgkn", "mmtnnefagaifsjajcdfajedbnfdpdmqnpeltcemhpiljhladqpsrgqicfjlnmnbjbqgkgagssmirciofrhokrmodlo", "fpsdmbrjlbiirrdnpjqlghiiaesqlmdoggdccstbcjnehitgjkmbajjrqrrbocclbqoabshklbtrpandfoganingdp", "ljlmtjkddhdpkqteqckmnallghipstiftglctcfktkhedfjjdpsipoejrsfdofhlflqaimtsjjjcclrfjqrbqiokkj", "icnhagcrqdesgjrjstdbpgqcsslolqffoknbtfmdtogiqftrrndtatcecdemccdfocmmfifnljspjrqjkiikhmcshm", "atmmknasfojhnpsqnjhkofjkanmjbitrbclhqjfpmblhofhermmdqfpfleplntfhjrfnqslmnaqhcmotqithpcpmhd", "nfebcjlibadtqnfiotmdiaodenetblohflbbehmqmidqdktehlqcdijdmgebfgdhdsbijokqetsfjddapppkralhgq", "kktophpichiqditdehbfcmmaprgbldgnpeakqhohctotkijldblpsaqbsjakniminglgapolatntmnedjtdobbcpmc", "jakejocankoerkripsdocerbracogetcmrjnhfnhqhkflctqlcfdoidqohclegbhbtckelkrkqfhssmrssbqinldjk", "maninqfrnemhgeorcamclsosikalsckjqsqslrgeoojgeiqorkaoaersaedqlnnjombbakokgmfcmdfikathefgpho", "igsjhjgonicefrmgdnslmckkpiglcqiqhatrpjnfkhpcjaqqtostnearkborljmlmopbrnkpscoqmeplssblehiqld", "mllhrkqjbsgjrsrgqpntqshpcicjshpmrkfffbsarkaohjbcmbadlbkmcfcilcqsebglmqgdatfpgfescasdiftrdg", "edroamrmmjmiaseaejqqrkgqrihnckllehhfopmlndtmnfdcscmaititperniejelsojgosdecdieradkepjbdrpar", "rjmkrkhglojacdcjdabbirmkmbbfqlcgdngqqeohpcgiqlfierabnleetghgbhhlaglgrhqpnlfhmhlmesksbgship", "ttbpkfthsjmrdqomctlethfpctqedcpmonlaplaelafiffibotqngdoeamehcfjrtplnhndoclflfkjdittkoeagfb", "tgheflscfoilrqaskohpqrmtmdbqfadcakclljhqneqsakpogcajklgamfofegiihjiteholmbtroensnnpdrndfoc", "csrehoifqhgbqkftlqpgsknfinpqgckidpitreaogatimfmgcpnnaoqgpttldaqttbckahsgdmodsncehpcejhjknc", "ngcjbjbkgfccgtrjaslnrtshfpaealjfgfatmaoamikfapckeorrlcmtktrnkqtenqigbhqgiqnebboljncnanphgh", "rgkjltcmcmapbblmctigpnkbcosarmlpliahghbashkcmqiqmlrildanhdlmkfldlkmfchpnbcogodpqchfobfngtm", "ciqcshjippbhdbbfqcroejfacmjhapbnlqrlomcctfpefcbqetlrmclqscpjnerppnebrfibcdkqfhkatltfkddfmj", "eaarpieffhkgjqqsmjiefiiaonoirgjippgcqgisfnbefdlbqoespitjcfqkfpdheogihrlecptkgkjmgmtqdfegjs", "nhmnistoqsadpjkkhecorqpbjokriakkqgnkisjgjqesclnafcqcdmsdaecgjcgtreeifcbolcrntjjebmeceqlbct", "fpomtreprhbsdkqrjnctilhpdgdicagfcbmodgkjfjhgqdiimmdjrsisadafmljjkhfnfqgjtmhtodnhaatmiodbtq", "lnqbjqpmlmgsimhtbqlbbtdqblskefibnkobqlllbtqieerkkfhgpkqafkrpiepsqakqndfjikblhhoonsiddghbqc",
	}

	fmt.Printf("%d\n", minDeletionSize(strs))
}

func main() {

	// test1()
	// test2()
	// test3()
	test4()

}
