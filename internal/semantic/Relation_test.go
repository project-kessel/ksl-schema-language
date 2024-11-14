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

	schemaDef, err := schema.ToZanzibar()
	assert.Equal(t, the_type.namespace.name+"/"+the_type.name, schemaDef[0].GetName())
	assert.Equal(t, other_type.namespace.name+"/"+other_type.name, schemaDef[1].GetName())
	assert.NoError(t, err)
}

func TestAssertSelfRelationExpressionToZanzibarTypeDoesNotExist(t *testing.T) {
	schema := NewSchema()
	namespace := NewNamespace("test_namespace", []string{})
	the_type := NewType("test_type", namespace, VisibilityPublic)

	selfRelation := NewSelfRelationExpression([]*TypeReference{NewTypeReference(namespace.name, "other", "", false)}, CardinalityAny)

	rel, err := NewRelation("test_relation", the_type, VisibilityPublic, selfRelation, nil)
	assert.NoError(t, err)
	namespace.AddType(the_type)
	schema.AddNamespace(namespace)
	the_type.AddRelation(rel)

	_, err = schema.ToZanzibar()
	assert.ErrorIs(t, err, ErrSymbolNotFound)
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
	assert.ErrorIs(t, err, ErrSymbolNotFound)
}

func TestAssertReferenceRelationExpressionToZanzibarFailsIfSubRelationIsTypo(t *testing.T) {
	schema := NewSchema()
	namespace := NewNamespace("test_namespace", []string{})
	principal := NewType("principal", namespace, VisibilityPublic)
	group := NewType("group", namespace, VisibilityPublic)

	principalTypeReference := NewTypeReference("", "principal", "", false)
	groupTypeReference := NewTypeReference("", "group", "membr", false)

	typeReferences := []*TypeReference{principalTypeReference, groupTypeReference}

	setrel := NewSelfRelationExpression(typeReferences, CardinalityAny)

	rel, err := NewRelation("member", group, VisibilityPublic, setrel, nil)
	assert.NoError(t, err)

	group.AddRelation(rel)
	namespace.AddType(principal)
	namespace.AddType(group)
	schema.AddNamespace(namespace)

	_, err = schema.ToZanzibar()
	assert.ErrorIs(t, err, ErrSymbolNotFound)
}
