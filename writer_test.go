package javast_test

import (
	"errors"
	"testing"

	"github.com/kapavkin/javast"
)

type SpaceWriter struct {
	buf []byte
}

func (sw *SpaceWriter) Write(p []byte) (n int, err error) {
	if sw != nil {
		if len(sw.buf) > 0 {
			sw.buf = append(sw.buf, ' ')
			n += 1
		}
		sw.buf = append(sw.buf, p...)
		n += len(p)
		return
	}
	err = errors.New("sw is nil")
	return
}

func (sw SpaceWriter) String() string {
	return string(sw.buf)
}

func TestAnnotatedType_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	at := javast.AnnotatedType{
		Annotations: []javast.AnnotationNode{
			javast.TypeAnnotation{
				AnnotationType: javast.Identifier{
					Name: "annotationType",
				},
				Arguments: []javast.ExpressionNode{
					&javast.Identifier{
						Name: "arguments",
					},
				},
			},
		},
		UnderlyingType: &javast.Identifier{
			Name: "Date",
		},
	}
	if _, err := at.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "@ annotationType ( arguments ) Date"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestAnnotation_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	a := javast.Annotation{
		AnnotationType: javast.Identifier{
			Name: "annotationType",
		},
		Arguments: []javast.ExpressionNode{
			&javast.Identifier{
				Name: "arguments",
			},
		},
	}
	if _, err := a.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "@ annotationType ( arguments )"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestTypeAnnotation_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	ta := javast.TypeAnnotation{
		AnnotationType: javast.Identifier{
			Name: "annotationType",
		},
		Arguments: []javast.ExpressionNode{
			&javast.Identifier{
				Name: "arguments",
			},
		},
	}
	if _, err := ta.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "@ annotationType ( arguments )"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestArrayAccess_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	aa := javast.ArrayAccess{
		Expression: javast.Identifier{
			Name: "array",
		},
		Index: javast.Identifier{
			Name: "i",
		},
	}
	if _, err := aa.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "array [ i ]"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestArrayType_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	at := javast.ArrayType{
		Type: javast.PrimitiveType{
			PrimitiveTypeKind: javast.INT_TYPE_KIND,
		},
	}
	if _, err := at.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "int []"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestAssert_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	a := javast.Assert{
		Condition: javast.GreaterThan{
			LeftOperand: javast.Identifier{
				Name: "a",
			},
			RightOperand: javast.Identifier{
				Name: "b",
			},
		},
		Detail: nil,
	}
	if _, err := a.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "assert a > b ;"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestAssignment_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	a := javast.Assignment{
		Variable: javast.Identifier{
			Name: "var",
		},
		Expression: javast.IntLiteral{
			Value: "1000L",
		},
	}
	if _, err := a.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "var = 1000L"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestBlock_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	b := javast.Block{
		Static: false,
		Statements: []javast.StatementNode{
			javast.ExpressionStatement{
				Expression: javast.MethodInvocation{
					TypeArguments: nil,
					MethodSelect: javast.MemberSelect{
						Expression: javast.Identifier{
							Name: "writer",
						},
						Identifier: "write",
					},
					Arguments: []javast.ExpressionNode{
						javast.Identifier{
							Name: "arguments",
						},
					},
				},
			},
		},
	}
	if _, err := b.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "{ writer . write ( arguments ) ; }"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestBreak_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	label := "label"
	b := javast.Break{
		Label: &label,
	}
	if _, err := b.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "break label ;"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestStatementCase_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	sc := javast.StatementCase{
		Expression: javast.Identifier{
			Name: "CONST",
		},
		Statements: []javast.StatementNode{
			javast.ExpressionStatement{
				Expression: javast.Assignment{
					Variable: javast.Identifier{
						Name: "a",
					},
					Expression: javast.Identifier{
						Name: "b",
					},
				},
			},
			javast.Break{
				Label: nil,
			},
		},
	}
	if _, err := sc.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "case CONST : a = b ; break ;"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestRuleCase_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	rc := javast.RuleCase{
		Labels: []javast.CaseLabelNode{
			javast.Identifier{
				Name: "A",
			},
			javast.Identifier{
				Name: "B",
			},
		},
		Body: javast.ExpressionStatement{
			Expression: javast.Assignment{
				Variable: javast.Identifier{
					Name: "a",
				},
				Expression: javast.Identifier{
					Name: "b",
				},
			},
		},
	}
	if _, err := rc.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "case A , B -> a = b ;"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestCatch_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	c := javast.Catch{
		Parameter: javast.Variable{
			Modifiers: javast.Modifiers{
				Flags:       nil,
				Annotations: nil,
			},
			Name:           "e",
			NameExpression: nil,
			Type: javast.Identifier{
				Name: "IllegalStateException",
			},
			Initializer: nil,
		},
		Block: javast.Block{
			Static: false,
			Statements: []javast.StatementNode{
				javast.ExpressionStatement{
					Expression: javast.Assignment{
						Variable: javast.Identifier{
							Name: "a",
						},
						Expression: javast.Identifier{
							Name: "b",
						},
					},
				},
			},
		},
	}
	if _, err := c.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "catch ( IllegalStateException e ) { a = b ; }"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestClass_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	c := javast.Class{
		Modifiers: javast.Modifiers{
			Flags: []javast.Modifier{
				javast.PUBLIC_MODIFIER,
				javast.FINAL_MODIFIER,
			},
			Annotations: nil,
		},
		SimpleName: "PrettyPrinter",
		TypeParameters: []javast.TypeParameterNode{
			javast.TypeParameter{
				Name: "T",
				Bounds: []javast.Node{
					javast.Identifier{
						Name: "Printable",
					},
				},
				Annotations: nil,
			},
		},
		ExtendsClause: javast.Identifier{
			Name: "ConsolePrinter",
		},
		ImplementsClause: []javast.Node{
			javast.Identifier{
				Name: "Printer",
			},
		},
		Members: nil,
	}
	if _, err := c.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "public final class PrettyPrinter < T extends Printable > extends ConsolePrinter implements Printer { }"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestCompilationUnit_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	cu := javast.CompilationUnit{
		Module: nil,
		Package: javast.Package{
			Annotations: nil,
			PackageName: javast.Identifier{
				Name: "printer",
			},
		},
		Imports: nil,
		TypeDecls: []javast.Node{
			javast.Class{
				Modifiers: javast.Modifiers{
					Flags: []javast.Modifier{
						javast.PUBLIC_MODIFIER,
						javast.FINAL_MODIFIER,
					},
					Annotations: nil,
				},
				SimpleName: "PrettyPrinter",
				TypeParameters: []javast.TypeParameterNode{
					javast.TypeParameter{
						Name: "T",
						Bounds: []javast.Node{
							javast.Identifier{
								Name: "Printable",
							},
						},
						Annotations: nil,
					},
				},
				ExtendsClause: javast.Identifier{
					Name: "ConsolePrinter",
				},
				ImplementsClause: []javast.Node{
					javast.Identifier{
						Name: "Printer",
					},
				},
				Members: nil,
			},
		},
	}
	if _, err := cu.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "package printer ; public final class PrettyPrinter < T extends Printable > extends ConsolePrinter implements Printer { }"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestConditionalExpression_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	cx := javast.ConditionalExpression{
		Condition: javast.GreaterThanEqual{
			LeftOperand: javast.Identifier{
				Name: "a",
			},
			RightOperand: javast.Identifier{
				Name: "b",
			},
		},
		TrueExpression: javast.Identifier{
			Name: "a",
		},
		FalseExpression: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := cx.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a >= b ? a : b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestContinue_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	label := "label"
	c := javast.Continue{
		Label: &label,
	}
	if _, err := c.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "continue label ;"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestDoWhileLoop_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	dwl := javast.DoWhileLoop{
		Condition: javast.LessThan{
			LeftOperand: javast.Identifier{
				Name: "a",
			},
			RightOperand: javast.Identifier{
				Name: "b",
			},
		},
		Statement: javast.Block{
			Static: false,
			Statements: []javast.StatementNode{
				javast.ExpressionStatement{
					Expression: javast.PlusAssignment{
						Variable: javast.Identifier{
							Name: "a",
						},
						Expression: javast.Identifier{
							Name: "c",
						},
					},
				},
				javast.ExpressionStatement{
					Expression: javast.PostfixIncrement{
						Expression: javast.Identifier{
							Name: "cnt",
						},
					},
				},
			},
		},
	}
	if _, err := dwl.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "do { a += c ; cnt ++ ; } while ( a < b ) ;"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestEnhancedForLoop_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	efl := javast.EnhancedForLoop{
		Variable: javast.Variable{
			Modifiers: javast.Modifiers{
				Flags:       nil,
				Annotations: nil,
			},
			Name:           "elem",
			NameExpression: nil,
			Type: javast.Identifier{
				Name: "Elem",
			},
			Initializer: nil,
		},
		Expression: javast.Identifier{
			Name: "elements",
		},
		Statement: javast.ExpressionStatement{
			Expression: javast.PlusAssignment{
				Variable: javast.Identifier{
					Name: "buf",
				},
				Expression: javast.Identifier{
					Name: "elem",
				},
			},
		},
	}
	if _, err := efl.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "for ( Elem elem : elements ) buf += elem ;"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestExpressionStatement_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	xs := javast.ExpressionStatement{
		Expression: javast.PlusAssignment{
			Variable: javast.Identifier{
				Name: "buf",
			},
			Expression: javast.Identifier{
				Name: "elem",
			},
		},
	}
	if _, err := xs.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "buf += elem ;"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestMemberSelect_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	ms := javast.MemberSelect{
		Expression: javast.MemberSelect{
			Expression: javast.Identifier{
				Name: "www",
			},
			Identifier: "google",
		},
		Identifier: "com",
	}
	if _, err := ms.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "www . google . com"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestInvokeMemberReference_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	imr := javast.InvokeMemberReference{
		QualifierExpression: javast.Identifier{
			Name: "String",
		},
		Name:          "replace",
		TypeArguments: nil,
	}
	if _, err := imr.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "String :: replace"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestNewMemberReference_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	nmr := javast.NewMemberReference{
		QualifierExpression: javast.Identifier{
			Name: "String",
		},
		TypeArguments: nil,
	}
	if _, err := nmr.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "String :: new"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestForLoop_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	fl := javast.ForLoop{
		Initializer: []javast.VariableNode{
			javast.Variable{
				Modifiers: javast.Modifiers{
					Flags:       nil,
					Annotations: nil,
				},
				Name:           "i",
				NameExpression: nil,
				Type: javast.PrimitiveType{
					PrimitiveTypeKind: javast.INT_TYPE_KIND,
				},
				Initializer: javast.IntLiteral{
					Value: "0",
				},
			},
		},
		Condition: javast.LessThan{
			LeftOperand: javast.Identifier{
				Name: "i",
			},
			RightOperand: javast.Identifier{
				Name: "n",
			},
		},
		Update: []javast.ExpressionNode{
			javast.PostfixIncrement{
				Expression: javast.Identifier{
					Name: "i",
				},
			},
		},
		Statement: javast.ExpressionStatement{
			Expression: javast.MethodInvocation{
				TypeArguments: nil,
				MethodSelect: javast.MemberSelect{
					Expression: javast.Identifier{
						Name: "buf",
					},
					Identifier: "addOne",
				},
				Arguments: []javast.ExpressionNode{
					javast.ArrayAccess{
						Expression: javast.Identifier{
							Name: "array",
						},
						Index: javast.Identifier{
							Name: "i",
						},
					},
				},
			},
		},
	}
	if _, err := fl.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "for ( int i = 0 ; i < n ; i ++ ) buf . addOne ( array [ i ] ) ;"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestIdentifier_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	i := javast.Identifier{
		Name: "name",
	}
	if _, err := i.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "name"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestIf_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	i := javast.If{
		Condition: javast.GreaterThan{
			LeftOperand: javast.Identifier{
				Name: "a",
			},
			RightOperand: javast.Identifier{
				Name: "b",
			},
		},
		ThenStatement: javast.ExpressionStatement{
			Expression: javast.Assignment{
				Variable: javast.Identifier{
					Name: "max",
				},
				Expression: javast.Identifier{
					Name: "a",
				},
			},
		},
		ElseStatement: javast.ExpressionStatement{
			Expression: javast.Assignment{
				Variable: javast.Identifier{
					Name: "max",
				},
				Expression: javast.Identifier{
					Name: "b",
				},
			},
		},
	}
	if _, err := i.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "if ( a > b ) max = a ; else max = b ;"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestImport_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	i := javast.Import{
		Static: false,
		QualifiedIdentifier: javast.MemberSelect{
			Expression: javast.MemberSelect{
				Expression: javast.MemberSelect{
					Expression: javast.MemberSelect{
						Expression: javast.Identifier{
							Name: "com",
						},
						Identifier: "sun",
					},
					Identifier: "source",
				},
				Identifier: "tree",
			},
			Identifier: "Tree",
		},
	}
	if _, err := i.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "import com . sun . source . tree . Tree ;"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestInstanceOf_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	io := javast.InstanceOf{
		Expression: javast.Identifier{
			Name: "g",
		},
		Type: javast.Identifier{
			Name: "String",
		},
		Pattern: nil,
	}
	if _, err := io.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "g instanceof String"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestLabeledStatement_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	label := "start"
	ls := javast.LabeledStatement{
		Label: label,
		Statement: javast.DoWhileLoop{
			Condition: javast.LessThan{
				LeftOperand: javast.Identifier{
					Name: "i",
				},
				RightOperand: javast.Identifier{
					Name: "n",
				},
			},
			Statement: javast.DoWhileLoop{
				Condition: javast.LessThan{
					LeftOperand: javast.Identifier{
						Name: "j",
					},
					RightOperand: javast.Identifier{
						Name: "n",
					},
				},
				Statement: javast.Block{
					Static: false,
					Statements: []javast.StatementNode{
						javast.ExpressionStatement{
							Expression: javast.PlusAssignment{
								Variable: javast.Identifier{
									Name: "c",
								},
								Expression: javast.Multiply{
									LeftOperand: javast.Identifier{
										Name: "i",
									},
									RightOperand: javast.Identifier{
										Name: "j",
									},
								},
							},
						},
						javast.If{
							Condition: javast.EqualTo{
								LeftOperand: javast.Identifier{
									Name: "i",
								},
								RightOperand: javast.Identifier{
									Name: "j",
								},
							},
							ThenStatement: javast.Break{
								Label: &label,
							},
							ElseStatement: nil,
						},
					},
				},
			},
		},
	}
	if _, err := ls.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "start : do do { c += i * j ; if ( i == j ) break start ; } while ( j < n ) ; while ( i < n ) ;"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestMethod_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	m := javast.Method{
		Modifiers: javast.Modifiers{
			Flags: []javast.Modifier{
				javast.PUBLIC_MODIFIER,
				javast.FINAL_MODIFIER,
			},
			Annotations: nil,
		},
		Name: "collect",
		ReturnType: javast.Identifier{
			Name: "void",
		},
		TypeParameters: nil,
		Parameters: []javast.VariableNode{
			javast.Variable{
				Modifiers: javast.Modifiers{
					Flags:       nil,
					Annotations: nil,
				},
				Name:           "elements",
				NameExpression: nil,
				Type: javast.ArrayType{
					Type: javast.PrimitiveType{
						PrimitiveTypeKind: javast.INT_TYPE_KIND,
					},
				},
				Initializer: nil,
			},
		},
		ReceiverParameter: nil,
		Throws: []javast.ExpressionNode{
			javast.Identifier{
				Name: "IllegalStateException",
			},
		},
		Body: javast.Block{
			Static: false,
			Statements: []javast.StatementNode{
				javast.ExpressionStatement{
					Expression: javast.MethodInvocation{
						TypeArguments: nil,
						MethodSelect: javast.MemberSelect{
							Expression: javast.Identifier{
								Name: "collector",
							},
							Identifier: "collect",
						},
						Arguments: []javast.ExpressionNode{
							javast.Identifier{
								Name: "elements",
							},
						},
					},
				},
			},
		},
		DefaultValue: nil,
	}
	if _, err := m.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "public final void collect ( int [] elements ) throws IllegalStateException { collector . collect ( elements ) ; }"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestMethodInvocation_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	mi := javast.MethodInvocation{
		TypeArguments: nil,
		MethodSelect: javast.MemberSelect{
			Expression: javast.Identifier{
				Name: "collector",
			},
			Identifier: "collect",
		},
		Arguments: []javast.ExpressionNode{
			javast.Identifier{
				Name: "elements",
			},
		},
	}
	if _, err := mi.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "collector . collect ( elements )"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestModifiers_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	m := javast.Modifiers{
		Flags: []javast.Modifier{
			javast.PUBLIC_MODIFIER,
			javast.FINAL_MODIFIER,
		},
		Annotations: nil,
	}
	if _, err := m.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "public final"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestNewArray_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	na := javast.NewArray{
		Type: javast.PrimitiveType{
			PrimitiveTypeKind: javast.INT_TYPE_KIND,
		},
		Dimensions: []javast.ExpressionNode{
			javast.IntLiteral{
				Value: "5",
			},
		},
		Initializers: []javast.ExpressionNode{
			javast.IntLiteral{
				Value: "1",
			},
			javast.IntLiteral{
				Value: "2",
			},
			javast.IntLiteral{
				Value: "3",
			},
			javast.IntLiteral{
				Value: "4",
			},
			javast.IntLiteral{
				Value: "5",
			},
		},
		Annotations:    nil,
		DimAnnotations: nil,
	}
	if _, err := na.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "new int [ 5 ] { 1 , 2 , 3 , 4 , 5 }"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestNewClass_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	nc := javast.NewClass{
		EnclosingExpression: nil,
		TypeArguments:       nil,
		Identifier: javast.Identifier{
			Name: "String",
		},
		Arguments: []javast.ExpressionNode{
			javast.Identifier{
				Name: "bytes",
			},
		},
		ClassBody: nil,
	}
	if _, err := nc.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "new String ( bytes )"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestExpressionLambdaExpression_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	xlx := javast.ExpressionLambdaExpression{
		Parameters: []javast.VariableNode{
			javast.Variable{
				Modifiers: javast.Modifiers{
					Flags:       nil,
					Annotations: nil,
				},
				Name:           "a",
				NameExpression: nil,
				Type: javast.PrimitiveType{
					PrimitiveTypeKind: javast.INT_TYPE_KIND,
				},
				Initializer: nil,
			},
			javast.Variable{
				Modifiers: javast.Modifiers{
					Flags:       nil,
					Annotations: nil,
				},
				Name:           "b",
				NameExpression: nil,
				Type: javast.PrimitiveType{
					PrimitiveTypeKind: javast.INT_TYPE_KIND,
				},
				Initializer: nil,
			},
		},
		Expression: javast.GreaterThan{
			LeftOperand: javast.Identifier{
				Name: "a",
			},
			RightOperand: javast.Identifier{
				Name: "b",
			},
		},
	}
	if _, err := xlx.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "( int a , int b ) -> a > b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestStatementLambdaExpression_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	slx := javast.StatementLambdaExpression{
		Parameters: []javast.VariableNode{
			javast.Variable{
				Modifiers: javast.Modifiers{
					Flags:       nil,
					Annotations: nil,
				},
				Name:           "a",
				NameExpression: nil,
				Type: javast.PrimitiveType{
					PrimitiveTypeKind: javast.INT_TYPE_KIND,
				},
				Initializer: nil,
			},
			javast.Variable{
				Modifiers: javast.Modifiers{
					Flags:       nil,
					Annotations: nil,
				},
				Name:           "b",
				NameExpression: nil,
				Type: javast.PrimitiveType{
					PrimitiveTypeKind: javast.INT_TYPE_KIND,
				},
				Initializer: nil,
			},
		},
		Block: javast.Block{
			Static: false,
			Statements: []javast.StatementNode{
				javast.ExpressionStatement{
					Expression: javast.GreaterThan{
						LeftOperand: javast.Identifier{
							Name: "a",
						},
						RightOperand: javast.Identifier{
							Name: "b",
						},
					},
				},
			},
		},
	}
	if _, err := slx.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "( int a , int b ) -> { a > b ; }"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestPackage_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	p := javast.Package{
		Annotations: nil,
		PackageName: javast.Identifier{
			Name: "printer",
		},
	}
	if _, err := p.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "package printer ;"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestParenthesized_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	p := javast.Parenthesized{
		Expression: javast.Plus{
			LeftOperand: javast.Identifier{
				Name: "a",
			},
			RightOperand: javast.Identifier{
				Name: "b",
			},
		},
	}
	if _, err := p.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "( a + b )"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestBindingPattern_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	bp := javast.BindingPattern{
		Variable: javast.Variable{
			Modifiers: javast.Modifiers{
				Flags:       nil,
				Annotations: nil,
			},
			Name:           "a",
			NameExpression: nil,
			Type: javast.PrimitiveType{
				PrimitiveTypeKind: javast.INT_TYPE_KIND,
			},
			Initializer: nil,
		},
	}
	if _, err := bp.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "int a"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestGuardedPattern_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	gp := javast.GuardedPattern{
		Pattern: javast.BindingPattern{
			Variable: javast.Variable{
				Modifiers: javast.Modifiers{
					Flags:       nil,
					Annotations: nil,
				},
				Name:           "a",
				NameExpression: nil,
				Type: javast.PrimitiveType{
					PrimitiveTypeKind: javast.INT_TYPE_KIND,
				},
				Initializer: nil,
			},
		},
		Expression: javast.GreaterThan{
			LeftOperand: javast.Identifier{
				Name: "a",
			},
			RightOperand: javast.IntLiteral{
				Value: "0",
			},
		},
	}
	if _, err := gp.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "int a & a > 0"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestParenthesizedPattern_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	pp := javast.ParenthesizedPattern{
		Pattern: javast.BindingPattern{
			Variable: javast.Variable{
				Modifiers: javast.Modifiers{
					Flags:       nil,
					Annotations: nil,
				},
				Name:           "a",
				NameExpression: nil,
				Type: javast.PrimitiveType{
					PrimitiveTypeKind: javast.INT_TYPE_KIND,
				},
				Initializer: nil,
			},
		},
	}
	if _, err := pp.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "( int a )"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestDefaultCaseLabel_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	dcl := javast.DefaultCaseLabel{}
	if _, err := dcl.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "default"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestPrimitiveType_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	pt := javast.PrimitiveType{
		PrimitiveTypeKind: javast.INT_TYPE_KIND,
	}
	if _, err := pt.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "int"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestReturn_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	r := javast.Return{
		Expression: javast.Identifier{
			Name: "n",
		},
	}
	if _, err := r.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "return n ;"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestEmptyStatement_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	es := javast.EmptyStatement{}
	if _, err := es.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := ";"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestSwitch_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	s := javast.Switch{
		Expression: javast.Identifier{
			Name: "n",
		},
		Cases: []javast.CaseNode{
			javast.StatementCase{
				Expression: javast.IntLiteral{
					Value: "1",
				},
				Statements: []javast.StatementNode{
					javast.ExpressionStatement{
						Expression: javast.Assignment{
							Variable: javast.Identifier{
								Name: "a",
							},
							Expression: javast.IntLiteral{
								Value: "5",
							},
						},
					},
					javast.Break{
						Label: nil,
					},
				},
			},
			javast.StatementCase{
				Expression: nil,
				Statements: []javast.StatementNode{
					javast.ExpressionStatement{
						Expression: javast.Assignment{
							Variable: javast.Identifier{
								Name: "a",
							},
							Expression: javast.Identifier{
								Name: "b",
							},
						},
					},
				},
			},
		},
	}
	if _, err := s.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "switch ( n ) { case 1 : a = 5 ; break ; default : a = b ; }"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestSwitchExpression_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	sx := javast.SwitchExpression{
		Expression: javast.Identifier{
			Name: "n",
		},
		Cases: []javast.CaseNode{
			javast.StatementCase{
				Expression: javast.IntLiteral{
					Value: "1",
				},
				Statements: []javast.StatementNode{
					javast.ExpressionStatement{
						Expression: javast.Assignment{
							Variable: javast.Identifier{
								Name: "a",
							},
							Expression: javast.IntLiteral{
								Value: "5",
							},
						},
					},
					javast.Break{
						Label: nil,
					},
				},
			},
			javast.StatementCase{
				Expression: nil,
				Statements: []javast.StatementNode{
					javast.ExpressionStatement{
						Expression: javast.Assignment{
							Variable: javast.Identifier{
								Name: "a",
							},
							Expression: javast.Identifier{
								Name: "b",
							},
						},
					},
				},
			},
		},
	}
	if _, err := sx.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "switch ( n ) { case 1 : a = 5 ; break ; default : a = b ; }"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestSynchronized_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	s := javast.Synchronized{
		Expression: javast.Identifier{
			Name: "obj",
		},
		Block: javast.Block{
			Static: false,
			Statements: []javast.StatementNode{
				javast.ExpressionStatement{
					Expression: javast.PlusAssignment{
						Variable: javast.MemberSelect{
							Expression: javast.Identifier{
								Name: "obj",
							},
							Identifier: "a",
						},
						Expression: javast.Identifier{
							Name: "b",
						},
					},
				},
			},
		},
	}
	if _, err := s.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "synchronized ( obj ) { obj . a += b ; }"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestThrow_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	tw := javast.Throw{
		Expression: javast.NewClass{
			EnclosingExpression: nil,
			TypeArguments:       nil,
			Identifier: javast.Identifier{
				Name: "IllegalStateException",
			},
			Arguments: []javast.ExpressionNode{
				javast.StringLiteral{
					Value: "couln't do something",
				},
			},
			ClassBody: nil,
		},
	}
	if _, err := tw.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "throw new IllegalStateException ( \"couln't do something\" ) ;"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestTry_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	try := javast.Try{
		Block: javast.Block{
			Static: false,
			Statements: []javast.StatementNode{
				javast.ExpressionStatement{
					Expression: javast.Assignment{
						Variable: javast.Identifier{
							Name: "a",
						},
						Expression: javast.Identifier{
							Name: "b",
						},
					},
				},
			},
		},
		Catches: []javast.CatchNode{
			javast.Catch{
				Parameter: javast.Variable{
					Modifiers: javast.Modifiers{
						Flags:       nil,
						Annotations: nil,
					},
					Name:           "e",
					NameExpression: nil,
					Type: javast.Identifier{
						Name: "IllegalStateException",
					},
					Initializer: nil,
				},
				Block: javast.Block{
					Static: false,
					Statements: []javast.StatementNode{
						javast.ExpressionStatement{
							Expression: javast.Assignment{
								Variable: javast.Identifier{
									Name: "a",
								},
								Expression: javast.Identifier{
									Name: "c",
								},
							},
						},
					},
				},
			},
		},
		FinallyBlock: javast.Block{
			Static: false,
			Statements: []javast.StatementNode{
				javast.ExpressionStatement{
					Expression: javast.MethodInvocation{
						TypeArguments: nil,
						MethodSelect: javast.MemberSelect{
							Expression: javast.Identifier{
								Name: "resource",
							},
							Identifier: "close",
						},
						Arguments: nil,
					},
				},
			},
		},
		Resources: nil,
	}
	if _, err := try.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "try { a = b ; } catch ( IllegalStateException e ) { a = c ; } finally { resource . close ( ) ; }"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestParameterizedType_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	pt := javast.ParameterizedType{
		Type: javast.Identifier{
			Name: "Map",
		},
		TypeArguments: []javast.Node{
			javast.Identifier{
				Name: "String",
			},
			javast.Identifier{
				Name: "String",
			},
		},
	}
	if _, err := pt.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "Map < String , String >"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestUnionType_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	ut := javast.UnionType{
		TypeAlternatives: []javast.Node{
			javast.PrimitiveType{
				PrimitiveTypeKind: javast.FLOAT_TYPE_KIND,
			},
			javast.PrimitiveType{
				PrimitiveTypeKind: javast.DOUBLE_TYPE_KIND,
			},
		},
	}
	if _, err := ut.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "float | double"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestIntersectionType_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	it := javast.IntersectionType{
		Bounds: []javast.Node{
			javast.Identifier{
				Name: "Reader",
			},
			javast.Identifier{
				Name: "Writer",
			},
		},
	}
	if _, err := it.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "Reader & Writer"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestTypeCast_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	tc := javast.TypeCast{
		Type: javast.PrimitiveType{
			PrimitiveTypeKind: javast.INT_TYPE_KIND,
		},
		Expression: javast.Identifier{
			Name: "count",
		},
	}
	if _, err := tc.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "( int ) count"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestTypeParameter_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	tp := javast.TypeParameter{
		Name: "W",
		Bounds: []javast.Node{
			javast.Identifier{
				Name: "Writer",
			},
		},
		Annotations: nil,
	}
	if _, err := tp.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "W extends Writer"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestVariable_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	v := javast.Variable{
		Modifiers: javast.Modifiers{
			Flags: []javast.Modifier{
				javast.PUBLIC_MODIFIER,
				javast.STATIC_MODIFIER,
			},
			Annotations: nil,
		},
		Name:           "HOST",
		NameExpression: nil,
		Type: javast.Identifier{
			Name: "String",
		},
		Initializer: javast.StringLiteral{
			Value: "google.com",
		},
	}
	if _, err := v.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "public static String HOST = \"google.com\" ;"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestWhileLoop_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	wl := javast.WhileLoop{
		Condition: javast.LessThan{
			LeftOperand: javast.Identifier{
				Name: "i",
			},
			RightOperand: javast.Identifier{
				Name: "n",
			},
		},
		Statement: javast.ExpressionStatement{
			Expression: javast.PostfixIncrement{
				Expression: javast.Identifier{
					Name: "i",
				},
			},
		},
	}
	if _, err := wl.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "while ( i < n ) i ++ ;"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestPostfixIncrement_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	pi := javast.PostfixIncrement{
		Expression: javast.Identifier{
			Name: "i",
		},
	}
	if _, err := pi.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "i ++"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestPostfixDecrement_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	pd := javast.PostfixDecrement{
		Expression: javast.Identifier{
			Name: "i",
		},
	}
	if _, err := pd.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "i --"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestPrefixIncrement_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	pi := javast.PrefixIncrement{
		Expression: javast.Identifier{
			Name: "i",
		},
	}
	if _, err := pi.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "++ i"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestPrefixDecrement_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	pd := javast.PrefixDecrement{
		Expression: javast.Identifier{
			Name: "i",
		},
	}
	if _, err := pd.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "-- i"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestUnaryPlus_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	up := javast.UnaryPlus{
		Expression: javast.Identifier{
			Name: "n",
		},
	}
	if _, err := up.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "+ n"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestUnaryMinus_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	um := javast.UnaryMinus{
		Expression: javast.Identifier{
			Name: "n",
		},
	}
	if _, err := um.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "- n"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestBitwiseComplement_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	bc := javast.BitwiseComplement{
		Expression: javast.Identifier{
			Name: "n",
		},
	}
	if _, err := bc.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "~ n"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestLogicalComplement_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	lc := javast.LogicalComplement{
		Expression: javast.Identifier{
			Name: "n",
		},
	}
	if _, err := lc.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "! n"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestMultiply_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	m := javast.Multiply{
		LeftOperand: javast.Identifier{
			Name: "a",
		},
		RightOperand: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := m.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a * b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestDivide_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	d := javast.Divide{
		LeftOperand: javast.Identifier{
			Name: "a",
		},
		RightOperand: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := d.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a / b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestRemainder_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	r := javast.Remainder{
		LeftOperand: javast.Identifier{
			Name: "a",
		},
		RightOperand: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := r.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a % b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestPlus_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	p := javast.Plus{
		LeftOperand: javast.Identifier{
			Name: "a",
		},
		RightOperand: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := p.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a + b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestMinus_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	m := javast.Minus{
		LeftOperand: javast.Identifier{
			Name: "a",
		},
		RightOperand: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := m.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a - b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestLeftShift_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	ls := javast.LeftShift{
		LeftOperand: javast.Identifier{
			Name: "a",
		},
		RightOperand: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := ls.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a << b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestRightShift_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	rs := javast.RightShift{
		LeftOperand: javast.Identifier{
			Name: "a",
		},
		RightOperand: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := rs.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a >> b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestUnsignedRightShift_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	urs := javast.UnsignedRightShift{
		LeftOperand: javast.Identifier{
			Name: "a",
		},
		RightOperand: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := urs.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a >>> b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestLessThan_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	lt := javast.LessThan{
		LeftOperand: javast.Identifier{
			Name: "a",
		},
		RightOperand: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := lt.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a < b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestGreaterThan_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	gt := javast.GreaterThan{
		LeftOperand: javast.Identifier{
			Name: "a",
		},
		RightOperand: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := gt.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a > b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestLessThanEqual_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	lte := javast.LessThanEqual{
		LeftOperand: javast.Identifier{
			Name: "a",
		},
		RightOperand: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := lte.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a <= b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestGreaterThanEqual_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	gte := javast.GreaterThanEqual{
		LeftOperand: javast.Identifier{
			Name: "a",
		},
		RightOperand: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := gte.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a >= b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestEqualTo_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	eq := javast.EqualTo{
		LeftOperand: javast.Identifier{
			Name: "a",
		},
		RightOperand: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := eq.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a == b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestNotEqualTo_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	neq := javast.NotEqualTo{
		LeftOperand: javast.Identifier{
			Name: "a",
		},
		RightOperand: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := neq.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a != b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestAnd_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	a := javast.And{
		LeftOperand: javast.Identifier{
			Name: "a",
		},
		RightOperand: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := a.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a & b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestXor_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	x := javast.Xor{
		LeftOperand: javast.Identifier{
			Name: "a",
		},
		RightOperand: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := x.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a ^ b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestOr_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	or := javast.Or{
		LeftOperand: javast.Identifier{
			Name: "a",
		},
		RightOperand: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := or.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a | b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestConditionalAnd_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	ca := javast.ConditionalAnd{
		LeftOperand: javast.Identifier{
			Name: "a",
		},
		RightOperand: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := ca.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a && b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestConditionalOr_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	or := javast.ConditionalOr{
		LeftOperand: javast.Identifier{
			Name: "a",
		},
		RightOperand: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := or.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a || b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestMultiplyAssignment_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	ma := javast.MultiplyAssignment{
		Variable: javast.Identifier{
			Name: "a",
		},
		Expression: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := ma.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a *= b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestDivideAssignment_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	da := javast.DivideAssignment{
		Variable: javast.Identifier{
			Name: "a",
		},
		Expression: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := da.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a /= b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestRemainderAssignment_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	ra := javast.RemainderAssignment{
		Variable: javast.Identifier{
			Name: "a",
		},
		Expression: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := ra.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a %= b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestPlusAssignment_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	pa := javast.PlusAssignment{
		Variable: javast.Identifier{
			Name: "a",
		},
		Expression: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := pa.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a += b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestMinusAssignment_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	ma := javast.MinusAssignment{
		Variable: javast.Identifier{
			Name: "a",
		},
		Expression: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := ma.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a -= b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestLeftShiftAssignment_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	lsa := javast.LeftShiftAssignment{
		Variable: javast.Identifier{
			Name: "a",
		},
		Expression: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := lsa.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a <<= b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestRightShiftAssignment_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	rsa := javast.RightShiftAssignment{
		Variable: javast.Identifier{
			Name: "a",
		},
		Expression: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := rsa.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a >>= b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestUnsignedRightShiftAssignment_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	ursa := javast.UnsignedRightShiftAssignment{
		Variable: javast.Identifier{
			Name: "a",
		},
		Expression: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := ursa.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a >>>= b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestAndAssignment_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	aa := javast.AndAssignment{
		Variable: javast.Identifier{
			Name: "a",
		},
		Expression: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := aa.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a &= b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestXorAssignment_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	xa := javast.XorAssignment{
		Variable: javast.Identifier{
			Name: "a",
		},
		Expression: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := xa.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a ^= b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestOrAssignment_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	or := javast.OrAssignment{
		Variable: javast.Identifier{
			Name: "a",
		},
		Expression: javast.Identifier{
			Name: "b",
		},
	}
	if _, err := or.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "a |= b"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestIntLiteral_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	il := javast.IntLiteral{
		Value: "1000",
	}
	if _, err := il.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "1000"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestLongLiteral_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	ll := javast.LongLiteral{
		Value: "1000L",
	}
	if _, err := ll.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "1000L"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestFloatLiteral_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	fl := javast.FloatLiteral{
		Value: "1000.0f",
	}
	if _, err := fl.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "1000.0f"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestDoubleLiteral_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	dl := javast.DoubleLiteral{
		Value: "1000.0d",
	}
	if _, err := dl.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "1000.0d"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestBooleanLiteral_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	bl := javast.BooleanLiteral{
		Value: true,
	}
	if _, err := bl.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "true"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestCharLiteral_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	cl := javast.CharLiteral{
		Value: "\\n",
	}
	if _, err := cl.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "'\\n'"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestStringLiteral_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	sl := javast.StringLiteral{
		Value: "string",
	}
	if _, err := sl.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "\"string\""
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestNullLiteral_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	nl := javast.NullLiteral{}
	if _, err := nl.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "null"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestUnboundedWildcard_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	uw := javast.UnboundedWildcard{}
	if _, err := uw.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "?"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestExtendsWildcard_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	xw := javast.ExtendsWildcard{
		Bound: javast.Identifier{
			Name: "Writer",
		},
	}
	if _, err := xw.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "? extends Writer"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestSuperWildcard_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	suw := javast.SuperWildcard{
		Bound: javast.Identifier{
			Name: "Writer",
		},
	}
	if _, err := suw.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "? super Writer"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestInterface_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	i := javast.Interface{
		Modifiers: javast.Modifiers{
			Flags: []javast.Modifier{
				javast.PUBLIC_MODIFIER,
			},
			Annotations: nil,
		},
		SimpleName:     "ExpressionNode",
		TypeParameters: nil,
		ExtendsClause: javast.Identifier{
			Name: "Node",
		},
		PermitsClause: nil,
		Members: []javast.Node{
			javast.Method{
				Modifiers: javast.Modifiers{
					Flags: []javast.Modifier{
						javast.PUBLIC_MODIFIER,
					},
					Annotations: nil,
				},
				Name: "expressionNode",
				ReturnType: javast.PrimitiveType{
					PrimitiveTypeKind: javast.VOID_TYPE_KIND,
				},
				Parameters:        nil,
				ReceiverParameter: nil,
				Throws:            nil,
				Body:              nil,
				DefaultValue:      nil,
			},
		},
	}
	if _, err := i.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "public interface ExpressionNode extends Node { public void expressionNode ( ) }"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestEnum_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	e := javast.Enum{
		Modifiers: javast.Modifiers{
			Flags: []javast.Modifier{
				javast.PUBLIC_MODIFIER,
			},
			Annotations: nil,
		},
		SimpleName: "DayOfWeek",
		Members: []javast.Node{
			javast.Identifier{
				Name: "SUNDAY,",
			},
			javast.Identifier{
				Name: "MONDAY,",
			},
			javast.Identifier{
				Name: "TUESDAY,",
			},
			javast.Identifier{
				Name: "WEDNESDAY,",
			},
			javast.Identifier{
				Name: "THURSDAY,",
			},
			javast.Identifier{
				Name: "FRIDAY,",
			},
			javast.Identifier{
				Name: "SATURDAY",
			},
		},
	}
	if _, err := e.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "public enum DayOfWeek { SUNDAY, MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY }"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestAnnotationType_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	at := javast.AnnotationType{
		Modifiers: javast.Modifiers{
			Flags: []javast.Modifier{
				javast.PUBLIC_MODIFIER,
			},
			Annotations: nil,
		},
		SimpleName: "Author",
		Members: []javast.Node{
			javast.Method{
				Modifiers: javast.Modifiers{
					Flags: []javast.Modifier{
						javast.PUBLIC_MODIFIER,
					},
					Annotations: nil,
				},
				Name: "setAuthor",
				ReturnType: javast.PrimitiveType{
					PrimitiveTypeKind: javast.VOID_TYPE_KIND,
				},
				Parameters: []javast.VariableNode{
					javast.Variable{
						Modifiers: javast.Modifiers{
							Flags:       nil,
							Annotations: nil,
						},
						Name:           "author",
						NameExpression: nil,
						Type: javast.Identifier{
							Name: "String",
						},
						Initializer: nil,
					},
				},
				ReceiverParameter: nil,
				Throws:            nil,
				Body:              nil,
				DefaultValue:      nil,
			},
		},
	}
	if _, err := at.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "public @interface Author { public void setAuthor ( String author ) }"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestModule_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	m := javast.Module{
		Annotations: nil,
		ModuleType:  javast.OPEN_MODULE_KIND,
		Name: javast.MemberSelect{
			Expression: javast.Identifier{
				Name: "java",
			},
			Identifier: "sql",
		},
		Directives: []javast.DirectiveNode{
			javast.Requires{
				Static:     false,
				Transitive: true,
				ModuleName: javast.MemberSelect{
					Expression: javast.Identifier{
						Name: "java",
					},
					Identifier: "logging",
				},
			},
			javast.Exports{
				PackageName: javast.MemberSelect{
					Expression: javast.Identifier{
						Name: "java",
					},
					Identifier: "sql",
				},
				ModuleNames: nil,
			},
		},
	}
	if _, err := m.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "open module java . sql { requires transitive java . logging ; exports java . sql ; }"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestExports_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	e := javast.Exports{
		PackageName: javast.MemberSelect{
			Expression: javast.Identifier{
				Name: "java",
			},
			Identifier: "sql",
		},
		ModuleNames: []javast.ExpressionNode{
			javast.MemberSelect{
				Expression: javast.Identifier{
					Name: "javax",
				},
				Identifier: "sql",
			},
		},
	}
	if _, err := e.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "exports java . sql to javax . sql ;"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestOpens_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	o := javast.Opens{
		PackageName: javast.MemberSelect{
			Expression: javast.MemberSelect{
				Expression: javast.Identifier{
					Name: "java",
				},
				Identifier: "sql",
			},
			Identifier: "tx",
		},
		ModuleNames: []javast.ExpressionNode{
			javast.MemberSelect{
				Expression: javast.Identifier{
					Name: "java",
				},
				Identifier: "jdbc",
			},
		},
	}
	if _, err := o.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "opens java . sql . tx to java . jdbc ;"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestProvides_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	p := javast.Provides{
		ServiceName: javast.MemberSelect{
			Expression: javast.MemberSelect{
				Expression: javast.Identifier{
					Name: "java",
				},
				Identifier: "transaction",
			},
			Identifier: "Tx",
		},
		ImplementationNames: []javast.ExpressionNode{
			javast.MemberSelect{
				Expression: javast.MemberSelect{
					Expression: javast.Identifier{
						Name: "java",
					},
					Identifier: "sql",
				},
				Identifier: "Transaction",
			},
		},
	}
	if _, err := p.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "provides java . transaction . Tx with java . sql . Transaction ;"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestRecord_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	r := javast.Record{
		Modifiers: javast.Modifiers{
			Flags: []javast.Modifier{
				javast.PUBLIC_MODIFIER,
			},
			Annotations: nil,
		},
		SimpleName:     "Cat",
		TypeParameters: nil,
		ImplementsClause: []javast.Node{
			javast.Identifier{
				Name: "Animal",
			},
		},
		Members: nil,
	}
	if _, err := r.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "public record Cat implements Animal { }"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestRequires_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	r := javast.Requires{
		Static:     false,
		Transitive: true,
		ModuleName: javast.MemberSelect{
			Expression: javast.Identifier{
				Name: "java",
			},
			Identifier: "logging",
		},
	}
	if _, err := r.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "requires transitive java . logging ;"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestUses_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	u := javast.Uses{
		ServiceName: javast.MemberSelect{
			Expression: javast.MemberSelect{
				Expression: javast.Identifier{
					Name: "java",
				},
				Identifier: "sql",
			},
			Identifier: "Driver",
		},
	}
	if _, err := u.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "uses java . sql . Driver ;"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}

func TestYield_WriteTo(t *testing.T) {
	t.Parallel()
	sw := SpaceWriter{}
	y := javast.Yield{
		Value: javast.Plus{
			LeftOperand: javast.Identifier{
				Name: "a",
			},
			RightOperand: javast.Identifier{
				Name: "b",
			},
		},
	}
	if _, err := y.WriteTo(&sw); err != nil {
		t.Error(err)
	}
	got := sw.String()
	want := "yield a + b ;"
	if got != want {
		t.Errorf("sw.String() = %s, want %s", got, want)
	}
}
