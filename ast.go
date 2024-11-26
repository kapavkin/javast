package javast

import (
	"io"
)

// A Kind enumerates all kinds of nodes.
type Kind int

const (
	ANNOTATED_TYPE                  Kind = iota // Used for instances of [AnnotatedTypeNode] representing annotated types.
	ANNOTATION                                  // Used for instances of [AnnotationNode] representing declaration annotations.
	TYPE_ANNOTATION                             // Used for instances of [AnnotationNode] representing type annotations.
	ARRAY_ACCESS                                // Used for instances of [ArrayAccessNode].
	ARRAY_TYPE                                  // Used for instances of [ArrayTypeNode].
	ASSERT                                      // Used for instances of [AssertNode].
	ASSIGNMENT                                  // Used for instances of [AssignmentNode].
	BLOCK                                       // Used for instances of [BlockNode].
	BREAK                                       // Used for instances of [BreakNode].
	CASE                                        // Used for instances of [CaseNode].
	CATCH                                       // Used for instances of [CatchNode].
	CLASS                                       // Used for instances of [ClassNode] representing classes.
	COMPILATION_UNIT                            // Used for instances of [CompilationUnitNode].
	CONDITIONAL_EXPRESSION                      // Used for instances of [ConditionalExpressionNode].
	CONTINUE                                    // Used for instances of [ContinueNode].
	DO_WHILE_LOOP                               // Used for instances of [DoWhileLoopNode].
	ENHANCED_FOR_LOOP                           // Used for instances of [EnhancedForLoopNode].
	EXPRESSION_STATEMENT                        // Used for instances of [ExpressionStatementNode].
	MEMBER_SELECT                               // Used for instances of [MemberSelectNode].
	MEMBER_REFERENCE                            // Used for instances of [MemberReferenceNode].
	FOR_LOOP                                    // Used for instances of [ForLoopNode].
	IDENTIFIER                                  // Used for instances of [IdentifierNode].
	IF                                          // Used for instances of [IfNode].
	IMPORT                                      // Used for instances of [ImportNode].
	INSTANCE_OF                                 // Used for instances of [InstanceOfNode].
	LABELED_STATEMENT                           // Used for instances of [LabeledStatementNode].
	METHOD                                      // Used for instances of [MethodNode].
	METHOD_INVOCATION                           // Used for instances of [MethodInvocationNode].
	MODIFIERS                                   // Used for instances of [ModifiersNode].
	NEW_ARRAY                                   // Used for instances of [NewArrayNode].
	NEW_CLASS                                   // Used for instances of [NewClassNode].
	LAMBDA_EXPRESSION                           // Used for instances of [LambdaExpressionNode].
	PACKAGE                                     // Used for instances of [PackageNode].
	PARENTHESIZED                               // Used for instances of [ParenthesizedNode].
	BINDING_PATTERN                             // Used for instances of [BindingPatternNode].
	GUARDED_PATTERN                             // Used for instances of [GuardedPatternNode].
	PARENTHESIZED_PATTERN                       // Used for instances of [ParenthesizedPatternNode].
	DEFAULT_CASE_LABEL                          // Used for instances of [DefaultCaseLabelNode].
	PRIMITIVE_TYPE                              // Used for instances of [PrimitiveTypeNode].
	RETURN                                      // Used for instances of [ReturnNode].
	EMPTY_STATEMENT                             // Used for instances of [EmptyStatementNode].
	EMPTY_EXPRESSION                            // Used for instances of [EmptyExpressionNode].
	SWITCH                                      // Used for instances of [SwitchNode].
	SWITCH_EXPRESSION                           // Used for instances of [SwitchExpressionNode].
	SYNCHRONIZED                                // Used for instances of [SynchronizedNode].
	THROW                                       // Used for instances of [ThrowNode].
	TRY                                         // Used for instances of [TryNode].
	PARAMETERIZED_TYPE                          // Used for instances of [ParameterizedTypeNode].
	UNION_TYPE                                  // Used for instances of [UnionTypeNode].
	INTERSECTION_TYPE                           // Used for instances of [IntersectionTypeNode].
	TYPE_CAST                                   // Used for instances of [TypeCastNode].
	TYPE_PARAMETER                              // Used for instances of [TypeParameterNode].
	VARIABLE                                    // Used for instances of [VariableNode].
	WHILE_LOOP                                  // Used for instances of [WhileLoopNode].
	POSTFIX_INCREMENT                           // Used for instances of [UnaryNode] representing postfix increment operator "++".
	POSTFIX_DECREMENT                           // Used for instances of [UnaryNode] representing postfix decrement operator "--".
	PREFIX_INCREMENT                            // Used for instances of [UnaryNode] representing prefix increment operator "++".
	PREFIX_DECREMENT                            // Used for instances of [UnaryNode] representing prefix decrement operator "--".
	UNARY_PLUS                                  // Used for instances of [UnaryNode] representing unary plus operator "+".
	UNARY_MINUS                                 // Used for instances of [UnaryNode] representing unary minus operator "-".
	BITWISE_COMPLEMENT                          // Used for instances of [UnaryNode] representing bitwise complement operator "~".
	LOGICAL_COMPLEMENT                          // Used for instances of [UnaryNode] representing logical complement operator "!".
	MULTIPLY                                    // Used for instances of [BinaryNode] representing multiplication "*".
	DIVIDE                                      // Used for instances of [BinaryNode] representing division "/".
	REMAINDER                                   // Used for instances of [BinaryNode] representing remainder "%".
	PLUS                                        // Used for instances of [BinaryNode] representing addition or string concatenation "+".
	MINUS                                       // Used for instances of [BinaryNode] representing subtraction "-".
	LEFT_SHIFT                                  // Used for instances of [BinaryNode] representing left shift "<<".
	RIGHT_SHIFT                                 // Used for instances of [BinaryNode] representing right shift ">>".
	UNSIGNED_RIGHT_SHIFT                        // Used for instances of [BinaryNode] representing unsigned right shift ">>>".
	LESS_THAN                                   // Used for instances of [BinaryNode] representing less-than "<".
	GREATER_THAN                                // Used for instances of [BinaryNode] representing greater-than ">".
	LESS_THAN_EQUAL                             // Used for instances of [BinaryNode] representing less-than-equal "<=".
	GREATER_THAN_EQUAL                          // Used for instances of [BinaryNode] representing greater-than-equal ">=".
	EQUAL_TO                                    // Used for instances of [BinaryNode] representing equal-to "==".
	NOT_EQUAL_TO                                // Used for instances of [BinaryNode] representing not-equal-to "!=".
	AND                                         // Used for instances of [BinaryNode] representing bitwise and logical "and" "&".
	XOR                                         // Used for instances of [BinaryNode] representing bitwise and logical "xor" "^".
	OR                                          // Used for instances of [BinaryNode] representing bitwise and logical "or" "|".
	CONDITIONAL_AND                             // Used for instances of [BinaryNode] representing conditional-and "&&".
	CONDITIONAL_OR                              // Used for instances of [BinaryNode] representing conditional-or "||".
	MULTIPLY_ASSIGNMENT                         // Used for instances of [CompoundAssignmentNode] representing multiplication assignment "*=".
	DIVIDE_ASSIGNMENT                           // Used for instances of [CompoundAssignmentNode] representing division assignment "/=".
	REMAINDER_ASSIGNMENT                        // Used for instances of [CompoundAssignmentNode] representing remainder assignment "%=".
	PLUS_ASSIGNMENT                             // Used for instances of [CompoundAssignmentNode] representing addition or string concatenation assignment "+=".
	MINUS_ASSIGNMENT                            // Used for instances of [CompoundAssignmentNode] representing subtraction assignment "-=".
	LEFT_SHIFT_ASSIGNMENT                       // Used for instances of [CompoundAssignmentNode] representing left shift assignment "<<=".
	RIGHT_SHIFT_ASSIGNMENT                      // Used for instances of [CompoundAssignmentNode] representing right shift assignment ">>=".
	UNSIGNED_RIGHT_SHIFT_ASSIGNMENT             // Used for instances of [CompoundAssignmentNode] representing unsigned right shift assignment ">>>=".
	AND_ASSIGNMENT                              // Used for instances of [CompoundAssignmentNode] representing bitwise and logical "and" assignment "&=".
	XOR_ASSIGNMENT                              // Used for instances of [CompoundAssignmentNode] representing bitwise and logical "xor" assignment "^=".
	OR_ASSIGNMENT                               // Used for instances of [CompoundAssignmentNode] representing bitwise and logical "or" assignment "|=".
	INT_LITERAL                                 // Used for instances of [LiteralNode] representing an integral literal expression of type "int".
	LONG_LITERAL                                // Used for instances of [LiteralNode] representing an integral literal expression of type "long".
	FLOAT_LITERAL                               // Used for instances of [LiteralNode] representing a floating-point literal expression of type "float".
	DOUBLE_LITERAL                              // Used for instances of [LiteralNode] representing a floating-point literal expression of type "double".
	BOOLEAN_LITERAL                             // Used for instances of [LiteralNode] representing a boolean literal expression of type "boolean".
	CHAR_LITERAL                                // Used for instances of [LiteralNode] representing a character literal expression of type "char".
	STRING_LITERAL                              // Used for instances of [LiteralNode] representing a string literal expression of type "String".
	NULL_LITERAL                                // Used for instances of [LiteralNode] representing the use of "null".
	UNBOUNDED_WILDCARD                          // Used for instances of [WildcardNode] representing an unbounded wildcard type argument.
	EXTENDS_WILDCARD                            // Used for instances of [WildcardNode] representing an extends bounded wildcard type argument.
	SUPER_WILDCARD                              // Used for instances of [WildcardNode] representing a super bounded wildcard type argument.
	ERRONEOUS                                   // Used for instances of [ErroneousNode].
	INTERFACE                                   // Used for instances of [ClassNode] representing interfaces.
	ENUM                                        // Used for instances of [ClassNode] representing enums.
	ANNOTATION_TYPE                             // Used for instances of [ClassNode] representing annotation types.
	MODULE                                      // Used for instances of [ModuleNode] representing module declarations.
	EXPORTS                                     // Used for instances of [ExportsNode] representing exports directives in a module declaration.
	OPENS                                       // Used for instances of [OpensNode] representing opens directives in a module declaration.
	PROVIDES                                    // Used for instances of [ProvidesNode] representing provides directives in a module declaration.
	RECORD                                      // Used for instances of [ClassNode] representing records.
	REQUIRES                                    // Used for instances of [RequiresNode] representing requires directives in a module declaration.
	USES                                        // Used for instances of [UsesNode] representing uses directives in a module declaration.
	YIELD                                       // Used for instances of [YieldNode].
	OTHER                                       // An implementation-reserved node. This is the not the node you are looking for.
)

// Lambda expressions come in two forms:
// - expression lambdas, whose body is an expression,
// - statement lambdas, whose body is a block.
type BodyKind int

const (
	EXPRESSION_BODY_KIND BodyKind = iota // Enum constant for expression lambdas.
	STATEMENT_BODY_KIND                  // Enum constant for statement lambdas.
)

// There are two kinds of member references:
// - method references,
// - constructor references.
type ReferenceMode int

const (
	INVOKE_REFERENCE_MODE ReferenceMode = iota // Enum constant for method references.
	NEW_REFERENCE_MODE                         // Enum constant for constructor references.
)

// The syntactic form of this case:
// STATEMENT: "case <expression>: <statements>"
// RULE: "case <expression> -> <expression>/<statement>"
type CaseKind int

const (
	STATEMENT_CASE_KIND CaseKind = iota // Case is in the form: "case <expression>: <statements>".
	RULE_CASE_KIND                      // Case is in the form: "case <expression> -> <expression>".
)

// Represents a modifier on a program element such as a class, method, or field.
type Modifier int

const (
	PUBLIC_MODIFIER       Modifier = iota // The modifier "public".
	PROTECTED_MODIFIER                    // The modifier "protected".
	PRIVATE_MODIFIER                      // The modifier "private".
	ABSTRACT_MODIFIER                     // The modifier "abstract".
	DEFAULT_MODIFIER                      // The modifier "default".
	STATIC_MODIFIER                       // The modifier "static".
	SEALED_MODIFIER                       // The modifier "sealed".
	NON_SEALED_MODIFIER                   // The modifier "non-sealed".
	FINAL_MODIFIER                        // The modifier "final".
	TRANSIENT_MODIFIER                    // The modifier "transient".
	VOLATILE_MODIFIER                     // The modifier "volatile".
	SYNCHRONIZED_MODIFIER                 // The modifier "synchronized".
	NATIVE_MODIFIER                       // The modifier "native".
	STRICTFP_MODIFIER                     // The modifier "strictfp".
)

// The kind of the module.
type ModuleKind int

const (
	OPEN_MODULE_KIND   ModuleKind = iota // Open module.
	STRONG_MODULE_KIND                   // Strong module.
)

// The kind of a type mirror.
type TypeKind int

const (
	BOOLEAN_TYPE_KIND      TypeKind = iota // The primitive type "boolean".
	BYTE_TYPE_KIND                         // The primitive type "byte".
	SHORT_TYPE_KIND                        // The primitive type "short".
	INT_TYPE_KIND                          // The primitive type "int".
	LONG_TYPE_KIND                         // The primitive type "long".
	CHAR_TYPE_KIND                         // The primitive type "char".
	FLOAT_TYPE_KIND                        // The primitive type "float".
	DOUBLE_TYPE_KIND                       // The primitive type "double".
	VOID_TYPE_KIND                         // The pseudo-type corresponding to the keyword "void".
	NONE_TYPE_KIND                         // A pseudo-type used where no actual type is appropriate.
	NULL_TYPE_KIND                         // The null type.
	ARRAY_TYPE_KIND                        // An array type.
	DECLARED_TYPE_KIND                     // A class or interface type.
	ERROR_TYPE_KIND                        // A class or interface type that could not be resolved.
	TYPEVAR_TYPE_KIND                      // A type variable.
	WILDCARD_TYPE_KIND                     // A wildcard type argument.
	PACKAGE_TYPE_KIND                      // A pseudo-type corresponding to a package element.
	EXECUTABLE_TYPE_KIND                   // A method, constructor, or initializer.
	UNION_TYPE_KIND                        // A union type.
	INTERSECTION_TYPE_KIND                 // An intersection type.
	MODULE_TYPE_KIND                       // A pseudo-type corresponding to a module element.
	OTHER_TYPE_KIND                        // An implementation-reserved type. This is not the type you are looking for.
)

// Common interface for all nodes in an abstract syntax tree.
type Node interface {
	io.WriterTo
	GetKind() Kind
}

// A tree node for an array type.
// For example:
//
//	type []
type ArrayTypeNode interface {
	Node
	GetType() Node  // Returns the element type of this array type.
	arrayTypeNode() // arrayTypeNode() ensures that only array type nodes can be assigned to an ArrayTypeNode.
}

// A marker interface for "Node"s that may be used as "Case" labels.
type CaseLabelNode interface {
	Node
	caseLabelNode() // caseLabelNode() ensures that only case label nodes can be assigned to a CaseLabelNode.
}

// A case label that marks "default" in "case null, default".
type DefaultCaseLabelNode interface {
	CaseLabelNode
	defaultCaseLabelNode() // defaultCaseLabelNode() ensures that only default case label nodes can be assigned to a DefaultCaseLabelNode.
}

// A tree node used as the base class for the different types of expressions.
type ExpressionNode interface {
	Node
	CaseLabelNode
	expressionNode() // expressionNode() ensures that only expression nodes can be assigned to an ExpressionNode.
}

// A tree node for an annotated type.
// For example:
//
//	@annotationType String
//
//	@annotationType ( arguments ) Date
type AnnotatedTypeNode interface {
	ExpressionNode
	GetAnnotations() []AnnotationNode  // Returns the annotations associated with this type expression.
	GetUnderlyingType() ExpressionNode // Returns the underlying type with which the annotations are associated.
	annotatedTypeNode()                // annotatedTypeNode() ensures that only annotated type nodes can be assigned to an AnnotatedTypeNode.
}

// A tree node for an annotation.
// For example:
//
//	@annotationType
//
//	@annotationType ( arguments )
type AnnotationNode interface {
	ExpressionNode
	GetAnnotationType() Node        // Returns the annotation type.
	GetArguments() []ExpressionNode // Returns the arguments, if any, for the annotation.
	annotationNode()                // annotationNode() ensures that only annotation nodes can be assigned to an AnnotationNode.
}

// A tree node for an array access expression.
// For example:
//
//	expression [ index ]
type ArrayAccessNode interface {
	ExpressionNode
	GetExpression() ExpressionNode // Returns the expression for the array being accessed.
	GetIndex() ExpressionNode      // Returns the expression for the index.
	arrayAccessNode()              // arrayAccessNode() ensures that only array access nodes can be assigned to an ArrayAccessNode.
}

// A tree node for an assignment expression.
// For example:
//
//	variable = expression
type AssignmentNode interface {
	ExpressionNode
	GetVariable() ExpressionNode   // Returns the variable being assigned to.
	GetExpression() ExpressionNode // Returns the expression being assigned to the variable.
	assignmentNode()               // assignmentNode() ensures that only assignment nodes can be assigned to an AssignmentNode.
}

// A tree node for a binary expression.
// Use [Node].GetKind to determine the kind of operator.
// For example:
//
//	leftOperand operator rightOperand
type BinaryNode interface {
	ExpressionNode
	GetLeftOperand() ExpressionNode  // Returns the left (first) operand of the expression.
	GetRightOperand() ExpressionNode // Returns the right (second) operand of the expression.
	binaryNode()                     // binaryNode() ensures that only binary nodes can be assigned to a BinaryNode.
}

// A tree node for compound assignment operator.
// Use [Node].GetKind to determine the kind of operator.
// For example:
//
//	variable operator expression
type CompoundAssignmentNode interface {
	ExpressionNode
	GetVariable() ExpressionNode   // Returns the variable on the left hand side of the compound assignment.
	GetExpression() ExpressionNode // Returns the expression on the right hand side of the compound assignment.
	compoundAssignmentNode()       // compoundAssignmentNode() ensures that only compound assignment nodes can be assigned to a CompoundAssignmentNode.
}

// A tree node for the conditional operator "? :".
// For example:
//
//	condition ? trueExpression : falseExpression
type ConditionalExpressionNode interface {
	ExpressionNode
	GetCondition() ExpressionNode       // Returns the condition.
	GetTrueExpression() ExpressionNode  // Returns the expression to be evaluated if the condition is true.
	GetFalseExpression() ExpressionNode // Returns the expression to be evaluated if the condition is false.
	conditionalExpressionNode()         // conditionalExpressionNode() ensures that only conditional expression nodes can be assigned to a ConditionalExpressionNode.
}

// A tree node to stand in for a malformed expression.
type ErroneousNode interface {
	ExpressionNode
	GetErrorNodes() []Node // Returns any nodes that were saved in this node.
	erroneousNode()        // erroneousNode() ensures that only erroneous nodes can be assigned to an ErroneousNode.
}

// A tree node for an identifier expression.
// For example:
//
//	name
type IdentifierNode interface {
	ExpressionNode
	GetName() string // Returns the name of the identifier.
	identifierNode() // identifierNode() ensures that only identifier nodes can be assigned to an IdentifierNode.
}

// A tree node for an "instanceof" expression.
// For example:
//
//	expression instanceof type
type InstanceOfNode interface {
	ExpressionNode
	GetExpression() ExpressionNode // Returns the expression to be tested.
	GetType() Node                 // Returns the type for which to check.
	GetPattern() PatternNode       // Returns the tested pattern, or nil if this instanceof does not use a pattern.
	instanceOfNode()               // instanceOfNode() ensures that only instance of nodes can be assigned to an InstanceOfNode.
}

// A tree node for a lambda expression.
// For example:
//
//	()->{}
//
//	(List<String> ls)->ls.size()
//
//	(x,y)-> { return x + y; }
type LambdaExpressionNode interface {
	ExpressionNode
	GetParameters() []VariableNode // Returns the parameters of this lambda expression.
	GetBody() Node                 // Returns the body of the lambda expression.
	GetBodyKind() BodyKind         // Returns the kind of the body of the lambda expression.
	lambdaExpressionNode()         // lambdaExpressionNode() ensures that only lambda expression nodes can be assigned to a LambdaExpressionNode.
}

// A tree node for a literal expression.
// Use [Node].GetKind to determine the kind of literal.
// For example:
//
//	value
type LiteralNode interface {
	ExpressionNode
	GetValue() string // Returns the value of the literal expression.
	literalNode()     // literalNode() ensures that only literal nodes can be assigned to a LiteralNode.
}

// A tree node for a member reference expression.
// For example:
//
//	expression # [ identifier | new ]
type MemberReferenceNode interface {
	ExpressionNode
	GetMode() ReferenceMode                 // Returns the mode of the reference.
	GetQualifierExpression() ExpressionNode // Returns the qualifier expression for the reference.
	GetName() *string                       // Returns the name of the reference.
	GetTypeArguments() []ExpressionNode     // Returns the type arguments for the reference.
	memberReferenceNode()                   // memberReference() ensures that only member reference nodes can be assigned to a MemberReferenceNode.
}

// A tree node for a member access expression.
// For example:
//
//	expression . identifier
type MemberSelectNode interface {
	ExpressionNode
	GetExpression() ExpressionNode // Returns the expression for which a member is to be selected.
	GetIdentifier() string         // Returns the name of the member to be selected.
	memberSelectNode()             // memberSelectNode() ensures that only member select nodes can be assigned to a MemberSelectNode.
}

// A tree node for a method invocation expression.
// For example:
//
//	identifier ( arguments )
//
//	this . typeArguments identifier ( arguments )
type MethodInvocationNode interface {
	ExpressionNode
	GetTypeArguments() []Node        // Returns the type arguments for this method invocation.
	GetMethodSelect() ExpressionNode // Returns the expression identifying the method to be invoked.
	GetArguments() []ExpressionNode  // Returns the arguments for the method invocation.
	methodInvocationNode()           // methodInvocationNode() ensures that only member invocation nodes can be assigned to a MethodInvocationNode.
}

// A tree node for an expression to create a new instance of an array.
// For example:
//
//	new type dimensions initializers
//
//	new type dimensions [ ] initializers
type NewArrayNode interface {
	ExpressionNode
	GetType() Node                         // Returns the base type of the expression. May be nil for an array initializer expression.
	GetDimensions() []ExpressionNode       // Returns the dimension expressions for the type.
	GetInitializers() []ExpressionNode     // Returns the initializer expressions.
	GetAnnotations() []AnnotationNode      // Returns the annotations on the base type.
	GetDimAnnotations() [][]AnnotationNode // Returns the annotations on each of the dimension expressions.
	newArrayNode()                         // newArrayNode() ensures that only new array nodes can be assigned to a NewArrayNode.
}

// A tree node to declare a new instance of a class.
// For example:
//
//	new identifier ( )
//
//	new identifier ( arguments )
//
//	new typeArguments identifier ( arguments )
//	    classBody
//
//	enclosingExpression.new identifier ( arguments )
type NewClassNode interface {
	ExpressionNode
	GetEnclosingExpression() ExpressionNode // Returns the enclosing expression, or nil if none.
	GetTypeArguments() []Node               // Returns the type arguments for the object being created.
	GetIdentifier() ExpressionNode          // Returns the name of the class being instantiated.
	GetArguments() []ExpressionNode         // Returns the arguments for the constructor to be invoked.
	GetClassBody() ClassNode                // Returns the class body if an anonymous class is being instantiated, and nil otherwise.
	newClassNode()                          // newClassNode() ensures that only new class nodes can be assigned to a NewClassNode.
}

// A tree node for a parenthesized expression.
// Note: parentheses not be preserved by the parser.
// For example:
//
//	( expression )
type ParenthesizedNode interface {
	ExpressionNode
	GetExpression() ExpressionNode // Returns the expression within the parentheses.
	parenthesizedNode()            // parenthesizedNode() ensures that only parenthesized nodes can be assigned to a ParenthesizedNode.
}

// A tree node for a "switch" expression.
// For example:
//
//	switch ( expression ) {
//	    cases
//	}
type SwitchExpressionNode interface {
	ExpressionNode
	GetExpression() ExpressionNode // Returns the expression for the "switch" expression.
	GetCases() []CaseNode          // Returns the cases for the "switch" expression.
	switchExpressionNode()         // switchExpressionNode() ensures that only switch expression nodes can be assigned to a SwitchExpressionNode.
}

// A tree node for a type cast expression.
// For example:
//
//	( type ) expression
type TypeCastNode interface {
	ExpressionNode
	GetType() Node                 // Returns the target type of the cast.
	GetExpression() ExpressionNode // Returns the expression being cast.
	typeCastNode()                 // typeCastNode() ensures that only type cast nodes can be assigned to a TypeCastNode.
}

// A tree node for postfix and unary expressions.
// Use [Node].GetKind to determine the kind of operator.
// For example:
//
//	operator expression
//
//	expression operator
type UnaryNode interface {
	ExpressionNode
	GetExpression() ExpressionNode // Returns the expression that is the operand of the unary operator.
	unaryNode()                    // unaryNode() ensures that only unary nodes can be assigned to an UnaryNode.
}

// A tree node used as the base class for the different kinds of patterns.
type PatternNode interface {
	Node
	CaseLabelNode
	patternNode() // patternNode() ensures that only pattern nodes can be assigned to a PatternNode.
}

// A tree node for a binding pattern.
type BindingPatternNode interface {
	PatternNode
	GetVariable() VariableNode // Returns the binding variable.
	bindingPatternNode()       // bindingPatternNode() ensures that only binding pattern nodes can be assigned to a BindingPatternNode.
}

// A tree node for a guard pattern.
type GuardedPatternNode interface {
	PatternNode
	GetPattern() PatternNode       // The guarded pattern expression.
	GetExpression() ExpressionNode // The guard expression.
	guardedPatternNode()           // guardedPatternNode() ensures that only guarded pattern nodes can be assigned to a GuardedPatternNode.
}

// A tree node for a parenthesized pattern.
// For example:
//
//	( pattern )
type ParenthesizedPatternNode interface {
	PatternNode
	GetPattern() PatternNode   // Returns the pattern within the parentheses.
	parenthesizedPatternNode() // parenthesizedPatternNode() ensures that only parenthesized pattern nodes can be assigned to a ParenthesizedPatternNode.
}

// A tree node for a "case" in a "switch" statement or expression.
// For example:
//
//	case expression :
//	    statements
//
//	default :
//	    statements
type CaseNode interface {
	Node
	GetExpressions() []ExpressionNode // Returns the labels for this case. For default case, returns an empty list.
	GetLabels() []CaseLabelNode       // Returns the labels for this case. For "default" case return a list with a single element, [DefaultCaseLabelNode].
	GetStatements() []StatementNode   // For case with kind [STATEMENT_CASE_KIND], returns the statements labeled by the case. Returns nil for case with kind [RULE_CASE_KIND].
	GetBody() Node                    // For case with kind [RULE_CASE_KIND], returns the statement or expression after the arrow. Returns nil for case with kind [STATEMENT_CASE_KIND].
	GetCaseKind() CaseKind            // Returns the kind of this case.
	caseNode()                        // caseNode() ensures that only case nodes can be assigned to a CaseNode.
}

// A tree node for a "catch" block in a "try" statement.
// For example:
//
//	catch ( parameter )
//	    block
type CatchNode interface {
	Node
	GetParameter() VariableNode // Returns the catch variable. A multi-catch variable will have a [UnionTypeNode] as the type of the variable.
	GetBlock() BlockNode        // Returns the catch block.
	catchNode()                 // catchNode() ensures that only catch nodes can be assigned to a CatchNode.
}

// Represents the abstract syntax tree for ordinary compilation units and modular compilation units.
type CompilationUnitNode interface {
	Node
	GetModule() ModuleNode    // Returns the module tree associated with this compilation unit, or nil if there is no module declaration.
	GetPackage() PackageNode  // Returns the package tree associated with this compilation unit, or nil if there is no package declaration.
	GetImports() []ImportNode // Returns the import declarations appearing in this compilation unit, or an empty list if there are no import declarations.
	GetTypeDecls() []Node     // Returns the type declarations appearing in this compilation unit, or an empty list if there are no type declarations. The list may also include empty statements resulting from extraneous semicolons. A modular compilation unit does not contain any type declarations.
	compilationUnitNode()     // compilationUnitNode() ensures that only compilation unit nodes can be assigned to a CompilationUnitNode.
}

// A super-type for all the directives in a ModuleTree.
type DirectiveNode interface {
	Node
	directiveNode() // directiveNode() ensures that only directive nodes can be assigned to a DirectiveNode.
}

// A tree node for an 'exports' directive in a module declaration.
// For example:
//
//	exports package-name
//
//	exports package-name to module-name
type ExportsNode interface {
	DirectiveNode
	GetPackageName() ExpressionNode   // Returns the name of the package to be exported.
	GetModuleNames() []ExpressionNode // Returns the names of the modules to which the package is exported, or nil, if the package is exported to all modules.
	exportsNode()                     // exportsNode() ensures that only exports nodes can be assigned to an ExportsNode.
}

// A tree node for an 'opens' directive in a module declaration.
// For example:
//
//	opens package-name
//
//	opens package-name to module-name
type OpensNode interface {
	DirectiveNode
	GetPackageName() ExpressionNode   // Returns the name of the package to be opened.
	GetModuleNames() []ExpressionNode // Returns the names of the modules to which the package is opened, or nil, if the package is opened to all modules.
	opensNode()                       // opensNode() ensures that only opens nodes can be assigned to an OpensNode.
}

// A tree node for a 'provides' directive in a module declaration.
// For example:
//
//	provides service-name with implementation-name
type ProvidesNode interface {
	DirectiveNode
	GetServiceName() ExpressionNode           // Returns the name of the service type being provided.
	GetImplementationNames() []ExpressionNode // Returns the names of the implementation types being provided.
	providesNode()                            // providesNode() ensures that only provides nodes can be assigned to a ProvidesNode.
}

// A tree node for a 'requires' directive in a module declaration.
// For example:
//
//	requires module-name
//
//	requires static module-name
//
//	requires transitive module-name
type RequiresNode interface {
	DirectiveNode
	IsStatic() bool                // Returns true if this is a "requires static" directive.
	IsTransitive() bool            // Returns true if this is a "requires transitive" directive.
	GetModuleName() ExpressionNode // Returns the name of the module that is required.
	requiresNode()                 // requiresNode() ensures that only requires nodes can be assigned to a RequiresNode.
}

// A tree node for a 'uses' directive in a module declaration.
// For example:
//
//	uses service-name
type UsesNode interface {
	DirectiveNode
	GetServiceName() ExpressionNode // Returns the name of the service type.
	usesNode()                      // usesNode() ensures that only uses nodes can be assigned to an UsesNode.
}

// A tree node for an import declaration.
// For example:
//
//	import qualifiedIdentifier
//
//	static import qualifiedIdentifier
type ImportNode interface {
	Node
	IsStatic() bool               // Returns true if this is a static import declaration.
	GetQualifiedIdentifier() Node // Returns the qualified identifier for the declaration(s) being imported. If this is an import-on-demand declaration, the qualified identifier will end in "*".
	importNode()                  // importNode() ensures that only import nodes can be assigned to an ImportNode.
}

// A tree node for an intersection type in a cast expression.
type IntersectionTypeNode interface {
	Node
	GetBounds() []Node     // Returns the bounds of the type.
	intersectionTypeNode() // intersectionTypeNode() ensures that only intersection type nodes can be assigned to an IntersectionTypeNode.
}

// A tree node for a method or annotation type element declaration.
// For example:
//
//	modifiers typeParameters type name ( parameters )
//	    body
//
//	modifiers type name () default defaultValue
type MethodNode interface {
	Node
	GetModifiers() ModifiersNode            // Returns the modifiers, including any annotations for the method being declared.
	GetName() string                        // Returns the name of the method being declared.
	GetReturnType() Node                    // Returns the return type of the method being declared. Returns nil for a constructor.
	GetTypeParameters() []TypeParameterNode // Returns the type parameters of the method being declared.
	GetParameters() []VariableNode          // Returns the parameters of the method being declared.
	GetReceiverParameter() VariableNode     // Return an explicit receiver parameter ("this" parameter), or nil if none.
	GetThrows() []ExpressionNode            // Returns the exceptions listed as being thrown by this method.
	GetBody() BlockNode                     // Returns the method body, or nil if this is an abstract or native method.
	GetDefaultValue() Node                  // Returns the default value, if this is an element within an annotation type declaration. Returns nil otherwise.
	methodNode()                            // methodNode() ensures that only method nodes can be assigned to a MethodNode.
}

// A tree node for the modifiers, including annotations, for a declaration.
// For example:
//
//	flags
//
//	flags annotations
type ModifiersNode interface {
	Node
	GetFlags() []Modifier             // Returns the flags in this modifiers tree.
	GetAnnotations() []AnnotationNode // Returns the annotations in this modifiers tree.
	modifiersNode()                   // modifiersNode() ensures that only modifiers nodes can be assigned to a ModifiersNode.
}

// A tree node for a module declaration.
// For example:
//
//	annotations
//	[open] module module-name {
//	    directives
//	}
type ModuleNode interface {
	Node
	GetAnnotations() []AnnotationNode // Returns the annotations associated with this module declaration.
	GetModuleType() ModuleKind        // Returns the type of this module.
	GetName() ExpressionNode          // Returns the name of the module.
	GetDirectives() []DirectiveNode   // Returns the directives in the module declaration.
	moduleNode()                      // moduleNode() ensures that only module nodes can be assigned to a ModuleNode.
}

// Represents the package declaration.
type PackageNode interface {
	Node
	GetAnnotations() []AnnotationNode // Returns the annotations associated with this package declaration.
	GetPackageName() ExpressionNode   // Returns the name of the package being declared.
	packageNode()                     // packageNode() ensures that only package nodes can be assigned to a PackageNode.
}

// A tree node for a type expression involving type parameters.
// For example:
//
//	type < typeArguments >
type ParameterizedTypeNode interface {
	Node
	GetType() Node            // Returns the base type.
	GetTypeArguments() []Node // Returns the type arguments.
	parameterizedTypeNode()   // parameterizedTypeNode() ensures that only parameterized type nodes can be assigned to a ParameterizedTypeNode.
}

// A tree node for a primitive type.
// For example:
//
//	primitiveTypeKind
type PrimitiveTypeNode interface {
	Node
	GetPrimitiveTypeKind() TypeKind // Returns the kind of this primitive type.
	primitiveTypeNode()             // primitiveTypeNode() ensures that only primitive type nodes can be assigned to a PrimitiveTypeNode.
}

// A tree node used as the base class for the different kinds of statements.
type StatementNode interface {
	Node
	statementNode() // statementNode() ensures that only statement nodes can be assigned to a StatementNode.
}

// A tree node for an "assert" statement.
// For example:
//
//	assert condition ;
//
//	assert condition : detail ;
type AssertNode interface {
	StatementNode
	GetCondition() ExpressionNode // Returns the condition being asserted.
	GetDetail() ExpressionNode    // Returns the detail expression.
	assertNode()                  // assertNode() ensures that only assert nodes can be assigned to an AssertNode.
}

// A tree node for a statement block.
// For example:
//
//	{ }
//
//	{ statements }
//
//	static { statements }
type BlockNode interface {
	StatementNode
	IsStatic() bool                 // Returns true if and only if this is a static initializer block.
	GetStatements() []StatementNode // Returns the statements comprising this block.
	blockNode()                     // blockNode() ensures that only block nodes can be assigned to a BlockNode.
}

// A tree node for a "break" statement.
// For example:
//
//	break;
//
//	break label ;
type BreakNode interface {
	StatementNode
	GetLabel() *string // Returns the label for this "break" statement.
	breakNode()        // breakNode() ensures that only break nodes can be assigned to a BreakNode.
}

// A tree node for a class, interface, enum, record, or annotation type declaration.
// For example:
//
//	modifiers class simpleName typeParameters
//	    extends extendsClause
//	    implements implementsClause
//	{
//	    members
//	}
type ClassNode interface {
	StatementNode
	GetModifiers() ModifiersNode            // Returns the modifiers, including any annotations, for this type declaration.
	GetSimpleName() string                  // Returns the simple name of this type declaration.
	GetTypeParameters() []TypeParameterNode // Returns any type parameters of this type declaration.
	GetExtendsClause() Node                 // Returns the supertype of this type declaration, or nil if none is provided.
	GetImplementsClause() []Node            // Returns the interfaces implemented by this type declaration.
	GetPermitsClause() []Node               // Returns the subclasses permitted by this type declaration.
	GetMembers() []Node                     // Returns the members declared in this type declaration.
	classNode()                             // classNode() ensures that only class nodes can be assigned to a ClassNode.
}

// A tree node for a "continue" statement.
// For example:
//
//	continue;
//
//	continue label ;
type ContinueNode interface {
	StatementNode
	GetLabel() *string // Returns the label for this "continue" statement.
	continueNode()     // continueNode() ensures that only continue nodes can be assigned to a ContinueNode.
}

// A tree node for a "do" statement.
// For example:
//
//	do
//	    statement
//	while ( expression );
type DoWhileLoopNode interface {
	StatementNode
	GetCondition() ExpressionNode // Returns the condition of the loop.
	GetStatement() StatementNode  // Returns the body of the loop.
	doWhileLoopNode()             // doWhileLoopNode() ensures that only do-while loop nodes can be assigned to a DoWhileLoopNode.
}

// A tree node for an empty (skip) statement.
// For example:
//
//	;
type EmptyStatementNode interface {
	StatementNode
	emptyStatementNode() // emptyStatementNode() ensures that only empty statement nodes can be assigned to an EmptyStatementNode.
}

// A tree node for an empty expression.
type EmptyExpressionNode interface {
	ExpressionNode
	emptyExpressionNode() // emptyExpressionNode() ensures that only empty expression nodes can be assigned to an EmptyExpressionNode.
}

// A tree node for an "enhanced" "for" loop statement.
// For example:
//
//	for ( variable : expression )
//	    statement
type EnhancedForLoopNode interface {
	StatementNode
	GetVariable() VariableNode     // Returns the control variable for the loop.
	GetExpression() ExpressionNode // Returns the expression yielding the values for the control variable.
	GetStatement() StatementNode   // Returns the body of the loop.
	enhancedForLoopNode()          // enhancedForLoopNode() ensures that only enhanced for-loop nodes can be assigned to an EnhancedForLoopNode.
}

// A tree node for an expression statement.
// For example:
//
//	expression ;
type ExpressionStatementNode interface {
	StatementNode
	GetExpression() ExpressionNode // Returns the expression constituting this statement.
	expressionStatementNode()      // expressionStatementNode() ensures that only expression statement nodes can be assigned to an ExpressionStatementNode.
}

// A tree node for a basic "for" loop statement.
// For example:
//
//	for ( initializer ; condition ; update )
//	    statement
type ForLoopNode interface {
	StatementNode
	GetInitializer() []VariableNode // Returns any initializers of the "for" statement. The result will be an empty list if there are no initializers.
	GetCondition() ExpressionNode   // Returns the condition of the "for" statement. May be nil if there is no condition.
	GetUpdate() []ExpressionNode    // Returns any update expressions of the "for" statement.
	GetStatement() StatementNode    // Returns the body of the "for" statement.
	forLoopNode()                   // forLoopNode() ensures that only for-loop nodes can be assigned to a ForLoopNode.
}

// A tree node for an "if" statement.
// For example:
//
//	if ( condition )
//	    statement
//
//	if ( condition )
//	    thenStatement
//	else
//	    elseStatement
type IfNode interface {
	StatementNode
	GetCondition() ExpressionNode    // Returns the condition of the if-statement.
	GetThenStatement() StatementNode // Returns the statement to be executed if the condition is true.
	GetElseStatement() StatementNode // Returns the statement to be executed if the condition is false, or nil if there is no such statement.
	ifNode()                         // ifNode() ensures that only if nodes can be assigned to an IfNode.
}

// A tree node for a labeled statement.
// For example:
//
//	label : statement
type LabeledStatementNode interface {
	StatementNode
	GetLabel() string            // Returns the label.
	GetStatement() StatementNode // Returns the statement that is labeled.
	labeledStatementNode()       // labeledStatementNode() ensures that only labeled statement nodes can be assigned to a LabeledStatementNode.
}

// A tree node for a "return" statement.
// For example:
//
//	return;
//
//	return expression;
type ReturnNode interface {
	StatementNode
	GetExpression() ExpressionNode // Returns the expression to be returned.
	returnNode()                   // returnNode() ensures that only return nodes can be assigned to a ReturnNode.
}

// A tree node for a "switch" statement.
// For example:
//
//	switch ( expression ) {
//	    cases
//	}
type SwitchNode interface {
	StatementNode
	GetExpression() ExpressionNode // Returns the expression for the "switch" statement.
	GetCases() []CaseNode          // Returns the cases for the "switch" statement.
	switchNode()                   // switchNode() ensures that only switch nodes can be assigned to a SwitchNode.
}

// A tree node for a "synchronized" statement.
// For example:
//
//	synchronized ( expression )
//	    block
type SynchronizedNode interface {
	StatementNode
	GetExpression() ExpressionNode // Returns the expression on which to synchronize.
	GetBlock() BlockNode           // Returns the block of the "synchronized" statement.
	synchronizedNode()             // synchronizedNode() ensures that only synchronized nodes can be assigned to a SynchronizedNode.
}

// A tree node for a "throw" statement.
// For example:
//
//	throw expression;
type ThrowNode interface {
	StatementNode
	GetExpression() ExpressionNode // Returns the expression to be thrown.
	throwNode()                    // throwNode() ensures that only throw nodes can be assigned to a ThrowNode.
}

// A tree node for a "try" statement.
// For example:
//
//	try
//	    block
//	catches
//	finally
//	    finallyBlock
type TryNode interface {
	StatementNode
	GetBlock() BlockNode        // Returns the block of the "try" statement.
	GetCatches() []CatchNode    // Returns any catch blocks provided in the "try" statement. The result will be an empty list if there are no catch blocks.
	GetFinallyBlock() BlockNode // Returns the finally block provided in the "try" statement, or nil if there is none.
	GetResources() []Node       // Returns any resource declarations provided in the "try" statement. The result will be an empty list if there are no resource declarations.
	tryNode()                   // tryNode() ensures that only try nodes can be assigned to a TryNode.
}

// A tree node for a variable declaration.
// For example:
//
//	modifiers type name initializer ;
//
//	modifiers type qualified-name.this
type VariableNode interface {
	StatementNode
	GetModifiers() ModifiersNode       // Returns the modifiers, including any annotations, on the declaration.
	GetName() string                   // Returns the name of the variable being declared.
	GetNameExpression() ExpressionNode // Returns the qualified identifier for the name being "declared". This is only used in certain cases for the receiver of a method declaration. Returns nil in all other cases.
	GetType() Node                     // Returns the type of the variable being declared.
	GetInitializer() ExpressionNode    // Returns the initializer for the variable, or nil if none.
	variableNode()                     // variableNode() ensures that only variable nodes can be assigned to a VariableNode.
}

// A tree node for a "while" loop statement.
// For example:
//
//	while ( condition )
//	    statement
type WhileLoopNode interface {
	StatementNode
	GetCondition() ExpressionNode // Returns the condition of the loop.
	GetStatement() StatementNode  // Returns the body of the loop.
	whileLoopNode()               // whileLoopNode() ensures that only while loop nodes can be assigned to a WhileLoopNode.
}

// A tree node for a "yield" statement.
// For example:
//
//	yield expression ;
type YieldNode interface {
	StatementNode
	GetValue() ExpressionNode // Returns the expression for this "yield" statement.
	yieldNode()               // yieldNode() ensures that only yield nodes can be assigned to an YieldNode.
}

// A tree node for a type parameter.
// For example:
//
//	name
//
//	name extends bounds
//
//	annotations name
type TypeParameterNode interface {
	Node
	GetName() string                  // Returns the name of the type parameter.
	GetBounds() []Node                // Returns the bounds of the type parameter.
	GetAnnotations() []AnnotationNode // Returns annotations on the type parameter declaration.
	typeParameterNode()               // typeParameterNode() ensures that only type parameter nodes can be assigned to a TypeParameterNode.
}

// A tree node for a union type expression in a multicatch variable declaration.
type UnionTypeNode interface {
	Node
	GetTypeAlternatives() []Node // Returns the alternative type expressions.
	unionTypeNode()              // unionTypeNode() ensures that only union type nodes can be assigned to an UnionTypeNode.
}

// A tree node for a wildcard type argument.
// Use [Node].GetKind to determine the kind of bound.
// For example:
//
//	?
//
//	? extends bound
//
//	? super bound
type WildcardNode interface {
	Node
	GetBound() Node // Returns the bound of the wildcard.
	wildcardNode()  // wildcardNode() ensures that only wildcard nodes can be assigned to a WildcardNode.
}
