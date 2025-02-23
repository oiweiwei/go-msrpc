%{
package midl

// parse.go contains the go-yacc definitions for the
// (M)IDL grammar parser.

import (
    "fmt"
    "math/big"

    "github.com/oiweiwei/go-msrpc/midl/uuid"
)

var (
    _ = fmt.Sprintf("")
)

%}

%union {
    File                  File
    Interface             Interface
    InterfaceBody         InterfaceBody
    Version               *Version
    Export                []*Export
    Type                  *Type
    Kind                  Kind
    Const                 *Const
    Operation             *Operation
    Operations            []*Operation
    Param                 *Param
    Params                []*Param
    Fields                []*Field
    Field                 *Field
    Declarators           []*pDeclarator
    Declarator            *pDeclarator
    ArrayBound            ArrayBound
    Int64                 int64
    UnionSwitch           *UnionSwitch
    UnionCases            []*UnionCase
    UnionCase             *UnionCase
    Expr                  Expr
    Exprs                 Exprs
    Ident                 string
    Char                  rune
    String                string
    Int                   *big.Int
    UUID                  *uuid.UUID
    Range                 *Range
    Strings               []string
    Token                 int
    Typedef               *pTypedef
    Attr                  pAttr
    AttrType              pAttrType
    TagID                 pTagID
    TagIDs                []pTagID
    ComClass              ComClass
    Library               Library
    ComInterfaces         []*ComInterface
    ComInterface          *ComInterface
    DispatchInterface     DispatchInterface
    DispatchInterfaceBody DispatchInterfaceBody
}

%type <File> external file
%type <Library> library library_header library_body
%type <Interface> interface interface_header
%type <ComClass> coclass coclass_header
%type <ComInterfaces> coclass_member_list
%type <ComInterface> coclass_member
%type <DispatchInterface> dispinterface dispinterface_header
%type <DispatchInterfaceBody> dispinterface_body dispinterface_component
%type <Operations> method_list
%type <Ident> interface_base
%type <InterfaceBody> interface_body
%type <InterfaceBody> interface_component interface_component1
%type <Version> version
%type <Strings> port_spec1
%type <String> port_spec
%type <Strings> excep_name1
%type <String> excep_name
%type <Strings> import_list
%type <Strings> import import1
%type <String> import_name
%type <Export> export
%type <Operation> op_declarator

%type <Const> const_declarator
%type <Export> pragma_declarator

%type <Typedef> type_declarator
%type <Type> struct_type
%type <Type> union_type
%type <Type> pipe_type
%type <Type> enumeration_type
%type <Type> type_spec coclass_type_spec
%type <Type> tagged_struct
%type <Type> tagged_struct_declarator
%type <Type> constructed_type_spec
%type <Type> tagged_declarator
// XXX: __midl__
%type <Type> tagged_enumeration
%type <Type> tagged_enumeration_declarator
// XXX: __midl_end__
%type <Type> switch_type_spec
%type <Type> union_type_switch_attr
%type <Type> tagged_union_declarator
%type <Type> tagged_union
%type <Type> xmit_type
%type <Type> simple_type_spec type_ident
%type <Type> predefined_type_spec

%type <Range> range_declarator

%type <Attr> attributes_list attributes0 attributes1 coclass_attributes0 coclass_attribute_list
%type <AttrType> attribute coclass_attribute
%type <AttrType> union_instance_switch_attr
%type <AttrType> directional_attribute
%type <AttrType> range_attribute

%type <Exprs> const_exp1

%type <Expr> const_exp
%type <Expr> integer_const_exp
%type <Expr> conditional_exp
%type <Expr> logical_or_exp
%type <Expr> logical_and_exp
%type <Expr> inclusive_or_exp
%type <Expr> exclusive_or_exp
%type <Expr> and_exp
%type <Expr> equality_exp
%type <Expr> relational_exp
%type <Expr> shift_exp
%type <Expr> additive_exp
%type <Expr> multiplicative_exp
%type <Expr> unary_exp
%type <Expr> primary_exp


%type <Declarator> declarator
%type <Declarator> direct_declarator
%type <Declarators> declarators

%type <Param> param_declarator
%type <Params> param_declarators1

%type <Fields> member_list property_list
%type <Fields> member
%type <Fields> field_declarator

%type <String> tag ident_or_keyword

// XXX: __rule_reduced__
// %type <String> rr_ident_const_exp
// XXX: __rule_reduced_end__
%type <Int64> pointer_opt
%type <Int64> pointer_opt0

%type <TagID> identifier_tag

%type <Declarator> array_declarator
%type <Declarator> function_declarator

%type <ArrayBound> array_bounds_declarator
%type <ArrayBound> array_bound_pair
%type <Int> array_bound

%type <Expr> field_attr_var
%type <Expr> attr_var
%type <Exprs> attr_var_list
%type <TagIDs> enumeration1

%type <UnionSwitch> union_switch
%type <UnionCases> union_body
%type <UnionCases> union_body_n_e
%type <UnionCases> union_body_c
%type <String> union_name
%type <UnionCase> union_case
%type <UnionCase> union_case_n_e
%type <Exprs> union_case_label_n_e
%type <Exprs> union_case_label1
%type <Exprs> union_case_label
%type <Fields> union_arm
%type <Fields> default_case
%type <Fields> default_case_n_e

%type <Kind> const_type_spec
%type <Kind> base_type_spec
%type <Kind> floating_pt_type
%type <Kind> integer_type
%type <Kind> char_type
%type <Kind> wchar_type
%type <Kind> boolean_type
%type <Kind> byte_type
%type <Kind> void_type
%type <Kind> handle_type
%type <Kind> primitive_integer_type
%type <Kind> signed_integer
%type <Kind> ms_integer
%type <Kind> unsigned_integer

%type <Token> usage_attribute
%type <Token> format_attribute
%type <Token> unsigned
%type <Token> sign
%type <Token> integer_size
%type <Token> international_character_type
%type <Token> ptr_attr

%token <Token> FLOAT
%token <Token> DOUBLE
%token <Token> HYPER
%token <Token> UNSIGNED
%token <Token> SIGNED
%token <Token> LONG
%token <Token> SHORT
%token <Token> SMALL
%token <Token> INT
%token <Token> CHAR
%token <Token> BOOLEAN
%token <Token> BYTE
%token <Token> VOID
%token <Token> HANDLE_T
%token <Token> ERROR_STATUS_T
%token <Token> ISO_LATIN_1
%token <Token> ISO_MULTILINGUAL
%token <Token> ISO_UCS
%token <Token> STRUCT
%token <Token> FIRST_IS
%token <Token> LAST_IS
%token <Token> LENGTH_IS
%token <Token> MAX_IS
%token <Token> MIN_IS
%token <Token> SIZE_IS
%token <Token> SWITCH_IS
%token <Token> USAGE_STRING
%token <Token> USAGE_CONTEXT_HANDLE
%token <Token> FORMAT
%token <Token> FORMAT_NULL_TERMINATED
%token <Token> FORMAT_MULTI_SIZE
%token <Token> FORMAT_UTF8
%token <Token> FORMAT_RUNE
%token <Token> FORMAT_HEX
%token <Token> IGNORE
%token <Token> POINTER
%token <Token> POINTER_REF
%token <Token> POINTER_UNIQUE
%token <Token> POINTER_PTR
%token <Token> CONST
%token <Token> NULL
%token <Token> TRUE
%token <Token> FALSE
%token <Token> IN
%token <Token> OUT
%token <Token> ENUM
%token <Token> PIPE
%token <Token> UNION
%token <Token> SWITCH
%token <Token> CASE
%token <Token> DEFAULT
%token <Token> SWITCH_TYPE
%token <Token> TRANSMIT_AS
%token <Token> HANDLE
%token <Token> IMPORT
%token <Token> TYPEDEF
%token <UUID> UUID
%token <Token> INTERFACE
%token <Token> IDEMPOTENT
%token <Token> BROADCAST
%token <Token> MAYBE
%token <Token> REFLECT_DELETIONS
%token <Token> VERSION
%token <Token> ENDPOINT
%token <Token> EXCEPTIONS
%token <Token> LOCAL
%token <Token> POINTER_DEFAULT
%token <Token> RETVAL
%token <Token> IID_IS

// XXX: __midl__
%token <Token> WCHAR_T
%token <Token> INT3264
%token <Token> INT8
%token <Token> INT16
%token <Token> INT32
%token <Token> INT64
%token <Token> RANGE
%token <Token> MS_UNION
%token <Token> OBJECT
%token <Token> V1_ENUM
%token <Token> STRICT_CONTEXT_HANDLE
%token <Token> TYPE_STRICT_CONTEXT_HANDLE
%token <Token> DISABLE_CONSISTENCY_CHECK
%token <Token> SIZEOF
%token <Token> PRAGMA_DEFINE
%token <Token> PRAGMA_CPP_QUOTE
%token <Token> CALLBACK
%token <Token> HELP_STRING
%token <Token> DUAL
%token <Token> PROPGET
%token <Token> PROPPUT
%token <Token> PROPPUTREF
%token <Token> ID
%token <Token> HIDDEN
%token <Token> NONEXTENSIBLE
%token <Token> RESTRICTED
%token <Token> DEFAULT_VALUE
%token <Token> ODL
%token <Token> OLEAUTOMATION
%token <Token> OPTIONAL
%token <Token> APPOBJECT
%token <Token> SAFEARRAY
%token <Token> PAD
%token <Token> GOEXT_LAYOUT
%token <Token> GOEXT_DEFAULT_NULL
%token <Token> GOEXT_NO_SIZE_LIMIT
%token <Token> CALL_AS
%token <Token> ANNOTATION
%token <Token> WIRE_MARSHAL
%token <Token> PUBLIC
%token <Token> SOURCE
%token <Token> DISPINTERFACE
%token <Token> METHODS
%token <Token> PROPERTIES
%token <Token> COCLASS
%token <Token> LIBRARY

%token <Token> ACS_BYTE_COUNT
// XXX: __midl_end__

%token <Ident> STRING

%token <Char> CHARACTER_LITERAL
%token <String> STRING_LITERAL
%token <Int> INT_LITERAL

%token <Ident> IDENT

%token TERNARY

%left RNG
%left LOGICAL_OR
%left LOGICAL_AND
%left LE GE LT GT EQ NE
%left OR XOR AND
%left LSH RSH
%left '+' '-'
%left '*' '/' '%'
%left UPLUS
%left UNEG
%left UNOT
%left CAST
%left UMINUS
%left UMUL

%start main

%%

main                    : /* empty */
                        | file
                            {
                                f := $1
                                setResult(RPClex, &f)
                            }
                        ;

file                    : external
                            {
                                $$ = $1
                                if len($$.Interfaces) > 0 {
                                    if $$.Export == nil {
                                        $$.Export = make(map[string]*Export)
                                    }
                                    for i := range $$.Interfaces {
                                        iff := $$.Interfaces[i]
                                        if iff.Attrs.Object || iff.BaseName != "" {
                                            $$.Export[iff.Name] = &Export{Type: InterfaceToExport(iff), Position: Hash(iff.Name)}
                                        }
                                    }
                                }
                            }
                        | file external
                            {
                                if len($2.Imports) > 0 {
                                    $$.Imports = append($$.Imports, $2.Imports...)
                                }
                                if len($2.Export) > 0 {
                                    if $$.Export == nil {
                                        $$.Export = make(map[string]*Export)
                                    }
                                    for k, v := range $2.Export {
                                        $$.Export[k] = v
                                    }
                                }
                                if len($2.Interfaces) > 0 {
                                    $$.Interfaces = append($$.Interfaces, $2.Interfaces...)
                                    if $$.Export == nil {
                                        $$.Export = make(map[string]*Export)
                                    }
                                    for i := range $2.Interfaces {
                                        iff := $2.Interfaces[i]
                                        if iff.Attrs.Object || iff.BaseName != "" {
                                            $$.Export[iff.Name] = &Export{Type: InterfaceToExport(iff), Position: Hash(iff.Name)}
                                        }
                                    }
                                }
                                if len($2.ComClasses) > 0 {
                                    $$.ComClasses = append($$.ComClasses, $2.ComClasses...)
                                }
                                if len($2.DispatchInterfaces) > 0 {
                                    $$.DispatchInterfaces = append($$.DispatchInterfaces, $2.DispatchInterfaces...)
                                }
                                if len($2.Libraries) > 0 {
                                    $$.Libraries = append($$.Libraries, $2.Libraries...)
                                }
                            }
                        ;

external                : import1
                            {
                                $$ = File{Imports: $1}
                            }
                        | export
                            {
                                $$ = File{Export: make(map[string]*Export)}
                                if len($1) > 0 {
                                    for _, e := range $1 {
                                        $$.Export[e.Name] = e
                                        exportSyms(RPClex, e.Name, e)
                                    }
                                }
                            }
                        | interface semicolon
                            {
                                iff := $1
                                if iff.Attrs != nil {
                                    $$ = File{Interfaces: []*Interface{&iff}}
                                }
                                if iff.BaseName != "" {
                                    typ, ok := lookupType(RPClex, iff.BaseName)
                                    if !ok {
                                        RPClex.Error("cannot find interface " + iff.BaseName)
                                        return 0
                                    }

                                    if typ.Kind == TypeAttribute {
                                        typ = typ.Elem
                                    }

                                    if iff.Base = typ.Interface; iff.Base != nil {
                                        inc := 0
                                        if l := len(iff.Base.Body.Operations); l > 0 {
                                            inc = iff.Base.Body.Operations[l-1].OpNum + 1
                                        }
                                        for i := range iff.Body.Operations {
                                            iff.Body.Operations[i].OpNum += inc
                                        }
                                    }
                                }
                                exportSyms(RPClex, iff.Name, &Export{Type: InterfaceToExport(&iff)})
                            }
                        | coclass semicolon
                            {
                                cc := $1
                                $$ = File{ComClasses: []*ComClass{&cc}}
                            }
                        | dispinterface semicolon
                            {
                                di := $1
                                $$ = File{DispatchInterfaces: []*DispatchInterface{&di}}
                            }
                        | library semicolon
                            {
                                lib := $1
                                $$ = File{Libraries: []*Library{&lib}}
                            }
                        ;

semicolon               : /* empty */
                        | ';'
                        ;

library                 : library_header '{' library_body '}'
                            {
                                $$ = Library{Attrs: $1.Attrs, Name: $1.Name, Body: $3.Body}
                            }
                        ;

library_header          : attributes0 LIBRARY IDENT
                            {
                                $$ = Library{Attrs: $1.Library(), Name: $3}
                            }
                        ;

library_body            : coclass semicolon
                            {
                                cc := $1
                                $$ = Library{Body: LibraryBody{ComClasses: []*ComClass{&cc}}}
                            }
                        | library_body coclass semicolon
                            {
                                cc := $2
                                $$.Body.ComClasses = append($$.Body.ComClasses, &cc)
                            }
                        ;

coclass                 : coclass_header '{' coclass_member_list '}'
                            {
                                $$ = ComClass{Name: $1.Name, Attrs: $1.Attrs, Interfaces: $3}
                            }
                        ;

coclass_header          : attributes0 COCLASS IDENT
                            {
                               $$ = ComClass{Attrs: $1.ComClass(), Name: $3}
                            }
                        ;

coclass_member_list     : coclass_member
                            {
                                $$ = []*ComInterface{$1}
                            }
                        | coclass_member_list coclass_member
                            {
                                $$ = append($$, $2)
                            }
                        ;

coclass_member          : coclass_attributes0 coclass_type_spec IDENT ';'
                            {
                                $$ = &ComInterface{Name: $3, Type: $2, Attrs: $1.ComInterface()}
                            }
                        ;

coclass_type_spec       : INTERFACE
                            {
                                $$ = &Type{Kind: TypeInterface}
                            }
                        | DISPINTERFACE
                            {
                                $$ = &Type{Kind: TypeDispInterface}
                            }
                        ;

coclass_attributes0     : /* empty */
                            {
                                $$ = pAttr{}
                            }
                        | '[' coclass_attribute_list comma ']'
                            {
                                $$ = $2
                            }
                        ;

coclass_attribute_list  : coclass_attribute
                            {
                                switch $1.Type {
                                case DEFAULT:
                                    $$.Default = true
                                case SOURCE:
                                    $$.Source = true
                                }
                            }
                        | coclass_attribute_list ',' coclass_attribute
                            {
                                switch $3.Type {
                                case DEFAULT:
                                    $$.Default = true
                                case SOURCE:
                                    $$.Source = true
                                }
                            }
                        ;


coclass_attribute       : DEFAULT
                            {
                                $$ = pAttrType{DEFAULT, pAttr{}}
                            }
                        | SOURCE
                            {
                                $$ = pAttrType{SOURCE, pAttr{}}
                            }
                        ;

dispinterface           : dispinterface_header '{' dispinterface_body '}'
                            {
                                $$ = DispatchInterface{Name: $1.Name, Attrs: $1.Attrs, Body: $3}
                                for i := range $3.Methods {
                                    $3.Methods[i].OpNum = i
                                }
                            }
                        ;

dispinterface_header    : attributes0 DISPINTERFACE IDENT
                            {
                                $$ = DispatchInterface{Attrs: $1.DispatchInterface(), Name: $3}
                            }
                        ;

dispinterface_body      : dispinterface_component
                            {
                                $$ = $1
                            }
                        | dispinterface_body dispinterface_component
                            {
                                if len($2.Properties) > 0 {
                                    $$.Properties = append($$.Properties, $2.Properties...)
                                }
                                if len($2.Methods) > 0 {
                                    $$.Methods = append($$.Methods, $2.Methods...)
                                }
                            }
                        ;

dispinterface_component : PROPERTIES ':' property_list
                            {
                                $$ = DispatchInterfaceBody{Properties: $3}
                            }
                        | METHODS ':' method_list
                            {
                                $$ = DispatchInterfaceBody{Methods: $3}
                            }
                        ;

property_list           : /* empty */
                            {
                                $$ = nil
                            }
                        | member_list
                            {
                                $$ = $1
                            }
                        ;

method_list             : op_declarator ';'
                            {
                                $$ = []*Operation{$1}
                            }
                        | method_list op_declarator ';'
                            {
                                $$ = append($$, $2)
                            }
                        ;

interface               : interface_header '{' interface_body '}'
                            {
                                $$ = Interface{Name: $1.Name, Attrs: $1.Attrs, Body: $3, BaseName: $1.BaseName}
                            }
                        | interface_header
                            {
                                $$ = Interface{Name: $1.Name, Attrs: $1.Attrs, BaseName: $1.BaseName}
                            }
                        | interface_header '{' '}'
                            {
                                $$ = Interface{Name: $1.Name, Attrs: $1.Attrs, BaseName: $1.BaseName}
                            }
                        ;

interface_header        : attributes0 INTERFACE IDENT interface_base
                            {
                                $$ = Interface{Attrs: $1.Interface(), Name: $3, BaseName: $4}
                            }
                        ;

interface_base          : /* empty */
                            {
                                $$ = ""
                            }
                        | ':' IDENT
                            {
                                $$ = $2
                            }
                        ;

attributes0             : /* empty */
                            {
                                $$ = pAttr{}
                            }
                        | attributes1
                            {
                                $$ = $1
                            }
                        ;

attributes1             : '[' attributes_list comma ']'
                            {
                                $$ = $2
                            }
                        | attributes1 '[' attributes_list comma ']'
                            {
                                $$ = $$.Merge($3)
                            }
                        ;

attributes_list         : attribute
                            {
                                $$ = $$.Set($1)
                            }
                        | attributes_list ',' attribute
                            {
                                $$ = $$.Set($3)
                            }
                        ;

// XXX: put all attrbiutes to the single statement. validate via code.
attribute               : FIRST_IS '(' attr_var_list ')'
                            {
                                $$ = pAttrType{FIRST_IS, pAttr{FirstIs: $3}}
                            }
                        | LAST_IS '(' attr_var_list ')'
                            {
                                $$ = pAttrType{LAST_IS, pAttr{LastIs: $3}}
                            }
                        | LENGTH_IS '(' attr_var_list ')'
                            {
                                $$ = pAttrType{LENGTH_IS, pAttr{LengthIs: $3}}
                            }
                        | MIN_IS '(' attr_var_list ')'
                            {
                                $$ = pAttrType{MIN_IS, pAttr{MinIs: $3}}
                            }
                        | MAX_IS '(' attr_var_list ')'
                            {
                                $$ = pAttrType{MAX_IS, pAttr{MaxIs: $3}}
                            }
                        | SIZE_IS '(' attr_var_list ')'
                            {
                                $$ = pAttrType{SIZE_IS, pAttr{SizeIs: $3}}
                            }
                        | ptr_attr
                            {
                                $$ = pAttrType{$1, pAttr{}}
                            }
                        | GOEXT_DEFAULT_NULL '(' attr_var ')'
                            {
                                if $3.Empty() {
                                    $$ = pAttrType{GOEXT_DEFAULT_NULL, pAttr{DefaultNull: []Expr{}}}
                                } else {
                                    $$ = pAttrType{GOEXT_DEFAULT_NULL, pAttr{DefaultNull: []Expr{$3}}}
                                }
                            }
                        | usage_attribute
                            {
                                $$ = pAttrType{$1, pAttr{}}
                            }
                        | FORMAT '(' format_attribute ')'
                            {
                                $$ = pAttrType{$3, pAttr{}}
                            }
                        | union_instance_switch_attr
                            {
                                $$ = $1
                            }
                        | IGNORE
                            {
                                $$ = pAttrType{IGNORE, pAttr{}}
                            }
                        | directional_attribute
                            {
                                $$ = $1
                            }
                        |  TRANSMIT_AS '(' xmit_type ')'
                            {
                                $$ = pAttrType{TRANSMIT_AS, pAttr{TransmitAs: $3}}
                            }
                        | union_type_switch_attr
                            {
                                $$ = pAttrType{SWITCH_TYPE, pAttr{SwitchType: $1}}
                            }
                        | HANDLE
                            {
                                $$ = pAttrType{HANDLE, pAttr{}}
                            }
                        | IDEMPOTENT
                            {
                                $$ = pAttrType{IDEMPOTENT, pAttr{}}
                            }
                        | BROADCAST
                            {
                                $$ = pAttrType{BROADCAST, pAttr{}}
                            }
                        | MAYBE
                            {
                                $$ = pAttrType{MAYBE, pAttr{}}
                            }
                        | REFLECT_DELETIONS
                            {
                                $$ = pAttrType{REFLECT_DELETIONS, pAttr{}}
                            }
                        | UUID
                            {
                                $$ = pAttrType{UUID, pAttr{UUID: $1}}
                            }
                        | VERSION '(' version ')'
                            {
                                $$ = pAttrType{VERSION, pAttr{Version: $3}}
                            }
                        | ENDPOINT '(' port_spec1 ')'
                            {
                                $$ = pAttrType{ENDPOINT, pAttr{Endpoints: $3}}
                            }
                        | EXCEPTIONS '(' excep_name1 ')'
                            {
                                $$ = pAttrType{EXCEPTIONS, pAttr{Exceptions: $3}}
                            }
                        | LOCAL
                            {
                                $$ = pAttrType{LOCAL, pAttr{}}
                            }
                        | POINTER_DEFAULT '(' ptr_attr ')'
                            {
                                switch $3 {
                                case POINTER_PTR:
                                    $$ = pAttrType{POINTER_DEFAULT, pAttr{PointerDefault: PointerTypePtr}}
                                case POINTER_REF:
                                    $$ = pAttrType{POINTER_DEFAULT, pAttr{PointerDefault: PointerTypeRef}}
                                case POINTER_UNIQUE:
                                    $$ = pAttrType{POINTER_DEFAULT, pAttr{PointerDefault: PointerTypeUnique}}
                                }
                            }
                        | V1_ENUM
                            {
                                $$ = pAttrType{V1_ENUM, pAttr{}}
                            }
                        | MS_UNION
                            {
                                $$ = pAttrType{MS_UNION, pAttr{}}
                            }
                        | range_attribute
                            {
                                $$ = $1
                            }
                        | DISABLE_CONSISTENCY_CHECK
                            {
                                $$ = pAttrType{DISABLE_CONSISTENCY_CHECK, pAttr{}}
                            }
                        | OBJECT
                            {
                                $$ = pAttrType{OBJECT, pAttr{}}
                            }
                        | CALLBACK
                            {
                                $$ = pAttrType{CALLBACK, pAttr{}}
                            }
                        | RETVAL
                            {
                                $$ = pAttrType{RETVAL, pAttr{}}
                            }
                        | IID_IS '(' attr_var ')'
                            {
                                $$ = pAttrType{IID_IS, pAttr{IIDIs: $3}}
                            }
                        | HELP_STRING '(' STRING_LITERAL ')'
                            {
                                $$ = pAttrType{HELP_STRING, pAttr{HelpString: $3}}
                            }
                        | DUAL
                            {
                                $$ = pAttrType{DUAL, pAttr{}}
                            }
                        | PROPGET
                            {
                                $$ = pAttrType{PROPGET, pAttr{}}
                            }
                        | PROPPUT
                            {
                                $$ = pAttrType{PROPPUT, pAttr{}}
                            }
                        | PROPPUTREF
                            {
                                $$ = pAttrType{PROPPUTREF, pAttr{}}
                            }
                        | ID '(' integer_const_exp ')'
                            {
                                val, ok := $3.Eval(RPClex.(ExprStore))
                                if !ok {
                                    RPClex.Error("cannot eval ID expression")
                                    return 0
                                }
                                if _, ok = val.BigInt(); !ok {
                                    RPClex.Error("cannot eval ID expression")
                                    return 0
                                }
                                $$ = pAttrType{ID, pAttr{ID: val}}
                            }
                        | HIDDEN
                            {
                                $$ = pAttrType{HIDDEN, pAttr{}}
                            }
                        | NONEXTENSIBLE
                            {
                                $$ = pAttrType{NONEXTENSIBLE, pAttr{}}
                            }
                        | RESTRICTED
                            {
                                $$ = pAttrType{RESTRICTED, pAttr{}}
                            }
                        | DEFAULT_VALUE '(' const_exp ')'
                            {
                                exp, ok := $3.Eval(RPClex.(ExprStore))
                                if !ok {
                                    RPClex.Error("cannot evaluate default value")
                                    return 0
                                }
                                $$ = pAttrType{DEFAULT_VALUE, pAttr{DefaultValue: exp}}
                            }
                        | ODL
                            {
                                $$ = pAttrType{ODL, pAttr{}}
                            }
                        | OLEAUTOMATION
                            {
                                $$ = pAttrType{OLEAUTOMATION, pAttr{}}
                            }
                        | OPTIONAL
                            {
                                $$ = pAttrType{OPTIONAL, pAttr{}}
                            }
                        | APPOBJECT
                            {
                                $$ = pAttrType{APPOBJECT, pAttr{}}
                            }
                        | ANNOTATION '(' STRING_LITERAL ')'
                            {
                                $$ = pAttrType{ANNOTATION, pAttr{Annotation: $3}}
                            }
                        | CALL_AS '(' IDENT ')'
                            {
                                $$ = pAttrType{CALL_AS, pAttr{CallAs: $3}}
                            }
                        | WIRE_MARSHAL '(' IDENT ')'
                            {
                                $$ = pAttrType{WIRE_MARSHAL, pAttr{WireMarshal: $3}}
                            }
                        | PUBLIC
                            {
                                $$ = pAttrType{PUBLIC, pAttr{}}
                            }
                        | SAFEARRAY '(' type_spec ')'
                            {
                                $$ = pAttrType{SAFEARRAY, pAttr{Safearray: $3}}
                            }
                        | PAD '(' INT_LITERAL ')'
                            {
                                $$ = pAttrType{PAD, pAttr{Pad: $3.Uint64()}}
                            }
                        | GOEXT_LAYOUT '(' field_declarator ')'
                            {
                                for i := range $3 {
                                    $3[i].Attrs.IsLayout = true
                                }
                                $$ = pAttrType{GOEXT_LAYOUT, pAttr{Layout: $3}}
                            }
                        | SIZE_IS '(' '*' ')'
                            {
                                $$ = pAttrType{GOEXT_NO_SIZE_LIMIT, pAttr{NoSizeLimit: true}}
                            }
                        ;

version                 : INT_LITERAL
                            {
                                $$ = &Version{Major: uint16($1.Uint64())}
                            }
                        | INT_LITERAL '.' INT_LITERAL
                            {
                                $$ = &Version{Major: uint16($1.Uint64()), Minor: uint16($3.Uint64())}
                            }
                        ;

port_spec1              : port_spec
                            {
                                $$ = []string{$1}
                            }
                        | port_spec1 ',' port_spec
                            {
                                $$ = append($$, $3)
                            }
                        ;

excep_name1             : excep_name
                            {
                                $$ = []string{$1}
                            }
                        | excep_name1 ',' excep_name
                            {
                                $$ = append($$, $3)
                            }
                        ;

port_spec               : STRING_LITERAL
                            {
                                $$ = $1
                            }
                        ;

excep_name              : IDENT
                            {
                                $$ = $1
                            }
                        ;

interface_body          : import interface_component1
                            {
                                $$ = InterfaceBody{Imports: $1, Export: $2.Export, Operations: $2.Operations}
                                for i, o := range $$.Operations {
                                    o.OpNum = i
                                }
                            }
                        ;

interface_component1    : interface_component
                            {
                                $$ = $1
                                export := make([]*Export, len($1.Export))
                                for _, e := range $1.Export {
                                    export[e.Position] = e
                                }
                                for _, e := range export {
                                    exportSyms(RPClex, e.Name, e)
                                }
                            }
                        | interface_component1 interface_component
                            {
                                if len($2.Export) > 0 {
                                    if $$.Export == nil {
                                        $$.Export = make(map[string]*Export)
                                    }
                                    export := make([]*Export, len($2.Export))
                                    for _, e := range $2.Export {
                                        export[e.Position] = e
                                    }
                                    for _, e := range export {
                                        $$.Export[e.Name] = e
                                        exportSyms(RPClex, e.Name, e)
                                    }
                                }
                                if len($2.Operations) > 0 {
                                    $$.Operations = append($$.Operations, $2.Operations...)
                                }
                            }
                        ;

import                  : /* empty */
                            {
                                $$ = []string{}
                            }
                        | import1
                            {
                                $$ = $1
                            }
                        ;

import1                 : IMPORT import_list ';'
                            {
                                $$ = $2
                                // XXX: load import files.
                                for _, f := range $2 {
                                    if _, err := NewFile(f, "").Load(); err != nil {
                                        RPClex.Error("unable to load file: " + err.Error())
                                        return 0
                                    }
                                }
                            }

interface_component     : export
                            {
                                $$ = InterfaceBody{Export: make(map[string]*Export)}
                                for _, e := range $1 {
                                    $$.Export[e.Name], e.Position = e, len($$.Export)
                                    
                                }
                            }
                        | op_declarator ';'
                            {
                                $$ = InterfaceBody{Operations: []*Operation{$1}}
                            }
                        ;

export                  : type_declarator ';'
                            {
                                $$ = $1.pToExport_()
                            }
                        | const_declarator ';'
                            {
                                $$ = []*Export{&Export{Name: $1.Name, Const: $1}}
                            }
                        | pragma_declarator
                            {
                                $$ = $1
                            }
                        | tagged_declarator ';'
                            {
                                $$ = []*Export{&Export{Name: $1.TypeName(), Type: $1}}
                            }
                        ;

import_list             : import_name
                            {
                                $$ = []string{$1}
                            }
                        | import_list ',' import_name
                            {
                                $$ = append($$, $3)
                            }
                        ;

import_name             : STRING_LITERAL
                            {
                                $$ = $1
                            }
                        ;

const_declarator        : CONST const_type_spec IDENT '=' const_exp
                            {
                                exp, err := $5.Coerce($2)
                                if err != nil {
                                    RPClex.Error(err.Error())
                                    return 0
                                }
                                $$ = &Const{Type: $2, Name: $3, Value: exp}
                            }
                        ;
// ???: this is quite weird trying to preprocess the data from pragmas.
pragma_declarator       : PRAGMA_DEFINE IDENT const_exp
                            {
                                switch v := $3.Value.(type) {
                                case *big.Int:
                                    $$ = []*Export{&Export{Name: $2, Const: &Const{Type: TypeInt64, Name: $2, Value: $3}}}
                                case string:
                                    if $3.CanEval() {
                                        $$ = []*Export{&Export{Name: $2, Const: &Const{Type: TypeString, Name: $2, Value: $3}}}
                                        break
                                    }
                                    ref := &Type{Kind: TypeRef, Name: v}
                                    decl := &Type{Kind: TypeAttribute, Attrs: &TypeAttr{Alias: $2}}
                                    typ, ok := lookupType(RPClex, ref.Name)
                                    if !ok {
                                        // XXX: defer type resolution.
                                        decl.Elem = pushRef(RPClex, ref)
                                    } else {
                                        decl.Elem = typ
                                    }

                                    $$ = []*Export{&Export{Name: $2, Type: decl}}

                                    // FIXME: dnsp.idl PRAGMA Types.
                                    // $$ = []*Export{&Export{Name: $2, Type: &Type{Kind: TypeRef, Name: v}}}
                                default:
                                    RPClex.Error("invalid #define statement")
                                    return 0
                                }
                            }
                        | PRAGMA_CPP_QUOTE '(' STRING_LITERAL ')'
                            {
                                // XXX: re-enter the STRING_LITERAL as PRAGMA_DEFINE.
                                pushLex(RPClex, $3)
                            }
                        ;

const_type_spec         : primitive_integer_type
                            {
                                $$ = $1
                            }
                        | IDENT
                            {
                                typ, ok := lookupType(RPClex, $1)
                                if !ok {
                                    RPClex.Error("cannot lookup type " + $1)
                                    return 0
                                }

                                switch typ = typ.Base(); typ.Kind {
                                case TypeChar,
                                    TypeUChar,
                                    TypeWChar,
                                    TypeBoolean,
                                    TypeInt8,
                                    TypeUint8,
                                    TypeInt16,
                                    TypeUint16,
                                    TypeInt32,
                                    TypeUint32,
                                    TypeInt32_64,
                                    TypeUint32_64,
                                    TypeInt64,
                                    TypeUint64,
                                    TypeVoid,
                                    TypeString,
                                    // FIXME: float, double not expected but used.
                                    TypeFloat32,
                                    TypeFloat64:
                                        $$ = typ.Kind
                                default:
                                    RPClex.Error("invalid const type " + $1)
                                    return 0
                                }
                            }
                        | CHAR
                            {
                                $$ = TypeChar
                            }
                        | UNSIGNED CHAR
                            {
                                $$ = TypeUChar
                            }
                        | BOOLEAN
                            {
                                $$ = TypeBoolean
                            }
                        | VOID '*'
                            {
                                $$ = TypeVoid
                            }
                        | CHAR '*'
                            {
                                $$ = TypeString
                            }
                        ;

const_exp               : integer_const_exp
                            {
                                $$ = $1
                            }
                        | STRING_LITERAL
                            {
                                $$ = NewValue($1)
                            }
                        | CHARACTER_LITERAL
                            {
                                $$ = NewValue($1)
                            }
                        | NULL
                            {
                                $$ = NewValue(nil)
                            }
                        | TRUE
                            {
                                $$ = NewValue(true)
                            }
                        | FALSE
                            {
                                $$ = NewValue(false)
                            }
                        ;

integer_const_exp       : conditional_exp
                            {
                                $$ = $1
                            }
                        ;

conditional_exp         : logical_or_exp
                            {
                                $$ = $1
                            }
                        | logical_or_exp '?' integer_const_exp ':' conditional_exp
                            {
                                exp, ok := $1.Ter($3, $5)
                                if !ok {
                                    RPClex.Error("cannot evaluate ternary expression")
                                    return 0
                                }
                                $$ = exp
                            }
                        ;

logical_or_exp          : logical_and_exp
                            {
                                $$ = $1
                            }
                        | logical_or_exp LOGICAL_OR logical_and_exp
                            {
                                exp, ok := $1.LogicalOr($3)
                                if !ok {
                                    RPClex.Error("cannot evaluate l-or expression")
                                    return 0
                                }
                                $$ = exp
                            }
                        ;

logical_and_exp         : inclusive_or_exp
                            {
                                $$ = $1
                            }
                        | logical_and_exp LOGICAL_AND inclusive_or_exp
                            {
                                exp, ok := $1.LogicalAnd($3)
                                if !ok {
                                    RPClex.Error("cannot evaluate l-and expression")
                                    return 0
                                }
                                $$ = exp
                            }
                        ;

inclusive_or_exp        : exclusive_or_exp
                            {
                                $$ = $1
                            }
                        | inclusive_or_exp OR exclusive_or_exp
                            {
                                exp, ok := $1.Or($3)
                                if !ok {
                                    RPClex.Error("cannot evaluate or expression")
                                    return 0
                                }
                                $$ = exp
                            }
                        ;

exclusive_or_exp        : and_exp
                            {
                                $$ = $1
                            }
                        | exclusive_or_exp XOR and_exp
                            {
                                exp, ok := $1.Xor($3)
                                if !ok {
                                    RPClex.Error("cannot evaluate xor expression")
                                    return 0
                                }
                                $$ = exp
                            }
                        ;

and_exp                 : equality_exp
                            {
                                $$ = $1
                            }
                        | and_exp AND equality_exp
                            {
                                exp, ok := $1.And($3)
                                if !ok {
                                    RPClex.Error("cannot evaluate and expression")
                                    return 0
                                }
                                $$ = exp
                            }
                        ;

equality_exp            : relational_exp
                            {
                                $$ = $1
                            }
                        | equality_exp EQ relational_exp
                            {
                                exp, ok := $1.Eq($3)
                                if !ok {
                                    RPClex.Error("cannot evaluate eq expression")
                                    return 0
                                }
                                $$ = exp
                            }
                        | equality_exp NE relational_exp
                            {
                                exp, ok := $1.Ne($3)
                                if !ok {
                                    RPClex.Error("cannot evaluate ne expression")
                                    return 0
                                }
                                $$ = exp
                            }
                        ;

relational_exp          : shift_exp
                            {
                                $$ = $1
                            }
                        | relational_exp LT shift_exp
                            {
                                exp, ok := $1.Lt($3)
                                if !ok {
                                    RPClex.Error("cannot evaluate lt expression")
                                    return 0
                                }
                                $$ = exp
                            }
                        | relational_exp GT shift_exp
                            {
                                exp, ok := $1.Gt($3)
                                if !ok {
                                    RPClex.Error("cannot evaluate gt expression")
                                    return 0
                                }
                                $$ = exp
                            }
                        | relational_exp LE shift_exp
                            {
                                exp, ok := $1.Le($3)
                                if !ok {
                                    RPClex.Error("cannot evaluate le expression")
                                    return 0
                                }
                                $$ = exp
                            }
                        | relational_exp GE shift_exp
                            {
                                exp, ok := $1.Ge($3)
                                if !ok {
                                    RPClex.Error("cannot evaluate ge expression")
                                    return 0
                                }
                                $$ = exp
                            }
                        ;

shift_exp               : additive_exp
                            {
                                $$ = $1
                            }
                        | shift_exp LSH additive_exp
                            {
                                exp, ok := $1.Lsh($3)
                                if !ok {
                                    RPClex.Error("cannot evaluate lsh expression")
                                    return 0
                                }
                                $$ = exp
                            }
                        | shift_exp RSH additive_exp
                            {
                                exp, ok := $1.Rsh($3)
                                if !ok {
                                    RPClex.Error("cannot evaluate rsh expression")
                                    return 0
                                }
                                $$ = exp
                            }
                        ;

additive_exp            : multiplicative_exp
                            {
                                $$ = $1
                            }
                        | additive_exp '-' multiplicative_exp
                            {
                                exp, ok := $1.Sub($3)
                                if !ok {
                                    RPClex.Error("cannot evaluate sub expression")
                                    return 0
                                }
                                $$ = exp
                            }
                        | additive_exp '+' multiplicative_exp
                            {
                                exp, ok := $1.Add($3)
                                if !ok {
                                    RPClex.Error("cannot evaluate add expression")
                                    return 0
                                }
                                $$ = exp
                            }
                        ;

multiplicative_exp      : unary_exp
                            {
                                $$ = $1
                            }
                        | multiplicative_exp '*' unary_exp
                            {
                                exp, ok := $1.Mul($3)
                                if !ok {
                                    RPClex.Error("cannot evaluate mul expression")
                                    return 0
                                }
                                $$ = exp
                            }
                        | multiplicative_exp '/' unary_exp
                            {
                                exp, ok := $1.Div($3)
                                if !ok {
                                    RPClex.Error("cannot evaluate div expression")
                                    return 0
                                }
                                $$ = exp
                            }
                        | multiplicative_exp '%' unary_exp
                            {
                                exp, ok := $1.Rem($3)
                                if !ok {
                                    RPClex.Error("cannot evaluate rem expression")
                                    return 0
                                }
                                $$ = exp
                            }
                        ;

unary_exp               : primary_exp
                            {
                                $$ = $1
                            }
                        | '+' primary_exp %prec UPLUS
                            {
                                exp, ok := $2.Positive()
                                if !ok {
                                    RPClex.Error("cannot evaluate u'+' expression")
                                    return 0
                                }
                                $$ = exp
                            }
                        | '-' primary_exp %prec UMINUS
                            {
                                exp, ok := $2.Negative()
                                if !ok {
                                    RPClex.Error("cannot evaluate u'-' expression")
                                    return 0
                                }
                                $$ = exp
                            }
                        | '~' primary_exp %prec UNEG
                            {
                                exp, ok := $2.Neg()
                                if !ok {
                                    RPClex.Error("cannot evaluate u'~' expression")
                                    return 0
                                }
                                $$ = exp
                            }
                        | '!' primary_exp %prec UNOT
                            {
                                exp, ok := $2.Not()
                                if !ok {
                                    RPClex.Error("cannot evaluate u'!' expression")
                                    return 0
                                }
                                $$ = exp
                            }
                        ;

primary_exp             : INT_LITERAL
                            {
                                $$ = NewValue($1)
                            }
                        | SIZEOF '(' type_spec ')'
                            {
                                sz := TypeSize(RPClex.(TypeStore), $3)
                                if sz == 0 {
                                    RPClex.Error(fmt.Sprintf("cannot determine size of type: %s (%s)", $3.Name, $3.Kind))
                                    return 0
                                }
                                $$ = NewValue(big.NewInt(int64(sz)))
                            }
                        | IDENT
                            {
                                if c, ok := lookupConst(RPClex, $1); !ok {
                                    $$ = NewIdent($1)
                                } else {
                                    $$ = c.Value
                                }
                            }
                        | '*' IDENT
                            {
                                exp, ok := NewIdent($2).Ptr()
                                if !ok {
                                    RPClex.Error("cannot evaluate ptr expression")
                                    return 0
                                }
                                $$ = exp
                            }
                        | '(' const_exp ')'
                            {
                                $$ = $2
                            }
                        ;

type_declarator         : attributes0 TYPEDEF attributes0 type_spec declarators
                            {
                                attrs := $1.Merge($3).Type()
                                if attrs.Usage.ContextHandle {
                                    ref := &Type{Kind: TypeRef, Name: "ndr_context_handle"}
                                    typ, ok := lookupType(RPClex, ref.Name)
                                    if !ok {
                                        // XXX: defer type resolution.
                                        RPClex.Error("ndr_context_handle not found")
                                        break
                                    }
                                    $4 = typ
                                }
                                    
                                $$ = &pTypedef{Type: $4, Attrs: attrs, Declarators: $5}
                            }
                        ;

type_spec               : simple_type_spec
                            {
                                $$ = $1
                            }
                        | CONST simple_type_spec
                            {
                                $$ = $2
                            }
                        | simple_type_spec CONST
                            {
                                $$ = $1
                            }
                        | constructed_type_spec
                            {
                                $$ = $1
                            }
                        ;

simple_type_spec        : base_type_spec
                            {
                                $$ = &Type{Kind: $1}
                            }
                        | predefined_type_spec
                            {
                                $$ = $1
                            }
                        | type_ident
                            {
                                $$ = $1
                            }
                        ;

type_ident              : IDENT
                            {
                                ref := &Type{Kind: TypeRef, Name: $1}
                                typ, ok := lookupType(RPClex, ref.Name)
                                if !ok {
                                    // XXX: defer type resolution.
                                    $$ = pushRef(RPClex, ref)
                                } else {
                                    $$ = typ
                                }
                            }
                        ;


declarators             : declarator
                            {
                                $$ = pDeclarators{$1}
                            }
                        | declarators ',' declarator
                            {
                                $$ = append($$, $3)
                            }
// XXX: added empty declarator jff.
                        | /* empty */
                            {
                                $$ = pDeclarators{&pDeclarator{Name: ""}}
                            }
                        ;

declarator              : pointer_opt0 direct_declarator
                            {
                                $$ = $2
                                for i := 0; i < int($1); i++ {
                                    $$ = &pDeclarator{Name: $$.Name, Type: &Type{Kind: TypePointer, Elem: $$.Type}}
                                }
                            }
                        ;

direct_declarator       : ident_or_keyword
                            {
                                $$ = &pDeclarator{Name: $1}
                            }
                        | '(' declarator ')'
                            {
                               $$ = $2
                            }
                        | array_declarator
                            {
                                $$ = $1
                            }
                        | function_declarator
                            {
                                $$ = $1
                            }
                        ;

ident_or_keyword        : IDENT
                            {
                                $$ = $1
                            }
                        | error
                            {
                                ident, ok := tokName(RPClex, RPCrcvr.char)
                                if !ok {
                                    return 0
                                }
                                $$ = ident
                            }
                        ;

tagged_declarator       : tagged_struct_declarator
                            {
                                $$ = $1
                            }
                        | tagged_union_declarator
                            {
                                $$ = $1
                            }
// XXX: __midl__
                        | tagged_enumeration_declarator
                            {
                                $$ = $1
                            }
// XXX: __midl_end__
                        ;

/*  The base types are the fundamental data types of the IDL.
    Any other data types in an interface definition are derived
    from these types. Syntax gives the syntax rules for the base
    types. Integer Types to The handle_t Type define the various
    types.
*/

base_type_spec          : floating_pt_type
                            {
                                $$ = $1
                            }
                        | integer_type
                            {
                                $$ = $1
                            }
                        | char_type
                            {
                                $$ = $1
                            }
                        | boolean_type
                            {
                                $$ = $1
                            }
                        | byte_type
                            {
                                $$ = $1
                            }
                        | void_type
                            {
                                $$ = $1
                            }
                        | handle_type
                            {
                                $$ = $1
                            }
// XXX: __midl__
                        | wchar_type
                            {
                                $$ = $1
                            }
// XXX: __midl_end__
                        ;

floating_pt_type        : FLOAT
                            {
                                $$ = TypeFloat32
                            }
                        | DOUBLE
                            {
                                $$ = TypeFloat64
                            }
                        ;

integer_type            : primitive_integer_type
                            {
                                $$ = $1
                            }
                        | HYPER unsigned int
                            {
                                switch $2 {
                                    case UNSIGNED:
                                        $$ = TypeUint64
                                    default:
                                        $$ = TypeInt64
                                }
                            }
                        | UNSIGNED HYPER int
                            {
                                $$ = TypeUint64
                            }
                        ;

primitive_integer_type  : signed_integer
                            {
                                $$ = $1
                            }
                        | unsigned_integer
                            {
                                $$ = $1
                            }
                        | ms_integer
                            {
                                $$ = $1
                            }
                        ;

signed_integer          : SIGNED integer_size int
                            {
                                switch $2 {
                                case LONG:
                                    $$ = TypeInt32
                                case SHORT:
                                    $$ = TypeInt16
                                case SMALL:
                                    $$ = TypeInt8
                                }
                            }
                        | integer_size int
                            {
                                switch $1 {
                                case LONG:
                                    $$ = TypeInt32
                                case SHORT:
                                    $$ = TypeInt16
                                case SMALL:
                                    $$ = TypeInt8
                                }
                            }
                        ;

unsigned_integer        : integer_size UNSIGNED int
                            {
                                switch $1 {
                                case LONG:
                                    $$ = TypeUint32
                                case SHORT:
                                    $$ = TypeUint16
                                case SMALL:
                                    $$ = TypeUint8
                                }
                            }
                        | UNSIGNED integer_size int
                            {
                                switch $2 {
                                case LONG:
                                    $$ = TypeUint32
                                case SHORT:
                                    $$ = TypeUint16
                                case SMALL:
                                    $$ = TypeUint8
                                }
                            }
                        ;

// XXX: __midl__
ms_integer              : sign INT
                            {
                                switch $1 {
                                case UNSIGNED:
                                    $$ = TypeUint32
                                default:
                                    $$ = TypeInt32
                                }
                            }
                        | sign INT8
                            {
                                switch $1 {
                                case UNSIGNED:
                                    $$ = TypeUint8
                                default:
                                    $$ = TypeInt8
                                }
                            }
                        | sign INT16
                            {
                                switch $1 {
                                case UNSIGNED:
                                    $$ = TypeUint16
                                default:
                                    $$ = TypeInt16
                                }
                            }
                        | sign INT32
                            {
                                switch $1 {
                                case UNSIGNED:
                                    $$ = TypeUint32
                                default:
                                    $$ = TypeInt32
                                }
                            }
                        | sign INT3264
                            {
                                switch $1 {
                                case UNSIGNED:
                                    $$ = TypeUint32_64
                                default:
                                    $$ = TypeInt32_64
                                }
                            }
                        | sign INT64
                            {
                                switch $1 {
                                case UNSIGNED:
                                    $$ = TypeUint64
                                default:
                                    $$ = TypeInt64
                                }
                            }
                        ;
// XXX: __midl_end__

integer_size            : LONG
                            {
                                $$ = LONG
                            }
                        | SHORT
                            {
                                $$ = SHORT
                            }
                        | SMALL
                            {
                                $$ = SMALL
                            }
                        ;

char_type               : sign CHAR
                            {
                                switch $1 {
                                    case UNSIGNED:
                                        $$ = TypeUChar
                                    default:
                                        $$ = TypeChar
                                }
                            }
                        ;

// XXX: __midl__
wchar_type              : WCHAR_T
                            {
                                $$ = TypeWChar
                            }
                        ;
// XXX: __midl_end__

boolean_type            : BOOLEAN
                            {
                                $$ = TypeBoolean
                            }
                        ;

byte_type               : BYTE
                            {
                                $$ = TypeUint8
                            }
                        ;

void_type               : VOID
                            {
                                $$ = TypeVoid
                            }
                        ;

handle_type             : HANDLE_T
                            {
                                $$ = TypeHandle
                            }
                        ;

unsigned                : /* empty */
                            {
                                $$ = 0
                            }
                        | UNSIGNED
                            {
                                $$ = UNSIGNED
                            }
                        ;

sign                    : /* empty */
                            {
                                $$ = 0
                            }
                        | UNSIGNED
                            {
                                $$ = UNSIGNED
                            }
                        | SIGNED
                            {
                                $$ = 0
                            }
                        ;

void                    : /* empty */
                        | VOID
                        ;

int                     : /* empty */
                        | INT
                        ;


constructed_type_spec   : struct_type
                            {
                                $$ = $1
                            }
                        | union_type
                            {
                                $$ = $1
                            }
                        | enumeration_type
                            {
                                $$ = $1
                            }
                        | tagged_declarator
                            {
                                $$ = $1
                            }
                        | pipe_type
                            {
                                $$ = $1
                            }
                        ;

tagged_struct_declarator : STRUCT tag
                            {
                                ref := &Type{Kind: TypeRef, Name: "_struct_" + $2}
                                typ, ok := lookupType(RPClex, ref.Name)
                                if !ok {
                                    // XXX: defer type resolution.
                                    $$ = pushRef(RPClex, ref)
                                } else {
                                    $$ = typ
                                }
                            }
                        | tagged_struct
                            {
                                $$ = $1
                            }
                        ;

struct_type             : STRUCT '{' member_list '}'
                            {
                                for i := range $3 {
                                    $3[i].Position = i+1
                                }
                                // XXX: set proper switch type if not specified.
                                SetSwitchType($3)
                                $$ = &Type{Kind: TypeStruct, Struct: &Struct{Fields: $3}}
                            }
                        ;

tagged_struct           : STRUCT tag '{' member_list '}'
                            {
                                for i := range $4 {
                                    $4[i].Position = i+1
                                }
                                // XXX: set proper switch type if not specified.
                                SetSwitchType($4)
                                $$ = &Type{Kind: TypeStruct, Tag: $2, Struct: &Struct{Fields: $4}}
                            }
                        ;

tag                     : IDENT
                            {
                                $$ = $1
                            }
                        ;

member_list             : member
                            {
                                $$ = $1
                            }
                        | member_list member
                            {
                                $$ = append($$, $2...)
                            }
                        ;

member                  : field_declarator ';'
                            {
                                $$ = $1
                            }
                        ;

field_declarator        : attributes0 type_spec declarators field_attr_var
                            {
                                $$ = make([]*Field, 0, len($3))
                                for _, decl := range $3 {
                                    decl.Type = decl.Type.Append($2)
                                    $$ = append($$, &Field{Attrs: $1.Field(), Name: decl.Name, Type: decl.Type, DefaultValue: $4})
                                }

                            }
                        ;

field_attr_var          : /* empty */
                            {
                                $$ = Expr{}
                            }
                        | '=' attr_var
                            {
                                $$ = $2
                            }
                        ;

tagged_union_declarator : UNION tag
                            {
                                $$ = &Type{Kind: TypeUnion, Tag: $2}
                            }
                        | tagged_union
                            {
                                $$ = $1
                            }
                        ;

union_type              : UNION union_switch union_name '{' union_body '}'
                            {
                                pos := 1
                                for i := range $5 {
                                    for j := range $5[i].Arms {
                                        $5[i].Arms[j].Position = pos
                                        pos++
                                    }
                                }
                                $$ = &Type{Kind: TypeUnion, Name: $3, Union: &Union{Switch: $2, Body: $5}}
                            }
                        | UNION '{' union_body_n_e '}'
                            {
                                pos := 1
                                for i := range $3 {
                                    for j := range $3[i].Arms {
                                        $3[i].Arms[j].Position = pos
                                        pos++
                                    }
                                }
                                $$ = &Type{Kind: TypeUnion, Union: &Union{Body: $3}}
                            }
                        | UNION '{' union_body_c '}'
                            {
                                pos := 1
                                for i := range $3 {
                                    for j := range $3[i].Arms {
                                        $3[i].Arms[j].Position = pos
                                        pos++
                                    }
                                }
                                $$ = &Type{Kind: TypeCUnion, Union: &Union{Body: $3}}
                            }
                        ;

union_switch            : SWITCH '(' switch_type_spec IDENT ')'
                            {
                                $$ = &UnionSwitch{Type: $3, Name: $4}
                            }
                        ;

switch_type_spec        : primitive_integer_type
                            {
                                $$ = &Type{Kind: $1}
                            }
                        | char_type
                            {
                                $$ = &Type{Kind: $1}
                            }
                        | boolean_type
                            {
                                $$ = &Type{Kind: $1}
                            }
                        | type_ident
                            {
                                $$ = $1
                            }
// FIXME: enum tag only for few cases.
                        | ENUM tag
                            {
                                ref := &Type{Kind: TypeRef, Name: "_enum_" + $2}
                                typ, ok := lookupType(RPClex, ref.Name)
                                if !ok {
                                    // XXX: defer type resolution.
                                    $$ = pushRef(RPClex, ref)
                                } else {
                                    $$ = typ
                                }
                            }
                        ;

tagged_union            : UNION tag union_switch union_name '{' union_body '}'
                            {
                                pos := 1
                                for i := range $6 {
                                    for j := range $6[i].Arms {
                                        $6[i].Arms[j].Position = pos
                                        pos++
                                    }
                                }
                                $$ = &Type{Kind: TypeUnion, Tag: $2, Union: &Union{Switch: $3, Body: $6}}
                            }
                        | UNION tag '{' union_body_n_e '}'
                            {
                                pos := 1
                                for i := range $4 {
                                    for j := range $4[i].Arms {
                                        $4[i].Arms[j].Position = pos
                                        pos++
                                    }
                                }
                                $$ = &Type{Kind: TypeUnion, Tag: $2, Union: &Union{Body: $4}}
                            }
                        | UNION tag '{' union_body_c '}'
                            {
                                pos := 1
                                for i := range $4 {
                                    for j := range $4[i].Arms {
                                        $4[i].Arms[j].Position = pos
                                        pos++
                                    }
                                }
                                $$ = &Type{Kind: TypeCUnion, Tag: $2, Union: &Union{Body: $4}}
                            }
                        ;

union_name              : /* empty */
                            {
                                // In encapsulated unions, if the <union_name> is
                                // omitted, the union is assigned the name tagged_union
                                // in the generated header source.
                                $$ = "tagged_union"
                            }
                        | IDENT
                            {
                                $$ = $1
                            }
                        ;

union_body              : union_case
                            {
                                $$ = []*UnionCase{$1}
                            }
                        | union_body union_case
                            {
                                $$ = append($$, $2)
                            }
                        ;

union_body_c            : union_arm
                            {
                                $$ = []*UnionCase{&UnionCase{Arms: $1}}
                            }
                        | union_body_c union_arm
                            {
                                $$ = append($$, &UnionCase{Arms: $2})
                            }
                        ;

union_body_n_e          : union_case_n_e
                            {
                                $$ = []*UnionCase{$1}
                            }
                        | union_body_n_e union_case_n_e
                            {
                                $$ = append($$, $2)
                            }
                        ;

union_case              : union_case_label1 union_arm
                            {
                                labels := make([]interface{}, 0, len($1))
                                for _, label := range $1 {
                                    labels = append(labels, label)
                                }
                                $$ = &UnionCase{Labels: labels, Arms: $2}
                            }
                        | default_case
                            {
                                $$ = &UnionCase{Arms: $1, IsDefault: true}
                            }
                        ;

union_case_label1       : union_case_label
                            {
                                $$ = $1
                            }
                        | union_case_label1 union_case_label
                            {
                                $$ = append($$, $2...)
                            }
                        ;

union_case_n_e          : union_case_label_n_e union_arm
                            {
                                labels := make([]interface{}, 0, len($1))
                                for _, label := range $1 {
                                    labels = append(labels, label)
                                }
                                $$ = &UnionCase{Labels: labels, Arms: $2}
                            }
                        | default_case_n_e
                            {
                                $$ = &UnionCase{Labels: nil, Arms: $1, IsDefault: true}
                            }
                        ;

union_case_label        : CASE const_exp ':'
                            {
                                $$ = Exprs{$2}
                            }
                        ;

union_case_label_n_e    : '[' CASE '(' const_exp1 ')' ']'
                            {
                                $$ = $4
                            }
                        ;

const_exp1              : const_exp
                            {
                                $$ = Exprs{$1}
                            }
                        | const_exp1 ',' const_exp
                            {
                                $$ = append($$, $3)
                            }
                        ;


default_case            : DEFAULT ':' union_arm
                            {
                                $$ = $3
                            }
                        ;

default_case_n_e        : '[' DEFAULT ']' union_arm
                            {
                                $$ = $4
                            }
                        ;

union_arm               : ';'
                            {
                                $$ = nil
                            }
                        | field_declarator ';'
                            {
                                for i := range $1 {
                                    $1[i].Position = i+1
                                }
                                $$ = $1
                            }
                        ;

union_type_switch_attr  : SWITCH_TYPE '(' switch_type_spec ')'
                            {
                                $$ = $3
                            }
                        ;

union_instance_switch_attr : SWITCH_IS '(' attr_var ')'
                           {
                                $$ = pAttrType{SWITCH_IS, pAttr{SwitchIs: $3}}
                           }
                        ;

enumeration_type        : ENUM '{' enumeration1 comma '}'
                            {
                                $$ = &Type{Kind: TypeEnum, Enum: &Enum{Elems: make([]*Element, 0, len($3))}}
                                for idx, i := 0, 0; i < len($3); i, idx = i+1, idx+1 {
                                    tag, exp := $3[i].Tag, $3[i].ID
                                    if !exp.Empty() {
                                        val, ok := exp.Eval(RPClex.(ExprStore))
                                        if !ok {
                                            RPClex.Error("cannot evaluate tag-id expression")
                                            return 0
                                        }
                                        bi, ok := val.BigInt()
                                        if !ok {
                                            RPClex.Error("enum: not an integer type")
                                            return 0
                                        }
                                        idx, exp = int(bi.Uint64()), val
                                    } else {
                                        exp = NewValue(big.NewInt(int64(idx)))
                                    }
                                    if uint32(uint16(idx)) != uint32(idx) {
                                        $$.Enum.Is32 = true
                                    }
                                    $$.Enum.Elems = append($$.Enum.Elems, &Element{tag, idx})
                                    storeConst(RPClex, tag, exp)
                                }
                            }
                        ;

// XXX: __midl__
tagged_enumeration_declarator : ENUM tag
                            {
                                $$ = &Type{Kind: TypeEnum, Tag: $2}
                            }
                        | tagged_enumeration
                            {
                                $$ = $1
                            }
                        ;

tagged_enumeration      : ENUM tag '{' enumeration1 comma '}'
                            {
                                $$ = &Type{Kind: TypeEnum, Tag: $2, Enum: &Enum{Elems: make([]*Element, 0, len($4))}}
                                for idx, i := 0, 0; i < len($4); i, idx = i+1, idx+1 {
                                    tag, exp := $4[i].Tag, $4[i].ID
                                    if !exp.Empty() {
                                        val, ok := exp.Eval(RPClex.(ExprStore))
                                        if !ok {
                                            RPClex.Error("cannot evaluate tag-id expression")
                                            return 0
                                        }
                                        bi, ok := val.BigInt()
                                        if !ok {
                                            RPClex.Error("enum: not an integer type")
                                            return 0
                                        }
                                        idx, exp = int(bi.Uint64()), val
                                    } else {
                                        exp = NewValue(big.NewInt(int64(idx)))
                                    }
                                    if uint32(uint16(idx)) != uint32(idx) {
                                        $$.Enum.Is32 = true
                                    }
                                    $$.Enum.Elems = append($$.Enum.Elems, &Element{tag, idx})
                                    storeConst(RPClex, tag, exp)
                                }
                            }
                        ;
// XXX: __midl_end__

enumeration1            : identifier_tag
                            {
                                $$ = pTagIDs{$1}
                            }
                        | enumeration1 ',' identifier_tag
                            {
                                $$ = append($$, $3)
                            }
                        ;

comma                   : /* empty */
                        | ','
                        ;

identifier_tag          : IDENT
                            {
                                $$ = pTagID{Tag: $1}
                            }
// XXX: __midl__
                        | IDENT '=' integer_const_exp
                            {
                                $$ = pTagID{Tag: $1, ID: $3}
                            }
// XXX: __midl_end__
                        ;

pipe_type               : PIPE type_spec
                            {
                                $$ = &Type{Kind: TypePipe, Elem: $2}
                            }
                        ;

array_declarator        : direct_declarator '[' array_bounds_declarator ']'
                            {

                                // XXX: array associativity is different from pointer.
                                // so we need to find last array in the chain and insert
                                // new array right after it (as it's element).

                                array := &Type{Kind: TypeArray, Array: &Array{Bound: $3}, Elem: $1.Type}
                                if $1.Type == nil || $1.Type.Kind != TypeArray {
                                    $$ = &pDeclarator{Name: $1.Name, Type: array}
                                    break
                                }

                                last := $1.Type
                                // shift to the last array in the element chain.
                                for last.Elem != nil && last.Elem.Kind == TypeArray {
                                    last = last.Elem
                                }

                                // insert array after last array and acuqire
                                // the last array's element.
                                array.Elem, last.Elem = last.Elem, array
                                $$ = &pDeclarator{Name: $1.Name, Type: $1.Type}
                            }
                        ;

array_bounds_declarator : /* empty */
                            {
                                $$ = ArrayBound{Lower: 0, Upper: -1}
                            }
                        | array_bound
                            {
                                if $1.Int64() == -1 {
                                    $$ = ArrayBound{Lower: -1, Upper: 0}
                                } else {
                                    $$ = ArrayBound{Lower: 0, Upper: $1.Int64()-1}
                                }

                            }
                        | array_bound_pair
                            {
                                $$ = $1
                            }
                        ;

array_bound_pair        : array_bound RNG array_bound
                            {
                                $$ = ArrayBound{Lower: $1.Int64(), Upper: $3.Int64()}
                            }
                        ;

array_bound             : '*'
                            {
                                $$ = big.NewInt(-1)
                            }
                        | integer_const_exp
                            {
                                if !$1.CanEval() {
                                    RPClex.Error("cannot evaluate integer bound")
                                    return 0
                                }
                                val, ok := $1.BigInt()
                                if !ok {
                                    RPClex.Error("cannot use non-integer as a bound declarator")
                                    return 0
                                }
                                $$ = val
                            }
                        ;

usage_attribute         : USAGE_STRING
                            {
                                $$ = USAGE_STRING
                            }
                        | USAGE_CONTEXT_HANDLE
                            {
                                $$ = USAGE_CONTEXT_HANDLE
                            }
                        ;

format_attribute        : FORMAT_UTF8
                            {
                                $$ = FORMAT_UTF8
                            }
                        | FORMAT_NULL_TERMINATED
                            {
                                $$ = FORMAT_NULL_TERMINATED
                            }
                        | FORMAT_MULTI_SIZE
                            {
                                $$ = FORMAT_MULTI_SIZE
                            }
                        | FORMAT_RUNE
                            {
                                $$ = FORMAT_RUNE
                            }
                        | FORMAT_HEX
                            {
                                $$ = FORMAT_HEX
                            }
                        ;

xmit_type               : simple_type_spec
                            {
                                $$ = $1
                            }
                        ;

range_attribute         : RANGE range_declarator
                            {
                                $$ = pAttrType{RANGE, pAttr{Range: $2}}
                            }
                        ;

range_declarator        : '(' integer_const_exp ',' integer_const_exp ')'
                            {
                                if !CanEval($2, $4) {
                                    RPClex.Error("cannot evaluate range declaration")
                                    return 0
                                }
                                min, ok := $2.BigInt()
                                if !ok {
                                    RPClex.Error("invalid min value for range declarator")
                                    return 0
                                }
                                max, ok := $4.BigInt()
                                if !ok {
                                    RPClex.Error("invalid max value for range declarator")
                                    return 0
                                }
                                $$ = &Range{Min: min.Int64(), Max: max.Int64()}
                            }
                        ;
// XXX: __midl_end__

attr_var_list           : attr_var
                            {
                                $$ = []Expr{$1}
                            }
                        | attr_var_list ',' attr_var
                            {
                                // see https://learn.microsoft.com/en-us/windows/win32/rpc/multiple-levels-of-pointers
                                // for (,Size) constructions.
                                $$ = append($$, $3)
                            }
                        ;

attr_var                : /* empty */
                            {
                                $$ = Expr{}
                            }
                        | integer_const_exp
                            {
                                $$ = $1
                            }
                        ;

pointer_opt0            : /* empty */
                            {
                                $$ = 0
                            }
                        | pointer_opt
                            {
                                $$ = $1
                            }
                        ;

pointer_opt             : '*'
                            {
                                $$++
                            }
                        | pointer_opt '*'
                            {
                                $$++
                            }
                        ;

/*
pointer                 : '*' ...
*/

ptr_attr                : POINTER_REF
                            {
                                $$ = POINTER_REF
                            }
                        | POINTER_UNIQUE
                            {
                                $$ = POINTER_UNIQUE
                            }
                        | POINTER_PTR
                            {
                                $$ = POINTER_PTR
                            }
                        ;

op_declarator           : attributes0 simple_type_spec IDENT '(' param_declarators1 ')'
                            {
                                for i := range $5 {
                                    if $5[i].Name == "" {
                                        $5[i].Name = fmt.Sprintf("Param%d", i)
                                    }
                                    if $5[i].Type.Is(TypePointer) && $5[i].Attrs.Pointer == PointerTypeNone {
                                        $5[i].Attrs.Pointer = PointerTypeRefWeak
                                    }
                                }
                                $$ = &Operation{Attrs: $1.Operation(), Type: $2, Name: $3, Params: $5}
                            }
                        | attributes0 simple_type_spec IDENT '(' void ')'
                            {
                                $$ = &Operation{Attrs: $1.Operation(), Type: $2, Name: $3}
                            }
                        ;

param_declarators1      : param_declarators1 ',' param_declarator
                            {
                                $$ = append($$, $3)
                            }
                        | param_declarator
                            {
                                $$ = []*Param{$1}
                            }
                        ;

param_declarator        : attributes1 type_spec declarator
                            {
                                $3.Type = $3.Type.Append($2)
                                $$ = &Param{Attrs: $1.Param(), Name: $3.Name, Type: $3.Type}
                                if $$.Attrs.Usage.ContextHandle && !$3.Type.Base().Is(TypeStruct) {
                                    ref := &Type{Kind: TypeRef, Name: "ndr_context_handle"}
                                    if typ, ok := lookupType(RPClex, ref.Name); !ok {
                                        // XXX: defer type resolution.
                                        $$.Type = $$.Type.AppendAfterPointer(pushRef(RPClex, ref))
                                    } else {
                                        $$.Type = $$.Type.AppendAfterPointer(typ)
                                    }
                                } 
                            }
                        | type_spec declarator
                            {
                                $2.Type = $2.Type.Append($1)
                                attrs := pAttr{Direction: Direction{In: true}}
                                $$ = &Param{Attrs: attrs.Param(), Name: $2.Name, Type: $2.Type}
                            }
                        ;

directional_attribute   : IN
                            {
                                $$ = pAttrType{IN, pAttr{}}
                            }
                        | OUT
                            {
                                $$ = pAttrType{OUT, pAttr{}}
                            }
                        ;

function_declarator     : direct_declarator '(' param_declarators1 ')'
                            {
                                for i := range $3 {
                                    if $3[i].Name == "" {
                                        $3[i].Name = fmt.Sprintf("Param%d", i)
                                    }
                                }
                                $$ = &pDeclarator{Name: $1.Name, Type: &Type{Kind: TypeFunc, Func: &Func{Params: $3}, Elem: $1.Type}}
                            }
                        | direct_declarator '(' void ')'
                            {
                                $$ = &pDeclarator{Name: $1.Name, Type: &Type{Kind: TypeFunc, Func: &Func{}, Elem: $1.Type}}
                            }
                        ;

/*  Predefined types are data types derived from the base
    types that are intrinsic to the IDL language. The syntax
    for predefined types is as follows:
*/

predefined_type_spec    : ERROR_STATUS_T
                            {
                                $$ = &Type{Kind: TypeError}
                            }
                        | international_character_type
                            {
                                $$ = &Type{Kind: TypeCharset}

                                switch $1 {
                                case ISO_LATIN_1:
                                    $$.Charset = CharsetISO_Latin_1
                                case ISO_MULTILINGUAL:
                                    $$.Charset = CharsetISO_Multilingual
                                case ISO_UCS:
                                    $$.Charset = CharsetISO_UCS
                                }
                            }
                        ;

international_character_type : ISO_LATIN_1
                            {
                                $$ = ISO_LATIN_1
                            }
                        | ISO_MULTILINGUAL
                            {
                                $$ = ISO_MULTILINGUAL
                            }
                        | ISO_UCS
                            {
                                $$ = ISO_UCS
                            }
                        ;
%%
