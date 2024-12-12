public type principal {}
public type workspace {
    relation parent: [AtMostOne workspace]
    relation owner: [Any principal] or parent.ownr
}

declaration 
    typeExpr (public type workspace)
        relation*
            relation parent
                relationBody [AtMostOne workspace] (SelfRelation)
            relation owner
                OR
                    relationBody 
                        typeReference -> [Any principal] (SelfRelation)
                        relationBody 
                            SubRelation -> parent.ownr (SubRelation)

relationBody(parent) -> SelfRelation(Cardinality=AtMostOne, typeName=workspace)
relationBody(owner) -> SelfRelation(Cardinality=Any, typeName=principal) OR SubRelation(relationName=parent, subRelationName=ownr)

file
    delcarations
        workspace
            relations
                owner
                    OR
                        SubRelation(relationName=parent, subRelationName=ownr)
                        SelfRelation(Cardinality=Any, typeName=principal)
