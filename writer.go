package javast

import (
	"fmt"
	"io"
)

var modifiers = [...]string{
	PUBLIC_MODIFIER:       "public",
	PROTECTED_MODIFIER:    "protected",
	PRIVATE_MODIFIER:      "private",
	ABSTRACT_MODIFIER:     "abstract",
	DEFAULT_MODIFIER:      "default",
	STATIC_MODIFIER:       "static",
	SEALED_MODIFIER:       "sealed",
	NON_SEALED_MODIFIER:   "non-sealed",
	FINAL_MODIFIER:        "final",
	TRANSIENT_MODIFIER:    "transient",
	VOLATILE_MODIFIER:     "volatile",
	SYNCHRONIZED_MODIFIER: "synchronized",
	NATIVE_MODIFIER:       "native",
	STRICTFP_MODIFIER:     "strictfp",
}

// Implements [io.WriterTo] interface for [AnnotatedType].
func (at *AnnotatedType) WriteTo(w io.Writer) (n int64, err error) {
	for _, annotation := range at.Annotations {
		if an, aerr := annotation.WriteTo(w); aerr != nil {
			err = aerr
			return
		} else {
			n += an
		}
	}
	if utn, uterr := at.UnderlyingType.WriteTo(w); uterr != nil {
		err = uterr
		return
	} else {
		n += utn
	}
	return
}

// Implements [io.WriterTo] interface for [Annotation].
func (a *Annotation) WriteTo(w io.Writer) (n int64, err error) {
	if an, aerr := w.Write([]byte(`@`)); aerr != nil {
		err = aerr
		return
	} else {
		n += int64(an)
	}
	if atn, aterr := a.AnnotationType.WriteTo(w); aterr != nil {
		err = aterr
		return
	} else {
		n += atn
	}
	alen := len(a.Arguments)
	if alen > 0 {
		if on, oerr := w.Write([]byte(`(`)); oerr != nil {
			err = oerr
			return
		} else {
			n += int64(on)
		}
	}
	for _, argument := range a.Arguments {
		if an, aerr := argument.WriteTo(w); aerr != nil {
			err = aerr
			return
		} else {
			n += an
		}
	}
	if alen > 0 {
		if cn, cerr := w.Write([]byte(`)`)); cerr != nil {
			err = cerr
			return
		} else {
			n += int64(cn)
		}
	}
	return
}

// Implements [io.WriterTo] interface for [TypeAnnotation].
func (ta *TypeAnnotation) WriteTo(w io.Writer) (n int64, err error) {
	if an, aerr := w.Write([]byte(`@`)); aerr != nil {
		err = aerr
		return
	} else {
		n += int64(an)
	}
	if atn, aterr := ta.AnnotationType.WriteTo(w); aterr != nil {
		err = aterr
		return
	} else {
		n += atn
	}
	alen := len(ta.Arguments)
	if alen > 0 {
		if on, oerr := w.Write([]byte(`(`)); oerr != nil {
			err = oerr
			return
		} else {
			n += int64(on)
		}
	}
	for _, argument := range ta.Arguments {
		if an, aerr := argument.WriteTo(w); aerr != nil {
			err = aerr
			return
		} else {
			n += an
		}
	}
	if alen > 0 {
		if cn, cerr := w.Write([]byte(`)`)); cerr != nil {
			err = cerr
			return
		} else {
			n += int64(cn)
		}
	}
	return
}

// Implements [io.WriterTo] interface for [ArrayAccess].
func (aa *ArrayAccess) WriteTo(w io.Writer) (n int64, err error) {
	if xn, xerr := aa.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	if on, oerr := w.Write([]byte(`[`)); oerr != nil {
		err = oerr
		return
	} else {
		n += int64(on)
	}
	if in, ierr := aa.Index.WriteTo(w); ierr != nil {
		err = ierr
		return
	} else {
		n += in
	}
	if cn, cerr := w.Write([]byte(`]`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	return
}

// Implements [io.WriterTo] interface for [ArrayType].
func (at *ArrayType) WriteTo(w io.Writer) (n int64, err error) {
	if tn, terr := at.Type.WriteTo(w); terr != nil {
		err = terr
		return
	} else {
		n += tn
	}
	if bn, berr := w.Write([]byte(`[]`)); berr != nil {
		err = berr
		return
	} else {
		n += int64(bn)
	}
	return
}

// Implements [io.WriterTo] interface for [Assert].
func (a *Assert) WriteTo(w io.Writer) (n int64, err error) {
	if an, aerr := w.Write([]byte(`assert`)); aerr != nil {
		err = aerr
		return
	} else {
		n += int64(an)
	}
	if cn, cerr := a.Condition.WriteTo(w); cerr != nil {
		err = cerr
		return
	} else {
		n += cn
	}
	if a.Detail != nil {
		if cn, cerr := w.Write([]byte(`:`)); cerr != nil {
			err = cerr
			return
		} else {
			n += int64(cn)
		}
		if dn, derr := (*a.Detail).WriteTo(w); derr != nil {
			err = derr
			return
		} else {
			n += dn
		}
	}
	if sn, serr := w.Write([]byte(`;`)); serr != nil {
		err = serr
		return
	} else {
		n += int64(sn)
	}
	return
}

// Implements [io.WriterTo] interface for [Assignment].
func (a *Assignment) WriteTo(w io.Writer) (n int64, err error) {
	if vn, verr := a.Variable.WriteTo(w); verr != nil {
		err = verr
		return
	} else {
		n += vn
	}
	if en, eerr := w.Write([]byte(`;`)); eerr != nil {
		err = eerr
		return
	} else {
		n += int64(en)
	}
	if xn, xerr := a.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	return
}

// Implements [io.WriterTo] interface for [Block].
func (b *Block) WriteTo(w io.Writer) (n int64, err error) {
	if b.Static {
		if sn, serr := w.Write([]byte(`static`)); serr != nil {
			err = serr
			return
		} else {
			n += int64(sn)
		}
	}
	if on, oerr := w.Write([]byte(`{`)); oerr != nil {
		err = oerr
		return
	} else {
		n += int64(on)
	}
	for _, statement := range b.Statements {
		if sn, serr := statement.WriteTo(w); serr != nil {
			err = serr
			return
		} else {
			n += sn
		}
	}
	if cn, cerr := w.Write([]byte(`}`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	return
}

// Implements [io.WriterTo] interface for [Break].
func (b *Break) WriteTo(w io.Writer) (n int64, err error) {
	if bn, berr := w.Write([]byte(`break`)); berr != nil {
		err = berr
		return
	} else {
		n += int64(bn)
	}
	if b.Label != nil {
		if ln, lerr := w.Write([]byte(*b.Label)); lerr != nil {
			err = lerr
			return
		} else {
			n += int64(ln)
		}
	}
	if sn, serr := w.Write([]byte(`;`)); serr != nil {
		err = serr
		return
	} else {
		n += int64(sn)
	}
	return
}

// Implements [io.WriterTo] interface for [StatementCase].
func (sc *StatementCase) WriteTo(w io.Writer) (n int64, err error) {
	if sc.Expression != nil {
		if cn, cerr := w.Write([]byte(`case`)); cerr != nil {
			err = cerr
			return
		} else {
			n += int64(cn)
		}
		if en, eerr := (*sc.Expression).WriteTo(w); eerr != nil {
			err = eerr
			return
		} else {
			n += en
		}
	} else {
		if dn, derr := w.Write([]byte(`default`)); derr != nil {
			err = derr
			return
		} else {
			n += int64(dn)
		}
	}
	if cn, cerr := w.Write([]byte(`:`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	for _, statement := range sc.Statements {
		if sn, serr := statement.WriteTo(w); serr != nil {
			err = serr
			return
		} else {
			n += sn
		}
	}
	return
}

// Implements [io.WriterTo] interface for [RuleCase].
func (rc *RuleCase) WriteTo(w io.Writer) (n int64, err error) {
	llen := len(rc.Labels)
	if llen == 0 || llen == 1 && rc.Labels[0].GetKind() == DEFAULT_CASE_LABEL {
		if dn, derr := w.Write([]byte(`default`)); derr != nil {
			err = derr
			return
		} else {
			n += int64(dn)
		}
	} else {
		if cn, cerr := w.Write([]byte(`case`)); cerr != nil {
			err = cerr
			return
		} else {
			n += int64(cn)
		}
		for i := 0; i < llen-1; i++ {
			if ln, lerr := rc.Labels[i].WriteTo(w); lerr != nil {
				err = lerr
				return
			} else {
				n += ln
			}
			if cn, cerr := w.Write([]byte(`,`)); cerr != nil {
				err = cerr
				return
			} else {
				n += int64(cn)
			}
		}
		if ln, lerr := rc.Labels[llen-1].WriteTo(w); lerr != nil {
			err = lerr
			return
		} else {
			n += ln
		}
	}
	if an, aerr := w.Write([]byte(`->`)); aerr != nil {
		err = aerr
		return
	} else {
		n += int64(an)
	}
	if bn, berr := rc.Body.WriteTo(w); berr != nil {
		err = berr
		return
	} else {
		n += bn
	}
	return
}

// Implements [io.WriterTo] interface for [Catch].
func (c *Catch) WriteTo(w io.Writer) (n int64, err error) {
	if cn, cerr := w.Write([]byte(`catch`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	if on, oerr := w.Write([]byte(`(`)); oerr != nil {
		err = oerr
		return
	} else {
		n += int64(on)
	}
	if pn, perr := c.Parameter.WriteTo(w); perr != nil {
		err = perr
		return
	} else {
		n += pn
	}
	if cn, cerr := w.Write([]byte(`)`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	if bn, berr := c.Block.WriteTo(w); berr != nil {
		err = berr
		return
	} else {
		n += bn
	}
	return
}

// Implements [io.WriterTo] interface for [Class].
func (c *Class) WriteTo(w io.Writer) (n int64, err error) {
	if mn, merr := c.Modifiers.WriteTo(w); merr != nil {
		err = merr
		return
	} else {
		n += mn
	}
	if cn, cerr := w.Write([]byte(`class`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	if snn, snerr := w.Write([]byte(c.SimpleName)); snerr != nil {
		err = snerr
		return
	} else {
		n += int64(snn)
	}
	if tplen := len(c.TypeParameters); tplen > 0 {
		if on, oerr := w.Write([]byte(`<`)); oerr != nil {
			err = oerr
			return
		} else {
			n += int64(on)
		}
		for i := 0; i < tplen-1; i++ {
			if tpn, tperr := c.TypeParameters[i].WriteTo(w); tperr != nil {
				err = tperr
				return
			} else {
				n += tpn
			}
			if cn, cerr := w.Write([]byte(`,`)); cerr != nil {
				err = cerr
				return
			} else {
				n += int64(cn)
			}
		}
		if tpn, tperr := c.TypeParameters[tplen-1].WriteTo(w); tperr != nil {
			err = tperr
			return
		} else {
			n += tpn
		}
		if cn, cerr := w.Write([]byte(`>`)); cerr != nil {
			err = cerr
			return
		} else {
			n += int64(cn)
		}
	}
	if c.ExtendsClause != nil {
		if en, eerr := w.Write([]byte(`extends`)); eerr != nil {
			err = eerr
			return
		} else {
			n += int64(en)
		}
		if ecn, ecerr := (*c.ExtendsClause).WriteTo(w); ecerr != nil {
			err = ecerr
			return
		} else {
			n += ecn
		}
	}
	if iclen := len(c.ImplementsClause); iclen > 0 {
		if in, ierr := w.Write([]byte(`implements`)); ierr != nil {
			err = ierr
			return
		} else {
			n += int64(in)
		}
		for i := 0; i < iclen-1; i++ {
			if icn, icerr := c.ImplementsClause[i].WriteTo(w); icerr != nil {
				err = icerr
				return
			} else {
				n += icn
			}
			if cn, cerr := w.Write([]byte(`,`)); cerr != nil {
				err = cerr
				return
			} else {
				n += int64(cn)
			}
		}
		if icn, icerr := c.ImplementsClause[iclen-1].WriteTo(w); icerr != nil {
			err = icerr
			return
		} else {
			n += icn
		}
	}
	if on, oerr := w.Write([]byte(`{`)); oerr != nil {
		err = oerr
		return
	} else {
		n += int64(on)
	}
	for _, member := range c.Members {
		if mn, merr := member.WriteTo(w); merr != nil {
			err = merr
			return
		} else {
			n += mn
		}
	}
	if cn, cerr := w.Write([]byte(`}`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	return
}

// Implements [io.WriterTo] interface for [CompilationUnit].
func (cu *CompilationUnit) WriteTo(w io.Writer) (n int64, err error) {
	if cu.Module != nil {
		if mn, merr := (*cu.Module).WriteTo(w); merr != nil {
			err = merr
			return
		} else {
			n += mn
		}
	}
	if cu.Package != nil {
		if pn, perr := (*cu.Package).WriteTo(w); perr != nil {
			err = perr
			return
		} else {
			n += pn
		}
	} else if cu.PackageName != nil {
		for _, annotation := range cu.PackageAnnotations {
			if an, aerr := annotation.WriteTo(w); aerr != nil {
				err = aerr
				return
			} else {
				n += an
			}
		}
		if pn, perr := w.Write([]byte(`package`)); perr != nil {
			err = perr
			return
		} else {
			n += int64(pn)
		}
		if pnn, pnerr := (*cu.PackageName).WriteTo(w); pnerr != nil {
			err = pnerr
			return
		} else {
			n += pnn
		}
		if sn, serr := w.Write([]byte(`;`)); serr != nil {
			err = serr
			return
		} else {
			n += int64(sn)
		}
	}
	for _, i := range cu.Imports {
		if in, ierr := i.WriteTo(w); ierr != nil {
			err = ierr
			return
		} else {
			n += in
		}
	}
	for _, decl := range cu.TypeDecls {
		if tdn, tderr := decl.WriteTo(w); tderr != nil {
			err = tderr
			return
		} else {
			n += tdn
		}
	}
	return
}

// Implements [io.WriterTo] interface for [ConditionalExpression].
func (cx *ConditionalExpression) WriteTo(w io.Writer) (n int64, err error) {
	if cn, cerr := cx.Condition.WriteTo(w); cerr != nil {
		err = cerr
		return
	} else {
		n += cn
	}
	if qn, qerr := w.Write([]byte(`?`)); qerr != nil {
		err = qerr
		return
	} else {
		n += int64(qn)
	}
	if txn, txerr := cx.TrueExpression.WriteTo(w); txerr != nil {
		err = txerr
		return
	} else {
		n += txn
	}
	if cn, cerr := w.Write([]byte(`:`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	if fxn, fxerr := cx.FalseExpression.WriteTo(w); fxerr != nil {
		err = fxerr
		return
	} else {
		n += fxn
	}
	return
}

// Implements [io.WriterTo] interface for [Continue].
func (c *Continue) WriteTo(w io.Writer) (n int64, err error) {
	if cn, cerr := w.Write([]byte(`continue`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	if c.Label != nil {
		if ln, lerr := w.Write([]byte(*c.Label)); lerr != nil {
			err = lerr
			return
		} else {
			n += int64(ln)
		}
	}
	if sn, serr := w.Write([]byte(`;`)); serr != nil {
		err = serr
		return
	} else {
		n += int64(sn)
	}
	return
}

// Implements [io.WriterTo] interface for [DoWhileLoop].
func (dwl *DoWhileLoop) WriteTo(w io.Writer) (n int64, err error) {
	if dn, derr := w.Write([]byte(`do`)); derr != nil {
		err = derr
		return
	} else {
		n += int64(dn)
	}
	if sn, serr := dwl.Statement.WriteTo(w); serr != nil {
		err = serr
		return
	} else {
		n += sn
	}
	if wn, werr := w.Write([]byte(`while`)); werr != nil {
		err = werr
		return
	} else {
		n += int64(wn)
	}
	if on, oerr := w.Write([]byte(`(`)); oerr != nil {
		err = oerr
		return
	} else {
		n += int64(on)
	}
	if cn, cerr := dwl.Condition.WriteTo(w); cerr != nil {
		err = cerr
		return
	} else {
		n += cn
	}
	if cn, cerr := w.Write([]byte(`)`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	return
}

// Implements [io.WriterTo] interface for [EnhancedForLoop].
func (efl *EnhancedForLoop) WriteTo(w io.Writer) (n int64, err error) {
	if fn, ferr := w.Write([]byte(`for`)); ferr != nil {
		err = ferr
		return
	} else {
		n += int64(fn)
	}

	if on, oerr := w.Write([]byte(`(`)); oerr != nil {
		err = oerr
		return
	} else {
		n += int64(on)
	}
	if vn, verr := efl.Variable.WriteTo(w); verr != nil {
		err = verr
		return
	} else {
		n += vn
	}
	if cn, cerr := w.Write([]byte(`:`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	if xn, xerr := efl.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	if cn, cerr := w.Write([]byte(`)`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	if sn, serr := efl.Statement.WriteTo(w); serr != nil {
		err = serr
		return
	} else {
		n += sn
	}
	return
}

// Implements [io.WriterTo] interface for [ExpressionStatement].
func (xs *ExpressionStatement) WriteTo(w io.Writer) (n int64, err error) {
	if xn, xerr := xs.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	if sn, serr := w.Write([]byte(`;`)); serr != nil {
		err = serr
		return
	} else {
		n += int64(sn)
	}
	return
}

// Implements [io.WriterTo] interface for [MemberSelect].
func (ms *MemberSelect) WriteTo(w io.Writer) (n int64, err error) {
	if xn, xerr := ms.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	if dn, derr := w.Write([]byte(`.`)); derr != nil {
		err = derr
		return
	} else {
		n += int64(dn)
	}
	if in, ierr := w.Write([]byte(ms.Identifier)); ierr != nil {
		err = ierr
		return
	} else {
		n += int64(in)
	}
	return
}

// Implements [io.WriterTo] interface for [InvokeMemberReference].
func (imr *InvokeMemberReference) WriteTo(w io.Writer) (n int64, err error) {
	if qen, qeerr := imr.QualifierExpression.WriteTo(w); qeerr != nil {
		err = qeerr
		return
	} else {
		n += qen
	}
	if cn, cerr := w.Write([]byte(`::`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	if talen := len(imr.TypeArguments); talen > 0 {
		if on, oerr := w.Write([]byte(`<`)); oerr != nil {
			err = oerr
			return
		} else {
			n += int64(on)
		}
		for i := 0; i < talen-1; i++ {
			if tan, taerr := imr.TypeArguments[i].WriteTo(w); taerr != nil {
				err = taerr
				return
			} else {
				n += tan
			}
			if cn, cerr := w.Write([]byte(`,`)); cerr != nil {
				err = cerr
				return
			} else {
				n += int64(cn)
			}
		}
		if tan, taerr := imr.TypeArguments[talen-1].WriteTo(w); taerr != nil {
			err = taerr
			return
		} else {
			n += tan
		}
		if cn, cerr := w.Write([]byte(`>`)); cerr != nil {
			err = cerr
			return
		} else {
			n += int64(cn)
		}
	}
	if nn, nerr := w.Write([]byte(imr.Name)); nerr != nil {
		err = nerr
		return
	} else {
		n += int64(nn)
	}
	return
}

// Implements [io.WriterTo] interface for [NewMemberReference].
func (nmr *NewMemberReference) WriteTo(w io.Writer) (n int64, err error) {
	if qen, qeerr := nmr.QualifierExpression.WriteTo(w); qeerr != nil {
		err = qeerr
		return
	} else {
		n += qen
	}
	if cn, cerr := w.Write([]byte(`::`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	if talen := len(nmr.TypeArguments); talen > 0 {
		if on, oerr := w.Write([]byte(`<`)); oerr != nil {
			err = oerr
			return
		} else {
			n += int64(on)
		}
		for i := 0; i < talen-1; i++ {
			if tan, taerr := nmr.TypeArguments[i].WriteTo(w); taerr != nil {
				err = taerr
				return
			} else {
				n += tan
			}
			if cn, cerr := w.Write([]byte(`,`)); cerr != nil {
				err = cerr
				return
			} else {
				n += int64(cn)
			}
		}
		if tan, taerr := nmr.TypeArguments[talen-1].WriteTo(w); taerr != nil {
			err = taerr
			return
		} else {
			n += tan
		}
		if cn, cerr := w.Write([]byte(`>`)); cerr != nil {
			err = cerr
			return
		} else {
			n += int64(cn)
		}
	}
	if nn, nerr := w.Write([]byte(`new`)); nerr != nil {
		err = nerr
		return
	} else {
		n += int64(nn)
	}
	return
}

// Implements [io.WriterTo] interface for [ForLoop].
func (fl *ForLoop) WriteTo(w io.Writer) (n int64, err error) {
	if fn, ferr := w.Write([]byte(`for`)); ferr != nil {
		err = ferr
		return
	} else {
		n += int64(fn)
	}
	if on, oerr := w.Write([]byte(`(`)); oerr != nil {
		err = oerr
		return
	} else {
		n += int64(on)
	}
	if ilen := len(fl.Initializer); ilen > 0 {
		for i := 0; i < ilen-1; i++ {
			if in, ierr := fl.Initializer[i].WriteTo(w); ierr != nil {
				err = ierr
				return
			} else {
				n += in
			}
			if cn, cerr := w.Write([]byte(`,`)); cerr != nil {
				err = cerr
				return
			} else {
				n += int64(cn)
			}
		}
		if in, ierr := fl.Initializer[ilen-1].WriteTo(w); ierr != nil {
			err = ierr
			return
		} else {
			n += in
		}
	}
	if sn, serr := w.Write([]byte(`;`)); serr != nil {
		err = serr
		return
	} else {
		n += int64(sn)
	}
	if fl.Condition != nil {
		if cn, cerr := (*fl.Condition).WriteTo(w); cerr != nil {
			err = cerr
			return
		} else {
			n += cn
		}
	}
	if sn, serr := w.Write([]byte(`;`)); serr != nil {
		err = serr
		return
	} else {
		n += int64(sn)
	}
	if ulen := len(fl.Update); ulen > 0 {
		for i := 0; i < ulen-1; i++ {
			if un, uerr := fl.Update[i].WriteTo(w); uerr != nil {
				err = uerr
				return
			} else {
				n += un
			}
			if cn, cerr := w.Write([]byte(`,`)); cerr != nil {
				err = cerr
				return
			} else {
				n += int64(cn)
			}
		}
		if un, uerr := fl.Update[ulen-1].WriteTo(w); uerr != nil {
			err = uerr
			return
		} else {
			n += un
		}
	}
	if cn, cerr := w.Write([]byte(`)`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	if sn, serr := fl.Statement.WriteTo(w); serr != nil {
		err = serr
		return
	} else {
		n += sn
	}
	return
}

// Implements [io.WriterTo] interface for [Identifier].
func (i *Identifier) WriteTo(w io.Writer) (n int64, err error) {
	if nn, nerr := w.Write([]byte(i.Name)); nerr != nil {
		err = nerr
		return
	} else {
		n += int64(nn)
	}
	return
}

// Implements [io.WriterTo] interface for [If].
func (i *If) WriteTo(w io.Writer) (n int64, err error) {
	if in, ierr := w.Write([]byte(`if`)); ierr != nil {
		err = ierr
		return
	} else {
		n += int64(in)
	}
	if on, oerr := w.Write([]byte(`(`)); oerr != nil {
		err = oerr
		return
	} else {
		n += int64(on)
	}
	if cn, cerr := i.Condition.WriteTo(w); cerr != nil {
		err = cerr
		return
	} else {
		n += cn
	}
	if cn, cerr := w.Write([]byte(`)`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	if tn, terr := i.ThenStatement.WriteTo(w); terr != nil {
		err = terr
		return
	} else {
		n += tn
	}
	if i.ElseStatement != nil {
		if en, eerr := w.Write([]byte(`else`)); eerr != nil {
			err = eerr
			return
		} else {
			n += int64(en)
		}
		if en, eerr := (*i.ElseStatement).WriteTo(w); eerr != nil {
			err = eerr
			return
		} else {
			n += en
		}
	}
	return
}

// Implements [io.WriterTo] interface for [Import].
func (i *Import) WriteTo(w io.Writer) (n int64, err error) {
	if i.Static {
		if sn, serr := w.Write([]byte(`static`)); serr != nil {
			err = serr
			return
		} else {
			n += int64(sn)
		}
	}
	if in, ierr := w.Write([]byte(`import`)); ierr != nil {
		err = ierr
		return
	} else {
		n += int64(in)
	}
	if qin, qierr := i.QualifiedIdentifier.WriteTo(w); qierr != nil {
		err = qierr
		return
	} else {
		n += qin
	}
	return
}

// Implements [io.WriterTo] interface for [InstanceOf].
func (io *InstanceOf) WriteTo(w io.Writer) (n int64, err error) {
	if xn, xerr := io.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	if ion, ioerr := w.Write([]byte(`instanceof`)); ioerr != nil {
		err = ioerr
		return
	} else {
		n += int64(ion)
	}
	if tn, terr := io.Type.WriteTo(w); terr != nil {
		err = terr
		return
	} else {
		n += tn
	}
	if io.Pattern != nil {
		if pn, perr := (*io.Pattern).WriteTo(w); perr != nil {
			err = perr
			return
		} else {
			n += pn
		}
	}
	return
}

// Implements [io.WriterTo] interface for [LabeledStatement].
func (ls *LabeledStatement) WriteTo(w io.Writer) (n int64, err error) {
	if ln, lerr := w.Write([]byte(ls.Label)); lerr != nil {
		err = lerr
		return
	} else {
		n += int64(ln)
	}
	if cn, cerr := w.Write([]byte(`:`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	if sn, serr := ls.Statement.WriteTo(w); serr != nil {
		err = serr
		return
	} else {
		n += sn
	}
	return
}

// Implements [io.WriterTo] interface for [Method].
func (m *Method) WriteTo(w io.Writer) (n int64, err error) {
	if mn, merr := m.Modifiers.WriteTo(w); merr != nil {
		err = merr
		return
	} else {
		n += mn
	}
	if tplen := len(m.TypeParameters); tplen > 0 {
		if on, oerr := w.Write([]byte(`<`)); oerr != nil {
			err = oerr
			return
		} else {
			n += int64(on)
		}
		for i := 0; i < tplen-1; i++ {
			if tpn, tperr := m.TypeParameters[i].WriteTo(w); tperr != nil {
				err = tperr
				return
			} else {
				n += tpn
			}
			if cn, cerr := w.Write([]byte(`,`)); cerr != nil {
				err = cerr
				return
			} else {
				n += int64(cn)
			}
		}
		if tpn, tperr := m.TypeParameters[tplen-1].WriteTo(w); tperr != nil {
			err = tperr
			return
		} else {
			n += tpn
		}
		if cn, cerr := w.Write([]byte(`>`)); cerr != nil {
			err = cerr
			return
		} else {
			n += int64(cn)
		}
	}
	if m.ReturnType != nil {
		if rtn, rterr := (*m.ReturnType).WriteTo(w); rterr != nil {
			err = rterr
			return
		} else {
			n += rtn
		}
	}
	if nn, nerr := w.Write([]byte(m.Name)); nerr != nil {
		err = nerr
		return
	} else {
		n += int64(nn)
	}
	if on, oerr := w.Write([]byte(`(`)); oerr != nil {
		err = oerr
		return
	} else {
		n += int64(on)
	}
	if m.ReceiverParameter != nil {
		if rpn, rperr := (*m.ReceiverParameter).WriteTo(w); rperr != nil {
			err = rperr
			return
		} else {
			n += rpn
		}
		if len(m.Parameters) > 0 {
			if cn, cerr := w.Write([]byte(`,`)); cerr != nil {
				err = cerr
				return
			} else {
				n += int64(cn)
			}
		}
	}
	if plen := len(m.Parameters); plen > 0 {
		for i := 0; i < plen-1; i++ {
			if pn, perr := m.Parameters[i].WriteTo(w); perr != nil {
				err = perr
				return
			} else {
				n += pn
			}
			if cn, cerr := w.Write([]byte(`,`)); cerr != nil {
				err = cerr
				return
			} else {
				n += int64(cn)
			}
		}
		if pn, perr := m.Parameters[plen-1].WriteTo(w); perr != nil {
			err = perr
			return
		} else {
			n += pn
		}
	}
	if cn, cerr := w.Write([]byte(`)`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	if tlen := len(m.Throws); tlen > 0 {
		if tn, terr := w.Write([]byte(`throws`)); terr != nil {
			err = terr
			return
		} else {
			n += int64(tn)
		}
		for i := 0; i < tlen-1; i++ {
			if tn, terr := m.Throws[i].WriteTo(w); terr != nil {
				err = terr
				return
			} else {
				n += tn
			}
			if cn, cerr := w.Write([]byte(`,`)); cerr != nil {
				err = cerr
				return
			} else {
				n += int64(cn)
			}
		}
		if tn, terr := m.Throws[tlen-1].WriteTo(w); terr != nil {
			err = terr
			return
		} else {
			n += tn
		}
	}
	if m.Body != nil {
		if bn, berr := (*m.Body).WriteTo(w); berr != nil {
			err = berr
			return
		} else {
			n += bn
		}
	}
	if m.DefaultValue != nil {
		if dn, derr := w.Write([]byte(`default`)); derr != nil {
			err = derr
			return
		} else {
			n += int64(dn)
		}
		if dvn, dverr := (*m.DefaultValue).WriteTo(w); dverr != nil {
			err = dverr
			return
		} else {
			n += dvn
		}
	}
	return
}

// Implements [io.WriterTo] interface for [MethodInvocation].
func (mi *MethodInvocation) WriteTo(w io.Writer) (n int64, err error) {
	if talen := len(mi.TypeArguments); talen > 0 {
		if on, oerr := w.Write([]byte(`<`)); oerr != nil {
			err = oerr
			return
		} else {
			n += int64(on)
		}
		for i := 0; i < talen-1; i++ {
			if tan, taerr := mi.TypeArguments[i].WriteTo(w); taerr != nil {
				err = taerr
				return
			} else {
				n += tan
			}
			if cn, cerr := w.Write([]byte(`,`)); cerr != nil {
				err = cerr
				return
			} else {
				n += int64(cn)
			}
		}
		if tan, taerr := mi.TypeArguments[talen-1].WriteTo(w); taerr != nil {
			err = taerr
			return
		} else {
			n += tan
		}
		if cn, cerr := w.Write([]byte(`>`)); cerr != nil {
			err = cerr
			return
		} else {
			n += int64(cn)
		}
	}
	if msn, mserr := mi.MethodSelect.WriteTo(w); mserr != nil {
		err = mserr
		return
	} else {
		n += msn
	}
	if on, oerr := w.Write([]byte(`(`)); oerr != nil {
		err = oerr
		return
	} else {
		n += int64(on)
	}
	if alen := len(mi.Arguments); alen > 0 {
		for i := 0; i < alen-1; i++ {
			if an, aerr := mi.Arguments[i].WriteTo(w); aerr != nil {
				err = aerr
				return
			} else {
				n += an
			}
			if cn, cerr := w.Write([]byte(`,`)); cerr != nil {
				err = cerr
				return
			} else {
				n += int64(cn)
			}
		}
		if an, aerr := mi.Arguments[alen-1].WriteTo(w); aerr != nil {
			err = aerr
			return
		} else {
			n += an
		}
	}
	if cn, cerr := w.Write([]byte(`)`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	return
}

// Implements [io.WriterTo] interface for [Modifiers].
func (m *Modifiers) WriteTo(w io.Writer) (n int64, err error) {
	for _, flag := range m.Flags {
		if fn, ferr := w.Write([]byte(modifiers[flag])); ferr != nil {
			err = ferr
			return
		} else {
			n += int64(fn)
		}
	}
	for _, annotation := range m.Annotations {
		if an, aerr := annotation.WriteTo(w); aerr != nil {
			err = aerr
			return
		} else {
			n += an
		}
	}
	return
}

// Implements [io.WriterTo] interface for [NewArray].
func (na *NewArray) WriteTo(w io.Writer) (n int64, err error) {
	if na.Type != nil {
		if nn, nerr := w.Write([]byte(`new`)); nerr != nil {
			err = nerr
			return
		} else {
			n += int64(nn)
		}
		if tn, terr := (*na.Type).WriteTo(w); terr != nil {
			err = terr
			return
		} else {
			n += tn
		}
	}
	for _, dimension := range na.Dimensions {
		if on, oerr := w.Write([]byte(`[`)); oerr != nil {
			err = oerr
			return
		} else {
			n += int64(on)
		}
		if dn, derr := dimension.WriteTo(w); derr != nil {
			err = derr
			return
		} else {
			n += dn
		}
		if cn, cerr := w.Write([]byte(`[`)); cerr != nil {
			err = cerr
			return
		} else {
			n += int64(cn)
		}
	}
	if ilen := len(na.Initializers); ilen > 0 {
		if on, oerr := w.Write([]byte(`{`)); oerr != nil {
			err = oerr
			return
		} else {
			n += int64(on)
		}
		for i := 0; i < ilen-1; i++ {
			if in, ierr := na.Initializers[i].WriteTo(w); ierr != nil {
				err = ierr
				return
			} else {
				n += in
			}
			if cn, cerr := w.Write([]byte(`,`)); cerr != nil {
				err = cerr
				return
			} else {
				n += int64(cn)
			}
		}
		if in, ierr := na.Initializers[ilen-1].WriteTo(w); ierr != nil {
			err = ierr
			return
		} else {
			n += in
		}
		if cn, cerr := w.Write([]byte(`}`)); cerr != nil {
			err = cerr
			return
		} else {
			n += int64(cn)
		}
	}
	return
}

// Implements [io.WriterTo] interface for [NewClass].
func (nc *NewClass) WriteTo(w io.Writer) (n int64, err error) {
	if nc.EnclosingExpression != nil {
		if exn, exerr := (*nc.EnclosingExpression).WriteTo(w); exerr != nil {
			err = exerr
			return
		} else {
			n += exn
		}
		if dn, derr := w.Write([]byte(`.`)); derr != nil {
			err = derr
			return
		} else {
			n += int64(dn)
		}
	}
	if nn, nerr := w.Write([]byte(`new`)); nerr != nil {
		err = nerr
		return
	} else {
		n += int64(nn)
	}
	if talen := len(nc.TypeArguments); talen > 0 {
		if on, oerr := w.Write([]byte(`<`)); oerr != nil {
			err = oerr
			return
		} else {
			n += int64(on)
		}
		for i := 0; i < talen-1; i++ {
			if tan, taerr := nc.TypeArguments[i].WriteTo(w); taerr != nil {
				err = taerr
				return
			} else {
				n += tan
			}
			if cn, cerr := w.Write([]byte(`,`)); cerr != nil {
				err = cerr
				return
			} else {
				n += int64(cn)
			}
		}
		if tan, taerr := nc.TypeArguments[talen-1].WriteTo(w); taerr != nil {
			err = taerr
			return
		} else {
			n += tan
		}
		if cn, cerr := w.Write([]byte(`>`)); cerr != nil {
			err = cerr
			return
		} else {
			n += int64(cn)
		}
	}
	if in, ierr := nc.Identifier.WriteTo(w); ierr != nil {
		err = ierr
		return
	} else {
		n += in
	}
	if on, oerr := w.Write([]byte(`(`)); oerr != nil {
		err = oerr
		return
	} else {
		n += int64(on)
	}
	if alen := len(nc.Arguments); alen > 0 {
		for i := 0; i < alen-1; i++ {
			if an, aerr := nc.Arguments[i].WriteTo(w); aerr != nil {
				err = aerr
				return
			} else {
				n += an
			}
			if cn, cerr := w.Write([]byte(`,`)); cerr != nil {
				err = cerr
				return
			} else {
				n += int64(cn)
			}
		}
		if an, aerr := nc.Arguments[alen-1].WriteTo(w); aerr != nil {
			err = aerr
			return
		} else {
			n += an
		}
	}
	if cn, cerr := w.Write([]byte(`)`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	if nc.ClassBody != nil {
		if cbn, cberr := (*nc.ClassBody).WriteTo(w); cberr != nil {
			err = cberr
			return
		} else {
			n += cbn
		}
	}
	return
}

// Implements [io.WriterTo] interface for [ExpressionLambdaExpression].
func (xlx *ExpressionLambdaExpression) WriteTo(w io.Writer) (n int64, err error) {
	if on, oerr := w.Write([]byte(`(`)); oerr != nil {
		err = oerr
		return
	} else {
		n += int64(on)
	}
	if plen := len(xlx.Parameters); plen > 0 {
		for i := 0; i < plen-1; i++ {
			if pn, perr := xlx.Parameters[i].WriteTo(w); perr != nil {
				err = perr
				return
			} else {
				n += pn
			}
			if cn, cerr := w.Write([]byte(`,`)); cerr != nil {
				err = cerr
				return
			} else {
				n += int64(cn)
			}
		}
		if pn, perr := xlx.Parameters[plen-1].WriteTo(w); perr != nil {
			err = perr
			return
		} else {
			n += pn
		}
	}
	if cn, cerr := w.Write([]byte(`)`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	if an, aerr := w.Write([]byte(`->`)); aerr != nil {
		err = aerr
		return
	} else {
		n += int64(an)
	}
	if xn, xerr := xlx.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	return
}

// Implements [io.WriterTo] interface for [StatementLambdaExpression].
func (slx *StatementLambdaExpression) WriteTo(w io.Writer) (n int64, err error) {
	if on, oerr := w.Write([]byte(`(`)); oerr != nil {
		err = oerr
		return
	} else {
		n += int64(on)
	}
	if plen := len(slx.Parameters); plen > 0 {
		for i := 0; i < plen-1; i++ {
			if pn, perr := slx.Parameters[i].WriteTo(w); perr != nil {
				err = perr
				return
			} else {
				n += pn
			}
			if cn, cerr := w.Write([]byte(`,`)); cerr != nil {
				err = cerr
				return
			} else {
				n += int64(cn)
			}
		}
		if pn, perr := slx.Parameters[plen-1].WriteTo(w); perr != nil {
			err = perr
			return
		} else {
			n += pn
		}
	}
	if cn, cerr := w.Write([]byte(`)`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	if an, aerr := w.Write([]byte(`->`)); aerr != nil {
		err = aerr
		return
	} else {
		n += int64(an)
	}
	if bn, berr := slx.Block.WriteTo(w); berr != nil {
		err = berr
		return
	} else {
		n += bn
	}
	return
}

// Implements [io.WriterTo] interface for [Package].
func (p *Package) WriteTo(w io.Writer) (n int64, err error) {
	for _, annotation := range p.Annotations {
		if an, aerr := annotation.WriteTo(w); aerr != nil {
			err = aerr
			return
		} else {
			n += an
		}
	}
	if pn, perr := w.Write([]byte(`package`)); perr != nil {
		err = perr
		return
	} else {
		n += int64(pn)
	}
	if pnn, pnerr := p.PackageName.WriteTo(w); pnerr != nil {
		err = pnerr
		return
	} else {
		n += pnn
	}
	if sn, serr := w.Write([]byte(`;`)); serr != nil {
		err = serr
		return
	} else {
		n += int64(sn)
	}
	return
}

// Implements [io.WriterTo] interface for [Parenthesized].
func (p *Parenthesized) WriteTo(w io.Writer) (n int64, err error) {
	if on, oerr := w.Write([]byte(`(`)); oerr != nil {
		err = oerr
		return
	} else {
		n += int64(on)
	}
	if xn, xerr := p.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	if cn, cerr := w.Write([]byte(`)`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	return
}

// Implements [io.WriterTo] interface for [BindingPattern].
func (bp *BindingPattern) WriteTo(w io.Writer) (n int64, err error) {
	if vn, verr := bp.Variable.WriteTo(w); verr != nil {
		err = verr
		return
	} else {
		n += vn
	}
	return
}

// Implements [io.WriterTo] interface for [GuardedPattern].
func (gp *GuardedPattern) WriteTo(w io.Writer) (n int64, err error) {
	if pn, perr := gp.Pattern.WriteTo(w); perr != nil {
		err = perr
		return
	} else {
		n += pn
	}
	if an, aerr := w.Write([]byte(`&`)); aerr != nil {
		err = aerr
		return
	} else {
		n += int64(an)
	}
	if xn, xerr := gp.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	return
}

// Implements [io.WriterTo] interface for [ParenthesizedPattern].
func (pp *ParenthesizedPattern) WriteTo(w io.Writer) (n int64, err error) {
	if on, oerr := w.Write([]byte(`(`)); oerr != nil {
		err = oerr
		return
	} else {
		n += int64(on)
	}
	if pn, perr := pp.Pattern.WriteTo(w); perr != nil {
		err = perr
		return
	} else {
		n += pn
	}
	if cn, cerr := w.Write([]byte(`)`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	return
}

// Implements [io.WriterTo] interface for [DefaultCaseLabel].
func (*DefaultCaseLabel) WriteTo(w io.Writer) (n int64, err error) {
	if dn, derr := w.Write([]byte(`default`)); derr != nil {
		err = derr
		return
	} else {
		n += int64(dn)
	}
	return
}

// Implements [io.WriterTo] interface for [PrimitiveType].
func (pt *PrimitiveType) WriteTo(w io.Writer) (n int64, err error) {
	switch pt.GetPrimitiveTypeKind() {
	case BOOLEAN_TYPE_KIND:
		if bn, berr := w.Write([]byte(`boolean`)); berr != nil {
			err = berr
			return
		} else {
			n += int64(bn)
		}
		return
	case BYTE_TYPE_KIND:
		if bn, berr := w.Write([]byte(`byte`)); berr != nil {
			err = berr
			return
		} else {
			n += int64(bn)
		}
		return
	case SHORT_TYPE_KIND:
		if sn, serr := w.Write([]byte(`short`)); serr != nil {
			err = serr
			return
		} else {
			n += int64(sn)
		}
		return
	case INT_TYPE_KIND:
		if in, ierr := w.Write([]byte(`int`)); ierr != nil {
			err = ierr
			return
		} else {
			n += int64(in)
		}
		return
	case LONG_TYPE_KIND:
		if ln, lerr := w.Write([]byte(`long`)); lerr != nil {
			err = lerr
			return
		} else {
			n += int64(ln)
		}
		return
	case CHAR_TYPE_KIND:
		if cn, cerr := w.Write([]byte(`char`)); cerr != nil {
			err = cerr
			return
		} else {
			n += int64(cn)
		}
		return
	case FLOAT_TYPE_KIND:
		if fn, ferr := w.Write([]byte(`float`)); ferr != nil {
			err = ferr
			return
		} else {
			n += int64(fn)
		}
		return
	case DOUBLE_TYPE_KIND:
		if dn, derr := w.Write([]byte(`double`)); derr != nil {
			err = derr
			return
		} else {
			n += int64(dn)
		}
		return
	case VOID_TYPE_KIND:
		if vn, verr := w.Write([]byte(`void`)); verr != nil {
			err = verr
			return
		} else {
			n += int64(vn)
		}
		return
	case NULL_TYPE_KIND:
		if nn, nerr := w.Write([]byte(`null`)); nerr != nil {
			err = nerr
			return
		} else {
			n += int64(nn)
		}
		return
	default:
		err = fmt.Errorf("unexpected primitive type kind")
		return
	}
}

// Implements [io.WriterTo] interface for [Return].
func (r *Return) WriteTo(w io.Writer) (n int64, err error) {
	if rn, rerr := w.Write([]byte(`return`)); rerr != nil {
		err = rerr
		return
	} else {
		n += int64(rn)
	}
	if r.Expression != nil {
		if xn, xerr := (*r.Expression).WriteTo(w); xerr != nil {
			err = xerr
			return
		} else {
			n += xn
		}
	}
	if sn, serr := w.Write([]byte(`;`)); serr != nil {
		err = serr
		return
	} else {
		n += int64(sn)
	}
	return
}

// Implements [io.WriterTo] interface for [EmptyStatement].
func (*EmptyStatement) WriteTo(w io.Writer) (n int64, err error) {
	if sn, serr := w.Write([]byte(`;`)); serr != nil {
		err = serr
		return
	} else {
		n += int64(sn)
	}
	return
}

// Implements [io.WriterTo] interface for [Switch].
func (s *Switch) WriteTo(w io.Writer) (n int64, err error) {
	if sn, serr := w.Write([]byte(`switch`)); serr != nil {
		err = serr
		return
	} else {
		n += int64(sn)
	}
	if on, oerr := w.Write([]byte(`(`)); oerr != nil {
		err = oerr
		return
	} else {
		n += int64(on)
	}
	if xn, xerr := s.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	if cn, cerr := w.Write([]byte(`)`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	if on, oerr := w.Write([]byte(`{`)); oerr != nil {
		err = oerr
		return
	} else {
		n += int64(on)
	}
	for _, c := range s.Cases {
		if cn, cerr := c.WriteTo(w); cerr != nil {
			err = cerr
			return
		} else {
			n += cn
		}
	}
	if cn, cerr := w.Write([]byte(`}`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	return
}

// Implements [io.WriterTo] interface for [SwitchExpression].
func (sx *SwitchExpression) WriteTo(w io.Writer) (n int64, err error) {
	if sn, serr := w.Write([]byte(`switch`)); serr != nil {
		err = serr
		return
	} else {
		n += int64(sn)
	}
	if on, oerr := w.Write([]byte(`(`)); oerr != nil {
		err = oerr
		return
	} else {
		n += int64(on)
	}
	if xn, xerr := sx.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	if cn, cerr := w.Write([]byte(`)`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	if on, oerr := w.Write([]byte(`{`)); oerr != nil {
		err = oerr
		return
	} else {
		n += int64(on)
	}
	for _, c := range sx.Cases {
		if cn, cerr := c.WriteTo(w); cerr != nil {
			err = cerr
			return
		} else {
			n += cn
		}
	}
	if cn, cerr := w.Write([]byte(`}`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	return
}

// Implements [io.WriterTo] interface for [Synchronized].
func (s *Synchronized) WriteTo(w io.Writer) (n int64, err error) {
	if sn, serr := w.Write([]byte(`synchronized`)); serr != nil {
		err = serr
		return
	} else {
		n += int64(sn)
	}

	if on, oerr := w.Write([]byte(`(`)); oerr != nil {
		err = oerr
		return
	} else {
		n += int64(on)
	}
	if xn, xerr := s.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	if cn, cerr := w.Write([]byte(`)`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	if bn, berr := s.Block.WriteTo(w); berr != nil {
		err = berr
		return
	} else {
		n += bn
	}
	return
}

// Implements [io.WriterTo] interface for [Throw].
func (t *Throw) WriteTo(w io.Writer) (n int64, err error) {
	if tn, terr := w.Write([]byte(`throw`)); terr != nil {
		err = terr
		return
	} else {
		n += int64(tn)
	}
	if xn, xerr := t.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	if sn, serr := w.Write([]byte(`;`)); serr != nil {
		err = serr
		return
	} else {
		n += int64(sn)
	}
	return
}

// Implements [io.WriterTo] interface for [Try].
func (t *Try) WriteTo(w io.Writer) (n int64, err error) {
	if tn, terr := w.Write([]byte(`try`)); terr != nil {
		err = terr
		return
	} else {
		n += int64(tn)
	}
	if rlen := len(t.Resources); rlen > 0 {
		if on, oerr := w.Write([]byte(`(`)); oerr != nil {
			err = oerr
			return
		} else {
			n += int64(on)
		}
		for i := 0; i < rlen-1; i++ {
			if rn, rerr := t.Resources[i].WriteTo(w); rerr != nil {
				err = rerr
				return
			} else {
				n += rn
			}
			if sn, serr := w.Write([]byte(`;`)); serr != nil {
				err = serr
				return
			} else {
				n += int64(sn)
			}
		}
		if rn, rerr := t.Resources[rlen-1].WriteTo(w); rerr != nil {
			err = rerr
			return
		} else {
			n += rn
		}
		if cn, cerr := w.Write([]byte(`)`)); cerr != nil {
			err = cerr
			return
		} else {
			n += int64(cn)
		}
	}
	if bn, berr := t.Block.WriteTo(w); berr != nil {
		err = berr
		return
	} else {
		n += bn
	}
	for _, catch := range t.Catches {
		if cn, cerr := catch.WriteTo(w); cerr != nil {
			err = cerr
			return
		} else {
			n += cn
		}
	}
	if t.FinallyBlock != nil {
		if fn, ferr := w.Write([]byte(`finally`)); ferr != nil {
			err = ferr
			return
		} else {
			n += int64(fn)
		}
		if fbn, fberr := (*t.FinallyBlock).WriteTo(w); fberr != nil {
			err = fberr
			return
		} else {
			n += fbn
		}
	}
	return
}

// Implements [io.WriterTo] interface for [ParameterizedType].
func (pt *ParameterizedType) WriteTo(w io.Writer) (n int64, err error) {
	if tn, terr := pt.Type.WriteTo(w); terr != nil {
		err = terr
		return
	} else {
		n += tn
	}
	if talen := len(pt.TypeArguments); talen > 0 {
		if on, oerr := w.Write([]byte(`<`)); oerr != nil {
			err = oerr
			return
		} else {
			n += int64(on)
		}
		for i := 0; i < talen-1; i++ {
			if tan, taerr := pt.TypeArguments[i].WriteTo(w); taerr != nil {
				err = taerr
				return
			} else {
				n += tan
			}
			if cn, cerr := w.Write([]byte(`,`)); cerr != nil {
				err = cerr
				return
			} else {
				n += int64(cn)
			}
		}
		if tan, taerr := pt.TypeArguments[talen-1].WriteTo(w); taerr != nil {
			err = taerr
			return
		} else {
			n += tan
		}
		if cn, cerr := w.Write([]byte(`>`)); cerr != nil {
			err = cerr
			return
		} else {
			n += int64(cn)
		}
	}
	return
}

// Implements [io.WriterTo] interface for [UnionType].
func (ut *UnionType) WriteTo(w io.Writer) (n int64, err error) {
	if talen := len(ut.TypeAlternatives); talen > 0 {
		for i := 0; i < talen-1; i++ {
			if tan, taerr := ut.TypeAlternatives[i].WriteTo(w); taerr != nil {
				err = taerr
				return
			} else {
				n += tan
			}
			if pn, perr := w.Write([]byte(`|`)); perr != nil {
				err = perr
				return
			} else {
				n += int64(pn)
			}
		}
		if tan, taerr := ut.TypeAlternatives[talen-1].WriteTo(w); taerr != nil {
			err = taerr
			return
		} else {
			n += tan
		}
	}
	return
}

// Implements [io.WriterTo] interface for [IntersectionType].
func (it *IntersectionType) WriteTo(w io.Writer) (n int64, err error) {
	if blen := len(it.Bounds); blen > 0 {
		for i := 0; i < blen-1; i++ {
			if bn, berr := it.Bounds[i].WriteTo(w); berr != nil {
				err = berr
				return
			} else {
				n += bn
			}
			if pn, perr := w.Write([]byte(`&`)); perr != nil {
				err = perr
				return
			} else {
				n += int64(pn)
			}
		}
		if bn, berr := it.Bounds[blen-1].WriteTo(w); berr != nil {
			err = berr
			return
		} else {
			n += bn
		}
	}
	return
}

// Implements [io.WriterTo] interface for [TypeCast].
func (tc *TypeCast) WriteTo(w io.Writer) (n int64, err error) {
	if on, oerr := w.Write([]byte(`(`)); oerr != nil {
		err = oerr
		return
	} else {
		n += int64(on)
	}
	if tn, terr := tc.Type.WriteTo(w); terr != nil {
		err = terr
		return
	} else {
		n += tn
	}
	if cn, cerr := w.Write([]byte(`)`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	if xn, xerr := tc.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	return
}

// Implements [io.WriterTo] interface for [TypeParameter].
func (tp *TypeParameter) WriteTo(w io.Writer) (n int64, err error) {
	for _, annotation := range tp.Annotations {
		if an, aerr := annotation.WriteTo(w); aerr != nil {
			err = aerr
			return
		} else {
			n += an
		}
	}
	if nn, nerr := w.Write([]byte(tp.Name)); nerr != nil {
		err = nerr
		return
	} else {
		n += int64(nn)
	}
	if alen := len(tp.Annotations); alen > 0 {
		if en, eerr := w.Write([]byte(`extends`)); eerr != nil {
			err = eerr
			return
		} else {
			n += int64(en)
		}
		for i := 0; i < alen-1; i++ {
			if bn, berr := tp.Bounds[i].WriteTo(w); berr != nil {
				err = berr
				return
			} else {
				n += bn
			}
			if cn, cerr := w.Write([]byte(`,`)); cerr != nil {
				err = cerr
				return
			} else {
				n += int64(cn)
			}
		}
		if bn, berr := tp.Bounds[alen-1].WriteTo(w); berr != nil {
			err = berr
			return
		} else {
			n += bn
		}
	}
	return
}

// Implements [io.WriterTo] interface for [Variable].
func (v *Variable) WriteTo(w io.Writer) (n int64, err error) {
	if mn, merr := v.Modifiers.WriteTo(w); merr != nil {
		err = merr
		return
	} else {
		n += mn
	}
	if tn, terr := v.Type.WriteTo(w); terr != nil {
		err = terr
		return
	} else {
		n += tn
	}
	if nn, nerr := w.Write([]byte(v.Name)); nerr != nil {
		err = nerr
		return
	} else {
		n += int64(nn)
	}
	if v.Initializer != nil {
		if en, eerr := w.Write([]byte(`=`)); eerr != nil {
			err = eerr
			return
		} else {
			n += int64(en)
		}
		if in, ierr := (*v.Initializer).WriteTo(w); ierr != nil {
			err = ierr
			return
		} else {
			n += in
		}
	}
	if sn, serr := w.Write([]byte(`;`)); serr != nil {
		err = serr
		return
	} else {
		n += int64(sn)
	}
	return
}

// Implements [io.WriterTo] interface for [WhileLoop].
func (wl *WhileLoop) WriteTo(w io.Writer) (n int64, err error) {
	if wn, werr := w.Write([]byte(`while`)); werr != nil {
		err = werr
		return
	} else {
		n += int64(wn)
	}
	if on, oerr := w.Write([]byte(`(`)); oerr != nil {
		err = oerr
		return
	} else {
		n += int64(on)
	}
	if cn, cerr := wl.Condition.WriteTo(w); cerr != nil {
		err = cerr
		return
	} else {
		n += cn
	}
	if cn, cerr := w.Write([]byte(`)`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	if sn, serr := wl.Statement.WriteTo(w); serr != nil {
		err = serr
		return
	} else {
		n += sn
	}
	return
}

// Implements [io.WriterTo] interface for [PostfixIncrement].
func (pi *PostfixIncrement) WriteTo(w io.Writer) (n int64, err error) {
	if xn, xerr := pi.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	if in, ierr := w.Write([]byte(`++`)); ierr != nil {
		err = ierr
		return
	} else {
		n += int64(in)
	}
	return
}

// Implements [io.WriterTo] interface for [PostfixDecrement].
func (pd *PostfixDecrement) WriteTo(w io.Writer) (n int64, err error) {
	if xn, xerr := pd.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	if in, ierr := w.Write([]byte(`--`)); ierr != nil {
		err = ierr
		return
	} else {
		n += int64(in)
	}
	return
}

// Implements [io.WriterTo] interface for [PrefixIncrement].
func (pi *PrefixIncrement) WriteTo(w io.Writer) (n int64, err error) {
	if dn, derr := w.Write([]byte(`++`)); derr != nil {
		err = derr
		return
	} else {
		n += int64(dn)
	}
	if xn, xerr := pi.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	return
}

// Implements [io.WriterTo] interface for [PrefixDecrement].
func (pd *PrefixDecrement) WriteTo(w io.Writer) (n int64, err error) {
	if dn, derr := w.Write([]byte(`--`)); derr != nil {
		err = derr
		return
	} else {
		n += int64(dn)
	}
	if xn, xerr := pd.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	return
}

// Implements [io.WriterTo] interface for [UnaryPlus].
func (up *UnaryPlus) WriteTo(w io.Writer) (n int64, err error) {
	if pn, perr := w.Write([]byte(`+`)); perr != nil {
		err = perr
		return
	} else {
		n += int64(pn)
	}
	if xn, xerr := up.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	return
}

// Implements [io.WriterTo] interface for [UnaryMinus].
func (um *UnaryMinus) WriteTo(w io.Writer) (n int64, err error) {
	if mn, merr := w.Write([]byte(`-`)); merr != nil {
		err = merr
		return
	} else {
		n += int64(mn)
	}
	if xn, xerr := um.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	return
}

// Implements [io.WriterTo] interface for [BitwiseComplement].
func (bc *BitwiseComplement) WriteTo(w io.Writer) (n int64, err error) {
	if bcn, bcerr := w.Write([]byte(`~`)); bcerr != nil {
		err = bcerr
		return
	} else {
		n += int64(bcn)
	}
	if xn, xerr := bc.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	return
}

// Implements [io.WriterTo] interface for [LogicalComplement].
func (lc *LogicalComplement) WriteTo(w io.Writer) (n int64, err error) {
	if lcn, lcerr := w.Write([]byte(`!`)); lcerr != nil {
		err = lcerr
		return
	} else {
		n += int64(lcn)
	}
	if xn, xerr := lc.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	return
}

// Implements [io.WriterTo] interface for [Multiply].
func (m *Multiply) WriteTo(w io.Writer) (n int64, err error) {
	if lon, loerr := m.LeftOperand.WriteTo(w); loerr != nil {
		err = loerr
		return
	} else {
		n += lon
	}
	if mn, merr := w.Write([]byte(`*`)); merr != nil {
		err = merr
		return
	} else {
		n += int64(mn)
	}
	if ron, roerr := m.RightOperand.WriteTo(w); roerr != nil {
		err = roerr
		return
	} else {
		n += ron
	}
	return
}

// Implements [io.WriterTo] interface for [Divide].
func (d *Divide) WriteTo(w io.Writer) (n int64, err error) {
	if lon, loerr := d.LeftOperand.WriteTo(w); loerr != nil {
		err = loerr
		return
	} else {
		n += lon
	}
	if dn, derr := w.Write([]byte(`/`)); derr != nil {
		err = derr
		return
	} else {
		n += int64(dn)
	}
	if ron, roerr := d.RightOperand.WriteTo(w); roerr != nil {
		err = roerr
		return
	} else {
		n += ron
	}
	return
}

// Implements [io.WriterTo] interface for [Remainder].
func (r *Remainder) WriteTo(w io.Writer) (n int64, err error) {
	if lon, loerr := r.LeftOperand.WriteTo(w); loerr != nil {
		err = loerr
		return
	} else {
		n += lon
	}
	if pn, perr := w.Write([]byte(`%`)); perr != nil {
		err = perr
		return
	} else {
		n += int64(pn)
	}
	if ron, roerr := r.RightOperand.WriteTo(w); roerr != nil {
		err = roerr
		return
	} else {
		n += ron
	}
	return
}

// Implements [io.WriterTo] interface for [Plus].
func (p *Plus) WriteTo(w io.Writer) (n int64, err error) {
	if lon, loerr := p.LeftOperand.WriteTo(w); loerr != nil {
		err = loerr
		return
	} else {
		n += lon
	}
	if pn, perr := w.Write([]byte(`+`)); perr != nil {
		err = perr
		return
	} else {
		n += int64(pn)
	}
	if ron, roerr := p.RightOperand.WriteTo(w); roerr != nil {
		err = roerr
		return
	} else {
		n += ron
	}
	return
}

// Implements [io.WriterTo] interface for [Minus].
func (m *Minus) WriteTo(w io.Writer) (n int64, err error) {
	if lon, loerr := m.LeftOperand.WriteTo(w); loerr != nil {
		err = loerr
		return
	} else {
		n += lon
	}
	if mn, merr := w.Write([]byte(`-`)); merr != nil {
		err = merr
		return
	} else {
		n += int64(mn)
	}
	if ron, roerr := m.RightOperand.WriteTo(w); roerr != nil {
		err = roerr
		return
	} else {
		n += ron
	}
	return
}

// Implements [io.WriterTo] interface for [LeftShift].
func (ls *LeftShift) WriteTo(w io.Writer) (n int64, err error) {
	if lon, loerr := ls.LeftOperand.WriteTo(w); loerr != nil {
		err = loerr
		return
	} else {
		n += lon
	}
	if lsn, lserr := w.Write([]byte(`<<`)); lserr != nil {
		err = lserr
		return
	} else {
		n += int64(lsn)
	}
	if ron, roerr := ls.RightOperand.WriteTo(w); roerr != nil {
		err = roerr
		return
	} else {
		n += ron
	}
	return
}

// Implements [io.WriterTo] interface for [RightShift].
func (rs *RightShift) WriteTo(w io.Writer) (n int64, err error) {
	if lon, loerr := rs.LeftOperand.WriteTo(w); loerr != nil {
		err = loerr
		return
	} else {
		n += lon
	}
	if rsn, rserr := w.Write([]byte(`>>`)); rserr != nil {
		err = rserr
		return
	} else {
		n += int64(rsn)
	}
	if ron, roerr := rs.RightOperand.WriteTo(w); roerr != nil {
		err = roerr
		return
	} else {
		n += ron
	}
	return
}

// Implements [io.WriterTo] interface for [UnsignedRightShift].
func (urs *UnsignedRightShift) WriteTo(w io.Writer) (n int64, err error) {
	if lon, loerr := urs.LeftOperand.WriteTo(w); loerr != nil {
		err = loerr
		return
	} else {
		n += lon
	}
	if ursn, urserr := w.Write([]byte(`>>>`)); urserr != nil {
		err = urserr
		return
	} else {
		n += int64(ursn)
	}
	if ron, roerr := urs.RightOperand.WriteTo(w); roerr != nil {
		err = roerr
		return
	} else {
		n += ron
	}
	return
}

// Implements [io.WriterTo] interface for [LessThan].
func (lt *LessThan) WriteTo(w io.Writer) (n int64, err error) {
	if lon, loerr := lt.LeftOperand.WriteTo(w); loerr != nil {
		err = loerr
		return
	} else {
		n += lon
	}
	if ltn, lterr := w.Write([]byte(`<`)); lterr != nil {
		err = lterr
		return
	} else {
		n += int64(ltn)
	}
	if ron, roerr := lt.RightOperand.WriteTo(w); roerr != nil {
		err = roerr
		return
	} else {
		n += ron
	}
	return
}

// Implements [io.WriterTo] interface for [GreaterThan].
func (gt *GreaterThan) WriteTo(w io.Writer) (n int64, err error) {
	if lon, loerr := gt.LeftOperand.WriteTo(w); loerr != nil {
		err = loerr
		return
	} else {
		n += lon
	}
	if gtn, gterr := w.Write([]byte(`>`)); gterr != nil {
		err = gterr
		return
	} else {
		n += int64(gtn)
	}
	if ron, roerr := gt.RightOperand.WriteTo(w); roerr != nil {
		err = roerr
		return
	} else {
		n += ron
	}
	return
}

// Implements [io.WriterTo] interface for [LessThanEqual].
func (lte *LessThanEqual) WriteTo(w io.Writer) (n int64, err error) {
	if lon, loerr := lte.LeftOperand.WriteTo(w); loerr != nil {
		err = loerr
		return
	} else {
		n += lon
	}
	if lten, lteerr := w.Write([]byte(`<=`)); lteerr != nil {
		err = lteerr
		return
	} else {
		n += int64(lten)
	}
	if ron, roerr := lte.RightOperand.WriteTo(w); roerr != nil {
		err = roerr
		return
	} else {
		n += ron
	}
	return
}

// Implements [io.WriterTo] interface for [GreaterThanEqual].
func (gte *GreaterThanEqual) WriteTo(w io.Writer) (n int64, err error) {
	if lon, loerr := gte.LeftOperand.WriteTo(w); loerr != nil {
		err = loerr
		return
	} else {
		n += lon
	}
	if gten, gteerr := w.Write([]byte(`>=`)); gteerr != nil {
		err = gteerr
		return
	} else {
		n += int64(gten)
	}
	if ron, roerr := gte.RightOperand.WriteTo(w); roerr != nil {
		err = roerr
		return
	} else {
		n += ron
	}
	return
}

// Implements [io.WriterTo] interface for [EqualTo].
func (eq *EqualTo) WriteTo(w io.Writer) (n int64, err error) {
	if lon, loerr := eq.LeftOperand.WriteTo(w); loerr != nil {
		err = loerr
		return
	} else {
		n += lon
	}
	if eqn, eqerr := w.Write([]byte(`==`)); eqerr != nil {
		err = eqerr
		return
	} else {
		n += int64(eqn)
	}
	if ron, roerr := eq.RightOperand.WriteTo(w); roerr != nil {
		err = roerr
		return
	} else {
		n += ron
	}
	return
}

// Implements [io.WriterTo] interface for [NotEqualTo].
func (neq *NotEqualTo) WriteTo(w io.Writer) (n int64, err error) {
	if lon, loerr := neq.LeftOperand.WriteTo(w); loerr != nil {
		err = loerr
		return
	} else {
		n += lon
	}
	if neqn, neqerr := w.Write([]byte(`!=`)); neqerr != nil {
		err = neqerr
		return
	} else {
		n += int64(neqn)
	}
	if ron, roerr := neq.RightOperand.WriteTo(w); roerr != nil {
		err = roerr
		return
	} else {
		n += ron
	}
	return
}

// Implements [io.WriterTo] interface for [And].
func (a *And) WriteTo(w io.Writer) (n int64, err error) {
	if lon, loerr := a.LeftOperand.WriteTo(w); loerr != nil {
		err = loerr
		return
	} else {
		n += lon
	}
	if an, aerr := w.Write([]byte(`&`)); aerr != nil {
		err = aerr
		return
	} else {
		n += int64(an)
	}
	if ron, roerr := a.RightOperand.WriteTo(w); roerr != nil {
		err = roerr
		return
	} else {
		n += ron
	}
	return
}

// Implements [io.WriterTo] interface for [Xor].
func (x *Xor) WriteTo(w io.Writer) (n int64, err error) {
	if lon, loerr := x.LeftOperand.WriteTo(w); loerr != nil {
		err = loerr
		return
	} else {
		n += lon
	}
	if xn, xerr := w.Write([]byte(`^`)); xerr != nil {
		err = xerr
		return
	} else {
		n += int64(xn)
	}
	if ron, roerr := x.RightOperand.WriteTo(w); roerr != nil {
		err = roerr
		return
	} else {
		n += ron
	}
	return
}

// Implements [io.WriterTo] interface for [Or].
func (or *Or) WriteTo(w io.Writer) (n int64, err error) {
	if lon, loerr := or.LeftOperand.WriteTo(w); loerr != nil {
		err = loerr
		return
	} else {
		n += lon
	}
	if pn, perr := w.Write([]byte(`|`)); perr != nil {
		err = perr
		return
	} else {
		n += int64(pn)
	}
	if ron, roerr := or.RightOperand.WriteTo(w); roerr != nil {
		err = roerr
		return
	} else {
		n += ron
	}
	return
}

// Implements [io.WriterTo] interface for [ConditionalAnd].
func (ca *ConditionalAnd) WriteTo(w io.Writer) (n int64, err error) {
	if lon, loerr := ca.LeftOperand.WriteTo(w); loerr != nil {
		err = loerr
		return
	} else {
		n += lon
	}
	if an, aerr := w.Write([]byte(`&&`)); aerr != nil {
		err = aerr
		return
	} else {
		n += int64(an)
	}
	if ron, roerr := ca.RightOperand.WriteTo(w); roerr != nil {
		err = roerr
		return
	} else {
		n += ron
	}
	return
}

// Implements [io.WriterTo] interface for [ConditionalOr].
func (or *ConditionalOr) WriteTo(w io.Writer) (n int64, err error) {
	if lon, loerr := or.LeftOperand.WriteTo(w); loerr != nil {
		err = loerr
		return
	} else {
		n += lon
	}
	if pn, perr := w.Write([]byte(`||`)); perr != nil {
		err = perr
		return
	} else {
		n += int64(pn)
	}
	if ron, roerr := or.RightOperand.WriteTo(w); roerr != nil {
		err = roerr
		return
	} else {
		n += ron
	}
	return
}

// Implements [io.WriterTo] interface for [MultiplyAssignment].
func (ma *MultiplyAssignment) WriteTo(w io.Writer) (n int64, err error) {
	if vn, verr := ma.Variable.WriteTo(w); verr != nil {
		err = verr
		return
	} else {
		n += vn
	}
	if man, maerr := w.Write([]byte(`*=`)); maerr != nil {
		err = maerr
		return
	} else {
		n += int64(man)
	}
	if xn, xerr := ma.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	return
}

// Implements [io.WriterTo] interface for [DivideAssignment].
func (da *DivideAssignment) WriteTo(w io.Writer) (n int64, err error) {
	if vn, verr := da.Variable.WriteTo(w); verr != nil {
		err = verr
		return
	} else {
		n += vn
	}
	if dan, daerr := w.Write([]byte(`/=`)); daerr != nil {
		err = daerr
		return
	} else {
		n += int64(dan)
	}
	if xn, xerr := da.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	return
}

// Implements [io.WriterTo] interface for [RemainderAssignment].
func (ra *RemainderAssignment) WriteTo(w io.Writer) (n int64, err error) {
	if vn, verr := ra.Variable.WriteTo(w); verr != nil {
		err = verr
		return
	} else {
		n += vn
	}
	if ran, raerr := w.Write([]byte(`%=`)); raerr != nil {
		err = raerr
		return
	} else {
		n += int64(ran)
	}
	if xn, xerr := ra.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	return
}

// Implements [io.WriterTo] interface for [PlusAssignment].
func (pa *PlusAssignment) WriteTo(w io.Writer) (n int64, err error) {
	if vn, verr := pa.Variable.WriteTo(w); verr != nil {
		err = verr
		return
	} else {
		n += vn
	}
	if pan, paerr := w.Write([]byte(`+=`)); paerr != nil {
		err = paerr
		return
	} else {
		n += int64(pan)
	}
	if xn, xerr := pa.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	return
}

// Implements [io.WriterTo] interface for [MinusAssignment].
func (ma *MinusAssignment) WriteTo(w io.Writer) (n int64, err error) {
	if vn, verr := ma.Variable.WriteTo(w); verr != nil {
		err = verr
		return
	} else {
		n += vn
	}
	if man, maerr := w.Write([]byte(`-=`)); maerr != nil {
		err = maerr
		return
	} else {
		n += int64(man)
	}
	if xn, xerr := ma.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	return
}

// Implements [io.WriterTo] interface for [LeftShiftAssignment].
func (lsa *LeftShiftAssignment) WriteTo(w io.Writer) (n int64, err error) {
	if vn, verr := lsa.Variable.WriteTo(w); verr != nil {
		err = verr
		return
	} else {
		n += vn
	}
	if lsan, lsaerr := w.Write([]byte(`<<=`)); lsaerr != nil {
		err = lsaerr
		return
	} else {
		n += int64(lsan)
	}
	if xn, xerr := lsa.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	return
}

// Implements [io.WriterTo] interface for [RightShiftAssignment].
func (rsa *RightShiftAssignment) WriteTo(w io.Writer) (n int64, err error) {
	if vn, verr := rsa.Variable.WriteTo(w); verr != nil {
		err = verr
		return
	} else {
		n += vn
	}
	if rsan, rsaerr := w.Write([]byte(`>>=`)); rsaerr != nil {
		err = rsaerr
		return
	} else {
		n += int64(rsan)
	}
	if xn, xerr := rsa.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	return
}

// Implements [io.WriterTo] interface for [UnsignedRightShiftAssignment].
func (ursa *UnsignedRightShiftAssignment) WriteTo(w io.Writer) (n int64, err error) {
	if vn, verr := ursa.Variable.WriteTo(w); verr != nil {
		err = verr
		return
	} else {
		n += vn
	}
	if ursan, ursaerr := w.Write([]byte(`>>=`)); ursaerr != nil {
		err = ursaerr
		return
	} else {
		n += int64(ursan)
	}
	if xn, xerr := ursa.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	return
}

// Implements [io.WriterTo] interface for [AndAssignment].
func (aa *AndAssignment) WriteTo(w io.Writer) (n int64, err error) {
	if vn, verr := aa.Variable.WriteTo(w); verr != nil {
		err = verr
		return
	} else {
		n += vn
	}
	if aan, aaerr := w.Write([]byte(`&=`)); aaerr != nil {
		err = aaerr
		return
	} else {
		n += int64(aan)
	}
	if xn, xerr := aa.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	return
}

// Implements [io.WriterTo] interface for [XorAssignment].
func (xa *XorAssignment) WriteTo(w io.Writer) (n int64, err error) {
	if vn, verr := xa.Variable.WriteTo(w); verr != nil {
		err = verr
		return
	} else {
		n += vn
	}
	if xan, xaerr := w.Write([]byte(`^=`)); xaerr != nil {
		err = xaerr
		return
	} else {
		n += int64(xan)
	}
	if xn, xerr := xa.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	return
}

// Implements [io.WriterTo] interface for [OrAssignment].
func (oa *OrAssignment) WriteTo(w io.Writer) (n int64, err error) {
	if vn, verr := oa.Variable.WriteTo(w); verr != nil {
		err = verr
		return
	} else {
		n += vn
	}
	if oan, oaerr := w.Write([]byte(`|=`)); oaerr != nil {
		err = oaerr
		return
	} else {
		n += int64(oan)
	}
	if xn, xerr := oa.Expression.WriteTo(w); xerr != nil {
		err = xerr
		return
	} else {
		n += xn
	}
	return
}

// Implements [io.WriterTo] interface for [IntLiteral].
func (il *IntLiteral) WriteTo(w io.Writer) (n int64, err error) {
	if iln, ilerr := w.Write([]byte(il.Value)); ilerr != nil {
		err = ilerr
		return
	} else {
		n += int64(iln)
	}
	return
}

// Implements [io.WriterTo] interface for [LongLiteral].
func (ll *LongLiteral) WriteTo(w io.Writer) (n int64, err error) {
	if lln, llerr := w.Write([]byte(ll.Value)); llerr != nil {
		err = llerr
		return
	} else {
		n += int64(lln)
	}
	return
}

// Implements [io.WriterTo] interface for [FloatLiteral].
func (fl *FloatLiteral) WriteTo(w io.Writer) (n int64, err error) {
	if fln, flerr := w.Write([]byte(fl.Value)); flerr != nil {
		err = flerr
		return
	} else {
		n += int64(fln)
	}
	return
}

// Implements [io.WriterTo] interface for [DoubleLiteral].
func (dl *DoubleLiteral) WriteTo(w io.Writer) (n int64, err error) {
	if dln, dlerr := w.Write([]byte(dl.Value)); dlerr != nil {
		err = dlerr
		return
	} else {
		n += int64(dln)
	}
	return
}

// Implements [io.WriterTo] interface for [BooleanLiteral].
func (bl *BooleanLiteral) WriteTo(w io.Writer) (n int64, err error) {
	if bl.Value {
		if bln, blerr := w.Write([]byte(`true`)); blerr != nil {
			err = blerr
			return
		} else {
			n += int64(bln)
		}
		return
	}
	if bln, blerr := w.Write([]byte(`false`)); blerr != nil {
		err = blerr
		return
	} else {
		n += int64(bln)
	}
	return
}

// Implements [io.WriterTo] interface for [CharLiteral].
func (cl *CharLiteral) WriteTo(w io.Writer) (n int64, err error) {
	if cln, clerr := w.Write([]byte(`'` + cl.Value + `'`)); clerr != nil {
		err = clerr
		return
	} else {
		n += int64(cln)
	}
	return
}

// Implements [io.WriterTo] interface for [StringLiteral].
func (sl *StringLiteral) WriteTo(w io.Writer) (n int64, err error) {
	if sln, slerr := w.Write([]byte(`"` + sl.Value + `"`)); slerr != nil {
		err = slerr
		return
	} else {
		n += int64(sln)
	}
	return
}

// Implements [io.WriterTo] interface for [NullLiteral].
func (*NullLiteral) WriteTo(w io.Writer) (n int64, err error) {
	if nln, nlerr := w.Write([]byte(`null`)); nlerr != nil {
		err = nlerr
		return
	} else {
		n += int64(nln)
	}
	return
}

// Implements [io.WriterTo] interface for [UnboundedWildcard].
func (*UnboundedWildcard) WriteTo(w io.Writer) (n int64, err error) {
	if qn, qerr := w.Write([]byte(`?`)); qerr != nil {
		err = qerr
		return
	} else {
		n += int64(qn)
	}
	return
}

// Implements [io.WriterTo] interface for [ExtendsWildcard].
func (xw *ExtendsWildcard) WriteTo(w io.Writer) (n int64, err error) {
	if qn, qerr := w.Write([]byte(`?`)); qerr != nil {
		err = qerr
		return
	} else {
		n += int64(qn)
	}
	if en, eerr := w.Write([]byte(`extends`)); eerr != nil {
		err = eerr
		return
	} else {
		n += int64(en)
	}
	if bn, berr := xw.Bound.WriteTo(w); berr != nil {
		err = berr
		return
	} else {
		n += bn
	}
	return
}

// Implements [io.WriterTo] interface for [SuperWildcard].
func (sw *SuperWildcard) WriteTo(w io.Writer) (n int64, err error) {
	if qn, qerr := w.Write([]byte(`?`)); qerr != nil {
		err = qerr
		return
	} else {
		n += int64(qn)
	}
	if sn, serr := w.Write([]byte(`super`)); serr != nil {
		err = serr
		return
	} else {
		n += int64(sn)
	}
	if bn, berr := sw.Bound.WriteTo(w); berr != nil {
		err = berr
		return
	} else {
		n += bn
	}
	return
}

// Implements [io.WriterTo] interface for [Erroneous].
func (*Erroneous) WriteTo(w io.Writer) (n int64, err error) {
	err = fmt.Errorf("unexpected erroneous")
	return
}

// Implements [io.WriterTo] interface for [Interface].
func (i *Interface) WriteTo(w io.Writer) (n int64, err error) {
	if mn, merr := i.Modifiers.WriteTo(w); merr != nil {
		err = merr
		return
	} else {
		n += mn
	}
	if cn, cerr := w.Write([]byte(`interface`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	if snn, snerr := w.Write([]byte(i.SimpleName)); snerr != nil {
		err = snerr
		return
	} else {
		n += int64(snn)
	}
	if tplen := len(i.TypeParameters); tplen > 0 {
		if on, oerr := w.Write([]byte(`<`)); oerr != nil {
			err = oerr
			return
		} else {
			n += int64(on)
		}
		for j := 0; j < tplen-1; j++ {
			if tpn, tperr := i.TypeParameters[j].WriteTo(w); tperr != nil {
				err = tperr
				return
			} else {
				n += tpn
			}
			if cn, cerr := w.Write([]byte(`,`)); cerr != nil {
				err = cerr
				return
			} else {
				n += int64(cn)
			}
		}
		if tpn, tperr := i.TypeParameters[tplen-1].WriteTo(w); tperr != nil {
			err = tperr
			return
		} else {
			n += tpn
		}
		if cn, cerr := w.Write([]byte(`>`)); cerr != nil {
			err = cerr
			return
		} else {
			n += int64(cn)
		}
	}
	if i.ExtendsClause != nil {
		if en, eerr := w.Write([]byte(`extends`)); eerr != nil {
			err = eerr
			return
		} else {
			n += int64(en)
		}
		if ecn, ecerr := (*i.ExtendsClause).WriteTo(w); ecerr != nil {
			err = ecerr
			return
		} else {
			n += ecn
		}
	}
	if pclen := len(i.PermitsClause); pclen > 0 {
		if pn, perr := w.Write([]byte(`permits`)); perr != nil {
			err = perr
			return
		} else {
			n += int64(pn)
		}
		for j := 0; j < pclen-1; j++ {
			if pcn, pcerr := i.PermitsClause[j].WriteTo(w); pcerr != nil {
				err = pcerr
				return
			} else {
				n += pcn
			}
			if cn, cerr := w.Write([]byte(`,`)); cerr != nil {
				err = cerr
				return
			} else {
				n += int64(cn)
			}
		}
		if pcn, pcerr := i.PermitsClause[pclen-1].WriteTo(w); pcerr != nil {
			err = pcerr
			return
		} else {
			n += pcn
		}
	}
	if on, oerr := w.Write([]byte(`{`)); oerr != nil {
		err = oerr
		return
	} else {
		n += int64(on)
	}
	for _, member := range i.Members {
		if mn, merr := member.WriteTo(w); merr != nil {
			err = merr
			return
		} else {
			n += mn
		}
	}
	if cn, cerr := w.Write([]byte(`}`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	return
}

// Implements [io.WriterTo] interface for [Enum].
func (e *Enum) WriteTo(w io.Writer) (n int64, err error) {
	if mn, merr := e.Modifiers.WriteTo(w); merr != nil {
		err = merr
		return
	} else {
		n += mn
	}
	if cn, cerr := w.Write([]byte(`enum`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	if snn, snerr := w.Write([]byte(e.SimpleName)); snerr != nil {
		err = snerr
		return
	} else {
		n += int64(snn)
	}
	if on, oerr := w.Write([]byte(`{`)); oerr != nil {
		err = oerr
		return
	} else {
		n += int64(on)
	}
	for _, member := range e.Members {
		if mn, merr := member.WriteTo(w); merr != nil {
			err = merr
			return
		} else {
			n += mn
		}
	}
	if cn, cerr := w.Write([]byte(`}`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	return
}

// Implements [io.WriterTo] interface for [AnnotationType].
func (at *AnnotationType) WriteTo(w io.Writer) (n int64, err error) {
	if mn, merr := at.Modifiers.WriteTo(w); merr != nil {
		err = merr
		return
	} else {
		n += mn
	}
	if cn, cerr := w.Write([]byte(`@interface`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	if snn, snerr := w.Write([]byte(at.SimpleName)); snerr != nil {
		err = snerr
		return
	} else {
		n += int64(snn)
	}
	if on, oerr := w.Write([]byte(`{`)); oerr != nil {
		err = oerr
		return
	} else {
		n += int64(on)
	}
	for _, member := range at.Members {
		if mn, merr := member.WriteTo(w); merr != nil {
			err = merr
			return
		} else {
			n += mn
		}
	}
	if cn, cerr := w.Write([]byte(`}`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	return
}

// Implements [io.WriterTo] interface for [Module].
func (m *Module) WriteTo(w io.Writer) (n int64, err error) {
	for _, annotation := range m.Annotations {
		if an, aerr := annotation.WriteTo(w); aerr != nil {
			err = aerr
			return
		} else {
			n += an
		}
	}
	if m.ModuleType == OPEN_MODULE_KIND {
		if on, oerr := w.Write([]byte(`open`)); oerr != nil {
			err = oerr
			return
		} else {
			n += int64(on)
		}
	}
	if mn, merr := w.Write([]byte(`module`)); merr != nil {
		err = merr
		return
	} else {
		n += int64(mn)
	}
	if nn, nerr := m.Name.WriteTo(w); nerr != nil {
		err = nerr
		return
	} else {
		n += nn
	}
	if on, oerr := w.Write([]byte(`{`)); oerr != nil {
		err = oerr
		return
	} else {
		n += int64(on)
	}
	for _, directive := range m.Directives {
		if dn, derr := directive.WriteTo(w); derr != nil {
			err = derr
			return
		} else {
			n += dn
		}
	}
	if cn, cerr := w.Write([]byte(`}`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	return
}

// Implements [io.WriterTo] interface for [Exports].
func (x *Exports) WriteTo(w io.Writer) (n int64, err error) {
	if xn, xerr := w.Write([]byte(`exports`)); xerr != nil {
		err = xerr
		return
	} else {
		n += int64(xn)
	}
	if pnn, pnerr := x.PackageName.WriteTo(w); pnerr != nil {
		err = pnerr
		return
	} else {
		n += int64(pnn)
	}
	if mnlen := len(x.ModuleNames); mnlen > 0 {
		if tn, terr := w.Write([]byte(`to`)); terr != nil {
			err = terr
			return
		} else {
			n += int64(tn)
		}
		for i := 0; i < mnlen-1; i++ {
			if mnn, mnerr := x.ModuleNames[i].WriteTo(w); mnerr != nil {
				err = mnerr
				return
			} else {
				n += mnn
			}
			if cn, cerr := w.Write([]byte(`,`)); cerr != nil {
				err = cerr
				return
			} else {
				n += int64(cn)
			}
		}
		if mnn, mnerr := x.ModuleNames[mnlen-1].WriteTo(w); mnerr != nil {
			err = mnerr
			return
		} else {
			n += mnn
		}
	}
	return
}

// Implements [io.WriterTo] interface for [Opens].
func (o *Opens) WriteTo(w io.Writer) (n int64, err error) {
	if on, oerr := w.Write([]byte(`opens`)); oerr != nil {
		err = oerr
		return
	} else {
		n += int64(on)
	}
	if pnn, pnerr := o.PackageName.WriteTo(w); pnerr != nil {
		err = pnerr
		return
	} else {
		n += int64(pnn)
	}
	if mnlen := len(o.ModuleNames); mnlen > 0 {
		if tn, terr := w.Write([]byte(`to`)); terr != nil {
			err = terr
			return
		} else {
			n += int64(tn)
		}
		for i := 0; i < mnlen-1; i++ {
			if mnn, mnerr := o.ModuleNames[i].WriteTo(w); mnerr != nil {
				err = mnerr
				return
			} else {
				n += mnn
			}
			if cn, cerr := w.Write([]byte(`,`)); cerr != nil {
				err = cerr
				return
			} else {
				n += int64(cn)
			}
		}
		if mnn, mnerr := o.ModuleNames[mnlen-1].WriteTo(w); mnerr != nil {
			err = mnerr
			return
		} else {
			n += mnn
		}
	}
	return
}

// Implements [io.WriterTo] interface for [Provides].
func (p *Provides) WriteTo(w io.Writer) (n int64, err error) {
	if pn, perr := w.Write([]byte(`provides`)); perr != nil {
		err = perr
		return
	} else {
		n += int64(pn)
	}
	if snn, snerr := p.ServiceName.WriteTo(w); snerr != nil {
		err = snerr
		return
	} else {
		n += int64(snn)
	}
	if inlen := len(p.ImplementationNames); inlen > 0 {
		if tn, terr := w.Write([]byte(`with`)); terr != nil {
			err = terr
			return
		} else {
			n += int64(tn)
		}
		for i := 0; i < inlen-1; i++ {
			if inn, inerr := p.ImplementationNames[i].WriteTo(w); inerr != nil {
				err = inerr
				return
			} else {
				n += inn
			}
			if cn, cerr := w.Write([]byte(`,`)); cerr != nil {
				err = cerr
				return
			} else {
				n += int64(cn)
			}
		}
		if inn, inerr := p.ImplementationNames[inlen-1].WriteTo(w); inerr != nil {
			err = inerr
			return
		} else {
			n += inn
		}
	}
	return
}

// Implements [io.WriterTo] interface for [Record].
func (r *Record) WriteTo(w io.Writer) (n int64, err error) {
	if mn, merr := r.Modifiers.WriteTo(w); merr != nil {
		err = merr
		return
	} else {
		n += mn
	}
	if cn, cerr := w.Write([]byte(`record`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	if snn, snerr := w.Write([]byte(r.SimpleName)); snerr != nil {
		err = snerr
		return
	} else {
		n += int64(snn)
	}
	if iclen := len(r.ImplementsClause); iclen > 0 {
		if in, ierr := w.Write([]byte(`implements`)); ierr != nil {
			err = ierr
			return
		} else {
			n += int64(in)
		}
		for i := 0; i < iclen-1; i++ {
			if icn, icerr := r.ImplementsClause[i].WriteTo(w); icerr != nil {
				err = icerr
				return
			} else {
				n += icn
			}
			if cn, cerr := w.Write([]byte(`,`)); cerr != nil {
				err = cerr
				return
			} else {
				n += int64(cn)
			}
		}
		if icn, icerr := r.ImplementsClause[iclen-1].WriteTo(w); icerr != nil {
			err = icerr
			return
		} else {
			n += icn
		}
	}
	if on, oerr := w.Write([]byte(`{`)); oerr != nil {
		err = oerr
		return
	} else {
		n += int64(on)
	}
	for _, member := range r.Members {
		if mn, merr := member.WriteTo(w); merr != nil {
			err = merr
			return
		} else {
			n += mn
		}
	}
	if cn, cerr := w.Write([]byte(`}`)); cerr != nil {
		err = cerr
		return
	} else {
		n += int64(cn)
	}
	return
}

// Implements [io.WriterTo] interface for [Requires].
func (r *Requires) WriteTo(w io.Writer) (n int64, err error) {
	if rn, rerr := w.Write([]byte(`requires`)); rerr != nil {
		err = rerr
		return
	} else {
		n += int64(rn)
	}
	if r.Static {
		if sn, serr := w.Write([]byte(`static`)); serr != nil {
			err = serr
			return
		} else {
			n += int64(sn)
		}
	}
	if r.Transitive {
		if tn, terr := w.Write([]byte(`transitive`)); terr != nil {
			err = terr
			return
		} else {
			n += int64(tn)
		}
	}
	if mnn, mnerr := r.ModuleName.WriteTo(w); mnerr != nil {
		err = mnerr
		return
	} else {
		n += int64(mnn)
	}
	return
}

// Implements [io.WriterTo] interface for [Uses].
func (u *Uses) WriteTo(w io.Writer) (n int64, err error) {
	if un, uerr := w.Write([]byte(`uses`)); uerr != nil {
		err = uerr
		return
	} else {
		n += int64(un)
	}
	if snn, snerr := u.ServiceName.WriteTo(w); snerr != nil {
		err = snerr
		return
	} else {
		n += int64(snn)
	}
	return
}

// Implements [io.WriterTo] interface for [Yield].
func (y *Yield) WriteTo(w io.Writer) (n int64, err error) {
	if yn, yerr := w.Write([]byte(`yield`)); yerr != nil {
		err = yerr
		return
	} else {
		n += int64(yn)
	}
	if vn, verr := y.Value.WriteTo(w); verr != nil {
		err = verr
		return
	} else {
		n += vn
	}
	if sn, serr := w.Write([]byte(`;`)); serr != nil {
		err = serr
		return
	} else {
		n += int64(sn)
	}
	return
}
