package ast

import (
	"github.com/ddddddO/monkey/token"
)

type Node interface {
	TokenLiteral() string
}

// 文(値を生成しないもの)
type Statement interface {
	Node
	statementNode()
}

// 式(値を生成するもの)
type Expression interface {
	Node
	expressionNode()
}

// AST(抽象構文木)のルートノード、、
type Program struct {
	Statements []Statement // ASTノードのスライス、、
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// let <identifier> = <expression>;
type LetStatement struct {
	Token token.Token // token.LETトークン
	Name  *Identifier // let文の左辺
	Value Expression  // let文の右辺
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // token.IDENTトークン
	Value string
}

func (i *Identifier) expressionNode()      {} // Monkeyでは識別子は値を生成する、から。p33
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
