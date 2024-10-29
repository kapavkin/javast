package javast

// Implements [AnnotatedTypeNode].
type AnnotatedType struct {
	Annotations    []AnnotationNode
	UnderlyingType ExpressionNode
}

func (AnnotatedType) GetKind() Kind { return ANNOTATED_TYPE }

func (at AnnotatedType) GetAnnotations() []AnnotationNode  { return at.Annotations }
func (at AnnotatedType) GetUnderlyingType() ExpressionNode { return at.UnderlyingType }

func (AnnotatedType) caseLabelNode()     {}
func (AnnotatedType) expressionNode()    {}
func (AnnotatedType) annotatedTypeNode() {}

// Implements [AnnotationNode] of kind [ANNOTATION].
type Annotation struct {
	AnnotationType Node
	Arguments      []ExpressionNode
}

func (Annotation) GetKind() Kind { return ANNOTATION }

func (a Annotation) GetAnnotationType() Node        { return a.AnnotationType }
func (a Annotation) GetArguments() []ExpressionNode { return a.Arguments }

func (Annotation) caseLabelNode()  {}
func (Annotation) expressionNode() {}
func (Annotation) annotationNode() {}

// Implements [AnnotationNode] of kind [TYPE_ANNOTATION].
type TypeAnnotation struct {
	AnnotationType Node
	Arguments      []ExpressionNode
}

func (TypeAnnotation) GetKind() Kind { return TYPE_ANNOTATION }

func (ta TypeAnnotation) GetAnnotationType() Node        { return ta.AnnotationType }
func (ta TypeAnnotation) GetArguments() []ExpressionNode { return ta.Arguments }

func (TypeAnnotation) caseLabelNode()  {}
func (TypeAnnotation) expressionNode() {}
func (TypeAnnotation) annotationNode() {}

// Implements [ArrayAccessNode].
type ArrayAccess struct {
	Expression ExpressionNode
	Index      ExpressionNode
}

func (ArrayAccess) GetKind() Kind { return ARRAY_ACCESS }

func (aa ArrayAccess) GetExpression() ExpressionNode { return aa.Expression }
func (aa ArrayAccess) GetIndex() ExpressionNode      { return aa.Index }

func (ArrayAccess) caseLabelNode()   {}
func (ArrayAccess) expressionNode()  {}
func (ArrayAccess) arrayAccessNode() {}

// Implements [ArrayTypeNode].
type ArrayType struct {
	Type Node
}

func (ArrayType) GetKind() Kind { return ARRAY_TYPE }

func (at ArrayType) GetType() Node { return at.Type }

func (ArrayType) arrayTypeNode() {}

// Implements [AssertNode].
type Assert struct {
	Condition ExpressionNode
	Detail    ExpressionNode
}

func (Assert) GetKind() Kind { return ASSERT }

func (a Assert) GetCondition() ExpressionNode { return a.Condition }
func (a Assert) GetDetail() ExpressionNode    { return a.Detail }

func (Assert) statementNode() {}
func (Assert) assertNode()    {}

// Implements [AssignmentNode].
type Assignment struct {
	Variable   ExpressionNode
	Expression ExpressionNode
}

func (Assignment) GetKind() Kind { return ASSIGNMENT }

func (a Assignment) GetVariable() ExpressionNode   { return a.Variable }
func (a Assignment) GetExpression() ExpressionNode { return a.Expression }

func (Assignment) caseLabelNode()  {}
func (Assignment) expressionNode() {}
func (Assignment) assignmentNode() {}

// Implements [BlockNode].
type Block struct {
	Static     bool
	Statements []StatementNode
}

func (Block) GetKind() Kind { return BLOCK }

func (b Block) IsStatic() bool                 { return b.Static }
func (b Block) GetStatements() []StatementNode { return b.Statements }

func (Block) statementNode() {}
func (Block) blockNode()     {}

// Implements [BreakNode].
type Break struct {
	Label *string
}

func (Break) GetKind() Kind { return BREAK }

func (b Break) GetLabel() *string { return b.Label }

func (Break) statementNode() {}
func (Break) breakNode()     {}

// Implements [CaseNode] of kind [STATEMENT_CASE_KIND].
type StatementCase struct {
	Expression ExpressionNode
	Statements []StatementNode
}

func (StatementCase) GetKind() Kind { return CASE }

func (sc StatementCase) GetExpressions() []ExpressionNode {
	if sc.Expression != nil {
		return []ExpressionNode{sc.Expression}
	}
	return []ExpressionNode{}
}

func (sc StatementCase) GetLabels() []CaseLabelNode     { return nil }
func (sc StatementCase) GetStatements() []StatementNode { return sc.Statements }
func (StatementCase) GetBody() Node                     { return nil }
func (StatementCase) GetCaseKind() CaseKind             { return STATEMENT_CASE_KIND }

func (StatementCase) caseNode() {}

// Implements [CaseNode] of kind [RULE_CASE_KIND].
type RuleCase struct {
	Labels []CaseLabelNode
	Body   Node
}

func (RuleCase) GetKind() Kind { return CASE }

func (rc RuleCase) GetExpressions() []ExpressionNode { return nil }
func (rc RuleCase) GetLabels() []CaseLabelNode       { return rc.Labels }
func (rc RuleCase) GetStatements() []StatementNode   { return nil }
func (rc RuleCase) GetBody() Node                    { return rc.Body }
func (rc RuleCase) GetCaseKind() CaseKind            { return RULE_CASE_KIND }

func (RuleCase) caseNode() {}

// Implements [CatchNode].
type Catch struct {
	Parameter VariableNode
	Block     BlockNode
}

func (Catch) GetKind() Kind { return CATCH }

func (c Catch) GetParameter() VariableNode { return c.Parameter }
func (c Catch) GetBlock() BlockNode        { return c.Block }

func (Catch) catchNode() {}

// Implements [ClassNode].
type Class struct {
	Modifiers        ModifiersNode
	SimpleName       string
	TypeParameters   []TypeParameterNode
	ExtendsClause    Node
	ImplementsClause []Node
	Members          []Node
}

func (Class) GetKind() Kind { return CLASS }

func (c Class) GetModifiers() ModifiersNode            { return c.Modifiers }
func (c Class) GetSimpleName() string                  { return c.SimpleName }
func (c Class) GetTypeParameters() []TypeParameterNode { return c.TypeParameters }
func (c Class) GetExtendsClause() Node                 { return c.ExtendsClause }
func (c Class) GetImplementsClause() []Node            { return c.ImplementsClause }
func (c Class) GetPermitsClause() []Node               { return nil }
func (c Class) GetMembers() []Node                     { return c.Members }

func (Class) statementNode() {}
func (Class) classNode()     {}

// Implements [CompilationUnitNode].
type CompilationUnit struct {
	Module    ModuleNode
	Package   PackageNode
	Imports   []ImportNode
	TypeDecls []Node
}

func (CompilationUnit) GetKind() Kind { return COMPILATION_UNIT }

func (cu CompilationUnit) GetModule() ModuleNode    { return cu.Module }
func (cu CompilationUnit) GetPackage() PackageNode  { return cu.Package }
func (cu CompilationUnit) GetImports() []ImportNode { return cu.Imports }
func (cu CompilationUnit) GetTypeDecls() []Node     { return cu.TypeDecls }

func (CompilationUnit) compilationUnitNode() {}

// Implements [ConditionalExpressionNode].
type ConditionalExpression struct {
	Condition       ExpressionNode
	TrueExpression  ExpressionNode
	FalseExpression ExpressionNode
}

func (ConditionalExpression) GetKind() Kind { return CONDITIONAL_EXPRESSION }

func (cx ConditionalExpression) GetCondition() ExpressionNode       { return cx.Condition }
func (cx ConditionalExpression) GetTrueExpression() ExpressionNode  { return cx.TrueExpression }
func (cx ConditionalExpression) GetFalseExpression() ExpressionNode { return cx.FalseExpression }

func (ConditionalExpression) caseLabelNode()             {}
func (ConditionalExpression) expressionNode()            {}
func (ConditionalExpression) conditionalExpressionNode() {}

// Implements [ContinueNode].
type Continue struct {
	Label *string
}

func (Continue) GetKind() Kind { return CONTINUE }

func (c Continue) GetLabel() *string { return c.Label }

func (Continue) statementNode() {}
func (Continue) continueNode()  {}

// Implements [DoWhileLoopNode].
type DoWhileLoop struct {
	Condition ExpressionNode
	Statement StatementNode
}

func (DoWhileLoop) GetKind() Kind { return DO_WHILE_LOOP }

func (dwl DoWhileLoop) GetCondition() ExpressionNode { return dwl.Condition }
func (dwl DoWhileLoop) GetStatement() StatementNode  { return dwl.Statement }

func (DoWhileLoop) statementNode()   {}
func (DoWhileLoop) doWhileLoopNode() {}

// Implements [EnhancedForLoopNode].
type EnhancedForLoop struct {
	Variable   VariableNode
	Expression ExpressionNode
	Statement  StatementNode
}

func (EnhancedForLoop) GetKind() Kind { return ENHANCED_FOR_LOOP }

func (efl EnhancedForLoop) GetVariable() VariableNode     { return efl.Variable }
func (efl EnhancedForLoop) GetExpression() ExpressionNode { return efl.Expression }
func (efl EnhancedForLoop) GetStatement() StatementNode   { return efl.Statement }

func (EnhancedForLoop) statementNode()       {}
func (EnhancedForLoop) enhancedForLoopNode() {}

// Implements [ExpressionStatementNode].
type ExpressionStatement struct {
	Expression ExpressionNode
}

func (ExpressionStatement) GetKind() Kind { return EXPRESSION_STATEMENT }

func (xs ExpressionStatement) GetExpression() ExpressionNode { return xs.Expression }

func (ExpressionStatement) statementNode()           {}
func (ExpressionStatement) expressionStatementNode() {}

// Implements [MemberSelectNode].
type MemberSelect struct {
	Expression ExpressionNode
	Identifier string
}

func (MemberSelect) GetKind() Kind { return MEMBER_SELECT }

func (ms MemberSelect) GetExpression() ExpressionNode { return ms.Expression }
func (ms MemberSelect) GetIdentifier() string         { return ms.Identifier }

func (MemberSelect) caseLabelNode()    {}
func (MemberSelect) expressionNode()   {}
func (MemberSelect) memberSelectNode() {}

// Implements [MemberReferenceNode] of mode [INVOKE_REFERENCE_MODE].
type InvokeMemberReference struct {
	QualifierExpression ExpressionNode
	Name                string
	TypeArguments       []ExpressionNode
}

func (InvokeMemberReference) GetKind() Kind { return MEMBER_REFERENCE }

func (imr InvokeMemberReference) GetMode() ReferenceMode { return INVOKE_REFERENCE_MODE }

func (imr InvokeMemberReference) GetQualifierExpression() ExpressionNode {
	return imr.QualifierExpression

}
func (imr InvokeMemberReference) GetName() *string                   { return &imr.Name }
func (imr InvokeMemberReference) GetTypeArguments() []ExpressionNode { return imr.TypeArguments }

func (InvokeMemberReference) caseLabelNode()       {}
func (InvokeMemberReference) expressionNode()      {}
func (InvokeMemberReference) memberReferenceNode() {}

// Implements [MemberReferenceNode] of mode [NEW_REFERENCE_MODE].
type NewMemberReference struct {
	QualifierExpression ExpressionNode
	TypeArguments       []ExpressionNode
}

func (NewMemberReference) GetKind() Kind { return MEMBER_REFERENCE }

func (nmr NewMemberReference) GetMode() ReferenceMode { return NEW_REFERENCE_MODE }

func (nmr NewMemberReference) GetQualifierExpression() ExpressionNode {
	return nmr.QualifierExpression
}

func (nmr NewMemberReference) GetName() *string                   { return nil }
func (nmr NewMemberReference) GetTypeArguments() []ExpressionNode { return nmr.TypeArguments }

func (NewMemberReference) caseLabelNode()       {}
func (NewMemberReference) expressionNode()      {}
func (NewMemberReference) memberReferenceNode() {}

// Implements [ForLoopNode].
type ForLoop struct {
	Initializer []VariableNode
	Condition   ExpressionNode
	Update      []ExpressionNode
	Statement   StatementNode
}

func (ForLoop) GetKind() Kind { return FOR_LOOP }

func (fl ForLoop) GetInitializer() []VariableNode { return fl.Initializer }
func (fl ForLoop) GetCondition() ExpressionNode   { return fl.Condition }
func (fl ForLoop) GetUpdate() []ExpressionNode    { return fl.Update }
func (fl ForLoop) GetStatement() StatementNode    { return fl.Statement }

func (ForLoop) statementNode() {}
func (ForLoop) forLoopNode()   {}

// Implements [IdentifierNode].
type Identifier struct {
	Name string
}

func (Identifier) GetKind() Kind { return IDENTIFIER }

func (i Identifier) GetName() string { return i.Name }

func (Identifier) caseLabelNode()  {}
func (Identifier) expressionNode() {}
func (Identifier) identifierNode() {}

// Implements [IfNode].
type If struct {
	Condition     ExpressionNode
	ThenStatement StatementNode
	ElseStatement StatementNode
}

func (If) GetKind() Kind { return IF }

func (i If) GetCondition() ExpressionNode    { return i.Condition }
func (i If) GetThenStatement() StatementNode { return i.ThenStatement }
func (i If) GetElseStatement() StatementNode { return i.ElseStatement }

func (If) statementNode() {}
func (If) ifNode()        {}

// Implements [ImportNode].
type Import struct {
	Static              bool
	QualifiedIdentifier Node
}

func (Import) GetKind() Kind { return IMPORT }

func (i Import) IsStatic() bool               { return i.Static }
func (i Import) GetQualifiedIdentifier() Node { return i.QualifiedIdentifier }

func (Import) importNode() {}

// Implements [InstanceOfNode].
type InstanceOf struct {
	Expression ExpressionNode
	Type       Node
	Pattern    PatternNode
}

func (InstanceOf) GetKind() Kind { return INSTANCE_OF }

func (io InstanceOf) GetExpression() ExpressionNode { return io.Expression }
func (io InstanceOf) GetType() Node                 { return io.Type }
func (io InstanceOf) GetPattern() PatternNode       { return io.Pattern }

func (InstanceOf) caseLabelNode()  {}
func (InstanceOf) expressionNode() {}
func (InstanceOf) instanceOfNode() {}

// Implements [LabeledStatementNode].
type LabeledStatement struct {
	Label     string
	Statement StatementNode
}

func (LabeledStatement) GetKind() Kind { return LABELED_STATEMENT }

func (ls LabeledStatement) GetLabel() string            { return ls.Label }
func (ls LabeledStatement) GetStatement() StatementNode { return ls.Statement }

func (LabeledStatement) statementNode()        {}
func (LabeledStatement) labeledStatementNode() {}

// Implements [MethodNode].
type Method struct {
	Modifiers         ModifiersNode
	Name              string
	ReturnType        Node
	TypeParameters    []TypeParameterNode
	Parameters        []VariableNode
	ReceiverParameter VariableNode
	Throws            []ExpressionNode
	Body              BlockNode
	DefaultValue      Node
}

func (Method) GetKind() Kind { return METHOD }

func (m Method) GetModifiers() ModifiersNode            { return m.Modifiers }
func (m Method) GetName() string                        { return m.Name }
func (m Method) GetReturnType() Node                    { return m.ReturnType }
func (m Method) GetTypeParameters() []TypeParameterNode { return m.TypeParameters }
func (m Method) GetParameters() []VariableNode          { return m.Parameters }
func (m Method) GetReceiverParameter() VariableNode     { return m.ReceiverParameter }
func (m Method) GetThrows() []ExpressionNode            { return m.Throws }
func (m Method) GetBody() BlockNode                     { return m.Body }
func (m Method) GetDefaultValue() Node                  { return m.DefaultValue }

func (Method) methodNode() {}

// Implements [MethodInvocationNode].
type MethodInvocation struct {
	TypeArguments []Node
	MethodSelect  ExpressionNode
	Arguments     []ExpressionNode
}

func (MethodInvocation) GetKind() Kind { return METHOD_INVOCATION }

func (mi MethodInvocation) GetTypeArguments() []Node        { return mi.TypeArguments }
func (mi MethodInvocation) GetMethodSelect() ExpressionNode { return mi.MethodSelect }
func (mi MethodInvocation) GetArguments() []ExpressionNode  { return mi.Arguments }

func (MethodInvocation) caseLabelNode()        {}
func (MethodInvocation) expressionNode()       {}
func (MethodInvocation) methodInvocationNode() {}

// Implements [ModifiersNode].
type Modifiers struct {
	Flags       []Modifier
	Annotations []AnnotationNode
}

func (Modifiers) GetKind() Kind { return MODIFIERS }

func (m Modifiers) GetFlags() []Modifier { return m.Flags }

func (m Modifiers) GetAnnotations() []AnnotationNode { return m.Annotations }

func (Modifiers) modifiersNode() {}

// Implements [NewArrayNode].
type NewArray struct {
	Type           Node
	Dimensions     []ExpressionNode
	Initializers   []ExpressionNode
	Annotations    []AnnotationNode
	DimAnnotations [][]AnnotationNode
}

func (NewArray) GetKind() Kind { return NEW_ARRAY }

func (na NewArray) GetType() Node                         { return na.Type }
func (na NewArray) GetDimensions() []ExpressionNode       { return na.Dimensions }
func (na NewArray) GetInitializers() []ExpressionNode     { return na.Initializers }
func (na NewArray) GetAnnotations() []AnnotationNode      { return na.Annotations }
func (na NewArray) GetDimAnnotations() [][]AnnotationNode { return na.DimAnnotations }

func (NewArray) caseLabelNode()  {}
func (NewArray) expressionNode() {}
func (NewArray) newArrayNode()   {}

// Implements [NewClassNode].
type NewClass struct {
	EnclosingExpression ExpressionNode
	TypeArguments       []Node
	Identifier          ExpressionNode
	Arguments           []ExpressionNode
	ClassBody           ClassNode
}

func (NewClass) GetKind() Kind { return NEW_CLASS }

func (nc NewClass) GetEnclosingExpression() ExpressionNode { return nc.EnclosingExpression }
func (nc NewClass) GetTypeArguments() []Node               { return nc.TypeArguments }
func (nc NewClass) GetIdentifier() ExpressionNode          { return nc.Identifier }
func (nc NewClass) GetArguments() []ExpressionNode         { return nc.Arguments }
func (nc NewClass) GetClassBody() ClassNode                { return nc.ClassBody }

func (NewClass) caseLabelNode()  {}
func (NewClass) expressionNode() {}
func (NewClass) newClassNode()   {}

// Implements [LambdaExpressionNode] of kind [EXPRESSION_BODY_KIND].
type ExpressionLambdaExpression struct {
	Parameters []VariableNode
	Expression ExpressionNode
}

func (ExpressionLambdaExpression) GetKind() Kind { return LAMBDA_EXPRESSION }

func (xlx ExpressionLambdaExpression) GetParameters() []VariableNode { return xlx.Parameters }
func (xlx ExpressionLambdaExpression) GetBody() Node                 { return xlx.Expression }
func (xlx ExpressionLambdaExpression) GetBodyKind() BodyKind         { return EXPRESSION_BODY_KIND }

func (ExpressionLambdaExpression) caseLabelNode()        {}
func (ExpressionLambdaExpression) expressionNode()       {}
func (ExpressionLambdaExpression) lambdaExpressionNode() {}

// Implements [LambdaExpressionNode] of kind [STATEMENT_BODY_KIND].
type StatementLambdaExpression struct {
	Parameters []VariableNode
	Block      BlockNode
}

func (StatementLambdaExpression) GetKind() Kind { return LAMBDA_EXPRESSION }

func (slx StatementLambdaExpression) GetParameters() []VariableNode { return slx.Parameters }
func (slx StatementLambdaExpression) GetBody() Node                 { return slx.Block }
func (slx StatementLambdaExpression) GetBodyKind() BodyKind         { return STATEMENT_BODY_KIND }

func (StatementLambdaExpression) caseLabelNode()        {}
func (StatementLambdaExpression) expressionNode()       {}
func (StatementLambdaExpression) lambdaExpressionNode() {}

// Implements [PackageNode].
type Package struct {
	Annotations []AnnotationNode
	PackageName ExpressionNode
}

func (Package) GetKind() Kind { return PACKAGE }

func (p Package) GetAnnotations() []AnnotationNode { return p.Annotations }
func (p Package) GetPackageName() ExpressionNode   { return p.PackageName }

func (Package) packageNode() {}

// Implements [ParenthesizedNode].
type Parenthesized struct {
	Expression ExpressionNode
}

func (Parenthesized) GetKind() Kind { return PARENTHESIZED }

func (p Parenthesized) GetExpression() ExpressionNode { return p.Expression }

func (Parenthesized) caseLabelNode()     {}
func (Parenthesized) expressionNode()    {}
func (Parenthesized) parenthesizedNode() {}

// Implements [BindingPatternNode].
type BindingPattern struct {
	Variable VariableNode
}

func (BindingPattern) GetKind() Kind { return BINDING_PATTERN }

func (bp BindingPattern) GetVariable() VariableNode { return bp.Variable }

func (BindingPattern) caseLabelNode()      {}
func (BindingPattern) patternNode()        {}
func (BindingPattern) bindingPatternNode() {}

// Implements [GuardedPatternNode].
type GuardedPattern struct {
	Pattern    PatternNode
	Expression ExpressionNode
}

func (GuardedPattern) GetKind() Kind { return GUARDED_PATTERN }

func (gp GuardedPattern) GetPattern() PatternNode       { return gp.Pattern }
func (gp GuardedPattern) GetExpression() ExpressionNode { return gp.Expression }

func (GuardedPattern) caseLabelNode()      {}
func (GuardedPattern) patternNode()        {}
func (GuardedPattern) guardedPatternNode() {}

// Implements [ParenthesizedPatternNode].
type ParenthesizedPattern struct {
	Pattern PatternNode
}

func (ParenthesizedPattern) GetKind() Kind { return PARENTHESIZED_PATTERN }

func (pp ParenthesizedPattern) GetPattern() PatternNode { return pp.Pattern }

func (ParenthesizedPattern) caseLabelNode()            {}
func (ParenthesizedPattern) patternNode()              {}
func (ParenthesizedPattern) parenthesizedPatternNode() {}

// Implements [DefaultCaseLabelNode].
type DefaultCaseLabel struct {
}

func (DefaultCaseLabel) GetKind() Kind { return DEFAULT_CASE_LABEL }

func (DefaultCaseLabel) caseLabelNode()        {}
func (DefaultCaseLabel) defaultCaseLabelNode() {}

// Implements [PrimitiveTypeNode].
type PrimitiveType struct {
	PrimitiveTypeKind TypeKind
}

func (PrimitiveType) GetKind() Kind { return PRIMITIVE_TYPE }

func (pt PrimitiveType) GetPrimitiveTypeKind() TypeKind { return pt.PrimitiveTypeKind }

func (PrimitiveType) primitiveTypeNode() {}

// Implements [ReturnNode].
type Return struct {
	Expression ExpressionNode
}

func (Return) GetKind() Kind { return RETURN }

func (r Return) GetExpression() ExpressionNode { return r.Expression }

func (Return) statementNode() {}
func (Return) returnNode()    {}

// Implements [EmptyStatementNode].
type EmptyStatement struct {
}

func (EmptyStatement) GetKind() Kind { return EMPTY_STATEMENT }

func (EmptyStatement) statementNode()      {}
func (EmptyStatement) emptyStatementNode() {}

// Implements [SwitchNode].
type Switch struct {
	Expression ExpressionNode
	Cases      []CaseNode
}

func (Switch) GetKind() Kind { return SWITCH }

func (s Switch) GetExpression() ExpressionNode { return s.Expression }
func (s Switch) GetCases() []CaseNode          { return s.Cases }

func (Switch) statementNode() {}
func (Switch) switchNode()    {}

// Implements [SwitchExpressionNode].
type SwitchExpression struct {
	Expression ExpressionNode
	Cases      []CaseNode
}

func (SwitchExpression) GetKind() Kind { return SWITCH_EXPRESSION }

func (sx SwitchExpression) GetExpression() ExpressionNode { return sx.Expression }
func (sx SwitchExpression) GetCases() []CaseNode          { return sx.Cases }

func (SwitchExpression) caseLabelNode()        {}
func (SwitchExpression) expressionNode()       {}
func (SwitchExpression) switchExpressionNode() {}

// Implements [SynchronizedNode].
type Synchronized struct {
	Expression ExpressionNode
	Block      BlockNode
}

func (Synchronized) GetKind() Kind { return SYNCHRONIZED }

func (s Synchronized) GetExpression() ExpressionNode { return s.Expression }
func (s Synchronized) GetBlock() BlockNode           { return s.Block }

func (Synchronized) statementNode()    {}
func (Synchronized) synchronizedNode() {}

// Implements [ThrowNode].
type Throw struct {
	Expression ExpressionNode
}

func (Throw) GetKind() Kind { return THROW }

func (t Throw) GetExpression() ExpressionNode { return t.Expression }

func (Throw) statementNode() {}
func (Throw) throwNode()     {}

// Implements [TryNode].
type Try struct {
	Block        BlockNode
	Catches      []CatchNode
	FinallyBlock BlockNode
	Resources    []Node
}

func (Try) GetKind() Kind { return TRY }

func (t Try) GetBlock() BlockNode        { return t.Block }
func (t Try) GetCatches() []CatchNode    { return t.Catches }
func (t Try) GetFinallyBlock() BlockNode { return t.FinallyBlock }
func (t Try) GetResources() []Node       { return t.Resources }

func (Try) statementNode() {}
func (Try) tryNode()       {}

// Implements [ParameterizedTypeNode].
type ParameterizedType struct {
	Type          Node
	TypeArguments []Node
}

func (ParameterizedType) GetKind() Kind { return PARAMETERIZED_TYPE }

func (pt ParameterizedType) GetType() Node            { return pt.Type }
func (pt ParameterizedType) GetTypeArguments() []Node { return pt.TypeArguments }

func (ParameterizedType) parameterizedTypeNode() {}

// Implements [UnionTypeNode].
type UnionType struct {
	TypeAlternatives []Node
}

func (UnionType) GetKind() Kind { return UNION_TYPE }

func (ut UnionType) GetTypeAlternatives() []Node { return ut.TypeAlternatives }

func (UnionType) unionTypeNode() {}

// Implements [IntersectionTypeNode].
type IntersectionType struct {
	Bounds []Node
}

func (IntersectionType) GetKind() Kind { return INTERSECTION_TYPE }

func (it IntersectionType) GetBounds() []Node { return it.Bounds }

func (IntersectionType) intersectionTypeNode() {}

// Implements [TypeCastNode].
type TypeCast struct {
	Type       Node
	Expression ExpressionNode
}

func (TypeCast) GetKind() Kind { return TYPE_CAST }

func (tc TypeCast) GetType() Node                 { return tc.Type }
func (tc TypeCast) GetExpression() ExpressionNode { return tc.Expression }

func (TypeCast) caseLabelNode()  {}
func (TypeCast) expressionNode() {}
func (TypeCast) typeCastNode()   {}

// Implements [TypeParameterNode].
type TypeParameter struct {
	Name        string
	Bounds      []Node
	Annotations []AnnotationNode
}

func (TypeParameter) GetKind() Kind { return TYPE_PARAMETER }

func (tp TypeParameter) GetName() string                  { return tp.Name }
func (tp TypeParameter) GetBounds() []Node                { return tp.Bounds }
func (tp TypeParameter) GetAnnotations() []AnnotationNode { return tp.Annotations }

func (TypeParameter) typeParameterNode() {}

// Implements [VariableNode].
type Variable struct {
	Modifiers      ModifiersNode
	Name           string
	NameExpression ExpressionNode
	Type           Node
	Initializer    ExpressionNode
}

func (Variable) GetKind() Kind { return VARIABLE }

func (v Variable) GetModifiers() ModifiersNode       { return v.Modifiers }
func (v Variable) GetName() string                   { return v.Name }
func (v Variable) GetNameExpression() ExpressionNode { return v.NameExpression }
func (v Variable) GetType() Node                     { return v.Type }
func (v Variable) GetInitializer() ExpressionNode    { return v.Initializer }

func (Variable) statementNode() {}
func (Variable) variableNode()  {}

// Implements [WhileLoopNode].
type WhileLoop struct {
	Condition ExpressionNode
	Statement StatementNode
}

func (WhileLoop) GetKind() Kind { return WHILE_LOOP }

func (wl WhileLoop) GetCondition() ExpressionNode { return wl.Condition }
func (wl WhileLoop) GetStatement() StatementNode  { return wl.Statement }

func (WhileLoop) statementNode() {}
func (WhileLoop) whileLoopNode() {}

// Implements [UnaryNode] of kind [POSTFIX_INCREMENT].
type PostfixIncrement struct {
	Expression ExpressionNode
}

func (PostfixIncrement) GetKind() Kind { return POSTFIX_INCREMENT }

func (pi PostfixIncrement) GetExpression() ExpressionNode { return pi.Expression }

func (PostfixIncrement) caseLabelNode()  {}
func (PostfixIncrement) expressionNode() {}
func (PostfixIncrement) unaryNode()      {}

// Implements [UnaryNode] of kind [POSTFIX_DECREMENT].
type PostfixDecrement struct {
	Expression ExpressionNode
}

func (PostfixDecrement) GetKind() Kind { return POSTFIX_DECREMENT }

func (pd PostfixDecrement) GetExpression() ExpressionNode { return pd.Expression }

func (PostfixDecrement) caseLabelNode()  {}
func (PostfixDecrement) expressionNode() {}
func (PostfixDecrement) unaryNode()      {}

// Implements [UnaryNode] of kind [PREFIX_INCREMENT].
type PrefixIncrement struct {
	Expression ExpressionNode
}

func (PrefixIncrement) GetKind() Kind { return PREFIX_INCREMENT }

func (pi PrefixIncrement) GetExpression() ExpressionNode { return pi.Expression }

func (PrefixIncrement) caseLabelNode()  {}
func (PrefixIncrement) expressionNode() {}
func (PrefixIncrement) unaryNode()      {}

// Implements [UnaryNode] of kind [PREFIX_DECREMENT].
type PrefixDecrement struct {
	Expression ExpressionNode
}

func (PrefixDecrement) GetKind() Kind { return PREFIX_DECREMENT }

func (pd PrefixDecrement) GetExpression() ExpressionNode { return pd.Expression }

func (PrefixDecrement) caseLabelNode()  {}
func (PrefixDecrement) expressionNode() {}
func (PrefixDecrement) unaryNode()      {}

// Implements [UnaryNode] of kind [UNARY_PLUS].
type UnaryPlus struct {
	Expression ExpressionNode
}

func (UnaryPlus) GetKind() Kind { return UNARY_PLUS }

func (up UnaryPlus) GetExpression() ExpressionNode { return up.Expression }

func (UnaryPlus) caseLabelNode()  {}
func (UnaryPlus) expressionNode() {}
func (UnaryPlus) unaryNode()      {}

// Implements [UnaryNode] of kind [UNARY_MINUS].
type UnaryMinus struct {
	Expression ExpressionNode
}

func (UnaryMinus) GetKind() Kind { return UNARY_MINUS }

func (um UnaryMinus) GetExpression() ExpressionNode { return um.Expression }

func (UnaryMinus) caseLabelNode()  {}
func (UnaryMinus) expressionNode() {}
func (UnaryMinus) unaryNode()      {}

// Implements [UnaryNode] of kind [BITWISE_COMPLEMENT].
type BitwiseComplement struct {
	Expression ExpressionNode
}

func (BitwiseComplement) GetKind() Kind { return BITWISE_COMPLEMENT }

func (bc BitwiseComplement) GetExpression() ExpressionNode { return bc.Expression }

func (BitwiseComplement) caseLabelNode()  {}
func (BitwiseComplement) expressionNode() {}
func (BitwiseComplement) unaryNode()      {}

// Implements [UnaryNode] of kind [LOGICAL_COMPLEMENT].
type LogicalComplement struct {
	Expression ExpressionNode
}

func (LogicalComplement) GetKind() Kind { return LOGICAL_COMPLEMENT }

func (lc LogicalComplement) GetExpression() ExpressionNode { return lc.Expression }

func (LogicalComplement) caseLabelNode()  {}
func (LogicalComplement) expressionNode() {}
func (LogicalComplement) unaryNode()      {}

// Implements [BinaryNode] of kind [MULTIPLY].
type Multiply struct {
	LeftOperand  ExpressionNode
	RightOperand ExpressionNode
}

func (Multiply) GetKind() Kind { return MULTIPLY }

func (m Multiply) GetLeftOperand() ExpressionNode  { return m.LeftOperand }
func (m Multiply) GetRightOperand() ExpressionNode { return m.RightOperand }

func (Multiply) caseLabelNode()  {}
func (Multiply) expressionNode() {}
func (Multiply) binaryNode()     {}

// Implements [BinaryNode] of kind [DIVIDE].
type Divide struct {
	LeftOperand  ExpressionNode
	RightOperand ExpressionNode
}

func (Divide) GetKind() Kind { return DIVIDE }

func (d Divide) GetLeftOperand() ExpressionNode  { return d.LeftOperand }
func (d Divide) GetRightOperand() ExpressionNode { return d.RightOperand }

func (Divide) caseLabelNode()  {}
func (Divide) expressionNode() {}
func (Divide) binaryNode()     {}

// Implements [BinaryNode] of kind [REMAINDER].
type Remainder struct {
	LeftOperand  ExpressionNode
	RightOperand ExpressionNode
}

func (Remainder) GetKind() Kind { return REMAINDER }

func (r Remainder) GetLeftOperand() ExpressionNode  { return r.LeftOperand }
func (r Remainder) GetRightOperand() ExpressionNode { return r.RightOperand }

func (Remainder) caseLabelNode()  {}
func (Remainder) expressionNode() {}
func (Remainder) binaryNode()     {}

// Implements [BinaryNode] of kind [PLUS].
type Plus struct {
	LeftOperand  ExpressionNode
	RightOperand ExpressionNode
}

func (Plus) GetKind() Kind { return PLUS }

func (p Plus) GetLeftOperand() ExpressionNode  { return p.LeftOperand }
func (p Plus) GetRightOperand() ExpressionNode { return p.RightOperand }

func (Plus) caseLabelNode()  {}
func (Plus) expressionNode() {}
func (Plus) binaryNode()     {}

// Implements [BinaryNode] of kind [MINUS].
type Minus struct {
	LeftOperand  ExpressionNode
	RightOperand ExpressionNode
}

func (Minus) GetKind() Kind { return MINUS }

func (m Minus) GetLeftOperand() ExpressionNode  { return m.LeftOperand }
func (m Minus) GetRightOperand() ExpressionNode { return m.RightOperand }

func (Minus) caseLabelNode()  {}
func (Minus) expressionNode() {}
func (Minus) binaryNode()     {}

// Implements [BinaryNode] of kind [LEFT_SHIFT].
type LeftShift struct {
	LeftOperand  ExpressionNode
	RightOperand ExpressionNode
}

func (LeftShift) GetKind() Kind { return LEFT_SHIFT }

func (ls LeftShift) GetLeftOperand() ExpressionNode  { return ls.LeftOperand }
func (ls LeftShift) GetRightOperand() ExpressionNode { return ls.RightOperand }

func (LeftShift) caseLabelNode()  {}
func (LeftShift) expressionNode() {}
func (LeftShift) binaryNode()     {}

// Implements [BinaryNode] of kind [RIGHT_SHIFT].
type RightShift struct {
	LeftOperand  ExpressionNode
	RightOperand ExpressionNode
}

func (RightShift) GetKind() Kind { return RIGHT_SHIFT }

func (rs RightShift) GetLeftOperand() ExpressionNode  { return rs.LeftOperand }
func (rs RightShift) GetRightOperand() ExpressionNode { return rs.RightOperand }

func (RightShift) caseLabelNode()  {}
func (RightShift) expressionNode() {}
func (RightShift) binaryNode()     {}

// Implements [BinaryNode] of kind [UNSIGNED_RIGHT_SHIFT].
type UnsignedRightShift struct {
	LeftOperand  ExpressionNode
	RightOperand ExpressionNode
}

func (UnsignedRightShift) GetKind() Kind { return UNSIGNED_RIGHT_SHIFT }

func (urs UnsignedRightShift) GetLeftOperand() ExpressionNode  { return urs.LeftOperand }
func (urs UnsignedRightShift) GetRightOperand() ExpressionNode { return urs.RightOperand }

func (UnsignedRightShift) caseLabelNode()  {}
func (UnsignedRightShift) expressionNode() {}
func (UnsignedRightShift) binaryNode()     {}

// Implements [BinaryNode] of kind [LESS_THAN].
type LessThan struct {
	LeftOperand  ExpressionNode
	RightOperand ExpressionNode
}

func (LessThan) GetKind() Kind { return LESS_THAN }

func (lt LessThan) GetLeftOperand() ExpressionNode  { return lt.LeftOperand }
func (lt LessThan) GetRightOperand() ExpressionNode { return lt.RightOperand }

func (LessThan) caseLabelNode()  {}
func (LessThan) expressionNode() {}
func (LessThan) binaryNode()     {}

// Implements [BinaryNode] of kind [GREATER_THAN].
type GreaterThan struct {
	LeftOperand  ExpressionNode
	RightOperand ExpressionNode
}

func (GreaterThan) GetKind() Kind { return GREATER_THAN }

func (gt GreaterThan) GetLeftOperand() ExpressionNode  { return gt.LeftOperand }
func (gt GreaterThan) GetRightOperand() ExpressionNode { return gt.RightOperand }

func (GreaterThan) caseLabelNode()  {}
func (GreaterThan) expressionNode() {}
func (GreaterThan) binaryNode()     {}

// Implements [BinaryNode] of kind [LESS_THAN_EQUAL].
type LessThanEqual struct {
	LeftOperand  ExpressionNode
	RightOperand ExpressionNode
}

func (LessThanEqual) GetKind() Kind { return LESS_THAN_EQUAL }

func (lte LessThanEqual) GetLeftOperand() ExpressionNode  { return lte.LeftOperand }
func (lte LessThanEqual) GetRightOperand() ExpressionNode { return lte.RightOperand }

func (LessThanEqual) caseLabelNode()  {}
func (LessThanEqual) expressionNode() {}
func (LessThanEqual) binaryNode()     {}

// Implements [BinaryNode] of kind [GREATER_THAN_EQUAL].
type GreaterThanEqual struct {
	LeftOperand  ExpressionNode
	RightOperand ExpressionNode
}

func (GreaterThanEqual) GetKind() Kind { return GREATER_THAN_EQUAL }

func (gte GreaterThanEqual) GetLeftOperand() ExpressionNode  { return gte.LeftOperand }
func (gte GreaterThanEqual) GetRightOperand() ExpressionNode { return gte.RightOperand }

func (GreaterThanEqual) caseLabelNode()  {}
func (GreaterThanEqual) expressionNode() {}
func (GreaterThanEqual) binaryNode()     {}

// Implements [BinaryNode] of kind [EQUAL_TO].
type EqualTo struct {
	LeftOperand  ExpressionNode
	RightOperand ExpressionNode
}

func (EqualTo) GetKind() Kind { return EQUAL_TO }

func (eq EqualTo) GetLeftOperand() ExpressionNode  { return eq.LeftOperand }
func (eq EqualTo) GetRightOperand() ExpressionNode { return eq.RightOperand }

func (EqualTo) caseLabelNode()  {}
func (EqualTo) expressionNode() {}
func (EqualTo) binaryNode()     {}

// Implements [BinaryNode] of kind [NOT_EQUAL_TO].
type NotEqualTo struct {
	LeftOperand  ExpressionNode
	RightOperand ExpressionNode
}

func (NotEqualTo) GetKind() Kind { return NOT_EQUAL_TO }

func (neq NotEqualTo) GetLeftOperand() ExpressionNode  { return neq.LeftOperand }
func (neq NotEqualTo) GetRightOperand() ExpressionNode { return neq.RightOperand }

func (NotEqualTo) caseLabelNode()  {}
func (NotEqualTo) expressionNode() {}
func (NotEqualTo) binaryNode()     {}

// Implements [BinaryNode] of kind [AND].
type And struct {
	LeftOperand  ExpressionNode
	RightOperand ExpressionNode
}

func (And) GetKind() Kind { return AND }

func (a And) GetLeftOperand() ExpressionNode  { return a.LeftOperand }
func (a And) GetRightOperand() ExpressionNode { return a.RightOperand }

func (And) caseLabelNode()  {}
func (And) expressionNode() {}
func (And) binaryNode()     {}

// Implements [BinaryNode] of kind [XOR].
type Xor struct {
	LeftOperand  ExpressionNode
	RightOperand ExpressionNode
}

func (Xor) GetKind() Kind { return XOR }

func (x Xor) GetLeftOperand() ExpressionNode  { return x.LeftOperand }
func (x Xor) GetRightOperand() ExpressionNode { return x.RightOperand }

func (Xor) caseLabelNode()  {}
func (Xor) expressionNode() {}
func (Xor) binaryNode()     {}

// Implements [BinaryNode] of kind [OR].
type Or struct {
	LeftOperand  ExpressionNode
	RightOperand ExpressionNode
}

func (Or) GetKind() Kind { return OR }

func (or Or) GetLeftOperand() ExpressionNode  { return or.LeftOperand }
func (or Or) GetRightOperand() ExpressionNode { return or.RightOperand }

func (Or) caseLabelNode()  {}
func (Or) expressionNode() {}
func (Or) binaryNode()     {}

// Implements [BinaryNode] of kind [CONDITIONAL_AND].
type ConditionalAnd struct {
	LeftOperand  ExpressionNode
	RightOperand ExpressionNode
}

func (ConditionalAnd) GetKind() Kind { return CONDITIONAL_AND }

func (ca ConditionalAnd) GetLeftOperand() ExpressionNode  { return ca.LeftOperand }
func (ca ConditionalAnd) GetRightOperand() ExpressionNode { return ca.RightOperand }

func (ConditionalAnd) caseLabelNode()  {}
func (ConditionalAnd) expressionNode() {}
func (ConditionalAnd) binaryNode()     {}

// Implements [BinaryNode] of kind [CONDITIONAL_OR].
type ConditionalOr struct {
	LeftOperand  ExpressionNode
	RightOperand ExpressionNode
}

func (ConditionalOr) GetKind() Kind { return CONDITIONAL_OR }

func (or ConditionalOr) GetLeftOperand() ExpressionNode  { return or.LeftOperand }
func (or ConditionalOr) GetRightOperand() ExpressionNode { return or.RightOperand }

func (ConditionalOr) caseLabelNode()  {}
func (ConditionalOr) expressionNode() {}
func (ConditionalOr) binaryNode()     {}

// Implements [CompoundAssignmentNode] of kind [MULTIPLY_ASSIGNMENT].
type MultiplyAssignment struct {
	Variable   ExpressionNode
	Expression ExpressionNode
}

func (MultiplyAssignment) GetKind() Kind { return MULTIPLY_ASSIGNMENT }

func (ma MultiplyAssignment) GetVariable() ExpressionNode   { return ma.Variable }
func (ma MultiplyAssignment) GetExpression() ExpressionNode { return ma.Expression }

func (MultiplyAssignment) caseLabelNode()          {}
func (MultiplyAssignment) expressionNode()         {}
func (MultiplyAssignment) compoundAssignmentNode() {}

// Implements [CompoundAssignmentNode] of kind [DIVIDE_ASSIGNMENT].
type DivideAssignment struct {
	Variable   ExpressionNode
	Expression ExpressionNode
}

func (DivideAssignment) GetKind() Kind { return DIVIDE_ASSIGNMENT }

func (da DivideAssignment) GetVariable() ExpressionNode   { return da.Variable }
func (da DivideAssignment) GetExpression() ExpressionNode { return da.Expression }

func (DivideAssignment) caseLabelNode()          {}
func (DivideAssignment) expressionNode()         {}
func (DivideAssignment) compoundAssignmentNode() {}

// Implements [CompoundAssignmentNode] of kind [REMAINDER_ASSIGNMENT].
type RemainderAssignment struct {
	Variable   ExpressionNode
	Expression ExpressionNode
}

func (RemainderAssignment) GetKind() Kind { return REMAINDER_ASSIGNMENT }

func (ra RemainderAssignment) GetVariable() ExpressionNode   { return ra.Variable }
func (ra RemainderAssignment) GetExpression() ExpressionNode { return ra.Expression }

func (RemainderAssignment) caseLabelNode()          {}
func (RemainderAssignment) expressionNode()         {}
func (RemainderAssignment) compoundAssignmentNode() {}

// Implements [CompoundAssignmentNode] of kind [PLUS_ASSIGNMENT].
type PlusAssignment struct {
	Variable   ExpressionNode
	Expression ExpressionNode
}

func (PlusAssignment) GetKind() Kind { return PLUS_ASSIGNMENT }

func (pa PlusAssignment) GetVariable() ExpressionNode   { return pa.Variable }
func (pa PlusAssignment) GetExpression() ExpressionNode { return pa.Expression }

func (PlusAssignment) caseLabelNode()          {}
func (PlusAssignment) expressionNode()         {}
func (PlusAssignment) compoundAssignmentNode() {}

// Implements [CompoundAssignmentNode] of kind [MINUS_ASSIGNMENT].
type MinusAssignment struct {
	Variable   ExpressionNode
	Expression ExpressionNode
}

func (MinusAssignment) GetKind() Kind { return MINUS_ASSIGNMENT }

func (ma MinusAssignment) GetVariable() ExpressionNode   { return ma.Variable }
func (ma MinusAssignment) GetExpression() ExpressionNode { return ma.Expression }

func (MinusAssignment) caseLabelNode()          {}
func (MinusAssignment) expressionNode()         {}
func (MinusAssignment) compoundAssignmentNode() {}

// Implements [CompoundAssignmentNode] of kind [LEFT_SHIFT_ASSIGNMENT].
type LeftShiftAssignment struct {
	Variable   ExpressionNode
	Expression ExpressionNode
}

func (LeftShiftAssignment) GetKind() Kind { return LEFT_SHIFT_ASSIGNMENT }

func (lsa LeftShiftAssignment) GetVariable() ExpressionNode   { return lsa.Variable }
func (lsa LeftShiftAssignment) GetExpression() ExpressionNode { return lsa.Expression }

func (LeftShiftAssignment) caseLabelNode()          {}
func (LeftShiftAssignment) expressionNode()         {}
func (LeftShiftAssignment) compoundAssignmentNode() {}

// Implements [CompoundAssignmentNode] of kind [RIGHT_SHIFT_ASSIGNMENT].
type RightShiftAssignment struct {
	Variable   ExpressionNode
	Expression ExpressionNode
}

func (RightShiftAssignment) GetKind() Kind { return RIGHT_SHIFT_ASSIGNMENT }

func (rsa RightShiftAssignment) GetVariable() ExpressionNode   { return rsa.Variable }
func (rsa RightShiftAssignment) GetExpression() ExpressionNode { return rsa.Expression }

func (RightShiftAssignment) caseLabelNode()          {}
func (RightShiftAssignment) expressionNode()         {}
func (RightShiftAssignment) compoundAssignmentNode() {}

// Implements [CompoundAssignmentNode] of kind [UNSIGNED_RIGHT_SHIFT_ASSIGNMENT].
type UnsignedRightShiftAssignment struct {
	Variable   ExpressionNode
	Expression ExpressionNode
}

func (UnsignedRightShiftAssignment) GetKind() Kind { return UNSIGNED_RIGHT_SHIFT_ASSIGNMENT }

func (ursa UnsignedRightShiftAssignment) GetVariable() ExpressionNode   { return ursa.Variable }
func (ursa UnsignedRightShiftAssignment) GetExpression() ExpressionNode { return ursa.Expression }

func (UnsignedRightShiftAssignment) caseLabelNode()          {}
func (UnsignedRightShiftAssignment) expressionNode()         {}
func (UnsignedRightShiftAssignment) compoundAssignmentNode() {}

// Implements [CompoundAssignmentNode] of kind [AND_ASSIGNMENT].
type AndAssignment struct {
	Variable   ExpressionNode
	Expression ExpressionNode
}

func (AndAssignment) GetKind() Kind { return AND_ASSIGNMENT }

func (aa AndAssignment) GetVariable() ExpressionNode   { return aa.Variable }
func (aa AndAssignment) GetExpression() ExpressionNode { return aa.Expression }

func (AndAssignment) caseLabelNode()          {}
func (AndAssignment) expressionNode()         {}
func (AndAssignment) compoundAssignmentNode() {}

// Implements [CompoundAssignmentNode] of kind [XOR_ASSIGNMENT].
type XorAssignment struct {
	Variable   ExpressionNode
	Expression ExpressionNode
}

func (XorAssignment) GetKind() Kind { return XOR_ASSIGNMENT }

func (xa XorAssignment) GetVariable() ExpressionNode   { return xa.Variable }
func (xa XorAssignment) GetExpression() ExpressionNode { return xa.Expression }

func (XorAssignment) caseLabelNode()          {}
func (XorAssignment) expressionNode()         {}
func (XorAssignment) compoundAssignmentNode() {}

// Implements [CompoundAssignmentNode] of kind [OR_ASSIGNMENT].
type OrAssignment struct {
	Variable   ExpressionNode
	Expression ExpressionNode
}

func (OrAssignment) GetKind() Kind { return OR_ASSIGNMENT }

func (oa OrAssignment) GetVariable() ExpressionNode   { return oa.Variable }
func (oa OrAssignment) GetExpression() ExpressionNode { return oa.Expression }

func (OrAssignment) caseLabelNode()          {}
func (OrAssignment) expressionNode()         {}
func (OrAssignment) compoundAssignmentNode() {}

// Implements [LiteralNode] of kind [INT_LITERAL].
type IntLiteral struct {
	Value string
}

func (IntLiteral) GetKind() Kind { return INT_LITERAL }

func (il IntLiteral) GetValue() string { return il.Value }

func (IntLiteral) caseLabelNode()  {}
func (IntLiteral) expressionNode() {}
func (IntLiteral) literalNode()    {}

// Implements [LiteralNode] of kind [LONG_LITERAL].
type LongLiteral struct {
	Value string
}

func (LongLiteral) GetKind() Kind { return LONG_LITERAL }

func (ll LongLiteral) GetValue() string { return ll.Value }

func (LongLiteral) caseLabelNode()  {}
func (LongLiteral) expressionNode() {}
func (LongLiteral) literalNode()    {}

// Implements [LiteralNode] of kind [FLOAT_LITERAL].
type FloatLiteral struct {
	Value string
}

func (FloatLiteral) GetKind() Kind { return FLOAT_LITERAL }

func (fl FloatLiteral) GetValue() string { return fl.Value }

func (FloatLiteral) caseLabelNode()  {}
func (FloatLiteral) expressionNode() {}
func (FloatLiteral) literalNode()    {}

// Implements [LiteralNode] of kind [DOUBLE_LITERAL].
type DoubleLiteral struct {
	Value string
}

func (DoubleLiteral) GetKind() Kind { return DOUBLE_LITERAL }

func (dl DoubleLiteral) GetValue() string { return dl.Value }

func (DoubleLiteral) caseLabelNode()  {}
func (DoubleLiteral) expressionNode() {}
func (DoubleLiteral) literalNode()    {}

// Implements [LiteralNode] of kind [BOOLEAN_LITERAL].
type BooleanLiteral struct {
	Value bool
}

func (BooleanLiteral) GetKind() Kind { return BOOLEAN_LITERAL }

func (bl BooleanLiteral) GetValue() string {
	if bl.Value {
		return "true"
	}
	return "false"
}

func (BooleanLiteral) caseLabelNode()  {}
func (BooleanLiteral) expressionNode() {}
func (BooleanLiteral) literalNode()    {}

// Implements [LiteralNode] of kind [CHAR_LITERAL].
type CharLiteral struct {
	Value string
}

func (CharLiteral) GetKind() Kind { return CHAR_LITERAL }

func (cl CharLiteral) GetValue() string { return cl.Value }

func (CharLiteral) caseLabelNode()  {}
func (CharLiteral) expressionNode() {}
func (CharLiteral) literalNode()    {}

// Implements [LiteralNode] of kind [STRING_LITERAL].
type StringLiteral struct {
	Value string
}

func (StringLiteral) GetKind() Kind { return STRING_LITERAL }

func (sl StringLiteral) GetValue() string { return sl.Value }

func (StringLiteral) caseLabelNode()  {}
func (StringLiteral) expressionNode() {}
func (StringLiteral) literalNode()    {}

// Implements [LiteralNode] of kind [NULL_LITERAL].
type NullLiteral struct {
}

func (NullLiteral) GetKind() Kind { return NULL_LITERAL }

func (NullLiteral) GetValue() string { return "null" }

func (NullLiteral) caseLabelNode()  {}
func (NullLiteral) expressionNode() {}
func (NullLiteral) literalNode()    {}

// Implements [WildcardNode] of kind [UNBOUNDED_WILDCARD].
type UnboundedWildcard struct {
}

func (UnboundedWildcard) GetKind() Kind { return UNBOUNDED_WILDCARD }

func (uw UnboundedWildcard) GetBound() Node { return nil }

func (UnboundedWildcard) wildcardNode() {}

// Implements [WildcardNode] of kind [EXTENDS_WILDCARD].
type ExtendsWildcard struct {
	Bound Node
}

func (ExtendsWildcard) GetKind() Kind { return EXTENDS_WILDCARD }

func (xw ExtendsWildcard) GetBound() Node { return xw.Bound }

func (ExtendsWildcard) wildcardNode() {}

// Implements [WildcardNode] of kind [SUPER_WILDCARD].
type SuperWildcard struct {
	Bound Node
}

func (SuperWildcard) GetKind() Kind { return SUPER_WILDCARD }

func (sw SuperWildcard) GetBound() Node { return sw.Bound }

func (SuperWildcard) wildcardNode() {}

// Implements [ErroneousNode].
type Erroneous struct {
	ErrorNodes []Node
}

func (Erroneous) GetKind() Kind { return ERRONEOUS }

func (e Erroneous) GetErrorNodes() []Node { return e.ErrorNodes }

func (Erroneous) caseLabelNode()  {}
func (Erroneous) expressionNode() {}
func (Erroneous) erroneousNode()  {}

// Implements [ClassNode] of kind [INTERFACE].
type Interface struct {
	Modifiers      ModifiersNode
	SimpleName     string
	TypeParameters []TypeParameterNode
	ExtendsClause  Node
	PermitsClause  []Node
	Members        []Node
}

func (Interface) GetKind() Kind { return INTERFACE }

func (i Interface) GetModifiers() ModifiersNode            { return i.Modifiers }
func (i Interface) GetSimpleName() string                  { return i.SimpleName }
func (i Interface) GetTypeParameters() []TypeParameterNode { return i.TypeParameters }
func (i Interface) GetExtendsClause() Node                 { return i.ExtendsClause }
func (i Interface) GetImplementsClause() []Node            { return nil }
func (i Interface) GetPermitsClause() []Node               { return i.PermitsClause }
func (i Interface) GetMembers() []Node                     { return i.Members }

func (Interface) statementNode() {}
func (Interface) classNode()     {}

// Implements [ClassNode] of kind [ENUM].
type Enum struct {
	Modifiers  ModifiersNode
	SimpleName string
	Members    []Node
}

func (Enum) GetKind() Kind { return ENUM }

func (e Enum) GetModifiers() ModifiersNode            { return e.Modifiers }
func (e Enum) GetSimpleName() string                  { return e.SimpleName }
func (e Enum) GetTypeParameters() []TypeParameterNode { return nil }
func (e Enum) GetExtendsClause() Node                 { return nil }
func (e Enum) GetImplementsClause() []Node            { return nil }
func (e Enum) GetPermitsClause() []Node               { return nil }
func (e Enum) GetMembers() []Node                     { return e.Members }

func (Enum) statementNode() {}
func (Enum) classNode()     {}

// Implements [ClassNode] of kind [ANNOTATION_TYPE].
type AnnotationType struct {
	Modifiers  ModifiersNode
	SimpleName string
	Members    []Node
}

func (AnnotationType) GetKind() Kind { return ANNOTATION_TYPE }

func (at AnnotationType) GetModifiers() ModifiersNode            { return at.Modifiers }
func (at AnnotationType) GetSimpleName() string                  { return at.SimpleName }
func (at AnnotationType) GetTypeParameters() []TypeParameterNode { return nil }
func (at AnnotationType) GetExtendsClause() Node                 { return nil }
func (at AnnotationType) GetImplementsClause() []Node            { return nil }
func (at AnnotationType) GetPermitsClause() []Node               { return nil }
func (at AnnotationType) GetMembers() []Node                     { return at.Members }

func (AnnotationType) statementNode() {}
func (AnnotationType) classNode()     {}

// Implements [ModuleNode].
type Module struct {
	Annotations []AnnotationNode
	ModuleType  ModuleKind
	Name        ExpressionNode
	Directives  []DirectiveNode
}

func (Module) GetKind() Kind { return MODULE }

func (m Module) GetAnnotations() []AnnotationNode { return m.Annotations }
func (m Module) GetModuleType() ModuleKind        { return m.ModuleType }
func (m Module) GetName() ExpressionNode          { return m.Name }
func (m Module) GetDirectives() []DirectiveNode   { return m.Directives }

func (Module) moduleNode() {}

// Implements [ExportsNode].
type Exports struct {
	PackageName ExpressionNode
	ModuleNames []ExpressionNode
}

func (Exports) GetKind() Kind { return EXPORTS }

func (x Exports) GetPackageName() ExpressionNode   { return x.PackageName }
func (x Exports) GetModuleNames() []ExpressionNode { return x.ModuleNames }

func (Exports) directiveNode() {}
func (Exports) exportsNode()   {}

// Implements [OpensNode].
type Opens struct {
	PackageName ExpressionNode
	ModuleNames []ExpressionNode
}

func (Opens) GetKind() Kind { return OPENS }

func (o Opens) GetPackageName() ExpressionNode   { return o.PackageName }
func (o Opens) GetModuleNames() []ExpressionNode { return o.ModuleNames }

func (Opens) directiveNode() {}
func (Opens) opensNode()     {}

// Implements [ProvidesNode].
type Provides struct {
	ServiceName         ExpressionNode
	ImplementationNames []ExpressionNode
}

func (Provides) GetKind() Kind { return PROVIDES }

func (p Provides) GetServiceName() ExpressionNode           { return p.ServiceName }
func (p Provides) GetImplementationNames() []ExpressionNode { return p.ImplementationNames }

func (Provides) directiveNode() {}
func (Provides) providesNode()  {}

// Implements [ClassNode] of kind [RECORD].
type Record struct {
	Modifiers        ModifiersNode
	SimpleName       string
	TypeParameters   []TypeParameterNode
	ImplementsClause []Node
	Members          []Node
}

func (Record) GetKind() Kind { return RECORD }

func (r Record) GetModifiers() ModifiersNode            { return r.Modifiers }
func (r Record) GetSimpleName() string                  { return r.SimpleName }
func (r Record) GetTypeParameters() []TypeParameterNode { return r.TypeParameters }
func (r Record) GetExtendsClause() Node                 { return nil }
func (r Record) GetImplementsClause() []Node            { return r.ImplementsClause }
func (r Record) GetPermitsClause() []Node               { return nil }
func (r Record) GetMembers() []Node                     { return r.Members }

func (Record) statementNode() {}
func (Record) classNode()     {}

// Implements [RequiresNode].
type Requires struct {
	Static     bool
	Transitive bool
	ModuleName ExpressionNode
}

func (Requires) GetKind() Kind { return REQUIRES }

func (r Requires) IsStatic() bool                { return r.Static }
func (r Requires) IsTransitive() bool            { return r.Transitive }
func (r Requires) GetModuleName() ExpressionNode { return r.ModuleName }

func (Requires) directiveNode() {}
func (Requires) requiresNode()  {}

// Implements [UsesNode].
type Uses struct {
	ServiceName ExpressionNode
}

func (Uses) GetKind() Kind { return USES }

func (u Uses) GetServiceName() ExpressionNode { return u.ServiceName }

func (Uses) directiveNode() {}
func (Uses) usesNode()      {}

// Implements [YieldNode].
type Yield struct {
	Value ExpressionNode
}

func (Yield) GetKind() Kind { return YIELD }

func (y Yield) GetValue() ExpressionNode { return y.Value }

func (Yield) statementNode() {}
func (Yield) yieldNode()     {}
