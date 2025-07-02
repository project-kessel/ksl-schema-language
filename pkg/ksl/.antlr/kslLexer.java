// Generated from /Users/jmarcant/Documents/ksl-schema-language/pkg/ksl/ksl.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class kslLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		KW_ESC=1, FORCED_NAME=2, VERSION=3, VERSIONNUM=4, RESOLVE=5, NAMESPACE=6, 
		ACCESS=7, PUBLIC=8, INTERNAL=9, PRIVATE=10, TYPE=11, RELATION=12, IMPORT=13, 
		EXTENSION=14, CARDINALITY=15, ATMOSTONE=16, EXACTLYONE=17, ATLEASTONE=18, 
		ANY=19, AS=20, AND=21, OR=22, UNLESS=23, ALLOW_DUPLICATES=24, EXPAND=25, 
		LBRACE=26, RBRACE=27, EXTENSION_CALL=28, LPAREN=29, RPAREN=30, LSQUARE=31, 
		RSQARE=32, VARREF=33, TEMPLATE_DELIM=34, STRING_DELIM=35, ARG_DELIM=36, 
		DECL_END=37, NAME=38, COMMENT=39, WS=40;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"KW_ESC", "FORCED_NAME", "VERSION", "VERSIONNUM", "RESOLVE", "NAMESPACE", 
			"ACCESS", "PUBLIC", "INTERNAL", "PRIVATE", "TYPE", "RELATION", "IMPORT", 
			"EXTENSION", "CARDINALITY", "ATMOSTONE", "EXACTLYONE", "ATLEASTONE", 
			"ANY", "AS", "AND", "OR", "UNLESS", "ALLOW_DUPLICATES", "EXPAND", "LBRACE", 
			"RBRACE", "EXTENSION_CALL", "LPAREN", "RPAREN", "LSQUARE", "RSQARE", 
			"VARREF", "TEMPLATE_DELIM", "STRING_DELIM", "ARG_DELIM", "DECL_END", 
			"NAME", "COMMENT", "WS"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'#'", null, "'version'", null, "'.'", "'namespace'", null, "'public'", 
			"'internal'", "'private'", "'type'", "'relation'", "'import'", "'extension'", 
			null, "'AtMostOne'", "'ExactlyOne'", "'AtLeastOne'", "'Any'", "'as'", 
			"'and'", "'or'", "'unless'", "'allow_duplicates'", "':'", "'{'", "'}'", 
			"'@'", "'('", "')'", "'['", "']'", "'$'", "'`'", "'''", "','", "';'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "KW_ESC", "FORCED_NAME", "VERSION", "VERSIONNUM", "RESOLVE", "NAMESPACE", 
			"ACCESS", "PUBLIC", "INTERNAL", "PRIVATE", "TYPE", "RELATION", "IMPORT", 
			"EXTENSION", "CARDINALITY", "ATMOSTONE", "EXACTLYONE", "ATLEASTONE", 
			"ANY", "AS", "AND", "OR", "UNLESS", "ALLOW_DUPLICATES", "EXPAND", "LBRACE", 
			"RBRACE", "EXTENSION_CALL", "LPAREN", "RPAREN", "LSQUARE", "RSQARE", 
			"VARREF", "TEMPLATE_DELIM", "STRING_DELIM", "ARG_DELIM", "DECL_END", 
			"NAME", "COMMENT", "WS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public kslLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "ksl.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2*\u012e\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\3\2\3\2\3\3\3"+
		"\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\6\5d\n\5\r\5\16\5e"+
		"\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\5\bw\n\b"+
		"\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3"+
		"\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3"+
		"\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\5\20\u00b4"+
		"\n\20\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\23\3\23\3\23\3\23\3\23\3\23"+
		"\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\26\3\26"+
		"\3\26\3\26\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\31\3\31"+
		"\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31"+
		"\3\31\3\32\3\32\3\33\3\33\3\34\3\34\3\35\3\35\3\36\3\36\3\37\3\37\3 \3"+
		" \3!\3!\3\"\3\"\3#\3#\3$\3$\3%\3%\3&\3&\3\'\3\'\7\'\u0118\n\'\f\'\16\'"+
		"\u011b\13\'\3(\3(\3(\3(\7(\u0121\n(\f(\16(\u0124\13(\3(\3(\3)\6)\u0129"+
		"\n)\r)\16)\u012a\3)\3)\2\2*\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25"+
		"\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32"+
		"\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*\3\2\7\3\2\62;\5\2C"+
		"\\aac|\6\2\62;C\\aac|\4\2\f\f\17\17\5\2\13\f\17\17\"\"\2\u0136\2\3\3\2"+
		"\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17"+
		"\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2"+
		"\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3"+
		"\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3"+
		"\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2"+
		"=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3"+
		"\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\3S\3\2\2\2\5U\3\2\2"+
		"\2\7X\3\2\2\2\t`\3\2\2\2\13g\3\2\2\2\ri\3\2\2\2\17v\3\2\2\2\21x\3\2\2"+
		"\2\23\177\3\2\2\2\25\u0088\3\2\2\2\27\u0090\3\2\2\2\31\u0095\3\2\2\2\33"+
		"\u009e\3\2\2\2\35\u00a5\3\2\2\2\37\u00b3\3\2\2\2!\u00b5\3\2\2\2#\u00bf"+
		"\3\2\2\2%\u00ca\3\2\2\2\'\u00d5\3\2\2\2)\u00d9\3\2\2\2+\u00dc\3\2\2\2"+
		"-\u00e0\3\2\2\2/\u00e3\3\2\2\2\61\u00ea\3\2\2\2\63\u00fb\3\2\2\2\65\u00fd"+
		"\3\2\2\2\67\u00ff\3\2\2\29\u0101\3\2\2\2;\u0103\3\2\2\2=\u0105\3\2\2\2"+
		"?\u0107\3\2\2\2A\u0109\3\2\2\2C\u010b\3\2\2\2E\u010d\3\2\2\2G\u010f\3"+
		"\2\2\2I\u0111\3\2\2\2K\u0113\3\2\2\2M\u0115\3\2\2\2O\u011c\3\2\2\2Q\u0128"+
		"\3\2\2\2ST\7%\2\2T\4\3\2\2\2UV\5\3\2\2VW\5M\'\2W\6\3\2\2\2XY\7x\2\2YZ"+
		"\7g\2\2Z[\7t\2\2[\\\7u\2\2\\]\7k\2\2]^\7q\2\2^_\7p\2\2_\b\3\2\2\2`a\t"+
		"\2\2\2ac\5\13\6\2bd\t\2\2\2cb\3\2\2\2de\3\2\2\2ec\3\2\2\2ef\3\2\2\2f\n"+
		"\3\2\2\2gh\7\60\2\2h\f\3\2\2\2ij\7p\2\2jk\7c\2\2kl\7o\2\2lm\7g\2\2mn\7"+
		"u\2\2no\7r\2\2op\7c\2\2pq\7e\2\2qr\7g\2\2r\16\3\2\2\2sw\5\21\t\2tw\5\23"+
		"\n\2uw\5\25\13\2vs\3\2\2\2vt\3\2\2\2vu\3\2\2\2w\20\3\2\2\2xy\7r\2\2yz"+
		"\7w\2\2z{\7d\2\2{|\7n\2\2|}\7k\2\2}~\7e\2\2~\22\3\2\2\2\177\u0080\7k\2"+
		"\2\u0080\u0081\7p\2\2\u0081\u0082\7v\2\2\u0082\u0083\7g\2\2\u0083\u0084"+
		"\7t\2\2\u0084\u0085\7p\2\2\u0085\u0086\7c\2\2\u0086\u0087\7n\2\2\u0087"+
		"\24\3\2\2\2\u0088\u0089\7r\2\2\u0089\u008a\7t\2\2\u008a\u008b\7k\2\2\u008b"+
		"\u008c\7x\2\2\u008c\u008d\7c\2\2\u008d\u008e\7v\2\2\u008e\u008f\7g\2\2"+
		"\u008f\26\3\2\2\2\u0090\u0091\7v\2\2\u0091\u0092\7{\2\2\u0092\u0093\7"+
		"r\2\2\u0093\u0094\7g\2\2\u0094\30\3\2\2\2\u0095\u0096\7t\2\2\u0096\u0097"+
		"\7g\2\2\u0097\u0098\7n\2\2\u0098\u0099\7c\2\2\u0099\u009a\7v\2\2\u009a"+
		"\u009b\7k\2\2\u009b\u009c\7q\2\2\u009c\u009d\7p\2\2\u009d\32\3\2\2\2\u009e"+
		"\u009f\7k\2\2\u009f\u00a0\7o\2\2\u00a0\u00a1\7r\2\2\u00a1\u00a2\7q\2\2"+
		"\u00a2\u00a3\7t\2\2\u00a3\u00a4\7v\2\2\u00a4\34\3\2\2\2\u00a5\u00a6\7"+
		"g\2\2\u00a6\u00a7\7z\2\2\u00a7\u00a8\7v\2\2\u00a8\u00a9\7g\2\2\u00a9\u00aa"+
		"\7p\2\2\u00aa\u00ab\7u\2\2\u00ab\u00ac\7k\2\2\u00ac\u00ad\7q\2\2\u00ad"+
		"\u00ae\7p\2\2\u00ae\36\3\2\2\2\u00af\u00b4\5!\21\2\u00b0\u00b4\5#\22\2"+
		"\u00b1\u00b4\5%\23\2\u00b2\u00b4\5\'\24\2\u00b3\u00af\3\2\2\2\u00b3\u00b0"+
		"\3\2\2\2\u00b3\u00b1\3\2\2\2\u00b3\u00b2\3\2\2\2\u00b4 \3\2\2\2\u00b5"+
		"\u00b6\7C\2\2\u00b6\u00b7\7v\2\2\u00b7\u00b8\7O\2\2\u00b8\u00b9\7q\2\2"+
		"\u00b9\u00ba\7u\2\2\u00ba\u00bb\7v\2\2\u00bb\u00bc\7Q\2\2\u00bc\u00bd"+
		"\7p\2\2\u00bd\u00be\7g\2\2\u00be\"\3\2\2\2\u00bf\u00c0\7G\2\2\u00c0\u00c1"+
		"\7z\2\2\u00c1\u00c2\7c\2\2\u00c2\u00c3\7e\2\2\u00c3\u00c4\7v\2\2\u00c4"+
		"\u00c5\7n\2\2\u00c5\u00c6\7{\2\2\u00c6\u00c7\7Q\2\2\u00c7\u00c8\7p\2\2"+
		"\u00c8\u00c9\7g\2\2\u00c9$\3\2\2\2\u00ca\u00cb\7C\2\2\u00cb\u00cc\7v\2"+
		"\2\u00cc\u00cd\7N\2\2\u00cd\u00ce\7g\2\2\u00ce\u00cf\7c\2\2\u00cf\u00d0"+
		"\7u\2\2\u00d0\u00d1\7v\2\2\u00d1\u00d2\7Q\2\2\u00d2\u00d3\7p\2\2\u00d3"+
		"\u00d4\7g\2\2\u00d4&\3\2\2\2\u00d5\u00d6\7C\2\2\u00d6\u00d7\7p\2\2\u00d7"+
		"\u00d8\7{\2\2\u00d8(\3\2\2\2\u00d9\u00da\7c\2\2\u00da\u00db\7u\2\2\u00db"+
		"*\3\2\2\2\u00dc\u00dd\7c\2\2\u00dd\u00de\7p\2\2\u00de\u00df\7f\2\2\u00df"+
		",\3\2\2\2\u00e0\u00e1\7q\2\2\u00e1\u00e2\7t\2\2\u00e2.\3\2\2\2\u00e3\u00e4"+
		"\7w\2\2\u00e4\u00e5\7p\2\2\u00e5\u00e6\7n\2\2\u00e6\u00e7\7g\2\2\u00e7"+
		"\u00e8\7u\2\2\u00e8\u00e9\7u\2\2\u00e9\60\3\2\2\2\u00ea\u00eb\7c\2\2\u00eb"+
		"\u00ec\7n\2\2\u00ec\u00ed\7n\2\2\u00ed\u00ee\7q\2\2\u00ee\u00ef\7y\2\2"+
		"\u00ef\u00f0\7a\2\2\u00f0\u00f1\7f\2\2\u00f1\u00f2\7w\2\2\u00f2\u00f3"+
		"\7r\2\2\u00f3\u00f4\7n\2\2\u00f4\u00f5\7k\2\2\u00f5\u00f6\7e\2\2\u00f6"+
		"\u00f7\7c\2\2\u00f7\u00f8\7v\2\2\u00f8\u00f9\7g\2\2\u00f9\u00fa\7u\2\2"+
		"\u00fa\62\3\2\2\2\u00fb\u00fc\7<\2\2\u00fc\64\3\2\2\2\u00fd\u00fe\7}\2"+
		"\2\u00fe\66\3\2\2\2\u00ff\u0100\7\177\2\2\u01008\3\2\2\2\u0101\u0102\7"+
		"B\2\2\u0102:\3\2\2\2\u0103\u0104\7*\2\2\u0104<\3\2\2\2\u0105\u0106\7+"+
		"\2\2\u0106>\3\2\2\2\u0107\u0108\7]\2\2\u0108@\3\2\2\2\u0109\u010a\7_\2"+
		"\2\u010aB\3\2\2\2\u010b\u010c\7&\2\2\u010cD\3\2\2\2\u010d\u010e\7b\2\2"+
		"\u010eF\3\2\2\2\u010f\u0110\7)\2\2\u0110H\3\2\2\2\u0111\u0112\7.\2\2\u0112"+
		"J\3\2\2\2\u0113\u0114\7=\2\2\u0114L\3\2\2\2\u0115\u0119\t\3\2\2\u0116"+
		"\u0118\t\4\2\2\u0117\u0116\3\2\2\2\u0118\u011b\3\2\2\2\u0119\u0117\3\2"+
		"\2\2\u0119\u011a\3\2\2\2\u011aN\3\2\2\2\u011b\u0119\3\2\2\2\u011c\u011d"+
		"\7\61\2\2\u011d\u011e\7\61\2\2\u011e\u0122\3\2\2\2\u011f\u0121\n\5\2\2"+
		"\u0120\u011f\3\2\2\2\u0121\u0124\3\2\2\2\u0122\u0120\3\2\2\2\u0122\u0123"+
		"\3\2\2\2\u0123\u0125\3\2\2\2\u0124\u0122\3\2\2\2\u0125\u0126\b(\2\2\u0126"+
		"P\3\2\2\2\u0127\u0129\t\6\2\2\u0128\u0127\3\2\2\2\u0129\u012a\3\2\2\2"+
		"\u012a\u0128\3\2\2\2\u012a\u012b\3\2\2\2\u012b\u012c\3\2\2\2\u012c\u012d"+
		"\b)\2\2\u012dR\3\2\2\2\t\2ev\u00b3\u0119\u0122\u012a\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}