// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package media

import (
	"github.com/briancabbott/entgo/dialect/sql"
	"github.com/briancabbott/entgo/entc/integration/migrate/entv2/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Source applies equality check predicate on the "source" field. It's identical to SourceEQ.
func Source(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSource), v))
	})
}

// SourceURI applies equality check predicate on the "source_uri" field. It's identical to SourceURIEQ.
func SourceURI(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSourceURI), v))
	})
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldText), v))
	})
}

// SourceEQ applies the EQ predicate on the "source" field.
func SourceEQ(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSource), v))
	})
}

// SourceNEQ applies the NEQ predicate on the "source" field.
func SourceNEQ(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSource), v))
	})
}

// SourceIn applies the In predicate on the "source" field.
func SourceIn(vs ...string) predicate.Media {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Media(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSource), v...))
	})
}

// SourceNotIn applies the NotIn predicate on the "source" field.
func SourceNotIn(vs ...string) predicate.Media {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Media(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSource), v...))
	})
}

// SourceGT applies the GT predicate on the "source" field.
func SourceGT(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSource), v))
	})
}

// SourceGTE applies the GTE predicate on the "source" field.
func SourceGTE(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSource), v))
	})
}

// SourceLT applies the LT predicate on the "source" field.
func SourceLT(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSource), v))
	})
}

// SourceLTE applies the LTE predicate on the "source" field.
func SourceLTE(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSource), v))
	})
}

// SourceContains applies the Contains predicate on the "source" field.
func SourceContains(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSource), v))
	})
}

// SourceHasPrefix applies the HasPrefix predicate on the "source" field.
func SourceHasPrefix(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSource), v))
	})
}

// SourceHasSuffix applies the HasSuffix predicate on the "source" field.
func SourceHasSuffix(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSource), v))
	})
}

// SourceIsNil applies the IsNil predicate on the "source" field.
func SourceIsNil() predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSource)))
	})
}

// SourceNotNil applies the NotNil predicate on the "source" field.
func SourceNotNil() predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSource)))
	})
}

// SourceEqualFold applies the EqualFold predicate on the "source" field.
func SourceEqualFold(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSource), v))
	})
}

// SourceContainsFold applies the ContainsFold predicate on the "source" field.
func SourceContainsFold(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSource), v))
	})
}

// SourceURIEQ applies the EQ predicate on the "source_uri" field.
func SourceURIEQ(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSourceURI), v))
	})
}

// SourceURINEQ applies the NEQ predicate on the "source_uri" field.
func SourceURINEQ(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSourceURI), v))
	})
}

// SourceURIIn applies the In predicate on the "source_uri" field.
func SourceURIIn(vs ...string) predicate.Media {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Media(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSourceURI), v...))
	})
}

// SourceURINotIn applies the NotIn predicate on the "source_uri" field.
func SourceURINotIn(vs ...string) predicate.Media {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Media(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSourceURI), v...))
	})
}

// SourceURIGT applies the GT predicate on the "source_uri" field.
func SourceURIGT(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSourceURI), v))
	})
}

// SourceURIGTE applies the GTE predicate on the "source_uri" field.
func SourceURIGTE(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSourceURI), v))
	})
}

// SourceURILT applies the LT predicate on the "source_uri" field.
func SourceURILT(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSourceURI), v))
	})
}

// SourceURILTE applies the LTE predicate on the "source_uri" field.
func SourceURILTE(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSourceURI), v))
	})
}

// SourceURIContains applies the Contains predicate on the "source_uri" field.
func SourceURIContains(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSourceURI), v))
	})
}

// SourceURIHasPrefix applies the HasPrefix predicate on the "source_uri" field.
func SourceURIHasPrefix(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSourceURI), v))
	})
}

// SourceURIHasSuffix applies the HasSuffix predicate on the "source_uri" field.
func SourceURIHasSuffix(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSourceURI), v))
	})
}

// SourceURIIsNil applies the IsNil predicate on the "source_uri" field.
func SourceURIIsNil() predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSourceURI)))
	})
}

// SourceURINotNil applies the NotNil predicate on the "source_uri" field.
func SourceURINotNil() predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSourceURI)))
	})
}

// SourceURIEqualFold applies the EqualFold predicate on the "source_uri" field.
func SourceURIEqualFold(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSourceURI), v))
	})
}

// SourceURIContainsFold applies the ContainsFold predicate on the "source_uri" field.
func SourceURIContainsFold(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSourceURI), v))
	})
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldText), v))
	})
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldText), v))
	})
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.Media {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Media(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldText), v...))
	})
}

// TextNotIn applies the NotIn predicate on the "text" field.
func TextNotIn(vs ...string) predicate.Media {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Media(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldText), v...))
	})
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldText), v))
	})
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldText), v))
	})
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldText), v))
	})
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldText), v))
	})
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldText), v))
	})
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldText), v))
	})
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldText), v))
	})
}

// TextIsNil applies the IsNil predicate on the "text" field.
func TextIsNil() predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldText)))
	})
}

// TextNotNil applies the NotNil predicate on the "text" field.
func TextNotNil() predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldText)))
	})
}

// TextEqualFold applies the EqualFold predicate on the "text" field.
func TextEqualFold(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldText), v))
	})
}

// TextContainsFold applies the ContainsFold predicate on the "text" field.
func TextContainsFold(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldText), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Media) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Media) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Media) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		p(s.Not())
	})
}
