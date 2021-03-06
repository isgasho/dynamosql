parser.row{
  Query: "SELECT * FROM movies WHERE UserId = :UserId DESC",
  AST: &parser.AST{
    Select: &parser.Select{
      Projection: &parser.ProjectionExpression{
        All: true,
      },
      From: "movies",
      Where: &parser.AndExpression{
        And: []*parser.Condition{
          {
            Operand: &parser.ConditionOperand{
              Operand: &parser.DocumentPath{
                Fragment: []*parser.PathFragment{
                  {
                    Symbol: "UserId",
                  },
                },
              },
              ConditionRHS: &parser.ConditionRHS{
                Compare: &parser.Compare{
                  Operator: "=",
                  Operand: &parser.Operand{
                    Value: &parser.Value{
                      PlaceHolder: &":UserId",
                    },
                  },
                },
              },
            },
          },
        },
      },
      Descending: &parser.ScanDescending(true),
    },
  },
}