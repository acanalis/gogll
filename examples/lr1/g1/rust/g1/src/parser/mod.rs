//! Generated by GoGLL. Do not edit.

mod action_table;
mod goto_table;
mod productions_table;

use action_table::{ACTION_TABLE, Action::{Accept, Reduce, Shift}};
use goto_table::GOTO_TABLE;
use productions_table::PROD_TABLE;
use crate::ast;
use crate::ast::{Node::{T}};
use crate::lexer;
use crate::token;

use std::rc::Rc;

/*** Stack ***/
struct Stack {
	state: Vec<usize>,
	attrib: Vec<Option<Box<ast::Node>>>,
}

impl Stack {
	fn new() -> Box<Stack> {
		Box::new(Stack{ 
			state: 	Vec::with_capacity(100),
			attrib: Vec::with_capacity(100),
		})
	}

	fn peek(&self, pos: usize) -> usize {
		return self.state[pos]
	}

	fn pop_n(&mut self, items: usize) -> Vec<Option<Box<ast::Node>>> {
		let from = self.state.len() - items;

		self.state.truncate(from);
		self.attrib.split_off(from)
	}

	fn push(&mut self, s: usize, a: Option<Box<ast::Node>>) {
		self.state.push(s);
		self.attrib.push(a);
	}	

	// TODO: unused, delete
	// fn reset(&mut self) {
	// 	self.state.clear();	
	// 	self.attrib.clear();
	// }
	
	fn top(&self) -> usize {
		self.state[self.state.len()-1]
	}		

	fn top_index(&self) -> usize {
		self.state.len() - 1
	}	

}				


// TODO: implement or delete
// func (S *stack) String() string {
// 	w := new(bytes.Buffer)
// 	fmt.Fprintf(w, "stack:\n")
// 	for i, st := range S.state {
// 		fmt.Fprintf(w, "\t%d:%d , ", i, st)
// 		if S.attrib[i] == nil {
// 			fmt.Fprintf(w, "nil")
// 		} else {
// 			fmt.Fprintf(w, "%v", S.attrib[i])
// 		}
// 		w.WriteString("\n")
// 	}
// 	return w.String()
// }


/*** Parser ***/

pub struct Parser {
	stack:     Box<Stack>,
	next_token: Rc<token::Token>, 

	lex: Rc<lexer::Lexer>,

	/// input position in token stream
	i: usize,
}

impl Parser {
	#[allow(dead_code)]
	pub fn new(lex: Rc<lexer::Lexer>) -> Box<Parser> {
		let mut p = Box::new(Parser{
				stack:  Stack::new(),
				next_token: lex.tokens[0].clone(),
				lex:    lex,
				i:      1,
		});
		p.stack.push(0, None);
		p
	}

	pub fn parse(&mut self) -> Result<Option<Box<ast::Node>>, String> {
		let mut acc = false;
		let mut res: Option<Box<ast::Node>> = None;

		while !acc {
			match ACTION_TABLE[self.stack.top()].actions.get(&self.next_token.typ) {
				None => return Err(self.error()),
				Some(act) => {
					match act {
						Accept => {
                            res = self.stack.pop_n(1).remove(0);
							acc = true;
						},
						Shift(state) => {
							self.stack.push(*state, Some(Box::new(T(self.next_token.clone()))));
							self.next();
						},
						Reduce(production) => {
							let prod = &PROD_TABLE[*production];
							match (prod.reduce_func)(self.stack.pop_n(prod.num_symbols)) {
								Err(e) => return Err(e),
								Ok(nd) => match GOTO_TABLE[self.stack.top()][prod.nt_type] {
									None => panic!("State {} NT {}", self.stack.top(), prod.nt_type),
									Some(s) => self.stack.push(s, nd)
								}
							}
						},
					}
				},
			}
	
			// println!("S{} {} {}", self.stack.top(), self.next_token, action)
	
	
		}
		Ok(res)
	}
	
	fn next(&mut self) {
		if self.next_token.typ != token::Type::EOF {
			self.next_token = self.lex.tokens[self.i].clone();
			self.i += 1;
		}
	}
	
	fn error(&self) -> String {
		let (ln, col) = self.next_token.get_line_column();
		let mut errs = format!("Error @ line {} col {}, token {}", ln, col, self.next_token.clone());

		errs

		// fmt.Fprintf(w, "Error in S%d: %s @ line %d col %d",
		// 	P.stack.top(), P.next_token.LiteralString(), ln, col)
		// if err != nil {
		// 	w.WriteString(err.Error())
		// } else {
		// 	w.WriteString(", expected one of: ")
		// 	actRow := ACTION_TABLE[P.stack.top()]
		// 	for tok, act := range actRow.actions {
		// 		if act != nil {
		// 			fmt.Fprintf(w, "%s ", tok.ID())
		// 		}
		// 	}
		// }
		// return errors.New(w.String())
	}
	
}	
