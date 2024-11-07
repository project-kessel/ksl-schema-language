package semantic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssertSelfRelationExpressionToZanzibarSuccess(t *testing.T) {
	schema := NewSchema()
	namespace := NewNamespace("test_namespace", []string{})
	the_type := NewType("test_type", namespace, VisibilityPublic)
	other_type := NewType("other_type", namespace, VisibilityPublic)
	namespace.AddType(the_type)
	namespace.AddType(other_type)

	selfRelation := NewSelfRelationExpression([]*TypeReference{NewTypeReference(namespace.name, other_type.name, "", false)}, CardinalityAny)

	rel, err := NewRelation("test_relation", the_type, VisibilityPublic, selfRelation, nil)
	assert.NoError(t, err)
	schema.AddNamespace(namespace)
	the_type.AddRelation(rel)

	_, err = schema.ToZanzibar()
	assert.NoError(t, err)
}

func TestAssertSelfRelationExpressionToZanzibarTypeDoesNotExist(t *testing.T) {
	schema := NewSchema()
	namespace := NewNamespace("test_namespace", []string{})
	the_type := NewType("test_type", namespace, VisibilityPublic)
	//other_type := NewType("other_type", namespace, VisibilityPublic)

	selfRelation := NewSelfRelationExpression([]*TypeReference{NewTypeReference(namespace.name, "other", "", false)}, CardinalityAny)

	rel, err := NewRelation("test_relation", the_type, VisibilityPublic, selfRelation, nil)
	assert.NoError(t, err)
	namespace.AddType(the_type)
	schema.AddNamespace(namespace)
	the_type.AddRelation(rel)

	_, err = schema.ToZanzibar()
	assert.Error(t, err)
}

func TestAssertReferenceRelationExpressionToZanzibarSucceedsIfSubRelationIsNilAndRelationTypeDoesNotExist(t *testing.T) {
	schema := NewSchema()
	namespace := NewNamespace("test_namespace", []string{})
	the_type := NewType("test_type", namespace, VisibilityPublic)

	referenceRelation := NewReferenceRelationExpression("other", nil)

	rel, err := NewRelation("test_relation", the_type, VisibilityPublic, referenceRelation, nil)
	assert.NoError(t, err)
	namespace.AddType(the_type)
	schema.AddNamespace(namespace)
	the_type.AddRelation(rel)
	_, err = schema.ToZanzibar()
	assert.Error(t, err)
}
