package main
type T interface {
	Read(b Buffer) bool
	Write(b Buffer) bool
	Close()
}

/**-----
Go file
  PackageDeclaration(main)
    PsiElement(KEYWORD_PACKAGE)('package')
    PsiWhiteSpace(' ')
    PsiElement(IDENTIFIER)('main')
  PsiElement(WS_NEW_LINES)('\n')
  TypeDeclarationsImpl
    PsiElement(KEYWORD_TYPE)('type')
    PsiWhiteSpace(' ')
    TypeSpecImpl
      TypeNameDeclaration(T)
        PsiElement(IDENTIFIER)('T')
      PsiWhiteSpace(' ')
      TypeInterfaceImpl
        PsiElement(KEYWORD_INTERFACE)('interface')
        PsiWhiteSpace(' ')
        PsiElement({)('{')
        PsiElement(WS_NEW_LINES)('\n')
        PsiWhiteSpace('\t')
        MethodDeclaration(Read)
          PsiElement(IDENTIFIER)('Read')
          PsiElement(()('(')
          FunctionParameterListImpl
            FunctionParameterImpl
              LiteralIdentifierImpl
                PsiElement(IDENTIFIER)('b')
              PsiWhiteSpace(' ')
              TypeNameImpl
                LiteralIdentifierImpl
                  PsiElement(IDENTIFIER)('Buffer')
          PsiElement())(')')
          PsiWhiteSpace(' ')
          FunctionResult
            FunctionParameterListImpl
              FunctionParameterImpl
                TypeNameImpl
                  LiteralIdentifierImpl
                    PsiElement(IDENTIFIER)('bool')
        PsiElement(WS_NEW_LINES)('\n')
        PsiWhiteSpace('\t')
        MethodDeclaration(Write)
          PsiElement(IDENTIFIER)('Write')
          PsiElement(()('(')
          FunctionParameterListImpl
            FunctionParameterImpl
              LiteralIdentifierImpl
                PsiElement(IDENTIFIER)('b')
              PsiWhiteSpace(' ')
              TypeNameImpl
                LiteralIdentifierImpl
                  PsiElement(IDENTIFIER)('Buffer')
          PsiElement())(')')
          PsiWhiteSpace(' ')
          FunctionResult
            FunctionParameterListImpl
              FunctionParameterImpl
                TypeNameImpl
                  LiteralIdentifierImpl
                    PsiElement(IDENTIFIER)('bool')
        PsiElement(WS_NEW_LINES)('\n')
        PsiWhiteSpace('\t')
        MethodDeclaration(Close)
          PsiElement(IDENTIFIER)('Close')
          PsiElement(()('(')
          PsiElement())(')')
        PsiElement(WS_NEW_LINES)('\n')
        PsiElement(})('}')
  PsiElement(WS_NEW_LINES)('\n')