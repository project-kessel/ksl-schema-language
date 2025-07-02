// Generated from /Users/jmarcant/Documents/ksl-schema-language/pkg/ksl/ksl.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class kslParser extends Parser {
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
	public static final int
		RULE_name = 0, RULE_file = 1, RULE_version = 2, RULE_namespace = 3, RULE_import_stmt = 4, 
		RULE_declaration = 5, RULE_typeExpr = 6, RULE_typeReference = 7, RULE_extensionParam = 8, 
		RULE_extensionParams = 9, RULE_extensionReference = 10, RULE_relation = 11, 
		RULE_relationBody = 12, RULE_paramNames = 13, RULE_extension = 14, RULE_dynamicType = 15, 
		RULE_dynamicRelation = 16, RULE_dynamicName = 17, RULE_dynamicBody = 18;
	private static String[] makeRuleNames() {
		return new String[] {
			"name", "file", "version", "namespace", "import_stmt", "declaration", 
			"typeExpr", "typeReference", "extensionParam", "extensionParams", "extensionReference", 
			"relation", "relationBody", "paramNames", "extension", "dynamicType", 
			"dynamicRelation", "dynamicName", "dynamicBody"
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

	@Override
	public String getGrammarFileName() { return "ksl.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public kslParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class NameContext extends ParserRuleContext {
		public TerminalNode NAME() { return getToken(kslParser.NAME, 0); }
		public TerminalNode FORCED_NAME() { return getToken(kslParser.FORCED_NAME, 0); }
		public NameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_name; }
	}

	public final NameContext name() throws RecognitionException {
		NameContext _localctx = new NameContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_name);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(38);
			_la = _input.LA(1);
			if ( !(_la==FORCED_NAME || _la==NAME) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FileContext extends ParserRuleContext {
		public VersionContext version() {
			return getRuleContext(VersionContext.class,0);
		}
		public NamespaceContext namespace() {
			return getRuleContext(NamespaceContext.class,0);
		}
		public List<Import_stmtContext> import_stmt() {
			return getRuleContexts(Import_stmtContext.class);
		}
		public Import_stmtContext import_stmt(int i) {
			return getRuleContext(Import_stmtContext.class,i);
		}
		public List<DeclarationContext> declaration() {
			return getRuleContexts(DeclarationContext.class);
		}
		public DeclarationContext declaration(int i) {
			return getRuleContext(DeclarationContext.class,i);
		}
		public FileContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_file; }
	}

	public final FileContext file() throws RecognitionException {
		FileContext _localctx = new FileContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_file);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(40);
			version();
			setState(41);
			namespace();
			setState(45);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IMPORT) {
				{
				{
				setState(42);
				import_stmt();
				}
				}
				setState(47);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(49); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(48);
				declaration();
				}
				}
				setState(51); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << ACCESS) | (1L << TYPE) | (1L << EXTENSION) | (1L << EXTENSION_CALL))) != 0) );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class VersionContext extends ParserRuleContext {
		public TerminalNode VERSION() { return getToken(kslParser.VERSION, 0); }
		public TerminalNode VERSIONNUM() { return getToken(kslParser.VERSIONNUM, 0); }
		public VersionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_version; }
	}

	public final VersionContext version() throws RecognitionException {
		VersionContext _localctx = new VersionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_version);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(53);
			match(VERSION);
			setState(54);
			match(VERSIONNUM);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class NamespaceContext extends ParserRuleContext {
		public TerminalNode NAMESPACE() { return getToken(kslParser.NAMESPACE, 0); }
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public NamespaceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_namespace; }
	}

	public final NamespaceContext namespace() throws RecognitionException {
		NamespaceContext _localctx = new NamespaceContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_namespace);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(56);
			match(NAMESPACE);
			setState(57);
			name();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Import_stmtContext extends ParserRuleContext {
		public TerminalNode IMPORT() { return getToken(kslParser.IMPORT, 0); }
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public Import_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_import_stmt; }
	}

	public final Import_stmtContext import_stmt() throws RecognitionException {
		Import_stmtContext _localctx = new Import_stmtContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_import_stmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(59);
			match(IMPORT);
			setState(60);
			name();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DeclarationContext extends ParserRuleContext {
		public TypeExprContext typeExpr() {
			return getRuleContext(TypeExprContext.class,0);
		}
		public TerminalNode DECL_END() { return getToken(kslParser.DECL_END, 0); }
		public ExtensionContext extension() {
			return getRuleContext(ExtensionContext.class,0);
		}
		public ExtensionReferenceContext extensionReference() {
			return getRuleContext(ExtensionReferenceContext.class,0);
		}
		public DeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaration; }
	}

	public final DeclarationContext declaration() throws RecognitionException {
		DeclarationContext _localctx = new DeclarationContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_declaration);
		int _la;
		try {
			setState(73);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(62);
				typeExpr();
				setState(64);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==DECL_END) {
					{
					setState(63);
					match(DECL_END);
					}
				}

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(66);
				extension();
				setState(68);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==DECL_END) {
					{
					setState(67);
					match(DECL_END);
					}
				}

				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(70);
				extensionReference();
				setState(71);
				match(DECL_END);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeExprContext extends ParserRuleContext {
		public TerminalNode TYPE() { return getToken(kslParser.TYPE, 0); }
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public TerminalNode LBRACE() { return getToken(kslParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(kslParser.RBRACE, 0); }
		public List<ExtensionReferenceContext> extensionReference() {
			return getRuleContexts(ExtensionReferenceContext.class);
		}
		public ExtensionReferenceContext extensionReference(int i) {
			return getRuleContext(ExtensionReferenceContext.class,i);
		}
		public TerminalNode ACCESS() { return getToken(kslParser.ACCESS, 0); }
		public List<RelationContext> relation() {
			return getRuleContexts(RelationContext.class);
		}
		public RelationContext relation(int i) {
			return getRuleContext(RelationContext.class,i);
		}
		public TypeExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeExpr; }
	}

	public final TypeExprContext typeExpr() throws RecognitionException {
		TypeExprContext _localctx = new TypeExprContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_typeExpr);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(78);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EXTENSION_CALL) {
				{
				{
				setState(75);
				extensionReference();
				}
				}
				setState(80);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(82);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ACCESS) {
				{
				setState(81);
				match(ACCESS);
				}
			}

			setState(84);
			match(TYPE);
			setState(85);
			name();
			setState(86);
			match(LBRACE);
			setState(90);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << ACCESS) | (1L << RELATION) | (1L << EXTENSION_CALL))) != 0)) {
				{
				{
				setState(87);
				relation();
				}
				}
				setState(92);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(93);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeReferenceContext extends ParserRuleContext {
		public List<NameContext> name() {
			return getRuleContexts(NameContext.class);
		}
		public NameContext name(int i) {
			return getRuleContext(NameContext.class,i);
		}
		public List<TerminalNode> RESOLVE() { return getTokens(kslParser.RESOLVE); }
		public TerminalNode RESOLVE(int i) {
			return getToken(kslParser.RESOLVE, i);
		}
		public TypeReferenceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeReference; }
	}

	public final TypeReferenceContext typeReference() throws RecognitionException {
		TypeReferenceContext _localctx = new TypeReferenceContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_typeReference);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(95);
			name();
			setState(100);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==RESOLVE) {
				{
				{
				setState(96);
				match(RESOLVE);
				setState(97);
				name();
				}
				}
				setState(102);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExtensionParamContext extends ParserRuleContext {
		public Token value;
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public TerminalNode EXPAND() { return getToken(kslParser.EXPAND, 0); }
		public List<TerminalNode> STRING_DELIM() { return getTokens(kslParser.STRING_DELIM); }
		public TerminalNode STRING_DELIM(int i) {
			return getToken(kslParser.STRING_DELIM, i);
		}
		public ExtensionParamContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_extensionParam; }
	}

	public final ExtensionParamContext extensionParam() throws RecognitionException {
		ExtensionParamContext _localctx = new ExtensionParamContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_extensionParam);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(103);
			name();
			setState(104);
			match(EXPAND);
			setState(105);
			match(STRING_DELIM);
			setState(106);
			((ExtensionParamContext)_localctx).value = _input.LT(1);
			_la = _input.LA(1);
			if ( _la <= 0 || (_la==STRING_DELIM) ) {
				((ExtensionParamContext)_localctx).value = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(107);
			match(STRING_DELIM);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExtensionParamsContext extends ParserRuleContext {
		public List<ExtensionParamContext> extensionParam() {
			return getRuleContexts(ExtensionParamContext.class);
		}
		public ExtensionParamContext extensionParam(int i) {
			return getRuleContext(ExtensionParamContext.class,i);
		}
		public List<TerminalNode> ARG_DELIM() { return getTokens(kslParser.ARG_DELIM); }
		public TerminalNode ARG_DELIM(int i) {
			return getToken(kslParser.ARG_DELIM, i);
		}
		public ExtensionParamsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_extensionParams; }
	}

	public final ExtensionParamsContext extensionParams() throws RecognitionException {
		ExtensionParamsContext _localctx = new ExtensionParamsContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_extensionParams);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(109);
			extensionParam();
			setState(114);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==ARG_DELIM) {
				{
				{
				setState(110);
				match(ARG_DELIM);
				setState(111);
				extensionParam();
				}
				}
				setState(116);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExtensionReferenceContext extends ParserRuleContext {
		public TerminalNode EXTENSION_CALL() { return getToken(kslParser.EXTENSION_CALL, 0); }
		public TypeReferenceContext typeReference() {
			return getRuleContext(TypeReferenceContext.class,0);
		}
		public TerminalNode LPAREN() { return getToken(kslParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(kslParser.RPAREN, 0); }
		public ExtensionParamsContext extensionParams() {
			return getRuleContext(ExtensionParamsContext.class,0);
		}
		public ExtensionReferenceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_extensionReference; }
	}

	public final ExtensionReferenceContext extensionReference() throws RecognitionException {
		ExtensionReferenceContext _localctx = new ExtensionReferenceContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_extensionReference);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(117);
			match(EXTENSION_CALL);
			setState(118);
			typeReference();
			setState(119);
			match(LPAREN);
			setState(121);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==FORCED_NAME || _la==NAME) {
				{
				setState(120);
				extensionParams();
				}
			}

			setState(123);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class RelationContext extends ParserRuleContext {
		public TerminalNode RELATION() { return getToken(kslParser.RELATION, 0); }
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public TerminalNode EXPAND() { return getToken(kslParser.EXPAND, 0); }
		public RelationBodyContext relationBody() {
			return getRuleContext(RelationBodyContext.class,0);
		}
		public List<ExtensionReferenceContext> extensionReference() {
			return getRuleContexts(ExtensionReferenceContext.class);
		}
		public ExtensionReferenceContext extensionReference(int i) {
			return getRuleContext(ExtensionReferenceContext.class,i);
		}
		public TerminalNode ACCESS() { return getToken(kslParser.ACCESS, 0); }
		public RelationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_relation; }
	}

	public final RelationContext relation() throws RecognitionException {
		RelationContext _localctx = new RelationContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_relation);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(128);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EXTENSION_CALL) {
				{
				{
				setState(125);
				extensionReference();
				}
				}
				setState(130);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(132);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ACCESS) {
				{
				setState(131);
				match(ACCESS);
				}
			}

			setState(134);
			match(RELATION);
			setState(135);
			name();
			setState(136);
			match(EXPAND);
			setState(137);
			relationBody(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class RelationBodyContext extends ParserRuleContext {
		public RelationBodyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_relationBody; }
	 
		public RelationBodyContext() { }
		public void copyFrom(RelationBodyContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class ORContext extends RelationBodyContext {
		public List<RelationBodyContext> relationBody() {
			return getRuleContexts(RelationBodyContext.class);
		}
		public RelationBodyContext relationBody(int i) {
			return getRuleContext(RelationBodyContext.class,i);
		}
		public TerminalNode OR() { return getToken(kslParser.OR, 0); }
		public ORContext(RelationBodyContext ctx) { copyFrom(ctx); }
	}
	public static class ReferenceContext extends RelationBodyContext {
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public ReferenceContext(RelationBodyContext ctx) { copyFrom(ctx); }
	}
	public static class AndContext extends RelationBodyContext {
		public List<RelationBodyContext> relationBody() {
			return getRuleContexts(RelationBodyContext.class);
		}
		public RelationBodyContext relationBody(int i) {
			return getRuleContext(RelationBodyContext.class,i);
		}
		public TerminalNode AND() { return getToken(kslParser.AND, 0); }
		public AndContext(RelationBodyContext ctx) { copyFrom(ctx); }
	}
	public static class SelfContext extends RelationBodyContext {
		public TerminalNode LSQUARE() { return getToken(kslParser.LSQUARE, 0); }
		public List<TypeReferenceContext> typeReference() {
			return getRuleContexts(TypeReferenceContext.class);
		}
		public TypeReferenceContext typeReference(int i) {
			return getRuleContext(TypeReferenceContext.class,i);
		}
		public TerminalNode RSQARE() { return getToken(kslParser.RSQARE, 0); }
		public TerminalNode CARDINALITY() { return getToken(kslParser.CARDINALITY, 0); }
		public List<TerminalNode> OR() { return getTokens(kslParser.OR); }
		public TerminalNode OR(int i) {
			return getToken(kslParser.OR, i);
		}
		public SelfContext(RelationBodyContext ctx) { copyFrom(ctx); }
	}
	public static class SubRelationContext extends RelationBodyContext {
		public NameContext relationName;
		public NameContext subrelationName;
		public TerminalNode RESOLVE() { return getToken(kslParser.RESOLVE, 0); }
		public List<NameContext> name() {
			return getRuleContexts(NameContext.class);
		}
		public NameContext name(int i) {
			return getRuleContext(NameContext.class,i);
		}
		public SubRelationContext(RelationBodyContext ctx) { copyFrom(ctx); }
	}
	public static class UnlessContext extends RelationBodyContext {
		public List<RelationBodyContext> relationBody() {
			return getRuleContexts(RelationBodyContext.class);
		}
		public RelationBodyContext relationBody(int i) {
			return getRuleContext(RelationBodyContext.class,i);
		}
		public TerminalNode UNLESS() { return getToken(kslParser.UNLESS, 0); }
		public UnlessContext(RelationBodyContext ctx) { copyFrom(ctx); }
	}
	public static class ParenContext extends RelationBodyContext {
		public TerminalNode LPAREN() { return getToken(kslParser.LPAREN, 0); }
		public RelationBodyContext relationBody() {
			return getRuleContext(RelationBodyContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(kslParser.RPAREN, 0); }
		public ParenContext(RelationBodyContext ctx) { copyFrom(ctx); }
	}

	public final RelationBodyContext relationBody() throws RecognitionException {
		return relationBody(0);
	}

	private RelationBodyContext relationBody(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		RelationBodyContext _localctx = new RelationBodyContext(_ctx, _parentState);
		RelationBodyContext _prevctx = _localctx;
		int _startState = 24;
		enterRecursionRule(_localctx, 24, RULE_relationBody, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(163);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				{
				_localctx = new SelfContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(140);
				match(LSQUARE);
				setState(142);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==CARDINALITY) {
					{
					setState(141);
					match(CARDINALITY);
					}
				}

				setState(144);
				typeReference();
				setState(149);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==OR) {
					{
					{
					setState(145);
					match(OR);
					setState(146);
					typeReference();
					}
					}
					setState(151);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(152);
				match(RSQARE);
				}
				break;
			case 2:
				{
				_localctx = new ReferenceContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(154);
				name();
				}
				break;
			case 3:
				{
				_localctx = new SubRelationContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(155);
				((SubRelationContext)_localctx).relationName = name();
				setState(156);
				match(RESOLVE);
				setState(157);
				((SubRelationContext)_localctx).subrelationName = name();
				}
				break;
			case 4:
				{
				_localctx = new ParenContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(159);
				match(LPAREN);
				setState(160);
				relationBody(0);
				setState(161);
				match(RPAREN);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(176);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(174);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
					case 1:
						{
						_localctx = new AndContext(new RelationBodyContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_relationBody);
						setState(165);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(166);
						match(AND);
						setState(167);
						relationBody(4);
						}
						break;
					case 2:
						{
						_localctx = new ORContext(new RelationBodyContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_relationBody);
						setState(168);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(169);
						match(OR);
						setState(170);
						relationBody(3);
						}
						break;
					case 3:
						{
						_localctx = new UnlessContext(new RelationBodyContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_relationBody);
						setState(171);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(172);
						match(UNLESS);
						setState(173);
						relationBody(2);
						}
						break;
					}
					} 
				}
				setState(178);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class ParamNamesContext extends ParserRuleContext {
		public List<NameContext> name() {
			return getRuleContexts(NameContext.class);
		}
		public NameContext name(int i) {
			return getRuleContext(NameContext.class,i);
		}
		public List<TerminalNode> ARG_DELIM() { return getTokens(kslParser.ARG_DELIM); }
		public TerminalNode ARG_DELIM(int i) {
			return getToken(kslParser.ARG_DELIM, i);
		}
		public ParamNamesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_paramNames; }
	}

	public final ParamNamesContext paramNames() throws RecognitionException {
		ParamNamesContext _localctx = new ParamNamesContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_paramNames);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(179);
			name();
			setState(184);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==ARG_DELIM) {
				{
				{
				setState(180);
				match(ARG_DELIM);
				setState(181);
				name();
				}
				}
				setState(186);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExtensionContext extends ParserRuleContext {
		public TerminalNode EXTENSION() { return getToken(kslParser.EXTENSION, 0); }
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public TerminalNode LPAREN() { return getToken(kslParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(kslParser.RPAREN, 0); }
		public TerminalNode LBRACE() { return getToken(kslParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(kslParser.RBRACE, 0); }
		public TerminalNode ACCESS() { return getToken(kslParser.ACCESS, 0); }
		public ParamNamesContext paramNames() {
			return getRuleContext(ParamNamesContext.class,0);
		}
		public List<DynamicTypeContext> dynamicType() {
			return getRuleContexts(DynamicTypeContext.class);
		}
		public DynamicTypeContext dynamicType(int i) {
			return getRuleContext(DynamicTypeContext.class,i);
		}
		public ExtensionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_extension; }
	}

	public final ExtensionContext extension() throws RecognitionException {
		ExtensionContext _localctx = new ExtensionContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_extension);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(188);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ACCESS) {
				{
				setState(187);
				match(ACCESS);
				}
			}

			setState(190);
			match(EXTENSION);
			setState(191);
			name();
			setState(192);
			match(LPAREN);
			setState(194);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==FORCED_NAME || _la==NAME) {
				{
				setState(193);
				paramNames();
				}
			}

			setState(196);
			match(RPAREN);
			setState(197);
			match(LBRACE);
			setState(199); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(198);
				dynamicType();
				}
				}
				setState(201); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==ACCESS || _la==TYPE );
			setState(203);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DynamicTypeContext extends ParserRuleContext {
		public TerminalNode TYPE() { return getToken(kslParser.TYPE, 0); }
		public DynamicNameContext dynamicName() {
			return getRuleContext(DynamicNameContext.class,0);
		}
		public TerminalNode LBRACE() { return getToken(kslParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(kslParser.RBRACE, 0); }
		public TerminalNode ACCESS() { return getToken(kslParser.ACCESS, 0); }
		public List<DynamicRelationContext> dynamicRelation() {
			return getRuleContexts(DynamicRelationContext.class);
		}
		public DynamicRelationContext dynamicRelation(int i) {
			return getRuleContext(DynamicRelationContext.class,i);
		}
		public DynamicTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dynamicType; }
	}

	public final DynamicTypeContext dynamicType() throws RecognitionException {
		DynamicTypeContext _localctx = new DynamicTypeContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_dynamicType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(206);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ACCESS) {
				{
				setState(205);
				match(ACCESS);
				}
			}

			setState(208);
			match(TYPE);
			setState(209);
			dynamicName();
			setState(210);
			match(LBRACE);
			setState(214);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << ACCESS) | (1L << RELATION) | (1L << ALLOW_DUPLICATES))) != 0)) {
				{
				{
				setState(211);
				dynamicRelation();
				}
				}
				setState(216);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(217);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DynamicRelationContext extends ParserRuleContext {
		public TerminalNode RELATION() { return getToken(kslParser.RELATION, 0); }
		public DynamicNameContext dynamicName() {
			return getRuleContext(DynamicNameContext.class,0);
		}
		public TerminalNode EXPAND() { return getToken(kslParser.EXPAND, 0); }
		public DynamicBodyContext dynamicBody() {
			return getRuleContext(DynamicBodyContext.class,0);
		}
		public TerminalNode ALLOW_DUPLICATES() { return getToken(kslParser.ALLOW_DUPLICATES, 0); }
		public TerminalNode ACCESS() { return getToken(kslParser.ACCESS, 0); }
		public DynamicRelationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dynamicRelation; }
	}

	public final DynamicRelationContext dynamicRelation() throws RecognitionException {
		DynamicRelationContext _localctx = new DynamicRelationContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_dynamicRelation);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(220);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ALLOW_DUPLICATES) {
				{
				setState(219);
				match(ALLOW_DUPLICATES);
				}
			}

			setState(223);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ACCESS) {
				{
				setState(222);
				match(ACCESS);
				}
			}

			setState(225);
			match(RELATION);
			setState(226);
			dynamicName();
			setState(227);
			match(EXPAND);
			setState(228);
			dynamicBody(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DynamicNameContext extends ParserRuleContext {
		public DynamicNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dynamicName; }
	 
		public DynamicNameContext() { }
		public void copyFrom(DynamicNameContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class VariableContext extends DynamicNameContext {
		public TerminalNode VARREF() { return getToken(kslParser.VARREF, 0); }
		public TerminalNode LBRACE() { return getToken(kslParser.LBRACE, 0); }
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public TerminalNode RBRACE() { return getToken(kslParser.RBRACE, 0); }
		public VariableContext(DynamicNameContext ctx) { copyFrom(ctx); }
	}
	public static class LiteralContext extends DynamicNameContext {
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public LiteralContext(DynamicNameContext ctx) { copyFrom(ctx); }
	}
	public static class TemplateContext extends DynamicNameContext {
		public List<TerminalNode> TEMPLATE_DELIM() { return getTokens(kslParser.TEMPLATE_DELIM); }
		public TerminalNode TEMPLATE_DELIM(int i) {
			return getToken(kslParser.TEMPLATE_DELIM, i);
		}
		public List<DynamicNameContext> dynamicName() {
			return getRuleContexts(DynamicNameContext.class);
		}
		public DynamicNameContext dynamicName(int i) {
			return getRuleContext(DynamicNameContext.class,i);
		}
		public TemplateContext(DynamicNameContext ctx) { copyFrom(ctx); }
	}

	public final DynamicNameContext dynamicName() throws RecognitionException {
		DynamicNameContext _localctx = new DynamicNameContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_dynamicName);
		try {
			int _alt;
			setState(244);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case FORCED_NAME:
			case NAME:
				_localctx = new LiteralContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(230);
				name();
				}
				break;
			case VARREF:
				_localctx = new VariableContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(231);
				match(VARREF);
				setState(232);
				match(LBRACE);
				setState(233);
				name();
				setState(234);
				match(RBRACE);
				}
				break;
			case TEMPLATE_DELIM:
				_localctx = new TemplateContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(236);
				match(TEMPLATE_DELIM);
				setState(238); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(237);
						dynamicName();
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(240); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,26,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				setState(242);
				match(TEMPLATE_DELIM);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DynamicBodyContext extends ParserRuleContext {
		public DynamicBodyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dynamicBody; }
	 
		public DynamicBodyContext() { }
		public void copyFrom(DynamicBodyContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class DynamicUnlessContext extends DynamicBodyContext {
		public List<DynamicBodyContext> dynamicBody() {
			return getRuleContexts(DynamicBodyContext.class);
		}
		public DynamicBodyContext dynamicBody(int i) {
			return getRuleContext(DynamicBodyContext.class,i);
		}
		public TerminalNode UNLESS() { return getToken(kslParser.UNLESS, 0); }
		public DynamicUnlessContext(DynamicBodyContext ctx) { copyFrom(ctx); }
	}
	public static class DynamicORContext extends DynamicBodyContext {
		public List<DynamicBodyContext> dynamicBody() {
			return getRuleContexts(DynamicBodyContext.class);
		}
		public DynamicBodyContext dynamicBody(int i) {
			return getRuleContext(DynamicBodyContext.class,i);
		}
		public TerminalNode OR() { return getToken(kslParser.OR, 0); }
		public DynamicORContext(DynamicBodyContext ctx) { copyFrom(ctx); }
	}
	public static class DynamicParenContext extends DynamicBodyContext {
		public TerminalNode LPAREN() { return getToken(kslParser.LPAREN, 0); }
		public DynamicBodyContext dynamicBody() {
			return getRuleContext(DynamicBodyContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(kslParser.RPAREN, 0); }
		public DynamicParenContext(DynamicBodyContext ctx) { copyFrom(ctx); }
	}
	public static class DynamicSelfContext extends DynamicBodyContext {
		public TerminalNode LSQUARE() { return getToken(kslParser.LSQUARE, 0); }
		public List<TypeReferenceContext> typeReference() {
			return getRuleContexts(TypeReferenceContext.class);
		}
		public TypeReferenceContext typeReference(int i) {
			return getRuleContext(TypeReferenceContext.class,i);
		}
		public TerminalNode RSQARE() { return getToken(kslParser.RSQARE, 0); }
		public TerminalNode CARDINALITY() { return getToken(kslParser.CARDINALITY, 0); }
		public List<TerminalNode> OR() { return getTokens(kslParser.OR); }
		public TerminalNode OR(int i) {
			return getToken(kslParser.OR, i);
		}
		public DynamicSelfContext(DynamicBodyContext ctx) { copyFrom(ctx); }
	}
	public static class DynamicAndContext extends DynamicBodyContext {
		public List<DynamicBodyContext> dynamicBody() {
			return getRuleContexts(DynamicBodyContext.class);
		}
		public DynamicBodyContext dynamicBody(int i) {
			return getRuleContext(DynamicBodyContext.class,i);
		}
		public TerminalNode AND() { return getToken(kslParser.AND, 0); }
		public DynamicAndContext(DynamicBodyContext ctx) { copyFrom(ctx); }
	}
	public static class DynamicReferenceContext extends DynamicBodyContext {
		public DynamicNameContext dynamicName() {
			return getRuleContext(DynamicNameContext.class,0);
		}
		public DynamicReferenceContext(DynamicBodyContext ctx) { copyFrom(ctx); }
	}
	public static class DynamicSubRelationContext extends DynamicBodyContext {
		public List<DynamicNameContext> dynamicName() {
			return getRuleContexts(DynamicNameContext.class);
		}
		public DynamicNameContext dynamicName(int i) {
			return getRuleContext(DynamicNameContext.class,i);
		}
		public TerminalNode RESOLVE() { return getToken(kslParser.RESOLVE, 0); }
		public DynamicSubRelationContext(DynamicBodyContext ctx) { copyFrom(ctx); }
	}

	public final DynamicBodyContext dynamicBody() throws RecognitionException {
		return dynamicBody(0);
	}

	private DynamicBodyContext dynamicBody(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		DynamicBodyContext _localctx = new DynamicBodyContext(_ctx, _parentState);
		DynamicBodyContext _prevctx = _localctx;
		int _startState = 36;
		enterRecursionRule(_localctx, 36, RULE_dynamicBody, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(270);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,30,_ctx) ) {
			case 1:
				{
				_localctx = new DynamicSelfContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(247);
				match(LSQUARE);
				setState(249);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==CARDINALITY) {
					{
					setState(248);
					match(CARDINALITY);
					}
				}

				setState(251);
				typeReference();
				setState(256);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==OR) {
					{
					{
					setState(252);
					match(OR);
					setState(253);
					typeReference();
					}
					}
					setState(258);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(259);
				match(RSQARE);
				}
				break;
			case 2:
				{
				_localctx = new DynamicReferenceContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(261);
				dynamicName();
				}
				break;
			case 3:
				{
				_localctx = new DynamicSubRelationContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(262);
				dynamicName();
				setState(263);
				match(RESOLVE);
				setState(264);
				dynamicName();
				}
				break;
			case 4:
				{
				_localctx = new DynamicParenContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(266);
				match(LPAREN);
				setState(267);
				dynamicBody(0);
				setState(268);
				match(RPAREN);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(283);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,32,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(281);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,31,_ctx) ) {
					case 1:
						{
						_localctx = new DynamicAndContext(new DynamicBodyContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_dynamicBody);
						setState(272);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(273);
						match(AND);
						setState(274);
						dynamicBody(4);
						}
						break;
					case 2:
						{
						_localctx = new DynamicORContext(new DynamicBodyContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_dynamicBody);
						setState(275);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(276);
						match(OR);
						setState(277);
						dynamicBody(3);
						}
						break;
					case 3:
						{
						_localctx = new DynamicUnlessContext(new DynamicBodyContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_dynamicBody);
						setState(278);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(279);
						match(UNLESS);
						setState(280);
						dynamicBody(2);
						}
						break;
					}
					} 
				}
				setState(285);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,32,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 12:
			return relationBody_sempred((RelationBodyContext)_localctx, predIndex);
		case 18:
			return dynamicBody_sempred((DynamicBodyContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean relationBody_sempred(RelationBodyContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 3);
		case 1:
			return precpred(_ctx, 2);
		case 2:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean dynamicBody_sempred(DynamicBodyContext _localctx, int predIndex) {
		switch (predIndex) {
		case 3:
			return precpred(_ctx, 3);
		case 4:
			return precpred(_ctx, 2);
		case 5:
			return precpred(_ctx, 1);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3*\u0121\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\3\2\3\2\3\3\3\3\3\3\7\3.\n\3\f\3\16\3\61\13\3\3\3"+
		"\6\3\64\n\3\r\3\16\3\65\3\4\3\4\3\4\3\5\3\5\3\5\3\6\3\6\3\6\3\7\3\7\5"+
		"\7C\n\7\3\7\3\7\5\7G\n\7\3\7\3\7\3\7\5\7L\n\7\3\b\7\bO\n\b\f\b\16\bR\13"+
		"\b\3\b\5\bU\n\b\3\b\3\b\3\b\3\b\7\b[\n\b\f\b\16\b^\13\b\3\b\3\b\3\t\3"+
		"\t\3\t\7\te\n\t\f\t\16\th\13\t\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13"+
		"\7\13s\n\13\f\13\16\13v\13\13\3\f\3\f\3\f\3\f\5\f|\n\f\3\f\3\f\3\r\7\r"+
		"\u0081\n\r\f\r\16\r\u0084\13\r\3\r\5\r\u0087\n\r\3\r\3\r\3\r\3\r\3\r\3"+
		"\16\3\16\3\16\5\16\u0091\n\16\3\16\3\16\3\16\7\16\u0096\n\16\f\16\16\16"+
		"\u0099\13\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\5"+
		"\16\u00a6\n\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\7\16\u00b1"+
		"\n\16\f\16\16\16\u00b4\13\16\3\17\3\17\3\17\7\17\u00b9\n\17\f\17\16\17"+
		"\u00bc\13\17\3\20\5\20\u00bf\n\20\3\20\3\20\3\20\3\20\5\20\u00c5\n\20"+
		"\3\20\3\20\3\20\6\20\u00ca\n\20\r\20\16\20\u00cb\3\20\3\20\3\21\5\21\u00d1"+
		"\n\21\3\21\3\21\3\21\3\21\7\21\u00d7\n\21\f\21\16\21\u00da\13\21\3\21"+
		"\3\21\3\22\5\22\u00df\n\22\3\22\5\22\u00e2\n\22\3\22\3\22\3\22\3\22\3"+
		"\22\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\6\23\u00f1\n\23\r\23\16\23"+
		"\u00f2\3\23\3\23\5\23\u00f7\n\23\3\24\3\24\3\24\5\24\u00fc\n\24\3\24\3"+
		"\24\3\24\7\24\u0101\n\24\f\24\16\24\u0104\13\24\3\24\3\24\3\24\3\24\3"+
		"\24\3\24\3\24\3\24\3\24\3\24\3\24\5\24\u0111\n\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\7\24\u011c\n\24\f\24\16\24\u011f\13\24\3\24"+
		"\2\4\32&\25\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&\2\4\4\2\4\4(("+
		"\3\2%%\2\u0136\2(\3\2\2\2\4*\3\2\2\2\6\67\3\2\2\2\b:\3\2\2\2\n=\3\2\2"+
		"\2\fK\3\2\2\2\16P\3\2\2\2\20a\3\2\2\2\22i\3\2\2\2\24o\3\2\2\2\26w\3\2"+
		"\2\2\30\u0082\3\2\2\2\32\u00a5\3\2\2\2\34\u00b5\3\2\2\2\36\u00be\3\2\2"+
		"\2 \u00d0\3\2\2\2\"\u00de\3\2\2\2$\u00f6\3\2\2\2&\u0110\3\2\2\2()\t\2"+
		"\2\2)\3\3\2\2\2*+\5\6\4\2+/\5\b\5\2,.\5\n\6\2-,\3\2\2\2.\61\3\2\2\2/-"+
		"\3\2\2\2/\60\3\2\2\2\60\63\3\2\2\2\61/\3\2\2\2\62\64\5\f\7\2\63\62\3\2"+
		"\2\2\64\65\3\2\2\2\65\63\3\2\2\2\65\66\3\2\2\2\66\5\3\2\2\2\678\7\5\2"+
		"\289\7\6\2\29\7\3\2\2\2:;\7\b\2\2;<\5\2\2\2<\t\3\2\2\2=>\7\17\2\2>?\5"+
		"\2\2\2?\13\3\2\2\2@B\5\16\b\2AC\7\'\2\2BA\3\2\2\2BC\3\2\2\2CL\3\2\2\2"+
		"DF\5\36\20\2EG\7\'\2\2FE\3\2\2\2FG\3\2\2\2GL\3\2\2\2HI\5\26\f\2IJ\7\'"+
		"\2\2JL\3\2\2\2K@\3\2\2\2KD\3\2\2\2KH\3\2\2\2L\r\3\2\2\2MO\5\26\f\2NM\3"+
		"\2\2\2OR\3\2\2\2PN\3\2\2\2PQ\3\2\2\2QT\3\2\2\2RP\3\2\2\2SU\7\t\2\2TS\3"+
		"\2\2\2TU\3\2\2\2UV\3\2\2\2VW\7\r\2\2WX\5\2\2\2X\\\7\34\2\2Y[\5\30\r\2"+
		"ZY\3\2\2\2[^\3\2\2\2\\Z\3\2\2\2\\]\3\2\2\2]_\3\2\2\2^\\\3\2\2\2_`\7\35"+
		"\2\2`\17\3\2\2\2af\5\2\2\2bc\7\7\2\2ce\5\2\2\2db\3\2\2\2eh\3\2\2\2fd\3"+
		"\2\2\2fg\3\2\2\2g\21\3\2\2\2hf\3\2\2\2ij\5\2\2\2jk\7\33\2\2kl\7%\2\2l"+
		"m\n\3\2\2mn\7%\2\2n\23\3\2\2\2ot\5\22\n\2pq\7&\2\2qs\5\22\n\2rp\3\2\2"+
		"\2sv\3\2\2\2tr\3\2\2\2tu\3\2\2\2u\25\3\2\2\2vt\3\2\2\2wx\7\36\2\2xy\5"+
		"\20\t\2y{\7\37\2\2z|\5\24\13\2{z\3\2\2\2{|\3\2\2\2|}\3\2\2\2}~\7 \2\2"+
		"~\27\3\2\2\2\177\u0081\5\26\f\2\u0080\177\3\2\2\2\u0081\u0084\3\2\2\2"+
		"\u0082\u0080\3\2\2\2\u0082\u0083\3\2\2\2\u0083\u0086\3\2\2\2\u0084\u0082"+
		"\3\2\2\2\u0085\u0087\7\t\2\2\u0086\u0085\3\2\2\2\u0086\u0087\3\2\2\2\u0087"+
		"\u0088\3\2\2\2\u0088\u0089\7\16\2\2\u0089\u008a\5\2\2\2\u008a\u008b\7"+
		"\33\2\2\u008b\u008c\5\32\16\2\u008c\31\3\2\2\2\u008d\u008e\b\16\1\2\u008e"+
		"\u0090\7!\2\2\u008f\u0091\7\21\2\2\u0090\u008f\3\2\2\2\u0090\u0091\3\2"+
		"\2\2\u0091\u0092\3\2\2\2\u0092\u0097\5\20\t\2\u0093\u0094\7\30\2\2\u0094"+
		"\u0096\5\20\t\2\u0095\u0093\3\2\2\2\u0096\u0099\3\2\2\2\u0097\u0095\3"+
		"\2\2\2\u0097\u0098\3\2\2\2\u0098\u009a\3\2\2\2\u0099\u0097\3\2\2\2\u009a"+
		"\u009b\7\"\2\2\u009b\u00a6\3\2\2\2\u009c\u00a6\5\2\2\2\u009d\u009e\5\2"+
		"\2\2\u009e\u009f\7\7\2\2\u009f\u00a0\5\2\2\2\u00a0\u00a6\3\2\2\2\u00a1"+
		"\u00a2\7\37\2\2\u00a2\u00a3\5\32\16\2\u00a3\u00a4\7 \2\2\u00a4\u00a6\3"+
		"\2\2\2\u00a5\u008d\3\2\2\2\u00a5\u009c\3\2\2\2\u00a5\u009d\3\2\2\2\u00a5"+
		"\u00a1\3\2\2\2\u00a6\u00b2\3\2\2\2\u00a7\u00a8\f\5\2\2\u00a8\u00a9\7\27"+
		"\2\2\u00a9\u00b1\5\32\16\6\u00aa\u00ab\f\4\2\2\u00ab\u00ac\7\30\2\2\u00ac"+
		"\u00b1\5\32\16\5\u00ad\u00ae\f\3\2\2\u00ae\u00af\7\31\2\2\u00af\u00b1"+
		"\5\32\16\4\u00b0\u00a7\3\2\2\2\u00b0\u00aa\3\2\2\2\u00b0\u00ad\3\2\2\2"+
		"\u00b1\u00b4\3\2\2\2\u00b2\u00b0\3\2\2\2\u00b2\u00b3\3\2\2\2\u00b3\33"+
		"\3\2\2\2\u00b4\u00b2\3\2\2\2\u00b5\u00ba\5\2\2\2\u00b6\u00b7\7&\2\2\u00b7"+
		"\u00b9\5\2\2\2\u00b8\u00b6\3\2\2\2\u00b9\u00bc\3\2\2\2\u00ba\u00b8\3\2"+
		"\2\2\u00ba\u00bb\3\2\2\2\u00bb\35\3\2\2\2\u00bc\u00ba\3\2\2\2\u00bd\u00bf"+
		"\7\t\2\2\u00be\u00bd\3\2\2\2\u00be\u00bf\3\2\2\2\u00bf\u00c0\3\2\2\2\u00c0"+
		"\u00c1\7\20\2\2\u00c1\u00c2\5\2\2\2\u00c2\u00c4\7\37\2\2\u00c3\u00c5\5"+
		"\34\17\2\u00c4\u00c3\3\2\2\2\u00c4\u00c5\3\2\2\2\u00c5\u00c6\3\2\2\2\u00c6"+
		"\u00c7\7 \2\2\u00c7\u00c9\7\34\2\2\u00c8\u00ca\5 \21\2\u00c9\u00c8\3\2"+
		"\2\2\u00ca\u00cb\3\2\2\2\u00cb\u00c9\3\2\2\2\u00cb\u00cc\3\2\2\2\u00cc"+
		"\u00cd\3\2\2\2\u00cd\u00ce\7\35\2\2\u00ce\37\3\2\2\2\u00cf\u00d1\7\t\2"+
		"\2\u00d0\u00cf\3\2\2\2\u00d0\u00d1\3\2\2\2\u00d1\u00d2\3\2\2\2\u00d2\u00d3"+
		"\7\r\2\2\u00d3\u00d4\5$\23\2\u00d4\u00d8\7\34\2\2\u00d5\u00d7\5\"\22\2"+
		"\u00d6\u00d5\3\2\2\2\u00d7\u00da\3\2\2\2\u00d8\u00d6\3\2\2\2\u00d8\u00d9"+
		"\3\2\2\2\u00d9\u00db\3\2\2\2\u00da\u00d8\3\2\2\2\u00db\u00dc\7\35\2\2"+
		"\u00dc!\3\2\2\2\u00dd\u00df\7\32\2\2\u00de\u00dd\3\2\2\2\u00de\u00df\3"+
		"\2\2\2\u00df\u00e1\3\2\2\2\u00e0\u00e2\7\t\2\2\u00e1\u00e0\3\2\2\2\u00e1"+
		"\u00e2\3\2\2\2\u00e2\u00e3\3\2\2\2\u00e3\u00e4\7\16\2\2\u00e4\u00e5\5"+
		"$\23\2\u00e5\u00e6\7\33\2\2\u00e6\u00e7\5&\24\2\u00e7#\3\2\2\2\u00e8\u00f7"+
		"\5\2\2\2\u00e9\u00ea\7#\2\2\u00ea\u00eb\7\34\2\2\u00eb\u00ec\5\2\2\2\u00ec"+
		"\u00ed\7\35\2\2\u00ed\u00f7\3\2\2\2\u00ee\u00f0\7$\2\2\u00ef\u00f1\5$"+
		"\23\2\u00f0\u00ef\3\2\2\2\u00f1\u00f2\3\2\2\2\u00f2\u00f0\3\2\2\2\u00f2"+
		"\u00f3\3\2\2\2\u00f3\u00f4\3\2\2\2\u00f4\u00f5\7$\2\2\u00f5\u00f7\3\2"+
		"\2\2\u00f6\u00e8\3\2\2\2\u00f6\u00e9\3\2\2\2\u00f6\u00ee\3\2\2\2\u00f7"+
		"%\3\2\2\2\u00f8\u00f9\b\24\1\2\u00f9\u00fb\7!\2\2\u00fa\u00fc\7\21\2\2"+
		"\u00fb\u00fa\3\2\2\2\u00fb\u00fc\3\2\2\2\u00fc\u00fd\3\2\2\2\u00fd\u0102"+
		"\5\20\t\2\u00fe\u00ff\7\30\2\2\u00ff\u0101\5\20\t\2\u0100\u00fe\3\2\2"+
		"\2\u0101\u0104\3\2\2\2\u0102\u0100\3\2\2\2\u0102\u0103\3\2\2\2\u0103\u0105"+
		"\3\2\2\2\u0104\u0102\3\2\2\2\u0105\u0106\7\"\2\2\u0106\u0111\3\2\2\2\u0107"+
		"\u0111\5$\23\2\u0108\u0109\5$\23\2\u0109\u010a\7\7\2\2\u010a\u010b\5$"+
		"\23\2\u010b\u0111\3\2\2\2\u010c\u010d\7\37\2\2\u010d\u010e\5&\24\2\u010e"+
		"\u010f\7 \2\2\u010f\u0111\3\2\2\2\u0110\u00f8\3\2\2\2\u0110\u0107\3\2"+
		"\2\2\u0110\u0108\3\2\2\2\u0110\u010c\3\2\2\2\u0111\u011d\3\2\2\2\u0112"+
		"\u0113\f\5\2\2\u0113\u0114\7\27\2\2\u0114\u011c\5&\24\6\u0115\u0116\f"+
		"\4\2\2\u0116\u0117\7\30\2\2\u0117\u011c\5&\24\5\u0118\u0119\f\3\2\2\u0119"+
		"\u011a\7\31\2\2\u011a\u011c\5&\24\4\u011b\u0112\3\2\2\2\u011b\u0115\3"+
		"\2\2\2\u011b\u0118\3\2\2\2\u011c\u011f\3\2\2\2\u011d\u011b\3\2\2\2\u011d"+
		"\u011e\3\2\2\2\u011e\'\3\2\2\2\u011f\u011d\3\2\2\2#/\65BFKPT\\ft{\u0082"+
		"\u0086\u0090\u0097\u00a5\u00b0\u00b2\u00ba\u00be\u00c4\u00cb\u00d0\u00d8"+
		"\u00de\u00e1\u00f2\u00f6\u00fb\u0102\u0110\u011b\u011d";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}