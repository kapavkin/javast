package javast_test

import (
	"fmt"
	"testing"

	"github.com/kapavkin/javast"
)

func TestFormatter_Write(t *testing.T) {
	t.Parallel()
	var buf []byte
	formatter := javast.Formatter{
		Writer: javast.WriterFunc(
			func(p []byte) (int, error) {
				n := len(p)
				buf = append(buf, p...)
				return n, nil
			},
		),
		Options: javast.FormatterOptions{
			Identation: javast.Identation,
			LineLength: javast.LineLength,
		},
	}
	node := javast.Class{
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
	if _, err := node.WriteTo(&formatter); err != nil {
		t.Error(err)
	}
	fmt.Println(string(buf))
	got := string(buf)
	want := `public final class PrettyPrinter <T extends Printable> extends ConsolePrinter implements Printer {
}`
	if got != want {
		t.Errorf("string(buf) = %s, want %s", got, want)
	}
}
