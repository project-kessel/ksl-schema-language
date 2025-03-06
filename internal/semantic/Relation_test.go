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
	assert.Equal(t, other_type.namespace.name+"/"+other_type.name, schemaDef[0].GetName())
	assert.Equal(t, the_type.namespace.name+"/"+the_type.name, schemaDef[1].GetName())
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

func TestAssertReferenceRelationExpressionToZanzibarFailsIfSelfRelationIsTypo(t *testing.T) {
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

func TestAssertReferenceRelationExpressionToZanzibarSucceedsIfSelfRelationIsValid(t *testing.T) {
	schema := NewSchema()
	namespace := NewNamespace("test_namespace", []string{})
	principal := NewType("principal", namespace, VisibilityPublic)
	group := NewType("group", namespace, VisibilityPublic)

	principalTypeReference := NewTypeReference("", "principal", "", false)
	groupTypeReference := NewTypeReference("", "group", "member", false)

	typeReferences := []*TypeReference{principalTypeReference, groupTypeReference}

	setrel := NewSelfRelationExpression(typeReferences, CardinalityAny)

	rel, err := NewRelation("member", group, VisibilityPublic, setrel, nil)
	assert.NoError(t, err)

	group.AddRelation(rel)
	namespace.AddType(principal)
	namespace.AddType(group)
	schema.AddNamespace(namespace)

	_, err = schema.ToZanzibar()
	assert.NoError(t, err)
}

func TestAssertReferenceRelationExpressionToZanzibarFailsIfSubRelationTypo(t *testing.T) {
	schema := NewSchema()
	namespace := NewNamespace("test_namespace", []string{})
	principal := NewType("principal", namespace, VisibilityPublic)
	workspace := NewType("workspace", namespace, VisibilityPublic)

	workspaceTypeReference := NewTypeReference("", "workspace", "", false)

	parentTypeReferences := []*TypeReference{workspaceTypeReference}

	parentRelationExpression := NewSelfRelationExpression(parentTypeReferences, CardinalityAtMostOne)

	parentRelation, err := NewRelation("parent", workspace, VisibilityPublic, parentRelationExpression, nil)
	assert.NoError(t, err)

	ownerPrincipalTypeReference := NewTypeReference("", "principal", "", false)
	ownerTypeReferences := []*TypeReference{ownerPrincipalTypeReference}
	ownerPrincipalRelationExpression := NewSelfRelationExpression(ownerTypeReferences, CardinalityAny)

	ownerTypo := "ownr"
	ownerParentRelationExpression := NewReferenceRelationExpression("parent", &ownerTypo)

	ownerRelationExpression := NewSetRelationExpression("union", ownerPrincipalRelationExpression, ownerParentRelationExpression)
	ownerRelation, err := NewRelation("owner", workspace, VisibilityPublic, ownerRelationExpression, nil)
	assert.NoError(t, err)

	workspace.AddRelation(parentRelation)
	workspace.AddRelation(ownerRelation)
	namespace.AddType(principal)
	namespace.AddType(workspace)
	schema.AddNamespace(namespace)

	_, err = schema.ToZanzibar()
	assert.ErrorIs(t, err, ErrSymbolNotFound)
}

func TestAssertReferenceRelationExpressionToZanzibarSucceedsIfSubRelationOkay(t *testing.T) {
	schema := NewSchema()
	namespace := NewNamespace("test_namespace", []string{})
	principal := NewType("principal", namespace, VisibilityPublic)
	workspace := NewType("workspace", namespace, VisibilityPublic)

	workspaceTypeReference := NewTypeReference("", "workspace", "", false)

	parentTypeReferences := []*TypeReference{workspaceTypeReference}

	parentRelationExpression := NewSelfRelationExpression(parentTypeReferences, CardinalityAtMostOne)

	parentRelation, err := NewRelation("parent", workspace, VisibilityPublic, parentRelationExpression, nil)
	assert.NoError(t, err)

	ownerPrincipalTypeReference := NewTypeReference("", "principal", "", false)
	ownerTypeReferences := []*TypeReference{ownerPrincipalTypeReference}
	ownerPrincipalRelationExpression := NewSelfRelationExpression(ownerTypeReferences, CardinalityAny)

	owner := "owner"
	ownerParentRelationExpression := NewReferenceRelationExpression("parent", &owner)

	ownerRelationExpression := NewSetRelationExpression("union", ownerPrincipalRelationExpression, ownerParentRelationExpression)
	ownerRelation, err := NewRelation("owner", workspace, VisibilityPublic, ownerRelationExpression, nil)
	assert.NoError(t, err)

	workspace.AddRelation(parentRelation)
	workspace.AddRelation(ownerRelation)
	namespace.AddType(principal)
	namespace.AddType(workspace)
	schema.AddNamespace(namespace)

	_, err = schema.ToZanzibar()
	assert.NoError(t, err)
}
